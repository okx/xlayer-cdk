// Code generated by mockery. DO NOT EDIT.

package sync

import (
	context "context"
	big "math/big"

	common "github.com/ethereum/go-ethereum/common"

	ethereum "github.com/ethereum/go-ethereum"

	mock "github.com/stretchr/testify/mock"

	types "github.com/ethereum/go-ethereum/core/types"
)

// L2Mock is an autogenerated mock type for the EthClienter type
type L2Mock struct {
	mock.Mock
}

type L2Mock_Expecter struct {
	mock *mock.Mock
}

func (_m *L2Mock) EXPECT() *L2Mock_Expecter {
	return &L2Mock_Expecter{mock: &_m.Mock}
}

// BlockByHash provides a mock function with given fields: ctx, hash
func (_m *L2Mock) BlockByHash(ctx context.Context, hash common.Hash) (*types.Block, error) {
	ret := _m.Called(ctx, hash)

	if len(ret) == 0 {
		panic("no return value specified for BlockByHash")
	}

	var r0 *types.Block
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, common.Hash) (*types.Block, error)); ok {
		return rf(ctx, hash)
	}
	if rf, ok := ret.Get(0).(func(context.Context, common.Hash) *types.Block); ok {
		r0 = rf(ctx, hash)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Block)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, common.Hash) error); ok {
		r1 = rf(ctx, hash)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// L2Mock_BlockByHash_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'BlockByHash'
type L2Mock_BlockByHash_Call struct {
	*mock.Call
}

// BlockByHash is a helper method to define mock.On call
//   - ctx context.Context
//   - hash common.Hash
func (_e *L2Mock_Expecter) BlockByHash(ctx interface{}, hash interface{}) *L2Mock_BlockByHash_Call {
	return &L2Mock_BlockByHash_Call{Call: _e.mock.On("BlockByHash", ctx, hash)}
}

func (_c *L2Mock_BlockByHash_Call) Run(run func(ctx context.Context, hash common.Hash)) *L2Mock_BlockByHash_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(common.Hash))
	})
	return _c
}

func (_c *L2Mock_BlockByHash_Call) Return(_a0 *types.Block, _a1 error) *L2Mock_BlockByHash_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *L2Mock_BlockByHash_Call) RunAndReturn(run func(context.Context, common.Hash) (*types.Block, error)) *L2Mock_BlockByHash_Call {
	_c.Call.Return(run)
	return _c
}

// BlockByNumber provides a mock function with given fields: ctx, number
func (_m *L2Mock) BlockByNumber(ctx context.Context, number *big.Int) (*types.Block, error) {
	ret := _m.Called(ctx, number)

	if len(ret) == 0 {
		panic("no return value specified for BlockByNumber")
	}

	var r0 *types.Block
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *big.Int) (*types.Block, error)); ok {
		return rf(ctx, number)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *big.Int) *types.Block); ok {
		r0 = rf(ctx, number)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Block)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *big.Int) error); ok {
		r1 = rf(ctx, number)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// L2Mock_BlockByNumber_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'BlockByNumber'
type L2Mock_BlockByNumber_Call struct {
	*mock.Call
}

// BlockByNumber is a helper method to define mock.On call
//   - ctx context.Context
//   - number *big.Int
func (_e *L2Mock_Expecter) BlockByNumber(ctx interface{}, number interface{}) *L2Mock_BlockByNumber_Call {
	return &L2Mock_BlockByNumber_Call{Call: _e.mock.On("BlockByNumber", ctx, number)}
}

func (_c *L2Mock_BlockByNumber_Call) Run(run func(ctx context.Context, number *big.Int)) *L2Mock_BlockByNumber_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*big.Int))
	})
	return _c
}

func (_c *L2Mock_BlockByNumber_Call) Return(_a0 *types.Block, _a1 error) *L2Mock_BlockByNumber_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *L2Mock_BlockByNumber_Call) RunAndReturn(run func(context.Context, *big.Int) (*types.Block, error)) *L2Mock_BlockByNumber_Call {
	_c.Call.Return(run)
	return _c
}

