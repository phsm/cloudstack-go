//
// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
//

// Code generated by MockGen. DO NOT EDIT.
// Source: ./cloudstack/StoragePoolService.go
//
// Generated by this command:
//
//	mockgen -destination=./cloudstack/StoragePoolService_mock.go -package=cloudstack -copyright_file=header.txt -source=./cloudstack/StoragePoolService.go
//

// Package cloudstack is a generated GoMock package.
package cloudstack

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockStoragePoolServiceIface is a mock of StoragePoolServiceIface interface.
type MockStoragePoolServiceIface struct {
	ctrl     *gomock.Controller
	recorder *MockStoragePoolServiceIfaceMockRecorder
	isgomock struct{}
}

// MockStoragePoolServiceIfaceMockRecorder is the mock recorder for MockStoragePoolServiceIface.
type MockStoragePoolServiceIfaceMockRecorder struct {
	mock *MockStoragePoolServiceIface
}

// NewMockStoragePoolServiceIface creates a new mock instance.
func NewMockStoragePoolServiceIface(ctrl *gomock.Controller) *MockStoragePoolServiceIface {
	mock := &MockStoragePoolServiceIface{ctrl: ctrl}
	mock.recorder = &MockStoragePoolServiceIfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStoragePoolServiceIface) EXPECT() *MockStoragePoolServiceIfaceMockRecorder {
	return m.recorder
}

