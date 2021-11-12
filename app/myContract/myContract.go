// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package chianService

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

// MyContractMetaData contains all meta data concerning the MyContract contract.
var MyContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"id\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"join\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"id\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"_value\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"_type\",\"type\":\"uint8\"}],\"name\":\"paylog\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"SYS_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TUSD_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"addrToId\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_type\",\"type\":\"uint8\"}],\"name\":\"get\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"get_balance_info\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"SYS_balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"TUSD_balance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"get_player_count\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"idToAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"_value\",\"type\":\"uint128\"},{\"internalType\":\"uint8\",\"name\":\"_type\",\"type\":\"uint8\"}],\"name\":\"pay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"register\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"opAddr\",\"type\":\"address\"}],\"name\":\"set_op_addr\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"val\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"_token_type\",\"type\":\"uint8\"}],\"name\":\"withdrawAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// MyContractABI is the input ABI used to generate the binding from.
// Deprecated: Use MyContractMetaData.ABI instead.
var MyContractABI = MyContractMetaData.ABI

// MyContract is an auto generated Go binding around an Ethereum contract.
type MyContract struct {
	MyContractCaller     // Read-only binding to the contract
	MyContractTransactor // Write-only binding to the contract
	MyContractFilterer   // Log filterer for contract events
}

// MyContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type MyContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MyContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MyContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MyContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MyContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MyContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MyContractSession struct {
	Contract     *MyContract       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MyContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MyContractCallerSession struct {
	Contract *MyContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// MyContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MyContractTransactorSession struct {
	Contract     *MyContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// MyContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type MyContractRaw struct {
	Contract *MyContract // Generic contract binding to access the raw methods on
}

// MyContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MyContractCallerRaw struct {
	Contract *MyContractCaller // Generic read-only contract binding to access the raw methods on
}

// MyContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MyContractTransactorRaw struct {
	Contract *MyContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMyContract creates a new instance of MyContract, bound to a specific deployed contract.
func NewMyContract(address common.Address, backend bind.ContractBackend) (*MyContract, error) {
	contract, err := bindMyContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MyContract{MyContractCaller: MyContractCaller{contract: contract}, MyContractTransactor: MyContractTransactor{contract: contract}, MyContractFilterer: MyContractFilterer{contract: contract}}, nil
}

// NewMyContractCaller creates a new read-only instance of MyContract, bound to a specific deployed contract.
func NewMyContractCaller(address common.Address, caller bind.ContractCaller) (*MyContractCaller, error) {
	contract, err := bindMyContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MyContractCaller{contract: contract}, nil
}

// NewMyContractTransactor creates a new write-only instance of MyContract, bound to a specific deployed contract.
func NewMyContractTransactor(address common.Address, transactor bind.ContractTransactor) (*MyContractTransactor, error) {
	contract, err := bindMyContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MyContractTransactor{contract: contract}, nil
}

// NewMyContractFilterer creates a new log filterer instance of MyContract, bound to a specific deployed contract.
func NewMyContractFilterer(address common.Address, filterer bind.ContractFilterer) (*MyContractFilterer, error) {
	contract, err := bindMyContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MyContractFilterer{contract: contract}, nil
}

// bindMyContract binds a generic wrapper to an already deployed contract.
func bindMyContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MyContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MyContract *MyContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MyContract.Contract.MyContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MyContract *MyContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MyContract.Contract.MyContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MyContract *MyContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MyContract.Contract.MyContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MyContract *MyContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MyContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MyContract *MyContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MyContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MyContract *MyContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MyContract.Contract.contract.Transact(opts, method, params...)
}

