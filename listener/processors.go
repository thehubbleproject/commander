package listener

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
)

func (s *Syncer) processDeposit(eventName string, abiObject *abi.ABI, vLog *ethTypes.Log) {
	s.Logger.Info("New deposit found")
	// add deposit to the account DB

}

func (s *Syncer) processNewBatch(eventName string, abiObject *abi.ABI, vLog *ethTypes.Log) {
	s.Logger.Info("New batch found")
	// event := new(rollupContract.RollupNewBatch)
	// types.NewBatch(event.UpdatedRoot, event.Comm)

}

func (s *Syncer) processNewAccount(eventName string, abiObject *abi.ABI, vLog *ethTypes.Log) {
	s.Logger.Info("New account found")
}
