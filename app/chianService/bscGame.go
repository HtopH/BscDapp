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

// BscGameMetaData contains all meta data concerning the BscGame contract.
var BscGameMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"id\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"join\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"id\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"_value\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"_type\",\"type\":\"uint8\"}],\"name\":\"paylog\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"SYS_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TUSD_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"addrToId\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_type\",\"type\":\"uint8\"}],\"name\":\"get\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"get_balance_info\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"SYS_balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"TUSD_balance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"get_player_count\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"idToAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"_value\",\"type\":\"uint128\"},{\"internalType\":\"uint8\",\"name\":\"_type\",\"type\":\"uint8\"}],\"name\":\"pay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"register\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"opAddr\",\"type\":\"address\"}],\"name\":\"set_op_addr\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"val\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"_token_type\",\"type\":\"uint8\"}],\"name\":\"withdrawAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// BscGameABI is the input ABI used to generate the binding from.
// Deprecated: Use BscGameMetaData.ABI instead.
var BscGameABI = BscGameMetaData.ABI

// BscGame is an auto generated Go binding around an Ethereum contract.
type BscGame struct {
	BscGameCaller     // Read-only binding to the contract
	BscGameTransactor // Write-only binding to the contract
	BscGameFilterer   // Log filterer for contract events
}

