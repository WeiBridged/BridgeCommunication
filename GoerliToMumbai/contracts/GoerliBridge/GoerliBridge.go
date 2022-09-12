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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addressWMATIC\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"bridgeEmpty\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"bridgeOnOtherSideNeedsLiqudity\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"bridgedAlready\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"msgValueNot1003\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"msgValueZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"notExternalBridge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"notOwnerAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ownerBridgeUsersBeforeWithdraw\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"queueIsEmpty\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"queueNotEmpty\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"callMATIC\",\"outputs\":[{\"internalType\":\"contractERC20TokenContractMATIC\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dequeue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"first\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"last\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lockTokensForOptimism\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ownerRemoveBridgeLiqudity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"userToBridge\",\"type\":\"address\"}],\"name\":\"ownerUnlockGoerliETH\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"queue\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60a0604052600160035534801561001557600080fd5b50604051610d78380380610d7883398181016040528101906100379190610114565b3373ffffffffffffffffffffffffffffffffffffffff1660808173ffffffffffffffffffffffffffffffffffffffff1681525050806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050610141565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006100e1826100b6565b9050919050565b6100f1816100d6565b81146100fc57600080fd5b50565b60008151905061010e816100e8565b92915050565b60006020828403121561012a576101296100b1565b5b6000610138848285016100ff565b91505092915050565b608051610bf961017f600039600081816101d70152818161036c0152818161048a015281816106020152818161066b01526107820152610bf96000f3fe6080604052600436106100865760003560e01c80634e22d208116100595780634e22d20814610121578063957908d11461012b578063b4a99a4e14610142578063cc7c7a7d1461016d578063ddf0b0091461019857610086565b806308dd057a1461008b57806315b5cbdf146100a25780633df4ddf4146100cb57806347799da8146100f6575b600080fd5b34801561009757600080fd5b506100a06101d5565b005b3480156100ae57600080fd5b506100c960048036038101906100c491906108ce565b610488565b005b3480156100d757600080fd5b506100e06105b1565b6040516100ed9190610914565b60405180910390f35b34801561010257600080fd5b5061010b6105b7565b6040516101189190610914565b60405180910390f35b6101296105bd565b005b34801561013757600080fd5b50610140610669565b005b34801561014e57600080fd5b50610157610780565b604051610164919061093e565b60405180910390f35b34801561017957600080fd5b506101826107a4565b60405161018f91906109b8565b60405180910390f35b3480156101a457600080fd5b506101bf60048036038101906101ba91906109ff565b6107c8565b6040516101cc919061093e565b60405180910390f35b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461025a576040517f1ee8575e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b81526004016102b6919061093e565b602060405180830381865afa1580156102d3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102f79190610a41565b0361032e576040517f7a1f291700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb7f000000000000000000000000000000000000000000000000000000000000000060008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b81526004016103e4919061093e565b602060405180830381865afa158015610401573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104259190610a41565b6040518363ffffffff1660e01b8152600401610442929190610a6e565b6020604051808303816000875af1158015610461573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104859190610acf565b50565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461050d576040517f1ee8575e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb826103e86040518363ffffffff1660e01b815260040161056a929190610b37565b6020604051808303816000875af1158015610589573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105ad9190610acf565b5050565b60035481565b60025481565b6103eb34146105f8576040517fa43d860800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6106006107fb565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166108fc349081150290604051600060405180830381858888f19350505050158015610666573d6000803e3d6000fd5b50565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146106ee576040517f1ee8575e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600354600254101561072c576040517f5e61bb8700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60016000600354815260200190815260200160002060006101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690556001600360008282546107779190610b8f565b92505081905550565b7f000000000000000000000000000000000000000000000000000000000000000081565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60016020528060005260406000206000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60016002600082825461080e9190610b8f565b925050819055503360016000600254815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061089b82610870565b9050919050565b6108ab81610890565b81146108b657600080fd5b50565b6000813590506108c8816108a2565b92915050565b6000602082840312156108e4576108e361086b565b5b60006108f2848285016108b9565b91505092915050565b6000819050919050565b61090e816108fb565b82525050565b60006020820190506109296000830184610905565b92915050565b61093881610890565b82525050565b6000602082019050610953600083018461092f565b92915050565b6000819050919050565b600061097e61097961097484610870565b610959565b610870565b9050919050565b600061099082610963565b9050919050565b60006109a282610985565b9050919050565b6109b281610997565b82525050565b60006020820190506109cd60008301846109a9565b92915050565b6109dc816108fb565b81146109e757600080fd5b50565b6000813590506109f9816109d3565b92915050565b600060208284031215610a1557610a1461086b565b5b6000610a23848285016109ea565b91505092915050565b600081519050610a3b816109d3565b92915050565b600060208284031215610a5757610a5661086b565b5b6000610a6584828501610a2c565b91505092915050565b6000604082019050610a83600083018561092f565b610a906020830184610905565b9392505050565b60008115159050919050565b610aac81610a97565b8114610ab757600080fd5b50565b600081519050610ac981610aa3565b92915050565b600060208284031215610ae557610ae461086b565b5b6000610af384828501610aba565b91505092915050565b6000819050919050565b6000610b21610b1c610b1784610afc565b610959565b6108fb565b9050919050565b610b3181610b06565b82525050565b6000604082019050610b4c600083018561092f565b610b596020830184610b28565b9392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000610b9a826108fb565b9150610ba5836108fb565b9250828201905080821115610bbd57610bbc610b60565b5b9291505056fea264697066735822122085a3b48104e3170ad48ca50b6a9f792fa0b88a289b129aa0b3861ff778bfcba764736f6c63430008110033",
}

// GoerliBridgeABI is the input ABI used to generate the binding from.
// Deprecated: Use GoerliBridgeMetaData.ABI instead.
var GoerliBridgeABI = GoerliBridgeMetaData.ABI

// GoerliBridgeBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use GoerliBridgeMetaData.Bin instead.
var GoerliBridgeBin = GoerliBridgeMetaData.Bin

// DeployGoerliBridge deploys a new Ethereum contract, binding an instance of GoerliBridge to it.
func DeployGoerliBridge(auth *bind.TransactOpts, backend bind.ContractBackend, addressWMATIC common.Address) (common.Address, *types.Transaction, *GoerliBridge, error) {
	parsed, err := GoerliBridgeMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(GoerliBridgeBin), backend, addressWMATIC)
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

// CallMATIC is a free data retrieval call binding the contract method 0xcc7c7a7d.
//
// Solidity: function callMATIC() view returns(address)
func (_GoerliBridge *GoerliBridgeCaller) CallMATIC(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GoerliBridge.contract.Call(opts, &out, "callMATIC")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CallMATIC is a free data retrieval call binding the contract method 0xcc7c7a7d.
//
// Solidity: function callMATIC() view returns(address)
func (_GoerliBridge *GoerliBridgeSession) CallMATIC() (common.Address, error) {
	return _GoerliBridge.Contract.CallMATIC(&_GoerliBridge.CallOpts)
}

// CallMATIC is a free data retrieval call binding the contract method 0xcc7c7a7d.
//
// Solidity: function callMATIC() view returns(address)
func (_GoerliBridge *GoerliBridgeCallerSession) CallMATIC() (common.Address, error) {
	return _GoerliBridge.Contract.CallMATIC(&_GoerliBridge.CallOpts)
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
