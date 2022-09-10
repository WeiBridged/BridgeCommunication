// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

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

// MainMetaData contains all meta data concerning the Main contract.
var MainMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"bridgeEmpty\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"bridgeOnOtherSideNeedsLiqudity\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"bridgedAlready\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"msgValueDoesNotCoverFee\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"msgValueLessThan1000\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"msgValueZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"notExternalBridge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"notOwnerAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ownerBridgeUsersBeforeWithdraw\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"queueIsEmpty\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"queueNotEmpty\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dequeue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"first\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"goerliBridgedETH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"last\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"bridgeAmount\",\"type\":\"uint256\"}],\"name\":\"lockTokensForOptimism\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"lockedForOptimismETH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ownerAddBridgeLiqudity\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ownerRemoveBridgeLiqudity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"userToBridge\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"lockedForGoerliETH\",\"type\":\"uint256\"}],\"name\":\"ownerUnlockGoerliETH\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"queue\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60a0604052600160045534801561001557600080fd5b503373ffffffffffffffffffffffffffffffffffffffff1660808173ffffffffffffffffffffffffffffffffffffffff1681525050608051610c3561008f6000396000818161025801528181610317015281816103800152818161051a015281816105db0152818161070a01526108670152610c356000f3fe60806040526004361061009c5760003560e01c80638521eebe116100645780638521eebe14610174578063957908d11461017e578063abb6b5a114610195578063b4a99a4e146101d2578063ddf0b009146101fd578063e76438ab1461023a5761009c565b806308dd057a146100a157806330b21283146100b85780633df4ddf4146100e157806347799da81461010c5780637bb1fdd214610137575b600080fd5b3480156100ad57600080fd5b506100b6610256565b005b3480156100c457600080fd5b506100df60048036038101906100da91906109d8565b61037e565b005b3480156100ed57600080fd5b506100f66104f4565b6040516101039190610a27565b60405180910390f35b34801561011857600080fd5b506101216104fa565b60405161012e9190610a27565b60405180910390f35b34801561014357600080fd5b5061015e60048036038101906101599190610a42565b610500565b60405161016b9190610a27565b60405180910390f35b61017c610518565b005b34801561018a57600080fd5b506101936105d9565b005b3480156101a157600080fd5b506101bc60048036038101906101b79190610a42565b6106f0565b6040516101c99190610a27565b60405180910390f35b3480156101de57600080fd5b506101e7610708565b6040516101f49190610a7e565b60405180910390f35b34801561020957600080fd5b50610224600480360381019061021f9190610a99565b61072c565b6040516102319190610a7e565b60405180910390f35b610254600480360381019061024f9190610a99565b61075f565b005b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146102db576040517f1ee8575e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60004703610315576040517f7a1f291700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166108fc479081150290604051600060405180830381858888f1935050505015801561037b573d6000803e3d6000fd5b50565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610403576040517f1ee8575e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054826104509190610af5565b905080600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546104a19190610b29565b925050819055508273ffffffffffffffffffffffffffffffffffffffff166108fc829081150290604051600060405180830381858888f193505050501580156104ee573d6000803e3d6000fd5b50505050565b60045481565b60035481565b60016020528060005260406000206000915090505481565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461059d576040517f1ee8575e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600034036105d7576040517f0cdb5d8500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461065e576040517f1ee8575e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600454600354101561069c576040517f5e61bb8700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60026000600454815260200190815260200160002060006101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690556001600460008282546106e79190610b29565b92505081905550565b60006020528060005260406000206000915090505481565b7f000000000000000000000000000000000000000000000000000000000000000081565b60026020528060005260406000206000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6103e881101561079b576040517fa9f2d61800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6103e8816103eb6107ac9190610b5d565b6107b69190610bce565b34146107ee576040517f400f9fb800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6103eb346103e86107ff9190610b5d565b6108099190610bce565b6000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546108569190610b29565b925050819055506108656108cf565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166108fc349081150290604051600060405180830381858888f193505050501580156108cb573d6000803e3d6000fd5b5050565b6001600360008282546108e29190610b29565b925050819055503360026000600354815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061096f82610944565b9050919050565b61097f81610964565b811461098a57600080fd5b50565b60008135905061099c81610976565b92915050565b6000819050919050565b6109b5816109a2565b81146109c057600080fd5b50565b6000813590506109d2816109ac565b92915050565b600080604083850312156109ef576109ee61093f565b5b60006109fd8582860161098d565b9250506020610a0e858286016109c3565b9150509250929050565b610a21816109a2565b82525050565b6000602082019050610a3c6000830184610a18565b92915050565b600060208284031215610a5857610a5761093f565b5b6000610a668482850161098d565b91505092915050565b610a7881610964565b82525050565b6000602082019050610a936000830184610a6f565b92915050565b600060208284031215610aaf57610aae61093f565b5b6000610abd848285016109c3565b91505092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000610b00826109a2565b9150610b0b836109a2565b9250828203905081811115610b2357610b22610ac6565b5b92915050565b6000610b34826109a2565b9150610b3f836109a2565b9250828201905080821115610b5757610b56610ac6565b5b92915050565b6000610b68826109a2565b9150610b73836109a2565b9250828202610b81816109a2565b91508282048414831517610b9857610b97610ac6565b5b5092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b6000610bd9826109a2565b9150610be4836109a2565b925082610bf457610bf3610b9f565b5b82820490509291505056fea2646970667358221220ef6cd583253406f0a5eef208cb9e67b4da7355007f4b8f1146245a28cec7971364736f6c63430008110033",
}

