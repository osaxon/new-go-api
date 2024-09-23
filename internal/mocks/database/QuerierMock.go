// Code generated by mockery v2.46.0. DO NOT EDIT.

package mocks

import (
	context "context"
	database "new-go-api/internal/database"

	mock "github.com/stretchr/testify/mock"
)

// QuerierMock is an autogenerated mock type for the Querier type
type QuerierMock struct {
	mock.Mock
}

type QuerierMock_Expecter struct {
	mock *mock.Mock
}

func (_m *QuerierMock) EXPECT() *QuerierMock_Expecter {
	return &QuerierMock_Expecter{mock: &_m.Mock}
}

// CreateTask provides a mock function with given fields: ctx, arg
func (_m *QuerierMock) CreateTask(ctx context.Context, arg database.CreateTaskParams) (database.Task, error) {
	ret := _m.Called(ctx, arg)

	if len(ret) == 0 {
		panic("no return value specified for CreateTask")
	}

	var r0 database.Task
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, database.CreateTaskParams) (database.Task, error)); ok {
		return rf(ctx, arg)
	}
	if rf, ok := ret.Get(0).(func(context.Context, database.CreateTaskParams) database.Task); ok {
		r0 = rf(ctx, arg)
	} else {
		r0 = ret.Get(0).(database.Task)
	}

	if rf, ok := ret.Get(1).(func(context.Context, database.CreateTaskParams) error); ok {
		r1 = rf(ctx, arg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// QuerierMock_CreateTask_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateTask'
type QuerierMock_CreateTask_Call struct {
	*mock.Call
}

// CreateTask is a helper method to define mock.On call
//   - ctx context.Context
//   - arg database.CreateTaskParams
func (_e *QuerierMock_Expecter) CreateTask(ctx interface{}, arg interface{}) *QuerierMock_CreateTask_Call {
	return &QuerierMock_CreateTask_Call{Call: _e.mock.On("CreateTask", ctx, arg)}
}

func (_c *QuerierMock_CreateTask_Call) Run(run func(ctx context.Context, arg database.CreateTaskParams)) *QuerierMock_CreateTask_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(database.CreateTaskParams))
	})
	return _c
}

func (_c *QuerierMock_CreateTask_Call) Return(_a0 database.Task, _a1 error) *QuerierMock_CreateTask_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *QuerierMock_CreateTask_Call) RunAndReturn(run func(context.Context, database.CreateTaskParams) (database.Task, error)) *QuerierMock_CreateTask_Call {
	_c.Call.Return(run)
	return _c
}

// CreateUser provides a mock function with given fields: ctx, arg
func (_m *QuerierMock) CreateUser(ctx context.Context, arg database.CreateUserParams) (database.User, error) {
	ret := _m.Called(ctx, arg)

	if len(ret) == 0 {
		panic("no return value specified for CreateUser")
	}

	var r0 database.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, database.CreateUserParams) (database.User, error)); ok {
		return rf(ctx, arg)
	}
	if rf, ok := ret.Get(0).(func(context.Context, database.CreateUserParams) database.User); ok {
		r0 = rf(ctx, arg)
	} else {
		r0 = ret.Get(0).(database.User)
	}

	if rf, ok := ret.Get(1).(func(context.Context, database.CreateUserParams) error); ok {
		r1 = rf(ctx, arg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// QuerierMock_CreateUser_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateUser'
type QuerierMock_CreateUser_Call struct {
	*mock.Call
}

// CreateUser is a helper method to define mock.On call
//   - ctx context.Context
//   - arg database.CreateUserParams
func (_e *QuerierMock_Expecter) CreateUser(ctx interface{}, arg interface{}) *QuerierMock_CreateUser_Call {
	return &QuerierMock_CreateUser_Call{Call: _e.mock.On("CreateUser", ctx, arg)}
}

func (_c *QuerierMock_CreateUser_Call) Run(run func(ctx context.Context, arg database.CreateUserParams)) *QuerierMock_CreateUser_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(database.CreateUserParams))
	})
	return _c
}

