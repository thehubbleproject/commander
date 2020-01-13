package db

import (
	"fmt"

	"github.com/BOPR/common"
	"github.com/BOPR/types"
)

// InsertBatchInfo inserts batch info to db
func InsertBatchInfo(root types.ByteArray, index uint64) error {
	session := MgoSession.Copy()
	defer session.Close()
	batch := types.BatchInfo{StateRoot: root, Index: index}
	if err := session.GetCollection(common.DATABASE, common.BATCH_COLLECTION).Insert(batch); err != nil {
		fmt.Println("Unable to insert", "error", err)
		return err
	}
	return nil
}

func GetAllBatches() (batches []types.Batch, err error) {
	session := MgoSession.Copy()
	defer session.Close()
	c := session.GetCollection(common.DATABASE, common.BATCH_COLLECTION)
	if err := c.Find(nil).All(batches); err != nil {
		return batches, err
	}
	return
}

func GetLatestBatch() (types.Batch, error) {
	session := MgoSession.Copy()
	defer session.Close()
	c := session.GetCollection(common.DATABASE, common.BATCH_COLLECTION)
	var batches []types.Batch
	if err := c.Find(nil).All(batches); err != nil {
		return batches[0], err
	}
	return batches[0], nil
}

func GetBatchCount() (int, error) {
	session := MgoSession.Copy()
	defer session.Close()
	coll := session.GetCollection(common.DATABASE, common.BATCH_COLLECTION)
	return coll.Count()
}
