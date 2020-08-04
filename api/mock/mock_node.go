// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/libopenstorage/openstorage/api (interfaces: OpenStorageNodeServer,OpenStorageNodeClient)

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	api "github.com/libopenstorage/openstorage/api"
	grpc "google.golang.org/grpc"
)

// MockOpenStorageNodeServer is a mock of OpenStorageNodeServer interface.
type MockOpenStorageNodeServer struct {
	ctrl     *gomock.Controller
	recorder *MockOpenStorageNodeServerMockRecorder
}

// MockOpenStorageNodeServerMockRecorder is the mock recorder for MockOpenStorageNodeServer.
type MockOpenStorageNodeServerMockRecorder struct {
	mock *MockOpenStorageNodeServer
}

// NewMockOpenStorageNodeServer creates a new mock instance.
func NewMockOpenStorageNodeServer(ctrl *gomock.Controller) *MockOpenStorageNodeServer {
	mock := &MockOpenStorageNodeServer{ctrl: ctrl}
	mock.recorder = &MockOpenStorageNodeServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOpenStorageNodeServer) EXPECT() *MockOpenStorageNodeServerMockRecorder {
	return m.recorder
}

// Enumerate mocks base method.
func (m *MockOpenStorageNodeServer) Enumerate(arg0 context.Context, arg1 *api.SdkNodeEnumerateRequest) (*api.SdkNodeEnumerateResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Enumerate", arg0, arg1)
	ret0, _ := ret[0].(*api.SdkNodeEnumerateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Enumerate indicates an expected call of Enumerate.
func (mr *MockOpenStorageNodeServerMockRecorder) Enumerate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Enumerate", reflect.TypeOf((*MockOpenStorageNodeServer)(nil).Enumerate), arg0, arg1)
}

// EnumerateWithFilters mocks base method.
func (m *MockOpenStorageNodeServer) EnumerateWithFilters(arg0 context.Context, arg1 *api.SdkNodeEnumerateWithFiltersRequest) (*api.SdkNodeEnumerateWithFiltersResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnumerateWithFilters", arg0, arg1)
	ret0, _ := ret[0].(*api.SdkNodeEnumerateWithFiltersResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EnumerateWithFilters indicates an expected call of EnumerateWithFilters.
func (mr *MockOpenStorageNodeServerMockRecorder) EnumerateWithFilters(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnumerateWithFilters", reflect.TypeOf((*MockOpenStorageNodeServer)(nil).EnumerateWithFilters), arg0, arg1)
}

// Inspect mocks base method.
func (m *MockOpenStorageNodeServer) Inspect(arg0 context.Context, arg1 *api.SdkNodeInspectRequest) (*api.SdkNodeInspectResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Inspect", arg0, arg1)
	ret0, _ := ret[0].(*api.SdkNodeInspectResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Inspect indicates an expected call of Inspect.
func (mr *MockOpenStorageNodeServerMockRecorder) Inspect(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Inspect", reflect.TypeOf((*MockOpenStorageNodeServer)(nil).Inspect), arg0, arg1)
}

// InspectCurrent mocks base method.
func (m *MockOpenStorageNodeServer) InspectCurrent(arg0 context.Context, arg1 *api.SdkNodeInspectCurrentRequest) (*api.SdkNodeInspectCurrentResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InspectCurrent", arg0, arg1)
	ret0, _ := ret[0].(*api.SdkNodeInspectCurrentResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InspectCurrent indicates an expected call of InspectCurrent.
func (mr *MockOpenStorageNodeServerMockRecorder) InspectCurrent(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InspectCurrent", reflect.TypeOf((*MockOpenStorageNodeServer)(nil).InspectCurrent), arg0, arg1)
}

// MockOpenStorageNodeClient is a mock of OpenStorageNodeClient interface.
type MockOpenStorageNodeClient struct {
	ctrl     *gomock.Controller
	recorder *MockOpenStorageNodeClientMockRecorder
}

// MockOpenStorageNodeClientMockRecorder is the mock recorder for MockOpenStorageNodeClient.
type MockOpenStorageNodeClientMockRecorder struct {
	mock *MockOpenStorageNodeClient
}

// NewMockOpenStorageNodeClient creates a new mock instance.
func NewMockOpenStorageNodeClient(ctrl *gomock.Controller) *MockOpenStorageNodeClient {
	mock := &MockOpenStorageNodeClient{ctrl: ctrl}
	mock.recorder = &MockOpenStorageNodeClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOpenStorageNodeClient) EXPECT() *MockOpenStorageNodeClientMockRecorder {
	return m.recorder
}

// Enumerate mocks base method.
func (m *MockOpenStorageNodeClient) Enumerate(arg0 context.Context, arg1 *api.SdkNodeEnumerateRequest, arg2 ...grpc.CallOption) (*api.SdkNodeEnumerateResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Enumerate", varargs...)
	ret0, _ := ret[0].(*api.SdkNodeEnumerateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Enumerate indicates an expected call of Enumerate.
func (mr *MockOpenStorageNodeClientMockRecorder) Enumerate(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Enumerate", reflect.TypeOf((*MockOpenStorageNodeClient)(nil).Enumerate), varargs...)
}

// EnumerateWithFilters mocks base method.
func (m *MockOpenStorageNodeClient) EnumerateWithFilters(arg0 context.Context, arg1 *api.SdkNodeEnumerateWithFiltersRequest, arg2 ...grpc.CallOption) (*api.SdkNodeEnumerateWithFiltersResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "EnumerateWithFilters", varargs...)
	ret0, _ := ret[0].(*api.SdkNodeEnumerateWithFiltersResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EnumerateWithFilters indicates an expected call of EnumerateWithFilters.
func (mr *MockOpenStorageNodeClientMockRecorder) EnumerateWithFilters(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnumerateWithFilters", reflect.TypeOf((*MockOpenStorageNodeClient)(nil).EnumerateWithFilters), varargs...)
}

// Inspect mocks base method.
func (m *MockOpenStorageNodeClient) Inspect(arg0 context.Context, arg1 *api.SdkNodeInspectRequest, arg2 ...grpc.CallOption) (*api.SdkNodeInspectResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Inspect", varargs...)
	ret0, _ := ret[0].(*api.SdkNodeInspectResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Inspect indicates an expected call of Inspect.
func (mr *MockOpenStorageNodeClientMockRecorder) Inspect(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Inspect", reflect.TypeOf((*MockOpenStorageNodeClient)(nil).Inspect), varargs...)
}

// InspectCurrent mocks base method.
func (m *MockOpenStorageNodeClient) InspectCurrent(arg0 context.Context, arg1 *api.SdkNodeInspectCurrentRequest, arg2 ...grpc.CallOption) (*api.SdkNodeInspectCurrentResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "InspectCurrent", varargs...)
	ret0, _ := ret[0].(*api.SdkNodeInspectCurrentResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InspectCurrent indicates an expected call of InspectCurrent.
func (mr *MockOpenStorageNodeClientMockRecorder) InspectCurrent(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InspectCurrent", reflect.TypeOf((*MockOpenStorageNodeClient)(nil).InspectCurrent), varargs...)
}
