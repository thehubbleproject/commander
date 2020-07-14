package core

import (
	"fmt"

	"encoding/hex"

	"github.com/BOPR/config"
	"github.com/BOPR/contracts/rollup"
	ethCmn "github.com/ethereum/go-ethereum/common"

	"github.com/ethereum/go-ethereum/rlp"
	"golang.org/x/crypto/sha3"
)

// Tx represets the transaction on hubble
type Tx struct {
	DBModel
	To        uint64 `json:"to"`
	From      uint64 `json:"from"`
	Data      []byte `json:"data"`
	Signature string `json:"sig" gorm:"not null"`
	TxHash    string `json:"hash" gorm:"not null"`
	Status    uint64 `json:"status"`
	Type      uint64 `json:"type"`
}

// NewTx creates a new transaction
func NewTx(from, to, txType uint64, message []byte, sig string) Tx {
	return Tx{
		From:      from,
		To:        to,
		Data:      message,
		Signature: sig,
		Type:      txType,
	}
}

// NewPendingTx creates a new transaction
func NewPendingTx(from, to, txType uint64, sig string, message []byte) Tx {
	tx := Tx{
		To:        to,
		From:      from,
		Data:      message,
		Signature: sig,
		Status:    TX_STATUS_PENDING,
		Type:      txType,
	}
	tx.AssignHash()
	return tx
}

// GetSignBytes returns the transaction data that has to be signed
func (tx Tx) GetSignBytes() (signBytes []byte) {
	return tx.Data
}

// AssignHash creates a tx hash and add it to the tx
func (t *Tx) AssignHash() {
	if t.TxHash != "" {
		return
	}
	hash := rlpHash(t)
	t.TxHash = hash.String()
}

func (tx *Tx) Apply(updatedFrom, updatedTo []byte) error {
	// get fresh copy of database
	db, err := NewDB()
	if err != nil {
		return err
	}

	// begin a transaction
	mysqlTx := db.Instance.Begin()
	defer func() {
		if r := recover(); r != nil {
			mysqlTx.Rollback()
		}
	}()

	// post this perform all ops on transaction
	db.Instance = mysqlTx

	// apply transaction on from account
	fromAcc, err := db.GetAccountByID(tx.From)
	if err != nil {
		mysqlTx.Rollback()
		return err
	}

	fromAcc.Data = updatedFrom

	err = db.UpdateAccount(fromAcc)
	if err != nil {
		mysqlTx.Rollback()
		return err
	}

	// apply transaction on to account
	toAcc, err := db.GetAccountByID(tx.To)
	if err != nil {
		mysqlTx.Rollback()
		return err
	}

	toAcc.Data = updatedTo

	err = db.UpdateAccount(toAcc)
	if err != nil {
		mysqlTx.Rollback()
		return err
	}

	tx.UpdateStatus(TX_STATUS_PROCESSED)

	// Or commit the transaction
	mysqlTx.Commit()
	return nil
}

func (t *Tx) String() string {
	return fmt.Sprintf("To: %v From: %v Status:%v Hash: %v Data: %v", t.To, t.From, t.Status, t.TxHash, hex.EncodeToString(t.Data))
}

// ToABIVersion converts a standard tx to the the DataTypesTransaction struct on the contract
func (t *Tx) ToABIVersion() (rollupTx rollup.TypesTransaction, err error) {
	decodedSignature, _ := hex.DecodeString(t.Signature)
	from, to, token, nonce, txType, amount, err := LoadedBazooka.DecodeTx(t.Data)
	if err != nil {
		return
	}
	rollupTx = rollup.TypesTransaction{
		FromIndex: from,
		ToIndex:   to,
		TokenType: token,
		Amount:    amount,
		Nonce:     nonce,
		TxType:    txType,
		Signature: decodedSignature,
	}
	return rollupTx, nil
}

func (tx *Tx) Compress() ([]byte, error) {
	return LoadedBazooka.CompressTx(*tx)
}

// Insert tx into the DB
func (db *DB) InsertTx(t *Tx) error {
	return db.Instance.Create(t).Error
}

func (db *DB) PopTxs() (txs []Tx, err error) {
	tx := db.Instance.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return txs, err
	}
	var pendingTxs []Tx

	// select N number of transactions which are pending in mempool and
	if err := tx.Limit(config.GlobalCfg.TxsPerBatch).Where(&Tx{Status: TX_STATUS_PENDING, Type: TX_TRANSFER_TYPE}).Find(&pendingTxs).Error; err != nil {
		db.Logger.Error("error while fetching pending transactions", err)
		return txs, err
	}

	db.Logger.Info("found txs", "pendingTxs", pendingTxs)

	var ids []string
	for _, tx := range pendingTxs {
		ids = append(ids, tx.ID)
	}

	// update the transactions from pending to processing
	errs := tx.Table("txes").Where("id IN (?)", ids).Updates(map[string]interface{}{"status": TX_STATUS_PROCESSING}).GetErrors()
	if err != nil {
		db.Logger.Error("errors while processing transactions", errs)
		return
	}

	return pendingTxs, tx.Commit().Error
}

func (db *DB) GetTx() (tx []Tx, err error) {
	err = db.Instance.First(&tx).Error
	if err != nil {
		return tx, err
	}
	return
}

func (tx *Tx) UpdateStatus(status uint64) error {
	return DBInstance.Instance.Model(&tx).Update("status", status).Error
}

func rlpHash(x interface{}) (h ethCmn.Hash) {
	hw := sha3.NewLegacyKeccak256()
	rlp.Encode(hw, x)
	hw.Sum(h[:0])
	return h
}
