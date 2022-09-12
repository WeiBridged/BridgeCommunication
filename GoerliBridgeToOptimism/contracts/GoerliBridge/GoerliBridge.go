// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package goerliBridge

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

// GoerliBridgeMetaData contains all meta data concerning the GoerliBridge contract.
var GoerliBridgeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"bridgeEmpty\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"bridgeOnOtherSideNeedsLiqudity\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"bridgedAlready\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"msgValueNot1003\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"msgValueZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"notExternalBridge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"notOwnerAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ownerBridgeUsersBeforeWithdraw\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"queueIsEmpty\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"queueNotEmpty\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dequeue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"first\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"last\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lockTokensForOptimism\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ownerAddBridgeLiqudity\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ownerRemoveBridgeLiqudity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"userToBridge\",\"type\":\"address\"}],\"name\":\"ownerUnlockGoerliETH\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"queue\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60a0604052600160025534801561001557600080fd5b503373ffffffffffffffffffffffffffffffffffffffff1660808173ffffffffffffffffffffffffffffffffffffffff16815250506080516108e161008f600039600081816101b601528181610275015281816102de015281816103fe0152818161046701528181610528015261063e01526108e16000f3fe6080604052600436106100865760003560e01c80634e22d208116100595780634e22d208146101215780638521eebe1461012b578063957908d114610135578063b4a99a4e1461014c578063ddf0b0091461017757610086565b806308dd057a1461008b57806315b5cbdf146100a25780633df4ddf4146100cb57806347799da8146100f6575b600080fd5b34801561009757600080fd5b506100a06101b4565b005b3480156100ae57600080fd5b506100c960048036038101906100c49190610764565b6102dc565b005b3480156100d757600080fd5b506100e06103ad565b6040516100ed91906107aa565b60405180910390f35b34801561010257600080fd5b5061010b6103b3565b60405161011891906107aa565b60405180910390f35b6101296103b9565b005b610133610465565b005b34801561014157600080fd5b5061014a610526565b005b34801561015857600080fd5b5061016161063c565b60405161016e91906107d4565b60405180910390f35b34801561018357600080fd5b5061019e6004803603810190610199919061081b565b610660565b6040516101ab91906107d4565b60405180910390f35b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610239576040517f1ee8575e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60004703610273576040517f7a1f291700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166108fc479081150290604051600060405180830381858888f193505050501580156102d9573d6000803e3d6000fd5b50565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610361576040517f1ee8575e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff166108fc6103e89081150290604051600060405180830381858888f193505050501580156103a9573d6000803e3d6000fd5b5050565b60025481565b60015481565b6103eb34146103f4576040517fa43d860800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6103fc610693565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166108fc349081150290604051600060405180830381858888f19350505050158015610462573d6000803e3d6000fd5b50565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146104ea576040517f1ee8575e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60003403610524576040517f0cdb5d8500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146105ab576040517f1ee8575e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60025460015410156105e9576040517f5e61bb8700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600080600254815260200190815260200160002060006101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690556001600260008282546106339190610877565b92505081905550565b7f000000000000000000000000000000000000000000000000000000000000000081565b60006020528060005260406000206000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60018060008282546106a59190610877565b9250508190555033600080600154815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061073182610706565b9050919050565b61074181610726565b811461074c57600080fd5b50565b60008135905061075e81610738565b92915050565b60006020828403121561077a57610779610701565b5b60006107888482850161074f565b91505092915050565b6000819050919050565b6107a481610791565b82525050565b60006020820190506107bf600083018461079b565b92915050565b6107ce81610726565b82525050565b60006020820190506107e960008301846107c5565b92915050565b6107f881610791565b811461080357600080fd5b50565b600081359050610815816107ef565b92915050565b60006020828403121561083157610830610701565b5b600061083f84828501610806565b91505092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600061088282610791565b915061088d83610791565b92508282019050808211156108a5576108a4610848565b5b9291505056fea2646970667358221220fc71edf8533f14c03ef4c4affdd594fc0bd32b85302c20bb5949066e9f7469dd64736f6c63430008110033",
}

// GoerliBridgeABI is the input ABI used to generate the binding from.
// Deprecated: Use GoerliBridgeMetaData.ABI instead.
var GoerliBridgeABI = GoerliBridgeMetaData.ABI

// GoerliBridgeBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use GoerliBridgeMetaData.Bin instead.
var GoerliBridgeBin = GoerliBridgeMetaData.Bin

