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
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/database-mesh/golang-sdk/aws"
	"github.com/database-mesh/golang-sdk/aws/client/rds"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var ctx = context.Background()

var _ = Describe("instance", func() {
	Context("describe instance", func() {
		It("should describe instance", func() {
			if region == "" || accessKeyId == "" || secretAccessKey == "" {
				Skip("region, accessKeyId, secretAccessKey are required")
			}
			sess := aws.NewSessions().SetCredential(region, accessKeyId, secretAccessKey).Build()
			instance := rds.NewService(sess[region]).Instance()

			instance.SetDBInstanceIdentifier("test1")
			ins, err := instance.Describe(context.Background())
			Expect(err).To(BeNil())
			Expect(ins).To(BeNil())
		})
	})

	It("should create instance", func() {
		if region == "" || accessKeyId == "" || secretAccessKey == "" {
			Skip("region, accessKeyId, secretAccessKey are required")
		}
		sess := aws.NewSessions().SetCredential(region, accessKeyId, secretAccessKey).Build()
		instance := rds.NewService(sess[region]).Instance()

		instance.SetEngine("mysql").
			SetEngineVersion("5.7").
			SetDBInstanceClass("db.t3.micro").
			SetDBInstanceIdentifier("test-public").
			SetMasterUsername("root").
			SetMasterUserPassword("password").
			SetAllocatedStorage(20).
			SetDBName("test_db").
			SetPublicAccessible(true)

		Expect(instance.Create(ctx)).To(BeNil())
	})

	It("should delete instance", func() {
		if region == "" || accessKeyId == "" || secretAccessKey == "" {
			Skip("region, accessKeyId, secretAccessKey are required")
		}
		sess := aws.NewSessions().SetCredential(region, accessKeyId, secretAccessKey).Build()
		instance := rds.NewService(sess[region]).Instance()

		instance.SetDeleteAutomateBackups(false).
			SetSkipFinalSnapshot(true).
			SetDBInstanceIdentifier("test2")
		err := instance.Delete(ctx)
		Expect(err).To(BeNil())
	})

	It("should create snapshot success", func() {
		if region == "" || accessKeyId == "" || secretAccessKey == "" {
			Skip("region, accessKeyId, secretAccessKey are required")
		}
		sess := aws.NewSessions().SetCredential(region, accessKeyId, secretAccessKey).Build()
		instance := rds.NewService(sess[region]).Instance()

		instance.SetDBInstanceIdentifier("test2")
		instance.SetSnapshotIdentifier(fmt.Sprintf("test2-snapshot-%s", time.Now().Format("20060102150405")))

		Expect(instance.CreateSnapshot(context.Background())).To(BeNil())

		ins, err := instance.Describe(context.Background())
		Expect(err).To(BeNil())
		Expect(ins).ToNot(BeNil())
		d, _ := json.MarshalIndent(ins, "", "  ")
		fmt.Println(string(d))

	})

	It("should get snapshot success", func() {
		if region == "" || accessKeyId == "" || secretAccessKey == "" {
			Skip("region, accessKeyId, secretAccessKey are required")
		}
		sess := aws.NewSessions().SetCredential(region, accessKeyId, secretAccessKey).Build()
		instance := rds.NewService(sess[region]).Instance()

		instance.SetDBInstanceIdentifier("test2")
		instance.SetSnapshotIdentifier("test2-snapshot-20230526163909")
		snapshot, err := instance.DescribeSnapshot(context.Background())
		Expect(err).To(BeNil())
		fmt.Printf("snapshot create time: %s", snapshot.SnapshotCreateTime.Format("20060102150405"))
	})

	It("should describe instances by filter", func() {
		if region == "" || accessKeyId == "" || secretAccessKey == "" {
			Skip("region, accessKeyId, secretAccessKey are required")
		}
		sess := aws.NewSessions().SetCredential(region, accessKeyId, secretAccessKey).Build()
		instance := rds.NewService(sess[region]).Instance()

		instance.SetFilter("db-cluster-id", []string{"database-1-op"})
		resp, err := instance.DescribeAll(ctx)
		Expect(err).To(BeNil())
		Expect(resp).ToNot(BeNil())
		d, _ := json.MarshalIndent(resp, "", "  ")
		fmt.Println(string(d))
	})

	It("should restore instance success", func() {
		sess := aws.NewSessions().SetCredential(region, accessKeyId, secretAccessKey).Build()
		instance := rds.NewService(sess[region]).Instance()

		instance.SetSnapshotIdentifier("test-public-snapshot-20230616").
			SetDBInstanceIdentifier("test-public-restore-3")

		Expect(instance.RestoreFromSnapshot(ctx)).To(BeNil())
	})
})
