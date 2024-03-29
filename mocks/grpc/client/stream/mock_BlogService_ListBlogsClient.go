// Code generated by mockery. DO NOT EDIT.

package clientstreamprotomock

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	metadata "google.golang.org/grpc/metadata"

	proto "github.com/grpc-mongo-go/gen/proto"
)

// ClientStreamProtoMock is an autogenerated mock type for the BlogService_ListBlogsClient type
type ClientStreamProtoMock struct {
	mock.Mock
}

type ClientStreamProtoMock_Expecter struct {
	mock *mock.Mock
}

func (_m *ClientStreamProtoMock) EXPECT() *ClientStreamProtoMock_Expecter {
	return &ClientStreamProtoMock_Expecter{mock: &_m.Mock}
}

// CloseSend provides a mock function with given fields:
func (_m *ClientStreamProtoMock) CloseSend() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for CloseSend")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ClientStreamProtoMock_CloseSend_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CloseSend'
type ClientStreamProtoMock_CloseSend_Call struct {
	*mock.Call
}

// CloseSend is a helper method to define mock.On call
func (_e *ClientStreamProtoMock_Expecter) CloseSend() *ClientStreamProtoMock_CloseSend_Call {
	return &ClientStreamProtoMock_CloseSend_Call{Call: _e.mock.On("CloseSend")}
}

func (_c *ClientStreamProtoMock_CloseSend_Call) Run(run func()) *ClientStreamProtoMock_CloseSend_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ClientStreamProtoMock_CloseSend_Call) Return(_a0 error) *ClientStreamProtoMock_CloseSend_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ClientStreamProtoMock_CloseSend_Call) RunAndReturn(run func() error) *ClientStreamProtoMock_CloseSend_Call {
	_c.Call.Return(run)
	return _c
}

// Context provides a mock function with given fields:
func (_m *ClientStreamProtoMock) Context() context.Context {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Context")
	}

	var r0 context.Context
	if rf, ok := ret.Get(0).(func() context.Context); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(context.Context)
		}
	}

	return r0
}

// ClientStreamProtoMock_Context_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Context'
type ClientStreamProtoMock_Context_Call struct {
	*mock.Call
}

// Context is a helper method to define mock.On call
func (_e *ClientStreamProtoMock_Expecter) Context() *ClientStreamProtoMock_Context_Call {
	return &ClientStreamProtoMock_Context_Call{Call: _e.mock.On("Context")}
}

func (_c *ClientStreamProtoMock_Context_Call) Run(run func()) *ClientStreamProtoMock_Context_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ClientStreamProtoMock_Context_Call) Return(_a0 context.Context) *ClientStreamProtoMock_Context_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ClientStreamProtoMock_Context_Call) RunAndReturn(run func() context.Context) *ClientStreamProtoMock_Context_Call {
	_c.Call.Return(run)
	return _c
}

// Header provides a mock function with given fields:
func (_m *ClientStreamProtoMock) Header() (metadata.MD, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Header")
	}

	var r0 metadata.MD
	var r1 error
	if rf, ok := ret.Get(0).(func() (metadata.MD, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() metadata.MD); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(metadata.MD)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ClientStreamProtoMock_Header_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Header'
type ClientStreamProtoMock_Header_Call struct {
	*mock.Call
}

// Header is a helper method to define mock.On call
func (_e *ClientStreamProtoMock_Expecter) Header() *ClientStreamProtoMock_Header_Call {
	return &ClientStreamProtoMock_Header_Call{Call: _e.mock.On("Header")}
}

func (_c *ClientStreamProtoMock_Header_Call) Run(run func()) *ClientStreamProtoMock_Header_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ClientStreamProtoMock_Header_Call) Return(_a0 metadata.MD, _a1 error) *ClientStreamProtoMock_Header_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ClientStreamProtoMock_Header_Call) RunAndReturn(run func() (metadata.MD, error)) *ClientStreamProtoMock_Header_Call {
	_c.Call.Return(run)
	return _c
}

// Recv provides a mock function with given fields:
func (_m *ClientStreamProtoMock) Recv() (*proto.Blog, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Recv")
	}

	var r0 *proto.Blog
	var r1 error
	if rf, ok := ret.Get(0).(func() (*proto.Blog, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() *proto.Blog); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*proto.Blog)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ClientStreamProtoMock_Recv_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Recv'
type ClientStreamProtoMock_Recv_Call struct {
	*mock.Call
}

// Recv is a helper method to define mock.On call
func (_e *ClientStreamProtoMock_Expecter) Recv() *ClientStreamProtoMock_Recv_Call {
	return &ClientStreamProtoMock_Recv_Call{Call: _e.mock.On("Recv")}
}

