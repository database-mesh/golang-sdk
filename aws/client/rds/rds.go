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
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/rds"
)

type RDS interface {
	Instance() Instance
	Cluster() Cluster
}

type service struct {
	instance *rdsInstance
	cluster  *rdsCluster
}

func (s *service) Instance() Instance {
	return s.instance
}

func (s *service) Cluster() Cluster {
	return s.cluster
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
		},
		cluster: &rdsCluster{
			core:                       rds.NewFromConfig(sess),
			createClusterParam:         &rds.CreateDBClusterInput{},
			deleteClusterParam:         &rds.DeleteDBClusterInput{},
			failoverClusterParam:       &rds.FailoverDBClusterInput{},
			failoverGlobalClusterParam: &rds.FailoverGlobalClusterInput{},
			rebootClusterParam:         &rds.RebootDBClusterInput{},
			describeClusterParam:       &rds.DescribeDBClustersInput{},
			restoreDBClusterPitrParam:  &rds.RestoreDBClusterToPointInTimeInput{},
		},
	}
}

type Instance interface {
	SetEngine(engine string) Instance
	SetEngineVersion(version string) Instance
	SetDBInstanceIdentifier(id string) Instance
	SetMasterUsername(username string) Instance
	SetMasterUserPassword(pass string) Instance
	SetDBInstanceClass(class string) Instance
	SetAllocatedStorage(size int32) Instance
	SetIOPS(iops int32) Instance
	SetDBName(name string) Instance
	SetVpcSecurityGroupIds(sgs []string) Instance
	SetDBSubnetGroup(name string) Instance
	SetMultiAZ(enable bool) Instance
	SetAvailabilityZones(az string) Instance
	SetDeleteAutomateBackups(enable bool) Instance
	SetFinalDBSnapshotIdentifier(id string) Instance
	SetSkipFinalSnapshot(skip bool) Instance
	SetForceFailover(force bool) Instance
	SetTargetDBInstanceIdentifier(id string) Instance
	SetRestoreTime(rt *time.Time) Instance
	SetSourceDBInstanceAutomatedBackupsArn(arn string) Instance
	SetSourceDBInstanceIdentifier(id string) Instance
	SetSourceDBiResourceId(dbi string) Instance
	SetUseLatestRestorableTime(enable bool) Instance
	SetDBClusterIdentifier(id string) Instance
	SetPublicAccessible(enable bool) Instance
	SetLicenseModel(model string) Instance

	Create(context.Context) error
	Delete(context.Context) error
	Reboot(context.Context) error
	Describe(context.Context) (*DescInstance, error)
	RestorePitr(context.Context) error
}

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

	Failover(context.Context) error
	FailoverGlobal(context.Context) error
	Create(context.Context) error
	Delete(context.Context) error
	Reboot(context.Context) error
	Describe(context.Context) (*DescCluster, error)
	RestorePitr(context.Context) error
}

type rdsInstance struct {
	core                     *rds.Client
	createInstanceParam      *rds.CreateDBInstanceInput
	deleteInstanceParam      *rds.DeleteDBInstanceInput
	rebootInstanceParam      *rds.RebootDBInstanceInput
	describeInstanceParam    *rds.DescribeDBInstancesInput
	restoreInstancePitrParam *rds.RestoreDBInstanceToPointInTimeInput
}

// CreateDBInstanceInput
func (s *rdsInstance) SetEngine(engine string) Instance {
	s.createInstanceParam.Engine = aws.String(engine)
	return s
}

func (s *rdsInstance) SetEngineVersion(version string) Instance {
	s.createInstanceParam.EngineVersion = aws.String(version)
	return s
}

func (s *rdsInstance) SetDBInstanceIdentifier(id string) Instance {
	s.createInstanceParam.DBInstanceIdentifier = aws.String(id)
	s.deleteInstanceParam.DBInstanceIdentifier = aws.String(id)
	s.rebootInstanceParam.DBInstanceIdentifier = aws.String(id)
	s.describeInstanceParam.DBInstanceIdentifier = aws.String(id)
	return s
}

func (s *rdsInstance) SetMasterUsername(username string) Instance {
	s.createInstanceParam.MasterUsername = aws.String(username)
	return s
}

func (s *rdsInstance) SetMasterUserPassword(pass string) Instance {
	s.createInstanceParam.MasterUserPassword = aws.String(pass)
	return s
}

func (s *rdsInstance) SetDBInstanceClass(class string) Instance {
	s.createInstanceParam.DBInstanceClass = aws.String(class)
	s.restoreInstancePitrParam.DBInstanceClass = aws.String(class)
	return s
}

func (s *rdsInstance) SetAllocatedStorage(size int32) Instance {
	s.createInstanceParam.AllocatedStorage = aws.Int32(size)
	// s.restoreInstancePitrParam.MaxAllocatedStorage = aws.Int32(size)
	return s
}

