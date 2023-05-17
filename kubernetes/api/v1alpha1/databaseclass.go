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

package v1alpha1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// +kubebuilder:resource:scope=Cluster,shortName="dc"
// +kubebuilder:object:root=true

type DatabaseClass struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DatabaseClassSpec   `json:"spec,omitempty"`
	Status            DatabaseClassStatus `json:"status,omitempty"`
}

type DatabaseClassSpec struct {
	Provisioner   string                `json:"provisioner"`
	Parameters    map[string]string     `json:"parameters"`
	ReclaimPolicy DatabaseReclaimPolicy `json:"reclaimPolicy"`
}

// DatabaseReclaimPolicy describes a policy for end-of-life maintenance of persistent volumes.
// +enum
type DatabaseReclaimPolicy string

const (
	// The database will be deleted with a final snapshot reserved.
	DatabaseReclaimDeleteWithFinalSnapshot DatabaseReclaimPolicy = "DeleteWithFinalSnapshot"
	// The database will be deleted.
	DatabaseReclaimDelete DatabaseReclaimPolicy = "Delete"
	// The database will be retained.
	// The default policy is Retain.
	DatabaseReclaimRetain DatabaseReclaimPolicy = "Retain"
)

const (
	AnnotationsVPCSecurityGroupIds = "databaseclass.database-mesh.io/vpc-security-group-ids"
	AnnotationsSubnetGroupName     = "databaseclass.database-mesh.io/vpc-subnet-group-name"
	AnnotationsAvailabilityZones   = "databaseclass.database-mesh.io/availability-zones"
	AnnotationsClusterIdentifier   = "databaseclass.database-mesh.io/cluster-identifier"
	AnnotationsInstanceIdentifier  = "databaseclass.database-mesh.io/instance-identifier"
	AnnotationsInstanceDBName      = "databaseclass.database-mesh.io/instance-db-name"
	AnnotationsSnapshotIdentifier  = "databaseclass.database-mesh.io/snapshot-identifier"
	AnnotationsMasterUsername      = "databaseclass.database-mesh.io/master-username"
	AnnotationsMasterUserPassword  = "databaseclass.database-mesh.io/master-user-password"
)

const (
	ProvisionerAWSRDSInstance = "databaseclass.database-mesh.io/aws-rds-instance"
	ProvisionerAWSRDSCluster  = "databaseclass.database-mesh.io/aws-rds-cluster"
	ProvisionerAWSAurora      = "databaseclass.database-mesh.io/aws-aurora"
)

type DatabaseStorage struct {
	// +optional
	AllocatedStorage int32 `json:"allocatedStorage"`
	// +optional
	IOPS int32 `json:"iops"`
}

type DatabaseClassStatus struct{}

// +kubebuilder:object:root=true
type DatabaseClassList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DatabaseClass `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DatabaseClass{}, &DatabaseClassList{})
}
