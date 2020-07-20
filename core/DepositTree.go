package core

import (
	"bytes"
	"errors"
)

func (db *DB) GetDepositNodeAndSiblings() (NodeToBeReplaced UserAccount, siblings []UserAccount, err error) {
	// get params
	params, err := db.GetParams()
	if err != nil {
		return
	}

	// get the deposit node
	expectedHash := defaultHashes[params.MaxDepositSubTreeHeight]

	// getNode with the expectedHash
	NodeToBeReplaced, err = db.GetDepositSubTreeRoot(expectedHash.String(), params.MaxDepth-params.MaxDepositSubTreeHeight)
	if err != nil {
		return
	}

	// get siblings for the path to node
	siblings, err = db.GetSiblings(NodeToBeReplaced.Path)
	if err != nil {
		return
	}

	return
}

func (db *DB) FinaliseDepositsAndAddBatch(depositRoot ByteArray, pathToDepositSubTree uint64) (string, error) {
	var root string
	db.Logger.Info("Finalising accounts", "depositRoot", depositRoot, "pathToDepositSubTree", pathToDepositSubTree)

	// update the empty leaves with new accounts
	err := db.FinaliseDeposits(pathToDepositSubTree, depositRoot)
	if err != nil {
		return root, err
	}

	rootAccount, err := db.GetRoot()
	if err != nil {
		return root, err
	}

	return rootAccount.Hash, nil
}

func (db *DB) FinaliseDeposits(pathToDepositSubTree uint64, depositRoot ByteArray) error {
	params, err := db.GetParams()
	if err != nil {
		return err
	}

	// find out the accounts that are finalised
	accounts, err := db.GetPendingAccByDepositRoot(depositRoot)
	if err != nil {
		return err
	}

	// find out where the insertion was made
	height := params.MaxDepth - 1
	getTerminalNodesOf, err := SolidityPathToNodePath(pathToDepositSubTree, height)
	if err != nil {
		return err
	}

	// TODO add error for if no account found
	terminalNodes, err := db.GetAllTerminalNodes(getTerminalNodesOf)
	if err != nil {
		return err
	}

	for i, acc := range accounts {
		acc.Status = STATUS_ACTIVE
		acc.UpdatePath(terminalNodes[i])
		acc.CreateAccountHash()
		err := db.UpdateAccount(acc)
		if err != nil {
			return err
		}

		// delete pending account
		err = db.DeletePendingAccount(acc.AccountID)
		if err != nil {
			return err
		}
	}

	return nil
}

func (db *DB) GetPendingDeposits(numberOfAccs uint64) ([]UserAccount, error) {
	var accounts []UserAccount
	err := db.Instance.Limit(numberOfAccs).Where("status = ?", 0).Find(&accounts).Error
	if err != nil {
		return accounts, err
	}
	return accounts, nil
}

func (db *DB) GetAllTerminalNodes(pathToDepositSubTree string) (terminalNodes []string, err error) {
	buf := bytes.Buffer{}
	buf.WriteString(pathToDepositSubTree)
	buf.WriteString("%")
	var accounts []UserAccount

	// LIKE query with search for terminal nodes to DB
	if err = db.Instance.Where("path LIKE ? AND type = ?", buf.String(), 1).Find(&accounts).Error; err != nil {
		return
	}

	// get all accounts while making sure they are empty and append to paths array
	for _, account := range accounts {
		if account.Hash != ZERO_VALUE_LEAF.String() {
			return terminalNodes, errors.New("Account not zero, aborting operation")
		}
		terminalNodes = append(terminalNodes, account.Path)
	}
	return
}
