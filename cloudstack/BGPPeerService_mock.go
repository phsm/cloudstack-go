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
// Source: ./cloudstack/BGPPeerService.go
//
// Generated by this command:
//
//	mockgen -destination=./cloudstack/BGPPeerService_mock.go -package=cloudstack -copyright_file=header.txt -source=./cloudstack/BGPPeerService.go
//

// Package cloudstack is a generated GoMock package.
package cloudstack

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockBGPPeerServiceIface is a mock of BGPPeerServiceIface interface.
type MockBGPPeerServiceIface struct {
	ctrl     *gomock.Controller
	recorder *MockBGPPeerServiceIfaceMockRecorder
	isgomock struct{}
}

// MockBGPPeerServiceIfaceMockRecorder is the mock recorder for MockBGPPeerServiceIface.
type MockBGPPeerServiceIfaceMockRecorder struct {
	mock *MockBGPPeerServiceIface
}

// NewMockBGPPeerServiceIface creates a new mock instance.
func NewMockBGPPeerServiceIface(ctrl *gomock.Controller) *MockBGPPeerServiceIface {
	mock := &MockBGPPeerServiceIface{ctrl: ctrl}
	mock.recorder = &MockBGPPeerServiceIfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBGPPeerServiceIface) EXPECT() *MockBGPPeerServiceIfaceMockRecorder {
	return m.recorder
}

