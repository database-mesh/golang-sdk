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
	"errors"
	"github.com/aws/smithy-go"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/rds"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
)

type DBClusterStatus string

const (
	DBClusterStatusCreating        DBClusterStatus = "creating"
	DBClusterStatusAvailable       DBClusterStatus = "available"
	DBClusterStatusDeleting        DBClusterStatus = "deleting"
	DBClusterStatusFailed          DBClusterStatus = "failed"
	DBClusterStatusBackingUp       DBClusterStatus = "backing-up"
	DBClusterStatusBacktracking    DBClusterStatus = "backtracking"
	DBClusterStatusCloningFailed   DBClusterStatus = "cloning-failed"
	DBClusterStatusFailingOver     DBClusterStatus = "failing-over"
	DBClusterStatusMaintenance     DBClusterStatus = "maintenance"
	DBClusterStatusMigrating       DBClusterStatus = "migrating"
	DBClusterStatusMigrationFailed DBClusterStatus = "migration-failed"
	DBClusterStatusModifying       DBClusterStatus = "modifying"
	DBClusterStatusPromoting       DBClusterStatus = "promoting"
	DBClusterStatusRenaming        DBClusterStatus = "renaming"
	DBClusterStatusStaring         DBClusterStatus = "starting"
	DBClusterStatusStopping        DBClusterStatus = "stopping"
	DBClusterStatusStopped         DBClusterStatus = "stopped"
	DBClusterStatusUpgrading       DBClusterStatus = "upgrading"
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
	SetFinalDBSnapshotIdentifier(id string) Cluster
	SetSkipSnapshot(bool) Cluster

	Failover(context.Context) error
	FailoverGlobal(context.Context) error
	Create(context.Context) error
	Delete(context.Context) error
	Reboot(context.Context) error
	Describe(context.Context) (*DescCluster, error)
	RestorePitr(context.Context) error
	CreateSnapshot(context.Context) error
	DescribeSnapshot(context.Context) (*DescSnapshot, error)
}

type rdsCluster struct {
	core                           *rds.Client
	createClusterParam             *rds.CreateDBClusterInput
	deleteClusterParam             *rds.DeleteDBClusterInput
	failoverClusterParam           *rds.FailoverDBClusterInput
	failoverGlobalClusterParam     *rds.FailoverGlobalClusterInput
	rebootClusterParam             *rds.RebootDBClusterInput
	describeClusterParam           *rds.DescribeDBClustersInput
	restoreDBClusterPitrParam      *rds.RestoreDBClusterToPointInTimeInput
	createDBClusterSnapshotParam   *rds.CreateDBClusterSnapshotInput
	describeDBClusterSnapshotParam *rds.DescribeDBClusterSnapshotsInput
}

func (s *rdsCluster) SetSkipSnapshot(enable bool) Cluster {
	s.deleteClusterParam.SkipFinalSnapshot = enable
	return s
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
	if err != nil {
		if _, ok := errors.Unwrap(err.(*smithy.OperationError).Err).(*types.DBClusterNotFoundFault); !ok {
			return err
		}
	}
	return nil
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
	s.describeDBClusterSnapshotParam.DBClusterSnapshotIdentifier = aws.String(id)
	return s
}

func (s *rdsCluster) SetFinalDBSnapshotIdentifier(id string) Cluster {
	s.deleteClusterParam.FinalDBSnapshotIdentifier = aws.String(id)
	return s
}

func (s *rdsCluster) CreateSnapshot(ctx context.Context) error {
	_, err := s.core.CreateDBClusterSnapshot(ctx, s.createDBClusterSnapshotParam)
	return err
}

func (s *rdsCluster) DescribeSnapshot(ctx context.Context) (*DescSnapshot, error) {
	snapshots, err := s.core.DescribeDBClusterSnapshots(ctx, s.describeDBClusterSnapshotParam)
	if err != nil {
		return nil, err
	}
	if len(snapshots.DBClusterSnapshots) == 0 {
		return nil, nil
	}
	snapshot := snapshots.DBClusterSnapshots[0]
	return convertDBClusterSnapshot(&snapshot), nil
}

