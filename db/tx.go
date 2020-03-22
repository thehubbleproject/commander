package db

import (
	"fmt"

	"github.com/BOPR/types"
)

// Insert tx into the DB
func (db *DB) InsertTx(t *types.Tx) error {
	db.Instance.Create(t)
	return nil
}

func (db *DB) PopTxs() (txs []types.Tx, err error) {
	tx := db.Instance.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return txs, err
	}
	var pendingTxs []types.Tx

	// select N number of transactions which are pending in mempool and
	if err := tx.Limit(3).Where(&types.Tx{Status: 100}).Find(&pendingTxs).Error; err != nil {
		fmt.Println("error while fetching pending transactions", err)
		return txs, err
	}

	fmt.Println("found txs", pendingTxs)

	var ids []uint
	for _, tx := range pendingTxs {
		ids = append(ids, tx.ID)
	}

	// update the transactions from pending to processing
	errs := tx.Table("txes").Where("id IN (?)", ids).Updates(map[string]interface{}{"status": "processing"}).GetErrors()
	fmt.Println("errors while processing transactions", errs)

	return pendingTxs, tx.Commit().Error
}

func (db *DB) GetTx() (tx []types.Tx) {
	db.Instance.First(&tx)
	return
}
