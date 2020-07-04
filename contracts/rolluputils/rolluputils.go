// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package rolluputils

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

// TypesPDALeaf is an auto generated low-level Go binding around an user-defined struct.
type TypesPDALeaf struct {
	Pubkey []byte
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

// RolluputilsABI is the input ABI used to generate the binding from.
const RolluputilsABI = "[{\"constant\":true,\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"}],\"internalType\":\"structTypes.PDALeaf\",\"name\":\"_PDA_Leaf\",\"type\":\"tuple\"}],\"name\":\"PDALeafToHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"ID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structTypes.UserAccount\",\"name\":\"original_account\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"new_balance\",\"type\":\"uint256\"}],\"name\":\"UpdateBalanceInAccount\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"ID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structTypes.UserAccount\",\"name\":\"updated_account\",\"type\":\"tuple\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"ID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structTypes.UserAccount\",\"name\":\"account\",\"type\":\"tuple\"}],\"name\":\"BalanceFromAccount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"accountBytes\",\"type\":\"bytes\"}],\"name\":\"AccountFromBytes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"ID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"ID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structTypes.UserAccount\",\"name\":\"account\",\"type\":\"tuple\"}],\"name\":\"BytesFromAccount\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"ID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"}],\"name\":\"BytesFromAccountDeconstructed\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"}],\"name\":\"getAccountHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"ID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structTypes.UserAccount\",\"name\":\"account\",\"type\":\"tuple\"}],\"name\":\"HashFromAccount\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"txType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structTypes.Transaction\",\"name\":\"_tx\",\"type\":\"tuple\"}],\"name\":\"CompressTx\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"txBytes\",\"type\":\"bytes\"}],\"name\":\"TxFromBytes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"from\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"to\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"txType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"txType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structTypes.Transaction\",\"name\":\"_tx\",\"type\":\"tuple\"}],\"name\":\"BytesFromTx\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"from\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"to\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"txType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"BytesFromTxDeconstructed\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"txType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structTypes.Transaction\",\"name\":\"_tx\",\"type\":\"tuple\"}],\"name\":\"HashFromTx\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fromIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"txType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"getTxSignBytes\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"pub\",\"type\":\"bytes\"}],\"name\":\"calculateAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// Rolluputils is an auto generated Go binding around an Ethereum contract.
type Rolluputils struct {
	RolluputilsCaller     // Read-only binding to the contract
	RolluputilsTransactor // Write-only binding to the contract
	RolluputilsFilterer   // Log filterer for contract events
}

