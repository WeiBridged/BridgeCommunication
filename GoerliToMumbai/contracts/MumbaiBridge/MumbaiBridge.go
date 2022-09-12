// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package mumbaiBridge

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

// MumbaiBridgeMetaData contains all meta data concerning the MumbaiBridge contract.
var MumbaiBridgeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"bridgeEmpty\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"bridgeOnOtherSideNeedsLiqudity\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"msgValueNot1003\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"msgValueZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"notExternalBridge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"notOwnerAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"queueIsEmpty\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"queueNotEmpty\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"dequeue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lockTokensForGoerli\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ownerRemoveBridgeLiqudity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"userToBridge\",\"type\":\"address\"}],\"name\":\"ownerUnlockOptimismETH\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addressWETH\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"callWETH\",\"outputs\":[{\"internalType\":\"contractERC20TokenContractWETH\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"first\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"last\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"queue\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60a0604052600160025534801561001557600080fd5b50604051610d7f380380610d7f83398181016040528101906100379190610115565b3373ffffffffffffffffffffffffffffffffffffffff1660808173ffffffffffffffffffffffffffffffffffffffff168152505080600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050610142565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006100e2826100b7565b9050919050565b6100f2816100d7565b81146100fd57600080fd5b50565b60008151905061010f816100e9565b92915050565b60006020828403121561012b5761012a6100b2565b5b600061013984828501610100565b91505092915050565b608051610bff610180600039600081816101d7015281816103020152818161049a015281816105c601528181610702015261079c0152610bff6000f3fe6080604052600436106100865760003560e01c8063957908d111610059578063957908d1146101215780639ee813d214610138578063b4a99a4e14610163578063ddf0b0091461018e578063f67744c4146101cb57610086565b806304c5366e1461008b57806308dd057a146100b45780633df4ddf4146100cb57806347799da8146100f6575b600080fd5b34801561009757600080fd5b506100b260048036038101906100ad91906108d4565b6101d5565b005b3480156100c057600080fd5b506100c9610300565b005b3480156100d757600080fd5b506100e06105b8565b6040516100ed919061091a565b60405180910390f35b34801561010257600080fd5b5061010b6105be565b604051610118919061091a565b60405180910390f35b34801561012d57600080fd5b506101366105c4565b005b34801561014457600080fd5b5061014d6106da565b60405161015a9190610994565b60405180910390f35b34801561016f57600080fd5b50610178610700565b60405161018591906109be565b60405180910390f35b34801561019a57600080fd5b506101b560048036038101906101b09190610a05565b610724565b6040516101c291906109be565b60405180910390f35b6101d3610757565b005b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461025a576040517f1ee8575e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb826103e86040518363ffffffff1660e01b81526004016102b9929190610a6d565b6020604051808303816000875af11580156102d8573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102fc9190610ace565b5050565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610385576040517f1ee8575e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b81526004016103e291906109be565b602060405180830381865afa1580156103ff573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104239190610b10565b0361045a576040517f7a1f291700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb7f0000000000000000000000000000000000000000000000000000000000000000600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b815260040161051491906109be565b602060405180830381865afa158015610531573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105559190610b10565b6040518363ffffffff1660e01b8152600401610572929190610b3d565b6020604051808303816000875af1158015610591573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105b59190610ace565b50565b60025481565b60015481565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610649576040517f1ee8575e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002546001541015610687576040517f5e61bb8700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600080600254815260200190815260200160002060006101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690556001600260008282546106d19190610b95565b92505081905550565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b7f000000000000000000000000000000000000000000000000000000000000000081565b60006020528060005260406000206000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6103eb3414610792576040517fa43d860800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61079a610803565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166108fc349081150290604051600060405180830381858888f19350505050158015610800573d6000803e3d6000fd5b50565b60018060008282546108159190610b95565b9250508190555033600080600154815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006108a182610876565b9050919050565b6108b181610896565b81146108bc57600080fd5b50565b6000813590506108ce816108a8565b92915050565b6000602082840312156108ea576108e9610871565b5b60006108f8848285016108bf565b91505092915050565b6000819050919050565b61091481610901565b82525050565b600060208201905061092f600083018461090b565b92915050565b6000819050919050565b600061095a61095561095084610876565b610935565b610876565b9050919050565b600061096c8261093f565b9050919050565b600061097e82610961565b9050919050565b61098e81610973565b82525050565b60006020820190506109a96000830184610985565b92915050565b6109b881610896565b82525050565b60006020820190506109d360008301846109af565b92915050565b6109e281610901565b81146109ed57600080fd5b50565b6000813590506109ff816109d9565b92915050565b600060208284031215610a1b57610a1a610871565b5b6000610a29848285016109f0565b91505092915050565b6000819050919050565b6000610a57610a52610a4d84610a32565b610935565b610901565b9050919050565b610a6781610a3c565b82525050565b6000604082019050610a8260008301856109af565b610a8f6020830184610a5e565b9392505050565b60008115159050919050565b610aab81610a96565b8114610ab657600080fd5b50565b600081519050610ac881610aa2565b92915050565b600060208284031215610ae457610ae3610871565b5b6000610af284828501610ab9565b91505092915050565b600081519050610b0a816109d9565b92915050565b600060208284031215610b2657610b25610871565b5b6000610b3484828501610afb565b91505092915050565b6000604082019050610b5260008301856109af565b610b5f602083018461090b565b9392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000610ba082610901565b9150610bab83610901565b9250828201905080821115610bc357610bc2610b66565b5b9291505056fea26469706673582212200436c5c57b3a0b56df6417595021fa721666f4d0ea7816e4e4daa10ac6656ea664736f6c63430008110033",
}

