// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/platform/database/database.go

// Package database is a generated GoMock package.
package database

import (
	reflect "reflect"

	openapi3 "github.com/getkin/kin-openapi/openapi3"
	gomock "github.com/golang/mock/gomock"
)

// MockDBOpenAPILoader is a mock of DBOpenAPILoader interface.
type MockDBOpenAPILoader struct {
	ctrl     *gomock.Controller
	recorder *MockDBOpenAPILoaderMockRecorder
}

// MockDBOpenAPILoaderMockRecorder is the mock recorder for MockDBOpenAPILoader.
type MockDBOpenAPILoaderMockRecorder struct {
	mock *MockDBOpenAPILoader
}

// NewMockDBOpenAPILoader creates a new mock instance.
func NewMockDBOpenAPILoader(ctrl *gomock.Controller) *MockDBOpenAPILoader {
	mock := &MockDBOpenAPILoader{ctrl: ctrl}
	mock.recorder = &MockDBOpenAPILoaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDBOpenAPILoader) EXPECT() *MockDBOpenAPILoaderMockRecorder {
	return m.recorder
}

// AfterLoad mocks base method.
func (m *MockDBOpenAPILoader) AfterLoad(dbStoragePath string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AfterLoad", dbStoragePath)
	ret0, _ := ret[0].(error)
	return ret0
}

// AfterLoad indicates an expected call of AfterLoad.
func (mr *MockDBOpenAPILoaderMockRecorder) AfterLoad(dbStoragePath interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AfterLoad", reflect.TypeOf((*MockDBOpenAPILoader)(nil).AfterLoad), dbStoragePath)
}

// IsLoaded mocks base method.
func (m *MockDBOpenAPILoader) IsLoaded(schemaID int) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsLoaded", schemaID)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsLoaded indicates an expected call of IsLoaded.
func (mr *MockDBOpenAPILoaderMockRecorder) IsLoaded(schemaID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsLoaded", reflect.TypeOf((*MockDBOpenAPILoader)(nil).IsLoaded), schemaID)
}

// IsReady mocks base method.
func (m *MockDBOpenAPILoader) IsReady() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsReady")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsReady indicates an expected call of IsReady.
func (mr *MockDBOpenAPILoaderMockRecorder) IsReady() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsReady", reflect.TypeOf((*MockDBOpenAPILoader)(nil).IsReady))
}

// Load mocks base method.
func (m *MockDBOpenAPILoader) Load(dbStoragePath string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Load", dbStoragePath)
	ret0, _ := ret[0].(error)
	return ret0
}

// Load indicates an expected call of Load.
func (mr *MockDBOpenAPILoaderMockRecorder) Load(dbStoragePath interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Load", reflect.TypeOf((*MockDBOpenAPILoader)(nil).Load), dbStoragePath)
}

// SchemaIDs mocks base method.
func (m *MockDBOpenAPILoader) SchemaIDs() []int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SchemaIDs")
	ret0, _ := ret[0].([]int)
	return ret0
}

// SchemaIDs indicates an expected call of SchemaIDs.
func (mr *MockDBOpenAPILoaderMockRecorder) SchemaIDs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SchemaIDs", reflect.TypeOf((*MockDBOpenAPILoader)(nil).SchemaIDs))
}

// ShouldUpdate mocks base method.
func (m *MockDBOpenAPILoader) ShouldUpdate(newStorage DBOpenAPILoader) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ShouldUpdate", newStorage)
	ret0, _ := ret[0].(bool)
	return ret0
}

// ShouldUpdate indicates an expected call of ShouldUpdate.
func (mr *MockDBOpenAPILoaderMockRecorder) ShouldUpdate(newStorage interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ShouldUpdate", reflect.TypeOf((*MockDBOpenAPILoader)(nil).ShouldUpdate), newStorage)
}

// Specification mocks base method.
func (m *MockDBOpenAPILoader) Specification(schemaID int) *openapi3.T {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Specification", schemaID)
	ret0, _ := ret[0].(*openapi3.T)
	return ret0
}

// Specification indicates an expected call of Specification.
func (mr *MockDBOpenAPILoaderMockRecorder) Specification(schemaID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Specification", reflect.TypeOf((*MockDBOpenAPILoader)(nil).Specification), schemaID)
}

// SpecificationRaw mocks base method.
func (m *MockDBOpenAPILoader) SpecificationRaw(schemaID int) interface{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SpecificationRaw", schemaID)
	ret0, _ := ret[0].(interface{})
	return ret0
}

// SpecificationRaw indicates an expected call of SpecificationRaw.
func (mr *MockDBOpenAPILoaderMockRecorder) SpecificationRaw(schemaID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SpecificationRaw", reflect.TypeOf((*MockDBOpenAPILoader)(nil).SpecificationRaw), schemaID)
}

// SpecificationRawContent mocks base method.
func (m *MockDBOpenAPILoader) SpecificationRawContent(schemaID int) []byte {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SpecificationRawContent", schemaID)
	ret0, _ := ret[0].([]byte)
	return ret0
}

// SpecificationRawContent indicates an expected call of SpecificationRawContent.
func (mr *MockDBOpenAPILoaderMockRecorder) SpecificationRawContent(schemaID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SpecificationRawContent", reflect.TypeOf((*MockDBOpenAPILoader)(nil).SpecificationRawContent), schemaID)
}

// SpecificationVersion mocks base method.
func (m *MockDBOpenAPILoader) SpecificationVersion(schemaID int) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SpecificationVersion", schemaID)
	ret0, _ := ret[0].(string)
	return ret0
}

// SpecificationVersion indicates an expected call of SpecificationVersion.
func (mr *MockDBOpenAPILoaderMockRecorder) SpecificationVersion(schemaID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SpecificationVersion", reflect.TypeOf((*MockDBOpenAPILoader)(nil).SpecificationVersion), schemaID)
}

// Version mocks base method.
func (m *MockDBOpenAPILoader) Version() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Version")
	ret0, _ := ret[0].(int)
	return ret0
}

// Version indicates an expected call of Version.
func (mr *MockDBOpenAPILoaderMockRecorder) Version() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Version", reflect.TypeOf((*MockDBOpenAPILoader)(nil).Version))
}
