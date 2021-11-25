// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package chainService

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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"doType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"id\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"_value\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"getTicket\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"percent\",\"type\":\"uint32\"}],\"name\":\"buyTicketLog\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"doType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"id\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"_value\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"_round\",\"type\":\"uint32\"}],\"name\":\"joinLog\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"doType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"id\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"refId\",\"type\":\"uint64\"}],\"name\":\"registerLog\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"doType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"id\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"_value\",\"type\":\"uint128\"}],\"name\":\"userGetLog\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"addrToId\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"val\",\"type\":\"uint128\"},{\"internalType\":\"uint8\",\"name\":\"_token_type\",\"type\":\"uint8\"}],\"name\":\"adminWithdraw\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"n\",\"type\":\"uint32\"}],\"name\":\"getBase\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNo\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"n\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPercent\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"percent\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_id\",\"type\":\"uint64\"}],\"name\":\"getUserInfo\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"_getTickets\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"_useTickets\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"_balanceU\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"_useU\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"_refId\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"get_balance_info\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"ticketBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"usdBalance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"get_player_count\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"idToAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"joinBase\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"_value\",\"type\":\"uint128\"}],\"name\":\"joinGame\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_id\",\"type\":\"uint64\"},{\"internalType\":\"uint128\",\"name\":\"_value\",\"type\":\"uint128\"}],\"name\":\"pay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"_value\",\"type\":\"uint128\"}],\"name\":\"payForTickets\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ref_addr\",\"type\":\"address\"}],\"name\":\"register\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"round\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"setRound\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"_value\",\"type\":\"uint128\"}],\"name\":\"setSpend\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"opAddr\",\"type\":\"address\"}],\"name\":\"set_op_addr\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"spendTickets\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"userId\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"_round\",\"type\":\"uint32\"}],\"name\":\"userOut\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"_value\",\"type\":\"uint128\"}],\"name\":\"userWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
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

