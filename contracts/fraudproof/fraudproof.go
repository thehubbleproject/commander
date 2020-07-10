// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package fraudproof

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

// TypesAccountInclusionProof is an auto generated low-level Go binding around an user-defined struct.
type TypesAccountInclusionProof struct {
	PathToAccount *big.Int
	Account       TypesUserAccount
}

// TypesAccountMerkleProof is an auto generated low-level Go binding around an user-defined struct.
type TypesAccountMerkleProof struct {
	AccountIP TypesAccountInclusionProof
	Siblings  [][32]byte
}

// TypesAccountProofs is an auto generated low-level Go binding around an user-defined struct.
type TypesAccountProofs struct {
	From TypesAccountMerkleProof
	To   TypesAccountMerkleProof
}

// TypesBatchValidationProofs is an auto generated low-level Go binding around an user-defined struct.
type TypesBatchValidationProofs struct {
	AccountProofs []TypesAccountProofs
	PdaProof      []TypesPDAMerkleProof
}

// TypesPDAInclusionProof is an auto generated low-level Go binding around an user-defined struct.
type TypesPDAInclusionProof struct {
	PathToPubkey *big.Int
	PubkeyLeaf   TypesPDALeaf
}

// TypesPDALeaf is an auto generated low-level Go binding around an user-defined struct.
type TypesPDALeaf struct {
	Pubkey []byte
}

// TypesPDAMerkleProof is an auto generated low-level Go binding around an user-defined struct.
type TypesPDAMerkleProof struct {
	Pda      TypesPDAInclusionProof
	Siblings [][32]byte
}

// TypesTransaction is an auto generated low-level Go binding around an user-defined struct.
type TypesTransaction struct {
	FromIndex *big.Int
	ToIndex   *big.Int
	TokenType *big.Int
	Nonce     *big.Int
	TxType    *big.Int
	Amount    *big.Int
	Signature []byte
}

// TypesUserAccount is an auto generated low-level Go binding around an user-defined struct.
type TypesUserAccount struct {
	ID        *big.Int
	TokenType *big.Int
	Balance   *big.Int
	Nonce     *big.Int
}

