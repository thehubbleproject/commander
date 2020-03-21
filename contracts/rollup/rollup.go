// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package rollup

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// DataTypesAccountInclusionProof is an auto generated low-level Go binding around an user-defined struct.
type DataTypesAccountInclusionProof struct {
	PathToAccount *big.Int
	Account       DataTypesUserAccount
}

// DataTypesAccountMerkleProof is an auto generated low-level Go binding around an user-defined struct.
type DataTypesAccountMerkleProof struct {
	AccountIP DataTypesAccountInclusionProof
	Siblings  [][32]byte
}

// DataTypesTransaction is an auto generated low-level Go binding around an user-defined struct.
type DataTypesTransaction struct {
	From      DataTypesUserAccount
	To        DataTypesUserAccount
	TokenType *big.Int
	Amount    uint32
	Signature []byte
}

// DataTypesUserAccount is an auto generated low-level Go binding around an user-defined struct.
type DataTypesUserAccount struct {
	ID        *big.Int
	TokenType *big.Int
	Balance   *big.Int
	Nonce     *big.Int
}

// RollupABI is the input ABI used to generate the binding from.
const RollupABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_balancesTree\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_accountsTree\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_merkleTreeLib\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tokenRegistryAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_coordinator\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"batch_id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"committer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"txRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stakeSlashed\",\"type\":\"uint256\"}],\"name\":\"BatchRollback\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"DepositLeafMerged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"DepositQueued\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"DepositsProcessed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"NewAccount\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"committer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"txroot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"updatedRoot\",\"type\":\"bytes32\"}],\"name\":\"NewBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tokenContract\",\"type\":\"address\"}],\"name\":\"RegisteredToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tokenContract\",\"type\":\"address\"}],\"name\":\"RegistrationRequest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalBatchesSlashed\",\"type\":\"uint256\"}],\"name\":\"RollbackFinalisation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"committed\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"batch_id\",\"type\":\"uint256\"}],\"name\":\"StakeWithdraw\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"ZERO_BYTES32\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"accountsTree\",\"outputs\":[{\"internalType\":\"contractMerkleTree\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"balancesTree\",\"outputs\":[{\"internalType\":\"contractMerkleTree\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"batches\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"committer\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"txRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"stakeCommitted\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"finalisesOn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"depositSubtreeHeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"merkleTreeLib\",\"outputs\":[{\"internalType\":\"contractMerkleTreeLib\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"pendingDeposits\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"queueNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tokenContract\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tokenRegistry\",\"outputs\":[{\"internalType\":\"contractITokenRegistry\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"_txs\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"_updatedRoot\",\"type\":\"bytes32\"}],\"name\":\"submitBatch\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_batch_id\",\"type\":\"uint256\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"ID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.UserAccount\",\"name\":\"from\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"ID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.UserAccount\",\"name\":\"to\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"amount\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structDataTypes.Transaction[]\",\"name\":\"_txs\",\"type\":\"tuple[]\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"pathToAccount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"ID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.UserAccount\",\"name\":\"account\",\"type\":\"tuple\"}],\"internalType\":\"structDataTypes.AccountInclusionProof\",\"name\":\"accountIP\",\"type\":\"tuple\"},{\"internalType\":\"bytes32[]\",\"name\":\"siblings\",\"type\":\"bytes32[]\"}],\"internalType\":\"structDataTypes.AccountMerkleProof[]\",\"name\":\"_from_proofs\",\"type\":\"tuple[]\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"pathToAccount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"ID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.UserAccount\",\"name\":\"account\",\"type\":\"tuple\"}],\"internalType\":\"structDataTypes.AccountInclusionProof\",\"name\":\"accountIP\",\"type\":\"tuple\"},{\"internalType\":\"bytes32[]\",\"name\":\"siblings\",\"type\":\"bytes32[]\"}],\"internalType\":\"structDataTypes.AccountMerkleProof[]\",\"name\":\"_to_proofs\",\"type\":\"tuple[]\"}],\"name\":\"disputeBatch\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_balanceRoot\",\"type\":\"bytes32\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"ID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.UserAccount\",\"name\":\"from\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"ID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.UserAccount\",\"name\":\"to\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"amount\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structDataTypes.Transaction\",\"name\":\"_tx\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"pathToAccount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"ID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.UserAccount\",\"name\":\"account\",\"type\":\"tuple\"}],\"internalType\":\"structDataTypes.AccountInclusionProof\",\"name\":\"accountIP\",\"type\":\"tuple\"},{\"internalType\":\"bytes32[]\",\"name\":\"siblings\",\"type\":\"bytes32[]\"}],\"internalType\":\"structDataTypes.AccountMerkleProof\",\"name\":\"_from_merkle_proof\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"pathToAccount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"ID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.UserAccount\",\"name\":\"account\",\"type\":\"tuple\"}],\"internalType\":\"structDataTypes.AccountInclusionProof\",\"name\":\"accountIP\",\"type\":\"tuple\"},{\"internalType\":\"bytes32[]\",\"name\":\"siblings\",\"type\":\"bytes32[]\"}],\"internalType\":\"structDataTypes.AccountMerkleProof\",\"name\":\"_to_merkle_proof\",\"type\":\"tuple\"}],\"name\":\"processTx\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"SlashAndRollback\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_tokenType\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_pubkey\",\"type\":\"bytes\"}],\"name\":\"deposit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_destination\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_tokenType\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_pubkey\",\"type\":\"bytes\"}],\"name\":\"depositFor\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_subTreeDepth\",\"type\":\"uint256\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"pathToAccount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"ID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.UserAccount\",\"name\":\"account\",\"type\":\"tuple\"}],\"internalType\":\"structDataTypes.AccountInclusionProof\",\"name\":\"accountIP\",\"type\":\"tuple\"},{\"internalType\":\"bytes32[]\",\"name\":\"siblings\",\"type\":\"bytes32[]\"}],\"internalType\":\"structDataTypes.AccountMerkleProof\",\"name\":\"_zero_account_mp\",\"type\":\"tuple\"}],\"name\":\"finaliseDeposits\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenContractAddress\",\"type\":\"address\"}],\"name\":\"requestTokenRegistration\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenContractAddress\",\"type\":\"address\"}],\"name\":\"finaliseTokenRegistration\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"batch_id\",\"type\":\"uint256\"}],\"name\":\"WithdrawStake\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"ID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.UserAccount\",\"name\":\"original_account\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"new_balance\",\"type\":\"uint256\"}],\"name\":\"UpdateBalanceInAccount\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"ID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.UserAccount\",\"name\":\"new_account\",\"type\":\"tuple\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"ID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.UserAccount\",\"name\":\"account\",\"type\":\"tuple\"}],\"name\":\"BalanceFromAccount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"ID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.UserAccount\",\"name\":\"account\",\"type\":\"tuple\"}],\"name\":\"HashFromAccount\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"ID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.UserAccount\",\"name\":\"account\",\"type\":\"tuple\"}],\"name\":\"BytesFromAccount\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"ID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.UserAccount\",\"name\":\"from\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"ID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.UserAccount\",\"name\":\"to\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"amount\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structDataTypes.Transaction\",\"name\":\"_tx\",\"type\":\"tuple\"}],\"name\":\"BytesFromTx\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"ID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.UserAccount\",\"name\":\"from\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"ID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.UserAccount\",\"name\":\"to\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"amount\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structDataTypes.Transaction\",\"name\":\"_tx\",\"type\":\"tuple\"}],\"name\":\"HashFromTx\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getBalanceTreeRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"a\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"b\",\"type\":\"bytes32\"}],\"name\":\"getDepositsHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"numberOfBatches\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"pub\",\"type\":\"bytes\"}],\"name\":\"calculateAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// Rollup is an auto generated Go binding around an Ethereum contract.
type Rollup struct {
	RollupCaller     // Read-only binding to the contract
	RollupTransactor // Write-only binding to the contract
	RollupFilterer   // Log filterer for contract events
}

