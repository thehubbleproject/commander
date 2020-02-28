package db

import (
	"github.com/BOPR/types"
)

// Insert tx into the DB
func (db *DB) InsertTx(t *types.Tx) error {
	db.Instance.Create(t)
	return nil
}

func (db *DB) PopTxs() (txs []types.Tx, err error) {
	// coll := db.GetTxCollection()

	// TODO fetch a limited set of transactions
	//
	// ids = db.collection.find(<condition>).limit(<limit>).map(
	// 	function(doc) {
	// 		return doc._id;
	// 	}
	// );
	// db.collection.updateMany({_id: {$in: ids}}, <update>})

	// query := bson.M{"status": "pending"}
	// updateTo := bson.M{"$set": bson.M{"status": "processed"}}

	// bulk := coll.Bulk()
	// bulk.UpdateAll(query, updateTo)
	// data, err := bulk.Run()
	// fmt.Println("data", data, err)
	return
}

func (db *DB) GetTx() (tx []types.Tx) {
	db.Instance.First(&tx)
	return
}
