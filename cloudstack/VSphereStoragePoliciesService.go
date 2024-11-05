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
	"net/url"
)

type VSphereStoragePoliciesServiceIface interface {
	ImportVsphereStoragePolicies(p *ImportVsphereStoragePoliciesParams) (*ImportVsphereStoragePoliciesResponse, error)
	NewImportVsphereStoragePoliciesParams() *ImportVsphereStoragePoliciesParams
	ListVsphereStoragePolicies(p *ListVsphereStoragePoliciesParams) (*ListVsphereStoragePoliciesResponse, error)
	NewListVsphereStoragePoliciesParams() *ListVsphereStoragePoliciesParams
}

type ImportVsphereStoragePoliciesParams struct {
	p map[string]interface{}
}

func (p *ImportVsphereStoragePoliciesParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *ImportVsphereStoragePoliciesParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

func (p *ImportVsphereStoragePoliciesParams) ResetZoneid() {
	if p.p != nil && p.p["zoneid"] != nil {
		delete(p.p, "zoneid")
	}
}

func (p *ImportVsphereStoragePoliciesParams) GetZoneid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new ImportVsphereStoragePoliciesParams instance,
// as then you are sure you have configured all required params
func (s *VSphereStoragePoliciesService) NewImportVsphereStoragePoliciesParams() *ImportVsphereStoragePoliciesParams {
	p := &ImportVsphereStoragePoliciesParams{}
	p.p = make(map[string]interface{})
	return p
}

// Import vSphere storage policies
func (s *VSphereStoragePoliciesService) ImportVsphereStoragePolicies(p *ImportVsphereStoragePoliciesParams) (*ImportVsphereStoragePoliciesResponse, error) {
	resp, err := s.cs.newRequest("importVsphereStoragePolicies", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ImportVsphereStoragePoliciesResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ImportVsphereStoragePoliciesResponse struct {
	Description string `json:"description"`
	Id          string `json:"id"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Name        string `json:"name"`
	Policyid    string `json:"policyid"`
	Zoneid      string `json:"zoneid"`
}

type ListVsphereStoragePoliciesParams struct {
	p map[string]interface{}
}

func (p *ListVsphereStoragePoliciesParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *ListVsphereStoragePoliciesParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

func (p *ListVsphereStoragePoliciesParams) ResetZoneid() {
	if p.p != nil && p.p["zoneid"] != nil {
		delete(p.p, "zoneid")
	}
}

func (p *ListVsphereStoragePoliciesParams) GetZoneid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new ListVsphereStoragePoliciesParams instance,
// as then you are sure you have configured all required params
func (s *VSphereStoragePoliciesService) NewListVsphereStoragePoliciesParams() *ListVsphereStoragePoliciesParams {
	p := &ListVsphereStoragePoliciesParams{}
	p.p = make(map[string]interface{})
	return p
}

// List vSphere storage policies
func (s *VSphereStoragePoliciesService) ListVsphereStoragePolicies(p *ListVsphereStoragePoliciesParams) (*ListVsphereStoragePoliciesResponse, error) {
	resp, err := s.cs.newRequest("listVsphereStoragePolicies", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListVsphereStoragePoliciesResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListVsphereStoragePoliciesResponse struct {
	Count                  int                     `json:"count"`
	VsphereStoragePolicies []*VsphereStoragePolicy `json:"vspherestoragepolicy"`
}

type VsphereStoragePolicy struct {
	Description string `json:"description"`
	Id          string `json:"id"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Name        string `json:"name"`
	Policyid    string `json:"policyid"`
	Zoneid      string `json:"zoneid"`
}
