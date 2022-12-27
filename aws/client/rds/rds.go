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
	Instance(context.Context) Instance
	Cluster(context.Context) Cluster
}

type service struct {
	instance *rdsInstance
	cluster  *rdsCluster
}

func (s *service) Instance() *rdsInstance {
	return s.instance
}

func (s *service) Cluster() *rdsCluster {
	return s.cluster
}

func NewService(sess aws.Config) *service {
	return &service{
		instance: &rdsInstance{
			core:                  rds.NewFromConfig(sess),
			createInstanceParam:   &rds.CreateDBInstanceInput{},
			deleteInstanceParam:   &rds.DeleteDBInstanceInput{},
			rebootInstanceParam:   &rds.RebootDBInstanceInput{},
			describeInstanceParam: &rds.DescribeDBInstancesInput{},
		},
		cluster: &rdsCluster{
			core:                       rds.NewFromConfig(sess),
			createClusterParam:         &rds.CreateDBClusterInput{},
			deleteClusterParam:         &rds.DeleteDBClusterInput{},
			failoverClusterParam:       &rds.FailoverDBClusterInput{},
			failoverGlobalClusterParam: &rds.FailoverGlobalClusterInput{},
			rebootClusterParam:         &rds.RebootDBClusterInput{},
			describeClusterParam:       &rds.DescribeDBClustersInput{},
		},
	}
}

type Instance interface {
	Create(context.Context) error
	Delete(context.Context) error
	Reboot(context.Context) error
	Describe(context.Context) (*DescInstance, error)
}

type Cluster interface {
	Failover(context.Context) error
	FailoverGlobal(context.Context) error
	Create(context.Context) error
	Delete(context.Context) error
	Reboot(context.Context) error
	Describe(context.Context) (*DescCluster, error)
}

type rdsInstance struct {
	core                  *rds.Client
	createInstanceParam   *rds.CreateDBInstanceInput
	deleteInstanceParam   *rds.DeleteDBInstanceInput
	rebootInstanceParam   *rds.RebootDBInstanceInput
	describeInstanceParam *rds.DescribeDBInstancesInput
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
	s.rebootInstanceParam.DBInstanceIdentifier = aws.String(id)
	s.describeInstanceParam.DBInstanceIdentifier = aws.String(id)
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

func (s *rdsInstance) SetIOPS(iops int32) *rdsInstance {
	s.createInstanceParam.Iops = aws.Int32(iops)
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

func (s *rdsInstance) SetMultiAZ(enable bool) *rdsInstance {
	s.createInstanceParam.MultiAZ = aws.Bool(enable)
	return s
}

func (s *rdsInstance) SetAvailabilityZones(az string) *rdsInstance {
	s.createInstanceParam.AvailabilityZone = aws.String(az)
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

// NOTE: ForceFailover cannot be specified since the instance is not configured for either MultiAZ or High Availability
// RebootDBInstanceInput
func (s *rdsInstance) SetForceFailover(force bool) *rdsInstance {
	s.rebootInstanceParam.ForceFailover = aws.Bool(force)
	return s
}

// NOTE: Can only reboot db instances with state in: available, storage-optimization, incompatible-credentials, incompatible-parameters.
func (s *rdsInstance) Reboot(ctx context.Context) error {
	_, err := s.core.RebootDBInstance(ctx, s.rebootInstanceParam)
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
}

// FailoverClusterInput
func (s *rdsCluster) SetDBClusterIdentifier(id string) *rdsCluster {
	s.createClusterParam.DBClusterIdentifier = aws.String(id)
	s.deleteClusterParam.DBClusterIdentifier = aws.String(id)
	s.failoverClusterParam.DBClusterIdentifier = aws.String(id)
	s.rebootClusterParam.DBClusterIdentifier = aws.String(id)
	s.describeClusterParam.DBClusterIdentifier = aws.String(id)
	return s
}

func (s *rdsCluster) SetTargetDBInstanceIdentifier(id string) *rdsCluster {
	s.failoverClusterParam.TargetDBInstanceIdentifier = aws.String(id)
	return s
}

func (s *rdsCluster) Failover(ctx context.Context) error {
	_, err := s.core.FailoverDBCluster(ctx, s.failoverClusterParam)
	return err
}

// FailoverGlobalClusterInput
func (s *rdsCluster) SetGlobalClusterIdentifier(id string) *rdsCluster {
	s.failoverGlobalClusterParam.GlobalClusterIdentifier = aws.String(id)
	return s
}

func (s *rdsCluster) SetTargetDbClusterIdentifier(id string) *rdsCluster {
	s.failoverGlobalClusterParam.TargetDbClusterIdentifier = aws.String(id)
	return s
}

func (s *rdsCluster) FailoverGlobal(ctx context.Context) error {
	_, err := s.core.FailoverGlobalCluster(ctx, s.failoverGlobalClusterParam)
	return err
}

// CreateDBClusterInput
func (s *rdsCluster) SetEngine(engine string) *rdsCluster {
	s.createClusterParam.Engine = aws.String(engine)
	return s
}

func (s *rdsCluster) SetAllocatedStorage(size int32) *rdsCluster {
	s.createClusterParam.AllocatedStorage = aws.Int32(size)
	return s
}

func (s *rdsCluster) SetAvailabilityZones(azs []string) *rdsCluster {
	s.createClusterParam.AvailabilityZones = azs
	return s
}

func (s *rdsCluster) SetDBClusterInstanceClass(class string) *rdsCluster {
	s.createClusterParam.DBClusterInstanceClass = aws.String(class)
	return s
}

func (s *rdsCluster) SetDBSubnetGroupName(name string) *rdsCluster {
	s.createClusterParam.DBSubnetGroupName = aws.String(name)
	return s
}

func (s *rdsCluster) SetDatabaseName(name string) *rdsCluster {
	s.createClusterParam.DatabaseName = aws.String(name)
	return s
}

func (s *rdsCluster) SetEngineVersion(version string) *rdsCluster {
	s.createClusterParam.EngineVersion = aws.String(version)
	return s
}

func (s *rdsCluster) SetEngineMode(mode string) *rdsCluster {
	s.createClusterParam.EngineMode = aws.String(mode)
	return s
}

func (s *rdsCluster) SetMasterUsername(username string) *rdsCluster {
	s.createClusterParam.MasterUsername = aws.String(username)
	return s
}

func (s *rdsCluster) SetMasterUserPassword(pass string) *rdsCluster {
	s.createClusterParam.MasterUserPassword = aws.String(pass)
	return s
}

func (s *rdsCluster) SetVpcSecurityGroupIds(sgs []string) *rdsCluster {
	s.createClusterParam.VpcSecurityGroupIds = sgs
	return s
}

func (s *rdsCluster) SetStorageType(t string) *rdsCluster {
	s.createClusterParam.StorageType = aws.String(t)
	return s
}

func (s *rdsCluster) SetIOPS(ps int32) *rdsCluster {
	s.createClusterParam.Iops = aws.Int32(ps)
	return s
}

func (s *rdsCluster) Create(ctx context.Context) error {
	_, err := s.core.CreateDBCluster(ctx, s.createClusterParam)
	return err
}

// DeleteDBClusterInput
func (s *rdsCluster) SetSkipFinalSnapshot(skip bool) *rdsCluster {
	s.deleteClusterParam.SkipFinalSnapshot = *aws.Bool(skip)
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
		desc.Status = aws.ToString(output.DBClusters[0].Status)
	}
	return desc, nil
}
