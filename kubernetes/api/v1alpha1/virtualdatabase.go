// Copyright 2022 SphereEx Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +kubebuilder:subresource:status
// +kubebuilder:resource:shortName="vdb"
// +kubebuilder:object:root=true
type VirtualDatabase struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VirtualDatabaseSpec   `json:"spec,omitempty"`
	Status            VirtualDatabaseStatus `json:"status,omitempty"`
}

// VirtualDatabaseSpec defines the desired state of VirtualDatabase
type VirtualDatabaseSpec struct {
	// +optional
	DatabaseClassName string                   `json:"databaseClassName"`
	Services          []VirtualDatabaseService `json:"services"`
}

// Service Defines the content of a VirtualDatabase
type VirtualDatabaseService struct {
	DatabaseService `json:",inline"`

	Name string `json:"name"`
	// +optional
	TrafficStrategy string `json:"trafficStrategy"`
	// +optional
	DataShard string `json:"dataShard,omitempty"`
	// +optional
	QoSClaim string `json:"qosClaim,omitempty"`
}

// DatabaseService The type of VirtualDatabase that needs to be applied for.
// Current support: databaseMySQL
type DatabaseService struct {
	DatabaseMySQL *DatabaseMySQL `json:"databaseMySQL"`
}

// DatabaseMySQL The type one of VirtualDatabase.Represents a virtual MySQL type
type DatabaseMySQL struct {
	Host     string `json:"host,omitempty"`
	Port     uint32 `json:"port,omitempty"`
	User     string `json:"user,omitempty"`
	Password string `json:"password,omitempty"`
	// +optional
	DB string `json:"db,omitempty"`
	// +optional
	PoolSize uint32 `json:"poolSize,omitempty"`
	// +optional
	ServerVersion string `json:"serverVersion,omitempty"`
}

// VirtualDatabaseStatus defines the observed state of VirtualDatabase
// Endpoints display the name of the associated DatabaseEndpoint
// TODO: Implement dynamic updates
type VirtualDatabaseStatus struct {
	Endpoints []string `json:"endpoints"`
}

// +kubebuilder:object:root=true
type VirtualDatabaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VirtualDatabase `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VirtualDatabase{}, &VirtualDatabaseList{})
}
