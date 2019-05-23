// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/kubeflow/katib/pkg/controller/v1alpha2/experiment/managerclient (interfaces: ManagerClient)

// Package mock is a generated GoMock package.
package mock

import (
	gomock "github.com/golang/mock/gomock"
	v1alpha2 "github.com/kubeflow/katib/pkg/api/operators/apis/experiment/v1alpha2"
	v1alpha20 "github.com/kubeflow/katib/pkg/api/v1alpha2"
	reflect "reflect"
)

// MockManagerClient is a mock of ManagerClient interface
type MockManagerClient struct {
	ctrl     *gomock.Controller
	recorder *MockManagerClientMockRecorder
}

// MockManagerClientMockRecorder is the mock recorder for MockManagerClient
type MockManagerClientMockRecorder struct {
	mock *MockManagerClient
}

// NewMockManagerClient creates a new mock instance
func NewMockManagerClient(ctrl *gomock.Controller) *MockManagerClient {
	mock := &MockManagerClient{ctrl: ctrl}
	mock.recorder = &MockManagerClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockManagerClient) EXPECT() *MockManagerClientMockRecorder {
	return m.recorder
}

// CreateExperimentInDB mocks base method
func (m *MockManagerClient) CreateExperimentInDB(arg0 *v1alpha2.Experiment) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateExperimentInDB", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateExperimentInDB indicates an expected call of CreateExperimentInDB
func (mr *MockManagerClientMockRecorder) CreateExperimentInDB(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateExperimentInDB", reflect.TypeOf((*MockManagerClient)(nil).CreateExperimentInDB), arg0)
}

// DeleteExperimentInDB mocks base method
func (m *MockManagerClient) DeleteExperimentInDB(arg0 *v1alpha2.Experiment) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteExperimentInDB", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteExperimentInDB indicates an expected call of DeleteExperimentInDB
func (mr *MockManagerClientMockRecorder) DeleteExperimentInDB(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteExperimentInDB", reflect.TypeOf((*MockManagerClient)(nil).DeleteExperimentInDB), arg0)
}

// PreCheckRegisterExperimentInDB mocks base method
func (m *MockManagerClient) PreCheckRegisterExperimentInDB(arg0 *v1alpha2.Experiment) (*v1alpha20.PreCheckRegisterExperimentReply, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PreCheckRegisterExperimentInDB", arg0)
	ret0, _ := ret[0].(*v1alpha20.PreCheckRegisterExperimentReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PreCheckRegisterExperimentInDB indicates an expected call of PreCheckRegisterExperimentInDB
func (mr *MockManagerClientMockRecorder) PreCheckRegisterExperimentInDB(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PreCheckRegisterExperimentInDB", reflect.TypeOf((*MockManagerClient)(nil).PreCheckRegisterExperimentInDB), arg0)
}

// UpdateExperimentStatusInDB mocks base method
func (m *MockManagerClient) UpdateExperimentStatusInDB(arg0 *v1alpha2.Experiment) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateExperimentStatusInDB", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateExperimentStatusInDB indicates an expected call of UpdateExperimentStatusInDB
func (mr *MockManagerClientMockRecorder) UpdateExperimentStatusInDB(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateExperimentStatusInDB", reflect.TypeOf((*MockManagerClient)(nil).UpdateExperimentStatusInDB), arg0)
}