// BlockNumber provides a mock function with given fields: ctx
func (_m *L2Mock) BlockNumber(ctx context.Context) (uint64, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for BlockNumber")
	}

	var r0 uint64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (uint64, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) uint64); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// L2Mock_BlockNumber_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'BlockNumber'
type L2Mock_BlockNumber_Call struct {
	*mock.Call
}

// BlockNumber is a helper method to define mock.On call
//   - ctx context.Context
func (_e *L2Mock_Expecter) BlockNumber(ctx interface{}) *L2Mock_BlockNumber_Call {
	return &L2Mock_BlockNumber_Call{Call: _e.mock.On("BlockNumber", ctx)}
}

func (_c *L2Mock_BlockNumber_Call) Run(run func(ctx context.Context)) *L2Mock_BlockNumber_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *L2Mock_BlockNumber_Call) Return(_a0 uint64, _a1 error) *L2Mock_BlockNumber_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *L2Mock_BlockNumber_Call) RunAndReturn(run func(context.Context) (uint64, error)) *L2Mock_BlockNumber_Call {
	_c.Call.Return(run)
	return _c
}

// CallContract provides a mock function with given fields: ctx, call, blockNumber
func (_m *L2Mock) CallContract(ctx context.Context, call ethereum.CallMsg, blockNumber *big.Int) ([]byte, error) {
	ret := _m.Called(ctx, call, blockNumber)

	if len(ret) == 0 {
		panic("no return value specified for CallContract")
	}

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, ethereum.CallMsg, *big.Int) ([]byte, error)); ok {
		return rf(ctx, call, blockNumber)
	}
	if rf, ok := ret.Get(0).(func(context.Context, ethereum.CallMsg, *big.Int) []byte); ok {
		r0 = rf(ctx, call, blockNumber)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, ethereum.CallMsg, *big.Int) error); ok {
		r1 = rf(ctx, call, blockNumber)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// L2Mock_CallContract_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CallContract'
type L2Mock_CallContract_Call struct {
	*mock.Call
}

// CallContract is a helper method to define mock.On call
//   - ctx context.Context
//   - call ethereum.CallMsg
//   - blockNumber *big.Int
func (_e *L2Mock_Expecter) CallContract(ctx interface{}, call interface{}, blockNumber interface{}) *L2Mock_CallContract_Call {
	return &L2Mock_CallContract_Call{Call: _e.mock.On("CallContract", ctx, call, blockNumber)}
}

func (_c *L2Mock_CallContract_Call) Run(run func(ctx context.Context, call ethereum.CallMsg, blockNumber *big.Int)) *L2Mock_CallContract_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(ethereum.CallMsg), args[2].(*big.Int))
	})
	return _c
}

func (_c *L2Mock_CallContract_Call) Return(_a0 []byte, _a1 error) *L2Mock_CallContract_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *L2Mock_CallContract_Call) RunAndReturn(run func(context.Context, ethereum.CallMsg, *big.Int) ([]byte, error)) *L2Mock_CallContract_Call {
	_c.Call.Return(run)
	return _c
}

// CodeAt provides a mock function with given fields: ctx, contract, blockNumber
func (_m *L2Mock) CodeAt(ctx context.Context, contract common.Address, blockNumber *big.Int) ([]byte, error) {
	ret := _m.Called(ctx, contract, blockNumber)

	if len(ret) == 0 {
		panic("no return value specified for CodeAt")
	}

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, common.Address, *big.Int) ([]byte, error)); ok {
		return rf(ctx, contract, blockNumber)
	}
	if rf, ok := ret.Get(0).(func(context.Context, common.Address, *big.Int) []byte); ok {
		r0 = rf(ctx, contract, blockNumber)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, common.Address, *big.Int) error); ok {
		r1 = rf(ctx, contract, blockNumber)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// L2Mock_CodeAt_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CodeAt'
type L2Mock_CodeAt_Call struct {
	*mock.Call
}

// CodeAt is a helper method to define mock.On call
//   - ctx context.Context
//   - contract common.Address
//   - blockNumber *big.Int
func (_e *L2Mock_Expecter) CodeAt(ctx interface{}, contract interface{}, blockNumber interface{}) *L2Mock_CodeAt_Call {
	return &L2Mock_CodeAt_Call{Call: _e.mock.On("CodeAt", ctx, contract, blockNumber)}
}

func (_c *L2Mock_CodeAt_Call) Run(run func(ctx context.Context, contract common.Address, blockNumber *big.Int)) *L2Mock_CodeAt_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(common.Address), args[2].(*big.Int))
	})
	return _c
}