// RollupCaller is an auto generated read-only Go binding around an Ethereum contract.
type RollupCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RollupTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RollupTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RollupFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RollupFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RollupSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RollupSession struct {
	Contract     *Rollup           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RollupCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RollupCallerSession struct {
	Contract *RollupCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// RollupTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RollupTransactorSession struct {
	Contract     *RollupTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RollupRaw is an auto generated low-level Go binding around an Ethereum contract.
type RollupRaw struct {
	Contract *Rollup // Generic contract binding to access the raw methods on
}

// RollupCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RollupCallerRaw struct {
	Contract *RollupCaller // Generic read-only contract binding to access the raw methods on
}

// RollupTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RollupTransactorRaw struct {
	Contract *RollupTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRollup creates a new instance of Rollup, bound to a specific deployed contract.
func NewRollup(address common.Address, backend bind.ContractBackend) (*Rollup, error) {
	contract, err := bindRollup(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Rollup{RollupCaller: RollupCaller{contract: contract}, RollupTransactor: RollupTransactor{contract: contract}, RollupFilterer: RollupFilterer{contract: contract}}, nil
}

// NewRollupCaller creates a new read-only instance of Rollup, bound to a specific deployed contract.
func NewRollupCaller(address common.Address, caller bind.ContractCaller) (*RollupCaller, error) {
	contract, err := bindRollup(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RollupCaller{contract: contract}, nil
}

// NewRollupTransactor creates a new write-only instance of Rollup, bound to a specific deployed contract.
func NewRollupTransactor(address common.Address, transactor bind.ContractTransactor) (*RollupTransactor, error) {
	contract, err := bindRollup(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RollupTransactor{contract: contract}, nil
}

// NewRollupFilterer creates a new log filterer instance of Rollup, bound to a specific deployed contract.
func NewRollupFilterer(address common.Address, filterer bind.ContractFilterer) (*RollupFilterer, error) {
	contract, err := bindRollup(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RollupFilterer{contract: contract}, nil
}

// bindRollup binds a generic wrapper to an already deployed contract.
func bindRollup(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RollupABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Rollup *RollupRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Rollup.Contract.RollupCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Rollup *RollupRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Rollup.Contract.RollupTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Rollup *RollupRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Rollup.Contract.RollupTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Rollup *RollupCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Rollup.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Rollup *RollupTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Rollup.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Rollup *RollupTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Rollup.Contract.contract.Transact(opts, method, params...)
}

// BalanceFromAccount is a free data retrieval call binding the contract method 0x0f9b2cb2.
//
// Solidity: function BalanceFromAccount(DataTypesUserAccount account) constant returns(uint256)
func (_Rollup *RollupCaller) BalanceFromAccount(opts *bind.CallOpts, account DataTypesUserAccount) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Rollup.contract.Call(opts, out, "BalanceFromAccount", account)
	return *ret0, err
}

// BalanceFromAccount is a free data retrieval call binding the contract method 0x0f9b2cb2.
//
// Solidity: function BalanceFromAccount(DataTypesUserAccount account) constant returns(uint256)
func (_Rollup *RollupSession) BalanceFromAccount(account DataTypesUserAccount) (*big.Int, error) {
	return _Rollup.Contract.BalanceFromAccount(&_Rollup.CallOpts, account)
}

// BalanceFromAccount is a free data retrieval call binding the contract method 0x0f9b2cb2.
//
// Solidity: function BalanceFromAccount(DataTypesUserAccount account) constant returns(uint256)
func (_Rollup *RollupCallerSession) BalanceFromAccount(account DataTypesUserAccount) (*big.Int, error) {
	return _Rollup.Contract.BalanceFromAccount(&_Rollup.CallOpts, account)
}

// BytesFromAccount is a free data retrieval call binding the contract method 0x3035226f.
//
// Solidity: function BytesFromAccount(DataTypesUserAccount account) constant returns(bytes)
func (_Rollup *RollupCaller) BytesFromAccount(opts *bind.CallOpts, account DataTypesUserAccount) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _Rollup.contract.Call(opts, out, "BytesFromAccount", account)
	return *ret0, err
}

// BytesFromAccount is a free data retrieval call binding the contract method 0x3035226f.
//
// Solidity: function BytesFromAccount(DataTypesUserAccount account) constant returns(bytes)
func (_Rollup *RollupSession) BytesFromAccount(account DataTypesUserAccount) ([]byte, error) {
	return _Rollup.Contract.BytesFromAccount(&_Rollup.CallOpts, account)
}

// BytesFromAccount is a free data retrieval call binding the contract method 0x3035226f.
//
// Solidity: function BytesFromAccount(DataTypesUserAccount account) constant returns(bytes)
func (_Rollup *RollupCallerSession) BytesFromAccount(account DataTypesUserAccount) ([]byte, error) {
	return _Rollup.Contract.BytesFromAccount(&_Rollup.CallOpts, account)
}

// BytesFromTx is a free data retrieval call binding the contract method 0x95f70e7c.
//
// Solidity: function BytesFromTx(DataTypesTransaction _tx) constant returns(bytes)
func (_Rollup *RollupCaller) BytesFromTx(opts *bind.CallOpts, _tx DataTypesTransaction) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _Rollup.contract.Call(opts, out, "BytesFromTx", _tx)
	return *ret0, err
}

// BytesFromTx is a free data retrieval call binding the contract method 0x95f70e7c.
//
// Solidity: function BytesFromTx(DataTypesTransaction _tx) constant returns(bytes)
func (_Rollup *RollupSession) BytesFromTx(_tx DataTypesTransaction) ([]byte, error) {
	return _Rollup.Contract.BytesFromTx(&_Rollup.CallOpts, _tx)
}

// BytesFromTx is a free data retrieval call binding the contract method 0x95f70e7c.
//
// Solidity: function BytesFromTx(DataTypesTransaction _tx) constant returns(bytes)
func (_Rollup *RollupCallerSession) BytesFromTx(_tx DataTypesTransaction) ([]byte, error) {
	return _Rollup.Contract.BytesFromTx(&_Rollup.CallOpts, _tx)
}

// HashFromAccount is a free data retrieval call binding the contract method 0xcadbd919.
//
// Solidity: function HashFromAccount(DataTypesUserAccount account) constant returns(bytes32)
func (_Rollup *RollupCaller) HashFromAccount(opts *bind.CallOpts, account DataTypesUserAccount) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Rollup.contract.Call(opts, out, "HashFromAccount", account)
	return *ret0, err
}

// HashFromAccount is a free data retrieval call binding the contract method 0xcadbd919.
//
// Solidity: function HashFromAccount(DataTypesUserAccount account) constant returns(bytes32)
func (_Rollup *RollupSession) HashFromAccount(account DataTypesUserAccount) ([32]byte, error) {
	return _Rollup.Contract.HashFromAccount(&_Rollup.CallOpts, account)
}

// HashFromAccount is a free data retrieval call binding the contract method 0xcadbd919.
//
// Solidity: function HashFromAccount(DataTypesUserAccount account) constant returns(bytes32)
func (_Rollup *RollupCallerSession) HashFromAccount(account DataTypesUserAccount) ([32]byte, error) {
	return _Rollup.Contract.HashFromAccount(&_Rollup.CallOpts, account)
}

// HashFromTx is a free data retrieval call binding the contract method 0x690ac0ec.
//
// Solidity: function HashFromTx(DataTypesTransaction _tx) constant returns(bytes32)
func (_Rollup *RollupCaller) HashFromTx(opts *bind.CallOpts, _tx DataTypesTransaction) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Rollup.contract.Call(opts, out, "HashFromTx", _tx)
	return *ret0, err
}

// HashFromTx is a free data retrieval call binding the contract method 0x690ac0ec.
//
// Solidity: function HashFromTx(DataTypesTransaction _tx) constant returns(bytes32)
func (_Rollup *RollupSession) HashFromTx(_tx DataTypesTransaction) ([32]byte, error) {
	return _Rollup.Contract.HashFromTx(&_Rollup.CallOpts, _tx)
}

// HashFromTx is a free data retrieval call binding the contract method 0x690ac0ec.
//
// Solidity: function HashFromTx(DataTypesTransaction _tx) constant returns(bytes32)
func (_Rollup *RollupCallerSession) HashFromTx(_tx DataTypesTransaction) ([32]byte, error) {
	return _Rollup.Contract.HashFromTx(&_Rollup.CallOpts, _tx)
}

// ZEROBYTES32 is a free data retrieval call binding the contract method 0x069321b0.
//
// Solidity: function ZERO_BYTES32() constant returns(bytes32)
func (_Rollup *RollupCaller) ZEROBYTES32(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Rollup.contract.Call(opts, out, "ZERO_BYTES32")
	return *ret0, err
}

// ZEROBYTES32 is a free data retrieval call binding the contract method 0x069321b0.
//
// Solidity: function ZERO_BYTES32() constant returns(bytes32)
func (_Rollup *RollupSession) ZEROBYTES32() ([32]byte, error) {
	return _Rollup.Contract.ZEROBYTES32(&_Rollup.CallOpts)
}

// ZEROBYTES32 is a free data retrieval call binding the contract method 0x069321b0.
//
// Solidity: function ZERO_BYTES32() constant returns(bytes32)
func (_Rollup *RollupCallerSession) ZEROBYTES32() ([32]byte, error) {
	return _Rollup.Contract.ZEROBYTES32(&_Rollup.CallOpts)
}

// AccountsTree is a free data retrieval call binding the contract method 0xae2926d4.
//
// Solidity: function accountsTree() constant returns(address)
func (_Rollup *RollupCaller) AccountsTree(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Rollup.contract.Call(opts, out, "accountsTree")
	return *ret0, err
}

// AccountsTree is a free data retrieval call binding the contract method 0xae2926d4.
//
// Solidity: function accountsTree() constant returns(address)
func (_Rollup *RollupSession) AccountsTree() (common.Address, error) {
	return _Rollup.Contract.AccountsTree(&_Rollup.CallOpts)
}

