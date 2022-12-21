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

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +kubebuilder:object:root=true

type DatabaseEndpointList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DatabaseEndpoint `json:"items"`
}

// +kubebuilder:object:root=true

type DatabaseEndpoint struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DatabaseEndpointSpec   `json:"spec,omitempty"`
	Status            DatabaseEndpointStatus `json:"status,omitempty"`
}

// DatabaseEndpointSpec defines the desired state of DatabaseEndpoint
type DatabaseEndpointSpec struct {
	Database Database `json:"database"`
}

// Database Backend data source type
type Database struct {
	MySQL *MySQL `json:"MySQL"`
}

// MySQL Configuration Definition
type MySQL struct {
	Host     string `json:"host"`
	Port     uint32 `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	DB       string `json:"db"`
}

type DatabaseEndpointStatus struct{}

func init() {
	SchemeBuilder.Register(&DatabaseEndpoint{}, &DatabaseEndpointList{})
}
