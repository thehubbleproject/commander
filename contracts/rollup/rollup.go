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

// DataTypesAccount is an auto generated low-level Go binding around an user-defined struct.
type DataTypesAccount struct {
	Path      *big.Int
	Balance   *big.Int
	TokenType *big.Int
	Nonce     *big.Int
}

// DataTypesMerkleProof is an auto generated low-level Go binding around an user-defined struct.
type DataTypesMerkleProof struct {
	Account  DataTypesAccount
	Siblings [][32]byte
}

// DataTypesTransaction is an auto generated low-level Go binding around an user-defined struct.
type DataTypesTransaction struct {
	From      DataTypesAccount
	To        DataTypesAccount
	TokenType *big.Int
	Amount    uint32
	Signature []byte
}

// RollupABI is the input ABI used to generate the binding from.
const RollupABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"merkleTreeLib\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"NewAccount\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"txroot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"updatedRoot\",\"type\":\"bytes32\"}],\"name\":\"NewBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32[]\",\"name\":\"to_siblings\",\"type\":\"bytes32[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"to_path\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32[]\",\"name\":\"from_siblings\",\"type\":\"bytes32[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"from_path\",\"type\":\"uint256\"}],\"name\":\"SiblingsGenerated\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"ZERO_BYTES32\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"batches\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"committer\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"txRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"numberOfBatches\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"initAccounts\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"_txs\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"_updatedRoot\",\"type\":\"bytes32\"}],\"name\":\"submitBatch\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_balanceRoot\",\"type\":\"bytes32\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"path\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.Account\",\"name\":\"from\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"path\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.Account\",\"name\":\"to\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"amount\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structDataTypes.Transaction\",\"name\":\"_tx\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"path\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.Account\",\"name\":\"account\",\"type\":\"tuple\"},{\"internalType\":\"bytes32[]\",\"name\":\"siblings\",\"type\":\"bytes32[]\"}],\"internalType\":\"structDataTypes.MerkleProof\",\"name\":\"_from_merkle_proof\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"path\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.Account\",\"name\":\"account\",\"type\":\"tuple\"},{\"internalType\":\"bytes32[]\",\"name\":\"siblings\",\"type\":\"bytes32[]\"}],\"internalType\":\"structDataTypes.MerkleProof\",\"name\":\"_to_merkle_proof\",\"type\":\"tuple\"}],\"name\":\"processTxUpdate\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"path\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.Account\",\"name\":\"account\",\"type\":\"tuple\"}],\"name\":\"getBalanceFromAccount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"path\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.Account\",\"name\":\"original_account\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"new_balance\",\"type\":\"uint256\"}],\"name\":\"updateBalanceInLeaf\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"path\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.Account\",\"name\":\"new_account\",\"type\":\"tuple\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"path\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.Account\",\"name\":\"account\",\"type\":\"tuple\"}],\"name\":\"getAccountBytes\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getBalanceTreeRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

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

// Batches is a free data retrieval call binding the contract method 0xb32c4d8d.
//
// Solidity: function batches(uint256 ) constant returns(bytes32 stateRoot, address committer, bytes32 txRoot, uint256 timestamp)
func (_Rollup *RollupCaller) Batches(opts *bind.CallOpts, arg0 *big.Int) (struct {
	StateRoot [32]byte
	Committer common.Address
	TxRoot    [32]byte
	Timestamp *big.Int
}, error) {
	ret := new(struct {
		StateRoot [32]byte
		Committer common.Address
		TxRoot    [32]byte
		Timestamp *big.Int
	})
	out := ret
	err := _Rollup.contract.Call(opts, out, "batches", arg0)
	return *ret, err
}

// Batches is a free data retrieval call binding the contract method 0xb32c4d8d.
//
// Solidity: function batches(uint256 ) constant returns(bytes32 stateRoot, address committer, bytes32 txRoot, uint256 timestamp)
func (_Rollup *RollupSession) Batches(arg0 *big.Int) (struct {
	StateRoot [32]byte
	Committer common.Address
	TxRoot    [32]byte
	Timestamp *big.Int
}, error) {
	return _Rollup.Contract.Batches(&_Rollup.CallOpts, arg0)
}