// MumbaiBridgeABI is the input ABI used to generate the binding from.
// Deprecated: Use MumbaiBridgeMetaData.ABI instead.
var MumbaiBridgeABI = MumbaiBridgeMetaData.ABI

// MumbaiBridgeBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MumbaiBridgeMetaData.Bin instead.
var MumbaiBridgeBin = MumbaiBridgeMetaData.Bin

// DeployMumbaiBridge deploys a new Ethereum contract, binding an instance of MumbaiBridge to it.
func DeployMumbaiBridge(auth *bind.TransactOpts, backend bind.ContractBackend, addressWETH common.Address) (common.Address, *types.Transaction, *MumbaiBridge, error) {
	parsed, err := MumbaiBridgeMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MumbaiBridgeBin), backend, addressWETH)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MumbaiBridge{MumbaiBridgeCaller: MumbaiBridgeCaller{contract: contract}, MumbaiBridgeTransactor: MumbaiBridgeTransactor{contract: contract}, MumbaiBridgeFilterer: MumbaiBridgeFilterer{contract: contract}}, nil
}

// MumbaiBridge is an auto generated Go binding around an Ethereum contract.
type MumbaiBridge struct {
	MumbaiBridgeCaller     // Read-only binding to the contract
	MumbaiBridgeTransactor // Write-only binding to the contract
	MumbaiBridgeFilterer   // Log filterer for contract events
}

// MumbaiBridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type MumbaiBridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MumbaiBridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MumbaiBridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MumbaiBridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MumbaiBridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MumbaiBridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MumbaiBridgeSession struct {
	Contract     *MumbaiBridge     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MumbaiBridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MumbaiBridgeCallerSession struct {
	Contract *MumbaiBridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// MumbaiBridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MumbaiBridgeTransactorSession struct {
	Contract     *MumbaiBridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// MumbaiBridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type MumbaiBridgeRaw struct {
	Contract *MumbaiBridge // Generic contract binding to access the raw methods on
}

// MumbaiBridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MumbaiBridgeCallerRaw struct {
	Contract *MumbaiBridgeCaller // Generic read-only contract binding to access the raw methods on
}

// MumbaiBridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MumbaiBridgeTransactorRaw struct {
	Contract *MumbaiBridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMumbaiBridge creates a new instance of MumbaiBridge, bound to a specific deployed contract.
func NewMumbaiBridge(address common.Address, backend bind.ContractBackend) (*MumbaiBridge, error) {
	contract, err := bindMumbaiBridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MumbaiBridge{MumbaiBridgeCaller: MumbaiBridgeCaller{contract: contract}, MumbaiBridgeTransactor: MumbaiBridgeTransactor{contract: contract}, MumbaiBridgeFilterer: MumbaiBridgeFilterer{contract: contract}}, nil
}

// NewMumbaiBridgeCaller creates a new read-only instance of MumbaiBridge, bound to a specific deployed contract.
func NewMumbaiBridgeCaller(address common.Address, caller bind.ContractCaller) (*MumbaiBridgeCaller, error) {
	contract, err := bindMumbaiBridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MumbaiBridgeCaller{contract: contract}, nil
}

// NewMumbaiBridgeTransactor creates a new write-only instance of MumbaiBridge, bound to a specific deployed contract.
func NewMumbaiBridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*MumbaiBridgeTransactor, error) {
	contract, err := bindMumbaiBridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MumbaiBridgeTransactor{contract: contract}, nil
}

// NewMumbaiBridgeFilterer creates a new log filterer instance of MumbaiBridge, bound to a specific deployed contract.
func NewMumbaiBridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*MumbaiBridgeFilterer, error) {
	contract, err := bindMumbaiBridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MumbaiBridgeFilterer{contract: contract}, nil
}

