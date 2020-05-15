package core

import (
	"errors"
	"fmt"
	"math"
	"math/big"

	"github.com/BOPR/common"
	"github.com/BOPR/contracts/rollup"
	"github.com/ethereum/go-ethereum/accounts/abi"
	gormbulk "github.com/t-tiger/gorm-bulk-insert"
)

// UserAccount is the user data stored on the node per user
type UserAccount struct {
	// DBModel

	// ID is the path of the user account in the PDA Tree
	// Cannot be changed once created
	AccountID uint64 `gorm:"not null;index:AccountID"`

	// Token type of the user account
	// Cannot be changed once creation
	TokenType uint64 `gorm:"not null;default:0"`

	// Balance of the user account
	Balance uint64 `gorm:"not null;"`

	// Nonce of the account
	Nonce uint64 `gorm:"not null;"`

	// Public key for the user
	PublicKey string `gorm:"size:1000"`

	// Path from root to leaf
	// NOTE: not a part of the leaf
	// Path is a string to that we can run LIKE queries
	Path string `gorm:"not null;index:Path"`

	// Pending = 0 means has deposit but not merged to balance tree
	// Active = 1
	// InActive = 2 => non leaf node
	// NonInitialised = 100
	Status uint64 `gorm:"not null;index:Status"`

	// Type of nodes
	// 1 => terminal
	// 0 => root
	// 2 => non terminal
	Type uint64 `gorm:"not null;index:Type"`

	// keccak hash of the node
	Hash string `gorm:"not null;index:Hash"`

	Level uint64 `gorm:"not null;index:Level"`
}

func NewUserAccount(id, balance, tokenType, nonce, status uint64, pubkey, path string) *UserAccount {
	newAcccount := &UserAccount{
		AccountID: id,
		PublicKey: pubkey,
		Balance:   balance,
		TokenType: tokenType,
		Nonce:     nonce,
		Path:      path,
		Status:    status,
		Type:      1,
	}
	newAcccount.UpdatePath(newAcccount.Path)
	newAcccount.CreateAccountHash()
	return newAcccount
}

func NewPendingUserAccount(id, balance, tokenType uint64, _pubkey string) *UserAccount {
	newAcccount := &UserAccount{
		AccountID: id,
		TokenType: tokenType,
		Balance:   balance,
		Nonce:     0,
		Path:      "0",
		Status:    0,
		PublicKey: _pubkey,
		Type:      1,
	}
	newAcccount.UpdatePath(newAcccount.Path)
	newAcccount.CreateAccountHash()
	return newAcccount
}

func NewAccountNode(path, hash string) *UserAccount {
	newAcccount := &UserAccount{
		AccountID: 0,
		PublicKey: "",
		Balance:   0,
		TokenType: 0,
		Nonce:     0,
		Path:      path,
		Status:    1,
		Type:      2,
	}
	newAcccount.UpdatePath(newAcccount.Path)
	newAcccount.Hash = hash
	return newAcccount
}

func (acc *UserAccount) UpdatePath(path string) {
	acc.Path = path
	acc.Level = uint64(len(path))
}

func EmptyAccount() UserAccount {
	return *NewUserAccount(0, 0, 0, 0, 1, "", "")
}

func (acc *UserAccount) String() string {
	return fmt.Sprintf("ID: %d Bal: %d Path: %v Nonce: %v TokenType:%v NodeType: %d %v", acc.AccountID, acc.Balance, acc.Path, acc.Nonce, acc.TokenType, acc.Type, acc.Hash)
}

func (acc *UserAccount) ToABIAccount() rollup.DataTypesUserAccount {
	return rollup.DataTypesUserAccount{
		ID:        UintToBigInt(acc.AccountID),
		Balance:   UintToBigInt(acc.Balance),
		TokenType: UintToBigInt(acc.TokenType),
		Nonce:     UintToBigInt(acc.Nonce),
	}
}

func (acc *UserAccount) HashToByteArray() ByteArray {
	ba, err := HexToByteArray(acc.Hash)
	if err != nil {
		panic(err)
	}
	return ba
}