// RolluputilsCaller is an auto generated read-only Go binding around an Ethereum contract.
type RolluputilsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RolluputilsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RolluputilsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RolluputilsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RolluputilsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RolluputilsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RolluputilsSession struct {
	Contract     *Rolluputils      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RolluputilsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RolluputilsCallerSession struct {
	Contract *RolluputilsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// RolluputilsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RolluputilsTransactorSession struct {
	Contract     *RolluputilsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// RolluputilsRaw is an auto generated low-level Go binding around an Ethereum contract.
type RolluputilsRaw struct {
	Contract *Rolluputils // Generic contract binding to access the raw methods on
}

// RolluputilsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RolluputilsCallerRaw struct {
	Contract *RolluputilsCaller // Generic read-only contract binding to access the raw methods on
}

// RolluputilsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RolluputilsTransactorRaw struct {
	Contract *RolluputilsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRolluputils creates a new instance of Rolluputils, bound to a specific deployed contract.
func NewRolluputils(address common.Address, backend bind.ContractBackend) (*Rolluputils, error) {
	contract, err := bindRolluputils(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Rolluputils{RolluputilsCaller: RolluputilsCaller{contract: contract}, RolluputilsTransactor: RolluputilsTransactor{contract: contract}, RolluputilsFilterer: RolluputilsFilterer{contract: contract}}, nil
}

// NewRolluputilsCaller creates a new read-only instance of Rolluputils, bound to a specific deployed contract.
func NewRolluputilsCaller(address common.Address, caller bind.ContractCaller) (*RolluputilsCaller, error) {
	contract, err := bindRolluputils(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RolluputilsCaller{contract: contract}, nil
}

// NewRolluputilsTransactor creates a new write-only instance of Rolluputils, bound to a specific deployed contract.
func NewRolluputilsTransactor(address common.Address, transactor bind.ContractTransactor) (*RolluputilsTransactor, error) {
	contract, err := bindRolluputils(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RolluputilsTransactor{contract: contract}, nil
}

// NewRolluputilsFilterer creates a new log filterer instance of Rolluputils, bound to a specific deployed contract.
func NewRolluputilsFilterer(address common.Address, filterer bind.ContractFilterer) (*RolluputilsFilterer, error) {
	contract, err := bindRolluputils(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RolluputilsFilterer{contract: contract}, nil
}

// bindRolluputils binds a generic wrapper to an already deployed contract.
func bindRolluputils(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RolluputilsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Rolluputils *RolluputilsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Rolluputils.Contract.RolluputilsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Rolluputils *RolluputilsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Rolluputils.Contract.RolluputilsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Rolluputils *RolluputilsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Rolluputils.Contract.RolluputilsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Rolluputils *RolluputilsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Rolluputils.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Rolluputils *RolluputilsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Rolluputils.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Rolluputils *RolluputilsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Rolluputils.Contract.contract.Transact(opts, method, params...)
}

// AccountFromBytes is a free data retrieval call binding the contract method 0x1a636e86.
//
// Solidity: function AccountFromBytes(bytes accountBytes) pure returns(uint256 ID, uint256 balance, uint256 nonce, uint256 tokenType)
func (_Rolluputils *RolluputilsCaller) AccountFromBytes(opts *bind.CallOpts, accountBytes []byte) (struct {
	ID        *big.Int
	Balance   *big.Int
	Nonce     *big.Int
	TokenType *big.Int
}, error) {
	ret := new(struct {
		ID        *big.Int
		Balance   *big.Int
		Nonce     *big.Int
		TokenType *big.Int
	})
	out := ret
	err := _Rolluputils.contract.Call(opts, out, "AccountFromBytes", accountBytes)
	return *ret, err
}

// AccountFromBytes is a free data retrieval call binding the contract method 0x1a636e86.
//
// Solidity: function AccountFromBytes(bytes accountBytes) pure returns(uint256 ID, uint256 balance, uint256 nonce, uint256 tokenType)
func (_Rolluputils *RolluputilsSession) AccountFromBytes(accountBytes []byte) (struct {
	ID        *big.Int
	Balance   *big.Int
	Nonce     *big.Int
	TokenType *big.Int
}, error) {
	return _Rolluputils.Contract.AccountFromBytes(&_Rolluputils.CallOpts, accountBytes)
}

// AccountFromBytes is a free data retrieval call binding the contract method 0x1a636e86.
//
// Solidity: function AccountFromBytes(bytes accountBytes) pure returns(uint256 ID, uint256 balance, uint256 nonce, uint256 tokenType)
func (_Rolluputils *RolluputilsCallerSession) AccountFromBytes(accountBytes []byte) (struct {
	ID        *big.Int
	Balance   *big.Int
	Nonce     *big.Int
	TokenType *big.Int
}, error) {
	return _Rolluputils.Contract.AccountFromBytes(&_Rolluputils.CallOpts, accountBytes)
}

// BalanceFromAccount is a free data retrieval call binding the contract method 0x0f9b2cb2.
//
// Solidity: function BalanceFromAccount(TypesUserAccount account) pure returns(uint256)
func (_Rolluputils *RolluputilsCaller) BalanceFromAccount(opts *bind.CallOpts, account TypesUserAccount) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Rolluputils.contract.Call(opts, out, "BalanceFromAccount", account)
	return *ret0, err
}

// BalanceFromAccount is a free data retrieval call binding the contract method 0x0f9b2cb2.
//
// Solidity: function BalanceFromAccount(TypesUserAccount account) pure returns(uint256)
func (_Rolluputils *RolluputilsSession) BalanceFromAccount(account TypesUserAccount) (*big.Int, error) {
	return _Rolluputils.Contract.BalanceFromAccount(&_Rolluputils.CallOpts, account)
}

// BalanceFromAccount is a free data retrieval call binding the contract method 0x0f9b2cb2.
//
// Solidity: function BalanceFromAccount(TypesUserAccount account) pure returns(uint256)
func (_Rolluputils *RolluputilsCallerSession) BalanceFromAccount(account TypesUserAccount) (*big.Int, error) {
	return _Rolluputils.Contract.BalanceFromAccount(&_Rolluputils.CallOpts, account)
}

// BytesFromAccount is a free data retrieval call binding the contract method 0x3035226f.
//
// Solidity: function BytesFromAccount(TypesUserAccount account) pure returns(bytes)
func (_Rolluputils *RolluputilsCaller) BytesFromAccount(opts *bind.CallOpts, account TypesUserAccount) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _Rolluputils.contract.Call(opts, out, "BytesFromAccount", account)
	return *ret0, err
}

// BytesFromAccount is a free data retrieval call binding the contract method 0x3035226f.
//
// Solidity: function BytesFromAccount(TypesUserAccount account) pure returns(bytes)
func (_Rolluputils *RolluputilsSession) BytesFromAccount(account TypesUserAccount) ([]byte, error) {
	return _Rolluputils.Contract.BytesFromAccount(&_Rolluputils.CallOpts, account)
}

// BytesFromAccount is a free data retrieval call binding the contract method 0x3035226f.
//
// Solidity: function BytesFromAccount(TypesUserAccount account) pure returns(bytes)
func (_Rolluputils *RolluputilsCallerSession) BytesFromAccount(account TypesUserAccount) ([]byte, error) {
	return _Rolluputils.Contract.BytesFromAccount(&_Rolluputils.CallOpts, account)
}

// BytesFromAccountDeconstructed is a free data retrieval call binding the contract method 0x5bc64a2b.
//
// Solidity: function BytesFromAccountDeconstructed(uint256 ID, uint256 balance, uint256 nonce, uint256 tokenType) pure returns(bytes)
func (_Rolluputils *RolluputilsCaller) BytesFromAccountDeconstructed(opts *bind.CallOpts, ID *big.Int, balance *big.Int, nonce *big.Int, tokenType *big.Int) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _Rolluputils.contract.Call(opts, out, "BytesFromAccountDeconstructed", ID, balance, nonce, tokenType)
	return *ret0, err
}

// BytesFromAccountDeconstructed is a free data retrieval call binding the contract method 0x5bc64a2b.
//
// Solidity: function BytesFromAccountDeconstructed(uint256 ID, uint256 balance, uint256 nonce, uint256 tokenType) pure returns(bytes)
func (_Rolluputils *RolluputilsSession) BytesFromAccountDeconstructed(ID *big.Int, balance *big.Int, nonce *big.Int, tokenType *big.Int) ([]byte, error) {
	return _Rolluputils.Contract.BytesFromAccountDeconstructed(&_Rolluputils.CallOpts, ID, balance, nonce, tokenType)
}

// BytesFromAccountDeconstructed is a free data retrieval call binding the contract method 0x5bc64a2b.
//
// Solidity: function BytesFromAccountDeconstructed(uint256 ID, uint256 balance, uint256 nonce, uint256 tokenType) pure returns(bytes)
func (_Rolluputils *RolluputilsCallerSession) BytesFromAccountDeconstructed(ID *big.Int, balance *big.Int, nonce *big.Int, tokenType *big.Int) ([]byte, error) {
	return _Rolluputils.Contract.BytesFromAccountDeconstructed(&_Rolluputils.CallOpts, ID, balance, nonce, tokenType)
}

// BytesFromTx is a free data retrieval call binding the contract method 0xd0dbc902.
//
// Solidity: function BytesFromTx(TypesTransaction _tx) pure returns(bytes)
func (_Rolluputils *RolluputilsCaller) BytesFromTx(opts *bind.CallOpts, _tx TypesTransaction) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _Rolluputils.contract.Call(opts, out, "BytesFromTx", _tx)
	return *ret0, err
}

// BytesFromTx is a free data retrieval call binding the contract method 0xd0dbc902.
//
// Solidity: function BytesFromTx(TypesTransaction _tx) pure returns(bytes)
func (_Rolluputils *RolluputilsSession) BytesFromTx(_tx TypesTransaction) ([]byte, error) {
	return _Rolluputils.Contract.BytesFromTx(&_Rolluputils.CallOpts, _tx)
}

// BytesFromTx is a free data retrieval call binding the contract method 0xd0dbc902.
//
// Solidity: function BytesFromTx(TypesTransaction _tx) pure returns(bytes)
func (_Rolluputils *RolluputilsCallerSession) BytesFromTx(_tx TypesTransaction) ([]byte, error) {
	return _Rolluputils.Contract.BytesFromTx(&_Rolluputils.CallOpts, _tx)
}

// BytesFromTxDeconstructed is a free data retrieval call binding the contract method 0xb1b84d99.
//
// Solidity: function BytesFromTxDeconstructed(uint256 from, uint256 to, uint256 tokenType, uint256 nonce, uint256 txType, uint256 amount) pure returns(bytes)
func (_Rolluputils *RolluputilsCaller) BytesFromTxDeconstructed(opts *bind.CallOpts, from *big.Int, to *big.Int, tokenType *big.Int, nonce *big.Int, txType *big.Int, amount *big.Int) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _Rolluputils.contract.Call(opts, out, "BytesFromTxDeconstructed", from, to, tokenType, nonce, txType, amount)
	return *ret0, err
}

