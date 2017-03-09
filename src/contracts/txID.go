// This file is an automatically generated Go binding. Do not modify as any
// change will likely be lost upon the next re-generation!

package contracts

import (
	"strings"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/core/types"
)

// TxRegABI is the input ABI used to generate the binding from.
const TxRegABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"txUPD\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"txRVK\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"verifier\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"txLKP\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"inputs\":[{\"name\":\"rootID\",\"type\":\"address\"},{\"name\":\"rootPKf\",\"type\":\"bytes\"},{\"name\":\"rootPKo\",\"type\":\"bytes\"},{\"name\":\"sig\",\"type\":\"bytes\"},{\"name\":\"rootPointer\",\"type\":\"bytes\"}],\"payable\":false,\"type\":\"constructor\"}]"

// TxRegBin is the compiled bytecode used for deploying new contracts.
const TxRegBin = `0x6060604052341561000c57fe5b6040516102e33803806102e38339810160409081528151602083015191830151606084015160808501519294938401939182019290820191015b60008054600160a060020a03191633600160a060020a0390811691909117825586168152600160208181526040909220865161008a93919092019190870190610126565b50600160a060020a038516600090815260016020908152604090912084516100ba92600290920191860190610126565b50600160a060020a038516600090815260016020908152604090912083516100ea92600390920191850190610126565b50600160a060020a0385166000908152600160209081526040909120825161011a92600490920191840190610126565b505b50505050506101c6565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061016757805160ff1916838001178555610194565b82800160010185558215610194579182015b82811115610194578251825591602001919060010190610179565b5b506101a19291506101a5565b5090565b6101c391905b808211156101a157600081556001016101ab565b5090565b90565b61010e806101d56000396000f300606060405263ffffffff60e060020a600035041663229389188114604857806328aec0ec1460485780632b7ac3f31460485780638da5cb5b146075578063c5476fbd146048575bfe5b3415604f57fe5b605560ba565b005b3415604f57fe5b605560ba565b005b3415604f57fe5b605560ba565b005b3415607c57fe5b608260c3565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b3415604f57fe5b605560ba565b005b5b565b5b565b5b565b60005473ffffffffffffffffffffffffffffffffffffffff1681565b5b5600a165627a7a72305820bbdd23d7b62d966ace21fb828a22a88fe0f314b80a183a7917b974f5055a61f50029`

// DeployTxReg deploys a new Ethereum contract, binding an instance of TxReg to it.
func DeployTxReg(auth *bind.TransactOpts, backend bind.ContractBackend, rootID common.Address, rootPKf []byte, rootPKo []byte, sig []byte, rootPointer []byte) (common.Address, *types.Transaction, *TxReg, error) {
	parsed, err := abi.JSON(strings.NewReader(TxRegABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TxRegBin), backend, rootID, rootPKf, rootPKo, sig, rootPointer)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TxReg{TxRegCaller: TxRegCaller{contract: contract}, TxRegTransactor: TxRegTransactor{contract: contract}}, nil
}

// TxReg is an auto generated Go binding around an Ethereum contract.
type TxReg struct {
	TxRegCaller     // Read-only binding to the contract
	TxRegTransactor // Write-only binding to the contract
}

