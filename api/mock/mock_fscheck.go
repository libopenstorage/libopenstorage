// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/libopenstorage/openstorage/api (interfaces: OpenStorageFilesystemCheckServer,OpenStorageFilesystemCheckClient)

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	api "github.com/libopenstorage/openstorage/api"
	grpc "google.golang.org/grpc"
	reflect "reflect"
)

// MockOpenStorageFilesystemCheckServer is a mock of OpenStorageFilesystemCheckServer interface.
type MockOpenStorageFilesystemCheckServer struct {
	ctrl     *gomock.Controller
	recorder *MockOpenStorageFilesystemCheckServerMockRecorder
}

// MockOpenStorageFilesystemCheckServerMockRecorder is the mock recorder for MockOpenStorageFilesystemCheckServer.
type MockOpenStorageFilesystemCheckServerMockRecorder struct {
	mock *MockOpenStorageFilesystemCheckServer
}

// NewMockOpenStorageFilesystemCheckServer creates a new mock instance.
func NewMockOpenStorageFilesystemCheckServer(ctrl *gomock.Controller) *MockOpenStorageFilesystemCheckServer {
	mock := &MockOpenStorageFilesystemCheckServer{ctrl: ctrl}
	mock.recorder = &MockOpenStorageFilesystemCheckServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOpenStorageFilesystemCheckServer) EXPECT() *MockOpenStorageFilesystemCheckServerMockRecorder {
	return m.recorder
}

