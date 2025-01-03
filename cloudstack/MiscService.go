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

type MiscServiceIface interface {
	ListElastistorInterface(p *ListElastistorInterfaceParams) (*ListElastistorInterfaceResponse, error)
	NewListElastistorInterfaceParams() *ListElastistorInterfaceParams
}

type ListElastistorInterfaceParams struct {
	p map[string]interface{}
}

func (p *ListElastistorInterfaceParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["controllerid"]; found {
		u.Set("controllerid", v.(string))
	}
	return u
}

func (p *ListElastistorInterfaceParams) SetControllerid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["controllerid"] = v
}

func (p *ListElastistorInterfaceParams) ResetControllerid() {
	if p.p != nil && p.p["controllerid"] != nil {
		delete(p.p, "controllerid")
	}
}

func (p *ListElastistorInterfaceParams) GetControllerid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["controllerid"].(string)
	return value, ok
}

// You should always use this function to get a new ListElastistorInterfaceParams instance,
// as then you are sure you have configured all required params
func (s *MiscService) NewListElastistorInterfaceParams() *ListElastistorInterfaceParams {
	p := &ListElastistorInterfaceParams{}
	p.p = make(map[string]interface{})
	return p
}

// Lists the network Interfaces of elastistor
func (s *MiscService) ListElastistorInterface(p *ListElastistorInterfaceParams) (*ListElastistorInterfaceResponse, error) {
	resp, err := s.cs.newRequest("listElastistorInterface", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListElastistorInterfaceResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListElastistorInterfaceResponse struct {
	Count               int                    `json:"count"`
	ElastistorInterface []*ElastistorInterface `json:"elastistorinterface"`
}

type ElastistorInterface struct {
	Compression   string `json:"compression"`
	Deduplication string `json:"deduplication"`
	Graceallowed  string `json:"graceallowed"`
	Id            string `json:"id"`
	JobID         string `json:"jobid"`
	Jobstatus     int    `json:"jobstatus"`
	Name          string `json:"name"`
	Sync          string `json:"sync"`
}
