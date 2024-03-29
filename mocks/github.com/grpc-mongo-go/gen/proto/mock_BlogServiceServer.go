// Code generated by mockery. DO NOT EDIT.

package proto

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	emptypb "google.golang.org/protobuf/types/known/emptypb"

	proto "github.com/grpc-mongo-go/gen/proto"
)

// MockBlogServiceServer is an autogenerated mock type for the BlogServiceServer type
type MockBlogServiceServer struct {
	mock.Mock
}

type MockBlogServiceServer_Expecter struct {
	mock *mock.Mock
}

func (_m *MockBlogServiceServer) EXPECT() *MockBlogServiceServer_Expecter {
	return &MockBlogServiceServer_Expecter{mock: &_m.Mock}
}

// CreateBlog provides a mock function with given fields: _a0, _a1
func (_m *MockBlogServiceServer) CreateBlog(_a0 context.Context, _a1 *proto.Blog) (*proto.BlogId, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for CreateBlog")
	}

	var r0 *proto.BlogId
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *proto.Blog) (*proto.BlogId, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *proto.Blog) *proto.BlogId); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*proto.BlogId)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *proto.Blog) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockBlogServiceServer_CreateBlog_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateBlog'
type MockBlogServiceServer_CreateBlog_Call struct {
	*mock.Call
}

// CreateBlog is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *proto.Blog
func (_e *MockBlogServiceServer_Expecter) CreateBlog(_a0 interface{}, _a1 interface{}) *MockBlogServiceServer_CreateBlog_Call {
	return &MockBlogServiceServer_CreateBlog_Call{Call: _e.mock.On("CreateBlog", _a0, _a1)}
}

func (_c *MockBlogServiceServer_CreateBlog_Call) Run(run func(_a0 context.Context, _a1 *proto.Blog)) *MockBlogServiceServer_CreateBlog_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*proto.Blog))
	})
	return _c
}

func (_c *MockBlogServiceServer_CreateBlog_Call) Return(_a0 *proto.BlogId, _a1 error) *MockBlogServiceServer_CreateBlog_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockBlogServiceServer_CreateBlog_Call) RunAndReturn(run func(context.Context, *proto.Blog) (*proto.BlogId, error)) *MockBlogServiceServer_CreateBlog_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteBlog provides a mock function with given fields: _a0, _a1
func (_m *MockBlogServiceServer) DeleteBlog(_a0 context.Context, _a1 *proto.BlogId) (*emptypb.Empty, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for DeleteBlog")
	}

	var r0 *emptypb.Empty
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *proto.BlogId) (*emptypb.Empty, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *proto.BlogId) *emptypb.Empty); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*emptypb.Empty)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *proto.BlogId) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockBlogServiceServer_DeleteBlog_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteBlog'
type MockBlogServiceServer_DeleteBlog_Call struct {
	*mock.Call
}

// DeleteBlog is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *proto.BlogId
func (_e *MockBlogServiceServer_Expecter) DeleteBlog(_a0 interface{}, _a1 interface{}) *MockBlogServiceServer_DeleteBlog_Call {
	return &MockBlogServiceServer_DeleteBlog_Call{Call: _e.mock.On("DeleteBlog", _a0, _a1)}
}

func (_c *MockBlogServiceServer_DeleteBlog_Call) Run(run func(_a0 context.Context, _a1 *proto.BlogId)) *MockBlogServiceServer_DeleteBlog_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*proto.BlogId))
	})
	return _c
}

func (_c *MockBlogServiceServer_DeleteBlog_Call) Return(_a0 *emptypb.Empty, _a1 error) *MockBlogServiceServer_DeleteBlog_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockBlogServiceServer_DeleteBlog_Call) RunAndReturn(run func(context.Context, *proto.BlogId) (*emptypb.Empty, error)) *MockBlogServiceServer_DeleteBlog_Call {
	_c.Call.Return(run)
	return _c
}

// ListBlogs provides a mock function with given fields: _a0, _a1
func (_m *MockBlogServiceServer) ListBlogs(_a0 *emptypb.Empty, _a1 proto.BlogService_ListBlogsServer) error {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for ListBlogs")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*emptypb.Empty, proto.BlogService_ListBlogsServer) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockBlogServiceServer_ListBlogs_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListBlogs'
type MockBlogServiceServer_ListBlogs_Call struct {
	*mock.Call
}

// ListBlogs is a helper method to define mock.On call
//   - _a0 *emptypb.Empty
//   - _a1 proto.BlogService_ListBlogsServer
func (_e *MockBlogServiceServer_Expecter) ListBlogs(_a0 interface{}, _a1 interface{}) *MockBlogServiceServer_ListBlogs_Call {
	return &MockBlogServiceServer_ListBlogs_Call{Call: _e.mock.On("ListBlogs", _a0, _a1)}
}

func (_c *MockBlogServiceServer_ListBlogs_Call) Run(run func(_a0 *emptypb.Empty, _a1 proto.BlogService_ListBlogsServer)) *MockBlogServiceServer_ListBlogs_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*emptypb.Empty), args[1].(proto.BlogService_ListBlogsServer))
	})
	return _c
}

