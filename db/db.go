package db

import (
	"github.com/BOPR/config"
	"github.com/BOPR/types"
	"github.com/globalsign/mgo"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type IDB interface {
	// Account related DB functions
	// FetchSiblings(accID uint64) (accs []types.UserAccount, err error)
	GetAllAccounts() (accs []types.UserAccount, err error)
	GetAccount(accID uint64) (types.UserAccount, error)
	InsertBulkAccounts(accounts []types.UserAccount) error
	InsertGenAccounts(genAccs []config.GenUserAccount) error
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

// NewDB creates a new DB instance
// NOTE: it uses the configrations present in the config.toml file
// returns error if not able to open the DB
func NewDB() (DB, error) {
	if err := config.ParseAndInitGlobalConfig(); err != nil {
		return DB{}, err
	}
	db, err := gorm.Open(config.GlobalCfg.DB, config.GlobalCfg.FormattedDBURL())
	if err != nil {
		return DB{}, err
	}
	db.LogMode(config.GlobalCfg.DBLogMode)
	return DB{Instance: db}, nil
}

func (db *DB) Close() {
	db.Instance.Close()
}
