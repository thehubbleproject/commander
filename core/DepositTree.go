package core

import (
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
