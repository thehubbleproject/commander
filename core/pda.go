package core

import (
	"encoding/hex"
	"errors"
	"fmt"
	"math"

	"github.com/BOPR/common"
	"github.com/ethereum/go-ethereum/accounts/abi"
	gormbulk "github.com/t-tiger/gorm-bulk-insert"
)

// PDA is the copy of the accounts tree
type PDA struct {
	// ID is the path of the user account in the PDA Tree
	// Cannot be changed once created
	AccountID uint64 `gorm:"not null;index:AccountID"`

	// Public key for the user
	PublicKey string `gorm:"type:varchar(1000)"`

	// Path from root to leaf
	// NOTE: not a part of the leaf
	// Path is a string to that we can run LIKE queries
	Path string `gorm:"not null;index:Path"`

	// Type of nodes
	// 1 => terminal
	// 0 => root
	// 2 => non terminal
	Type uint64 `gorm:"not null;index:Type"`

	// keccak hash of the node
	Hash string `gorm:"not null;index:Hash"`

	Level uint64 `gorm:"not null;index:Level"`
}

func NewPDA(id uint64, pubkey, path string) (*PDA, error) {
	newPDALeaf := &PDA{
		AccountID: id,
		PublicKey: pubkey,
		Path:      path,
		Type:      TYPE_TERMINAL,
	}
	err := newPDALeaf.PopulateHash()
	if err != nil {
		return nil, err
	}
	return newPDALeaf, nil
}

func NewEmptyPDA() *PDA {
	return &PDA{AccountID: ZERO, PublicKey: "", Type: TYPE_TERMINAL}
}

func NewPDANode(path, hash string) *PDA {
	newPDALeaf := &PDA{
		AccountID: ZERO,
		Path:      path,
		Type:      TYPE_NON_TERMINAL,
	}
	newPDALeaf.UpdatePath(path)
	newPDALeaf.Hash = hash
	return newPDALeaf
}

func (p *PDA) UpdatePath(path string) {
	p.Path = path
	p.Level = uint64(len(path))
}

func (p *PDA) HashToByteArray() ByteArray {
	ba, err := HexToByteArray(p.Hash)
	if err != nil {
		panic(err)
	}
	return ba
}

func (p *PDA) PopulateHash() error {
	if p.PublicKey == "" {
		p.Hash = ZERO_VALUE_LEAF.String()
		return nil
	}
	bz, err := abiEncodePubkey(p.PublicKey)
	if err != nil {
		return err
	}
	p.Hash = common.Keccak256(bz).String()
	return nil
}

func (db *DB) UpdatePDALeaf(leaf PDA) error {
	db.Logger.Info("Updated account pubkey", "ID", leaf.AccountID)
	leaf.PopulateHash()
	siblings, err := db.GetPDASiblings(leaf.Path)
	if err != nil {
		return err
	}

	db.Logger.Debug("Updating account", "Hash", leaf.Hash, "Path", leaf.Path, "siblings", siblings, "countOfSiblings", len(siblings))
	return db.StorePDALeaf(leaf, leaf.Path, siblings)
}

func (db *DB) GetPDASiblings(path string) ([]PDA, error) {
	var relativePath = path
	var siblings []PDA
	for i := len(path); i > 0; i-- {
		otherChild := GetOtherChild(relativePath)
		otherNode, err := db.GetPDALeafByPath(otherChild)
		if err != nil {
			return siblings, err
		}
		siblings = append(siblings, otherNode)
		relativePath = GetParentPath(relativePath)
	}
	return siblings, nil
}

func (db *DB) StorePDALeaf(pdaLeaf PDA, path string, siblings []PDA) error {
	var err error
	computedNode := pdaLeaf
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
			// Store the node!
			err = db.StorePDANode(parentHash, computedNode, sibling)
			if err != nil {
				return err
			}
		} else {
			parentHash, err = GetParent(sibling.HashToByteArray(), computedNode.HashToByteArray())
			if err != nil {
				return err
			}
			// Store the node!
			err = db.StorePDANode(parentHash, sibling, computedNode)
			if err != nil {
				return err
			}
		}

		parentAccount, err := db.GetPDALeafByPath(GetParentPath(computedNode.Path))
		if err != nil {
			return err
		}
		computedNode = parentAccount
	}

	// Store the new root
	err = db.UpdatePDARootNodes(computedNode.HashToByteArray())
	if err != nil {
		return err
	}

	return nil
}

// StoreNode updates the nodes given the parent hash
func (db *DB) StorePDANode(parentHash ByteArray, leftNode, rightNode PDA) (err error) {
	// update left account
	err = db.updatePDALeaf(leftNode, leftNode.Path)
	if err != nil {
		return err
	}
	// update right account
	err = db.updatePDALeaf(rightNode, rightNode.Path)
	if err != nil {
		return err
	}
	// update the parent with the new hashes
	return db.UpdateParentPDAWithHash(GetParentPath(leftNode.Path), parentHash)
}

// updateAccount will simply replace all the changed fields
func (db *DB) updatePDALeaf(newPDA PDA, path string) error {
	return db.Instance.Model(&newPDA).Where("path = ?", path).Update(newPDA).Error
}