// FraudproofABI is the input ABI used to generate the binding from.
const FraudproofABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_registryAddr\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"constant\":true,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"ID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structTypes.UserAccount\",\"name\":\"account\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"numOfTokens\",\"type\":\"uint256\"}],\"name\":\"AddTokensToAccount\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"ID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structTypes.UserAccount\",\"name\":\"updatedAccount\",\"type\":\"tuple\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"pathToAccount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"ID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structTypes.UserAccount\",\"name\":\"account\",\"type\":\"tuple\"}],\"internalType\":\"structTypes.AccountInclusionProof\",\"name\":\"accountIP\",\"type\":\"tuple\"},{\"internalType\":\"bytes32[]\",\"name\":\"siblings\",\"type\":\"bytes32[]\"}],\"internalType\":\"structTypes.AccountMerkleProof\",\"name\":\"_merkle_proof\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"txType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structTypes.Transaction\",\"name\":\"transaction\",\"type\":\"tuple\"}],\"name\":\"ApplyTx\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"updatedAccount\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"newRoot\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ERR_FROM_TOKEN_TYPE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ERR_TOKEN_ADDR_INVAILD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ERR_TOKEN_AMT_INVAILD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ERR_TOKEN_NOT_ENOUGH_BAL\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ERR_TO_TOKEN_TYPE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"NO_ERR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"ID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structTypes.UserAccount\",\"name\":\"account\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"numOfTokens\",\"type\":\"uint256\"}],\"name\":\"RemoveTokensFromAccount\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"ID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structTypes.UserAccount\",\"name\":\"updatedAccount\",\"type\":\"tuple\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"ID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structTypes.UserAccount\",\"name\":\"new_account\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"pathToAccount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"ID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structTypes.UserAccount\",\"name\":\"account\",\"type\":\"tuple\"}],\"internalType\":\"structTypes.AccountInclusionProof\",\"name\":\"accountIP\",\"type\":\"tuple\"},{\"internalType\":\"bytes32[]\",\"name\":\"siblings\",\"type\":\"bytes32[]\"}],\"internalType\":\"structTypes.AccountMerkleProof\",\"name\":\"_merkle_proof\",\"type\":\"tuple\"}],\"name\":\"UpdateAccountWithSiblings\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"pathToAccount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"ID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structTypes.UserAccount\",\"name\":\"account\",\"type\":\"tuple\"}],\"internalType\":\"structTypes.AccountInclusionProof\",\"name\":\"accountIP\",\"type\":\"tuple\"},{\"internalType\":\"bytes32[]\",\"name\":\"siblings\",\"type\":\"bytes32[]\"}],\"internalType\":\"structTypes.AccountMerkleProof\",\"name\":\"merkle_proof\",\"type\":\"tuple\"}],\"name\":\"ValidateAccountMP\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_accountsRoot\",\"type\":\"bytes32\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"pathToPubkey\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"}],\"internalType\":\"structTypes.PDALeaf\",\"name\":\"pubkey_leaf\",\"type\":\"tuple\"}],\"internalType\":\"structTypes.PDAInclusionProof\",\"name\":\"_pda\",\"type\":\"tuple\"},{\"internalType\":\"bytes32[]\",\"name\":\"siblings\",\"type\":\"bytes32[]\"}],\"internalType\":\"structTypes.PDAMerkleProof\",\"name\":\"_from_pda_proof\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"from_index\",\"type\":\"uint256\"}],\"name\":\"ValidatePubkeyAvailability\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"txType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structTypes.Transaction\",\"name\":\"_tx\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"pathToPubkey\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"}],\"internalType\":\"structTypes.PDALeaf\",\"name\":\"pubkey_leaf\",\"type\":\"tuple\"}],\"internalType\":\"structTypes.PDAInclusionProof\",\"name\":\"_pda\",\"type\":\"tuple\"},{\"internalType\":\"bytes32[]\",\"name\":\"siblings\",\"type\":\"bytes32[]\"}],\"internalType\":\"structTypes.PDAMerkleProof\",\"name\":\"_from_pda_proof\",\"type\":\"tuple\"}],\"name\":\"ValidateSignature\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ZERO_BYTES32\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"governance\",\"outputs\":[{\"internalType\":\"contractGovernance\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"merkleUtils\",\"outputs\":[{\"internalType\":\"contractMerkleTreeUtils\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"nameRegistry\",\"outputs\":[{\"internalType\":\"contractNameRegistry\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tokenRegistry\",\"outputs\":[{\"internalType\":\"contractITokenRegistry\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"txType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structTypes.Transaction\",\"name\":\"_tx\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"ID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structTypes.UserAccount\",\"name\":\"_from_account\",\"type\":\"tuple\"}],\"name\":\"validateTxBasic\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"txType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structTypes.Transaction[]\",\"name\":\"_txs\",\"type\":\"tuple[]\"}],\"name\":\"generateTxRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"txRoot\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"accountsRoot\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"txType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structTypes.Transaction[]\",\"name\":\"_txs\",\"type\":\"tuple[]\"},{\"components\":[{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"pathToAccount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"ID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structTypes.UserAccount\",\"name\":\"account\",\"type\":\"tuple\"}],\"internalType\":\"structTypes.AccountInclusionProof\",\"name\":\"accountIP\",\"type\":\"tuple\"},{\"internalType\":\"bytes32[]\",\"name\":\"siblings\",\"type\":\"bytes32[]\"}],\"internalType\":\"structTypes.AccountMerkleProof\",\"name\":\"from\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"pathToAccount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"ID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structTypes.UserAccount\",\"name\":\"account\",\"type\":\"tuple\"}],\"internalType\":\"structTypes.AccountInclusionProof\",\"name\":\"accountIP\",\"type\":\"tuple\"},{\"internalType\":\"bytes32[]\",\"name\":\"siblings\",\"type\":\"bytes32[]\"}],\"internalType\":\"structTypes.AccountMerkleProof\",\"name\":\"to\",\"type\":\"tuple\"}],\"internalType\":\"structTypes.AccountProofs[]\",\"name\":\"accountProofs\",\"type\":\"tuple[]\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"pathToPubkey\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"}],\"internalType\":\"structTypes.PDALeaf\",\"name\":\"pubkey_leaf\",\"type\":\"tuple\"}],\"internalType\":\"structTypes.PDAInclusionProof\",\"name\":\"_pda\",\"type\":\"tuple\"},{\"internalType\":\"bytes32[]\",\"name\":\"siblings\",\"type\":\"bytes32[]\"}],\"internalType\":\"structTypes.PDAMerkleProof[]\",\"name\":\"pdaProof\",\"type\":\"tuple[]\"}],\"internalType\":\"structTypes.BatchValidationProofs\",\"name\":\"batchProofs\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"expectedTxRoot\",\"type\":\"bytes32\"}],\"name\":\"processBatch\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_balanceRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_accountsRoot\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"txType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structTypes.Transaction\",\"name\":\"_tx\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"pathToPubkey\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"}],\"internalType\":\"structTypes.PDALeaf\",\"name\":\"pubkey_leaf\",\"type\":\"tuple\"}],\"internalType\":\"structTypes.PDAInclusionProof\",\"name\":\"_pda\",\"type\":\"tuple\"},{\"internalType\":\"bytes32[]\",\"name\":\"siblings\",\"type\":\"bytes32[]\"}],\"internalType\":\"structTypes.PDAMerkleProof\",\"name\":\"_from_pda_proof\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"pathToAccount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"ID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structTypes.UserAccount\",\"name\":\"account\",\"type\":\"tuple\"}],\"internalType\":\"structTypes.AccountInclusionProof\",\"name\":\"accountIP\",\"type\":\"tuple\"},{\"internalType\":\"bytes32[]\",\"name\":\"siblings\",\"type\":\"bytes32[]\"}],\"internalType\":\"structTypes.AccountMerkleProof\",\"name\":\"from\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"pathToAccount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"ID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structTypes.UserAccount\",\"name\":\"account\",\"type\":\"tuple\"}],\"internalType\":\"structTypes.AccountInclusionProof\",\"name\":\"accountIP\",\"type\":\"tuple\"},{\"internalType\":\"bytes32[]\",\"name\":\"siblings\",\"type\":\"bytes32[]\"}],\"internalType\":\"structTypes.AccountMerkleProof\",\"name\":\"to\",\"type\":\"tuple\"}],\"internalType\":\"structTypes.AccountProofs\",\"name\":\"accountProofs\",\"type\":\"tuple\"}],\"name\":\"processTx\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// Fraudproof is an auto generated Go binding around an Ethereum contract.
type Fraudproof struct {
	FraudproofCaller     // Read-only binding to the contract
	FraudproofTransactor // Write-only binding to the contract
	FraudproofFilterer   // Log filterer for contract events
}

