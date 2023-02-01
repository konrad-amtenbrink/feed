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

// GetDocuments mocks base method.
func (m *MockDatabase) GetDocuments(ctx context.Context) ([]Document, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDocuments", ctx)
	ret0, _ := ret[0].([]Document)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDocuments indicates an expected call of GetDocuments.
func (mr *MockDatabaseMockRecorder) GetDocuments(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDocuments", reflect.TypeOf((*MockDatabase)(nil).GetDocuments), ctx)
}