// GetBase is a free data retrieval call binding the contract method 0xd9ff56bd.
//
// Solidity: function getBase(uint32 n) pure returns(uint32)
func (_BscGame *BscGameCaller) GetBase(opts *bind.CallOpts, n uint32) (uint32, error) {
	var out []interface{}
	err := _BscGame.contract.Call(opts, &out, "getBase", n)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetBase is a free data retrieval call binding the contract method 0xd9ff56bd.
//
// Solidity: function getBase(uint32 n) pure returns(uint32)
func (_BscGame *BscGameSession) GetBase(n uint32) (uint32, error) {
	return _BscGame.Contract.GetBase(&_BscGame.CallOpts, n)
}

// GetBase is a free data retrieval call binding the contract method 0xd9ff56bd.
//
// Solidity: function getBase(uint32 n) pure returns(uint32)
func (_BscGame *BscGameCallerSession) GetBase(n uint32) (uint32, error) {
	return _BscGame.Contract.GetBase(&_BscGame.CallOpts, n)
}

// GetNo is a free data retrieval call binding the contract method 0x9f0314fc.
//
// Solidity: function getNo() view returns(uint32 n)
func (_BscGame *BscGameCaller) GetNo(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _BscGame.contract.Call(opts, &out, "getNo")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetNo is a free data retrieval call binding the contract method 0x9f0314fc.
//
// Solidity: function getNo() view returns(uint32 n)
func (_BscGame *BscGameSession) GetNo() (uint32, error) {
	return _BscGame.Contract.GetNo(&_BscGame.CallOpts)
}

// GetNo is a free data retrieval call binding the contract method 0x9f0314fc.
//
// Solidity: function getNo() view returns(uint32 n)
func (_BscGame *BscGameCallerSession) GetNo() (uint32, error) {
	return _BscGame.Contract.GetNo(&_BscGame.CallOpts)
}

// GetPercent is a free data retrieval call binding the contract method 0x2a6dd8c9.
//
// Solidity: function getPercent() view returns(uint32 percent)
func (_BscGame *BscGameCaller) GetPercent(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _BscGame.contract.Call(opts, &out, "getPercent")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetPercent is a free data retrieval call binding the contract method 0x2a6dd8c9.
//
// Solidity: function getPercent() view returns(uint32 percent)
func (_BscGame *BscGameSession) GetPercent() (uint32, error) {
	return _BscGame.Contract.GetPercent(&_BscGame.CallOpts)
}

// GetPercent is a free data retrieval call binding the contract method 0x2a6dd8c9.
//
// Solidity: function getPercent() view returns(uint32 percent)
func (_BscGame *BscGameCallerSession) GetPercent() (uint32, error) {
	return _BscGame.Contract.GetPercent(&_BscGame.CallOpts)
}

// GetUserInfo is a free data retrieval call binding the contract method 0xb495763b.
//
// Solidity: function getUserInfo(uint64 _id) view returns(uint128 _getTickets, uint128 _useTickets, uint128 _balanceU, uint128 _useU, uint128 _refId)
func (_BscGame *BscGameCaller) GetUserInfo(opts *bind.CallOpts, _id uint64) (struct {
	GetTickets *big.Int
	UseTickets *big.Int
	BalanceU   *big.Int
	UseU       *big.Int
	RefId      *big.Int
}, error) {
	var out []interface{}
	err := _BscGame.contract.Call(opts, &out, "getUserInfo", _id)

	outstruct := new(struct {
		GetTickets *big.Int
		UseTickets *big.Int
		BalanceU   *big.Int
		UseU       *big.Int
		RefId      *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.GetTickets = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.UseTickets = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.BalanceU = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.UseU = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.RefId = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetUserInfo is a free data retrieval call binding the contract method 0xb495763b.
//
// Solidity: function getUserInfo(uint64 _id) view returns(uint128 _getTickets, uint128 _useTickets, uint128 _balanceU, uint128 _useU, uint128 _refId)
func (_BscGame *BscGameSession) GetUserInfo(_id uint64) (struct {
	GetTickets *big.Int
	UseTickets *big.Int
	BalanceU   *big.Int
	UseU       *big.Int
	RefId      *big.Int
}, error) {
	return _BscGame.Contract.GetUserInfo(&_BscGame.CallOpts, _id)
}

// GetUserInfo is a free data retrieval call binding the contract method 0xb495763b.
//
// Solidity: function getUserInfo(uint64 _id) view returns(uint128 _getTickets, uint128 _useTickets, uint128 _balanceU, uint128 _useU, uint128 _refId)
func (_BscGame *BscGameCallerSession) GetUserInfo(_id uint64) (struct {
	GetTickets *big.Int
	UseTickets *big.Int
	BalanceU   *big.Int
	UseU       *big.Int
	RefId      *big.Int
}, error) {
	return _BscGame.Contract.GetUserInfo(&_BscGame.CallOpts, _id)
}

// GetBalanceInfo is a free data retrieval call binding the contract method 0x8a63891e.
//
// Solidity: function get_balance_info() view returns(uint256 ticketBalance, uint256 usdBalance)
func (_BscGame *BscGameCaller) GetBalanceInfo(opts *bind.CallOpts) (struct {
	TicketBalance *big.Int
	UsdBalance    *big.Int
}, error) {
	var out []interface{}
	err := _BscGame.contract.Call(opts, &out, "get_balance_info")

	outstruct := new(struct {
		TicketBalance *big.Int
		UsdBalance    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TicketBalance = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.UsdBalance = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetBalanceInfo is a free data retrieval call binding the contract method 0x8a63891e.
//
// Solidity: function get_balance_info() view returns(uint256 ticketBalance, uint256 usdBalance)
func (_BscGame *BscGameSession) GetBalanceInfo() (struct {
	TicketBalance *big.Int
	UsdBalance    *big.Int
}, error) {
	return _BscGame.Contract.GetBalanceInfo(&_BscGame.CallOpts)
}

// GetBalanceInfo is a free data retrieval call binding the contract method 0x8a63891e.
//
// Solidity: function get_balance_info() view returns(uint256 ticketBalance, uint256 usdBalance)
func (_BscGame *BscGameCallerSession) GetBalanceInfo() (struct {
	TicketBalance *big.Int
	UsdBalance    *big.Int
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

// JoinBase is a free data retrieval call binding the contract method 0x9d254940.
//
// Solidity: function joinBase() view returns(uint128)
func (_BscGame *BscGameCaller) JoinBase(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BscGame.contract.Call(opts, &out, "joinBase")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// JoinBase is a free data retrieval call binding the contract method 0x9d254940.
//
// Solidity: function joinBase() view returns(uint128)
func (_BscGame *BscGameSession) JoinBase() (*big.Int, error) {
	return _BscGame.Contract.JoinBase(&_BscGame.CallOpts)
}

// JoinBase is a free data retrieval call binding the contract method 0x9d254940.
//
// Solidity: function joinBase() view returns(uint128)
func (_BscGame *BscGameCallerSession) JoinBase() (*big.Int, error) {
	return _BscGame.Contract.JoinBase(&_BscGame.CallOpts)
}

// Round is a free data retrieval call binding the contract method 0x146ca531.
//
// Solidity: function round() view returns(uint32)
func (_BscGame *BscGameCaller) Round(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _BscGame.contract.Call(opts, &out, "round")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// Round is a free data retrieval call binding the contract method 0x146ca531.
//
// Solidity: function round() view returns(uint32)
func (_BscGame *BscGameSession) Round() (uint32, error) {
	return _BscGame.Contract.Round(&_BscGame.CallOpts)
}

// Round is a free data retrieval call binding the contract method 0x146ca531.
//
// Solidity: function round() view returns(uint32)
func (_BscGame *BscGameCallerSession) Round() (uint32, error) {
	return _BscGame.Contract.Round(&_BscGame.CallOpts)
}

// SpendTickets is a free data retrieval call binding the contract method 0xe9c4c783.
//
// Solidity: function spendTickets() view returns(uint128)
func (_BscGame *BscGameCaller) SpendTickets(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BscGame.contract.Call(opts, &out, "spendTickets")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SpendTickets is a free data retrieval call binding the contract method 0xe9c4c783.
//
// Solidity: function spendTickets() view returns(uint128)
func (_BscGame *BscGameSession) SpendTickets() (*big.Int, error) {
	return _BscGame.Contract.SpendTickets(&_BscGame.CallOpts)
}

// SpendTickets is a free data retrieval call binding the contract method 0xe9c4c783.
//
// Solidity: function spendTickets() view returns(uint128)
func (_BscGame *BscGameCallerSession) SpendTickets() (*big.Int, error) {
	return _BscGame.Contract.SpendTickets(&_BscGame.CallOpts)
}

// AdminWithdraw is a paid mutator transaction binding the contract method 0x10144354.
//
// Solidity: function adminWithdraw(uint128 val, uint8 _token_type) returns(bool)
func (_BscGame *BscGameTransactor) AdminWithdraw(opts *bind.TransactOpts, val *big.Int, _token_type uint8) (*types.Transaction, error) {
	return _BscGame.contract.Transact(opts, "adminWithdraw", val, _token_type)
}

// AdminWithdraw is a paid mutator transaction binding the contract method 0x10144354.
//
// Solidity: function adminWithdraw(uint128 val, uint8 _token_type) returns(bool)
func (_BscGame *BscGameSession) AdminWithdraw(val *big.Int, _token_type uint8) (*types.Transaction, error) {
	return _BscGame.Contract.AdminWithdraw(&_BscGame.TransactOpts, val, _token_type)
}

// AdminWithdraw is a paid mutator transaction binding the contract method 0x10144354.
//
// Solidity: function adminWithdraw(uint128 val, uint8 _token_type) returns(bool)
func (_BscGame *BscGameTransactorSession) AdminWithdraw(val *big.Int, _token_type uint8) (*types.Transaction, error) {
	return _BscGame.Contract.AdminWithdraw(&_BscGame.TransactOpts, val, _token_type)
}

// JoinGame is a paid mutator transaction binding the contract method 0xb4fde155.
//
// Solidity: function joinGame(uint128 _value) returns()
func (_BscGame *BscGameTransactor) JoinGame(opts *bind.TransactOpts, _value *big.Int) (*types.Transaction, error) {
	return _BscGame.contract.Transact(opts, "joinGame", _value)
}

// JoinGame is a paid mutator transaction binding the contract method 0xb4fde155.
//
// Solidity: function joinGame(uint128 _value) returns()
func (_BscGame *BscGameSession) JoinGame(_value *big.Int) (*types.Transaction, error) {
	return _BscGame.Contract.JoinGame(&_BscGame.TransactOpts, _value)
}

// JoinGame is a paid mutator transaction binding the contract method 0xb4fde155.
//
// Solidity: function joinGame(uint128 _value) returns()
func (_BscGame *BscGameTransactorSession) JoinGame(_value *big.Int) (*types.Transaction, error) {
	return _BscGame.Contract.JoinGame(&_BscGame.TransactOpts, _value)
}

// Pay is a paid mutator transaction binding the contract method 0xdbc3051b.
//
// Solidity: function pay(uint64 _id, uint128 _value) returns()
func (_BscGame *BscGameTransactor) Pay(opts *bind.TransactOpts, _id uint64, _value *big.Int) (*types.Transaction, error) {
	return _BscGame.contract.Transact(opts, "pay", _id, _value)
}

// Pay is a paid mutator transaction binding the contract method 0xdbc3051b.
//
// Solidity: function pay(uint64 _id, uint128 _value) returns()
func (_BscGame *BscGameSession) Pay(_id uint64, _value *big.Int) (*types.Transaction, error) {
	return _BscGame.Contract.Pay(&_BscGame.TransactOpts, _id, _value)
}

// Pay is a paid mutator transaction binding the contract method 0xdbc3051b.
//
// Solidity: function pay(uint64 _id, uint128 _value) returns()
func (_BscGame *BscGameTransactorSession) Pay(_id uint64, _value *big.Int) (*types.Transaction, error) {
	return _BscGame.Contract.Pay(&_BscGame.TransactOpts, _id, _value)
}

// PayForTickets is a paid mutator transaction binding the contract method 0xbc553d8f.
//
// Solidity: function payForTickets(uint128 _value) returns()
func (_BscGame *BscGameTransactor) PayForTickets(opts *bind.TransactOpts, _value *big.Int) (*types.Transaction, error) {
	return _BscGame.contract.Transact(opts, "payForTickets", _value)
}

// PayForTickets is a paid mutator transaction binding the contract method 0xbc553d8f.
//
// Solidity: function payForTickets(uint128 _value) returns()
func (_BscGame *BscGameSession) PayForTickets(_value *big.Int) (*types.Transaction, error) {
	return _BscGame.Contract.PayForTickets(&_BscGame.TransactOpts, _value)
}

// PayForTickets is a paid mutator transaction binding the contract method 0xbc553d8f.
//
// Solidity: function payForTickets(uint128 _value) returns()
func (_BscGame *BscGameTransactorSession) PayForTickets(_value *big.Int) (*types.Transaction, error) {
	return _BscGame.Contract.PayForTickets(&_BscGame.TransactOpts, _value)
}

// Register is a paid mutator transaction binding the contract method 0x4420e486.
//
// Solidity: function register(address ref_addr) returns()
func (_BscGame *BscGameTransactor) Register(opts *bind.TransactOpts, ref_addr common.Address) (*types.Transaction, error) {
	return _BscGame.contract.Transact(opts, "register", ref_addr)
}

// Register is a paid mutator transaction binding the contract method 0x4420e486.
//
// Solidity: function register(address ref_addr) returns()
func (_BscGame *BscGameSession) Register(ref_addr common.Address) (*types.Transaction, error) {
	return _BscGame.Contract.Register(&_BscGame.TransactOpts, ref_addr)
}

// Register is a paid mutator transaction binding the contract method 0x4420e486.
//
// Solidity: function register(address ref_addr) returns()
func (_BscGame *BscGameTransactorSession) Register(ref_addr common.Address) (*types.Transaction, error) {
	return _BscGame.Contract.Register(&_BscGame.TransactOpts, ref_addr)
}

// SetRound is a paid mutator transaction binding the contract method 0x6de38bec.
//
// Solidity: function setRound() returns(bool)
func (_BscGame *BscGameTransactor) SetRound(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BscGame.contract.Transact(opts, "setRound")
}

// SetRound is a paid mutator transaction binding the contract method 0x6de38bec.
//
// Solidity: function setRound() returns(bool)
func (_BscGame *BscGameSession) SetRound() (*types.Transaction, error) {
	return _BscGame.Contract.SetRound(&_BscGame.TransactOpts)
}

// SetRound is a paid mutator transaction binding the contract method 0x6de38bec.
//
// Solidity: function setRound() returns(bool)
func (_BscGame *BscGameTransactorSession) SetRound() (*types.Transaction, error) {
	return _BscGame.Contract.SetRound(&_BscGame.TransactOpts)
}

// SetSpend is a paid mutator transaction binding the contract method 0xb54f6121.
//
// Solidity: function setSpend(uint128 _value) returns(bool)
func (_BscGame *BscGameTransactor) SetSpend(opts *bind.TransactOpts, _value *big.Int) (*types.Transaction, error) {
	return _BscGame.contract.Transact(opts, "setSpend", _value)
}

// SetSpend is a paid mutator transaction binding the contract method 0xb54f6121.
//
// Solidity: function setSpend(uint128 _value) returns(bool)
func (_BscGame *BscGameSession) SetSpend(_value *big.Int) (*types.Transaction, error) {
	return _BscGame.Contract.SetSpend(&_BscGame.TransactOpts, _value)
}

// SetSpend is a paid mutator transaction binding the contract method 0xb54f6121.
//
// Solidity: function setSpend(uint128 _value) returns(bool)
func (_BscGame *BscGameTransactorSession) SetSpend(_value *big.Int) (*types.Transaction, error) {
	return _BscGame.Contract.SetSpend(&_BscGame.TransactOpts, _value)
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

// UserOut is a paid mutator transaction binding the contract method 0xb62a0b39.
//
// Solidity: function userOut(uint64 userId, uint32 _round) returns(bool)
func (_BscGame *BscGameTransactor) UserOut(opts *bind.TransactOpts, userId uint64, _round uint32) (*types.Transaction, error) {
	return _BscGame.contract.Transact(opts, "userOut", userId, _round)
}

// UserOut is a paid mutator transaction binding the contract method 0xb62a0b39.
//
// Solidity: function userOut(uint64 userId, uint32 _round) returns(bool)
func (_BscGame *BscGameSession) UserOut(userId uint64, _round uint32) (*types.Transaction, error) {
	return _BscGame.Contract.UserOut(&_BscGame.TransactOpts, userId, _round)
}

// UserOut is a paid mutator transaction binding the contract method 0xb62a0b39.
//
// Solidity: function userOut(uint64 userId, uint32 _round) returns(bool)
func (_BscGame *BscGameTransactorSession) UserOut(userId uint64, _round uint32) (*types.Transaction, error) {
	return _BscGame.Contract.UserOut(&_BscGame.TransactOpts, userId, _round)
}

// UserWithdraw is a paid mutator transaction binding the contract method 0x7ccf264d.
//
// Solidity: function userWithdraw(uint128 _value) returns()
func (_BscGame *BscGameTransactor) UserWithdraw(opts *bind.TransactOpts, _value *big.Int) (*types.Transaction, error) {
	return _BscGame.contract.Transact(opts, "userWithdraw", _value)
}

// UserWithdraw is a paid mutator transaction binding the contract method 0x7ccf264d.
//
// Solidity: function userWithdraw(uint128 _value) returns()
func (_BscGame *BscGameSession) UserWithdraw(_value *big.Int) (*types.Transaction, error) {
	return _BscGame.Contract.UserWithdraw(&_BscGame.TransactOpts, _value)
}

// UserWithdraw is a paid mutator transaction binding the contract method 0x7ccf264d.
//
// Solidity: function userWithdraw(uint128 _value) returns()
func (_BscGame *BscGameTransactorSession) UserWithdraw(_value *big.Int) (*types.Transaction, error) {
	return _BscGame.Contract.UserWithdraw(&_BscGame.TransactOpts, _value)
}

// BscGameBuyTicketLogIterator is returned from FilterBuyTicketLog and is used to iterate over the raw logs and unpacked data for BuyTicketLog events raised by the BscGame contract.
type BscGameBuyTicketLogIterator struct {
	Event *BscGameBuyTicketLog // Event containing the contract specifics and raw log

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
func (it *BscGameBuyTicketLogIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BscGameBuyTicketLog)
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
		it.Event = new(BscGameBuyTicketLog)
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
func (it *BscGameBuyTicketLogIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BscGameBuyTicketLogIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BscGameBuyTicketLog represents a BuyTicketLog event raised by the BscGame contract.
type BscGameBuyTicketLog struct {
	DoType    uint8
	Id        uint64
	Value     *big.Int
	GetTicket *big.Int
	Percent   uint32
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterBuyTicketLog is a free log retrieval operation binding the contract event 0x49cc93a02c7f8ee300b9532f8fa1682942f92d52b8be40d385a43f66873ffe26.
//
// Solidity: event buyTicketLog(uint8 doType, uint64 id, uint128 _value, uint128 getTicket, uint32 percent)
func (_BscGame *BscGameFilterer) FilterBuyTicketLog(opts *bind.FilterOpts) (*BscGameBuyTicketLogIterator, error) {

	logs, sub, err := _BscGame.contract.FilterLogs(opts, "buyTicketLog")
	if err != nil {
		return nil, err
	}
	return &BscGameBuyTicketLogIterator{contract: _BscGame.contract, event: "buyTicketLog", logs: logs, sub: sub}, nil
}

// WatchBuyTicketLog is a free log subscription operation binding the contract event 0x49cc93a02c7f8ee300b9532f8fa1682942f92d52b8be40d385a43f66873ffe26.
//
// Solidity: event buyTicketLog(uint8 doType, uint64 id, uint128 _value, uint128 getTicket, uint32 percent)
func (_BscGame *BscGameFilterer) WatchBuyTicketLog(opts *bind.WatchOpts, sink chan<- *BscGameBuyTicketLog) (event.Subscription, error) {

	logs, sub, err := _BscGame.contract.WatchLogs(opts, "buyTicketLog")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BscGameBuyTicketLog)
				if err := _BscGame.contract.UnpackLog(event, "buyTicketLog", log); err != nil {
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

// ParseBuyTicketLog is a log parse operation binding the contract event 0x49cc93a02c7f8ee300b9532f8fa1682942f92d52b8be40d385a43f66873ffe26.
//
// Solidity: event buyTicketLog(uint8 doType, uint64 id, uint128 _value, uint128 getTicket, uint32 percent)
func (_BscGame *BscGameFilterer) ParseBuyTicketLog(log types.Log) (*BscGameBuyTicketLog, error) {
	event := new(BscGameBuyTicketLog)
	if err := _BscGame.contract.UnpackLog(event, "buyTicketLog", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BscGameJoinLogIterator is returned from FilterJoinLog and is used to iterate over the raw logs and unpacked data for JoinLog events raised by the BscGame contract.
type BscGameJoinLogIterator struct {
	Event *BscGameJoinLog // Event containing the contract specifics and raw log

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
func (it *BscGameJoinLogIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BscGameJoinLog)
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
		it.Event = new(BscGameJoinLog)
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
func (it *BscGameJoinLogIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BscGameJoinLogIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BscGameJoinLog represents a JoinLog event raised by the BscGame contract.
type BscGameJoinLog struct {
	DoType uint8
	Id     uint64
	Value  *big.Int
	Round  uint32
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterJoinLog is a free log retrieval operation binding the contract event 0x3f0a3a2cbe132c424167bcd3a5cf80c5e7dca8ffbb4f72c72eb9ed2b6b1be526.
//
// Solidity: event joinLog(uint8 doType, uint64 id, uint128 _value, uint32 _round)
func (_BscGame *BscGameFilterer) FilterJoinLog(opts *bind.FilterOpts) (*BscGameJoinLogIterator, error) {

	logs, sub, err := _BscGame.contract.FilterLogs(opts, "joinLog")
	if err != nil {
		return nil, err
	}
	return &BscGameJoinLogIterator{contract: _BscGame.contract, event: "joinLog", logs: logs, sub: sub}, nil
}

// WatchJoinLog is a free log subscription operation binding the contract event 0x3f0a3a2cbe132c424167bcd3a5cf80c5e7dca8ffbb4f72c72eb9ed2b6b1be526.
//
// Solidity: event joinLog(uint8 doType, uint64 id, uint128 _value, uint32 _round)
func (_BscGame *BscGameFilterer) WatchJoinLog(opts *bind.WatchOpts, sink chan<- *BscGameJoinLog) (event.Subscription, error) {

	logs, sub, err := _BscGame.contract.WatchLogs(opts, "joinLog")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BscGameJoinLog)
				if err := _BscGame.contract.UnpackLog(event, "joinLog", log); err != nil {
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

// ParseJoinLog is a log parse operation binding the contract event 0x3f0a3a2cbe132c424167bcd3a5cf80c5e7dca8ffbb4f72c72eb9ed2b6b1be526.
//
// Solidity: event joinLog(uint8 doType, uint64 id, uint128 _value, uint32 _round)
func (_BscGame *BscGameFilterer) ParseJoinLog(log types.Log) (*BscGameJoinLog, error) {
	event := new(BscGameJoinLog)
	if err := _BscGame.contract.UnpackLog(event, "joinLog", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BscGameRegisterLogIterator is returned from FilterRegisterLog and is used to iterate over the raw logs and unpacked data for RegisterLog events raised by the BscGame contract.
type BscGameRegisterLogIterator struct {
	Event *BscGameRegisterLog // Event containing the contract specifics and raw log

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
func (it *BscGameRegisterLogIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BscGameRegisterLog)
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
		it.Event = new(BscGameRegisterLog)
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
func (it *BscGameRegisterLogIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BscGameRegisterLogIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BscGameRegisterLog represents a RegisterLog event raised by the BscGame contract.
type BscGameRegisterLog struct {
	DoType uint8
	Id     uint64
	Addr   common.Address
	RefId  uint64
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRegisterLog is a free log retrieval operation binding the contract event 0x7e43c07e66fa0b2486fc06807f3944ca195fe188fae0e7cfa9b768f6fc352dee.
//
// Solidity: event registerLog(uint8 doType, uint64 id, address indexed addr, uint64 refId)
func (_BscGame *BscGameFilterer) FilterRegisterLog(opts *bind.FilterOpts, addr []common.Address) (*BscGameRegisterLogIterator, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _BscGame.contract.FilterLogs(opts, "registerLog", addrRule)
	if err != nil {
		return nil, err
	}
	return &BscGameRegisterLogIterator{contract: _BscGame.contract, event: "registerLog", logs: logs, sub: sub}, nil
}

// WatchRegisterLog is a free log subscription operation binding the contract event 0x7e43c07e66fa0b2486fc06807f3944ca195fe188fae0e7cfa9b768f6fc352dee.
//
// Solidity: event registerLog(uint8 doType, uint64 id, address indexed addr, uint64 refId)
func (_BscGame *BscGameFilterer) WatchRegisterLog(opts *bind.WatchOpts, sink chan<- *BscGameRegisterLog, addr []common.Address) (event.Subscription, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _BscGame.contract.WatchLogs(opts, "registerLog", addrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BscGameRegisterLog)
				if err := _BscGame.contract.UnpackLog(event, "registerLog", log); err != nil {
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

// ParseRegisterLog is a log parse operation binding the contract event 0x7e43c07e66fa0b2486fc06807f3944ca195fe188fae0e7cfa9b768f6fc352dee.
//
// Solidity: event registerLog(uint8 doType, uint64 id, address indexed addr, uint64 refId)
func (_BscGame *BscGameFilterer) ParseRegisterLog(log types.Log) (*BscGameRegisterLog, error) {
	event := new(BscGameRegisterLog)
	if err := _BscGame.contract.UnpackLog(event, "registerLog", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BscGameUserGetLogIterator is returned from FilterUserGetLog and is used to iterate over the raw logs and unpacked data for UserGetLog events raised by the BscGame contract.
type BscGameUserGetLogIterator struct {
	Event *BscGameUserGetLog // Event containing the contract specifics and raw log

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
func (it *BscGameUserGetLogIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BscGameUserGetLog)
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
		it.Event = new(BscGameUserGetLog)
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
func (it *BscGameUserGetLogIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BscGameUserGetLogIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BscGameUserGetLog represents a UserGetLog event raised by the BscGame contract.
type BscGameUserGetLog struct {
	DoType uint8
	Id     uint64
	Value  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterUserGetLog is a free log retrieval operation binding the contract event 0xbae0e0291354dfeef36374d68e0b2cb53638cce668f7073a4a65f92a61557dd3.
//
// Solidity: event userGetLog(uint8 doType, uint64 id, uint128 _value)
func (_BscGame *BscGameFilterer) FilterUserGetLog(opts *bind.FilterOpts) (*BscGameUserGetLogIterator, error) {

	logs, sub, err := _BscGame.contract.FilterLogs(opts, "userGetLog")
	if err != nil {
		return nil, err
	}
	return &BscGameUserGetLogIterator{contract: _BscGame.contract, event: "userGetLog", logs: logs, sub: sub}, nil
}

// WatchUserGetLog is a free log subscription operation binding the contract event 0xbae0e0291354dfeef36374d68e0b2cb53638cce668f7073a4a65f92a61557dd3.
//
// Solidity: event userGetLog(uint8 doType, uint64 id, uint128 _value)
func (_BscGame *BscGameFilterer) WatchUserGetLog(opts *bind.WatchOpts, sink chan<- *BscGameUserGetLog) (event.Subscription, error) {

	logs, sub, err := _BscGame.contract.WatchLogs(opts, "userGetLog")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BscGameUserGetLog)
				if err := _BscGame.contract.UnpackLog(event, "userGetLog", log); err != nil {
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

// ParseUserGetLog is a log parse operation binding the contract event 0xbae0e0291354dfeef36374d68e0b2cb53638cce668f7073a4a65f92a61557dd3.
//
// Solidity: event userGetLog(uint8 doType, uint64 id, uint128 _value)
func (_BscGame *BscGameFilterer) ParseUserGetLog(log types.Log) (*BscGameUserGetLog, error) {
	event := new(BscGameUserGetLog)
	if err := _BscGame.contract.UnpackLog(event, "userGetLog", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