// DeployGoerliBridge deploys a new Ethereum contract, binding an instance of GoerliBridge to it.
func DeployGoerliBridge(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *GoerliBridge, error) {
	parsed, err := GoerliBridgeMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(GoerliBridgeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &GoerliBridge{GoerliBridgeCaller: GoerliBridgeCaller{contract: contract}, GoerliBridgeTransactor: GoerliBridgeTransactor{contract: contract}, GoerliBridgeFilterer: GoerliBridgeFilterer{contract: contract}}, nil
}

// GoerliBridge is an auto generated Go binding around an Ethereum contract.
type GoerliBridge struct {
	GoerliBridgeCaller     // Read-only binding to the contract
	GoerliBridgeTransactor // Write-only binding to the contract
	GoerliBridgeFilterer   // Log filterer for contract events
}

// GoerliBridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type GoerliBridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GoerliBridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GoerliBridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GoerliBridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GoerliBridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GoerliBridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GoerliBridgeSession struct {
	Contract     *GoerliBridge     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GoerliBridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GoerliBridgeCallerSession struct {
	Contract *GoerliBridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// GoerliBridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GoerliBridgeTransactorSession struct {
	Contract     *GoerliBridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// GoerliBridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type GoerliBridgeRaw struct {
	Contract *GoerliBridge // Generic contract binding to access the raw methods on
}

// GoerliBridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GoerliBridgeCallerRaw struct {
	Contract *GoerliBridgeCaller // Generic read-only contract binding to access the raw methods on
}

// GoerliBridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GoerliBridgeTransactorRaw struct {
	Contract *GoerliBridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGoerliBridge creates a new instance of GoerliBridge, bound to a specific deployed contract.
func NewGoerliBridge(address common.Address, backend bind.ContractBackend) (*GoerliBridge, error) {
	contract, err := bindGoerliBridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GoerliBridge{GoerliBridgeCaller: GoerliBridgeCaller{contract: contract}, GoerliBridgeTransactor: GoerliBridgeTransactor{contract: contract}, GoerliBridgeFilterer: GoerliBridgeFilterer{contract: contract}}, nil
}

// NewGoerliBridgeCaller creates a new read-only instance of GoerliBridge, bound to a specific deployed contract.
func NewGoerliBridgeCaller(address common.Address, caller bind.ContractCaller) (*GoerliBridgeCaller, error) {
	contract, err := bindGoerliBridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GoerliBridgeCaller{contract: contract}, nil
}

// NewGoerliBridgeTransactor creates a new write-only instance of GoerliBridge, bound to a specific deployed contract.
func NewGoerliBridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*GoerliBridgeTransactor, error) {
	contract, err := bindGoerliBridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GoerliBridgeTransactor{contract: contract}, nil
}

// NewGoerliBridgeFilterer creates a new log filterer instance of GoerliBridge, bound to a specific deployed contract.
func NewGoerliBridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*GoerliBridgeFilterer, error) {
	contract, err := bindGoerliBridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GoerliBridgeFilterer{contract: contract}, nil
}