// AddObjectStoragePool mocks base method.
func (m *MockStoragePoolServiceIface) AddObjectStoragePool(p *AddObjectStoragePoolParams) (*AddObjectStoragePoolResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddObjectStoragePool", p)
	ret0, _ := ret[0].(*AddObjectStoragePoolResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddObjectStoragePool indicates an expected call of AddObjectStoragePool.
func (mr *MockStoragePoolServiceIfaceMockRecorder) AddObjectStoragePool(p any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddObjectStoragePool", reflect.TypeOf((*MockStoragePoolServiceIface)(nil).AddObjectStoragePool), p)
}

// CancelStorageMaintenance mocks base method.
func (m *MockStoragePoolServiceIface) CancelStorageMaintenance(p *CancelStorageMaintenanceParams) (*CancelStorageMaintenanceResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CancelStorageMaintenance", p)
	ret0, _ := ret[0].(*CancelStorageMaintenanceResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CancelStorageMaintenance indicates an expected call of CancelStorageMaintenance.
func (mr *MockStoragePoolServiceIfaceMockRecorder) CancelStorageMaintenance(p any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CancelStorageMaintenance", reflect.TypeOf((*MockStoragePoolServiceIface)(nil).CancelStorageMaintenance), p)
}

// ChangeStoragePoolScope mocks base method.
func (m *MockStoragePoolServiceIface) ChangeStoragePoolScope(p *ChangeStoragePoolScopeParams) (*ChangeStoragePoolScopeResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangeStoragePoolScope", p)
	ret0, _ := ret[0].(*ChangeStoragePoolScopeResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ChangeStoragePoolScope indicates an expected call of ChangeStoragePoolScope.
func (mr *MockStoragePoolServiceIfaceMockRecorder) ChangeStoragePoolScope(p any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeStoragePoolScope", reflect.TypeOf((*MockStoragePoolServiceIface)(nil).ChangeStoragePoolScope), p)
}

// DeleteObjectStoragePool mocks base method.
func (m *MockStoragePoolServiceIface) DeleteObjectStoragePool(p *DeleteObjectStoragePoolParams) (*DeleteObjectStoragePoolResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteObjectStoragePool", p)
	ret0, _ := ret[0].(*DeleteObjectStoragePoolResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteObjectStoragePool indicates an expected call of DeleteObjectStoragePool.
func (mr *MockStoragePoolServiceIfaceMockRecorder) DeleteObjectStoragePool(p any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteObjectStoragePool", reflect.TypeOf((*MockStoragePoolServiceIface)(nil).DeleteObjectStoragePool), p)
}

// EnableStorageMaintenance mocks base method.
func (m *MockStoragePoolServiceIface) EnableStorageMaintenance(p *EnableStorageMaintenanceParams) (*EnableStorageMaintenanceResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnableStorageMaintenance", p)
	ret0, _ := ret[0].(*EnableStorageMaintenanceResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EnableStorageMaintenance indicates an expected call of EnableStorageMaintenance.
func (mr *MockStoragePoolServiceIfaceMockRecorder) EnableStorageMaintenance(p any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnableStorageMaintenance", reflect.TypeOf((*MockStoragePoolServiceIface)(nil).EnableStorageMaintenance), p)
}

// GetAffectedVmsForStorageScopeChangeID mocks base method.
func (m *MockStoragePoolServiceIface) GetAffectedVmsForStorageScopeChangeID(keyword, clusterid, storageid string, opts ...OptionFunc) (string, int, error) {
	m.ctrl.T.Helper()
	varargs := []any{keyword, clusterid, storageid}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetAffectedVmsForStorageScopeChangeID", varargs...)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetAffectedVmsForStorageScopeChangeID indicates an expected call of GetAffectedVmsForStorageScopeChangeID.
func (mr *MockStoragePoolServiceIfaceMockRecorder) GetAffectedVmsForStorageScopeChangeID(keyword, clusterid, storageid any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{keyword, clusterid, storageid}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAffectedVmsForStorageScopeChangeID", reflect.TypeOf((*MockStoragePoolServiceIface)(nil).GetAffectedVmsForStorageScopeChangeID), varargs...)
}

// GetObjectStoragePoolByID mocks base method.
func (m *MockStoragePoolServiceIface) GetObjectStoragePoolByID(id string, opts ...OptionFunc) (*ObjectStoragePool, int, error) {
	m.ctrl.T.Helper()
	varargs := []any{id}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetObjectStoragePoolByID", varargs...)
	ret0, _ := ret[0].(*ObjectStoragePool)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetObjectStoragePoolByID indicates an expected call of GetObjectStoragePoolByID.
func (mr *MockStoragePoolServiceIfaceMockRecorder) GetObjectStoragePoolByID(id any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{id}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetObjectStoragePoolByID", reflect.TypeOf((*MockStoragePoolServiceIface)(nil).GetObjectStoragePoolByID), varargs...)
}

// GetObjectStoragePoolByName mocks base method.
func (m *MockStoragePoolServiceIface) GetObjectStoragePoolByName(name string, opts ...OptionFunc) (*ObjectStoragePool, int, error) {
	m.ctrl.T.Helper()
	varargs := []any{name}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetObjectStoragePoolByName", varargs...)
	ret0, _ := ret[0].(*ObjectStoragePool)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetObjectStoragePoolByName indicates an expected call of GetObjectStoragePoolByName.
func (mr *MockStoragePoolServiceIfaceMockRecorder) GetObjectStoragePoolByName(name any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{name}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetObjectStoragePoolByName", reflect.TypeOf((*MockStoragePoolServiceIface)(nil).GetObjectStoragePoolByName), varargs...)
}

// GetObjectStoragePoolID mocks base method.
func (m *MockStoragePoolServiceIface) GetObjectStoragePoolID(name string, opts ...OptionFunc) (string, int, error) {
	m.ctrl.T.Helper()
	varargs := []any{name}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetObjectStoragePoolID", varargs...)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetObjectStoragePoolID indicates an expected call of GetObjectStoragePoolID.
func (mr *MockStoragePoolServiceIfaceMockRecorder) GetObjectStoragePoolID(name any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{name}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetObjectStoragePoolID", reflect.TypeOf((*MockStoragePoolServiceIface)(nil).GetObjectStoragePoolID), varargs...)
}

// GetStoragePoolObjectByID mocks base method.
func (m *MockStoragePoolServiceIface) GetStoragePoolObjectByID(id string, opts ...OptionFunc) (*StoragePoolObject, int, error) {
	m.ctrl.T.Helper()
	varargs := []any{id}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetStoragePoolObjectByID", varargs...)
	ret0, _ := ret[0].(*StoragePoolObject)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetStoragePoolObjectByID indicates an expected call of GetStoragePoolObjectByID.
func (mr *MockStoragePoolServiceIfaceMockRecorder) GetStoragePoolObjectByID(id any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{id}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStoragePoolObjectByID", reflect.TypeOf((*MockStoragePoolServiceIface)(nil).GetStoragePoolObjectByID), varargs...)
}

// GetStoragePoolsMetricByID mocks base method.
func (m *MockStoragePoolServiceIface) GetStoragePoolsMetricByID(id string, opts ...OptionFunc) (*StoragePoolsMetric, int, error) {
	m.ctrl.T.Helper()
	varargs := []any{id}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetStoragePoolsMetricByID", varargs...)
	ret0, _ := ret[0].(*StoragePoolsMetric)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetStoragePoolsMetricByID indicates an expected call of GetStoragePoolsMetricByID.
func (mr *MockStoragePoolServiceIfaceMockRecorder) GetStoragePoolsMetricByID(id any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{id}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStoragePoolsMetricByID", reflect.TypeOf((*MockStoragePoolServiceIface)(nil).GetStoragePoolsMetricByID), varargs...)
}

// GetStoragePoolsMetricByName mocks base method.
func (m *MockStoragePoolServiceIface) GetStoragePoolsMetricByName(name string, opts ...OptionFunc) (*StoragePoolsMetric, int, error) {
	m.ctrl.T.Helper()
	varargs := []any{name}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetStoragePoolsMetricByName", varargs...)
	ret0, _ := ret[0].(*StoragePoolsMetric)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetStoragePoolsMetricByName indicates an expected call of GetStoragePoolsMetricByName.
func (mr *MockStoragePoolServiceIfaceMockRecorder) GetStoragePoolsMetricByName(name any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{name}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStoragePoolsMetricByName", reflect.TypeOf((*MockStoragePoolServiceIface)(nil).GetStoragePoolsMetricByName), varargs...)
}

// GetStoragePoolsMetricID mocks base method.
func (m *MockStoragePoolServiceIface) GetStoragePoolsMetricID(name string, opts ...OptionFunc) (string, int, error) {
	m.ctrl.T.Helper()
	varargs := []any{name}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetStoragePoolsMetricID", varargs...)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetStoragePoolsMetricID indicates an expected call of GetStoragePoolsMetricID.
func (mr *MockStoragePoolServiceIfaceMockRecorder) GetStoragePoolsMetricID(name any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{name}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStoragePoolsMetricID", reflect.TypeOf((*MockStoragePoolServiceIface)(nil).GetStoragePoolsMetricID), varargs...)
}

// ListAffectedVmsForStorageScopeChange mocks base method.
func (m *MockStoragePoolServiceIface) ListAffectedVmsForStorageScopeChange(p *ListAffectedVmsForStorageScopeChangeParams) (*ListAffectedVmsForStorageScopeChangeResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAffectedVmsForStorageScopeChange", p)
	ret0, _ := ret[0].(*ListAffectedVmsForStorageScopeChangeResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAffectedVmsForStorageScopeChange indicates an expected call of ListAffectedVmsForStorageScopeChange.
func (mr *MockStoragePoolServiceIfaceMockRecorder) ListAffectedVmsForStorageScopeChange(p any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAffectedVmsForStorageScopeChange", reflect.TypeOf((*MockStoragePoolServiceIface)(nil).ListAffectedVmsForStorageScopeChange), p)
}

// ListObjectStoragePools mocks base method.
func (m *MockStoragePoolServiceIface) ListObjectStoragePools(p *ListObjectStoragePoolsParams) (*ListObjectStoragePoolsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListObjectStoragePools", p)
	ret0, _ := ret[0].(*ListObjectStoragePoolsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListObjectStoragePools indicates an expected call of ListObjectStoragePools.
func (mr *MockStoragePoolServiceIfaceMockRecorder) ListObjectStoragePools(p any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListObjectStoragePools", reflect.TypeOf((*MockStoragePoolServiceIface)(nil).ListObjectStoragePools), p)
}

// ListStoragePoolObjects mocks base method.
func (m *MockStoragePoolServiceIface) ListStoragePoolObjects(p *ListStoragePoolObjectsParams) (*ListStoragePoolObjectsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListStoragePoolObjects", p)
	ret0, _ := ret[0].(*ListStoragePoolObjectsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListStoragePoolObjects indicates an expected call of ListStoragePoolObjects.
func (mr *MockStoragePoolServiceIfaceMockRecorder) ListStoragePoolObjects(p any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListStoragePoolObjects", reflect.TypeOf((*MockStoragePoolServiceIface)(nil).ListStoragePoolObjects), p)
}

// ListStoragePoolsMetrics mocks base method.
func (m *MockStoragePoolServiceIface) ListStoragePoolsMetrics(p *ListStoragePoolsMetricsParams) (*ListStoragePoolsMetricsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListStoragePoolsMetrics", p)
	ret0, _ := ret[0].(*ListStoragePoolsMetricsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListStoragePoolsMetrics indicates an expected call of ListStoragePoolsMetrics.
func (mr *MockStoragePoolServiceIfaceMockRecorder) ListStoragePoolsMetrics(p any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListStoragePoolsMetrics", reflect.TypeOf((*MockStoragePoolServiceIface)(nil).ListStoragePoolsMetrics), p)
}

// ListStorageProviders mocks base method.
func (m *MockStoragePoolServiceIface) ListStorageProviders(p *ListStorageProvidersParams) (*ListStorageProvidersResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListStorageProviders", p)
	ret0, _ := ret[0].(*ListStorageProvidersResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListStorageProviders indicates an expected call of ListStorageProviders.
func (mr *MockStoragePoolServiceIfaceMockRecorder) ListStorageProviders(p any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListStorageProviders", reflect.TypeOf((*MockStoragePoolServiceIface)(nil).ListStorageProviders), p)
}

// NewAddObjectStoragePoolParams mocks base method.
func (m *MockStoragePoolServiceIface) NewAddObjectStoragePoolParams(name, provider, url string) *AddObjectStoragePoolParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewAddObjectStoragePoolParams", name, provider, url)
	ret0, _ := ret[0].(*AddObjectStoragePoolParams)
	return ret0
}

// NewAddObjectStoragePoolParams indicates an expected call of NewAddObjectStoragePoolParams.
func (mr *MockStoragePoolServiceIfaceMockRecorder) NewAddObjectStoragePoolParams(name, provider, url any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewAddObjectStoragePoolParams", reflect.TypeOf((*MockStoragePoolServiceIface)(nil).NewAddObjectStoragePoolParams), name, provider, url)
}

// NewCancelStorageMaintenanceParams mocks base method.
func (m *MockStoragePoolServiceIface) NewCancelStorageMaintenanceParams(id string) *CancelStorageMaintenanceParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewCancelStorageMaintenanceParams", id)
	ret0, _ := ret[0].(*CancelStorageMaintenanceParams)
	return ret0
}

// NewCancelStorageMaintenanceParams indicates an expected call of NewCancelStorageMaintenanceParams.
func (mr *MockStoragePoolServiceIfaceMockRecorder) NewCancelStorageMaintenanceParams(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewCancelStorageMaintenanceParams", reflect.TypeOf((*MockStoragePoolServiceIface)(nil).NewCancelStorageMaintenanceParams), id)
}

// NewChangeStoragePoolScopeParams mocks base method.
func (m *MockStoragePoolServiceIface) NewChangeStoragePoolScopeParams(id, scope string) *ChangeStoragePoolScopeParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewChangeStoragePoolScopeParams", id, scope)
	ret0, _ := ret[0].(*ChangeStoragePoolScopeParams)
	return ret0
}

// NewChangeStoragePoolScopeParams indicates an expected call of NewChangeStoragePoolScopeParams.
func (mr *MockStoragePoolServiceIfaceMockRecorder) NewChangeStoragePoolScopeParams(id, scope any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewChangeStoragePoolScopeParams", reflect.TypeOf((*MockStoragePoolServiceIface)(nil).NewChangeStoragePoolScopeParams), id, scope)
}

// NewDeleteObjectStoragePoolParams mocks base method.
func (m *MockStoragePoolServiceIface) NewDeleteObjectStoragePoolParams(id string) *DeleteObjectStoragePoolParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewDeleteObjectStoragePoolParams", id)
	ret0, _ := ret[0].(*DeleteObjectStoragePoolParams)
	return ret0
}

// NewDeleteObjectStoragePoolParams indicates an expected call of NewDeleteObjectStoragePoolParams.
func (mr *MockStoragePoolServiceIfaceMockRecorder) NewDeleteObjectStoragePoolParams(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewDeleteObjectStoragePoolParams", reflect.TypeOf((*MockStoragePoolServiceIface)(nil).NewDeleteObjectStoragePoolParams), id)
}

// NewEnableStorageMaintenanceParams mocks base method.
func (m *MockStoragePoolServiceIface) NewEnableStorageMaintenanceParams(id string) *EnableStorageMaintenanceParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewEnableStorageMaintenanceParams", id)
	ret0, _ := ret[0].(*EnableStorageMaintenanceParams)
	return ret0
}

// NewEnableStorageMaintenanceParams indicates an expected call of NewEnableStorageMaintenanceParams.
func (mr *MockStoragePoolServiceIfaceMockRecorder) NewEnableStorageMaintenanceParams(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewEnableStorageMaintenanceParams", reflect.TypeOf((*MockStoragePoolServiceIface)(nil).NewEnableStorageMaintenanceParams), id)
}

// NewListAffectedVmsForStorageScopeChangeParams mocks base method.
func (m *MockStoragePoolServiceIface) NewListAffectedVmsForStorageScopeChangeParams(clusterid, storageid string) *ListAffectedVmsForStorageScopeChangeParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewListAffectedVmsForStorageScopeChangeParams", clusterid, storageid)
	ret0, _ := ret[0].(*ListAffectedVmsForStorageScopeChangeParams)
	return ret0
}

// NewListAffectedVmsForStorageScopeChangeParams indicates an expected call of NewListAffectedVmsForStorageScopeChangeParams.
func (mr *MockStoragePoolServiceIfaceMockRecorder) NewListAffectedVmsForStorageScopeChangeParams(clusterid, storageid any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewListAffectedVmsForStorageScopeChangeParams", reflect.TypeOf((*MockStoragePoolServiceIface)(nil).NewListAffectedVmsForStorageScopeChangeParams), clusterid, storageid)
}

// NewListObjectStoragePoolsParams mocks base method.
func (m *MockStoragePoolServiceIface) NewListObjectStoragePoolsParams() *ListObjectStoragePoolsParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewListObjectStoragePoolsParams")
	ret0, _ := ret[0].(*ListObjectStoragePoolsParams)
	return ret0
}

// NewListObjectStoragePoolsParams indicates an expected call of NewListObjectStoragePoolsParams.
func (mr *MockStoragePoolServiceIfaceMockRecorder) NewListObjectStoragePoolsParams() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewListObjectStoragePoolsParams", reflect.TypeOf((*MockStoragePoolServiceIface)(nil).NewListObjectStoragePoolsParams))
}

// NewListStoragePoolObjectsParams mocks base method.
func (m *MockStoragePoolServiceIface) NewListStoragePoolObjectsParams(id string) *ListStoragePoolObjectsParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewListStoragePoolObjectsParams", id)
	ret0, _ := ret[0].(*ListStoragePoolObjectsParams)
	return ret0
}

// NewListStoragePoolObjectsParams indicates an expected call of NewListStoragePoolObjectsParams.
func (mr *MockStoragePoolServiceIfaceMockRecorder) NewListStoragePoolObjectsParams(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewListStoragePoolObjectsParams", reflect.TypeOf((*MockStoragePoolServiceIface)(nil).NewListStoragePoolObjectsParams), id)
}

// NewListStoragePoolsMetricsParams mocks base method.
func (m *MockStoragePoolServiceIface) NewListStoragePoolsMetricsParams() *ListStoragePoolsMetricsParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewListStoragePoolsMetricsParams")
	ret0, _ := ret[0].(*ListStoragePoolsMetricsParams)
	return ret0
}

// NewListStoragePoolsMetricsParams indicates an expected call of NewListStoragePoolsMetricsParams.
func (mr *MockStoragePoolServiceIfaceMockRecorder) NewListStoragePoolsMetricsParams() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewListStoragePoolsMetricsParams", reflect.TypeOf((*MockStoragePoolServiceIface)(nil).NewListStoragePoolsMetricsParams))
}

// NewListStorageProvidersParams mocks base method.
func (m *MockStoragePoolServiceIface) NewListStorageProvidersParams(storagePoolType string) *ListStorageProvidersParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewListStorageProvidersParams", storagePoolType)
	ret0, _ := ret[0].(*ListStorageProvidersParams)
	return ret0
}

// NewListStorageProvidersParams indicates an expected call of NewListStorageProvidersParams.
func (mr *MockStoragePoolServiceIfaceMockRecorder) NewListStorageProvidersParams(storagePoolType any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewListStorageProvidersParams", reflect.TypeOf((*MockStoragePoolServiceIface)(nil).NewListStorageProvidersParams), storagePoolType)
}

// NewUpdateObjectStoragePoolParams mocks base method.
func (m *MockStoragePoolServiceIface) NewUpdateObjectStoragePoolParams(id string) *UpdateObjectStoragePoolParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewUpdateObjectStoragePoolParams", id)
	ret0, _ := ret[0].(*UpdateObjectStoragePoolParams)
	return ret0
}

// NewUpdateObjectStoragePoolParams indicates an expected call of NewUpdateObjectStoragePoolParams.
func (mr *MockStoragePoolServiceIfaceMockRecorder) NewUpdateObjectStoragePoolParams(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewUpdateObjectStoragePoolParams", reflect.TypeOf((*MockStoragePoolServiceIface)(nil).NewUpdateObjectStoragePoolParams), id)
}

// UpdateObjectStoragePool mocks base method.
func (m *MockStoragePoolServiceIface) UpdateObjectStoragePool(p *UpdateObjectStoragePoolParams) (*UpdateObjectStoragePoolResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateObjectStoragePool", p)
	ret0, _ := ret[0].(*UpdateObjectStoragePoolResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateObjectStoragePool indicates an expected call of UpdateObjectStoragePool.
func (mr *MockStoragePoolServiceIfaceMockRecorder) UpdateObjectStoragePool(p any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateObjectStoragePool", reflect.TypeOf((*MockStoragePoolServiceIface)(nil).UpdateObjectStoragePool), p)
}