// BytesFromTxDeconstructed is a free data retrieval call binding the contract method 0xb1b84d99.
//
// Solidity: function BytesFromTxDeconstructed(uint256 from, uint256 to, uint256 tokenType, uint256 nonce, uint256 txType, uint256 amount) pure returns(bytes)
func (_Rolluputils *RolluputilsSession) BytesFromTxDeconstructed(from *big.Int, to *big.Int, tokenType *big.Int, nonce *big.Int, txType *big.Int, amount *big.Int) ([]byte, error) {
	return _Rolluputils.Contract.BytesFromTxDeconstructed(&_Rolluputils.CallOpts, from, to, tokenType, nonce, txType, amount)
}

// BytesFromTxDeconstructed is a free data retrieval call binding the contract method 0xb1b84d99.
//
// Solidity: function BytesFromTxDeconstructed(uint256 from, uint256 to, uint256 tokenType, uint256 nonce, uint256 txType, uint256 amount) pure returns(bytes)
func (_Rolluputils *RolluputilsCallerSession) BytesFromTxDeconstructed(from *big.Int, to *big.Int, tokenType *big.Int, nonce *big.Int, txType *big.Int, amount *big.Int) ([]byte, error) {
	return _Rolluputils.Contract.BytesFromTxDeconstructed(&_Rolluputils.CallOpts, from, to, tokenType, nonce, txType, amount)
}

