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
	DefaultMasterUsername       string `json:"defaultMasterUsername"`
	AutoGeneratedMasterPassword bool   `json:"autoGeneratedMasterPassword"`

	MultiAZ     bool                `json:"multiAZ"`
	Provisioner DatabaseProvisioner `json:"provisioner"`
	Engine      DatabaseEngine      `json:"engine"`
	Instance    DatabaseInstance    `json:"instance"`
	Storage     DatabaseStorage     `json:"storage"`
}

type DatabaseProvisioner string

const (
	DatabaseProvisionerAWSRdsInstance DatabaseProvisioner = "AWSRdsInstance"
	DatabaseProvisionerAWSRdsCluster  DatabaseProvisioner = "AWSRdsCluster"
	DatabaseProvisionerAWSRdsAurora   DatabaseProvisioner = "AWSRdsAurora"
)

type DatabaseEngine struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	// +optional
	Mode string `json:"mode"`
}

type DatabaseInstance struct {
	// +optional
	Class string `json:"class"`
}

const (
	AnnotationsVPCSecurityGroupIds = "databaseclass.database-mesh.io/vpc-security-group-ids"
	AnnotationsSubnetGroupName     = "databaseclass.database-mesh.io/vpc-subnet-group-name"
	AnnotationsAvailabilityZones   = "databaseclass.database-mesh.io/availability-zones"
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
