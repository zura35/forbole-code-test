// Code generated by mockery v2.32.2. DO NOT EDIT.

package mock

import (
	context "context"
	model "forbole_code_test/model"

	mock "github.com/stretchr/testify/mock"
)

// UserStore is an autogenerated mock type for the UserStore type
type UserStore struct {
	mock.Mock
}

type UserStore_Expecter struct {
	mock *mock.Mock
}

func (_m *UserStore) EXPECT() *UserStore_Expecter {
	return &UserStore_Expecter{mock: &_m.Mock}
}

// CreateUser provides a mock function with given fields: ctx, user
func (_m *UserStore) CreateUser(ctx context.Context, user *model.User) (*model.User, error) {
	ret := _m.Called(ctx, user)

	var r0 *model.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.User) (*model.User, error)); ok {
		return rf(ctx, user)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *model.User) *model.User); ok {
		r0 = rf(ctx, user)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *model.User) error); ok {
		r1 = rf(ctx, user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UserStore_CreateUser_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateUser'
type UserStore_CreateUser_Call struct {
	*mock.Call
}

// CreateUser is a helper method to define mock.On call
//   - ctx context.Context
//   - user *model.User
func (_e *UserStore_Expecter) CreateUser(ctx interface{}, user interface{}) *UserStore_CreateUser_Call {
	return &UserStore_CreateUser_Call{Call: _e.mock.On("CreateUser", ctx, user)}
}

func (_c *UserStore_CreateUser_Call) Run(run func(ctx context.Context, user *model.User)) *UserStore_CreateUser_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*model.User))
	})
	return _c
}

func (_c *UserStore_CreateUser_Call) Return(_a0 *model.User, _a1 error) *UserStore_CreateUser_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *UserStore_CreateUser_Call) RunAndReturn(run func(context.Context, *model.User) (*model.User, error)) *UserStore_CreateUser_Call {
	_c.Call.Return(run)
	return _c
}

// NewUserStore creates a new instance of UserStore. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUserStore(t interface {
	mock.TestingT
	Cleanup(func())
}) *UserStore {
	mock := &UserStore{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