// FraudproofCaller is an auto generated read-only Go binding around an Ethereum contract.
type FraudproofCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FraudproofTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FraudproofTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FraudproofFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FraudproofFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FraudproofSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FraudproofSession struct {
	Contract     *Fraudproof       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FraudproofCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FraudproofCallerSession struct {
	Contract *FraudproofCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// FraudproofTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FraudproofTransactorSession struct {
	Contract     *FraudproofTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// FraudproofRaw is an auto generated low-level Go binding around an Ethereum contract.
type FraudproofRaw struct {
	Contract *Fraudproof // Generic contract binding to access the raw methods on
}

// FraudproofCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FraudproofCallerRaw struct {
	Contract *FraudproofCaller // Generic read-only contract binding to access the raw methods on
}

// FraudproofTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FraudproofTransactorRaw struct {
	Contract *FraudproofTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFraudproof creates a new instance of Fraudproof, bound to a specific deployed contract.
func NewFraudproof(address common.Address, backend bind.ContractBackend) (*Fraudproof, error) {
	contract, err := bindFraudproof(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Fraudproof{FraudproofCaller: FraudproofCaller{contract: contract}, FraudproofTransactor: FraudproofTransactor{contract: contract}, FraudproofFilterer: FraudproofFilterer{contract: contract}}, nil
}

// NewFraudproofCaller creates a new read-only instance of Fraudproof, bound to a specific deployed contract.
func NewFraudproofCaller(address common.Address, caller bind.ContractCaller) (*FraudproofCaller, error) {
	contract, err := bindFraudproof(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FraudproofCaller{contract: contract}, nil
}

// NewFraudproofTransactor creates a new write-only instance of Fraudproof, bound to a specific deployed contract.
func NewFraudproofTransactor(address common.Address, transactor bind.ContractTransactor) (*FraudproofTransactor, error) {
	contract, err := bindFraudproof(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FraudproofTransactor{contract: contract}, nil
}

// NewFraudproofFilterer creates a new log filterer instance of Fraudproof, bound to a specific deployed contract.
func NewFraudproofFilterer(address common.Address, filterer bind.ContractFilterer) (*FraudproofFilterer, error) {
	contract, err := bindFraudproof(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FraudproofFilterer{contract: contract}, nil
}

// bindFraudproof binds a generic wrapper to an already deployed contract.
func bindFraudproof(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FraudproofABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Fraudproof *FraudproofRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Fraudproof.Contract.FraudproofCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Fraudproof *FraudproofRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Fraudproof.Contract.FraudproofTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Fraudproof *FraudproofRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Fraudproof.Contract.FraudproofTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Fraudproof *FraudproofCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Fraudproof.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Fraudproof *FraudproofTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Fraudproof.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Fraudproof *FraudproofTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Fraudproof.Contract.contract.Transact(opts, method, params...)
}

// AddTokensToAccount is a free data retrieval call binding the contract method 0xb581273f.
//
// Solidity: function AddTokensToAccount(TypesUserAccount account, uint256 numOfTokens) pure returns(TypesUserAccount updatedAccount)
func (_Fraudproof *FraudproofCaller) AddTokensToAccount(opts *bind.CallOpts, account TypesUserAccount, numOfTokens *big.Int) (TypesUserAccount, error) {
	var (
		ret0 = new(TypesUserAccount)
	)
	out := ret0
	err := _Fraudproof.contract.Call(opts, out, "AddTokensToAccount", account, numOfTokens)
	return *ret0, err
}

// AddTokensToAccount is a free data retrieval call binding the contract method 0xb581273f.
//
// Solidity: function AddTokensToAccount(TypesUserAccount account, uint256 numOfTokens) pure returns(TypesUserAccount updatedAccount)
func (_Fraudproof *FraudproofSession) AddTokensToAccount(account TypesUserAccount, numOfTokens *big.Int) (TypesUserAccount, error) {
	return _Fraudproof.Contract.AddTokensToAccount(&_Fraudproof.CallOpts, account, numOfTokens)
}

// AddTokensToAccount is a free data retrieval call binding the contract method 0xb581273f.
//
// Solidity: function AddTokensToAccount(TypesUserAccount account, uint256 numOfTokens) pure returns(TypesUserAccount updatedAccount)
func (_Fraudproof *FraudproofCallerSession) AddTokensToAccount(account TypesUserAccount, numOfTokens *big.Int) (TypesUserAccount, error) {
	return _Fraudproof.Contract.AddTokensToAccount(&_Fraudproof.CallOpts, account, numOfTokens)
}

// ApplyTx is a free data retrieval call binding the contract method 0x3100ef1c.
//
// Solidity: function ApplyTx(TypesAccountMerkleProof _merkle_proof, TypesTransaction transaction) view returns(bytes updatedAccount, bytes32 newRoot)
func (_Fraudproof *FraudproofCaller) ApplyTx(opts *bind.CallOpts, _merkle_proof TypesAccountMerkleProof, transaction TypesTransaction) (struct {
	UpdatedAccount []byte
	NewRoot        [32]byte
}, error) {
	ret := new(struct {
		UpdatedAccount []byte
		NewRoot        [32]byte
	})
	out := ret
	err := _Fraudproof.contract.Call(opts, out, "ApplyTx", _merkle_proof, transaction)
	return *ret, err
}

// ApplyTx is a free data retrieval call binding the contract method 0x3100ef1c.
//
// Solidity: function ApplyTx(TypesAccountMerkleProof _merkle_proof, TypesTransaction transaction) view returns(bytes updatedAccount, bytes32 newRoot)
func (_Fraudproof *FraudproofSession) ApplyTx(_merkle_proof TypesAccountMerkleProof, transaction TypesTransaction) (struct {
	UpdatedAccount []byte
	NewRoot        [32]byte
}, error) {
	return _Fraudproof.Contract.ApplyTx(&_Fraudproof.CallOpts, _merkle_proof, transaction)
}

// ApplyTx is a free data retrieval call binding the contract method 0x3100ef1c.
//
// Solidity: function ApplyTx(TypesAccountMerkleProof _merkle_proof, TypesTransaction transaction) view returns(bytes updatedAccount, bytes32 newRoot)
func (_Fraudproof *FraudproofCallerSession) ApplyTx(_merkle_proof TypesAccountMerkleProof, transaction TypesTransaction) (struct {
	UpdatedAccount []byte
	NewRoot        [32]byte
}, error) {
	return _Fraudproof.Contract.ApplyTx(&_Fraudproof.CallOpts, _merkle_proof, transaction)
}

// ERRFROMTOKENTYPE is a free data retrieval call binding the contract method 0xe5484755.
//
// Solidity: function ERR_FROM_TOKEN_TYPE() view returns(uint256)
func (_Fraudproof *FraudproofCaller) ERRFROMTOKENTYPE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Fraudproof.contract.Call(opts, out, "ERR_FROM_TOKEN_TYPE")
	return *ret0, err
}

// ERRFROMTOKENTYPE is a free data retrieval call binding the contract method 0xe5484755.
//
// Solidity: function ERR_FROM_TOKEN_TYPE() view returns(uint256)
func (_Fraudproof *FraudproofSession) ERRFROMTOKENTYPE() (*big.Int, error) {
	return _Fraudproof.Contract.ERRFROMTOKENTYPE(&_Fraudproof.CallOpts)
}

// ERRFROMTOKENTYPE is a free data retrieval call binding the contract method 0xe5484755.
//
// Solidity: function ERR_FROM_TOKEN_TYPE() view returns(uint256)
func (_Fraudproof *FraudproofCallerSession) ERRFROMTOKENTYPE() (*big.Int, error) {
	return _Fraudproof.Contract.ERRFROMTOKENTYPE(&_Fraudproof.CallOpts)
}

// ERRTOKENADDRINVAILD is a free data retrieval call binding the contract method 0xd79b6711.
//
// Solidity: function ERR_TOKEN_ADDR_INVAILD() view returns(uint256)
func (_Fraudproof *FraudproofCaller) ERRTOKENADDRINVAILD(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Fraudproof.contract.Call(opts, out, "ERR_TOKEN_ADDR_INVAILD")
	return *ret0, err
}

// ERRTOKENADDRINVAILD is a free data retrieval call binding the contract method 0xd79b6711.
//
// Solidity: function ERR_TOKEN_ADDR_INVAILD() view returns(uint256)
func (_Fraudproof *FraudproofSession) ERRTOKENADDRINVAILD() (*big.Int, error) {
	return _Fraudproof.Contract.ERRTOKENADDRINVAILD(&_Fraudproof.CallOpts)
}

// ERRTOKENADDRINVAILD is a free data retrieval call binding the contract method 0xd79b6711.
//
// Solidity: function ERR_TOKEN_ADDR_INVAILD() view returns(uint256)
func (_Fraudproof *FraudproofCallerSession) ERRTOKENADDRINVAILD() (*big.Int, error) {
	return _Fraudproof.Contract.ERRTOKENADDRINVAILD(&_Fraudproof.CallOpts)
}

// ERRTOKENAMTINVAILD is a free data retrieval call binding the contract method 0x060319ba.
//
// Solidity: function ERR_TOKEN_AMT_INVAILD() view returns(uint256)
func (_Fraudproof *FraudproofCaller) ERRTOKENAMTINVAILD(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Fraudproof.contract.Call(opts, out, "ERR_TOKEN_AMT_INVAILD")
	return *ret0, err
}

// ERRTOKENAMTINVAILD is a free data retrieval call binding the contract method 0x060319ba.
//
// Solidity: function ERR_TOKEN_AMT_INVAILD() view returns(uint256)
func (_Fraudproof *FraudproofSession) ERRTOKENAMTINVAILD() (*big.Int, error) {
	return _Fraudproof.Contract.ERRTOKENAMTINVAILD(&_Fraudproof.CallOpts)
}

// ERRTOKENAMTINVAILD is a free data retrieval call binding the contract method 0x060319ba.
//
// Solidity: function ERR_TOKEN_AMT_INVAILD() view returns(uint256)
func (_Fraudproof *FraudproofCallerSession) ERRTOKENAMTINVAILD() (*big.Int, error) {
	return _Fraudproof.Contract.ERRTOKENAMTINVAILD(&_Fraudproof.CallOpts)
}

// ERRTOKENNOTENOUGHBAL is a free data retrieval call binding the contract method 0xd15df193.
//
// Solidity: function ERR_TOKEN_NOT_ENOUGH_BAL() view returns(uint256)
func (_Fraudproof *FraudproofCaller) ERRTOKENNOTENOUGHBAL(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Fraudproof.contract.Call(opts, out, "ERR_TOKEN_NOT_ENOUGH_BAL")
	return *ret0, err
}

// ERRTOKENNOTENOUGHBAL is a free data retrieval call binding the contract method 0xd15df193.
//
// Solidity: function ERR_TOKEN_NOT_ENOUGH_BAL() view returns(uint256)
func (_Fraudproof *FraudproofSession) ERRTOKENNOTENOUGHBAL() (*big.Int, error) {
	return _Fraudproof.Contract.ERRTOKENNOTENOUGHBAL(&_Fraudproof.CallOpts)
}

// ERRTOKENNOTENOUGHBAL is a free data retrieval call binding the contract method 0xd15df193.
//
// Solidity: function ERR_TOKEN_NOT_ENOUGH_BAL() view returns(uint256)
func (_Fraudproof *FraudproofCallerSession) ERRTOKENNOTENOUGHBAL() (*big.Int, error) {
	return _Fraudproof.Contract.ERRTOKENNOTENOUGHBAL(&_Fraudproof.CallOpts)
}

// ERRTOTOKENTYPE is a free data retrieval call binding the contract method 0x85ac3744.
//
// Solidity: function ERR_TO_TOKEN_TYPE() view returns(uint256)
func (_Fraudproof *FraudproofCaller) ERRTOTOKENTYPE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Fraudproof.contract.Call(opts, out, "ERR_TO_TOKEN_TYPE")
	return *ret0, err
}

// ERRTOTOKENTYPE is a free data retrieval call binding the contract method 0x85ac3744.
//
// Solidity: function ERR_TO_TOKEN_TYPE() view returns(uint256)
func (_Fraudproof *FraudproofSession) ERRTOTOKENTYPE() (*big.Int, error) {
	return _Fraudproof.Contract.ERRTOTOKENTYPE(&_Fraudproof.CallOpts)
}

// ERRTOTOKENTYPE is a free data retrieval call binding the contract method 0x85ac3744.
//
// Solidity: function ERR_TO_TOKEN_TYPE() view returns(uint256)
func (_Fraudproof *FraudproofCallerSession) ERRTOTOKENTYPE() (*big.Int, error) {
	return _Fraudproof.Contract.ERRTOTOKENTYPE(&_Fraudproof.CallOpts)
}

// NOERR is a free data retrieval call binding the contract method 0x7932736b.
//
// Solidity: function NO_ERR() view returns(uint256)
func (_Fraudproof *FraudproofCaller) NOERR(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Fraudproof.contract.Call(opts, out, "NO_ERR")
	return *ret0, err
}

// NOERR is a free data retrieval call binding the contract method 0x7932736b.
//
// Solidity: function NO_ERR() view returns(uint256)
func (_Fraudproof *FraudproofSession) NOERR() (*big.Int, error) {
	return _Fraudproof.Contract.NOERR(&_Fraudproof.CallOpts)
}

// NOERR is a free data retrieval call binding the contract method 0x7932736b.
//
// Solidity: function NO_ERR() view returns(uint256)
func (_Fraudproof *FraudproofCallerSession) NOERR() (*big.Int, error) {
	return _Fraudproof.Contract.NOERR(&_Fraudproof.CallOpts)
}

// RemoveTokensFromAccount is a free data retrieval call binding the contract method 0xe664f2bc.
//
// Solidity: function RemoveTokensFromAccount(TypesUserAccount account, uint256 numOfTokens) pure returns(TypesUserAccount updatedAccount)
func (_Fraudproof *FraudproofCaller) RemoveTokensFromAccount(opts *bind.CallOpts, account TypesUserAccount, numOfTokens *big.Int) (TypesUserAccount, error) {
	var (
		ret0 = new(TypesUserAccount)
	)
	out := ret0
	err := _Fraudproof.contract.Call(opts, out, "RemoveTokensFromAccount", account, numOfTokens)
	return *ret0, err
}

// RemoveTokensFromAccount is a free data retrieval call binding the contract method 0xe664f2bc.
//
// Solidity: function RemoveTokensFromAccount(TypesUserAccount account, uint256 numOfTokens) pure returns(TypesUserAccount updatedAccount)
func (_Fraudproof *FraudproofSession) RemoveTokensFromAccount(account TypesUserAccount, numOfTokens *big.Int) (TypesUserAccount, error) {
	return _Fraudproof.Contract.RemoveTokensFromAccount(&_Fraudproof.CallOpts, account, numOfTokens)
}

// RemoveTokensFromAccount is a free data retrieval call binding the contract method 0xe664f2bc.
//
// Solidity: function RemoveTokensFromAccount(TypesUserAccount account, uint256 numOfTokens) pure returns(TypesUserAccount updatedAccount)
func (_Fraudproof *FraudproofCallerSession) RemoveTokensFromAccount(account TypesUserAccount, numOfTokens *big.Int) (TypesUserAccount, error) {
	return _Fraudproof.Contract.RemoveTokensFromAccount(&_Fraudproof.CallOpts, account, numOfTokens)
}

// UpdateAccountWithSiblings is a free data retrieval call binding the contract method 0x17610f46.
//
// Solidity: function UpdateAccountWithSiblings(TypesUserAccount new_account, TypesAccountMerkleProof _merkle_proof) view returns(bytes32)
func (_Fraudproof *FraudproofCaller) UpdateAccountWithSiblings(opts *bind.CallOpts, new_account TypesUserAccount, _merkle_proof TypesAccountMerkleProof) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Fraudproof.contract.Call(opts, out, "UpdateAccountWithSiblings", new_account, _merkle_proof)
	return *ret0, err
}

// UpdateAccountWithSiblings is a free data retrieval call binding the contract method 0x17610f46.
//
// Solidity: function UpdateAccountWithSiblings(TypesUserAccount new_account, TypesAccountMerkleProof _merkle_proof) view returns(bytes32)
func (_Fraudproof *FraudproofSession) UpdateAccountWithSiblings(new_account TypesUserAccount, _merkle_proof TypesAccountMerkleProof) ([32]byte, error) {
	return _Fraudproof.Contract.UpdateAccountWithSiblings(&_Fraudproof.CallOpts, new_account, _merkle_proof)
}

// UpdateAccountWithSiblings is a free data retrieval call binding the contract method 0x17610f46.
//
// Solidity: function UpdateAccountWithSiblings(TypesUserAccount new_account, TypesAccountMerkleProof _merkle_proof) view returns(bytes32)
func (_Fraudproof *FraudproofCallerSession) UpdateAccountWithSiblings(new_account TypesUserAccount, _merkle_proof TypesAccountMerkleProof) ([32]byte, error) {
	return _Fraudproof.Contract.UpdateAccountWithSiblings(&_Fraudproof.CallOpts, new_account, _merkle_proof)
}

// ValidateAccountMP is a free data retrieval call binding the contract method 0x1a6964bc.
//
// Solidity: function ValidateAccountMP(bytes32 root, TypesAccountMerkleProof merkle_proof) view returns()
func (_Fraudproof *FraudproofCaller) ValidateAccountMP(opts *bind.CallOpts, root [32]byte, merkle_proof TypesAccountMerkleProof) error {
	var ()
	out := &[]interface{}{}
	err := _Fraudproof.contract.Call(opts, out, "ValidateAccountMP", root, merkle_proof)
	return err
}

// ValidateAccountMP is a free data retrieval call binding the contract method 0x1a6964bc.
//
// Solidity: function ValidateAccountMP(bytes32 root, TypesAccountMerkleProof merkle_proof) view returns()
func (_Fraudproof *FraudproofSession) ValidateAccountMP(root [32]byte, merkle_proof TypesAccountMerkleProof) error {
	return _Fraudproof.Contract.ValidateAccountMP(&_Fraudproof.CallOpts, root, merkle_proof)
}

// ValidateAccountMP is a free data retrieval call binding the contract method 0x1a6964bc.
//
// Solidity: function ValidateAccountMP(bytes32 root, TypesAccountMerkleProof merkle_proof) view returns()
func (_Fraudproof *FraudproofCallerSession) ValidateAccountMP(root [32]byte, merkle_proof TypesAccountMerkleProof) error {
	return _Fraudproof.Contract.ValidateAccountMP(&_Fraudproof.CallOpts, root, merkle_proof)
}

// ValidatePubkeyAvailability is a free data retrieval call binding the contract method 0x5b684df6.
//
// Solidity: function ValidatePubkeyAvailability(bytes32 _accountsRoot, TypesPDAMerkleProof _from_pda_proof, uint256 from_index) view returns()
func (_Fraudproof *FraudproofCaller) ValidatePubkeyAvailability(opts *bind.CallOpts, _accountsRoot [32]byte, _from_pda_proof TypesPDAMerkleProof, from_index *big.Int) error {
	var ()
	out := &[]interface{}{}
	err := _Fraudproof.contract.Call(opts, out, "ValidatePubkeyAvailability", _accountsRoot, _from_pda_proof, from_index)
	return err
}

// ValidatePubkeyAvailability is a free data retrieval call binding the contract method 0x5b684df6.
//
// Solidity: function ValidatePubkeyAvailability(bytes32 _accountsRoot, TypesPDAMerkleProof _from_pda_proof, uint256 from_index) view returns()
func (_Fraudproof *FraudproofSession) ValidatePubkeyAvailability(_accountsRoot [32]byte, _from_pda_proof TypesPDAMerkleProof, from_index *big.Int) error {
	return _Fraudproof.Contract.ValidatePubkeyAvailability(&_Fraudproof.CallOpts, _accountsRoot, _from_pda_proof, from_index)
}

// ValidatePubkeyAvailability is a free data retrieval call binding the contract method 0x5b684df6.
//
// Solidity: function ValidatePubkeyAvailability(bytes32 _accountsRoot, TypesPDAMerkleProof _from_pda_proof, uint256 from_index) view returns()
func (_Fraudproof *FraudproofCallerSession) ValidatePubkeyAvailability(_accountsRoot [32]byte, _from_pda_proof TypesPDAMerkleProof, from_index *big.Int) error {
	return _Fraudproof.Contract.ValidatePubkeyAvailability(&_Fraudproof.CallOpts, _accountsRoot, _from_pda_proof, from_index)
}

// ValidateSignature is a free data retrieval call binding the contract method 0x4c983eee.
//
// Solidity: function ValidateSignature(TypesTransaction _tx, TypesPDAMerkleProof _from_pda_proof) pure returns(bool)
func (_Fraudproof *FraudproofCaller) ValidateSignature(opts *bind.CallOpts, _tx TypesTransaction, _from_pda_proof TypesPDAMerkleProof) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Fraudproof.contract.Call(opts, out, "ValidateSignature", _tx, _from_pda_proof)
	return *ret0, err
}

// ValidateSignature is a free data retrieval call binding the contract method 0x4c983eee.
//
// Solidity: function ValidateSignature(TypesTransaction _tx, TypesPDAMerkleProof _from_pda_proof) pure returns(bool)
func (_Fraudproof *FraudproofSession) ValidateSignature(_tx TypesTransaction, _from_pda_proof TypesPDAMerkleProof) (bool, error) {
	return _Fraudproof.Contract.ValidateSignature(&_Fraudproof.CallOpts, _tx, _from_pda_proof)
}

// ValidateSignature is a free data retrieval call binding the contract method 0x4c983eee.
//
// Solidity: function ValidateSignature(TypesTransaction _tx, TypesPDAMerkleProof _from_pda_proof) pure returns(bool)
func (_Fraudproof *FraudproofCallerSession) ValidateSignature(_tx TypesTransaction, _from_pda_proof TypesPDAMerkleProof) (bool, error) {
	return _Fraudproof.Contract.ValidateSignature(&_Fraudproof.CallOpts, _tx, _from_pda_proof)
}

// ZEROBYTES32 is a free data retrieval call binding the contract method 0x069321b0.
//
// Solidity: function ZERO_BYTES32() view returns(bytes32)
func (_Fraudproof *FraudproofCaller) ZEROBYTES32(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Fraudproof.contract.Call(opts, out, "ZERO_BYTES32")
	return *ret0, err
}

// ZEROBYTES32 is a free data retrieval call binding the contract method 0x069321b0.
//
// Solidity: function ZERO_BYTES32() view returns(bytes32)
func (_Fraudproof *FraudproofSession) ZEROBYTES32() ([32]byte, error) {
	return _Fraudproof.Contract.ZEROBYTES32(&_Fraudproof.CallOpts)
}

// ZEROBYTES32 is a free data retrieval call binding the contract method 0x069321b0.
//
// Solidity: function ZERO_BYTES32() view returns(bytes32)
func (_Fraudproof *FraudproofCallerSession) ZEROBYTES32() ([32]byte, error) {
	return _Fraudproof.Contract.ZEROBYTES32(&_Fraudproof.CallOpts)
}

// GenerateTxRoot is a free data retrieval call binding the contract method 0xbac089e1.
//
// Solidity: function generateTxRoot([]TypesTransaction _txs) view returns(bytes32 txRoot)
func (_Fraudproof *FraudproofCaller) GenerateTxRoot(opts *bind.CallOpts, _txs []TypesTransaction) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Fraudproof.contract.Call(opts, out, "generateTxRoot", _txs)
	return *ret0, err
}

