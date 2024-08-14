// Code generated by mockery. DO NOT EDIT.

package mocks_da

import (
	context "context"

	etherman "github.com/0xPolygon/cdk/etherman"

	mock "github.com/stretchr/testify/mock"
)

// SequenceSender is an autogenerated mock type for the SequenceSender type
type SequenceSender struct {
	mock.Mock
}

type SequenceSender_Expecter struct {
	mock *mock.Mock
}

func (_m *SequenceSender) EXPECT() *SequenceSender_Expecter {
	return &SequenceSender_Expecter{mock: &_m.Mock}
}

// PostSequenceBanana provides a mock function with given fields: ctx, sequence
func (_m *SequenceSender) PostSequenceBanana(ctx context.Context, sequence etherman.SequenceBanana) ([]byte, error) {
	ret := _m.Called(ctx, sequence)

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, etherman.SequenceBanana) ([]byte, error)); ok {
		return rf(ctx, sequence)
	}
	if rf, ok := ret.Get(0).(func(context.Context, etherman.SequenceBanana) []byte); ok {
		r0 = rf(ctx, sequence)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, etherman.SequenceBanana) error); ok {
		r1 = rf(ctx, sequence)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SequenceSender_PostSequenceBanana_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PostSequenceBanana'
type SequenceSender_PostSequenceBanana_Call struct {
	*mock.Call
}

// PostSequenceBanana is a helper method to define mock.On call
//   - ctx context.Context
//   - sequence etherman.SequenceBanana
func (_e *SequenceSender_Expecter) PostSequenceBanana(ctx interface{}, sequence interface{}) *SequenceSender_PostSequenceBanana_Call {
	return &SequenceSender_PostSequenceBanana_Call{Call: _e.mock.On("PostSequenceBanana", ctx, sequence)}
}

func (_c *SequenceSender_PostSequenceBanana_Call) Run(run func(ctx context.Context, sequence etherman.SequenceBanana)) *SequenceSender_PostSequenceBanana_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(etherman.SequenceBanana))
	})
	return _c
}

func (_c *SequenceSender_PostSequenceBanana_Call) Return(_a0 []byte, _a1 error) *SequenceSender_PostSequenceBanana_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *SequenceSender_PostSequenceBanana_Call) RunAndReturn(run func(context.Context, etherman.SequenceBanana) ([]byte, error)) *SequenceSender_PostSequenceBanana_Call {
	_c.Call.Return(run)
	return _c
}

// PostSequenceElderberry provides a mock function with given fields: ctx, batchesData
func (_m *SequenceSender) PostSequenceElderberry(ctx context.Context, batchesData [][]byte) ([]byte, error) {
	ret := _m.Called(ctx, batchesData)

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, [][]byte) ([]byte, error)); ok {
		return rf(ctx, batchesData)
	}
	if rf, ok := ret.Get(0).(func(context.Context, [][]byte) []byte); ok {
		r0 = rf(ctx, batchesData)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, [][]byte) error); ok {
		r1 = rf(ctx, batchesData)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SequenceSender_PostSequenceElderberry_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PostSequenceElderberry'
type SequenceSender_PostSequenceElderberry_Call struct {
	*mock.Call
}

// PostSequenceElderberry is a helper method to define mock.On call
//   - ctx context.Context
//   - batchesData [][]byte
func (_e *SequenceSender_Expecter) PostSequenceElderberry(ctx interface{}, batchesData interface{}) *SequenceSender_PostSequenceElderberry_Call {
	return &SequenceSender_PostSequenceElderberry_Call{Call: _e.mock.On("PostSequenceElderberry", ctx, batchesData)}
}

func (_c *SequenceSender_PostSequenceElderberry_Call) Run(run func(ctx context.Context, batchesData [][]byte)) *SequenceSender_PostSequenceElderberry_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([][]byte))
	})
	return _c
}

func (_c *SequenceSender_PostSequenceElderberry_Call) Return(_a0 []byte, _a1 error) *SequenceSender_PostSequenceElderberry_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *SequenceSender_PostSequenceElderberry_Call) RunAndReturn(run func(context.Context, [][]byte) ([]byte, error)) *SequenceSender_PostSequenceElderberry_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewSequenceSender interface {
	mock.TestingT
	Cleanup(func())
}

// NewSequenceSender creates a new instance of SequenceSender. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewSequenceSender(t mockConstructorTestingTNewSequenceSender) *SequenceSender {
	mock := &SequenceSender{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}