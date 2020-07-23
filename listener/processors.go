package listener

import (
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"

	"github.com/BOPR/common"
	"github.com/BOPR/core"
	"github.com/ethereum/go-ethereum/accounts/abi"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/jinzhu/gorm"

	"github.com/BOPR/contracts/logger"
)

// ZEROROOT
const ZEROROOT = "0x0000000000000000000000000000000000000000000000000000000000000000"

func (s *Syncer) processNewPubkeyAddition(eventName string, abiObject *abi.ABI, vLog *ethTypes.Log) {
	s.Logger.Info("New deposit found")

	// unpack event
	event := new(logger.LoggerNewPubkeyAdded)
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
		"pubkey", event.Pubkey,
	)

	// add new account in pending state to DB and
	pathToNode, err := core.SolidityPathToNodePath(event.AccountID.Uint64(), 4)
	if err != nil {
		// TODO do something with this error
		fmt.Println("Unable to convert path", err)
		panic(err)
	}
	newPDAAccount, err := core.NewPDA(event.AccountID.Uint64(), hex.EncodeToString(event.Pubkey), pathToNode)
	if err != nil {
		fmt.Println("unable to create new account")
		panic(err)
	}
	if err := s.DBInstance.UpdatePDALeaf(*newPDAAccount); err != nil {
		panic(err)
	}
}
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
		"Amount", hex.EncodeToString(event.Data),
		"pubkey", event.Pubkey,
	)

	// add new account in pending state to DB and
	newAccount := core.NewPendingUserAccount(event.AccountID.Uint64(), event.Data)
	if err := s.DBInstance.AddNewPendingAccount(*newAccount); err != nil {
		panic(err)
	}
}

func (s *Syncer) processDepositSubtreeCreated(eventName string, abiObject *abi.ABI, vLog *ethTypes.Log) {
	s.Logger.Info("New deposit subtree created")
	// unpack event
	event := new(logger.LoggerDepositSubTreeReady)
	err := common.UnpackLog(abiObject, event, eventName, vLog)
	if err != nil {
		// TODO do something with this error
		fmt.Println("Unable to unpack log:", err)
		panic(err)
	}
	err = s.DBInstance.AttachDepositInfo(event.Root)
	if err != nil {
		// TODO do something with this error
		fmt.Println("Unable to attack deposit information:", err)
		panic(err)
	}
	// TODO add a sync flag, do not send transactions when in sync mode

	// send deposit finalisation transction to ethereum chain
	s.SendDepositFinalisationTx()
}

func (s *Syncer) processDepositFinalised(eventName string, abiObject *abi.ABI, vLog *ethTypes.Log) {
	s.Logger.Info("Deposit batch finalised!")

	// unpack event
	event := new(logger.LoggerDepositsFinalised)

	err := common.UnpackLog(abiObject, event, eventName, vLog)
	if err != nil {
		// TODO do something with this error
		fmt.Println("Unable to unpack log:", err)
		panic(err)
	}

	depositRoot := core.ByteArray(event.DepositSubTreeRoot)
	pathToDepositSubTree := event.PathToSubTree

	s.Logger.Info(
		"⬜ New event found",
		"event", eventName,
		"DepositSubTreeRoot", depositRoot.String(),
		"PathToDepositSubTreeInserted", pathToDepositSubTree.String(),
	)

	// TODO handle error
	newRoot, err := s.DBInstance.FinaliseDepositsAndAddBatch(depositRoot, pathToDepositSubTree.Uint64())
	if err != nil {
		fmt.Println("Error while finalising deposits", err)
	}

	fmt.Println("new root", newRoot)
}

