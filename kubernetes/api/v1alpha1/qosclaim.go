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

package client

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type QoSClaim struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              QoSClaimSpec   `json:"spec,omitempty"`
	Status            QoSClaimStatus `json:"status,omitempty"`
}

type QoSClaimSpec struct {
	TrafficQoS TrafficQoS `json:"trafficQoS"`
}

type TrafficQoS struct {
	Name     string   `json:"name"`
	QoSGroup QoSGroup `json:"qos_group,omitempty"`
}

type QoSGroup struct {
	Rate string `json:"rate,omitempty"`
	Ceil string `json:"ceil,omitempty"`
}

type QoSClaimStatus struct{}

type QoSClaimList struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Items             []QoSClaim `json:"items"`
}
