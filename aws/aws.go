// Copyright 2022 SphereEx Authors
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

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	awscfg "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
)

type Sessions map[string]aws.Config

type awsCreds struct {
	credentials []credential
}

type credential struct {
	region          string
	accessKeyId       string
	secretAccessKey string
}

func NewSessions() *awsCreds {
	return &awsCreds{credentials: []credential{}}
}

func (s *awsCreds) SetCredential(region, accessKeyId, secretAccessKey string) *awsCreds {
	s.credentials = append(s.credentials, credential{
		region:          region,
		accessKeyId:       accessKeyId,
		secretAccessKey: secretAccessKey,
	})
	return s
}

func (s *awsCreds) Build() Sessions {
	sess := map[string]aws.Config{}
	for _, v := range s.credentials {
		as, err := newAWSSession(v.region, v.accessKeyId, v.secretAccessKey)
		if err != nil {
			continue
		}
		sess[v.region] = as
	}
	return sess
}

func newAWSSession(region, ak, sk string) (aws.Config, error) {
	opts := []func(*awscfg.LoadOptions) error{
		awscfg.WithRegion(region),
	}
	opts = append(opts, awscfg.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
		ak,
		sk,
		"",
	)))
	return awscfg.LoadDefaultConfig(context.Background(), opts...)
}
