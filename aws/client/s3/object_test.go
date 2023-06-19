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

var _ = Describe("Object", func() {
	Context("Put Object", func() {
		It("should put object", func() {
			sess := aws.NewSessions().SetCredential(region, accessKeyId, secretAccessKey).Build()
			object := s3.NewService(sess[region]).Object()
			object.SetBucket("test-for-create-bucket").
				SetKey("test").
				SetValue("test")
			Expect(object.Put(ctx)).To(BeNil())
		})
		Context("Get Object", func() {
			It("should get object", func() {
				sess := aws.NewSessions().SetCredential(region, accessKeyId, secretAccessKey).Build()
				object := s3.NewService(sess[region]).Object()
				object.SetBucket("test-for-create-bucket").
					SetKey("test")
				data, err := object.Get(ctx)
				Expect(err).To(BeNil())
				Expect(data).To(Equal("test"))
			})
		})
		Context("List Objects", func() {
			It("should list objects", func() {
				sess := aws.NewSessions().SetCredential(region, accessKeyId, secretAccessKey).Build()
				object := s3.NewService(sess[region]).Object()

				object.SetBucket("test-for-create-bucket")
				object.SetPrefix("namespace1-clusterbackupname1-202306121356")
				fileNames, err := object.List(ctx)
				Expect(err).To(BeNil())
				for _, fileName := range fileNames {
					fmt.Println(fileName)
				}
			})
		})

		Context("Delete Folder", func() {
			It("should delete folder", func() {
				sess := aws.NewSessions().SetCredential(region, accessKeyId, secretAccessKey).Build()
				object := s3.NewService(sess[region]).Object()

				object.SetBucket("test-for-create-bucket")
				object.SetFolderName("namespace1-clusterbackupname1-202306121356")
				err = object.DeleteFolder(ctx)
				if err != nil {
					fmt.Println(err.Error())
				}
			})
		})
	})
})
