package bazooka

import (
	"context"
	"errors"
	big "math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/rpc"

	"github.com/BOPR/common"
	"github.com/BOPR/config"

	"github.com/BOPR/contracts/logger"
	"github.com/BOPR/contracts/merkleTree"
	"github.com/BOPR/contracts/rollup"

	ethTypes "github.com/ethereum/go-ethereum/core/types"

	"github.com/BOPR/core"
	"github.com/ethereum/go-ethereum/accounts/abi"
	ethCmn "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/tendermint/tendermint/libs/log"
)

// IContractCaller is the common interface using which we will interact with the contracts
// and the ethereum chain
type IBazooka interface {
	FetchBatchInputData(txHash ethCmn.Hash) (txs [][]byte, err error)
}

// TODO use context to remove this completely
// Global Contract Caller Object
var LoadedBazooka Bazooka

// ContractCaller satisfies the IContractCaller interface and contains all the variables required to interact
// With the ethereum chain along with contract addresses and ABI's
type Bazooka struct {
	Logger log.Logger

	EthClient *ethclient.Client

	ContractABI map[string]abi.ABI

	// Rollup contract
	RollupContract *rollup.Rollup
	BalanceTree    *merkleTree.MerkleTree
	EventLogger    *logger.Logger
}

// NewContractCaller contract caller
// NOTE: Reads configration from the config.toml file
func NewPreLoadedBazooka() (bazooka Bazooka, err error) {
	// TODO remove
	err = config.SetOperatorKeys(config.GlobalCfg.OperatorKey)
	if err != nil {
		return
	}
	err = config.ParseAndInitGlobalConfig()
	if err != nil {
		return
	}

	if RPCClient, err := rpc.Dial(config.GlobalCfg.EthRPC); err != nil {
		return bazooka, err
	} else {
		bazooka.EthClient = ethclient.NewClient(RPCClient)
	}

	bazooka.ContractABI = make(map[string]abi.ABI)
	// initialise all variables for rollup contract
	rollupContractAddress := ethCmn.HexToAddress(config.GlobalCfg.RollupAddress)
	if bazooka.RollupContract, err = rollup.NewRollup(rollupContractAddress, bazooka.EthClient); err != nil {
		return bazooka, err
	}
	if bazooka.ContractABI[common.ROLLUP_CONTRACT_KEY], err = abi.JSON(strings.NewReader(rollup.RollupABI)); err != nil {
		return bazooka, err
	}

	// initialise all variables for merkle tree contract
	balanceTreeContractAddress := ethCmn.HexToAddress(config.GlobalCfg.BalanceTreeAddress)
	if bazooka.BalanceTree, err = merkleTree.NewMerkleTree(balanceTreeContractAddress, bazooka.EthClient); err != nil {
		return bazooka, err
	}
	if bazooka.ContractABI[common.BALANCE_TREE_KEY], err = abi.JSON(strings.NewReader(merkleTree.MerkleTreeABI)); err != nil {
		return bazooka, err
	}

	// initialise all variables for event logger contract
	loggerAddress := ethCmn.HexToAddress(config.GlobalCfg.LoggerAddress)
	if bazooka.EventLogger, err = logger.NewLogger(loggerAddress, bazooka.EthClient); err != nil {
		return bazooka, err
	}
	if bazooka.ContractABI[common.LOGGER_KEY], err = abi.JSON(strings.NewReader(logger.LoggerABI)); err != nil {
		return bazooka, err
	}

	bazooka.Logger = common.Logger.With("module", "bazooka")
	return bazooka, nil
}

// get main chain block header
func (b *Bazooka) GetMainChainBlock(blockNum *big.Int) (header *ethTypes.Header, err error) {
	latestBlock, err := b.EthClient.HeaderByNumber(context.Background(), blockNum)
	if err != nil {
		return
	}
	return latestBlock, nil
}

// TotalBatches returns the total number of batches that have been submitted on chain
func (b *Bazooka) TotalBatches() (uint64, error) {
	totalBatches, err := b.RollupContract.NumberOfBatches(nil)
	if err != nil {
		return 0, err
	}
	return totalBatches.Uint64(), nil
}

// FetchBatchWithIndex fetched a block with index
// func (c *ContractCaller) FetchBatchWithIndex(index uint64) (Batch, error) {
// 	crudeBatch, err := c.RollupContract.Batches(nil, big.NewInt(0).SetUint64(index))
// 	if err != nil {
// 		return Batch{}, err
// 	}
// 	batch := Batch{
// 		Index: ,
// 		StateRoot: ,
// 		Committer: ,
// 		TxRoot: ,
// 		StakeCommitted: ,
// 		FinalisesOn: ,

// 	}
// 	return NewBatch(crudeBatch.StateRoot, Address(crudeBatch.Committer), crudeBatch.TxRoot), nil
// }

func (b *Bazooka) FetchBalanceTreeRoot() (core.ByteArray, error) {
	root, err := b.RollupContract.GetBalanceTreeRoot(nil)
	if err != nil {
		return core.ByteArray{}, err
	}
	return root, nil
}

func (b *Bazooka) FetchBatchInputData(txHash ethCmn.Hash) (txs [][]byte, err error) {
	tx, isPending, err := b.EthClient.TransactionByHash(context.Background(), txHash)
	if err != nil {
		b.Logger.Error("Cannot fetch transaction from hash", "Error", err)
		return
	}

	if isPending {
		err := errors.New("Transaction is pending")
		b.Logger.Error("Transaction is still pending, cannot process", "Error", err)
		return txs, err
	}

	payload := tx.Data()
	decodedPayload := payload[4:]

	inputDataMap := make(map[string]interface{})
	method := b.ContractABI[common.ROLLUP_CONTRACT_KEY].Methods["submitBatch"]
	err = method.Inputs.UnpackIntoMap(inputDataMap, decodedPayload)
	if err != nil {
		b.Logger.Error("Error unpacking payload", "Error", err)
		return
	}
	b.Logger.Debug("Created input data map", "InputData", inputDataMap)

	return GetTxsFromInput(inputDataMap), nil
}

// ABIEncodeTx encodes a transaction account to abi.encode parameter in solidity
// func (c *ContractCaller) ABIEncodeTx(tx *core.Transaction) []byte {

// }

func GenerateAuthObj(client *ethclient.Client, callMsg ethereum.CallMsg) (auth *bind.TransactOpts, err error) {
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
func (b *Bazooka) ProcessTx(balanceTreeRoot core.ByteArray, tx core.Tx, fromMerkleProof, toMerkleProof core.MerkleProof) (newBalanceRoot core.ByteArray, from, to core.UserAccount, err error) {
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

func (b *Bazooka) FireDepositFinalisation(TBreplaced core.UserAccount, siblings []core.UserAccount, subTreeDepth uint64) {

}