func (acc *UserAccount) IsCoordinator() bool {
	if acc.Path != "" {
		return false
	}

	if acc.Status != 1 {
		return false
	}

	if acc.Type != 0 {
		return false
	}

	return true
}

func (acc *UserAccount) AccountInclusionProof(path int64) rollup.DataTypesAccountInclusionProof {
	return rollup.DataTypesAccountInclusionProof{
		PathToAccount: big.NewInt(path),
		Account:       acc.ToABIAccount(),
	}
}

func (acc *UserAccount) ABIEncode() ([]byte, error) {
	uint256Ty, err := abi.NewType("uint256", "uint256", nil)
	if err != nil {
		return []byte(""), err
	}

	arguments := abi.Arguments{
		{
			Type: uint256Ty,
		},
		{
			Type: uint256Ty,
		},
		{
			Type: uint256Ty,
		},
		{
			Type: uint256Ty,
		},
	}
	bytes, err := arguments.Pack(
		big.NewInt(int64(acc.AccountID)),
		big.NewInt(int64(acc.Balance)),
		big.NewInt(int64(acc.Nonce)),
		big.NewInt(int64(acc.TokenType)),
	)
	if err != nil {
		return []byte(""), err
	}

	return bytes, nil
}

func (acc *UserAccount) CreateAccountHash() {
	data, err := acc.ABIEncode()
	if err != nil {
		return
	}
	accountHash := common.Keccak256(data)
	acc.Hash = accountHash.String()
}

// InitBalancesTree initialises the balances tree
func (db *DB) InitBalancesTree(depth uint64, genesisAccounts []UserAccount) error {
	// calculate total number of leaves
	totalLeaves := math.Exp2(float64(depth))
	if int(totalLeaves) != len(genesisAccounts) {
		return errors.New("Depth and number of leaves do not match")
	}
	db.Logger.Debug("Attempting to init balance tree", "totalAccounts", totalLeaves)

	var err error

	// insert coodinator leaf
	err = db.InsertCoordinatorAccount(&genesisAccounts[0], depth)
	if err != nil {
		db.Logger.Error("Unable to insert coodinator account", "err", err)
		return err
	}

	var insertRecords []interface{}
	prevNodePath := genesisAccounts[0].Path

	for i := 1; i < len(genesisAccounts); i++ {
		pathToAdjacentNode, err := GetAdjacentNodePath(prevNodePath)
		if err != nil {
			return err
		}
		db.Logger.Debug("Obtained adjacent node path", "adjacentNodePath", pathToAdjacentNode, "currentNodePath", prevNodePath)
		genesisAccounts[i].UpdatePath(pathToAdjacentNode)
		// TODO set type and other node level data as well
		insertRecords = append(insertRecords, genesisAccounts[i])
		prevNodePath = genesisAccounts[i].Path
	}

	db.Logger.Info("Inserting all accounts to DB", "count", len(insertRecords))
	err = gormbulk.BulkInsert(db.Instance, insertRecords, len(insertRecords))
	if err != nil {
		db.Logger.Error("Unable to insert accounts to DB", "err", err)
		return errors.New("Unable to insert accounts")
	}

	// merkelise
	// 1. Pick all leaves at level depth
	// 2. Iterate 2 of them and create parents and store
	// 3. Persist all parents to database
	// 4. Start with next round
	for i := depth; i > 0; i-- {
		// get all leaves at depth N
		accs, err := db.GetAccountsAtDepth(i)
		if err != nil {
			return err
		}
		db.Logger.Info("Fetched nodes from DB", "depth", i, "count", len(accs))

		var nextLevelAccounts []interface{}

		// iterate over 2 at a time and create next level
		for i := 0; i < len(accs); i += 2 {
			db.Logger.Debug("Creating parent node", "leftAccount", accs[i].String(), "rightAccount", accs[i+1].String())
			left, err := HexToByteArray(accs[i].Hash)
			if err != nil {
				return err
			}
			right, err := HexToByteArray(accs[i+1].Hash)
			if err != nil {
				return err
			}
			parentHash, err := GetParent(left, right)
			if err != nil {
				return err
			}
			parentPath := GetParentPath(accs[i].Path)
			db.Logger.Debug("adding parent with path", "obtainedParentPath", parentPath, "child", accs[i].Path)
			newAccNode := *NewAccountNode(parentPath, parentHash.String())
			nextLevelAccounts = append(nextLevelAccounts, newAccNode)
		}

		err = gormbulk.BulkInsert(db.Instance, nextLevelAccounts, len(nextLevelAccounts))
		if err != nil {
			db.Logger.Error("Unable to insert accounts to DB", "err", err)
			return errors.New("Unable to insert accounts")
		}
	}

	// mark the root node type correctly
	return nil
}

