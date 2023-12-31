// Code generated by MockGen. DO NOT EDIT.
// Source: interfaces.go

// Package get_user_segments is a generated GoMock package.
package get_user_segments

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockGetUserSegmentsService is a mock of GetUserSegmentsService interface.
type MockGetUserSegmentsService struct {
	ctrl     *gomock.Controller
	recorder *MockGetUserSegmentsServiceMockRecorder
}

// MockGetUserSegmentsServiceMockRecorder is the mock recorder for MockGetUserSegmentsService.
type MockGetUserSegmentsServiceMockRecorder struct {
	mock *MockGetUserSegmentsService
}

// NewMockGetUserSegmentsService creates a new mock instance.
func NewMockGetUserSegmentsService(ctrl *gomock.Controller) *MockGetUserSegmentsService {
	mock := &MockGetUserSegmentsService{ctrl: ctrl}
	mock.recorder = &MockGetUserSegmentsServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGetUserSegmentsService) EXPECT() *MockGetUserSegmentsServiceMockRecorder {
	return m.recorder
}

// GetUserSegments mocks base method.
func (m *MockGetUserSegmentsService) GetUserSegments(data *GetUserSegmentsData) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserSegments", data)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserSegments indicates an expected call of GetUserSegments.
func (mr *MockGetUserSegmentsServiceMockRecorder) GetUserSegments(data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserSegments", reflect.TypeOf((*MockGetUserSegmentsService)(nil).GetUserSegments), data)
}