func (_c *QuerierMock_CreateUser_Call) Return(_a0 database.User, _a1 error) *QuerierMock_CreateUser_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *QuerierMock_CreateUser_Call) RunAndReturn(run func(context.Context, database.CreateUserParams) (database.User, error)) *QuerierMock_CreateUser_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteTask provides a mock function with given fields: ctx, id
func (_m *QuerierMock) DeleteTask(ctx context.Context, id int64) error {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for DeleteTask")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// QuerierMock_DeleteTask_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteTask'
type QuerierMock_DeleteTask_Call struct {
	*mock.Call
}

// DeleteTask is a helper method to define mock.On call
//   - ctx context.Context
//   - id int64
func (_e *QuerierMock_Expecter) DeleteTask(ctx interface{}, id interface{}) *QuerierMock_DeleteTask_Call {
	return &QuerierMock_DeleteTask_Call{Call: _e.mock.On("DeleteTask", ctx, id)}
}

func (_c *QuerierMock_DeleteTask_Call) Run(run func(ctx context.Context, id int64)) *QuerierMock_DeleteTask_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64))
	})
	return _c
}

func (_c *QuerierMock_DeleteTask_Call) Return(_a0 error) *QuerierMock_DeleteTask_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *QuerierMock_DeleteTask_Call) RunAndReturn(run func(context.Context, int64) error) *QuerierMock_DeleteTask_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteUser provides a mock function with given fields: ctx, id
func (_m *QuerierMock) DeleteUser(ctx context.Context, id int64) error {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for DeleteUser")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// QuerierMock_DeleteUser_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteUser'
type QuerierMock_DeleteUser_Call struct {
	*mock.Call
}

// DeleteUser is a helper method to define mock.On call
//   - ctx context.Context
//   - id int64
func (_e *QuerierMock_Expecter) DeleteUser(ctx interface{}, id interface{}) *QuerierMock_DeleteUser_Call {
	return &QuerierMock_DeleteUser_Call{Call: _e.mock.On("DeleteUser", ctx, id)}
}

func (_c *QuerierMock_DeleteUser_Call) Run(run func(ctx context.Context, id int64)) *QuerierMock_DeleteUser_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64))
	})
	return _c
}

func (_c *QuerierMock_DeleteUser_Call) Return(_a0 error) *QuerierMock_DeleteUser_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *QuerierMock_DeleteUser_Call) RunAndReturn(run func(context.Context, int64) error) *QuerierMock_DeleteUser_Call {
	_c.Call.Return(run)
	return _c
}

// GetTask provides a mock function with given fields: ctx, id
func (_m *QuerierMock) GetTask(ctx context.Context, id int64) (database.Task, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetTask")
	}

	var r0 database.Task
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) (database.Task, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) database.Task); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(database.Task)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// QuerierMock_GetTask_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetTask'
type QuerierMock_GetTask_Call struct {
	*mock.Call
}

// GetTask is a helper method to define mock.On call
//   - ctx context.Context
//   - id int64
func (_e *QuerierMock_Expecter) GetTask(ctx interface{}, id interface{}) *QuerierMock_GetTask_Call {
	return &QuerierMock_GetTask_Call{Call: _e.mock.On("GetTask", ctx, id)}
}

func (_c *QuerierMock_GetTask_Call) Run(run func(ctx context.Context, id int64)) *QuerierMock_GetTask_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64))
	})
	return _c
}

func (_c *QuerierMock_GetTask_Call) Return(_a0 database.Task, _a1 error) *QuerierMock_GetTask_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *QuerierMock_GetTask_Call) RunAndReturn(run func(context.Context, int64) (database.Task, error)) *QuerierMock_GetTask_Call {
	_c.Call.Return(run)
	return _c
}

// GetUser provides a mock function with given fields: ctx, id
func (_m *QuerierMock) GetUser(ctx context.Context, id int64) (database.User, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetUser")
	}

	var r0 database.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) (database.User, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) database.User); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(database.User)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// QuerierMock_GetUser_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetUser'
type QuerierMock_GetUser_Call struct {
	*mock.Call
}

// GetUser is a helper method to define mock.On call
//   - ctx context.Context
//   - id int64
func (_e *QuerierMock_Expecter) GetUser(ctx interface{}, id interface{}) *QuerierMock_GetUser_Call {
	return &QuerierMock_GetUser_Call{Call: _e.mock.On("GetUser", ctx, id)}
}

func (_c *QuerierMock_GetUser_Call) Run(run func(ctx context.Context, id int64)) *QuerierMock_GetUser_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64))
	})
	return _c
}

