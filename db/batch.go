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
