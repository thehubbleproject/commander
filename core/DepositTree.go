package core

import (
	"errors"
	"fmt"
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

	generatedRoot, err := GetParent(left, right)
	if err != nil {
		return 0, err
	}

	if generatedRoot.String() != newRoot.String() {
		return 0, errors.New("Unable to update deposit tree, deposit tree root doesnt match")
	}
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
