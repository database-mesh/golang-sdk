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
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/rds"
)

type Cluster interface {
	SetDBClusterIdentifier(id string) Cluster
	SetTargetDBInstanceIdentifier(id string) Cluster
	SetGlobalClusterIdentifier(id string) Cluster
	SetTargetDbClusterIdentifier(id string) Cluster
	SetEngine(engine string) Cluster
	SetAllocatedStorage(size int32) Cluster
	SetAvailabilityZones(azs []string) Cluster
	SetDBClusterInstanceClass(class string) Cluster
	SetDBSubnetGroupName(name string) Cluster
	SetDatabaseName(name string) Cluster
	SetEngineVersion(version string) Cluster
	SetEngineMode(mode string) Cluster
	SetMasterUsername(username string) Cluster
	SetMasterUserPassword(pass string) Cluster
	SetVpcSecurityGroupIds(sgs []string) Cluster
	SetStorageType(t string) Cluster
	SetIOPS(iops int32) Cluster
	SetSkipFinalSnapshot(skip bool) Cluster
	SetSourceDBClusterIdentifier(sid string) Cluster
	SetBacktraceWindow(w int64) Cluster
	SetRestoreToTime(rt *time.Time) Cluster
	SetRestoreType(t string) Cluster
	SetUseLatestRestorableTime(enable bool) Cluster
	SetPublicAccessible(enable bool) Cluster
	SetSnapshotIdentifier(id string) Cluster

	Failover(context.Context) error
	FailoverGlobal(context.Context) error
	Create(context.Context) error
	Delete(context.Context) error
	Reboot(context.Context) error
	Describe(context.Context) (*DescCluster, error)
	RestorePitr(context.Context) error
	CreateSnapshot(context.Context) error
}

type rdsCluster struct {
	core                         *rds.Client
	createClusterParam           *rds.CreateDBClusterInput
	deleteClusterParam           *rds.DeleteDBClusterInput
	failoverClusterParam         *rds.FailoverDBClusterInput
	failoverGlobalClusterParam   *rds.FailoverGlobalClusterInput
	rebootClusterParam           *rds.RebootDBClusterInput
	describeClusterParam         *rds.DescribeDBClustersInput
	restoreDBClusterPitrParam    *rds.RestoreDBClusterToPointInTimeInput
	createDBClusterSnapshotParam *rds.CreateDBClusterSnapshotInput
}

// FailoverClusterInput
func (s *rdsCluster) SetDBClusterIdentifier(id string) Cluster {
	s.createClusterParam.DBClusterIdentifier = aws.String(id)
	s.deleteClusterParam.DBClusterIdentifier = aws.String(id)
	s.failoverClusterParam.DBClusterIdentifier = aws.String(id)
	s.rebootClusterParam.DBClusterIdentifier = aws.String(id)
	s.describeClusterParam.DBClusterIdentifier = aws.String(id)
	s.restoreDBClusterPitrParam.DBClusterIdentifier = aws.String(id)
	s.createDBClusterSnapshotParam.DBClusterIdentifier = aws.String(id)
	return s
}

func (s *rdsCluster) SetTargetDBInstanceIdentifier(id string) Cluster {
	s.failoverClusterParam.TargetDBInstanceIdentifier = aws.String(id)
	return s
}

func (s *rdsCluster) Failover(ctx context.Context) error {
	_, err := s.core.FailoverDBCluster(ctx, s.failoverClusterParam)
	return err
}

// FailoverGlobalClusterInput
func (s *rdsCluster) SetGlobalClusterIdentifier(id string) Cluster {
	s.failoverGlobalClusterParam.GlobalClusterIdentifier = aws.String(id)
	return s
}

func (s *rdsCluster) SetTargetDbClusterIdentifier(id string) Cluster {
	s.failoverGlobalClusterParam.TargetDbClusterIdentifier = aws.String(id)
	return s
}

