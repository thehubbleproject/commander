package listener

import (
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/BOPR/common"
	"github.com/BOPR/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	ethTypes "github.com/ethereum/go-ethereum/core/types"

	"github.com/BOPR/contracts/logger"
	"github.com/BOPR/contracts/rollup"
)

func (s *Syncer) processDepositQueued(eventName string, abiObject *abi.ABI, vLog *ethTypes.Log) {
	s.Logger.Info("New deposit found")

	// unpack event
	event := new(logger.LoggerDepositQueued)

	err := common.UnpackLog(abiObject, event, eventName, vLog)
	if err != nil {
		// TODO do something with this error
		fmt.Println("Unable to unpack log:", err)
		panic(err)
	}

	s.Logger.Info(
		"⬜ New event found",
		"event", eventName,
		"accountID", event.AccountID.String(),
		"Amount", event.Amount.String(),
		"TokenID", event.Token.String(),
		"AccountHash", event.AccountHash,
		"pubkey", event.Pubkey,
	)

	// add new account in pending state to DB and
	newAccount := types.NewPendingUserAccount(event.AccountID.Uint64(), event.Amount.Uint64(), event.Token.Uint64(), hex.EncodeToString(event.Pubkey))
	if err := s.DBInstance.InsertAccount(newAccount); err != nil {
		panic(err)
	}
}

func (s *Syncer) processDepositLeafMerged(eventName string, abiObject *abi.ABI, vLog *ethTypes.Log) {
	s.Logger.Info("Deposit Leaf merged")
	// unpack event
	event := new(logger.LoggerDepositLeafMerged)

	err := common.UnpackLog(abiObject, event, eventName, vLog)
	if err != nil {
		// TODO do something with this error
		fmt.Println("Unable to unpack log:", err)
		panic(err)
	}

	leftLeaf := types.ByteArray(event.Left)
	rightLeaf := types.ByteArray(event.Right)
	newRoot := types.ByteArray(event.NewRoot)

	s.Logger.Info(
		"⬜ New event found",
		"event", eventName,
		"prevDepositRoot", leftLeaf.String(),
		"incomingLeaf", rightLeaf.String(),
		"newDepositRoot", newRoot.String(),
	)

	// update deposit sub tree root
	newheight, err := s.DBInstance.OnDepositLeafMerge(leftLeaf, rightLeaf, newRoot)
	if err != nil {
		panic(err)
	}
	params, err := s.DBInstance.GetParams()
	if err != nil {
		panic(err)
	}
	// if deposit subtree height = deposit finalisation height then
	if newheight == params.MaxDepositSubTreeHeight {
		// send deposit finalisation transction to ethereum chain
		go s.sendDepositFinalisationTx()
	}
}

func (s *Syncer) processDepositFinalised(eventName string, abiObject *abi.ABI, vLog *ethTypes.Log) {
	s.Logger.Info("Deposits finalised")

	// unpack event
	event := new(logger.LoggerDepositsFinalised)

	err := common.UnpackLog(abiObject, event, eventName, vLog)
	if err != nil {
		// TODO do something with this error
		fmt.Println("Unable to unpack log:", err)
		panic(err)
	}
	accountsRoot := types.ByteArray(event.DepositSubTreeRoot)
	pathToDepositSubTree := event.PathToSubTree
	newBalanceRoot := types.ByteArray(event.NewBalanceRoot)

	s.Logger.Info(
		"⬜ New event found",
		"event", eventName,
		"DepositSubTreeRoot", accountsRoot.String(),
		"PathToDepositSubTreeInserted", pathToDepositSubTree.String(),
		"NewBalanceTreeRoot", newBalanceRoot.String(),
	)

	// TODO handle error
	s.DBInstance.FinaliseDeposits(accountsRoot, pathToDepositSubTree.Uint64(), newBalanceRoot)
}

func (s *Syncer) processNewBatch(eventName string, abiObject *abi.ABI, vLog *ethTypes.Log) {
	s.Logger.Info("New batch submitted on eth chain")

	event := new(rollup.RollupNewBatch)

	err := common.UnpackLog(abiObject, event, eventName, vLog)
	if err != nil {
		// TODO do something with this error
		fmt.Println("Unable to unpack log:", err)
		panic(err)
	}
	s.Logger.Info(
		"⬜ New event found",
		"event", eventName,
		"BatchNumber", event.Index.String(),
		"TxRoot", types.ByteArray(event.Txroot).String(),
		"NewStateRoot", types.ByteArray(event.UpdatedRoot).String(),
		"Committer", event.Committer.String(),
	)

	// TODO run the transactions through ProcessTx present on-chain
	// if any tx is fraud, challenge

	// pick the calldata for the batch
	txHash := vLog.TxHash
	txs, err := s.loadedBazooka.FetchBatchInputData(txHash)
	if err != nil {
		// TODO do something with this error
		panic(err)
	}

	newBatch := types.Batch{
		Index:                event.Index.Uint64(),
		StateRoot:            types.ByteArray(event.UpdatedRoot),
		TxRoot:               types.ByteArray(event.Txroot),
		TransactionsIncluded: txs,
		Committer:            event.Committer.String(),
		StakeAmount:          32,
		FinalisesOn:          *big.NewInt(100),
	}

	err = s.DBInstance.AddNewBatch(newBatch)
	if err != nil {
		// TODO do something with this error
		panic(err)
	}
}
func (s *Syncer) processRegisteredToken(eventName string, abiObject *abi.ABI, vLog *ethTypes.Log) {
	s.Logger.Info("New token registered")
	event := new(logger.LoggerRegisteredToken)

	err := common.UnpackLog(abiObject, event, eventName, vLog)
	if err != nil {
		// TODO do something with this error
		fmt.Println("Unable to unpack log:", err)
		panic(err)
	}
	s.Logger.Info(
		"⬜ New event found",
		"event", eventName,
		"TokenAddress", event.TokenContract.String(),
		"TokenID", event.TokenType,
	)
	newToken := types.Token{TokenID: event.TokenType.Uint64(), Address: types.Address(event.TokenContract)}
	if err := s.DBInstance.AddToken(newToken); err != nil {
		panic(err)
	}
}

func (s *Syncer) sendDepositFinalisationTx() {
	// get all 2**MaxDepositTreeHeight uninitialised accounts and find paths
	pathToNode, err := s.DBInstance.GetDepositNodePath()
	if err != nil {
		panic(err)
	}

	numberOfSiblings := len(pathToNode)
	var siblingsPath []string
	fmt.Println("data", numberOfSiblings, siblingsPath)
	i := numberOfSiblings
	for i > 0 {
		path := types.FlipBitInString(pathToNode, i-1)
		siblingsPath = append(siblingsPath, path)
		i--
	}
	fmt.Println("pathsneeded", siblingsPath)

	/*
		*** Optimise sibling nodes creation ***
		// we will need to find parent of only leaves behind the selected leaf node(where the deposit rooot is going to be)
		// for all the others we can find out the root using the default hashes
	*/

}