// ChangeBgpPeersForVpc mocks base method.
func (m *MockBGPPeerServiceIface) ChangeBgpPeersForVpc(p *ChangeBgpPeersForVpcParams) (*ChangeBgpPeersForVpcResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangeBgpPeersForVpc", p)
	ret0, _ := ret[0].(*ChangeBgpPeersForVpcResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ChangeBgpPeersForVpc indicates an expected call of ChangeBgpPeersForVpc.
func (mr *MockBGPPeerServiceIfaceMockRecorder) ChangeBgpPeersForVpc(p any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeBgpPeersForVpc", reflect.TypeOf((*MockBGPPeerServiceIface)(nil).ChangeBgpPeersForVpc), p)
}

// CreateBgpPeer mocks base method.
func (m *MockBGPPeerServiceIface) CreateBgpPeer(p *CreateBgpPeerParams) (*CreateBgpPeerResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateBgpPeer", p)
	ret0, _ := ret[0].(*CreateBgpPeerResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateBgpPeer indicates an expected call of CreateBgpPeer.
func (mr *MockBGPPeerServiceIfaceMockRecorder) CreateBgpPeer(p any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBgpPeer", reflect.TypeOf((*MockBGPPeerServiceIface)(nil).CreateBgpPeer), p)
}

// DedicateBgpPeer mocks base method.
func (m *MockBGPPeerServiceIface) DedicateBgpPeer(p *DedicateBgpPeerParams) (*DedicateBgpPeerResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DedicateBgpPeer", p)
	ret0, _ := ret[0].(*DedicateBgpPeerResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DedicateBgpPeer indicates an expected call of DedicateBgpPeer.
func (mr *MockBGPPeerServiceIfaceMockRecorder) DedicateBgpPeer(p any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DedicateBgpPeer", reflect.TypeOf((*MockBGPPeerServiceIface)(nil).DedicateBgpPeer), p)
}

// DeleteBgpPeer mocks base method.
func (m *MockBGPPeerServiceIface) DeleteBgpPeer(p *DeleteBgpPeerParams) (*DeleteBgpPeerResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteBgpPeer", p)
	ret0, _ := ret[0].(*DeleteBgpPeerResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteBgpPeer indicates an expected call of DeleteBgpPeer.
func (mr *MockBGPPeerServiceIfaceMockRecorder) DeleteBgpPeer(p any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteBgpPeer", reflect.TypeOf((*MockBGPPeerServiceIface)(nil).DeleteBgpPeer), p)
}

// GetBgpPeerByID mocks base method.
func (m *MockBGPPeerServiceIface) GetBgpPeerByID(id string, opts ...OptionFunc) (*BgpPeer, int, error) {
	m.ctrl.T.Helper()
	varargs := []any{id}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetBgpPeerByID", varargs...)
	ret0, _ := ret[0].(*BgpPeer)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetBgpPeerByID indicates an expected call of GetBgpPeerByID.
func (mr *MockBGPPeerServiceIfaceMockRecorder) GetBgpPeerByID(id any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{id}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBgpPeerByID", reflect.TypeOf((*MockBGPPeerServiceIface)(nil).GetBgpPeerByID), varargs...)
}

// ListBgpPeers mocks base method.
func (m *MockBGPPeerServiceIface) ListBgpPeers(p *ListBgpPeersParams) (*ListBgpPeersResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListBgpPeers", p)
	ret0, _ := ret[0].(*ListBgpPeersResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListBgpPeers indicates an expected call of ListBgpPeers.
func (mr *MockBGPPeerServiceIfaceMockRecorder) ListBgpPeers(p any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListBgpPeers", reflect.TypeOf((*MockBGPPeerServiceIface)(nil).ListBgpPeers), p)
}

// NewChangeBgpPeersForVpcParams mocks base method.
func (m *MockBGPPeerServiceIface) NewChangeBgpPeersForVpcParams(vpcid string) *ChangeBgpPeersForVpcParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewChangeBgpPeersForVpcParams", vpcid)
	ret0, _ := ret[0].(*ChangeBgpPeersForVpcParams)
	return ret0
}

// NewChangeBgpPeersForVpcParams indicates an expected call of NewChangeBgpPeersForVpcParams.
func (mr *MockBGPPeerServiceIfaceMockRecorder) NewChangeBgpPeersForVpcParams(vpcid any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewChangeBgpPeersForVpcParams", reflect.TypeOf((*MockBGPPeerServiceIface)(nil).NewChangeBgpPeersForVpcParams), vpcid)
}

// NewCreateBgpPeerParams mocks base method.
func (m *MockBGPPeerServiceIface) NewCreateBgpPeerParams(asnumber int64, zoneid string) *CreateBgpPeerParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewCreateBgpPeerParams", asnumber, zoneid)
	ret0, _ := ret[0].(*CreateBgpPeerParams)
	return ret0
}

// NewCreateBgpPeerParams indicates an expected call of NewCreateBgpPeerParams.
func (mr *MockBGPPeerServiceIfaceMockRecorder) NewCreateBgpPeerParams(asnumber, zoneid any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewCreateBgpPeerParams", reflect.TypeOf((*MockBGPPeerServiceIface)(nil).NewCreateBgpPeerParams), asnumber, zoneid)
}

// NewDedicateBgpPeerParams mocks base method.
func (m *MockBGPPeerServiceIface) NewDedicateBgpPeerParams(id string) *DedicateBgpPeerParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewDedicateBgpPeerParams", id)
	ret0, _ := ret[0].(*DedicateBgpPeerParams)
	return ret0
}

// NewDedicateBgpPeerParams indicates an expected call of NewDedicateBgpPeerParams.
func (mr *MockBGPPeerServiceIfaceMockRecorder) NewDedicateBgpPeerParams(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewDedicateBgpPeerParams", reflect.TypeOf((*MockBGPPeerServiceIface)(nil).NewDedicateBgpPeerParams), id)
}

// NewDeleteBgpPeerParams mocks base method.
func (m *MockBGPPeerServiceIface) NewDeleteBgpPeerParams(id string) *DeleteBgpPeerParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewDeleteBgpPeerParams", id)
	ret0, _ := ret[0].(*DeleteBgpPeerParams)
	return ret0
}

// NewDeleteBgpPeerParams indicates an expected call of NewDeleteBgpPeerParams.
func (mr *MockBGPPeerServiceIfaceMockRecorder) NewDeleteBgpPeerParams(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewDeleteBgpPeerParams", reflect.TypeOf((*MockBGPPeerServiceIface)(nil).NewDeleteBgpPeerParams), id)
}

// NewListBgpPeersParams mocks base method.
func (m *MockBGPPeerServiceIface) NewListBgpPeersParams() *ListBgpPeersParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewListBgpPeersParams")
	ret0, _ := ret[0].(*ListBgpPeersParams)
	return ret0
}

// NewListBgpPeersParams indicates an expected call of NewListBgpPeersParams.
func (mr *MockBGPPeerServiceIfaceMockRecorder) NewListBgpPeersParams() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewListBgpPeersParams", reflect.TypeOf((*MockBGPPeerServiceIface)(nil).NewListBgpPeersParams))
}

// NewReleaseBgpPeerParams mocks base method.
func (m *MockBGPPeerServiceIface) NewReleaseBgpPeerParams(id string) *ReleaseBgpPeerParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewReleaseBgpPeerParams", id)
	ret0, _ := ret[0].(*ReleaseBgpPeerParams)
	return ret0
}

// NewReleaseBgpPeerParams indicates an expected call of NewReleaseBgpPeerParams.
func (mr *MockBGPPeerServiceIfaceMockRecorder) NewReleaseBgpPeerParams(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewReleaseBgpPeerParams", reflect.TypeOf((*MockBGPPeerServiceIface)(nil).NewReleaseBgpPeerParams), id)
}

// NewUpdateBgpPeerParams mocks base method.
func (m *MockBGPPeerServiceIface) NewUpdateBgpPeerParams(id string) *UpdateBgpPeerParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewUpdateBgpPeerParams", id)
	ret0, _ := ret[0].(*UpdateBgpPeerParams)
	return ret0
}

// NewUpdateBgpPeerParams indicates an expected call of NewUpdateBgpPeerParams.
func (mr *MockBGPPeerServiceIfaceMockRecorder) NewUpdateBgpPeerParams(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewUpdateBgpPeerParams", reflect.TypeOf((*MockBGPPeerServiceIface)(nil).NewUpdateBgpPeerParams), id)
}

// ReleaseBgpPeer mocks base method.
func (m *MockBGPPeerServiceIface) ReleaseBgpPeer(p *ReleaseBgpPeerParams) (*ReleaseBgpPeerResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReleaseBgpPeer", p)
	ret0, _ := ret[0].(*ReleaseBgpPeerResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReleaseBgpPeer indicates an expected call of ReleaseBgpPeer.
func (mr *MockBGPPeerServiceIfaceMockRecorder) ReleaseBgpPeer(p any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReleaseBgpPeer", reflect.TypeOf((*MockBGPPeerServiceIface)(nil).ReleaseBgpPeer), p)
}

// UpdateBgpPeer mocks base method.
func (m *MockBGPPeerServiceIface) UpdateBgpPeer(p *UpdateBgpPeerParams) (*UpdateBgpPeerResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateBgpPeer", p)
	ret0, _ := ret[0].(*UpdateBgpPeerResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateBgpPeer indicates an expected call of UpdateBgpPeer.
func (mr *MockBGPPeerServiceIfaceMockRecorder) UpdateBgpPeer(p any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateBgpPeer", reflect.TypeOf((*MockBGPPeerServiceIface)(nil).UpdateBgpPeer), p)
}