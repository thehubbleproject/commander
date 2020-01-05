package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

//  GenAccountLeaf Exists to allow remove circular dependency with types
// and to allow storing more data than account leaf
type GenAccountLeaf struct {
	Path      uint64
	Balance   uint64
	TokenType uint64
	Nonce     uint64
}

func NewGenAccountLeaf(path, balance, tokenType, nonce uint64) GenAccountLeaf {
	return GenAccountLeaf{
		Path:      path,
		Balance:   balance,
		TokenType: tokenType,
		Nonce:     nonce,
	}
}

type GenesisAccounts struct {
	Accounts []GenAccountLeaf `json:"gen_accounts"`
}

func NewGenesisAccounts(accounts []GenAccountLeaf) GenesisAccounts {
	return GenesisAccounts{Accounts: accounts}
}

func DefaultGenesisAccounts() GenesisAccounts {
	var accounts []GenAccountLeaf
	acc := NewGenAccountLeaf(1223123, 10, 1, 0)
	accounts = append(accounts, acc)
	return NewGenesisAccounts(accounts)
}

func ReadGenesisFile() (GenesisAccounts, error) {
	var genAccounts GenesisAccounts

	genesisFile, err := os.Open("genesis.json")
	if err != nil {
		fmt.Println(err)
		return genAccounts, err
	}
	defer genesisFile.Close()

	genBytes, err := ioutil.ReadAll(genesisFile)
	if err != nil {
		fmt.Println(err)
		return genAccounts, err
	}
	err = json.Unmarshal(genBytes, &genAccounts)
	return genAccounts, err
}

func WriteGenesisFile(genesis GenesisAccounts) error {
	bz, err := json.Marshal(genesis)
	if err != nil {
		return err
	}
	return ioutil.WriteFile("genesis.json", bz, 0644)
}
