package db

import (
	"fmt"
	"github.com/BOPR/common"
	"github.com/BOPR/types"
	"gopkg.in/mgo.v2/bson"
)

// Insert tx into the DB
func InsertTx(t *types.Tx) error {
	session := MgoSession.Copy()
	defer session.Close()
	if err := session.GetCollection(common.DATABASE, common.TRANSACTION_COLLECTION).Insert(t); err != nil {
		fmt.Println("Unable to insert", "error", err)
		return err
	}
	return nil
}

func PopTxs() (txs []types.Tx, err error) {
	session := MgoSession.Copy()
	defer session.Close()
	// selector := bson.M{"status": "pending"}
	// update := bson.M{"$set": bson.M{"status": "processed"}}

	query := bson.M{"status": "pending"}

	//Select Limit
	iter := session.GetCollection(common.DATABASE, common.TRANSACTION_COLLECTION).Find(query).Limit(2).Iter()
	err = iter.All(&txs)
	if err != nil {
		fmt.Println("Error iterating over transactions", "Error", err)
		return
	}
	return
}