func (db *DB) UpdateParentPDAWithHash(pathToParent string, newHash ByteArray) error {
	// Update the root hash
	var tempAccount PDA
	tempAccount.Path = pathToParent
	tempAccount.Hash = newHash.String()
	return db.updatePDALeaf(tempAccount, pathToParent)
}

// GetAccount gets the account of the given path from the DB
func (db *DB) GetPDALeafByPath(path string) (PDA, error) {
	var pdaLeaf PDA
	err := db.Instance.Where("path = ?", path).Find(&pdaLeaf).GetErrors()
	if len(err) != 0 {
		return pdaLeaf, ErrRecordNotFound(fmt.Sprintf("unable to find record for path: %v err:%v", path, err))
	}
	return pdaLeaf, nil
}

func (db *DB) UpdatePDARootNodes(newRoot ByteArray) error {
	var tempPDALeaf PDA
	tempPDALeaf.Path = ""
	tempPDALeaf.Hash = newRoot.String()
	return db.updatePDALeaf(tempPDALeaf, tempPDALeaf.Path)
}

func (db *DB) InitPDATree(depth uint64, genesisPDA []PDA) error {
	// calculate total number of leaves
	totalLeaves := math.Exp2(float64(depth))
	if int(totalLeaves) != len(genesisPDA) {
		return errors.New("Depth and number of leaves do not match")
	}
	db.Logger.Debug("Attempting to init balance tree", "totalAccounts", totalLeaves)

	var err error

	// insert coodinator leaf
	err = db.InsertCoordinatorPubkeyAccounts(&genesisPDA[0], depth)
	if err != nil {
		db.Logger.Error("Unable to insert coodinator account", "err", err)
		return err
	}

	var insertRecords []interface{}
	prevNodePath := genesisPDA[0].Path

	for i := 1; i < len(genesisPDA); i++ {
		pathToAdjacentNode, err := GetAdjacentNodePath(prevNodePath)
		if err != nil {
			return err
		}
		genesisPDA[i].UpdatePath(pathToAdjacentNode)
		insertRecords = append(insertRecords, genesisPDA[i])
		prevNodePath = genesisPDA[i].Path
	}

	db.Logger.Info("Inserting all leaves to DB", "count", len(insertRecords))
	err = gormbulk.BulkInsert(db.Instance, insertRecords, len(insertRecords))
	if err != nil {
		db.Logger.Error("Unable to insert leaves to DB", "err", err)
		return errors.New("Unable to insert leaves")
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
		var nextLevelAccounts []interface{}

		// iterate over 2 at a time and create next level
		for i := 0; i < len(accs); i += 2 {
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
			newAccNode := *NewPDANode(parentPath, parentHash.String())
			nextLevelAccounts = append(nextLevelAccounts, newAccNode)
		}

		err = gormbulk.BulkInsert(db.Instance, nextLevelAccounts, len(nextLevelAccounts))
		if err != nil {
			db.Logger.Error("Unable to insert PDA leaves to DB", "err", err)
			return errors.New("Unable to insert PDA leaves")
		}
	}
	// mark the root node type correctly
	return nil
}

// InsertCoordinatorPubkeyAccounts inserts the coordinator accounts
func (db *DB) InsertCoordinatorPubkeyAccounts(coordinatorPDA *PDA, depth uint64) error {
	coordinatorPDA.UpdatePath(GenCoordinatorPath(depth))
	fmt.Println("coordinator PDA", coordinatorPDA.Hash, coordinatorPDA.PublicKey)

	coordinatorPDA.PopulateHash()

	fmt.Println("coordinator PDA", coordinatorPDA.Hash, coordinatorPDA.PublicKey)
	coordinatorPDA.Type = TYPE_TERMINAL
	return db.Instance.Create(&coordinatorPDA).Error
}

func abiEncodePubkey(pubkey string) ([]byte, error) {
	pubkeyBytes, err := hex.DecodeString(pubkey)
	if err != nil {
		panic(err)
	}
	uint256Ty, err := abi.NewType("bytes", "bytes", nil)
	if err != nil {
		return []byte(""), err
	}

	arguments := abi.Arguments{
		{
			Type: uint256Ty,
		},
	}

	bytes, err := arguments.Pack(
		pubkeyBytes,
	)

	if err != nil {
		return []byte(""), err
	}

	return bytes, nil
}

func (db *DB) GetPDALeafByID(ID uint64) (PDA, error) {
	var pda PDA
	if err := db.Instance.Where("account_id = ?", ID).Find(&pda).Error; err != nil {
		return pda, ErrRecordNotFound(fmt.Sprintf("unable to find record for ID: %v", ID))
	}
	return pda, nil
}

func (db *DB) GetPDARoot() (PDA, error) {
	var PDAAccount PDA
	err := db.Instance.Where("level = ?", 0).Find(&PDAAccount).GetErrors()
	if len(err) != 0 {
		return PDAAccount, ErrRecordNotFound(fmt.Sprintf("unable to find record. err:%v", err))
	}
	return PDAAccount, nil
}
