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

package rds

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/rds"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
)

type Aurora interface {
	// RDSCluster for Aurora
	SetEngine(engine string) Aurora
	SetEngineVersion(version string) Aurora
	SetDBClusterIdentifier(id string) Aurora
	SetMasterUsername(user string) Aurora
	SetMasterUserPassword(pass string) Aurora
	SetVpcSecurityGroupIds(sgids []string) Aurora
	SetDBSubnetGroup(sbg string) Aurora
	SetSkipFinalSnapshot(enable bool) Aurora
	SetInstanceNumber(num int32) Aurora

	// RDSInstance for Aurora
	SetDBInstanceIdentifier(id string) Aurora
	SetDBInstanceClass(class string) Aurora
	SetPublicAccessible(enable bool) Aurora
	SetDeleteAutomateBackups(enable bool) Aurora

	Create(context.Context) error
	CreateWithPrimary(context.Context) error
	FailoverPrimary(context.Context) error
	FailoverRandomOneReadonlyEndpoint(context.Context) error
	NewReadonlyEndpoint(context.Context) error
	Delete(context.Context) error
	Describe(context.Context) (*DescCluster, error)
}

type rdsAurora struct {
	core *rds.Client

	createClusterParam         *rds.CreateDBClusterInput
	deleteClusterParam         *rds.DeleteDBClusterInput
	failoverClusterParam       *rds.FailoverDBClusterInput
	failoverGlobalClusterParam *rds.FailoverGlobalClusterInput
	rebootClusterParam         *rds.RebootDBClusterInput
	describeClusterParam       *rds.DescribeDBClustersInput
	restoreDBClusterPitrParam  *rds.RestoreDBClusterToPointInTimeInput
	instanceNumber             int32

	createInstanceParam      *rds.CreateDBInstanceInput
	deleteInstanceParam      *rds.DeleteDBInstanceInput
	rebootInstanceParam      *rds.RebootDBInstanceInput
	describeInstanceParam    *rds.DescribeDBInstancesInput
	restoreInstancePitrParam *rds.RestoreDBInstanceToPointInTimeInput
}

var _ Aurora = &rdsAurora{}

func (s *rdsAurora) SetEngine(engine string) Aurora {
	s.createClusterParam.Engine = aws.String(engine)
	s.createInstanceParam.Engine = aws.String(engine)
	return s
}

func (s *rdsAurora) SetEngineVersion(version string) Aurora {
	s.createClusterParam.EngineVersion = aws.String(version)
	return s
}

func (s *rdsAurora) SetDBClusterIdentifier(id string) Aurora {
	s.createClusterParam.DBClusterIdentifier = aws.String(id)
	s.createInstanceParam.DBClusterIdentifier = aws.String(id)
	s.failoverClusterParam.DBClusterIdentifier = aws.String(id)
	s.deleteClusterParam.DBClusterIdentifier = aws.String(id)
	s.describeClusterParam.DBClusterIdentifier = aws.String(id)
	return s
}

func (s *rdsAurora) SetInstanceNumber(num int32) Aurora {
	s.instanceNumber = num
	return s
}

func (s *rdsAurora) SetVpcSecurityGroupIds(ids []string) Aurora {
	s.createClusterParam.VpcSecurityGroupIds = ids
	return s
}

func (s *rdsAurora) SetDBSubnetGroup(sbg string) Aurora {
	s.createClusterParam.DBSubnetGroupName = aws.String(sbg)
	return s
}

func (s *rdsAurora) SetDBInstanceIdentifier(id string) Aurora {
	s.createInstanceParam.DBInstanceIdentifier = aws.String(id)
	s.deleteInstanceParam.DBInstanceIdentifier = aws.String(id)
	return s
}

func (s *rdsAurora) SetDBInstanceClass(class string) Aurora {
	s.createInstanceParam.DBInstanceClass = aws.String(class)
	return s
}

