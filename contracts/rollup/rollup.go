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

// RollupABI is the input ABI used to generate the binding from.
const RollupABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"ZERO_BYTES32\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"tx_sig\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"tx_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tx_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tx_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"proof_from\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"proof_to\",\"type\":\"bytes\"}],\"name\":\"updateTx\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"_txs\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"updatedRoot\",\"type\":\"bytes32\"}],\"name\":\"submitBatch\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"batches\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"withdraw_root\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"committer\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"account_tree_state\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"txRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"GENESIS_BALANCE_TREE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ACCOUNT_TREE_STATE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"merkleTreeLib\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"txroot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"updatedRoot\",\"type\":\"bytes32\"}],\"name\":\"NewBatch\",\"type\":\"event\"}]"

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

// ACCOUNTTREESTATE is a free data retrieval call binding the contract method 0xecf24838.
//
// Solidity: function ACCOUNT_TREE_STATE() constant returns(bytes32)
func (_Rollup *RollupCaller) ACCOUNTTREESTATE(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Rollup.contract.Call(opts, out, "ACCOUNT_TREE_STATE")
	return *ret0, err
}

// ACCOUNTTREESTATE is a free data retrieval call binding the contract method 0xecf24838.
//
// Solidity: function ACCOUNT_TREE_STATE() constant returns(bytes32)
func (_Rollup *RollupSession) ACCOUNTTREESTATE() ([32]byte, error) {
	return _Rollup.Contract.ACCOUNTTREESTATE(&_Rollup.CallOpts)
}

// ACCOUNTTREESTATE is a free data retrieval call binding the contract method 0xecf24838.
//
// Solidity: function ACCOUNT_TREE_STATE() constant returns(bytes32)
func (_Rollup *RollupCallerSession) ACCOUNTTREESTATE() ([32]byte, error) {
	return _Rollup.Contract.ACCOUNTTREESTATE(&_Rollup.CallOpts)
}

// GENESISBALANCETREE is a free data retrieval call binding the contract method 0xc6a6996a.
//
// Solidity: function GENESIS_BALANCE_TREE() constant returns(bytes32)
func (_Rollup *RollupCaller) GENESISBALANCETREE(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Rollup.contract.Call(opts, out, "GENESIS_BALANCE_TREE")
	return *ret0, err
}

// GENESISBALANCETREE is a free data retrieval call binding the contract method 0xc6a6996a.
//
// Solidity: function GENESIS_BALANCE_TREE() constant returns(bytes32)
func (_Rollup *RollupSession) GENESISBALANCETREE() ([32]byte, error) {
	return _Rollup.Contract.GENESISBALANCETREE(&_Rollup.CallOpts)
}

// GENESISBALANCETREE is a free data retrieval call binding the contract method 0xc6a6996a.
//
// Solidity: function GENESIS_BALANCE_TREE() constant returns(bytes32)
func (_Rollup *RollupCallerSession) GENESISBALANCETREE() ([32]byte, error) {
	return _Rollup.Contract.GENESISBALANCETREE(&_Rollup.CallOpts)
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
// Solidity: function batches(uint256 ) constant returns(bytes32 stateRoot, bytes32 withdraw_root, address committer, bytes32 account_tree_state, bytes32 txRoot, uint256 timestamp)
func (_Rollup *RollupCaller) Batches(opts *bind.CallOpts, arg0 *big.Int) (struct {
	StateRoot        [32]byte
	WithdrawRoot     [32]byte
	Committer        common.Address
	AccountTreeState [32]byte
	TxRoot           [32]byte
	Timestamp        *big.Int
}, error) {
	ret := new(struct {
		StateRoot        [32]byte
		WithdrawRoot     [32]byte
		Committer        common.Address
		AccountTreeState [32]byte
		TxRoot           [32]byte
		Timestamp        *big.Int
	})
	out := ret
	err := _Rollup.contract.Call(opts, out, "batches", arg0)
	return *ret, err
}

// Batches is a free data retrieval call binding the contract method 0xb32c4d8d.
//
// Solidity: function batches(uint256 ) constant returns(bytes32 stateRoot, bytes32 withdraw_root, address committer, bytes32 account_tree_state, bytes32 txRoot, uint256 timestamp)
func (_Rollup *RollupSession) Batches(arg0 *big.Int) (struct {
	StateRoot        [32]byte
	WithdrawRoot     [32]byte
	Committer        common.Address
	AccountTreeState [32]byte
	TxRoot           [32]byte
	Timestamp        *big.Int
}, error) {
	return _Rollup.Contract.Batches(&_Rollup.CallOpts, arg0)
}

// Batches is a free data retrieval call binding the contract method 0xb32c4d8d.
//
// Solidity: function batches(uint256 ) constant returns(bytes32 stateRoot, bytes32 withdraw_root, address committer, bytes32 account_tree_state, bytes32 txRoot, uint256 timestamp)
func (_Rollup *RollupCallerSession) Batches(arg0 *big.Int) (struct {
	StateRoot        [32]byte
	WithdrawRoot     [32]byte
	Committer        common.Address
	AccountTreeState [32]byte
	TxRoot           [32]byte
	Timestamp        *big.Int
}, error) {
	return _Rollup.Contract.Batches(&_Rollup.CallOpts, arg0)
}

// SubmitBatch is a paid mutator transaction binding the contract method 0x0e981757.
//
// Solidity: function submitBatch(bytes[] _txs, bytes32 updatedRoot) returns(bytes32)
func (_Rollup *RollupTransactor) SubmitBatch(opts *bind.TransactOpts, _txs [][]byte, updatedRoot [32]byte) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "submitBatch", _txs, updatedRoot)
}

