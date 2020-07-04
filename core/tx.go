package core

import (
	"context"
	"errors"
	"fmt"
	"math/big"

	"encoding/hex"

	"github.com/BOPR/common"
	"github.com/BOPR/config"
	"github.com/BOPR/contracts/rollup"
	"github.com/ethereum/go-ethereum"
	ethCmn "github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/rlp"
	"golang.org/x/crypto/sha3"
)

// Tx represets the transaction on hubble
type Tx struct {
	DBModel
	To        uint64 `json:"to"`
	From      uint64 `json:"from"`
	Data      []byte `json:"data"`
	Signature string `json:"sig" gorm:"not null"`
	TxHash    string `json:"hash" gorm:"not null"`
	Status    uint64 `json:"status"`
}

// NewTx creates a new transaction
func NewTx(from, to uint64, message []byte, sig string) Tx {
	return Tx{
		From:      from,
		To:        to,
		Data:      message,
		Signature: sig,
	}
}

// NewPendingTx creates a new transaction
func NewPendingTx(from, to uint64, sig string, message []byte) Tx {
	return Tx{
		To:        to,
		From:      from,
		Data:      message,
		Signature: sig,
		Status:    TX_STATUS_PENDING,
	}
}

// GetSignBytes returns the transaction data that has to be signed
func (tx Tx) GetSignBytes() (signBytes []byte, err error) {
	// TODO: [apply-tx] call bazooka
	return
}

// AssignHash creates a tx hash and add it to the tx
func (t *Tx) AssignHash() {
	if t.TxHash != "" {
		return
	}
	hash := rlpHash(t)
	t.TxHash = hash.String()
}

func (tx *Tx) Apply(updatedFrom, updatedTo []byte) error {
	// get fresh copy of database
	db, err := NewDB()
	if err != nil {
		return err
	}

	// begin a transaction
	mysqlTx := db.Instance.Begin()
	defer func() {
		if r := recover(); r != nil {
			mysqlTx.Rollback()
		}
	}()

	// post this perform all ops on transaction
	db.Instance = mysqlTx

	// apply transaction on from account
	fromAcc, err := db.GetAccountByID(tx.From)
	if err != nil {
		mysqlTx.Rollback()
		return err
	}

	fromAcc.Data = updatedFrom

	err = db.UpdateAccount(fromAcc)
	if err != nil {
		mysqlTx.Rollback()
		return err
	}

	// apply transaction on to account
	toAcc, err := db.GetAccountByID(tx.To)
	if err != nil {
		mysqlTx.Rollback()
		return err
	}

	toAcc.Data = updatedTo

	err = db.UpdateAccount(toAcc)
	if err != nil {
		mysqlTx.Rollback()
		return err
	}

	tx.UpdateStatus(TX_STATUS_PROCESSED)

	// Or commit the transaction
	mysqlTx.Commit()
	return nil
}

func (t *Tx) String() string {
	return fmt.Sprintf("To: %v From: %v Status:%v Hash: %v Data: %v", t.To, t.From, t.Status, t.TxHash, hex.EncodeToString(t.Data))
}

// ToABIVersion converts a standard tx to the the DataTypesTransaction struct on the contract
func (t *Tx) ToABIVersion(from, to int64) rollup.TypesTransaction {
	decodedSignature, _ := hex.DecodeString(t.Signature)
	// TODO call decode transaction with bazooka
	return rollup.TypesTransaction{
		ToIndex:   big.NewInt(to),
		FromIndex: big.NewInt(from),
		Signature: decodedSignature,
	}
}

func (tx *Tx) Compress() ([]byte, error) {
	// TODO : [apply-tx] call bazooka

	var bytes []byte
	return bytes, nil
}

// Insert tx into the DB
func (db *DB) InsertTx(t *Tx) error {
	return db.Instance.Create(t).Error
}

func (db *DB) PopTxs() (txs []Tx, err error) {
	tx := db.Instance.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return txs, err
	}
	var pendingTxs []Tx

	// select N number of transactions which are pending in mempool and
	if err := tx.Limit(config.GlobalCfg.TxsPerBatch).Where(&Tx{Status: 100}).Find(&pendingTxs).Error; err != nil {
		db.Logger.Error("error while fetching pending transactions", err)
		return txs, err
	}

	db.Logger.Info("found txs", "pendingTxs", pendingTxs)

	var ids []string
	for _, tx := range pendingTxs {
		ids = append(ids, tx.ID)
	}

	// update the transactions from pending to processing
	errs := tx.Table("txes").Where("id IN (?)", ids).Updates(map[string]interface{}{"status": 200}).GetErrors()
	if err != nil {
		db.Logger.Error("errors while processing transactions", errs)
		return
	}

	return pendingTxs, tx.Commit().Error
}

func (db *DB) GetTx() (tx []Tx, err error) {
	err = db.Instance.First(&tx).Error
	if err != nil {
		return tx, err
	}
	return
}

func (tx *Tx) UpdateStatus(status uint64) error {
	return DBInstance.Instance.Model(&tx).Update("status", status).Error
}

func rlpHash(x interface{}) (h ethCmn.Hash) {
	hw := sha3.NewLegacyKeccak256()
	rlp.Encode(hw, x)
	hw.Sum(h[:0])
	return h
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
	totalBatches, err := b.RollupContract.NumOfBatchesSubmitted(nil)
	if err != nil {
		return 0, err
	}
	return totalBatches.Uint64(), nil
}

