// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// Console is an autogenerated mock type for the Console type
type Console struct {
	mock.Mock
}

type Console_Expecter struct {
	mock *mock.Mock
}

func (_m *Console) EXPECT() *Console_Expecter {
	return &Console_Expecter{mock: &_m.Mock}
}

// RunCmd provides a mock function with given fields: ctx, dir, command, commandParams
func (_m *Console) RunCmd(ctx context.Context, dir string, command string, commandParams ...string) (string, error) {
	_va := make([]interface{}, len(commandParams))
	for _i := range commandParams {
		_va[_i] = commandParams[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, dir, command)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for RunCmd")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, ...string) (string, error)); ok {
		return rf(ctx, dir, command, commandParams...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, ...string) string); ok {
		r0 = rf(ctx, dir, command, commandParams...)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, ...string) error); ok {
		r1 = rf(ctx, dir, command, commandParams...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Console_RunCmd_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RunCmd'
type Console_RunCmd_Call struct {
	*mock.Call
}

// RunCmd is a helper method to define mock.On call
//   - ctx context.Context
//   - dir string
//   - command string
//   - commandParams ...string
func (_e *Console_Expecter) RunCmd(ctx interface{}, dir interface{}, command interface{}, commandParams ...interface{}) *Console_RunCmd_Call {
	return &Console_RunCmd_Call{Call: _e.mock.On("RunCmd",
		append([]interface{}{ctx, dir, command}, commandParams...)...)}
}

func (_c *Console_RunCmd_Call) Run(run func(ctx context.Context, dir string, command string, commandParams ...string)) *Console_RunCmd_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]string, len(args)-3)
		for i, a := range args[3:] {
			if a != nil {
				variadicArgs[i] = a.(string)
			}
		}
		run(args[0].(context.Context), args[1].(string), args[2].(string), variadicArgs...)
	})
	return _c
}

func (_c *Console_RunCmd_Call) Return(_a0 string, _a1 error) *Console_RunCmd_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Console_RunCmd_Call) RunAndReturn(run func(context.Context, string, string, ...string) (string, error)) *Console_RunCmd_Call {
	_c.Call.Return(run)
	return _c
}

// NewConsole creates a new instance of Console. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewConsole(t interface {
	mock.TestingT
	Cleanup(func())
}) *Console {
	mock := &Console{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