// CheckHealth mocks base method.
func (m *MockOpenStorageFilesystemCheckServer) CheckHealth(arg0 context.Context, arg1 *api.SdkFilesystemCheckCheckHealthRequest) (*api.SdkFilesystemCheckCheckHealthResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckHealth", arg0, arg1)
	ret0, _ := ret[0].(*api.SdkFilesystemCheckCheckHealthResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckHealth indicates an expected call of CheckHealth.
func (mr *MockOpenStorageFilesystemCheckServerMockRecorder) CheckHealth(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckHealth", reflect.TypeOf((*MockOpenStorageFilesystemCheckServer)(nil).CheckHealth), arg0, arg1)
}

// CheckHealthGetStatus mocks base method.
func (m *MockOpenStorageFilesystemCheckServer) CheckHealthGetStatus(arg0 context.Context, arg1 *api.SdkFilesystemCheckCheckHealthGetStatusRequest) (*api.SdkFilesystemCheckCheckHealthGetStatusResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckHealthGetStatus", arg0, arg1)
	ret0, _ := ret[0].(*api.SdkFilesystemCheckCheckHealthGetStatusResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckHealthGetStatus indicates an expected call of CheckHealthGetStatus.
func (mr *MockOpenStorageFilesystemCheckServerMockRecorder) CheckHealthGetStatus(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckHealthGetStatus", reflect.TypeOf((*MockOpenStorageFilesystemCheckServer)(nil).CheckHealthGetStatus), arg0, arg1)
}

// FixAll mocks base method.
func (m *MockOpenStorageFilesystemCheckServer) FixAll(arg0 context.Context, arg1 *api.SdkFilesystemCheckFixAllRequest) (*api.SdkFilesystemCheckFixAllResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FixAll", arg0, arg1)
	ret0, _ := ret[0].(*api.SdkFilesystemCheckFixAllResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FixAll indicates an expected call of FixAll.
func (mr *MockOpenStorageFilesystemCheckServerMockRecorder) FixAll(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FixAll", reflect.TypeOf((*MockOpenStorageFilesystemCheckServer)(nil).FixAll), arg0, arg1)
}

// FixAllGetStatus mocks base method.
func (m *MockOpenStorageFilesystemCheckServer) FixAllGetStatus(arg0 context.Context, arg1 *api.SdkFilesystemCheckFixAllGetStatusRequest) (*api.SdkFilesystemCheckFixAllGetStatusResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FixAllGetStatus", arg0, arg1)
	ret0, _ := ret[0].(*api.SdkFilesystemCheckFixAllGetStatusResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FixAllGetStatus indicates an expected call of FixAllGetStatus.
func (mr *MockOpenStorageFilesystemCheckServerMockRecorder) FixAllGetStatus(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FixAllGetStatus", reflect.TypeOf((*MockOpenStorageFilesystemCheckServer)(nil).FixAllGetStatus), arg0, arg1)
}

// Stop mocks base method.
func (m *MockOpenStorageFilesystemCheckServer) Stop(arg0 context.Context, arg1 *api.SdkFilesystemCheckStopRequest) (*api.SdkFilesystemCheckStopResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stop", arg0, arg1)
	ret0, _ := ret[0].(*api.SdkFilesystemCheckStopResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Stop indicates an expected call of Stop.
func (mr *MockOpenStorageFilesystemCheckServerMockRecorder) Stop(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockOpenStorageFilesystemCheckServer)(nil).Stop), arg0, arg1)
}

// MockOpenStorageFilesystemCheckClient is a mock of OpenStorageFilesystemCheckClient interface.
type MockOpenStorageFilesystemCheckClient struct {
	ctrl     *gomock.Controller
	recorder *MockOpenStorageFilesystemCheckClientMockRecorder
}

// MockOpenStorageFilesystemCheckClientMockRecorder is the mock recorder for MockOpenStorageFilesystemCheckClient.
type MockOpenStorageFilesystemCheckClientMockRecorder struct {
	mock *MockOpenStorageFilesystemCheckClient
}

// NewMockOpenStorageFilesystemCheckClient creates a new mock instance.
func NewMockOpenStorageFilesystemCheckClient(ctrl *gomock.Controller) *MockOpenStorageFilesystemCheckClient {
	mock := &MockOpenStorageFilesystemCheckClient{ctrl: ctrl}
	mock.recorder = &MockOpenStorageFilesystemCheckClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOpenStorageFilesystemCheckClient) EXPECT() *MockOpenStorageFilesystemCheckClientMockRecorder {
	return m.recorder
}

// CheckHealth mocks base method.
func (m *MockOpenStorageFilesystemCheckClient) CheckHealth(arg0 context.Context, arg1 *api.SdkFilesystemCheckCheckHealthRequest, arg2 ...grpc.CallOption) (*api.SdkFilesystemCheckCheckHealthResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CheckHealth", varargs...)
	ret0, _ := ret[0].(*api.SdkFilesystemCheckCheckHealthResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckHealth indicates an expected call of CheckHealth.
func (mr *MockOpenStorageFilesystemCheckClientMockRecorder) CheckHealth(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckHealth", reflect.TypeOf((*MockOpenStorageFilesystemCheckClient)(nil).CheckHealth), varargs...)
}

// CheckHealthGetStatus mocks base method.
func (m *MockOpenStorageFilesystemCheckClient) CheckHealthGetStatus(arg0 context.Context, arg1 *api.SdkFilesystemCheckCheckHealthGetStatusRequest, arg2 ...grpc.CallOption) (*api.SdkFilesystemCheckCheckHealthGetStatusResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CheckHealthGetStatus", varargs...)
	ret0, _ := ret[0].(*api.SdkFilesystemCheckCheckHealthGetStatusResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckHealthGetStatus indicates an expected call of CheckHealthGetStatus.
func (mr *MockOpenStorageFilesystemCheckClientMockRecorder) CheckHealthGetStatus(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckHealthGetStatus", reflect.TypeOf((*MockOpenStorageFilesystemCheckClient)(nil).CheckHealthGetStatus), varargs...)
}

// FixAll mocks base method.
func (m *MockOpenStorageFilesystemCheckClient) FixAll(arg0 context.Context, arg1 *api.SdkFilesystemCheckFixAllRequest, arg2 ...grpc.CallOption) (*api.SdkFilesystemCheckFixAllResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FixAll", varargs...)
	ret0, _ := ret[0].(*api.SdkFilesystemCheckFixAllResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FixAll indicates an expected call of FixAll.
func (mr *MockOpenStorageFilesystemCheckClientMockRecorder) FixAll(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FixAll", reflect.TypeOf((*MockOpenStorageFilesystemCheckClient)(nil).FixAll), varargs...)
}

// FixAllGetStatus mocks base method.
func (m *MockOpenStorageFilesystemCheckClient) FixAllGetStatus(arg0 context.Context, arg1 *api.SdkFilesystemCheckFixAllGetStatusRequest, arg2 ...grpc.CallOption) (*api.SdkFilesystemCheckFixAllGetStatusResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FixAllGetStatus", varargs...)
	ret0, _ := ret[0].(*api.SdkFilesystemCheckFixAllGetStatusResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FixAllGetStatus indicates an expected call of FixAllGetStatus.
func (mr *MockOpenStorageFilesystemCheckClientMockRecorder) FixAllGetStatus(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FixAllGetStatus", reflect.TypeOf((*MockOpenStorageFilesystemCheckClient)(nil).FixAllGetStatus), varargs...)
}

// Stop mocks base method.
func (m *MockOpenStorageFilesystemCheckClient) Stop(arg0 context.Context, arg1 *api.SdkFilesystemCheckStopRequest, arg2 ...grpc.CallOption) (*api.SdkFilesystemCheckStopResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Stop", varargs...)
	ret0, _ := ret[0].(*api.SdkFilesystemCheckStopResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Stop indicates an expected call of Stop.
func (mr *MockOpenStorageFilesystemCheckClientMockRecorder) Stop(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockOpenStorageFilesystemCheckClient)(nil).Stop), varargs...)
}
