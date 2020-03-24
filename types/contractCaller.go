package types

import (
	"context"
	"errors"
	"fmt"
	big "math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/rpc"

	"github.com/BOPR/common"
	"github.com/BOPR/config"

	"github.com/BOPR/contracts/merkleTree"
	"github.com/BOPR/contracts/rollup"

	ethTypes "github.com/ethereum/go-ethereum/core/types"

	"github.com/ethereum/go-ethereum/accounts/abi"
	ethCmn "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/tendermint/tendermint/libs/log"
)

// IContractCaller is the common interface using which we will interact with the contracts
// and the ethereum chain
type IContractCaller interface {
}

// TODO use context to remove this completely
// Global Contract Caller Object
var ContractCallerObj ContractCaller

// ContractCaller satisfies the IContractCaller interface and contains all the variables required to interact
// With the ethereum chain along with contract addresses and ABI's
type ContractCaller struct {
	Logger log.Logger

	EthClient *ethclient.Client

	// Rollup contract
	RollupContract        *rollup.Rollup
	RollupContractABI     abi.ABI
	RollupContractAddress ethCmn.Address

	// Merkle Tree libs
	MerkleTreeContract   *merkleTree.MerkleTree
	MerkleTreeABI        abi.ABI
	MerkleTreeLibAddress ethCmn.Address
}

// NewContractCaller contract caller
// NOTE: Reads configration from the config.toml file
func NewContractCaller() (contractCaller ContractCaller, err error) {
	config.ParseAndInitGlobalConfig()
	if RPCClient, err := rpc.Dial(config.GlobalCfg.EthRPC); err != nil {
		return contractCaller, err
	} else {
		contractCaller.EthClient = ethclient.NewClient(RPCClient)
	}

	// initialise all variables for rollup contract
	rollupContractAddress := ethCmn.HexToAddress(config.GlobalCfg.RollupAddress)
	if contractCaller.RollupContract, err = rollup.NewRollup(rollupContractAddress, contractCaller.EthClient); err != nil {
		return contractCaller, err
	}
	if contractCaller.RollupContractABI, err = abi.JSON(strings.NewReader(rollup.RollupABI)); err != nil {
		return contractCaller, err
	}
	contractCaller.RollupContractAddress = rollupContractAddress

	// initialise all variables for merkle tree contract
	merkleTreeContractAddress := ethCmn.HexToAddress(config.GlobalCfg.MerkleTreeLibAddress)
	if contractCaller.MerkleTreeContract, err = merkleTree.NewMerkleTree(merkleTreeContractAddress, contractCaller.EthClient); err != nil {
		return contractCaller, err
	}
	if contractCaller.MerkleTreeABI, err = abi.JSON(strings.NewReader(merkleTree.MerkleTreeABI)); err != nil {
		return contractCaller, err
	}
	contractCaller.MerkleTreeLibAddress = merkleTreeContractAddress

	contractCaller.Logger = common.Logger.With("contractCaller")
	return contractCaller, nil
}

// get main chain block header
func (c *ContractCaller) GetMainChainBlock(blockNum *big.Int) (header *ethTypes.Header, err error) {
	latestBlock, err := c.EthClient.HeaderByNumber(context.Background(), blockNum)
	if err != nil {
		return
	}
	return latestBlock, nil
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

func (c *ContractCaller) FetchBatchInputData(txHash ethCmn.Hash) (txs [][]byte, err error) {
	tx, isPending, err := c.EthClient.TransactionByHash(context.Background(), txHash)
	if err != nil {
		c.Logger.Error("Cannot fetch transaction from hash", "Error", err)
		return
	}

	if isPending {
		err := errors.New("Transaction is pending")
		c.Logger.Error("Transaction is still pending, cannot process", "Error", err)
		return txs, err
	}

	payload := tx.Data()
	decodedPayload := payload[4:]

	inputDataMap := make(map[string]interface{})

	method := c.RollupContractABI.Methods["SubmitBatch"]

	err = method.Inputs.UnpackIntoMap(inputDataMap, decodedPayload)
	if err != nil {
		fmt.Println("Error unpacking payload", "Error", err)
		return
	}

	return GetTxsFromInput(inputDataMap), nil
}

// ABIEncodeTx encodes a transaction account to abi.encode parameter in solidity
// func (c *ContractCaller) ABIEncodeTx(tx *types.Transaction) []byte {

// }

func GenerateAuthObj(client *ethclient.Client, callMsg ethereum.CallMsg) (auth *bind.TransactOpts, err error) {
	// TODO remove
	config.ParseAndInitGlobalConfig()
	config.SetOperatorKeys(config.GlobalCfg.OperatorKey)

	// from address
	fromAddress := config.OperatorAddress()

	// fetch gas price
	gasprice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return
	}
	// fetch nonce
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return
	}

	// fetch gas limit
	callMsg.From = fromAddress
	gasLimit, err := client.EstimateGas(context.Background(), callMsg)

	// create auth
	auth = bind.NewKeyedTransactor(config.OperatorKey)
	auth.GasPrice = gasprice
	auth.Nonce = big.NewInt(int64(nonce))
	auth.GasLimit = uint64(gasLimit) // uint64(gasLimit)

	return
}

// ProcessTx calls the ProcessTx function on the contract to verify the tx
// returns the updated accounts and the new balance root
func (c *ContractCaller) ProcessTx(balanceTreeRoot ByteArray, tx Tx, fromMerkleProof, toMerkleProof MerkleProof) (newBalanceRoot ByteArray, from, to UserAccount, err error) {
	// txReceipt, err := c.RollupContract.ProcessTxUpdate(nil, balanceTreeRoot,
	// 	tx.ToABIVersion(fromMerkleProof.Account.ToABIAccount(), toMerkleProof.Account.ToABIAccount()),
	// 	fromMerkleProof.ToABIVersion(),
	// 	toMerkleProof.ToABIVersion())
	// if err != nil {
	// 	return
	// }

	// data, err := c.RollupContractABI.Pack("processTxUpdate", balanceTreeRoot,
	// 	tx.ToABIVersion(fromMerkleProof.Account.ToABIAccount(), toMerkleProof.Account.ToABIAccount()),
	// 	fromMerkleProof.ToABIVersion(),
	// 	toMerkleProof.ToABIVersion())
	// if err != nil {
	// 	fmt.Println("Unable to pack tx for submitHeaderBlock", "error", err)
	// 	return
	// }
	// callMsg := ethereum.CallMsg{
	// 	To:   &c.RollupContractAddress,
	// 	Data: data,
	// }
	// data, err = c.EthClient.CallContract(context.Background(), callMsg, nil)
	// if err != nil {
	// 	fmt.Println("Unable to pack tx for submitHeaderBlock", "error", err)
	// 	return
	// }
	// fmt.Println("data", data)

	return
}

func GetTxsFromInput(input map[string]interface{}) (txs [][]byte) {
	return input["_txs"].([][]byte)
}
