package db

import (
	"encoding/hex"
	"fmt"

	"github.com/BOPR/types"
)

// InsertBatchInfo inserts batch info to db
func (db *DB) InsertBatchInfo(root types.ByteArray, index uint64) error {
	batch := types.BatchInfo{Index: index, StateRoot: hex.EncodeToString(root[:])}

	err := db.Instance.Create(&batch).Error
	if err != nil {
		return ErrUnableToCreateRecord(fmt.Sprintf("unable to insert new batch with index:%v and root:%v", index, root))
	}
	return nil
}

func (db *DB) GetAllBatches() (batches []types.Batch, err error) {
	errs := db.Instance.Find(&batches).GetErrors()
	for _, err := range errs {
		if err != nil {
			return batches, GenericError("got error while fetch all batches")
		}
	}
	return
}

func (db *DB) GetLatestBatch() (batch types.Batch, err error) {
	if err := db.Instance.First(&batch).Error; err != nil {
		return batch, ErrRecordNotFound(fmt.Sprintf("unable to find latest batch"))
	}
	return batch, nil
}

func (db *DB) GetBatchCount() (int, error) {
	var count int
	db.Instance.Table("batches").Count(&count)
	return count, nil
}
