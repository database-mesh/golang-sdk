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
			accessKey = v
		}
		if v, ok := os.LookupEnv("AWS_SECRET_ACCESS_KEY"); ok {
			secretKey = v
		}
	})

	It("should be able to describe an aurora cluster", func() {
		sess := aws.NewSessions().SetCredential(region, accessKey, secretKey).Build()
		aurora := rds.NewService(sess[region]).Aurora()

		aurora.SetEngineVersion("5.7").
			SetEngine("aurora-mysql").
			SetDBClusterIdentifier("test").
			SetMasterUsername("root").
			SetMasterUserPassword("12345678")
		//Expect(aurora.Create(context.Background())).To(BeNil())
		aurora.SetDBClusterIdentifier("test")
		cluster, err := aurora.Describe(context.Background())
		Expect(err).To(BeNil())
		Expect(cluster).ToNot(BeNil())

		b, _ := json.MarshalIndent(cluster, "", "  ")
		fmt.Printf("cluster: %+v\n", string(b))

		Expect(aurora.Delete(context.Background())).To(BeNil())
	})

	It("should create aws aurora with 3 replicas", func() {
		sess := aws.NewSessions().SetCredential(region, accessKey, secretKey).Build()
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
		sess := aws.NewSessions().SetCredential(region, accessKey, secretKey).Build()
		aurora := rds.NewService(sess[region]).Aurora()

		aurora.SetDBClusterIdentifier("test-create-aws-aurora-with-replicas3").
			SetDeleteAutomateBackups(true).
			SetSkipFinalSnapshot(true)
		Expect(aurora.Delete(context.Background())).To(BeNil())
	})
})
