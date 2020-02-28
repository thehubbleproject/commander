package db

import (
	"fmt"

	"github.com/BOPR/config"
	"github.com/BOPR/types"
	merkle "github.com/cbergoon/merkletree"
)

func (db *DB) StoreMT(mt merkle.MerkleTree) error {
	// coll, session := db.GetAccountCollection()
	// defer session.Close()

	// if err := coll.Insert(mt); err != nil {
	// 	fmt.Println("Unable to insert", "error", err)
	// 	return err
	// }
	return nil
}

// GetAllAccounts fetches all accounts from the database
func (db *DB) GetAllAccounts() (accs []types.AccountLeaf, err error) {
	// TODO add limits here
	errs := db.Instance.Find(&accs).GetErrors()
	for _, err := range errs {
		if err != nil {
			return accs, GenericError("got error while fetch all accounts")
		}
	}
	return accs, nil
}

// GetAccount gets the account of the given path from the DB
func (db *DB) GetAccount(accID uint64) (types.AccountLeaf, error) {
	var account types.AccountLeaf
	if db.Instance.First(&account, accID).RecordNotFound() {
		return account, ErrRecordNotFound(fmt.Sprintf("unable to find record for accountID: %d", accID))
	}
	return account, nil
}

func (db *DB) InsertAccount(account types.AccountLeaf) error {
	db.Instance.Create(account)
	return nil
}

func (db *DB) InsertBulkAccounts(accounts []types.AccountLeaf) error {
	for _, account := range accounts {
		err := db.InsertAccount(account)
		if err != nil {
			return ErrUnableToCreateRecord(fmt.Sprintf("Unable to add account with ID:%v to DB", account.Path))
		}
	}
	return nil
}

func (db *DB) InsertGenAccounts(genAccs []config.GenAccountLeaf) error {
	var accLeafs []types.AccountLeaf
	for _, acc := range genAccs {
		newAccLeaf := types.NewAccountLeaf(acc.Path, acc.Balance, acc.TokenType, acc.Nonce)
		accLeafs = append(accLeafs, newAccLeaf)
	}
	return db.InsertBulkAccounts(accLeafs)
}

func (db *DB) GetAccountCount() (int, error) {
	var count int
	db.Instance.Table("account_leafs").Count(&count)
	return count, nil
}

// FetchSiblings retuns the siblings of an account leaf till root
// TODO make this more performannt by using bulk account fetch or using groutines to fetch in parerell
func FetchSiblings(accID uint64, db DB) (accs []types.AccountLeaf, err error) {
	// For eg: for account ID 1111 => 1110, 110X, 10XX
	var siblings []types.AccountLeaf

	return siblings, nil
}