// AccountsTree is a free data retrieval call binding the contract method 0xae2926d4.
//
// Solidity: function accountsTree() constant returns(address)
func (_Rollup *RollupCallerSession) AccountsTree() (common.Address, error) {
	return _Rollup.Contract.AccountsTree(&_Rollup.CallOpts)
}

// BalancesTree is a free data retrieval call binding the contract method 0xb2692d77.
//
// Solidity: function balancesTree() constant returns(address)
func (_Rollup *RollupCaller) BalancesTree(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Rollup.contract.Call(opts, out, "balancesTree")
	return *ret0, err
}

// BalancesTree is a free data retrieval call binding the contract method 0xb2692d77.
//
// Solidity: function balancesTree() constant returns(address)
func (_Rollup *RollupSession) BalancesTree() (common.Address, error) {
	return _Rollup.Contract.BalancesTree(&_Rollup.CallOpts)
}

// BalancesTree is a free data retrieval call binding the contract method 0xb2692d77.
//
// Solidity: function balancesTree() constant returns(address)
func (_Rollup *RollupCallerSession) BalancesTree() (common.Address, error) {
	return _Rollup.Contract.BalancesTree(&_Rollup.CallOpts)
}

// Batches is a free data retrieval call binding the contract method 0xb32c4d8d.
//
// Solidity: function batches(uint256 ) constant returns(bytes32 stateRoot, address committer, bytes32 txRoot, uint256 stakeCommitted, uint256 finalisesOn, uint256 timestamp)
func (_Rollup *RollupCaller) Batches(opts *bind.CallOpts, arg0 *big.Int) (struct {
	StateRoot      [32]byte
	Committer      common.Address
	TxRoot         [32]byte
	StakeCommitted *big.Int
	FinalisesOn    *big.Int
	Timestamp      *big.Int
}, error) {
	ret := new(struct {
		StateRoot      [32]byte
		Committer      common.Address
		TxRoot         [32]byte
		StakeCommitted *big.Int
		FinalisesOn    *big.Int
		Timestamp      *big.Int
	})
	out := ret
	err := _Rollup.contract.Call(opts, out, "batches", arg0)
	return *ret, err
}

// Batches is a free data retrieval call binding the contract method 0xb32c4d8d.
//
// Solidity: function batches(uint256 ) constant returns(bytes32 stateRoot, address committer, bytes32 txRoot, uint256 stakeCommitted, uint256 finalisesOn, uint256 timestamp)
func (_Rollup *RollupSession) Batches(arg0 *big.Int) (struct {
	StateRoot      [32]byte
	Committer      common.Address
	TxRoot         [32]byte
	StakeCommitted *big.Int
	FinalisesOn    *big.Int
	Timestamp      *big.Int
}, error) {
	return _Rollup.Contract.Batches(&_Rollup.CallOpts, arg0)
}

// Batches is a free data retrieval call binding the contract method 0xb32c4d8d.
//
// Solidity: function batches(uint256 ) constant returns(bytes32 stateRoot, address committer, bytes32 txRoot, uint256 stakeCommitted, uint256 finalisesOn, uint256 timestamp)
func (_Rollup *RollupCallerSession) Batches(arg0 *big.Int) (struct {
	StateRoot      [32]byte
	Committer      common.Address
	TxRoot         [32]byte
	StakeCommitted *big.Int
	FinalisesOn    *big.Int
	Timestamp      *big.Int
}, error) {
	return _Rollup.Contract.Batches(&_Rollup.CallOpts, arg0)
}

// CalculateAddress is a free data retrieval call binding the contract method 0xe8a4c04e.
//
// Solidity: function calculateAddress(bytes pub) constant returns(address addr)
func (_Rollup *RollupCaller) CalculateAddress(opts *bind.CallOpts, pub []byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Rollup.contract.Call(opts, out, "calculateAddress", pub)
	return *ret0, err
}

// CalculateAddress is a free data retrieval call binding the contract method 0xe8a4c04e.
//
// Solidity: function calculateAddress(bytes pub) constant returns(address addr)
func (_Rollup *RollupSession) CalculateAddress(pub []byte) (common.Address, error) {
	return _Rollup.Contract.CalculateAddress(&_Rollup.CallOpts, pub)
}

// CalculateAddress is a free data retrieval call binding the contract method 0xe8a4c04e.
//
// Solidity: function calculateAddress(bytes pub) constant returns(address addr)
func (_Rollup *RollupCallerSession) CalculateAddress(pub []byte) (common.Address, error) {
	return _Rollup.Contract.CalculateAddress(&_Rollup.CallOpts, pub)
}

// DepositSubtreeHeight is a free data retrieval call binding the contract method 0x2882dd98.
//
// Solidity: function depositSubtreeHeight() constant returns(uint256)
func (_Rollup *RollupCaller) DepositSubtreeHeight(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Rollup.contract.Call(opts, out, "depositSubtreeHeight")
	return *ret0, err
}

// DepositSubtreeHeight is a free data retrieval call binding the contract method 0x2882dd98.
//
// Solidity: function depositSubtreeHeight() constant returns(uint256)
func (_Rollup *RollupSession) DepositSubtreeHeight() (*big.Int, error) {
	return _Rollup.Contract.DepositSubtreeHeight(&_Rollup.CallOpts)
}

// DepositSubtreeHeight is a free data retrieval call binding the contract method 0x2882dd98.
//
// Solidity: function depositSubtreeHeight() constant returns(uint256)
func (_Rollup *RollupCallerSession) DepositSubtreeHeight() (*big.Int, error) {
	return _Rollup.Contract.DepositSubtreeHeight(&_Rollup.CallOpts)
}

// GetBalanceTreeRoot is a free data retrieval call binding the contract method 0x652eb691.
//
// Solidity: function getBalanceTreeRoot() constant returns(bytes32)
func (_Rollup *RollupCaller) GetBalanceTreeRoot(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Rollup.contract.Call(opts, out, "getBalanceTreeRoot")
	return *ret0, err
}

// GetBalanceTreeRoot is a free data retrieval call binding the contract method 0x652eb691.
//
// Solidity: function getBalanceTreeRoot() constant returns(bytes32)
func (_Rollup *RollupSession) GetBalanceTreeRoot() ([32]byte, error) {
	return _Rollup.Contract.GetBalanceTreeRoot(&_Rollup.CallOpts)
}

// GetBalanceTreeRoot is a free data retrieval call binding the contract method 0x652eb691.
//
// Solidity: function getBalanceTreeRoot() constant returns(bytes32)
func (_Rollup *RollupCallerSession) GetBalanceTreeRoot() ([32]byte, error) {
	return _Rollup.Contract.GetBalanceTreeRoot(&_Rollup.CallOpts)
}

// MerkleTreeLib is a free data retrieval call binding the contract method 0x78ad68c6.
//
// Solidity: function merkleTreeLib() constant returns(address)
func (_Rollup *RollupCaller) MerkleTreeLib(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Rollup.contract.Call(opts, out, "merkleTreeLib")
	return *ret0, err
}

// MerkleTreeLib is a free data retrieval call binding the contract method 0x78ad68c6.
//
// Solidity: function merkleTreeLib() constant returns(address)
func (_Rollup *RollupSession) MerkleTreeLib() (common.Address, error) {
	return _Rollup.Contract.MerkleTreeLib(&_Rollup.CallOpts)
}

// MerkleTreeLib is a free data retrieval call binding the contract method 0x78ad68c6.
//
// Solidity: function merkleTreeLib() constant returns(address)
func (_Rollup *RollupCallerSession) MerkleTreeLib() (common.Address, error) {
	return _Rollup.Contract.MerkleTreeLib(&_Rollup.CallOpts)
}

// NumberOfBatches is a free data retrieval call binding the contract method 0x88c46fb8.
//
// Solidity: function numberOfBatches() constant returns(uint256)
func (_Rollup *RollupCaller) NumberOfBatches(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Rollup.contract.Call(opts, out, "numberOfBatches")
	return *ret0, err
}

// NumberOfBatches is a free data retrieval call binding the contract method 0x88c46fb8.
//
// Solidity: function numberOfBatches() constant returns(uint256)
func (_Rollup *RollupSession) NumberOfBatches() (*big.Int, error) {
	return _Rollup.Contract.NumberOfBatches(&_Rollup.CallOpts)
}

// NumberOfBatches is a free data retrieval call binding the contract method 0x88c46fb8.
//
// Solidity: function numberOfBatches() constant returns(uint256)
func (_Rollup *RollupCallerSession) NumberOfBatches() (*big.Int, error) {
	return _Rollup.Contract.NumberOfBatches(&_Rollup.CallOpts)
}

// PendingDeposits is a free data retrieval call binding the contract method 0xa7932794.
//
// Solidity: function pendingDeposits(uint256 ) constant returns(bytes32)
func (_Rollup *RollupCaller) PendingDeposits(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Rollup.contract.Call(opts, out, "pendingDeposits", arg0)
	return *ret0, err
}

