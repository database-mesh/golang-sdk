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

package rds

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/rds"
)

type RDS interface {
	Instance(context.Context) Instance
}

type service struct {
	instance *rdsInstance
}

func (s *service) Instance() *rdsInstance {
	return s.instance
}

func NewService(sess aws.Config) *service {
	return &service{
		instance: &rdsInstance{
			core:  rds.NewFromConfig(sess),
			param: &rds.CreateDBInstanceInput{},
		},
	}
}

type Instance interface {
	Create() error
}

type rdsInstance struct {
	core  *rds.Client
	param *rds.CreateDBInstanceInput
}

func (s *rdsInstance) SetEngine(engine string) *rdsInstance {
	s.param.Engine = aws.String(engine)
	return s
}

func (s *rdsInstance) SetEngineVersion(version string) *rdsInstance {
	s.param.EngineVersion = aws.String(version)
	return s
}

func (s *rdsInstance) SetDBInstanceIdentifier(id string) *rdsInstance {
	s.param.DBInstanceIdentifier = aws.String(id)
	return s
}

func (s *rdsInstance) SetMasterUsername(username string) *rdsInstance {
	s.param.MasterUsername = aws.String(username)
	return s
}

func (s *rdsInstance) SetMasterUserPassword(pass string) *rdsInstance {
	s.param.MasterUserPassword = aws.String(pass)
	return s
}

func (s *rdsInstance) SetDBInstanceClass(class string) *rdsInstance {
	s.param.DBInstanceClass = aws.String(class)
	return s
}

func (s *rdsInstance) SetAllocatedStorage(size int32) *rdsInstance {
	s.param.AllocatedStorage = aws.Int32(size)
	return s
}

func (s *rdsInstance) SetDBName(name string) *rdsInstance {
	s.param.DBName = aws.String(name)
	return s
}

func (s *rdsInstance) SetVpcSecurityGroupIds(sgs []string) *rdsInstance {
	s.param.VpcSecurityGroupIds = sgs
	return s
}

func (s *rdsInstance) SetDBSubnetGroup(name string) *rdsInstance {
	s.param.DBSubnetGroupName = aws.String(name)
	return s
}

func (s *rdsInstance) Create(ctx context.Context) error {
	_, err := s.core.CreateDBInstance(ctx, s.param)
	return err
}
