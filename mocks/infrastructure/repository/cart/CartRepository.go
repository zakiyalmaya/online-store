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
func (m *MockRepository) Create(cart *model.CartEntity) (*model.CartEntity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", cart)
	ret0, _ := ret[0].(*model.CartEntity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockRepositoryMockRecorder) Create(cart interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockRepository)(nil).Create), cart)
}

// Delete mocks base method.
func (m *MockRepository) Delete(request *model.DeleteCartRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", request)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockRepositoryMockRecorder) Delete(request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockRepository)(nil).Delete), request)
}

// GetByID mocks base method.
func (m *MockRepository) GetByID(cartID int) (*model.CartEntity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", cartID)
	ret0, _ := ret[0].(*model.CartEntity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockRepositoryMockRecorder) GetByID(cartID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockRepository)(nil).GetByID), cartID)
}

// GetByParams mocks base method.
func (m *MockRepository) GetByParams(request *model.GetCartRequest) ([]*model.CartEntity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByParams", request)
	ret0, _ := ret[0].([]*model.CartEntity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByParams indicates an expected call of GetByParams.
func (mr *MockRepositoryMockRecorder) GetByParams(request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByParams", reflect.TypeOf((*MockRepository)(nil).GetByParams), request)
}

// GetItemByID mocks base method.
func (m *MockRepository) GetItemByID(cartItemID int) (*model.CartItemEntity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetItemByID", cartItemID)
	ret0, _ := ret[0].(*model.CartItemEntity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetItemByID indicates an expected call of GetItemByID.
func (mr *MockRepositoryMockRecorder) GetItemByID(cartItemID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetItemByID", reflect.TypeOf((*MockRepository)(nil).GetItemByID), cartItemID)
}

// Upsert mocks base method.
func (m *MockRepository) Upsert(cartID int, items []*model.CartItemEntity) (*model.CartEntity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upsert", cartID, items)
	ret0, _ := ret[0].(*model.CartEntity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Upsert indicates an expected call of Upsert.
func (mr *MockRepositoryMockRecorder) Upsert(cartID, items interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upsert", reflect.TypeOf((*MockRepository)(nil).Upsert), cartID, items)
}