// Code generated by MockGen. DO NOT EDIT.
// Source: repository/interfaces.go

// Package repository is a generated GoMock package.
package repository

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
)

// MockRepositoryInterface is a mock of RepositoryInterface interface.
type MockRepositoryInterface struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryInterfaceMockRecorder
}

// MockRepositoryInterfaceMockRecorder is the mock recorder for MockRepositoryInterface.
type MockRepositoryInterfaceMockRecorder struct {
	mock *MockRepositoryInterface
}

// NewMockRepositoryInterface creates a new mock instance.
func NewMockRepositoryInterface(ctrl *gomock.Controller) *MockRepositoryInterface {
	mock := &MockRepositoryInterface{ctrl: ctrl}
	mock.recorder = &MockRepositoryInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepositoryInterface) EXPECT() *MockRepositoryInterfaceMockRecorder {
	return m.recorder
}

// FetchEstateTrees mocks base method.
func (m *MockRepositoryInterface) FetchEstateTrees(ctx context.Context, estateId uuid.UUID) ([]EstateTree, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchEstateTrees", ctx, estateId)
	ret0, _ := ret[0].([]EstateTree)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchEstateTrees indicates an expected call of FetchEstateTrees.
func (mr *MockRepositoryInterfaceMockRecorder) FetchEstateTrees(ctx, estateId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchEstateTrees", reflect.TypeOf((*MockRepositoryInterface)(nil).FetchEstateTrees), ctx, estateId)
}

// GetEstateByID mocks base method.
func (m *MockRepositoryInterface) GetEstateByID(ctx context.Context, input GetEstateByIDInput) (Estate, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEstateByID", ctx, input)
	ret0, _ := ret[0].(Estate)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEstateByID indicates an expected call of GetEstateByID.
func (mr *MockRepositoryInterfaceMockRecorder) GetEstateByID(ctx, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEstateByID", reflect.TypeOf((*MockRepositoryInterface)(nil).GetEstateByID), ctx, input)
}

// StoreEstate mocks base method.
func (m *MockRepositoryInterface) StoreEstate(ctx context.Context, input StoreEstateInput) (StoreEstateOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StoreEstate", ctx, input)
	ret0, _ := ret[0].(StoreEstateOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StoreEstate indicates an expected call of StoreEstate.
func (mr *MockRepositoryInterfaceMockRecorder) StoreEstate(ctx, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StoreEstate", reflect.TypeOf((*MockRepositoryInterface)(nil).StoreEstate), ctx, input)
}

// StoreEstateIdTree mocks base method.
func (m *MockRepositoryInterface) StoreEstateIdTree(ctx context.Context, input StoreEstateIdTreeInput) (StoreEstateIdTreeOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StoreEstateIdTree", ctx, input)
	ret0, _ := ret[0].(StoreEstateIdTreeOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StoreEstateIdTree indicates an expected call of StoreEstateIdTree.
func (mr *MockRepositoryInterfaceMockRecorder) StoreEstateIdTree(ctx, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StoreEstateIdTree", reflect.TypeOf((*MockRepositoryInterface)(nil).StoreEstateIdTree), ctx, input)
}
