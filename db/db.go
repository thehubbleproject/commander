package db

import (
	"github.com/BOPR/config"
	"github.com/BOPR/types"
	"github.com/globalsign/mgo"
)

type IDB interface {
	// Account related DB functions
	// FetchSiblings(accID uint64) (accs []types.AccountLeaf, err error)
	GetAllAccounts() (accs []types.AccountLeaf, err error)
	GetAccount(accID uint64) (types.AccountLeaf, error)
	InsertBulkAccounts(accounts []types.AccountLeaf) error
	InsertGenAccounts(genAccs []config.GenAccountLeaf) error
	GetAccountCount() (int, error)

	// Tx related functions
	InsertTx(t *types.Tx) error
	PopTxs() (txs []types.Tx, err error)

	// Batch related functions
	InsertBatchInfo(root types.ByteArray, index uint64) error
	GetAllBatches() (batches []types.Batch, err error)
	GetLatestBatch() (types.Batch, error)
	GetBatchCount() (int, error)

	// common functions
	GetBatchCollection() *mgo.Collection
	GetTransactionCollection() *mgo.Collection
	GetAccountCollection() *mgo.Collection
}

// global DB instance created while doing init
var DBInstance DB

type DB struct {
	MgoSession Session
}

func NewDB() (DB, error) {
	session, err := NewSession("url")
	if err != nil {
		return DB{}, err
	}
	return DB{MgoSession: *session}, nil
}
