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
	"os"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/rds"
	dbmesh "github.com/database-mesh/golang-sdk/aws"
)

const (
	EnvAWSRegion          = "AWS_REGION"
	EnvAWSAccessKey       = "AWS_ACCESS_KEY"
	EnvAWSSecretAccessKey = "AWS_SECRET_ACCESS_KEY"
)

func Test_CreateRDSInstance(t *testing.T) {
	region, _ := os.LookupEnv(EnvAWSRegion)
	accessKey, _ := os.LookupEnv(EnvAWSAccessKey)
	secretAccessKey, _ := os.LookupEnv(EnvAWSSecretAccessKey)
	sess := dbmesh.NewSessions().SetCredential(region, accessKey, secretAccessKey).Build()
	err := NewService(sess[region]).Instance().
		SetEngine("mysql").
		SetEngineVersion("8.0.28").
		SetDBInstanceIdentifier("foo").
		SetMasterUsername("admin").
		SetMasterUserPassword("sp").
		SetDBInstanceClass("db.m5.large").
		SetAllocatedStorage(40).
		SetDBName("foo").
		SetVpcSecurityGroupIds([]string{"sg-gg"}).
		SetDBSubnetGroup("unknown-sng").
		Create(context.TODO())

	if err != nil {
		t.Fatalf("%+v\n", err)
	}
	t.Logf("succ\n")
}

func Test_CreateRDSInstanceWithMultiAZ(t *testing.T) {
	region, _ := os.LookupEnv(EnvAWSRegion)
	accessKey, _ := os.LookupEnv(EnvAWSAccessKey)
	secretAccessKey, _ := os.LookupEnv(EnvAWSSecretAccessKey)
	sess := dbmesh.NewSessions().SetCredential(region, accessKey, secretAccessKey).Build()
	err := NewService(sess[region]).Instance().
		SetEngine("mysql").
		SetEngineVersion("8.0.28").
		SetDBInstanceIdentifier("fooaaa").
		SetMasterUsername("admin").
		SetMasterUserPassword("sp").
		SetDBInstanceClass("db.m5.large").
		SetAllocatedStorage(40).
		SetDBName("foomaz").
		SetVpcSecurityGroupIds([]string{"sg-gg"}).
		SetDBSubnetGroup("default").
		SetMultiAZ(true).
		Create(context.TODO())

	if err != nil {
		t.Fatalf("%+v\n", err)
	}
	t.Logf("succ\n")
}

func Test_DeleteRDSInstance(t *testing.T) {
	region, _ := os.LookupEnv(EnvAWSRegion)
	accessKey, _ := os.LookupEnv(EnvAWSAccessKey)
	secretAccessKey, _ := os.LookupEnv(EnvAWSSecretAccessKey)
	sess := dbmesh.NewSessions().SetCredential(region, accessKey, secretAccessKey).Build()
	err := NewService(sess[region]).Instance().SetDBInstanceIdentifier("fooaaa").SetSkipFinalSnapshot(true).SetDeleteAutomateBackups(false).Delete(context.TODO())
	if err != nil {
		t.Fatalf("%+v\n", err)
	}
	t.Logf("succ\n")
}

func Test_RebootRDSInstance(t *testing.T) {
	region, _ := os.LookupEnv(EnvAWSRegion)
	accessKey, _ := os.LookupEnv(EnvAWSAccessKey)
	secretAccessKey, _ := os.LookupEnv(EnvAWSSecretAccessKey)
	sess := dbmesh.NewSessions().SetCredential(region, accessKey, secretAccessKey).Build()
	err := NewService(sess[region]).Instance().SetDBInstanceIdentifier("fooaaa").SetForceFailover(true).RebootDBInstance(context.TODO())
	if err != nil {
		t.Fatalf("%+v\n", err)
	}
	t.Logf("succ\n")
}

func Test_CreateRDSCluster(t *testing.T) {
	region, _ := os.LookupEnv(EnvAWSRegion)
	accessKey, _ := os.LookupEnv(EnvAWSAccessKey)
	secretAccessKey, _ := os.LookupEnv(EnvAWSSecretAccessKey)
	sess := dbmesh.NewSessions().SetCredential(region, accessKey, secretAccessKey).Build()
	err := NewService(sess[region]).Cluster().
		SetEngine("mysql").
		SetEngineVersion("8.0.28").
		SetDBClusterIdentifier("foobbb").
		SetMasterUsername("admin").
		SetMasterUserPassword("sp").
		SetDBClusterInstanceClass("db.m5d.large").
		SetAllocatedStorage(100).
		SetDatabaseName("foomaz").
		SetVpcSecurityGroupIds([]string{"sg-gg"}).
		SetStorageType("io1").
		SetIOPS(1000).
		SetDBSubnetGroupName("test").
		SetAvailabilityZones([]string{"ap-southeast-1a", "ap-southeast-1c"}).
		Create(context.TODO())

	if err != nil {
		t.Fatalf("%+v\n", err)
	}
	t.Logf("succ\n")
}

func Test_DeleteRDSCluster(t *testing.T) {
	region, _ := os.LookupEnv(EnvAWSRegion)
	accessKey, _ := os.LookupEnv(EnvAWSAccessKey)
	secretAccessKey, _ := os.LookupEnv(EnvAWSSecretAccessKey)
	sess := dbmesh.NewSessions().SetCredential(region, accessKey, secretAccessKey).Build()
	err := NewService(sess[region]).Cluster().SetDBClusterIdentifier("foobbb").SetSkipFinalSnapshot(true).Delete(context.TODO())
	if err != nil {
		t.Fatalf("%+v\n", err)
	}
	t.Logf("succ\n")
}

func Test_FailoverRDSCluster(t *testing.T) {
	region, _ := os.LookupEnv(EnvAWSRegion)
	accessKey, _ := os.LookupEnv(EnvAWSAccessKey)
	secretAccessKey, _ := os.LookupEnv(EnvAWSSecretAccessKey)
	sess := dbmesh.NewSessions().SetCredential(region, accessKey, secretAccessKey).Build()
	err := NewService(sess[region]).Cluster().SetDBClusterIdentifier("foobbb").Failover(context.TODO())
	if err != nil {
		t.Fatalf("%+v\n", err)
	}
	t.Logf("succ\n")
}

func Test_CreateRDSSubnetsGroup(t *testing.T) {
	region, _ := os.LookupEnv(EnvAWSRegion)
	accessKey, _ := os.LookupEnv(EnvAWSAccessKey)
	secretAccessKey, _ := os.LookupEnv(EnvAWSSecretAccessKey)
	sess := dbmesh.NewSessions().SetCredential(region, accessKey, secretAccessKey).Build()
	client := NewService(sess[region]).Cluster().core
	snginput := &rds.CreateDBSubnetGroupInput{
		SubnetIds:                []string{"subnet-gg", "subnet-gg", "subnet-gg"},
		DBSubnetGroupName:        aws.String("test"),
		DBSubnetGroupDescription: aws.String("test"),
	}
	_, err := client.CreateDBSubnetGroup(context.TODO(), snginput)
	if err != nil {
		t.Fatalf("%+v\n", err)
	}
	t.Logf("succ\n")
}
