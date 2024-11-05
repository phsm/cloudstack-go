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

type CertificateServiceIface interface {
	ListCaCertificate(p *ListCaCertificateParams) (*ListCaCertificateResponse, error)
	NewListCaCertificateParams() *ListCaCertificateParams
	IssueCertificate(p *IssueCertificateParams) (*IssueCertificateResponse, error)
	NewIssueCertificateParams() *IssueCertificateParams
	ListCAProviders(p *ListCAProvidersParams) (*ListCAProvidersResponse, error)
	NewListCAProvidersParams() *ListCAProvidersParams
	ProvisionCertificate(p *ProvisionCertificateParams) (*ProvisionCertificateResponse, error)
	NewProvisionCertificateParams(hostid string) *ProvisionCertificateParams
	RevokeCertificate(p *RevokeCertificateParams) (*RevokeCertificateResponse, error)
	NewRevokeCertificateParams(serial string) *RevokeCertificateParams
	UploadCustomCertificate(p *UploadCustomCertificateParams) (*UploadCustomCertificateResponse, error)
	NewUploadCustomCertificateParams(certificate string, domainsuffix string) *UploadCustomCertificateParams
}

type ListCaCertificateParams struct {
	p map[string]interface{}
}

func (p *ListCaCertificateParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["provider"]; found {
		u.Set("provider", v.(string))
	}
	return u
}

func (p *ListCaCertificateParams) SetProvider(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["provider"] = v
}

func (p *ListCaCertificateParams) ResetProvider() {
	if p.p != nil && p.p["provider"] != nil {
		delete(p.p, "provider")
	}
}

func (p *ListCaCertificateParams) GetProvider() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["provider"].(string)
	return value, ok
}

// You should always use this function to get a new ListCaCertificateParams instance,
// as then you are sure you have configured all required params
func (s *CertificateService) NewListCaCertificateParams() *ListCaCertificateParams {
	p := &ListCaCertificateParams{}
	p.p = make(map[string]interface{})
	return p
}

// Lists the CA public certificate(s) as support by the configured/provided CA plugin
func (s *CertificateService) ListCaCertificate(p *ListCaCertificateParams) (*ListCaCertificateResponse, error) {
	resp, err := s.cs.newRequest("listCaCertificate", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListCaCertificateResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListCaCertificateResponse struct {
	Count         int              `json:"count"`
	CaCertificate []*CaCertificate `json:"cacertificate"`
}

type CaCertificate struct {
	Cacertificates string `json:"cacertificates"`
	Certificate    string `json:"certificate"`
	JobID          string `json:"jobid"`
	Jobstatus      int    `json:"jobstatus"`
	Privatekey     string `json:"privatekey"`
}

type IssueCertificateParams struct {
	p map[string]interface{}
}

func (p *IssueCertificateParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["csr"]; found {
		u.Set("csr", v.(string))
	}
	if v, found := p.p["domain"]; found {
		u.Set("domain", v.(string))
	}
	if v, found := p.p["duration"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("duration", vv)
	}
	if v, found := p.p["ipaddress"]; found {
		u.Set("ipaddress", v.(string))
	}
	if v, found := p.p["provider"]; found {
		u.Set("provider", v.(string))
	}
	return u
}

func (p *IssueCertificateParams) SetCsr(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["csr"] = v
}

func (p *IssueCertificateParams) ResetCsr() {
	if p.p != nil && p.p["csr"] != nil {
		delete(p.p, "csr")
	}
}

func (p *IssueCertificateParams) GetCsr() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["csr"].(string)
	return value, ok
}

func (p *IssueCertificateParams) SetDomain(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domain"] = v
}

func (p *IssueCertificateParams) ResetDomain() {
	if p.p != nil && p.p["domain"] != nil {
		delete(p.p, "domain")
	}
}

func (p *IssueCertificateParams) GetDomain() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domain"].(string)
	return value, ok
}

func (p *IssueCertificateParams) SetDuration(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["duration"] = v
}

func (p *IssueCertificateParams) ResetDuration() {
	if p.p != nil && p.p["duration"] != nil {
		delete(p.p, "duration")
	}
}

func (p *IssueCertificateParams) GetDuration() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["duration"].(int)
	return value, ok
}

func (p *IssueCertificateParams) SetIpaddress(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ipaddress"] = v
}

func (p *IssueCertificateParams) ResetIpaddress() {
	if p.p != nil && p.p["ipaddress"] != nil {
		delete(p.p, "ipaddress")
	}
}

func (p *IssueCertificateParams) GetIpaddress() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["ipaddress"].(string)
	return value, ok
}

func (p *IssueCertificateParams) SetProvider(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["provider"] = v
}

