package db

import (
	"encoding/json"
	"fmt"

	"github.com/BOPR/types"
)

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

func (db *DB) AddNewBatch(batch types.Batch, transactionsIncluded [][]byte) error {
	bz, err := json.Marshal(transactionsIncluded)
	if err != nil {
		return err
	}
	batch.TransactionsIncluded = bz
	return db.Instance.Create(batch).Error
}