func (_c *L2Mock_CodeAt_Call) Return(_a0 []byte, _a1 error) *L2Mock_CodeAt_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *L2Mock_CodeAt_Call) RunAndReturn(run func(context.Context, common.Address, *big.Int) ([]byte, error)) *L2Mock_CodeAt_Call {
	_c.Call.Return(run)
	return _c
}

// EstimateGas provides a mock function with given fields: ctx, call
func (_m *L2Mock) EstimateGas(ctx context.Context, call ethereum.CallMsg) (uint64, error) {
	ret := _m.Called(ctx, call)

	if len(ret) == 0 {
		panic("no return value specified for EstimateGas")
	}

	var r0 uint64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, ethereum.CallMsg) (uint64, error)); ok {
		return rf(ctx, call)
	}
	if rf, ok := ret.Get(0).(func(context.Context, ethereum.CallMsg) uint64); ok {
		r0 = rf(ctx, call)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, ethereum.CallMsg) error); ok {
		r1 = rf(ctx, call)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// L2Mock_EstimateGas_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'EstimateGas'
type L2Mock_EstimateGas_Call struct {
	*mock.Call
}

// EstimateGas is a helper method to define mock.On call
//   - ctx context.Context
//   - call ethereum.CallMsg
func (_e *L2Mock_Expecter) EstimateGas(ctx interface{}, call interface{}) *L2Mock_EstimateGas_Call {
	return &L2Mock_EstimateGas_Call{Call: _e.mock.On("EstimateGas", ctx, call)}
}

func (_c *L2Mock_EstimateGas_Call) Run(run func(ctx context.Context, call ethereum.CallMsg)) *L2Mock_EstimateGas_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(ethereum.CallMsg))
	})
	return _c
}

func (_c *L2Mock_EstimateGas_Call) Return(_a0 uint64, _a1 error) *L2Mock_EstimateGas_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *L2Mock_EstimateGas_Call) RunAndReturn(run func(context.Context, ethereum.CallMsg) (uint64, error)) *L2Mock_EstimateGas_Call {
	_c.Call.Return(run)
	return _c
}

// FilterLogs provides a mock function with given fields: ctx, q
func (_m *L2Mock) FilterLogs(ctx context.Context, q ethereum.FilterQuery) ([]types.Log, error) {
	ret := _m.Called(ctx, q)

	if len(ret) == 0 {
		panic("no return value specified for FilterLogs")
	}

	var r0 []types.Log
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, ethereum.FilterQuery) ([]types.Log, error)); ok {
		return rf(ctx, q)
	}
	if rf, ok := ret.Get(0).(func(context.Context, ethereum.FilterQuery) []types.Log); ok {
		r0 = rf(ctx, q)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]types.Log)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, ethereum.FilterQuery) error); ok {
		r1 = rf(ctx, q)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// L2Mock_FilterLogs_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FilterLogs'
type L2Mock_FilterLogs_Call struct {
	*mock.Call
}

// FilterLogs is a helper method to define mock.On call
//   - ctx context.Context
//   - q ethereum.FilterQuery
func (_e *L2Mock_Expecter) FilterLogs(ctx interface{}, q interface{}) *L2Mock_FilterLogs_Call {
	return &L2Mock_FilterLogs_Call{Call: _e.mock.On("FilterLogs", ctx, q)}
}

func (_c *L2Mock_FilterLogs_Call) Run(run func(ctx context.Context, q ethereum.FilterQuery)) *L2Mock_FilterLogs_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(ethereum.FilterQuery))
	})
	return _c
}