// BscGameCaller is an auto generated read-only Go binding around an Ethereum contract.
type BscGameCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BscGameTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BscGameTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BscGameFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BscGameFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BscGameSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BscGameSession struct {
	Contract     *BscGame          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BscGameCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BscGameCallerSession struct {
	Contract *BscGameCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// BscGameTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BscGameTransactorSession struct {
	Contract     *BscGameTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// BscGameRaw is an auto generated low-level Go binding around an Ethereum contract.
type BscGameRaw struct {
	Contract *BscGame // Generic contract binding to access the raw methods on
}

// BscGameCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BscGameCallerRaw struct {
	Contract *BscGameCaller // Generic read-only contract binding to access the raw methods on
}

// BscGameTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BscGameTransactorRaw struct {
	Contract *BscGameTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBscGame creates a new instance of BscGame, bound to a specific deployed contract.
func NewBscGame(address common.Address, backend bind.ContractBackend) (*BscGame, error) {
	contract, err := bindBscGame(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BscGame{BscGameCaller: BscGameCaller{contract: contract}, BscGameTransactor: BscGameTransactor{contract: contract}, BscGameFilterer: BscGameFilterer{contract: contract}}, nil
}

// NewBscGameCaller creates a new read-only instance of BscGame, bound to a specific deployed contract.
func NewBscGameCaller(address common.Address, caller bind.ContractCaller) (*BscGameCaller, error) {
	contract, err := bindBscGame(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BscGameCaller{contract: contract}, nil
}

// NewBscGameTransactor creates a new write-only instance of BscGame, bound to a specific deployed contract.
func NewBscGameTransactor(address common.Address, transactor bind.ContractTransactor) (*BscGameTransactor, error) {
	contract, err := bindBscGame(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BscGameTransactor{contract: contract}, nil
}

// NewBscGameFilterer creates a new log filterer instance of BscGame, bound to a specific deployed contract.
func NewBscGameFilterer(address common.Address, filterer bind.ContractFilterer) (*BscGameFilterer, error) {
	contract, err := bindBscGame(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BscGameFilterer{contract: contract}, nil
}

// bindBscGame binds a generic wrapper to an already deployed contract.
func bindBscGame(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BscGameABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BscGame *BscGameRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BscGame.Contract.BscGameCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BscGame *BscGameRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BscGame.Contract.BscGameTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BscGame *BscGameRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BscGame.Contract.BscGameTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BscGame *BscGameCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BscGame.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BscGame *BscGameTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BscGame.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BscGame *BscGameTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BscGame.Contract.contract.Transact(opts, method, params...)
}

// SYSADDR is a free data retrieval call binding the contract method 0xde230e32.
//
// Solidity: function SYS_ADDR() view returns(address)
func (_BscGame *BscGameCaller) SYSADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BscGame.contract.Call(opts, &out, "SYS_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SYSADDR is a free data retrieval call binding the contract method 0xde230e32.
//
// Solidity: function SYS_ADDR() view returns(address)
func (_BscGame *BscGameSession) SYSADDR() (common.Address, error) {
	return _BscGame.Contract.SYSADDR(&_BscGame.CallOpts)
}

// SYSADDR is a free data retrieval call binding the contract method 0xde230e32.
//
// Solidity: function SYS_ADDR() view returns(address)
func (_BscGame *BscGameCallerSession) SYSADDR() (common.Address, error) {
	return _BscGame.Contract.SYSADDR(&_BscGame.CallOpts)
}

// TUSDADDR is a free data retrieval call binding the contract method 0x3989533f.
//
// Solidity: function TUSD_ADDR() view returns(address)
func (_BscGame *BscGameCaller) TUSDADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BscGame.contract.Call(opts, &out, "TUSD_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TUSDADDR is a free data retrieval call binding the contract method 0x3989533f.
//
// Solidity: function TUSD_ADDR() view returns(address)
func (_BscGame *BscGameSession) TUSDADDR() (common.Address, error) {
	return _BscGame.Contract.TUSDADDR(&_BscGame.CallOpts)
}

// TUSDADDR is a free data retrieval call binding the contract method 0x3989533f.
//
// Solidity: function TUSD_ADDR() view returns(address)
func (_BscGame *BscGameCallerSession) TUSDADDR() (common.Address, error) {
	return _BscGame.Contract.TUSDADDR(&_BscGame.CallOpts)
}

// AddrToId is a free data retrieval call binding the contract method 0xa33c4df1.
//
// Solidity: function addrToId(address ) view returns(uint64)
func (_BscGame *BscGameCaller) AddrToId(opts *bind.CallOpts, arg0 common.Address) (uint64, error) {
	var out []interface{}
	err := _BscGame.contract.Call(opts, &out, "addrToId", arg0)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// AddrToId is a free data retrieval call binding the contract method 0xa33c4df1.
//
// Solidity: function addrToId(address ) view returns(uint64)
func (_BscGame *BscGameSession) AddrToId(arg0 common.Address) (uint64, error) {
	return _BscGame.Contract.AddrToId(&_BscGame.CallOpts, arg0)
}

// AddrToId is a free data retrieval call binding the contract method 0xa33c4df1.
//
// Solidity: function addrToId(address ) view returns(uint64)
func (_BscGame *BscGameCallerSession) AddrToId(arg0 common.Address) (uint64, error) {
	return _BscGame.Contract.AddrToId(&_BscGame.CallOpts, arg0)
}

// GetBalanceInfo is a free data retrieval call binding the contract method 0x8a63891e.
//
// Solidity: function get_balance_info() view returns(uint256 SYS_balance, uint256 TUSD_balance)
func (_BscGame *BscGameCaller) GetBalanceInfo(opts *bind.CallOpts) (struct {
	SYSBalance  *big.Int
	TUSDBalance *big.Int
}, error) {
	var out []interface{}
	err := _BscGame.contract.Call(opts, &out, "get_balance_info")

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
func (_BscGame *BscGameSession) GetBalanceInfo() (struct {
	SYSBalance  *big.Int
	TUSDBalance *big.Int
}, error) {
	return _BscGame.Contract.GetBalanceInfo(&_BscGame.CallOpts)
}

// GetBalanceInfo is a free data retrieval call binding the contract method 0x8a63891e.
//
// Solidity: function get_balance_info() view returns(uint256 SYS_balance, uint256 TUSD_balance)
func (_BscGame *BscGameCallerSession) GetBalanceInfo() (struct {
	SYSBalance  *big.Int
	TUSDBalance *big.Int
}, error) {
	return _BscGame.Contract.GetBalanceInfo(&_BscGame.CallOpts)
}

// GetPlayerCount is a free data retrieval call binding the contract method 0x5dd94772.
//
// Solidity: function get_player_count() view returns(uint64)
func (_BscGame *BscGameCaller) GetPlayerCount(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _BscGame.contract.Call(opts, &out, "get_player_count")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetPlayerCount is a free data retrieval call binding the contract method 0x5dd94772.
//
// Solidity: function get_player_count() view returns(uint64)
func (_BscGame *BscGameSession) GetPlayerCount() (uint64, error) {
	return _BscGame.Contract.GetPlayerCount(&_BscGame.CallOpts)
}

// GetPlayerCount is a free data retrieval call binding the contract method 0x5dd94772.
//
// Solidity: function get_player_count() view returns(uint64)
func (_BscGame *BscGameCallerSession) GetPlayerCount() (uint64, error) {
	return _BscGame.Contract.GetPlayerCount(&_BscGame.CallOpts)
}

// IdToAddr is a free data retrieval call binding the contract method 0x80b6b967.
//
// Solidity: function idToAddr(uint64 ) view returns(address)
func (_BscGame *BscGameCaller) IdToAddr(opts *bind.CallOpts, arg0 uint64) (common.Address, error) {
	var out []interface{}
	err := _BscGame.contract.Call(opts, &out, "idToAddr", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// IdToAddr is a free data retrieval call binding the contract method 0x80b6b967.
//
// Solidity: function idToAddr(uint64 ) view returns(address)
func (_BscGame *BscGameSession) IdToAddr(arg0 uint64) (common.Address, error) {
	return _BscGame.Contract.IdToAddr(&_BscGame.CallOpts, arg0)
}

// IdToAddr is a free data retrieval call binding the contract method 0x80b6b967.
//
// Solidity: function idToAddr(uint64 ) view returns(address)
func (_BscGame *BscGameCallerSession) IdToAddr(arg0 uint64) (common.Address, error) {
	return _BscGame.Contract.IdToAddr(&_BscGame.CallOpts, arg0)
}

// Get is a paid mutator transaction binding the contract method 0xf72d0b3b.
//
// Solidity: function get(uint8 _type) returns()
func (_BscGame *BscGameTransactor) Get(opts *bind.TransactOpts, _type uint8) (*types.Transaction, error) {
	return _BscGame.contract.Transact(opts, "get", _type)
}

// Get is a paid mutator transaction binding the contract method 0xf72d0b3b.
//
// Solidity: function get(uint8 _type) returns()
func (_BscGame *BscGameSession) Get(_type uint8) (*types.Transaction, error) {
	return _BscGame.Contract.Get(&_BscGame.TransactOpts, _type)
}

// Get is a paid mutator transaction binding the contract method 0xf72d0b3b.
//
// Solidity: function get(uint8 _type) returns()
func (_BscGame *BscGameTransactorSession) Get(_type uint8) (*types.Transaction, error) {
	return _BscGame.Contract.Get(&_BscGame.TransactOpts, _type)
}

// Pay is a paid mutator transaction binding the contract method 0xabab2edd.
//
// Solidity: function pay(uint128 _value, uint8 _type) returns()
func (_BscGame *BscGameTransactor) Pay(opts *bind.TransactOpts, _value *big.Int, _type uint8) (*types.Transaction, error) {
	return _BscGame.contract.Transact(opts, "pay", _value, _type)
}

// Pay is a paid mutator transaction binding the contract method 0xabab2edd.
//
// Solidity: function pay(uint128 _value, uint8 _type) returns()
func (_BscGame *BscGameSession) Pay(_value *big.Int, _type uint8) (*types.Transaction, error) {
	return _BscGame.Contract.Pay(&_BscGame.TransactOpts, _value, _type)
}

// Pay is a paid mutator transaction binding the contract method 0xabab2edd.
//
// Solidity: function pay(uint128 _value, uint8 _type) returns()
func (_BscGame *BscGameTransactorSession) Pay(_value *big.Int, _type uint8) (*types.Transaction, error) {
	return _BscGame.Contract.Pay(&_BscGame.TransactOpts, _value, _type)
}

// Register is a paid mutator transaction binding the contract method 0x1aa3a008.
//
// Solidity: function register() returns()
func (_BscGame *BscGameTransactor) Register(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BscGame.contract.Transact(opts, "register")
}

// Register is a paid mutator transaction binding the contract method 0x1aa3a008.
//
// Solidity: function register() returns()
func (_BscGame *BscGameSession) Register() (*types.Transaction, error) {
	return _BscGame.Contract.Register(&_BscGame.TransactOpts)
}

// Register is a paid mutator transaction binding the contract method 0x1aa3a008.
//
// Solidity: function register() returns()
func (_BscGame *BscGameTransactorSession) Register() (*types.Transaction, error) {
	return _BscGame.Contract.Register(&_BscGame.TransactOpts)
}

// SetOpAddr is a paid mutator transaction binding the contract method 0xc5fd955b.
//
// Solidity: function set_op_addr(address opAddr) returns()
func (_BscGame *BscGameTransactor) SetOpAddr(opts *bind.TransactOpts, opAddr common.Address) (*types.Transaction, error) {
	return _BscGame.contract.Transact(opts, "set_op_addr", opAddr)
}

// SetOpAddr is a paid mutator transaction binding the contract method 0xc5fd955b.
//
// Solidity: function set_op_addr(address opAddr) returns()
func (_BscGame *BscGameSession) SetOpAddr(opAddr common.Address) (*types.Transaction, error) {
	return _BscGame.Contract.SetOpAddr(&_BscGame.TransactOpts, opAddr)
}

// SetOpAddr is a paid mutator transaction binding the contract method 0xc5fd955b.
//
// Solidity: function set_op_addr(address opAddr) returns()
func (_BscGame *BscGameTransactorSession) SetOpAddr(opAddr common.Address) (*types.Transaction, error) {
	return _BscGame.Contract.SetOpAddr(&_BscGame.TransactOpts, opAddr)
}

// WithdrawAdmin is a paid mutator transaction binding the contract method 0x5e144d2c.
//
// Solidity: function withdrawAdmin(uint256 val, uint8 _token_type) returns()
func (_BscGame *BscGameTransactor) WithdrawAdmin(opts *bind.TransactOpts, val *big.Int, _token_type uint8) (*types.Transaction, error) {
	return _BscGame.contract.Transact(opts, "withdrawAdmin", val, _token_type)
}

// WithdrawAdmin is a paid mutator transaction binding the contract method 0x5e144d2c.
//
// Solidity: function withdrawAdmin(uint256 val, uint8 _token_type) returns()
func (_BscGame *BscGameSession) WithdrawAdmin(val *big.Int, _token_type uint8) (*types.Transaction, error) {
	return _BscGame.Contract.WithdrawAdmin(&_BscGame.TransactOpts, val, _token_type)
}

// WithdrawAdmin is a paid mutator transaction binding the contract method 0x5e144d2c.
//
// Solidity: function withdrawAdmin(uint256 val, uint8 _token_type) returns()
func (_BscGame *BscGameTransactorSession) WithdrawAdmin(val *big.Int, _token_type uint8) (*types.Transaction, error) {
	return _BscGame.Contract.WithdrawAdmin(&_BscGame.TransactOpts, val, _token_type)
}

// BscGameJoinIterator is returned from FilterJoin and is used to iterate over the raw logs and unpacked data for Join events raised by the BscGame contract.
type BscGameJoinIterator struct {
	Event *BscGameJoin // Event containing the contract specifics and raw log

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
func (it *BscGameJoinIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BscGameJoin)
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
		it.Event = new(BscGameJoin)
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
func (it *BscGameJoinIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BscGameJoinIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BscGameJoin represents a Join event raised by the BscGame contract.
type BscGameJoin struct {
	Id   uint64
	Addr common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterJoin is a free log retrieval operation binding the contract event 0x17c4fb9bfda0c2290c44441da6ae66874c1a70f483e08fc4dac554bec05f5bb7.
//
// Solidity: event join(uint64 id, address indexed addr)
func (_BscGame *BscGameFilterer) FilterJoin(opts *bind.FilterOpts, addr []common.Address) (*BscGameJoinIterator, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _BscGame.contract.FilterLogs(opts, "join", addrRule)
	if err != nil {
		return nil, err
	}
	return &BscGameJoinIterator{contract: _BscGame.contract, event: "join", logs: logs, sub: sub}, nil
}

// WatchJoin is a free log subscription operation binding the contract event 0x17c4fb9bfda0c2290c44441da6ae66874c1a70f483e08fc4dac554bec05f5bb7.
//
// Solidity: event join(uint64 id, address indexed addr)
func (_BscGame *BscGameFilterer) WatchJoin(opts *bind.WatchOpts, sink chan<- *BscGameJoin, addr []common.Address) (event.Subscription, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _BscGame.contract.WatchLogs(opts, "join", addrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BscGameJoin)
				if err := _BscGame.contract.UnpackLog(event, "join", log); err != nil {
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
func (_BscGame *BscGameFilterer) ParseJoin(log types.Log) (*BscGameJoin, error) {
	event := new(BscGameJoin)
	if err := _BscGame.contract.UnpackLog(event, "join", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BscGamePaylogIterator is returned from FilterPaylog and is used to iterate over the raw logs and unpacked data for Paylog events raised by the BscGame contract.
type BscGamePaylogIterator struct {
	Event *BscGamePaylog // Event containing the contract specifics and raw log

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
func (it *BscGamePaylogIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BscGamePaylog)
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
		it.Event = new(BscGamePaylog)
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
func (it *BscGamePaylogIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BscGamePaylogIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BscGamePaylog represents a Paylog event raised by the BscGame contract.
type BscGamePaylog struct {
	Id    uint64
	Value *big.Int
	Type  uint8
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterPaylog is a free log retrieval operation binding the contract event 0x37501fbcebe25e5c825eeb72a852927e128e4b9b35cbc162045c112dba89e5db.
//
// Solidity: event paylog(uint64 id, uint128 _value, uint8 _type)
func (_BscGame *BscGameFilterer) FilterPaylog(opts *bind.FilterOpts) (*BscGamePaylogIterator, error) {

	logs, sub, err := _BscGame.contract.FilterLogs(opts, "paylog")
	if err != nil {
		return nil, err
	}
	return &BscGamePaylogIterator{contract: _BscGame.contract, event: "paylog", logs: logs, sub: sub}, nil
}

// WatchPaylog is a free log subscription operation binding the contract event 0x37501fbcebe25e5c825eeb72a852927e128e4b9b35cbc162045c112dba89e5db.
//
// Solidity: event paylog(uint64 id, uint128 _value, uint8 _type)
func (_BscGame *BscGameFilterer) WatchPaylog(opts *bind.WatchOpts, sink chan<- *BscGamePaylog) (event.Subscription, error) {

	logs, sub, err := _BscGame.contract.WatchLogs(opts, "paylog")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BscGamePaylog)
				if err := _BscGame.contract.UnpackLog(event, "paylog", log); err != nil {
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
func (_BscGame *BscGameFilterer) ParsePaylog(log types.Log) (*BscGamePaylog, error) {
	event := new(BscGamePaylog)
	if err := _BscGame.contract.UnpackLog(event, "paylog", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