// CompressTx is a free data retrieval call binding the contract method 0x02c36512.
//
// Solidity: function CompressTx(TypesTransaction _tx) pure returns(bytes)
func (_Rolluputils *RolluputilsCaller) CompressTx(opts *bind.CallOpts, _tx TypesTransaction) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _Rolluputils.contract.Call(opts, out, "CompressTx", _tx)
	return *ret0, err
}

// CompressTx is a free data retrieval call binding the contract method 0x02c36512.
//
// Solidity: function CompressTx(TypesTransaction _tx) pure returns(bytes)
func (_Rolluputils *RolluputilsSession) CompressTx(_tx TypesTransaction) ([]byte, error) {
	return _Rolluputils.Contract.CompressTx(&_Rolluputils.CallOpts, _tx)
}

// CompressTx is a free data retrieval call binding the contract method 0x02c36512.
//
// Solidity: function CompressTx(TypesTransaction _tx) pure returns(bytes)
func (_Rolluputils *RolluputilsCallerSession) CompressTx(_tx TypesTransaction) ([]byte, error) {
	return _Rolluputils.Contract.CompressTx(&_Rolluputils.CallOpts, _tx)
}

// HashFromAccount is a free data retrieval call binding the contract method 0xcadbd919.
//
// Solidity: function HashFromAccount(TypesUserAccount account) pure returns(bytes32)
func (_Rolluputils *RolluputilsCaller) HashFromAccount(opts *bind.CallOpts, account TypesUserAccount) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Rolluputils.contract.Call(opts, out, "HashFromAccount", account)
	return *ret0, err
}