// PendingDeposits is a free data retrieval call binding the contract method 0xa7932794.
//
// Solidity: function pendingDeposits(uint256 ) constant returns(bytes32)
func (_Rollup *RollupSession) PendingDeposits(arg0 *big.Int) ([32]byte, error) {
	return _Rollup.Contract.PendingDeposits(&_Rollup.CallOpts, arg0)
}

// PendingDeposits is a free data retrieval call binding the contract method 0xa7932794.
//
// Solidity: function pendingDeposits(uint256 ) constant returns(bytes32)
func (_Rollup *RollupCallerSession) PendingDeposits(arg0 *big.Int) ([32]byte, error) {
	return _Rollup.Contract.PendingDeposits(&_Rollup.CallOpts, arg0)
}

// QueueNumber is a free data retrieval call binding the contract method 0x2fa6779a.
//
// Solidity: function queueNumber() constant returns(uint256)
func (_Rollup *RollupCaller) QueueNumber(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Rollup.contract.Call(opts, out, "queueNumber")
	return *ret0, err
}

// QueueNumber is a free data retrieval call binding the contract method 0x2fa6779a.
//
// Solidity: function queueNumber() constant returns(uint256)
func (_Rollup *RollupSession) QueueNumber() (*big.Int, error) {
	return _Rollup.Contract.QueueNumber(&_Rollup.CallOpts)
}

// QueueNumber is a free data retrieval call binding the contract method 0x2fa6779a.
//
// Solidity: function queueNumber() constant returns(uint256)
func (_Rollup *RollupCallerSession) QueueNumber() (*big.Int, error) {
	return _Rollup.Contract.QueueNumber(&_Rollup.CallOpts)
}

// TokenContract is a free data retrieval call binding the contract method 0x55a373d6.
//
// Solidity: function tokenContract() constant returns(address)
func (_Rollup *RollupCaller) TokenContract(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Rollup.contract.Call(opts, out, "tokenContract")
	return *ret0, err
}

// TokenContract is a free data retrieval call binding the contract method 0x55a373d6.
//
// Solidity: function tokenContract() constant returns(address)
func (_Rollup *RollupSession) TokenContract() (common.Address, error) {
	return _Rollup.Contract.TokenContract(&_Rollup.CallOpts)
}

// TokenContract is a free data retrieval call binding the contract method 0x55a373d6.
//
// Solidity: function tokenContract() constant returns(address)
func (_Rollup *RollupCallerSession) TokenContract() (common.Address, error) {
	return _Rollup.Contract.TokenContract(&_Rollup.CallOpts)
}

// TokenRegistry is a free data retrieval call binding the contract method 0x9d23c4c7.
//
// Solidity: function tokenRegistry() constant returns(address)
func (_Rollup *RollupCaller) TokenRegistry(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Rollup.contract.Call(opts, out, "tokenRegistry")
	return *ret0, err
}

// TokenRegistry is a free data retrieval call binding the contract method 0x9d23c4c7.
//
// Solidity: function tokenRegistry() constant returns(address)
func (_Rollup *RollupSession) TokenRegistry() (common.Address, error) {
	return _Rollup.Contract.TokenRegistry(&_Rollup.CallOpts)
}

// TokenRegistry is a free data retrieval call binding the contract method 0x9d23c4c7.
//
// Solidity: function tokenRegistry() constant returns(address)
func (_Rollup *RollupCallerSession) TokenRegistry() (common.Address, error) {
	return _Rollup.Contract.TokenRegistry(&_Rollup.CallOpts)
}

// SlashAndRollback is a paid mutator transaction binding the contract method 0xdf070983.
//
// Solidity: function SlashAndRollback() returns()
func (_Rollup *RollupTransactor) SlashAndRollback(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "SlashAndRollback")
}

// SlashAndRollback is a paid mutator transaction binding the contract method 0xdf070983.
//
// Solidity: function SlashAndRollback() returns()
func (_Rollup *RollupSession) SlashAndRollback() (*types.Transaction, error) {
	return _Rollup.Contract.SlashAndRollback(&_Rollup.TransactOpts)
}

// SlashAndRollback is a paid mutator transaction binding the contract method 0xdf070983.
//
// Solidity: function SlashAndRollback() returns()
func (_Rollup *RollupTransactorSession) SlashAndRollback() (*types.Transaction, error) {
	return _Rollup.Contract.SlashAndRollback(&_Rollup.TransactOpts)
}

// UpdateBalanceInAccount is a paid mutator transaction binding the contract method 0x2e9be49a.
//
// Solidity: function UpdateBalanceInAccount(DataTypesUserAccount original_account, uint256 new_balance) returns(DataTypesUserAccount new_account)
func (_Rollup *RollupTransactor) UpdateBalanceInAccount(opts *bind.TransactOpts, original_account DataTypesUserAccount, new_balance *big.Int) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "UpdateBalanceInAccount", original_account, new_balance)
}

// UpdateBalanceInAccount is a paid mutator transaction binding the contract method 0x2e9be49a.
//
// Solidity: function UpdateBalanceInAccount(DataTypesUserAccount original_account, uint256 new_balance) returns(DataTypesUserAccount new_account)
func (_Rollup *RollupSession) UpdateBalanceInAccount(original_account DataTypesUserAccount, new_balance *big.Int) (*types.Transaction, error) {
	return _Rollup.Contract.UpdateBalanceInAccount(&_Rollup.TransactOpts, original_account, new_balance)
}

// UpdateBalanceInAccount is a paid mutator transaction binding the contract method 0x2e9be49a.
//
// Solidity: function UpdateBalanceInAccount(DataTypesUserAccount original_account, uint256 new_balance) returns(DataTypesUserAccount new_account)
func (_Rollup *RollupTransactorSession) UpdateBalanceInAccount(original_account DataTypesUserAccount, new_balance *big.Int) (*types.Transaction, error) {
	return _Rollup.Contract.UpdateBalanceInAccount(&_Rollup.TransactOpts, original_account, new_balance)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xff34585d.
//
// Solidity: function WithdrawStake(uint256 batch_id) returns()
func (_Rollup *RollupTransactor) WithdrawStake(opts *bind.TransactOpts, batch_id *big.Int) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "WithdrawStake", batch_id)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xff34585d.
//
// Solidity: function WithdrawStake(uint256 batch_id) returns()
func (_Rollup *RollupSession) WithdrawStake(batch_id *big.Int) (*types.Transaction, error) {
	return _Rollup.Contract.WithdrawStake(&_Rollup.TransactOpts, batch_id)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xff34585d.
//
// Solidity: function WithdrawStake(uint256 batch_id) returns()
func (_Rollup *RollupTransactorSession) WithdrawStake(batch_id *big.Int) (*types.Transaction, error) {
	return _Rollup.Contract.WithdrawStake(&_Rollup.TransactOpts, batch_id)
}

// Deposit is a paid mutator transaction binding the contract method 0xaa0b7db7.
//
// Solidity: function deposit(uint256 _amount, uint256 _tokenType, bytes _pubkey) returns()
func (_Rollup *RollupTransactor) Deposit(opts *bind.TransactOpts, _amount *big.Int, _tokenType *big.Int, _pubkey []byte) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "deposit", _amount, _tokenType, _pubkey)
}

// Deposit is a paid mutator transaction binding the contract method 0xaa0b7db7.
//
// Solidity: function deposit(uint256 _amount, uint256 _tokenType, bytes _pubkey) returns()
func (_Rollup *RollupSession) Deposit(_amount *big.Int, _tokenType *big.Int, _pubkey []byte) (*types.Transaction, error) {
	return _Rollup.Contract.Deposit(&_Rollup.TransactOpts, _amount, _tokenType, _pubkey)
}

// Deposit is a paid mutator transaction binding the contract method 0xaa0b7db7.
//
// Solidity: function deposit(uint256 _amount, uint256 _tokenType, bytes _pubkey) returns()
func (_Rollup *RollupTransactorSession) Deposit(_amount *big.Int, _tokenType *big.Int, _pubkey []byte) (*types.Transaction, error) {
	return _Rollup.Contract.Deposit(&_Rollup.TransactOpts, _amount, _tokenType, _pubkey)
}

// DepositFor is a paid mutator transaction binding the contract method 0x1f8eb3e8.
//
// Solidity: function depositFor(address _destination, uint256 _amount, uint256 _tokenType, bytes _pubkey) returns()
func (_Rollup *RollupTransactor) DepositFor(opts *bind.TransactOpts, _destination common.Address, _amount *big.Int, _tokenType *big.Int, _pubkey []byte) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "depositFor", _destination, _amount, _tokenType, _pubkey)
}

