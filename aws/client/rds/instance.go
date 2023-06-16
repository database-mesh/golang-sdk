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
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/rds"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/aws/smithy-go"
)

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
	SetSnapshotIdentifier(id string) Instance
	SetFilter(name string, values []string) Instance

	Create(context.Context) error
	Delete(context.Context) error
	Reboot(context.Context) error
	Describe(context.Context) (*DescInstance, error)
	DescribeAll(ctx context.Context) ([]*DescInstance, error)
	RestorePitr(context.Context) error
	CreateSnapshot(context.Context) error
	DescribeSnapshot(context.Context) (*DescSnapshot, error)
	RestoreFromSnapshot(context.Context) error
}

type rdsInstance struct {
	core                     *rds.Client
	createInstanceParam      *rds.CreateDBInstanceInput
	deleteInstanceParam      *rds.DeleteDBInstanceInput
	rebootInstanceParam      *rds.RebootDBInstanceInput
	describeInstanceParam    *rds.DescribeDBInstancesInput
	restoreInstancePitrParam *rds.RestoreDBInstanceToPointInTimeInput
	createSnapshotParam      *rds.CreateDBSnapshotInput
	describeSnapshotParam    *rds.DescribeDBSnapshotsInput
	restoreFromSnapshotParam *rds.RestoreDBInstanceFromDBSnapshotInput
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
	DBInstanceStatus                      DBInstanceStatus
	DBName                                string
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

type DescSnapshot struct {
	DBInstanceIdentifier string
	DBSnapshotArn        string
	DBSnapshotIdentifier string
	Engine               string
	EngineVersion        string
	InstanceCreateTime   time.Time
	PercentProgress      int32
	SnapshotCreateTime   time.Time
	SnapshotDatabaseTime time.Time
	SnapshotType         string
	Status               string
}

type DBInstanceStatus string

const (
	DBInstanceStatusAvailable DBInstanceStatus = "available"
	DBInstanceStatusBackingUp DBInstanceStatus = "backing-up"
	DBInstanceStatusCreating  DBInstanceStatus = "creating"
	DBInstanceStatusDeleting  DBInstanceStatus = "deleting"
	DBInstanceStatusFailed    DBInstanceStatus = "failed"
	DBInstanceStatusModifying DBInstanceStatus = "modifying"
	DBInstanceStatusRebooting DBInstanceStatus = "rebooting"
	DBInstanceStatusRenaming  DBInstanceStatus = "renaming"
	DBInstanceStatusStarting  DBInstanceStatus = "starting"
	DBInstanceStatusStopped   DBInstanceStatus = "stopped"
	DBInstanceStatusStopping  DBInstanceStatus = "stopping"
	DBInstanceStatusReady     DBInstanceStatus = "Ready"
)

type ParameterGroupStatus struct {
	Name        string
	ApplyStatus string
}

// CreateDBInstanceInput
func (s *rdsInstance) SetEngine(engine string) Instance {
	s.createInstanceParam.Engine = aws.String(engine)
	s.restoreFromSnapshotParam.Engine = aws.String(engine)
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
	s.createSnapshotParam.DBInstanceIdentifier = aws.String(id)
	s.describeInstanceParam.DBInstanceIdentifier = aws.String(id)
	s.restoreFromSnapshotParam.DBInstanceIdentifier = aws.String(id)
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
	s.restoreFromSnapshotParam.DBInstanceClass = aws.String(class)
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

	// restore from snapshot does not support SetDBName
	// s.restoreFromSnapshotParam.DBName = aws.String(name)
	return s
}

func (s *rdsInstance) SetVpcSecurityGroupIds(sgs []string) Instance {
	s.createInstanceParam.VpcSecurityGroupIds = sgs
	s.restoreInstancePitrParam.VpcSecurityGroupIds = sgs
	s.restoreFromSnapshotParam.VpcSecurityGroupIds = sgs
	return s
}

func (s *rdsInstance) SetDBSubnetGroup(name string) Instance {
	s.createInstanceParam.DBSubnetGroupName = aws.String(name)
	s.restoreInstancePitrParam.DBSubnetGroupName = aws.String(name)
	s.restoreFromSnapshotParam.DBSubnetGroupName = aws.String(name)
	return s
}

func (s *rdsInstance) SetMultiAZ(enable bool) Instance {
	s.createInstanceParam.MultiAZ = aws.Bool(enable)
	s.restoreInstancePitrParam.MultiAZ = aws.Bool(enable)
	s.restoreFromSnapshotParam.MultiAZ = aws.Bool(enable)
	return s
}

func (s *rdsInstance) SetAvailabilityZones(az string) Instance {
	s.createInstanceParam.AvailabilityZone = aws.String(az)
	s.restoreInstancePitrParam.AvailabilityZone = aws.String(az)
	s.restoreFromSnapshotParam.AvailabilityZone = aws.String(az)
	return s
}

func (s *rdsInstance) Create(ctx context.Context) error {
	_, err := s.core.CreateDBInstance(ctx, s.createInstanceParam)
	return err
}

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
	if err != nil {
		if _, ok := errors.Unwrap(err.(*smithy.OperationError).Err).(*types.DBInstanceNotFoundFault); ok {
			return nil
		}
		return err
	}
	return nil
}

