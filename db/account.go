package db

import (
	"fmt"

	"github.com/BOPR/common"
	"github.com/BOPR/types"
	merkle "github.com/cbergoon/merkletree"
	"gopkg.in/mgo.v2/bson"
)

// FetchSiblings retuns the siblings of an account leaf till root
func FetchSiblings(accID uint64) {

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