func (_c *QuerierMock_GetUser_Call) Return(_a0 database.User, _a1 error) *QuerierMock_GetUser_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *QuerierMock_GetUser_Call) RunAndReturn(run func(context.Context, int64) (database.User, error)) *QuerierMock_GetUser_Call {
	_c.Call.Return(run)
	return _c
}

// GetUserByEmail provides a mock function with given fields: ctx, email
func (_m *QuerierMock) GetUserByEmail(ctx context.Context, email string) (database.User, error) {
	ret := _m.Called(ctx, email)

	if len(ret) == 0 {
		panic("no return value specified for GetUserByEmail")
	}

	var r0 database.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (database.User, error)); ok {
		return rf(ctx, email)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) database.User); ok {
		r0 = rf(ctx, email)
	} else {
		r0 = ret.Get(0).(database.User)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// QuerierMock_GetUserByEmail_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetUserByEmail'
type QuerierMock_GetUserByEmail_Call struct {
	*mock.Call
}

// GetUserByEmail is a helper method to define mock.On call
//   - ctx context.Context
//   - email string
func (_e *QuerierMock_Expecter) GetUserByEmail(ctx interface{}, email interface{}) *QuerierMock_GetUserByEmail_Call {
	return &QuerierMock_GetUserByEmail_Call{Call: _e.mock.On("GetUserByEmail", ctx, email)}
}

func (_c *QuerierMock_GetUserByEmail_Call) Run(run func(ctx context.Context, email string)) *QuerierMock_GetUserByEmail_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *QuerierMock_GetUserByEmail_Call) Return(_a0 database.User, _a1 error) *QuerierMock_GetUserByEmail_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *QuerierMock_GetUserByEmail_Call) RunAndReturn(run func(context.Context, string) (database.User, error)) *QuerierMock_GetUserByEmail_Call {
	_c.Call.Return(run)
	return _c
}

// ListTasks provides a mock function with given fields: ctx
func (_m *QuerierMock) ListTasks(ctx context.Context) ([]database.Task, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for ListTasks")
	}

	var r0 []database.Task
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]database.Task, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []database.Task); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]database.Task)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// QuerierMock_ListTasks_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListTasks'
type QuerierMock_ListTasks_Call struct {
	*mock.Call
}

// ListTasks is a helper method to define mock.On call
//   - ctx context.Context
func (_e *QuerierMock_Expecter) ListTasks(ctx interface{}) *QuerierMock_ListTasks_Call {
	return &QuerierMock_ListTasks_Call{Call: _e.mock.On("ListTasks", ctx)}
}

func (_c *QuerierMock_ListTasks_Call) Run(run func(ctx context.Context)) *QuerierMock_ListTasks_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *QuerierMock_ListTasks_Call) Return(_a0 []database.Task, _a1 error) *QuerierMock_ListTasks_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *QuerierMock_ListTasks_Call) RunAndReturn(run func(context.Context) ([]database.Task, error)) *QuerierMock_ListTasks_Call {
	_c.Call.Return(run)
	return _c
}

// ListUsers provides a mock function with given fields: ctx
func (_m *QuerierMock) ListUsers(ctx context.Context) ([]database.User, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for ListUsers")
	}

	var r0 []database.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]database.User, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []database.User); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]database.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// QuerierMock_ListUsers_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListUsers'
type QuerierMock_ListUsers_Call struct {
	*mock.Call
}

// ListUsers is a helper method to define mock.On call
//   - ctx context.Context
func (_e *QuerierMock_Expecter) ListUsers(ctx interface{}) *QuerierMock_ListUsers_Call {
	return &QuerierMock_ListUsers_Call{Call: _e.mock.On("ListUsers", ctx)}
}

func (_c *QuerierMock_ListUsers_Call) Run(run func(ctx context.Context)) *QuerierMock_ListUsers_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *QuerierMock_ListUsers_Call) Return(_a0 []database.User, _a1 error) *QuerierMock_ListUsers_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *QuerierMock_ListUsers_Call) RunAndReturn(run func(context.Context) ([]database.User, error)) *QuerierMock_ListUsers_Call {
	_c.Call.Return(run)
	return _c
}

