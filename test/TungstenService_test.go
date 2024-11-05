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

package test

import (
	"testing"

	"github.com/apache/cloudstack-go/v2/cloudstack"
)

func TestTungstenService(t *testing.T) {
	service := "TungstenService"
	response, err := readData(service)
	if err != nil {
		t.Skipf("Skipping test as %v", err)
	}
	server := CreateTestServer(t, response)
	client := cloudstack.NewClient(server.URL, "APIKEY", "SECRETKEY", true)
	defer server.Close()

	testaddTungstenFabricPolicyRule := func(t *testing.T) {
		if _, ok := response["addTungstenFabricPolicyRule"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Tungsten.NewAddTungstenFabricPolicyRuleParams("action", 0, "destipprefix", 0, "destnetwork", 0, "direction", "policyuuid", "protocol", 0, "srcipprefix", 0, "srcnetwork", 0, "zoneid")
		_, err := client.Tungsten.AddTungstenFabricPolicyRule(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("AddTungstenFabricPolicyRule", testaddTungstenFabricPolicyRule)

	testapplyTungstenFabricPolicy := func(t *testing.T) {
		if _, ok := response["applyTungstenFabricPolicy"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Tungsten.NewApplyTungstenFabricPolicyParams(0, 0, "networkuuid", "policyuuid", "zoneid")
		_, err := client.Tungsten.ApplyTungstenFabricPolicy(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ApplyTungstenFabricPolicy", testapplyTungstenFabricPolicy)

	testapplyTungstenFabricTag := func(t *testing.T) {
		if _, ok := response["applyTungstenFabricTag"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Tungsten.NewApplyTungstenFabricTagParams("taguuid", "zoneid")
		_, err := client.Tungsten.ApplyTungstenFabricTag(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ApplyTungstenFabricTag", testapplyTungstenFabricTag)

	testcreateTungstenFabricAddressGroup := func(t *testing.T) {
		if _, ok := response["createTungstenFabricAddressGroup"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Tungsten.NewCreateTungstenFabricAddressGroupParams("ipprefix", 0, "name", "zoneid")
		_, err := client.Tungsten.CreateTungstenFabricAddressGroup(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("CreateTungstenFabricAddressGroup", testcreateTungstenFabricAddressGroup)

	testcreateTungstenFabricApplicationPolicySet := func(t *testing.T) {
		if _, ok := response["createTungstenFabricApplicationPolicySet"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Tungsten.NewCreateTungstenFabricApplicationPolicySetParams("name", "zoneid")
		_, err := client.Tungsten.CreateTungstenFabricApplicationPolicySet(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("CreateTungstenFabricApplicationPolicySet", testcreateTungstenFabricApplicationPolicySet)

	testcreateTungstenFabricServiceGroup := func(t *testing.T) {
		if _, ok := response["createTungstenFabricServiceGroup"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Tungsten.NewCreateTungstenFabricServiceGroupParams(0, "name", "protocol", 0, "zoneid")
		_, err := client.Tungsten.CreateTungstenFabricServiceGroup(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("CreateTungstenFabricServiceGroup", testcreateTungstenFabricServiceGroup)

	testcreateTungstenFabricPolicy := func(t *testing.T) {
		if _, ok := response["createTungstenFabricPolicy"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Tungsten.NewCreateTungstenFabricPolicyParams("name", "zoneid")
		_, err := client.Tungsten.CreateTungstenFabricPolicy(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("CreateTungstenFabricPolicy", testcreateTungstenFabricPolicy)

	testcreateTungstenFabricProvider := func(t *testing.T) {
		if _, ok := response["createTungstenFabricProvider"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Tungsten.NewCreateTungstenFabricProviderParams("name", "tungstengateway", "tungstenproviderhostname", "zoneid")
		_, err := client.Tungsten.CreateTungstenFabricProvider(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("CreateTungstenFabricProvider", testcreateTungstenFabricProvider)

	testcreateTungstenFabricTag := func(t *testing.T) {
		if _, ok := response["createTungstenFabricTag"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Tungsten.NewCreateTungstenFabricTagParams("tagtype", "tagvalue", "zoneid")
		_, err := client.Tungsten.CreateTungstenFabricTag(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("CreateTungstenFabricTag", testcreateTungstenFabricTag)

	testcreateTungstenFabricTagType := func(t *testing.T) {
		if _, ok := response["createTungstenFabricTagType"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Tungsten.NewCreateTungstenFabricTagTypeParams("name", "zoneid")
		_, err := client.Tungsten.CreateTungstenFabricTagType(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("CreateTungstenFabricTagType", testcreateTungstenFabricTagType)

	testconfigTungstenFabricService := func(t *testing.T) {
		if _, ok := response["configTungstenFabricService"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Tungsten.NewConfigTungstenFabricServiceParams("physicalnetworkid", "zoneid")
		_, err := client.Tungsten.ConfigTungstenFabricService(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ConfigTungstenFabricService", testconfigTungstenFabricService)

	testdeleteTungstenFabricAddressGroup := func(t *testing.T) {
		if _, ok := response["deleteTungstenFabricAddressGroup"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Tungsten.NewDeleteTungstenFabricAddressGroupParams("addressgroupuuid", "zoneid")
		_, err := client.Tungsten.DeleteTungstenFabricAddressGroup(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("DeleteTungstenFabricAddressGroup", testdeleteTungstenFabricAddressGroup)

	testdeleteTungstenFabricApplicationPolicySet := func(t *testing.T) {
		if _, ok := response["deleteTungstenFabricApplicationPolicySet"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Tungsten.NewDeleteTungstenFabricApplicationPolicySetParams("applicationpolicysetuuid", "zoneid")
		_, err := client.Tungsten.DeleteTungstenFabricApplicationPolicySet(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("DeleteTungstenFabricApplicationPolicySet", testdeleteTungstenFabricApplicationPolicySet)

	testdeleteTungstenFabricPolicy := func(t *testing.T) {
		if _, ok := response["deleteTungstenFabricPolicy"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Tungsten.NewDeleteTungstenFabricPolicyParams("policyuuid", "zoneid")
		_, err := client.Tungsten.DeleteTungstenFabricPolicy(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("DeleteTungstenFabricPolicy", testdeleteTungstenFabricPolicy)

	testdeleteTungstenFabricTag := func(t *testing.T) {
		if _, ok := response["deleteTungstenFabricTag"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Tungsten.NewDeleteTungstenFabricTagParams("taguuid", "zoneid")
		_, err := client.Tungsten.DeleteTungstenFabricTag(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("DeleteTungstenFabricTag", testdeleteTungstenFabricTag)

	testdeleteTungstenFabricTagType := func(t *testing.T) {
		if _, ok := response["deleteTungstenFabricTagType"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Tungsten.NewDeleteTungstenFabricTagTypeParams("tagtypeuuid", "zoneid")
		_, err := client.Tungsten.DeleteTungstenFabricTagType(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("DeleteTungstenFabricTagType", testdeleteTungstenFabricTagType)

	testdeleteTungstenFabricServiceGroup := func(t *testing.T) {
		if _, ok := response["deleteTungstenFabricServiceGroup"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Tungsten.NewDeleteTungstenFabricServiceGroupParams("servicegroupuuid", "zoneid")
		_, err := client.Tungsten.DeleteTungstenFabricServiceGroup(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("DeleteTungstenFabricServiceGroup", testdeleteTungstenFabricServiceGroup)

	testlistTungstenFabricAddressGroup := func(t *testing.T) {
		if _, ok := response["listTungstenFabricAddressGroup"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Tungsten.NewListTungstenFabricAddressGroupParams()
		_, err := client.Tungsten.ListTungstenFabricAddressGroup(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListTungstenFabricAddressGroup", testlistTungstenFabricAddressGroup)

	testlistTungstenFabricApplicationPolicySet := func(t *testing.T) {
		if _, ok := response["listTungstenFabricApplicationPolicySet"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Tungsten.NewListTungstenFabricApplicationPolicySetParams()
		_, err := client.Tungsten.ListTungstenFabricApplicationPolicySet(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListTungstenFabricApplicationPolicySet", testlistTungstenFabricApplicationPolicySet)

	testlistTungstenFabricLBHealthMonitor := func(t *testing.T) {
		if _, ok := response["listTungstenFabricLBHealthMonitor"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Tungsten.NewListTungstenFabricLBHealthMonitorParams("lbruleid")
		_, err := client.Tungsten.ListTungstenFabricLBHealthMonitor(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListTungstenFabricLBHealthMonitor", testlistTungstenFabricLBHealthMonitor)

	testlistTungstenFabricNic := func(t *testing.T) {
		if _, ok := response["listTungstenFabricNic"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Tungsten.NewListTungstenFabricNicParams()
		_, err := client.Tungsten.ListTungstenFabricNic(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListTungstenFabricNic", testlistTungstenFabricNic)

	testlistTungstenFabricPolicyRule := func(t *testing.T) {
		if _, ok := response["listTungstenFabricPolicyRule"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Tungsten.NewListTungstenFabricPolicyRuleParams("policyuuid")
		_, err := client.Tungsten.ListTungstenFabricPolicyRule(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListTungstenFabricPolicyRule", testlistTungstenFabricPolicyRule)

	testlistTungstenFabricPolicy := func(t *testing.T) {
		if _, ok := response["listTungstenFabricPolicy"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Tungsten.NewListTungstenFabricPolicyParams()
		_, err := client.Tungsten.ListTungstenFabricPolicy(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListTungstenFabricPolicy", testlistTungstenFabricPolicy)

	testlistTungstenFabricProviders := func(t *testing.T) {
		if _, ok := response["listTungstenFabricProviders"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Tungsten.NewListTungstenFabricProvidersParams()
		_, err := client.Tungsten.ListTungstenFabricProviders(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListTungstenFabricProviders", testlistTungstenFabricProviders)

	testlistTungstenFabricServiceGroup := func(t *testing.T) {
		if _, ok := response["listTungstenFabricServiceGroup"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Tungsten.NewListTungstenFabricServiceGroupParams()
		_, err := client.Tungsten.ListTungstenFabricServiceGroup(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListTungstenFabricServiceGroup", testlistTungstenFabricServiceGroup)

	testlistTungstenFabricTag := func(t *testing.T) {
		if _, ok := response["listTungstenFabricTag"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Tungsten.NewListTungstenFabricTagParams()
		_, err := client.Tungsten.ListTungstenFabricTag(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListTungstenFabricTag", testlistTungstenFabricTag)

	testlistTungstenFabricTagType := func(t *testing.T) {
		if _, ok := response["listTungstenFabricTagType"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Tungsten.NewListTungstenFabricTagTypeParams()
		_, err := client.Tungsten.ListTungstenFabricTagType(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListTungstenFabricTagType", testlistTungstenFabricTagType)

	testlistTungstenFabricVm := func(t *testing.T) {
		if _, ok := response["listTungstenFabricVm"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Tungsten.NewListTungstenFabricVmParams()
		_, err := client.Tungsten.ListTungstenFabricVm(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListTungstenFabricVm", testlistTungstenFabricVm)

	testremoveTungstenFabricPolicy := func(t *testing.T) {
		if _, ok := response["removeTungstenFabricPolicy"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Tungsten.NewRemoveTungstenFabricPolicyParams("networkuuid", "policyuuid", "zoneid")
		_, err := client.Tungsten.RemoveTungstenFabricPolicy(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("RemoveTungstenFabricPolicy", testremoveTungstenFabricPolicy)

	testremoveTungstenFabricPolicyRule := func(t *testing.T) {
		if _, ok := response["removeTungstenFabricPolicyRule"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Tungsten.NewRemoveTungstenFabricPolicyRuleParams("policyuuid", "ruleuuid", "zoneid")
		_, err := client.Tungsten.RemoveTungstenFabricPolicyRule(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("RemoveTungstenFabricPolicyRule", testremoveTungstenFabricPolicyRule)

	testremoveTungstenFabricTag := func(t *testing.T) {
		if _, ok := response["removeTungstenFabricTag"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Tungsten.NewRemoveTungstenFabricTagParams("taguuid", "zoneid")
		_, err := client.Tungsten.RemoveTungstenFabricTag(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("RemoveTungstenFabricTag", testremoveTungstenFabricTag)

	testsynchronizeTungstenFabricData := func(t *testing.T) {
		if _, ok := response["synchronizeTungstenFabricData"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Tungsten.NewSynchronizeTungstenFabricDataParams("id")
		_, err := client.Tungsten.SynchronizeTungstenFabricData(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("SynchronizeTungstenFabricData", testsynchronizeTungstenFabricData)

	testupdateTungstenFabricLBHealthMonitor := func(t *testing.T) {
		if _, ok := response["updateTungstenFabricLBHealthMonitor"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Tungsten.NewUpdateTungstenFabricLBHealthMonitorParams(0, "lbruleid", 0, 0, "type")
		_, err := client.Tungsten.UpdateTungstenFabricLBHealthMonitor(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("UpdateTungstenFabricLBHealthMonitor", testupdateTungstenFabricLBHealthMonitor)

}
