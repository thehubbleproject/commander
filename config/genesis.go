package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/BOPR/common"
)

type Genesis struct {
	StartEthBlock           string          `json:"startEthBlock"`
	MaxTreeDepth            string          `json:"maxTreeDepth"`
	MaxDepositSubTreeHeight string          `json:"maxDepositSubTreeHeight"`
	GenesisAccounts         GenesisAccounts `json:"genesisAccounts"`
}

//  GenUserAccount exists to allow remove circular dependency with types
// and to allow storing more data about the account than the data in UserAccount
type GenUserAccount struct {
	ID        uint64 `json:"ID"`
	Path      uint64
	Balance   uint64
	TokenType uint64
	Nonce     uint64
}

func NewGenUserAccount(path, balance, tokenType, nonce uint64) GenUserAccount {
	return GenUserAccount{
		Path:      path,
		Balance:   balance,
		TokenType: tokenType,
		Nonce:     nonce,
	}
}

type GenesisAccounts struct {
	Accounts []GenUserAccount `json:"gen_accounts"`
}

func NewGenesisAccounts(accounts []GenUserAccount) GenesisAccounts {
	return GenesisAccounts{Accounts: accounts}
}

func DefaultGenesisAccounts() GenesisAccounts {
	var accounts []GenUserAccount
	acc := NewGenUserAccount(common.DEAFULT_PATH, common.DEFAULT_BALANCE, common.DEFAULT_TOKEN_TYPE, common.DEFAULT_TOKEN_TYPE)
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
