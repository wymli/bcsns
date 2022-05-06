// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bcsns

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

// BcsnsMetaData contains all meta data concerning the Bcsns contract.
var BcsnsMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"send_uid\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"int64\",\"name\":\"server_msg_id\",\"type\":\"int64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"momentsPersistedEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"room_uid\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"send_uid\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"int64\",\"name\":\"server_msg_id\",\"type\":\"int64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"roomMsgPersistedEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"send_uid\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"recv_uid\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"int64\",\"name\":\"server_msg_id\",\"type\":\"int64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"userMsgPersistedEvent\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"send_uid\",\"type\":\"uint64\"},{\"internalType\":\"int64\",\"name\":\"server_msg_id\",\"type\":\"int64\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"PersistMoments\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"room_id\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"send_uid\",\"type\":\"uint64\"},{\"internalType\":\"int64\",\"name\":\"server_msg_id\",\"type\":\"int64\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"PersistRoomMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"send_uid\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"recv_uid\",\"type\":\"uint64\"},{\"internalType\":\"int64\",\"name\":\"server_msg_id\",\"type\":\"int64\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"PersistUserMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// BcsnsABI is the input ABI used to generate the binding from.
// Deprecated: Use BcsnsMetaData.ABI instead.
var BcsnsABI = BcsnsMetaData.ABI

// Bcsns is an auto generated Go binding around an Ethereum contract.
type Bcsns struct {
	BcsnsCaller     // Read-only binding to the contract
	BcsnsTransactor // Write-only binding to the contract
	BcsnsFilterer   // Log filterer for contract events
}

