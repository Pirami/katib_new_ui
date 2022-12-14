// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/kubeflow/katib/pkg/db/v1beta1/common (interfaces: KatibDBInterface)

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	api_v1_beta1 "github.com/kubeflow/katib/pkg/apis/manager/v1beta1"
)

// MockKatibDBInterface is a mock of KatibDBInterface interface.
type MockKatibDBInterface struct {
	ctrl     *gomock.Controller
	recorder *MockKatibDBInterfaceMockRecorder
}

// MockKatibDBInterfaceMockRecorder is the mock recorder for MockKatibDBInterface.
type MockKatibDBInterfaceMockRecorder struct {
	mock *MockKatibDBInterface
}

// NewMockKatibDBInterface creates a new mock instance.
func NewMockKatibDBInterface(ctrl *gomock.Controller) *MockKatibDBInterface {
	mock := &MockKatibDBInterface{ctrl: ctrl}
	mock.recorder = &MockKatibDBInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockKatibDBInterface) EXPECT() *MockKatibDBInterfaceMockRecorder {
	return m.recorder
}

// DBInit mocks base method.
func (m *MockKatibDBInterface) DBInit() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DBInit")
}

// DBInit indicates an expected call of DBInit.
func (mr *MockKatibDBInterfaceMockRecorder) DBInit() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DBInit", reflect.TypeOf((*MockKatibDBInterface)(nil).DBInit))
}

// DeleteObservationLog mocks base method.
func (m *MockKatibDBInterface) DeleteObservationLog(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteObservationLog", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteObservationLog indicates an expected call of DeleteObservationLog.
func (mr *MockKatibDBInterfaceMockRecorder) DeleteObservationLog(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteObservationLog", reflect.TypeOf((*MockKatibDBInterface)(nil).DeleteObservationLog), arg0)
}

// GetObservationLog mocks base method.
func (m *MockKatibDBInterface) GetObservationLog(arg0, arg1, arg2, arg3 string) (*api_v1_beta1.ObservationLog, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetObservationLog", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*api_v1_beta1.ObservationLog)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetObservationLog indicates an expected call of GetObservationLog.
func (mr *MockKatibDBInterfaceMockRecorder) GetObservationLog(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetObservationLog", reflect.TypeOf((*MockKatibDBInterface)(nil).GetObservationLog), arg0, arg1, arg2, arg3)
}

// RegisterObservationLog mocks base method.
func (m *MockKatibDBInterface) RegisterObservationLog(arg0 string, arg1 *api_v1_beta1.ObservationLog) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterObservationLog", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RegisterObservationLog indicates an expected call of RegisterObservationLog.
func (mr *MockKatibDBInterfaceMockRecorder) RegisterObservationLog(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterObservationLog", reflect.TypeOf((*MockKatibDBInterface)(nil).RegisterObservationLog), arg0, arg1)
}

// SelectOne mocks base method.
func (m *MockKatibDBInterface) SelectOne() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectOne")
	ret0, _ := ret[0].(error)
	return ret0
}

// SelectOne indicates an expected call of SelectOne.
func (mr *MockKatibDBInterfaceMockRecorder) SelectOne() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectOne", reflect.TypeOf((*MockKatibDBInterface)(nil).SelectOne))
}