// DepositFor is a paid mutator transaction binding the contract method 0x1f8eb3e8.
//
// Solidity: function depositFor(address _destination, uint256 _amount, uint256 _tokenType, bytes _pubkey) returns()
func (_Rollup *RollupSession) DepositFor(_destination common.Address, _amount *big.Int, _tokenType *big.Int, _pubkey []byte) (*types.Transaction, error) {
	return _Rollup.Contract.DepositFor(&_Rollup.TransactOpts, _destination, _amount, _tokenType, _pubkey)
}

// DepositFor is a paid mutator transaction binding the contract method 0x1f8eb3e8.
//
// Solidity: function depositFor(address _destination, uint256 _amount, uint256 _tokenType, bytes _pubkey) returns()
func (_Rollup *RollupTransactorSession) DepositFor(_destination common.Address, _amount *big.Int, _tokenType *big.Int, _pubkey []byte) (*types.Transaction, error) {
	return _Rollup.Contract.DepositFor(&_Rollup.TransactOpts, _destination, _amount, _tokenType, _pubkey)
}

// DisputeBatch is a paid mutator transaction binding the contract method 0x36050974.
//
// Solidity: function disputeBatch(uint256 _batch_id, []DataTypesTransaction _txs, []DataTypesAccountMerkleProof _from_proofs, []DataTypesAccountMerkleProof _to_proofs) returns()
func (_Rollup *RollupTransactor) DisputeBatch(opts *bind.TransactOpts, _batch_id *big.Int, _txs []DataTypesTransaction, _from_proofs []DataTypesAccountMerkleProof, _to_proofs []DataTypesAccountMerkleProof) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "disputeBatch", _batch_id, _txs, _from_proofs, _to_proofs)
}

// DisputeBatch is a paid mutator transaction binding the contract method 0x36050974.
//
// Solidity: function disputeBatch(uint256 _batch_id, []DataTypesTransaction _txs, []DataTypesAccountMerkleProof _from_proofs, []DataTypesAccountMerkleProof _to_proofs) returns()
func (_Rollup *RollupSession) DisputeBatch(_batch_id *big.Int, _txs []DataTypesTransaction, _from_proofs []DataTypesAccountMerkleProof, _to_proofs []DataTypesAccountMerkleProof) (*types.Transaction, error) {
	return _Rollup.Contract.DisputeBatch(&_Rollup.TransactOpts, _batch_id, _txs, _from_proofs, _to_proofs)
}

// DisputeBatch is a paid mutator transaction binding the contract method 0x36050974.
//
// Solidity: function disputeBatch(uint256 _batch_id, []DataTypesTransaction _txs, []DataTypesAccountMerkleProof _from_proofs, []DataTypesAccountMerkleProof _to_proofs) returns()
func (_Rollup *RollupTransactorSession) DisputeBatch(_batch_id *big.Int, _txs []DataTypesTransaction, _from_proofs []DataTypesAccountMerkleProof, _to_proofs []DataTypesAccountMerkleProof) (*types.Transaction, error) {
	return _Rollup.Contract.DisputeBatch(&_Rollup.TransactOpts, _batch_id, _txs, _from_proofs, _to_proofs)
}

// FinaliseDeposits is a paid mutator transaction binding the contract method 0x1450e6ec.
//
// Solidity: function finaliseDeposits(uint256 _subTreeDepth, DataTypesAccountMerkleProof _zero_account_mp) returns(bytes32)
func (_Rollup *RollupTransactor) FinaliseDeposits(opts *bind.TransactOpts, _subTreeDepth *big.Int, _zero_account_mp DataTypesAccountMerkleProof) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "finaliseDeposits", _subTreeDepth, _zero_account_mp)
}

// FinaliseDeposits is a paid mutator transaction binding the contract method 0x1450e6ec.
//
// Solidity: function finaliseDeposits(uint256 _subTreeDepth, DataTypesAccountMerkleProof _zero_account_mp) returns(bytes32)
func (_Rollup *RollupSession) FinaliseDeposits(_subTreeDepth *big.Int, _zero_account_mp DataTypesAccountMerkleProof) (*types.Transaction, error) {
	return _Rollup.Contract.FinaliseDeposits(&_Rollup.TransactOpts, _subTreeDepth, _zero_account_mp)
}

// FinaliseDeposits is a paid mutator transaction binding the contract method 0x1450e6ec.
//
// Solidity: function finaliseDeposits(uint256 _subTreeDepth, DataTypesAccountMerkleProof _zero_account_mp) returns(bytes32)
func (_Rollup *RollupTransactorSession) FinaliseDeposits(_subTreeDepth *big.Int, _zero_account_mp DataTypesAccountMerkleProof) (*types.Transaction, error) {
	return _Rollup.Contract.FinaliseDeposits(&_Rollup.TransactOpts, _subTreeDepth, _zero_account_mp)
}

// FinaliseTokenRegistration is a paid mutator transaction binding the contract method 0x50de1b0f.
//
// Solidity: function finaliseTokenRegistration(address _tokenContractAddress) returns()
func (_Rollup *RollupTransactor) FinaliseTokenRegistration(opts *bind.TransactOpts, _tokenContractAddress common.Address) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "finaliseTokenRegistration", _tokenContractAddress)
}

// FinaliseTokenRegistration is a paid mutator transaction binding the contract method 0x50de1b0f.
//
// Solidity: function finaliseTokenRegistration(address _tokenContractAddress) returns()
func (_Rollup *RollupSession) FinaliseTokenRegistration(_tokenContractAddress common.Address) (*types.Transaction, error) {
	return _Rollup.Contract.FinaliseTokenRegistration(&_Rollup.TransactOpts, _tokenContractAddress)
}

// FinaliseTokenRegistration is a paid mutator transaction binding the contract method 0x50de1b0f.
//
// Solidity: function finaliseTokenRegistration(address _tokenContractAddress) returns()
func (_Rollup *RollupTransactorSession) FinaliseTokenRegistration(_tokenContractAddress common.Address) (*types.Transaction, error) {
	return _Rollup.Contract.FinaliseTokenRegistration(&_Rollup.TransactOpts, _tokenContractAddress)
}

// GetDepositsHash is a paid mutator transaction binding the contract method 0x31133280.
//
// Solidity: function getDepositsHash(bytes32 a, bytes32 b) returns(bytes32)
func (_Rollup *RollupTransactor) GetDepositsHash(opts *bind.TransactOpts, a [32]byte, b [32]byte) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "getDepositsHash", a, b)
}

// GetDepositsHash is a paid mutator transaction binding the contract method 0x31133280.
//
// Solidity: function getDepositsHash(bytes32 a, bytes32 b) returns(bytes32)
func (_Rollup *RollupSession) GetDepositsHash(a [32]byte, b [32]byte) (*types.Transaction, error) {
	return _Rollup.Contract.GetDepositsHash(&_Rollup.TransactOpts, a, b)
}

// GetDepositsHash is a paid mutator transaction binding the contract method 0x31133280.
//
// Solidity: function getDepositsHash(bytes32 a, bytes32 b) returns(bytes32)
func (_Rollup *RollupTransactorSession) GetDepositsHash(a [32]byte, b [32]byte) (*types.Transaction, error) {
	return _Rollup.Contract.GetDepositsHash(&_Rollup.TransactOpts, a, b)
}

// ProcessTx is a paid mutator transaction binding the contract method 0xfb372f70.
//
// Solidity: function processTx(bytes32 _balanceRoot, DataTypesTransaction _tx, DataTypesAccountMerkleProof _from_merkle_proof, DataTypesAccountMerkleProof _to_merkle_proof) returns(bytes32, uint256, uint256, bool)
func (_Rollup *RollupTransactor) ProcessTx(opts *bind.TransactOpts, _balanceRoot [32]byte, _tx DataTypesTransaction, _from_merkle_proof DataTypesAccountMerkleProof, _to_merkle_proof DataTypesAccountMerkleProof) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "processTx", _balanceRoot, _tx, _from_merkle_proof, _to_merkle_proof)
}

// ProcessTx is a paid mutator transaction binding the contract method 0xfb372f70.
//
// Solidity: function processTx(bytes32 _balanceRoot, DataTypesTransaction _tx, DataTypesAccountMerkleProof _from_merkle_proof, DataTypesAccountMerkleProof _to_merkle_proof) returns(bytes32, uint256, uint256, bool)
func (_Rollup *RollupSession) ProcessTx(_balanceRoot [32]byte, _tx DataTypesTransaction, _from_merkle_proof DataTypesAccountMerkleProof, _to_merkle_proof DataTypesAccountMerkleProof) (*types.Transaction, error) {
	return _Rollup.Contract.ProcessTx(&_Rollup.TransactOpts, _balanceRoot, _tx, _from_merkle_proof, _to_merkle_proof)
}