func (_c *L2Mock_FilterLogs_Call) Return(_a0 []types.Log, _a1 error) *L2Mock_FilterLogs_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *L2Mock_FilterLogs_Call) RunAndReturn(run func(context.Context, ethereum.FilterQuery) ([]types.Log, error)) *L2Mock_FilterLogs_Call {
	_c.Call.Return(run)
	return _c
}

// HeaderByHash provides a mock function with given fields: ctx, hash
func (_m *L2Mock) HeaderByHash(ctx context.Context, hash common.Hash) (*types.Header, error) {
	ret := _m.Called(ctx, hash)

	if len(ret) == 0 {
		panic("no return value specified for HeaderByHash")
	}

	var r0 *types.Header
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, common.Hash) (*types.Header, error)); ok {
		return rf(ctx, hash)
	}
	if rf, ok := ret.Get(0).(func(context.Context, common.Hash) *types.Header); ok {
		r0 = rf(ctx, hash)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Header)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, common.Hash) error); ok {
		r1 = rf(ctx, hash)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// L2Mock_HeaderByHash_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'HeaderByHash'
type L2Mock_HeaderByHash_Call struct {
	*mock.Call
}

// HeaderByHash is a helper method to define mock.On call
//   - ctx context.Context
//   - hash common.Hash
func (_e *L2Mock_Expecter) HeaderByHash(ctx interface{}, hash interface{}) *L2Mock_HeaderByHash_Call {
	return &L2Mock_HeaderByHash_Call{Call: _e.mock.On("HeaderByHash", ctx, hash)}
}

func (_c *L2Mock_HeaderByHash_Call) Run(run func(ctx context.Context, hash common.Hash)) *L2Mock_HeaderByHash_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(common.Hash))
	})
	return _c
}

func (_c *L2Mock_HeaderByHash_Call) Return(_a0 *types.Header, _a1 error) *L2Mock_HeaderByHash_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *L2Mock_HeaderByHash_Call) RunAndReturn(run func(context.Context, common.Hash) (*types.Header, error)) *L2Mock_HeaderByHash_Call {
	_c.Call.Return(run)
	return _c
}

// HeaderByNumber provides a mock function with given fields: ctx, number
func (_m *L2Mock) HeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error) {
	ret := _m.Called(ctx, number)

	if len(ret) == 0 {
		panic("no return value specified for HeaderByNumber")
	}

	var r0 *types.Header
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *big.Int) (*types.Header, error)); ok {
		return rf(ctx, number)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *big.Int) *types.Header); ok {
		r0 = rf(ctx, number)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Header)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *big.Int) error); ok {
		r1 = rf(ctx, number)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// L2Mock_HeaderByNumber_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'HeaderByNumber'
type L2Mock_HeaderByNumber_Call struct {
	*mock.Call
}

// HeaderByNumber is a helper method to define mock.On call
//   - ctx context.Context
//   - number *big.Int
func (_e *L2Mock_Expecter) HeaderByNumber(ctx interface{}, number interface{}) *L2Mock_HeaderByNumber_Call {
	return &L2Mock_HeaderByNumber_Call{Call: _e.mock.On("HeaderByNumber", ctx, number)}
}

func (_c *L2Mock_HeaderByNumber_Call) Run(run func(ctx context.Context, number *big.Int)) *L2Mock_HeaderByNumber_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*big.Int))
	})
	return _c
}

func (_c *L2Mock_HeaderByNumber_Call) Return(_a0 *types.Header, _a1 error) *L2Mock_HeaderByNumber_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *L2Mock_HeaderByNumber_Call) RunAndReturn(run func(context.Context, *big.Int) (*types.Header, error)) *L2Mock_HeaderByNumber_Call {
	_c.Call.Return(run)
	return _c
}

