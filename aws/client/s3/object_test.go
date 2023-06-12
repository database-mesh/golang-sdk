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
	"github.com/database-mesh/golang-sdk/aws"
	"github.com/database-mesh/golang-sdk/aws/client/s3"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Object", func() {
	Context("Put Object", func() {
		It("should put object", func() {
			sess := aws.NewSessions().SetCredential(region, accessKey, secretKey).Build()
			object := s3.NewService(sess[region]).Object()
			object.SetBucket("test-for-create-bucket").
				SetKey("test").
				SetValue("test")
			Expect(object.Put(ctx)).To(BeNil())
		})
		Context("Get Object", func() {
			It("should get object", func() {
				sess := aws.NewSessions().SetCredential(region, accessKey, secretKey).Build()
				object := s3.NewService(sess[region]).Object()
				object.SetBucket("test-for-create-bucket").
					SetKey("test")
				data, err := object.Get(ctx)
				Expect(err).To(BeNil())
				Expect(data).To(Equal("test"))
			})
		})
	})
})