// GenerateTxRoot is a free data retrieval call binding the contract method 0xbac089e1.
//
// Solidity: function generateTxRoot([]TypesTransaction _txs) view returns(bytes32 txRoot)
func (_Fraudproof *FraudproofSession) GenerateTxRoot(_txs []TypesTransaction) ([32]byte, error) {
	return _Fraudproof.Contract.GenerateTxRoot(&_Fraudproof.CallOpts, _txs)
}

// GenerateTxRoot is a free data retrieval call binding the contract method 0xbac089e1.
//
// Solidity: function generateTxRoot([]TypesTransaction _txs) view returns(bytes32 txRoot)
func (_Fraudproof *FraudproofCallerSession) GenerateTxRoot(_txs []TypesTransaction) ([32]byte, error) {
	return _Fraudproof.Contract.GenerateTxRoot(&_Fraudproof.CallOpts, _txs)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_Fraudproof *FraudproofCaller) Governance(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Fraudproof.contract.Call(opts, out, "governance")
	return *ret0, err
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_Fraudproof *FraudproofSession) Governance() (common.Address, error) {
	return _Fraudproof.Contract.Governance(&_Fraudproof.CallOpts)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_Fraudproof *FraudproofCallerSession) Governance() (common.Address, error) {
	return _Fraudproof.Contract.Governance(&_Fraudproof.CallOpts)
}

