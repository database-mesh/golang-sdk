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
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

type Object interface {
	SetBucket(bucket string) Object
	SetACL(acl types.ObjectCannedACL) Object
	SetReader(reader io.Reader) Object

	Put(context.Context) error
	Get(context.Context) error
	List(context.Context) error
	Delete(context.Context) error
	Head(context.Context) error
}

type object struct {
	core              *s3.Client
	putObjectParam    *s3.PutObjectInput
	getObjectParam    *s3.GetObjectInput
	listObjectsParam  *s3.ListObjectsInput
	deleteObjectParam *s3.DeleteObjectInput
	headObjectParam   *s3.HeadObjectInput
}

func (s *object) SetBucket(bucket string) Object {
	s.putObjectParam.Bucket = aws.String(bucket)
	s.getObjectParam.Bucket = aws.String(bucket)
	s.listObjectsParam.Bucket = aws.String(bucket)
	s.deleteObjectParam.Bucket = aws.String(bucket)
	s.headObjectParam.Bucket = aws.String(bucket)
	return s
}

// TODO: add support for ACL
func (s *object) SetACL(acl types.ObjectCannedACL) Object {
	s.putObjectParam.ACL = acl
	return s
}

func (s *object) SetReader(reader io.Reader) Object {
	s.putObjectParam.Body = reader
	return s
}

func (s *object) Put(ctx context.Context) error {
	_, err := s.core.PutObject(ctx, s.putObjectParam)
	return err
}

func (s *object) Get(ctx context.Context) error {
	_, err := s.core.GetObject(ctx, s.getObjectParam)
	return err
}

func (s *object) List(ctx context.Context) error {
	//TODO
	_, err := s.core.ListObjects(ctx, s.listObjectsParam)
	return err
}

func (s *object) Delete(ctx context.Context) error {
	_, err := s.core.DeleteObject(ctx, s.deleteObjectParam)
	return err
}

func (s *object) Head(ctx context.Context) error {
	_, err := s.core.HeadObject(ctx, s.headObjectParam)
	return err
}
