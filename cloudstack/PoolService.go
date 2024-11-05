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
)

type PoolServiceIface interface {
	ListElastistorPool(p *ListElastistorPoolParams) (*ListElastistorPoolResponse, error)
	NewListElastistorPoolParams() *ListElastistorPoolParams
	ListVsphereStoragePolicyCompatiblePools(p *ListVsphereStoragePolicyCompatiblePoolsParams) (*ListVsphereStoragePolicyCompatiblePoolsResponse, error)
	NewListVsphereStoragePolicyCompatiblePoolsParams() *ListVsphereStoragePolicyCompatiblePoolsParams
	GetVsphereStoragePolicyCompatiblePoolID(keyword string, opts ...OptionFunc) (string, int, error)
}

type ListElastistorPoolParams struct {
	p map[string]interface{}
}

func (p *ListElastistorPoolParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("id", vv)
	}
	return u
}

func (p *ListElastistorPoolParams) SetId(v int64) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *ListElastistorPoolParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *ListElastistorPoolParams) GetId() (int64, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(int64)
	return value, ok
}

// You should always use this function to get a new ListElastistorPoolParams instance,
// as then you are sure you have configured all required params
func (s *PoolService) NewListElastistorPoolParams() *ListElastistorPoolParams {
	p := &ListElastistorPoolParams{}
	p.p = make(map[string]interface{})
	return p
}

// Lists the pools of elastistor
func (s *PoolService) ListElastistorPool(p *ListElastistorPoolParams) (*ListElastistorPoolResponse, error) {
	resp, err := s.cs.newRequest("listElastistorPool", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListElastistorPoolResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListElastistorPoolResponse struct {
	Count          int               `json:"count"`
	ElastistorPool []*ElastistorPool `json:"elastistorpool"`
}

type ElastistorPool struct {
	Controllerid string `json:"controllerid"`
	Gateway      string `json:"gateway"`
	Id           string `json:"id"`
	JobID        string `json:"jobid"`
	Jobstatus    int    `json:"jobstatus"`
	Maxiops      int64  `json:"maxiops"`
	Name         string `json:"name"`
	Size         int64  `json:"size"`
	State        string `json:"state"`
}

type ListVsphereStoragePolicyCompatiblePoolsParams struct {
	p map[string]interface{}
}

func (p *ListVsphereStoragePolicyCompatiblePoolsParams) toURLValues() url.Values {
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
	if v, found := p.p["policyid"]; found {
		u.Set("policyid", v.(string))
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *ListVsphereStoragePolicyCompatiblePoolsParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListVsphereStoragePolicyCompatiblePoolsParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
}

func (p *ListVsphereStoragePolicyCompatiblePoolsParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListVsphereStoragePolicyCompatiblePoolsParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListVsphereStoragePolicyCompatiblePoolsParams) ResetPage() {
	if p.p != nil && p.p["page"] != nil {
		delete(p.p, "page")
	}
}

func (p *ListVsphereStoragePolicyCompatiblePoolsParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListVsphereStoragePolicyCompatiblePoolsParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListVsphereStoragePolicyCompatiblePoolsParams) ResetPagesize() {
	if p.p != nil && p.p["pagesize"] != nil {
		delete(p.p, "pagesize")
	}
}

func (p *ListVsphereStoragePolicyCompatiblePoolsParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

func (p *ListVsphereStoragePolicyCompatiblePoolsParams) SetPolicyid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["policyid"] = v
}

func (p *ListVsphereStoragePolicyCompatiblePoolsParams) ResetPolicyid() {
	if p.p != nil && p.p["policyid"] != nil {
		delete(p.p, "policyid")
	}
}

func (p *ListVsphereStoragePolicyCompatiblePoolsParams) GetPolicyid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["policyid"].(string)
	return value, ok
}

func (p *ListVsphereStoragePolicyCompatiblePoolsParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

func (p *ListVsphereStoragePolicyCompatiblePoolsParams) ResetZoneid() {
	if p.p != nil && p.p["zoneid"] != nil {
		delete(p.p, "zoneid")
	}
}

func (p *ListVsphereStoragePolicyCompatiblePoolsParams) GetZoneid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new ListVsphereStoragePolicyCompatiblePoolsParams instance,
// as then you are sure you have configured all required params
func (s *PoolService) NewListVsphereStoragePolicyCompatiblePoolsParams() *ListVsphereStoragePolicyCompatiblePoolsParams {
	p := &ListVsphereStoragePolicyCompatiblePoolsParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *PoolService) GetVsphereStoragePolicyCompatiblePoolID(keyword string, opts ...OptionFunc) (string, int, error) {
	p := &ListVsphereStoragePolicyCompatiblePoolsParams{}
	p.p = make(map[string]interface{})

	p.p["keyword"] = keyword

	for _, fn := range append(s.cs.options, opts...) {
		if err := fn(s.cs, p); err != nil {
			return "", -1, err
		}
	}

	l, err := s.ListVsphereStoragePolicyCompatiblePools(p)
	if err != nil {
		return "", -1, err
	}

	if l.Count == 0 {
		return "", l.Count, fmt.Errorf("No match found for %s: %+v", keyword, l)
	}

	if l.Count == 1 {
		return l.VsphereStoragePolicyCompatiblePools[0].Id, l.Count, nil
	}

	if l.Count > 1 {
		for _, v := range l.VsphereStoragePolicyCompatiblePools {
			if v.Name == keyword {
				return v.Id, l.Count, nil
			}
		}
	}
	return "", l.Count, fmt.Errorf("Could not find an exact match for %s: %+v", keyword, l)
}

// List storage pools compatible with a vSphere storage policy
func (s *PoolService) ListVsphereStoragePolicyCompatiblePools(p *ListVsphereStoragePolicyCompatiblePoolsParams) (*ListVsphereStoragePolicyCompatiblePoolsResponse, error) {
	resp, err := s.cs.newRequest("listVsphereStoragePolicyCompatiblePools", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListVsphereStoragePolicyCompatiblePoolsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListVsphereStoragePolicyCompatiblePoolsResponse struct {
	Count                               int                                   `json:"count"`
	VsphereStoragePolicyCompatiblePools []*VsphereStoragePolicyCompatiblePool `json:"vspherestoragepolicycompatiblepool"`
}

type VsphereStoragePolicyCompatiblePool struct {
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
