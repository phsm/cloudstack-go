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

type ShutdownServiceIface interface {
	CancelShutdown(p *CancelShutdownParams) (*CancelShutdownResponse, error)
	NewCancelShutdownParams(managementserverid UUID) *CancelShutdownParams
	PrepareForShutdown(p *PrepareForShutdownParams) (*PrepareForShutdownResponse, error)
	NewPrepareForShutdownParams(managementserverid UUID) *PrepareForShutdownParams
	ReadyForShutdown(p *ReadyForShutdownParams) (*ReadyForShutdownResponse, error)
	NewReadyForShutdownParams() *ReadyForShutdownParams
	TriggerShutdown(p *TriggerShutdownParams) (*TriggerShutdownResponse, error)
	NewTriggerShutdownParams(managementserverid UUID) *TriggerShutdownParams
}

type CancelShutdownParams struct {
	p map[string]interface{}
}

func (p *CancelShutdownParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["managementserverid"]; found {
		u.Set("managementserverid", v.(string))
	}
	return u
}

func (p *CancelShutdownParams) SetManagementserverid(v UUID) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["managementserverid"] = v
}

func (p *CancelShutdownParams) ResetManagementserverid() {
	if p.p != nil && p.p["managementserverid"] != nil {
		delete(p.p, "managementserverid")
	}
}

func (p *CancelShutdownParams) GetManagementserverid() (UUID, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["managementserverid"].(UUID)
	return value, ok
}

// You should always use this function to get a new CancelShutdownParams instance,
// as then you are sure you have configured all required params
func (s *ShutdownService) NewCancelShutdownParams(managementserverid UUID) *CancelShutdownParams {
	p := &CancelShutdownParams{}
	p.p = make(map[string]interface{})
	p.p["managementserverid"] = managementserverid
	return p
}

// Cancels a triggered shutdown
func (s *ShutdownService) CancelShutdown(p *CancelShutdownParams) (*CancelShutdownResponse, error) {
	resp, err := s.cs.newRequest("cancelShutdown", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CancelShutdownResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type CancelShutdownResponse struct {
	JobID              string `json:"jobid"`
	Jobstatus          int    `json:"jobstatus"`
	Managementserverid UUID   `json:"managementserverid"`
	Pendingjobscount   int64  `json:"pendingjobscount"`
	Readyforshutdown   bool   `json:"readyforshutdown"`
	Shutdowntriggered  bool   `json:"shutdowntriggered"`
}

type PrepareForShutdownParams struct {
	p map[string]interface{}
}

func (p *PrepareForShutdownParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["managementserverid"]; found {
		u.Set("managementserverid", v.(string))
	}
	return u
}

func (p *PrepareForShutdownParams) SetManagementserverid(v UUID) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["managementserverid"] = v
}

func (p *PrepareForShutdownParams) ResetManagementserverid() {
	if p.p != nil && p.p["managementserverid"] != nil {
		delete(p.p, "managementserverid")
	}
}

func (p *PrepareForShutdownParams) GetManagementserverid() (UUID, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["managementserverid"].(UUID)
	return value, ok
}

// You should always use this function to get a new PrepareForShutdownParams instance,
// as then you are sure you have configured all required params
func (s *ShutdownService) NewPrepareForShutdownParams(managementserverid UUID) *PrepareForShutdownParams {
	p := &PrepareForShutdownParams{}
	p.p = make(map[string]interface{})
	p.p["managementserverid"] = managementserverid
	return p
}

// Prepares CloudStack for a safe manual shutdown by preventing new jobs from being accepted
func (s *ShutdownService) PrepareForShutdown(p *PrepareForShutdownParams) (*PrepareForShutdownResponse, error) {
	resp, err := s.cs.newRequest("prepareForShutdown", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r PrepareForShutdownResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type PrepareForShutdownResponse struct {
	JobID              string `json:"jobid"`
	Jobstatus          int    `json:"jobstatus"`
	Managementserverid UUID   `json:"managementserverid"`
	Pendingjobscount   int64  `json:"pendingjobscount"`
	Readyforshutdown   bool   `json:"readyforshutdown"`
	Shutdowntriggered  bool   `json:"shutdowntriggered"`
}

type ReadyForShutdownParams struct {
	p map[string]interface{}
}

func (p *ReadyForShutdownParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["managementserverid"]; found {
		u.Set("managementserverid", v.(string))
	}
	return u
}

func (p *ReadyForShutdownParams) SetManagementserverid(v UUID) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["managementserverid"] = v
}

func (p *ReadyForShutdownParams) ResetManagementserverid() {
	if p.p != nil && p.p["managementserverid"] != nil {
		delete(p.p, "managementserverid")
	}
}

func (p *ReadyForShutdownParams) GetManagementserverid() (UUID, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["managementserverid"].(UUID)
	return value, ok
}

// You should always use this function to get a new ReadyForShutdownParams instance,
// as then you are sure you have configured all required params
func (s *ShutdownService) NewReadyForShutdownParams() *ReadyForShutdownParams {
	p := &ReadyForShutdownParams{}
	p.p = make(map[string]interface{})
	return p
}

// Returns the status of CloudStack, whether a shutdown has been triggered and if ready to shutdown
func (s *ShutdownService) ReadyForShutdown(p *ReadyForShutdownParams) (*ReadyForShutdownResponse, error) {
	resp, err := s.cs.newRequest("readyForShutdown", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ReadyForShutdownResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ReadyForShutdownResponse struct {
	JobID              string `json:"jobid"`
	Jobstatus          int    `json:"jobstatus"`
	Managementserverid UUID   `json:"managementserverid"`
	Pendingjobscount   int64  `json:"pendingjobscount"`
	Readyforshutdown   bool   `json:"readyforshutdown"`
	Shutdowntriggered  bool   `json:"shutdowntriggered"`
}

type TriggerShutdownParams struct {
	p map[string]interface{}
}

func (p *TriggerShutdownParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["managementserverid"]; found {
		u.Set("managementserverid", v.(string))
	}
	return u
}

func (p *TriggerShutdownParams) SetManagementserverid(v UUID) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["managementserverid"] = v
}

func (p *TriggerShutdownParams) ResetManagementserverid() {
	if p.p != nil && p.p["managementserverid"] != nil {
		delete(p.p, "managementserverid")
	}
}

func (p *TriggerShutdownParams) GetManagementserverid() (UUID, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["managementserverid"].(UUID)
	return value, ok
}

// You should always use this function to get a new TriggerShutdownParams instance,
// as then you are sure you have configured all required params
func (s *ShutdownService) NewTriggerShutdownParams(managementserverid UUID) *TriggerShutdownParams {
	p := &TriggerShutdownParams{}
	p.p = make(map[string]interface{})
	p.p["managementserverid"] = managementserverid
	return p
}

// Triggers an automatic safe shutdown of CloudStack by not accepting new jobs and shutting down when all pending jobbs have been completed. Triggers an immediate shutdown if forced
func (s *ShutdownService) TriggerShutdown(p *TriggerShutdownParams) (*TriggerShutdownResponse, error) {
	resp, err := s.cs.newRequest("triggerShutdown", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r TriggerShutdownResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type TriggerShutdownResponse struct {
	JobID              string `json:"jobid"`
	Jobstatus          int    `json:"jobstatus"`
	Managementserverid UUID   `json:"managementserverid"`
	Pendingjobscount   int64  `json:"pendingjobscount"`
	Readyforshutdown   bool   `json:"readyforshutdown"`
	Shutdowntriggered  bool   `json:"shutdowntriggered"`
}
