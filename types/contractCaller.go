package types

import (
	"strings"

	"github.com/ethereum/go-ethereum/rpc"

	"github.com/BOPR/config"
	"github.com/BOPR/contracts/rollup"
	"github.com/ethereum/go-ethereum/accounts/abi"
	ethCmn "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var ContractCallerObj ContractCaller

// IContractCaller represents contract caller
type IContractCaller interface {
	SubmitBatch(txs []Tx)
}

// ContractCaller contract caller
type ContractCaller struct {
	ethClient *ethclient.Client

	RollupContract *rollup.Rollup

	RollupContractABI abi.ABI
}

// NewContractCaller contract caller
func NewContractCaller() (contractCaller ContractCaller, err error) {
	if RPCClient, err := rpc.Dial(config.GlobalCfg.EthRPC); err != nil {
		return contractCaller, err
	} else {
		contractCaller.ethClient = ethclient.NewClient(RPCClient)
	}

	if contractCaller.RollupContract, err = rollup.NewRollup(ethCmn.HexToAddress("0x51579Fe655a6F14475Cf233503c9FD9a24F3d40c"), contractCaller.ethClient); err != nil {
		return contractCaller, err
	}
	if contractCaller.RollupContractABI, err = abi.JSON(strings.NewReader(rollup.RollupABI)); err != nil {
		return contractCaller, err
	}
	return contractCaller, nil
}
