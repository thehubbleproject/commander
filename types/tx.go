package types

import (
	"bytes"

	"github.com/BOPR/common"
	ethCommon "github.com/ethereum/go-ethereum/common"
)

// Tx represets the transaction on BOPRU
type Tx struct {
	To     string
	From   string
	Amount uint64
	Nonce  uint64
	// Fee       uint64
	TxType    uint64
	Signature string
}

// NewTx creates a new transaction
func NewTx(to ethCommon.Address, from ethCommon.Address, amount uint64, nonce uint64, sig string) Tx {
	if bytes.Equal(to.Bytes(), common.WithdrawAddress) {
		txType := 1
	}
	return Tx{
		To:     to,
		From:   from,
		Amount: amount,
		Nonce:  nonce,
		TxType: 1,
	}
}

// ValidateTx validates a transaction
func (t *Tx) ValidateTx() bool {
	// signature len verification

	// amount > 0 verification

	// signature to from account verification
}
