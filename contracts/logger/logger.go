// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package trial

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

// TrialABI is the input ABI used to generate the binding from.
const TrialABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"batch_id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"committer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"txRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stakeSlashed\",\"type\":\"uint256\"}],\"name\":\"BatchRollback\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"DepositLeafMerged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"token\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"accountHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"}],\"name\":\"DepositQueued\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"DepositsProcessed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"committer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"txroot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"updatedRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"NewBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tokenContract\",\"type\":\"address\"}],\"name\":\"RegisteredToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tokenContract\",\"type\":\"address\"}],\"name\":\"RegistrationRequest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalBatchesSlashed\",\"type\":\"uint256\"}],\"name\":\"RollbackFinalisation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"committed\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"batch_id\",\"type\":\"uint256\"}],\"name\":\"StakeWithdraw\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"committer\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"txroot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"updatedRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"logNewBatch\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"committed\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"batch_id\",\"type\":\"uint256\"}],\"name\":\"logStakeWithdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"batch_id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"committer\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"txRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"stakeSlashed\",\"type\":\"uint256\"}],\"name\":\"logBatchRollback\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"totalBatchesSlashed\",\"type\":\"uint256\"}],\"name\":\"logRollbackFinalisation\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenContract\",\"type\":\"address\"}],\"name\":\"logRegisteredToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenContract\",\"type\":\"address\"}],\"name\":\"logRegistrationRequest\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"token\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"accountHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"}],\"name\":\"logDepositQueued\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"logDepositLeafMerged\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"logDepositsProcessed\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Trial is an auto generated Go binding around an Ethereum contract.
type Trial struct {
	TrialCaller     // Read-only binding to the contract
	TrialTransactor // Write-only binding to the contract
	TrialFilterer   // Log filterer for contract events
}