// PendingCodeAt provides a mock function with given fields: ctx, account
func (_m *L2Mock) PendingCodeAt(ctx context.Context, account common.Address) ([]byte, error) {
	ret := _m.Called(ctx, account)

	if len(ret) == 0 {
		panic("no return value specified for PendingCodeAt")
	}

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, common.Address) ([]byte, error)); ok {
		return rf(ctx, account)
	}
	if rf, ok := ret.Get(0).(func(context.Context, common.Address) []byte); ok {
		r0 = rf(ctx, account)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, common.Address) error); ok {
		r1 = rf(ctx, account)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// L2Mock_PendingCodeAt_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PendingCodeAt'
type L2Mock_PendingCodeAt_Call struct {
	*mock.Call
}

// PendingCodeAt is a helper method to define mock.On call
//   - ctx context.Context
//   - account common.Address
func (_e *L2Mock_Expecter) PendingCodeAt(ctx interface{}, account interface{}) *L2Mock_PendingCodeAt_Call {
	return &L2Mock_PendingCodeAt_Call{Call: _e.mock.On("PendingCodeAt", ctx, account)}
}

func (_c *L2Mock_PendingCodeAt_Call) Run(run func(ctx context.Context, account common.Address)) *L2Mock_PendingCodeAt_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(common.Address))
	})
	return _c
}

func (_c *L2Mock_PendingCodeAt_Call) Return(_a0 []byte, _a1 error) *L2Mock_PendingCodeAt_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *L2Mock_PendingCodeAt_Call) RunAndReturn(run func(context.Context, common.Address) ([]byte, error)) *L2Mock_PendingCodeAt_Call {
	_c.Call.Return(run)
	return _c
}

// PendingNonceAt provides a mock function with given fields: ctx, account
func (_m *L2Mock) PendingNonceAt(ctx context.Context, account common.Address) (uint64, error) {
	ret := _m.Called(ctx, account)

	if len(ret) == 0 {
		panic("no return value specified for PendingNonceAt")
	}

	var r0 uint64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, common.Address) (uint64, error)); ok {
		return rf(ctx, account)
	}
	if rf, ok := ret.Get(0).(func(context.Context, common.Address) uint64); ok {
		r0 = rf(ctx, account)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, common.Address) error); ok {
		r1 = rf(ctx, account)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// L2Mock_PendingNonceAt_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PendingNonceAt'
type L2Mock_PendingNonceAt_Call struct {
	*mock.Call
}

// PendingNonceAt is a helper method to define mock.On call
//   - ctx context.Context
//   - account common.Address
func (_e *L2Mock_Expecter) PendingNonceAt(ctx interface{}, account interface{}) *L2Mock_PendingNonceAt_Call {
	return &L2Mock_PendingNonceAt_Call{Call: _e.mock.On("PendingNonceAt", ctx, account)}
}

func (_c *L2Mock_PendingNonceAt_Call) Run(run func(ctx context.Context, account common.Address)) *L2Mock_PendingNonceAt_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(common.Address))
	})
	return _c
}

func (_c *L2Mock_PendingNonceAt_Call) Return(_a0 uint64, _a1 error) *L2Mock_PendingNonceAt_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *L2Mock_PendingNonceAt_Call) RunAndReturn(run func(context.Context, common.Address) (uint64, error)) *L2Mock_PendingNonceAt_Call {
	_c.Call.Return(run)
	return _c
}

