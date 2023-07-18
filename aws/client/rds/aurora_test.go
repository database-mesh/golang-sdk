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
	"os"
	"time"

	"github.com/database-mesh/golang-sdk/aws"
	"github.com/database-mesh/golang-sdk/aws/client/rds"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Aurora", func() {
	BeforeEach(func() {
		// load env variables if exist
		if v, ok := os.LookupEnv("AWS_REGION"); ok {
			region = v
		}
		if v, ok := os.LookupEnv("AWS_ACCESS_KEY_ID"); ok {
			accessKeyId = v
		}
		if v, ok := os.LookupEnv("AWS_SECRET_ACCESS_KEY"); ok {
			secretAccessKey = v
		}
	})

	It("should be able to describe an aurora cluster", func() {
		sess := aws.NewSessions().SetCredential(region, accessKeyId, secretAccessKey).Build()
		aurora := rds.NewService(sess[region]).Aurora()

		aurora.SetDBClusterIdentifier("test-dataabse-pitr")
		cluster, err := aurora.Describe(context.Background())
		Expect(err).To(BeNil())
		Expect(cluster).ToNot(BeNil())

		b, _ := json.MarshalIndent(cluster, "", "  ")
		fmt.Printf("cluster: %+v\n", string(b))
	})

	It("should create aws aurora with 3 replicas", func() {
		sess := aws.NewSessions().SetCredential(region, accessKeyId, secretAccessKey).Build()
		aurora := rds.NewService(sess[region]).Aurora()

		aurora.SetEngineVersion("5.7").
			SetEngine("aurora-mysql").
			SetDBClusterIdentifier("test-create-aws-aurora-with-replicas3").
			SetMasterUsername("root").
			SetMasterUserPassword("12345678").
			SetDBInstanceClass("db.t3.medium").
			SetInstanceNumber(3)

		Expect(aurora.Create(context.Background())).To(BeNil())
	})

	It("should delete aws aurora with 3 replicas", func() {
		sess := aws.NewSessions().SetCredential(region, accessKeyId, secretAccessKey).Build()
		aurora := rds.NewService(sess[region]).Aurora()

		aurora.SetDBClusterIdentifier("test-create-aws-aurora-with-replicas3").
			SetDeleteAutomateBackups(true).
			SetSkipFinalSnapshot(true)
		Expect(aurora.Delete(context.Background())).To(BeNil())
	})

	It("should create aurora cluster snapshot", func() {
		sess := aws.NewSessions().SetCredential(region, accessKeyId, secretAccessKey).Build()
		aurora := rds.NewService(sess[region]).Aurora()
		aurora.SetSnapshotIdentifier("storagenode-sample-snapshot-20230718063124")
		Expect(aurora.CreateSnapshot(ctx)).To(BeNil())
	})

	It("should get aurora cluster snapshot", func() {
		sess := aws.NewSessions().SetCredential(region, accessKeyId, secretAccessKey).Build()
		aurora := rds.NewService(sess[region]).Aurora()

		aurora.SetSnapshotIdentifier("database-for-console-test-1-instance-1-snapshot")
		snapshot, err := aurora.DescribeSnapshot(ctx)
		Expect(err).To(BeNil())
		Expect(snapshot).ToNot(BeNil())
		d, _ := json.MarshalIndent(snapshot, "", "  ")
		fmt.Printf("snapshot: %+v\n", string(d))
	})

	It("should restore aurora cluster from snapshot", func() {
		sess := aws.NewSessions().SetCredential(region, accessKeyId, secretAccessKey).Build()
		aurora := rds.NewService(sess[region]).Aurora()

		aurora.SetDBClusterIdentifier("test-restore-aws-aurora-from-snapshot").
			SetSnapshotIdentifier("database-for-console-test-1-instance-1-snapshot").
			SetEngine("aurora-mysql").
			SetInstanceNumber(2).
			SetDBInstanceClass("db.t3.medium")

		err := aurora.RestoreFromSnapshot(ctx)

		Expect(err).To(BeNil())
	})

	It("should restore aurora cluster to pitr", func() {
		sess := aws.NewSessions().SetCredential(region, accessKeyId, secretAccessKey).Build()
		aurora := rds.NewService(sess[region]).Aurora()

		restoreTime := "2023-06-28T07:43:29Z"
		t, err := time.Parse(time.RFC3339, restoreTime)
		Expect(err).To(BeNil())
		aurora.SetDBClusterIdentifier("test-restore-aws-aurora-from-pitr").
			SetSourceDBClusterIdentifier("test-dataabse-pitr").
			SetRestoreType(rds.DBClusterRestoreTypeFullCopy).
			SetRestoreToTime(t).
			SetInstanceNumber(1).
			SetDBInstanceClass("db.t3.medium").
			SetEngine("aurora-mysql")

		Expect(aurora.RestoreToPitr(ctx)).To(BeNil())
	})
})
