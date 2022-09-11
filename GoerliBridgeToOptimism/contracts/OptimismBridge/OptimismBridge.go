// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package optimismBridge

import (
	"errors"
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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// OptimismBridgeMetaData contains all meta data concerning the OptimismBridge contract.
var OptimismBridgeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"bridgeEmpty\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"bridgeOnOtherSideNeedsLiqudity\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"msgValueDoesNotCoverFee\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"msgValueLessThan1000\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"msgValueZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"notExternalBridge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"notOwnerAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"queueIsEmpty\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"queueNotEmpty\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dequeue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"first\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"last\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"bridgeAmount\",\"type\":\"uint256\"}],\"name\":\"lockTokensForGoerli\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"lockedForGoerliETH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"optimismBridgedETH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ownerAddBridgeLiqudity\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ownerRemoveBridgeLiqudity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"userToBridge\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"lockedForOptimismETH\",\"type\":\"uint256\"}],\"name\":\"ownerUnlockOptimismETH\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"queue\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60a0604052600160045534801561001557600080fd5b503373ffffffffffffffffffffffffffffffffffffffff1660808173ffffffffffffffffffffffffffffffffffffffff1681525050608051610c3561008f60003960008181610258015281816103170152818161038c0152818161044d01528181610564015281816106a601526107100152610c356000f3fe60806040526004361061009c5760003560e01c8063b4a99a4e11610064578063b4a99a4e1461012f578063c5aaccf71461015a578063c637aa9714610197578063cfd121e6146101b3578063ddf0b009146101dc578063e3d66f68146102195761009c565b806308dd057a146100a15780633df4ddf4146100b857806347799da8146100e35780638521eebe1461010e578063957908d114610118575b600080fd5b3480156100ad57600080fd5b506100b6610256565b005b3480156100c457600080fd5b506100cd61037e565b6040516100da9190610958565b60405180910390f35b3480156100ef57600080fd5b506100f8610384565b6040516101059190610958565b60405180910390f35b61011661038a565b005b34801561012457600080fd5b5061012d61044b565b005b34801561013b57600080fd5b50610144610562565b60405161015191906109b4565b60405180910390f35b34801561016657600080fd5b50610181600480360381019061017c9190610a00565b610586565b60405161018e9190610958565b60405180910390f35b6101b160048036038101906101ac9190610a59565b61059e565b005b3480156101bf57600080fd5b506101da60048036038101906101d59190610a86565b61070e565b005b3480156101e857600080fd5b5061020360048036038101906101fe9190610a59565b610884565b60405161021091906109b4565b60405180910390f35b34801561022557600080fd5b50610240600480360381019061023b9190610a00565b6108b7565b60405161024d9190610958565b60405180910390f35b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146102db576040517f1ee8575e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60004703610315576040517f7a1f291700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166108fc479081150290604051600060405180830381858888f1935050505015801561037b573d6000803e3d6000fd5b50565b60045481565b60035481565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461040f576040517f1ee8575e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60003403610449576040517f0cdb5d8500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146104d0576040517f1ee8575e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600454600354101561050e576040517f5e61bb8700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60026000600454815260200190815260200160002060006101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690556001600460008282546105599190610af5565b92505081905550565b7f000000000000000000000000000000000000000000000000000000000000000081565b60016020528060005260406000206000915090505481565b6103e88110156105da576040517fa9f2d61800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6103e8816103eb6105eb9190610b29565b6105f59190610b9a565b341461062d576040517f400f9fb800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6103eb346103e861063e9190610b29565b6106489190610b9a565b6000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546106959190610af5565b925050819055506106a46108cf565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166108fc349081150290604051600060405180830381858888f1935050505015801561070a573d6000803e3d6000fd5b5050565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610793576040517f1ee8575e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054826107e09190610bcb565b905080600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546108319190610af5565b925050819055508273ffffffffffffffffffffffffffffffffffffffff166108fc829081150290604051600060405180830381858888f1935050505015801561087e573d6000803e3d6000fd5b50505050565b60026020528060005260406000206000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60006020528060005260406000206000915090505481565b6001600360008282546108e29190610af5565b925050819055503360026000600354815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550565b6000819050919050565b6109528161093f565b82525050565b600060208201905061096d6000830184610949565b92915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061099e82610973565b9050919050565b6109ae81610993565b82525050565b60006020820190506109c960008301846109a5565b92915050565b600080fd5b6109dd81610993565b81146109e857600080fd5b50565b6000813590506109fa816109d4565b92915050565b600060208284031215610a1657610a156109cf565b5b6000610a24848285016109eb565b91505092915050565b610a368161093f565b8114610a4157600080fd5b50565b600081359050610a5381610a2d565b92915050565b600060208284031215610a6f57610a6e6109cf565b5b6000610a7d84828501610a44565b91505092915050565b60008060408385031215610a9d57610a9c6109cf565b5b6000610aab858286016109eb565b9250506020610abc85828601610a44565b9150509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000610b008261093f565b9150610b0b8361093f565b9250828201905080821115610b2357610b22610ac6565b5b92915050565b6000610b348261093f565b9150610b3f8361093f565b9250828202610b4d8161093f565b91508282048414831517610b6457610b63610ac6565b5b5092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b6000610ba58261093f565b9150610bb08361093f565b925082610bc057610bbf610b6b565b5b828204905092915050565b6000610bd68261093f565b9150610be18361093f565b9250828203905081811115610bf957610bf8610ac6565b5b9291505056fea264697066735822122031bb092eacd12b8ec7655d3c2fa02f011add28bb4f697f15aea395cedf2cb20364736f6c63430008110033",
}

