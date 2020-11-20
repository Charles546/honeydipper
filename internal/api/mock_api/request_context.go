// Copyright 2019 Honey Science Corporation
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, you can obtain one at http://mozilla.org/MPL/2.0/.
//

// Code generated by MockGen. DO NOT EDIT.
// Source: internal/api/request_context.go

// Package mock_api is a generated GoMock package.
package mock_api

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockRequestContext is a mock of RequestContext interface.
type MockRequestContext struct {
	ctrl     *gomock.Controller
	recorder *MockRequestContextMockRecorder
}

// MockRequestContextMockRecorder is the mock recorder for MockRequestContext.
type MockRequestContextMockRecorder struct {
	mock *MockRequestContext
}

// NewMockRequestContext creates a new mock instance.
func NewMockRequestContext(ctrl *gomock.Controller) *MockRequestContext {
	mock := &MockRequestContext{ctrl: ctrl}
	mock.recorder = &MockRequestContextMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRequestContext) EXPECT() *MockRequestContextMockRecorder {
	return m.recorder
}

// AbortWithStatusJSON mocks base method.
func (m *MockRequestContext) AbortWithStatusJSON(arg0 int, arg1 interface{}) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AbortWithStatusJSON", arg0, arg1)
}

// AbortWithStatusJSON indicates an expected call of AbortWithStatusJSON.
func (mr *MockRequestContextMockRecorder) AbortWithStatusJSON(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AbortWithStatusJSON", reflect.TypeOf((*MockRequestContext)(nil).AbortWithStatusJSON), arg0, arg1)
}

// IndentedJSON mocks base method.
func (m *MockRequestContext) IndentedJSON(arg0 int, arg1 interface{}) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "IndentedJSON", arg0, arg1)
}

// IndentedJSON indicates an expected call of IndentedJSON.
func (mr *MockRequestContextMockRecorder) IndentedJSON(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IndentedJSON", reflect.TypeOf((*MockRequestContext)(nil).IndentedJSON), arg0, arg1)
}

// ContentType mocks base method.
func (m *MockRequestContext) ContentType() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ContentType")
	ret0, _ := ret[0].(string)
	return ret0
}

// ContentType indicates an expected call of ContentType.
func (mr *MockRequestContextMockRecorder) ContentType() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ContentType", reflect.TypeOf((*MockRequestContext)(nil).ContentType))
}

// Get mocks base method.
func (m *MockRequestContext) Get(arg0 string) (interface{}, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockRequestContextMockRecorder) Get(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockRequestContext)(nil).Get), arg0)
}

// Set mocks base method.
func (m *MockRequestContext) Set(arg0 string, arg1 interface{}) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Set", arg0, arg1)
}

// Set indicates an expected call of Set.
func (mr *MockRequestContextMockRecorder) Set(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockRequestContext)(nil).Set), arg0, arg1)
}

// GetPath mocks base method.
func (m *MockRequestContext) GetPath() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPath")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetPath indicates an expected call of GetPath.
func (mr *MockRequestContextMockRecorder) GetPath() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPath", reflect.TypeOf((*MockRequestContext)(nil).GetPath))
}

// GetPayload mocks base method.
func (m *MockRequestContext) GetPayload(method string) map[string]interface{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPayload", method)
	ret0, _ := ret[0].(map[string]interface{})
	return ret0
}

// GetPayload indicates an expected call of GetPayload.
func (mr *MockRequestContextMockRecorder) GetPayload(method interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPayload", reflect.TypeOf((*MockRequestContext)(nil).GetPayload), method)
}