func (s *rdsInstance) SetIOPS(iops int32) Instance {
	s.createInstanceParam.Iops = aws.Int32(iops)
	s.restoreInstancePitrParam.Iops = aws.Int32(iops)
	return s
}

func (s *rdsInstance) SetDBName(name string) Instance {
	s.createInstanceParam.DBName = aws.String(name)
	s.restoreInstancePitrParam.DBName = aws.String(name)
	return s
}

func (s *rdsInstance) SetVpcSecurityGroupIds(sgs []string) Instance {
	s.createInstanceParam.VpcSecurityGroupIds = sgs
	s.restoreInstancePitrParam.VpcSecurityGroupIds = sgs
	return s
}

func (s *rdsInstance) SetDBSubnetGroup(name string) Instance {
	s.createInstanceParam.DBSubnetGroupName = aws.String(name)
	s.restoreInstancePitrParam.DBSubnetGroupName = aws.String(name)
	return s
}

func (s *rdsInstance) SetMultiAZ(enable bool) Instance {
	s.createInstanceParam.MultiAZ = aws.Bool(enable)
	s.restoreInstancePitrParam.MultiAZ = aws.Bool(enable)
	return s
}

func (s *rdsInstance) SetAvailabilityZones(az string) Instance {
	s.createInstanceParam.AvailabilityZone = aws.String(az)
	s.restoreInstancePitrParam.AvailabilityZone = aws.String(az)
	return s
}

func (s *rdsInstance) Create(ctx context.Context) error {
	_, err := s.core.CreateDBInstance(ctx, s.createInstanceParam)
	return err
}

// DeleteDBInstanceInput
func (s *rdsInstance) SetDeleteAutomateBackups(enable bool) Instance {
	s.deleteInstanceParam.DeleteAutomatedBackups = aws.Bool(enable)
	return s
}

func (s *rdsInstance) SetFinalDBSnapshotIdentifier(id string) Instance {
	s.deleteInstanceParam.DBInstanceIdentifier = aws.String(id)
	return s
}

func (s *rdsInstance) SetSkipFinalSnapshot(skip bool) Instance {
	s.deleteInstanceParam.SkipFinalSnapshot = skip
	return s
}

func (s *rdsInstance) Delete(ctx context.Context) error {
	_, err := s.core.DeleteDBInstance(ctx, s.deleteInstanceParam)
	return err
}

// NOTE: ForceFailover cannot be specified since the instance is not configured for either MultiAZ or High Availability
// RebootDBInstanceInput
func (s *rdsInstance) SetForceFailover(force bool) Instance {
	s.rebootInstanceParam.ForceFailover = aws.Bool(force)
	return s
}

// NOTE: Can only reboot db instances with state in: available, storage-optimization, incompatible-credentials, incompatible-parameters.
func (s *rdsInstance) Reboot(ctx context.Context) error {
	_, err := s.core.RebootDBInstance(ctx, s.rebootInstanceParam)
	return err
}

func (s *rdsInstance) SetTargetDBInstanceIdentifier(tid string) Instance {
	s.restoreInstancePitrParam.TargetDBInstanceIdentifier = aws.String(tid)
	return s
}

// func (s *rdsInstance) SetAutoMinorVersionUpgrade(enable bool) Instance {
// 	s.restoreInstancePitrParam.AutoMinorVersionUpgrade = aws.Bool(enable)
// 	return s
// }

// func (s *rdsInstance) SetBackupTarget(target string) Instance {
// 	s.restoreInstancePitrParam.BackupTarget = aws.String(target)
// 	return s
// }

func (s *rdsInstance) SetRestoreTime(rt *time.Time) Instance {
	s.restoreInstancePitrParam.RestoreTime = rt
	return s
}

func (s *rdsInstance) SetSourceDBInstanceAutomatedBackupsArn(arn string) Instance {
	s.restoreInstancePitrParam.SourceDBInstanceAutomatedBackupsArn = aws.String(arn)
	return s
}

func (s *rdsInstance) SetSourceDBInstanceIdentifier(sid string) Instance {
	s.restoreInstancePitrParam.SourceDBInstanceIdentifier = aws.String(sid)
	return s
}

func (s *rdsInstance) SetSourceDBiResourceId(dbi string) Instance {
	s.restoreInstancePitrParam.SourceDbiResourceId = aws.String(dbi)
	return s
}

func (s *rdsInstance) SetUseLatestRestorableTime(enable bool) Instance {
	s.restoreInstancePitrParam.UseLatestRestorableTime = enable
	return s
}

func (s *rdsInstance) SetDBClusterIdentifier(id string) Instance {
	s.createInstanceParam.DBClusterIdentifier = aws.String(id)
	return s
}

func (s *rdsInstance) SetPublicAccessible(enable bool) Instance {
	s.createInstanceParam.PubliclyAccessible = aws.Bool(enable)
	return s
}

func (s *rdsInstance) SetLicenseModel(model string) Instance {
	s.createInstanceParam.LicenseModel = aws.String(model)
	return s
}

