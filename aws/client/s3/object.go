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
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

type Object interface {
	SetBucket(bucket string) Object
	SetKey(key string) Object
	SetValue(value string) Object
	SetACL(acl types.ObjectCannedACL) Object
	SetReader(reader io.Reader) Object
	SetPrefix(prefix string) Object
	SetFolderName(folderName string) Object

	Put(context.Context) error
	Get(context.Context) (string, error)
	List(context.Context) (fileNames []string, err error)
	Delete(context.Context) error
	Head(context.Context) error
	DeleteFolder(context.Context) error
}

type object struct {
	core              *s3.Client
	putObjectParam    *s3.PutObjectInput
	getObjectParam    *s3.GetObjectInput
	listObjectsParam  *s3.ListObjectsInput
	deleteObjectParam *s3.DeleteObjectInput
	headObjectParam   *s3.HeadObjectInput

	folderName string
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

func (s *object) SetKey(key string) Object {
	s.putObjectParam.Key = aws.String(key)
	s.getObjectParam.Key = aws.String(key)
	s.listObjectsParam.Prefix = aws.String(key)
	s.deleteObjectParam.Key = aws.String(key)
	s.headObjectParam.Key = aws.String(key)
	return s
}

func (s *object) SetValue(value string) Object {
	s.putObjectParam.Body = strings.NewReader(value)
	return s
}

func (s *object) SetPrefix(prefix string) Object {
	if !strings.HasSuffix(prefix, "/") {
		prefix += "/"
	}
	s.listObjectsParam.Prefix = aws.String(prefix)
	return s
}

func (s *object) Put(ctx context.Context) error {
	_, err := s.core.PutObject(ctx, s.putObjectParam)
	return err
}
func (s *object) SetFolderName(folderName string) Object {
	if !strings.HasSuffix(folderName, "/") {
		folderName += "/"
	}
	s.folderName = folderName
	return s
}

func (s *object) Get(ctx context.Context) (string, error) {
	obj, err := s.core.GetObject(ctx, s.getObjectParam)
	if err != nil {
		return "", err
	}
	defer obj.Body.Close()
	data, err := io.ReadAll(obj.Body)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (s *object) List(ctx context.Context) (fileNames []string, err error) {
	objs, err := s.core.ListObjects(ctx, s.listObjectsParam)
	if err != nil {
		return nil, err
	}

	prefix := aws.ToString(objs.Prefix)
	for _, obj := range objs.Contents {
		if *obj.Key == prefix {
			continue
		}
		fileNames = append(fileNames, *obj.Key)
	}

	return fileNames, err
}

func (s *object) Delete(ctx context.Context) error {
	_, err := s.core.DeleteObject(ctx, s.deleteObjectParam)
	return err
}

func (s *object) Head(ctx context.Context) error {
	_, err := s.core.HeadObject(ctx, s.headObjectParam)
	return err
}

func (s *object) DeleteFolder(ctx context.Context) error {
	s.SetPrefix(s.folderName)
	fileNames, err := s.List(ctx)
	if err != nil {
		return err
	}
	if len(fileNames) == 0 {
		return nil
	}

	objs := make([]types.ObjectIdentifier, len(fileNames))
	for i, fileName := range fileNames {
		objs[i] = types.ObjectIdentifier{
			Key: aws.String(fileName),
		}
	}

	_, err = s.core.DeleteObjects(ctx, &s3.DeleteObjectsInput{
		Bucket: s.deleteObjectParam.Bucket,
		Delete: &types.Delete{
			Objects: objs,
		},
	})
	return err
}