// SYSADDR is a free data retrieval call binding the contract method 0xde230e32.
//
// Solidity: function SYS_ADDR() view returns(address)
func (_MyContract *MyContractCaller) SYSADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MyContract.contract.Call(opts, &out, "SYS_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SYSADDR is a free data retrieval call binding the contract method 0xde230e32.
//
// Solidity: function SYS_ADDR() view returns(address)
func (_MyContract *MyContractSession) SYSADDR() (common.Address, error) {
	return _MyContract.Contract.SYSADDR(&_MyContract.CallOpts)
}

// SYSADDR is a free data retrieval call binding the contract method 0xde230e32.
//
// Solidity: function SYS_ADDR() view returns(address)
func (_MyContract *MyContractCallerSession) SYSADDR() (common.Address, error) {
	return _MyContract.Contract.SYSADDR(&_MyContract.CallOpts)
}

// TUSDADDR is a free data retrieval call binding the contract method 0x3989533f.
//
// Solidity: function TUSD_ADDR() view returns(address)
func (_MyContract *MyContractCaller) TUSDADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MyContract.contract.Call(opts, &out, "TUSD_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TUSDADDR is a free data retrieval call binding the contract method 0x3989533f.
//
// Solidity: function TUSD_ADDR() view returns(address)
func (_MyContract *MyContractSession) TUSDADDR() (common.Address, error) {
	return _MyContract.Contract.TUSDADDR(&_MyContract.CallOpts)
}

// TUSDADDR is a free data retrieval call binding the contract method 0x3989533f.
//
// Solidity: function TUSD_ADDR() view returns(address)
func (_MyContract *MyContractCallerSession) TUSDADDR() (common.Address, error) {
	return _MyContract.Contract.TUSDADDR(&_MyContract.CallOpts)
}

// AddrToId is a free data retrieval call binding the contract method 0xa33c4df1.
//
// Solidity: function addrToId(address ) view returns(uint64)
func (_MyContract *MyContractCaller) AddrToId(opts *bind.CallOpts, arg0 common.Address) (uint64, error) {
	var out []interface{}
	err := _MyContract.contract.Call(opts, &out, "addrToId", arg0)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// AddrToId is a free data retrieval call binding the contract method 0xa33c4df1.
//
// Solidity: function addrToId(address ) view returns(uint64)
func (_MyContract *MyContractSession) AddrToId(arg0 common.Address) (uint64, error) {
	return _MyContract.Contract.AddrToId(&_MyContract.CallOpts, arg0)
}

// AddrToId is a free data retrieval call binding the contract method 0xa33c4df1.
//
// Solidity: function addrToId(address ) view returns(uint64)
func (_MyContract *MyContractCallerSession) AddrToId(arg0 common.Address) (uint64, error) {
	return _MyContract.Contract.AddrToId(&_MyContract.CallOpts, arg0)
}

// GetBalanceInfo is a free data retrieval call binding the contract method 0x8a63891e.
//
// Solidity: function get_balance_info() view returns(uint256 SYS_balance, uint256 TUSD_balance)
func (_MyContract *MyContractCaller) GetBalanceInfo(opts *bind.CallOpts) (struct {
	SYSBalance  *big.Int
	TUSDBalance *big.Int
}, error) {
	var out []interface{}
	err := _MyContract.contract.Call(opts, &out, "get_balance_info")

	outstruct := new(struct {
		SYSBalance  *big.Int
		TUSDBalance *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.SYSBalance = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.TUSDBalance = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetBalanceInfo is a free data retrieval call binding the contract method 0x8a63891e.
//
// Solidity: function get_balance_info() view returns(uint256 SYS_balance, uint256 TUSD_balance)
func (_MyContract *MyContractSession) GetBalanceInfo() (struct {
	SYSBalance  *big.Int
	TUSDBalance *big.Int
}, error) {
	return _MyContract.Contract.GetBalanceInfo(&_MyContract.CallOpts)
}

// GetBalanceInfo is a free data retrieval call binding the contract method 0x8a63891e.
//
// Solidity: function get_balance_info() view returns(uint256 SYS_balance, uint256 TUSD_balance)
func (_MyContract *MyContractCallerSession) GetBalanceInfo() (struct {
	SYSBalance  *big.Int
	TUSDBalance *big.Int
}, error) {
	return _MyContract.Contract.GetBalanceInfo(&_MyContract.CallOpts)
}

// GetPlayerCount is a free data retrieval call binding the contract method 0x5dd94772.
//
// Solidity: function get_player_count() view returns(uint64)
func (_MyContract *MyContractCaller) GetPlayerCount(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _MyContract.contract.Call(opts, &out, "get_player_count")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetPlayerCount is a free data retrieval call binding the contract method 0x5dd94772.
//
// Solidity: function get_player_count() view returns(uint64)
func (_MyContract *MyContractSession) GetPlayerCount() (uint64, error) {
	return _MyContract.Contract.GetPlayerCount(&_MyContract.CallOpts)
}

// GetPlayerCount is a free data retrieval call binding the contract method 0x5dd94772.
//
// Solidity: function get_player_count() view returns(uint64)
func (_MyContract *MyContractCallerSession) GetPlayerCount() (uint64, error) {
	return _MyContract.Contract.GetPlayerCount(&_MyContract.CallOpts)
}

// IdToAddr is a free data retrieval call binding the contract method 0x80b6b967.
//
// Solidity: function idToAddr(uint64 ) view returns(address)
func (_MyContract *MyContractCaller) IdToAddr(opts *bind.CallOpts, arg0 uint64) (common.Address, error) {
	var out []interface{}
	err := _MyContract.contract.Call(opts, &out, "idToAddr", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// IdToAddr is a free data retrieval call binding the contract method 0x80b6b967.
//
// Solidity: function idToAddr(uint64 ) view returns(address)
func (_MyContract *MyContractSession) IdToAddr(arg0 uint64) (common.Address, error) {
	return _MyContract.Contract.IdToAddr(&_MyContract.CallOpts, arg0)
}

// IdToAddr is a free data retrieval call binding the contract method 0x80b6b967.
//
// Solidity: function idToAddr(uint64 ) view returns(address)
func (_MyContract *MyContractCallerSession) IdToAddr(arg0 uint64) (common.Address, error) {
	return _MyContract.Contract.IdToAddr(&_MyContract.CallOpts, arg0)
}

// Get is a paid mutator transaction binding the contract method 0xf72d0b3b.
//
// Solidity: function get(uint8 _type) returns()
func (_MyContract *MyContractTransactor) Get(opts *bind.TransactOpts, _type uint8) (*types.Transaction, error) {
	return _MyContract.contract.Transact(opts, "get", _type)
}

// Get is a paid mutator transaction binding the contract method 0xf72d0b3b.
//
// Solidity: function get(uint8 _type) returns()
func (_MyContract *MyContractSession) Get(_type uint8) (*types.Transaction, error) {
	return _MyContract.Contract.Get(&_MyContract.TransactOpts, _type)
}

// Get is a paid mutator transaction binding the contract method 0xf72d0b3b.
//
// Solidity: function get(uint8 _type) returns()
func (_MyContract *MyContractTransactorSession) Get(_type uint8) (*types.Transaction, error) {
	return _MyContract.Contract.Get(&_MyContract.TransactOpts, _type)
}

// Pay is a paid mutator transaction binding the contract method 0xabab2edd.
//
// Solidity: function pay(uint128 _value, uint8 _type) returns()
func (_MyContract *MyContractTransactor) Pay(opts *bind.TransactOpts, _value *big.Int, _type uint8) (*types.Transaction, error) {
	return _MyContract.contract.Transact(opts, "pay", _value, _type)
}

// Pay is a paid mutator transaction binding the contract method 0xabab2edd.
//
// Solidity: function pay(uint128 _value, uint8 _type) returns()
func (_MyContract *MyContractSession) Pay(_value *big.Int, _type uint8) (*types.Transaction, error) {
	return _MyContract.Contract.Pay(&_MyContract.TransactOpts, _value, _type)
}

// Pay is a paid mutator transaction binding the contract method 0xabab2edd.
//
// Solidity: function pay(uint128 _value, uint8 _type) returns()
func (_MyContract *MyContractTransactorSession) Pay(_value *big.Int, _type uint8) (*types.Transaction, error) {
	return _MyContract.Contract.Pay(&_MyContract.TransactOpts, _value, _type)
}

// Register is a paid mutator transaction binding the contract method 0x1aa3a008.
//
// Solidity: function register() returns()
func (_MyContract *MyContractTransactor) Register(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MyContract.contract.Transact(opts, "register")
}

// Register is a paid mutator transaction binding the contract method 0x1aa3a008.
//
// Solidity: function register() returns()
func (_MyContract *MyContractSession) Register() (*types.Transaction, error) {
	return _MyContract.Contract.Register(&_MyContract.TransactOpts)
}

// Register is a paid mutator transaction binding the contract method 0x1aa3a008.
//
// Solidity: function register() returns()
func (_MyContract *MyContractTransactorSession) Register() (*types.Transaction, error) {
	return _MyContract.Contract.Register(&_MyContract.TransactOpts)
}

// SetOpAddr is a paid mutator transaction binding the contract method 0xc5fd955b.
//
// Solidity: function set_op_addr(address opAddr) returns()
func (_MyContract *MyContractTransactor) SetOpAddr(opts *bind.TransactOpts, opAddr common.Address) (*types.Transaction, error) {
	return _MyContract.contract.Transact(opts, "set_op_addr", opAddr)
}

// SetOpAddr is a paid mutator transaction binding the contract method 0xc5fd955b.
//
// Solidity: function set_op_addr(address opAddr) returns()
func (_MyContract *MyContractSession) SetOpAddr(opAddr common.Address) (*types.Transaction, error) {
	return _MyContract.Contract.SetOpAddr(&_MyContract.TransactOpts, opAddr)
}

// SetOpAddr is a paid mutator transaction binding the contract method 0xc5fd955b.
//
// Solidity: function set_op_addr(address opAddr) returns()
func (_MyContract *MyContractTransactorSession) SetOpAddr(opAddr common.Address) (*types.Transaction, error) {
	return _MyContract.Contract.SetOpAddr(&_MyContract.TransactOpts, opAddr)
}

// WithdrawAdmin is a paid mutator transaction binding the contract method 0x5e144d2c.
//
// Solidity: function withdrawAdmin(uint256 val, uint8 _token_type) returns()
func (_MyContract *MyContractTransactor) WithdrawAdmin(opts *bind.TransactOpts, val *big.Int, _token_type uint8) (*types.Transaction, error) {
	return _MyContract.contract.Transact(opts, "withdrawAdmin", val, _token_type)
}

// WithdrawAdmin is a paid mutator transaction binding the contract method 0x5e144d2c.
//
// Solidity: function withdrawAdmin(uint256 val, uint8 _token_type) returns()
func (_MyContract *MyContractSession) WithdrawAdmin(val *big.Int, _token_type uint8) (*types.Transaction, error) {
	return _MyContract.Contract.WithdrawAdmin(&_MyContract.TransactOpts, val, _token_type)
}

// WithdrawAdmin is a paid mutator transaction binding the contract method 0x5e144d2c.
//
// Solidity: function withdrawAdmin(uint256 val, uint8 _token_type) returns()
func (_MyContract *MyContractTransactorSession) WithdrawAdmin(val *big.Int, _token_type uint8) (*types.Transaction, error) {
	return _MyContract.Contract.WithdrawAdmin(&_MyContract.TransactOpts, val, _token_type)
}

// MyContractJoinIterator is returned from FilterJoin and is used to iterate over the raw logs and unpacked data for Join events raised by the MyContract contract.
type MyContractJoinIterator struct {
	Event *MyContractJoin // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MyContractJoinIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MyContractJoin)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MyContractJoin)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MyContractJoinIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MyContractJoinIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MyContractJoin represents a Join event raised by the MyContract contract.
type MyContractJoin struct {
	Id   uint64
	Addr common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterJoin is a free log retrieval operation binding the contract event 0x17c4fb9bfda0c2290c44441da6ae66874c1a70f483e08fc4dac554bec05f5bb7.
//
// Solidity: event join(uint64 id, address indexed addr)
func (_MyContract *MyContractFilterer) FilterJoin(opts *bind.FilterOpts, addr []common.Address) (*MyContractJoinIterator, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _MyContract.contract.FilterLogs(opts, "join", addrRule)
	if err != nil {
		return nil, err
	}
	return &MyContractJoinIterator{contract: _MyContract.contract, event: "join", logs: logs, sub: sub}, nil
}

// WatchJoin is a free log subscription operation binding the contract event 0x17c4fb9bfda0c2290c44441da6ae66874c1a70f483e08fc4dac554bec05f5bb7.
//
// Solidity: event join(uint64 id, address indexed addr)
func (_MyContract *MyContractFilterer) WatchJoin(opts *bind.WatchOpts, sink chan<- *MyContractJoin, addr []common.Address) (event.Subscription, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _MyContract.contract.WatchLogs(opts, "join", addrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MyContractJoin)
				if err := _MyContract.contract.UnpackLog(event, "join", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseJoin is a log parse operation binding the contract event 0x17c4fb9bfda0c2290c44441da6ae66874c1a70f483e08fc4dac554bec05f5bb7.
//
// Solidity: event join(uint64 id, address indexed addr)
func (_MyContract *MyContractFilterer) ParseJoin(log types.Log) (*MyContractJoin, error) {
	event := new(MyContractJoin)
	if err := _MyContract.contract.UnpackLog(event, "join", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MyContractPaylogIterator is returned from FilterPaylog and is used to iterate over the raw logs and unpacked data for Paylog events raised by the MyContract contract.
type MyContractPaylogIterator struct {
	Event *MyContractPaylog // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MyContractPaylogIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MyContractPaylog)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MyContractPaylog)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MyContractPaylogIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MyContractPaylogIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MyContractPaylog represents a Paylog event raised by the MyContract contract.
type MyContractPaylog struct {
	Id    uint64
	Value *big.Int
	Type  uint8
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterPaylog is a free log retrieval operation binding the contract event 0x37501fbcebe25e5c825eeb72a852927e128e4b9b35cbc162045c112dba89e5db.
//
// Solidity: event paylog(uint64 id, uint128 _value, uint8 _type)
func (_MyContract *MyContractFilterer) FilterPaylog(opts *bind.FilterOpts) (*MyContractPaylogIterator, error) {

	logs, sub, err := _MyContract.contract.FilterLogs(opts, "paylog")
	if err != nil {
		return nil, err
	}
	return &MyContractPaylogIterator{contract: _MyContract.contract, event: "paylog", logs: logs, sub: sub}, nil
}

// WatchPaylog is a free log subscription operation binding the contract event 0x37501fbcebe25e5c825eeb72a852927e128e4b9b35cbc162045c112dba89e5db.
//
// Solidity: event paylog(uint64 id, uint128 _value, uint8 _type)
func (_MyContract *MyContractFilterer) WatchPaylog(opts *bind.WatchOpts, sink chan<- *MyContractPaylog) (event.Subscription, error) {

	logs, sub, err := _MyContract.contract.WatchLogs(opts, "paylog")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MyContractPaylog)
				if err := _MyContract.contract.UnpackLog(event, "paylog", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePaylog is a log parse operation binding the contract event 0x37501fbcebe25e5c825eeb72a852927e128e4b9b35cbc162045c112dba89e5db.
//
// Solidity: event paylog(uint64 id, uint128 _value, uint8 _type)
func (_MyContract *MyContractFilterer) ParsePaylog(log types.Log) (*MyContractPaylog, error) {
	event := new(MyContractPaylog)
	if err := _MyContract.contract.UnpackLog(event, "paylog", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
