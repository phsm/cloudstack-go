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

func TestRouterService(t *testing.T) {
	service := "RouterService"
	response, err := readData(service)
	if err != nil {
		t.Skipf("Skipping test as %v", err)
	}
	server := CreateTestServer(t, response)
	client := cloudstack.NewClient(server.URL, "APIKEY", "SECRETKEY", true)
	defer server.Close()

	testaddTungstenFabricNetworkGatewayToLogicalRouter := func(t *testing.T) {
		if _, ok := response["addTungstenFabricNetworkGatewayToLogicalRouter"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Router.NewAddTungstenFabricNetworkGatewayToLogicalRouterParams("logicalrouteruuid", "networkuuid", "zoneid")
		_, err := client.Router.AddTungstenFabricNetworkGatewayToLogicalRouter(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("AddTungstenFabricNetworkGatewayToLogicalRouter", testaddTungstenFabricNetworkGatewayToLogicalRouter)

	testchangeServiceForRouter := func(t *testing.T) {
		if _, ok := response["changeServiceForRouter"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Router.NewChangeServiceForRouterParams("id", "serviceofferingid")
		r, err := client.Router.ChangeServiceForRouter(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("ChangeServiceForRouter", testchangeServiceForRouter)

	testconfigureVirtualRouterElement := func(t *testing.T) {
		if _, ok := response["configureVirtualRouterElement"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Router.NewConfigureVirtualRouterElementParams(true, "id")
		r, err := client.Router.ConfigureVirtualRouterElement(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("ConfigureVirtualRouterElement", testconfigureVirtualRouterElement)

	testcreateTungstenFabricLogicalRouter := func(t *testing.T) {
		if _, ok := response["createTungstenFabricLogicalRouter"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Router.NewCreateTungstenFabricLogicalRouterParams("name", "zoneid")
		_, err := client.Router.CreateTungstenFabricLogicalRouter(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("CreateTungstenFabricLogicalRouter", testcreateTungstenFabricLogicalRouter)

	testcreateVirtualRouterElement := func(t *testing.T) {
		if _, ok := response["createVirtualRouterElement"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Router.NewCreateVirtualRouterElementParams("nspid")
		r, err := client.Router.CreateVirtualRouterElement(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("CreateVirtualRouterElement", testcreateVirtualRouterElement)

	testdeleteTungstenFabricLogicalRouter := func(t *testing.T) {
		if _, ok := response["deleteTungstenFabricLogicalRouter"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Router.NewDeleteTungstenFabricLogicalRouterParams("logicalrouteruuid", "zoneid")
		_, err := client.Router.DeleteTungstenFabricLogicalRouter(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("DeleteTungstenFabricLogicalRouter", testdeleteTungstenFabricLogicalRouter)

	testdestroyRouter := func(t *testing.T) {
		if _, ok := response["destroyRouter"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Router.NewDestroyRouterParams("id")
		r, err := client.Router.DestroyRouter(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("DestroyRouter", testdestroyRouter)

	testgetRouterHealthCheckResults := func(t *testing.T) {
		if _, ok := response["getRouterHealthCheckResults"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Router.NewGetRouterHealthCheckResultsParams("routerid")
		_, err := client.Router.GetRouterHealthCheckResults(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("GetRouterHealthCheckResults", testgetRouterHealthCheckResults)

	testlistRouters := func(t *testing.T) {
		if _, ok := response["listRouters"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Router.NewListRoutersParams()
		_, err := client.Router.ListRouters(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListRouters", testlistRouters)

	testlistTungstenFabricLogicalRouter := func(t *testing.T) {
		if _, ok := response["listTungstenFabricLogicalRouter"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Router.NewListTungstenFabricLogicalRouterParams()
		_, err := client.Router.ListTungstenFabricLogicalRouter(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListTungstenFabricLogicalRouter", testlistTungstenFabricLogicalRouter)

	testlistVirtualRouterElements := func(t *testing.T) {
		if _, ok := response["listVirtualRouterElements"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Router.NewListVirtualRouterElementsParams()
		_, err := client.Router.ListVirtualRouterElements(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListVirtualRouterElements", testlistVirtualRouterElements)

	testrebootRouter := func(t *testing.T) {
		if _, ok := response["rebootRouter"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Router.NewRebootRouterParams("id")
		r, err := client.Router.RebootRouter(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("RebootRouter", testrebootRouter)

	testremoveTungstenFabricNetworkGatewayFromLogicalRouter := func(t *testing.T) {
		if _, ok := response["removeTungstenFabricNetworkGatewayFromLogicalRouter"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Router.NewRemoveTungstenFabricNetworkGatewayFromLogicalRouterParams("logicalrouteruuid", "networkuuid", "zoneid")
		_, err := client.Router.RemoveTungstenFabricNetworkGatewayFromLogicalRouter(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("RemoveTungstenFabricNetworkGatewayFromLogicalRouter", testremoveTungstenFabricNetworkGatewayFromLogicalRouter)

	teststartRouter := func(t *testing.T) {
		if _, ok := response["startRouter"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Router.NewStartRouterParams("id")
		r, err := client.Router.StartRouter(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("StartRouter", teststartRouter)

	teststopRouter := func(t *testing.T) {
		if _, ok := response["stopRouter"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.Router.NewStopRouterParams("id")
		r, err := client.Router.StopRouter(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("StopRouter", teststopRouter)

}