// MerkleUtils is a free data retrieval call binding the contract method 0x47b0f08e.
//
// Solidity: function merkleUtils() view returns(address)
func (_Fraudproof *FraudproofCaller) MerkleUtils(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Fraudproof.contract.Call(opts, out, "merkleUtils")
	return *ret0, err
}

// MerkleUtils is a free data retrieval call binding the contract method 0x47b0f08e.
//
// Solidity: function merkleUtils() view returns(address)
func (_Fraudproof *FraudproofSession) MerkleUtils() (common.Address, error) {
	return _Fraudproof.Contract.MerkleUtils(&_Fraudproof.CallOpts)
}

// MerkleUtils is a free data retrieval call binding the contract method 0x47b0f08e.
//
// Solidity: function merkleUtils() view returns(address)
func (_Fraudproof *FraudproofCallerSession) MerkleUtils() (common.Address, error) {
	return _Fraudproof.Contract.MerkleUtils(&_Fraudproof.CallOpts)
}

// NameRegistry is a free data retrieval call binding the contract method 0x4eb7221a.
//
// Solidity: function nameRegistry() view returns(address)
func (_Fraudproof *FraudproofCaller) NameRegistry(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Fraudproof.contract.Call(opts, out, "nameRegistry")
	return *ret0, err
}