// bindGoerliBridge binds a generic wrapper to an already deployed contract.
func bindGoerliBridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GoerliBridgeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GoerliBridge *GoerliBridgeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GoerliBridge.Contract.GoerliBridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GoerliBridge *GoerliBridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GoerliBridge.Contract.GoerliBridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GoerliBridge *GoerliBridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GoerliBridge.Contract.GoerliBridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GoerliBridge *GoerliBridgeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GoerliBridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GoerliBridge *GoerliBridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GoerliBridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GoerliBridge *GoerliBridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GoerliBridge.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0xb4a99a4e.
//
// Solidity: function Owner() view returns(address)
func (_GoerliBridge *GoerliBridgeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GoerliBridge.contract.Call(opts, &out, "Owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0xb4a99a4e.
//
// Solidity: function Owner() view returns(address)
func (_GoerliBridge *GoerliBridgeSession) Owner() (common.Address, error) {
	return _GoerliBridge.Contract.Owner(&_GoerliBridge.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0xb4a99a4e.
//
// Solidity: function Owner() view returns(address)
func (_GoerliBridge *GoerliBridgeCallerSession) Owner() (common.Address, error) {
	return _GoerliBridge.Contract.Owner(&_GoerliBridge.CallOpts)
}

// First is a free data retrieval call binding the contract method 0x3df4ddf4.
//
// Solidity: function first() view returns(uint256)
func (_GoerliBridge *GoerliBridgeCaller) First(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GoerliBridge.contract.Call(opts, &out, "first")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// First is a free data retrieval call binding the contract method 0x3df4ddf4.
//
// Solidity: function first() view returns(uint256)
func (_GoerliBridge *GoerliBridgeSession) First() (*big.Int, error) {
	return _GoerliBridge.Contract.First(&_GoerliBridge.CallOpts)
}

// First is a free data retrieval call binding the contract method 0x3df4ddf4.
//
// Solidity: function first() view returns(uint256)
func (_GoerliBridge *GoerliBridgeCallerSession) First() (*big.Int, error) {
	return _GoerliBridge.Contract.First(&_GoerliBridge.CallOpts)
}

// Last is a free data retrieval call binding the contract method 0x47799da8.
//
// Solidity: function last() view returns(uint256)
func (_GoerliBridge *GoerliBridgeCaller) Last(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GoerliBridge.contract.Call(opts, &out, "last")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Last is a free data retrieval call binding the contract method 0x47799da8.
//
// Solidity: function last() view returns(uint256)
func (_GoerliBridge *GoerliBridgeSession) Last() (*big.Int, error) {
	return _GoerliBridge.Contract.Last(&_GoerliBridge.CallOpts)
}

// Last is a free data retrieval call binding the contract method 0x47799da8.
//
// Solidity: function last() view returns(uint256)
func (_GoerliBridge *GoerliBridgeCallerSession) Last() (*big.Int, error) {
	return _GoerliBridge.Contract.Last(&_GoerliBridge.CallOpts)
}

// Queue is a free data retrieval call binding the contract method 0xddf0b009.
//
// Solidity: function queue(uint256 ) view returns(address)
func (_GoerliBridge *GoerliBridgeCaller) Queue(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _GoerliBridge.contract.Call(opts, &out, "queue", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Queue is a free data retrieval call binding the contract method 0xddf0b009.
//
// Solidity: function queue(uint256 ) view returns(address)
func (_GoerliBridge *GoerliBridgeSession) Queue(arg0 *big.Int) (common.Address, error) {
	return _GoerliBridge.Contract.Queue(&_GoerliBridge.CallOpts, arg0)
}

// Queue is a free data retrieval call binding the contract method 0xddf0b009.
//
// Solidity: function queue(uint256 ) view returns(address)
func (_GoerliBridge *GoerliBridgeCallerSession) Queue(arg0 *big.Int) (common.Address, error) {
	return _GoerliBridge.Contract.Queue(&_GoerliBridge.CallOpts, arg0)
}

// Dequeue is a paid mutator transaction binding the contract method 0x957908d1.
//
// Solidity: function dequeue() returns()
func (_GoerliBridge *GoerliBridgeTransactor) Dequeue(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GoerliBridge.contract.Transact(opts, "dequeue")
}

// Dequeue is a paid mutator transaction binding the contract method 0x957908d1.
//
// Solidity: function dequeue() returns()
func (_GoerliBridge *GoerliBridgeSession) Dequeue() (*types.Transaction, error) {
	return _GoerliBridge.Contract.Dequeue(&_GoerliBridge.TransactOpts)
}

// Dequeue is a paid mutator transaction binding the contract method 0x957908d1.
//
// Solidity: function dequeue() returns()
func (_GoerliBridge *GoerliBridgeTransactorSession) Dequeue() (*types.Transaction, error) {
	return _GoerliBridge.Contract.Dequeue(&_GoerliBridge.TransactOpts)
}

// LockTokensForOptimism is a paid mutator transaction binding the contract method 0x4e22d208.
//
// Solidity: function lockTokensForOptimism() payable returns()
func (_GoerliBridge *GoerliBridgeTransactor) LockTokensForOptimism(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GoerliBridge.contract.Transact(opts, "lockTokensForOptimism")
}

// LockTokensForOptimism is a paid mutator transaction binding the contract method 0x4e22d208.
//
// Solidity: function lockTokensForOptimism() payable returns()
func (_GoerliBridge *GoerliBridgeSession) LockTokensForOptimism() (*types.Transaction, error) {
	return _GoerliBridge.Contract.LockTokensForOptimism(&_GoerliBridge.TransactOpts)
}

// LockTokensForOptimism is a paid mutator transaction binding the contract method 0x4e22d208.
//
// Solidity: function lockTokensForOptimism() payable returns()
func (_GoerliBridge *GoerliBridgeTransactorSession) LockTokensForOptimism() (*types.Transaction, error) {
	return _GoerliBridge.Contract.LockTokensForOptimism(&_GoerliBridge.TransactOpts)
}

// OwnerAddBridgeLiqudity is a paid mutator transaction binding the contract method 0x8521eebe.
//
// Solidity: function ownerAddBridgeLiqudity() payable returns()
func (_GoerliBridge *GoerliBridgeTransactor) OwnerAddBridgeLiqudity(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GoerliBridge.contract.Transact(opts, "ownerAddBridgeLiqudity")
}

// OwnerAddBridgeLiqudity is a paid mutator transaction binding the contract method 0x8521eebe.
//
// Solidity: function ownerAddBridgeLiqudity() payable returns()
func (_GoerliBridge *GoerliBridgeSession) OwnerAddBridgeLiqudity() (*types.Transaction, error) {
	return _GoerliBridge.Contract.OwnerAddBridgeLiqudity(&_GoerliBridge.TransactOpts)
}

// OwnerAddBridgeLiqudity is a paid mutator transaction binding the contract method 0x8521eebe.
//
// Solidity: function ownerAddBridgeLiqudity() payable returns()
func (_GoerliBridge *GoerliBridgeTransactorSession) OwnerAddBridgeLiqudity() (*types.Transaction, error) {
	return _GoerliBridge.Contract.OwnerAddBridgeLiqudity(&_GoerliBridge.TransactOpts)
}

// OwnerRemoveBridgeLiqudity is a paid mutator transaction binding the contract method 0x08dd057a.
//
// Solidity: function ownerRemoveBridgeLiqudity() returns()
func (_GoerliBridge *GoerliBridgeTransactor) OwnerRemoveBridgeLiqudity(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GoerliBridge.contract.Transact(opts, "ownerRemoveBridgeLiqudity")
}

// OwnerRemoveBridgeLiqudity is a paid mutator transaction binding the contract method 0x08dd057a.
//
// Solidity: function ownerRemoveBridgeLiqudity() returns()
func (_GoerliBridge *GoerliBridgeSession) OwnerRemoveBridgeLiqudity() (*types.Transaction, error) {
	return _GoerliBridge.Contract.OwnerRemoveBridgeLiqudity(&_GoerliBridge.TransactOpts)
}

// OwnerRemoveBridgeLiqudity is a paid mutator transaction binding the contract method 0x08dd057a.
//
// Solidity: function ownerRemoveBridgeLiqudity() returns()
func (_GoerliBridge *GoerliBridgeTransactorSession) OwnerRemoveBridgeLiqudity() (*types.Transaction, error) {
	return _GoerliBridge.Contract.OwnerRemoveBridgeLiqudity(&_GoerliBridge.TransactOpts)
}

// OwnerUnlockGoerliETH is a paid mutator transaction binding the contract method 0x15b5cbdf.
//
// Solidity: function ownerUnlockGoerliETH(address userToBridge) returns()
func (_GoerliBridge *GoerliBridgeTransactor) OwnerUnlockGoerliETH(opts *bind.TransactOpts, userToBridge common.Address) (*types.Transaction, error) {
	return _GoerliBridge.contract.Transact(opts, "ownerUnlockGoerliETH", userToBridge)
}

// OwnerUnlockGoerliETH is a paid mutator transaction binding the contract method 0x15b5cbdf.
//
// Solidity: function ownerUnlockGoerliETH(address userToBridge) returns()
func (_GoerliBridge *GoerliBridgeSession) OwnerUnlockGoerliETH(userToBridge common.Address) (*types.Transaction, error) {
	return _GoerliBridge.Contract.OwnerUnlockGoerliETH(&_GoerliBridge.TransactOpts, userToBridge)
}

// OwnerUnlockGoerliETH is a paid mutator transaction binding the contract method 0x15b5cbdf.
//
// Solidity: function ownerUnlockGoerliETH(address userToBridge) returns()
func (_GoerliBridge *GoerliBridgeTransactorSession) OwnerUnlockGoerliETH(userToBridge common.Address) (*types.Transaction, error) {
	return _GoerliBridge.Contract.OwnerUnlockGoerliETH(&_GoerliBridge.TransactOpts, userToBridge)
}