// Batches is a free data retrieval call binding the contract method 0xb32c4d8d.
//
// Solidity: function batches(uint256 ) constant returns(bytes32 stateRoot, address committer, bytes32 txRoot, uint256 timestamp)
func (_Rollup *RollupCallerSession) Batches(arg0 *big.Int) (struct {
	StateRoot [32]byte
	Committer common.Address
	TxRoot    [32]byte
	Timestamp *big.Int
}, error) {
	return _Rollup.Contract.Batches(&_Rollup.CallOpts, arg0)
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

// GetAccountBytes is a paid mutator transaction binding the contract method 0x7d3b5c3a.
//
// Solidity: function getAccountBytes(DataTypesAccount account) returns(bytes)
func (_Rollup *RollupTransactor) GetAccountBytes(opts *bind.TransactOpts, account DataTypesAccount) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "getAccountBytes", account)
}

// GetAccountBytes is a paid mutator transaction binding the contract method 0x7d3b5c3a.
//
// Solidity: function getAccountBytes(DataTypesAccount account) returns(bytes)
func (_Rollup *RollupSession) GetAccountBytes(account DataTypesAccount) (*types.Transaction, error) {
	return _Rollup.Contract.GetAccountBytes(&_Rollup.TransactOpts, account)
}

// GetAccountBytes is a paid mutator transaction binding the contract method 0x7d3b5c3a.
//
// Solidity: function getAccountBytes(DataTypesAccount account) returns(bytes)
func (_Rollup *RollupTransactorSession) GetAccountBytes(account DataTypesAccount) (*types.Transaction, error) {
	return _Rollup.Contract.GetAccountBytes(&_Rollup.TransactOpts, account)
}

// GetBalanceFromAccount is a paid mutator transaction binding the contract method 0xbe599579.
//
// Solidity: function getBalanceFromAccount(DataTypesAccount account) returns(uint256)
func (_Rollup *RollupTransactor) GetBalanceFromAccount(opts *bind.TransactOpts, account DataTypesAccount) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "getBalanceFromAccount", account)
}

// GetBalanceFromAccount is a paid mutator transaction binding the contract method 0xbe599579.
//
// Solidity: function getBalanceFromAccount(DataTypesAccount account) returns(uint256)
func (_Rollup *RollupSession) GetBalanceFromAccount(account DataTypesAccount) (*types.Transaction, error) {
	return _Rollup.Contract.GetBalanceFromAccount(&_Rollup.TransactOpts, account)
}

// GetBalanceFromAccount is a paid mutator transaction binding the contract method 0xbe599579.
//
// Solidity: function getBalanceFromAccount(DataTypesAccount account) returns(uint256)
func (_Rollup *RollupTransactorSession) GetBalanceFromAccount(account DataTypesAccount) (*types.Transaction, error) {
	return _Rollup.Contract.GetBalanceFromAccount(&_Rollup.TransactOpts, account)
}

// InitAccounts is a paid mutator transaction binding the contract method 0x49f2545b.
//
// Solidity: function initAccounts() returns()
func (_Rollup *RollupTransactor) InitAccounts(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "initAccounts")
}

// InitAccounts is a paid mutator transaction binding the contract method 0x49f2545b.
//
// Solidity: function initAccounts() returns()
func (_Rollup *RollupSession) InitAccounts() (*types.Transaction, error) {
	return _Rollup.Contract.InitAccounts(&_Rollup.TransactOpts)
}

// InitAccounts is a paid mutator transaction binding the contract method 0x49f2545b.
//
// Solidity: function initAccounts() returns()
func (_Rollup *RollupTransactorSession) InitAccounts() (*types.Transaction, error) {
	return _Rollup.Contract.InitAccounts(&_Rollup.TransactOpts)
}

