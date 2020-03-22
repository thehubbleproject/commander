package types

import (
	"errors"
	"fmt"

	"github.com/BOPR/contracts/rollup"
	"github.com/jinzhu/gorm"
)

// Tx represets the transaction on BOPRU
type Tx struct {
	gorm.Model

	To        uint64 `json:"to"`
	From      uint64 `json:"from"`
	Amount    uint64 `json:"amount"`
	Nonce     uint64 `json:"nonce"`
	TokenID   uint64 `json:"tokenID"`
	Signature string `json:"sig"`

	// 100 Pending
	// 200 Processing
	// 300 Processed
	Status uint `json:"status"`
}

// NewTx creates a new transaction
func NewTx(to uint64, from uint64, amount uint64, nonce uint64, sig string) Tx {
	return Tx{
		To:        to,
		From:      from,
		Amount:    amount,
		Nonce:     nonce,
		TokenID:   1,
		Signature: sig,
		Status:    100,
	}
}

// ValidateTx validates a transaction
// NOTE: This is a stateless op, should be run before adding txs to mempool
func (t *Tx) ValidateBasic() error {
	// signature len verification
	if len(t.Signature) != 65 {
		return errors.New("Signature invalid")
	}

	// check status is within the permissible status codes
	if t.Status < 100 {
		return errors.New("Invalid status code for the transaction found")
	}

	// check amount is greater than 0
	if t.Amount <= 0 {
		return errors.New("Invalid amount. Cannot be less than 0")
	}

	// Check nonce is greater than 0
	if t.Nonce < 0 {
		return errors.New("Invalid nonce for the transaction found. Cannot be less than 0")
	}

	// check the ID's are greater than 0
	if t.From == 0 {
		return errors.New("Invalid from address found. From cannot be 0")
	}

	// check the ID's are greater than 0
	if t.From < 0 || t.To < 0 {
		return errors.New("Invalid To or From found. Cannot be less than 0")
	}

	return nil
}

func (t *Tx) String() string {
	return fmt.Sprintf("To: %v From: %v Amount: %v Nonce: %v Type:%v Status:%v", t.To, t.From, t.Amount, t.Nonce, t.TxType, t.Status)
}

// MinimalTx constructs minimal tx from normal tx
// func (t *Tx) MinimalTx() (tx MinimalTx, err error) {
// 	sig, err := hex.DecodeString(t.Signature)
// 	if err != nil {
// 		return tx, err
// 	}
// 	var trimmedSig [64]byte
// 	copy(trimmedSig[:], sig)
// 	return NewMinimalTx(
// 		uint32(t.To),
// 		uint32(t.From),
// 		uint32(t.Amount),
// 		uint32(t.Nonce),
// 		common.UintTo2Byte(uint32(t.TxType)),
// 		trimmedSig,
// 	), nil
// }

// ToABIVersion converts a standard tx to the the DataTypesTransaction struct on the contract
func (t *Tx) ToABIVersion(from, to rollup.DataTypesUserAccount) rollup.DataTypesTransaction {
	return rollup.DataTypesTransaction{
		To:        to,
		From:      from,
		Amount:    uint32(t.Amount),
		TokenType: from.TokenType,
		Signature: []byte(t.Signature),
	}
}

// // MinimalTx is the transaction that needs to be pushed on-chain
// type MinimalTx struct {
// 	To     uint32 `bson:"to"`
// 	From   uint32 `bson:"from"`
// 	Amount uint32 `bson:"amount"`
// 	Nonce  uint32 `bson:"nonce"`
// 	// Fee       uint64
// 	TxType    [2]byte  `bson:"type"`
// 	Signature [64]byte `bson:"sig"`
// }

// // NewMinimalTx generates new minimal tx
// func NewMinimalTx(to uint32, from uint32, amount uint32, nonce uint32, txType [2]byte, sig [64]byte) MinimalTx {
// 	return MinimalTx{
// 		To:        to,
// 		From:      from,
// 		Amount:    amount,
// 		Nonce:     nonce,
// 		TxType:    txType,
// 		Signature: sig,
// 	}
// }

// // Serialise serialises the minmialTx to be pushed on chain
// func (mtx *MinimalTx) Serialise() []byte {
// 	var dataSlices = [][]byte{
// 		common.UintToByte(mtx.To)[:],
// 		common.UintToByte(mtx.From)[:],
// 		common.UintToByte(mtx.Amount)[:],
// 		common.UintToByte(mtx.Nonce)[:],
// 		mtx.TxType[:],
// 		mtx.Signature[:],
// 	}
// 	return common.AppendSlices(dataSlices)
// }

// func DeserialiseMinimalTx(tx []byte) MinimalTx {
// 	to := binary.LittleEndian.Uint32(tx[:4])
// 	from := binary.LittleEndian.Uint32(tx[4:8])
// 	amount := binary.LittleEndian.Uint32(tx[8:12])
// 	nonce := binary.LittleEndian.Uint32(tx[12:16])

// 	txType := tx[16:18]
// 	var txType2Byte [2]byte
// 	copy(txType2Byte[:], txType[:])

// 	sig := tx[18:]
// 	var sig64Byte [64]byte
// 	copy(sig64Byte[:], sig[:])

// 	// var txType4Byte [4]byte
// 	// copy(txType2Byte[:], txType4Byte[:])
// 	// txType := binary.LittleEndian.Uint32(txType4Byte[:])

// 	return MinimalTx{To: to, From: from, Amount: amount, Nonce: nonce, TxType: txType2Byte, Signature: sig64Byte}
// }

// func (t *MinimalTx) String() string {
// 	return fmt.Sprintf("To: %v From: %v Amount: %v Nonce: %v Type:%v Sig:%v", t.To, t.From, t.Amount, t.Nonce, t.TxType, t.Signature)
// }
