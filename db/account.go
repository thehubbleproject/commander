package db

import (
	"errors"
	"fmt"
	"math"

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
			return ErrUnableToCreateRecord(fmt.Sprintf("Unable to add account with ID:%v to DB. Error: %v", account.AccountID, err))
		}
	}
	return nil
}

func (db *DB) InsertGenAccounts(genAccs []config.GenUserAccount) error {
	var accLeafs []types.UserAccount
	for _, acc := range genAccs {
		newAccLeaf := types.NewUserAccount(acc.ID, acc.Balance, acc.TokenType, acc.Nonce, acc.Path, int(acc.Status), acc.PublicKey)
		accLeafs = append(accLeafs, *newAccLeaf)
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

func (db *DB) InitEmptyDepositTree() error {
	var depositTree types.DepositTree
	depositTree.Root = types.ZERO_VALUE_LEAF.String()
	return db.Instance.Create(&depositTree).Error
}

func (db *DB) OnDepositLeafMerge(left, right, newRoot types.ByteArray) (uint64, error) {
	// get last deposit from deposit tree
	var lastDeposit types.DepositTree
	err := db.Instance.First(&lastDeposit).Error
	if err != nil {
		return 0, err
	}

	// update the deposit tree stored
	var updatedDepositTreeInfo types.DepositTree
	updatedDepositTreeInfo.Height = lastDeposit.Height + 1
	updatedDepositTreeInfo.NumberOfDeposits = lastDeposit.NumberOfDeposits + 2
	updatedDepositTreeInfo.Root = newRoot.String()

	generatedRoot, err := types.GetParent(left, right)
	if err != nil {
		return 0, err
	}

	if generatedRoot.String() != newRoot.String() {
		return 0, errors.New("Unable to update deposit tree, deposit tree root doesnt match")
	}
	if err := db.Instance.Model(&lastDeposit).Update(&updatedDepositTreeInfo).Error; err != nil {
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

func (db *DB) GetPendingDeposits(numberOfAccs uint64) ([]types.UserAccount, error) {
	var accounts []types.UserAccount
	err := db.Instance.Limit(numberOfAccs).Where("status = ?", 0).Find(&accounts).Error
	if err != nil {
		return accounts, err
	}
	return accounts, nil
}

func (db *DB) FinaliseDeposits(accountsRoot types.ByteArray, pathToDepositSubTree uint64, newBalanceRoot types.ByteArray) error {
	db.Logger.Info("Finalising accounts", "accountRoot", accountsRoot, "NewBalanceRoot", newBalanceRoot, "pathToDepositSubTree", pathToDepositSubTree)

	// get params
	params, err := db.GetParams()
	if err != nil {
		return err
	}

	// number of new deposits = 2**MaxDepthOfDepositTree
	depositCount := uint64(math.Exp2(float64(params.MaxDepositSubTreeHeight)))

	// get all pending accounts
	pendingAccs, err := db.GetPendingDeposits(depositCount)
	if err != nil {
		return err
	}
	db.Logger.Debug("Fetched pending deposits", "count", len(pendingAccs), "data", pendingAccs)

	// update the empty leaves with new accounts
	err = db.AddPendingDeposits(pendingAccs)
	if err != nil {
		return err
	}

	// TODO ensure the accounts are inserted at pathToDepositSubTree

	//TODO  make sure all the accounts root match to accountsRoot

	return nil
}

func (db *DB) AddPendingDeposits(pendingAccs []types.UserAccount) error {
	var accounts []types.UserAccount
	// fetch 2**DepositSubTree inactive accounts ordered by path
	err := db.Instance.Limit(len(pendingAccs)).Order("path").Where("status = ?", 100).Find(&accounts).Error
	if err != nil {
		return err
	}
	// TODO add error for if no account found

	// update the accounts
	for i, acc := range accounts {
		// acc.Balance = pendingAccs[i].Balance
		// acc.TokenType = pendingAccs[i].TokenType
		// acc.AccountID = pendingAccs[i].AccountID
		// acc.PublicKey = pendingAccs[i].PublicKey
		err := db.Instance.Model(&acc).Updates(types.UserAccount{Balance: pendingAccs[i].Balance,
			TokenType: pendingAccs[i].TokenType,
			AccountID: pendingAccs[i].AccountID,
			PublicKey: pendingAccs[i].PublicKey,
			Status:    1,
		}).Error
		if err != nil {
			return err
		}
	}
	return nil
}

// create merkle proof for finalisation of deposits
// send transaction to etherum chain using contract caller
func (db *DB) sendDepositFinalisationTx() {

}