// NameRegistry is a free data retrieval call binding the contract method 0x4eb7221a.
//
// Solidity: function nameRegistry() view returns(address)
func (_Fraudproof *FraudproofSession) NameRegistry() (common.Address, error) {
	return _Fraudproof.Contract.NameRegistry(&_Fraudproof.CallOpts)
}

// NameRegistry is a free data retrieval call binding the contract method 0x4eb7221a.
//
// Solidity: function nameRegistry() view returns(address)
func (_Fraudproof *FraudproofCallerSession) NameRegistry() (common.Address, error) {
	return _Fraudproof.Contract.NameRegistry(&_Fraudproof.CallOpts)
}

// ProcessBatch is a free data retrieval call binding the contract method 0x5b7f2470.
//
// Solidity: function processBatch(bytes32 stateRoot, bytes32 accountsRoot, []TypesTransaction _txs, TypesBatchValidationProofs batchProofs, bytes32 expectedTxRoot) view returns(bytes32, bytes32, bool)
func (_Fraudproof *FraudproofCaller) ProcessBatch(opts *bind.CallOpts, stateRoot [32]byte, accountsRoot [32]byte, _txs []TypesTransaction, batchProofs TypesBatchValidationProofs, expectedTxRoot [32]byte) ([32]byte, [32]byte, bool, error) {
	var (
		ret0 = new([32]byte)
		ret1 = new([32]byte)
		ret2 = new(bool)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
	}
	err := _Fraudproof.contract.Call(opts, out, "processBatch", stateRoot, accountsRoot, _txs, batchProofs, expectedTxRoot)
	return *ret0, *ret1, *ret2, err
}