type DescCluster struct {
	CharSetName                 string
	ClusterCreateTime           time.Time
	AvailabilityZones           []string
	CustomEndpoints             []string
	DBClusterArn                string
	DBClusterIdentifier         string
	DBClusterMembers            []ClusterMember
	DBClusterParameterGroup     string
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

type DescSnapshot struct {
	ClusterCreateTime           time.Time
	DBClusterIdentifier         string
	DBClusterSnapshotArn        string
	DBClusterSnapshotIdentifier string
	Engine                      string
	EngineVersion               string
	PercentProgress             int32
	SnapshotCreateTime          time.Time
	SnapshotType                string
	Status                      string
}

// convertDBCluster converts aws types.DBCluster to DescCluster
func convertDBCluster(in *types.DBCluster) *DescCluster {
	return &DescCluster{
		CharSetName:                 aws.ToString(in.CharacterSetName),
		ClusterCreateTime:           aws.ToTime(in.ClusterCreateTime),
		AvailabilityZones:           in.AvailabilityZones,
		CustomEndpoints:             in.CustomEndpoints,
		DBClusterArn:                aws.ToString(in.DBClusterArn),
		DBClusterIdentifier:         aws.ToString(in.DBClusterIdentifier),
		DBClusterMembers:            convertDBClusterMembers(in.DBClusterMembers),
		DBClusterParameterGroup:     aws.ToString(in.DBClusterParameterGroup),
		DeletionProtection:          aws.ToBool(in.DeletionProtection),
		PrimaryEndpoint:             aws.ToString(in.Endpoint),
		ReadReplicaIdentifiers:      in.ReadReplicaIdentifiers,
		ReaderEndpoint:              aws.ToString(in.ReaderEndpoint),
		ReplicationSourceIdentifier: aws.ToString(in.ReplicationSourceIdentifier),
		Status:                      aws.ToString(in.Status),
		Port:                        aws.ToInt32(in.Port),
	}
}

func convertDBClusterMembers(in []types.DBClusterMember) []ClusterMember {
	var out []ClusterMember
	for _, m := range in {
		out = append(out, ClusterMember{
			DBClusterParameterGroupStatus: aws.ToString(m.DBClusterParameterGroupStatus),
			DBInstanceIdentifier:          aws.ToString(m.DBInstanceIdentifier),
			IsClusterWrite:                m.IsClusterWriter,
		})
	}
	return out
}

func (s *rdsCluster) Describe(ctx context.Context) (*DescCluster, error) {
	output, err := s.core.DescribeDBClusters(ctx, s.describeClusterParam)

	if err != nil {
		if _, ok := errors.Unwrap(err.(*smithy.OperationError).Err).(*types.DBClusterNotFoundFault); ok {
			return nil, nil
		}
		return nil, err
	}

	desc := &DescCluster{}
	if len(output.DBClusters) > 0 {
		desc = convertDBCluster(&output.DBClusters[0])
	}
	return desc, nil
}

func convertDBClusterSnapshot(in *types.DBClusterSnapshot) *DescSnapshot {
	return &DescSnapshot{
		ClusterCreateTime:           aws.ToTime(in.ClusterCreateTime),
		DBClusterIdentifier:         aws.ToString(in.DBClusterIdentifier),
		DBClusterSnapshotArn:        aws.ToString(in.DBClusterSnapshotArn),
		DBClusterSnapshotIdentifier: aws.ToString(in.DBClusterSnapshotIdentifier),
		Engine:                      aws.ToString(in.Engine),
		EngineVersion:               aws.ToString(in.EngineVersion),
		PercentProgress:             in.PercentProgress,
		SnapshotCreateTime:          aws.ToTime(in.SnapshotCreateTime),
		SnapshotType:                aws.ToString(in.SnapshotType),
		Status:                      aws.ToString(in.Status),
	}
}