// ProcessTx is a paid mutator transaction binding the contract method 0xfb372f70.
//
// Solidity: function processTx(bytes32 _balanceRoot, DataTypesTransaction _tx, DataTypesAccountMerkleProof _from_merkle_proof, DataTypesAccountMerkleProof _to_merkle_proof) returns(bytes32, uint256, uint256, bool)
func (_Rollup *RollupTransactorSession) ProcessTx(_balanceRoot [32]byte, _tx DataTypesTransaction, _from_merkle_proof DataTypesAccountMerkleProof, _to_merkle_proof DataTypesAccountMerkleProof) (*types.Transaction, error) {
	return _Rollup.Contract.ProcessTx(&_Rollup.TransactOpts, _balanceRoot, _tx, _from_merkle_proof, _to_merkle_proof)
}

// RequestTokenRegistration is a paid mutator transaction binding the contract method 0x55857ed9.
//
// Solidity: function requestTokenRegistration(address _tokenContractAddress) returns()
func (_Rollup *RollupTransactor) RequestTokenRegistration(opts *bind.TransactOpts, _tokenContractAddress common.Address) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "requestTokenRegistration", _tokenContractAddress)
}

// RequestTokenRegistration is a paid mutator transaction binding the contract method 0x55857ed9.
//
// Solidity: function requestTokenRegistration(address _tokenContractAddress) returns()
func (_Rollup *RollupSession) RequestTokenRegistration(_tokenContractAddress common.Address) (*types.Transaction, error) {
	return _Rollup.Contract.RequestTokenRegistration(&_Rollup.TransactOpts, _tokenContractAddress)
}

// RequestTokenRegistration is a paid mutator transaction binding the contract method 0x55857ed9.
//
// Solidity: function requestTokenRegistration(address _tokenContractAddress) returns()
func (_Rollup *RollupTransactorSession) RequestTokenRegistration(_tokenContractAddress common.Address) (*types.Transaction, error) {
	return _Rollup.Contract.RequestTokenRegistration(&_Rollup.TransactOpts, _tokenContractAddress)
}

// SubmitBatch is a paid mutator transaction binding the contract method 0x0e981757.
//
// Solidity: function submitBatch(bytes[] _txs, bytes32 _updatedRoot) returns()
func (_Rollup *RollupTransactor) SubmitBatch(opts *bind.TransactOpts, _txs [][]byte, _updatedRoot [32]byte) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "submitBatch", _txs, _updatedRoot)
}

// SubmitBatch is a paid mutator transaction binding the contract method 0x0e981757.
//
// Solidity: function submitBatch(bytes[] _txs, bytes32 _updatedRoot) returns()
func (_Rollup *RollupSession) SubmitBatch(_txs [][]byte, _updatedRoot [32]byte) (*types.Transaction, error) {
	return _Rollup.Contract.SubmitBatch(&_Rollup.TransactOpts, _txs, _updatedRoot)
}

// SubmitBatch is a paid mutator transaction binding the contract method 0x0e981757.
//
// Solidity: function submitBatch(bytes[] _txs, bytes32 _updatedRoot) returns()
func (_Rollup *RollupTransactorSession) SubmitBatch(_txs [][]byte, _updatedRoot [32]byte) (*types.Transaction, error) {
	return _Rollup.Contract.SubmitBatch(&_Rollup.TransactOpts, _txs, _updatedRoot)
}

