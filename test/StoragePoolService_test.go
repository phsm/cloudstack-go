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

func TestStoragePoolService(t *testing.T) {
	service := "StoragePoolService"
	response, err := readData(service)
	if err != nil {
		t.Skipf("Skipping test as %v", err)
	}
	server := CreateTestServer(t, response)
	client := cloudstack.NewClient(server.URL, "APIKEY", "SECRETKEY", true)
	defer server.Close()

	testaddObjectStoragePool := func(t *testing.T) {
		if _, ok := response["addObjectStoragePool"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.StoragePool.NewAddObjectStoragePoolParams("name", "provider", "url")
		r, err := client.StoragePool.AddObjectStoragePool(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("AddObjectStoragePool", testaddObjectStoragePool)

	testcancelStorageMaintenance := func(t *testing.T) {
		if _, ok := response["cancelStorageMaintenance"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.StoragePool.NewCancelStorageMaintenanceParams("id")
		r, err := client.StoragePool.CancelStorageMaintenance(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("CancelStorageMaintenance", testcancelStorageMaintenance)

	testchangeStoragePoolScope := func(t *testing.T) {
		if _, ok := response["changeStoragePoolScope"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.StoragePool.NewChangeStoragePoolScopeParams("id", "scope")
		_, err := client.StoragePool.ChangeStoragePoolScope(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ChangeStoragePoolScope", testchangeStoragePoolScope)

	testcreateStoragePool := func(t *testing.T) {
		if _, ok := response["createStoragePool"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.StoragePool.NewCreateStoragePoolParams("name", "url", "zoneid")
		r, err := client.StoragePool.CreateStoragePool(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("CreateStoragePool", testcreateStoragePool)

	testdeleteObjectStoragePool := func(t *testing.T) {
		if _, ok := response["deleteObjectStoragePool"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.StoragePool.NewDeleteObjectStoragePoolParams("id")
		_, err := client.StoragePool.DeleteObjectStoragePool(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("DeleteObjectStoragePool", testdeleteObjectStoragePool)

	testdeleteStoragePool := func(t *testing.T) {
		if _, ok := response["deleteStoragePool"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.StoragePool.NewDeleteStoragePoolParams("id")
		_, err := client.StoragePool.DeleteStoragePool(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("DeleteStoragePool", testdeleteStoragePool)

	testenableStorageMaintenance := func(t *testing.T) {
		if _, ok := response["enableStorageMaintenance"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.StoragePool.NewEnableStorageMaintenanceParams("id")
		r, err := client.StoragePool.EnableStorageMaintenance(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("EnableStorageMaintenance", testenableStorageMaintenance)

	testfindStoragePoolsForMigration := func(t *testing.T) {
		if _, ok := response["findStoragePoolsForMigration"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.StoragePool.NewFindStoragePoolsForMigrationParams("id")
		r, err := client.StoragePool.FindStoragePoolsForMigration(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("FindStoragePoolsForMigration", testfindStoragePoolsForMigration)

	testlistObjectStoragePools := func(t *testing.T) {
		if _, ok := response["listObjectStoragePools"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.StoragePool.NewListObjectStoragePoolsParams()
		_, err := client.StoragePool.ListObjectStoragePools(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListObjectStoragePools", testlistObjectStoragePools)

	testlistStoragePools := func(t *testing.T) {
		if _, ok := response["listStoragePools"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.StoragePool.NewListStoragePoolsParams()
		_, err := client.StoragePool.ListStoragePools(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListStoragePools", testlistStoragePools)

	testlistStoragePoolsMetrics := func(t *testing.T) {
		if _, ok := response["listStoragePoolsMetrics"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.StoragePool.NewListStoragePoolsMetricsParams()
		_, err := client.StoragePool.ListStoragePoolsMetrics(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListStoragePoolsMetrics", testlistStoragePoolsMetrics)

	testlistStoragePoolObjects := func(t *testing.T) {
		if _, ok := response["listStoragePoolObjects"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.StoragePool.NewListStoragePoolObjectsParams("id")
		_, err := client.StoragePool.ListStoragePoolObjects(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListStoragePoolObjects", testlistStoragePoolObjects)

	testlistStorageProviders := func(t *testing.T) {
		if _, ok := response["listStorageProviders"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.StoragePool.NewListStorageProvidersParams("type")
		_, err := client.StoragePool.ListStorageProviders(p)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	t.Run("ListStorageProviders", testlistStorageProviders)

	testsyncStoragePool := func(t *testing.T) {
		if _, ok := response["syncStoragePool"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.StoragePool.NewSyncStoragePoolParams("id")
		r, err := client.StoragePool.SyncStoragePool(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("SyncStoragePool", testsyncStoragePool)

	testupdateObjectStoragePool := func(t *testing.T) {
		if _, ok := response["updateObjectStoragePool"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.StoragePool.NewUpdateObjectStoragePoolParams("id")
		r, err := client.StoragePool.UpdateObjectStoragePool(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("UpdateObjectStoragePool", testupdateObjectStoragePool)

	testupdateStoragePool := func(t *testing.T) {
		if _, ok := response["updateStoragePool"]; !ok {
			t.Skipf("Skipping as no json response is provided in testdata")
		}
		p := client.StoragePool.NewUpdateStoragePoolParams("id")
		r, err := client.StoragePool.UpdateStoragePool(p)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r.Id == "" {
			t.Errorf("Failed to parse response. ID not found")
		}
	}
	t.Run("UpdateStoragePool", testupdateStoragePool)

}
