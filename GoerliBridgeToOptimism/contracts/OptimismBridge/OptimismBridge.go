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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"bridgeEmpty\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"bridgeOnOtherSideNeedsLiqudity\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"msgValueNot1003\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"msgValueZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"notExternalBridge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"notOwnerAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"queueIsEmpty\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"queueNotEmpty\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dequeue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"first\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"last\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lockTokensForGoerli\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ownerAddBridgeLiqudity\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ownerRemoveBridgeLiqudity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"userToBridge\",\"type\":\"address\"}],\"name\":\"ownerUnlockOptimismETH\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"queue\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60a0604052600160025534801561001557600080fd5b503373ffffffffffffffffffffffffffffffffffffffff1660808173ffffffffffffffffffffffffffffffffffffffff16815250506080516108e161008f600039600081816101b60152818161028701528181610346015281816103bb0152818161047c01528181610592015261062c01526108e16000f3fe6080604052600436106100865760003560e01c80638521eebe116100595780638521eebe14610121578063957908d11461012b578063b4a99a4e14610142578063ddf0b0091461016d578063f67744c4146101aa57610086565b806304c5366e1461008b57806308dd057a146100b45780633df4ddf4146100cb57806347799da8146100f6575b600080fd5b34801561009757600080fd5b506100b260048036038101906100ad9190610764565b6101b4565b005b3480156100c057600080fd5b506100c9610285565b005b3480156100d757600080fd5b506100e06103ad565b6040516100ed91906107aa565b60405180910390f35b34801561010257600080fd5b5061010b6103b3565b60405161011891906107aa565b60405180910390f35b6101296103b9565b005b34801561013757600080fd5b5061014061047a565b005b34801561014e57600080fd5b50610157610590565b60405161016491906107d4565b60405180910390f35b34801561017957600080fd5b50610194600480360381019061018f919061081b565b6105b4565b6040516101a191906107d4565b60405180910390f35b6101b26105e7565b005b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610239576040517f1ee8575e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff166108fc6103e89081150290604051600060405180830381858888f19350505050158015610281573d6000803e3d6000fd5b5050565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461030a576040517f1ee8575e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60004703610344576040517f7a1f291700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166108fc479081150290604051600060405180830381858888f193505050501580156103aa573d6000803e3d6000fd5b50565b60025481565b60015481565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461043e576040517f1ee8575e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60003403610478576040517f0cdb5d8500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146104ff576040517f1ee8575e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600254600154101561053d576040517f5e61bb8700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600080600254815260200190815260200160002060006101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690556001600260008282546105879190610877565b92505081905550565b7f000000000000000000000000000000000000000000000000000000000000000081565b60006020528060005260406000206000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6103eb3414610622576040517fa43d860800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61062a610693565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166108fc349081150290604051600060405180830381858888f19350505050158015610690573d6000803e3d6000fd5b50565b60018060008282546106a59190610877565b9250508190555033600080600154815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061073182610706565b9050919050565b61074181610726565b811461074c57600080fd5b50565b60008135905061075e81610738565b92915050565b60006020828403121561077a57610779610701565b5b60006107888482850161074f565b91505092915050565b6000819050919050565b6107a481610791565b82525050565b60006020820190506107bf600083018461079b565b92915050565b6107ce81610726565b82525050565b60006020820190506107e960008301846107c5565b92915050565b6107f881610791565b811461080357600080fd5b50565b600081359050610815816107ef565b92915050565b60006020828403121561083157610830610701565b5b600061083f84828501610806565b91505092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600061088282610791565b915061088d83610791565b92508282019050808211156108a5576108a4610848565b5b9291505056fea2646970667358221220bc14f3d238d8a256cc578b1b37850d613449acf133b7ab4e736024f5fa03d64364736f6c63430008110033",
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

// LockTokensForGoerli is a paid mutator transaction binding the contract method 0xf67744c4.
//
// Solidity: function lockTokensForGoerli() payable returns()
func (_OptimismBridge *OptimismBridgeTransactor) LockTokensForGoerli(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OptimismBridge.contract.Transact(opts, "lockTokensForGoerli")
}

// LockTokensForGoerli is a paid mutator transaction binding the contract method 0xf67744c4.
//
// Solidity: function lockTokensForGoerli() payable returns()
func (_OptimismBridge *OptimismBridgeSession) LockTokensForGoerli() (*types.Transaction, error) {
	return _OptimismBridge.Contract.LockTokensForGoerli(&_OptimismBridge.TransactOpts)
}

// LockTokensForGoerli is a paid mutator transaction binding the contract method 0xf67744c4.
//
// Solidity: function lockTokensForGoerli() payable returns()
func (_OptimismBridge *OptimismBridgeTransactorSession) LockTokensForGoerli() (*types.Transaction, error) {
	return _OptimismBridge.Contract.LockTokensForGoerli(&_OptimismBridge.TransactOpts)
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

// OwnerUnlockOptimismETH is a paid mutator transaction binding the contract method 0x04c5366e.
//
// Solidity: function ownerUnlockOptimismETH(address userToBridge) returns()
func (_OptimismBridge *OptimismBridgeTransactor) OwnerUnlockOptimismETH(opts *bind.TransactOpts, userToBridge common.Address) (*types.Transaction, error) {
	return _OptimismBridge.contract.Transact(opts, "ownerUnlockOptimismETH", userToBridge)
}

// OwnerUnlockOptimismETH is a paid mutator transaction binding the contract method 0x04c5366e.
//
// Solidity: function ownerUnlockOptimismETH(address userToBridge) returns()
func (_OptimismBridge *OptimismBridgeSession) OwnerUnlockOptimismETH(userToBridge common.Address) (*types.Transaction, error) {
	return _OptimismBridge.Contract.OwnerUnlockOptimismETH(&_OptimismBridge.TransactOpts, userToBridge)
}

// OwnerUnlockOptimismETH is a paid mutator transaction binding the contract method 0x04c5366e.
//
// Solidity: function ownerUnlockOptimismETH(address userToBridge) returns()
func (_OptimismBridge *OptimismBridgeTransactorSession) OwnerUnlockOptimismETH(userToBridge common.Address) (*types.Transaction, error) {
	return _OptimismBridge.Contract.OwnerUnlockOptimismETH(&_OptimismBridge.TransactOpts, userToBridge)
}
