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

package cloudstack

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

type StoragePoolServiceIface interface {
	AddObjectStoragePool(p *AddObjectStoragePoolParams) (*AddObjectStoragePoolResponse, error)
	NewAddObjectStoragePoolParams(name string, provider string, url string) *AddObjectStoragePoolParams
	CancelStorageMaintenance(p *CancelStorageMaintenanceParams) (*CancelStorageMaintenanceResponse, error)
	NewCancelStorageMaintenanceParams(id string) *CancelStorageMaintenanceParams
	ChangeStoragePoolScope(p *ChangeStoragePoolScopeParams) (*ChangeStoragePoolScopeResponse, error)
	NewChangeStoragePoolScopeParams(id string, scope string) *ChangeStoragePoolScopeParams
	CreateStoragePool(p *CreateStoragePoolParams) (*CreateStoragePoolResponse, error)
	NewCreateStoragePoolParams(name string, url string, zoneid string) *CreateStoragePoolParams
	DeleteObjectStoragePool(p *DeleteObjectStoragePoolParams) (*DeleteObjectStoragePoolResponse, error)
	NewDeleteObjectStoragePoolParams(id string) *DeleteObjectStoragePoolParams
	DeleteStoragePool(p *DeleteStoragePoolParams) (*DeleteStoragePoolResponse, error)
	NewDeleteStoragePoolParams(id string) *DeleteStoragePoolParams
	EnableStorageMaintenance(p *EnableStorageMaintenanceParams) (*EnableStorageMaintenanceResponse, error)
	NewEnableStorageMaintenanceParams(id string) *EnableStorageMaintenanceParams
	FindStoragePoolsForMigration(p *FindStoragePoolsForMigrationParams) (*FindStoragePoolsForMigrationResponse, error)
	NewFindStoragePoolsForMigrationParams(id string) *FindStoragePoolsForMigrationParams
	ListObjectStoragePools(p *ListObjectStoragePoolsParams) (*ListObjectStoragePoolsResponse, error)
	NewListObjectStoragePoolsParams() *ListObjectStoragePoolsParams
	GetObjectStoragePoolID(name string, opts ...OptionFunc) (string, int, error)
	GetObjectStoragePoolByName(name string, opts ...OptionFunc) (*ObjectStoragePool, int, error)
	GetObjectStoragePoolByID(id string, opts ...OptionFunc) (*ObjectStoragePool, int, error)
	ListStoragePools(p *ListStoragePoolsParams) (*ListStoragePoolsResponse, error)
	NewListStoragePoolsParams() *ListStoragePoolsParams
	GetStoragePoolID(name string, opts ...OptionFunc) (string, int, error)
	GetStoragePoolByName(name string, opts ...OptionFunc) (*StoragePool, int, error)
	GetStoragePoolByID(id string, opts ...OptionFunc) (*StoragePool, int, error)
	ListStoragePoolsMetrics(p *ListStoragePoolsMetricsParams) (*ListStoragePoolsMetricsResponse, error)
	NewListStoragePoolsMetricsParams() *ListStoragePoolsMetricsParams
	GetStoragePoolsMetricID(name string, opts ...OptionFunc) (string, int, error)
	GetStoragePoolsMetricByName(name string, opts ...OptionFunc) (*StoragePoolsMetric, int, error)
	GetStoragePoolsMetricByID(id string, opts ...OptionFunc) (*StoragePoolsMetric, int, error)
	ListStoragePoolObjects(p *ListStoragePoolObjectsParams) (*ListStoragePoolObjectsResponse, error)
	NewListStoragePoolObjectsParams(id string) *ListStoragePoolObjectsParams
	GetStoragePoolObjectByID(id string, opts ...OptionFunc) (*StoragePoolObject, int, error)
	ListStorageProviders(p *ListStorageProvidersParams) (*ListStorageProvidersResponse, error)
	NewListStorageProvidersParams(storagePoolType string) *ListStorageProvidersParams
	SyncStoragePool(p *SyncStoragePoolParams) (*SyncStoragePoolResponse, error)
	NewSyncStoragePoolParams(id string) *SyncStoragePoolParams
	UpdateObjectStoragePool(p *UpdateObjectStoragePoolParams) (*UpdateObjectStoragePoolResponse, error)
	NewUpdateObjectStoragePoolParams(id string) *UpdateObjectStoragePoolParams
	UpdateStoragePool(p *UpdateStoragePoolParams) (*UpdateStoragePoolResponse, error)
	NewUpdateStoragePoolParams(id string) *UpdateStoragePoolParams
}

type AddObjectStoragePoolParams struct {
	p map[string]interface{}
}

func (p *AddObjectStoragePoolParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["details"]; found {
		m := v.(map[string]string)
		for i, k := range getSortedKeysFromMap(m) {
			u.Set(fmt.Sprintf("details[%d].%s", i, k), m[k])
		}
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["provider"]; found {
		u.Set("provider", v.(string))
	}
	if v, found := p.p["tags"]; found {
		u.Set("tags", v.(string))
	}
	if v, found := p.p["url"]; found {
		u.Set("url", v.(string))
	}
	return u
}

func (p *AddObjectStoragePoolParams) SetDetails(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["details"] = v
}

func (p *AddObjectStoragePoolParams) ResetDetails() {
	if p.p != nil && p.p["details"] != nil {
		delete(p.p, "details")
	}
}

func (p *AddObjectStoragePoolParams) GetDetails() (map[string]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["details"].(map[string]string)
	return value, ok
}

func (p *AddObjectStoragePoolParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *AddObjectStoragePoolParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
}

func (p *AddObjectStoragePoolParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *AddObjectStoragePoolParams) SetProvider(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["provider"] = v
}

func (p *AddObjectStoragePoolParams) ResetProvider() {
	if p.p != nil && p.p["provider"] != nil {
		delete(p.p, "provider")
	}
}

func (p *AddObjectStoragePoolParams) GetProvider() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["provider"].(string)
	return value, ok
}

func (p *AddObjectStoragePoolParams) SetTags(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["tags"] = v
}

func (p *AddObjectStoragePoolParams) ResetTags() {
	if p.p != nil && p.p["tags"] != nil {
		delete(p.p, "tags")
	}
}

func (p *AddObjectStoragePoolParams) GetTags() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["tags"].(string)
	return value, ok
}

func (p *AddObjectStoragePoolParams) SetUrl(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["url"] = v
}

func (p *AddObjectStoragePoolParams) ResetUrl() {
	if p.p != nil && p.p["url"] != nil {
		delete(p.p, "url")
	}
}

func (p *AddObjectStoragePoolParams) GetUrl() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["url"].(string)
	return value, ok
}

// You should always use this function to get a new AddObjectStoragePoolParams instance,
// as then you are sure you have configured all required params
func (s *StoragePoolService) NewAddObjectStoragePoolParams(name string, provider string, url string) *AddObjectStoragePoolParams {
	p := &AddObjectStoragePoolParams{}
	p.p = make(map[string]interface{})
	p.p["name"] = name
	p.p["provider"] = provider
	p.p["url"] = url
	return p
}