// SendTransaction provides a mock function with given fields: ctx, tx
func (_m *L2Mock) SendTransaction(ctx context.Context, tx *types.Transaction) error {
	ret := _m.Called(ctx, tx)

	if len(ret) == 0 {
		panic("no return value specified for SendTransaction")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *types.Transaction) error); ok {
		r0 = rf(ctx, tx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// L2Mock_SendTransaction_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SendTransaction'
type L2Mock_SendTransaction_Call struct {
	*mock.Call
}

// SendTransaction is a helper method to define mock.On call
//   - ctx context.Context
//   - tx *types.Transaction
func (_e *L2Mock_Expecter) SendTransaction(ctx interface{}, tx interface{}) *L2Mock_SendTransaction_Call {
	return &L2Mock_SendTransaction_Call{Call: _e.mock.On("SendTransaction", ctx, tx)}
}

func (_c *L2Mock_SendTransaction_Call) Run(run func(ctx context.Context, tx *types.Transaction)) *L2Mock_SendTransaction_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*types.Transaction))
	})
	return _c
}

func (_c *L2Mock_SendTransaction_Call) Return(_a0 error) *L2Mock_SendTransaction_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *L2Mock_SendTransaction_Call) RunAndReturn(run func(context.Context, *types.Transaction) error) *L2Mock_SendTransaction_Call {
	_c.Call.Return(run)
	return _c
}

// SubscribeFilterLogs provides a mock function with given fields: ctx, q, ch
func (_m *L2Mock) SubscribeFilterLogs(ctx context.Context, q ethereum.FilterQuery, ch chan<- types.Log) (ethereum.Subscription, error) {
	ret := _m.Called(ctx, q, ch)

	if len(ret) == 0 {
		panic("no return value specified for SubscribeFilterLogs")
	}

	var r0 ethereum.Subscription
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, ethereum.FilterQuery, chan<- types.Log) (ethereum.Subscription, error)); ok {
		return rf(ctx, q, ch)
	}
	if rf, ok := ret.Get(0).(func(context.Context, ethereum.FilterQuery, chan<- types.Log) ethereum.Subscription); ok {
		r0 = rf(ctx, q, ch)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ethereum.Subscription)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, ethereum.FilterQuery, chan<- types.Log) error); ok {
		r1 = rf(ctx, q, ch)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// L2Mock_SubscribeFilterLogs_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SubscribeFilterLogs'
type L2Mock_SubscribeFilterLogs_Call struct {
	*mock.Call
}

// SubscribeFilterLogs is a helper method to define mock.On call
//   - ctx context.Context
//   - q ethereum.FilterQuery
//   - ch chan<- types.Log
func (_e *L2Mock_Expecter) SubscribeFilterLogs(ctx interface{}, q interface{}, ch interface{}) *L2Mock_SubscribeFilterLogs_Call {
	return &L2Mock_SubscribeFilterLogs_Call{Call: _e.mock.On("SubscribeFilterLogs", ctx, q, ch)}
}

func (_c *L2Mock_SubscribeFilterLogs_Call) Run(run func(ctx context.Context, q ethereum.FilterQuery, ch chan<- types.Log)) *L2Mock_SubscribeFilterLogs_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(ethereum.FilterQuery), args[2].(chan<- types.Log))
	})
	return _c
}

func (_c *L2Mock_SubscribeFilterLogs_Call) Return(_a0 ethereum.Subscription, _a1 error) *L2Mock_SubscribeFilterLogs_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *L2Mock_SubscribeFilterLogs_Call) RunAndReturn(run func(context.Context, ethereum.FilterQuery, chan<- types.Log) (ethereum.Subscription, error)) *L2Mock_SubscribeFilterLogs_Call {
	_c.Call.Return(run)
	return _c
}