// SetForceFailover
// NOTE: ForceFailover cannot be specified since the instance is not configured for either MultiAZ or High Availability
// RebootDBInstanceInput
func (s *rdsInstance) SetForceFailover(force bool) Instance {
	s.rebootInstanceParam.ForceFailover = aws.Bool(force)
	return s
}

// Reboot
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

func (s *rdsInstance) SetSnapshotIdentifier(id string) Instance {
	s.createSnapshotParam.DBSnapshotIdentifier = aws.String(id)
	s.describeSnapshotParam.DBSnapshotIdentifier = aws.String(id)
	s.restoreFromSnapshotParam.DBSnapshotIdentifier = aws.String(id)
	return s
}

func (s *rdsInstance) CreateSnapshot(ctx context.Context) error {
	_, err := s.core.CreateDBSnapshot(ctx, s.createSnapshotParam)
	return err
}

func (s *rdsInstance) DescribeSnapshot(ctx context.Context) (*DescSnapshot, error) {
	resp, err := s.core.DescribeDBSnapshots(ctx, s.describeSnapshotParam)
	if err != nil {
		return nil, err
	}
	if len(resp.DBSnapshots) == 0 {
		return nil, nil
	}
	return convertDBSnapshot(&resp.DBSnapshots[0]), nil
}

func convertDBSnapshot(in *types.DBSnapshot) *DescSnapshot {
	return &DescSnapshot{
		DBInstanceIdentifier: aws.ToString(in.DBInstanceIdentifier),
		DBSnapshotArn:        aws.ToString(in.DBSnapshotArn),
		DBSnapshotIdentifier: aws.ToString(in.DBSnapshotIdentifier),
		Engine:               aws.ToString(in.Engine),
		EngineVersion:        aws.ToString(in.EngineVersion),
		InstanceCreateTime:   aws.ToTime(in.InstanceCreateTime),
		PercentProgress:      in.PercentProgress,
		SnapshotCreateTime:   aws.ToTime(in.SnapshotCreateTime),
		SnapshotDatabaseTime: aws.ToTime(in.SnapshotDatabaseTime),
		SnapshotType:         aws.ToString(in.SnapshotType),
		Status:               aws.ToString(in.Status),
	}
}