func (db *DB) GetAccountsAtDepth(depth uint64) ([]UserAccount, error) {
	var accs []UserAccount
	err := db.Instance.Where("level = ?", depth).Find(&accs).Error
	if err != nil {
		return accs, err
	}
	return accs, nil
}

func (db *DB) UpdateAccount(account UserAccount) error {
	siblings, err := db.GetSiblings(account.Path)
	if err != nil {
		return err
	}

	db.Logger.Debug("Updating account", "Hash", account.Hash, "Path", account.Path, "siblings", siblings, "countOfSiblings", len(siblings))
	return db.StoreLeaf(account, account.Path, siblings)
}

func (db *DB) StoreLeaf(account UserAccount, path string, siblings []UserAccount) error {
	var err error
	computedNode := account
	for i := 0; i < len(siblings); i++ {
		var parentHash ByteArray
		sibling := siblings[i]
		isComputedRightSibling := GetNthBitFromRight(
			path,
			i,
		)
		if isComputedRightSibling == 0 {
			parentHash, err = GetParent(computedNode.HashToByteArray(), sibling.HashToByteArray())
			if err != nil {
				return err
			}
			db.Logger.Info("Updating account", "ComputedNode", computedNode.String(), "Sibling", sibling.String(), "ParentHash", parentHash)
			// Store the node!
			err = db.StoreNode(parentHash, computedNode, sibling)
			if err != nil {
				return err
			}
		} else {
			parentHash, err = GetParent(sibling.HashToByteArray(), computedNode.HashToByteArray())
			if err != nil {
				return err
			}
			db.Logger.Info("Updating account", "Sibling", sibling.String(), "ComputedNode", computedNode.String(), "ParentHash", parentHash)

			// Store the node!
			err = db.StoreNode(parentHash, sibling, computedNode)
			if err != nil {
				return err
			}
		}

		parentAccount, err := db.GetAccountByPath(GetParentPath(computedNode.Path))
		if err != nil {
			return err
		}
		computedNode = parentAccount
	}

	// Store the new root
	err = db.UpdateRootNode(computedNode.HashToByteArray())
	if err != nil {
		return err
	}

	fmt.Println("Updating whole balance tree root to", computedNode.Hash)
	return nil
}

// StoreNode updates the nodes given the parent hash
func (db *DB) StoreNode(parentHash ByteArray, leftNode UserAccount, rightNode UserAccount) (err error) {
	db.Logger.Info("Storing account", "Account", leftNode.String(), "path", leftNode.Path)
	// update left account
	err = db.updateAccount(leftNode, leftNode.Path)
	if err != nil {
		return err
	}
	db.Logger.Info("Storing account", "Account", rightNode.String(), "path", rightNode.Path)
	// update right account
	err = db.updateAccount(rightNode, rightNode.Path)
	if err != nil {
		return err
	}
	// update the parent with the new hash
	return db.UpdateParentWithHash(GetParentPath(leftNode.Path), parentHash)
}

func (db *DB) UpdateParentWithHash(pathToParent string, newHash ByteArray) error {
	// Update the root hash
	var tempAccount UserAccount
	tempAccount.Path = pathToParent
	tempAccount.Hash = newHash.String()
	db.Logger.Debug("Updating parent account", "hash", tempAccount.Hash, "path", pathToParent)
	return db.updateAccount(tempAccount, pathToParent)
}

