package bazooka

import (
	"context"
	"errors"
	"fmt"
	big "math/big"

	"github.com/BOPR/common"
	"github.com/BOPR/core"
	ethCmn "github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
)

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
	totalBatches, err := b.RollupContract.NumOfBatchesSubmitted(nil)
	if err != nil {
		return 0, err
	}
	return totalBatches.Uint64(), nil
}

func (b *Bazooka) FetchBalanceTreeRoot() (core.ByteArray, error) {
	root, err := b.RollupContract.GetLatestBalanceTreeRoot(nil)
	if err != nil {
		return core.ByteArray{}, err
	}
	return root, nil
}

func (b *Bazooka) FetchBatchInputData(txHash ethCmn.Hash) (txs [][]byte, err error) {
	tx, isPending, err := b.EthClient.TransactionByHash(context.Background(), txHash)
	if err != nil {
		b.log.Error("Cannot fetch transaction from hash", "Error", err)
		return
	}

	if isPending {
		err := errors.New("Transaction is pending")
		b.log.Error("Transaction is still pending, cannot process", "Error", err)
		return txs, err
	}

	payload := tx.Data()
	decodedPayload := payload[4:]
	fmt.Println("decoded payload", decodedPayload)
	inputDataMap := make(map[string]interface{})
	method := b.ContractABI[common.ROLLUP_CONTRACT_KEY].Methods["submitBatch"]
	err = method.Inputs.UnpackIntoMap(inputDataMap, decodedPayload)
	if err != nil {
		b.log.Error("Error unpacking payload", "Error", err)
		return
	}
	b.log.Debug("Created input data map", "InputData", inputDataMap)

	return GetTxsFromInput(inputDataMap), nil
}

// ProcessTx calls the ProcessTx function on the contract to verify the tx
// returns the updated accounts and the new balance root
func (b *Bazooka) ProcessTx(balanceTreeRoot, accountTreeRoot core.ByteArray, tx core.Tx, fromMerkleProof, toMerkleProof core.AccountMerkleProof, pdaProof core.PDAMerkleProof) (newBalanceRoot core.ByteArray, from, to core.UserAccount, err error) {
	txABIVersion := tx.ToABIVersion(int64(tx.From), int64(tx.To))

	updatedRoot, newBalFrom, newBalTo, IsValidTx, err := b.RollupContract.ProcessTx(nil,
		balanceTreeRoot,
		accountTreeRoot,
		txABIVersion,
		pdaProof.ToABIVersion(),
		fromMerkleProof.ToABIVersion(),
		toMerkleProof.ToABIVersion(),
	)
	if err != nil {
		return
	}

	b.log.Info("Processed transaction", "success", IsValidTx, "newRoot", updatedRoot)

	if !IsValidTx {
		return newBalanceRoot, from, to, errors.New("Tx is invalid")
	}

	newBalanceRoot = core.BytesToByteArray(updatedRoot[:])

	fromMerkleProof.Account.Balance = newBalFrom.Uint64()
	from = fromMerkleProof.Account

	toMerkleProof.Account.Balance = newBalTo.Uint64()
	to = toMerkleProof.Account

	return newBalanceRoot, from, to, nil
}

func (b *Bazooka) VerifyPDAProof(accountsRoot core.ByteArray, leaf core.ByteArray, proofpath *big.Int, siblings [][32]byte) (bool, error) {
	return b.BalanceTree.VerifyLeaf(nil, accountsRoot, leaf, proofpath, siblings)
}

func (b *Bazooka) ValidateSignature(tx core.Tx, pdaProof core.PDAMerkleProof) error {
	return b.RollupContract.ValidateSignature(nil, tx.ToABIVersion(int64(tx.From), int64(tx.To)), pdaProof.ToABIVersion())
}

func (b *Bazooka) ValidateAccountMP(root core.ByteArray, accountMP core.AccountMerkleProof) error {
	return b.RollupContract.ValidateAccountMP(nil, root, accountMP.ToABIVersion())
}
