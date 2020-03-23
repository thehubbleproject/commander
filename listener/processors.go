package listener

import (
	"fmt"

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
		fmt.Println("Unable to unpack log:", err)
		return
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
	_ = vLog.TxHash

}

func (s *Syncer) processNewAccount(eventName string, abiObject *abi.ABI, vLog *ethTypes.Log) {
	s.Logger.Info("New account found")
}

func (s *Syncer) processDeposit(eventName string, abiObject *abi.ABI, vLog *ethTypes.Log) {
	s.Logger.Info("New deposit found")
	// add deposit to the account DB

}
