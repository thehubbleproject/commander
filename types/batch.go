package types

import (
	"math/big"

	"github.com/jinzhu/gorm"
)

// Batch is the batches that need to be submitted on-chain periodically
type Batch struct {
	gorm.Model

	Index                uint64
	StateRoot            ByteArray
	Committer            Address
	TxRoot               ByteArray
	StakeAmount          uint64
	FinalisesOn          big.Int
	SubmissionHash       Hash
	TransactionsIncluded [][]byte
}