func (s *rdsAurora) SetPublicAccessible(enable bool) Aurora {
	s.createInstanceParam.PubliclyAccessible = aws.Bool(enable)
	return s
}

func (s *rdsAurora) SetMasterUsername(user string) Aurora {
	s.createClusterParam.MasterUsername = aws.String(user)
	return s
}

func (s *rdsAurora) SetMasterUserPassword(pass string) Aurora {
	s.createClusterParam.MasterUserPassword = aws.String(pass)
	return s
}

func (s *rdsAurora) SetSkipFinalSnapshot(enable bool) Aurora {
	s.deleteClusterParam.SkipFinalSnapshot = enable
	s.deleteInstanceParam.SkipFinalSnapshot = enable
	return s
}

func (s *rdsAurora) SetDeleteAutomateBackups(enable bool) Aurora {
	s.deleteInstanceParam.DeleteAutomatedBackups = aws.Bool(enable)
	return s
}

func (s *rdsAurora) Create(ctx context.Context) error {
	if _, err := s.core.CreateDBCluster(ctx, s.createClusterParam); err != nil {
		return err
	}

	for i := 0; i < int(s.instanceNumber); i++ {
		instanceIdentifierName := fmt.Sprintf("%s-instance-%d", *s.createClusterParam.DBClusterIdentifier, i)
		s.SetDBInstanceIdentifier(instanceIdentifierName)
		if _, err := s.core.CreateDBInstance(ctx, s.createInstanceParam); err != nil {
			return err
		}
	}
	return nil
}

func (s *rdsAurora) CreateWithPrimary(ctx context.Context) error {
	if _, err := s.core.CreateDBCluster(ctx, s.createClusterParam); err != nil {
		return err
	}

	if _, err := s.core.CreateDBInstance(ctx, s.createInstanceParam); err != nil {
		return err
	}
	return nil
}

func (s *rdsAurora) NewReadonlyEndpoint(ctx context.Context) error {
	return nil
}

func (s *rdsAurora) FailoverPrimary(ctx context.Context) error {
	_, err := s.core.FailoverDBCluster(ctx, s.failoverClusterParam)
	return err
}

func (s *rdsAurora) FailoverRandomOneReadonlyEndpoint(ctx context.Context) error {
	return nil
}

func (s *rdsAurora) Delete(ctx context.Context) error {
	if s.deleteInstanceParam.SkipFinalSnapshot == false && (s.deleteInstanceParam.FinalDBSnapshotIdentifier == nil) {
		return fmt.Errorf("final snapshot identifier is required when skip final snapshot is false")
	}

	// delete instances of cluster
	var exists = false
	for _, filter := range s.describeClusterParam.Filters {
		if *filter.Name == "db-cluster-id" {
			filter.Values = []string{*s.createClusterParam.DBClusterIdentifier}
			exists = true
			break
		}
	}
	if !exists {
		s.describeInstanceParam.Filters = append(s.describeInstanceParam.Filters, types.Filter{
			Name:   aws.String("db-cluster-id"),
			Values: []string{*s.createClusterParam.DBClusterIdentifier},
		})
	}

	instances, err := s.core.DescribeDBInstances(ctx, s.describeInstanceParam)
	if err != nil {
		return err
	}

	for _, instance := range instances.DBInstances {
		s.deleteInstanceParam.DBInstanceIdentifier = instance.DBInstanceIdentifier
		if _, err := s.core.DeleteDBInstance(ctx, s.deleteInstanceParam); err != nil {
			return err
		}
	}

	// delete cluster
	if _, err := s.core.DeleteDBCluster(ctx, s.deleteClusterParam); err != nil {
		return err
	}

	return nil
}

func (s *rdsAurora) Describe(ctx context.Context) (*DescCluster, error) {
	out, err := s.core.DescribeDBClusters(ctx, s.describeClusterParam)
	// if cluster not found, aws api will return error.
	if err != nil {
		return nil, err
	}

	return convertDBCluster(&out.DBClusters[0]), nil
}
