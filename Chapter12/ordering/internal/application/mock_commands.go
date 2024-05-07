// Code generated by mockery v2.14.0. DO NOT EDIT.

package application

import (
	context "context"
	commands "github.com/ibiscum/Event-Driven-Architecture-in-Golang/ordering/internal/application/commands"

	mock "github.com/stretchr/testify/mock"
)

// MockCommands is an autogenerated mock type for the Commands type
type MockCommands struct {
	mock.Mock
}

// ApproveOrder provides a mock function with given fields: ctx, cmd
func (_m *MockCommands) ApproveOrder(ctx context.Context, cmd commands.ApproveOrder) error {
	ret := _m.Called(ctx, cmd)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, commands.ApproveOrder) error); ok {
		r0 = rf(ctx, cmd)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CancelOrder provides a mock function with given fields: ctx, cmd
func (_m *MockCommands) CancelOrder(ctx context.Context, cmd commands.CancelOrder) error {
	ret := _m.Called(ctx, cmd)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, commands.CancelOrder) error); ok {
		r0 = rf(ctx, cmd)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CompleteOrder provides a mock function with given fields: ctx, cmd
func (_m *MockCommands) CompleteOrder(ctx context.Context, cmd commands.CompleteOrder) error {
	ret := _m.Called(ctx, cmd)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, commands.CompleteOrder) error); ok {
		r0 = rf(ctx, cmd)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateOrder provides a mock function with given fields: ctx, cmd
func (_m *MockCommands) CreateOrder(ctx context.Context, cmd commands.CreateOrder) error {
	ret := _m.Called(ctx, cmd)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, commands.CreateOrder) error); ok {
		r0 = rf(ctx, cmd)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ReadyOrder provides a mock function with given fields: ctx, cmd
func (_m *MockCommands) ReadyOrder(ctx context.Context, cmd commands.ReadyOrder) error {
	ret := _m.Called(ctx, cmd)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, commands.ReadyOrder) error); ok {
		r0 = rf(ctx, cmd)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RejectOrder provides a mock function with given fields: ctx, cmd
func (_m *MockCommands) RejectOrder(ctx context.Context, cmd commands.RejectOrder) error {
	ret := _m.Called(ctx, cmd)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, commands.RejectOrder) error); ok {
		r0 = rf(ctx, cmd)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewMockCommands interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockCommands creates a new instance of MockCommands. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockCommands(t mockConstructorTestingTNewMockCommands) *MockCommands {
	mock := &MockCommands{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
