package db

import (
	"fmt"

	"github.com/BOPR/common"
	"github.com/BOPR/config"
	"github.com/BOPR/types"
	merkle "github.com/cbergoon/merkletree"
	"github.com/globalsign/mgo"
	"gopkg.in/mgo.v2/bson"
)

func (db *DB) GetAccountCollection() (*mgo.Collection, *Session) {
	session := db.MgoSession.Copy()
	return session.GetCollection(common.DATABASE, common.ACCOUNT_COLLECTION), session
}

func (db *DB) StoreMT(mt merkle.MerkleTree) error {
	coll, session := db.GetAccountCollection()
	defer session.Close()

	if err := coll.Insert(mt); err != nil {
		fmt.Println("Unable to insert", "error", err)
		return err
	}
	return nil
}

func (db *DB) GetAllAccounts() (accs []types.AccountLeaf, err error) {
	coll, session := db.GetAccountCollection()
	defer session.Close()

	err = coll.Find(nil).All(accs)
	if err != nil {
		return accs, err
	}
	return accs, nil
}

// GetAccount gets the account of the given path from the DB
func (db *DB) GetAccount(accID uint64) (types.AccountLeaf, error) {
	coll, session := db.GetAccountCollection()
	defer session.Close()

	query := bson.M{"path": accID}

	var account types.AccountLeaf
	iter := coll.Find(query).Limit(1).Iter()
	err := iter.All(&account)
	if err != nil {
		return account, err
	}

	return account, nil
}

func (db *DB) InsertBulkAccounts(accounts []types.AccountLeaf) error {
	coll, session := db.GetAccountCollection()
	defer session.Close()

	var AccI []interface{}
	for _, acc := range accounts {
		AccI = append(AccI, acc)
	}
	bulk := coll.Bulk()
	bulk.Insert(AccI...)
	_, err := bulk.Run()
	if err != nil {
		return err
	}
	return nil
}

func (db *DB) InsertGenAccounts(genAccs []config.GenAccountLeaf) error {
	var accLeafs []types.AccountLeaf
	for _, acc := range genAccs {
		newAccLeaf := types.NewAccountLeaf(acc.Path, acc.Balance, acc.TokenType, acc.Nonce)
		accLeafs = append(accLeafs, newAccLeaf)
	}
	return db.InsertBulkAccounts(accLeafs)
}

func (db *DB) GetAccountCount() (int, error) {
	coll, session := db.GetAccountCollection()
	defer session.Close()

	return coll.Count()
}

// FetchSiblings retuns the siblings of an account leaf till root
// TODO make this more performannt by using bulk account fetch or using groutines to fetch in parerell
func FetchSiblings(accID uint64, db DB) (accs []types.AccountLeaf, err error) {
	// For eg: for account ID 1111 => 1110, 110X, 10XX
	var siblings []types.AccountLeaf

	return siblings, nil
}