// OptimismBridgeABI is the input ABI used to generate the binding from.
// Deprecated: Use OptimismBridgeMetaData.ABI instead.
var OptimismBridgeABI = OptimismBridgeMetaData.ABI

// OptimismBridgeBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use OptimismBridgeMetaData.Bin instead.
var OptimismBridgeBin = OptimismBridgeMetaData.Bin

// DeployOptimismBridge deploys a new Ethereum contract, binding an instance of OptimismBridge to it.
func DeployOptimismBridge(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *OptimismBridge, error) {
	parsed, err := OptimismBridgeMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OptimismBridgeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OptimismBridge{OptimismBridgeCaller: OptimismBridgeCaller{contract: contract}, OptimismBridgeTransactor: OptimismBridgeTransactor{contract: contract}, OptimismBridgeFilterer: OptimismBridgeFilterer{contract: contract}}, nil
}

// OptimismBridge is an auto generated Go binding around an Ethereum contract.
type OptimismBridge struct {
	OptimismBridgeCaller     // Read-only binding to the contract
	OptimismBridgeTransactor // Write-only binding to the contract
	OptimismBridgeFilterer   // Log filterer for contract events
}

// OptimismBridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type OptimismBridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OptimismBridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OptimismBridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OptimismBridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OptimismBridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OptimismBridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OptimismBridgeSession struct {
	Contract     *OptimismBridge   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OptimismBridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OptimismBridgeCallerSession struct {
	Contract *OptimismBridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// OptimismBridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OptimismBridgeTransactorSession struct {
	Contract     *OptimismBridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// OptimismBridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type OptimismBridgeRaw struct {
	Contract *OptimismBridge // Generic contract binding to access the raw methods on
}

// OptimismBridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OptimismBridgeCallerRaw struct {
	Contract *OptimismBridgeCaller // Generic read-only contract binding to access the raw methods on
}

// OptimismBridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OptimismBridgeTransactorRaw struct {
	Contract *OptimismBridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOptimismBridge creates a new instance of OptimismBridge, bound to a specific deployed contract.
func NewOptimismBridge(address common.Address, backend bind.ContractBackend) (*OptimismBridge, error) {
	contract, err := bindOptimismBridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OptimismBridge{OptimismBridgeCaller: OptimismBridgeCaller{contract: contract}, OptimismBridgeTransactor: OptimismBridgeTransactor{contract: contract}, OptimismBridgeFilterer: OptimismBridgeFilterer{contract: contract}}, nil
}

// NewOptimismBridgeCaller creates a new read-only instance of OptimismBridge, bound to a specific deployed contract.
func NewOptimismBridgeCaller(address common.Address, caller bind.ContractCaller) (*OptimismBridgeCaller, error) {
	contract, err := bindOptimismBridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OptimismBridgeCaller{contract: contract}, nil
}

// NewOptimismBridgeTransactor creates a new write-only instance of OptimismBridge, bound to a specific deployed contract.
func NewOptimismBridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*OptimismBridgeTransactor, error) {
	contract, err := bindOptimismBridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OptimismBridgeTransactor{contract: contract}, nil
}

