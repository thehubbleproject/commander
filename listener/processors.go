package listener

import (
	"fmt"
	"math/big"

	"github.com/BOPR/common"
	"github.com/BOPR/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	ethTypes "github.com/ethereum/go-ethereum/core/types"

	"github.com/BOPR/contracts/rollup"
)

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
	txs, err := s.contractCaller.FetchBatchInputData(txHash)
	if err != nil {
		// TODO do something with this error
		panic(err)
	}

	newBatch := types.Batch{
		Index:       event.Index.Uint64(),
		StateRoot:   types.ByteArray(event.UpdatedRoot),
		TxRoot:      types.ByteArray(event.Txroot),
		Committer:   types.Address(event.Committer),
		StakeAmount: 100,
		FinalisesOn: *big.NewInt(100),
	}

	err = s.DBInstance.AddNewBatch(newBatch, txs)
	if err != nil {
		// TODO do something with this error
		panic(err)
	}
}

func (s *Syncer) processNewAccount(eventName string, abiObject *abi.ABI, vLog *ethTypes.Log) {
	s.Logger.Info("New account found")
}

func (s *Syncer) processDeposit(eventName string, abiObject *abi.ABI, vLog *ethTypes.Log) {
	s.Logger.Info("New deposit found")
	// add deposit to the account DB

}