// ProcessBatch is a free data retrieval call binding the contract method 0x5b7f2470.
//
// Solidity: function processBatch(bytes32 stateRoot, bytes32 accountsRoot, []TypesTransaction _txs, TypesBatchValidationProofs batchProofs, bytes32 expectedTxRoot) view returns(bytes32, bytes32, bool)
func (_Fraudproof *FraudproofSession) ProcessBatch(stateRoot [32]byte, accountsRoot [32]byte, _txs []TypesTransaction, batchProofs TypesBatchValidationProofs, expectedTxRoot [32]byte) ([32]byte, [32]byte, bool, error) {
	return _Fraudproof.Contract.ProcessBatch(&_Fraudproof.CallOpts, stateRoot, accountsRoot, _txs, batchProofs, expectedTxRoot)
}

// ProcessBatch is a free data retrieval call binding the contract method 0x5b7f2470.
//
// Solidity: function processBatch(bytes32 stateRoot, bytes32 accountsRoot, []TypesTransaction _txs, TypesBatchValidationProofs batchProofs, bytes32 expectedTxRoot) view returns(bytes32, bytes32, bool)
func (_Fraudproof *FraudproofCallerSession) ProcessBatch(stateRoot [32]byte, accountsRoot [32]byte, _txs []TypesTransaction, batchProofs TypesBatchValidationProofs, expectedTxRoot [32]byte) ([32]byte, [32]byte, bool, error) {
	return _Fraudproof.Contract.ProcessBatch(&_Fraudproof.CallOpts, stateRoot, accountsRoot, _txs, batchProofs, expectedTxRoot)
}

