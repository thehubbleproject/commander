package types

import "encoding/hex"

type Batch struct {
	StateRoot string
	Committer Address
	TxRoot    string
}

func NewBatch(stateRoot ByteArray, committer Address, txRoot ByteArray) Batch {
	return Batch{
		StateRoot: hex.EncodeToString(stateRoot[:]),
		Committer: committer,
		TxRoot:    hex.EncodeToString(txRoot[:]),
	}
}

type BatchInfo struct {
	StateRoot string
	Index     uint64
}
