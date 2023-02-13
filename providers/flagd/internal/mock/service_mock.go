// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/service/iservice.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	schemav1 "buf.build/gen/go/open-feature/flagd/protocolbuffers/go/schema/v1"
	gomock "github.com/golang/mock/gomock"
)

// MockIService is a mock of IService interface.
type MockIService struct {
	ctrl     *gomock.Controller
	recorder *MockIServiceMockRecorder
}

// MockIServiceMockRecorder is the mock recorder for MockIService.
type MockIServiceMockRecorder struct {
	mock *MockIService
}

// NewMockIService creates a new mock instance.
func NewMockIService(ctrl *gomock.Controller) *MockIService {
	mock := &MockIService{ctrl: ctrl}
	mock.recorder = &MockIServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIService) EXPECT() *MockIServiceMockRecorder {
	return m.recorder
}

// EventStream mocks base method.
func (m *MockIService) EventStream(ctx context.Context, eventChan chan<- *schemav1.EventStreamResponse, maxAttempts int, errChan chan<- error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "EventStream", ctx, eventChan, maxAttempts, errChan)
}

// EventStream indicates an expected call of EventStream.
func (mr *MockIServiceMockRecorder) EventStream(ctx, eventChan, maxAttempts, errChan interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EventStream", reflect.TypeOf((*MockIService)(nil).EventStream), ctx, eventChan, maxAttempts, errChan)
}

// IsEventStreamAlive mocks base method.
func (m *MockIService) IsEventStreamAlive() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsEventStreamAlive")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsEventStreamAlive indicates an expected call of IsEventStreamAlive.
func (mr *MockIServiceMockRecorder) IsEventStreamAlive() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsEventStreamAlive", reflect.TypeOf((*MockIService)(nil).IsEventStreamAlive))
}

// ResolveBoolean mocks base method.
func (m *MockIService) ResolveBoolean(arg0 context.Context, arg1 string, arg2 map[string]interface{}) (*schemav1.ResolveBooleanResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResolveBoolean", arg0, arg1, arg2)
	ret0, _ := ret[0].(*schemav1.ResolveBooleanResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ResolveBoolean indicates an expected call of ResolveBoolean.
func (mr *MockIServiceMockRecorder) ResolveBoolean(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResolveBoolean", reflect.TypeOf((*MockIService)(nil).ResolveBoolean), arg0, arg1, arg2)
}

// ResolveFloat mocks base method.
func (m *MockIService) ResolveFloat(arg0 context.Context, arg1 string, arg2 map[string]interface{}) (*schemav1.ResolveFloatResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResolveFloat", arg0, arg1, arg2)
	ret0, _ := ret[0].(*schemav1.ResolveFloatResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ResolveFloat indicates an expected call of ResolveFloat.
func (mr *MockIServiceMockRecorder) ResolveFloat(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResolveFloat", reflect.TypeOf((*MockIService)(nil).ResolveFloat), arg0, arg1, arg2)
}

// ResolveInt mocks base method.
func (m *MockIService) ResolveInt(arg0 context.Context, arg1 string, arg2 map[string]interface{}) (*schemav1.ResolveIntResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResolveInt", arg0, arg1, arg2)
	ret0, _ := ret[0].(*schemav1.ResolveIntResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ResolveInt indicates an expected call of ResolveInt.
func (mr *MockIServiceMockRecorder) ResolveInt(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResolveInt", reflect.TypeOf((*MockIService)(nil).ResolveInt), arg0, arg1, arg2)
}

// ResolveObject mocks base method.
func (m *MockIService) ResolveObject(arg0 context.Context, arg1 string, arg2 map[string]interface{}) (*schemav1.ResolveObjectResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResolveObject", arg0, arg1, arg2)
	ret0, _ := ret[0].(*schemav1.ResolveObjectResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ResolveObject indicates an expected call of ResolveObject.
func (mr *MockIServiceMockRecorder) ResolveObject(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResolveObject", reflect.TypeOf((*MockIService)(nil).ResolveObject), arg0, arg1, arg2)
}

// ResolveString mocks base method.
func (m *MockIService) ResolveString(arg0 context.Context, arg1 string, arg2 map[string]interface{}) (*schemav1.ResolveStringResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResolveString", arg0, arg1, arg2)
	ret0, _ := ret[0].(*schemav1.ResolveStringResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ResolveString indicates an expected call of ResolveString.
func (mr *MockIServiceMockRecorder) ResolveString(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResolveString", reflect.TypeOf((*MockIService)(nil).ResolveString), arg0, arg1, arg2)
}