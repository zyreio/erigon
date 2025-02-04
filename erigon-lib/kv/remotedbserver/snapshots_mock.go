// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ledgerwatch/erigon-lib/kv/remotedbserver (interfaces: Snapshots)
//
// Generated by this command:
//
//	mockgen -destination=./snapshots_mock.go -package=remotedbserver . Snapshots
//

// Package remotedbserver is a generated GoMock package.
package remotedbserver

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockSnapshots is a mock of Snapshots interface.
type MockSnapshots struct {
	ctrl     *gomock.Controller
	recorder *MockSnapshotsMockRecorder
}

// MockSnapshotsMockRecorder is the mock recorder for MockSnapshots.
type MockSnapshotsMockRecorder struct {
	mock *MockSnapshots
}

// NewMockSnapshots creates a new mock instance.
func NewMockSnapshots(ctrl *gomock.Controller) *MockSnapshots {
	mock := &MockSnapshots{ctrl: ctrl}
	mock.recorder = &MockSnapshotsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSnapshots) EXPECT() *MockSnapshotsMockRecorder {
	return m.recorder
}

// Files mocks base method.
func (m *MockSnapshots) Files() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Files")
	ret0, _ := ret[0].([]string)
	return ret0
}

// Files indicates an expected call of Files.
func (mr *MockSnapshotsMockRecorder) Files() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Files", reflect.TypeOf((*MockSnapshots)(nil).Files))
}
