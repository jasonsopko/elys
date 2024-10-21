// Code generated by mockery v2.45.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"

	types "github.com/cosmos/cosmos-sdk/types"
)

// AuthorizationChecker is an autogenerated mock type for the AuthorizationChecker type
type AuthorizationChecker struct {
	mock.Mock
}

type AuthorizationChecker_Expecter struct {
	mock *mock.Mock
}

func (_m *AuthorizationChecker) EXPECT() *AuthorizationChecker_Expecter {
	return &AuthorizationChecker_Expecter{mock: &_m.Mock}
}

// CheckIfWhitelisted provides a mock function with given fields: ctx, creator
func (_m *AuthorizationChecker) CheckIfWhitelisted(ctx types.Context, creator types.AccAddress) bool {
	ret := _m.Called(ctx, creator)

	if len(ret) == 0 {
		panic("no return value specified for CheckIfWhitelisted")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func(types.Context, types.AccAddress) bool); ok {
		r0 = rf(ctx, creator)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// AuthorizationChecker_CheckIfWhitelisted_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CheckIfWhitelisted'
type AuthorizationChecker_CheckIfWhitelisted_Call struct {
	*mock.Call
}

// CheckIfWhitelisted is a helper method to define mock.On call
//   - ctx types.Context
//   - creator types.AccAddress
func (_e *AuthorizationChecker_Expecter) CheckIfWhitelisted(ctx interface{}, creator interface{}) *AuthorizationChecker_CheckIfWhitelisted_Call {
	return &AuthorizationChecker_CheckIfWhitelisted_Call{Call: _e.mock.On("CheckIfWhitelisted", ctx, creator)}
}

func (_c *AuthorizationChecker_CheckIfWhitelisted_Call) Run(run func(ctx types.Context, creator types.AccAddress)) *AuthorizationChecker_CheckIfWhitelisted_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(types.Context), args[1].(types.AccAddress))
	})
	return _c
}

func (_c *AuthorizationChecker_CheckIfWhitelisted_Call) Return(_a0 bool) *AuthorizationChecker_CheckIfWhitelisted_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *AuthorizationChecker_CheckIfWhitelisted_Call) RunAndReturn(run func(types.Context, types.AccAddress) bool) *AuthorizationChecker_CheckIfWhitelisted_Call {
	_c.Call.Return(run)
	return _c
}

// IsWhitelistingEnabled provides a mock function with given fields: ctx
func (_m *AuthorizationChecker) IsWhitelistingEnabled(ctx types.Context) bool {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for IsWhitelistingEnabled")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func(types.Context) bool); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// AuthorizationChecker_IsWhitelistingEnabled_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsWhitelistingEnabled'
type AuthorizationChecker_IsWhitelistingEnabled_Call struct {
	*mock.Call
}

// IsWhitelistingEnabled is a helper method to define mock.On call
//   - ctx types.Context
func (_e *AuthorizationChecker_Expecter) IsWhitelistingEnabled(ctx interface{}) *AuthorizationChecker_IsWhitelistingEnabled_Call {
	return &AuthorizationChecker_IsWhitelistingEnabled_Call{Call: _e.mock.On("IsWhitelistingEnabled", ctx)}
}

func (_c *AuthorizationChecker_IsWhitelistingEnabled_Call) Run(run func(ctx types.Context)) *AuthorizationChecker_IsWhitelistingEnabled_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(types.Context))
	})
	return _c
}

func (_c *AuthorizationChecker_IsWhitelistingEnabled_Call) Return(_a0 bool) *AuthorizationChecker_IsWhitelistingEnabled_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *AuthorizationChecker_IsWhitelistingEnabled_Call) RunAndReturn(run func(types.Context) bool) *AuthorizationChecker_IsWhitelistingEnabled_Call {
	_c.Call.Return(run)
	return _c
}

// NewAuthorizationChecker creates a new instance of AuthorizationChecker. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAuthorizationChecker(t interface {
	mock.TestingT
	Cleanup(func())
}) *AuthorizationChecker {
	mock := &AuthorizationChecker{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
