// Copyright 2023 SphereEx Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package s3

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3 interface {
	Object() Object
	Bucket() Bucket
}

type service struct {
	object *object
	bucket *bucket
}

func (s *service) Object() Object {
	return s.object
}

func (s *service) Bucket() Bucket {
	return s.bucket
}

func NewService(sess aws.Config) *service {
	return &service{
		bucket: &bucket{
			core:              s3.NewFromConfig(sess),
			createBucketParam: &s3.CreateBucketInput{},
			deleteBucketParam: &s3.DeleteBucketInput{},
			uploadPartParam:   &s3.UploadPartInput{},
		},
		object: &object{
			core:              s3.NewFromConfig(sess),
			putObjectParam:    &s3.PutObjectInput{},
			getObjectParam:    &s3.GetObjectInput{},
			listObjectsParam:  &s3.ListObjectsInput{},
			deleteObjectParam: &s3.DeleteObjectInput{},
			headObjectParam:   &s3.HeadObjectInput{},
		},
	}
}
