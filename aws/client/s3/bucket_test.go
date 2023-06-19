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

package s3_test

import (
	"fmt"

	"github.com/database-mesh/golang-sdk/aws"
	"github.com/database-mesh/golang-sdk/aws/client/s3"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Bucket", func() {
	Context("Create Bucket", func() {
		It("should create bucket", func() {
			sess := aws.NewSessions().SetCredential(region, accessKeyId, secretAccessKey).Build()
			bucket := s3.NewService(sess[region]).Bucket()
			bucket.SetBucket("test-for-create-bucket").
				SetBucketLocationConstraint("ap-southeast-1")
			err = bucket.Create(ctx)
			if err != nil {
				fmt.Println(err.Error())
			}
		})

		It("should create bucket fail with already exits error", func() {
			sess := aws.NewSessions().SetCredential(region, accessKeyId, secretAccessKey).Build()
			bucket := s3.NewService(sess[region]).Bucket()
			bucket.SetBucket("test-for-create-bucket").
				SetBucketLocationConstraint("ap-southeast-1")
			err = bucket.Create(ctx)
			if err != nil {
				fmt.Println(err.Error())
			}
		})
	})
	Context("List buckets", func() {
		It("should list buckets", func() {
			sess := aws.NewSessions().SetCredential(region, accessKeyId, secretAccessKey).Build()
			bucket := s3.NewService(sess[region]).Bucket()
			buckets, err := bucket.List(ctx)
			Expect(err).To(BeNil())
			Expect(buckets).ToNot(BeNil())
		})
	})
})
