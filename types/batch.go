package types

import (
	"encoding/hex"
	"math/big"

	"github.com/jinzhu/gorm"
)

// Batch is the batches that need to be submitted on-chain periodically
type Batch struct {
	gorm.Model

	Index          uint64
	StateRoot      string
	Committer      Address
	TxRoot         string
	StakeAmount    uint64
	FinalisesOn    big.Int
	SubmissionHash Hash
}

// NewBatch creates new batch
func NewBatch(stateRoot ByteArray, committer Address, txRoot ByteArray) Batch {
	return Batch{
		StateRoot: hex.EncodeToString(stateRoot[:]),
		Committer: committer,
		TxRoot:    hex.EncodeToString(txRoot[:]),
	}
}
