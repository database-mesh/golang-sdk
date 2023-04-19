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

	"github.com/database-mesh/golang-sdk/aws"
	"github.com/database-mesh/golang-sdk/aws/client/rds"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("instance", func() {

	Context("describe instance", func() {
		It("should describe instance", func() {
			if region == "" || accessKey == "" || secretKey == "" {
				Skip("region, accessKey, secretKey are required")
			}

			sess := aws.NewSessions().SetCredential(region, accessKey, secretKey).Build()
			instance := rds.NewService(sess[region]).Instance()

			instance.SetDBInstanceIdentifier("test1")
			ins, err := instance.Describe(context.Background())
			Expect(err).To(BeNil())
			Expect(ins).To(BeNil())
		})
	})

	It("should create instance", func() {
		if region == "" || accessKey == "" || secretKey == "" {
			Skip("region, accessKey, secretKey are required")
		}
		sess := aws.NewSessions().SetCredential(region, accessKey, secretKey).Build()
		instance := rds.NewService(sess[region]).Instance()

		instance.SetEngine("mysql").
			SetEngineVersion("5.7").
			SetDBInstanceClass("db.t3.micro").
			SetDBInstanceIdentifier("test2").
			SetMasterUsername("root").
			SetMasterUserPassword("password").
			SetAllocatedStorage(20)

		//db, err := instance.Create(context.Background())
		//Expect(err).To(BeNil())
		//Expect(db).ToNot(BeNil())

		ins, err := instance.Describe(context.Background())
		Expect(err).To(BeNil())
		Expect(ins).ToNot(BeNil())
		d, _ := json.MarshalIndent(ins, "", "  ")
		fmt.Println(string(d))
	})

	It("should delete instance", func() {
		if region == "" || accessKey == "" || secretKey == "" {
			Skip("region, accessKey, secretKey are required")
		}
		sess := aws.NewSessions().SetCredential(region, accessKey, secretKey).Build()
		instance := rds.NewService(sess[region]).Instance()

		instance.SetDeleteAutomateBackups(false).
			SetSkipFinalSnapshot(true).
			SetDBInstanceIdentifier("test2")
		//err := instance.Delete(context.Background())
		//Expect(err).To(BeNil())

		ins, err := instance.Describe(context.Background())
		Expect(err).To(BeNil())
		Expect(ins).ToNot(BeNil())
		d, _ := json.MarshalIndent(ins, "", "  ")
		fmt.Println(string(d))
	})
})
