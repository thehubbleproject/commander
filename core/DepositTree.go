package core

import (
	"fmt"
	"math"
)

type DepositTree struct {
	Height           uint64
	NumberOfDeposits uint64
	Root             string
}

func (t *DepositTree) String() string {
	return fmt.Sprintf("DepositTree: H:%v Count:%v Root:%v", t.Height, t.NumberOfDeposits, t.Root)
}

func (db *DB) InitEmptyDepositTree() error {
	var depositTree DepositTree
	depositTree.Root = ZERO_VALUE_LEAF.String()
	return db.Instance.Create(&depositTree).Error
}

func (db *DB) OnDepositLeafMerge(left, right, newRoot ByteArray) (uint64, error) {
	// get last deposit from deposit tree
	var lastDeposit DepositTree
	err := db.Instance.First(&lastDeposit).Error
	if err != nil {
		return 0, err
	}

	// update the deposit tree stored
	var updatedDepositTreeInfo DepositTree
	updatedDepositTreeInfo.Height = lastDeposit.Height + 1
	updatedDepositTreeInfo.NumberOfDeposits = lastDeposit.NumberOfDeposits + 2
	updatedDepositTreeInfo.Root = newRoot.String()

	if err := db.Instance.Model(&lastDeposit).Update(&updatedDepositTreeInfo).Error; err != nil {
		return 0, err
	}

	return updatedDepositTreeInfo.Height, nil
}

func (db *DB) GetDepositNodeAndSiblings() (NodeToBeReplaced UserAccount, siblings []UserAccount, err error) {
	// get params
	params, err := db.GetParams()
	if err != nil {
		return
	}

	// get the deposit node
	expectedHash := defaultHashes[params.MaxDepositSubTreeHeight]

	// getNode with the expectedHash
	NodeToBeReplaced, err = db.GetAccountByHash(expectedHash.String())
	if err != nil {
		return
	}

	// get siblings for the path to node
	siblings, err = db.GetSiblings(NodeToBeReplaced.Path)
	if err != nil {
		return
	}

	return
}

func (db *DB) FinaliseDepositsAndAddBatch(accountsRoot ByteArray, pathToDepositSubTree uint64, newBalanceRoot ByteArray) error {
	db.Logger.Info("Finalising accounts", "accountRoot", accountsRoot, "NewBalanceRoot", newBalanceRoot, "pathToDepositSubTree", pathToDepositSubTree)

	// get params
	params, err := db.GetParams()
	if err != nil {
		return err
	}

	// number of new deposits = 2**MaxDepthOfDepositTree
	depositCount := uint64(math.Exp2(float64(params.MaxDepositSubTreeHeight)))

	// get all pending accounts
	pendingAccs, err := db.GetPendingDeposits(depositCount)
	if err != nil {
		return err
	}

	db.Logger.Debug("Fetched pending deposits", "count", len(pendingAccs), "data", pendingAccs)

	// update the empty leaves with new accounts
	err = db.AddPendingDeposits(pendingAccs)
	if err != nil {
		return err
	}

	// TODO ensure the accounts are inserted at pathToDepositSubTree

	//TODO  make sure all the accounts root match to accountsRoot

	return nil
}

func (db *DB) AddPendingDeposits(pendingAccs []UserAccount) error {
	var accounts []UserAccount

	// fetch 2**DepositSubTree inactive accounts ordered by path
	err := db.Instance.Limit(len(pendingAccs)).Order("path").Where("status = ?", 100).Find(&accounts).Error
	if err != nil {
		return err
	}
	// TODO add error for if no account found

	// update the accounts
	for i, acc := range accounts {
		err := db.Instance.Model(&acc).Updates(UserAccount{Balance: pendingAccs[i].Balance,
			TokenType: pendingAccs[i].TokenType,
			AccountID: pendingAccs[i].AccountID,
			PublicKey: pendingAccs[i].PublicKey,
			Status:    1,
		}).Error
		if err != nil {
			return err
		}
	}
	return nil
}

// create merkle proof for finalisation of deposits
// send transaction to etherum chain using contract caller
func (db *DB) sendDepositFinalisationTx() {

}

func (db *DB) GetPendingDeposits(numberOfAccs uint64) ([]UserAccount, error) {
	var accounts []UserAccount
	err := db.Instance.Limit(numberOfAccs).Where("status = ?", 0).Find(&accounts).Error
	if err != nil {
		return accounts, err
	}
	return accounts, nil
}