// HashFromAccount is a free data retrieval call binding the contract method 0xcadbd919.
//
// Solidity: function HashFromAccount(TypesUserAccount account) pure returns(bytes32)
func (_Rolluputils *RolluputilsSession) HashFromAccount(account TypesUserAccount) ([32]byte, error) {
	return _Rolluputils.Contract.HashFromAccount(&_Rolluputils.CallOpts, account)
}

// HashFromAccount is a free data retrieval call binding the contract method 0xcadbd919.
//
// Solidity: function HashFromAccount(TypesUserAccount account) pure returns(bytes32)
func (_Rolluputils *RolluputilsCallerSession) HashFromAccount(account TypesUserAccount) ([32]byte, error) {
	return _Rolluputils.Contract.HashFromAccount(&_Rolluputils.CallOpts, account)
}

// HashFromTx is a free data retrieval call binding the contract method 0xb90cbf51.
//
// Solidity: function HashFromTx(TypesTransaction _tx) pure returns(bytes32)
func (_Rolluputils *RolluputilsCaller) HashFromTx(opts *bind.CallOpts, _tx TypesTransaction) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Rolluputils.contract.Call(opts, out, "HashFromTx", _tx)
	return *ret0, err
}

// HashFromTx is a free data retrieval call binding the contract method 0xb90cbf51.
//
// Solidity: function HashFromTx(TypesTransaction _tx) pure returns(bytes32)
func (_Rolluputils *RolluputilsSession) HashFromTx(_tx TypesTransaction) ([32]byte, error) {
	return _Rolluputils.Contract.HashFromTx(&_Rolluputils.CallOpts, _tx)
}

// HashFromTx is a free data retrieval call binding the contract method 0xb90cbf51.
//
// Solidity: function HashFromTx(TypesTransaction _tx) pure returns(bytes32)
func (_Rolluputils *RolluputilsCallerSession) HashFromTx(_tx TypesTransaction) ([32]byte, error) {
	return _Rolluputils.Contract.HashFromTx(&_Rolluputils.CallOpts, _tx)
}

// PDALeafToHash is a free data retrieval call binding the contract method 0xc2ddab33.
//
// Solidity: function PDALeafToHash(TypesPDALeaf _PDA_Leaf) pure returns(bytes32)
func (_Rolluputils *RolluputilsCaller) PDALeafToHash(opts *bind.CallOpts, _PDA_Leaf TypesPDALeaf) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Rolluputils.contract.Call(opts, out, "PDALeafToHash", _PDA_Leaf)
	return *ret0, err
}

// PDALeafToHash is a free data retrieval call binding the contract method 0xc2ddab33.
//
// Solidity: function PDALeafToHash(TypesPDALeaf _PDA_Leaf) pure returns(bytes32)
func (_Rolluputils *RolluputilsSession) PDALeafToHash(_PDA_Leaf TypesPDALeaf) ([32]byte, error) {
	return _Rolluputils.Contract.PDALeafToHash(&_Rolluputils.CallOpts, _PDA_Leaf)
}

// PDALeafToHash is a free data retrieval call binding the contract method 0xc2ddab33.
//
// Solidity: function PDALeafToHash(TypesPDALeaf _PDA_Leaf) pure returns(bytes32)
func (_Rolluputils *RolluputilsCallerSession) PDALeafToHash(_PDA_Leaf TypesPDALeaf) ([32]byte, error) {
	return _Rolluputils.Contract.PDALeafToHash(&_Rolluputils.CallOpts, _PDA_Leaf)
}

// TxFromBytes is a free data retrieval call binding the contract method 0xbdbf417a.
//
// Solidity: function TxFromBytes(bytes txBytes) pure returns(uint256 from, uint256 to, uint256 tokenType, uint256 nonce, uint256 txType, uint256 amount)
func (_Rolluputils *RolluputilsCaller) TxFromBytes(opts *bind.CallOpts, txBytes []byte) (struct {
	From      *big.Int
	To        *big.Int
	TokenType *big.Int
	Nonce     *big.Int
	TxType    *big.Int
	Amount    *big.Int
}, error) {
	ret := new(struct {
		From      *big.Int
		To        *big.Int
		TokenType *big.Int
		Nonce     *big.Int
		TxType    *big.Int
		Amount    *big.Int
	})
	out := ret
	err := _Rolluputils.contract.Call(opts, out, "TxFromBytes", txBytes)
	return *ret, err
}

