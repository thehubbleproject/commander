// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package merkleTree

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

// MerkleTreeABI is the input ABI used to generate the binding from.
const MerkleTreeABI = "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"defaultHashes\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tree\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"height\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"_dataBlocks\",\"type\":\"bytes[]\"}],\"name\":\"getMerkleRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_dataBlock\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_path\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"_siblings\",\"type\":\"bytes32[]\"}],\"name\":\"computeInclusionProofRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_root\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_dataBlock\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_path\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"_siblings\",\"type\":\"bytes32[]\"}],\"name\":\"verify\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_dataBlock\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_path\",\"type\":\"uint256\"}],\"name\":\"update\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_leaf\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_path\",\"type\":\"uint256\"}],\"name\":\"updateLeaf\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_dataBlock\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_path\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"_siblings\",\"type\":\"bytes32[]\"}],\"name\":\"verifyAndStore\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_dataBlock\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_path\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"_siblings\",\"type\":\"bytes32[]\"}],\"name\":\"store\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_leaf\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_path\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"_siblings\",\"type\":\"bytes32[]\"}],\"name\":\"storeLeaf\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_leaf\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_path\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_root\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"_siblings\",\"type\":\"bytes32[]\"}],\"name\":\"updateLeafWithSiblings\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_path\",\"type\":\"uint256\"}],\"name\":\"getSiblings\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_root\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_height\",\"type\":\"uint256\"}],\"name\":\"setMerkleRootAndHeight\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_parent\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_leftChild\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_rightChild\",\"type\":\"bytes32\"}],\"name\":\"storeNode\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_intVal\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getNthBitFromRight\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_parent\",\"type\":\"bytes32\"}],\"name\":\"getChildren\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_parent\",\"type\":\"bytes32\"}],\"name\":\"getLeftSiblingKey\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_parent\",\"type\":\"bytes32\"}],\"name\":\"getRightSiblingKey\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// MerkleTree is an auto generated Go binding around an Ethereum contract.
type MerkleTree struct {
	MerkleTreeCaller     // Read-only binding to the contract
	MerkleTreeTransactor // Write-only binding to the contract
	MerkleTreeFilterer   // Log filterer for contract events
}