func convertDBInstance(dbInstance *types.DBInstance) *DescInstance {
	desc := &DescInstance{}
	desc.CharSetName = aws.ToString(dbInstance.CharacterSetName)
	desc.DBInstanceArn = aws.ToString(dbInstance.DBInstanceArn)
	desc.DBInstanceIdentifier = aws.ToString(dbInstance.DBInstanceIdentifier)
	if dbInstance.DBInstanceStatus != nil {
		desc.DBInstanceStatus = convertDBInstanceStatus(dbInstance.DBInstanceStatus)
	}
	desc.DBName = aws.ToString(dbInstance.DBName)
	desc.DeletionProtection = dbInstance.DeletionProtection
	desc.InstanceCreateTime = aws.ToTime(dbInstance.InstanceCreateTime)
	desc.Timezone = aws.ToString(dbInstance.Timezone)
	desc.SecondaryAZ = aws.ToString(dbInstance.SecondaryAvailabilityZone)
	desc.ReadReplicaSourceDBInstanceIdentifier = aws.ToString(dbInstance.ReadReplicaSourceDBInstanceIdentifier)
	desc.ReadReplicaDBInstanceIdentifiers = dbInstance.ReadReplicaDBInstanceIdentifiers
	desc.ReadReplicaStatusInfos = convertReadReplicaStatus(dbInstance.StatusInfos)
	if dbInstance.Endpoint != nil {
		desc.Endpoint = convertEndpoint(dbInstance.Endpoint)
	}
	desc.DBParameterGroups = convertParameterGroupStatus(dbInstance.DBParameterGroups)
	desc.DBClusterIdentifier = aws.ToString(dbInstance.DBClusterIdentifier)
	desc.ReadReplicaDBClusterIdentifiers = dbInstance.ReadReplicaDBClusterIdentifiers
	return desc
}

func convertReadReplicaStatus(infos []types.DBInstanceStatusInfo) []ReadReplicaStatus {
	var readReplicaStatusInfos []ReadReplicaStatus
	for _, info := range infos {
		readReplicaStatusInfos = append(readReplicaStatusInfos, ReadReplicaStatus{
			Message:    aws.ToString(info.Message),
			Normal:     info.Normal,
			Status:     aws.ToString(info.Status),
			StatusType: aws.ToString(info.StatusType),
		})
	}
	return readReplicaStatusInfos
}

func convertEndpoint(endpoint *types.Endpoint) Endpoint {
	return Endpoint{
		Address: aws.ToString(endpoint.Address),
		Port:    endpoint.Port,
	}
}

func convertDBInstanceStatus(dbInstanceStatus *string) DBInstanceStatus {
	return DBInstanceStatus(aws.ToString(dbInstanceStatus))
}

func convertParameterGroupStatus(dbParameterGroups []types.DBParameterGroupStatus) []ParameterGroupStatus {
	var parameterGroupStatus []ParameterGroupStatus
	for _, group := range dbParameterGroups {
		parameterGroupStatus = append(parameterGroupStatus, ParameterGroupStatus{
			Name:        aws.ToString(group.DBParameterGroupName),
			ApplyStatus: aws.ToString(group.ParameterApplyStatus),
		})
	}
	return parameterGroupStatus
}

func (s *rdsInstance) Describe(ctx context.Context) (*DescInstance, error) {
	output, err := s.core.DescribeDBInstances(ctx, s.describeInstanceParam)
	if err != nil {
		if _, ok := errors.Unwrap(err.(*smithy.OperationError).Err).(*types.DBInstanceNotFoundFault); ok {
			return nil, nil
		}
		return nil, err
	}

	desc := &DescInstance{}
	if len(output.DBInstances) > 0 {
		desc = convertDBInstance(&output.DBInstances[0])
	}
	return desc, nil
}

func (s *rdsInstance) SetFilter(name string, values []string) Instance {
	var exists = false
	for _, filter := range s.describeInstanceParam.Filters {
		if filter.Name == nil {
			continue
		}
		if *filter.Name == name {
			filter.Values = values
			exists = true
		}
	}
	if !exists {
		s.describeInstanceParam.Filters = append(s.describeInstanceParam.Filters, types.Filter{
			Name:   aws.String(name),
			Values: values,
		})
	}
	return s
}

func (s *rdsInstance) DescribeAll(ctx context.Context) ([]*DescInstance, error) {
	output, err := s.core.DescribeDBInstances(ctx, s.describeInstanceParam)
	if err != nil {
		if _, ok := errors.Unwrap(err.(*smithy.OperationError).Err).(*types.DBInstanceNotFoundFault); ok {
			return nil, nil
		}
		return nil, err
	}
	var descs []*DescInstance
	for _, dbInstance := range output.DBInstances {
		descs = append(descs, convertDBInstance(&dbInstance))
	}
	return descs, nil
}

func (s *rdsInstance) RestoreFromSnapshot(ctx context.Context) error {
	_, err := s.core.RestoreDBInstanceFromDBSnapshot(ctx, s.restoreFromSnapshotParam)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	return nil
}
