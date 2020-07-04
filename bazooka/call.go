package bazooka

import (
	"context"
	"errors"
	big "math/big"

	"github.com/BOPR/common"
	"github.com/BOPR/config"
	"github.com/BOPR/contracts/rollup"
	"github.com/BOPR/core"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
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
	inputDataMap := make(map[string]interface{})
	method := b.ContractABI[common.ROLLUP_CONTRACT_KEY].Methods["submitBatch"]
	err = method.Inputs.UnpackIntoMap(inputDataMap, decodedPayload)
	if err != nil {
		b.log.Error("Error unpacking payload", "Error", err)
		return
	}

	return GetTxsFromInput(inputDataMap), nil
}

// ProcessTx calls the ProcessTx function on the contract to verify the tx
// returns the updated accounts and the new balance root
func (b *Bazooka) ProcessTx(balanceTreeRoot, accountTreeRoot core.ByteArray, tx core.Tx, fromMerkleProof, toMerkleProof core.AccountMerkleProof, pdaProof core.PDAMerkleProof) (newBalanceRoot core.ByteArray, from, to []byte, err error) {
	txABIVersion := tx.ToABIVersion(int64(tx.From), int64(tx.To))
	opts := bind.CallOpts{From: config.OperatorAddress()}
	typesAccountProofs := rollup.TypesAccountProofs{From: fromMerkleProof.ToABIVersion(), To: toMerkleProof.ToABIVersion()}
	updatedRoot, newFromAccount, newToAccount, err_code, IsValidTx, err := b.RollupContract.ProcessTx(&opts,
		balanceTreeRoot,
		accountTreeRoot,
		txABIVersion,
		pdaProof.ToABIVersion(),
		typesAccountProofs,
	)
	if err != nil {
		return
	}

	b.log.Info("Processed transaction", "IsSuccess", IsValidTx, "newRoot", updatedRoot)

	if !IsValidTx {
		b.log.Error("Invalid transaction", "error_code", err_code)
		return newBalanceRoot, from, to, errors.New("Tx is invalid")
	}
	newBalanceRoot = core.BytesToByteArray(updatedRoot[:])
	return newBalanceRoot, newFromAccount, newToAccount, nil
}

func (b *Bazooka) ApplyTransferTx(account core.AccountMerkleProof, tx core.Tx) ([]byte, core.ByteArray, error) {
	txABIVersion := tx.ToABIVersion(int64(tx.From), int64(tx.To))
	updatedAccountBytes, updatedRoot, err := b.RollupContract.ApplyTx(nil, account.ToABIVersion(), txABIVersion)
	if err != nil {
		return updatedAccountBytes, updatedRoot, err
	}

	return updatedAccountBytes, updatedRoot, nil
}

// func (b *Bazooka) CompressTx(tx core.Tx) ([]byte, core.ByteArray, error) {
// 	txABIVersion := tx.ToABIVersion(int64(tx.From), int64(tx.To))

// 	b.RollupUtils.CompressTx(nil, txABIVersion)
// }

// func (b *Bazooka) EncodeTx(from, to, token, amount int64) ([]byte, error) {
// 	return b.RollupUtils.BytesFromTxDeconstructed(nil, big.NewInt(from), big.NewInt(to), big.NewInt(token), big.NewInt(amount))
// }

// func (b *Bazooka) DecodeTx(txBytes []byte) (rolluputils.TypesTransaction, error) {
// 	tx, err := b.RollupUtils.TxFromBytes(nil, txBytes)
// 	return tx, err
// }

func (b *Bazooka) EncodeAccount(id, balance, nonce, token int64) (accountBytes []byte, err error) {
	accountBytes, err = b.RollupUtils.BytesFromAccountDeconstructed(nil, big.NewInt(id), big.NewInt(balance), big.NewInt(nonce), big.NewInt(token))
	if err != nil {
		return
	}
	return accountBytes, nil
}

func (b *Bazooka) DecodeAccount(accountBytes []byte) (ID, balance, nonce, token uint64, err error) {
	account, err := b.RollupUtils.AccountFromBytes(nil, accountBytes)
	if err != nil {
		return
	}

	return account.ID.Uint64(), account.Balance.Uint64(), account.Nonce.Uint64(), account.TokenType.Uint64(), nil
}
