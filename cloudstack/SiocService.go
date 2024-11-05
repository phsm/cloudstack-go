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
	"strconv"
)

type SiocServiceIface interface {
	UpdateSiocInfo(p *UpdateSiocInfoParams) (*UpdateSiocInfoResponse, error)
	NewUpdateSiocInfoParams(iopsnotifythreshold int, limitiopspergb int, sharespergb int, storageid string, zoneid string) *UpdateSiocInfoParams
}

type UpdateSiocInfoParams struct {
	p map[string]interface{}
}

func (p *UpdateSiocInfoParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["iopsnotifythreshold"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("iopsnotifythreshold", vv)
	}
	if v, found := p.p["limitiopspergb"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("limitiopspergb", vv)
	}
	if v, found := p.p["sharespergb"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("sharespergb", vv)
	}
	if v, found := p.p["storageid"]; found {
		u.Set("storageid", v.(string))
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *UpdateSiocInfoParams) SetIopsnotifythreshold(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["iopsnotifythreshold"] = v
}

func (p *UpdateSiocInfoParams) ResetIopsnotifythreshold() {
	if p.p != nil && p.p["iopsnotifythreshold"] != nil {
		delete(p.p, "iopsnotifythreshold")
	}
}

func (p *UpdateSiocInfoParams) GetIopsnotifythreshold() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["iopsnotifythreshold"].(int)
	return value, ok
}

func (p *UpdateSiocInfoParams) SetLimitiopspergb(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["limitiopspergb"] = v
}

func (p *UpdateSiocInfoParams) ResetLimitiopspergb() {
	if p.p != nil && p.p["limitiopspergb"] != nil {
		delete(p.p, "limitiopspergb")
	}
}

func (p *UpdateSiocInfoParams) GetLimitiopspergb() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["limitiopspergb"].(int)
	return value, ok
}

func (p *UpdateSiocInfoParams) SetSharespergb(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["sharespergb"] = v
}

func (p *UpdateSiocInfoParams) ResetSharespergb() {
	if p.p != nil && p.p["sharespergb"] != nil {
		delete(p.p, "sharespergb")
	}
}

func (p *UpdateSiocInfoParams) GetSharespergb() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["sharespergb"].(int)
	return value, ok
}

func (p *UpdateSiocInfoParams) SetStorageid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["storageid"] = v
}

func (p *UpdateSiocInfoParams) ResetStorageid() {
	if p.p != nil && p.p["storageid"] != nil {
		delete(p.p, "storageid")
	}
}

func (p *UpdateSiocInfoParams) GetStorageid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["storageid"].(string)
	return value, ok
}

func (p *UpdateSiocInfoParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

func (p *UpdateSiocInfoParams) ResetZoneid() {
	if p.p != nil && p.p["zoneid"] != nil {
		delete(p.p, "zoneid")
	}
}

func (p *UpdateSiocInfoParams) GetZoneid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new UpdateSiocInfoParams instance,
// as then you are sure you have configured all required params
func (s *SiocService) NewUpdateSiocInfoParams(iopsnotifythreshold int, limitiopspergb int, sharespergb int, storageid string, zoneid string) *UpdateSiocInfoParams {
	p := &UpdateSiocInfoParams{}
	p.p = make(map[string]interface{})
	p.p["iopsnotifythreshold"] = iopsnotifythreshold
	p.p["limitiopspergb"] = limitiopspergb
	p.p["sharespergb"] = sharespergb
	p.p["storageid"] = storageid
	p.p["zoneid"] = zoneid
	return p
}

// Update SIOC info
func (s *SiocService) UpdateSiocInfo(p *UpdateSiocInfoParams) (*UpdateSiocInfoResponse, error) {
	resp, err := s.cs.newRequest("updateSiocInfo", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateSiocInfoResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type UpdateSiocInfoResponse struct {
	JobID     string `json:"jobid"`
	Jobstatus int    `json:"jobstatus"`
	Msg       string `json:"msg"`
}
