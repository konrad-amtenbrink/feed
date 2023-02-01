// Code generated by MockGen. DO NOT EDIT.
// Source: db/go

// Package mock_db is a generated GoMock package.
package db

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
)

// MockDatabase is a mock of Database interface.
type MockDatabase struct {
	ctrl     *gomock.Controller
	recorder *MockDatabaseMockRecorder
}

// MockDatabaseMockRecorder is the mock recorder for MockDatabase.
type MockDatabaseMockRecorder struct {
	mock *MockDatabase
}

// NewMockDatabase creates a new mock instance.
func NewMockDatabase(ctrl *gomock.Controller) *MockDatabase {
	mock := &MockDatabase{ctrl: ctrl}
	mock.recorder = &MockDatabaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDatabase) EXPECT() *MockDatabaseMockRecorder {
	return m.recorder
}

// CreateDocument mocks base method.
func (m *MockDatabase) CreateDocument(ctx context.Context, doc Document) (uuid.UUID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateDocument", ctx, doc)
	ret0, _ := ret[0].(uuid.UUID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateDocument indicates an expected call of CreateDocument.
func (mr *MockDatabaseMockRecorder) CreateDocument(ctx, doc interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateDocument", reflect.TypeOf((*MockDatabase)(nil).CreateDocument), ctx, doc)
}

// CreateUser mocks base method.
func (m *MockDatabase) CreateUser(ctx context.Context, user User) (uuid.UUID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", ctx, user)
	ret0, _ := ret[0].(uuid.UUID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockDatabaseMockRecorder) CreateUser(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockDatabase)(nil).CreateUser), ctx, user)
}

// DeleteDocumentById mocks base method.
func (m *MockDatabase) DeleteDocumentById(ctx context.Context, id uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteDocumentById", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteDocumentById indicates an expected call of DeleteDocumentById.
func (mr *MockDatabaseMockRecorder) DeleteDocumentById(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteDocumentById", reflect.TypeOf((*MockDatabase)(nil).DeleteDocumentById), ctx, id)
}

// DeleteUserById mocks base method.
func (m *MockDatabase) DeleteUserById(ctx context.Context, id uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUserById", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUserById indicates an expected call of DeleteUserById.
func (mr *MockDatabaseMockRecorder) DeleteUserById(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUserById", reflect.TypeOf((*MockDatabase)(nil).DeleteUserById), ctx, id)
}

// GetDocumentById mocks base method.
func (m *MockDatabase) GetDocumentById(ctx context.Context, id uuid.UUID) (Document, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDocumentById", ctx, id)
	ret0, _ := ret[0].(Document)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDocumentById indicates an expected call of GetDocumentById.
func (mr *MockDatabaseMockRecorder) GetDocumentById(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDocumentById", reflect.TypeOf((*MockDatabase)(nil).GetDocumentById), ctx, id)
}

// GetDocumentsByUserId mocks base method.
func (m *MockDatabase) GetDocumentsByUserId(ctx context.Context, userId uuid.UUID) ([]Document, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDocumentsByUserId", ctx, userId)
	ret0, _ := ret[0].([]Document)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDocumentsByUserId indicates an expected call of GetDocumentsByUserId.
func (mr *MockDatabaseMockRecorder) GetDocumentsByUserId(ctx, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDocumentsByUserId", reflect.TypeOf((*MockDatabase)(nil).GetDocumentsByUserId), ctx, userId)
}

// GetUserById mocks base method.
func (m *MockDatabase) GetUserById(ctx context.Context, id uuid.UUID) (User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserById", ctx, id)
	ret0, _ := ret[0].(User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserById indicates an expected call of GetUserById.
func (mr *MockDatabaseMockRecorder) GetUserById(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserById", reflect.TypeOf((*MockDatabase)(nil).GetUserById), ctx, id)
}

// GetUserByUsername mocks base method.
func (m *MockDatabase) GetUserByUsername(ctx context.Context, username string) (User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByUsername", ctx, username)
	ret0, _ := ret[0].(User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByUsername indicates an expected call of GetUserByUsername.
func (mr *MockDatabaseMockRecorder) GetUserByUsername(ctx, username interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByUsername", reflect.TypeOf((*MockDatabase)(nil).GetUserByUsername), ctx, username)
}