func (_c *MockBlogServiceServer_ListBlogs_Call) Return(_a0 error) *MockBlogServiceServer_ListBlogs_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockBlogServiceServer_ListBlogs_Call) RunAndReturn(run func(*emptypb.Empty, proto.BlogService_ListBlogsServer) error) *MockBlogServiceServer_ListBlogs_Call {
	_c.Call.Return(run)
	return _c
}

// ReadBlog provides a mock function with given fields: _a0, _a1
func (_m *MockBlogServiceServer) ReadBlog(_a0 context.Context, _a1 *proto.BlogId) (*proto.Blog, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for ReadBlog")
	}

	var r0 *proto.Blog
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *proto.BlogId) (*proto.Blog, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *proto.BlogId) *proto.Blog); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*proto.Blog)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *proto.BlogId) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockBlogServiceServer_ReadBlog_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ReadBlog'
type MockBlogServiceServer_ReadBlog_Call struct {
	*mock.Call
}

// ReadBlog is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *proto.BlogId
func (_e *MockBlogServiceServer_Expecter) ReadBlog(_a0 interface{}, _a1 interface{}) *MockBlogServiceServer_ReadBlog_Call {
	return &MockBlogServiceServer_ReadBlog_Call{Call: _e.mock.On("ReadBlog", _a0, _a1)}
}

func (_c *MockBlogServiceServer_ReadBlog_Call) Run(run func(_a0 context.Context, _a1 *proto.BlogId)) *MockBlogServiceServer_ReadBlog_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*proto.BlogId))
	})
	return _c
}

func (_c *MockBlogServiceServer_ReadBlog_Call) Return(_a0 *proto.Blog, _a1 error) *MockBlogServiceServer_ReadBlog_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockBlogServiceServer_ReadBlog_Call) RunAndReturn(run func(context.Context, *proto.BlogId) (*proto.Blog, error)) *MockBlogServiceServer_ReadBlog_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateBlog provides a mock function with given fields: _a0, _a1
func (_m *MockBlogServiceServer) UpdateBlog(_a0 context.Context, _a1 *proto.Blog) (*emptypb.Empty, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for UpdateBlog")
	}

	var r0 *emptypb.Empty
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *proto.Blog) (*emptypb.Empty, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *proto.Blog) *emptypb.Empty); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*emptypb.Empty)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *proto.Blog) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockBlogServiceServer_UpdateBlog_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateBlog'
type MockBlogServiceServer_UpdateBlog_Call struct {
	*mock.Call
}

// UpdateBlog is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *proto.Blog
func (_e *MockBlogServiceServer_Expecter) UpdateBlog(_a0 interface{}, _a1 interface{}) *MockBlogServiceServer_UpdateBlog_Call {
	return &MockBlogServiceServer_UpdateBlog_Call{Call: _e.mock.On("UpdateBlog", _a0, _a1)}
}

func (_c *MockBlogServiceServer_UpdateBlog_Call) Run(run func(_a0 context.Context, _a1 *proto.Blog)) *MockBlogServiceServer_UpdateBlog_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*proto.Blog))
	})
	return _c
}

func (_c *MockBlogServiceServer_UpdateBlog_Call) Return(_a0 *emptypb.Empty, _a1 error) *MockBlogServiceServer_UpdateBlog_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockBlogServiceServer_UpdateBlog_Call) RunAndReturn(run func(context.Context, *proto.Blog) (*emptypb.Empty, error)) *MockBlogServiceServer_UpdateBlog_Call {
	_c.Call.Return(run)
	return _c
}

// mustEmbedUnimplementedBlogServiceServer provides a mock function with given fields:
func (_m *MockBlogServiceServer) mustEmbedUnimplementedBlogServiceServer() {
	_m.Called()
}

// MockBlogServiceServer_mustEmbedUnimplementedBlogServiceServer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'mustEmbedUnimplementedBlogServiceServer'
type MockBlogServiceServer_mustEmbedUnimplementedBlogServiceServer_Call struct {
	*mock.Call
}

// mustEmbedUnimplementedBlogServiceServer is a helper method to define mock.On call
func (_e *MockBlogServiceServer_Expecter) mustEmbedUnimplementedBlogServiceServer() *MockBlogServiceServer_mustEmbedUnimplementedBlogServiceServer_Call {
	return &MockBlogServiceServer_mustEmbedUnimplementedBlogServiceServer_Call{Call: _e.mock.On("mustEmbedUnimplementedBlogServiceServer")}
}

func (_c *MockBlogServiceServer_mustEmbedUnimplementedBlogServiceServer_Call) Run(run func()) *MockBlogServiceServer_mustEmbedUnimplementedBlogServiceServer_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockBlogServiceServer_mustEmbedUnimplementedBlogServiceServer_Call) Return() *MockBlogServiceServer_mustEmbedUnimplementedBlogServiceServer_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockBlogServiceServer_mustEmbedUnimplementedBlogServiceServer_Call) RunAndReturn(run func()) *MockBlogServiceServer_mustEmbedUnimplementedBlogServiceServer_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockBlogServiceServer creates a new instance of MockBlogServiceServer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockBlogServiceServer(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockBlogServiceServer {
	mock := &MockBlogServiceServer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