// TxRegCaller is an auto generated read-only Go binding around an Ethereum contract.
type TxRegCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TxRegTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TxRegTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TxRegSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TxRegSession struct {
	Contract     *TxReg            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TxRegCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TxRegCallerSession struct {
	Contract *TxRegCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TxRegTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TxRegTransactorSession struct {
	Contract     *TxRegTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TxRegRaw is an auto generated low-level Go binding around an Ethereum contract.
type TxRegRaw struct {
	Contract *TxReg // Generic contract binding to access the raw methods on
}

// TxRegCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TxRegCallerRaw struct {
	Contract *TxRegCaller // Generic read-only contract binding to access the raw methods on
}

// TxRegTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TxRegTransactorRaw struct {
	Contract *TxRegTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTxReg creates a new instance of TxReg, bound to a specific deployed contract.
func NewTxReg(address common.Address, backend bind.ContractBackend) (*TxReg, error) {
	contract, err := bindTxReg(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TxReg{TxRegCaller: TxRegCaller{contract: contract}, TxRegTransactor: TxRegTransactor{contract: contract}}, nil
}

// NewTxRegCaller creates a new read-only instance of TxReg, bound to a specific deployed contract.
func NewTxRegCaller(address common.Address, caller bind.ContractCaller) (*TxRegCaller, error) {
	contract, err := bindTxReg(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &TxRegCaller{contract: contract}, nil
}

// NewTxRegTransactor creates a new write-only instance of TxReg, bound to a specific deployed contract.
func NewTxRegTransactor(address common.Address, transactor bind.ContractTransactor) (*TxRegTransactor, error) {
	contract, err := bindTxReg(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &TxRegTransactor{contract: contract}, nil
}

// bindTxReg binds a generic wrapper to an already deployed contract.
func bindTxReg(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TxRegABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TxReg *TxRegRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TxReg.Contract.TxRegCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TxReg *TxRegRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TxReg.Contract.TxRegTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TxReg *TxRegRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TxReg.Contract.TxRegTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TxReg *TxRegCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TxReg.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TxReg *TxRegTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TxReg.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TxReg *TxRegTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TxReg.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_TxReg *TxRegCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _TxReg.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_TxReg *TxRegSession) Owner() (common.Address, error) {
	return _TxReg.Contract.Owner(&_TxReg.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_TxReg *TxRegCallerSession) Owner() (common.Address, error) {
	return _TxReg.Contract.Owner(&_TxReg.CallOpts)
}

// TxLKP is a paid mutator transaction binding the contract method 0xc5476fbd.
//
// Solidity: function txLKP() returns()
func (_TxReg *TxRegTransactor) TxLKP(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TxReg.contract.Transact(opts, "txLKP")
}

// TxLKP is a paid mutator transaction binding the contract method 0xc5476fbd.
//
// Solidity: function txLKP() returns()
func (_TxReg *TxRegSession) TxLKP() (*types.Transaction, error) {
	return _TxReg.Contract.TxLKP(&_TxReg.TransactOpts)
}

// TxLKP is a paid mutator transaction binding the contract method 0xc5476fbd.
//
// Solidity: function txLKP() returns()
func (_TxReg *TxRegTransactorSession) TxLKP() (*types.Transaction, error) {
	return _TxReg.Contract.TxLKP(&_TxReg.TransactOpts)
}

// TxRVK is a paid mutator transaction binding the contract method 0x28aec0ec.
//
// Solidity: function txRVK() returns()
func (_TxReg *TxRegTransactor) TxRVK(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TxReg.contract.Transact(opts, "txRVK")
}

// TxRVK is a paid mutator transaction binding the contract method 0x28aec0ec.
//
// Solidity: function txRVK() returns()
func (_TxReg *TxRegSession) TxRVK() (*types.Transaction, error) {
	return _TxReg.Contract.TxRVK(&_TxReg.TransactOpts)
}

// TxRVK is a paid mutator transaction binding the contract method 0x28aec0ec.
//
// Solidity: function txRVK() returns()
func (_TxReg *TxRegTransactorSession) TxRVK() (*types.Transaction, error) {
	return _TxReg.Contract.TxRVK(&_TxReg.TransactOpts)
}

// TxUPD is a paid mutator transaction binding the contract method 0x22938918.
//
// Solidity: function txUPD() returns()
func (_TxReg *TxRegTransactor) TxUPD(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TxReg.contract.Transact(opts, "txUPD")
}

// TxUPD is a paid mutator transaction binding the contract method 0x22938918.
//
// Solidity: function txUPD() returns()
func (_TxReg *TxRegSession) TxUPD() (*types.Transaction, error) {
	return _TxReg.Contract.TxUPD(&_TxReg.TransactOpts)
}

// TxUPD is a paid mutator transaction binding the contract method 0x22938918.
//
// Solidity: function txUPD() returns()
func (_TxReg *TxRegTransactorSession) TxUPD() (*types.Transaction, error) {
	return _TxReg.Contract.TxUPD(&_TxReg.TransactOpts)
}

// Verifier is a paid mutator transaction binding the contract method 0x2b7ac3f3.
//
// Solidity: function verifier() returns()
func (_TxReg *TxRegTransactor) Verifier(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TxReg.contract.Transact(opts, "verifier")
}

// Verifier is a paid mutator transaction binding the contract method 0x2b7ac3f3.
//
// Solidity: function verifier() returns()
func (_TxReg *TxRegSession) Verifier() (*types.Transaction, error) {
	return _TxReg.Contract.Verifier(&_TxReg.TransactOpts)
}

// Verifier is a paid mutator transaction binding the contract method 0x2b7ac3f3.
//
// Solidity: function verifier() returns()
func (_TxReg *TxRegTransactorSession) Verifier() (*types.Transaction, error) {
	return _TxReg.Contract.Verifier(&_TxReg.TransactOpts)
}
