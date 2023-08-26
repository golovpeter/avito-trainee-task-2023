// Code generated by MockGen. DO NOT EDIT.
// Source: interfaces.go

// Package user_segments is a generated GoMock package.
package user_segments

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// ChangeUserSegments mocks base method.
func (m *MockRepository) ChangeUserSegments(changeData ChangeUserSegmentsData) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangeUserSegments", changeData)
	ret0, _ := ret[0].(error)
	return ret0
}

// ChangeUserSegments indicates an expected call of ChangeUserSegments.
func (mr *MockRepositoryMockRecorder) ChangeUserSegments(changeData interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeUserSegments", reflect.TypeOf((*MockRepository)(nil).ChangeUserSegments), changeData)
}
