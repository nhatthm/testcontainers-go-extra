// Code generated by mockery v2.45.1. DO NOT EDIT.

package wait

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	wait "github.com/testcontainers/testcontainers-go/wait"
)

// Strategy is an autogenerated mock type for the Strategy type
type Strategy struct {
	mock.Mock
}

// WaitUntilReady provides a mock function with given fields: _a0, _a1
func (_m *Strategy) WaitUntilReady(_a0 context.Context, _a1 wait.StrategyTarget) error {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for WaitUntilReady")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, wait.StrategyTarget) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewStrategy creates a new instance of Strategy. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewStrategy(t interface {
	mock.TestingT
	Cleanup(func())
}) *Strategy {
	mock := &Strategy{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
