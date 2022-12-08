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

	dbmesh "github.com/database-mesh/golang-sdk/aws"
)

const (
	EnvAWSRegion          = "AWS_REGOIN"
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
		SetMasterUserPassword("admin").
		SetDBInstanceClass("db.m5.large").
		SetAllocatedStorage(40).
		SetDBName("foo").
		SetVpcSecurityGroupIds([]string{"sg-xxx"}).
		SetDBSubnetGroup("default").
		Create(context.TODO())

	if err != nil {
		t.Fatalf("%+v\n", err)
	}
	t.Logf("succ\n")
}