// ListUsersTasks provides a mock function with given fields: ctx, userID
func (_m *QuerierMock) ListUsersTasks(ctx context.Context, userID int64) ([]database.Task, error) {
	ret := _m.Called(ctx, userID)

	if len(ret) == 0 {
		panic("no return value specified for ListUsersTasks")
	}

	var r0 []database.Task
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) ([]database.Task, error)); ok {
		return rf(ctx, userID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) []database.Task); ok {
		r0 = rf(ctx, userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]database.Task)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// QuerierMock_ListUsersTasks_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListUsersTasks'
type QuerierMock_ListUsersTasks_Call struct {
	*mock.Call
}

// ListUsersTasks is a helper method to define mock.On call
//   - ctx context.Context
//   - userID int64
func (_e *QuerierMock_Expecter) ListUsersTasks(ctx interface{}, userID interface{}) *QuerierMock_ListUsersTasks_Call {
	return &QuerierMock_ListUsersTasks_Call{Call: _e.mock.On("ListUsersTasks", ctx, userID)}
}

func (_c *QuerierMock_ListUsersTasks_Call) Run(run func(ctx context.Context, userID int64)) *QuerierMock_ListUsersTasks_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64))
	})
	return _c
}

func (_c *QuerierMock_ListUsersTasks_Call) Return(_a0 []database.Task, _a1 error) *QuerierMock_ListUsersTasks_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *QuerierMock_ListUsersTasks_Call) RunAndReturn(run func(context.Context, int64) ([]database.Task, error)) *QuerierMock_ListUsersTasks_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateTask provides a mock function with given fields: ctx, arg
func (_m *QuerierMock) UpdateTask(ctx context.Context, arg database.UpdateTaskParams) error {
	ret := _m.Called(ctx, arg)

	if len(ret) == 0 {
		panic("no return value specified for UpdateTask")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, database.UpdateTaskParams) error); ok {
		r0 = rf(ctx, arg)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// QuerierMock_UpdateTask_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateTask'
type QuerierMock_UpdateTask_Call struct {
	*mock.Call
}

// UpdateTask is a helper method to define mock.On call
//   - ctx context.Context
//   - arg database.UpdateTaskParams
func (_e *QuerierMock_Expecter) UpdateTask(ctx interface{}, arg interface{}) *QuerierMock_UpdateTask_Call {
	return &QuerierMock_UpdateTask_Call{Call: _e.mock.On("UpdateTask", ctx, arg)}
}

func (_c *QuerierMock_UpdateTask_Call) Run(run func(ctx context.Context, arg database.UpdateTaskParams)) *QuerierMock_UpdateTask_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(database.UpdateTaskParams))
	})
	return _c
}

func (_c *QuerierMock_UpdateTask_Call) Return(_a0 error) *QuerierMock_UpdateTask_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *QuerierMock_UpdateTask_Call) RunAndReturn(run func(context.Context, database.UpdateTaskParams) error) *QuerierMock_UpdateTask_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateUser provides a mock function with given fields: ctx, arg
func (_m *QuerierMock) UpdateUser(ctx context.Context, arg database.UpdateUserParams) error {
	ret := _m.Called(ctx, arg)

	if len(ret) == 0 {
		panic("no return value specified for UpdateUser")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, database.UpdateUserParams) error); ok {
		r0 = rf(ctx, arg)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// QuerierMock_UpdateUser_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateUser'
type QuerierMock_UpdateUser_Call struct {
	*mock.Call
}

// UpdateUser is a helper method to define mock.On call
//   - ctx context.Context
//   - arg database.UpdateUserParams
func (_e *QuerierMock_Expecter) UpdateUser(ctx interface{}, arg interface{}) *QuerierMock_UpdateUser_Call {
	return &QuerierMock_UpdateUser_Call{Call: _e.mock.On("UpdateUser", ctx, arg)}
}

func (_c *QuerierMock_UpdateUser_Call) Run(run func(ctx context.Context, arg database.UpdateUserParams)) *QuerierMock_UpdateUser_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(database.UpdateUserParams))
	})
	return _c
}

func (_c *QuerierMock_UpdateUser_Call) Return(_a0 error) *QuerierMock_UpdateUser_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *QuerierMock_UpdateUser_Call) RunAndReturn(run func(context.Context, database.UpdateUserParams) error) *QuerierMock_UpdateUser_Call {
	_c.Call.Return(run)
	return _c
}

// NewQuerierMock creates a new instance of QuerierMock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewQuerierMock(t interface {
	mock.TestingT
	Cleanup(func())
}) *QuerierMock {
	mock := &QuerierMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}