// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/service/iservice.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	openfeature "github.com/open-feature/go-sdk/pkg/openfeature"
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

// EventChannel mocks base method.
func (m *MockIService) EventChannel() <-chan openfeature.Event {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EventChannel")
	ret0, _ := ret[0].(<-chan openfeature.Event)
	return ret0
}

// EventChannel indicates an expected call of EventChannel.
func (mr *MockIServiceMockRecorder) EventChannel() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EventChannel", reflect.TypeOf((*MockIService)(nil).EventChannel))
}

// Init mocks base method.
func (m *MockIService) Init() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Init")
	ret0, _ := ret[0].(error)
	return ret0
}

// Init indicates an expected call of Init.
func (mr *MockIServiceMockRecorder) Init() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Init", reflect.TypeOf((*MockIService)(nil).Init))
}

// ResolveBoolean mocks base method.
func (m *MockIService) ResolveBoolean(ctx context.Context, key string, defaultValue bool, evalCtx map[string]interface{}) openfeature.BoolResolutionDetail {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResolveBoolean", ctx, key, defaultValue, evalCtx)
	ret0, _ := ret[0].(openfeature.BoolResolutionDetail)
	return ret0
}

// ResolveBoolean indicates an expected call of ResolveBoolean.
func (mr *MockIServiceMockRecorder) ResolveBoolean(ctx, key, defaultValue, evalCtx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResolveBoolean", reflect.TypeOf((*MockIService)(nil).ResolveBoolean), ctx, key, defaultValue, evalCtx)
}

// ResolveFloat mocks base method.
func (m *MockIService) ResolveFloat(ctx context.Context, key string, defaultValue float64, evalCtx map[string]interface{}) openfeature.FloatResolutionDetail {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResolveFloat", ctx, key, defaultValue, evalCtx)
	ret0, _ := ret[0].(openfeature.FloatResolutionDetail)
	return ret0
}

// ResolveFloat indicates an expected call of ResolveFloat.
func (mr *MockIServiceMockRecorder) ResolveFloat(ctx, key, defaultValue, evalCtx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResolveFloat", reflect.TypeOf((*MockIService)(nil).ResolveFloat), ctx, key, defaultValue, evalCtx)
}

// ResolveInt mocks base method.
func (m *MockIService) ResolveInt(ctx context.Context, key string, defaultValue int64, evalCtx map[string]interface{}) openfeature.IntResolutionDetail {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResolveInt", ctx, key, defaultValue, evalCtx)
	ret0, _ := ret[0].(openfeature.IntResolutionDetail)
	return ret0
}

// ResolveInt indicates an expected call of ResolveInt.
func (mr *MockIServiceMockRecorder) ResolveInt(ctx, key, defaultValue, evalCtx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResolveInt", reflect.TypeOf((*MockIService)(nil).ResolveInt), ctx, key, defaultValue, evalCtx)
}

// ResolveObject mocks base method.
func (m *MockIService) ResolveObject(ctx context.Context, key string, defaultValue interface{}, evalCtx map[string]interface{}) openfeature.InterfaceResolutionDetail {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResolveObject", ctx, key, defaultValue, evalCtx)
	ret0, _ := ret[0].(openfeature.InterfaceResolutionDetail)
	return ret0
}

// ResolveObject indicates an expected call of ResolveObject.
func (mr *MockIServiceMockRecorder) ResolveObject(ctx, key, defaultValue, evalCtx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResolveObject", reflect.TypeOf((*MockIService)(nil).ResolveObject), ctx, key, defaultValue, evalCtx)
}

// ResolveString mocks base method.
func (m *MockIService) ResolveString(ctx context.Context, key, defaultValue string, evalCtx map[string]interface{}) openfeature.StringResolutionDetail {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResolveString", ctx, key, defaultValue, evalCtx)
	ret0, _ := ret[0].(openfeature.StringResolutionDetail)
	return ret0
}

// ResolveString indicates an expected call of ResolveString.
func (mr *MockIServiceMockRecorder) ResolveString(ctx, key, defaultValue, evalCtx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResolveString", reflect.TypeOf((*MockIService)(nil).ResolveString), ctx, key, defaultValue, evalCtx)
}

// Shutdown mocks base method.
func (m *MockIService) Shutdown() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Shutdown")
}

// Shutdown indicates an expected call of Shutdown.
func (mr *MockIServiceMockRecorder) Shutdown() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Shutdown", reflect.TypeOf((*MockIService)(nil).Shutdown))
}
