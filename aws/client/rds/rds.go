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
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/rds"
)

type RDS interface {
	Instance() Instance
	Cluster() Cluster
	Aurora() Aurora
}

type service struct {
	instance *rdsInstance
	cluster  *rdsCluster
	aurora   *rdsAurora
}

func (s *service) Instance() Instance {
	return s.instance
}

func (s *service) Cluster() Cluster {
	return s.cluster
}

func (s *service) Aurora() Aurora {
	return s.aurora
}

func NewService(sess aws.Config) *service {
	return &service{
		instance: &rdsInstance{
			core:                     rds.NewFromConfig(sess),
			createInstanceParam:      &rds.CreateDBInstanceInput{},
			deleteInstanceParam:      &rds.DeleteDBInstanceInput{},
			rebootInstanceParam:      &rds.RebootDBInstanceInput{},
			describeInstanceParam:    &rds.DescribeDBInstancesInput{},
			restoreInstancePitrParam: &rds.RestoreDBInstanceToPointInTimeInput{},
			createSnapshotParam:      &rds.CreateDBSnapshotInput{},
			describeSnapshotParam:    &rds.DescribeDBSnapshotsInput{},
			restoreFromSnapshotParam: &rds.RestoreDBInstanceFromDBSnapshotInput{},
		},
		cluster: &rdsCluster{
			core:                              rds.NewFromConfig(sess),
			createClusterParam:                &rds.CreateDBClusterInput{},
			deleteClusterParam:                &rds.DeleteDBClusterInput{},
			failoverClusterParam:              &rds.FailoverDBClusterInput{},
			failoverGlobalClusterParam:        &rds.FailoverGlobalClusterInput{},
			rebootClusterParam:                &rds.RebootDBClusterInput{},
			describeClusterParam:              &rds.DescribeDBClustersInput{},
			restoreDBClusterPitrParam:         &rds.RestoreDBClusterToPointInTimeInput{},
			restoreDBClusterFromSnapshotParam: &rds.RestoreDBClusterFromSnapshotInput{},
			createDBClusterSnapshotParam:      &rds.CreateDBClusterSnapshotInput{},
			describeDBClusterSnapshotParam:    &rds.DescribeDBClusterSnapshotsInput{},
		},
		aurora: &rdsAurora{
			core:                            rds.NewFromConfig(sess),
			createClusterParam:              &rds.CreateDBClusterInput{},
			deleteClusterParam:              &rds.DeleteDBClusterInput{},
			failoverClusterParam:            &rds.FailoverDBClusterInput{},
			failoverGlobalClusterParam:      &rds.FailoverGlobalClusterInput{},
			rebootClusterParam:              &rds.RebootDBClusterInput{},
			describeClusterParam:            &rds.DescribeDBClustersInput{},
			restoreClusterPitrParam:         &rds.RestoreDBClusterToPointInTimeInput{},
			createInstanceParam:             &rds.CreateDBInstanceInput{},
			deleteInstanceParam:             &rds.DeleteDBInstanceInput{},
			rebootInstanceParam:             &rds.RebootDBInstanceInput{},
			describeInstanceParam:           &rds.DescribeDBInstancesInput{},
			restoreInstancePitrParam:        &rds.RestoreDBInstanceToPointInTimeInput{},
			createClusterSnapshotParam:      &rds.CreateDBClusterSnapshotInput{},
			describeClusterSnapshotParam:    &rds.DescribeDBClusterSnapshotsInput{},
			restoreClusterFromSnapshotParam: &rds.RestoreDBClusterFromSnapshotInput{},
		},
	}
}
