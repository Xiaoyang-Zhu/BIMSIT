// This file is an automatically generated Go binding. Do not modify as any
// change will likely be lost upon the next re-generation!

package contracts

import (
	"strings"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
)

// TxIDABI is the input ABI used to generate the binding from.
const TxIDABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"txUPD\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"txRVK\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"verifier\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"rootID\",\"type\":\"address\"},{\"name\":\"rootPKf\",\"type\":\"bytes\"},{\"name\":\"rootPKo\",\"type\":\"bytes\"},{\"name\":\"sig\",\"type\":\"bytes\"},{\"name\":\"rootPointer\",\"type\":\"bytes\"}],\"name\":\"txReg\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"txLKP\",\"outputs\":[],\"payable\":false,\"type\":\"function\"}]"

// TxIDBin is the compiled bytecode used for deploying new contracts.
const TxIDBin = `0x6060604052341561000c57fe5b5b6103bc8061001c6000396000f3006060604052361561005c5763ffffffff60e060020a60003504166322938918811461005e57806328aec0ec1461005e5780632b7ac3f31461005e5780638da5cb5b146100945780639c9704ba146100c0578063c5476fbd1461005e575bfe5b341561006657fe5b61006e6101eb565b005b341561006657fe5b61006e6101eb565b005b341561006657fe5b61006e6101eb565b005b341561009c57fe5b6100a46101f4565b60408051600160a060020a039092168252519081900360200190f35b34156100c857fe5b60408051602060046024803582810135601f810185900485028601850190965285855261006e958335600160a060020a0316959394604494939290920191819084018382808284375050604080516020601f89358b0180359182018390048302840183019094528083529799988101979196509182019450925082915084018382808284375050604080516020601f89358b0180359182018390048302840183019094528083529799988101979196509182019450925082915084018382808284375050604080516020601f89358b0180359182018390048302840183019094528083529799988101979196509182019450925082915084018382808284375094965061020395505050505050565b005b341561006657fe5b61006e6101eb565b005b5b565b5b565b5b565b600054600160a060020a031681565b6000805473ffffffffffffffffffffffffffffffffffffffff191633600160a060020a03908116919091178255861681526001602081815260409092208651610254939190920191908701906102f0565b50600160a060020a03851660009081526001602090815260409091208451610284926002909201918601906102f0565b50600160a060020a038516600090815260016020908152604090912083516102b4926003909201918501906102f0565b50600160a060020a038516600090815260016020908152604090912082516102e4926004909201918401906102f0565b505b5050505050565b5b565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061033157805160ff191683800117855561035e565b8280016001018555821561035e579182015b8281111561035e578251825591602001919060010190610343565b5b5061036b92915061036f565b5090565b61038d91905b8082111561036b5760008155600101610375565b5090565b905600a165627a7a72305820dd2f629e9b24ee036f3c02bfca56afedcbee4f94d83b25d820d07c21d8478ed20029`