// ProcessTxUpdate is a paid mutator transaction binding the contract method 0x592dcef0.
//
// Solidity: function processTxUpdate(bytes32 _balanceRoot, DataTypesTransaction _tx, DataTypesMerkleProof _from_merkle_proof, DataTypesMerkleProof _to_merkle_proof) returns(bytes32, uint256, uint256)
func (_Rollup *RollupTransactor) ProcessTxUpdate(opts *bind.TransactOpts, _balanceRoot [32]byte, _tx DataTypesTransaction, _from_merkle_proof DataTypesMerkleProof, _to_merkle_proof DataTypesMerkleProof) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "processTxUpdate", _balanceRoot, _tx, _from_merkle_proof, _to_merkle_proof)
}

// ProcessTxUpdate is a paid mutator transaction binding the contract method 0x592dcef0.
//
// Solidity: function processTxUpdate(bytes32 _balanceRoot, DataTypesTransaction _tx, DataTypesMerkleProof _from_merkle_proof, DataTypesMerkleProof _to_merkle_proof) returns(bytes32, uint256, uint256)
func (_Rollup *RollupSession) ProcessTxUpdate(_balanceRoot [32]byte, _tx DataTypesTransaction, _from_merkle_proof DataTypesMerkleProof, _to_merkle_proof DataTypesMerkleProof) (*types.Transaction, error) {
	return _Rollup.Contract.ProcessTxUpdate(&_Rollup.TransactOpts, _balanceRoot, _tx, _from_merkle_proof, _to_merkle_proof)
}

// ProcessTxUpdate is a paid mutator transaction binding the contract method 0x592dcef0.
//
// Solidity: function processTxUpdate(bytes32 _balanceRoot, DataTypesTransaction _tx, DataTypesMerkleProof _from_merkle_proof, DataTypesMerkleProof _to_merkle_proof) returns(bytes32, uint256, uint256)
func (_Rollup *RollupTransactorSession) ProcessTxUpdate(_balanceRoot [32]byte, _tx DataTypesTransaction, _from_merkle_proof DataTypesMerkleProof, _to_merkle_proof DataTypesMerkleProof) (*types.Transaction, error) {
	return _Rollup.Contract.ProcessTxUpdate(&_Rollup.TransactOpts, _balanceRoot, _tx, _from_merkle_proof, _to_merkle_proof)
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

// UpdateBalanceInLeaf is a paid mutator transaction binding the contract method 0xe255f0c2.
//
// Solidity: function updateBalanceInLeaf(DataTypesAccount original_account, uint256 new_balance) returns(DataTypesAccount new_account)
func (_Rollup *RollupTransactor) UpdateBalanceInLeaf(opts *bind.TransactOpts, original_account DataTypesAccount, new_balance *big.Int) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "updateBalanceInLeaf", original_account, new_balance)
}

// UpdateBalanceInLeaf is a paid mutator transaction binding the contract method 0xe255f0c2.
//
// Solidity: function updateBalanceInLeaf(DataTypesAccount original_account, uint256 new_balance) returns(DataTypesAccount new_account)
func (_Rollup *RollupSession) UpdateBalanceInLeaf(original_account DataTypesAccount, new_balance *big.Int) (*types.Transaction, error) {
	return _Rollup.Contract.UpdateBalanceInLeaf(&_Rollup.TransactOpts, original_account, new_balance)
}

