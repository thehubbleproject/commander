package types

type Batch struct {
	StateRoot ByteArray
	Committer Address
	TxRoot    ByteArray
}

func NewBatch(stateRoot ByteArray, committer Address, txRoot ByteArray) Batch {
	return Batch{
		StateRoot: stateRoot,
		Committer: committer,
		TxRoot:    txRoot,
	}
}

type BatchInfo struct {
	StateRoot ByteArray
	Index     uint64
}
