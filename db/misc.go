package db

import (
	"github.com/BOPR/types"
)

func (db *DB) AddToken(t types.Token) error {
	return db.Instance.Create(t).Error
}

func (db *DB) GetTokenByID(id uint) (token types.Token, err error) {
	err = db.Instance.Where("token_id = ?", id).First(&token).Error
	if err != nil {
		return
	}
	return
}

func (db *DB) UpdateDepositTreeInfo(dt types.DepositTree) error {
	return db.Instance.Create(dt).Error
}

func (db *DB) OnDepositLeafMerge(left, right, newRoot types.ByteArray) (uint64, error) {
	var lastDeposit types.DepositTree
	err := db.Instance.First(&lastDeposit).Error
	if err != nil {
		return 0, err
	}
	var updatedDepositTreeInfo types.DepositTree
	updatedDepositTreeInfo.Height = lastDeposit.Height + 1
	updatedDepositTreeInfo.NumberOfDeposits = lastDeposit.NumberOfDeposits + 2
	updatedDepositTreeInfo.Root = newRoot
	// TODO utils.getParent(left,right)==newRoot, else error
	if err := db.UpdateDepositTreeInfo(updatedDepositTreeInfo); err != nil {
		return 0, err
	}

	return updatedDepositTreeInfo.Height, nil
}

func (db *DB) GetDepositTreeInfo() (dt types.DepositTree, err error) {
	err = db.Instance.First(&dt).Error
	if err != nil {
		return
	}
	return
}

func (db *DB) FinaliseDeposits(accountsRoot types.ByteArray, pathToDepositSubTree uint64, newBalanceRoot types.ByteArray) {
	// get all pending accounts

	// update paths

	// make sure all the accounts root match to accountsRoot
}

// create merkle proof for finalisation of deposits
// send transaction to etherum chain using contract caller
func (db *DB) sendDepositFinalisationTx() {

}