func (s *Syncer) processNewBatch(eventName string, abiObject *abi.ABI, vLog *ethTypes.Log) {
	s.Logger.Info("New batch submitted on eth chain")

	event := new(logger.LoggerNewBatch)

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
		"TxRoot", core.ByteArray(event.Txroot).String(),
		"NewStateRoot", core.ByteArray(event.UpdatedRoot).String(),
		"Committer", event.Committer.String(),
	)

	params, err := s.DBInstance.GetParams()
	if err != nil {
		return
	}

	// if the batch has some txs, parse them
	var txs [][]byte
	if ZEROROOT != core.ByteArray(event.Txroot).String() {
		// pick the calldata for the batch
		txs, err = s.loadedBazooka.FetchBatchInputData(vLog.TxHash)
		if err != nil {
			// TODO do something with this error
			panic(err)
		}
	}

	batch, err := s.DBInstance.GetBatchByIndex(event.Index.Uint64())
	// if we havent seen the batch, apply txs and store batch
	if err != nil && gorm.IsRecordNotFoundError(err) {
		s.Logger.Info("Found a new batch, applying transactions and adding new batch", "index", event.Index.Uint64)
		err := s.ApplyTxsFromBatch(txs, uint64(event.BatchType))
		if err != nil {
			panic(err)
		}

		newBatch := core.Batch{
			BatchID:              event.Index.Uint64(),
			StateRoot:            core.ByteArray(event.UpdatedRoot).String(),
			TxRoot:               core.ByteArray(event.Txroot).String(),
			TransactionsIncluded: core.ConcatTxs(txs),
			Committer:            event.Committer.String(),
			StakeAmount:          params.StakeAmount,
			FinalisesOn:          *big.NewInt(int64(params.FinalisationTime)),
			Status:               core.BATCH_COMMITTED,
		}

		err = s.DBInstance.AddNewBatch(newBatch)
		if err != nil {
			// TODO do something with this error
			panic(err)
		}
		return
	} else if err != nil {
		s.Logger.Error("Unable to fetch batch", "index", event.Index, "err", err)
		return
	}

	// if batch is present but in a non committed state we parse txs and commit batch
	if batch.Status != core.BATCH_COMMITTED {
		s.Logger.Info("Found a non committed batch")
		if batch.StateRoot != core.ByteArray(event.UpdatedRoot).String() {
			// State root mismatch error
		}
		// batch broadcasted by us
		// txs applied but batch needs to be committed
		// TODO add batch type
		newBatch := core.Batch{
			BatchID:              event.Index.Uint64(),
			StateRoot:            core.ByteArray(event.UpdatedRoot).String(),
			TxRoot:               core.ByteArray(event.Txroot).String(),
			TransactionsIncluded: core.ConcatTxs(txs),
			Committer:            event.Committer.String(),
			StakeAmount:          params.StakeAmount,
			FinalisesOn:          *big.NewInt(int64(params.FinalisationTime)),
			Status:               core.BATCH_COMMITTED,
		}
		s.DBInstance.CommitBatch(newBatch)
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
	newToken := core.Token{TokenID: event.TokenType.Uint64(), Address: event.TokenContract.String()}
	if err := s.DBInstance.AddToken(newToken); err != nil {
		panic(err)
	}
}

func (s *Syncer) SendDepositFinalisationTx() {
	params, err := s.DBInstance.GetParams()
	if err != nil {
		return
	}
	nodeToBeReplaced, siblings, err := s.DBInstance.GetDepositNodeAndSiblings()
	if err != nil {
		return
	}

	err = s.loadedBazooka.FireDepositFinalisation(nodeToBeReplaced, siblings, params.MaxDepositSubTreeHeight)
}

func (s *Syncer) ApplyTxsFromBatch(txs [][]byte, txType uint64) error {
	if len(txs) == 0 {
		s.Logger.Info("No txs to apply")
		return nil
	}
	var coreTxs []core.Tx
	for i := range txs {
		var from, to uint64
		var txData, sig []byte

		switch txType {
		case core.TX_TRANSFER_TYPE:
			from, to, amount, sig, err := s.loadedBazooka.DecompressTransferTx(txs[i])
			if err != nil {
				return err
			}
			s.Logger.Debug("Fetched tx data", "from", from, "to", to, "amount", amount, "sig", sig)
			fromAccount, err := s.DBInstance.GetAccountByID(from.Uint64())
			if err != nil {
				return err
			}
			_, _, nonce, token, err := s.loadedBazooka.DecodeAccount(fromAccount.Data)
			if err != nil {
				return err
			}
			s.Logger.Debug("Decoded account", "nonce", nonce, "token", token)
			txData, err = s.loadedBazooka.EncodeTransferTx(from.Int64(), to.Int64(), token.Int64(), nonce.Int64(), amount.Int64(), core.TX_TRANSFER_TYPE)
			if err != nil {
				return err
			}
		default:
			fmt.Println("TxType didnt match any options", txType)
			return errors.New("Didn't match any options")
		}
		coreTx := core.NewTx(from, to, txType, txData, hex.EncodeToString(sig))
		coreTxs = append(coreTxs, coreTx)
		fromMP, toMP, _, err := coreTx.GetVerificationData()
		if err != nil {
			return err
		}

		updatedFromAccData, _, err := s.loadedBazooka.ApplyTx(fromMP, coreTx)
		if err != nil {
			return err
		}

		updatedToAccData, _, err := s.loadedBazooka.ApplyTx(toMP, coreTx)
		if err != nil {
			return err
		}

		err = coreTx.Apply(updatedFromAccData, updatedToAccData)
		if err != nil {
			return err
		}
		// TODO validate updated root post application
		// root, err := s.DBInstance.GetRoot()
		// if err != nil {
		// 	return err
		// }
	}
	return nil
}
