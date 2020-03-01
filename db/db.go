package db

import (
	"fmt"

	"github.com/BOPR/config"
	"github.com/BOPR/types"
	"github.com/globalsign/mgo"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
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
	Instance *gorm.DB
}

func NewDB() (DB, error) {
	config.ParseAndInitGlobalConfig()
	fmt.Println("Connecting to DB", "type", config.GlobalCfg.DB, "URL", config.GlobalCfg.FormattedDBURL())
	db, err := gorm.Open(config.GlobalCfg.DB, config.GlobalCfg.FormattedDBURL())
	if err != nil {
		return DB{}, err
	}
	db.LogMode(config.GlobalCfg.DBLogMode)
	fmt.Println("database", db)
	return DB{Instance: db}, nil
}

func (db *DB) Close() {
	db.Instance.Close()
}
