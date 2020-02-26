package db

import (
	"fmt"

	"github.com/BOPR/common"
	"github.com/BOPR/types"
	"github.com/globalsign/mgo"
)

// GetBatchCollection fetches the batch collection
func (db *DB) GetBatchCollection() (*mgo.Collection, *Session) {
	session := db.MgoSession.Copy()
	return session.GetCollection(common.DATABASE, common.BATCH_COLLECTION), session
}

// InsertBatchInfo inserts batch info to db
func (db *DB) InsertBatchInfo(root types.ByteArray, index uint64) error {
	coll, session := db.GetBatchCollection()
	defer session.Close()

	batch := types.BatchInfo{StateRoot: root, Index: index}
	if err := coll.Insert(batch); err != nil {
		fmt.Println("Unable to insert", "error", err)
		return err
	}
	return nil
}

func (db *DB) GetAllBatches() (batches []types.Batch, err error) {
	coll, session := db.GetBatchCollection()
	defer session.Close()

	if err := coll.Find(nil).All(batches); err != nil {
		return batches, err
	}
	return
}

func (db *DB) GetLatestBatch() (types.Batch, error) {
	coll, session := db.GetBatchCollection()
	defer session.Close()
	var batches []types.Batch
	if err := coll.Find(nil).All(batches); err != nil {
		return batches[0], err
	}
	return batches[0], nil
}

func (db *DB) GetBatchCount() (int, error) {
	coll, session := db.GetBatchCollection()
	defer session.Close()

	return coll.Count()
}