func (s *rdsCluster) FailoverGlobal(ctx context.Context) error {
	_, err := s.core.FailoverGlobalCluster(ctx, s.failoverGlobalClusterParam)
	return err
}

// CreateDBClusterInput
func (s *rdsCluster) SetEngine(engine string) Cluster {
	s.createClusterParam.Engine = aws.String(engine)
	return s
}

func (s *rdsCluster) SetAllocatedStorage(size int32) Cluster {
	s.createClusterParam.AllocatedStorage = aws.Int32(size)
	return s
}

func (s *rdsCluster) SetAvailabilityZones(azs []string) Cluster {
	s.createClusterParam.AvailabilityZones = azs
	return s
}

func (s *rdsCluster) SetDBClusterInstanceClass(class string) Cluster {
	s.createClusterParam.DBClusterInstanceClass = aws.String(class)
	s.restoreDBClusterPitrParam.DBClusterInstanceClass = aws.String(class)
	return s
}

func (s *rdsCluster) SetDBSubnetGroupName(name string) Cluster {
	s.createClusterParam.DBSubnetGroupName = aws.String(name)
	s.restoreDBClusterPitrParam.DBSubnetGroupName = aws.String(name)
	return s
}

func (s *rdsCluster) SetDatabaseName(name string) Cluster {
	s.createClusterParam.DatabaseName = aws.String(name)
	return s
}

func (s *rdsCluster) SetEngineVersion(version string) Cluster {
	s.createClusterParam.EngineVersion = aws.String(version)
	return s
}

func (s *rdsCluster) SetEngineMode(mode string) Cluster {
	s.createClusterParam.EngineMode = aws.String(mode)
	return s
}

func (s *rdsCluster) SetMasterUsername(username string) Cluster {
	s.createClusterParam.MasterUsername = aws.String(username)
	return s
}

func (s *rdsCluster) SetMasterUserPassword(pass string) Cluster {
	s.createClusterParam.MasterUserPassword = aws.String(pass)
	return s
}

func (s *rdsCluster) SetVpcSecurityGroupIds(sgs []string) Cluster {
	s.createClusterParam.VpcSecurityGroupIds = sgs
	return s
}

func (s *rdsCluster) SetStorageType(t string) Cluster {
	s.createClusterParam.StorageType = aws.String(t)
	return s
}

func (s *rdsCluster) SetIOPS(ps int32) Cluster {
	s.createClusterParam.Iops = aws.Int32(ps)
	s.restoreDBClusterPitrParam.Iops = aws.Int32(ps)
	return s
}

func (s *rdsCluster) Create(ctx context.Context) error {
	_, err := s.core.CreateDBCluster(ctx, s.createClusterParam)
	return err
}

// DeleteDBClusterInput
func (s *rdsCluster) SetSkipFinalSnapshot(skip bool) Cluster {
	s.deleteClusterParam.SkipFinalSnapshot = skip
	return s
}

func (s *rdsCluster) SetPublicAccessible(enable bool) Cluster {
	s.createClusterParam.PubliclyAccessible = aws.Bool(enable)
	return s
}

func (s *rdsCluster) Delete(ctx context.Context) error {
	_, err := s.core.DeleteDBCluster(ctx, s.deleteClusterParam)
	return err
}

// RebootDBClusterInput
func (s *rdsCluster) Reboot(ctx context.Context) error {
	_, err := s.core.RebootDBCluster(ctx, s.rebootClusterParam)
	return err
}

func (s *rdsCluster) SetSourceDBClusterIdentifier(sid string) Cluster {
	s.restoreDBClusterPitrParam.SourceDBClusterIdentifier = aws.String(sid)
	return s
}

func (s *rdsCluster) SetBacktraceWindow(w int64) Cluster {
	s.restoreDBClusterPitrParam.BacktrackWindow = aws.Int64(w)
	return s
}

