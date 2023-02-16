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

	TestAWSRegion          = "region-test"
	TestAWSAccessKey       = "ak-test"
	TestAWSSecretAccessKey = "sk-test"
	TestSubnetGroup        = "test"
	TestVpcSecurityGroupId = "sg-008e74936b3f9de19"
	TestDBName             = "foo"
	TestDBPass             = "admin123"
	TestDBIdentifier       = "foo"
)

func Test_CreateRDSInstance(t *testing.T) {
	region, _ := os.LookupEnv(EnvAWSRegion)
	accessKey, _ := os.LookupEnv(EnvAWSAccessKey)
	secretAccessKey, _ := os.LookupEnv(EnvAWSSecretAccessKey)
	sess := dbmesh.NewSessions().SetCredential(region, accessKey, secretAccessKey).Build()
	err := NewService(sess[region]).Instance().
		SetEngine("mysql").
		SetEngineVersion("8.0.28").
		SetDBInstanceIdentifier(TestDBIdentifier).
		SetDBInstanceClass("db.m5.large").
		SetMasterUsername("admin").
		SetMasterUserPassword(TestDBPass).
		SetAllocatedStorage(40).
		SetDBName(TestDBName).
		SetVpcSecurityGroupIds([]string{TestVpcSecurityGroupId}).
		SetDBSubnetGroup(TestSubnetGroup).
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
		SetDBInstanceIdentifier(TestDBIdentifier).
		SetDBInstanceClass("db.m5.large").
		SetMasterUsername("admin").
		SetMasterUserPassword(TestDBPass).
		SetAllocatedStorage(40).
		SetDBName(TestDBName).
		SetVpcSecurityGroupIds([]string{TestVpcSecurityGroupId}).
		SetDBSubnetGroup(TestSubnetGroup).
		SetMultiAZ(true).
		Create(context.TODO())

	if err != nil {
		t.Fatalf("%+v\n", err)
	}
	t.Logf("succ\n")
}

func Test_DescribeRDSInstance(t *testing.T) {
	region, _ := os.LookupEnv(EnvAWSRegion)
	accessKey, _ := os.LookupEnv(EnvAWSAccessKey)
	secretAccessKey, _ := os.LookupEnv(EnvAWSSecretAccessKey)
	sess := dbmesh.NewSessions().SetCredential(region, accessKey, secretAccessKey).Build()
	output, err := NewService(sess[region]).Instance().
		SetDBInstanceIdentifier(TestDBIdentifier).
		Describe(context.TODO())

	if err != nil {
		t.Fatalf("%+v\n", err)
	}
	t.Logf("succ\n")
	t.Logf("%#v\n", output)
}

func Test_DeleteRDSInstance(t *testing.T) {
	region, _ := os.LookupEnv(EnvAWSRegion)
	accessKey, _ := os.LookupEnv(EnvAWSAccessKey)
	secretAccessKey, _ := os.LookupEnv(EnvAWSSecretAccessKey)
	sess := dbmesh.NewSessions().SetCredential(region, accessKey, secretAccessKey).Build()
	err := NewService(sess[region]).Instance().SetDBInstanceIdentifier("foo2-instance-1").SetSkipFinalSnapshot(true).SetDeleteAutomateBackups(false).Delete(context.TODO())
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
	err := NewService(sess[region]).Instance().SetDBInstanceIdentifier(TestDBIdentifier).SetForceFailover(true).Reboot(context.TODO())
	if err != nil {
		t.Fatalf("%+v\n", err)
	}
	t.Logf("succ\n")
}

func Test_DescRDSInstance(t *testing.T) {
	region, _ := os.LookupEnv(EnvAWSRegion)
	accessKey, _ := os.LookupEnv(EnvAWSAccessKey)
	secretAccessKey, _ := os.LookupEnv(EnvAWSSecretAccessKey)
	sess := dbmesh.NewSessions().SetCredential(region, accessKey, secretAccessKey).Build()
	desc, err := NewService(sess[region]).Instance().SetDBInstanceIdentifier(TestDBIdentifier).Describe(context.TODO())
	if err != nil {
		t.Fatalf("%+v\n", err)
	}
	t.Logf("succ\n")
	t.Logf("%#v\n", desc)
}

func Test_CreateRDSCluster(t *testing.T) {
	region, _ := os.LookupEnv(EnvAWSRegion)
	accessKey, _ := os.LookupEnv(EnvAWSAccessKey)
	secretAccessKey, _ := os.LookupEnv(EnvAWSSecretAccessKey)
	sess := dbmesh.NewSessions().SetCredential(region, accessKey, secretAccessKey).Build()
	err := NewService(sess[region]).Cluster().
		SetEngine("mysql").
		SetEngineVersion("8.0.28").
		SetDBClusterIdentifier(TestDBIdentifier).
		SetMasterUsername("admin").
		SetMasterUserPassword(TestDBPass).
		SetDBClusterInstanceClass("db.m5d.large").
		SetAllocatedStorage(100).
		SetDatabaseName(TestDBName).
		SetVpcSecurityGroupIds([]string{TestVpcSecurityGroupId}).
		SetStorageType("io1").
		SetIOPS(1000).
		SetDBSubnetGroupName(TestSubnetGroup).
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
	err := NewService(sess[region]).Cluster().SetDBClusterIdentifier("foo2").SetSkipFinalSnapshot(true).Delete(context.TODO())
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
	err := NewService(sess[region]).Cluster().SetDBClusterIdentifier(TestDBIdentifier).Failover(context.TODO())
	if err != nil {
		t.Fatalf("%+v\n", err)
	}
	t.Logf("succ\n")
}