// NewOptimismBridgeFilterer creates a new log filterer instance of OptimismBridge, bound to a specific deployed contract.
func NewOptimismBridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*OptimismBridgeFilterer, error) {
	contract, err := bindOptimismBridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OptimismBridgeFilterer{contract: contract}, nil
}

// bindOptimismBridge binds a generic wrapper to an already deployed contract.
func bindOptimismBridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OptimismBridgeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OptimismBridge *OptimismBridgeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OptimismBridge.Contract.OptimismBridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OptimismBridge *OptimismBridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OptimismBridge.Contract.OptimismBridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OptimismBridge *OptimismBridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OptimismBridge.Contract.OptimismBridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OptimismBridge *OptimismBridgeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OptimismBridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OptimismBridge *OptimismBridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OptimismBridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OptimismBridge *OptimismBridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OptimismBridge.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0xb4a99a4e.
//
// Solidity: function Owner() view returns(address)
func (_OptimismBridge *OptimismBridgeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OptimismBridge.contract.Call(opts, &out, "Owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0xb4a99a4e.
//
// Solidity: function Owner() view returns(address)
func (_OptimismBridge *OptimismBridgeSession) Owner() (common.Address, error) {
	return _OptimismBridge.Contract.Owner(&_OptimismBridge.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0xb4a99a4e.
//
// Solidity: function Owner() view returns(address)
func (_OptimismBridge *OptimismBridgeCallerSession) Owner() (common.Address, error) {
	return _OptimismBridge.Contract.Owner(&_OptimismBridge.CallOpts)
}

// First is a free data retrieval call binding the contract method 0x3df4ddf4.
//
// Solidity: function first() view returns(uint256)
func (_OptimismBridge *OptimismBridgeCaller) First(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OptimismBridge.contract.Call(opts, &out, "first")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// First is a free data retrieval call binding the contract method 0x3df4ddf4.
//
// Solidity: function first() view returns(uint256)
func (_OptimismBridge *OptimismBridgeSession) First() (*big.Int, error) {
	return _OptimismBridge.Contract.First(&_OptimismBridge.CallOpts)
}

// First is a free data retrieval call binding the contract method 0x3df4ddf4.
//
// Solidity: function first() view returns(uint256)
func (_OptimismBridge *OptimismBridgeCallerSession) First() (*big.Int, error) {
	return _OptimismBridge.Contract.First(&_OptimismBridge.CallOpts)
}

// Last is a free data retrieval call binding the contract method 0x47799da8.
//
// Solidity: function last() view returns(uint256)
func (_OptimismBridge *OptimismBridgeCaller) Last(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OptimismBridge.contract.Call(opts, &out, "last")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Last is a free data retrieval call binding the contract method 0x47799da8.
//
// Solidity: function last() view returns(uint256)
func (_OptimismBridge *OptimismBridgeSession) Last() (*big.Int, error) {
	return _OptimismBridge.Contract.Last(&_OptimismBridge.CallOpts)
}

// Last is a free data retrieval call binding the contract method 0x47799da8.
//
// Solidity: function last() view returns(uint256)
func (_OptimismBridge *OptimismBridgeCallerSession) Last() (*big.Int, error) {
	return _OptimismBridge.Contract.Last(&_OptimismBridge.CallOpts)
}

// LockedForGoerliETH is a free data retrieval call binding the contract method 0xe3d66f68.
//
// Solidity: function lockedForGoerliETH(address ) view returns(uint256)
func (_OptimismBridge *OptimismBridgeCaller) LockedForGoerliETH(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _OptimismBridge.contract.Call(opts, &out, "lockedForGoerliETH", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LockedForGoerliETH is a free data retrieval call binding the contract method 0xe3d66f68.
//
// Solidity: function lockedForGoerliETH(address ) view returns(uint256)
func (_OptimismBridge *OptimismBridgeSession) LockedForGoerliETH(arg0 common.Address) (*big.Int, error) {
	return _OptimismBridge.Contract.LockedForGoerliETH(&_OptimismBridge.CallOpts, arg0)
}

// LockedForGoerliETH is a free data retrieval call binding the contract method 0xe3d66f68.
//
// Solidity: function lockedForGoerliETH(address ) view returns(uint256)
func (_OptimismBridge *OptimismBridgeCallerSession) LockedForGoerliETH(arg0 common.Address) (*big.Int, error) {
	return _OptimismBridge.Contract.LockedForGoerliETH(&_OptimismBridge.CallOpts, arg0)
}

// OptimismBridgedETH is a free data retrieval call binding the contract method 0xc5aaccf7.
//
// Solidity: function optimismBridgedETH(address ) view returns(uint256)
func (_OptimismBridge *OptimismBridgeCaller) OptimismBridgedETH(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _OptimismBridge.contract.Call(opts, &out, "optimismBridgedETH", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OptimismBridgedETH is a free data retrieval call binding the contract method 0xc5aaccf7.
//
// Solidity: function optimismBridgedETH(address ) view returns(uint256)
func (_OptimismBridge *OptimismBridgeSession) OptimismBridgedETH(arg0 common.Address) (*big.Int, error) {
	return _OptimismBridge.Contract.OptimismBridgedETH(&_OptimismBridge.CallOpts, arg0)
}

// OptimismBridgedETH is a free data retrieval call binding the contract method 0xc5aaccf7.
//
// Solidity: function optimismBridgedETH(address ) view returns(uint256)
func (_OptimismBridge *OptimismBridgeCallerSession) OptimismBridgedETH(arg0 common.Address) (*big.Int, error) {
	return _OptimismBridge.Contract.OptimismBridgedETH(&_OptimismBridge.CallOpts, arg0)
}

// Queue is a free data retrieval call binding the contract method 0xddf0b009.
//
// Solidity: function queue(uint256 ) view returns(address)
func (_OptimismBridge *OptimismBridgeCaller) Queue(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _OptimismBridge.contract.Call(opts, &out, "queue", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Queue is a free data retrieval call binding the contract method 0xddf0b009.
//
// Solidity: function queue(uint256 ) view returns(address)
func (_OptimismBridge *OptimismBridgeSession) Queue(arg0 *big.Int) (common.Address, error) {
	return _OptimismBridge.Contract.Queue(&_OptimismBridge.CallOpts, arg0)
}

// Queue is a free data retrieval call binding the contract method 0xddf0b009.
//
// Solidity: function queue(uint256 ) view returns(address)
func (_OptimismBridge *OptimismBridgeCallerSession) Queue(arg0 *big.Int) (common.Address, error) {
	return _OptimismBridge.Contract.Queue(&_OptimismBridge.CallOpts, arg0)
}

// Dequeue is a paid mutator transaction binding the contract method 0x957908d1.
//
// Solidity: function dequeue() returns()
func (_OptimismBridge *OptimismBridgeTransactor) Dequeue(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OptimismBridge.contract.Transact(opts, "dequeue")
}

// Dequeue is a paid mutator transaction binding the contract method 0x957908d1.
//
// Solidity: function dequeue() returns()
func (_OptimismBridge *OptimismBridgeSession) Dequeue() (*types.Transaction, error) {
	return _OptimismBridge.Contract.Dequeue(&_OptimismBridge.TransactOpts)
}

// Dequeue is a paid mutator transaction binding the contract method 0x957908d1.
//
// Solidity: function dequeue() returns()
func (_OptimismBridge *OptimismBridgeTransactorSession) Dequeue() (*types.Transaction, error) {
	return _OptimismBridge.Contract.Dequeue(&_OptimismBridge.TransactOpts)
}

// LockTokensForGoerli is a paid mutator transaction binding the contract method 0xc637aa97.
//
// Solidity: function lockTokensForGoerli(uint256 bridgeAmount) payable returns()
func (_OptimismBridge *OptimismBridgeTransactor) LockTokensForGoerli(opts *bind.TransactOpts, bridgeAmount *big.Int) (*types.Transaction, error) {
	return _OptimismBridge.contract.Transact(opts, "lockTokensForGoerli", bridgeAmount)
}

// LockTokensForGoerli is a paid mutator transaction binding the contract method 0xc637aa97.
//
// Solidity: function lockTokensForGoerli(uint256 bridgeAmount) payable returns()
func (_OptimismBridge *OptimismBridgeSession) LockTokensForGoerli(bridgeAmount *big.Int) (*types.Transaction, error) {
	return _OptimismBridge.Contract.LockTokensForGoerli(&_OptimismBridge.TransactOpts, bridgeAmount)
}

// LockTokensForGoerli is a paid mutator transaction binding the contract method 0xc637aa97.
//
// Solidity: function lockTokensForGoerli(uint256 bridgeAmount) payable returns()
func (_OptimismBridge *OptimismBridgeTransactorSession) LockTokensForGoerli(bridgeAmount *big.Int) (*types.Transaction, error) {
	return _OptimismBridge.Contract.LockTokensForGoerli(&_OptimismBridge.TransactOpts, bridgeAmount)
}

// OwnerAddBridgeLiqudity is a paid mutator transaction binding the contract method 0x8521eebe.
//
// Solidity: function ownerAddBridgeLiqudity() payable returns()
func (_OptimismBridge *OptimismBridgeTransactor) OwnerAddBridgeLiqudity(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OptimismBridge.contract.Transact(opts, "ownerAddBridgeLiqudity")
}

// OwnerAddBridgeLiqudity is a paid mutator transaction binding the contract method 0x8521eebe.
//
// Solidity: function ownerAddBridgeLiqudity() payable returns()
func (_OptimismBridge *OptimismBridgeSession) OwnerAddBridgeLiqudity() (*types.Transaction, error) {
	return _OptimismBridge.Contract.OwnerAddBridgeLiqudity(&_OptimismBridge.TransactOpts)
}

// OwnerAddBridgeLiqudity is a paid mutator transaction binding the contract method 0x8521eebe.
//
// Solidity: function ownerAddBridgeLiqudity() payable returns()
func (_OptimismBridge *OptimismBridgeTransactorSession) OwnerAddBridgeLiqudity() (*types.Transaction, error) {
	return _OptimismBridge.Contract.OwnerAddBridgeLiqudity(&_OptimismBridge.TransactOpts)
}

// OwnerRemoveBridgeLiqudity is a paid mutator transaction binding the contract method 0x08dd057a.
//
// Solidity: function ownerRemoveBridgeLiqudity() returns()
func (_OptimismBridge *OptimismBridgeTransactor) OwnerRemoveBridgeLiqudity(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OptimismBridge.contract.Transact(opts, "ownerRemoveBridgeLiqudity")
}

// OwnerRemoveBridgeLiqudity is a paid mutator transaction binding the contract method 0x08dd057a.
//
// Solidity: function ownerRemoveBridgeLiqudity() returns()
func (_OptimismBridge *OptimismBridgeSession) OwnerRemoveBridgeLiqudity() (*types.Transaction, error) {
	return _OptimismBridge.Contract.OwnerRemoveBridgeLiqudity(&_OptimismBridge.TransactOpts)
}

// OwnerRemoveBridgeLiqudity is a paid mutator transaction binding the contract method 0x08dd057a.
//
// Solidity: function ownerRemoveBridgeLiqudity() returns()
func (_OptimismBridge *OptimismBridgeTransactorSession) OwnerRemoveBridgeLiqudity() (*types.Transaction, error) {
	return _OptimismBridge.Contract.OwnerRemoveBridgeLiqudity(&_OptimismBridge.TransactOpts)
}

// OwnerUnlockOptimismETH is a paid mutator transaction binding the contract method 0xcfd121e6.
//
// Solidity: function ownerUnlockOptimismETH(address userToBridge, uint256 lockedForOptimismETH) returns()
func (_OptimismBridge *OptimismBridgeTransactor) OwnerUnlockOptimismETH(opts *bind.TransactOpts, userToBridge common.Address, lockedForOptimismETH *big.Int) (*types.Transaction, error) {
	return _OptimismBridge.contract.Transact(opts, "ownerUnlockOptimismETH", userToBridge, lockedForOptimismETH)
}

// OwnerUnlockOptimismETH is a paid mutator transaction binding the contract method 0xcfd121e6.
//
// Solidity: function ownerUnlockOptimismETH(address userToBridge, uint256 lockedForOptimismETH) returns()
func (_OptimismBridge *OptimismBridgeSession) OwnerUnlockOptimismETH(userToBridge common.Address, lockedForOptimismETH *big.Int) (*types.Transaction, error) {
	return _OptimismBridge.Contract.OwnerUnlockOptimismETH(&_OptimismBridge.TransactOpts, userToBridge, lockedForOptimismETH)
}

// OwnerUnlockOptimismETH is a paid mutator transaction binding the contract method 0xcfd121e6.
//
// Solidity: function ownerUnlockOptimismETH(address userToBridge, uint256 lockedForOptimismETH) returns()
func (_OptimismBridge *OptimismBridgeTransactorSession) OwnerUnlockOptimismETH(userToBridge common.Address, lockedForOptimismETH *big.Int) (*types.Transaction, error) {
	return _OptimismBridge.Contract.OwnerUnlockOptimismETH(&_OptimismBridge.TransactOpts, userToBridge, lockedForOptimismETH)
}
