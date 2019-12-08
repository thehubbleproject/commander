package types

import (
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	"time"

	"github.com/BOPR/common"
	db "github.com/BOPR/db"
)

// Tx represets the transaction on BOPRU
type Tx struct {
	To     uint64 `bson:"to"`
	From   uint64 `bson:"from"`
	Amount uint64 `bson:"amount"`
	Nonce  uint64 `bson:"nonce"`
	// Fee       uint64
	TxType    uint64    `bson:"type"`
	Signature string    `bson:"sig"`
	Status    string    `bson:"status"`
	Timestamp time.Time `bson:"timestamp"`
}

// NewTx creates a new transaction
func NewTx(to uint64, from uint64, amount uint64, nonce uint64, sig string) Tx {
	if to == 0 {
		// its a withdraw tx
		// txType := 1
	}
	return Tx{
		To:        to,
		From:      from,
		Amount:    amount,
		Nonce:     nonce,
		TxType:    1,
		Signature: sig,
		Status:    "pending",
		Timestamp: time.Now(),
	}
}

// ValidateTx validates a transaction
func (t *Tx) ValidateTx() error {
	// signature len verification
	if len(t.Signature) != 65 {
		return errors.New("Signature invalid")
	}

	// signature to from account verification
	return nil
}

// Insert tx into the DB
func (t *Tx) Insert() error {
	session := db.MgoSession.Copy()
	defer session.Close()
	if err := session.GetCollection(common.DATABASE, common.TRANSACTION_COLLECTION).Insert(t); err != nil {
		fmt.Println("Unable to insert", "error", err)
		return err
	}
	return nil
}

func (t *Tx) String() string {
	return fmt.Sprintf("To: %v From: %v Amount: %v Nonce: %v Type:%v Status:%v", t.To, t.From, t.Amount, t.Nonce, t.Status)
}

// MinimalTx constructs minimal tx from normal tx
func (t *Tx) MinimalTx() (tx MinimalTx, err error) {
	sig, err := hex.DecodeString(t.Signature)
	if err != nil {
		return tx, err
	}
	var trimmedSig [64]byte
	copy(trimmedSig[:], sig)
	return NewMinimalTx(
		uint32(t.To),
		uint32(t.From),
		uint32(t.Amount),
		uint32(t.Nonce),
		common.UintTo2Byte(uint32(t.TxType)),
		trimmedSig,
	), nil
}

// MinimalTx is the transaction that needs to be pushed on-chain
type MinimalTx struct {
	To     uint32 `bson:"to"`
	From   uint32 `bson:"from"`
	Amount uint32 `bson:"amount"`
	Nonce  uint32 `bson:"nonce"`
	// Fee       uint64
	TxType    [2]byte  `bson:"type"`
	Signature [64]byte `bson:"sig"`
}

// NewMinimalTx generates new minimal tx
func NewMinimalTx(to uint32, from uint32, amount uint32, nonce uint32, txType [2]byte, sig [64]byte) MinimalTx {
	return MinimalTx{
		To:        to,
		From:      from,
		Amount:    amount,
		Nonce:     nonce,
		TxType:    txType,
		Signature: sig,
	}
}

// Serialise serialises the minmialTx to be pushed on chain
func (mtx *MinimalTx) Serialise() []byte {
	var dataSlices = [][]byte{
		common.UintToByte(mtx.To)[:],
		common.UintToByte(mtx.From)[:],
		common.UintToByte(mtx.Amount)[:],
		common.UintToByte(mtx.Nonce)[:],
		mtx.TxType[:],
		mtx.Signature[:],
	}
	return common.AppendSlices(dataSlices)
}

func DeserialiseMinimalTx(tx []byte) MinimalTx {
	to := binary.LittleEndian.Uint32(tx[:4])
	from := binary.LittleEndian.Uint32(tx[4:8])
	amount := binary.LittleEndian.Uint32(tx[8:12])
	nonce := binary.LittleEndian.Uint32(tx[12:16])

	txType := tx[16:18]
	var txType2Byte [2]byte
	copy(txType2Byte[:], txType[:])

	sig := tx[18:]
	var sig64Byte [64]byte
	copy(sig64Byte[:], sig[:])

	// var txType4Byte [4]byte
	// copy(txType2Byte[:], txType4Byte[:])
	// txType := binary.LittleEndian.Uint32(txType4Byte[:])

	return MinimalTx{To: to, From: from, Amount: amount, Nonce: nonce, TxType: txType2Byte, Signature: sig64Byte}
}

func (t *MinimalTx) String() string {
	return fmt.Sprintf("To: %v From: %v Amount: %v Nonce: %v Type:%v Sig:%v", t.To, t.From, t.Amount, t.Nonce, t.TxType, t.Signature)
}