// Adds a object storage pool
func (s *StoragePoolService) AddObjectStoragePool(p *AddObjectStoragePoolParams) (*AddObjectStoragePoolResponse, error) {
	resp, err := s.cs.newRequest("addObjectStoragePool", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r AddObjectStoragePoolResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type AddObjectStoragePoolResponse struct {
	Hasannotations bool   `json:"hasannotations"`
	Id             string `json:"id"`
	JobID          string `json:"jobid"`
	Jobstatus      int    `json:"jobstatus"`
	Name           string `json:"name"`
	Providername   string `json:"providername"`
	Storagetotal   int64  `json:"storagetotal"`
	Storageused    int64  `json:"storageused"`
	Url            string `json:"url"`
}

type CancelStorageMaintenanceParams struct {
	p map[string]interface{}
}

func (p *CancelStorageMaintenanceParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *CancelStorageMaintenanceParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *CancelStorageMaintenanceParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *CancelStorageMaintenanceParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

// You should always use this function to get a new CancelStorageMaintenanceParams instance,
// as then you are sure you have configured all required params
func (s *StoragePoolService) NewCancelStorageMaintenanceParams(id string) *CancelStorageMaintenanceParams {
	p := &CancelStorageMaintenanceParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Cancels maintenance for primary storage
func (s *StoragePoolService) CancelStorageMaintenance(p *CancelStorageMaintenanceParams) (*CancelStorageMaintenanceResponse, error) {
	resp, err := s.cs.newRequest("cancelStorageMaintenance", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CancelStorageMaintenanceResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	// If we have a async client, we need to wait for the async result
	if s.cs.async {
		b, err := s.cs.GetAsyncJobResult(r.JobID, s.cs.timeout)
		if err != nil {
			if err == AsyncTimeoutErr {
				return &r, err
			}
			return nil, err
		}

		b, err = getRawValue(b)
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}

	return &r, nil
}

type CancelStorageMaintenanceResponse struct {
	Allocatediops        int64             `json:"allocatediops"`
	Capacityiops         int64             `json:"capacityiops"`
	Clusterid            string            `json:"clusterid"`
	Clustername          string            `json:"clustername"`
	Created              string            `json:"created"`
	Disksizeallocated    int64             `json:"disksizeallocated"`
	Disksizetotal        int64             `json:"disksizetotal"`
	Disksizeused         int64             `json:"disksizeused"`
	Hasannotations       bool              `json:"hasannotations"`
	Hypervisor           string            `json:"hypervisor"`
	Id                   string            `json:"id"`
	Ipaddress            string            `json:"ipaddress"`
	Istagarule           bool              `json:"istagarule"`
	JobID                string            `json:"jobid"`
	Jobstatus            int               `json:"jobstatus"`
	Name                 string            `json:"name"`
	Nfsmountopts         string            `json:"nfsmountopts"`
	Overprovisionfactor  string            `json:"overprovisionfactor"`
	Path                 string            `json:"path"`
	Podid                string            `json:"podid"`
	Podname              string            `json:"podname"`
	Provider             string            `json:"provider"`
	Scope                string            `json:"scope"`
	State                string            `json:"state"`
	Storagecapabilities  map[string]string `json:"storagecapabilities"`
	Storagecustomstats   map[string]string `json:"storagecustomstats"`
	Suitableformigration bool              `json:"suitableformigration"`
	Tags                 string            `json:"tags"`
	Type                 string            `json:"type"`
	Zoneid               string            `json:"zoneid"`
	Zonename             string            `json:"zonename"`
}

type ChangeStoragePoolScopeParams struct {
	p map[string]interface{}
}

func (p *ChangeStoragePoolScopeParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["clusterid"]; found {
		u.Set("clusterid", v.(string))
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["scope"]; found {
		u.Set("scope", v.(string))
	}
	return u
}

func (p *ChangeStoragePoolScopeParams) SetClusterid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["clusterid"] = v
}

func (p *ChangeStoragePoolScopeParams) ResetClusterid() {
	if p.p != nil && p.p["clusterid"] != nil {
		delete(p.p, "clusterid")
	}
}

func (p *ChangeStoragePoolScopeParams) GetClusterid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["clusterid"].(string)
	return value, ok
}

func (p *ChangeStoragePoolScopeParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ChangeStoragePoolScopeParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *ChangeStoragePoolScopeParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *ChangeStoragePoolScopeParams) SetScope(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["scope"] = v
}

func (p *ChangeStoragePoolScopeParams) ResetScope() {
	if p.p != nil && p.p["scope"] != nil {
		delete(p.p, "scope")
	}
}

func (p *ChangeStoragePoolScopeParams) GetScope() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["scope"].(string)
	return value, ok
}

// You should always use this function to get a new ChangeStoragePoolScopeParams instance,
// as then you are sure you have configured all required params
func (s *StoragePoolService) NewChangeStoragePoolScopeParams(id string, scope string) *ChangeStoragePoolScopeParams {
	p := &ChangeStoragePoolScopeParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	p.p["scope"] = scope
	return p
}

// Changes the scope of a storage pool when the pool is in Disabled state.This feature is officially tested and supported for Hypervisors: KVM and VMware, Protocols: NFS and Ceph, and Storage Provider: DefaultPrimary. There might be extra steps involved to make this work for other hypervisors and storage options.
func (s *StoragePoolService) ChangeStoragePoolScope(p *ChangeStoragePoolScopeParams) (*ChangeStoragePoolScopeResponse, error) {
	resp, err := s.cs.newRequest("changeStoragePoolScope", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ChangeStoragePoolScopeResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	// If we have a async client, we need to wait for the async result
	if s.cs.async {
		b, err := s.cs.GetAsyncJobResult(r.JobID, s.cs.timeout)
		if err != nil {
			if err == AsyncTimeoutErr {
				return &r, err
			}
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}

	return &r, nil
}

type ChangeStoragePoolScopeResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type CreateStoragePoolParams struct {
	p map[string]interface{}
}

func (p *CreateStoragePoolParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["capacitybytes"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("capacitybytes", vv)
	}
	if v, found := p.p["capacityiops"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("capacityiops", vv)
	}
	if v, found := p.p["clusterid"]; found {
		u.Set("clusterid", v.(string))
	}
	if v, found := p.p["details"]; found {
		m := v.(map[string]string)
		for i, k := range getSortedKeysFromMap(m) {
			u.Set(fmt.Sprintf("details[%d].%s", i, k), m[k])
		}
	}
	if v, found := p.p["hypervisor"]; found {
		u.Set("hypervisor", v.(string))
	}
	if v, found := p.p["istagarule"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("istagarule", vv)
	}
	if v, found := p.p["managed"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("managed", vv)
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["podid"]; found {
		u.Set("podid", v.(string))
	}
	if v, found := p.p["provider"]; found {
		u.Set("provider", v.(string))
	}
	if v, found := p.p["scope"]; found {
		u.Set("scope", v.(string))
	}
	if v, found := p.p["tags"]; found {
		u.Set("tags", v.(string))
	}
	if v, found := p.p["url"]; found {
		u.Set("url", v.(string))
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *CreateStoragePoolParams) SetCapacitybytes(v int64) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["capacitybytes"] = v
}

func (p *CreateStoragePoolParams) ResetCapacitybytes() {
	if p.p != nil && p.p["capacitybytes"] != nil {
		delete(p.p, "capacitybytes")
	}
}

func (p *CreateStoragePoolParams) GetCapacitybytes() (int64, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["capacitybytes"].(int64)
	return value, ok
}

func (p *CreateStoragePoolParams) SetCapacityiops(v int64) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["capacityiops"] = v
}

func (p *CreateStoragePoolParams) ResetCapacityiops() {
	if p.p != nil && p.p["capacityiops"] != nil {
		delete(p.p, "capacityiops")
	}
}

func (p *CreateStoragePoolParams) GetCapacityiops() (int64, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["capacityiops"].(int64)
	return value, ok
}

func (p *CreateStoragePoolParams) SetClusterid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["clusterid"] = v
}

func (p *CreateStoragePoolParams) ResetClusterid() {
	if p.p != nil && p.p["clusterid"] != nil {
		delete(p.p, "clusterid")
	}
}

func (p *CreateStoragePoolParams) GetClusterid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["clusterid"].(string)
	return value, ok
}

func (p *CreateStoragePoolParams) SetDetails(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["details"] = v
}

func (p *CreateStoragePoolParams) ResetDetails() {
	if p.p != nil && p.p["details"] != nil {
		delete(p.p, "details")
	}
}

func (p *CreateStoragePoolParams) GetDetails() (map[string]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["details"].(map[string]string)
	return value, ok
}

func (p *CreateStoragePoolParams) SetHypervisor(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["hypervisor"] = v
}

func (p *CreateStoragePoolParams) ResetHypervisor() {
	if p.p != nil && p.p["hypervisor"] != nil {
		delete(p.p, "hypervisor")
	}
}

func (p *CreateStoragePoolParams) GetHypervisor() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["hypervisor"].(string)
	return value, ok
}

func (p *CreateStoragePoolParams) SetIstagarule(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["istagarule"] = v
}

func (p *CreateStoragePoolParams) ResetIstagarule() {
	if p.p != nil && p.p["istagarule"] != nil {
		delete(p.p, "istagarule")
	}
}

func (p *CreateStoragePoolParams) GetIstagarule() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["istagarule"].(bool)
	return value, ok
}

func (p *CreateStoragePoolParams) SetManaged(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["managed"] = v
}

func (p *CreateStoragePoolParams) ResetManaged() {
	if p.p != nil && p.p["managed"] != nil {
		delete(p.p, "managed")
	}
}

func (p *CreateStoragePoolParams) GetManaged() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["managed"].(bool)
	return value, ok
}

func (p *CreateStoragePoolParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *CreateStoragePoolParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
}

func (p *CreateStoragePoolParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *CreateStoragePoolParams) SetPodid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["podid"] = v
}

func (p *CreateStoragePoolParams) ResetPodid() {
	if p.p != nil && p.p["podid"] != nil {
		delete(p.p, "podid")
	}
}

func (p *CreateStoragePoolParams) GetPodid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["podid"].(string)
	return value, ok
}

func (p *CreateStoragePoolParams) SetProvider(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["provider"] = v
}

func (p *CreateStoragePoolParams) ResetProvider() {
	if p.p != nil && p.p["provider"] != nil {
		delete(p.p, "provider")
	}
}

func (p *CreateStoragePoolParams) GetProvider() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["provider"].(string)
	return value, ok
}

func (p *CreateStoragePoolParams) SetScope(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["scope"] = v
}

func (p *CreateStoragePoolParams) ResetScope() {
	if p.p != nil && p.p["scope"] != nil {
		delete(p.p, "scope")
	}
}

