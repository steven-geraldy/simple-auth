// Code generated by MockGen. DO NOT EDIT.
// Source: repository/interfaces.go

// Package repository is a generated GoMock package.
package repository

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockUserInterface is a mock of UserInterface interface.
type MockUserInterface struct {
	ctrl     *gomock.Controller
	recorder *MockUserInterfaceMockRecorder
}

// MockUserInterfaceMockRecorder is the mock recorder for MockUserInterface.
type MockUserInterfaceMockRecorder struct {
	mock *MockUserInterface
}

// NewMockUserInterface creates a new mock instance.
func NewMockUserInterface(ctrl *gomock.Controller) *MockUserInterface {
	mock := &MockUserInterface{ctrl: ctrl}
	mock.recorder = &MockUserInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserInterface) EXPECT() *MockUserInterfaceMockRecorder {
	return m.recorder
}

// GetUser mocks base method.
func (m *MockUserInterface) GetUser(ctx context.Context, input GetUserInput) (GetUserOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", ctx, input)
	ret0, _ := ret[0].(GetUserOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockUserInterfaceMockRecorder) GetUser(ctx, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockUserInterface)(nil).GetUser), ctx, input)
}

// Login mocks base method.
func (m *MockUserInterface) Login(ctx context.Context, input LoginInput) (LoginOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Login", ctx, input)
	ret0, _ := ret[0].(LoginOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Login indicates an expected call of Login.
func (mr *MockUserInterfaceMockRecorder) Login(ctx, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockUserInterface)(nil).Login), ctx, input)
}

// Register mocks base method.
func (m *MockUserInterface) Register(ctx context.Context, input RegisterUserInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register", ctx, input)
	ret0, _ := ret[0].(error)
	return ret0
}

// Register indicates an expected call of Register.
func (mr *MockUserInterfaceMockRecorder) Register(ctx, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockUserInterface)(nil).Register), ctx, input)
}

// UpdateUser mocks base method.
func (m *MockUserInterface) UpdateUser(ctx context.Context, input UpdateUserInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", ctx, input)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockUserInterfaceMockRecorder) UpdateUser(ctx, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockUserInterface)(nil).UpdateUser), ctx, input)
}

// MockAuthInterface is a mock of AuthInterface interface.
type MockAuthInterface struct {
	ctrl     *gomock.Controller
	recorder *MockAuthInterfaceMockRecorder
}

// MockAuthInterfaceMockRecorder is the mock recorder for MockAuthInterface.
type MockAuthInterfaceMockRecorder struct {
	mock *MockAuthInterface
}

// NewMockAuthInterface creates a new mock instance.
func NewMockAuthInterface(ctrl *gomock.Controller) *MockAuthInterface {
	mock := &MockAuthInterface{ctrl: ctrl}
	mock.recorder = &MockAuthInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthInterface) EXPECT() *MockAuthInterfaceMockRecorder {
	return m.recorder
}

// GenerateToken mocks base method.
func (m *MockAuthInterface) GenerateToken(input GenerateTokenInput) (GenerateTokenOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateToken", input)
	ret0, _ := ret[0].(GenerateTokenOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateToken indicates an expected call of GenerateToken.
func (mr *MockAuthInterfaceMockRecorder) GenerateToken(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateToken", reflect.TypeOf((*MockAuthInterface)(nil).GenerateToken), input)
}

// ParseToken mocks base method.
func (m *MockAuthInterface) ParseToken(input ParseTokenInput) (ParseTokenOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ParseToken", input)
	ret0, _ := ret[0].(ParseTokenOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ParseToken indicates an expected call of ParseToken.
func (mr *MockAuthInterfaceMockRecorder) ParseToken(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParseToken", reflect.TypeOf((*MockAuthInterface)(nil).ParseToken), input)
}
