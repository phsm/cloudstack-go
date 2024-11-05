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
	"strings"
)

type TungstenServiceIface interface {
	AddTungstenFabricPolicyRule(p *AddTungstenFabricPolicyRuleParams) (*AddTungstenFabricPolicyRuleResponse, error)
	NewAddTungstenFabricPolicyRuleParams(action string, destendport int, destipprefix string, destipprefixlen int, destnetwork string, deststartport int, direction string, policyuuid string, protocol string, srcendport int, srcipprefix string, srcipprefixlen int, srcnetwork string, srcstartport int, zoneid string) *AddTungstenFabricPolicyRuleParams
	ApplyTungstenFabricPolicy(p *ApplyTungstenFabricPolicyParams) (*ApplyTungstenFabricPolicyResponse, error)
	NewApplyTungstenFabricPolicyParams(majorsequence int, minorsequence int, networkuuid string, policyuuid string, zoneid string) *ApplyTungstenFabricPolicyParams
	ApplyTungstenFabricTag(p *ApplyTungstenFabricTagParams) (*ApplyTungstenFabricTagResponse, error)
	NewApplyTungstenFabricTagParams(taguuid string, zoneid string) *ApplyTungstenFabricTagParams
	CreateTungstenFabricAddressGroup(p *CreateTungstenFabricAddressGroupParams) (*CreateTungstenFabricAddressGroupResponse, error)
	NewCreateTungstenFabricAddressGroupParams(ipprefix string, ipprefixlen int, name string, zoneid string) *CreateTungstenFabricAddressGroupParams
	CreateTungstenFabricApplicationPolicySet(p *CreateTungstenFabricApplicationPolicySetParams) (*CreateTungstenFabricApplicationPolicySetResponse, error)
	NewCreateTungstenFabricApplicationPolicySetParams(name string, zoneid string) *CreateTungstenFabricApplicationPolicySetParams
	CreateTungstenFabricServiceGroup(p *CreateTungstenFabricServiceGroupParams) (*CreateTungstenFabricServiceGroupResponse, error)
	NewCreateTungstenFabricServiceGroupParams(endport int, name string, protocol string, startport int, zoneid string) *CreateTungstenFabricServiceGroupParams
	CreateTungstenFabricPolicy(p *CreateTungstenFabricPolicyParams) (*CreateTungstenFabricPolicyResponse, error)
	NewCreateTungstenFabricPolicyParams(name string, zoneid string) *CreateTungstenFabricPolicyParams
	CreateTungstenFabricProvider(p *CreateTungstenFabricProviderParams) (*CreateTungstenFabricProviderResponse, error)
	NewCreateTungstenFabricProviderParams(name string, tungstengateway string, tungstenproviderhostname string, zoneid string) *CreateTungstenFabricProviderParams
	CreateTungstenFabricTag(p *CreateTungstenFabricTagParams) (*CreateTungstenFabricTagResponse, error)
	NewCreateTungstenFabricTagParams(tagtype string, tagvalue string, zoneid string) *CreateTungstenFabricTagParams
	CreateTungstenFabricTagType(p *CreateTungstenFabricTagTypeParams) (*CreateTungstenFabricTagTypeResponse, error)
	NewCreateTungstenFabricTagTypeParams(name string, zoneid string) *CreateTungstenFabricTagTypeParams
	ConfigTungstenFabricService(p *ConfigTungstenFabricServiceParams) (*ConfigTungstenFabricServiceResponse, error)
	NewConfigTungstenFabricServiceParams(physicalnetworkid string, zoneid string) *ConfigTungstenFabricServiceParams
	DeleteTungstenFabricAddressGroup(p *DeleteTungstenFabricAddressGroupParams) (*DeleteTungstenFabricAddressGroupResponse, error)
	NewDeleteTungstenFabricAddressGroupParams(addressgroupuuid string, zoneid string) *DeleteTungstenFabricAddressGroupParams
	DeleteTungstenFabricApplicationPolicySet(p *DeleteTungstenFabricApplicationPolicySetParams) (*DeleteTungstenFabricApplicationPolicySetResponse, error)
	NewDeleteTungstenFabricApplicationPolicySetParams(applicationpolicysetuuid string, zoneid string) *DeleteTungstenFabricApplicationPolicySetParams
	DeleteTungstenFabricPolicy(p *DeleteTungstenFabricPolicyParams) (*DeleteTungstenFabricPolicyResponse, error)
	NewDeleteTungstenFabricPolicyParams(policyuuid string, zoneid string) *DeleteTungstenFabricPolicyParams
	DeleteTungstenFabricTag(p *DeleteTungstenFabricTagParams) (*DeleteTungstenFabricTagResponse, error)
	NewDeleteTungstenFabricTagParams(taguuid string, zoneid string) *DeleteTungstenFabricTagParams
	DeleteTungstenFabricTagType(p *DeleteTungstenFabricTagTypeParams) (*DeleteTungstenFabricTagTypeResponse, error)
	NewDeleteTungstenFabricTagTypeParams(tagtypeuuid string, zoneid string) *DeleteTungstenFabricTagTypeParams
	DeleteTungstenFabricServiceGroup(p *DeleteTungstenFabricServiceGroupParams) (*DeleteTungstenFabricServiceGroupResponse, error)
	NewDeleteTungstenFabricServiceGroupParams(servicegroupuuid string, zoneid string) *DeleteTungstenFabricServiceGroupParams
	ListTungstenFabricAddressGroup(p *ListTungstenFabricAddressGroupParams) (*ListTungstenFabricAddressGroupResponse, error)
	NewListTungstenFabricAddressGroupParams() *ListTungstenFabricAddressGroupParams
	ListTungstenFabricApplicationPolicySet(p *ListTungstenFabricApplicationPolicySetParams) (*ListTungstenFabricApplicationPolicySetResponse, error)
	NewListTungstenFabricApplicationPolicySetParams() *ListTungstenFabricApplicationPolicySetParams
	ListTungstenFabricLBHealthMonitor(p *ListTungstenFabricLBHealthMonitorParams) (*ListTungstenFabricLBHealthMonitorResponse, error)
	NewListTungstenFabricLBHealthMonitorParams(lbruleid string) *ListTungstenFabricLBHealthMonitorParams
	ListTungstenFabricNic(p *ListTungstenFabricNicParams) (*ListTungstenFabricNicResponse, error)
	NewListTungstenFabricNicParams() *ListTungstenFabricNicParams
	ListTungstenFabricPolicyRule(p *ListTungstenFabricPolicyRuleParams) (*ListTungstenFabricPolicyRuleResponse, error)
	NewListTungstenFabricPolicyRuleParams(policyuuid string) *ListTungstenFabricPolicyRuleParams
	ListTungstenFabricPolicy(p *ListTungstenFabricPolicyParams) (*ListTungstenFabricPolicyResponse, error)
	NewListTungstenFabricPolicyParams() *ListTungstenFabricPolicyParams
	ListTungstenFabricProviders(p *ListTungstenFabricProvidersParams) (*ListTungstenFabricProvidersResponse, error)
	NewListTungstenFabricProvidersParams() *ListTungstenFabricProvidersParams
	ListTungstenFabricServiceGroup(p *ListTungstenFabricServiceGroupParams) (*ListTungstenFabricServiceGroupResponse, error)
	NewListTungstenFabricServiceGroupParams() *ListTungstenFabricServiceGroupParams
	ListTungstenFabricTag(p *ListTungstenFabricTagParams) (*ListTungstenFabricTagResponse, error)
	NewListTungstenFabricTagParams() *ListTungstenFabricTagParams
	ListTungstenFabricTagType(p *ListTungstenFabricTagTypeParams) (*ListTungstenFabricTagTypeResponse, error)
	NewListTungstenFabricTagTypeParams() *ListTungstenFabricTagTypeParams
	ListTungstenFabricVm(p *ListTungstenFabricVmParams) (*ListTungstenFabricVmResponse, error)
	NewListTungstenFabricVmParams() *ListTungstenFabricVmParams
	RemoveTungstenFabricPolicy(p *RemoveTungstenFabricPolicyParams) (*RemoveTungstenFabricPolicyResponse, error)
	NewRemoveTungstenFabricPolicyParams(networkuuid string, policyuuid string, zoneid string) *RemoveTungstenFabricPolicyParams
	RemoveTungstenFabricPolicyRule(p *RemoveTungstenFabricPolicyRuleParams) (*RemoveTungstenFabricPolicyRuleResponse, error)
	NewRemoveTungstenFabricPolicyRuleParams(policyuuid string, ruleuuid string, zoneid string) *RemoveTungstenFabricPolicyRuleParams
	RemoveTungstenFabricTag(p *RemoveTungstenFabricTagParams) (*RemoveTungstenFabricTagResponse, error)
	NewRemoveTungstenFabricTagParams(taguuid string, zoneid string) *RemoveTungstenFabricTagParams
	SynchronizeTungstenFabricData(p *SynchronizeTungstenFabricDataParams) (*SynchronizeTungstenFabricDataResponse, error)
	NewSynchronizeTungstenFabricDataParams(id string) *SynchronizeTungstenFabricDataParams
	UpdateTungstenFabricLBHealthMonitor(p *UpdateTungstenFabricLBHealthMonitorParams) (*UpdateTungstenFabricLBHealthMonitorResponse, error)
	NewUpdateTungstenFabricLBHealthMonitorParams(interval int, lbruleid string, retry int, timeout int, tungstenType string) *UpdateTungstenFabricLBHealthMonitorParams
}

type AddTungstenFabricPolicyRuleParams struct {
	p map[string]interface{}
}

func (p *AddTungstenFabricPolicyRuleParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["action"]; found {
		u.Set("action", v.(string))
	}
	if v, found := p.p["destendport"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("destendport", vv)
	}
	if v, found := p.p["destipprefix"]; found {
		u.Set("destipprefix", v.(string))
	}
	if v, found := p.p["destipprefixlen"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("destipprefixlen", vv)
	}
	if v, found := p.p["destnetwork"]; found {
		u.Set("destnetwork", v.(string))
	}
	if v, found := p.p["deststartport"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("deststartport", vv)
	}
	if v, found := p.p["direction"]; found {
		u.Set("direction", v.(string))
	}
	if v, found := p.p["policyuuid"]; found {
		u.Set("policyuuid", v.(string))
	}
	if v, found := p.p["protocol"]; found {
		u.Set("protocol", v.(string))
	}
	if v, found := p.p["srcendport"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("srcendport", vv)
	}
	if v, found := p.p["srcipprefix"]; found {
		u.Set("srcipprefix", v.(string))
	}
	if v, found := p.p["srcipprefixlen"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("srcipprefixlen", vv)
	}
	if v, found := p.p["srcnetwork"]; found {
		u.Set("srcnetwork", v.(string))
	}
	if v, found := p.p["srcstartport"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("srcstartport", vv)
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *AddTungstenFabricPolicyRuleParams) SetAction(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["action"] = v
}

func (p *AddTungstenFabricPolicyRuleParams) ResetAction() {
	if p.p != nil && p.p["action"] != nil {
		delete(p.p, "action")
	}
}

func (p *AddTungstenFabricPolicyRuleParams) GetAction() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["action"].(string)
	return value, ok
}

func (p *AddTungstenFabricPolicyRuleParams) SetDestendport(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["destendport"] = v
}

func (p *AddTungstenFabricPolicyRuleParams) ResetDestendport() {
	if p.p != nil && p.p["destendport"] != nil {
		delete(p.p, "destendport")
	}
}

func (p *AddTungstenFabricPolicyRuleParams) GetDestendport() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["destendport"].(int)
	return value, ok
}

func (p *AddTungstenFabricPolicyRuleParams) SetDestipprefix(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["destipprefix"] = v
}

func (p *AddTungstenFabricPolicyRuleParams) ResetDestipprefix() {
	if p.p != nil && p.p["destipprefix"] != nil {
		delete(p.p, "destipprefix")
	}
}

func (p *AddTungstenFabricPolicyRuleParams) GetDestipprefix() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["destipprefix"].(string)
	return value, ok
}

func (p *AddTungstenFabricPolicyRuleParams) SetDestipprefixlen(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["destipprefixlen"] = v
}

func (p *AddTungstenFabricPolicyRuleParams) ResetDestipprefixlen() {
	if p.p != nil && p.p["destipprefixlen"] != nil {
		delete(p.p, "destipprefixlen")
	}
}

func (p *AddTungstenFabricPolicyRuleParams) GetDestipprefixlen() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["destipprefixlen"].(int)
	return value, ok
}

func (p *AddTungstenFabricPolicyRuleParams) SetDestnetwork(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["destnetwork"] = v
}

func (p *AddTungstenFabricPolicyRuleParams) ResetDestnetwork() {
	if p.p != nil && p.p["destnetwork"] != nil {
		delete(p.p, "destnetwork")
	}
}

func (p *AddTungstenFabricPolicyRuleParams) GetDestnetwork() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["destnetwork"].(string)
	return value, ok
}

func (p *AddTungstenFabricPolicyRuleParams) SetDeststartport(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["deststartport"] = v
}

func (p *AddTungstenFabricPolicyRuleParams) ResetDeststartport() {
	if p.p != nil && p.p["deststartport"] != nil {
		delete(p.p, "deststartport")
	}
}

func (p *AddTungstenFabricPolicyRuleParams) GetDeststartport() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["deststartport"].(int)
	return value, ok
}

func (p *AddTungstenFabricPolicyRuleParams) SetDirection(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["direction"] = v
}

func (p *AddTungstenFabricPolicyRuleParams) ResetDirection() {
	if p.p != nil && p.p["direction"] != nil {
		delete(p.p, "direction")
	}
}

func (p *AddTungstenFabricPolicyRuleParams) GetDirection() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["direction"].(string)
	return value, ok
}

func (p *AddTungstenFabricPolicyRuleParams) SetPolicyuuid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["policyuuid"] = v
}

func (p *AddTungstenFabricPolicyRuleParams) ResetPolicyuuid() {
	if p.p != nil && p.p["policyuuid"] != nil {
		delete(p.p, "policyuuid")
	}
}

func (p *AddTungstenFabricPolicyRuleParams) GetPolicyuuid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["policyuuid"].(string)
	return value, ok
}

func (p *AddTungstenFabricPolicyRuleParams) SetProtocol(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["protocol"] = v
}

func (p *AddTungstenFabricPolicyRuleParams) ResetProtocol() {
	if p.p != nil && p.p["protocol"] != nil {
		delete(p.p, "protocol")
	}
}

func (p *AddTungstenFabricPolicyRuleParams) GetProtocol() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["protocol"].(string)
	return value, ok
}

func (p *AddTungstenFabricPolicyRuleParams) SetSrcendport(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["srcendport"] = v
}

