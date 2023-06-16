/*
 * Copyright 2022 SphereEx Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package rds_test

import (
	"encoding/json"
	"fmt"
	"github.com/database-mesh/golang-sdk/aws"
	"github.com/database-mesh/golang-sdk/aws/client/rds"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Test Cluster", func() {
	Context("Test create cluster", func() {
		It("should success", func() {
			if region == "" || accessKey == "" || secretKey == "" {
				Skip("region, accessKey, secretKey are required")
			}
			sess := aws.NewSessions().SetCredential(region, accessKey, secretKey).Build()
			cc := rds.NewService(sess[region]).Cluster()

			cc.SetDBClusterIdentifier("test-cluster-1").
				SetEngine("mysql").
				SetAllocatedStorage(int32(100)).
				SetEngineVersion("8.0.32").
				SetStorageType("io1").
				SetIOPS(1000).
				SetDBClusterInstanceClass("db.m5d.large").
				SetMasterUsername("root").
				SetMasterUserPassword("password").
				SetDatabaseName("test_db")
			err := cc.Create(ctx)
			if err != nil {
				fmt.Println(err.Error())
			}
			Expect(err).To(BeNil())
		})
	})

	Context("Test describe cluster", func() {
		It("should success", func() {
			if region == "" || accessKey == "" || secretKey == "" {
				Skip("region, accessKey, secretKey are required")
			}
			sess := aws.NewSessions().SetCredential(region, accessKey, secretKey).Build()
			cc := rds.NewService(sess[region]).Cluster()

			cc.SetDBClusterIdentifier("test-cluster-1")
			cluster, err := cc.Describe(ctx)
			Expect(err).To(BeNil())
			d, _ := json.MarshalIndent(cluster, "", "  ")
			fmt.Println(string(d))
		})
	})

	Context("Test delete cluster", func() {
		It("should success when skip final snapshot", func() {
			if region == "" || accessKey == "" || secretKey == "" {
				Skip("region, accessKey, secretKey are required")
			}
			sess := aws.NewSessions().SetCredential(region, accessKey, secretKey).Build()
			cc := rds.NewService(sess[region]).Cluster()
			cc.SetDBClusterIdentifier("test-cluster-1").
				SetSkipFinalSnapshot(true)

			err := cc.Delete(ctx)
			if err != nil {
				fmt.Println(err.Error())
			}
			Expect(err).To(BeNil())
		})

		It("should success when set final snapshot", func() {
			if region == "" || accessKey == "" || secretKey == "" {
				Skip("region, accessKey, secretKey are required")
			}
			sess := aws.NewSessions().SetCredential(region, accessKey, secretKey).Build()
			cc := rds.NewService(sess[region]).Cluster()

			cc.SetDBClusterIdentifier("test-cluster-1").
				SetFinalDBSnapshotIdentifier("test-cluster-1-final-snapshot")
			err := cc.Delete(ctx)
			if err != nil {
				fmt.Println(err.Error())
			}
			Expect(err).To(BeNil())
		})
	})
	Context("Test restore cluster", func() {
		It("should success", func() {
			sess := aws.NewSessions().SetCredential(region, accessKey, secretKey).Build()
			cc := rds.NewService(sess[region]).Cluster()

			cc.SetDBClusterIdentifier("test-cluster-1").
				SetSnapshotIdentifier("test-cluster-1-final-snapshot").
				SetEngine("mysql")

			err := cc.RestoreFromSnapshot(ctx)
			if err != nil {
				fmt.Println(err.Error())
			}
		})
	})
})