// bindMumbaiBridge binds a generic wrapper to an already deployed contract.
func bindMumbaiBridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MumbaiBridgeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MumbaiBridge *MumbaiBridgeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MumbaiBridge.Contract.MumbaiBridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MumbaiBridge *MumbaiBridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MumbaiBridge.Contract.MumbaiBridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MumbaiBridge *MumbaiBridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MumbaiBridge.Contract.MumbaiBridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MumbaiBridge *MumbaiBridgeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MumbaiBridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MumbaiBridge *MumbaiBridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MumbaiBridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MumbaiBridge *MumbaiBridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MumbaiBridge.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0xb4a99a4e.
//
// Solidity: function Owner() view returns(address)
func (_MumbaiBridge *MumbaiBridgeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MumbaiBridge.contract.Call(opts, &out, "Owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0xb4a99a4e.
//
// Solidity: function Owner() view returns(address)
func (_MumbaiBridge *MumbaiBridgeSession) Owner() (common.Address, error) {
	return _MumbaiBridge.Contract.Owner(&_MumbaiBridge.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0xb4a99a4e.
//
// Solidity: function Owner() view returns(address)
func (_MumbaiBridge *MumbaiBridgeCallerSession) Owner() (common.Address, error) {
	return _MumbaiBridge.Contract.Owner(&_MumbaiBridge.CallOpts)
}

// CallWETH is a free data retrieval call binding the contract method 0x9ee813d2.
//
// Solidity: function callWETH() view returns(address)
func (_MumbaiBridge *MumbaiBridgeCaller) CallWETH(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MumbaiBridge.contract.Call(opts, &out, "callWETH")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CallWETH is a free data retrieval call binding the contract method 0x9ee813d2.
//
// Solidity: function callWETH() view returns(address)
func (_MumbaiBridge *MumbaiBridgeSession) CallWETH() (common.Address, error) {
	return _MumbaiBridge.Contract.CallWETH(&_MumbaiBridge.CallOpts)
}

// CallWETH is a free data retrieval call binding the contract method 0x9ee813d2.
//
// Solidity: function callWETH() view returns(address)
func (_MumbaiBridge *MumbaiBridgeCallerSession) CallWETH() (common.Address, error) {
	return _MumbaiBridge.Contract.CallWETH(&_MumbaiBridge.CallOpts)
}

// First is a free data retrieval call binding the contract method 0x3df4ddf4.
//
// Solidity: function first() view returns(uint256)
func (_MumbaiBridge *MumbaiBridgeCaller) First(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MumbaiBridge.contract.Call(opts, &out, "first")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// First is a free data retrieval call binding the contract method 0x3df4ddf4.
//
// Solidity: function first() view returns(uint256)
func (_MumbaiBridge *MumbaiBridgeSession) First() (*big.Int, error) {
	return _MumbaiBridge.Contract.First(&_MumbaiBridge.CallOpts)
}

// First is a free data retrieval call binding the contract method 0x3df4ddf4.
//
// Solidity: function first() view returns(uint256)
func (_MumbaiBridge *MumbaiBridgeCallerSession) First() (*big.Int, error) {
	return _MumbaiBridge.Contract.First(&_MumbaiBridge.CallOpts)
}

// Last is a free data retrieval call binding the contract method 0x47799da8.
//
// Solidity: function last() view returns(uint256)
func (_MumbaiBridge *MumbaiBridgeCaller) Last(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MumbaiBridge.contract.Call(opts, &out, "last")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Last is a free data retrieval call binding the contract method 0x47799da8.
//
// Solidity: function last() view returns(uint256)
func (_MumbaiBridge *MumbaiBridgeSession) Last() (*big.Int, error) {
	return _MumbaiBridge.Contract.Last(&_MumbaiBridge.CallOpts)
}

// Last is a free data retrieval call binding the contract method 0x47799da8.
//
// Solidity: function last() view returns(uint256)
func (_MumbaiBridge *MumbaiBridgeCallerSession) Last() (*big.Int, error) {
	return _MumbaiBridge.Contract.Last(&_MumbaiBridge.CallOpts)
}

// Queue is a free data retrieval call binding the contract method 0xddf0b009.
//
// Solidity: function queue(uint256 ) view returns(address)
func (_MumbaiBridge *MumbaiBridgeCaller) Queue(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _MumbaiBridge.contract.Call(opts, &out, "queue", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Queue is a free data retrieval call binding the contract method 0xddf0b009.
//
// Solidity: function queue(uint256 ) view returns(address)
func (_MumbaiBridge *MumbaiBridgeSession) Queue(arg0 *big.Int) (common.Address, error) {
	return _MumbaiBridge.Contract.Queue(&_MumbaiBridge.CallOpts, arg0)
}

// Queue is a free data retrieval call binding the contract method 0xddf0b009.
//
// Solidity: function queue(uint256 ) view returns(address)
func (_MumbaiBridge *MumbaiBridgeCallerSession) Queue(arg0 *big.Int) (common.Address, error) {
	return _MumbaiBridge.Contract.Queue(&_MumbaiBridge.CallOpts, arg0)
}

// Dequeue is a paid mutator transaction binding the contract method 0x957908d1.
//
// Solidity: function dequeue() returns()
func (_MumbaiBridge *MumbaiBridgeTransactor) Dequeue(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MumbaiBridge.contract.Transact(opts, "dequeue")
}

// Dequeue is a paid mutator transaction binding the contract method 0x957908d1.
//
// Solidity: function dequeue() returns()
func (_MumbaiBridge *MumbaiBridgeSession) Dequeue() (*types.Transaction, error) {
	return _MumbaiBridge.Contract.Dequeue(&_MumbaiBridge.TransactOpts)
}

// Dequeue is a paid mutator transaction binding the contract method 0x957908d1.
//
// Solidity: function dequeue() returns()
func (_MumbaiBridge *MumbaiBridgeTransactorSession) Dequeue() (*types.Transaction, error) {
	return _MumbaiBridge.Contract.Dequeue(&_MumbaiBridge.TransactOpts)
}

// LockTokensForGoerli is a paid mutator transaction binding the contract method 0xf67744c4.
//
// Solidity: function lockTokensForGoerli() payable returns()
func (_MumbaiBridge *MumbaiBridgeTransactor) LockTokensForGoerli(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MumbaiBridge.contract.Transact(opts, "lockTokensForGoerli")
}

// LockTokensForGoerli is a paid mutator transaction binding the contract method 0xf67744c4.
//
// Solidity: function lockTokensForGoerli() payable returns()
func (_MumbaiBridge *MumbaiBridgeSession) LockTokensForGoerli() (*types.Transaction, error) {
	return _MumbaiBridge.Contract.LockTokensForGoerli(&_MumbaiBridge.TransactOpts)
}

// LockTokensForGoerli is a paid mutator transaction binding the contract method 0xf67744c4.
//
// Solidity: function lockTokensForGoerli() payable returns()
func (_MumbaiBridge *MumbaiBridgeTransactorSession) LockTokensForGoerli() (*types.Transaction, error) {
	return _MumbaiBridge.Contract.LockTokensForGoerli(&_MumbaiBridge.TransactOpts)
}

// OwnerRemoveBridgeLiqudity is a paid mutator transaction binding the contract method 0x08dd057a.
//
// Solidity: function ownerRemoveBridgeLiqudity() returns()
func (_MumbaiBridge *MumbaiBridgeTransactor) OwnerRemoveBridgeLiqudity(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MumbaiBridge.contract.Transact(opts, "ownerRemoveBridgeLiqudity")
}

// OwnerRemoveBridgeLiqudity is a paid mutator transaction binding the contract method 0x08dd057a.
//
// Solidity: function ownerRemoveBridgeLiqudity() returns()
func (_MumbaiBridge *MumbaiBridgeSession) OwnerRemoveBridgeLiqudity() (*types.Transaction, error) {
	return _MumbaiBridge.Contract.OwnerRemoveBridgeLiqudity(&_MumbaiBridge.TransactOpts)
}

// OwnerRemoveBridgeLiqudity is a paid mutator transaction binding the contract method 0x08dd057a.
//
// Solidity: function ownerRemoveBridgeLiqudity() returns()
func (_MumbaiBridge *MumbaiBridgeTransactorSession) OwnerRemoveBridgeLiqudity() (*types.Transaction, error) {
	return _MumbaiBridge.Contract.OwnerRemoveBridgeLiqudity(&_MumbaiBridge.TransactOpts)
}

// OwnerUnlockOptimismETH is a paid mutator transaction binding the contract method 0x04c5366e.
//
// Solidity: function ownerUnlockOptimismETH(address userToBridge) returns()
func (_MumbaiBridge *MumbaiBridgeTransactor) OwnerUnlockOptimismETH(opts *bind.TransactOpts, userToBridge common.Address) (*types.Transaction, error) {
	return _MumbaiBridge.contract.Transact(opts, "ownerUnlockOptimismETH", userToBridge)
}

// OwnerUnlockOptimismETH is a paid mutator transaction binding the contract method 0x04c5366e.
//
// Solidity: function ownerUnlockOptimismETH(address userToBridge) returns()
func (_MumbaiBridge *MumbaiBridgeSession) OwnerUnlockOptimismETH(userToBridge common.Address) (*types.Transaction, error) {
	return _MumbaiBridge.Contract.OwnerUnlockOptimismETH(&_MumbaiBridge.TransactOpts, userToBridge)
}

// OwnerUnlockOptimismETH is a paid mutator transaction binding the contract method 0x04c5366e.
//
// Solidity: function ownerUnlockOptimismETH(address userToBridge) returns()
func (_MumbaiBridge *MumbaiBridgeTransactorSession) OwnerUnlockOptimismETH(userToBridge common.Address) (*types.Transaction, error) {
	return _MumbaiBridge.Contract.OwnerUnlockOptimismETH(&_MumbaiBridge.TransactOpts, userToBridge)
}
