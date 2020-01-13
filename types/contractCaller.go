package types

import (
	big "math/big"
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
	TotalBatches() uint64
	FetchBatchWithIndex(uint64) (Batch, error)
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

	if contractCaller.RollupContract, err = rollup.NewRollup(ethCmn.HexToAddress(config.GlobalCfg.RollupAddress), contractCaller.ethClient); err != nil {
		return contractCaller, err
	}
	if contractCaller.RollupContractABI, err = abi.JSON(strings.NewReader(rollup.RollupABI)); err != nil {
		return contractCaller, err
	}
	return contractCaller, nil
}

// TotalBatches returns the total number of batches that have been submitted on chain
func (c *ContractCaller) TotalBatches() (uint64, error) {
	totalBatches, err := c.RollupContract.NumberOfBatches(nil)
	if err != nil {
		return 0, err
	}
	return totalBatches.Uint64(), nil
}

// FetchBatchWithIndex fetched a block with index
func (c *ContractCaller) FetchBatchWithIndex(index uint64) (Batch, error) {
	crudeBatch, err := c.RollupContract.Batches(nil, big.NewInt(0).SetUint64(index))
	if err != nil {
		return Batch{}, err
	}
	return NewBatch(crudeBatch.StateRoot, Address(crudeBatch.Committer), crudeBatch.TxRoot), nil
}

func (c *ContractCaller) FetchBalanceTreeRoot() (ByteArray, error) {
	root, err := c.RollupContract.GetBalanceTreeRoot(nil)
	if err != nil {
		return ByteArray{}, err
	}
	return root, nil
}

// ProcessTx calls the ProcessTx function on the contract to verify the tx
// returns the updated accounts and the new balance root
func (c *ContractCaller) ProcessTx(balanceTreeRoot ByteArray, tx Tx, fromMerkleProof, toMerkleProof MerkleProof) (newBalanceRoot ByteArray, from, to AccountLeaf, err error) {
	txReceipt, err := c.RollupContract.ProcessTxUpdate(nil, balanceTreeRoot, tx.ToABIVersion(), fromMerkleProof.ToABIVersion(), toMerkleProof.ToABIVersion()).CallOpts()
	if err != nil {
		return
	}
	return
}