func (p *CreateStoragePoolParams) GetScope() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["scope"].(string)
	return value, ok
}

func (p *CreateStoragePoolParams) SetTags(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["tags"] = v
}

func (p *CreateStoragePoolParams) ResetTags() {
	if p.p != nil && p.p["tags"] != nil {
		delete(p.p, "tags")
	}
}

func (p *CreateStoragePoolParams) GetTags() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["tags"].(string)
	return value, ok
}

func (p *CreateStoragePoolParams) SetUrl(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["url"] = v
}

func (p *CreateStoragePoolParams) ResetUrl() {
	if p.p != nil && p.p["url"] != nil {
		delete(p.p, "url")
	}
}

func (p *CreateStoragePoolParams) GetUrl() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["url"].(string)
	return value, ok
}

func (p *CreateStoragePoolParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

func (p *CreateStoragePoolParams) ResetZoneid() {
	if p.p != nil && p.p["zoneid"] != nil {
		delete(p.p, "zoneid")
	}
}

func (p *CreateStoragePoolParams) GetZoneid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new CreateStoragePoolParams instance,
// as then you are sure you have configured all required params
func (s *StoragePoolService) NewCreateStoragePoolParams(name string, url string, zoneid string) *CreateStoragePoolParams {
	p := &CreateStoragePoolParams{}
	p.p = make(map[string]interface{})
	p.p["name"] = name
	p.p["url"] = url
	p.p["zoneid"] = zoneid
	return p
}

// Creates a storage pool.
func (s *StoragePoolService) CreateStoragePool(p *CreateStoragePoolParams) (*CreateStoragePoolResponse, error) {
	resp, err := s.cs.newRequest("createStoragePool", p.toURLValues())
	if err != nil {
		return nil, err
	}

	if resp, err = getRawValue(resp); err != nil {
		return nil, err
	}

	var r CreateStoragePoolResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type CreateStoragePoolResponse struct {
	Allocatediops        int64             `json:"allocatediops"`
	Capacityiops         int64             `json:"capacityiops"`
	Clusterid            string            `json:"clusterid"`
	Clustername          string            `json:"clustername"`
	Created              string            `json:"created"`
	Disksizeallocated    int64             `json:"disksizeallocated"`
	Disksizetotal        int64             `json:"disksizetotal"`
	Disksizeused         int64             `json:"disksizeused"`
	Hasannotations       bool              `json:"hasannotations"`
	Hypervisor           string            `json:"hypervisor"`
	Id                   string            `json:"id"`
	Ipaddress            string            `json:"ipaddress"`
	Istagarule           bool              `json:"istagarule"`
	JobID                string            `json:"jobid"`
	Jobstatus            int               `json:"jobstatus"`
	Name                 string            `json:"name"`
	Nfsmountopts         string            `json:"nfsmountopts"`
	Overprovisionfactor  string            `json:"overprovisionfactor"`
	Path                 string            `json:"path"`
	Podid                string            `json:"podid"`
	Podname              string            `json:"podname"`
	Provider             string            `json:"provider"`
	Scope                string            `json:"scope"`
	State                string            `json:"state"`
	Storagecapabilities  map[string]string `json:"storagecapabilities"`
	Storagecustomstats   map[string]string `json:"storagecustomstats"`
	Suitableformigration bool              `json:"suitableformigration"`
	Tags                 string            `json:"tags"`
	Type                 string            `json:"type"`
	Zoneid               string            `json:"zoneid"`
	Zonename             string            `json:"zonename"`
}

type DeleteObjectStoragePoolParams struct {
	p map[string]interface{}
}

func (p *DeleteObjectStoragePoolParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *DeleteObjectStoragePoolParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *DeleteObjectStoragePoolParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *DeleteObjectStoragePoolParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

// You should always use this function to get a new DeleteObjectStoragePoolParams instance,
// as then you are sure you have configured all required params
func (s *StoragePoolService) NewDeleteObjectStoragePoolParams(id string) *DeleteObjectStoragePoolParams {
	p := &DeleteObjectStoragePoolParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Deletes an Object Storage Pool
func (s *StoragePoolService) DeleteObjectStoragePool(p *DeleteObjectStoragePoolParams) (*DeleteObjectStoragePoolResponse, error) {
	resp, err := s.cs.newRequest("deleteObjectStoragePool", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteObjectStoragePoolResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type DeleteObjectStoragePoolResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

func (r *DeleteObjectStoragePoolResponse) UnmarshalJSON(b []byte) error {
	var m map[string]interface{}
	err := json.Unmarshal(b, &m)
	if err != nil {
		return err
	}

	if success, ok := m["success"].(string); ok {
		m["success"] = success == "true"
		b, err = json.Marshal(m)
		if err != nil {
			return err
		}
	}

	if ostypeid, ok := m["ostypeid"].(float64); ok {
		m["ostypeid"] = strconv.Itoa(int(ostypeid))
		b, err = json.Marshal(m)
		if err != nil {
			return err
		}
	}

	type alias DeleteObjectStoragePoolResponse
	return json.Unmarshal(b, (*alias)(r))
}

type DeleteStoragePoolParams struct {
	p map[string]interface{}
}

func (p *DeleteStoragePoolParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["forced"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("forced", vv)
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *DeleteStoragePoolParams) SetForced(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["forced"] = v
}

func (p *DeleteStoragePoolParams) ResetForced() {
	if p.p != nil && p.p["forced"] != nil {
		delete(p.p, "forced")
	}
}

func (p *DeleteStoragePoolParams) GetForced() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["forced"].(bool)
	return value, ok
}

func (p *DeleteStoragePoolParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *DeleteStoragePoolParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *DeleteStoragePoolParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

// You should always use this function to get a new DeleteStoragePoolParams instance,
// as then you are sure you have configured all required params
func (s *StoragePoolService) NewDeleteStoragePoolParams(id string) *DeleteStoragePoolParams {
	p := &DeleteStoragePoolParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Deletes a storage pool.
func (s *StoragePoolService) DeleteStoragePool(p *DeleteStoragePoolParams) (*DeleteStoragePoolResponse, error) {
	resp, err := s.cs.newRequest("deleteStoragePool", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteStoragePoolResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type DeleteStoragePoolResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

func (r *DeleteStoragePoolResponse) UnmarshalJSON(b []byte) error {
	var m map[string]interface{}
	err := json.Unmarshal(b, &m)
	if err != nil {
		return err
	}

	if success, ok := m["success"].(string); ok {
		m["success"] = success == "true"
		b, err = json.Marshal(m)
		if err != nil {
			return err
		}
	}

	if ostypeid, ok := m["ostypeid"].(float64); ok {
		m["ostypeid"] = strconv.Itoa(int(ostypeid))
		b, err = json.Marshal(m)
		if err != nil {
			return err
		}
	}

	type alias DeleteStoragePoolResponse
	return json.Unmarshal(b, (*alias)(r))
}

type EnableStorageMaintenanceParams struct {
	p map[string]interface{}
}

func (p *EnableStorageMaintenanceParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *EnableStorageMaintenanceParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *EnableStorageMaintenanceParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *EnableStorageMaintenanceParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

// You should always use this function to get a new EnableStorageMaintenanceParams instance,
// as then you are sure you have configured all required params
func (s *StoragePoolService) NewEnableStorageMaintenanceParams(id string) *EnableStorageMaintenanceParams {
	p := &EnableStorageMaintenanceParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Puts storage pool into maintenance state
func (s *StoragePoolService) EnableStorageMaintenance(p *EnableStorageMaintenanceParams) (*EnableStorageMaintenanceResponse, error) {
	resp, err := s.cs.newRequest("enableStorageMaintenance", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r EnableStorageMaintenanceResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	// If we have a async client, we need to wait for the async result
	if s.cs.async {
		b, err := s.cs.GetAsyncJobResult(r.JobID, s.cs.timeout)
		if err != nil {
			if err == AsyncTimeoutErr {
				return &r, err
			}
			return nil, err
		}

		b, err = getRawValue(b)
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}

	return &r, nil
}

type EnableStorageMaintenanceResponse struct {
	Allocatediops        int64             `json:"allocatediops"`
	Capacityiops         int64             `json:"capacityiops"`
	Clusterid            string            `json:"clusterid"`
	Clustername          string            `json:"clustername"`
	Created              string            `json:"created"`
	Disksizeallocated    int64             `json:"disksizeallocated"`
	Disksizetotal        int64             `json:"disksizetotal"`
	Disksizeused         int64             `json:"disksizeused"`
	Hasannotations       bool              `json:"hasannotations"`
	Hypervisor           string            `json:"hypervisor"`
	Id                   string            `json:"id"`
	Ipaddress            string            `json:"ipaddress"`
	Istagarule           bool              `json:"istagarule"`
	JobID                string            `json:"jobid"`
	Jobstatus            int               `json:"jobstatus"`
	Name                 string            `json:"name"`
	Nfsmountopts         string            `json:"nfsmountopts"`
	Overprovisionfactor  string            `json:"overprovisionfactor"`
	Path                 string            `json:"path"`
	Podid                string            `json:"podid"`
	Podname              string            `json:"podname"`
	Provider             string            `json:"provider"`
	Scope                string            `json:"scope"`
	State                string            `json:"state"`
	Storagecapabilities  map[string]string `json:"storagecapabilities"`
	Storagecustomstats   map[string]string `json:"storagecustomstats"`
	Suitableformigration bool              `json:"suitableformigration"`
	Tags                 string            `json:"tags"`
	Type                 string            `json:"type"`
	Zoneid               string            `json:"zoneid"`
	Zonename             string            `json:"zonename"`
}

type FindStoragePoolsForMigrationParams struct {
	p map[string]interface{}
}

func (p *FindStoragePoolsForMigrationParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["keyword"]; found {
		u.Set("keyword", v.(string))
	}
	if v, found := p.p["page"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("page", vv)
	}
	if v, found := p.p["pagesize"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("pagesize", vv)
	}
	return u
}

func (p *FindStoragePoolsForMigrationParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *FindStoragePoolsForMigrationParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *FindStoragePoolsForMigrationParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *FindStoragePoolsForMigrationParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *FindStoragePoolsForMigrationParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
}

func (p *FindStoragePoolsForMigrationParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *FindStoragePoolsForMigrationParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *FindStoragePoolsForMigrationParams) ResetPage() {
	if p.p != nil && p.p["page"] != nil {
		delete(p.p, "page")
	}
}

func (p *FindStoragePoolsForMigrationParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *FindStoragePoolsForMigrationParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *FindStoragePoolsForMigrationParams) ResetPagesize() {
	if p.p != nil && p.p["pagesize"] != nil {
		delete(p.p, "pagesize")
	}
}

func (p *FindStoragePoolsForMigrationParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

// You should always use this function to get a new FindStoragePoolsForMigrationParams instance,
// as then you are sure you have configured all required params
func (s *StoragePoolService) NewFindStoragePoolsForMigrationParams(id string) *FindStoragePoolsForMigrationParams {
	p := &FindStoragePoolsForMigrationParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Lists storage pools available for migration of a volume.
func (s *StoragePoolService) FindStoragePoolsForMigration(p *FindStoragePoolsForMigrationParams) (*FindStoragePoolsForMigrationResponse, error) {
	resp, err := s.cs.newRequest("findStoragePoolsForMigration", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r FindStoragePoolsForMigrationResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type FindStoragePoolsForMigrationResponse struct {
	Allocatediops        int64             `json:"allocatediops"`
	Capacityiops         int64             `json:"capacityiops"`
	Clusterid            string            `json:"clusterid"`
	Clustername          string            `json:"clustername"`
	Created              string            `json:"created"`
	Disksizeallocated    int64             `json:"disksizeallocated"`
	Disksizetotal        int64             `json:"disksizetotal"`
	Disksizeused         int64             `json:"disksizeused"`
	Hasannotations       bool              `json:"hasannotations"`
	Hypervisor           string            `json:"hypervisor"`
	Id                   string            `json:"id"`
	Ipaddress            string            `json:"ipaddress"`
	Istagarule           bool              `json:"istagarule"`
	JobID                string            `json:"jobid"`
	Jobstatus            int               `json:"jobstatus"`
	Name                 string            `json:"name"`
	Nfsmountopts         string            `json:"nfsmountopts"`
	Overprovisionfactor  string            `json:"overprovisionfactor"`
	Path                 string            `json:"path"`
	Podid                string            `json:"podid"`
	Podname              string            `json:"podname"`
	Provider             string            `json:"provider"`
	Scope                string            `json:"scope"`
	State                string            `json:"state"`
	Storagecapabilities  map[string]string `json:"storagecapabilities"`
	Storagecustomstats   map[string]string `json:"storagecustomstats"`
	Suitableformigration bool              `json:"suitableformigration"`
	Tags                 string            `json:"tags"`
	Type                 string            `json:"type"`
	Zoneid               string            `json:"zoneid"`
	Zonename             string            `json:"zonename"`
}

type ListObjectStoragePoolsParams struct {
	p map[string]interface{}
}

func (p *ListObjectStoragePoolsParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["keyword"]; found {
		u.Set("keyword", v.(string))
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["page"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("page", vv)
	}
	if v, found := p.p["pagesize"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("pagesize", vv)
	}
	if v, found := p.p["provider"]; found {
		u.Set("provider", v.(string))
	}
	return u
}

func (p *ListObjectStoragePoolsParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ListObjectStoragePoolsParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *ListObjectStoragePoolsParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *ListObjectStoragePoolsParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListObjectStoragePoolsParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
}

func (p *ListObjectStoragePoolsParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListObjectStoragePoolsParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *ListObjectStoragePoolsParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
}

func (p *ListObjectStoragePoolsParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *ListObjectStoragePoolsParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListObjectStoragePoolsParams) ResetPage() {
	if p.p != nil && p.p["page"] != nil {
		delete(p.p, "page")
	}
}

func (p *ListObjectStoragePoolsParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListObjectStoragePoolsParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListObjectStoragePoolsParams) ResetPagesize() {
	if p.p != nil && p.p["pagesize"] != nil {
		delete(p.p, "pagesize")
	}
}

func (p *ListObjectStoragePoolsParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

func (p *ListObjectStoragePoolsParams) SetProvider(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["provider"] = v
}

func (p *ListObjectStoragePoolsParams) ResetProvider() {
	if p.p != nil && p.p["provider"] != nil {
		delete(p.p, "provider")
	}
}

func (p *ListObjectStoragePoolsParams) GetProvider() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["provider"].(string)
	return value, ok
}

// You should always use this function to get a new ListObjectStoragePoolsParams instance,
// as then you are sure you have configured all required params
func (s *StoragePoolService) NewListObjectStoragePoolsParams() *ListObjectStoragePoolsParams {
	p := &ListObjectStoragePoolsParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *StoragePoolService) GetObjectStoragePoolID(name string, opts ...OptionFunc) (string, int, error) {
	p := &ListObjectStoragePoolsParams{}
	p.p = make(map[string]interface{})

	p.p["name"] = name

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return "", -1, err
		}
	}

	l, err := s.ListObjectStoragePools(p)
	if err != nil {
		return "", -1, err
	}

	if l.Count == 0 {
		return "", l.Count, fmt.Errorf("No match found for %s: %+v", name, l)
	}

	if l.Count == 1 {
		return l.ObjectStoragePools[0].Id, l.Count, nil
	}

	if l.Count > 1 {
		for _, v := range l.ObjectStoragePools {
			if v.Name == name {
				return v.Id, l.Count, nil
			}
		}
	}
	return "", l.Count, fmt.Errorf("Could not find an exact match for %s: %+v", name, l)
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *StoragePoolService) GetObjectStoragePoolByName(name string, opts ...OptionFunc) (*ObjectStoragePool, int, error) {
	id, count, err := s.GetObjectStoragePoolID(name, opts...)
	if err != nil {
		return nil, count, err
	}

	r, count, err := s.GetObjectStoragePoolByID(id, opts...)
	if err != nil {
		return nil, count, err
	}
	return r, count, nil
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *StoragePoolService) GetObjectStoragePoolByID(id string, opts ...OptionFunc) (*ObjectStoragePool, int, error) {
	p := &ListObjectStoragePoolsParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListObjectStoragePools(p)
	if err != nil {
		if strings.Contains(err.Error(), fmt.Sprintf(
			"Invalid parameter id value=%s due to incorrect long value format, "+
				"or entity does not exist", id)) {
			return nil, 0, fmt.Errorf("No match found for %s: %+v", id, l)
		}
		return nil, -1, err
	}

	if l.Count == 0 {
		return nil, l.Count, fmt.Errorf("No match found for %s: %+v", id, l)
	}

	if l.Count == 1 {
		return l.ObjectStoragePools[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for ObjectStoragePool UUID: %s!", id)
}

// Lists object storage pools.
func (s *StoragePoolService) ListObjectStoragePools(p *ListObjectStoragePoolsParams) (*ListObjectStoragePoolsResponse, error) {
	resp, err := s.cs.newRequest("listObjectStoragePools", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListObjectStoragePoolsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListObjectStoragePoolsResponse struct {
	Count              int                  `json:"count"`
	ObjectStoragePools []*ObjectStoragePool `json:"objectstoragepool"`
}

type ObjectStoragePool struct {
	Hasannotations bool   `json:"hasannotations"`
	Id             string `json:"id"`
	JobID          string `json:"jobid"`
	Jobstatus      int    `json:"jobstatus"`
	Name           string `json:"name"`
	Providername   string `json:"providername"`
	Storagetotal   int64  `json:"storagetotal"`
	Storageused    int64  `json:"storageused"`
	Url            string `json:"url"`
}

type ListStoragePoolsParams struct {
	p map[string]interface{}
}

func (p *ListStoragePoolsParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["clusterid"]; found {
		u.Set("clusterid", v.(string))
	}
	if v, found := p.p["hostid"]; found {
		u.Set("hostid", v.(string))
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["ipaddress"]; found {
		u.Set("ipaddress", v.(string))
	}
	if v, found := p.p["keyword"]; found {
		u.Set("keyword", v.(string))
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["page"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("page", vv)
	}
	if v, found := p.p["pagesize"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("pagesize", vv)
	}
	if v, found := p.p["path"]; found {
		u.Set("path", v.(string))
	}
	if v, found := p.p["podid"]; found {
		u.Set("podid", v.(string))
	}
	if v, found := p.p["scope"]; found {
		u.Set("scope", v.(string))
	}
	if v, found := p.p["status"]; found {
		u.Set("status", v.(string))
	}
	if v, found := p.p["storagecustomstats"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("storagecustomstats", vv)
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *ListStoragePoolsParams) SetClusterid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["clusterid"] = v
}

func (p *ListStoragePoolsParams) ResetClusterid() {
	if p.p != nil && p.p["clusterid"] != nil {
		delete(p.p, "clusterid")
	}
}

func (p *ListStoragePoolsParams) GetClusterid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["clusterid"].(string)
	return value, ok
}

func (p *ListStoragePoolsParams) SetHostid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["hostid"] = v
}

func (p *ListStoragePoolsParams) ResetHostid() {
	if p.p != nil && p.p["hostid"] != nil {
		delete(p.p, "hostid")
	}
}

func (p *ListStoragePoolsParams) GetHostid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["hostid"].(string)
	return value, ok
}

func (p *ListStoragePoolsParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ListStoragePoolsParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *ListStoragePoolsParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *ListStoragePoolsParams) SetIpaddress(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ipaddress"] = v
}

func (p *ListStoragePoolsParams) ResetIpaddress() {
	if p.p != nil && p.p["ipaddress"] != nil {
		delete(p.p, "ipaddress")
	}
}

func (p *ListStoragePoolsParams) GetIpaddress() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["ipaddress"].(string)
	return value, ok
}

func (p *ListStoragePoolsParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListStoragePoolsParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
}

func (p *ListStoragePoolsParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListStoragePoolsParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *ListStoragePoolsParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
}

func (p *ListStoragePoolsParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *ListStoragePoolsParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListStoragePoolsParams) ResetPage() {
	if p.p != nil && p.p["page"] != nil {
		delete(p.p, "page")
	}
}

func (p *ListStoragePoolsParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListStoragePoolsParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListStoragePoolsParams) ResetPagesize() {
	if p.p != nil && p.p["pagesize"] != nil {
		delete(p.p, "pagesize")
	}
}

func (p *ListStoragePoolsParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

func (p *ListStoragePoolsParams) SetPath(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["path"] = v
}

func (p *ListStoragePoolsParams) ResetPath() {
	if p.p != nil && p.p["path"] != nil {
		delete(p.p, "path")
	}
}

func (p *ListStoragePoolsParams) GetPath() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["path"].(string)
	return value, ok
}

func (p *ListStoragePoolsParams) SetPodid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["podid"] = v
}

func (p *ListStoragePoolsParams) ResetPodid() {
	if p.p != nil && p.p["podid"] != nil {
		delete(p.p, "podid")
	}
}

func (p *ListStoragePoolsParams) GetPodid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["podid"].(string)
	return value, ok
}

func (p *ListStoragePoolsParams) SetScope(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["scope"] = v
}

func (p *ListStoragePoolsParams) ResetScope() {
	if p.p != nil && p.p["scope"] != nil {
		delete(p.p, "scope")
	}
}

func (p *ListStoragePoolsParams) GetScope() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["scope"].(string)
	return value, ok
}

func (p *ListStoragePoolsParams) SetStatus(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["status"] = v
}

func (p *ListStoragePoolsParams) ResetStatus() {
	if p.p != nil && p.p["status"] != nil {
		delete(p.p, "status")
	}
}

func (p *ListStoragePoolsParams) GetStatus() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["status"].(string)
	return value, ok
}

func (p *ListStoragePoolsParams) SetStoragecustomstats(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["storagecustomstats"] = v
}

func (p *ListStoragePoolsParams) ResetStoragecustomstats() {
	if p.p != nil && p.p["storagecustomstats"] != nil {
		delete(p.p, "storagecustomstats")
	}
}

func (p *ListStoragePoolsParams) GetStoragecustomstats() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["storagecustomstats"].(bool)
	return value, ok
}

func (p *ListStoragePoolsParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

func (p *ListStoragePoolsParams) ResetZoneid() {
	if p.p != nil && p.p["zoneid"] != nil {
		delete(p.p, "zoneid")
	}
}

func (p *ListStoragePoolsParams) GetZoneid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new ListStoragePoolsParams instance,
// as then you are sure you have configured all required params
func (s *StoragePoolService) NewListStoragePoolsParams() *ListStoragePoolsParams {
	p := &ListStoragePoolsParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *StoragePoolService) GetStoragePoolID(name string, opts ...OptionFunc) (string, int, error) {
	p := &ListStoragePoolsParams{}
	p.p = make(map[string]interface{})

	p.p["name"] = name

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return "", -1, err
		}
	}

	l, err := s.ListStoragePools(p)
	if err != nil {
		return "", -1, err
	}

	if l.Count == 0 {
		return "", l.Count, fmt.Errorf("No match found for %s: %+v", name, l)
	}

	if l.Count == 1 {
		return l.StoragePools[0].Id, l.Count, nil
	}

	if l.Count > 1 {
		for _, v := range l.StoragePools {
			if v.Name == name {
				return v.Id, l.Count, nil
			}
		}
	}
	return "", l.Count, fmt.Errorf("Could not find an exact match for %s: %+v", name, l)
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *StoragePoolService) GetStoragePoolByName(name string, opts ...OptionFunc) (*StoragePool, int, error) {
	id, count, err := s.GetStoragePoolID(name, opts...)
	if err != nil {
		return nil, count, err
	}

	r, count, err := s.GetStoragePoolByID(id, opts...)
	if err != nil {
		return nil, count, err
	}
	return r, count, nil
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *StoragePoolService) GetStoragePoolByID(id string, opts ...OptionFunc) (*StoragePool, int, error) {
	p := &ListStoragePoolsParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListStoragePools(p)
	if err != nil {
		if strings.Contains(err.Error(), fmt.Sprintf(
			"Invalid parameter id value=%s due to incorrect long value format, "+
				"or entity does not exist", id)) {
			return nil, 0, fmt.Errorf("No match found for %s: %+v", id, l)
		}
		return nil, -1, err
	}

	if l.Count == 0 {
		return nil, l.Count, fmt.Errorf("No match found for %s: %+v", id, l)
	}

	if l.Count == 1 {
		return l.StoragePools[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for StoragePool UUID: %s!", id)
}

// Lists storage pools.
func (s *StoragePoolService) ListStoragePools(p *ListStoragePoolsParams) (*ListStoragePoolsResponse, error) {
	resp, err := s.cs.newRequest("listStoragePools", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListStoragePoolsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListStoragePoolsResponse struct {
	Count        int            `json:"count"`
	StoragePools []*StoragePool `json:"storagepool"`
}

type StoragePool struct {
	Allocatediops        int64             `json:"allocatediops"`
	Capacityiops         int64             `json:"capacityiops"`
	Clusterid            string            `json:"clusterid"`
	Clustername          string            `json:"clustername"`
	Created              string            `json:"created"`
	Disksizeallocated    int64             `json:"disksizeallocated"`
	Disksizetotal        int64             `json:"disksizetotal"`
	Disksizeused         int64             `json:"disksizeused"`
	Hasannotations       bool              `json:"hasannotations"`
	Hypervisor           string            `json:"hypervisor"`
	Id                   string            `json:"id"`
	Ipaddress            string            `json:"ipaddress"`
	Istagarule           bool              `json:"istagarule"`
	JobID                string            `json:"jobid"`
	Jobstatus            int               `json:"jobstatus"`
	Name                 string            `json:"name"`
	Nfsmountopts         string            `json:"nfsmountopts"`
	Overprovisionfactor  string            `json:"overprovisionfactor"`
	Path                 string            `json:"path"`
	Podid                string            `json:"podid"`
	Podname              string            `json:"podname"`
	Provider             string            `json:"provider"`
	Scope                string            `json:"scope"`
	State                string            `json:"state"`
	Storagecapabilities  map[string]string `json:"storagecapabilities"`
	Storagecustomstats   map[string]string `json:"storagecustomstats"`
	Suitableformigration bool              `json:"suitableformigration"`
	Tags                 string            `json:"tags"`
	Type                 string            `json:"type"`
	Zoneid               string            `json:"zoneid"`
	Zonename             string            `json:"zonename"`
}

type ListStoragePoolsMetricsParams struct {
	p map[string]interface{}
}

func (p *ListStoragePoolsMetricsParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["clusterid"]; found {
		u.Set("clusterid", v.(string))
	}
	if v, found := p.p["hostid"]; found {
		u.Set("hostid", v.(string))
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["ipaddress"]; found {
		u.Set("ipaddress", v.(string))
	}
	if v, found := p.p["keyword"]; found {
		u.Set("keyword", v.(string))
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["page"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("page", vv)
	}
	if v, found := p.p["pagesize"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("pagesize", vv)
	}
	if v, found := p.p["path"]; found {
		u.Set("path", v.(string))
	}
	if v, found := p.p["podid"]; found {
		u.Set("podid", v.(string))
	}
	if v, found := p.p["scope"]; found {
		u.Set("scope", v.(string))
	}
	if v, found := p.p["status"]; found {
		u.Set("status", v.(string))
	}
	if v, found := p.p["storagecustomstats"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("storagecustomstats", vv)
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *ListStoragePoolsMetricsParams) SetClusterid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["clusterid"] = v
}

func (p *ListStoragePoolsMetricsParams) ResetClusterid() {
	if p.p != nil && p.p["clusterid"] != nil {
		delete(p.p, "clusterid")
	}
}

func (p *ListStoragePoolsMetricsParams) GetClusterid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["clusterid"].(string)
	return value, ok
}

func (p *ListStoragePoolsMetricsParams) SetHostid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["hostid"] = v
}

func (p *ListStoragePoolsMetricsParams) ResetHostid() {
	if p.p != nil && p.p["hostid"] != nil {
		delete(p.p, "hostid")
	}
}

func (p *ListStoragePoolsMetricsParams) GetHostid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["hostid"].(string)
	return value, ok
}

func (p *ListStoragePoolsMetricsParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ListStoragePoolsMetricsParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *ListStoragePoolsMetricsParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *ListStoragePoolsMetricsParams) SetIpaddress(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ipaddress"] = v
}

func (p *ListStoragePoolsMetricsParams) ResetIpaddress() {
	if p.p != nil && p.p["ipaddress"] != nil {
		delete(p.p, "ipaddress")
	}
}

func (p *ListStoragePoolsMetricsParams) GetIpaddress() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["ipaddress"].(string)
	return value, ok
}

func (p *ListStoragePoolsMetricsParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListStoragePoolsMetricsParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
}

func (p *ListStoragePoolsMetricsParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListStoragePoolsMetricsParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *ListStoragePoolsMetricsParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
}

func (p *ListStoragePoolsMetricsParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *ListStoragePoolsMetricsParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListStoragePoolsMetricsParams) ResetPage() {
	if p.p != nil && p.p["page"] != nil {
		delete(p.p, "page")
	}
}

func (p *ListStoragePoolsMetricsParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListStoragePoolsMetricsParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListStoragePoolsMetricsParams) ResetPagesize() {
	if p.p != nil && p.p["pagesize"] != nil {
		delete(p.p, "pagesize")
	}
}

func (p *ListStoragePoolsMetricsParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

func (p *ListStoragePoolsMetricsParams) SetPath(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["path"] = v
}

func (p *ListStoragePoolsMetricsParams) ResetPath() {
	if p.p != nil && p.p["path"] != nil {
		delete(p.p, "path")
	}
}

func (p *ListStoragePoolsMetricsParams) GetPath() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["path"].(string)
	return value, ok
}

func (p *ListStoragePoolsMetricsParams) SetPodid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["podid"] = v
}

func (p *ListStoragePoolsMetricsParams) ResetPodid() {
	if p.p != nil && p.p["podid"] != nil {
		delete(p.p, "podid")
	}
}

func (p *ListStoragePoolsMetricsParams) GetPodid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["podid"].(string)
	return value, ok
}

func (p *ListStoragePoolsMetricsParams) SetScope(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["scope"] = v
}

func (p *ListStoragePoolsMetricsParams) ResetScope() {
	if p.p != nil && p.p["scope"] != nil {
		delete(p.p, "scope")
	}
}

func (p *ListStoragePoolsMetricsParams) GetScope() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["scope"].(string)
	return value, ok
}

func (p *ListStoragePoolsMetricsParams) SetStatus(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["status"] = v
}

func (p *ListStoragePoolsMetricsParams) ResetStatus() {
	if p.p != nil && p.p["status"] != nil {
		delete(p.p, "status")
	}
}

func (p *ListStoragePoolsMetricsParams) GetStatus() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["status"].(string)
	return value, ok
}

func (p *ListStoragePoolsMetricsParams) SetStoragecustomstats(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["storagecustomstats"] = v
}

func (p *ListStoragePoolsMetricsParams) ResetStoragecustomstats() {
	if p.p != nil && p.p["storagecustomstats"] != nil {
		delete(p.p, "storagecustomstats")
	}
}

func (p *ListStoragePoolsMetricsParams) GetStoragecustomstats() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["storagecustomstats"].(bool)
	return value, ok
}

func (p *ListStoragePoolsMetricsParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

func (p *ListStoragePoolsMetricsParams) ResetZoneid() {
	if p.p != nil && p.p["zoneid"] != nil {
		delete(p.p, "zoneid")
	}
}

func (p *ListStoragePoolsMetricsParams) GetZoneid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new ListStoragePoolsMetricsParams instance,
// as then you are sure you have configured all required params
func (s *StoragePoolService) NewListStoragePoolsMetricsParams() *ListStoragePoolsMetricsParams {
	p := &ListStoragePoolsMetricsParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *StoragePoolService) GetStoragePoolsMetricID(name string, opts ...OptionFunc) (string, int, error) {
	p := &ListStoragePoolsMetricsParams{}
	p.p = make(map[string]interface{})

	p.p["name"] = name

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return "", -1, err
		}
	}

	l, err := s.ListStoragePoolsMetrics(p)
	if err != nil {
		return "", -1, err
	}

	if l.Count == 0 {
		return "", l.Count, fmt.Errorf("No match found for %s: %+v", name, l)
	}

	if l.Count == 1 {
		return l.StoragePoolsMetrics[0].Id, l.Count, nil
	}

	if l.Count > 1 {
		for _, v := range l.StoragePoolsMetrics {
			if v.Name == name {
				return v.Id, l.Count, nil
			}
		}
	}
	return "", l.Count, fmt.Errorf("Could not find an exact match for %s: %+v", name, l)
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *StoragePoolService) GetStoragePoolsMetricByName(name string, opts ...OptionFunc) (*StoragePoolsMetric, int, error) {
	id, count, err := s.GetStoragePoolsMetricID(name, opts...)
	if err != nil {
		return nil, count, err
	}

	r, count, err := s.GetStoragePoolsMetricByID(id, opts...)
	if err != nil {
		return nil, count, err
	}
	return r, count, nil
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *StoragePoolService) GetStoragePoolsMetricByID(id string, opts ...OptionFunc) (*StoragePoolsMetric, int, error) {
	p := &ListStoragePoolsMetricsParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListStoragePoolsMetrics(p)
	if err != nil {
		if strings.Contains(err.Error(), fmt.Sprintf(
			"Invalid parameter id value=%s due to incorrect long value format, "+
				"or entity does not exist", id)) {
			return nil, 0, fmt.Errorf("No match found for %s: %+v", id, l)
		}
		return nil, -1, err
	}

	if l.Count == 0 {
		return nil, l.Count, fmt.Errorf("No match found for %s: %+v", id, l)
	}

	if l.Count == 1 {
		return l.StoragePoolsMetrics[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for StoragePoolsMetric UUID: %s!", id)
}

// Lists storage pool metrics
func (s *StoragePoolService) ListStoragePoolsMetrics(p *ListStoragePoolsMetricsParams) (*ListStoragePoolsMetricsResponse, error) {
	resp, err := s.cs.newRequest("listStoragePoolsMetrics", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListStoragePoolsMetricsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListStoragePoolsMetricsResponse struct {
	Count               int                   `json:"count"`
	StoragePoolsMetrics []*StoragePoolsMetric `json:"storagepoolsmetric"`
}

type StoragePoolsMetric struct {
	Allocatediops                    int64             `json:"allocatediops"`
	Capacityiops                     int64             `json:"capacityiops"`
	Clusterid                        string            `json:"clusterid"`
	Clustername                      string            `json:"clustername"`
	Created                          string            `json:"created"`
	Disksizeallocated                int64             `json:"disksizeallocated"`
	Disksizeallocatedgb              string            `json:"disksizeallocatedgb"`
	Disksizetotal                    int64             `json:"disksizetotal"`
	Disksizetotalgb                  string            `json:"disksizetotalgb"`
	Disksizeunallocatedgb            string            `json:"disksizeunallocatedgb"`
	Disksizeused                     int64             `json:"disksizeused"`
	Disksizeusedgb                   string            `json:"disksizeusedgb"`
	Hasannotations                   bool              `json:"hasannotations"`
	Hypervisor                       string            `json:"hypervisor"`
	Id                               string            `json:"id"`
	Ipaddress                        string            `json:"ipaddress"`
	Istagarule                       bool              `json:"istagarule"`
	JobID                            string            `json:"jobid"`
	Jobstatus                        int               `json:"jobstatus"`
	Name                             string            `json:"name"`
	Nfsmountopts                     string            `json:"nfsmountopts"`
	Overprovisionfactor              string            `json:"overprovisionfactor"`
	Path                             string            `json:"path"`
	Podid                            string            `json:"podid"`
	Podname                          string            `json:"podname"`
	Provider                         string            `json:"provider"`
	Scope                            string            `json:"scope"`
	State                            string            `json:"state"`
	Storageallocateddisablethreshold bool              `json:"storageallocateddisablethreshold"`
	Storageallocatedthreshold        bool              `json:"storageallocatedthreshold"`
	Storagecapabilities              map[string]string `json:"storagecapabilities"`
	Storagecustomstats               map[string]string `json:"storagecustomstats"`
	Storageusagedisablethreshold     bool              `json:"storageusagedisablethreshold"`
	Storageusagethreshold            bool              `json:"storageusagethreshold"`
	Suitableformigration             bool              `json:"suitableformigration"`
	Tags                             string            `json:"tags"`
	Type                             string            `json:"type"`
	Zoneid                           string            `json:"zoneid"`
	Zonename                         string            `json:"zonename"`
}

type ListStoragePoolObjectsParams struct {
	p map[string]interface{}
}

func (p *ListStoragePoolObjectsParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["keyword"]; found {
		u.Set("keyword", v.(string))
	}
	if v, found := p.p["page"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("page", vv)
	}
	if v, found := p.p["pagesize"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("pagesize", vv)
	}
	if v, found := p.p["path"]; found {
		u.Set("path", v.(string))
	}
	return u
}

func (p *ListStoragePoolObjectsParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ListStoragePoolObjectsParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *ListStoragePoolObjectsParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *ListStoragePoolObjectsParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListStoragePoolObjectsParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
}

func (p *ListStoragePoolObjectsParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListStoragePoolObjectsParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListStoragePoolObjectsParams) ResetPage() {
	if p.p != nil && p.p["page"] != nil {
		delete(p.p, "page")
	}
}

func (p *ListStoragePoolObjectsParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListStoragePoolObjectsParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListStoragePoolObjectsParams) ResetPagesize() {
	if p.p != nil && p.p["pagesize"] != nil {
		delete(p.p, "pagesize")
	}
}

func (p *ListStoragePoolObjectsParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

func (p *ListStoragePoolObjectsParams) SetPath(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["path"] = v
}

func (p *ListStoragePoolObjectsParams) ResetPath() {
	if p.p != nil && p.p["path"] != nil {
		delete(p.p, "path")
	}
}

func (p *ListStoragePoolObjectsParams) GetPath() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["path"].(string)
	return value, ok
}

// You should always use this function to get a new ListStoragePoolObjectsParams instance,
// as then you are sure you have configured all required params
func (s *StoragePoolService) NewListStoragePoolObjectsParams(id string) *ListStoragePoolObjectsParams {
	p := &ListStoragePoolObjectsParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *StoragePoolService) GetStoragePoolObjectByID(id string, opts ...OptionFunc) (*StoragePoolObject, int, error) {
	p := &ListStoragePoolObjectsParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return nil, -1, err
		}
	}

	l, err := s.ListStoragePoolObjects(p)
	if err != nil {
		if strings.Contains(err.Error(), fmt.Sprintf(
			"Invalid parameter id value=%s due to incorrect long value format, "+
				"or entity does not exist", id)) {
			return nil, 0, fmt.Errorf("No match found for %s: %+v", id, l)
		}
		return nil, -1, err
	}

	if l.Count == 0 {
		return nil, l.Count, fmt.Errorf("No match found for %s: %+v", id, l)
	}

	if l.Count == 1 {
		return l.StoragePoolObjects[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for StoragePoolObject UUID: %s!", id)
}

// Lists objects at specified path on a storage pool.
func (s *StoragePoolService) ListStoragePoolObjects(p *ListStoragePoolObjectsParams) (*ListStoragePoolObjectsResponse, error) {
	resp, err := s.cs.newRequest("listStoragePoolObjects", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListStoragePoolObjectsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListStoragePoolObjectsResponse struct {
	Count              int                  `json:"count"`
	StoragePoolObjects []*StoragePoolObject `json:"storagepoolobject"`
}

type StoragePoolObject struct {
	Format      string `json:"format"`
	Isdirectory bool   `json:"isdirectory"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Lastupdated string `json:"lastupdated"`
	Name        string `json:"name"`
	Size        int64  `json:"size"`
	Snapshotid  string `json:"snapshotid"`
	Templateid  string `json:"templateid"`
	Volumeid    string `json:"volumeid"`
}

type ListStorageProvidersParams struct {
	p map[string]interface{}
}

func (p *ListStorageProvidersParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["keyword"]; found {
		u.Set("keyword", v.(string))
	}
	if v, found := p.p["page"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("page", vv)
	}
	if v, found := p.p["pagesize"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("pagesize", vv)
	}
	if v, found := p.p["type"]; found {
		u.Set("type", v.(string))
	}
	return u
}

func (p *ListStorageProvidersParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListStorageProvidersParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
}

func (p *ListStorageProvidersParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListStorageProvidersParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListStorageProvidersParams) ResetPage() {
	if p.p != nil && p.p["page"] != nil {
		delete(p.p, "page")
	}
}

func (p *ListStorageProvidersParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListStorageProvidersParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListStorageProvidersParams) ResetPagesize() {
	if p.p != nil && p.p["pagesize"] != nil {
		delete(p.p, "pagesize")
	}
}

func (p *ListStorageProvidersParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

func (p *ListStorageProvidersParams) SetType(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["type"] = v
}

func (p *ListStorageProvidersParams) ResetType() {
	if p.p != nil && p.p["type"] != nil {
		delete(p.p, "type")
	}
}

func (p *ListStorageProvidersParams) GetType() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["type"].(string)
	return value, ok
}

// You should always use this function to get a new ListStorageProvidersParams instance,
// as then you are sure you have configured all required params
func (s *StoragePoolService) NewListStorageProvidersParams(storagePoolType string) *ListStorageProvidersParams {
	p := &ListStorageProvidersParams{}
	p.p = make(map[string]interface{})
	p.p["type"] = storagePoolType
	return p
}

// Lists storage providers.
func (s *StoragePoolService) ListStorageProviders(p *ListStorageProvidersParams) (*ListStorageProvidersResponse, error) {
	resp, err := s.cs.newRequest("listStorageProviders", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListStorageProvidersResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListStorageProvidersResponse struct {
	Count            int                `json:"count"`
	StorageProviders []*StorageProvider `json:"storageprovider"`
}

type StorageProvider struct {
	JobID     string `json:"jobid"`
	Jobstatus int    `json:"jobstatus"`
	Name      string `json:"name"`
	Type      string `json:"type"`
}

type SyncStoragePoolParams struct {
	p map[string]interface{}
}

func (p *SyncStoragePoolParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *SyncStoragePoolParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *SyncStoragePoolParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *SyncStoragePoolParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

// You should always use this function to get a new SyncStoragePoolParams instance,
// as then you are sure you have configured all required params
func (s *StoragePoolService) NewSyncStoragePoolParams(id string) *SyncStoragePoolParams {
	p := &SyncStoragePoolParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Sync storage pool with management server (currently supported for Datastore Cluster in VMware and syncs the datastores in it)
func (s *StoragePoolService) SyncStoragePool(p *SyncStoragePoolParams) (*SyncStoragePoolResponse, error) {
	resp, err := s.cs.newRequest("syncStoragePool", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r SyncStoragePoolResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	// If we have a async client, we need to wait for the async result
	if s.cs.async {
		b, err := s.cs.GetAsyncJobResult(r.JobID, s.cs.timeout)
		if err != nil {
			if err == AsyncTimeoutErr {
				return &r, err
			}
			return nil, err
		}

		b, err = getRawValue(b)
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}

	return &r, nil
}

type SyncStoragePoolResponse struct {
	Allocatediops        int64             `json:"allocatediops"`
	Capacityiops         int64             `json:"capacityiops"`
	Clusterid            string            `json:"clusterid"`
	Clustername          string            `json:"clustername"`
	Created              string            `json:"created"`
	Disksizeallocated    int64             `json:"disksizeallocated"`
	Disksizetotal        int64             `json:"disksizetotal"`
	Disksizeused         int64             `json:"disksizeused"`
	Hasannotations       bool              `json:"hasannotations"`
	Hypervisor           string            `json:"hypervisor"`
	Id                   string            `json:"id"`
	Ipaddress            string            `json:"ipaddress"`
	Istagarule           bool              `json:"istagarule"`
	JobID                string            `json:"jobid"`
	Jobstatus            int               `json:"jobstatus"`
	Name                 string            `json:"name"`
	Nfsmountopts         string            `json:"nfsmountopts"`
	Overprovisionfactor  string            `json:"overprovisionfactor"`
	Path                 string            `json:"path"`
	Podid                string            `json:"podid"`
	Podname              string            `json:"podname"`
	Provider             string            `json:"provider"`
	Scope                string            `json:"scope"`
	State                string            `json:"state"`
	Storagecapabilities  map[string]string `json:"storagecapabilities"`
	Storagecustomstats   map[string]string `json:"storagecustomstats"`
	Suitableformigration bool              `json:"suitableformigration"`
	Tags                 string            `json:"tags"`
	Type                 string            `json:"type"`
	Zoneid               string            `json:"zoneid"`
	Zonename             string            `json:"zonename"`
}

type UpdateObjectStoragePoolParams struct {
	p map[string]interface{}
}

func (p *UpdateObjectStoragePoolParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["url"]; found {
		u.Set("url", v.(string))
	}
	return u
}

func (p *UpdateObjectStoragePoolParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *UpdateObjectStoragePoolParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *UpdateObjectStoragePoolParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *UpdateObjectStoragePoolParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *UpdateObjectStoragePoolParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
}

func (p *UpdateObjectStoragePoolParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *UpdateObjectStoragePoolParams) SetUrl(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["url"] = v
}

func (p *UpdateObjectStoragePoolParams) ResetUrl() {
	if p.p != nil && p.p["url"] != nil {
		delete(p.p, "url")
	}
}

func (p *UpdateObjectStoragePoolParams) GetUrl() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["url"].(string)
	return value, ok
}

// You should always use this function to get a new UpdateObjectStoragePoolParams instance,
// as then you are sure you have configured all required params
func (s *StoragePoolService) NewUpdateObjectStoragePoolParams(id string) *UpdateObjectStoragePoolParams {
	p := &UpdateObjectStoragePoolParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Updates object storage pool
func (s *StoragePoolService) UpdateObjectStoragePool(p *UpdateObjectStoragePoolParams) (*UpdateObjectStoragePoolResponse, error) {
	resp, err := s.cs.newRequest("updateObjectStoragePool", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateObjectStoragePoolResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type UpdateObjectStoragePoolResponse struct {
	Hasannotations bool   `json:"hasannotations"`
	Id             string `json:"id"`
	JobID          string `json:"jobid"`
	Jobstatus      int    `json:"jobstatus"`
	Name           string `json:"name"`
	Providername   string `json:"providername"`
	Storagetotal   int64  `json:"storagetotal"`
	Storageused    int64  `json:"storageused"`
	Url            string `json:"url"`
}

type UpdateStoragePoolParams struct {
	p map[string]interface{}
}

func (p *UpdateStoragePoolParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["capacitybytes"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("capacitybytes", vv)
	}
	if v, found := p.p["capacityiops"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("capacityiops", vv)
	}
	if v, found := p.p["details"]; found {
		m := v.(map[string]string)
		for i, k := range getSortedKeysFromMap(m) {
			u.Set(fmt.Sprintf("details[%d].%s", i, k), m[k])
		}
	}
	if v, found := p.p["enabled"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("enabled", vv)
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["istagarule"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("istagarule", vv)
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["tags"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("tags", vv)
	}
	if v, found := p.p["url"]; found {
		u.Set("url", v.(string))
	}
	return u
}

func (p *UpdateStoragePoolParams) SetCapacitybytes(v int64) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["capacitybytes"] = v
}

func (p *UpdateStoragePoolParams) ResetCapacitybytes() {
	if p.p != nil && p.p["capacitybytes"] != nil {
		delete(p.p, "capacitybytes")
	}
}

func (p *UpdateStoragePoolParams) GetCapacitybytes() (int64, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["capacitybytes"].(int64)
	return value, ok
}

func (p *UpdateStoragePoolParams) SetCapacityiops(v int64) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["capacityiops"] = v
}

func (p *UpdateStoragePoolParams) ResetCapacityiops() {
	if p.p != nil && p.p["capacityiops"] != nil {
		delete(p.p, "capacityiops")
	}
}

func (p *UpdateStoragePoolParams) GetCapacityiops() (int64, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["capacityiops"].(int64)
	return value, ok
}

func (p *UpdateStoragePoolParams) SetDetails(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["details"] = v
}

func (p *UpdateStoragePoolParams) ResetDetails() {
	if p.p != nil && p.p["details"] != nil {
		delete(p.p, "details")
	}
}

func (p *UpdateStoragePoolParams) GetDetails() (map[string]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["details"].(map[string]string)
	return value, ok
}

func (p *UpdateStoragePoolParams) SetEnabled(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["enabled"] = v
}

func (p *UpdateStoragePoolParams) ResetEnabled() {
	if p.p != nil && p.p["enabled"] != nil {
		delete(p.p, "enabled")
	}
}

func (p *UpdateStoragePoolParams) GetEnabled() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["enabled"].(bool)
	return value, ok
}

func (p *UpdateStoragePoolParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *UpdateStoragePoolParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *UpdateStoragePoolParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

func (p *UpdateStoragePoolParams) SetIstagarule(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["istagarule"] = v
}

func (p *UpdateStoragePoolParams) ResetIstagarule() {
	if p.p != nil && p.p["istagarule"] != nil {
		delete(p.p, "istagarule")
	}
}

func (p *UpdateStoragePoolParams) GetIstagarule() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["istagarule"].(bool)
	return value, ok
}

func (p *UpdateStoragePoolParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *UpdateStoragePoolParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
}

func (p *UpdateStoragePoolParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *UpdateStoragePoolParams) SetTags(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["tags"] = v
}

func (p *UpdateStoragePoolParams) ResetTags() {
	if p.p != nil && p.p["tags"] != nil {
		delete(p.p, "tags")
	}
}

func (p *UpdateStoragePoolParams) GetTags() ([]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["tags"].([]string)
	return value, ok
}

func (p *UpdateStoragePoolParams) SetUrl(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["url"] = v
}

func (p *UpdateStoragePoolParams) ResetUrl() {
	if p.p != nil && p.p["url"] != nil {
		delete(p.p, "url")
	}
}

func (p *UpdateStoragePoolParams) GetUrl() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["url"].(string)
	return value, ok
}

// You should always use this function to get a new UpdateStoragePoolParams instance,
// as then you are sure you have configured all required params
func (s *StoragePoolService) NewUpdateStoragePoolParams(id string) *UpdateStoragePoolParams {
	p := &UpdateStoragePoolParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Updates a storage pool.
func (s *StoragePoolService) UpdateStoragePool(p *UpdateStoragePoolParams) (*UpdateStoragePoolResponse, error) {
	resp, err := s.cs.newRequest("updateStoragePool", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateStoragePoolResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type UpdateStoragePoolResponse struct {
	Allocatediops        int64             `json:"allocatediops"`
	Capacityiops         int64             `json:"capacityiops"`
	Clusterid            string            `json:"clusterid"`
	Clustername          string            `json:"clustername"`
	Created              string            `json:"created"`
	Disksizeallocated    int64             `json:"disksizeallocated"`
	Disksizetotal        int64             `json:"disksizetotal"`
	Disksizeused         int64             `json:"disksizeused"`
	Hasannotations       bool              `json:"hasannotations"`
	Hypervisor           string            `json:"hypervisor"`
	Id                   string            `json:"id"`
	Ipaddress            string            `json:"ipaddress"`
	Istagarule           bool              `json:"istagarule"`
	JobID                string            `json:"jobid"`
	Jobstatus            int               `json:"jobstatus"`
	Name                 string            `json:"name"`
	Nfsmountopts         string            `json:"nfsmountopts"`
	Overprovisionfactor  string            `json:"overprovisionfactor"`
	Path                 string            `json:"path"`
	Podid                string            `json:"podid"`
	Podname              string            `json:"podname"`
	Provider             string            `json:"provider"`
	Scope                string            `json:"scope"`
	State                string            `json:"state"`
	Storagecapabilities  map[string]string `json:"storagecapabilities"`
	Storagecustomstats   map[string]string `json:"storagecustomstats"`
	Suitableformigration bool              `json:"suitableformigration"`
	Tags                 string            `json:"tags"`
	Type                 string            `json:"type"`
	Zoneid               string            `json:"zoneid"`
	Zonename             string            `json:"zonename"`
}