func (p *AddTungstenFabricPolicyRuleParams) ResetSrcendport() {
	if p.p != nil && p.p["srcendport"] != nil {
		delete(p.p, "srcendport")
	}
}

func (p *AddTungstenFabricPolicyRuleParams) GetSrcendport() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["srcendport"].(int)
	return value, ok
}

func (p *AddTungstenFabricPolicyRuleParams) SetSrcipprefix(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["srcipprefix"] = v
}

func (p *AddTungstenFabricPolicyRuleParams) ResetSrcipprefix() {
	if p.p != nil && p.p["srcipprefix"] != nil {
		delete(p.p, "srcipprefix")
	}
}

func (p *AddTungstenFabricPolicyRuleParams) GetSrcipprefix() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["srcipprefix"].(string)
	return value, ok
}

func (p *AddTungstenFabricPolicyRuleParams) SetSrcipprefixlen(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["srcipprefixlen"] = v
}

func (p *AddTungstenFabricPolicyRuleParams) ResetSrcipprefixlen() {
	if p.p != nil && p.p["srcipprefixlen"] != nil {
		delete(p.p, "srcipprefixlen")
	}
}

func (p *AddTungstenFabricPolicyRuleParams) GetSrcipprefixlen() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["srcipprefixlen"].(int)
	return value, ok
}

func (p *AddTungstenFabricPolicyRuleParams) SetSrcnetwork(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["srcnetwork"] = v
}

func (p *AddTungstenFabricPolicyRuleParams) ResetSrcnetwork() {
	if p.p != nil && p.p["srcnetwork"] != nil {
		delete(p.p, "srcnetwork")
	}
}

func (p *AddTungstenFabricPolicyRuleParams) GetSrcnetwork() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["srcnetwork"].(string)
	return value, ok
}

func (p *AddTungstenFabricPolicyRuleParams) SetSrcstartport(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["srcstartport"] = v
}

func (p *AddTungstenFabricPolicyRuleParams) ResetSrcstartport() {
	if p.p != nil && p.p["srcstartport"] != nil {
		delete(p.p, "srcstartport")
	}
}

func (p *AddTungstenFabricPolicyRuleParams) GetSrcstartport() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["srcstartport"].(int)
	return value, ok
}

func (p *AddTungstenFabricPolicyRuleParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

func (p *AddTungstenFabricPolicyRuleParams) ResetZoneid() {
	if p.p != nil && p.p["zoneid"] != nil {
		delete(p.p, "zoneid")
	}
}

func (p *AddTungstenFabricPolicyRuleParams) GetZoneid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new AddTungstenFabricPolicyRuleParams instance,
// as then you are sure you have configured all required params
func (s *TungstenService) NewAddTungstenFabricPolicyRuleParams(action string, destendport int, destipprefix string, destipprefixlen int, destnetwork string, deststartport int, direction string, policyuuid string, protocol string, srcendport int, srcipprefix string, srcipprefixlen int, srcnetwork string, srcstartport int, zoneid string) *AddTungstenFabricPolicyRuleParams {
	p := &AddTungstenFabricPolicyRuleParams{}
	p.p = make(map[string]interface{})
	p.p["action"] = action
	p.p["destendport"] = destendport
	p.p["destipprefix"] = destipprefix
	p.p["destipprefixlen"] = destipprefixlen
	p.p["destnetwork"] = destnetwork
	p.p["deststartport"] = deststartport
	p.p["direction"] = direction
	p.p["policyuuid"] = policyuuid
	p.p["protocol"] = protocol
	p.p["srcendport"] = srcendport
	p.p["srcipprefix"] = srcipprefix
	p.p["srcipprefixlen"] = srcipprefixlen
	p.p["srcnetwork"] = srcnetwork
	p.p["srcstartport"] = srcstartport
	p.p["zoneid"] = zoneid
	return p
}

// add Tungsten-Fabric policy rule
func (s *TungstenService) AddTungstenFabricPolicyRule(p *AddTungstenFabricPolicyRuleParams) (*AddTungstenFabricPolicyRuleResponse, error) {
	resp, err := s.cs.newRequest("addTungstenFabricPolicyRule", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r AddTungstenFabricPolicyRuleResponse
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

type AddTungstenFabricPolicyRuleResponse struct {
	Action          string `json:"action"`
	Destendport     int    `json:"destendport"`
	Destipprefix    string `json:"destipprefix"`
	Destipprefixlen int    `json:"destipprefixlen"`
	Destnetwork     string `json:"destnetwork"`
	Deststartport   int    `json:"deststartport"`
	Direction       string `json:"direction"`
	JobID           string `json:"jobid"`
	Jobstatus       int    `json:"jobstatus"`
	Policyuuid      string `json:"policyuuid"`
	Protocol        string `json:"protocol"`
	Srcendport      int    `json:"srcendport"`
	Srcipprefix     string `json:"srcipprefix"`
	Srcipprefixlen  int    `json:"srcipprefixlen"`
	Srcnetwork      string `json:"srcnetwork"`
	Srcstartport    int    `json:"srcstartport"`
	Uuid            string `json:"uuid"`
	Zoneid          int64  `json:"zoneid"`
	Zonename        string `json:"zonename"`
}

type ApplyTungstenFabricPolicyParams struct {
	p map[string]interface{}
}

func (p *ApplyTungstenFabricPolicyParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["majorsequence"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("majorsequence", vv)
	}
	if v, found := p.p["minorsequence"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("minorsequence", vv)
	}
	if v, found := p.p["networkuuid"]; found {
		u.Set("networkuuid", v.(string))
	}
	if v, found := p.p["policyuuid"]; found {
		u.Set("policyuuid", v.(string))
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *ApplyTungstenFabricPolicyParams) SetMajorsequence(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["majorsequence"] = v
}

func (p *ApplyTungstenFabricPolicyParams) ResetMajorsequence() {
	if p.p != nil && p.p["majorsequence"] != nil {
		delete(p.p, "majorsequence")
	}
}

func (p *ApplyTungstenFabricPolicyParams) GetMajorsequence() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["majorsequence"].(int)
	return value, ok
}

func (p *ApplyTungstenFabricPolicyParams) SetMinorsequence(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["minorsequence"] = v
}

func (p *ApplyTungstenFabricPolicyParams) ResetMinorsequence() {
	if p.p != nil && p.p["minorsequence"] != nil {
		delete(p.p, "minorsequence")
	}
}

func (p *ApplyTungstenFabricPolicyParams) GetMinorsequence() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["minorsequence"].(int)
	return value, ok
}

func (p *ApplyTungstenFabricPolicyParams) SetNetworkuuid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["networkuuid"] = v
}

func (p *ApplyTungstenFabricPolicyParams) ResetNetworkuuid() {
	if p.p != nil && p.p["networkuuid"] != nil {
		delete(p.p, "networkuuid")
	}
}

func (p *ApplyTungstenFabricPolicyParams) GetNetworkuuid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["networkuuid"].(string)
	return value, ok
}

func (p *ApplyTungstenFabricPolicyParams) SetPolicyuuid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["policyuuid"] = v
}

func (p *ApplyTungstenFabricPolicyParams) ResetPolicyuuid() {
	if p.p != nil && p.p["policyuuid"] != nil {
		delete(p.p, "policyuuid")
	}
}

func (p *ApplyTungstenFabricPolicyParams) GetPolicyuuid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["policyuuid"].(string)
	return value, ok
}

func (p *ApplyTungstenFabricPolicyParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

func (p *ApplyTungstenFabricPolicyParams) ResetZoneid() {
	if p.p != nil && p.p["zoneid"] != nil {
		delete(p.p, "zoneid")
	}
}

func (p *ApplyTungstenFabricPolicyParams) GetZoneid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new ApplyTungstenFabricPolicyParams instance,
// as then you are sure you have configured all required params
func (s *TungstenService) NewApplyTungstenFabricPolicyParams(majorsequence int, minorsequence int, networkuuid string, policyuuid string, zoneid string) *ApplyTungstenFabricPolicyParams {
	p := &ApplyTungstenFabricPolicyParams{}
	p.p = make(map[string]interface{})
	p.p["majorsequence"] = majorsequence
	p.p["minorsequence"] = minorsequence
	p.p["networkuuid"] = networkuuid
	p.p["policyuuid"] = policyuuid
	p.p["zoneid"] = zoneid
	return p
}

// apply Tungsten-Fabric policy
func (s *TungstenService) ApplyTungstenFabricPolicy(p *ApplyTungstenFabricPolicyParams) (*ApplyTungstenFabricPolicyResponse, error) {
	resp, err := s.cs.newRequest("applyTungstenFabricPolicy", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ApplyTungstenFabricPolicyResponse
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

type ApplyTungstenFabricPolicyResponse struct {
	JobID     string     `json:"jobid"`
	Jobstatus int        `json:"jobstatus"`
	Name      string     `json:"name"`
	Network   []*Network `json:"network"`
	Uuid      string     `json:"uuid"`
	Zoneid    int64      `json:"zoneid"`
	Zonename  string     `json:"zonename"`
}

type ApplyTungstenFabricTagParams struct {
	p map[string]interface{}
}

func (p *ApplyTungstenFabricTagParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["applicationpolicysetuuid"]; found {
		u.Set("applicationpolicysetuuid", v.(string))
	}
	if v, found := p.p["networkuuid"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("networkuuid", vv)
	}
	if v, found := p.p["nicuuid"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("nicuuid", vv)
	}
	if v, found := p.p["policyuuid"]; found {
		u.Set("policyuuid", v.(string))
	}
	if v, found := p.p["taguuid"]; found {
		u.Set("taguuid", v.(string))
	}
	if v, found := p.p["vmuuid"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("vmuuid", vv)
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *ApplyTungstenFabricTagParams) SetApplicationpolicysetuuid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["applicationpolicysetuuid"] = v
}

func (p *ApplyTungstenFabricTagParams) ResetApplicationpolicysetuuid() {
	if p.p != nil && p.p["applicationpolicysetuuid"] != nil {
		delete(p.p, "applicationpolicysetuuid")
	}
}

func (p *ApplyTungstenFabricTagParams) GetApplicationpolicysetuuid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["applicationpolicysetuuid"].(string)
	return value, ok
}

func (p *ApplyTungstenFabricTagParams) SetNetworkuuid(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["networkuuid"] = v
}

func (p *ApplyTungstenFabricTagParams) ResetNetworkuuid() {
	if p.p != nil && p.p["networkuuid"] != nil {
		delete(p.p, "networkuuid")
	}
}

func (p *ApplyTungstenFabricTagParams) GetNetworkuuid() ([]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["networkuuid"].([]string)
	return value, ok
}

func (p *ApplyTungstenFabricTagParams) SetNicuuid(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["nicuuid"] = v
}

func (p *ApplyTungstenFabricTagParams) ResetNicuuid() {
	if p.p != nil && p.p["nicuuid"] != nil {
		delete(p.p, "nicuuid")
	}
}

func (p *ApplyTungstenFabricTagParams) GetNicuuid() ([]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["nicuuid"].([]string)
	return value, ok
}

func (p *ApplyTungstenFabricTagParams) SetPolicyuuid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["policyuuid"] = v
}

func (p *ApplyTungstenFabricTagParams) ResetPolicyuuid() {
	if p.p != nil && p.p["policyuuid"] != nil {
		delete(p.p, "policyuuid")
	}
}

func (p *ApplyTungstenFabricTagParams) GetPolicyuuid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["policyuuid"].(string)
	return value, ok
}

func (p *ApplyTungstenFabricTagParams) SetTaguuid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["taguuid"] = v
}

func (p *ApplyTungstenFabricTagParams) ResetTaguuid() {
	if p.p != nil && p.p["taguuid"] != nil {
		delete(p.p, "taguuid")
	}
}

func (p *ApplyTungstenFabricTagParams) GetTaguuid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["taguuid"].(string)
	return value, ok
}

func (p *ApplyTungstenFabricTagParams) SetVmuuid(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["vmuuid"] = v
}

func (p *ApplyTungstenFabricTagParams) ResetVmuuid() {
	if p.p != nil && p.p["vmuuid"] != nil {
		delete(p.p, "vmuuid")
	}
}

func (p *ApplyTungstenFabricTagParams) GetVmuuid() ([]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["vmuuid"].([]string)
	return value, ok
}

func (p *ApplyTungstenFabricTagParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

func (p *ApplyTungstenFabricTagParams) ResetZoneid() {
	if p.p != nil && p.p["zoneid"] != nil {
		delete(p.p, "zoneid")
	}
}

func (p *ApplyTungstenFabricTagParams) GetZoneid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new ApplyTungstenFabricTagParams instance,
// as then you are sure you have configured all required params
func (s *TungstenService) NewApplyTungstenFabricTagParams(taguuid string, zoneid string) *ApplyTungstenFabricTagParams {
	p := &ApplyTungstenFabricTagParams{}
	p.p = make(map[string]interface{})
	p.p["taguuid"] = taguuid
	p.p["zoneid"] = zoneid
	return p
}

// apply Tungsten-Fabric tag
func (s *TungstenService) ApplyTungstenFabricTag(p *ApplyTungstenFabricTagParams) (*ApplyTungstenFabricTagResponse, error) {
	resp, err := s.cs.newRequest("applyTungstenFabricTag", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ApplyTungstenFabricTagResponse
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

type ApplyTungstenFabricTagResponse struct {
	JobID     string     `json:"jobid"`
	Jobstatus int        `json:"jobstatus"`
	Name      string     `json:"name"`
	Network   []*Network `json:"network"`
	Nic       []string   `json:"nic"`
	Policy    []string   `json:"policy"`
	Uuid      string     `json:"uuid"`
	Vm        []string   `json:"vm"`
	Zoneid    int64      `json:"zoneid"`
	Zonename  string     `json:"zonename"`
}

type CreateTungstenFabricAddressGroupParams struct {
	p map[string]interface{}
}

func (p *CreateTungstenFabricAddressGroupParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["ipprefix"]; found {
		u.Set("ipprefix", v.(string))
	}
	if v, found := p.p["ipprefixlen"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("ipprefixlen", vv)
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *CreateTungstenFabricAddressGroupParams) SetIpprefix(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ipprefix"] = v
}

func (p *CreateTungstenFabricAddressGroupParams) ResetIpprefix() {
	if p.p != nil && p.p["ipprefix"] != nil {
		delete(p.p, "ipprefix")
	}
}

func (p *CreateTungstenFabricAddressGroupParams) GetIpprefix() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["ipprefix"].(string)
	return value, ok
}

func (p *CreateTungstenFabricAddressGroupParams) SetIpprefixlen(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ipprefixlen"] = v
}

func (p *CreateTungstenFabricAddressGroupParams) ResetIpprefixlen() {
	if p.p != nil && p.p["ipprefixlen"] != nil {
		delete(p.p, "ipprefixlen")
	}
}

func (p *CreateTungstenFabricAddressGroupParams) GetIpprefixlen() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["ipprefixlen"].(int)
	return value, ok
}

func (p *CreateTungstenFabricAddressGroupParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *CreateTungstenFabricAddressGroupParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
}

func (p *CreateTungstenFabricAddressGroupParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *CreateTungstenFabricAddressGroupParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

func (p *CreateTungstenFabricAddressGroupParams) ResetZoneid() {
	if p.p != nil && p.p["zoneid"] != nil {
		delete(p.p, "zoneid")
	}
}

func (p *CreateTungstenFabricAddressGroupParams) GetZoneid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new CreateTungstenFabricAddressGroupParams instance,
// as then you are sure you have configured all required params
func (s *TungstenService) NewCreateTungstenFabricAddressGroupParams(ipprefix string, ipprefixlen int, name string, zoneid string) *CreateTungstenFabricAddressGroupParams {
	p := &CreateTungstenFabricAddressGroupParams{}
	p.p = make(map[string]interface{})
	p.p["ipprefix"] = ipprefix
	p.p["ipprefixlen"] = ipprefixlen
	p.p["name"] = name
	p.p["zoneid"] = zoneid
	return p
}

// create Tungsten-Fabric address group
func (s *TungstenService) CreateTungstenFabricAddressGroup(p *CreateTungstenFabricAddressGroupParams) (*CreateTungstenFabricAddressGroupResponse, error) {
	resp, err := s.cs.newRequest("createTungstenFabricAddressGroup", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreateTungstenFabricAddressGroupResponse
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

type CreateTungstenFabricAddressGroupResponse struct {
	Ipprefix    string `json:"ipprefix"`
	Ipprefixlen int    `json:"ipprefixlen"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Name        string `json:"name"`
	Uuid        string `json:"uuid"`
	Zoneid      int64  `json:"zoneid"`
	Zonename    string `json:"zonename"`
}

type CreateTungstenFabricApplicationPolicySetParams struct {
	p map[string]interface{}
}

func (p *CreateTungstenFabricApplicationPolicySetParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *CreateTungstenFabricApplicationPolicySetParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *CreateTungstenFabricApplicationPolicySetParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
}

func (p *CreateTungstenFabricApplicationPolicySetParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *CreateTungstenFabricApplicationPolicySetParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

func (p *CreateTungstenFabricApplicationPolicySetParams) ResetZoneid() {
	if p.p != nil && p.p["zoneid"] != nil {
		delete(p.p, "zoneid")
	}
}

func (p *CreateTungstenFabricApplicationPolicySetParams) GetZoneid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new CreateTungstenFabricApplicationPolicySetParams instance,
// as then you are sure you have configured all required params
func (s *TungstenService) NewCreateTungstenFabricApplicationPolicySetParams(name string, zoneid string) *CreateTungstenFabricApplicationPolicySetParams {
	p := &CreateTungstenFabricApplicationPolicySetParams{}
	p.p = make(map[string]interface{})
	p.p["name"] = name
	p.p["zoneid"] = zoneid
	return p
}

// create Tungsten-Fabric application policy set
func (s *TungstenService) CreateTungstenFabricApplicationPolicySet(p *CreateTungstenFabricApplicationPolicySetParams) (*CreateTungstenFabricApplicationPolicySetResponse, error) {
	resp, err := s.cs.newRequest("createTungstenFabricApplicationPolicySet", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreateTungstenFabricApplicationPolicySetResponse
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

type CreateTungstenFabricApplicationPolicySetResponse struct {
	Firewallpolicy []string `json:"firewallpolicy"`
	JobID          string   `json:"jobid"`
	Jobstatus      int      `json:"jobstatus"`
	Name           string   `json:"name"`
	Tag            []string `json:"tag"`
	Uuid           string   `json:"uuid"`
	Zoneid         int64    `json:"zoneid"`
	Zonename       string   `json:"zonename"`
}

type CreateTungstenFabricServiceGroupParams struct {
	p map[string]interface{}
}

func (p *CreateTungstenFabricServiceGroupParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["endport"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("endport", vv)
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["protocol"]; found {
		u.Set("protocol", v.(string))
	}
	if v, found := p.p["startport"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("startport", vv)
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *CreateTungstenFabricServiceGroupParams) SetEndport(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["endport"] = v
}

func (p *CreateTungstenFabricServiceGroupParams) ResetEndport() {
	if p.p != nil && p.p["endport"] != nil {
		delete(p.p, "endport")
	}
}

func (p *CreateTungstenFabricServiceGroupParams) GetEndport() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["endport"].(int)
	return value, ok
}

func (p *CreateTungstenFabricServiceGroupParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *CreateTungstenFabricServiceGroupParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
}

func (p *CreateTungstenFabricServiceGroupParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *CreateTungstenFabricServiceGroupParams) SetProtocol(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["protocol"] = v
}

func (p *CreateTungstenFabricServiceGroupParams) ResetProtocol() {
	if p.p != nil && p.p["protocol"] != nil {
		delete(p.p, "protocol")
	}
}

func (p *CreateTungstenFabricServiceGroupParams) GetProtocol() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["protocol"].(string)
	return value, ok
}

func (p *CreateTungstenFabricServiceGroupParams) SetStartport(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["startport"] = v
}

func (p *CreateTungstenFabricServiceGroupParams) ResetStartport() {
	if p.p != nil && p.p["startport"] != nil {
		delete(p.p, "startport")
	}
}

func (p *CreateTungstenFabricServiceGroupParams) GetStartport() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["startport"].(int)
	return value, ok
}

func (p *CreateTungstenFabricServiceGroupParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

func (p *CreateTungstenFabricServiceGroupParams) ResetZoneid() {
	if p.p != nil && p.p["zoneid"] != nil {
		delete(p.p, "zoneid")
	}
}

func (p *CreateTungstenFabricServiceGroupParams) GetZoneid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new CreateTungstenFabricServiceGroupParams instance,
// as then you are sure you have configured all required params
func (s *TungstenService) NewCreateTungstenFabricServiceGroupParams(endport int, name string, protocol string, startport int, zoneid string) *CreateTungstenFabricServiceGroupParams {
	p := &CreateTungstenFabricServiceGroupParams{}
	p.p = make(map[string]interface{})
	p.p["endport"] = endport
	p.p["name"] = name
	p.p["protocol"] = protocol
	p.p["startport"] = startport
	p.p["zoneid"] = zoneid
	return p
}

// create Tungsten-Fabric service group
func (s *TungstenService) CreateTungstenFabricServiceGroup(p *CreateTungstenFabricServiceGroupParams) (*CreateTungstenFabricServiceGroupResponse, error) {
	resp, err := s.cs.newRequest("createTungstenFabricServiceGroup", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreateTungstenFabricServiceGroupResponse
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

type CreateTungstenFabricServiceGroupResponse struct {
	Endport   int    `json:"endport"`
	JobID     string `json:"jobid"`
	Jobstatus int    `json:"jobstatus"`
	Name      string `json:"name"`
	Protocol  string `json:"protocol"`
	Startport int    `json:"startport"`
	Uuid      string `json:"uuid"`
	Zoneid    int64  `json:"zoneid"`
	Zonename  string `json:"zonename"`
}

type CreateTungstenFabricPolicyParams struct {
	p map[string]interface{}
}

func (p *CreateTungstenFabricPolicyParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *CreateTungstenFabricPolicyParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *CreateTungstenFabricPolicyParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
}

func (p *CreateTungstenFabricPolicyParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *CreateTungstenFabricPolicyParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

func (p *CreateTungstenFabricPolicyParams) ResetZoneid() {
	if p.p != nil && p.p["zoneid"] != nil {
		delete(p.p, "zoneid")
	}
}

func (p *CreateTungstenFabricPolicyParams) GetZoneid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new CreateTungstenFabricPolicyParams instance,
// as then you are sure you have configured all required params
func (s *TungstenService) NewCreateTungstenFabricPolicyParams(name string, zoneid string) *CreateTungstenFabricPolicyParams {
	p := &CreateTungstenFabricPolicyParams{}
	p.p = make(map[string]interface{})
	p.p["name"] = name
	p.p["zoneid"] = zoneid
	return p
}

// create Tungsten-Fabric policy
func (s *TungstenService) CreateTungstenFabricPolicy(p *CreateTungstenFabricPolicyParams) (*CreateTungstenFabricPolicyResponse, error) {
	resp, err := s.cs.newRequest("createTungstenFabricPolicy", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreateTungstenFabricPolicyResponse
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

type CreateTungstenFabricPolicyResponse struct {
	JobID     string     `json:"jobid"`
	Jobstatus int        `json:"jobstatus"`
	Name      string     `json:"name"`
	Network   []*Network `json:"network"`
	Uuid      string     `json:"uuid"`
	Zoneid    int64      `json:"zoneid"`
	Zonename  string     `json:"zonename"`
}

type CreateTungstenFabricProviderParams struct {
	p map[string]interface{}
}

func (p *CreateTungstenFabricProviderParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["tungstengateway"]; found {
		u.Set("tungstengateway", v.(string))
	}
	if v, found := p.p["tungstenproviderhostname"]; found {
		u.Set("tungstenproviderhostname", v.(string))
	}
	if v, found := p.p["tungstenproviderintrospectport"]; found {
		u.Set("tungstenproviderintrospectport", v.(string))
	}
	if v, found := p.p["tungstenproviderport"]; found {
		u.Set("tungstenproviderport", v.(string))
	}
	if v, found := p.p["tungstenprovidervrouterport"]; found {
		u.Set("tungstenprovidervrouterport", v.(string))
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *CreateTungstenFabricProviderParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *CreateTungstenFabricProviderParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
}

func (p *CreateTungstenFabricProviderParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *CreateTungstenFabricProviderParams) SetTungstengateway(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["tungstengateway"] = v
}

func (p *CreateTungstenFabricProviderParams) ResetTungstengateway() {
	if p.p != nil && p.p["tungstengateway"] != nil {
		delete(p.p, "tungstengateway")
	}
}

func (p *CreateTungstenFabricProviderParams) GetTungstengateway() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["tungstengateway"].(string)
	return value, ok
}

func (p *CreateTungstenFabricProviderParams) SetTungstenproviderhostname(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["tungstenproviderhostname"] = v
}

func (p *CreateTungstenFabricProviderParams) ResetTungstenproviderhostname() {
	if p.p != nil && p.p["tungstenproviderhostname"] != nil {
		delete(p.p, "tungstenproviderhostname")
	}
}

func (p *CreateTungstenFabricProviderParams) GetTungstenproviderhostname() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["tungstenproviderhostname"].(string)
	return value, ok
}

func (p *CreateTungstenFabricProviderParams) SetTungstenproviderintrospectport(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["tungstenproviderintrospectport"] = v
}

func (p *CreateTungstenFabricProviderParams) ResetTungstenproviderintrospectport() {
	if p.p != nil && p.p["tungstenproviderintrospectport"] != nil {
		delete(p.p, "tungstenproviderintrospectport")
	}
}

func (p *CreateTungstenFabricProviderParams) GetTungstenproviderintrospectport() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["tungstenproviderintrospectport"].(string)
	return value, ok
}

func (p *CreateTungstenFabricProviderParams) SetTungstenproviderport(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["tungstenproviderport"] = v
}

func (p *CreateTungstenFabricProviderParams) ResetTungstenproviderport() {
	if p.p != nil && p.p["tungstenproviderport"] != nil {
		delete(p.p, "tungstenproviderport")
	}
}

func (p *CreateTungstenFabricProviderParams) GetTungstenproviderport() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["tungstenproviderport"].(string)
	return value, ok
}

func (p *CreateTungstenFabricProviderParams) SetTungstenprovidervrouterport(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["tungstenprovidervrouterport"] = v
}

func (p *CreateTungstenFabricProviderParams) ResetTungstenprovidervrouterport() {
	if p.p != nil && p.p["tungstenprovidervrouterport"] != nil {
		delete(p.p, "tungstenprovidervrouterport")
	}
}

func (p *CreateTungstenFabricProviderParams) GetTungstenprovidervrouterport() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["tungstenprovidervrouterport"].(string)
	return value, ok
}

func (p *CreateTungstenFabricProviderParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

func (p *CreateTungstenFabricProviderParams) ResetZoneid() {
	if p.p != nil && p.p["zoneid"] != nil {
		delete(p.p, "zoneid")
	}
}

func (p *CreateTungstenFabricProviderParams) GetZoneid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new CreateTungstenFabricProviderParams instance,
// as then you are sure you have configured all required params
func (s *TungstenService) NewCreateTungstenFabricProviderParams(name string, tungstengateway string, tungstenproviderhostname string, zoneid string) *CreateTungstenFabricProviderParams {
	p := &CreateTungstenFabricProviderParams{}
	p.p = make(map[string]interface{})
	p.p["name"] = name
	p.p["tungstengateway"] = tungstengateway
	p.p["tungstenproviderhostname"] = tungstenproviderhostname
	p.p["zoneid"] = zoneid
	return p
}

// Create Tungsten-Fabric provider in cloudstack
func (s *TungstenService) CreateTungstenFabricProvider(p *CreateTungstenFabricProviderParams) (*CreateTungstenFabricProviderResponse, error) {
	resp, err := s.cs.newRequest("createTungstenFabricProvider", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreateTungstenFabricProviderResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type CreateTungstenFabricProviderResponse struct {
	JobID                          string `json:"jobid"`
	Jobstatus                      int    `json:"jobstatus"`
	Name                           string `json:"name"`
	Securitygroupsenabled          bool   `json:"securitygroupsenabled"`
	Tungstengateway                string `json:"tungstengateway"`
	Tungstenproviderhostname       string `json:"tungstenproviderhostname"`
	Tungstenproviderintrospectport string `json:"tungstenproviderintrospectport"`
	Tungstenproviderport           string `json:"tungstenproviderport"`
	Tungstenprovideruuid           string `json:"tungstenprovideruuid"`
	Tungstenprovidervrouterport    string `json:"tungstenprovidervrouterport"`
	Zoneid                         int64  `json:"zoneid"`
	Zonename                       string `json:"zonename"`
}

type CreateTungstenFabricTagParams struct {
	p map[string]interface{}
}

func (p *CreateTungstenFabricTagParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["tagtype"]; found {
		u.Set("tagtype", v.(string))
	}
	if v, found := p.p["tagvalue"]; found {
		u.Set("tagvalue", v.(string))
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *CreateTungstenFabricTagParams) SetTagtype(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["tagtype"] = v
}

func (p *CreateTungstenFabricTagParams) ResetTagtype() {
	if p.p != nil && p.p["tagtype"] != nil {
		delete(p.p, "tagtype")
	}
}

func (p *CreateTungstenFabricTagParams) GetTagtype() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["tagtype"].(string)
	return value, ok
}

func (p *CreateTungstenFabricTagParams) SetTagvalue(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["tagvalue"] = v
}

func (p *CreateTungstenFabricTagParams) ResetTagvalue() {
	if p.p != nil && p.p["tagvalue"] != nil {
		delete(p.p, "tagvalue")
	}
}

func (p *CreateTungstenFabricTagParams) GetTagvalue() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["tagvalue"].(string)
	return value, ok
}

func (p *CreateTungstenFabricTagParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

func (p *CreateTungstenFabricTagParams) ResetZoneid() {
	if p.p != nil && p.p["zoneid"] != nil {
		delete(p.p, "zoneid")
	}
}

func (p *CreateTungstenFabricTagParams) GetZoneid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new CreateTungstenFabricTagParams instance,
// as then you are sure you have configured all required params
func (s *TungstenService) NewCreateTungstenFabricTagParams(tagtype string, tagvalue string, zoneid string) *CreateTungstenFabricTagParams {
	p := &CreateTungstenFabricTagParams{}
	p.p = make(map[string]interface{})
	p.p["tagtype"] = tagtype
	p.p["tagvalue"] = tagvalue
	p.p["zoneid"] = zoneid
	return p
}

// create Tungsten-Fabric tag
func (s *TungstenService) CreateTungstenFabricTag(p *CreateTungstenFabricTagParams) (*CreateTungstenFabricTagResponse, error) {
	resp, err := s.cs.newRequest("createTungstenFabricTag", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreateTungstenFabricTagResponse
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

type CreateTungstenFabricTagResponse struct {
	JobID     string     `json:"jobid"`
	Jobstatus int        `json:"jobstatus"`
	Name      string     `json:"name"`
	Network   []*Network `json:"network"`
	Nic       []string   `json:"nic"`
	Policy    []string   `json:"policy"`
	Uuid      string     `json:"uuid"`
	Vm        []string   `json:"vm"`
	Zoneid    int64      `json:"zoneid"`
	Zonename  string     `json:"zonename"`
}

type CreateTungstenFabricTagTypeParams struct {
	p map[string]interface{}
}

func (p *CreateTungstenFabricTagTypeParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *CreateTungstenFabricTagTypeParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
}

func (p *CreateTungstenFabricTagTypeParams) ResetName() {
	if p.p != nil && p.p["name"] != nil {
		delete(p.p, "name")
	}
}

func (p *CreateTungstenFabricTagTypeParams) GetName() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["name"].(string)
	return value, ok
}

func (p *CreateTungstenFabricTagTypeParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

func (p *CreateTungstenFabricTagTypeParams) ResetZoneid() {
	if p.p != nil && p.p["zoneid"] != nil {
		delete(p.p, "zoneid")
	}
}

func (p *CreateTungstenFabricTagTypeParams) GetZoneid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new CreateTungstenFabricTagTypeParams instance,
// as then you are sure you have configured all required params
func (s *TungstenService) NewCreateTungstenFabricTagTypeParams(name string, zoneid string) *CreateTungstenFabricTagTypeParams {
	p := &CreateTungstenFabricTagTypeParams{}
	p.p = make(map[string]interface{})
	p.p["name"] = name
	p.p["zoneid"] = zoneid
	return p
}

// create Tungsten-Fabric tag type
func (s *TungstenService) CreateTungstenFabricTagType(p *CreateTungstenFabricTagTypeParams) (*CreateTungstenFabricTagTypeResponse, error) {
	resp, err := s.cs.newRequest("createTungstenFabricTagType", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreateTungstenFabricTagTypeResponse
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

type CreateTungstenFabricTagTypeResponse struct {
	JobID     string `json:"jobid"`
	Jobstatus int    `json:"jobstatus"`
	Name      string `json:"name"`
	Uuid      string `json:"uuid"`
	Zoneid    int64  `json:"zoneid"`
	Zonename  string `json:"zonename"`
}

type ConfigTungstenFabricServiceParams struct {
	p map[string]interface{}
}

func (p *ConfigTungstenFabricServiceParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["physicalnetworkid"]; found {
		u.Set("physicalnetworkid", v.(string))
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *ConfigTungstenFabricServiceParams) SetPhysicalnetworkid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["physicalnetworkid"] = v
}

func (p *ConfigTungstenFabricServiceParams) ResetPhysicalnetworkid() {
	if p.p != nil && p.p["physicalnetworkid"] != nil {
		delete(p.p, "physicalnetworkid")
	}
}

func (p *ConfigTungstenFabricServiceParams) GetPhysicalnetworkid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["physicalnetworkid"].(string)
	return value, ok
}

func (p *ConfigTungstenFabricServiceParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

func (p *ConfigTungstenFabricServiceParams) ResetZoneid() {
	if p.p != nil && p.p["zoneid"] != nil {
		delete(p.p, "zoneid")
	}
}

func (p *ConfigTungstenFabricServiceParams) GetZoneid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new ConfigTungstenFabricServiceParams instance,
// as then you are sure you have configured all required params
func (s *TungstenService) NewConfigTungstenFabricServiceParams(physicalnetworkid string, zoneid string) *ConfigTungstenFabricServiceParams {
	p := &ConfigTungstenFabricServiceParams{}
	p.p = make(map[string]interface{})
	p.p["physicalnetworkid"] = physicalnetworkid
	p.p["zoneid"] = zoneid
	return p
}

// config Tungsten-Fabric service
func (s *TungstenService) ConfigTungstenFabricService(p *ConfigTungstenFabricServiceParams) (*ConfigTungstenFabricServiceResponse, error) {
	resp, err := s.cs.newRequest("configTungstenFabricService", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ConfigTungstenFabricServiceResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ConfigTungstenFabricServiceResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

func (r *ConfigTungstenFabricServiceResponse) UnmarshalJSON(b []byte) error {
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

	type alias ConfigTungstenFabricServiceResponse
	return json.Unmarshal(b, (*alias)(r))
}

type DeleteTungstenFabricAddressGroupParams struct {
	p map[string]interface{}
}

func (p *DeleteTungstenFabricAddressGroupParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["addressgroupuuid"]; found {
		u.Set("addressgroupuuid", v.(string))
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *DeleteTungstenFabricAddressGroupParams) SetAddressgroupuuid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["addressgroupuuid"] = v
}

func (p *DeleteTungstenFabricAddressGroupParams) ResetAddressgroupuuid() {
	if p.p != nil && p.p["addressgroupuuid"] != nil {
		delete(p.p, "addressgroupuuid")
	}
}

func (p *DeleteTungstenFabricAddressGroupParams) GetAddressgroupuuid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["addressgroupuuid"].(string)
	return value, ok
}

func (p *DeleteTungstenFabricAddressGroupParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

func (p *DeleteTungstenFabricAddressGroupParams) ResetZoneid() {
	if p.p != nil && p.p["zoneid"] != nil {
		delete(p.p, "zoneid")
	}
}

func (p *DeleteTungstenFabricAddressGroupParams) GetZoneid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new DeleteTungstenFabricAddressGroupParams instance,
// as then you are sure you have configured all required params
func (s *TungstenService) NewDeleteTungstenFabricAddressGroupParams(addressgroupuuid string, zoneid string) *DeleteTungstenFabricAddressGroupParams {
	p := &DeleteTungstenFabricAddressGroupParams{}
	p.p = make(map[string]interface{})
	p.p["addressgroupuuid"] = addressgroupuuid
	p.p["zoneid"] = zoneid
	return p
}

// delete Tungsten-Fabric address group
func (s *TungstenService) DeleteTungstenFabricAddressGroup(p *DeleteTungstenFabricAddressGroupParams) (*DeleteTungstenFabricAddressGroupResponse, error) {
	resp, err := s.cs.newRequest("deleteTungstenFabricAddressGroup", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteTungstenFabricAddressGroupResponse
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

type DeleteTungstenFabricAddressGroupResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type DeleteTungstenFabricApplicationPolicySetParams struct {
	p map[string]interface{}
}

func (p *DeleteTungstenFabricApplicationPolicySetParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["applicationpolicysetuuid"]; found {
		u.Set("applicationpolicysetuuid", v.(string))
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *DeleteTungstenFabricApplicationPolicySetParams) SetApplicationpolicysetuuid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["applicationpolicysetuuid"] = v
}

func (p *DeleteTungstenFabricApplicationPolicySetParams) ResetApplicationpolicysetuuid() {
	if p.p != nil && p.p["applicationpolicysetuuid"] != nil {
		delete(p.p, "applicationpolicysetuuid")
	}
}

func (p *DeleteTungstenFabricApplicationPolicySetParams) GetApplicationpolicysetuuid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["applicationpolicysetuuid"].(string)
	return value, ok
}

func (p *DeleteTungstenFabricApplicationPolicySetParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

func (p *DeleteTungstenFabricApplicationPolicySetParams) ResetZoneid() {
	if p.p != nil && p.p["zoneid"] != nil {
		delete(p.p, "zoneid")
	}
}

func (p *DeleteTungstenFabricApplicationPolicySetParams) GetZoneid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new DeleteTungstenFabricApplicationPolicySetParams instance,
// as then you are sure you have configured all required params
func (s *TungstenService) NewDeleteTungstenFabricApplicationPolicySetParams(applicationpolicysetuuid string, zoneid string) *DeleteTungstenFabricApplicationPolicySetParams {
	p := &DeleteTungstenFabricApplicationPolicySetParams{}
	p.p = make(map[string]interface{})
	p.p["applicationpolicysetuuid"] = applicationpolicysetuuid
	p.p["zoneid"] = zoneid
	return p
}

// delete Tungsten-Fabric application policy set
func (s *TungstenService) DeleteTungstenFabricApplicationPolicySet(p *DeleteTungstenFabricApplicationPolicySetParams) (*DeleteTungstenFabricApplicationPolicySetResponse, error) {
	resp, err := s.cs.newRequest("deleteTungstenFabricApplicationPolicySet", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteTungstenFabricApplicationPolicySetResponse
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

type DeleteTungstenFabricApplicationPolicySetResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type DeleteTungstenFabricPolicyParams struct {
	p map[string]interface{}
}

func (p *DeleteTungstenFabricPolicyParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["policyuuid"]; found {
		u.Set("policyuuid", v.(string))
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *DeleteTungstenFabricPolicyParams) SetPolicyuuid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["policyuuid"] = v
}

func (p *DeleteTungstenFabricPolicyParams) ResetPolicyuuid() {
	if p.p != nil && p.p["policyuuid"] != nil {
		delete(p.p, "policyuuid")
	}
}

func (p *DeleteTungstenFabricPolicyParams) GetPolicyuuid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["policyuuid"].(string)
	return value, ok
}

func (p *DeleteTungstenFabricPolicyParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

func (p *DeleteTungstenFabricPolicyParams) ResetZoneid() {
	if p.p != nil && p.p["zoneid"] != nil {
		delete(p.p, "zoneid")
	}
}

func (p *DeleteTungstenFabricPolicyParams) GetZoneid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new DeleteTungstenFabricPolicyParams instance,
// as then you are sure you have configured all required params
func (s *TungstenService) NewDeleteTungstenFabricPolicyParams(policyuuid string, zoneid string) *DeleteTungstenFabricPolicyParams {
	p := &DeleteTungstenFabricPolicyParams{}
	p.p = make(map[string]interface{})
	p.p["policyuuid"] = policyuuid
	p.p["zoneid"] = zoneid
	return p
}

// delete Tungsten-Fabric policy
func (s *TungstenService) DeleteTungstenFabricPolicy(p *DeleteTungstenFabricPolicyParams) (*DeleteTungstenFabricPolicyResponse, error) {
	resp, err := s.cs.newRequest("deleteTungstenFabricPolicy", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteTungstenFabricPolicyResponse
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

type DeleteTungstenFabricPolicyResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type DeleteTungstenFabricTagParams struct {
	p map[string]interface{}
}

func (p *DeleteTungstenFabricTagParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["taguuid"]; found {
		u.Set("taguuid", v.(string))
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *DeleteTungstenFabricTagParams) SetTaguuid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["taguuid"] = v
}

func (p *DeleteTungstenFabricTagParams) ResetTaguuid() {
	if p.p != nil && p.p["taguuid"] != nil {
		delete(p.p, "taguuid")
	}
}

func (p *DeleteTungstenFabricTagParams) GetTaguuid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["taguuid"].(string)
	return value, ok
}

func (p *DeleteTungstenFabricTagParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

func (p *DeleteTungstenFabricTagParams) ResetZoneid() {
	if p.p != nil && p.p["zoneid"] != nil {
		delete(p.p, "zoneid")
	}
}

func (p *DeleteTungstenFabricTagParams) GetZoneid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new DeleteTungstenFabricTagParams instance,
// as then you are sure you have configured all required params
func (s *TungstenService) NewDeleteTungstenFabricTagParams(taguuid string, zoneid string) *DeleteTungstenFabricTagParams {
	p := &DeleteTungstenFabricTagParams{}
	p.p = make(map[string]interface{})
	p.p["taguuid"] = taguuid
	p.p["zoneid"] = zoneid
	return p
}

// delete Tungsten-Fabric tag
func (s *TungstenService) DeleteTungstenFabricTag(p *DeleteTungstenFabricTagParams) (*DeleteTungstenFabricTagResponse, error) {
	resp, err := s.cs.newRequest("deleteTungstenFabricTag", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteTungstenFabricTagResponse
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

type DeleteTungstenFabricTagResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type DeleteTungstenFabricTagTypeParams struct {
	p map[string]interface{}
}

func (p *DeleteTungstenFabricTagTypeParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["tagtypeuuid"]; found {
		u.Set("tagtypeuuid", v.(string))
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *DeleteTungstenFabricTagTypeParams) SetTagtypeuuid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["tagtypeuuid"] = v
}

func (p *DeleteTungstenFabricTagTypeParams) ResetTagtypeuuid() {
	if p.p != nil && p.p["tagtypeuuid"] != nil {
		delete(p.p, "tagtypeuuid")
	}
}

func (p *DeleteTungstenFabricTagTypeParams) GetTagtypeuuid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["tagtypeuuid"].(string)
	return value, ok
}

func (p *DeleteTungstenFabricTagTypeParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

func (p *DeleteTungstenFabricTagTypeParams) ResetZoneid() {
	if p.p != nil && p.p["zoneid"] != nil {
		delete(p.p, "zoneid")
	}
}

func (p *DeleteTungstenFabricTagTypeParams) GetZoneid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new DeleteTungstenFabricTagTypeParams instance,
// as then you are sure you have configured all required params
func (s *TungstenService) NewDeleteTungstenFabricTagTypeParams(tagtypeuuid string, zoneid string) *DeleteTungstenFabricTagTypeParams {
	p := &DeleteTungstenFabricTagTypeParams{}
	p.p = make(map[string]interface{})
	p.p["tagtypeuuid"] = tagtypeuuid
	p.p["zoneid"] = zoneid
	return p
}

// delete Tungsten-Fabric tag type
func (s *TungstenService) DeleteTungstenFabricTagType(p *DeleteTungstenFabricTagTypeParams) (*DeleteTungstenFabricTagTypeResponse, error) {
	resp, err := s.cs.newRequest("deleteTungstenFabricTagType", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteTungstenFabricTagTypeResponse
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

type DeleteTungstenFabricTagTypeResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type DeleteTungstenFabricServiceGroupParams struct {
	p map[string]interface{}
}

func (p *DeleteTungstenFabricServiceGroupParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["servicegroupuuid"]; found {
		u.Set("servicegroupuuid", v.(string))
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *DeleteTungstenFabricServiceGroupParams) SetServicegroupuuid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["servicegroupuuid"] = v
}

func (p *DeleteTungstenFabricServiceGroupParams) ResetServicegroupuuid() {
	if p.p != nil && p.p["servicegroupuuid"] != nil {
		delete(p.p, "servicegroupuuid")
	}
}

func (p *DeleteTungstenFabricServiceGroupParams) GetServicegroupuuid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["servicegroupuuid"].(string)
	return value, ok
}

func (p *DeleteTungstenFabricServiceGroupParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

func (p *DeleteTungstenFabricServiceGroupParams) ResetZoneid() {
	if p.p != nil && p.p["zoneid"] != nil {
		delete(p.p, "zoneid")
	}
}

func (p *DeleteTungstenFabricServiceGroupParams) GetZoneid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new DeleteTungstenFabricServiceGroupParams instance,
// as then you are sure you have configured all required params
func (s *TungstenService) NewDeleteTungstenFabricServiceGroupParams(servicegroupuuid string, zoneid string) *DeleteTungstenFabricServiceGroupParams {
	p := &DeleteTungstenFabricServiceGroupParams{}
	p.p = make(map[string]interface{})
	p.p["servicegroupuuid"] = servicegroupuuid
	p.p["zoneid"] = zoneid
	return p
}

// delete Tungsten-Fabric service group
func (s *TungstenService) DeleteTungstenFabricServiceGroup(p *DeleteTungstenFabricServiceGroupParams) (*DeleteTungstenFabricServiceGroupResponse, error) {
	resp, err := s.cs.newRequest("deleteTungstenFabricServiceGroup", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteTungstenFabricServiceGroupResponse
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

type DeleteTungstenFabricServiceGroupResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

type ListTungstenFabricAddressGroupParams struct {
	p map[string]interface{}
}

func (p *ListTungstenFabricAddressGroupParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["addressgroupuuid"]; found {
		u.Set("addressgroupuuid", v.(string))
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
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *ListTungstenFabricAddressGroupParams) SetAddressgroupuuid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["addressgroupuuid"] = v
}

func (p *ListTungstenFabricAddressGroupParams) ResetAddressgroupuuid() {
	if p.p != nil && p.p["addressgroupuuid"] != nil {
		delete(p.p, "addressgroupuuid")
	}
}

func (p *ListTungstenFabricAddressGroupParams) GetAddressgroupuuid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["addressgroupuuid"].(string)
	return value, ok
}

func (p *ListTungstenFabricAddressGroupParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListTungstenFabricAddressGroupParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
}

func (p *ListTungstenFabricAddressGroupParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListTungstenFabricAddressGroupParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListTungstenFabricAddressGroupParams) ResetPage() {
	if p.p != nil && p.p["page"] != nil {
		delete(p.p, "page")
	}
}

func (p *ListTungstenFabricAddressGroupParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListTungstenFabricAddressGroupParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListTungstenFabricAddressGroupParams) ResetPagesize() {
	if p.p != nil && p.p["pagesize"] != nil {
		delete(p.p, "pagesize")
	}
}

func (p *ListTungstenFabricAddressGroupParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

func (p *ListTungstenFabricAddressGroupParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

func (p *ListTungstenFabricAddressGroupParams) ResetZoneid() {
	if p.p != nil && p.p["zoneid"] != nil {
		delete(p.p, "zoneid")
	}
}

func (p *ListTungstenFabricAddressGroupParams) GetZoneid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new ListTungstenFabricAddressGroupParams instance,
// as then you are sure you have configured all required params
func (s *TungstenService) NewListTungstenFabricAddressGroupParams() *ListTungstenFabricAddressGroupParams {
	p := &ListTungstenFabricAddressGroupParams{}
	p.p = make(map[string]interface{})
	return p
}

// list Tungsten-Fabric address group
func (s *TungstenService) ListTungstenFabricAddressGroup(p *ListTungstenFabricAddressGroupParams) (*ListTungstenFabricAddressGroupResponse, error) {
	resp, err := s.cs.newRequest("listTungstenFabricAddressGroup", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListTungstenFabricAddressGroupResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListTungstenFabricAddressGroupResponse struct {
	Count                      int                           `json:"count"`
	TungstenFabricAddressGroup []*TungstenFabricAddressGroup `json:"tungstenfabricaddressgroup"`
}

type TungstenFabricAddressGroup struct {
	Ipprefix    string `json:"ipprefix"`
	Ipprefixlen int    `json:"ipprefixlen"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Name        string `json:"name"`
	Uuid        string `json:"uuid"`
	Zoneid      int64  `json:"zoneid"`
	Zonename    string `json:"zonename"`
}

type ListTungstenFabricApplicationPolicySetParams struct {
	p map[string]interface{}
}

func (p *ListTungstenFabricApplicationPolicySetParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["applicationpolicysetuuid"]; found {
		u.Set("applicationpolicysetuuid", v.(string))
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
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *ListTungstenFabricApplicationPolicySetParams) SetApplicationpolicysetuuid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["applicationpolicysetuuid"] = v
}

func (p *ListTungstenFabricApplicationPolicySetParams) ResetApplicationpolicysetuuid() {
	if p.p != nil && p.p["applicationpolicysetuuid"] != nil {
		delete(p.p, "applicationpolicysetuuid")
	}
}

func (p *ListTungstenFabricApplicationPolicySetParams) GetApplicationpolicysetuuid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["applicationpolicysetuuid"].(string)
	return value, ok
}

func (p *ListTungstenFabricApplicationPolicySetParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListTungstenFabricApplicationPolicySetParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
}

func (p *ListTungstenFabricApplicationPolicySetParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListTungstenFabricApplicationPolicySetParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListTungstenFabricApplicationPolicySetParams) ResetPage() {
	if p.p != nil && p.p["page"] != nil {
		delete(p.p, "page")
	}
}

func (p *ListTungstenFabricApplicationPolicySetParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListTungstenFabricApplicationPolicySetParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListTungstenFabricApplicationPolicySetParams) ResetPagesize() {
	if p.p != nil && p.p["pagesize"] != nil {
		delete(p.p, "pagesize")
	}
}

func (p *ListTungstenFabricApplicationPolicySetParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

func (p *ListTungstenFabricApplicationPolicySetParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

func (p *ListTungstenFabricApplicationPolicySetParams) ResetZoneid() {
	if p.p != nil && p.p["zoneid"] != nil {
		delete(p.p, "zoneid")
	}
}

func (p *ListTungstenFabricApplicationPolicySetParams) GetZoneid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new ListTungstenFabricApplicationPolicySetParams instance,
// as then you are sure you have configured all required params
func (s *TungstenService) NewListTungstenFabricApplicationPolicySetParams() *ListTungstenFabricApplicationPolicySetParams {
	p := &ListTungstenFabricApplicationPolicySetParams{}
	p.p = make(map[string]interface{})
	return p
}

// list Tungsten-Fabric application policy set
func (s *TungstenService) ListTungstenFabricApplicationPolicySet(p *ListTungstenFabricApplicationPolicySetParams) (*ListTungstenFabricApplicationPolicySetResponse, error) {
	resp, err := s.cs.newRequest("listTungstenFabricApplicationPolicySet", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListTungstenFabricApplicationPolicySetResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListTungstenFabricApplicationPolicySetResponse struct {
	Count                              int                                   `json:"count"`
	TungstenFabricApplicationPolicySet []*TungstenFabricApplicationPolicySet `json:"tungstenfabricapplicationpolicyset"`
}

type TungstenFabricApplicationPolicySet struct {
	Firewallpolicy []string `json:"firewallpolicy"`
	JobID          string   `json:"jobid"`
	Jobstatus      int      `json:"jobstatus"`
	Name           string   `json:"name"`
	Tag            []string `json:"tag"`
	Uuid           string   `json:"uuid"`
	Zoneid         int64    `json:"zoneid"`
	Zonename       string   `json:"zonename"`
}

type ListTungstenFabricLBHealthMonitorParams struct {
	p map[string]interface{}
}

func (p *ListTungstenFabricLBHealthMonitorParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["keyword"]; found {
		u.Set("keyword", v.(string))
	}
	if v, found := p.p["lbruleid"]; found {
		u.Set("lbruleid", v.(string))
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

func (p *ListTungstenFabricLBHealthMonitorParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListTungstenFabricLBHealthMonitorParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
}

func (p *ListTungstenFabricLBHealthMonitorParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListTungstenFabricLBHealthMonitorParams) SetLbruleid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["lbruleid"] = v
}

func (p *ListTungstenFabricLBHealthMonitorParams) ResetLbruleid() {
	if p.p != nil && p.p["lbruleid"] != nil {
		delete(p.p, "lbruleid")
	}
}

func (p *ListTungstenFabricLBHealthMonitorParams) GetLbruleid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["lbruleid"].(string)
	return value, ok
}

func (p *ListTungstenFabricLBHealthMonitorParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListTungstenFabricLBHealthMonitorParams) ResetPage() {
	if p.p != nil && p.p["page"] != nil {
		delete(p.p, "page")
	}
}

func (p *ListTungstenFabricLBHealthMonitorParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListTungstenFabricLBHealthMonitorParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListTungstenFabricLBHealthMonitorParams) ResetPagesize() {
	if p.p != nil && p.p["pagesize"] != nil {
		delete(p.p, "pagesize")
	}
}

func (p *ListTungstenFabricLBHealthMonitorParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

// You should always use this function to get a new ListTungstenFabricLBHealthMonitorParams instance,
// as then you are sure you have configured all required params
func (s *TungstenService) NewListTungstenFabricLBHealthMonitorParams(lbruleid string) *ListTungstenFabricLBHealthMonitorParams {
	p := &ListTungstenFabricLBHealthMonitorParams{}
	p.p = make(map[string]interface{})
	p.p["lbruleid"] = lbruleid
	return p
}

// list Tungsten-Fabric LB health monitor
func (s *TungstenService) ListTungstenFabricLBHealthMonitor(p *ListTungstenFabricLBHealthMonitorParams) (*ListTungstenFabricLBHealthMonitorResponse, error) {
	resp, err := s.cs.newRequest("listTungstenFabricLBHealthMonitor", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListTungstenFabricLBHealthMonitorResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListTungstenFabricLBHealthMonitorResponse struct {
	Count                         int                              `json:"count"`
	TungstenFabricLBHealthMonitor []*TungstenFabricLBHealthMonitor `json:"tungstenfabriclbhealthmonitor"`
}

type TungstenFabricLBHealthMonitor struct {
	Expectedcode string `json:"expectedcode"`
	Httpmethod   string `json:"httpmethod"`
	Id           int64  `json:"id"`
	Interval     int    `json:"interval"`
	JobID        string `json:"jobid"`
	Jobstatus    int    `json:"jobstatus"`
	Lbruleid     string `json:"lbruleid"`
	Retry        int    `json:"retry"`
	Timeout      int    `json:"timeout"`
	Type         string `json:"type"`
	Urlpath      string `json:"urlpath"`
	Uuid         string `json:"uuid"`
	Zoneid       int64  `json:"zoneid"`
	Zonename     string `json:"zonename"`
}

type ListTungstenFabricNicParams struct {
	p map[string]interface{}
}

func (p *ListTungstenFabricNicParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["keyword"]; found {
		u.Set("keyword", v.(string))
	}
	if v, found := p.p["nicuuid"]; found {
		u.Set("nicuuid", v.(string))
	}
	if v, found := p.p["page"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("page", vv)
	}
	if v, found := p.p["pagesize"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("pagesize", vv)
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *ListTungstenFabricNicParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListTungstenFabricNicParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
}

func (p *ListTungstenFabricNicParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListTungstenFabricNicParams) SetNicuuid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["nicuuid"] = v
}

func (p *ListTungstenFabricNicParams) ResetNicuuid() {
	if p.p != nil && p.p["nicuuid"] != nil {
		delete(p.p, "nicuuid")
	}
}

func (p *ListTungstenFabricNicParams) GetNicuuid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["nicuuid"].(string)
	return value, ok
}

func (p *ListTungstenFabricNicParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListTungstenFabricNicParams) ResetPage() {
	if p.p != nil && p.p["page"] != nil {
		delete(p.p, "page")
	}
}

func (p *ListTungstenFabricNicParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListTungstenFabricNicParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListTungstenFabricNicParams) ResetPagesize() {
	if p.p != nil && p.p["pagesize"] != nil {
		delete(p.p, "pagesize")
	}
}

func (p *ListTungstenFabricNicParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

func (p *ListTungstenFabricNicParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

func (p *ListTungstenFabricNicParams) ResetZoneid() {
	if p.p != nil && p.p["zoneid"] != nil {
		delete(p.p, "zoneid")
	}
}

func (p *ListTungstenFabricNicParams) GetZoneid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new ListTungstenFabricNicParams instance,
// as then you are sure you have configured all required params
func (s *TungstenService) NewListTungstenFabricNicParams() *ListTungstenFabricNicParams {
	p := &ListTungstenFabricNicParams{}
	p.p = make(map[string]interface{})
	return p
}

// list Tungsten-Fabric nic
func (s *TungstenService) ListTungstenFabricNic(p *ListTungstenFabricNicParams) (*ListTungstenFabricNicResponse, error) {
	resp, err := s.cs.newRequest("listTungstenFabricNic", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListTungstenFabricNicResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListTungstenFabricNicResponse struct {
	Count             int                  `json:"count"`
	TungstenFabricNic []*TungstenFabricNic `json:"tungstenfabricnic"`
}

type TungstenFabricNic struct {
	JobID     string `json:"jobid"`
	Jobstatus int    `json:"jobstatus"`
	Name      string `json:"name"`
	Uuid      string `json:"uuid"`
	Zoneid    int64  `json:"zoneid"`
	Zonename  string `json:"zonename"`
}

type ListTungstenFabricPolicyRuleParams struct {
	p map[string]interface{}
}

func (p *ListTungstenFabricPolicyRuleParams) toURLValues() url.Values {
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
	if v, found := p.p["policyuuid"]; found {
		u.Set("policyuuid", v.(string))
	}
	if v, found := p.p["ruleuuid"]; found {
		u.Set("ruleuuid", v.(string))
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *ListTungstenFabricPolicyRuleParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListTungstenFabricPolicyRuleParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
}

func (p *ListTungstenFabricPolicyRuleParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListTungstenFabricPolicyRuleParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListTungstenFabricPolicyRuleParams) ResetPage() {
	if p.p != nil && p.p["page"] != nil {
		delete(p.p, "page")
	}
}

func (p *ListTungstenFabricPolicyRuleParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListTungstenFabricPolicyRuleParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListTungstenFabricPolicyRuleParams) ResetPagesize() {
	if p.p != nil && p.p["pagesize"] != nil {
		delete(p.p, "pagesize")
	}
}

func (p *ListTungstenFabricPolicyRuleParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

func (p *ListTungstenFabricPolicyRuleParams) SetPolicyuuid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["policyuuid"] = v
}

func (p *ListTungstenFabricPolicyRuleParams) ResetPolicyuuid() {
	if p.p != nil && p.p["policyuuid"] != nil {
		delete(p.p, "policyuuid")
	}
}

func (p *ListTungstenFabricPolicyRuleParams) GetPolicyuuid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["policyuuid"].(string)
	return value, ok
}

func (p *ListTungstenFabricPolicyRuleParams) SetRuleuuid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ruleuuid"] = v
}

func (p *ListTungstenFabricPolicyRuleParams) ResetRuleuuid() {
	if p.p != nil && p.p["ruleuuid"] != nil {
		delete(p.p, "ruleuuid")
	}
}

func (p *ListTungstenFabricPolicyRuleParams) GetRuleuuid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["ruleuuid"].(string)
	return value, ok
}

func (p *ListTungstenFabricPolicyRuleParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

func (p *ListTungstenFabricPolicyRuleParams) ResetZoneid() {
	if p.p != nil && p.p["zoneid"] != nil {
		delete(p.p, "zoneid")
	}
}

func (p *ListTungstenFabricPolicyRuleParams) GetZoneid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new ListTungstenFabricPolicyRuleParams instance,
// as then you are sure you have configured all required params
func (s *TungstenService) NewListTungstenFabricPolicyRuleParams(policyuuid string) *ListTungstenFabricPolicyRuleParams {
	p := &ListTungstenFabricPolicyRuleParams{}
	p.p = make(map[string]interface{})
	p.p["policyuuid"] = policyuuid
	return p
}

// list Tungsten-Fabric policy
func (s *TungstenService) ListTungstenFabricPolicyRule(p *ListTungstenFabricPolicyRuleParams) (*ListTungstenFabricPolicyRuleResponse, error) {
	resp, err := s.cs.newRequest("listTungstenFabricPolicyRule", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListTungstenFabricPolicyRuleResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListTungstenFabricPolicyRuleResponse struct {
	Count                    int                         `json:"count"`
	TungstenFabricPolicyRule []*TungstenFabricPolicyRule `json:"tungstenfabricpolicyrule"`
}

type TungstenFabricPolicyRule struct {
	Action          string `json:"action"`
	Destendport     int    `json:"destendport"`
	Destipprefix    string `json:"destipprefix"`
	Destipprefixlen int    `json:"destipprefixlen"`
	Destnetwork     string `json:"destnetwork"`
	Deststartport   int    `json:"deststartport"`
	Direction       string `json:"direction"`
	JobID           string `json:"jobid"`
	Jobstatus       int    `json:"jobstatus"`
	Policyuuid      string `json:"policyuuid"`
	Protocol        string `json:"protocol"`
	Srcendport      int    `json:"srcendport"`
	Srcipprefix     string `json:"srcipprefix"`
	Srcipprefixlen  int    `json:"srcipprefixlen"`
	Srcnetwork      string `json:"srcnetwork"`
	Srcstartport    int    `json:"srcstartport"`
	Uuid            string `json:"uuid"`
	Zoneid          int64  `json:"zoneid"`
	Zonename        string `json:"zonename"`
}

type ListTungstenFabricPolicyParams struct {
	p map[string]interface{}
}

func (p *ListTungstenFabricPolicyParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["ipaddressid"]; found {
		u.Set("ipaddressid", v.(string))
	}
	if v, found := p.p["keyword"]; found {
		u.Set("keyword", v.(string))
	}
	if v, found := p.p["networkid"]; found {
		u.Set("networkid", v.(string))
	}
	if v, found := p.p["page"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("page", vv)
	}
	if v, found := p.p["pagesize"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("pagesize", vv)
	}
	if v, found := p.p["policyuuid"]; found {
		u.Set("policyuuid", v.(string))
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *ListTungstenFabricPolicyParams) SetIpaddressid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ipaddressid"] = v
}

func (p *ListTungstenFabricPolicyParams) ResetIpaddressid() {
	if p.p != nil && p.p["ipaddressid"] != nil {
		delete(p.p, "ipaddressid")
	}
}

func (p *ListTungstenFabricPolicyParams) GetIpaddressid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["ipaddressid"].(string)
	return value, ok
}

func (p *ListTungstenFabricPolicyParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListTungstenFabricPolicyParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
}

func (p *ListTungstenFabricPolicyParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListTungstenFabricPolicyParams) SetNetworkid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["networkid"] = v
}

func (p *ListTungstenFabricPolicyParams) ResetNetworkid() {
	if p.p != nil && p.p["networkid"] != nil {
		delete(p.p, "networkid")
	}
}

func (p *ListTungstenFabricPolicyParams) GetNetworkid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["networkid"].(string)
	return value, ok
}

func (p *ListTungstenFabricPolicyParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListTungstenFabricPolicyParams) ResetPage() {
	if p.p != nil && p.p["page"] != nil {
		delete(p.p, "page")
	}
}

func (p *ListTungstenFabricPolicyParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListTungstenFabricPolicyParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListTungstenFabricPolicyParams) ResetPagesize() {
	if p.p != nil && p.p["pagesize"] != nil {
		delete(p.p, "pagesize")
	}
}

func (p *ListTungstenFabricPolicyParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

func (p *ListTungstenFabricPolicyParams) SetPolicyuuid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["policyuuid"] = v
}

func (p *ListTungstenFabricPolicyParams) ResetPolicyuuid() {
	if p.p != nil && p.p["policyuuid"] != nil {
		delete(p.p, "policyuuid")
	}
}

func (p *ListTungstenFabricPolicyParams) GetPolicyuuid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["policyuuid"].(string)
	return value, ok
}

func (p *ListTungstenFabricPolicyParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

func (p *ListTungstenFabricPolicyParams) ResetZoneid() {
	if p.p != nil && p.p["zoneid"] != nil {
		delete(p.p, "zoneid")
	}
}

func (p *ListTungstenFabricPolicyParams) GetZoneid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new ListTungstenFabricPolicyParams instance,
// as then you are sure you have configured all required params
func (s *TungstenService) NewListTungstenFabricPolicyParams() *ListTungstenFabricPolicyParams {
	p := &ListTungstenFabricPolicyParams{}
	p.p = make(map[string]interface{})
	return p
}

// list Tungsten-Fabric policy
func (s *TungstenService) ListTungstenFabricPolicy(p *ListTungstenFabricPolicyParams) (*ListTungstenFabricPolicyResponse, error) {
	resp, err := s.cs.newRequest("listTungstenFabricPolicy", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListTungstenFabricPolicyResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListTungstenFabricPolicyResponse struct {
	Count                int                     `json:"count"`
	TungstenFabricPolicy []*TungstenFabricPolicy `json:"tungstenfabricpolicy"`
}

type TungstenFabricPolicy struct {
	JobID     string     `json:"jobid"`
	Jobstatus int        `json:"jobstatus"`
	Name      string     `json:"name"`
	Network   []*Network `json:"network"`
	Uuid      string     `json:"uuid"`
	Zoneid    int64      `json:"zoneid"`
	Zonename  string     `json:"zonename"`
}

type ListTungstenFabricProvidersParams struct {
	p map[string]interface{}
}

func (p *ListTungstenFabricProvidersParams) toURLValues() url.Values {
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
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *ListTungstenFabricProvidersParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListTungstenFabricProvidersParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
}

func (p *ListTungstenFabricProvidersParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListTungstenFabricProvidersParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListTungstenFabricProvidersParams) ResetPage() {
	if p.p != nil && p.p["page"] != nil {
		delete(p.p, "page")
	}
}

func (p *ListTungstenFabricProvidersParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListTungstenFabricProvidersParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListTungstenFabricProvidersParams) ResetPagesize() {
	if p.p != nil && p.p["pagesize"] != nil {
		delete(p.p, "pagesize")
	}
}

func (p *ListTungstenFabricProvidersParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

func (p *ListTungstenFabricProvidersParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

func (p *ListTungstenFabricProvidersParams) ResetZoneid() {
	if p.p != nil && p.p["zoneid"] != nil {
		delete(p.p, "zoneid")
	}
}

func (p *ListTungstenFabricProvidersParams) GetZoneid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new ListTungstenFabricProvidersParams instance,
// as then you are sure you have configured all required params
func (s *TungstenService) NewListTungstenFabricProvidersParams() *ListTungstenFabricProvidersParams {
	p := &ListTungstenFabricProvidersParams{}
	p.p = make(map[string]interface{})
	return p
}

// Lists Tungsten-Fabric providers
func (s *TungstenService) ListTungstenFabricProviders(p *ListTungstenFabricProvidersParams) (*ListTungstenFabricProvidersResponse, error) {
	resp, err := s.cs.newRequest("listTungstenFabricProviders", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListTungstenFabricProvidersResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListTungstenFabricProvidersResponse struct {
	Count                   int                       `json:"count"`
	TungstenFabricProviders []*TungstenFabricProvider `json:"tungstenfabricprovider"`
}

type TungstenFabricProvider struct {
	JobID                          string `json:"jobid"`
	Jobstatus                      int    `json:"jobstatus"`
	Name                           string `json:"name"`
	Securitygroupsenabled          bool   `json:"securitygroupsenabled"`
	Tungstengateway                string `json:"tungstengateway"`
	Tungstenproviderhostname       string `json:"tungstenproviderhostname"`
	Tungstenproviderintrospectport string `json:"tungstenproviderintrospectport"`
	Tungstenproviderport           string `json:"tungstenproviderport"`
	Tungstenprovideruuid           string `json:"tungstenprovideruuid"`
	Tungstenprovidervrouterport    string `json:"tungstenprovidervrouterport"`
	Zoneid                         int64  `json:"zoneid"`
	Zonename                       string `json:"zonename"`
}

type ListTungstenFabricServiceGroupParams struct {
	p map[string]interface{}
}

func (p *ListTungstenFabricServiceGroupParams) toURLValues() url.Values {
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
	if v, found := p.p["servicegroupuuid"]; found {
		u.Set("servicegroupuuid", v.(string))
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *ListTungstenFabricServiceGroupParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListTungstenFabricServiceGroupParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
}

func (p *ListTungstenFabricServiceGroupParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListTungstenFabricServiceGroupParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListTungstenFabricServiceGroupParams) ResetPage() {
	if p.p != nil && p.p["page"] != nil {
		delete(p.p, "page")
	}
}

func (p *ListTungstenFabricServiceGroupParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListTungstenFabricServiceGroupParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListTungstenFabricServiceGroupParams) ResetPagesize() {
	if p.p != nil && p.p["pagesize"] != nil {
		delete(p.p, "pagesize")
	}
}

func (p *ListTungstenFabricServiceGroupParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

func (p *ListTungstenFabricServiceGroupParams) SetServicegroupuuid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["servicegroupuuid"] = v
}

func (p *ListTungstenFabricServiceGroupParams) ResetServicegroupuuid() {
	if p.p != nil && p.p["servicegroupuuid"] != nil {
		delete(p.p, "servicegroupuuid")
	}
}

func (p *ListTungstenFabricServiceGroupParams) GetServicegroupuuid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["servicegroupuuid"].(string)
	return value, ok
}

func (p *ListTungstenFabricServiceGroupParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

func (p *ListTungstenFabricServiceGroupParams) ResetZoneid() {
	if p.p != nil && p.p["zoneid"] != nil {
		delete(p.p, "zoneid")
	}
}

func (p *ListTungstenFabricServiceGroupParams) GetZoneid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new ListTungstenFabricServiceGroupParams instance,
// as then you are sure you have configured all required params
func (s *TungstenService) NewListTungstenFabricServiceGroupParams() *ListTungstenFabricServiceGroupParams {
	p := &ListTungstenFabricServiceGroupParams{}
	p.p = make(map[string]interface{})
	return p
}

// list Tungsten-Fabric service group
func (s *TungstenService) ListTungstenFabricServiceGroup(p *ListTungstenFabricServiceGroupParams) (*ListTungstenFabricServiceGroupResponse, error) {
	resp, err := s.cs.newRequest("listTungstenFabricServiceGroup", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListTungstenFabricServiceGroupResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListTungstenFabricServiceGroupResponse struct {
	Count                      int                           `json:"count"`
	TungstenFabricServiceGroup []*TungstenFabricServiceGroup `json:"tungstenfabricservicegroup"`
}

type TungstenFabricServiceGroup struct {
	Endport   int    `json:"endport"`
	JobID     string `json:"jobid"`
	Jobstatus int    `json:"jobstatus"`
	Name      string `json:"name"`
	Protocol  string `json:"protocol"`
	Startport int    `json:"startport"`
	Uuid      string `json:"uuid"`
	Zoneid    int64  `json:"zoneid"`
	Zonename  string `json:"zonename"`
}

type ListTungstenFabricTagParams struct {
	p map[string]interface{}
}

func (p *ListTungstenFabricTagParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["applicationpolicysetuuid"]; found {
		u.Set("applicationpolicysetuuid", v.(string))
	}
	if v, found := p.p["keyword"]; found {
		u.Set("keyword", v.(string))
	}
	if v, found := p.p["networkuuid"]; found {
		u.Set("networkuuid", v.(string))
	}
	if v, found := p.p["nicuuid"]; found {
		u.Set("nicuuid", v.(string))
	}
	if v, found := p.p["page"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("page", vv)
	}
	if v, found := p.p["pagesize"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("pagesize", vv)
	}
	if v, found := p.p["policyuuid"]; found {
		u.Set("policyuuid", v.(string))
	}
	if v, found := p.p["taguuid"]; found {
		u.Set("taguuid", v.(string))
	}
	if v, found := p.p["vmuuid"]; found {
		u.Set("vmuuid", v.(string))
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *ListTungstenFabricTagParams) SetApplicationpolicysetuuid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["applicationpolicysetuuid"] = v
}

func (p *ListTungstenFabricTagParams) ResetApplicationpolicysetuuid() {
	if p.p != nil && p.p["applicationpolicysetuuid"] != nil {
		delete(p.p, "applicationpolicysetuuid")
	}
}

func (p *ListTungstenFabricTagParams) GetApplicationpolicysetuuid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["applicationpolicysetuuid"].(string)
	return value, ok
}

func (p *ListTungstenFabricTagParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListTungstenFabricTagParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
}

func (p *ListTungstenFabricTagParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListTungstenFabricTagParams) SetNetworkuuid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["networkuuid"] = v
}

func (p *ListTungstenFabricTagParams) ResetNetworkuuid() {
	if p.p != nil && p.p["networkuuid"] != nil {
		delete(p.p, "networkuuid")
	}
}

func (p *ListTungstenFabricTagParams) GetNetworkuuid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["networkuuid"].(string)
	return value, ok
}

func (p *ListTungstenFabricTagParams) SetNicuuid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["nicuuid"] = v
}

func (p *ListTungstenFabricTagParams) ResetNicuuid() {
	if p.p != nil && p.p["nicuuid"] != nil {
		delete(p.p, "nicuuid")
	}
}

func (p *ListTungstenFabricTagParams) GetNicuuid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["nicuuid"].(string)
	return value, ok
}

func (p *ListTungstenFabricTagParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListTungstenFabricTagParams) ResetPage() {
	if p.p != nil && p.p["page"] != nil {
		delete(p.p, "page")
	}
}

func (p *ListTungstenFabricTagParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListTungstenFabricTagParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListTungstenFabricTagParams) ResetPagesize() {
	if p.p != nil && p.p["pagesize"] != nil {
		delete(p.p, "pagesize")
	}
}

func (p *ListTungstenFabricTagParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

func (p *ListTungstenFabricTagParams) SetPolicyuuid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["policyuuid"] = v
}

func (p *ListTungstenFabricTagParams) ResetPolicyuuid() {
	if p.p != nil && p.p["policyuuid"] != nil {
		delete(p.p, "policyuuid")
	}
}

func (p *ListTungstenFabricTagParams) GetPolicyuuid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["policyuuid"].(string)
	return value, ok
}

func (p *ListTungstenFabricTagParams) SetTaguuid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["taguuid"] = v
}

func (p *ListTungstenFabricTagParams) ResetTaguuid() {
	if p.p != nil && p.p["taguuid"] != nil {
		delete(p.p, "taguuid")
	}
}

func (p *ListTungstenFabricTagParams) GetTaguuid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["taguuid"].(string)
	return value, ok
}

func (p *ListTungstenFabricTagParams) SetVmuuid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["vmuuid"] = v
}

func (p *ListTungstenFabricTagParams) ResetVmuuid() {
	if p.p != nil && p.p["vmuuid"] != nil {
		delete(p.p, "vmuuid")
	}
}

func (p *ListTungstenFabricTagParams) GetVmuuid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["vmuuid"].(string)
	return value, ok
}

func (p *ListTungstenFabricTagParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

func (p *ListTungstenFabricTagParams) ResetZoneid() {
	if p.p != nil && p.p["zoneid"] != nil {
		delete(p.p, "zoneid")
	}
}

func (p *ListTungstenFabricTagParams) GetZoneid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new ListTungstenFabricTagParams instance,
// as then you are sure you have configured all required params
func (s *TungstenService) NewListTungstenFabricTagParams() *ListTungstenFabricTagParams {
	p := &ListTungstenFabricTagParams{}
	p.p = make(map[string]interface{})
	return p
}

// Lists Tungsten-Fabric tags
func (s *TungstenService) ListTungstenFabricTag(p *ListTungstenFabricTagParams) (*ListTungstenFabricTagResponse, error) {
	resp, err := s.cs.newRequest("listTungstenFabricTag", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListTungstenFabricTagResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListTungstenFabricTagResponse struct {
	Count             int                  `json:"count"`
	TungstenFabricTag []*TungstenFabricTag `json:"tungstenfabrictag"`
}

type TungstenFabricTag struct {
	JobID     string     `json:"jobid"`
	Jobstatus int        `json:"jobstatus"`
	Name      string     `json:"name"`
	Network   []*Network `json:"network"`
	Nic       []string   `json:"nic"`
	Policy    []string   `json:"policy"`
	Uuid      string     `json:"uuid"`
	Vm        []string   `json:"vm"`
	Zoneid    int64      `json:"zoneid"`
	Zonename  string     `json:"zonename"`
}

type ListTungstenFabricTagTypeParams struct {
	p map[string]interface{}
}

func (p *ListTungstenFabricTagTypeParams) toURLValues() url.Values {
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
	if v, found := p.p["tagtypeuuid"]; found {
		u.Set("tagtypeuuid", v.(string))
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *ListTungstenFabricTagTypeParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListTungstenFabricTagTypeParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
}

func (p *ListTungstenFabricTagTypeParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListTungstenFabricTagTypeParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListTungstenFabricTagTypeParams) ResetPage() {
	if p.p != nil && p.p["page"] != nil {
		delete(p.p, "page")
	}
}

func (p *ListTungstenFabricTagTypeParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListTungstenFabricTagTypeParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListTungstenFabricTagTypeParams) ResetPagesize() {
	if p.p != nil && p.p["pagesize"] != nil {
		delete(p.p, "pagesize")
	}
}

func (p *ListTungstenFabricTagTypeParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

func (p *ListTungstenFabricTagTypeParams) SetTagtypeuuid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["tagtypeuuid"] = v
}

func (p *ListTungstenFabricTagTypeParams) ResetTagtypeuuid() {
	if p.p != nil && p.p["tagtypeuuid"] != nil {
		delete(p.p, "tagtypeuuid")
	}
}

func (p *ListTungstenFabricTagTypeParams) GetTagtypeuuid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["tagtypeuuid"].(string)
	return value, ok
}

func (p *ListTungstenFabricTagTypeParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

func (p *ListTungstenFabricTagTypeParams) ResetZoneid() {
	if p.p != nil && p.p["zoneid"] != nil {
		delete(p.p, "zoneid")
	}
}

func (p *ListTungstenFabricTagTypeParams) GetZoneid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new ListTungstenFabricTagTypeParams instance,
// as then you are sure you have configured all required params
func (s *TungstenService) NewListTungstenFabricTagTypeParams() *ListTungstenFabricTagTypeParams {
	p := &ListTungstenFabricTagTypeParams{}
	p.p = make(map[string]interface{})
	return p
}

// Lists Tungsten-Fabric tags
func (s *TungstenService) ListTungstenFabricTagType(p *ListTungstenFabricTagTypeParams) (*ListTungstenFabricTagTypeResponse, error) {
	resp, err := s.cs.newRequest("listTungstenFabricTagType", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListTungstenFabricTagTypeResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListTungstenFabricTagTypeResponse struct {
	Count                 int                      `json:"count"`
	TungstenFabricTagType []*TungstenFabricTagType `json:"tungstenfabrictagtype"`
}

type TungstenFabricTagType struct {
	JobID     string `json:"jobid"`
	Jobstatus int    `json:"jobstatus"`
	Name      string `json:"name"`
	Uuid      string `json:"uuid"`
	Zoneid    int64  `json:"zoneid"`
	Zonename  string `json:"zonename"`
}

type ListTungstenFabricVmParams struct {
	p map[string]interface{}
}

func (p *ListTungstenFabricVmParams) toURLValues() url.Values {
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
	if v, found := p.p["vmuuid"]; found {
		u.Set("vmuuid", v.(string))
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *ListTungstenFabricVmParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
}

func (p *ListTungstenFabricVmParams) ResetKeyword() {
	if p.p != nil && p.p["keyword"] != nil {
		delete(p.p, "keyword")
	}
}

func (p *ListTungstenFabricVmParams) GetKeyword() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["keyword"].(string)
	return value, ok
}

func (p *ListTungstenFabricVmParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
}

func (p *ListTungstenFabricVmParams) ResetPage() {
	if p.p != nil && p.p["page"] != nil {
		delete(p.p, "page")
	}
}

func (p *ListTungstenFabricVmParams) GetPage() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["page"].(int)
	return value, ok
}

func (p *ListTungstenFabricVmParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
}

func (p *ListTungstenFabricVmParams) ResetPagesize() {
	if p.p != nil && p.p["pagesize"] != nil {
		delete(p.p, "pagesize")
	}
}

func (p *ListTungstenFabricVmParams) GetPagesize() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["pagesize"].(int)
	return value, ok
}

func (p *ListTungstenFabricVmParams) SetVmuuid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["vmuuid"] = v
}

func (p *ListTungstenFabricVmParams) ResetVmuuid() {
	if p.p != nil && p.p["vmuuid"] != nil {
		delete(p.p, "vmuuid")
	}
}

func (p *ListTungstenFabricVmParams) GetVmuuid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["vmuuid"].(string)
	return value, ok
}

func (p *ListTungstenFabricVmParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

func (p *ListTungstenFabricVmParams) ResetZoneid() {
	if p.p != nil && p.p["zoneid"] != nil {
		delete(p.p, "zoneid")
	}
}

func (p *ListTungstenFabricVmParams) GetZoneid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new ListTungstenFabricVmParams instance,
// as then you are sure you have configured all required params
func (s *TungstenService) NewListTungstenFabricVmParams() *ListTungstenFabricVmParams {
	p := &ListTungstenFabricVmParams{}
	p.p = make(map[string]interface{})
	return p
}

// list Tungsten-Fabric vm
func (s *TungstenService) ListTungstenFabricVm(p *ListTungstenFabricVmParams) (*ListTungstenFabricVmResponse, error) {
	resp, err := s.cs.newRequest("listTungstenFabricVm", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListTungstenFabricVmResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type ListTungstenFabricVmResponse struct {
	Count            int                 `json:"count"`
	TungstenFabricVm []*TungstenFabricVm `json:"tungstenfabricvm"`
}

type TungstenFabricVm struct {
	JobID     string `json:"jobid"`
	Jobstatus int    `json:"jobstatus"`
	Name      string `json:"name"`
	Uuid      string `json:"uuid"`
	Zoneid    int64  `json:"zoneid"`
	Zonename  string `json:"zonename"`
}

type RemoveTungstenFabricPolicyParams struct {
	p map[string]interface{}
}

func (p *RemoveTungstenFabricPolicyParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["networkuuid"]; found {
		u.Set("networkuuid", v.(string))
	}
	if v, found := p.p["policyuuid"]; found {
		u.Set("policyuuid", v.(string))
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *RemoveTungstenFabricPolicyParams) SetNetworkuuid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["networkuuid"] = v
}

func (p *RemoveTungstenFabricPolicyParams) ResetNetworkuuid() {
	if p.p != nil && p.p["networkuuid"] != nil {
		delete(p.p, "networkuuid")
	}
}

func (p *RemoveTungstenFabricPolicyParams) GetNetworkuuid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["networkuuid"].(string)
	return value, ok
}

func (p *RemoveTungstenFabricPolicyParams) SetPolicyuuid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["policyuuid"] = v
}

func (p *RemoveTungstenFabricPolicyParams) ResetPolicyuuid() {
	if p.p != nil && p.p["policyuuid"] != nil {
		delete(p.p, "policyuuid")
	}
}

func (p *RemoveTungstenFabricPolicyParams) GetPolicyuuid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["policyuuid"].(string)
	return value, ok
}

func (p *RemoveTungstenFabricPolicyParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

func (p *RemoveTungstenFabricPolicyParams) ResetZoneid() {
	if p.p != nil && p.p["zoneid"] != nil {
		delete(p.p, "zoneid")
	}
}

func (p *RemoveTungstenFabricPolicyParams) GetZoneid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new RemoveTungstenFabricPolicyParams instance,
// as then you are sure you have configured all required params
func (s *TungstenService) NewRemoveTungstenFabricPolicyParams(networkuuid string, policyuuid string, zoneid string) *RemoveTungstenFabricPolicyParams {
	p := &RemoveTungstenFabricPolicyParams{}
	p.p = make(map[string]interface{})
	p.p["networkuuid"] = networkuuid
	p.p["policyuuid"] = policyuuid
	p.p["zoneid"] = zoneid
	return p
}

// remove Tungsten-Fabric policy
func (s *TungstenService) RemoveTungstenFabricPolicy(p *RemoveTungstenFabricPolicyParams) (*RemoveTungstenFabricPolicyResponse, error) {
	resp, err := s.cs.newRequest("removeTungstenFabricPolicy", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r RemoveTungstenFabricPolicyResponse
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

type RemoveTungstenFabricPolicyResponse struct {
	JobID     string     `json:"jobid"`
	Jobstatus int        `json:"jobstatus"`
	Name      string     `json:"name"`
	Network   []*Network `json:"network"`
	Uuid      string     `json:"uuid"`
	Zoneid    int64      `json:"zoneid"`
	Zonename  string     `json:"zonename"`
}

type RemoveTungstenFabricPolicyRuleParams struct {
	p map[string]interface{}
}

func (p *RemoveTungstenFabricPolicyRuleParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["policyuuid"]; found {
		u.Set("policyuuid", v.(string))
	}
	if v, found := p.p["ruleuuid"]; found {
		u.Set("ruleuuid", v.(string))
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *RemoveTungstenFabricPolicyRuleParams) SetPolicyuuid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["policyuuid"] = v
}

func (p *RemoveTungstenFabricPolicyRuleParams) ResetPolicyuuid() {
	if p.p != nil && p.p["policyuuid"] != nil {
		delete(p.p, "policyuuid")
	}
}

func (p *RemoveTungstenFabricPolicyRuleParams) GetPolicyuuid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["policyuuid"].(string)
	return value, ok
}

func (p *RemoveTungstenFabricPolicyRuleParams) SetRuleuuid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ruleuuid"] = v
}

func (p *RemoveTungstenFabricPolicyRuleParams) ResetRuleuuid() {
	if p.p != nil && p.p["ruleuuid"] != nil {
		delete(p.p, "ruleuuid")
	}
}

func (p *RemoveTungstenFabricPolicyRuleParams) GetRuleuuid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["ruleuuid"].(string)
	return value, ok
}

func (p *RemoveTungstenFabricPolicyRuleParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

func (p *RemoveTungstenFabricPolicyRuleParams) ResetZoneid() {
	if p.p != nil && p.p["zoneid"] != nil {
		delete(p.p, "zoneid")
	}
}

func (p *RemoveTungstenFabricPolicyRuleParams) GetZoneid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new RemoveTungstenFabricPolicyRuleParams instance,
// as then you are sure you have configured all required params
func (s *TungstenService) NewRemoveTungstenFabricPolicyRuleParams(policyuuid string, ruleuuid string, zoneid string) *RemoveTungstenFabricPolicyRuleParams {
	p := &RemoveTungstenFabricPolicyRuleParams{}
	p.p = make(map[string]interface{})
	p.p["policyuuid"] = policyuuid
	p.p["ruleuuid"] = ruleuuid
	p.p["zoneid"] = zoneid
	return p
}

// remove Tungsten-Fabric policy
func (s *TungstenService) RemoveTungstenFabricPolicyRule(p *RemoveTungstenFabricPolicyRuleParams) (*RemoveTungstenFabricPolicyRuleResponse, error) {
	resp, err := s.cs.newRequest("removeTungstenFabricPolicyRule", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r RemoveTungstenFabricPolicyRuleResponse
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

type RemoveTungstenFabricPolicyRuleResponse struct {
	JobID     string     `json:"jobid"`
	Jobstatus int        `json:"jobstatus"`
	Name      string     `json:"name"`
	Network   []*Network `json:"network"`
	Uuid      string     `json:"uuid"`
	Zoneid    int64      `json:"zoneid"`
	Zonename  string     `json:"zonename"`
}

type RemoveTungstenFabricTagParams struct {
	p map[string]interface{}
}

func (p *RemoveTungstenFabricTagParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["applicationpolicysetuuid"]; found {
		u.Set("applicationpolicysetuuid", v.(string))
	}
	if v, found := p.p["networkuuid"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("networkuuid", vv)
	}
	if v, found := p.p["nicuuid"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("nicuuid", vv)
	}
	if v, found := p.p["policyuuid"]; found {
		u.Set("policyuuid", v.(string))
	}
	if v, found := p.p["taguuid"]; found {
		u.Set("taguuid", v.(string))
	}
	if v, found := p.p["vmuuid"]; found {
		vv := strings.Join(v.([]string), ",")
		u.Set("vmuuid", vv)
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *RemoveTungstenFabricTagParams) SetApplicationpolicysetuuid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["applicationpolicysetuuid"] = v
}

func (p *RemoveTungstenFabricTagParams) ResetApplicationpolicysetuuid() {
	if p.p != nil && p.p["applicationpolicysetuuid"] != nil {
		delete(p.p, "applicationpolicysetuuid")
	}
}

func (p *RemoveTungstenFabricTagParams) GetApplicationpolicysetuuid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["applicationpolicysetuuid"].(string)
	return value, ok
}

func (p *RemoveTungstenFabricTagParams) SetNetworkuuid(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["networkuuid"] = v
}

func (p *RemoveTungstenFabricTagParams) ResetNetworkuuid() {
	if p.p != nil && p.p["networkuuid"] != nil {
		delete(p.p, "networkuuid")
	}
}

func (p *RemoveTungstenFabricTagParams) GetNetworkuuid() ([]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["networkuuid"].([]string)
	return value, ok
}

func (p *RemoveTungstenFabricTagParams) SetNicuuid(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["nicuuid"] = v
}

func (p *RemoveTungstenFabricTagParams) ResetNicuuid() {
	if p.p != nil && p.p["nicuuid"] != nil {
		delete(p.p, "nicuuid")
	}
}

func (p *RemoveTungstenFabricTagParams) GetNicuuid() ([]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["nicuuid"].([]string)
	return value, ok
}

func (p *RemoveTungstenFabricTagParams) SetPolicyuuid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["policyuuid"] = v
}

func (p *RemoveTungstenFabricTagParams) ResetPolicyuuid() {
	if p.p != nil && p.p["policyuuid"] != nil {
		delete(p.p, "policyuuid")
	}
}

func (p *RemoveTungstenFabricTagParams) GetPolicyuuid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["policyuuid"].(string)
	return value, ok
}

func (p *RemoveTungstenFabricTagParams) SetTaguuid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["taguuid"] = v
}

func (p *RemoveTungstenFabricTagParams) ResetTaguuid() {
	if p.p != nil && p.p["taguuid"] != nil {
		delete(p.p, "taguuid")
	}
}

func (p *RemoveTungstenFabricTagParams) GetTaguuid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["taguuid"].(string)
	return value, ok
}

func (p *RemoveTungstenFabricTagParams) SetVmuuid(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["vmuuid"] = v
}

func (p *RemoveTungstenFabricTagParams) ResetVmuuid() {
	if p.p != nil && p.p["vmuuid"] != nil {
		delete(p.p, "vmuuid")
	}
}

func (p *RemoveTungstenFabricTagParams) GetVmuuid() ([]string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["vmuuid"].([]string)
	return value, ok
}

func (p *RemoveTungstenFabricTagParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
}

func (p *RemoveTungstenFabricTagParams) ResetZoneid() {
	if p.p != nil && p.p["zoneid"] != nil {
		delete(p.p, "zoneid")
	}
}

func (p *RemoveTungstenFabricTagParams) GetZoneid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["zoneid"].(string)
	return value, ok
}

// You should always use this function to get a new RemoveTungstenFabricTagParams instance,
// as then you are sure you have configured all required params
func (s *TungstenService) NewRemoveTungstenFabricTagParams(taguuid string, zoneid string) *RemoveTungstenFabricTagParams {
	p := &RemoveTungstenFabricTagParams{}
	p.p = make(map[string]interface{})
	p.p["taguuid"] = taguuid
	p.p["zoneid"] = zoneid
	return p
}

// remove Tungsten-Fabric tag
func (s *TungstenService) RemoveTungstenFabricTag(p *RemoveTungstenFabricTagParams) (*RemoveTungstenFabricTagResponse, error) {
	resp, err := s.cs.newRequest("removeTungstenFabricTag", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r RemoveTungstenFabricTagResponse
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

type RemoveTungstenFabricTagResponse struct {
	JobID     string     `json:"jobid"`
	Jobstatus int        `json:"jobstatus"`
	Name      string     `json:"name"`
	Network   []*Network `json:"network"`
	Nic       []string   `json:"nic"`
	Policy    []string   `json:"policy"`
	Uuid      string     `json:"uuid"`
	Vm        []string   `json:"vm"`
	Zoneid    int64      `json:"zoneid"`
	Zonename  string     `json:"zonename"`
}

type SynchronizeTungstenFabricDataParams struct {
	p map[string]interface{}
}

func (p *SynchronizeTungstenFabricDataParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *SynchronizeTungstenFabricDataParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
}

func (p *SynchronizeTungstenFabricDataParams) ResetId() {
	if p.p != nil && p.p["id"] != nil {
		delete(p.p, "id")
	}
}

func (p *SynchronizeTungstenFabricDataParams) GetId() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["id"].(string)
	return value, ok
}

// You should always use this function to get a new SynchronizeTungstenFabricDataParams instance,
// as then you are sure you have configured all required params
func (s *TungstenService) NewSynchronizeTungstenFabricDataParams(id string) *SynchronizeTungstenFabricDataParams {
	p := &SynchronizeTungstenFabricDataParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Synchronize Tungsten-Fabric data
func (s *TungstenService) SynchronizeTungstenFabricData(p *SynchronizeTungstenFabricDataParams) (*SynchronizeTungstenFabricDataResponse, error) {
	resp, err := s.cs.newRequest("synchronizeTungstenFabricData", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r SynchronizeTungstenFabricDataResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

type SynchronizeTungstenFabricDataResponse struct {
	Displaytext string `json:"displaytext"`
	JobID       string `json:"jobid"`
	Jobstatus   int    `json:"jobstatus"`
	Success     bool   `json:"success"`
}

func (r *SynchronizeTungstenFabricDataResponse) UnmarshalJSON(b []byte) error {
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

	type alias SynchronizeTungstenFabricDataResponse
	return json.Unmarshal(b, (*alias)(r))
}

type UpdateTungstenFabricLBHealthMonitorParams struct {
	p map[string]interface{}
}

func (p *UpdateTungstenFabricLBHealthMonitorParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["expectedcode"]; found {
		u.Set("expectedcode", v.(string))
	}
	if v, found := p.p["httpmethodtype"]; found {
		u.Set("httpmethodtype", v.(string))
	}
	if v, found := p.p["interval"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("interval", vv)
	}
	if v, found := p.p["lbruleid"]; found {
		u.Set("lbruleid", v.(string))
	}
	if v, found := p.p["retry"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("retry", vv)
	}
	if v, found := p.p["timeout"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("timeout", vv)
	}
	if v, found := p.p["type"]; found {
		u.Set("type", v.(string))
	}
	if v, found := p.p["urlpath"]; found {
		u.Set("urlpath", v.(string))
	}
	return u
}

func (p *UpdateTungstenFabricLBHealthMonitorParams) SetExpectedcode(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["expectedcode"] = v
}

func (p *UpdateTungstenFabricLBHealthMonitorParams) ResetExpectedcode() {
	if p.p != nil && p.p["expectedcode"] != nil {
		delete(p.p, "expectedcode")
	}
}

func (p *UpdateTungstenFabricLBHealthMonitorParams) GetExpectedcode() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["expectedcode"].(string)
	return value, ok
}

func (p *UpdateTungstenFabricLBHealthMonitorParams) SetHttpmethodtype(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["httpmethodtype"] = v
}

func (p *UpdateTungstenFabricLBHealthMonitorParams) ResetHttpmethodtype() {
	if p.p != nil && p.p["httpmethodtype"] != nil {
		delete(p.p, "httpmethodtype")
	}
}

func (p *UpdateTungstenFabricLBHealthMonitorParams) GetHttpmethodtype() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["httpmethodtype"].(string)
	return value, ok
}

func (p *UpdateTungstenFabricLBHealthMonitorParams) SetInterval(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["interval"] = v
}

func (p *UpdateTungstenFabricLBHealthMonitorParams) ResetInterval() {
	if p.p != nil && p.p["interval"] != nil {
		delete(p.p, "interval")
	}
}

func (p *UpdateTungstenFabricLBHealthMonitorParams) GetInterval() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["interval"].(int)
	return value, ok
}

func (p *UpdateTungstenFabricLBHealthMonitorParams) SetLbruleid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["lbruleid"] = v
}

func (p *UpdateTungstenFabricLBHealthMonitorParams) ResetLbruleid() {
	if p.p != nil && p.p["lbruleid"] != nil {
		delete(p.p, "lbruleid")
	}
}

func (p *UpdateTungstenFabricLBHealthMonitorParams) GetLbruleid() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["lbruleid"].(string)
	return value, ok
}

func (p *UpdateTungstenFabricLBHealthMonitorParams) SetRetry(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["retry"] = v
}

func (p *UpdateTungstenFabricLBHealthMonitorParams) ResetRetry() {
	if p.p != nil && p.p["retry"] != nil {
		delete(p.p, "retry")
	}
}

func (p *UpdateTungstenFabricLBHealthMonitorParams) GetRetry() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["retry"].(int)
	return value, ok
}

func (p *UpdateTungstenFabricLBHealthMonitorParams) SetTimeout(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["timeout"] = v
}

func (p *UpdateTungstenFabricLBHealthMonitorParams) ResetTimeout() {
	if p.p != nil && p.p["timeout"] != nil {
		delete(p.p, "timeout")
	}
}

func (p *UpdateTungstenFabricLBHealthMonitorParams) GetTimeout() (int, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["timeout"].(int)
	return value, ok
}

func (p *UpdateTungstenFabricLBHealthMonitorParams) SetType(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["type"] = v
}

func (p *UpdateTungstenFabricLBHealthMonitorParams) ResetType() {
	if p.p != nil && p.p["type"] != nil {
		delete(p.p, "type")
	}
}

func (p *UpdateTungstenFabricLBHealthMonitorParams) GetType() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["type"].(string)
	return value, ok
}

func (p *UpdateTungstenFabricLBHealthMonitorParams) SetUrlpath(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["urlpath"] = v
}

func (p *UpdateTungstenFabricLBHealthMonitorParams) ResetUrlpath() {
	if p.p != nil && p.p["urlpath"] != nil {
		delete(p.p, "urlpath")
	}
}

func (p *UpdateTungstenFabricLBHealthMonitorParams) GetUrlpath() (string, bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	value, ok := p.p["urlpath"].(string)
	return value, ok
}

// You should always use this function to get a new UpdateTungstenFabricLBHealthMonitorParams instance,
// as then you are sure you have configured all required params
func (s *TungstenService) NewUpdateTungstenFabricLBHealthMonitorParams(interval int, lbruleid string, retry int, timeout int, tungstenType string) *UpdateTungstenFabricLBHealthMonitorParams {
	p := &UpdateTungstenFabricLBHealthMonitorParams{}
	p.p = make(map[string]interface{})
	p.p["interval"] = interval
	p.p["lbruleid"] = lbruleid
	p.p["retry"] = retry
	p.p["timeout"] = timeout
	p.p["type"] = tungstenType
	return p
}

// update Tungsten-Fabric loadbalancer health monitor
func (s *TungstenService) UpdateTungstenFabricLBHealthMonitor(p *UpdateTungstenFabricLBHealthMonitorParams) (*UpdateTungstenFabricLBHealthMonitorResponse, error) {
	resp, err := s.cs.newRequest("updateTungstenFabricLBHealthMonitor", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateTungstenFabricLBHealthMonitorResponse
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

type UpdateTungstenFabricLBHealthMonitorResponse struct {
	Expectedcode string `json:"expectedcode"`
	Httpmethod   string `json:"httpmethod"`
	Id           int64  `json:"id"`
	Interval     int    `json:"interval"`
	JobID        string `json:"jobid"`
	Jobstatus    int    `json:"jobstatus"`
	Lbruleid     string `json:"lbruleid"`
	Retry        int    `json:"retry"`
	Timeout      int    `json:"timeout"`
	Type         string `json:"type"`
	Urlpath      string `json:"urlpath"`
	Uuid         string `json:"uuid"`
	Zoneid       int64  `json:"zoneid"`
	Zonename     string `json:"zonename"`
}
