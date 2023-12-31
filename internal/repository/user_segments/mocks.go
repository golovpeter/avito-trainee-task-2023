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

// AddOneUserSegment mocks base method.
func (m *MockRepository) AddOneUserSegment(userId, segmentId int64, addedSegment bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddOneUserSegment", userId, segmentId, addedSegment)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddOneUserSegment indicates an expected call of AddOneUserSegment.
func (mr *MockRepositoryMockRecorder) AddOneUserSegment(userId, segmentId, addedSegment interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddOneUserSegment", reflect.TypeOf((*MockRepository)(nil).AddOneUserSegment), userId, segmentId, addedSegment)
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

// DeleteExpiredUserSegments mocks base method.
func (m *MockRepository) DeleteExpiredUserSegments() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteExpiredUserSegments")
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteExpiredUserSegments indicates an expected call of DeleteExpiredUserSegments.
func (mr *MockRepositoryMockRecorder) DeleteExpiredUserSegments() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteExpiredUserSegments", reflect.TypeOf((*MockRepository)(nil).DeleteExpiredUserSegments))
}

// GetUserSegments mocks base method.
func (m *MockRepository) GetUserSegments(id int64) (map[string]SegmentInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserSegments", id)
	ret0, _ := ret[0].(map[string]SegmentInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserSegments indicates an expected call of GetUserSegments.
func (mr *MockRepositoryMockRecorder) GetUserSegments(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserSegments", reflect.TypeOf((*MockRepository)(nil).GetUserSegments), id)
}
