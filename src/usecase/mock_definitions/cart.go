// Code generated by MockGen. DO NOT EDIT.
// Source: src/usecase/definitions/cart.go

// Package mock_definitions is a generated GoMock package.
package mock_definitions

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	models "github.com/itsmeadi/cart/src/entities/models"
	reflect "reflect"
)

// MockCart is a mock of Cart interface
type MockCart struct {
	ctrl     *gomock.Controller
	recorder *MockCartMockRecorder
}

// MockCartMockRecorder is the mock recorder for MockCart
type MockCartMockRecorder struct {
	mock *MockCart
}

// NewMockCart creates a new mock instance
func NewMockCart(ctrl *gomock.Controller) *MockCart {
	mock := &MockCart{ctrl: ctrl}
	mock.recorder = &MockCartMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCart) EXPECT() *MockCartMockRecorder {
	return m.recorder
}

// UpdateCart mocks base method
func (m *MockCart) UpdateCart(ctx context.Context, userId, productId, qty int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCart", ctx, userId, productId, qty)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateCart indicates an expected call of UpdateCart
func (mr *MockCartMockRecorder) UpdateCart(ctx, userId, productId, qty interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCart", reflect.TypeOf((*MockCart)(nil).UpdateCart), ctx, userId, productId, qty)
}

// AddToCart mocks base method
func (m *MockCart) AddToCart(ctx context.Context, userId, productId, qty int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddToCart", ctx, userId, productId, qty)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddToCart indicates an expected call of AddToCart
func (mr *MockCartMockRecorder) AddToCart(ctx, userId, productId, qty interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddToCart", reflect.TypeOf((*MockCart)(nil).AddToCart), ctx, userId, productId, qty)
}

// GetCart mocks base method
func (m *MockCart) GetCart(ctx context.Context, userId int64) (models.CartDetail, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCart", ctx, userId)
	ret0, _ := ret[0].(models.CartDetail)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCart indicates an expected call of GetCart
func (mr *MockCartMockRecorder) GetCart(ctx, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCart", reflect.TypeOf((*MockCart)(nil).GetCart), ctx, userId)
}

// RemoveProductFromCart mocks base method
func (m *MockCart) RemoveProductFromCart(ctx context.Context, productId, userId int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveProductFromCart", ctx, productId, userId)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveProductFromCart indicates an expected call of RemoveProductFromCart
func (mr *MockCartMockRecorder) RemoveProductFromCart(ctx, productId, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveProductFromCart", reflect.TypeOf((*MockCart)(nil).RemoveProductFromCart), ctx, productId, userId)
}
