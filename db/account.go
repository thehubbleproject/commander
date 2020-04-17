package db

import (
	"fmt"

	"github.com/BOPR/config"
	"github.com/BOPR/types"
	merkle "github.com/cbergoon/merkletree"
)

func (db *DB) StoreMT(mt merkle.MerkleTree) error {

	return nil
}

// GetAllAccounts fetches all accounts from the database
func (db *DB) GetAllAccounts() (accs []types.UserAccount, err error) {
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
func (db *DB) GetAccount(ID uint64) (types.UserAccount, error) {
	var account types.UserAccount
	if db.Instance.First(&account, ID).RecordNotFound() {
		return account, ErrRecordNotFound(fmt.Sprintf("unable to find record for accountID: %d", ID))
	}
	return account, nil
}

func (db *DB) InsertAccount(account types.UserAccount) error {
	return db.Instance.Create(account).Error
}

func (db *DB) InsertBulkAccounts(accounts []types.UserAccount) error {
	for _, account := range accounts {
		err := db.InsertAccount(account)
		if err != nil {
			return ErrUnableToCreateRecord(fmt.Sprintf("Unable to add account with ID:%v to DB", account.ID))
		}
	}
	return nil
}

func (db *DB) InsertGenAccounts(genAccs []config.GenUserAccount) error {
	var accLeafs []types.UserAccount
	for _, acc := range genAccs {
		newAccLeaf := types.NewUserAccount(acc.Path, acc.Balance, acc.TokenType, acc.Nonce)
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
func FetchSiblings(accID uint64, db DB) (accs []types.UserAccount, err error) {
	// For eg: for account ID 1111 => 1110, 110X, 10XX
	var siblings []types.UserAccount

	return siblings, nil
}

func (db *DB) UpdateDepositTreeInfo(dt types.DepositTree) error {
	return db.Instance.Create(dt).Error
}

func (db *DB) OnDepositLeafMerge(left, right, newRoot types.ByteArray) (uint64, error) {
	var lastDeposit types.DepositTree
	err := db.Instance.First(&lastDeposit).Error
	if err != nil {
		return 0, err
	}
	var updatedDepositTreeInfo types.DepositTree
	updatedDepositTreeInfo.Height = lastDeposit.Height + 1
	updatedDepositTreeInfo.NumberOfDeposits = lastDeposit.NumberOfDeposits + 2
	updatedDepositTreeInfo.Root = newRoot
	// TODO utils.getParent(left,right)==newRoot, else error
	if err := db.UpdateDepositTreeInfo(updatedDepositTreeInfo); err != nil {
		return 0, err
	}

	return updatedDepositTreeInfo.Height, nil
}

func (db *DB) GetDepositTreeInfo() (dt types.DepositTree, err error) {
	err = db.Instance.First(&dt).Error
	if err != nil {
		return
	}
	return
}

func (db *DB) FinaliseDeposits(accountsRoot types.ByteArray, pathToDepositSubTree uint64, newBalanceRoot types.ByteArray) {
	// get all pending accounts

	// update paths

	// make sure all the accounts root match to accountsRoot
}

// create merkle proof for finalisation of deposits
// send transaction to etherum chain using contract caller
func (db *DB) sendDepositFinalisationTx() {

}