func (s *rdsInstance) RestorePitr(ctx context.Context) error {
	_, err := s.core.RestoreDBInstanceToPointInTime(ctx, s.restoreInstancePitrParam)
	return err
}

type ReadReplicaStatus struct {
	Message    string
	Normal     bool
	Status     string
	StatusType string
}
type Endpoint struct {
	Address string
	Port    int32
}

type DescInstance struct {
	CharSetName                           string
	DBInstanceArn                         string
	DBInstanceIdentifier                  string
	DBInstanceStatus                      string
	DeletionProtection                    bool
	InstanceCreateTime                    time.Time
	Timezone                              string
	SecondaryAZ                           string
	ReadReplicaSourceDBInstanceIdentifier string
	ReadReplicaDBInstanceIdentifiers      []string
	ReadReplicaStatusInfos                []ReadReplicaStatus
	Endpoint                              Endpoint
	DBParameterGroups                     []ParameterGroupStatus
	DBClusterIdentifier                   string
	ReadReplicaDBClusterIdentifiers       []string
}

type ParameterGroupStatus struct {
	Name        string
	ApplyStatus string
}

func (s *rdsInstance) Describe(ctx context.Context) (*DescInstance, error) {
	output, err := s.core.DescribeDBInstances(ctx, s.describeInstanceParam)
	if err != nil {
		return nil, err
	}
	desc := &DescInstance{}
	if len(output.DBInstances) > 0 {
		desc.CharSetName = aws.ToString(output.DBInstances[0].CharacterSetName)
		desc.DBInstanceArn = aws.ToString(output.DBInstances[0].DBInstanceArn)
		desc.DBInstanceIdentifier = aws.ToString(output.DBInstances[0].DBInstanceIdentifier)
		desc.DeletionProtection = output.DBInstances[0].DeletionProtection
		desc.InstanceCreateTime = aws.ToTime(output.DBInstances[0].InstanceCreateTime)
		desc.Timezone = aws.ToString(output.DBInstances[0].Timezone)
		desc.SecondaryAZ = aws.ToString(output.DBInstances[0].SecondaryAvailabilityZone)
		desc.ReadReplicaSourceDBInstanceIdentifier = aws.ToString(output.DBInstances[0].ReadReplicaSourceDBInstanceIdentifier)
		desc.ReadReplicaDBInstanceIdentifiers = output.DBInstances[0].ReadReplicaDBInstanceIdentifiers

		for _, s := range output.DBInstances[0].StatusInfos {
			desc.ReadReplicaStatusInfos = append(desc.ReadReplicaStatusInfos, ReadReplicaStatus{
				Message:    aws.ToString(s.Message),
				Normal:     s.Normal,
				Status:     aws.ToString(s.Status),
				StatusType: aws.ToString(s.StatusType),
			})
		}

		if output.DBInstances[0].DBInstanceStatus != nil {
			desc.DBInstanceStatus = aws.ToString(output.DBInstances[0].DBInstanceStatus)
		}

		if output.DBInstances[0].Endpoint != nil {
			desc.Endpoint = Endpoint{
				Address: aws.ToString(output.DBInstances[0].Endpoint.Address),
				Port:    aws.ToInt32(&output.DBInstances[0].Endpoint.Port),
			}
		}

		for _, g := range output.DBInstances[0].DBParameterGroups {
			desc.DBParameterGroups = append(desc.DBParameterGroups, ParameterGroupStatus{
				Name:        aws.ToString(g.DBParameterGroupName),
				ApplyStatus: aws.ToString(g.ParameterApplyStatus),
			})
		}

		desc.ReadReplicaDBClusterIdentifiers = output.DBInstances[0].ReadReplicaDBClusterIdentifiers
		desc.DBClusterIdentifier = aws.ToString(output.DBInstances[0].DBClusterIdentifier)
	}
	return desc, nil
}

type rdsCluster struct {
	core                       *rds.Client
	createClusterParam         *rds.CreateDBClusterInput
	deleteClusterParam         *rds.DeleteDBClusterInput
	failoverClusterParam       *rds.FailoverDBClusterInput
	failoverGlobalClusterParam *rds.FailoverGlobalClusterInput
	rebootClusterParam         *rds.RebootDBClusterInput
	describeClusterParam       *rds.DescribeDBClustersInput
	restoreDBClusterPitrParam  *rds.RestoreDBClusterToPointInTimeInput
}

// FailoverClusterInput
func (s *rdsCluster) SetDBClusterIdentifier(id string) Cluster {
	s.createClusterParam.DBClusterIdentifier = aws.String(id)
	s.deleteClusterParam.DBClusterIdentifier = aws.String(id)
	s.failoverClusterParam.DBClusterIdentifier = aws.String(id)
	s.rebootClusterParam.DBClusterIdentifier = aws.String(id)
	s.describeClusterParam.DBClusterIdentifier = aws.String(id)
	s.restoreDBClusterPitrParam.DBClusterIdentifier = aws.String(id)
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
