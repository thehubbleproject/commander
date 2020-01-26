package db

import (
	"fmt"

	"github.com/BOPR/common"
	"github.com/BOPR/types"
	"github.com/globalsign/mgo"
	"gopkg.in/mgo.v2/bson"
)

func (db *DB) GetTxCollection() *mgo.Collection {
	session := db.MgoSession.Copy()
	defer session.Close()
	return session.GetCollection(common.DATABASE, common.TRANSACTION_COLLECTION)
}

// Insert tx into the DB
func (db *DB) InsertTx(t *types.Tx) error {
	coll := db.GetTxCollection()
	if err := coll.Insert(t); err != nil {
		fmt.Println("Unable to insert", "error", err)
		return err
	}
	return nil
}

func (db *DB) PopTxs() (txs []types.Tx, err error) {
	coll := db.GetTxCollection()

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

	bulk := coll.Bulk()
	bulk.UpdateAll(query, updateTo)
	data, err := bulk.Run()
	fmt.Println("data", data, err)
	return
}
