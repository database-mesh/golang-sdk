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
	"context"
	"io"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type Bucket interface {
	SetBucket(bucket string) Bucket
	SetUploadBodyReader(reader io.Reader) Bucket

	Create(context.Context) error
	Delete(context.Context) error
	UploadPart(context.Context) error
}

type bucket struct {
	core              *s3.Client
	createBucketParam *s3.CreateBucketInput
	deleteBucketParam *s3.DeleteBucketInput
	uploadPartParam   *s3.UploadPartInput
}

func (s *bucket) SetBucket(bucket string) Bucket {
	s.createBucketParam.Bucket = aws.String(bucket)
	s.deleteBucketParam.Bucket = aws.String(bucket)
	s.uploadPartParam.Bucket = aws.String(bucket)
	return s
}

func (s *bucket) SetUploadBodyReader(r io.Reader) Bucket {
	s.uploadPartParam.Body = r
	return s
}

func (s *bucket) Create(ctx context.Context) error {
	_, err := s.core.CreateBucket(ctx, s.createBucketParam)
	return err
}

func (s *bucket) Delete(ctx context.Context) error {
	_, err := s.core.DeleteBucket(ctx, s.deleteBucketParam)
	return err
}

func (s *bucket) UploadPart(ctx context.Context) error {
	_, err := s.core.UploadPart(ctx, s.uploadPartParam)
	return err
}
