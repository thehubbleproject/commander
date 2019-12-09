package types

type Batch struct {
	StateRoot        ByteArray
	WithdrawRoot     ByteArray
	Committer        Address
	AccountTreeState ByteArray
	TxRoot           ByteArray
}

func NewBatch(stateRoot ByteArray, withdrawRoot ByteArray, committer Address, accountRoot ByteArray, txRoot ByteArray) Batch {
	return Batch{
		StateRoot:        stateRoot,
		WithdrawRoot:     withdrawRoot,
		Committer:        committer,
		AccountTreeState: accountRoot,
		TxRoot:           txRoot,
	}
}

type BatchInfo struct {
	StateRoot ByteArray
	Index     uint64
}
