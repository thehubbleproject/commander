package config

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/BOPR/common"
)

type Genesis struct {
	StartEthBlock           uint64          `json:"startEthBlock"`
	MaxTreeDepth            uint64          `json:"maxTreeDepth"`
	MaxDepositSubTreeHeight uint64          `json:"maxDepositSubTreeHeight"`
	StakeAmount             uint64          `json:"stakeAmount"`
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
	PublicKey string
}

func NewGenUserAccount(id, path, balance, tokenType, nonce uint64, publicKey string) GenUserAccount {
	return GenUserAccount{
		ID:        id,
		Path:      path,
		Balance:   balance,
		TokenType: tokenType,
		Nonce:     nonce,
		PublicKey: publicKey,
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

	// add coordinator account
	acc := NewGenUserAccount(common.COORDINATOR, common.COORDINATOR, common.COORDINATOR, common.COORDINATOR, common.COORDINATOR, common.COORDINATOR_PUBKEY)
	accounts = append(accounts, acc)
	return NewGenesisAccounts(accounts)
}

func DefaultGenesis() Genesis {
	return Genesis{
		StartEthBlock:           0,
		MaxTreeDepth:            common.DEFAULT_DEPTH,
		MaxDepositSubTreeHeight: common.DEFAULT_DEPTH,
		StakeAmount:             32,
		GenesisAccounts:         DefaultGenesisAccounts(),
	}
}

func ReadGenesisFile() (Genesis, error) {
	var genesis Genesis

	genesisFile, err := os.Open("genesis.json")
	if err != nil {
		return genesis, err
	}
	defer genesisFile.Close()

	genBytes, err := ioutil.ReadAll(genesisFile)
	if err != nil {
		return genesis, err
	}

	err = json.Unmarshal(genBytes, &genesis)
	return genesis, err
}

func WriteGenesisFile(genesis Genesis) error {
	bz, err := json.MarshalIndent(genesis, "", "    ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile("genesis.json", bz, 0644)
}
