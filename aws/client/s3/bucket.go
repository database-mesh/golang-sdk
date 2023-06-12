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
	"errors"
	"io"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/aws/smithy-go"
)

type Bucket interface {
	SetBucket(bucket string) Bucket
	SetUploadBodyReader(reader io.Reader) Bucket
	SetBucketLocationConstraint(location string) Bucket

	Create(context.Context) error
	List(ctx context.Context) ([]*DescBucket, error)
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
func (s *bucket) SetBucketLocationConstraint(location string) Bucket {
	s.createBucketParam.CreateBucketConfiguration = &types.CreateBucketConfiguration{
		LocationConstraint: types.BucketLocationConstraint(location),
	}
	return s
}

func (s *bucket) Create(ctx context.Context) error {
	_, err := s.core.CreateBucket(ctx, s.createBucketParam)
	if err != nil {
		if _, ok := errors.Unwrap(err.(*smithy.OperationError).Err).(*types.BucketAlreadyOwnedByYou); ok {
			return nil
		}
		// todo: handle BucketAlreadyExists error
	}
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

// List a function list to list all buckets
func (s *bucket) List(ctx context.Context) ([]*DescBucket, error) {
	buckets, err := s.core.ListBuckets(ctx, &s3.ListBucketsInput{})
	if err != nil {
		return nil, err
	}
	descBuckets := make([]*DescBucket, 0, len(buckets.Buckets))
	for _, b := range buckets.Buckets {
		descBuckets = append(descBuckets, &DescBucket{
			Name:         aws.ToString(b.Name),
			CreationDate: aws.ToTime(b.CreationDate),
		})
	}
	return descBuckets, nil
}

type DescBucket struct {
	Name         string
	CreationDate time.Time
}
