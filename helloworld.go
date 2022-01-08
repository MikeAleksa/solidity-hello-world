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

// HelloworldMetaData contains all meta data concerning the Helloworld contract.
var HelloworldMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"say\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"954ab4b2": "say()",
	},
	Bin: "0x608060405234801561001057600080fd5b5060ee8061001f6000396000f3fe6080604052348015600f57600080fd5b506004361060285760003560e01c8063954ab4b214602d575b600080fd5b604080518082018252601081526f1a195b1b1bc8195d1a195c9ddbdc9b1960821b60208201529051605d91906066565b60405180910390f35b600060208083528351808285015260005b818110156091578581018301518582016040015282016077565b8181111560a2576000604083870101525b50601f01601f191692909201604001939250505056fea26469706673582212201c2228ff2aef15d44a5e8f24534370ea3734993045a365aae16ccaeed398717964736f6c634300080b0033",
}

// HelloworldABI is the input ABI used to generate the binding from.
// Deprecated: Use HelloworldMetaData.ABI instead.
var HelloworldABI = HelloworldMetaData.ABI

// Deprecated: Use HelloworldMetaData.Sigs instead.
// HelloworldFuncSigs maps the 4-byte function signature to its string representation.
var HelloworldFuncSigs = HelloworldMetaData.Sigs

// HelloworldBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use HelloworldMetaData.Bin instead.
var HelloworldBin = HelloworldMetaData.Bin

// DeployHelloworld deploys a new Ethereum contract, binding an instance of Helloworld to it.
func DeployHelloworld(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Helloworld, error) {
	parsed, err := HelloworldMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(HelloworldBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Helloworld{HelloworldCaller: HelloworldCaller{contract: contract}, HelloworldTransactor: HelloworldTransactor{contract: contract}, HelloworldFilterer: HelloworldFilterer{contract: contract}}, nil
}

// Helloworld is an auto generated Go binding around an Ethereum contract.
type Helloworld struct {
	HelloworldCaller     // Read-only binding to the contract
	HelloworldTransactor // Write-only binding to the contract
	HelloworldFilterer   // Log filterer for contract events
}

// HelloworldCaller is an auto generated read-only Go binding around an Ethereum contract.
type HelloworldCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HelloworldTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HelloworldTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HelloworldFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HelloworldFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HelloworldSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HelloworldSession struct {
	Contract     *Helloworld       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HelloworldCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HelloworldCallerSession struct {
	Contract *HelloworldCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// HelloworldTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HelloworldTransactorSession struct {
	Contract     *HelloworldTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// HelloworldRaw is an auto generated low-level Go binding around an Ethereum contract.
type HelloworldRaw struct {
	Contract *Helloworld // Generic contract binding to access the raw methods on
}

// HelloworldCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HelloworldCallerRaw struct {
	Contract *HelloworldCaller // Generic read-only contract binding to access the raw methods on
}

// HelloworldTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HelloworldTransactorRaw struct {
	Contract *HelloworldTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHelloworld creates a new instance of Helloworld, bound to a specific deployed contract.
func NewHelloworld(address common.Address, backend bind.ContractBackend) (*Helloworld, error) {
	contract, err := bindHelloworld(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Helloworld{HelloworldCaller: HelloworldCaller{contract: contract}, HelloworldTransactor: HelloworldTransactor{contract: contract}, HelloworldFilterer: HelloworldFilterer{contract: contract}}, nil
}

// NewHelloworldCaller creates a new read-only instance of Helloworld, bound to a specific deployed contract.
func NewHelloworldCaller(address common.Address, caller bind.ContractCaller) (*HelloworldCaller, error) {
	contract, err := bindHelloworld(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HelloworldCaller{contract: contract}, nil
}

// NewHelloworldTransactor creates a new write-only instance of Helloworld, bound to a specific deployed contract.
func NewHelloworldTransactor(address common.Address, transactor bind.ContractTransactor) (*HelloworldTransactor, error) {
	contract, err := bindHelloworld(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HelloworldTransactor{contract: contract}, nil
}

// NewHelloworldFilterer creates a new log filterer instance of Helloworld, bound to a specific deployed contract.
func NewHelloworldFilterer(address common.Address, filterer bind.ContractFilterer) (*HelloworldFilterer, error) {
	contract, err := bindHelloworld(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HelloworldFilterer{contract: contract}, nil
}

// bindHelloworld binds a generic wrapper to an already deployed contract.
func bindHelloworld(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(HelloworldABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Helloworld *HelloworldRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Helloworld.Contract.HelloworldCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Helloworld *HelloworldRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Helloworld.Contract.HelloworldTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Helloworld *HelloworldRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Helloworld.Contract.HelloworldTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Helloworld *HelloworldCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Helloworld.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Helloworld *HelloworldTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Helloworld.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Helloworld *HelloworldTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Helloworld.Contract.contract.Transact(opts, method, params...)
}

// Say is a free data retrieval call binding the contract method 0x954ab4b2.
//
// Solidity: function say() pure returns(string)
func (_Helloworld *HelloworldCaller) Say(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Helloworld.contract.Call(opts, &out, "say")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Say is a free data retrieval call binding the contract method 0x954ab4b2.
//
// Solidity: function say() pure returns(string)
func (_Helloworld *HelloworldSession) Say() (string, error) {
	return _Helloworld.Contract.Say(&_Helloworld.CallOpts)
}

// Say is a free data retrieval call binding the contract method 0x954ab4b2.
//
// Solidity: function say() pure returns(string)
func (_Helloworld *HelloworldCallerSession) Say() (string, error) {
	return _Helloworld.Contract.Say(&_Helloworld.CallOpts)
}