// MainABI is the input ABI used to generate the binding from.
// Deprecated: Use MainMetaData.ABI instead.
var MainABI = MainMetaData.ABI

// MainBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MainMetaData.Bin instead.
var MainBin = MainMetaData.Bin

// DeployMain deploys a new Ethereum contract, binding an instance of Main to it.
func DeployMain(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Main, error) {
	parsed, err := MainMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MainBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Main{MainCaller: MainCaller{contract: contract}, MainTransactor: MainTransactor{contract: contract}, MainFilterer: MainFilterer{contract: contract}}, nil
}

// Main is an auto generated Go binding around an Ethereum contract.
type Main struct {
	MainCaller     // Read-only binding to the contract
	MainTransactor // Write-only binding to the contract
	MainFilterer   // Log filterer for contract events
}

// MainCaller is an auto generated read-only Go binding around an Ethereum contract.
type MainCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MainTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MainTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MainFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MainFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MainSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MainSession struct {
	Contract     *Main             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MainCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MainCallerSession struct {
	Contract *MainCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MainTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MainTransactorSession struct {
	Contract     *MainTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MainRaw is an auto generated low-level Go binding around an Ethereum contract.
type MainRaw struct {
	Contract *Main // Generic contract binding to access the raw methods on
}

// MainCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MainCallerRaw struct {
	Contract *MainCaller // Generic read-only contract binding to access the raw methods on
}

// MainTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MainTransactorRaw struct {
	Contract *MainTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMain creates a new instance of Main, bound to a specific deployed contract.
func NewMain(address common.Address, backend bind.ContractBackend) (*Main, error) {
	contract, err := bindMain(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Main{MainCaller: MainCaller{contract: contract}, MainTransactor: MainTransactor{contract: contract}, MainFilterer: MainFilterer{contract: contract}}, nil
}

// NewMainCaller creates a new read-only instance of Main, bound to a specific deployed contract.
func NewMainCaller(address common.Address, caller bind.ContractCaller) (*MainCaller, error) {
	contract, err := bindMain(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MainCaller{contract: contract}, nil
}

// NewMainTransactor creates a new write-only instance of Main, bound to a specific deployed contract.
func NewMainTransactor(address common.Address, transactor bind.ContractTransactor) (*MainTransactor, error) {
	contract, err := bindMain(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MainTransactor{contract: contract}, nil
}

// NewMainFilterer creates a new log filterer instance of Main, bound to a specific deployed contract.
func NewMainFilterer(address common.Address, filterer bind.ContractFilterer) (*MainFilterer, error) {
	contract, err := bindMain(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MainFilterer{contract: contract}, nil
}

// bindMain binds a generic wrapper to an already deployed contract.
func bindMain(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MainABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Main *MainRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Main.Contract.MainCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Main *MainRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Main.Contract.MainTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Main *MainRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Main.Contract.MainTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Main *MainCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Main.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Main *MainTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Main.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Main *MainTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Main.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0xb4a99a4e.
//
// Solidity: function Owner() view returns(address)
func (_Main *MainCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "Owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0xb4a99a4e.
//
// Solidity: function Owner() view returns(address)
func (_Main *MainSession) Owner() (common.Address, error) {
	return _Main.Contract.Owner(&_Main.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0xb4a99a4e.
//
// Solidity: function Owner() view returns(address)
func (_Main *MainCallerSession) Owner() (common.Address, error) {
	return _Main.Contract.Owner(&_Main.CallOpts)
}

// First is a free data retrieval call binding the contract method 0x3df4ddf4.
//
// Solidity: function first() view returns(uint256)
func (_Main *MainCaller) First(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "first")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// First is a free data retrieval call binding the contract method 0x3df4ddf4.
//
// Solidity: function first() view returns(uint256)
func (_Main *MainSession) First() (*big.Int, error) {
	return _Main.Contract.First(&_Main.CallOpts)
}

// First is a free data retrieval call binding the contract method 0x3df4ddf4.
//
// Solidity: function first() view returns(uint256)
func (_Main *MainCallerSession) First() (*big.Int, error) {
	return _Main.Contract.First(&_Main.CallOpts)
}

// GoerliBridgedETH is a free data retrieval call binding the contract method 0x7bb1fdd2.
//
// Solidity: function goerliBridgedETH(address ) view returns(uint256)
func (_Main *MainCaller) GoerliBridgedETH(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "goerliBridgedETH", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GoerliBridgedETH is a free data retrieval call binding the contract method 0x7bb1fdd2.
//
// Solidity: function goerliBridgedETH(address ) view returns(uint256)
func (_Main *MainSession) GoerliBridgedETH(arg0 common.Address) (*big.Int, error) {
	return _Main.Contract.GoerliBridgedETH(&_Main.CallOpts, arg0)
}

// GoerliBridgedETH is a free data retrieval call binding the contract method 0x7bb1fdd2.
//
// Solidity: function goerliBridgedETH(address ) view returns(uint256)
func (_Main *MainCallerSession) GoerliBridgedETH(arg0 common.Address) (*big.Int, error) {
	return _Main.Contract.GoerliBridgedETH(&_Main.CallOpts, arg0)
}

// Last is a free data retrieval call binding the contract method 0x47799da8.
//
// Solidity: function last() view returns(uint256)
func (_Main *MainCaller) Last(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "last")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Last is a free data retrieval call binding the contract method 0x47799da8.
//
// Solidity: function last() view returns(uint256)
func (_Main *MainSession) Last() (*big.Int, error) {
	return _Main.Contract.Last(&_Main.CallOpts)
}

// Last is a free data retrieval call binding the contract method 0x47799da8.
//
// Solidity: function last() view returns(uint256)
func (_Main *MainCallerSession) Last() (*big.Int, error) {
	return _Main.Contract.Last(&_Main.CallOpts)
}

// LockedForOptimismETH is a free data retrieval call binding the contract method 0xabb6b5a1.
//
// Solidity: function lockedForOptimismETH(address ) view returns(uint256)
func (_Main *MainCaller) LockedForOptimismETH(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "lockedForOptimismETH", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LockedForOptimismETH is a free data retrieval call binding the contract method 0xabb6b5a1.
//
// Solidity: function lockedForOptimismETH(address ) view returns(uint256)
func (_Main *MainSession) LockedForOptimismETH(arg0 common.Address) (*big.Int, error) {
	return _Main.Contract.LockedForOptimismETH(&_Main.CallOpts, arg0)
}

// LockedForOptimismETH is a free data retrieval call binding the contract method 0xabb6b5a1.
//
// Solidity: function lockedForOptimismETH(address ) view returns(uint256)
func (_Main *MainCallerSession) LockedForOptimismETH(arg0 common.Address) (*big.Int, error) {
	return _Main.Contract.LockedForOptimismETH(&_Main.CallOpts, arg0)
}

// Queue is a free data retrieval call binding the contract method 0xddf0b009.
//
// Solidity: function queue(uint256 ) view returns(address)
func (_Main *MainCaller) Queue(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "queue", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Queue is a free data retrieval call binding the contract method 0xddf0b009.
//
// Solidity: function queue(uint256 ) view returns(address)
func (_Main *MainSession) Queue(arg0 *big.Int) (common.Address, error) {
	return _Main.Contract.Queue(&_Main.CallOpts, arg0)
}

// Queue is a free data retrieval call binding the contract method 0xddf0b009.
//
// Solidity: function queue(uint256 ) view returns(address)
func (_Main *MainCallerSession) Queue(arg0 *big.Int) (common.Address, error) {
	return _Main.Contract.Queue(&_Main.CallOpts, arg0)
}

// Dequeue is a paid mutator transaction binding the contract method 0x957908d1.
//
// Solidity: function dequeue() returns()
func (_Main *MainTransactor) Dequeue(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "dequeue")
}

// Dequeue is a paid mutator transaction binding the contract method 0x957908d1.
//
// Solidity: function dequeue() returns()
func (_Main *MainSession) Dequeue() (*types.Transaction, error) {
	return _Main.Contract.Dequeue(&_Main.TransactOpts)
}

// Dequeue is a paid mutator transaction binding the contract method 0x957908d1.
//
// Solidity: function dequeue() returns()
func (_Main *MainTransactorSession) Dequeue() (*types.Transaction, error) {
	return _Main.Contract.Dequeue(&_Main.TransactOpts)
}

// LockTokensForOptimism is a paid mutator transaction binding the contract method 0xe76438ab.
//
// Solidity: function lockTokensForOptimism(uint256 bridgeAmount) payable returns()
func (_Main *MainTransactor) LockTokensForOptimism(opts *bind.TransactOpts, bridgeAmount *big.Int) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "lockTokensForOptimism", bridgeAmount)
}

// LockTokensForOptimism is a paid mutator transaction binding the contract method 0xe76438ab.
//
// Solidity: function lockTokensForOptimism(uint256 bridgeAmount) payable returns()
func (_Main *MainSession) LockTokensForOptimism(bridgeAmount *big.Int) (*types.Transaction, error) {
	return _Main.Contract.LockTokensForOptimism(&_Main.TransactOpts, bridgeAmount)
}

// LockTokensForOptimism is a paid mutator transaction binding the contract method 0xe76438ab.
//
// Solidity: function lockTokensForOptimism(uint256 bridgeAmount) payable returns()
func (_Main *MainTransactorSession) LockTokensForOptimism(bridgeAmount *big.Int) (*types.Transaction, error) {
	return _Main.Contract.LockTokensForOptimism(&_Main.TransactOpts, bridgeAmount)
}

// OwnerAddBridgeLiqudity is a paid mutator transaction binding the contract method 0x8521eebe.
//
// Solidity: function ownerAddBridgeLiqudity() payable returns()
func (_Main *MainTransactor) OwnerAddBridgeLiqudity(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "ownerAddBridgeLiqudity")
}

// OwnerAddBridgeLiqudity is a paid mutator transaction binding the contract method 0x8521eebe.
//
// Solidity: function ownerAddBridgeLiqudity() payable returns()
func (_Main *MainSession) OwnerAddBridgeLiqudity() (*types.Transaction, error) {
	return _Main.Contract.OwnerAddBridgeLiqudity(&_Main.TransactOpts)
}

// OwnerAddBridgeLiqudity is a paid mutator transaction binding the contract method 0x8521eebe.
//
// Solidity: function ownerAddBridgeLiqudity() payable returns()
func (_Main *MainTransactorSession) OwnerAddBridgeLiqudity() (*types.Transaction, error) {
	return _Main.Contract.OwnerAddBridgeLiqudity(&_Main.TransactOpts)
}

// OwnerRemoveBridgeLiqudity is a paid mutator transaction binding the contract method 0x08dd057a.
//
// Solidity: function ownerRemoveBridgeLiqudity() returns()
func (_Main *MainTransactor) OwnerRemoveBridgeLiqudity(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "ownerRemoveBridgeLiqudity")
}

// OwnerRemoveBridgeLiqudity is a paid mutator transaction binding the contract method 0x08dd057a.
//
// Solidity: function ownerRemoveBridgeLiqudity() returns()
func (_Main *MainSession) OwnerRemoveBridgeLiqudity() (*types.Transaction, error) {
	return _Main.Contract.OwnerRemoveBridgeLiqudity(&_Main.TransactOpts)
}

// OwnerRemoveBridgeLiqudity is a paid mutator transaction binding the contract method 0x08dd057a.
//
// Solidity: function ownerRemoveBridgeLiqudity() returns()
func (_Main *MainTransactorSession) OwnerRemoveBridgeLiqudity() (*types.Transaction, error) {
	return _Main.Contract.OwnerRemoveBridgeLiqudity(&_Main.TransactOpts)
}

// OwnerUnlockGoerliETH is a paid mutator transaction binding the contract method 0x30b21283.
//
// Solidity: function ownerUnlockGoerliETH(address userToBridge, uint256 lockedForGoerliETH) returns()
func (_Main *MainTransactor) OwnerUnlockGoerliETH(opts *bind.TransactOpts, userToBridge common.Address, lockedForGoerliETH *big.Int) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "ownerUnlockGoerliETH", userToBridge, lockedForGoerliETH)
}

// OwnerUnlockGoerliETH is a paid mutator transaction binding the contract method 0x30b21283.
//
// Solidity: function ownerUnlockGoerliETH(address userToBridge, uint256 lockedForGoerliETH) returns()
func (_Main *MainSession) OwnerUnlockGoerliETH(userToBridge common.Address, lockedForGoerliETH *big.Int) (*types.Transaction, error) {
	return _Main.Contract.OwnerUnlockGoerliETH(&_Main.TransactOpts, userToBridge, lockedForGoerliETH)
}

// OwnerUnlockGoerliETH is a paid mutator transaction binding the contract method 0x30b21283.
//
// Solidity: function ownerUnlockGoerliETH(address userToBridge, uint256 lockedForGoerliETH) returns()
func (_Main *MainTransactorSession) OwnerUnlockGoerliETH(userToBridge common.Address, lockedForGoerliETH *big.Int) (*types.Transaction, error) {
	return _Main.Contract.OwnerUnlockGoerliETH(&_Main.TransactOpts, userToBridge, lockedForGoerliETH)
}