// DeployTxID deploys a new Ethereum contract, binding an instance of TxID to it.
func DeployTxID(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TxID, error) {
	parsed, err := abi.JSON(strings.NewReader(TxIDABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TxIDBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TxID{TxIDCaller: TxIDCaller{contract: contract}, TxIDTransactor: TxIDTransactor{contract: contract}}, nil
}

// TxID is an auto generated Go binding around an Ethereum contract.
type TxID struct {
	TxIDCaller     // Read-only binding to the contract
	TxIDTransactor // Write-only binding to the contract
}

// TxIDCaller is an auto generated read-only Go binding around an Ethereum contract.
type TxIDCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TxIDTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TxIDTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TxIDSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TxIDSession struct {
	Contract     *TxID             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TxIDCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TxIDCallerSession struct {
	Contract *TxIDCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TxIDTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TxIDTransactorSession struct {
	Contract     *TxIDTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TxIDRaw is an auto generated low-level Go binding around an Ethereum contract.
type TxIDRaw struct {
	Contract *TxID // Generic contract binding to access the raw methods on
}

// TxIDCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TxIDCallerRaw struct {
	Contract *TxIDCaller // Generic read-only contract binding to access the raw methods on
}

// TxIDTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TxIDTransactorRaw struct {
	Contract *TxIDTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTxID creates a new instance of TxID, bound to a specific deployed contract.
func NewTxID(address common.Address, backend bind.ContractBackend) (*TxID, error) {
	contract, err := bindTxID(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TxID{TxIDCaller: TxIDCaller{contract: contract}, TxIDTransactor: TxIDTransactor{contract: contract}}, nil
}

// NewTxIDCaller creates a new read-only instance of TxID, bound to a specific deployed contract.
func NewTxIDCaller(address common.Address, caller bind.ContractCaller) (*TxIDCaller, error) {
	contract, err := bindTxID(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &TxIDCaller{contract: contract}, nil
}

// NewTxIDTransactor creates a new write-only instance of TxID, bound to a specific deployed contract.
func NewTxIDTransactor(address common.Address, transactor bind.ContractTransactor) (*TxIDTransactor, error) {
	contract, err := bindTxID(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &TxIDTransactor{contract: contract}, nil
}

// bindTxID binds a generic wrapper to an already deployed contract.
func bindTxID(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TxIDABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TxID *TxIDRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TxID.Contract.TxIDCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TxID *TxIDRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TxID.Contract.TxIDTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TxID *TxIDRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TxID.Contract.TxIDTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TxID *TxIDCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TxID.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TxID *TxIDTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TxID.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TxID *TxIDTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TxID.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_TxID *TxIDCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _TxID.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_TxID *TxIDSession) Owner() (common.Address, error) {
	return _TxID.Contract.Owner(&_TxID.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_TxID *TxIDCallerSession) Owner() (common.Address, error) {
	return _TxID.Contract.Owner(&_TxID.CallOpts)
}

// TxLKP is a paid mutator transaction binding the contract method 0xc5476fbd.
//
// Solidity: function txLKP() returns()
func (_TxID *TxIDTransactor) TxLKP(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TxID.contract.Transact(opts, "txLKP")
}

// TxLKP is a paid mutator transaction binding the contract method 0xc5476fbd.
//
// Solidity: function txLKP() returns()
func (_TxID *TxIDSession) TxLKP() (*types.Transaction, error) {
	return _TxID.Contract.TxLKP(&_TxID.TransactOpts)
}

// TxLKP is a paid mutator transaction binding the contract method 0xc5476fbd.
//
// Solidity: function txLKP() returns()
func (_TxID *TxIDTransactorSession) TxLKP() (*types.Transaction, error) {
	return _TxID.Contract.TxLKP(&_TxID.TransactOpts)
}

// TxRVK is a paid mutator transaction binding the contract method 0x28aec0ec.
//
// Solidity: function txRVK() returns()
func (_TxID *TxIDTransactor) TxRVK(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TxID.contract.Transact(opts, "txRVK")
}

// TxRVK is a paid mutator transaction binding the contract method 0x28aec0ec.
//
// Solidity: function txRVK() returns()
func (_TxID *TxIDSession) TxRVK() (*types.Transaction, error) {
	return _TxID.Contract.TxRVK(&_TxID.TransactOpts)
}

// TxRVK is a paid mutator transaction binding the contract method 0x28aec0ec.
//
// Solidity: function txRVK() returns()
func (_TxID *TxIDTransactorSession) TxRVK() (*types.Transaction, error) {
	return _TxID.Contract.TxRVK(&_TxID.TransactOpts)
}

// TxReg is a paid mutator transaction binding the contract method 0x9c9704ba.
//
// Solidity: function txReg(rootID address, rootPKf bytes, rootPKo bytes, sig bytes, rootPointer bytes) returns()
func (_TxID *TxIDTransactor) TxReg(opts *bind.TransactOpts, rootID common.Address, rootPKf []byte, rootPKo []byte, sig []byte, rootPointer []byte) (*types.Transaction, error) {
	return _TxID.contract.Transact(opts, "txReg", rootID, rootPKf, rootPKo, sig, rootPointer)
}

// TxReg is a paid mutator transaction binding the contract method 0x9c9704ba.
//
// Solidity: function txReg(rootID address, rootPKf bytes, rootPKo bytes, sig bytes, rootPointer bytes) returns()
func (_TxID *TxIDSession) TxReg(rootID common.Address, rootPKf []byte, rootPKo []byte, sig []byte, rootPointer []byte) (*types.Transaction, error) {
	return _TxID.Contract.TxReg(&_TxID.TransactOpts, rootID, rootPKf, rootPKo, sig, rootPointer)
}

// TxReg is a paid mutator transaction binding the contract method 0x9c9704ba.
//
// Solidity: function txReg(rootID address, rootPKf bytes, rootPKo bytes, sig bytes, rootPointer bytes) returns()
func (_TxID *TxIDTransactorSession) TxReg(rootID common.Address, rootPKf []byte, rootPKo []byte, sig []byte, rootPointer []byte) (*types.Transaction, error) {
	return _TxID.Contract.TxReg(&_TxID.TransactOpts, rootID, rootPKf, rootPKo, sig, rootPointer)
}

// TxUPD is a paid mutator transaction binding the contract method 0x22938918.
//
// Solidity: function txUPD() returns()
func (_TxID *TxIDTransactor) TxUPD(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TxID.contract.Transact(opts, "txUPD")
}

// TxUPD is a paid mutator transaction binding the contract method 0x22938918.
//
// Solidity: function txUPD() returns()
func (_TxID *TxIDSession) TxUPD() (*types.Transaction, error) {
	return _TxID.Contract.TxUPD(&_TxID.TransactOpts)
}

// TxUPD is a paid mutator transaction binding the contract method 0x22938918.
//
// Solidity: function txUPD() returns()
func (_TxID *TxIDTransactorSession) TxUPD() (*types.Transaction, error) {
	return _TxID.Contract.TxUPD(&_TxID.TransactOpts)
}

// Verifier is a paid mutator transaction binding the contract method 0x2b7ac3f3.
//
// Solidity: function verifier() returns()
func (_TxID *TxIDTransactor) Verifier(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TxID.contract.Transact(opts, "verifier")
}

// Verifier is a paid mutator transaction binding the contract method 0x2b7ac3f3.
//
// Solidity: function verifier() returns()
func (_TxID *TxIDSession) Verifier() (*types.Transaction, error) {
	return _TxID.Contract.Verifier(&_TxID.TransactOpts)
}

// Verifier is a paid mutator transaction binding the contract method 0x2b7ac3f3.
//
// Solidity: function verifier() returns()
func (_TxID *TxIDTransactorSession) Verifier() (*types.Transaction, error) {
	return _TxID.Contract.Verifier(&_TxID.TransactOpts)
}