// SubscribeNewHead provides a mock function with given fields: ctx, ch
func (_m *L2Mock) SubscribeNewHead(ctx context.Context, ch chan<- *types.Header) (ethereum.Subscription, error) {
	ret := _m.Called(ctx, ch)

	if len(ret) == 0 {
		panic("no return value specified for SubscribeNewHead")
	}

	var r0 ethereum.Subscription
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, chan<- *types.Header) (ethereum.Subscription, error)); ok {
		return rf(ctx, ch)
	}
	if rf, ok := ret.Get(0).(func(context.Context, chan<- *types.Header) ethereum.Subscription); ok {
		r0 = rf(ctx, ch)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ethereum.Subscription)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, chan<- *types.Header) error); ok {
		r1 = rf(ctx, ch)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// L2Mock_SubscribeNewHead_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SubscribeNewHead'
type L2Mock_SubscribeNewHead_Call struct {
	*mock.Call
}

// SubscribeNewHead is a helper method to define mock.On call
//   - ctx context.Context
//   - ch chan<- *types.Header
func (_e *L2Mock_Expecter) SubscribeNewHead(ctx interface{}, ch interface{}) *L2Mock_SubscribeNewHead_Call {
	return &L2Mock_SubscribeNewHead_Call{Call: _e.mock.On("SubscribeNewHead", ctx, ch)}
}

func (_c *L2Mock_SubscribeNewHead_Call) Run(run func(ctx context.Context, ch chan<- *types.Header)) *L2Mock_SubscribeNewHead_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(chan<- *types.Header))
	})
	return _c
}

func (_c *L2Mock_SubscribeNewHead_Call) Return(_a0 ethereum.Subscription, _a1 error) *L2Mock_SubscribeNewHead_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *L2Mock_SubscribeNewHead_Call) RunAndReturn(run func(context.Context, chan<- *types.Header) (ethereum.Subscription, error)) *L2Mock_SubscribeNewHead_Call {
	_c.Call.Return(run)
	return _c
}

// SuggestGasPrice provides a mock function with given fields: ctx
func (_m *L2Mock) SuggestGasPrice(ctx context.Context) (*big.Int, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for SuggestGasPrice")
	}

	var r0 *big.Int
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (*big.Int, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) *big.Int); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// L2Mock_SuggestGasPrice_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SuggestGasPrice'
type L2Mock_SuggestGasPrice_Call struct {
	*mock.Call
}

// SuggestGasPrice is a helper method to define mock.On call
//   - ctx context.Context
func (_e *L2Mock_Expecter) SuggestGasPrice(ctx interface{}) *L2Mock_SuggestGasPrice_Call {
	return &L2Mock_SuggestGasPrice_Call{Call: _e.mock.On("SuggestGasPrice", ctx)}
}

func (_c *L2Mock_SuggestGasPrice_Call) Run(run func(ctx context.Context)) *L2Mock_SuggestGasPrice_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *L2Mock_SuggestGasPrice_Call) Return(_a0 *big.Int, _a1 error) *L2Mock_SuggestGasPrice_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *L2Mock_SuggestGasPrice_Call) RunAndReturn(run func(context.Context) (*big.Int, error)) *L2Mock_SuggestGasPrice_Call {
	_c.Call.Return(run)
	return _c
}

// SuggestGasTipCap provides a mock function with given fields: ctx
func (_m *L2Mock) SuggestGasTipCap(ctx context.Context) (*big.Int, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for SuggestGasTipCap")
	}

	var r0 *big.Int
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (*big.Int, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) *big.Int); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// L2Mock_SuggestGasTipCap_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SuggestGasTipCap'
type L2Mock_SuggestGasTipCap_Call struct {
	*mock.Call
}

// SuggestGasTipCap is a helper method to define mock.On call
//   - ctx context.Context
func (_e *L2Mock_Expecter) SuggestGasTipCap(ctx interface{}) *L2Mock_SuggestGasTipCap_Call {
	return &L2Mock_SuggestGasTipCap_Call{Call: _e.mock.On("SuggestGasTipCap", ctx)}
}

func (_c *L2Mock_SuggestGasTipCap_Call) Run(run func(ctx context.Context)) *L2Mock_SuggestGasTipCap_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *L2Mock_SuggestGasTipCap_Call) Return(_a0 *big.Int, _a1 error) *L2Mock_SuggestGasTipCap_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *L2Mock_SuggestGasTipCap_Call) RunAndReturn(run func(context.Context) (*big.Int, error)) *L2Mock_SuggestGasTipCap_Call {
	_c.Call.Return(run)
	return _c
}