func (s *rdsCluster) SetRestoreToTime(rt *time.Time) Cluster {
	s.restoreDBClusterPitrParam.RestoreToTime = rt
	return s
}

func (s *rdsCluster) SetRestoreType(t string) Cluster {
	s.restoreDBClusterPitrParam.RestoreType = aws.String(t)
	return s
}

func (s *rdsCluster) SetUseLatestRestorableTime(enable bool) Cluster {
	s.restoreDBClusterPitrParam.UseLatestRestorableTime = enable
	return s
}

func (s *rdsCluster) RestorePitr(ctx context.Context) error {
	_, err := s.core.RestoreDBClusterToPointInTime(ctx, s.restoreDBClusterPitrParam)
	return err
}

func (s *rdsCluster) SetSnapshotIdentifier(id string) Cluster {
	s.createDBClusterSnapshotParam.DBClusterSnapshotIdentifier = aws.String(id)
	return s
}

func (s *rdsCluster) CreateSnapshot(ctx context.Context) error {
	_, err := s.core.CreateDBClusterSnapshot(ctx, s.createDBClusterSnapshotParam)
	return err
}

type DescCluster struct {
	CharSetName                 string
	ClusterCreateTime           time.Time
	AvailabilityZones           []string
	CustomEndpoints             []string
	DBClusterArn                string
	DBClusterIdentifier         string
	DBClusterMembers            []ClusterMember
	DBClusterParamterGroup      string
	DeletionProtection          bool
	PrimaryEndpoint             string
	ReadReplicaIdentifiers      []string
	ReaderEndpoint              string
	ReplicationSourceIdentifier string
	Status                      string
	Port                        int32
}

type ClusterMember struct {
	DBClusterParameterGroupStatus string
	DBInstanceIdentifier          string
	IsClusterWrite                bool
}

func (s *rdsCluster) Describe(ctx context.Context) (*DescCluster, error) {
	output, err := s.core.DescribeDBClusters(ctx, s.describeClusterParam)
	if err != nil {
		return nil, err
	}
	desc := &DescCluster{}
	if len(output.DBClusters) > 0 {
		desc.AvailabilityZones = output.DBClusters[0].AvailabilityZones
		desc.CharSetName = aws.ToString(output.DBClusters[0].CharacterSetName)
		desc.ClusterCreateTime = aws.ToTime(output.DBClusters[0].ClusterCreateTime)
		desc.CustomEndpoints = output.DBClusters[0].CustomEndpoints
		desc.DBClusterArn = aws.ToString(output.DBClusters[0].DBClusterArn)
		desc.DBClusterIdentifier = aws.ToString(output.DBClusters[0].DBClusterIdentifier)
		for _, m := range output.DBClusters[0].DBClusterMembers {
			desc.DBClusterMembers = append(desc.DBClusterMembers, ClusterMember{
				DBClusterParameterGroupStatus: aws.ToString(m.DBClusterParameterGroupStatus),
				DBInstanceIdentifier:          aws.ToString(m.DBInstanceIdentifier),
				IsClusterWrite:                m.IsClusterWriter,
			})
		}
		desc.DBClusterParamterGroup = aws.ToString(output.DBClusters[0].DBClusterParameterGroup)
		desc.DeletionProtection = aws.ToBool(output.DBClusters[0].DeletionProtection)
		desc.PrimaryEndpoint = aws.ToString(output.DBClusters[0].Endpoint)
		desc.ReadReplicaIdentifiers = output.DBClusters[0].ReadReplicaIdentifiers
		desc.ReaderEndpoint = aws.ToString(output.DBClusters[0].ReaderEndpoint)
		desc.ReplicationSourceIdentifier = aws.ToString(output.DBClusters[0].ReplicationSourceIdentifier)
		desc.Port = aws.ToInt32(output.DBClusters[0].Port)
		desc.Status = aws.ToString(output.DBClusters[0].Status)
	}
	return desc, nil
}