// RollupBatchRollbackIterator is returned from FilterBatchRollback and is used to iterate over the raw logs and unpacked data for BatchRollback events raised by the Rollup contract.
type RollupBatchRollbackIterator struct {
	Event *RollupBatchRollback // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RollupBatchRollbackIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupBatchRollback)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RollupBatchRollback)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RollupBatchRollbackIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupBatchRollbackIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupBatchRollback represents a BatchRollback event raised by the Rollup contract.
type RollupBatchRollback struct {
	BatchId      *big.Int
	Committer    common.Address
	StateRoot    [32]byte
	TxRoot       [32]byte
	StakeSlashed *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterBatchRollback is a free log retrieval operation binding the contract event 0xff0d01a3ea09eec8d9dc11b606f70fbf5aa373d5647db792e5886fc4285c38fc.
//
// Solidity: event BatchRollback(uint256 batch_id, address committer, bytes32 stateRoot, bytes32 txRoot, uint256 stakeSlashed)
func (_Rollup *RollupFilterer) FilterBatchRollback(opts *bind.FilterOpts) (*RollupBatchRollbackIterator, error) {

	logs, sub, err := _Rollup.contract.FilterLogs(opts, "BatchRollback")
	if err != nil {
		return nil, err
	}
	return &RollupBatchRollbackIterator{contract: _Rollup.contract, event: "BatchRollback", logs: logs, sub: sub}, nil
}

// WatchBatchRollback is a free log subscription operation binding the contract event 0xff0d01a3ea09eec8d9dc11b606f70fbf5aa373d5647db792e5886fc4285c38fc.
//
// Solidity: event BatchRollback(uint256 batch_id, address committer, bytes32 stateRoot, bytes32 txRoot, uint256 stakeSlashed)
func (_Rollup *RollupFilterer) WatchBatchRollback(opts *bind.WatchOpts, sink chan<- *RollupBatchRollback) (event.Subscription, error) {

	logs, sub, err := _Rollup.contract.WatchLogs(opts, "BatchRollback")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupBatchRollback)
				if err := _Rollup.contract.UnpackLog(event, "BatchRollback", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBatchRollback is a log parse operation binding the contract event 0xff0d01a3ea09eec8d9dc11b606f70fbf5aa373d5647db792e5886fc4285c38fc.
//
// Solidity: event BatchRollback(uint256 batch_id, address committer, bytes32 stateRoot, bytes32 txRoot, uint256 stakeSlashed)
func (_Rollup *RollupFilterer) ParseBatchRollback(log types.Log) (*RollupBatchRollback, error) {
	event := new(RollupBatchRollback)
	if err := _Rollup.contract.UnpackLog(event, "BatchRollback", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RollupDepositLeafMergedIterator is returned from FilterDepositLeafMerged and is used to iterate over the raw logs and unpacked data for DepositLeafMerged events raised by the Rollup contract.
type RollupDepositLeafMergedIterator struct {
	Event *RollupDepositLeafMerged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RollupDepositLeafMergedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupDepositLeafMerged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RollupDepositLeafMerged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RollupDepositLeafMergedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupDepositLeafMergedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupDepositLeafMerged represents a DepositLeafMerged event raised by the Rollup contract.
type RollupDepositLeafMerged struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDepositLeafMerged is a free log retrieval operation binding the contract event 0xe08d4467ae1c0717bba6e77edc5bc5394ea260848bee1f763aba623e6e1dc864.
//
// Solidity: event DepositLeafMerged()
func (_Rollup *RollupFilterer) FilterDepositLeafMerged(opts *bind.FilterOpts) (*RollupDepositLeafMergedIterator, error) {

	logs, sub, err := _Rollup.contract.FilterLogs(opts, "DepositLeafMerged")
	if err != nil {
		return nil, err
	}
	return &RollupDepositLeafMergedIterator{contract: _Rollup.contract, event: "DepositLeafMerged", logs: logs, sub: sub}, nil
}

// WatchDepositLeafMerged is a free log subscription operation binding the contract event 0xe08d4467ae1c0717bba6e77edc5bc5394ea260848bee1f763aba623e6e1dc864.
//
// Solidity: event DepositLeafMerged()
func (_Rollup *RollupFilterer) WatchDepositLeafMerged(opts *bind.WatchOpts, sink chan<- *RollupDepositLeafMerged) (event.Subscription, error) {

	logs, sub, err := _Rollup.contract.WatchLogs(opts, "DepositLeafMerged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupDepositLeafMerged)
				if err := _Rollup.contract.UnpackLog(event, "DepositLeafMerged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDepositLeafMerged is a log parse operation binding the contract event 0xe08d4467ae1c0717bba6e77edc5bc5394ea260848bee1f763aba623e6e1dc864.
//
// Solidity: event DepositLeafMerged()
func (_Rollup *RollupFilterer) ParseDepositLeafMerged(log types.Log) (*RollupDepositLeafMerged, error) {
	event := new(RollupDepositLeafMerged)
	if err := _Rollup.contract.UnpackLog(event, "DepositLeafMerged", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RollupDepositQueuedIterator is returned from FilterDepositQueued and is used to iterate over the raw logs and unpacked data for DepositQueued events raised by the Rollup contract.
type RollupDepositQueuedIterator struct {
	Event *RollupDepositQueued // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RollupDepositQueuedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupDepositQueued)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RollupDepositQueued)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RollupDepositQueuedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupDepositQueuedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupDepositQueued represents a DepositQueued event raised by the Rollup contract.
type RollupDepositQueued struct {
	Arg0 common.Address
	Arg1 *big.Int
	Arg2 *big.Int
	Arg3 [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterDepositQueued is a free log retrieval operation binding the contract event 0x5a835bed28b01b1737787e74965ded7f46cbbeb76095bd80ca6375a5e3c15a0f.
//
// Solidity: event DepositQueued(address , uint256 , uint256 , bytes32 )
func (_Rollup *RollupFilterer) FilterDepositQueued(opts *bind.FilterOpts) (*RollupDepositQueuedIterator, error) {

	logs, sub, err := _Rollup.contract.FilterLogs(opts, "DepositQueued")
	if err != nil {
		return nil, err
	}
	return &RollupDepositQueuedIterator{contract: _Rollup.contract, event: "DepositQueued", logs: logs, sub: sub}, nil
}

// WatchDepositQueued is a free log subscription operation binding the contract event 0x5a835bed28b01b1737787e74965ded7f46cbbeb76095bd80ca6375a5e3c15a0f.
//
// Solidity: event DepositQueued(address , uint256 , uint256 , bytes32 )
func (_Rollup *RollupFilterer) WatchDepositQueued(opts *bind.WatchOpts, sink chan<- *RollupDepositQueued) (event.Subscription, error) {

	logs, sub, err := _Rollup.contract.WatchLogs(opts, "DepositQueued")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupDepositQueued)
				if err := _Rollup.contract.UnpackLog(event, "DepositQueued", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDepositQueued is a log parse operation binding the contract event 0x5a835bed28b01b1737787e74965ded7f46cbbeb76095bd80ca6375a5e3c15a0f.
//
// Solidity: event DepositQueued(address , uint256 , uint256 , bytes32 )
func (_Rollup *RollupFilterer) ParseDepositQueued(log types.Log) (*RollupDepositQueued, error) {
	event := new(RollupDepositQueued)
	if err := _Rollup.contract.UnpackLog(event, "DepositQueued", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RollupDepositsProcessedIterator is returned from FilterDepositsProcessed and is used to iterate over the raw logs and unpacked data for DepositsProcessed events raised by the Rollup contract.
type RollupDepositsProcessedIterator struct {
	Event *RollupDepositsProcessed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RollupDepositsProcessedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupDepositsProcessed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RollupDepositsProcessed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RollupDepositsProcessedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupDepositsProcessedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupDepositsProcessed represents a DepositsProcessed event raised by the Rollup contract.
type RollupDepositsProcessed struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDepositsProcessed is a free log retrieval operation binding the contract event 0x663ea3b3fa88c877f70d364fe14a3b85f410a60df236737b37fcc0796225fe33.
//
// Solidity: event DepositsProcessed()
func (_Rollup *RollupFilterer) FilterDepositsProcessed(opts *bind.FilterOpts) (*RollupDepositsProcessedIterator, error) {

	logs, sub, err := _Rollup.contract.FilterLogs(opts, "DepositsProcessed")
	if err != nil {
		return nil, err
	}
	return &RollupDepositsProcessedIterator{contract: _Rollup.contract, event: "DepositsProcessed", logs: logs, sub: sub}, nil
}

// WatchDepositsProcessed is a free log subscription operation binding the contract event 0x663ea3b3fa88c877f70d364fe14a3b85f410a60df236737b37fcc0796225fe33.
//
// Solidity: event DepositsProcessed()
func (_Rollup *RollupFilterer) WatchDepositsProcessed(opts *bind.WatchOpts, sink chan<- *RollupDepositsProcessed) (event.Subscription, error) {

	logs, sub, err := _Rollup.contract.WatchLogs(opts, "DepositsProcessed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupDepositsProcessed)
				if err := _Rollup.contract.UnpackLog(event, "DepositsProcessed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDepositsProcessed is a log parse operation binding the contract event 0x663ea3b3fa88c877f70d364fe14a3b85f410a60df236737b37fcc0796225fe33.
//
// Solidity: event DepositsProcessed()
func (_Rollup *RollupFilterer) ParseDepositsProcessed(log types.Log) (*RollupDepositsProcessed, error) {
	event := new(RollupDepositsProcessed)
	if err := _Rollup.contract.UnpackLog(event, "DepositsProcessed", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RollupNewAccountIterator is returned from FilterNewAccount and is used to iterate over the raw logs and unpacked data for NewAccount events raised by the Rollup contract.
type RollupNewAccountIterator struct {
	Event *RollupNewAccount // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RollupNewAccountIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupNewAccount)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RollupNewAccount)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RollupNewAccountIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupNewAccountIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupNewAccount represents a NewAccount event raised by the Rollup contract.
type RollupNewAccount struct {
	Root  [32]byte
	Index *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterNewAccount is a free log retrieval operation binding the contract event 0x8b7959e43d6da6de7ea88053c29ed23dc7c9f55e140402407c430152562c14b3.
//
// Solidity: event NewAccount(bytes32 root, uint256 index)
func (_Rollup *RollupFilterer) FilterNewAccount(opts *bind.FilterOpts) (*RollupNewAccountIterator, error) {

	logs, sub, err := _Rollup.contract.FilterLogs(opts, "NewAccount")
	if err != nil {
		return nil, err
	}
	return &RollupNewAccountIterator{contract: _Rollup.contract, event: "NewAccount", logs: logs, sub: sub}, nil
}

// WatchNewAccount is a free log subscription operation binding the contract event 0x8b7959e43d6da6de7ea88053c29ed23dc7c9f55e140402407c430152562c14b3.
//
// Solidity: event NewAccount(bytes32 root, uint256 index)
func (_Rollup *RollupFilterer) WatchNewAccount(opts *bind.WatchOpts, sink chan<- *RollupNewAccount) (event.Subscription, error) {

	logs, sub, err := _Rollup.contract.WatchLogs(opts, "NewAccount")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupNewAccount)
				if err := _Rollup.contract.UnpackLog(event, "NewAccount", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewAccount is a log parse operation binding the contract event 0x8b7959e43d6da6de7ea88053c29ed23dc7c9f55e140402407c430152562c14b3.
//
// Solidity: event NewAccount(bytes32 root, uint256 index)
func (_Rollup *RollupFilterer) ParseNewAccount(log types.Log) (*RollupNewAccount, error) {
	event := new(RollupNewAccount)
	if err := _Rollup.contract.UnpackLog(event, "NewAccount", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RollupNewBatchIterator is returned from FilterNewBatch and is used to iterate over the raw logs and unpacked data for NewBatch events raised by the Rollup contract.
type RollupNewBatchIterator struct {
	Event *RollupNewBatch // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RollupNewBatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupNewBatch)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RollupNewBatch)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RollupNewBatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupNewBatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupNewBatch represents a NewBatch event raised by the Rollup contract.
type RollupNewBatch struct {
	Committer   common.Address
	Txroot      [32]byte
	UpdatedRoot [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterNewBatch is a free log retrieval operation binding the contract event 0x308f524b10e621c7661f602d9a43157983896901b7ef7f53b30c7f1572c8ff0d.
//
// Solidity: event NewBatch(address committer, bytes32 txroot, bytes32 updatedRoot)
func (_Rollup *RollupFilterer) FilterNewBatch(opts *bind.FilterOpts) (*RollupNewBatchIterator, error) {

	logs, sub, err := _Rollup.contract.FilterLogs(opts, "NewBatch")
	if err != nil {
		return nil, err
	}
	return &RollupNewBatchIterator{contract: _Rollup.contract, event: "NewBatch", logs: logs, sub: sub}, nil
}

// WatchNewBatch is a free log subscription operation binding the contract event 0x308f524b10e621c7661f602d9a43157983896901b7ef7f53b30c7f1572c8ff0d.
//
// Solidity: event NewBatch(address committer, bytes32 txroot, bytes32 updatedRoot)
func (_Rollup *RollupFilterer) WatchNewBatch(opts *bind.WatchOpts, sink chan<- *RollupNewBatch) (event.Subscription, error) {

	logs, sub, err := _Rollup.contract.WatchLogs(opts, "NewBatch")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupNewBatch)
				if err := _Rollup.contract.UnpackLog(event, "NewBatch", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewBatch is a log parse operation binding the contract event 0x308f524b10e621c7661f602d9a43157983896901b7ef7f53b30c7f1572c8ff0d.
//
// Solidity: event NewBatch(address committer, bytes32 txroot, bytes32 updatedRoot)
func (_Rollup *RollupFilterer) ParseNewBatch(log types.Log) (*RollupNewBatch, error) {
	event := new(RollupNewBatch)
	if err := _Rollup.contract.UnpackLog(event, "NewBatch", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RollupRegisteredTokenIterator is returned from FilterRegisteredToken and is used to iterate over the raw logs and unpacked data for RegisteredToken events raised by the Rollup contract.
type RollupRegisteredTokenIterator struct {
	Event *RollupRegisteredToken // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RollupRegisteredTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupRegisteredToken)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RollupRegisteredToken)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RollupRegisteredTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupRegisteredTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupRegisteredToken represents a RegisteredToken event raised by the Rollup contract.
type RollupRegisteredToken struct {
	TokenType     *big.Int
	TokenContract common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterRegisteredToken is a free log retrieval operation binding the contract event 0x5dbaa701a7acef513f72a61799f7e50f4653f462b9f780d88d1b9bec89de2168.
//
// Solidity: event RegisteredToken(uint256 tokenType, address tokenContract)
func (_Rollup *RollupFilterer) FilterRegisteredToken(opts *bind.FilterOpts) (*RollupRegisteredTokenIterator, error) {

	logs, sub, err := _Rollup.contract.FilterLogs(opts, "RegisteredToken")
	if err != nil {
		return nil, err
	}
	return &RollupRegisteredTokenIterator{contract: _Rollup.contract, event: "RegisteredToken", logs: logs, sub: sub}, nil
}

// WatchRegisteredToken is a free log subscription operation binding the contract event 0x5dbaa701a7acef513f72a61799f7e50f4653f462b9f780d88d1b9bec89de2168.
//
// Solidity: event RegisteredToken(uint256 tokenType, address tokenContract)
func (_Rollup *RollupFilterer) WatchRegisteredToken(opts *bind.WatchOpts, sink chan<- *RollupRegisteredToken) (event.Subscription, error) {

	logs, sub, err := _Rollup.contract.WatchLogs(opts, "RegisteredToken")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupRegisteredToken)
				if err := _Rollup.contract.UnpackLog(event, "RegisteredToken", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRegisteredToken is a log parse operation binding the contract event 0x5dbaa701a7acef513f72a61799f7e50f4653f462b9f780d88d1b9bec89de2168.
//
// Solidity: event RegisteredToken(uint256 tokenType, address tokenContract)
func (_Rollup *RollupFilterer) ParseRegisteredToken(log types.Log) (*RollupRegisteredToken, error) {
	event := new(RollupRegisteredToken)
	if err := _Rollup.contract.UnpackLog(event, "RegisteredToken", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RollupRegistrationRequestIterator is returned from FilterRegistrationRequest and is used to iterate over the raw logs and unpacked data for RegistrationRequest events raised by the Rollup contract.
type RollupRegistrationRequestIterator struct {
	Event *RollupRegistrationRequest // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RollupRegistrationRequestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupRegistrationRequest)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RollupRegistrationRequest)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RollupRegistrationRequestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupRegistrationRequestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupRegistrationRequest represents a RegistrationRequest event raised by the Rollup contract.
type RollupRegistrationRequest struct {
	TokenContract common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterRegistrationRequest is a free log retrieval operation binding the contract event 0xdc79fc57451962cfe3916e686997a49229af75ce2055deb4c0f0fdf3d5d2e7c1.
//
// Solidity: event RegistrationRequest(address tokenContract)
func (_Rollup *RollupFilterer) FilterRegistrationRequest(opts *bind.FilterOpts) (*RollupRegistrationRequestIterator, error) {

	logs, sub, err := _Rollup.contract.FilterLogs(opts, "RegistrationRequest")
	if err != nil {
		return nil, err
	}
	return &RollupRegistrationRequestIterator{contract: _Rollup.contract, event: "RegistrationRequest", logs: logs, sub: sub}, nil
}

// WatchRegistrationRequest is a free log subscription operation binding the contract event 0xdc79fc57451962cfe3916e686997a49229af75ce2055deb4c0f0fdf3d5d2e7c1.
//
// Solidity: event RegistrationRequest(address tokenContract)
func (_Rollup *RollupFilterer) WatchRegistrationRequest(opts *bind.WatchOpts, sink chan<- *RollupRegistrationRequest) (event.Subscription, error) {

	logs, sub, err := _Rollup.contract.WatchLogs(opts, "RegistrationRequest")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupRegistrationRequest)
				if err := _Rollup.contract.UnpackLog(event, "RegistrationRequest", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRegistrationRequest is a log parse operation binding the contract event 0xdc79fc57451962cfe3916e686997a49229af75ce2055deb4c0f0fdf3d5d2e7c1.
//
// Solidity: event RegistrationRequest(address tokenContract)
func (_Rollup *RollupFilterer) ParseRegistrationRequest(log types.Log) (*RollupRegistrationRequest, error) {
	event := new(RollupRegistrationRequest)
	if err := _Rollup.contract.UnpackLog(event, "RegistrationRequest", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RollupRollbackFinalisationIterator is returned from FilterRollbackFinalisation and is used to iterate over the raw logs and unpacked data for RollbackFinalisation events raised by the Rollup contract.
type RollupRollbackFinalisationIterator struct {
	Event *RollupRollbackFinalisation // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RollupRollbackFinalisationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupRollbackFinalisation)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RollupRollbackFinalisation)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RollupRollbackFinalisationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupRollbackFinalisationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupRollbackFinalisation represents a RollbackFinalisation event raised by the Rollup contract.
type RollupRollbackFinalisation struct {
	TotalBatchesSlashed *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterRollbackFinalisation is a free log retrieval operation binding the contract event 0x8efd02dfe309f172ea08425236aa43cec05abc25ebbb9cd4ed1de0d5048fc91a.
//
// Solidity: event RollbackFinalisation(uint256 totalBatchesSlashed)
func (_Rollup *RollupFilterer) FilterRollbackFinalisation(opts *bind.FilterOpts) (*RollupRollbackFinalisationIterator, error) {

	logs, sub, err := _Rollup.contract.FilterLogs(opts, "RollbackFinalisation")
	if err != nil {
		return nil, err
	}
	return &RollupRollbackFinalisationIterator{contract: _Rollup.contract, event: "RollbackFinalisation", logs: logs, sub: sub}, nil
}

// WatchRollbackFinalisation is a free log subscription operation binding the contract event 0x8efd02dfe309f172ea08425236aa43cec05abc25ebbb9cd4ed1de0d5048fc91a.
//
// Solidity: event RollbackFinalisation(uint256 totalBatchesSlashed)
func (_Rollup *RollupFilterer) WatchRollbackFinalisation(opts *bind.WatchOpts, sink chan<- *RollupRollbackFinalisation) (event.Subscription, error) {

	logs, sub, err := _Rollup.contract.WatchLogs(opts, "RollbackFinalisation")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupRollbackFinalisation)
				if err := _Rollup.contract.UnpackLog(event, "RollbackFinalisation", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRollbackFinalisation is a log parse operation binding the contract event 0x8efd02dfe309f172ea08425236aa43cec05abc25ebbb9cd4ed1de0d5048fc91a.
//
// Solidity: event RollbackFinalisation(uint256 totalBatchesSlashed)
func (_Rollup *RollupFilterer) ParseRollbackFinalisation(log types.Log) (*RollupRollbackFinalisation, error) {
	event := new(RollupRollbackFinalisation)
	if err := _Rollup.contract.UnpackLog(event, "RollbackFinalisation", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RollupStakeWithdrawIterator is returned from FilterStakeWithdraw and is used to iterate over the raw logs and unpacked data for StakeWithdraw events raised by the Rollup contract.
type RollupStakeWithdrawIterator struct {
	Event *RollupStakeWithdraw // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RollupStakeWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupStakeWithdraw)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RollupStakeWithdraw)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RollupStakeWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupStakeWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupStakeWithdraw represents a StakeWithdraw event raised by the Rollup contract.
type RollupStakeWithdraw struct {
	Committed common.Address
	Amount    *big.Int
	BatchId   *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterStakeWithdraw is a free log retrieval operation binding the contract event 0xb00cb3d8b4494806e139df1041902f1984526ffe765fa4d7e12d64f4a47362fe.
//
// Solidity: event StakeWithdraw(address committed, uint256 amount, uint256 batch_id)
func (_Rollup *RollupFilterer) FilterStakeWithdraw(opts *bind.FilterOpts) (*RollupStakeWithdrawIterator, error) {

	logs, sub, err := _Rollup.contract.FilterLogs(opts, "StakeWithdraw")
	if err != nil {
		return nil, err
	}
	return &RollupStakeWithdrawIterator{contract: _Rollup.contract, event: "StakeWithdraw", logs: logs, sub: sub}, nil
}

// WatchStakeWithdraw is a free log subscription operation binding the contract event 0xb00cb3d8b4494806e139df1041902f1984526ffe765fa4d7e12d64f4a47362fe.
//
// Solidity: event StakeWithdraw(address committed, uint256 amount, uint256 batch_id)
func (_Rollup *RollupFilterer) WatchStakeWithdraw(opts *bind.WatchOpts, sink chan<- *RollupStakeWithdraw) (event.Subscription, error) {

	logs, sub, err := _Rollup.contract.WatchLogs(opts, "StakeWithdraw")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupStakeWithdraw)
				if err := _Rollup.contract.UnpackLog(event, "StakeWithdraw", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStakeWithdraw is a log parse operation binding the contract event 0xb00cb3d8b4494806e139df1041902f1984526ffe765fa4d7e12d64f4a47362fe.
//
// Solidity: event StakeWithdraw(address committed, uint256 amount, uint256 batch_id)
func (_Rollup *RollupFilterer) ParseStakeWithdraw(log types.Log) (*RollupStakeWithdraw, error) {
	event := new(RollupStakeWithdraw)
	if err := _Rollup.contract.UnpackLog(event, "StakeWithdraw", log); err != nil {
		return nil, err
	}
	return event, nil
}