// TransactionCount provides a mock function with given fields: ctx, blockHash
func (_m *L2Mock) TransactionCount(ctx context.Context, blockHash common.Hash) (uint, error) {
	ret := _m.Called(ctx, blockHash)

	if len(ret) == 0 {
		panic("no return value specified for TransactionCount")
	}

	var r0 uint
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, common.Hash) (uint, error)); ok {
		return rf(ctx, blockHash)
	}
	if rf, ok := ret.Get(0).(func(context.Context, common.Hash) uint); ok {
		r0 = rf(ctx, blockHash)
	} else {
		r0 = ret.Get(0).(uint)
	}

	if rf, ok := ret.Get(1).(func(context.Context, common.Hash) error); ok {
		r1 = rf(ctx, blockHash)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// L2Mock_TransactionCount_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'TransactionCount'
type L2Mock_TransactionCount_Call struct {
	*mock.Call
}

// TransactionCount is a helper method to define mock.On call
//   - ctx context.Context
//   - blockHash common.Hash
func (_e *L2Mock_Expecter) TransactionCount(ctx interface{}, blockHash interface{}) *L2Mock_TransactionCount_Call {
	return &L2Mock_TransactionCount_Call{Call: _e.mock.On("TransactionCount", ctx, blockHash)}
}

func (_c *L2Mock_TransactionCount_Call) Run(run func(ctx context.Context, blockHash common.Hash)) *L2Mock_TransactionCount_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(common.Hash))
	})
	return _c
}

func (_c *L2Mock_TransactionCount_Call) Return(_a0 uint, _a1 error) *L2Mock_TransactionCount_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *L2Mock_TransactionCount_Call) RunAndReturn(run func(context.Context, common.Hash) (uint, error)) *L2Mock_TransactionCount_Call {
	_c.Call.Return(run)
	return _c
}

// TransactionInBlock provides a mock function with given fields: ctx, blockHash, index
func (_m *L2Mock) TransactionInBlock(ctx context.Context, blockHash common.Hash, index uint) (*types.Transaction, error) {
	ret := _m.Called(ctx, blockHash, index)

	if len(ret) == 0 {
		panic("no return value specified for TransactionInBlock")
	}

	var r0 *types.Transaction
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, common.Hash, uint) (*types.Transaction, error)); ok {
		return rf(ctx, blockHash, index)
	}
	if rf, ok := ret.Get(0).(func(context.Context, common.Hash, uint) *types.Transaction); ok {
		r0 = rf(ctx, blockHash, index)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Transaction)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, common.Hash, uint) error); ok {
		r1 = rf(ctx, blockHash, index)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// L2Mock_TransactionInBlock_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'TransactionInBlock'
type L2Mock_TransactionInBlock_Call struct {
	*mock.Call
}

// TransactionInBlock is a helper method to define mock.On call
//   - ctx context.Context
//   - blockHash common.Hash
//   - index uint
func (_e *L2Mock_Expecter) TransactionInBlock(ctx interface{}, blockHash interface{}, index interface{}) *L2Mock_TransactionInBlock_Call {
	return &L2Mock_TransactionInBlock_Call{Call: _e.mock.On("TransactionInBlock", ctx, blockHash, index)}
}

func (_c *L2Mock_TransactionInBlock_Call) Run(run func(ctx context.Context, blockHash common.Hash, index uint)) *L2Mock_TransactionInBlock_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(common.Hash), args[2].(uint))
	})
	return _c
}

func (_c *L2Mock_TransactionInBlock_Call) Return(_a0 *types.Transaction, _a1 error) *L2Mock_TransactionInBlock_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *L2Mock_TransactionInBlock_Call) RunAndReturn(run func(context.Context, common.Hash, uint) (*types.Transaction, error)) *L2Mock_TransactionInBlock_Call {
	_c.Call.Return(run)
	return _c
}

// NewL2Mock creates a new instance of L2Mock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewL2Mock(t interface {
	mock.TestingT
	Cleanup(func())
}) *L2Mock {
	mock := &L2Mock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