func (_c *ClientStreamProtoMock_Recv_Call) Run(run func()) *ClientStreamProtoMock_Recv_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ClientStreamProtoMock_Recv_Call) Return(_a0 *proto.Blog, _a1 error) *ClientStreamProtoMock_Recv_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ClientStreamProtoMock_Recv_Call) RunAndReturn(run func() (*proto.Blog, error)) *ClientStreamProtoMock_Recv_Call {
	_c.Call.Return(run)
	return _c
}

// RecvMsg provides a mock function with given fields: m
func (_m *ClientStreamProtoMock) RecvMsg(m interface{}) error {
	ret := _m.Called(m)

	if len(ret) == 0 {
		panic("no return value specified for RecvMsg")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(m)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ClientStreamProtoMock_RecvMsg_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RecvMsg'
type ClientStreamProtoMock_RecvMsg_Call struct {
	*mock.Call
}

// RecvMsg is a helper method to define mock.On call
//   - m interface{}
func (_e *ClientStreamProtoMock_Expecter) RecvMsg(m interface{}) *ClientStreamProtoMock_RecvMsg_Call {
	return &ClientStreamProtoMock_RecvMsg_Call{Call: _e.mock.On("RecvMsg", m)}
}

func (_c *ClientStreamProtoMock_RecvMsg_Call) Run(run func(m interface{})) *ClientStreamProtoMock_RecvMsg_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(interface{}))
	})
	return _c
}

func (_c *ClientStreamProtoMock_RecvMsg_Call) Return(_a0 error) *ClientStreamProtoMock_RecvMsg_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ClientStreamProtoMock_RecvMsg_Call) RunAndReturn(run func(interface{}) error) *ClientStreamProtoMock_RecvMsg_Call {
	_c.Call.Return(run)
	return _c
}

// SendMsg provides a mock function with given fields: m
func (_m *ClientStreamProtoMock) SendMsg(m interface{}) error {
	ret := _m.Called(m)

	if len(ret) == 0 {
		panic("no return value specified for SendMsg")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(m)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ClientStreamProtoMock_SendMsg_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SendMsg'
type ClientStreamProtoMock_SendMsg_Call struct {
	*mock.Call
}

// SendMsg is a helper method to define mock.On call
//   - m interface{}
func (_e *ClientStreamProtoMock_Expecter) SendMsg(m interface{}) *ClientStreamProtoMock_SendMsg_Call {
	return &ClientStreamProtoMock_SendMsg_Call{Call: _e.mock.On("SendMsg", m)}
}

func (_c *ClientStreamProtoMock_SendMsg_Call) Run(run func(m interface{})) *ClientStreamProtoMock_SendMsg_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(interface{}))
	})
	return _c
}

func (_c *ClientStreamProtoMock_SendMsg_Call) Return(_a0 error) *ClientStreamProtoMock_SendMsg_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ClientStreamProtoMock_SendMsg_Call) RunAndReturn(run func(interface{}) error) *ClientStreamProtoMock_SendMsg_Call {
	_c.Call.Return(run)
	return _c
}

// Trailer provides a mock function with given fields:
func (_m *ClientStreamProtoMock) Trailer() metadata.MD {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Trailer")
	}

	var r0 metadata.MD
	if rf, ok := ret.Get(0).(func() metadata.MD); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(metadata.MD)
		}
	}

	return r0
}

// ClientStreamProtoMock_Trailer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Trailer'
type ClientStreamProtoMock_Trailer_Call struct {
	*mock.Call
}

// Trailer is a helper method to define mock.On call
func (_e *ClientStreamProtoMock_Expecter) Trailer() *ClientStreamProtoMock_Trailer_Call {
	return &ClientStreamProtoMock_Trailer_Call{Call: _e.mock.On("Trailer")}
}

func (_c *ClientStreamProtoMock_Trailer_Call) Run(run func()) *ClientStreamProtoMock_Trailer_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ClientStreamProtoMock_Trailer_Call) Return(_a0 metadata.MD) *ClientStreamProtoMock_Trailer_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ClientStreamProtoMock_Trailer_Call) RunAndReturn(run func() metadata.MD) *ClientStreamProtoMock_Trailer_Call {
	_c.Call.Return(run)
	return _c
}

// NewClientStreamProtoMock creates a new instance of ClientStreamProtoMock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewClientStreamProtoMock(t interface {
	mock.TestingT
	Cleanup(func())
}) *ClientStreamProtoMock {
	mock := &ClientStreamProtoMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
