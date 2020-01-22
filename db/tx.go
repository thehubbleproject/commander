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

	// TODO fetch a limited set of transactions
	//
	// ids = db.collection.find(<condition>).limit(<limit>).map(
	// 	function(doc) {
	// 		return doc._id;
	// 	}
	// );
	// db.collection.updateMany({_id: {$in: ids}}, <update>})

	query := bson.M{"status": "pending"}
	updateTo := bson.M{"$set": bson.M{"status": "processed"}}
	collection := session.GetCollection(common.DATABASE, common.TRANSACTION_COLLECTION)
	bulk := collection.Bulk()
	bulk.UpdateAll(query, updateTo)
	data, err := bulk.Run()
	fmt.Println("data", data, err)
	return
}
