// Code generated by MockGen. DO NOT EDIT.
// Source: fs.go

// Package mock_git is a generated GoMock package.
package mock_git

import (
	os "os"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// Mockfs is a mock of fs interface.
type Mockfs struct {
	ctrl     *gomock.Controller
	recorder *MockfsMockRecorder
}

// MockfsMockRecorder is the mock recorder for Mockfs.
type MockfsMockRecorder struct {
	mock *Mockfs
}

// NewMockfs creates a new mock instance.
func NewMockfs(ctrl *gomock.Controller) *Mockfs {
	mock := &Mockfs{ctrl: ctrl}
	mock.recorder = &MockfsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *Mockfs) EXPECT() *MockfsMockRecorder {
	return m.recorder
}

// MkdirAll mocks base method.
func (m *Mockfs) MkdirAll(path string, perm os.FileMode) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MkdirAll", path, perm)
	ret0, _ := ret[0].(error)
	return ret0
}

// MkdirAll indicates an expected call of MkdirAll.
func (mr *MockfsMockRecorder) MkdirAll(path, perm interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MkdirAll", reflect.TypeOf((*Mockfs)(nil).MkdirAll), path, perm)
}

// MkdirTemp mocks base method.
func (m *Mockfs) MkdirTemp(dir, prefix string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MkdirTemp", dir, prefix)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MkdirTemp indicates an expected call of MkdirTemp.
func (mr *MockfsMockRecorder) MkdirTemp(dir, prefix interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MkdirTemp", reflect.TypeOf((*Mockfs)(nil).MkdirTemp), dir, prefix)
}

// RemoveAll mocks base method.
func (m *Mockfs) RemoveAll(path string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveAll", path)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveAll indicates an expected call of RemoveAll.
func (mr *MockfsMockRecorder) RemoveAll(path interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveAll", reflect.TypeOf((*Mockfs)(nil).RemoveAll), path)
}

// Stat mocks base method.
func (m *Mockfs) Stat(name string) (os.FileInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stat", name)
	ret0, _ := ret[0].(os.FileInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Stat indicates an expected call of Stat.
func (mr *MockfsMockRecorder) Stat(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stat", reflect.TypeOf((*Mockfs)(nil).Stat), name)
}