func Test_DescribeRDSCluster(t *testing.T) {
	region, _ := os.LookupEnv(EnvAWSRegion)
	accessKey, _ := os.LookupEnv(EnvAWSAccessKey)
	secretAccessKey, _ := os.LookupEnv(EnvAWSSecretAccessKey)
	sess := dbmesh.NewSessions().SetCredential(region, accessKey, secretAccessKey).Build()
	output, err := NewService(sess[region]).Cluster().
		SetDBClusterIdentifier("test").
		Describe(context.TODO())

	if err != nil {
		t.Fatalf("%+v\n", err)
	}
	t.Logf("succ\n")
	t.Logf("%#v\n", output)
}

func Test_CreateRDSSubnetsGroup(t *testing.T) {
	region, _ := os.LookupEnv(EnvAWSRegion)
	accessKey, _ := os.LookupEnv(EnvAWSAccessKey)
	secretAccessKey, _ := os.LookupEnv(EnvAWSSecretAccessKey)
	sess := dbmesh.NewSessions().SetCredential(region, accessKey, secretAccessKey).Build()
	client := NewService(sess[region]).cluster.core
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

func Test_CreateRDSAurora(t *testing.T) {
	region, _ := os.LookupEnv(EnvAWSRegion)
	accessKey, _ := os.LookupEnv(EnvAWSAccessKey)
	secretAccessKey, _ := os.LookupEnv(EnvAWSSecretAccessKey)
	sess := dbmesh.NewSessions().SetCredential(region, accessKey, secretAccessKey).Build()

	err := NewService(sess[region]).Cluster().
		SetEngine("aurora-mysql").
		SetEngineVersion("5.7.mysql_aurora.2.07.0").
		SetDBClusterIdentifier("foo2").
		SetMasterUsername("admin").
		SetMasterUserPassword(TestDBPass).
		SetVpcSecurityGroupIds([]string{TestVpcSecurityGroupId}).
		SetDBSubnetGroupName("test").
		// SetAvailabilityZones([]string{"ap-southeast-1a", "ap-southeast-1c", "ap-southeast-1b"}).
		Create(context.TODO())

	if err != nil {
		t.Fatalf("%+v\n", err)
	}
}

func Test_CreateRDSInstanceForAurora(t *testing.T) {
	region, _ := os.LookupEnv(EnvAWSRegion)
	accessKey, _ := os.LookupEnv(EnvAWSAccessKey)
	secretAccessKey, _ := os.LookupEnv(EnvAWSSecretAccessKey)
	sess := dbmesh.NewSessions().SetCredential(region, accessKey, secretAccessKey).Build()
	err := NewService(sess[region]).Instance().
		SetEngine("aurora-mysql").
		SetDBInstanceIdentifier("foo2-instance-1").
		SetDBInstanceClass("db.r5.large").
		SetPublicAccessible(true).
		SetDBClusterIdentifier("foo2").
		Create(context.TODO())

	if err != nil {
		t.Fatalf("%+v\n", err)
	}

	t.Logf("succ\n")
}

func Test_DescribeRDSAurora(t *testing.T) {
	region, _ := os.LookupEnv(EnvAWSRegion)
	accessKey, _ := os.LookupEnv(EnvAWSAccessKey)
	secretAccessKey, _ := os.LookupEnv(EnvAWSSecretAccessKey)
	sess := dbmesh.NewSessions().SetCredential(region, accessKey, secretAccessKey).Build()
	input := &rds.DescribeDBClustersInput{
		DBClusterIdentifier: aws.String(TestDBIdentifier),
	}

	output, err := NewService(sess[region]).cluster.core.DescribeDBClusters(context.TODO(), input)

	if err != nil {
		t.Fatalf("%+v\n", err)
	}

	t.Logf("%+v\n", *output.DBClusters[0].Engine)
	t.Logf("%+v\n", *output.DBClusters[0].EngineMode)
	t.Logf("%+v\n", *output.DBClusters[0].EngineVersion)
	t.Logf("%+v\n", *output.DBClusters[0].Status)
	t.Logf("%+v\n", *output.DBClusters[0].Endpoint)
	t.Logf("%+v\n", *output.DBClusters[0].MultiAZ)
	// t.Logf("%+v\n", *output.DBClusters[0].StorageType)
	t.Logf("%+v\n", *output.DBClusters[0].ReaderEndpoint)
	// t.Logf("%+v\n", *output.DBClusters[0].PubliclyAccessible)
}
