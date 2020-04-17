package db

import "github.com/BOPR/types"

func (db *DB) UpdateSyncStatusWithBatchNumber(batchIndex uint64) error {
	var updatedSyncStatus types.SyncStatus
	updatedSyncStatus.LastBatchRecorded = batchIndex
	if err := db.Instance.Table("params").Assign(types.SyncStatus{LastBatchRecorded: batchIndex}).FirstOrCreate(&updatedSyncStatus).Error; err != nil {
		return err
	}
	return nil
}
func (db *DB) UpdateSyncStatusWithBlockNumber(blkNum uint64) error {
	var updatedSyncStatus types.SyncStatus
	updatedSyncStatus.LastEthBlockRecorded = blkNum
	if err := db.Instance.Table("params").Assign(types.SyncStatus{LastEthBlockRecorded: blkNum}).FirstOrCreate(&updatedSyncStatus).Error; err != nil {
		return err
	}
	return nil
}

func (db *DB) GetSyncStatus() (status types.SyncStatus, err error) {
	if err := db.Instance.First(&status).Error; err != nil {
		return status, err
	}
	return status, nil
}
