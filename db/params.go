package db

import "github.com/BOPR/types"

func (db *DB) UpdateSyncStatusWithBatchNumber(batchIndex uint64) error {
	var updatedSyncStatus types.SyncStatus
	updatedSyncStatus.LastBatchRecorded = batchIndex
	if err := db.Instance.Table("sync_statuses").Assign(types.SyncStatus{LastBatchRecorded: batchIndex}).FirstOrCreate(&updatedSyncStatus).Error; err != nil {
		return err
	}
	return nil
}
func (db *DB) UpdateSyncStatusWithBlockNumber(blkNum uint64) error {
	var updatedSyncStatus types.SyncStatus
	updatedSyncStatus.LastEthBlockRecorded = blkNum
	if err := db.Instance.Table("sync_statuses").Assign(types.SyncStatus{LastEthBlockRecorded: blkNum}).FirstOrCreate(&updatedSyncStatus).Error; err != nil {
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

// UpdateStakeAmount updates the stake amount
func (db *DB) UpdateStakeAmount(newStakeAmount uint64) error {
	var updatedParams types.Params
	updatedParams.StakeAmount = newStakeAmount
	if err := db.Instance.Table("params").Assign(types.Params{StakeAmount: newStakeAmount}).FirstOrCreate(&updatedParams).Error; err != nil {
		return err
	}
	return nil
}

// UpdateMaxDepth updates the max depth
func (db *DB) UpdateMaxDepth(newDepth uint64) error {
	var updatedParams types.Params
	updatedParams.MaxDepth = newDepth
	if err := db.Instance.Table("params").Assign(types.Params{MaxDepth: newDepth}).FirstOrCreate(&updatedParams).Error; err != nil {
		return err
	}
	return nil
}

// UpdateDepositSubTreeHeight updates the max height of deposit sub tree
func (db *DB) UpdateDepositSubTreeHeight(newHeight uint64) error {
	var updatedParams types.Params
	updatedParams.MaxDepositSubTreeHeight = newHeight
	if err := db.Instance.Table("params").Assign(types.Params{MaxDepositSubTreeHeight: newHeight}).FirstOrCreate(&updatedParams).Error; err != nil {
		return err
	}
	return nil
}

// GetParams gets params from the DB
func (db *DB) GetParams() (params types.Params, err error) {
	if err := db.Instance.First(&params).Error; err != nil {
		return params, err
	}
	return params, nil
}
