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
	Delete() error
	Failover() error
	FailoverGlobal() error
}

type rdsInstance struct {
	core                  *rds.Client
	createInstanceParam   *rds.CreateDBInstanceInput
	deleteInstanceParam   *rds.DeleteDBInstanceInput
	failoverCluster       *rds.FailoverDBClusterInput
	failoverGlobalCluster *rds.FailoverGlobalClusterInput
}

// CreateDBInstanceInput
func (s *rdsInstance) SetEngine(engine string) *rdsInstance {
	s.createInstanceParam.Engine = aws.String(engine)
	return s
}

func (s *rdsInstance) SetEngineVersion(version string) *rdsInstance {
	s.createInstanceParam.EngineVersion = aws.String(version)
	return s
}

func (s *rdsInstance) SetDBInstanceIdentifier(id string) *rdsInstance {
	s.createInstanceParam.DBInstanceIdentifier = aws.String(id)
	s.deleteInstanceParam.DBInstanceIdentifier = aws.String(id)
	return s
}

func (s *rdsInstance) SetMasterUsername(username string) *rdsInstance {
	s.createInstanceParam.MasterUsername = aws.String(username)
	return s
}

func (s *rdsInstance) SetMasterUserPassword(pass string) *rdsInstance {
	s.createInstanceParam.MasterUserPassword = aws.String(pass)
	return s
}

func (s *rdsInstance) SetDBInstanceClass(class string) *rdsInstance {
	s.createInstanceParam.DBInstanceClass = aws.String(class)
	return s
}

func (s *rdsInstance) SetAllocatedStorage(size int32) *rdsInstance {
	s.createInstanceParam.AllocatedStorage = aws.Int32(size)
	return s
}

func (s *rdsInstance) SetDBName(name string) *rdsInstance {
	s.createInstanceParam.DBName = aws.String(name)
	return s
}

func (s *rdsInstance) SetVpcSecurityGroupIds(sgs []string) *rdsInstance {
	s.createInstanceParam.VpcSecurityGroupIds = sgs
	return s
}

func (s *rdsInstance) SetDBSubnetGroup(name string) *rdsInstance {
	s.createInstanceParam.DBSubnetGroupName = aws.String(name)
	return s
}

func (s *rdsInstance) Create(ctx context.Context) error {
	_, err := s.core.CreateDBInstance(ctx, s.createInstanceParam)
	return err
}

// DeleteDBInstanceInput
func (s *rdsInstance) SetDeleteAutomateBackups(enable bool) *rdsInstance {
	s.deleteInstanceParam.DeleteAutomatedBackups = aws.Bool(enable)
	return s
}

func (s *rdsInstance) SetFinalDBSnapshotIdentifier(id string) *rdsInstance {
	s.deleteInstanceParam.DBInstanceIdentifier = aws.String(id)
	return s
}

func (s *rdsInstance) SetSkipFinalSnapshot(skip bool) *rdsInstance {
	s.deleteInstanceParam.SkipFinalSnapshot = skip
	return s
}

func (s *rdsInstance) Delete(ctx context.Context) error {
	_, err := s.core.DeleteDBInstance(ctx, s.deleteInstanceParam)
	return err
}

// FailoverClusterInput
func (s *rdsInstance) SetDBClusterIdentifier(id string) *rdsInstance {
	s.failoverCluster.DBClusterIdentifier = aws.String(id)
	return s
}

func (s *rdsInstance) SetTargetDBInstanceIdentifier(id string) *rdsInstance {
	s.failoverCluster.TargetDBInstanceIdentifier = aws.String(id)
	return s
}

func (s *rdsInstance) Failover(ctx context.Context) error {
	_, err := s.core.FailoverDBCluster(ctx, s.failoverCluster)
	return err
}

// FailoverGlobalClusterInput
func (s *rdsInstance) SetGlobalClusterIdentifier(id string) *rdsInstance {
	s.failoverGlobalCluster.GlobalClusterIdentifier = aws.String(id)
	return s
}

func (s *rdsInstance) SetTargetDbClusterIdentifier(id string) *rdsInstance {
	s.failoverGlobalCluster.TargetDbClusterIdentifier = aws.String(id)
	return s
}

func (s *rdsInstance) FailoverGlobal(ctx context.Context) error {
	_, err := s.core.FailoverGlobalCluster(ctx, s.failoverGlobalCluster)
	return err
}
