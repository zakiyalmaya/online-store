// Code generated by MockGen. DO NOT EDIT.
// Source: repo.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	model "github.com/zakiyalmaya/online-store/model"
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

// Create mocks base method.
func (m *MockRepository) Create(customer *model.CustomerEntity) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", customer)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockRepositoryMockRecorder) Create(customer interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockRepository)(nil).Create), customer)
}

// GetByUsername mocks base method.
func (m *MockRepository) GetByUsername(username string) (*model.CustomerEntity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByUsername", username)
	ret0, _ := ret[0].(*model.CustomerEntity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByUsername indicates an expected call of GetByUsername.
func (mr *MockRepositoryMockRecorder) GetByUsername(username interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByUsername", reflect.TypeOf((*MockRepository)(nil).GetByUsername), username)
}