func (b *Bazooka) FetchBalanceTreeRoot() (ByteArray, error) {
	root, err := b.RollupContract.GetLatestBalanceTreeRoot(nil)
	if err != nil {
		return ByteArray{}, err
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
func (b *Bazooka) ProcessTx(balanceTreeRoot, accountTreeRoot ByteArray, tx Tx, fromMerkleProof, toMerkleProof AccountMerkleProof, pdaProof PDAMerkleProof) (newBalanceRoot ByteArray, from, to []byte, err error) {
	txABIVersion := tx.ToABIVersion(int64(tx.From), int64(tx.To))
	opts := bind.CallOpts{From: config.OperatorAddress}
	typesAccountProofs := rollup.TypesAccountProofs{From: fromMerkleProof.ToABIVersion(), To: toMerkleProof.ToABIVersion()}
	updatedRoot, newFromAccount, newToAccount, errCode, IsValidTx, err := b.RollupContract.ProcessTx(&opts,
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
		b.log.Error("Invalid transaction", "error_code", errCode)
		return newBalanceRoot, from, to, errors.New("Tx is invalid")
	}
	newBalanceRoot = BytesToByteArray(updatedRoot[:])
	return newBalanceRoot, newFromAccount, newToAccount, nil
}

func (b *Bazooka) ApplyTransferTx(account AccountMerkleProof, tx Tx) ([]byte, ByteArray, error) {
	txABIVersion := tx.ToABIVersion(int64(tx.From), int64(tx.To))
	updatedAccountBytes, updatedRoot, err := b.RollupContract.ApplyTx(nil, account.ToABIVersion(), txABIVersion)
	if err != nil {
		return updatedAccountBytes, updatedRoot, err
	}

	return updatedAccountBytes, updatedRoot, nil
}

// func (b *Bazooka) CompressTx(tx Tx) ([]byte, ByteArray, error) {
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

func (b *Bazooka) FireDepositFinalisation(TBreplaced UserAccount, siblings []UserAccount, subTreeHeight uint64) error {
	b.log.Info(
		"Attempting to finalise deposits",
		"NodeToBeReplaced",
		TBreplaced.String(),
		"NumberOfSiblings",
		len(siblings),
		"atDepth",
		subTreeHeight,
	)

	// TODO check latest batch on-chain and if we need to push new batch

	depositSubTreeHeight := big.NewInt(0)
	depositSubTreeHeight.SetUint64(subTreeHeight)
	var siblingData [][32]byte
	for _, sibling := range siblings {
		data, err := HexToByteArray(sibling.Hash)
		if err != nil {
			return err
		}
		siblingData = append(siblingData, data)
	}
	accountProof := rollup.TypesAccountMerkleProof{}
	accountProof.AccountIP.PathToAccount = StringToBigInt(TBreplaced.Path)
	accountProof.AccountIP.Account = TBreplaced.ToABIAccount()
	accountProof.Siblings = siblingData
	data, err := b.ContractABI[common.ROLLUP_CONTRACT_KEY].Pack("finaliseDepositsAndSubmitBatch", depositSubTreeHeight, accountProof)
	if err != nil {
		return err
	}

	rollupAddress := ethCmn.HexToAddress(config.GlobalCfg.RollupAddress)
	stakeAmount := big.NewInt(0)
	stakeAmount.SetString("32000000000000000000", 10)

	// generate call msg
	callMsg := ethereum.CallMsg{
		To:    &rollupAddress,
		Data:  data,
		Value: stakeAmount,
	}

	auth, err := b.GenerateAuthObj(b.EthClient, callMsg)
	if err != nil {
		return err
	}
	b.log.Info("Broadcasting deposit finalisation transaction")
	tx, err := b.RollupContract.FinaliseDepositsAndSubmitBatch(auth, depositSubTreeHeight, accountProof)
	if err != nil {
		return err
	}
	b.log.Info("Deposits successfully finalized!", "TxHash", tx.Hash())
	return nil
}

// SubmitBatch submits the batch on chain with updated root and compressed transactions
func (b *Bazooka) SubmitBatch(updatedRoot ByteArray, txs []Tx) error {
	b.log.Info(
		"Attempting to submit a new batch",
		"UpdatedRoot",
		updatedRoot.String(),
		"txs",
		len(txs),
	)

	var compressedTxs [][]byte
	for _, tx := range txs {
		compressedTx, err := tx.Compress()
		if err != nil {
			return err
		}
		compressedTxs = append(compressedTxs, compressedTx)
	}

	data, err := b.ContractABI[common.ROLLUP_CONTRACT_KEY].Pack("submitBatch", compressedTxs, updatedRoot)
	if err != nil {
		return err
	}

	rollupAddress := ethCmn.HexToAddress(config.GlobalCfg.RollupAddress)

	// generate call msg
	callMsg := ethereum.CallMsg{
		To:   &rollupAddress,
		Data: data,
	}

	auth, err := b.GenerateAuthObj(b.EthClient, callMsg)
	if err != nil {
		return err
	}

	tx, err := b.RollupContract.SubmitBatch(auth, compressedTxs, updatedRoot)
	if err != nil {
		return err
	}
	b.log.Info("Sent a new batch!", "txHash", tx.Hash().String())
	return nil
}
