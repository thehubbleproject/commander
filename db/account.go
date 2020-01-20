package db

import (
	"fmt"

	"github.com/BOPR/common"
	"github.com/BOPR/config"
	"github.com/BOPR/types"
	merkle "github.com/cbergoon/merkletree"
	"gopkg.in/mgo.v2/bson"
)

// FetchSiblings retuns the siblings of an account leaf till root
// TODO make this more performannt by using bulk account fetch or using groutines to fetch in parerell
func FetchSiblings(accID uint64) (accs []types.AccountLeaf, err error) {
	// For eg: for account ID 1111 => 1110, 110X, 10XX
	var siblings []types.AccountLeaf
	for i := common.DEFAULT_HEIGHT - 1; i < 0; i++ {
		accID := uint64(common.FlipBit(common.ExtractBit(int(accID), i)))
		acc, err := GetAccount(accID)
		if err != nil {
			return nil, err
		}
		siblings = append(siblings, acc)
	}
	return siblings, nil
}

func StoreMT(mt merkle.MerkleTree) error {
	session := MgoSession.Copy()
	defer session.Close()
	if err := session.GetCollection(common.DATABASE, common.ACCOUNT_COLLECTION).Insert(mt); err != nil {
		fmt.Println("Unable to insert", "error", err)
		return err
	}
	return nil
}

func CreateAndStoreMT(accounts []types.AccountLeaf) {
	// types.CreateTree()
}

func GetAllAccounts() (accs []types.AccountLeaf, err error) {
	session := MgoSession.Copy()
	defer session.Close()
	collection := session.GetCollection(common.DATABASE, common.ACCOUNT_COLLECTION)

	err = collection.Find(nil).All(accs)
	if err != nil {
		return accs, err
	}
	return accs, nil
}

// GetAccount gets the account of the given path from the DB
func GetAccount(accID uint64) (types.AccountLeaf, error) {
	session := MgoSession.Copy()
	defer session.Close()

	query := bson.M{"path": accID}

	var account types.AccountLeaf
	iter := session.GetCollection(common.DATABASE, common.ACCOUNT_COLLECTION).Find(query).Limit(1).Iter()
	err := iter.All(&account)
	if err != nil {
		return account, err
	}

	return account, nil
}

func InsertBulkAccounts(accounts []types.AccountLeaf) error {
	session := MgoSession.Copy()
	defer session.Close()
	collection := session.GetCollection(common.DATABASE, common.ACCOUNT_COLLECTION)
	var AccI []interface{}
	for _, acc := range accounts {
		AccI = append(AccI, acc)
	}
	bulk := collection.Bulk()
	bulk.Insert(AccI...)
	_, err := bulk.Run()
	if err != nil {
		return err
	}
	return nil
}

func InsertGenAccounts(genAccs []config.GenAccountLeaf) error {
	var accLeafs []types.AccountLeaf
	for _, acc := range genAccs {
		newAccLeaf := types.NewAccountLeaf(acc.Path, acc.Balance, acc.TokenType, acc.Nonce)
		accLeafs = append(accLeafs, newAccLeaf)
	}
	return InsertBulkAccounts(accLeafs)
}

func GetAccountCount() (int, error) {
	session := MgoSession.Copy()
	defer session.Close()
	coll := session.GetCollection(common.DATABASE, common.ACCOUNT_COLLECTION)
	return coll.Count()
}