// TrialCaller is an auto generated read-only Go binding around an Ethereum contract.
type TrialCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TrialTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TrialTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TrialFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TrialFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TrialSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TrialSession struct {
	Contract     *Trial            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TrialCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TrialCallerSession struct {
	Contract *TrialCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TrialTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TrialTransactorSession struct {
	Contract     *TrialTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TrialRaw is an auto generated low-level Go binding around an Ethereum contract.
type TrialRaw struct {
	Contract *Trial // Generic contract binding to access the raw methods on
}

// TrialCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TrialCallerRaw struct {
	Contract *TrialCaller // Generic read-only contract binding to access the raw methods on
}

// TrialTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TrialTransactorRaw struct {
	Contract *TrialTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTrial creates a new instance of Trial, bound to a specific deployed contract.
func NewTrial(address common.Address, backend bind.ContractBackend) (*Trial, error) {
	contract, err := bindTrial(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Trial{TrialCaller: TrialCaller{contract: contract}, TrialTransactor: TrialTransactor{contract: contract}, TrialFilterer: TrialFilterer{contract: contract}}, nil
}

// NewTrialCaller creates a new read-only instance of Trial, bound to a specific deployed contract.
func NewTrialCaller(address common.Address, caller bind.ContractCaller) (*TrialCaller, error) {
	contract, err := bindTrial(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TrialCaller{contract: contract}, nil
}

// NewTrialTransactor creates a new write-only instance of Trial, bound to a specific deployed contract.
func NewTrialTransactor(address common.Address, transactor bind.ContractTransactor) (*TrialTransactor, error) {
	contract, err := bindTrial(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TrialTransactor{contract: contract}, nil
}

// NewTrialFilterer creates a new log filterer instance of Trial, bound to a specific deployed contract.
func NewTrialFilterer(address common.Address, filterer bind.ContractFilterer) (*TrialFilterer, error) {
	contract, err := bindTrial(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TrialFilterer{contract: contract}, nil
}

// bindTrial binds a generic wrapper to an already deployed contract.
func bindTrial(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TrialABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Trial *TrialRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Trial.Contract.TrialCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Trial *TrialRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Trial.Contract.TrialTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Trial *TrialRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Trial.Contract.TrialTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Trial *TrialCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Trial.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Trial *TrialTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Trial.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Trial *TrialTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Trial.Contract.contract.Transact(opts, method, params...)
}

// LogBatchRollback is a paid mutator transaction binding the contract method 0xb4ba86b3.
//
// Solidity: function logBatchRollback(uint256 batch_id, address committer, bytes32 stateRoot, bytes32 txRoot, uint256 stakeSlashed) returns()
func (_Trial *TrialTransactor) LogBatchRollback(opts *bind.TransactOpts, batch_id *big.Int, committer common.Address, stateRoot [32]byte, txRoot [32]byte, stakeSlashed *big.Int) (*types.Transaction, error) {
	return _Trial.contract.Transact(opts, "logBatchRollback", batch_id, committer, stateRoot, txRoot, stakeSlashed)
}

// LogBatchRollback is a paid mutator transaction binding the contract method 0xb4ba86b3.
//
// Solidity: function logBatchRollback(uint256 batch_id, address committer, bytes32 stateRoot, bytes32 txRoot, uint256 stakeSlashed) returns()
func (_Trial *TrialSession) LogBatchRollback(batch_id *big.Int, committer common.Address, stateRoot [32]byte, txRoot [32]byte, stakeSlashed *big.Int) (*types.Transaction, error) {
	return _Trial.Contract.LogBatchRollback(&_Trial.TransactOpts, batch_id, committer, stateRoot, txRoot, stakeSlashed)
}

// LogBatchRollback is a paid mutator transaction binding the contract method 0xb4ba86b3.
//
// Solidity: function logBatchRollback(uint256 batch_id, address committer, bytes32 stateRoot, bytes32 txRoot, uint256 stakeSlashed) returns()
func (_Trial *TrialTransactorSession) LogBatchRollback(batch_id *big.Int, committer common.Address, stateRoot [32]byte, txRoot [32]byte, stakeSlashed *big.Int) (*types.Transaction, error) {
	return _Trial.Contract.LogBatchRollback(&_Trial.TransactOpts, batch_id, committer, stateRoot, txRoot, stakeSlashed)
}

// LogDepositLeafMerged is a paid mutator transaction binding the contract method 0xc5abe3c1.
//
// Solidity: function logDepositLeafMerged() returns()
func (_Trial *TrialTransactor) LogDepositLeafMerged(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Trial.contract.Transact(opts, "logDepositLeafMerged")
}

// LogDepositLeafMerged is a paid mutator transaction binding the contract method 0xc5abe3c1.
//
// Solidity: function logDepositLeafMerged() returns()
func (_Trial *TrialSession) LogDepositLeafMerged() (*types.Transaction, error) {
	return _Trial.Contract.LogDepositLeafMerged(&_Trial.TransactOpts)
}

// LogDepositLeafMerged is a paid mutator transaction binding the contract method 0xc5abe3c1.
//
// Solidity: function logDepositLeafMerged() returns()
func (_Trial *TrialTransactorSession) LogDepositLeafMerged() (*types.Transaction, error) {
	return _Trial.Contract.LogDepositLeafMerged(&_Trial.TransactOpts)
}

// LogDepositQueued is a paid mutator transaction binding the contract method 0x10e4bacc.
//
// Solidity: function logDepositQueued(address destination, uint256 amount, uint256 token, bytes32 accountHash, bytes pubkey) returns()
func (_Trial *TrialTransactor) LogDepositQueued(opts *bind.TransactOpts, destination common.Address, amount *big.Int, token *big.Int, accountHash [32]byte, pubkey []byte) (*types.Transaction, error) {
	return _Trial.contract.Transact(opts, "logDepositQueued", destination, amount, token, accountHash, pubkey)
}

// LogDepositQueued is a paid mutator transaction binding the contract method 0x10e4bacc.
//
// Solidity: function logDepositQueued(address destination, uint256 amount, uint256 token, bytes32 accountHash, bytes pubkey) returns()
func (_Trial *TrialSession) LogDepositQueued(destination common.Address, amount *big.Int, token *big.Int, accountHash [32]byte, pubkey []byte) (*types.Transaction, error) {
	return _Trial.Contract.LogDepositQueued(&_Trial.TransactOpts, destination, amount, token, accountHash, pubkey)
}

// LogDepositQueued is a paid mutator transaction binding the contract method 0x10e4bacc.
//
// Solidity: function logDepositQueued(address destination, uint256 amount, uint256 token, bytes32 accountHash, bytes pubkey) returns()
func (_Trial *TrialTransactorSession) LogDepositQueued(destination common.Address, amount *big.Int, token *big.Int, accountHash [32]byte, pubkey []byte) (*types.Transaction, error) {
	return _Trial.Contract.LogDepositQueued(&_Trial.TransactOpts, destination, amount, token, accountHash, pubkey)
}

// LogDepositsProcessed is a paid mutator transaction binding the contract method 0xd45cff88.
//
// Solidity: function logDepositsProcessed() returns()
func (_Trial *TrialTransactor) LogDepositsProcessed(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Trial.contract.Transact(opts, "logDepositsProcessed")
}

// LogDepositsProcessed is a paid mutator transaction binding the contract method 0xd45cff88.
//
// Solidity: function logDepositsProcessed() returns()
func (_Trial *TrialSession) LogDepositsProcessed() (*types.Transaction, error) {
	return _Trial.Contract.LogDepositsProcessed(&_Trial.TransactOpts)
}

// LogDepositsProcessed is a paid mutator transaction binding the contract method 0xd45cff88.
//
// Solidity: function logDepositsProcessed() returns()
func (_Trial *TrialTransactorSession) LogDepositsProcessed() (*types.Transaction, error) {
	return _Trial.Contract.LogDepositsProcessed(&_Trial.TransactOpts)
}

// LogNewBatch is a paid mutator transaction binding the contract method 0xcea4592f.
//
// Solidity: function logNewBatch(address committer, bytes32 txroot, bytes32 updatedRoot, uint256 index) returns()
func (_Trial *TrialTransactor) LogNewBatch(opts *bind.TransactOpts, committer common.Address, txroot [32]byte, updatedRoot [32]byte, index *big.Int) (*types.Transaction, error) {
	return _Trial.contract.Transact(opts, "logNewBatch", committer, txroot, updatedRoot, index)
}

// LogNewBatch is a paid mutator transaction binding the contract method 0xcea4592f.
//
// Solidity: function logNewBatch(address committer, bytes32 txroot, bytes32 updatedRoot, uint256 index) returns()
func (_Trial *TrialSession) LogNewBatch(committer common.Address, txroot [32]byte, updatedRoot [32]byte, index *big.Int) (*types.Transaction, error) {
	return _Trial.Contract.LogNewBatch(&_Trial.TransactOpts, committer, txroot, updatedRoot, index)
}

// LogNewBatch is a paid mutator transaction binding the contract method 0xcea4592f.
//
// Solidity: function logNewBatch(address committer, bytes32 txroot, bytes32 updatedRoot, uint256 index) returns()
func (_Trial *TrialTransactorSession) LogNewBatch(committer common.Address, txroot [32]byte, updatedRoot [32]byte, index *big.Int) (*types.Transaction, error) {
	return _Trial.Contract.LogNewBatch(&_Trial.TransactOpts, committer, txroot, updatedRoot, index)
}

// LogRegisteredToken is a paid mutator transaction binding the contract method 0x88870639.
//
// Solidity: function logRegisteredToken(uint256 tokenType, address tokenContract) returns()
func (_Trial *TrialTransactor) LogRegisteredToken(opts *bind.TransactOpts, tokenType *big.Int, tokenContract common.Address) (*types.Transaction, error) {
	return _Trial.contract.Transact(opts, "logRegisteredToken", tokenType, tokenContract)
}

// LogRegisteredToken is a paid mutator transaction binding the contract method 0x88870639.
//
// Solidity: function logRegisteredToken(uint256 tokenType, address tokenContract) returns()
func (_Trial *TrialSession) LogRegisteredToken(tokenType *big.Int, tokenContract common.Address) (*types.Transaction, error) {
	return _Trial.Contract.LogRegisteredToken(&_Trial.TransactOpts, tokenType, tokenContract)
}

// LogRegisteredToken is a paid mutator transaction binding the contract method 0x88870639.
//
// Solidity: function logRegisteredToken(uint256 tokenType, address tokenContract) returns()
func (_Trial *TrialTransactorSession) LogRegisteredToken(tokenType *big.Int, tokenContract common.Address) (*types.Transaction, error) {
	return _Trial.Contract.LogRegisteredToken(&_Trial.TransactOpts, tokenType, tokenContract)
}

// LogRegistrationRequest is a paid mutator transaction binding the contract method 0x4c7637d2.
//
// Solidity: function logRegistrationRequest(address tokenContract) returns()
func (_Trial *TrialTransactor) LogRegistrationRequest(opts *bind.TransactOpts, tokenContract common.Address) (*types.Transaction, error) {
	return _Trial.contract.Transact(opts, "logRegistrationRequest", tokenContract)
}

// LogRegistrationRequest is a paid mutator transaction binding the contract method 0x4c7637d2.
//
// Solidity: function logRegistrationRequest(address tokenContract) returns()
func (_Trial *TrialSession) LogRegistrationRequest(tokenContract common.Address) (*types.Transaction, error) {
	return _Trial.Contract.LogRegistrationRequest(&_Trial.TransactOpts, tokenContract)
}

// LogRegistrationRequest is a paid mutator transaction binding the contract method 0x4c7637d2.
//
// Solidity: function logRegistrationRequest(address tokenContract) returns()
func (_Trial *TrialTransactorSession) LogRegistrationRequest(tokenContract common.Address) (*types.Transaction, error) {
	return _Trial.Contract.LogRegistrationRequest(&_Trial.TransactOpts, tokenContract)
}

// LogRollbackFinalisation is a paid mutator transaction binding the contract method 0xaeedf0c7.
//
// Solidity: function logRollbackFinalisation(uint256 totalBatchesSlashed) returns()
func (_Trial *TrialTransactor) LogRollbackFinalisation(opts *bind.TransactOpts, totalBatchesSlashed *big.Int) (*types.Transaction, error) {
	return _Trial.contract.Transact(opts, "logRollbackFinalisation", totalBatchesSlashed)
}

// LogRollbackFinalisation is a paid mutator transaction binding the contract method 0xaeedf0c7.
//
// Solidity: function logRollbackFinalisation(uint256 totalBatchesSlashed) returns()
func (_Trial *TrialSession) LogRollbackFinalisation(totalBatchesSlashed *big.Int) (*types.Transaction, error) {
	return _Trial.Contract.LogRollbackFinalisation(&_Trial.TransactOpts, totalBatchesSlashed)
}

// LogRollbackFinalisation is a paid mutator transaction binding the contract method 0xaeedf0c7.
//
// Solidity: function logRollbackFinalisation(uint256 totalBatchesSlashed) returns()
func (_Trial *TrialTransactorSession) LogRollbackFinalisation(totalBatchesSlashed *big.Int) (*types.Transaction, error) {
	return _Trial.Contract.LogRollbackFinalisation(&_Trial.TransactOpts, totalBatchesSlashed)
}

// LogStakeWithdraw is a paid mutator transaction binding the contract method 0xd54404ae.
//
// Solidity: function logStakeWithdraw(address committed, uint256 amount, uint256 batch_id) returns()
func (_Trial *TrialTransactor) LogStakeWithdraw(opts *bind.TransactOpts, committed common.Address, amount *big.Int, batch_id *big.Int) (*types.Transaction, error) {
	return _Trial.contract.Transact(opts, "logStakeWithdraw", committed, amount, batch_id)
}

// LogStakeWithdraw is a paid mutator transaction binding the contract method 0xd54404ae.
//
// Solidity: function logStakeWithdraw(address committed, uint256 amount, uint256 batch_id) returns()
func (_Trial *TrialSession) LogStakeWithdraw(committed common.Address, amount *big.Int, batch_id *big.Int) (*types.Transaction, error) {
	return _Trial.Contract.LogStakeWithdraw(&_Trial.TransactOpts, committed, amount, batch_id)
}

// LogStakeWithdraw is a paid mutator transaction binding the contract method 0xd54404ae.
//
// Solidity: function logStakeWithdraw(address committed, uint256 amount, uint256 batch_id) returns()
func (_Trial *TrialTransactorSession) LogStakeWithdraw(committed common.Address, amount *big.Int, batch_id *big.Int) (*types.Transaction, error) {
	return _Trial.Contract.LogStakeWithdraw(&_Trial.TransactOpts, committed, amount, batch_id)
}

// TrialBatchRollbackIterator is returned from FilterBatchRollback and is used to iterate over the raw logs and unpacked data for BatchRollback events raised by the Trial contract.
type TrialBatchRollbackIterator struct {
	Event *TrialBatchRollback // Event containing the contract specifics and raw log

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
func (it *TrialBatchRollbackIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TrialBatchRollback)
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
		it.Event = new(TrialBatchRollback)
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
func (it *TrialBatchRollbackIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TrialBatchRollbackIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TrialBatchRollback represents a BatchRollback event raised by the Trial contract.
type TrialBatchRollback struct {
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
func (_Trial *TrialFilterer) FilterBatchRollback(opts *bind.FilterOpts) (*TrialBatchRollbackIterator, error) {

	logs, sub, err := _Trial.contract.FilterLogs(opts, "BatchRollback")
	if err != nil {
		return nil, err
	}
	return &TrialBatchRollbackIterator{contract: _Trial.contract, event: "BatchRollback", logs: logs, sub: sub}, nil
}

// WatchBatchRollback is a free log subscription operation binding the contract event 0xff0d01a3ea09eec8d9dc11b606f70fbf5aa373d5647db792e5886fc4285c38fc.
//
// Solidity: event BatchRollback(uint256 batch_id, address committer, bytes32 stateRoot, bytes32 txRoot, uint256 stakeSlashed)
func (_Trial *TrialFilterer) WatchBatchRollback(opts *bind.WatchOpts, sink chan<- *TrialBatchRollback) (event.Subscription, error) {

	logs, sub, err := _Trial.contract.WatchLogs(opts, "BatchRollback")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TrialBatchRollback)
				if err := _Trial.contract.UnpackLog(event, "BatchRollback", log); err != nil {
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
func (_Trial *TrialFilterer) ParseBatchRollback(log types.Log) (*TrialBatchRollback, error) {
	event := new(TrialBatchRollback)
	if err := _Trial.contract.UnpackLog(event, "BatchRollback", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TrialDepositLeafMergedIterator is returned from FilterDepositLeafMerged and is used to iterate over the raw logs and unpacked data for DepositLeafMerged events raised by the Trial contract.
type TrialDepositLeafMergedIterator struct {
	Event *TrialDepositLeafMerged // Event containing the contract specifics and raw log

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
func (it *TrialDepositLeafMergedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TrialDepositLeafMerged)
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
		it.Event = new(TrialDepositLeafMerged)
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
func (it *TrialDepositLeafMergedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TrialDepositLeafMergedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TrialDepositLeafMerged represents a DepositLeafMerged event raised by the Trial contract.
type TrialDepositLeafMerged struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDepositLeafMerged is a free log retrieval operation binding the contract event 0xe08d4467ae1c0717bba6e77edc5bc5394ea260848bee1f763aba623e6e1dc864.
//
// Solidity: event DepositLeafMerged()
func (_Trial *TrialFilterer) FilterDepositLeafMerged(opts *bind.FilterOpts) (*TrialDepositLeafMergedIterator, error) {

	logs, sub, err := _Trial.contract.FilterLogs(opts, "DepositLeafMerged")
	if err != nil {
		return nil, err
	}
	return &TrialDepositLeafMergedIterator{contract: _Trial.contract, event: "DepositLeafMerged", logs: logs, sub: sub}, nil
}

// WatchDepositLeafMerged is a free log subscription operation binding the contract event 0xe08d4467ae1c0717bba6e77edc5bc5394ea260848bee1f763aba623e6e1dc864.
//
// Solidity: event DepositLeafMerged()
func (_Trial *TrialFilterer) WatchDepositLeafMerged(opts *bind.WatchOpts, sink chan<- *TrialDepositLeafMerged) (event.Subscription, error) {

	logs, sub, err := _Trial.contract.WatchLogs(opts, "DepositLeafMerged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TrialDepositLeafMerged)
				if err := _Trial.contract.UnpackLog(event, "DepositLeafMerged", log); err != nil {
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
func (_Trial *TrialFilterer) ParseDepositLeafMerged(log types.Log) (*TrialDepositLeafMerged, error) {
	event := new(TrialDepositLeafMerged)
	if err := _Trial.contract.UnpackLog(event, "DepositLeafMerged", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TrialDepositQueuedIterator is returned from FilterDepositQueued and is used to iterate over the raw logs and unpacked data for DepositQueued events raised by the Trial contract.
type TrialDepositQueuedIterator struct {
	Event *TrialDepositQueued // Event containing the contract specifics and raw log

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
func (it *TrialDepositQueuedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TrialDepositQueued)
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
		it.Event = new(TrialDepositQueued)
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
func (it *TrialDepositQueuedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TrialDepositQueuedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TrialDepositQueued represents a DepositQueued event raised by the Trial contract.
type TrialDepositQueued struct {
	Destination common.Address
	Amount      *big.Int
	Token       *big.Int
	AccountHash [32]byte
	Pubkey      []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterDepositQueued is a free log retrieval operation binding the contract event 0xcb241fd5af8b0d42828bb79881d3965c2443fd98dc31f73604f538aa7a609c9b.
//
// Solidity: event DepositQueued(address destination, uint256 amount, uint256 token, bytes32 accountHash, bytes pubkey)
func (_Trial *TrialFilterer) FilterDepositQueued(opts *bind.FilterOpts) (*TrialDepositQueuedIterator, error) {

	logs, sub, err := _Trial.contract.FilterLogs(opts, "DepositQueued")
	if err != nil {
		return nil, err
	}
	return &TrialDepositQueuedIterator{contract: _Trial.contract, event: "DepositQueued", logs: logs, sub: sub}, nil
}

// WatchDepositQueued is a free log subscription operation binding the contract event 0xcb241fd5af8b0d42828bb79881d3965c2443fd98dc31f73604f538aa7a609c9b.
//
// Solidity: event DepositQueued(address destination, uint256 amount, uint256 token, bytes32 accountHash, bytes pubkey)
func (_Trial *TrialFilterer) WatchDepositQueued(opts *bind.WatchOpts, sink chan<- *TrialDepositQueued) (event.Subscription, error) {

	logs, sub, err := _Trial.contract.WatchLogs(opts, "DepositQueued")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TrialDepositQueued)
				if err := _Trial.contract.UnpackLog(event, "DepositQueued", log); err != nil {
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

// ParseDepositQueued is a log parse operation binding the contract event 0xcb241fd5af8b0d42828bb79881d3965c2443fd98dc31f73604f538aa7a609c9b.
//
// Solidity: event DepositQueued(address destination, uint256 amount, uint256 token, bytes32 accountHash, bytes pubkey)
func (_Trial *TrialFilterer) ParseDepositQueued(log types.Log) (*TrialDepositQueued, error) {
	event := new(TrialDepositQueued)
	if err := _Trial.contract.UnpackLog(event, "DepositQueued", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TrialDepositsProcessedIterator is returned from FilterDepositsProcessed and is used to iterate over the raw logs and unpacked data for DepositsProcessed events raised by the Trial contract.
type TrialDepositsProcessedIterator struct {
	Event *TrialDepositsProcessed // Event containing the contract specifics and raw log

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
func (it *TrialDepositsProcessedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TrialDepositsProcessed)
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
		it.Event = new(TrialDepositsProcessed)
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
func (it *TrialDepositsProcessedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TrialDepositsProcessedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TrialDepositsProcessed represents a DepositsProcessed event raised by the Trial contract.
type TrialDepositsProcessed struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDepositsProcessed is a free log retrieval operation binding the contract event 0x663ea3b3fa88c877f70d364fe14a3b85f410a60df236737b37fcc0796225fe33.
//
// Solidity: event DepositsProcessed()
func (_Trial *TrialFilterer) FilterDepositsProcessed(opts *bind.FilterOpts) (*TrialDepositsProcessedIterator, error) {

	logs, sub, err := _Trial.contract.FilterLogs(opts, "DepositsProcessed")
	if err != nil {
		return nil, err
	}
	return &TrialDepositsProcessedIterator{contract: _Trial.contract, event: "DepositsProcessed", logs: logs, sub: sub}, nil
}

// WatchDepositsProcessed is a free log subscription operation binding the contract event 0x663ea3b3fa88c877f70d364fe14a3b85f410a60df236737b37fcc0796225fe33.
//
// Solidity: event DepositsProcessed()
func (_Trial *TrialFilterer) WatchDepositsProcessed(opts *bind.WatchOpts, sink chan<- *TrialDepositsProcessed) (event.Subscription, error) {

	logs, sub, err := _Trial.contract.WatchLogs(opts, "DepositsProcessed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TrialDepositsProcessed)
				if err := _Trial.contract.UnpackLog(event, "DepositsProcessed", log); err != nil {
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
func (_Trial *TrialFilterer) ParseDepositsProcessed(log types.Log) (*TrialDepositsProcessed, error) {
	event := new(TrialDepositsProcessed)
	if err := _Trial.contract.UnpackLog(event, "DepositsProcessed", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TrialNewBatchIterator is returned from FilterNewBatch and is used to iterate over the raw logs and unpacked data for NewBatch events raised by the Trial contract.
type TrialNewBatchIterator struct {
	Event *TrialNewBatch // Event containing the contract specifics and raw log

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
func (it *TrialNewBatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TrialNewBatch)
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
		it.Event = new(TrialNewBatch)
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
func (it *TrialNewBatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TrialNewBatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TrialNewBatch represents a NewBatch event raised by the Trial contract.
type TrialNewBatch struct {
	Committer   common.Address
	Txroot      [32]byte
	UpdatedRoot [32]byte
	Index       *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterNewBatch is a free log retrieval operation binding the contract event 0x1914387dd72c107202016df3d09390ff777fa6399bf12f0e5a0ed11a0b0650bb.
//
// Solidity: event NewBatch(address committer, bytes32 txroot, bytes32 updatedRoot, uint256 index)
func (_Trial *TrialFilterer) FilterNewBatch(opts *bind.FilterOpts) (*TrialNewBatchIterator, error) {

	logs, sub, err := _Trial.contract.FilterLogs(opts, "NewBatch")
	if err != nil {
		return nil, err
	}
	return &TrialNewBatchIterator{contract: _Trial.contract, event: "NewBatch", logs: logs, sub: sub}, nil
}

// WatchNewBatch is a free log subscription operation binding the contract event 0x1914387dd72c107202016df3d09390ff777fa6399bf12f0e5a0ed11a0b0650bb.
//
// Solidity: event NewBatch(address committer, bytes32 txroot, bytes32 updatedRoot, uint256 index)
func (_Trial *TrialFilterer) WatchNewBatch(opts *bind.WatchOpts, sink chan<- *TrialNewBatch) (event.Subscription, error) {

	logs, sub, err := _Trial.contract.WatchLogs(opts, "NewBatch")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TrialNewBatch)
				if err := _Trial.contract.UnpackLog(event, "NewBatch", log); err != nil {
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

// ParseNewBatch is a log parse operation binding the contract event 0x1914387dd72c107202016df3d09390ff777fa6399bf12f0e5a0ed11a0b0650bb.
//
// Solidity: event NewBatch(address committer, bytes32 txroot, bytes32 updatedRoot, uint256 index)
func (_Trial *TrialFilterer) ParseNewBatch(log types.Log) (*TrialNewBatch, error) {
	event := new(TrialNewBatch)
	if err := _Trial.contract.UnpackLog(event, "NewBatch", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TrialRegisteredTokenIterator is returned from FilterRegisteredToken and is used to iterate over the raw logs and unpacked data for RegisteredToken events raised by the Trial contract.
type TrialRegisteredTokenIterator struct {
	Event *TrialRegisteredToken // Event containing the contract specifics and raw log

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
func (it *TrialRegisteredTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TrialRegisteredToken)
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
		it.Event = new(TrialRegisteredToken)
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
func (it *TrialRegisteredTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TrialRegisteredTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TrialRegisteredToken represents a RegisteredToken event raised by the Trial contract.
type TrialRegisteredToken struct {
	TokenType     *big.Int
	TokenContract common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterRegisteredToken is a free log retrieval operation binding the contract event 0x5dbaa701a7acef513f72a61799f7e50f4653f462b9f780d88d1b9bec89de2168.
//
// Solidity: event RegisteredToken(uint256 tokenType, address tokenContract)
func (_Trial *TrialFilterer) FilterRegisteredToken(opts *bind.FilterOpts) (*TrialRegisteredTokenIterator, error) {

	logs, sub, err := _Trial.contract.FilterLogs(opts, "RegisteredToken")
	if err != nil {
		return nil, err
	}
	return &TrialRegisteredTokenIterator{contract: _Trial.contract, event: "RegisteredToken", logs: logs, sub: sub}, nil
}

// WatchRegisteredToken is a free log subscription operation binding the contract event 0x5dbaa701a7acef513f72a61799f7e50f4653f462b9f780d88d1b9bec89de2168.
//
// Solidity: event RegisteredToken(uint256 tokenType, address tokenContract)
func (_Trial *TrialFilterer) WatchRegisteredToken(opts *bind.WatchOpts, sink chan<- *TrialRegisteredToken) (event.Subscription, error) {

	logs, sub, err := _Trial.contract.WatchLogs(opts, "RegisteredToken")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TrialRegisteredToken)
				if err := _Trial.contract.UnpackLog(event, "RegisteredToken", log); err != nil {
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
func (_Trial *TrialFilterer) ParseRegisteredToken(log types.Log) (*TrialRegisteredToken, error) {
	event := new(TrialRegisteredToken)
	if err := _Trial.contract.UnpackLog(event, "RegisteredToken", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TrialRegistrationRequestIterator is returned from FilterRegistrationRequest and is used to iterate over the raw logs and unpacked data for RegistrationRequest events raised by the Trial contract.
type TrialRegistrationRequestIterator struct {
	Event *TrialRegistrationRequest // Event containing the contract specifics and raw log

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
func (it *TrialRegistrationRequestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TrialRegistrationRequest)
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
		it.Event = new(TrialRegistrationRequest)
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
func (it *TrialRegistrationRequestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TrialRegistrationRequestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TrialRegistrationRequest represents a RegistrationRequest event raised by the Trial contract.
type TrialRegistrationRequest struct {
	TokenContract common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterRegistrationRequest is a free log retrieval operation binding the contract event 0xdc79fc57451962cfe3916e686997a49229af75ce2055deb4c0f0fdf3d5d2e7c1.
//
// Solidity: event RegistrationRequest(address tokenContract)
func (_Trial *TrialFilterer) FilterRegistrationRequest(opts *bind.FilterOpts) (*TrialRegistrationRequestIterator, error) {

	logs, sub, err := _Trial.contract.FilterLogs(opts, "RegistrationRequest")
	if err != nil {
		return nil, err
	}
	return &TrialRegistrationRequestIterator{contract: _Trial.contract, event: "RegistrationRequest", logs: logs, sub: sub}, nil
}

// WatchRegistrationRequest is a free log subscription operation binding the contract event 0xdc79fc57451962cfe3916e686997a49229af75ce2055deb4c0f0fdf3d5d2e7c1.
//
// Solidity: event RegistrationRequest(address tokenContract)
func (_Trial *TrialFilterer) WatchRegistrationRequest(opts *bind.WatchOpts, sink chan<- *TrialRegistrationRequest) (event.Subscription, error) {

	logs, sub, err := _Trial.contract.WatchLogs(opts, "RegistrationRequest")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TrialRegistrationRequest)
				if err := _Trial.contract.UnpackLog(event, "RegistrationRequest", log); err != nil {
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
func (_Trial *TrialFilterer) ParseRegistrationRequest(log types.Log) (*TrialRegistrationRequest, error) {
	event := new(TrialRegistrationRequest)
	if err := _Trial.contract.UnpackLog(event, "RegistrationRequest", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TrialRollbackFinalisationIterator is returned from FilterRollbackFinalisation and is used to iterate over the raw logs and unpacked data for RollbackFinalisation events raised by the Trial contract.
type TrialRollbackFinalisationIterator struct {
	Event *TrialRollbackFinalisation // Event containing the contract specifics and raw log

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
func (it *TrialRollbackFinalisationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TrialRollbackFinalisation)
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
		it.Event = new(TrialRollbackFinalisation)
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
func (it *TrialRollbackFinalisationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TrialRollbackFinalisationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TrialRollbackFinalisation represents a RollbackFinalisation event raised by the Trial contract.
type TrialRollbackFinalisation struct {
	TotalBatchesSlashed *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterRollbackFinalisation is a free log retrieval operation binding the contract event 0x8efd02dfe309f172ea08425236aa43cec05abc25ebbb9cd4ed1de0d5048fc91a.
//
// Solidity: event RollbackFinalisation(uint256 totalBatchesSlashed)
func (_Trial *TrialFilterer) FilterRollbackFinalisation(opts *bind.FilterOpts) (*TrialRollbackFinalisationIterator, error) {

	logs, sub, err := _Trial.contract.FilterLogs(opts, "RollbackFinalisation")
	if err != nil {
		return nil, err
	}
	return &TrialRollbackFinalisationIterator{contract: _Trial.contract, event: "RollbackFinalisation", logs: logs, sub: sub}, nil
}

// WatchRollbackFinalisation is a free log subscription operation binding the contract event 0x8efd02dfe309f172ea08425236aa43cec05abc25ebbb9cd4ed1de0d5048fc91a.
//
// Solidity: event RollbackFinalisation(uint256 totalBatchesSlashed)
func (_Trial *TrialFilterer) WatchRollbackFinalisation(opts *bind.WatchOpts, sink chan<- *TrialRollbackFinalisation) (event.Subscription, error) {

	logs, sub, err := _Trial.contract.WatchLogs(opts, "RollbackFinalisation")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TrialRollbackFinalisation)
				if err := _Trial.contract.UnpackLog(event, "RollbackFinalisation", log); err != nil {
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
func (_Trial *TrialFilterer) ParseRollbackFinalisation(log types.Log) (*TrialRollbackFinalisation, error) {
	event := new(TrialRollbackFinalisation)
	if err := _Trial.contract.UnpackLog(event, "RollbackFinalisation", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TrialStakeWithdrawIterator is returned from FilterStakeWithdraw and is used to iterate over the raw logs and unpacked data for StakeWithdraw events raised by the Trial contract.
type TrialStakeWithdrawIterator struct {
	Event *TrialStakeWithdraw // Event containing the contract specifics and raw log

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
func (it *TrialStakeWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TrialStakeWithdraw)
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
		it.Event = new(TrialStakeWithdraw)
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
func (it *TrialStakeWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TrialStakeWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TrialStakeWithdraw represents a StakeWithdraw event raised by the Trial contract.
type TrialStakeWithdraw struct {
	Committed common.Address
	Amount    *big.Int
	BatchId   *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterStakeWithdraw is a free log retrieval operation binding the contract event 0xb00cb3d8b4494806e139df1041902f1984526ffe765fa4d7e12d64f4a47362fe.
//
// Solidity: event StakeWithdraw(address committed, uint256 amount, uint256 batch_id)
func (_Trial *TrialFilterer) FilterStakeWithdraw(opts *bind.FilterOpts) (*TrialStakeWithdrawIterator, error) {

	logs, sub, err := _Trial.contract.FilterLogs(opts, "StakeWithdraw")
	if err != nil {
		return nil, err
	}
	return &TrialStakeWithdrawIterator{contract: _Trial.contract, event: "StakeWithdraw", logs: logs, sub: sub}, nil
}

// WatchStakeWithdraw is a free log subscription operation binding the contract event 0xb00cb3d8b4494806e139df1041902f1984526ffe765fa4d7e12d64f4a47362fe.
//
// Solidity: event StakeWithdraw(address committed, uint256 amount, uint256 batch_id)
func (_Trial *TrialFilterer) WatchStakeWithdraw(opts *bind.WatchOpts, sink chan<- *TrialStakeWithdraw) (event.Subscription, error) {

	logs, sub, err := _Trial.contract.WatchLogs(opts, "StakeWithdraw")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TrialStakeWithdraw)
				if err := _Trial.contract.UnpackLog(event, "StakeWithdraw", log); err != nil {
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
func (_Trial *TrialFilterer) ParseStakeWithdraw(log types.Log) (*TrialStakeWithdraw, error) {
	event := new(TrialStakeWithdraw)
	if err := _Trial.contract.UnpackLog(event, "StakeWithdraw", log); err != nil {
		return nil, err
	}
	return event, nil
}
