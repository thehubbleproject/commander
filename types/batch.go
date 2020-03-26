package types

import (
	"math/big"

	"encoding/json"
)

// Batch is the batches that need to be submitted on-chain periodically
type Batch struct {
	Index                uint64
	StateRoot            ByteArray
	Committer            string
	TxRoot               ByteArray
	StakeAmount          uint64
	FinalisesOn          big.Int
	SubmissionHash       string
	TransactionsIncluded [][]byte
}

func (b *Batch) DBModel() (BatchModel, error) {
	encodedTxs, err := json.Marshal(b.TransactionsIncluded)
	if err != nil {
		return BatchModel{}, err
	}
	encodedStateRoot, err := json.Marshal(b.StateRoot)
	if err != nil {
		return BatchModel{}, err
	}
	encodedTxRoot, err := json.Marshal(b.TxRoot)
	if err != nil {
		return BatchModel{}, err
	}
	finalisationBlockBytes := b.FinalisesOn.Bytes()
	newBatchModel := BatchModel{
		Index:                b.Index,
		StateRoot:            encodedStateRoot,
		Committer:            b.Committer,
		TxRoot:               encodedTxRoot,
		StakeAmount:          b.StakeAmount,
		FinalisesOn:          finalisationBlockBytes,
		SubmissionHash:       b.SubmissionHash,
		TransactionsIncluded: encodedTxs,
	}
	return newBatchModel, nil
}

// BatchModel represents the actual stuff stored in the DB
// We are encoding the whole struct because we will save some operations
// if we can just read the model for data in some cases and only decode when we need the encoded data
type BatchModel struct {
	DBModel

	Index                uint64
	StateRoot            []byte
	Committer            string
	TxRoot               []byte
	StakeAmount          uint64
	FinalisesOn          []byte
	SubmissionHash       string
	TransactionsIncluded []byte
}

func (b *BatchModel) Batch() (Batch, error) {
	var decodedTxs [][]byte
	err := json.Unmarshal(b.TransactionsIncluded, &decodedTxs)
	if err != nil {
		return Batch{}, err
	}
	var decodedStateRoot ByteArray
	err = json.Unmarshal(b.StateRoot, &decodedStateRoot)
	if err != nil {
		return Batch{}, err
	}
	var decodedTxRoot ByteArray
	err = json.Unmarshal(b.TxRoot, &decodedTxRoot)
	if err != nil {
		return Batch{}, err
	}
	finalisationBlockBN := big.NewInt(0)
	finalisationBlockBN.SetBytes(b.FinalisesOn)
	newBatch := Batch{
		Index:                b.Index,
		StateRoot:            decodedStateRoot,
		Committer:            b.Committer,
		TxRoot:               decodedTxRoot,
		StakeAmount:          b.StakeAmount,
		FinalisesOn:          *finalisationBlockBN,
		SubmissionHash:       b.SubmissionHash,
		TransactionsIncluded: decodedTxs,
	}
	return newBatch, nil
}
