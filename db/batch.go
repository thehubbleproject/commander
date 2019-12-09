package db

import (
	"fmt"
	"github.com/BOPR/common"
	"github.com/BOPR/types"
)

// InsertBatchRoot inserts a batch root to db
func InsertBatchRoot(root types.ByteArray, index uint64) error {
	session := MgoSession.Copy()
	defer session.Close()
	if err := session.GetCollection(common.DATABASE, common.BATCH_COLLECTION).Insert(types.BatchInfo{StateRoot: root, Index: index}); err != nil {
		fmt.Println("Unable to insert", "error", err)
		return err
	}
	return nil
}
