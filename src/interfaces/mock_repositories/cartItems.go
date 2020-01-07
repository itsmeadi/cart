// Code generated by MockGen. DO NOT EDIT.
// Source: src/domain/repositories/cartItems.go

// Package mock_repositories is a generated GoMock package.
package mock_repositories

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	models "github.com/itsmeadi/cart/src/entities/models"
	reflect "reflect"
)

// MockCartItems is a mock of CartItems interface
type MockCartItems struct {
	ctrl     *gomock.Controller
	recorder *MockCartItemsMockRecorder
}

// MockCartItemsMockRecorder is the mock recorder for MockCartItems
type MockCartItemsMockRecorder struct {
	mock *MockCartItems
}

// NewMockCartItems creates a new mock instance
func NewMockCartItems(ctrl *gomock.Controller) *MockCartItems {
	mock := &MockCartItems{ctrl: ctrl}
	mock.recorder = &MockCartItemsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCartItems) EXPECT() *MockCartItemsMockRecorder {
	return m.recorder
}

// InsertItemInCart mocks base method
func (m *MockCartItems) InsertItemInCart(ctx context.Context, item models.CartItems) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertItemInCart", ctx, item)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertItemInCart indicates an expected call of InsertItemInCart
func (mr *MockCartItemsMockRecorder) InsertItemInCart(ctx, item interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertItemInCart", reflect.TypeOf((*MockCartItems)(nil).InsertItemInCart), ctx, item)
}

// GetItemsInCart mocks base method
func (m *MockCartItems) GetItemsInCart(ctx context.Context, cartId, status, limit int64) ([]models.CartItems, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetItemsInCart", ctx, cartId, status, limit)
	ret0, _ := ret[0].([]models.CartItems)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetItemsInCart indicates an expected call of GetItemsInCart
func (mr *MockCartItemsMockRecorder) GetItemsInCart(ctx, cartId, status, limit interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetItemsInCart", reflect.TypeOf((*MockCartItems)(nil).GetItemsInCart), ctx, cartId, status, limit)
}

// GetProductFromCart mocks base method
func (m *MockCartItems) GetProductFromCart(ctx context.Context, cartId, productId int64, status int) ([]models.CartItems, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProductFromCart", ctx, cartId, productId, status)
	ret0, _ := ret[0].([]models.CartItems)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProductFromCart indicates an expected call of GetProductFromCart
func (mr *MockCartItemsMockRecorder) GetProductFromCart(ctx, cartId, productId, status interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProductFromCart", reflect.TypeOf((*MockCartItems)(nil).GetProductFromCart), ctx, cartId, productId, status)
}

// UpdateCartQty mocks base method
func (m *MockCartItems) UpdateCartQty(ctx context.Context, qty, cartId, productId int64) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCartQty", ctx, qty, cartId, productId)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateCartQty indicates an expected call of UpdateCartQty
func (mr *MockCartItemsMockRecorder) UpdateCartQty(ctx, qty, cartId, productId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCartQty", reflect.TypeOf((*MockCartItems)(nil).UpdateCartQty), ctx, qty, cartId, productId)
}

// UpdateCartItemStatus mocks base method
func (m *MockCartItems) UpdateCartItemStatus(ctx context.Context, status, cartId, productId int64) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCartItemStatus", ctx, status, cartId, productId)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateCartItemStatus indicates an expected call of UpdateCartItemStatus
func (mr *MockCartItemsMockRecorder) UpdateCartItemStatus(ctx, status, cartId, productId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCartItemStatus", reflect.TypeOf((*MockCartItems)(nil).UpdateCartItemStatus), ctx, status, cartId, productId)
}