// TxFromBytes is a free data retrieval call binding the contract method 0xbdbf417a.
//
// Solidity: function TxFromBytes(bytes txBytes) pure returns(uint256 from, uint256 to, uint256 tokenType, uint256 nonce, uint256 txType, uint256 amount)
func (_Rolluputils *RolluputilsSession) TxFromBytes(txBytes []byte) (struct {
	From      *big.Int
	To        *big.Int
	TokenType *big.Int
	Nonce     *big.Int
	TxType    *big.Int
	Amount    *big.Int
}, error) {
	return _Rolluputils.Contract.TxFromBytes(&_Rolluputils.CallOpts, txBytes)
}

// TxFromBytes is a free data retrieval call binding the contract method 0xbdbf417a.
//
// Solidity: function TxFromBytes(bytes txBytes) pure returns(uint256 from, uint256 to, uint256 tokenType, uint256 nonce, uint256 txType, uint256 amount)
func (_Rolluputils *RolluputilsCallerSession) TxFromBytes(txBytes []byte) (struct {
	From      *big.Int
	To        *big.Int
	TokenType *big.Int
	Nonce     *big.Int
	TxType    *big.Int
	Amount    *big.Int
}, error) {
	return _Rolluputils.Contract.TxFromBytes(&_Rolluputils.CallOpts, txBytes)
}

// UpdateBalanceInAccount is a free data retrieval call binding the contract method 0x2e9be49a.
//
// Solidity: function UpdateBalanceInAccount(TypesUserAccount original_account, uint256 new_balance) pure returns(TypesUserAccount updated_account)
func (_Rolluputils *RolluputilsCaller) UpdateBalanceInAccount(opts *bind.CallOpts, original_account TypesUserAccount, new_balance *big.Int) (TypesUserAccount, error) {
	var (
		ret0 = new(TypesUserAccount)
	)
	out := ret0
	err := _Rolluputils.contract.Call(opts, out, "UpdateBalanceInAccount", original_account, new_balance)
	return *ret0, err
}

// UpdateBalanceInAccount is a free data retrieval call binding the contract method 0x2e9be49a.
//
// Solidity: function UpdateBalanceInAccount(TypesUserAccount original_account, uint256 new_balance) pure returns(TypesUserAccount updated_account)
func (_Rolluputils *RolluputilsSession) UpdateBalanceInAccount(original_account TypesUserAccount, new_balance *big.Int) (TypesUserAccount, error) {
	return _Rolluputils.Contract.UpdateBalanceInAccount(&_Rolluputils.CallOpts, original_account, new_balance)
}

// UpdateBalanceInAccount is a free data retrieval call binding the contract method 0x2e9be49a.
//
// Solidity: function UpdateBalanceInAccount(TypesUserAccount original_account, uint256 new_balance) pure returns(TypesUserAccount updated_account)
func (_Rolluputils *RolluputilsCallerSession) UpdateBalanceInAccount(original_account TypesUserAccount, new_balance *big.Int) (TypesUserAccount, error) {
	return _Rolluputils.Contract.UpdateBalanceInAccount(&_Rolluputils.CallOpts, original_account, new_balance)
}

// CalculateAddress is a free data retrieval call binding the contract method 0xe8a4c04e.
//
// Solidity: function calculateAddress(bytes pub) pure returns(address addr)
func (_Rolluputils *RolluputilsCaller) CalculateAddress(opts *bind.CallOpts, pub []byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Rolluputils.contract.Call(opts, out, "calculateAddress", pub)
	return *ret0, err
}

// CalculateAddress is a free data retrieval call binding the contract method 0xe8a4c04e.
//
// Solidity: function calculateAddress(bytes pub) pure returns(address addr)
func (_Rolluputils *RolluputilsSession) CalculateAddress(pub []byte) (common.Address, error) {
	return _Rolluputils.Contract.CalculateAddress(&_Rolluputils.CallOpts, pub)
}