// UpdateBalanceInLeaf is a paid mutator transaction binding the contract method 0xe255f0c2.
//
// Solidity: function updateBalanceInLeaf(DataTypesAccount original_account, uint256 new_balance) returns(DataTypesAccount new_account)
func (_Rollup *RollupTransactorSession) UpdateBalanceInLeaf(original_account DataTypesAccount, new_balance *big.Int) (*types.Transaction, error) {
	return _Rollup.Contract.UpdateBalanceInLeaf(&_Rollup.TransactOpts, original_account, new_balance)
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
	Txroot      [32]byte
	UpdatedRoot [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterNewBatch is a free log retrieval operation binding the contract event 0xa7104c958bfef2602dc7918ca62e74b832ed15c2c06945db1300ef1e0715ada8.
//
// Solidity: event NewBatch(bytes32 txroot, bytes32 updatedRoot)
func (_Rollup *RollupFilterer) FilterNewBatch(opts *bind.FilterOpts) (*RollupNewBatchIterator, error) {

	logs, sub, err := _Rollup.contract.FilterLogs(opts, "NewBatch")
	if err != nil {
		return nil, err
	}
	return &RollupNewBatchIterator{contract: _Rollup.contract, event: "NewBatch", logs: logs, sub: sub}, nil
}

// WatchNewBatch is a free log subscription operation binding the contract event 0xa7104c958bfef2602dc7918ca62e74b832ed15c2c06945db1300ef1e0715ada8.
//
// Solidity: event NewBatch(bytes32 txroot, bytes32 updatedRoot)
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

// ParseNewBatch is a log parse operation binding the contract event 0xa7104c958bfef2602dc7918ca62e74b832ed15c2c06945db1300ef1e0715ada8.
//
// Solidity: event NewBatch(bytes32 txroot, bytes32 updatedRoot)
func (_Rollup *RollupFilterer) ParseNewBatch(log types.Log) (*RollupNewBatch, error) {
	event := new(RollupNewBatch)
	if err := _Rollup.contract.UnpackLog(event, "NewBatch", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RollupSiblingsGeneratedIterator is returned from FilterSiblingsGenerated and is used to iterate over the raw logs and unpacked data for SiblingsGenerated events raised by the Rollup contract.
type RollupSiblingsGeneratedIterator struct {
	Event *RollupSiblingsGenerated // Event containing the contract specifics and raw log

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
func (it *RollupSiblingsGeneratedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupSiblingsGenerated)
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
		it.Event = new(RollupSiblingsGenerated)
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
func (it *RollupSiblingsGeneratedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupSiblingsGeneratedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupSiblingsGenerated represents a SiblingsGenerated event raised by the Rollup contract.
type RollupSiblingsGenerated struct {
	ToSiblings   [][32]byte
	ToPath       *big.Int
	FromSiblings [][32]byte
	FromPath     *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterSiblingsGenerated is a free log retrieval operation binding the contract event 0x4a1ebd1ec289fff3c24a3775b2e9d999af4865038d164896e73e0a6b298c494b.
//
// Solidity: event SiblingsGenerated(bytes32[] to_siblings, uint256 to_path, bytes32[] from_siblings, uint256 from_path)
func (_Rollup *RollupFilterer) FilterSiblingsGenerated(opts *bind.FilterOpts) (*RollupSiblingsGeneratedIterator, error) {

	logs, sub, err := _Rollup.contract.FilterLogs(opts, "SiblingsGenerated")
	if err != nil {
		return nil, err
	}
	return &RollupSiblingsGeneratedIterator{contract: _Rollup.contract, event: "SiblingsGenerated", logs: logs, sub: sub}, nil
}

// WatchSiblingsGenerated is a free log subscription operation binding the contract event 0x4a1ebd1ec289fff3c24a3775b2e9d999af4865038d164896e73e0a6b298c494b.
//
// Solidity: event SiblingsGenerated(bytes32[] to_siblings, uint256 to_path, bytes32[] from_siblings, uint256 from_path)
func (_Rollup *RollupFilterer) WatchSiblingsGenerated(opts *bind.WatchOpts, sink chan<- *RollupSiblingsGenerated) (event.Subscription, error) {

	logs, sub, err := _Rollup.contract.WatchLogs(opts, "SiblingsGenerated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupSiblingsGenerated)
				if err := _Rollup.contract.UnpackLog(event, "SiblingsGenerated", log); err != nil {
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

// ParseSiblingsGenerated is a log parse operation binding the contract event 0x4a1ebd1ec289fff3c24a3775b2e9d999af4865038d164896e73e0a6b298c494b.
//
// Solidity: event SiblingsGenerated(bytes32[] to_siblings, uint256 to_path, bytes32[] from_siblings, uint256 from_path)
func (_Rollup *RollupFilterer) ParseSiblingsGenerated(log types.Log) (*RollupSiblingsGenerated, error) {
	event := new(RollupSiblingsGenerated)
	if err := _Rollup.contract.UnpackLog(event, "SiblingsGenerated", log); err != nil {
		return nil, err
	}
	return event, nil
}
