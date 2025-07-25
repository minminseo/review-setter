// Code generated by MockGen. DO NOT EDIT.
// Source: domain/user/user.go
//
// Generated by this command:
//
//	mockgen -source=domain/user/user.go -destination=domain/user/mock_user.go -package user
//

// Package user is a generated GoMock package.
package user

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockIHasher is a mock of IHasher interface.
type MockIHasher struct {
	ctrl     *gomock.Controller
	recorder *MockIHasherMockRecorder
	isgomock struct{}
}

// MockIHasherMockRecorder is the mock recorder for MockIHasher.
type MockIHasherMockRecorder struct {
	mock *MockIHasher
}

// NewMockIHasher creates a new mock instance.
func NewMockIHasher(ctrl *gomock.Controller) *MockIHasher {
	mock := &MockIHasher{ctrl: ctrl}
	mock.recorder = &MockIHasherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIHasher) EXPECT() *MockIHasherMockRecorder {
	return m.recorder
}

// GenerateSearchKey mocks base method.
func (m *MockIHasher) GenerateSearchKey(email string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateSearchKey", email)
	ret0, _ := ret[0].(string)
	return ret0
}

// GenerateSearchKey indicates an expected call of GenerateSearchKey.
func (mr *MockIHasherMockRecorder) GenerateSearchKey(email any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateSearchKey", reflect.TypeOf((*MockIHasher)(nil).GenerateSearchKey), email)
}
