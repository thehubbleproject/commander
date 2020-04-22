package core

import (
	"math/big"

	"github.com/BOPR/common"
	"github.com/BOPR/contracts/rollup"
	"github.com/ethereum/go-ethereum/accounts/abi"
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
	PublicKey string `gorm:"size:128"`

	// Path from root to leaf
	// NOTE: not a part of the leaf
	// Path is a string to that we can run LIKE queries
	Path string `gorm:"not null;index:Path"`

	// Pending = 0 means has deposit but not merged to balance tree
	// Active = 1
	// InActive = 2 => non leaf node
	// NonInitialised = 100
	Status uint64 `gorm:"not null;index:Status"`

	// level at which this node is stored in the balance tree
	// Level uint64 `gorm:"not null;index:Level"`

	// Type of nodes
	// 1 => terminal
	// 0 => root
	// 2 => non terminal
	Type uint64 `gorm:"not null;index:Type"`

	// keccak hash of the node
	Hash string `gorm:"not null;index:Hash"`
}

func NewUserAccount(id, balance, tokenType, nonce, status, level, nodeType uint64, pubkey, path string) *UserAccount {
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
	newAcccount.CreateAccountHash()
	return newAcccount
}

func (acc *UserAccount) ToABIAccount() rollup.DataTypesUserAccount {
	return rollup.DataTypesUserAccount{
		ID:        UintToBigInt(acc.AccountID),
		Balance:   UintToBigInt(acc.Balance),
		TokenType: UintToBigInt(acc.TokenType),
		Nonce:     UintToBigInt(acc.Nonce),
	}
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
		big.NewInt(int64(acc.TokenType)),
		big.NewInt(int64(acc.Nonce)),
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

func (db *DB) InitEmptyDepositTree() error {
	var depositTree DepositTree
	depositTree.Root = ZERO_VALUE_LEAF.String()
	return db.Instance.Create(&depositTree).Error
}

// GetAllAccounts fetches all accounts from the database
// func (db *DB) GetAllAccounts() (accs []UserAccount, err error) {
// 	// TODO add limits here
// 	errs := db.Instance.Find(&accs).GetErrors()
// 	for _, err := range errs {
// 		if err != nil {
// 			return accs, GenericError("got error while fetch all accounts")
// 		}
// 	}
// 	return accs, nil
// }

// // GetAccount gets the account of the given path from the DB
// func (db *DB) GetAccount(ID uint64) (UserAccount, error) {
// 	var account UserAccount
// 	if db.Instance.First(&account, ID).RecordNotFound() {
// 		return account, ErrRecordNotFound(fmt.Sprintf("unable to find record for accountID: %d", ID))
// 	}
// 	return account, nil
// }

// func (db *DB) InsertAccount(account UserAccount) error {
// 	return db.Instance.Create(account).Error
// }

// func (db *DB) InsertBulkAccounts(accounts []UserAccount) error {
// 	for _, account := range accounts {
// 		err := db.InsertAccount(account)
// 		if err != nil {
// 			return ErrUnableToCreateRecord(fmt.Sprintf("Unable to add account with ID:%v to DB. Error: %v", account.AccountID, err))
// 		}
// 	}
// 	return nil
// }

// func (db *DB) InsertGenAccounts(genAccs []config.GenUserAccount) error {
// 	var accLeafs []UserAccount
// 	for _, acc := range genAccs {
// 		newAccLeaf := NewUserAccount(acc.ID, acc.Balance, acc.TokenType, acc.Path, acc.Nonce, int(acc.Status), acc.PublicKey)
// 		accLeafs = append(accLeafs, *newAccLeaf)
// 	}
// 	return db.InsertBulkAccounts(accLeafs)
// }

// func (db *DB) GetAccountCount() (int, error) {
// 	var count int
// 	db.Instance.Table("account_leafs").Count(&count)
// 	return count, nil
// }

// // FetchSiblings retuns the siblings of an account leaf till root
// // TODO make this more performannt by using bulk account fetch or using groutines to fetch in parerell
// func FetchSiblings(accID uint64, db DB) (accs []UserAccount, err error) {
// 	// For eg: for account ID 1111 => 1110, 110X, 10XX
// 	var siblings []UserAccount

// 	return siblings, nil
// }

// func (db *DB) InitEmptyDepositTree() error {
// 	var depositTree DepositTree
// 	depositTree.Root = ZERO_VALUE_LEAF.String()
// 	return db.Instance.Create(&depositTree).Error
// }

// func (db *DB) OnDepositLeafMerge(left, right, newRoot ByteArray) (uint64, error) {
// 	// get last deposit from deposit tree
// 	var lastDeposit DepositTree
// 	err := db.Instance.First(&lastDeposit).Error
// 	if err != nil {
// 		return 0, err
// 	}

// 	// update the deposit tree stored
// 	var updatedDepositTreeInfo DepositTree
// 	updatedDepositTreeInfo.Height = lastDeposit.Height + 1
// 	updatedDepositTreeInfo.NumberOfDeposits = lastDeposit.NumberOfDeposits + 2
// 	updatedDepositTreeInfo.Root = newRoot.String()

// 	generatedRoot, err := GetParent(left, right)
// 	if err != nil {
// 		return 0, err
// 	}

// 	if generatedRoot.String() != newRoot.String() {
// 		return 0, errors.New("Unable to update deposit tree, deposit tree root doesnt match")
// 	}
// 	if err := db.Instance.Model(&lastDeposit).Update(&updatedDepositTreeInfo).Error; err != nil {
// 		return 0, err
// 	}

// 	return updatedDepositTreeInfo.Height, nil
// }

// func (db *DB) GetDepositTreeInfo() (dt DepositTree, err error) {
// 	err = db.Instance.First(&dt).Error
// 	if err != nil {
// 		return
// 	}
// 	return
// }

// func (db *DB) GetPendingDeposits(numberOfAccs uint64) ([]UserAccount, error) {
// 	var accounts []UserAccount
// 	err := db.Instance.Limit(numberOfAccs).Where("status = ?", 0).Find(&accounts).Error
// 	if err != nil {
// 		return accounts, err
// 	}
// 	return accounts, nil
// }

// func (db *DB) FinaliseDeposits(accountsRoot ByteArray, pathToDepositSubTree uint64, newBalanceRoot ByteArray) error {
// 	db.Logger.Info("Finalising accounts", "accountRoot", accountsRoot, "NewBalanceRoot", newBalanceRoot, "pathToDepositSubTree", pathToDepositSubTree)

// 	// get params
// 	params, err := db.GetParams()
// 	if err != nil {
// 		return err
// 	}

// 	// number of new deposits = 2**MaxDepthOfDepositTree
// 	depositCount := uint64(math.Exp2(float64(params.MaxDepositSubTreeHeight)))

// 	// get all pending accounts
// 	pendingAccs, err := db.GetPendingDeposits(depositCount)
// 	if err != nil {
// 		return err
// 	}
// 	db.Logger.Debug("Fetched pending deposits", "count", len(pendingAccs), "data", pendingAccs)

// 	// update the empty leaves with new accounts
// 	err = db.AddPendingDeposits(pendingAccs)
// 	if err != nil {
// 		return err
// 	}

// 	// TODO ensure the accounts are inserted at pathToDepositSubTree

// 	//TODO  make sure all the accounts root match to accountsRoot

// 	return nil
// }

// func (db *DB) AddPendingDeposits(pendingAccs []UserAccount) error {
// 	var accounts []UserAccount
// 	// fetch 2**DepositSubTree inactive accounts ordered by path
// 	err := db.Instance.Limit(len(pendingAccs)).Order("path").Where("status = ?", 100).Find(&accounts).Error
// 	if err != nil {
// 		return err
// 	}
// 	// TODO add error for if no account found

// 	// update the accounts
// 	for i, acc := range accounts {
// 		// acc.Balance = pendingAccs[i].Balance
// 		// acc.TokenType = pendingAccs[i].TokenType
// 		// acc.AccountID = pendingAccs[i].AccountID
// 		// acc.PublicKey = pendingAccs[i].PublicKey
// 		err := db.Instance.Model(&acc).Updates(UserAccount{Balance: pendingAccs[i].Balance,
// 			TokenType: pendingAccs[i].TokenType,
// 			AccountID: pendingAccs[i].AccountID,
// 			PublicKey: pendingAccs[i].PublicKey,
// 			Status:    1,
// 		}).Error
// 		if err != nil {
// 			return err
// 		}
// 	}
// 	return nil
// }

// // create merkle proof for finalisation of deposits
// // send transaction to etherum chain using contract caller
// func (db *DB) sendDepositFinalisationTx() {

// }

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

// func (db *DB) GetAccountByStatus(status uint64) (UserAccount, error) {
// 	var leaf UserAccount
// 	err := db.Instance.Order("path").Where("status = ?", status).First(&leaf).Error
// 	return leaf, err
// }

// func (db *DB) GetAccountLessThanPath(path uint64) ([]UserAccount, error) {
// 	var accounts []UserAccount
// 	err := db.Instance.Where("path BETWEEN ? AND ?", 0, path, 100).Find(&accounts).Error
// 	return accounts, err
// }