// SubmitBatch is a paid mutator transaction binding the contract method 0x0e981757.
//
// Solidity: function submitBatch(bytes[] _txs, bytes32 updatedRoot) returns(bytes32)
func (_Rollup *RollupSession) SubmitBatch(_txs [][]byte, updatedRoot [32]byte) (*types.Transaction, error) {
	return _Rollup.Contract.SubmitBatch(&_Rollup.TransactOpts, _txs, updatedRoot)
}

// SubmitBatch is a paid mutator transaction binding the contract method 0x0e981757.
//
// Solidity: function submitBatch(bytes[] _txs, bytes32 updatedRoot) returns(bytes32)
func (_Rollup *RollupTransactorSession) SubmitBatch(_txs [][]byte, updatedRoot [32]byte) (*types.Transaction, error) {
	return _Rollup.Contract.SubmitBatch(&_Rollup.TransactOpts, _txs, updatedRoot)
}

// UpdateTx is a paid mutator transaction binding the contract method 0x0736a825.
//
// Solidity: function updateTx(bytes tx_sig, address tx_from, address tx_to, uint256 tx_amount, bytes proof_from, bytes proof_to) returns()
func (_Rollup *RollupTransactor) UpdateTx(opts *bind.TransactOpts, tx_sig []byte, tx_from common.Address, tx_to common.Address, tx_amount *big.Int, proof_from []byte, proof_to []byte) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "updateTx", tx_sig, tx_from, tx_to, tx_amount, proof_from, proof_to)
}

// UpdateTx is a paid mutator transaction binding the contract method 0x0736a825.
//
// Solidity: function updateTx(bytes tx_sig, address tx_from, address tx_to, uint256 tx_amount, bytes proof_from, bytes proof_to) returns()
func (_Rollup *RollupSession) UpdateTx(tx_sig []byte, tx_from common.Address, tx_to common.Address, tx_amount *big.Int, proof_from []byte, proof_to []byte) (*types.Transaction, error) {
	return _Rollup.Contract.UpdateTx(&_Rollup.TransactOpts, tx_sig, tx_from, tx_to, tx_amount, proof_from, proof_to)
}

// UpdateTx is a paid mutator transaction binding the contract method 0x0736a825.
//
// Solidity: function updateTx(bytes tx_sig, address tx_from, address tx_to, uint256 tx_amount, bytes proof_from, bytes proof_to) returns()
func (_Rollup *RollupTransactorSession) UpdateTx(tx_sig []byte, tx_from common.Address, tx_to common.Address, tx_amount *big.Int, proof_from []byte, proof_to []byte) (*types.Transaction, error) {
	return _Rollup.Contract.UpdateTx(&_Rollup.TransactOpts, tx_sig, tx_from, tx_to, tx_amount, proof_from, proof_to)
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