func (p *IssueCertificateParams) ResetProvider() {
	if p.p != nil && p.p["provider"] != nil {
		delete(p.p, "provider")
	}
}

func (p *IssueCertificateParams) GetProvider() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["provider"].(string)
	return value, ok
}

// You should always use this function to get a new IssueCertificateParams instance,
// as then you are sure you have configured all required params
func (s *CertificateService) NewIssueCertificateParams() *IssueCertificateParams {
	p := &IssueCertificateParams{}
	p.p = make(map[string]interface{})
	return p
}

// Issues a client certificate using configured or provided CA plugin
func (s *CertificateService) IssueCertificate(p *IssueCertificateParams) (*IssueCertificateResponse, error) {
	resp, err := s.cs.newRequest("issueCertificate", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r IssueCertificateResponse
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

type IssueCertificateResponse struct {
	Cacertificates string `json:"cacertificates"`
	Certificate    string `json:"certificate"`
	JobID          string `json:"jobid"`
	Jobstatus      int    `json:"jobstatus"`
	Privatekey     string `json:"privatekey"`
}

type ListCAProvidersParams struct {
	p map[string]interface{}
}

func (p *ListCAProvidersParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	return u
}

func (p *ListCAProvidersParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *ListCAProvidersParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
}

func (p *ListCAProvidersParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

// You should always use this function to get a new ListCAProvidersParams instance,
// as then you are sure you have configured all required params
func (s *CertificateService) NewListCAProvidersParams() *ListCAProvidersParams {
	p := &ListCAProvidersParams{}
	p.p = make(map[string]interface{})
	return p
}

// Lists available certificate authority providers in CloudStack
func (s *CertificateService) ListCAProviders(p *ListCAProvidersParams) (*ListCAProvidersResponse, error) {
	resp, err := s.cs.newRequest("listCAProviders", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListCAProvidersResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListCAProvidersResponse struct {
	Count       int           `json:"count"`
	CAProviders []*CAProvider `json:"caprovider"`
}

type CAProvider struct {
	Description string `json:"description"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Name        string `json:"name"`
}

type ProvisionCertificateParams struct {
	p map[string]interface{}
}

func (p *ProvisionCertificateParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["hostid"]; found {
		u.Set("hostid", v.(string))
	}
	if v, found := p.p["provider"]; found {
		u.Set("provider", v.(string))
	}
	if v, found := p.p["reconnect"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("reconnect", vv)
	}
	return u
}

func (p *ProvisionCertificateParams) SetHostid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["hostid"] = v
}

func (p *ProvisionCertificateParams) ResetHostid() {
	if p.p != nil && p.p["hostid"] != nil {
		delete(p.p, "hostid")
	}
}

func (p *ProvisionCertificateParams) GetHostid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["hostid"].(string)
	return value, ok
}

func (p *ProvisionCertificateParams) SetProvider(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["provider"] = v
}

func (p *ProvisionCertificateParams) ResetProvider() {
	if p.p != nil && p.p["provider"] != nil {
		delete(p.p, "provider")
	}
}

func (p *ProvisionCertificateParams) GetProvider() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["provider"].(string)
	return value, ok
}

func (p *ProvisionCertificateParams) SetReconnect(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["reconnect"] = v
}

func (p *ProvisionCertificateParams) ResetReconnect() {
	if p.p != nil && p.p["reconnect"] != nil {
		delete(p.p, "reconnect")
	}
}

func (p *ProvisionCertificateParams) GetReconnect() (bool, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["reconnect"].(bool)
	return value, ok
}

// You should always use this function to get a new ProvisionCertificateParams instance,
// as then you are sure you have configured all required params
func (s *CertificateService) NewProvisionCertificateParams(hostid string) *ProvisionCertificateParams {
	p := &ProvisionCertificateParams{}
	p.p = make(map[string]interface{})
	p.p["hostid"] = hostid
	return p
}

// Issues and propagates client certificate on a connected host/agent using configured CA plugin
func (s *CertificateService) ProvisionCertificate(p *ProvisionCertificateParams) (*ProvisionCertificateResponse, error) {
	resp, err := s.cs.newRequest("provisionCertificate", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ProvisionCertificateResponse
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

type ProvisionCertificateResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type RevokeCertificateParams struct {
	p map[string]interface{}
}

func (p *RevokeCertificateParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["cn"]; found {
		u.Set("cn", v.(string))
	}
	if v, found := p.p["provider"]; found {
		u.Set("provider", v.(string))
	}
	if v, found := p.p["serial"]; found {
		u.Set("serial", v.(string))
	}
	return u
}

func (p *RevokeCertificateParams) SetCn(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["cn"] = v
}

func (p *RevokeCertificateParams) ResetCn() {
	if p.p != nil && p.p["cn"] != nil {
		delete(p.p, "cn")
	}
}

func (p *RevokeCertificateParams) GetCn() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["cn"].(string)
	return value, ok
}

func (p *RevokeCertificateParams) SetProvider(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["provider"] = v
}

func (p *RevokeCertificateParams) ResetProvider() {
	if p.p != nil && p.p["provider"] != nil {
		delete(p.p, "provider")
	}
}

func (p *RevokeCertificateParams) GetProvider() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["provider"].(string)
	return value, ok
}

func (p *RevokeCertificateParams) SetSerial(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["serial"] = v
}

func (p *RevokeCertificateParams) ResetSerial() {
	if p.p != nil && p.p["serial"] != nil {
		delete(p.p, "serial")
	}
}

func (p *RevokeCertificateParams) GetSerial() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["serial"].(string)
	return value, ok
}

// You should always use this function to get a new RevokeCertificateParams instance,
// as then you are sure you have configured all required params
func (s *CertificateService) NewRevokeCertificateParams(serial string) *RevokeCertificateParams {
	p := &RevokeCertificateParams{}
	p.p = make(map[string]interface{})
	p.p["serial"] = serial
	return p
}

// Revokes certificate using configured CA plugin
func (s *CertificateService) RevokeCertificate(p *RevokeCertificateParams) (*RevokeCertificateResponse, error) {
	resp, err := s.cs.newRequest("revokeCertificate", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r RevokeCertificateResponse
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

type RevokeCertificateResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type UploadCustomCertificateParams struct {
	p map[string]interface{}
}

func (p *UploadCustomCertificateParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["certificate"]; found {
		u.Set("certificate", v.(string))
	}
	if v, found := p.p["domainsuffix"]; found {
		u.Set("domainsuffix", v.(string))
	}
	if v, found := p.p["id"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("id", vv)
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["privatekey"]; found {
		u.Set("privatekey", v.(string))
	}
	return u
}

func (p *UploadCustomCertificateParams) SetCertificate(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["certificate"] = v
}

func (p *UploadCustomCertificateParams) ResetCertificate() {
	if p.p != nil && p.p["certificate"] != nil {
		delete(p.p, "certificate")
	}
}

func (p *UploadCustomCertificateParams) GetCertificate() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["certificate"].(string)
	return value, ok
}

func (p *UploadCustomCertificateParams) SetDomainsuffix(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainsuffix"] = v
}

func (p *UploadCustomCertificateParams) ResetDomainsuffix() {
	if p.p != nil && p.p["domainsuffix"] != nil {
		delete(p.p, "domainsuffix")
	}
}

func (p *UploadCustomCertificateParams) GetDomainsuffix() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["domainsuffix"].(string)
	return value, ok
}

func (p *UploadCustomCertificateParams) SetId(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *UploadCustomCertificateParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *UploadCustomCertificateParams) GetId() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(int)
	return value, ok
}

func (p *UploadCustomCertificateParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *UploadCustomCertificateParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
}

func (p *UploadCustomCertificateParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *UploadCustomCertificateParams) SetPrivatekey(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["privatekey"] = v
}

func (p *UploadCustomCertificateParams) ResetPrivatekey() {
	if p.p != nil && p.p["privatekey"] != nil {
		delete(p.p, "privatekey")
	}
}

func (p *UploadCustomCertificateParams) GetPrivatekey() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["privatekey"].(string)
	return value, ok
}

// You should always use this function to get a new UploadCustomCertificateParams instance,
// as then you are sure you have configured all required params
func (s *CertificateService) NewUploadCustomCertificateParams(certificate string, domainsuffix string) *UploadCustomCertificateParams {
	p := &UploadCustomCertificateParams{}
	p.p = make(map[string]interface{})
	p.p["certificate"] = certificate
	p.p["domainsuffix"] = domainsuffix
	return p
}

// Uploads a custom certificate for the console proxy VMs to use for SSL. Can be used to upload a single certificate signed by a known CA. Can also be used, through multiple calls, to upload a chain of certificates from CA to the custom certificate itself.
func (s *CertificateService) UploadCustomCertificate(p *UploadCustomCertificateParams) (*UploadCustomCertificateResponse, error) {
	resp, err := s.cs.newRequest("uploadCustomCertificate", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UploadCustomCertificateResponse
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

type UploadCustomCertificateResponse struct {
	JobID     string `json:"jobid"`
	Jobstatus int    `json:"jobstatus"`
	Message   string `json:"message"`
}