// ProcessTx is a free data retrieval call binding the contract method 0x8cb750a8.
//
// Solidity: function processTx(bytes32 _balanceRoot, bytes32 _accountsRoot, TypesTransaction _tx, TypesPDAMerkleProof _from_pda_proof, TypesAccountProofs accountProofs) view returns(bytes32, bytes, bytes, uint256, bool)
func (_Fraudproof *FraudproofCaller) ProcessTx(opts *bind.CallOpts, _balanceRoot [32]byte, _accountsRoot [32]byte, _tx TypesTransaction, _from_pda_proof TypesPDAMerkleProof, accountProofs TypesAccountProofs) ([32]byte, []byte, []byte, *big.Int, bool, error) {
	var (
		ret0 = new([32]byte)
		ret1 = new([]byte)
		ret2 = new([]byte)
		ret3 = new(*big.Int)
		ret4 = new(bool)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
	}
	err := _Fraudproof.contract.Call(opts, out, "processTx", _balanceRoot, _accountsRoot, _tx, _from_pda_proof, accountProofs)
	return *ret0, *ret1, *ret2, *ret3, *ret4, err
}

// ProcessTx is a free data retrieval call binding the contract method 0x8cb750a8.
//
// Solidity: function processTx(bytes32 _balanceRoot, bytes32 _accountsRoot, TypesTransaction _tx, TypesPDAMerkleProof _from_pda_proof, TypesAccountProofs accountProofs) view returns(bytes32, bytes, bytes, uint256, bool)
func (_Fraudproof *FraudproofSession) ProcessTx(_balanceRoot [32]byte, _accountsRoot [32]byte, _tx TypesTransaction, _from_pda_proof TypesPDAMerkleProof, accountProofs TypesAccountProofs) ([32]byte, []byte, []byte, *big.Int, bool, error) {
	return _Fraudproof.Contract.ProcessTx(&_Fraudproof.CallOpts, _balanceRoot, _accountsRoot, _tx, _from_pda_proof, accountProofs)
}

// ProcessTx is a free data retrieval call binding the contract method 0x8cb750a8.
//
// Solidity: function processTx(bytes32 _balanceRoot, bytes32 _accountsRoot, TypesTransaction _tx, TypesPDAMerkleProof _from_pda_proof, TypesAccountProofs accountProofs) view returns(bytes32, bytes, bytes, uint256, bool)
func (_Fraudproof *FraudproofCallerSession) ProcessTx(_balanceRoot [32]byte, _accountsRoot [32]byte, _tx TypesTransaction, _from_pda_proof TypesPDAMerkleProof, accountProofs TypesAccountProofs) ([32]byte, []byte, []byte, *big.Int, bool, error) {
	return _Fraudproof.Contract.ProcessTx(&_Fraudproof.CallOpts, _balanceRoot, _accountsRoot, _tx, _from_pda_proof, accountProofs)
}

// TokenRegistry is a free data retrieval call binding the contract method 0x9d23c4c7.
//
// Solidity: function tokenRegistry() view returns(address)
func (_Fraudproof *FraudproofCaller) TokenRegistry(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Fraudproof.contract.Call(opts, out, "tokenRegistry")
	return *ret0, err
}

// TokenRegistry is a free data retrieval call binding the contract method 0x9d23c4c7.
//
// Solidity: function tokenRegistry() view returns(address)
func (_Fraudproof *FraudproofSession) TokenRegistry() (common.Address, error) {
	return _Fraudproof.Contract.TokenRegistry(&_Fraudproof.CallOpts)
}

// TokenRegistry is a free data retrieval call binding the contract method 0x9d23c4c7.
//
// Solidity: function tokenRegistry() view returns(address)
func (_Fraudproof *FraudproofCallerSession) TokenRegistry() (common.Address, error) {
	return _Fraudproof.Contract.TokenRegistry(&_Fraudproof.CallOpts)
}

// ValidateTxBasic is a free data retrieval call binding the contract method 0x165002a8.
//
// Solidity: function validateTxBasic(TypesTransaction _tx, TypesUserAccount _from_account) view returns(uint256)
func (_Fraudproof *FraudproofCaller) ValidateTxBasic(opts *bind.CallOpts, _tx TypesTransaction, _from_account TypesUserAccount) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Fraudproof.contract.Call(opts, out, "validateTxBasic", _tx, _from_account)
	return *ret0, err
}

// ValidateTxBasic is a free data retrieval call binding the contract method 0x165002a8.
//
// Solidity: function validateTxBasic(TypesTransaction _tx, TypesUserAccount _from_account) view returns(uint256)
func (_Fraudproof *FraudproofSession) ValidateTxBasic(_tx TypesTransaction, _from_account TypesUserAccount) (*big.Int, error) {
	return _Fraudproof.Contract.ValidateTxBasic(&_Fraudproof.CallOpts, _tx, _from_account)
}

// ValidateTxBasic is a free data retrieval call binding the contract method 0x165002a8.
//
// Solidity: function validateTxBasic(TypesTransaction _tx, TypesUserAccount _from_account) view returns(uint256)
func (_Fraudproof *FraudproofCallerSession) ValidateTxBasic(_tx TypesTransaction, _from_account TypesUserAccount) (*big.Int, error) {
	return _Fraudproof.Contract.ValidateTxBasic(&_Fraudproof.CallOpts, _tx, _from_account)
}