// CalculateAddress is a free data retrieval call binding the contract method 0xe8a4c04e.
//
// Solidity: function calculateAddress(bytes pub) pure returns(address addr)
func (_Rolluputils *RolluputilsCallerSession) CalculateAddress(pub []byte) (common.Address, error) {
	return _Rolluputils.Contract.CalculateAddress(&_Rolluputils.CallOpts, pub)
}

// GetAccountHash is a free data retrieval call binding the contract method 0x61a8c3c2.
//
// Solidity: function getAccountHash(uint256 id, uint256 balance, uint256 nonce, uint256 tokenType) pure returns(bytes32)
func (_Rolluputils *RolluputilsCaller) GetAccountHash(opts *bind.CallOpts, id *big.Int, balance *big.Int, nonce *big.Int, tokenType *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Rolluputils.contract.Call(opts, out, "getAccountHash", id, balance, nonce, tokenType)
	return *ret0, err
}

// GetAccountHash is a free data retrieval call binding the contract method 0x61a8c3c2.
//
// Solidity: function getAccountHash(uint256 id, uint256 balance, uint256 nonce, uint256 tokenType) pure returns(bytes32)
func (_Rolluputils *RolluputilsSession) GetAccountHash(id *big.Int, balance *big.Int, nonce *big.Int, tokenType *big.Int) ([32]byte, error) {
	return _Rolluputils.Contract.GetAccountHash(&_Rolluputils.CallOpts, id, balance, nonce, tokenType)
}

// GetAccountHash is a free data retrieval call binding the contract method 0x61a8c3c2.
//
// Solidity: function getAccountHash(uint256 id, uint256 balance, uint256 nonce, uint256 tokenType) pure returns(bytes32)
func (_Rolluputils *RolluputilsCallerSession) GetAccountHash(id *big.Int, balance *big.Int, nonce *big.Int, tokenType *big.Int) ([32]byte, error) {
	return _Rolluputils.Contract.GetAccountHash(&_Rolluputils.CallOpts, id, balance, nonce, tokenType)
}

// GetTxSignBytes is a free data retrieval call binding the contract method 0x3ff55544.
//
// Solidity: function getTxSignBytes(uint256 fromIndex, uint256 toIndex, uint256 tokenType, uint256 txType, uint256 nonce, uint256 amount) pure returns(bytes32)
func (_Rolluputils *RolluputilsCaller) GetTxSignBytes(opts *bind.CallOpts, fromIndex *big.Int, toIndex *big.Int, tokenType *big.Int, txType *big.Int, nonce *big.Int, amount *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Rolluputils.contract.Call(opts, out, "getTxSignBytes", fromIndex, toIndex, tokenType, txType, nonce, amount)
	return *ret0, err
}

// GetTxSignBytes is a free data retrieval call binding the contract method 0x3ff55544.
//
// Solidity: function getTxSignBytes(uint256 fromIndex, uint256 toIndex, uint256 tokenType, uint256 txType, uint256 nonce, uint256 amount) pure returns(bytes32)
func (_Rolluputils *RolluputilsSession) GetTxSignBytes(fromIndex *big.Int, toIndex *big.Int, tokenType *big.Int, txType *big.Int, nonce *big.Int, amount *big.Int) ([32]byte, error) {
	return _Rolluputils.Contract.GetTxSignBytes(&_Rolluputils.CallOpts, fromIndex, toIndex, tokenType, txType, nonce, amount)
}

// GetTxSignBytes is a free data retrieval call binding the contract method 0x3ff55544.
//
// Solidity: function getTxSignBytes(uint256 fromIndex, uint256 toIndex, uint256 tokenType, uint256 txType, uint256 nonce, uint256 amount) pure returns(bytes32)
func (_Rolluputils *RolluputilsCallerSession) GetTxSignBytes(fromIndex *big.Int, toIndex *big.Int, tokenType *big.Int, txType *big.Int, nonce *big.Int, amount *big.Int) ([32]byte, error) {
	return _Rolluputils.Contract.GetTxSignBytes(&_Rolluputils.CallOpts, fromIndex, toIndex, tokenType, txType, nonce, amount)
}