func (db *DB) UpdateRootNode(newRoot ByteArray) error {
	var tempAccount UserAccount
	tempAccount.Path = ""
	tempAccount.Hash = newRoot.String()
	return db.updateAccount(tempAccount, tempAccount.Path)
}

func (db *DB) AddNewPendingAccount(acc UserAccount) error {
	return db.Instance.Create(&acc).Error
}

func (db *DB) GetSiblings(path string) ([]UserAccount, error) {
	var relativePath = path
	var siblings []UserAccount
	for i := len(path); i > 0; i-- {
		otherChild := GetOtherChild(relativePath)
		otherNode, err := db.GetAccountByPath(otherChild)
		if err != nil {
			return siblings, err
		}
		siblings = append(siblings, otherNode)
		relativePath = GetParentPath(relativePath)
	}
	return siblings, nil
}

// GetAccount gets the account of the given path from the DB
func (db *DB) GetAccountByPath(path string) (UserAccount, error) {
	var account UserAccount
	// err := db.Instance.Model(&account).Where("path = ?", path).GetErrors()
	err := db.Instance.Where("path = ?", path).Find(&account).GetErrors()
	if len(err) != 0 {
		return account, ErrRecordNotFound(fmt.Sprintf("unable to find record for path: %v err:%v", path, err))
	}
	return account, nil
}

func (db *DB) GetAccountByHash(hash string) (UserAccount, error) {
	var account UserAccount
	if db.Instance.First(&account, hash).RecordNotFound() {
		return account, ErrRecordNotFound(fmt.Sprintf("unable to find record for hash: %v", hash))
	}
	return account, nil
}

func (db *DB) GetRoot() (UserAccount, error) {
	var account UserAccount
	// err := db.Instance.Model(&account).Where("path = ?", path).GetErrors()
	err := db.Instance.Where("level = ?", 0).Find(&account).GetErrors()
	if len(err) != 0 {
		return account, ErrRecordNotFound(fmt.Sprintf("unable to find record. err:%v", err))
	}
	return account, nil
}

func (db *DB) InsertCoordinatorAccount(acc *UserAccount, depth uint64) error {
	acc.UpdatePath(GenCoordinatorPath(depth))
	acc.CreateAccountHash()
	acc.Type = 1
	return db.Instance.Create(&acc).Error
}

// updateAccount will simply replace all the changed fields
func (db *DB) updateAccount(newAcc UserAccount, path string) error {
	return db.Instance.Model(&newAcc).Where("path = ?", path).Update(newAcc).Error
}

func (db *DB) GetAccountCount() (int, error) {
	var count int
	db.Instance.Table("user_accounts").Count(&count)
	return count, nil
}

// // GetDepositNodePath is supposed to get a set of uninitialised leaves
// // number of uninitialised nodes have to be == 2**MaxDepositSubTreeHeight
// // Then we return the path of this node
// func (db *DB) GetDepositNodePath() (path string, err error) {
// 	// get first uninitialised leaf
// 	firstLeaf, err := db.GetAccountByStatus(100)
// 	if err != nil {
// 		return
// 	}
// 	params, err := db.GetParams()
// 	if err != nil {
// 		return
// 	}
// 	numberOfLeaves := math.Exp2(float64(params.MaxDepositSubTreeHeight))

// 	lastLeafPath := firstLeaf.Path + uint64(numberOfLeaves)

// 	var accounts []UserAccount
// 	err = db.Instance.Where("path BETWEEN ? AND ? AND status==?", firstLeaf.Path, lastLeafPath, 100).Find(&accounts).Error
// 	if err != nil {
// 		return
// 	}

// 	fmt.Println("found accounts", accounts)
// 	pathToFirstLeafStr := UintToBigInt(firstLeaf.Path).String()
// 	depth := params.MaxDepth - params.MaxDepositSubTreeHeight
// 	return pathToFirstLeafStr[:depth], nil
// }