// BcsnsCaller is an auto generated read-only Go binding around an Ethereum contract.
type BcsnsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BcsnsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BcsnsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BcsnsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BcsnsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BcsnsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BcsnsSession struct {
	Contract     *Bcsns            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BcsnsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BcsnsCallerSession struct {
	Contract *BcsnsCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BcsnsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BcsnsTransactorSession struct {
	Contract     *BcsnsTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BcsnsRaw is an auto generated low-level Go binding around an Ethereum contract.
type BcsnsRaw struct {
	Contract *Bcsns // Generic contract binding to access the raw methods on
}

// BcsnsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BcsnsCallerRaw struct {
	Contract *BcsnsCaller // Generic read-only contract binding to access the raw methods on
}

// BcsnsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BcsnsTransactorRaw struct {
	Contract *BcsnsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBcsns creates a new instance of Bcsns, bound to a specific deployed contract.
func NewBcsns(address common.Address, backend bind.ContractBackend) (*Bcsns, error) {
	contract, err := bindBcsns(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Bcsns{BcsnsCaller: BcsnsCaller{contract: contract}, BcsnsTransactor: BcsnsTransactor{contract: contract}, BcsnsFilterer: BcsnsFilterer{contract: contract}}, nil
}

// NewBcsnsCaller creates a new read-only instance of Bcsns, bound to a specific deployed contract.
func NewBcsnsCaller(address common.Address, caller bind.ContractCaller) (*BcsnsCaller, error) {
	contract, err := bindBcsns(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BcsnsCaller{contract: contract}, nil
}

// NewBcsnsTransactor creates a new write-only instance of Bcsns, bound to a specific deployed contract.
func NewBcsnsTransactor(address common.Address, transactor bind.ContractTransactor) (*BcsnsTransactor, error) {
	contract, err := bindBcsns(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BcsnsTransactor{contract: contract}, nil
}

// NewBcsnsFilterer creates a new log filterer instance of Bcsns, bound to a specific deployed contract.
func NewBcsnsFilterer(address common.Address, filterer bind.ContractFilterer) (*BcsnsFilterer, error) {
	contract, err := bindBcsns(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BcsnsFilterer{contract: contract}, nil
}

// bindBcsns binds a generic wrapper to an already deployed contract.
func bindBcsns(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BcsnsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bcsns *BcsnsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bcsns.Contract.BcsnsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bcsns *BcsnsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bcsns.Contract.BcsnsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bcsns *BcsnsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bcsns.Contract.BcsnsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bcsns *BcsnsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bcsns.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bcsns *BcsnsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bcsns.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bcsns *BcsnsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bcsns.Contract.contract.Transact(opts, method, params...)
}

// PersistMoments is a paid mutator transaction binding the contract method 0xfd23a5ba.
//
// Solidity: function PersistMoments(uint64 send_uid, int64 server_msg_id, bytes message) returns()
func (_Bcsns *BcsnsTransactor) PersistMoments(opts *bind.TransactOpts, send_uid uint64, server_msg_id int64, message []byte) (*types.Transaction, error) {
	return _Bcsns.contract.Transact(opts, "PersistMoments", send_uid, server_msg_id, message)
}

// PersistMoments is a paid mutator transaction binding the contract method 0xfd23a5ba.
//
// Solidity: function PersistMoments(uint64 send_uid, int64 server_msg_id, bytes message) returns()
func (_Bcsns *BcsnsSession) PersistMoments(send_uid uint64, server_msg_id int64, message []byte) (*types.Transaction, error) {
	return _Bcsns.Contract.PersistMoments(&_Bcsns.TransactOpts, send_uid, server_msg_id, message)
}

// PersistMoments is a paid mutator transaction binding the contract method 0xfd23a5ba.
//
// Solidity: function PersistMoments(uint64 send_uid, int64 server_msg_id, bytes message) returns()
func (_Bcsns *BcsnsTransactorSession) PersistMoments(send_uid uint64, server_msg_id int64, message []byte) (*types.Transaction, error) {
	return _Bcsns.Contract.PersistMoments(&_Bcsns.TransactOpts, send_uid, server_msg_id, message)
}

// PersistRoomMessage is a paid mutator transaction binding the contract method 0x44bb36c0.
//
// Solidity: function PersistRoomMessage(uint64 room_id, uint64 send_uid, int64 server_msg_id, bytes message) returns()
func (_Bcsns *BcsnsTransactor) PersistRoomMessage(opts *bind.TransactOpts, room_id uint64, send_uid uint64, server_msg_id int64, message []byte) (*types.Transaction, error) {
	return _Bcsns.contract.Transact(opts, "PersistRoomMessage", room_id, send_uid, server_msg_id, message)
}

// PersistRoomMessage is a paid mutator transaction binding the contract method 0x44bb36c0.
//
// Solidity: function PersistRoomMessage(uint64 room_id, uint64 send_uid, int64 server_msg_id, bytes message) returns()
func (_Bcsns *BcsnsSession) PersistRoomMessage(room_id uint64, send_uid uint64, server_msg_id int64, message []byte) (*types.Transaction, error) {
	return _Bcsns.Contract.PersistRoomMessage(&_Bcsns.TransactOpts, room_id, send_uid, server_msg_id, message)
}

// PersistRoomMessage is a paid mutator transaction binding the contract method 0x44bb36c0.
//
// Solidity: function PersistRoomMessage(uint64 room_id, uint64 send_uid, int64 server_msg_id, bytes message) returns()
func (_Bcsns *BcsnsTransactorSession) PersistRoomMessage(room_id uint64, send_uid uint64, server_msg_id int64, message []byte) (*types.Transaction, error) {
	return _Bcsns.Contract.PersistRoomMessage(&_Bcsns.TransactOpts, room_id, send_uid, server_msg_id, message)
}

// PersistUserMessage is a paid mutator transaction binding the contract method 0xdaf023ef.
//
// Solidity: function PersistUserMessage(uint64 send_uid, uint64 recv_uid, int64 server_msg_id, bytes message) returns()
func (_Bcsns *BcsnsTransactor) PersistUserMessage(opts *bind.TransactOpts, send_uid uint64, recv_uid uint64, server_msg_id int64, message []byte) (*types.Transaction, error) {
	return _Bcsns.contract.Transact(opts, "PersistUserMessage", send_uid, recv_uid, server_msg_id, message)
}

// PersistUserMessage is a paid mutator transaction binding the contract method 0xdaf023ef.
//
// Solidity: function PersistUserMessage(uint64 send_uid, uint64 recv_uid, int64 server_msg_id, bytes message) returns()
func (_Bcsns *BcsnsSession) PersistUserMessage(send_uid uint64, recv_uid uint64, server_msg_id int64, message []byte) (*types.Transaction, error) {
	return _Bcsns.Contract.PersistUserMessage(&_Bcsns.TransactOpts, send_uid, recv_uid, server_msg_id, message)
}

// PersistUserMessage is a paid mutator transaction binding the contract method 0xdaf023ef.
//
// Solidity: function PersistUserMessage(uint64 send_uid, uint64 recv_uid, int64 server_msg_id, bytes message) returns()
func (_Bcsns *BcsnsTransactorSession) PersistUserMessage(send_uid uint64, recv_uid uint64, server_msg_id int64, message []byte) (*types.Transaction, error) {
	return _Bcsns.Contract.PersistUserMessage(&_Bcsns.TransactOpts, send_uid, recv_uid, server_msg_id, message)
}

// BcsnsMomentsPersistedEventIterator is returned from FilterMomentsPersistedEvent and is used to iterate over the raw logs and unpacked data for MomentsPersistedEvent events raised by the Bcsns contract.
type BcsnsMomentsPersistedEventIterator struct {
	Event *BcsnsMomentsPersistedEvent // Event containing the contract specifics and raw log

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
func (it *BcsnsMomentsPersistedEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BcsnsMomentsPersistedEvent)
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
		it.Event = new(BcsnsMomentsPersistedEvent)
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
func (it *BcsnsMomentsPersistedEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BcsnsMomentsPersistedEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BcsnsMomentsPersistedEvent represents a MomentsPersistedEvent event raised by the Bcsns contract.
type BcsnsMomentsPersistedEvent struct {
	SendUid     uint64
	ServerMsgId int64
	Message     []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterMomentsPersistedEvent is a free log retrieval operation binding the contract event 0x0e29dcc2bbcf38e4d69b858d462a24135d6d0488ef0173fb8fcd279c3c1f0dae.
//
// Solidity: event momentsPersistedEvent(uint64 indexed send_uid, int64 server_msg_id, bytes message)
func (_Bcsns *BcsnsFilterer) FilterMomentsPersistedEvent(opts *bind.FilterOpts, send_uid []uint64) (*BcsnsMomentsPersistedEventIterator, error) {

	var send_uidRule []interface{}
	for _, send_uidItem := range send_uid {
		send_uidRule = append(send_uidRule, send_uidItem)
	}

	logs, sub, err := _Bcsns.contract.FilterLogs(opts, "momentsPersistedEvent", send_uidRule)
	if err != nil {
		return nil, err
	}
	return &BcsnsMomentsPersistedEventIterator{contract: _Bcsns.contract, event: "momentsPersistedEvent", logs: logs, sub: sub}, nil
}

// WatchMomentsPersistedEvent is a free log subscription operation binding the contract event 0x0e29dcc2bbcf38e4d69b858d462a24135d6d0488ef0173fb8fcd279c3c1f0dae.
//
// Solidity: event momentsPersistedEvent(uint64 indexed send_uid, int64 server_msg_id, bytes message)
func (_Bcsns *BcsnsFilterer) WatchMomentsPersistedEvent(opts *bind.WatchOpts, sink chan<- *BcsnsMomentsPersistedEvent, send_uid []uint64) (event.Subscription, error) {

	var send_uidRule []interface{}
	for _, send_uidItem := range send_uid {
		send_uidRule = append(send_uidRule, send_uidItem)
	}

	logs, sub, err := _Bcsns.contract.WatchLogs(opts, "momentsPersistedEvent", send_uidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BcsnsMomentsPersistedEvent)
				if err := _Bcsns.contract.UnpackLog(event, "momentsPersistedEvent", log); err != nil {
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

// ParseMomentsPersistedEvent is a log parse operation binding the contract event 0x0e29dcc2bbcf38e4d69b858d462a24135d6d0488ef0173fb8fcd279c3c1f0dae.
//
// Solidity: event momentsPersistedEvent(uint64 indexed send_uid, int64 server_msg_id, bytes message)
func (_Bcsns *BcsnsFilterer) ParseMomentsPersistedEvent(log types.Log) (*BcsnsMomentsPersistedEvent, error) {
	event := new(BcsnsMomentsPersistedEvent)
	if err := _Bcsns.contract.UnpackLog(event, "momentsPersistedEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BcsnsRoomMsgPersistedEventIterator is returned from FilterRoomMsgPersistedEvent and is used to iterate over the raw logs and unpacked data for RoomMsgPersistedEvent events raised by the Bcsns contract.
type BcsnsRoomMsgPersistedEventIterator struct {
	Event *BcsnsRoomMsgPersistedEvent // Event containing the contract specifics and raw log

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
func (it *BcsnsRoomMsgPersistedEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BcsnsRoomMsgPersistedEvent)
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
		it.Event = new(BcsnsRoomMsgPersistedEvent)
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
func (it *BcsnsRoomMsgPersistedEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BcsnsRoomMsgPersistedEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BcsnsRoomMsgPersistedEvent represents a RoomMsgPersistedEvent event raised by the Bcsns contract.
type BcsnsRoomMsgPersistedEvent struct {
	RoomUid     uint64
	SendUid     uint64
	ServerMsgId int64
	Message     []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRoomMsgPersistedEvent is a free log retrieval operation binding the contract event 0xc0b8dc5582fb0acd66048cdd462f7862cb7d24b0d43fa4a72e079e6dda79fe6a.
//
// Solidity: event roomMsgPersistedEvent(uint64 indexed room_uid, uint64 send_uid, int64 server_msg_id, bytes message)
func (_Bcsns *BcsnsFilterer) FilterRoomMsgPersistedEvent(opts *bind.FilterOpts, room_uid []uint64) (*BcsnsRoomMsgPersistedEventIterator, error) {

	var room_uidRule []interface{}
	for _, room_uidItem := range room_uid {
		room_uidRule = append(room_uidRule, room_uidItem)
	}

	logs, sub, err := _Bcsns.contract.FilterLogs(opts, "roomMsgPersistedEvent", room_uidRule)
	if err != nil {
		return nil, err
	}
	return &BcsnsRoomMsgPersistedEventIterator{contract: _Bcsns.contract, event: "roomMsgPersistedEvent", logs: logs, sub: sub}, nil
}

// WatchRoomMsgPersistedEvent is a free log subscription operation binding the contract event 0xc0b8dc5582fb0acd66048cdd462f7862cb7d24b0d43fa4a72e079e6dda79fe6a.
//
// Solidity: event roomMsgPersistedEvent(uint64 indexed room_uid, uint64 send_uid, int64 server_msg_id, bytes message)
func (_Bcsns *BcsnsFilterer) WatchRoomMsgPersistedEvent(opts *bind.WatchOpts, sink chan<- *BcsnsRoomMsgPersistedEvent, room_uid []uint64) (event.Subscription, error) {

	var room_uidRule []interface{}
	for _, room_uidItem := range room_uid {
		room_uidRule = append(room_uidRule, room_uidItem)
	}

	logs, sub, err := _Bcsns.contract.WatchLogs(opts, "roomMsgPersistedEvent", room_uidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BcsnsRoomMsgPersistedEvent)
				if err := _Bcsns.contract.UnpackLog(event, "roomMsgPersistedEvent", log); err != nil {
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

// ParseRoomMsgPersistedEvent is a log parse operation binding the contract event 0xc0b8dc5582fb0acd66048cdd462f7862cb7d24b0d43fa4a72e079e6dda79fe6a.
//
// Solidity: event roomMsgPersistedEvent(uint64 indexed room_uid, uint64 send_uid, int64 server_msg_id, bytes message)
func (_Bcsns *BcsnsFilterer) ParseRoomMsgPersistedEvent(log types.Log) (*BcsnsRoomMsgPersistedEvent, error) {
	event := new(BcsnsRoomMsgPersistedEvent)
	if err := _Bcsns.contract.UnpackLog(event, "roomMsgPersistedEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BcsnsUserMsgPersistedEventIterator is returned from FilterUserMsgPersistedEvent and is used to iterate over the raw logs and unpacked data for UserMsgPersistedEvent events raised by the Bcsns contract.
type BcsnsUserMsgPersistedEventIterator struct {
	Event *BcsnsUserMsgPersistedEvent // Event containing the contract specifics and raw log

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
func (it *BcsnsUserMsgPersistedEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BcsnsUserMsgPersistedEvent)
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
		it.Event = new(BcsnsUserMsgPersistedEvent)
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
func (it *BcsnsUserMsgPersistedEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BcsnsUserMsgPersistedEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BcsnsUserMsgPersistedEvent represents a UserMsgPersistedEvent event raised by the Bcsns contract.
type BcsnsUserMsgPersistedEvent struct {
	SendUid     uint64
	RecvUid     uint64
	ServerMsgId int64
	Message     []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterUserMsgPersistedEvent is a free log retrieval operation binding the contract event 0xf15798036872d30bde5df397f6d0835e6b9cfb538df80313b372ee5e56820553.
//
// Solidity: event userMsgPersistedEvent(uint64 indexed send_uid, uint64 indexed recv_uid, int64 server_msg_id, bytes message)
func (_Bcsns *BcsnsFilterer) FilterUserMsgPersistedEvent(opts *bind.FilterOpts, send_uid []uint64, recv_uid []uint64) (*BcsnsUserMsgPersistedEventIterator, error) {

	var send_uidRule []interface{}
	for _, send_uidItem := range send_uid {
		send_uidRule = append(send_uidRule, send_uidItem)
	}
	var recv_uidRule []interface{}
	for _, recv_uidItem := range recv_uid {
		recv_uidRule = append(recv_uidRule, recv_uidItem)
	}

	logs, sub, err := _Bcsns.contract.FilterLogs(opts, "userMsgPersistedEvent", send_uidRule, recv_uidRule)
	if err != nil {
		return nil, err
	}
	return &BcsnsUserMsgPersistedEventIterator{contract: _Bcsns.contract, event: "userMsgPersistedEvent", logs: logs, sub: sub}, nil
}

// WatchUserMsgPersistedEvent is a free log subscription operation binding the contract event 0xf15798036872d30bde5df397f6d0835e6b9cfb538df80313b372ee5e56820553.
//
// Solidity: event userMsgPersistedEvent(uint64 indexed send_uid, uint64 indexed recv_uid, int64 server_msg_id, bytes message)
func (_Bcsns *BcsnsFilterer) WatchUserMsgPersistedEvent(opts *bind.WatchOpts, sink chan<- *BcsnsUserMsgPersistedEvent, send_uid []uint64, recv_uid []uint64) (event.Subscription, error) {

	var send_uidRule []interface{}
	for _, send_uidItem := range send_uid {
		send_uidRule = append(send_uidRule, send_uidItem)
	}
	var recv_uidRule []interface{}
	for _, recv_uidItem := range recv_uid {
		recv_uidRule = append(recv_uidRule, recv_uidItem)
	}

	logs, sub, err := _Bcsns.contract.WatchLogs(opts, "userMsgPersistedEvent", send_uidRule, recv_uidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BcsnsUserMsgPersistedEvent)
				if err := _Bcsns.contract.UnpackLog(event, "userMsgPersistedEvent", log); err != nil {
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

// ParseUserMsgPersistedEvent is a log parse operation binding the contract event 0xf15798036872d30bde5df397f6d0835e6b9cfb538df80313b372ee5e56820553.
//
// Solidity: event userMsgPersistedEvent(uint64 indexed send_uid, uint64 indexed recv_uid, int64 server_msg_id, bytes message)
func (_Bcsns *BcsnsFilterer) ParseUserMsgPersistedEvent(log types.Log) (*BcsnsUserMsgPersistedEvent, error) {
	event := new(BcsnsUserMsgPersistedEvent)
	if err := _Bcsns.contract.UnpackLog(event, "userMsgPersistedEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
