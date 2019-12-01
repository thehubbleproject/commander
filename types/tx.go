package types

import (
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