// MerkleTreeCaller is an auto generated read-only Go binding around an Ethereum contract.
type MerkleTreeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MerkleTreeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MerkleTreeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MerkleTreeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MerkleTreeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MerkleTreeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MerkleTreeSession struct {
	Contract     *MerkleTree       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MerkleTreeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MerkleTreeCallerSession struct {
	Contract *MerkleTreeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// MerkleTreeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MerkleTreeTransactorSession struct {
	Contract     *MerkleTreeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// MerkleTreeRaw is an auto generated low-level Go binding around an Ethereum contract.
type MerkleTreeRaw struct {
	Contract *MerkleTree // Generic contract binding to access the raw methods on
}

// MerkleTreeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MerkleTreeCallerRaw struct {
	Contract *MerkleTreeCaller // Generic read-only contract binding to access the raw methods on
}

// MerkleTreeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MerkleTreeTransactorRaw struct {
	Contract *MerkleTreeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMerkleTree creates a new instance of MerkleTree, bound to a specific deployed contract.
func NewMerkleTree(address common.Address, backend bind.ContractBackend) (*MerkleTree, error) {
	contract, err := bindMerkleTree(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MerkleTree{MerkleTreeCaller: MerkleTreeCaller{contract: contract}, MerkleTreeTransactor: MerkleTreeTransactor{contract: contract}, MerkleTreeFilterer: MerkleTreeFilterer{contract: contract}}, nil
}

// NewMerkleTreeCaller creates a new read-only instance of MerkleTree, bound to a specific deployed contract.
func NewMerkleTreeCaller(address common.Address, caller bind.ContractCaller) (*MerkleTreeCaller, error) {
	contract, err := bindMerkleTree(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MerkleTreeCaller{contract: contract}, nil
}

// NewMerkleTreeTransactor creates a new write-only instance of MerkleTree, bound to a specific deployed contract.
func NewMerkleTreeTransactor(address common.Address, transactor bind.ContractTransactor) (*MerkleTreeTransactor, error) {
	contract, err := bindMerkleTree(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MerkleTreeTransactor{contract: contract}, nil
}

// NewMerkleTreeFilterer creates a new log filterer instance of MerkleTree, bound to a specific deployed contract.
func NewMerkleTreeFilterer(address common.Address, filterer bind.ContractFilterer) (*MerkleTreeFilterer, error) {
	contract, err := bindMerkleTree(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MerkleTreeFilterer{contract: contract}, nil
}

// bindMerkleTree binds a generic wrapper to an already deployed contract.
func bindMerkleTree(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MerkleTreeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MerkleTree *MerkleTreeRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MerkleTree.Contract.MerkleTreeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MerkleTree *MerkleTreeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MerkleTree.Contract.MerkleTreeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MerkleTree *MerkleTreeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MerkleTree.Contract.MerkleTreeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MerkleTree *MerkleTreeCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MerkleTree.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MerkleTree *MerkleTreeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MerkleTree.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MerkleTree *MerkleTreeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MerkleTree.Contract.contract.Transact(opts, method, params...)
}

// ComputeInclusionProofRoot is a free data retrieval call binding the contract method 0x99ae9309.
//
// Solidity: function computeInclusionProofRoot(bytes _dataBlock, uint256 _path, bytes32[] _siblings) constant returns(bytes32)
func (_MerkleTree *MerkleTreeCaller) ComputeInclusionProofRoot(opts *bind.CallOpts, _dataBlock []byte, _path *big.Int, _siblings [][32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _MerkleTree.contract.Call(opts, out, "computeInclusionProofRoot", _dataBlock, _path, _siblings)
	return *ret0, err
}

// ComputeInclusionProofRoot is a free data retrieval call binding the contract method 0x99ae9309.
//
// Solidity: function computeInclusionProofRoot(bytes _dataBlock, uint256 _path, bytes32[] _siblings) constant returns(bytes32)
func (_MerkleTree *MerkleTreeSession) ComputeInclusionProofRoot(_dataBlock []byte, _path *big.Int, _siblings [][32]byte) ([32]byte, error) {
	return _MerkleTree.Contract.ComputeInclusionProofRoot(&_MerkleTree.CallOpts, _dataBlock, _path, _siblings)
}

// ComputeInclusionProofRoot is a free data retrieval call binding the contract method 0x99ae9309.
//
// Solidity: function computeInclusionProofRoot(bytes _dataBlock, uint256 _path, bytes32[] _siblings) constant returns(bytes32)
func (_MerkleTree *MerkleTreeCallerSession) ComputeInclusionProofRoot(_dataBlock []byte, _path *big.Int, _siblings [][32]byte) ([32]byte, error) {
	return _MerkleTree.Contract.ComputeInclusionProofRoot(&_MerkleTree.CallOpts, _dataBlock, _path, _siblings)
}

// DefaultHashes is a free data retrieval call binding the contract method 0x48419ad8.
//
// Solidity: function defaultHashes(uint256 ) constant returns(bytes32)
func (_MerkleTree *MerkleTreeCaller) DefaultHashes(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _MerkleTree.contract.Call(opts, out, "defaultHashes", arg0)
	return *ret0, err
}

// DefaultHashes is a free data retrieval call binding the contract method 0x48419ad8.
//
// Solidity: function defaultHashes(uint256 ) constant returns(bytes32)
func (_MerkleTree *MerkleTreeSession) DefaultHashes(arg0 *big.Int) ([32]byte, error) {
	return _MerkleTree.Contract.DefaultHashes(&_MerkleTree.CallOpts, arg0)
}

// DefaultHashes is a free data retrieval call binding the contract method 0x48419ad8.
//
// Solidity: function defaultHashes(uint256 ) constant returns(bytes32)
func (_MerkleTree *MerkleTreeCallerSession) DefaultHashes(arg0 *big.Int) ([32]byte, error) {
	return _MerkleTree.Contract.DefaultHashes(&_MerkleTree.CallOpts, arg0)
}

// GetChildren is a free data retrieval call binding the contract method 0xd37684ff.
//
// Solidity: function getChildren(bytes32 _parent) constant returns(bytes32, bytes32)
func (_MerkleTree *MerkleTreeCaller) GetChildren(opts *bind.CallOpts, _parent [32]byte) ([32]byte, [32]byte, error) {
	var (
		ret0 = new([32]byte)
		ret1 = new([32]byte)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _MerkleTree.contract.Call(opts, out, "getChildren", _parent)
	return *ret0, *ret1, err
}

// GetChildren is a free data retrieval call binding the contract method 0xd37684ff.
//
// Solidity: function getChildren(bytes32 _parent) constant returns(bytes32, bytes32)
func (_MerkleTree *MerkleTreeSession) GetChildren(_parent [32]byte) ([32]byte, [32]byte, error) {
	return _MerkleTree.Contract.GetChildren(&_MerkleTree.CallOpts, _parent)
}

// GetChildren is a free data retrieval call binding the contract method 0xd37684ff.
//
// Solidity: function getChildren(bytes32 _parent) constant returns(bytes32, bytes32)
func (_MerkleTree *MerkleTreeCallerSession) GetChildren(_parent [32]byte) ([32]byte, [32]byte, error) {
	return _MerkleTree.Contract.GetChildren(&_MerkleTree.CallOpts, _parent)
}

// GetLeftSiblingKey is a free data retrieval call binding the contract method 0xdf7c7263.
//
// Solidity: function getLeftSiblingKey(bytes32 _parent) constant returns(bytes32)
func (_MerkleTree *MerkleTreeCaller) GetLeftSiblingKey(opts *bind.CallOpts, _parent [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _MerkleTree.contract.Call(opts, out, "getLeftSiblingKey", _parent)
	return *ret0, err
}

// GetLeftSiblingKey is a free data retrieval call binding the contract method 0xdf7c7263.
//
// Solidity: function getLeftSiblingKey(bytes32 _parent) constant returns(bytes32)
func (_MerkleTree *MerkleTreeSession) GetLeftSiblingKey(_parent [32]byte) ([32]byte, error) {
	return _MerkleTree.Contract.GetLeftSiblingKey(&_MerkleTree.CallOpts, _parent)
}

// GetLeftSiblingKey is a free data retrieval call binding the contract method 0xdf7c7263.
//
// Solidity: function getLeftSiblingKey(bytes32 _parent) constant returns(bytes32)
func (_MerkleTree *MerkleTreeCallerSession) GetLeftSiblingKey(_parent [32]byte) ([32]byte, error) {
	return _MerkleTree.Contract.GetLeftSiblingKey(&_MerkleTree.CallOpts, _parent)
}

// GetMerkleRoot is a free data retrieval call binding the contract method 0x40ff34ef.
//
// Solidity: function getMerkleRoot(bytes[] _dataBlocks) constant returns(bytes32)
func (_MerkleTree *MerkleTreeCaller) GetMerkleRoot(opts *bind.CallOpts, _dataBlocks [][]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _MerkleTree.contract.Call(opts, out, "getMerkleRoot", _dataBlocks)
	return *ret0, err
}

// GetMerkleRoot is a free data retrieval call binding the contract method 0x40ff34ef.
//
// Solidity: function getMerkleRoot(bytes[] _dataBlocks) constant returns(bytes32)
func (_MerkleTree *MerkleTreeSession) GetMerkleRoot(_dataBlocks [][]byte) ([32]byte, error) {
	return _MerkleTree.Contract.GetMerkleRoot(&_MerkleTree.CallOpts, _dataBlocks)
}

// GetMerkleRoot is a free data retrieval call binding the contract method 0x40ff34ef.
//
// Solidity: function getMerkleRoot(bytes[] _dataBlocks) constant returns(bytes32)
func (_MerkleTree *MerkleTreeCallerSession) GetMerkleRoot(_dataBlocks [][]byte) ([32]byte, error) {
	return _MerkleTree.Contract.GetMerkleRoot(&_MerkleTree.CallOpts, _dataBlocks)
}

// GetNthBitFromRight is a free data retrieval call binding the contract method 0xdb0787cb.
//
// Solidity: function getNthBitFromRight(uint256 _intVal, uint256 _index) constant returns(uint8)
func (_MerkleTree *MerkleTreeCaller) GetNthBitFromRight(opts *bind.CallOpts, _intVal *big.Int, _index *big.Int) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _MerkleTree.contract.Call(opts, out, "getNthBitFromRight", _intVal, _index)
	return *ret0, err
}

// GetNthBitFromRight is a free data retrieval call binding the contract method 0xdb0787cb.
//
// Solidity: function getNthBitFromRight(uint256 _intVal, uint256 _index) constant returns(uint8)
func (_MerkleTree *MerkleTreeSession) GetNthBitFromRight(_intVal *big.Int, _index *big.Int) (uint8, error) {
	return _MerkleTree.Contract.GetNthBitFromRight(&_MerkleTree.CallOpts, _intVal, _index)
}

// GetNthBitFromRight is a free data retrieval call binding the contract method 0xdb0787cb.
//
// Solidity: function getNthBitFromRight(uint256 _intVal, uint256 _index) constant returns(uint8)
func (_MerkleTree *MerkleTreeCallerSession) GetNthBitFromRight(_intVal *big.Int, _index *big.Int) (uint8, error) {
	return _MerkleTree.Contract.GetNthBitFromRight(&_MerkleTree.CallOpts, _intVal, _index)
}

// GetRightSiblingKey is a free data retrieval call binding the contract method 0xe913e47f.
//
// Solidity: function getRightSiblingKey(bytes32 _parent) constant returns(bytes32)
func (_MerkleTree *MerkleTreeCaller) GetRightSiblingKey(opts *bind.CallOpts, _parent [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _MerkleTree.contract.Call(opts, out, "getRightSiblingKey", _parent)
	return *ret0, err
}

// GetRightSiblingKey is a free data retrieval call binding the contract method 0xe913e47f.
//
// Solidity: function getRightSiblingKey(bytes32 _parent) constant returns(bytes32)
func (_MerkleTree *MerkleTreeSession) GetRightSiblingKey(_parent [32]byte) ([32]byte, error) {
	return _MerkleTree.Contract.GetRightSiblingKey(&_MerkleTree.CallOpts, _parent)
}

// GetRightSiblingKey is a free data retrieval call binding the contract method 0xe913e47f.
//
// Solidity: function getRightSiblingKey(bytes32 _parent) constant returns(bytes32)
func (_MerkleTree *MerkleTreeCallerSession) GetRightSiblingKey(_parent [32]byte) ([32]byte, error) {
	return _MerkleTree.Contract.GetRightSiblingKey(&_MerkleTree.CallOpts, _parent)
}

// GetRoot is a free data retrieval call binding the contract method 0x5ca1e165.
//
// Solidity: function getRoot() constant returns(bytes32)
func (_MerkleTree *MerkleTreeCaller) GetRoot(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _MerkleTree.contract.Call(opts, out, "getRoot")
	return *ret0, err
}

// GetRoot is a free data retrieval call binding the contract method 0x5ca1e165.
//
// Solidity: function getRoot() constant returns(bytes32)
func (_MerkleTree *MerkleTreeSession) GetRoot() ([32]byte, error) {
	return _MerkleTree.Contract.GetRoot(&_MerkleTree.CallOpts)
}

// GetRoot is a free data retrieval call binding the contract method 0x5ca1e165.
//
// Solidity: function getRoot() constant returns(bytes32)
func (_MerkleTree *MerkleTreeCallerSession) GetRoot() ([32]byte, error) {
	return _MerkleTree.Contract.GetRoot(&_MerkleTree.CallOpts)
}

// GetSiblings is a free data retrieval call binding the contract method 0x101b166c.
//
// Solidity: function getSiblings(uint256 _path) constant returns(bytes32[])
func (_MerkleTree *MerkleTreeCaller) GetSiblings(opts *bind.CallOpts, _path *big.Int) ([][32]byte, error) {
	var (
		ret0 = new([][32]byte)
	)
	out := ret0
	err := _MerkleTree.contract.Call(opts, out, "getSiblings", _path)
	return *ret0, err
}

// GetSiblings is a free data retrieval call binding the contract method 0x101b166c.
//
// Solidity: function getSiblings(uint256 _path) constant returns(bytes32[])
func (_MerkleTree *MerkleTreeSession) GetSiblings(_path *big.Int) ([][32]byte, error) {
	return _MerkleTree.Contract.GetSiblings(&_MerkleTree.CallOpts, _path)
}

// GetSiblings is a free data retrieval call binding the contract method 0x101b166c.
//
// Solidity: function getSiblings(uint256 _path) constant returns(bytes32[])
func (_MerkleTree *MerkleTreeCallerSession) GetSiblings(_path *big.Int) ([][32]byte, error) {
	return _MerkleTree.Contract.GetSiblings(&_MerkleTree.CallOpts, _path)
}

// Tree is a free data retrieval call binding the contract method 0xfd54b228.
//
// Solidity: function tree() constant returns(bytes32 root, uint256 height)
func (_MerkleTree *MerkleTreeCaller) Tree(opts *bind.CallOpts) (struct {
	Root   [32]byte
	Height *big.Int
}, error) {
	ret := new(struct {
		Root   [32]byte
		Height *big.Int
	})
	out := ret
	err := _MerkleTree.contract.Call(opts, out, "tree")
	return *ret, err
}

// Tree is a free data retrieval call binding the contract method 0xfd54b228.
//
// Solidity: function tree() constant returns(bytes32 root, uint256 height)
func (_MerkleTree *MerkleTreeSession) Tree() (struct {
	Root   [32]byte
	Height *big.Int
}, error) {
	return _MerkleTree.Contract.Tree(&_MerkleTree.CallOpts)
}

// Tree is a free data retrieval call binding the contract method 0xfd54b228.
//
// Solidity: function tree() constant returns(bytes32 root, uint256 height)
func (_MerkleTree *MerkleTreeCallerSession) Tree() (struct {
	Root   [32]byte
	Height *big.Int
}, error) {
	return _MerkleTree.Contract.Tree(&_MerkleTree.CallOpts)
}

// Verify is a free data retrieval call binding the contract method 0x30d90a76.
//
// Solidity: function verify(bytes32 _root, bytes _dataBlock, uint256 _path, bytes32[] _siblings) constant returns(bool)
func (_MerkleTree *MerkleTreeCaller) Verify(opts *bind.CallOpts, _root [32]byte, _dataBlock []byte, _path *big.Int, _siblings [][32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _MerkleTree.contract.Call(opts, out, "verify", _root, _dataBlock, _path, _siblings)
	return *ret0, err
}

// Verify is a free data retrieval call binding the contract method 0x30d90a76.
//
// Solidity: function verify(bytes32 _root, bytes _dataBlock, uint256 _path, bytes32[] _siblings) constant returns(bool)
func (_MerkleTree *MerkleTreeSession) Verify(_root [32]byte, _dataBlock []byte, _path *big.Int, _siblings [][32]byte) (bool, error) {
	return _MerkleTree.Contract.Verify(&_MerkleTree.CallOpts, _root, _dataBlock, _path, _siblings)
}

// Verify is a free data retrieval call binding the contract method 0x30d90a76.
//
// Solidity: function verify(bytes32 _root, bytes _dataBlock, uint256 _path, bytes32[] _siblings) constant returns(bool)
func (_MerkleTree *MerkleTreeCallerSession) Verify(_root [32]byte, _dataBlock []byte, _path *big.Int, _siblings [][32]byte) (bool, error) {
	return _MerkleTree.Contract.Verify(&_MerkleTree.CallOpts, _root, _dataBlock, _path, _siblings)
}

// SetMerkleRootAndHeight is a paid mutator transaction binding the contract method 0x5c22b6d9.
//
// Solidity: function setMerkleRootAndHeight(bytes32 _root, uint256 _height) returns()
func (_MerkleTree *MerkleTreeTransactor) SetMerkleRootAndHeight(opts *bind.TransactOpts, _root [32]byte, _height *big.Int) (*types.Transaction, error) {
	return _MerkleTree.contract.Transact(opts, "setMerkleRootAndHeight", _root, _height)
}

// SetMerkleRootAndHeight is a paid mutator transaction binding the contract method 0x5c22b6d9.
//
// Solidity: function setMerkleRootAndHeight(bytes32 _root, uint256 _height) returns()
func (_MerkleTree *MerkleTreeSession) SetMerkleRootAndHeight(_root [32]byte, _height *big.Int) (*types.Transaction, error) {
	return _MerkleTree.Contract.SetMerkleRootAndHeight(&_MerkleTree.TransactOpts, _root, _height)
}

// SetMerkleRootAndHeight is a paid mutator transaction binding the contract method 0x5c22b6d9.
//
// Solidity: function setMerkleRootAndHeight(bytes32 _root, uint256 _height) returns()
func (_MerkleTree *MerkleTreeTransactorSession) SetMerkleRootAndHeight(_root [32]byte, _height *big.Int) (*types.Transaction, error) {
	return _MerkleTree.Contract.SetMerkleRootAndHeight(&_MerkleTree.TransactOpts, _root, _height)
}

// Store is a paid mutator transaction binding the contract method 0x158933ad.
//
// Solidity: function store(bytes _dataBlock, uint256 _path, bytes32[] _siblings) returns()
func (_MerkleTree *MerkleTreeTransactor) Store(opts *bind.TransactOpts, _dataBlock []byte, _path *big.Int, _siblings [][32]byte) (*types.Transaction, error) {
	return _MerkleTree.contract.Transact(opts, "store", _dataBlock, _path, _siblings)
}

// Store is a paid mutator transaction binding the contract method 0x158933ad.
//
// Solidity: function store(bytes _dataBlock, uint256 _path, bytes32[] _siblings) returns()
func (_MerkleTree *MerkleTreeSession) Store(_dataBlock []byte, _path *big.Int, _siblings [][32]byte) (*types.Transaction, error) {
	return _MerkleTree.Contract.Store(&_MerkleTree.TransactOpts, _dataBlock, _path, _siblings)
}

// Store is a paid mutator transaction binding the contract method 0x158933ad.
//
// Solidity: function store(bytes _dataBlock, uint256 _path, bytes32[] _siblings) returns()
func (_MerkleTree *MerkleTreeTransactorSession) Store(_dataBlock []byte, _path *big.Int, _siblings [][32]byte) (*types.Transaction, error) {
	return _MerkleTree.Contract.Store(&_MerkleTree.TransactOpts, _dataBlock, _path, _siblings)
}

// StoreLeaf is a paid mutator transaction binding the contract method 0x9c0de520.
//
// Solidity: function storeLeaf(bytes32 _leaf, uint256 _path, bytes32[] _siblings) returns()
func (_MerkleTree *MerkleTreeTransactor) StoreLeaf(opts *bind.TransactOpts, _leaf [32]byte, _path *big.Int, _siblings [][32]byte) (*types.Transaction, error) {
	return _MerkleTree.contract.Transact(opts, "storeLeaf", _leaf, _path, _siblings)
}

// StoreLeaf is a paid mutator transaction binding the contract method 0x9c0de520.
//
// Solidity: function storeLeaf(bytes32 _leaf, uint256 _path, bytes32[] _siblings) returns()
func (_MerkleTree *MerkleTreeSession) StoreLeaf(_leaf [32]byte, _path *big.Int, _siblings [][32]byte) (*types.Transaction, error) {
	return _MerkleTree.Contract.StoreLeaf(&_MerkleTree.TransactOpts, _leaf, _path, _siblings)
}

// StoreLeaf is a paid mutator transaction binding the contract method 0x9c0de520.
//
// Solidity: function storeLeaf(bytes32 _leaf, uint256 _path, bytes32[] _siblings) returns()
func (_MerkleTree *MerkleTreeTransactorSession) StoreLeaf(_leaf [32]byte, _path *big.Int, _siblings [][32]byte) (*types.Transaction, error) {
	return _MerkleTree.Contract.StoreLeaf(&_MerkleTree.TransactOpts, _leaf, _path, _siblings)
}

// StoreNode is a paid mutator transaction binding the contract method 0x63327f89.
//
// Solidity: function storeNode(bytes32 _parent, bytes32 _leftChild, bytes32 _rightChild) returns()
func (_MerkleTree *MerkleTreeTransactor) StoreNode(opts *bind.TransactOpts, _parent [32]byte, _leftChild [32]byte, _rightChild [32]byte) (*types.Transaction, error) {
	return _MerkleTree.contract.Transact(opts, "storeNode", _parent, _leftChild, _rightChild)
}

// StoreNode is a paid mutator transaction binding the contract method 0x63327f89.
//
// Solidity: function storeNode(bytes32 _parent, bytes32 _leftChild, bytes32 _rightChild) returns()
func (_MerkleTree *MerkleTreeSession) StoreNode(_parent [32]byte, _leftChild [32]byte, _rightChild [32]byte) (*types.Transaction, error) {
	return _MerkleTree.Contract.StoreNode(&_MerkleTree.TransactOpts, _parent, _leftChild, _rightChild)
}

// StoreNode is a paid mutator transaction binding the contract method 0x63327f89.
//
// Solidity: function storeNode(bytes32 _parent, bytes32 _leftChild, bytes32 _rightChild) returns()
func (_MerkleTree *MerkleTreeTransactorSession) StoreNode(_parent [32]byte, _leftChild [32]byte, _rightChild [32]byte) (*types.Transaction, error) {
	return _MerkleTree.Contract.StoreNode(&_MerkleTree.TransactOpts, _parent, _leftChild, _rightChild)
}

// Update is a paid mutator transaction binding the contract method 0xc3b45234.
//
// Solidity: function update(bytes _dataBlock, uint256 _path) returns()
func (_MerkleTree *MerkleTreeTransactor) Update(opts *bind.TransactOpts, _dataBlock []byte, _path *big.Int) (*types.Transaction, error) {
	return _MerkleTree.contract.Transact(opts, "update", _dataBlock, _path)
}

// Update is a paid mutator transaction binding the contract method 0xc3b45234.
//
// Solidity: function update(bytes _dataBlock, uint256 _path) returns()
func (_MerkleTree *MerkleTreeSession) Update(_dataBlock []byte, _path *big.Int) (*types.Transaction, error) {
	return _MerkleTree.Contract.Update(&_MerkleTree.TransactOpts, _dataBlock, _path)
}

// Update is a paid mutator transaction binding the contract method 0xc3b45234.
//
// Solidity: function update(bytes _dataBlock, uint256 _path) returns()
func (_MerkleTree *MerkleTreeTransactorSession) Update(_dataBlock []byte, _path *big.Int) (*types.Transaction, error) {
	return _MerkleTree.Contract.Update(&_MerkleTree.TransactOpts, _dataBlock, _path)
}

// UpdateLeaf is a paid mutator transaction binding the contract method 0x272684b5.
//
// Solidity: function updateLeaf(bytes32 _leaf, uint256 _path) returns()
func (_MerkleTree *MerkleTreeTransactor) UpdateLeaf(opts *bind.TransactOpts, _leaf [32]byte, _path *big.Int) (*types.Transaction, error) {
	return _MerkleTree.contract.Transact(opts, "updateLeaf", _leaf, _path)
}

// UpdateLeaf is a paid mutator transaction binding the contract method 0x272684b5.
//
// Solidity: function updateLeaf(bytes32 _leaf, uint256 _path) returns()
func (_MerkleTree *MerkleTreeSession) UpdateLeaf(_leaf [32]byte, _path *big.Int) (*types.Transaction, error) {
	return _MerkleTree.Contract.UpdateLeaf(&_MerkleTree.TransactOpts, _leaf, _path)
}

// UpdateLeaf is a paid mutator transaction binding the contract method 0x272684b5.
//
// Solidity: function updateLeaf(bytes32 _leaf, uint256 _path) returns()
func (_MerkleTree *MerkleTreeTransactorSession) UpdateLeaf(_leaf [32]byte, _path *big.Int) (*types.Transaction, error) {
	return _MerkleTree.Contract.UpdateLeaf(&_MerkleTree.TransactOpts, _leaf, _path)
}

// UpdateLeafWithSiblings is a paid mutator transaction binding the contract method 0x240d1a5a.
//
// Solidity: function updateLeafWithSiblings(bytes32 _leaf, uint256 _path, bytes32 _root, bytes32[] _siblings) returns(bytes32)
func (_MerkleTree *MerkleTreeTransactor) UpdateLeafWithSiblings(opts *bind.TransactOpts, _leaf [32]byte, _path *big.Int, _root [32]byte, _siblings [][32]byte) (*types.Transaction, error) {
	return _MerkleTree.contract.Transact(opts, "updateLeafWithSiblings", _leaf, _path, _root, _siblings)
}

// UpdateLeafWithSiblings is a paid mutator transaction binding the contract method 0x240d1a5a.
//
// Solidity: function updateLeafWithSiblings(bytes32 _leaf, uint256 _path, bytes32 _root, bytes32[] _siblings) returns(bytes32)
func (_MerkleTree *MerkleTreeSession) UpdateLeafWithSiblings(_leaf [32]byte, _path *big.Int, _root [32]byte, _siblings [][32]byte) (*types.Transaction, error) {
	return _MerkleTree.Contract.UpdateLeafWithSiblings(&_MerkleTree.TransactOpts, _leaf, _path, _root, _siblings)
}

// UpdateLeafWithSiblings is a paid mutator transaction binding the contract method 0x240d1a5a.
//
// Solidity: function updateLeafWithSiblings(bytes32 _leaf, uint256 _path, bytes32 _root, bytes32[] _siblings) returns(bytes32)
func (_MerkleTree *MerkleTreeTransactorSession) UpdateLeafWithSiblings(_leaf [32]byte, _path *big.Int, _root [32]byte, _siblings [][32]byte) (*types.Transaction, error) {
	return _MerkleTree.Contract.UpdateLeafWithSiblings(&_MerkleTree.TransactOpts, _leaf, _path, _root, _siblings)
}

// VerifyAndStore is a paid mutator transaction binding the contract method 0x4359356d.
//
// Solidity: function verifyAndStore(bytes _dataBlock, uint256 _path, bytes32[] _siblings) returns()
func (_MerkleTree *MerkleTreeTransactor) VerifyAndStore(opts *bind.TransactOpts, _dataBlock []byte, _path *big.Int, _siblings [][32]byte) (*types.Transaction, error) {
	return _MerkleTree.contract.Transact(opts, "verifyAndStore", _dataBlock, _path, _siblings)
}

// VerifyAndStore is a paid mutator transaction binding the contract method 0x4359356d.
//
// Solidity: function verifyAndStore(bytes _dataBlock, uint256 _path, bytes32[] _siblings) returns()
func (_MerkleTree *MerkleTreeSession) VerifyAndStore(_dataBlock []byte, _path *big.Int, _siblings [][32]byte) (*types.Transaction, error) {
	return _MerkleTree.Contract.VerifyAndStore(&_MerkleTree.TransactOpts, _dataBlock, _path, _siblings)
}

// VerifyAndStore is a paid mutator transaction binding the contract method 0x4359356d.
//
// Solidity: function verifyAndStore(bytes _dataBlock, uint256 _path, bytes32[] _siblings) returns()
func (_MerkleTree *MerkleTreeTransactorSession) VerifyAndStore(_dataBlock []byte, _path *big.Int, _siblings [][32]byte) (*types.Transaction, error) {
	return _MerkleTree.Contract.VerifyAndStore(&_MerkleTree.TransactOpts, _dataBlock, _path, _siblings)
}
