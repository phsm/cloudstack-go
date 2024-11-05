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

func TestExternalDeviceService(t *testing.T) {
	service := "ExternalDeviceService"
	response, err := readData(service)
	if err != nil {
		t.Skipf("Skipping test as %v", err)
	}
	server := CreateTestServer(t, response)
	client := cloudstack.NewClient(server.URL, "APIKEY", "SECRETKEY", true)
	defer server.Close()

	testaddCiscoAsa1000vResource := func(t *testing.T) {
		if _, ok := response["addCiscoAsa1000vResource"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.ExternalDevice.NewAddCiscoAsa1000vResourceParams("clusterid", "hostname", "insideportprofile", "physicalnetworkid")
		_, err := client.ExternalDevice.AddCiscoAsa1000vResource(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("AddCiscoAsa1000vResource", testaddCiscoAsa1000vResource)

	testaddCiscoVnmcResource := func(t *testing.T) {
		if _, ok := response["addCiscoVnmcResource"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.ExternalDevice.NewAddCiscoVnmcResourceParams("hostname", "password", "physicalnetworkid", "username")
		_, err := client.ExternalDevice.AddCiscoVnmcResource(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("AddCiscoVnmcResource", testaddCiscoVnmcResource)

	testdeleteCiscoAsa1000vResource := func(t *testing.T) {
		if _, ok := response["deleteCiscoAsa1000vResource"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.ExternalDevice.NewDeleteCiscoAsa1000vResourceParams("resourceid")
		_, err := client.ExternalDevice.DeleteCiscoAsa1000vResource(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("DeleteCiscoAsa1000vResource", testdeleteCiscoAsa1000vResource)

	testdeleteCiscoNexusVSM := func(t *testing.T) {
		if _, ok := response["deleteCiscoNexusVSM"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.ExternalDevice.NewDeleteCiscoNexusVSMParams("id")
		_, err := client.ExternalDevice.DeleteCiscoNexusVSM(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("DeleteCiscoNexusVSM", testdeleteCiscoNexusVSM)

	testdeleteCiscoVnmcResource := func(t *testing.T) {
		if _, ok := response["deleteCiscoVnmcResource"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.ExternalDevice.NewDeleteCiscoVnmcResourceParams("resourceid")
		_, err := client.ExternalDevice.DeleteCiscoVnmcResource(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("DeleteCiscoVnmcResource", testdeleteCiscoVnmcResource)

	testdisableCiscoNexusVSM := func(t *testing.T) {
		if _, ok := response["disableCiscoNexusVSM"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.ExternalDevice.NewDisableCiscoNexusVSMParams("id")
		_, err := client.ExternalDevice.DisableCiscoNexusVSM(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("DisableCiscoNexusVSM", testdisableCiscoNexusVSM)

	testenableCiscoNexusVSM := func(t *testing.T) {
		if _, ok := response["enableCiscoNexusVSM"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.ExternalDevice.NewEnableCiscoNexusVSMParams("id")
		_, err := client.ExternalDevice.EnableCiscoNexusVSM(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("EnableCiscoNexusVSM", testenableCiscoNexusVSM)

	testlistCiscoAsa1000vResources := func(t *testing.T) {
		if _, ok := response["listCiscoAsa1000vResources"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.ExternalDevice.NewListCiscoAsa1000vResourcesParams()
		_, err := client.ExternalDevice.ListCiscoAsa1000vResources(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListCiscoAsa1000vResources", testlistCiscoAsa1000vResources)

	testlistCiscoNexusVSMs := func(t *testing.T) {
		if _, ok := response["listCiscoNexusVSMs"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.ExternalDevice.NewListCiscoNexusVSMsParams()
		_, err := client.ExternalDevice.ListCiscoNexusVSMs(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListCiscoNexusVSMs", testlistCiscoNexusVSMs)

	testlistCiscoVnmcResources := func(t *testing.T) {
		if _, ok := response["listCiscoVnmcResources"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.ExternalDevice.NewListCiscoVnmcResourcesParams()
		_, err := client.ExternalDevice.ListCiscoVnmcResources(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListCiscoVnmcResources", testlistCiscoVnmcResources)

}
