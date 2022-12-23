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
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:JSONPath=".spec.schedule",name=Schedule,type=string
// +kubebuilder:resource:shortName="dbchaos"
type DatabaseChaos struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DatabaseChaosSpec   `json:"spec,omitempty"`
	Status            DatabaseChaosStatus `json:"status,omitempty"`
}

type DatabaseChaosSpec struct {
	Selector metav1.LabelSelector `json:"selector"`
	Action   DatabaseChaosAction  `json:"action"`
	Schedule string               `json:"schedule"`
	// +optional
	Suspend bool `json:"suspend"`
}

type DatabaseChaosAction string

const (
	AWSRDSInstanceReboot  DatabaseChaosAction = "aws-rds-instance-reboot"
	AWSRDSClusterFailover DatabaseChaosAction = "aws-rds-cluster-failover"
)

type DatabaseChaosStatus struct {
	Conditions []DatabaseChaosCondition `json:"conditions,omitempty"`
	Records    []*DatabaseChaosRecord   `json:"records,omitempty"`
}

type DatabaseChaosConditionType string

const (
	ConditionSelected  DatabaseChaosConditionType = "Selected"
	ConditionExecuted  DatabaseChaosConditionType = "Executed"
	ConditionPaused    DatabaseChaosConditionType = "Paused"
	ConditionRecovered DatabaseChaosConditionType = "Recovered"
)

type DatabaseChaosCondition struct {
	Type   DatabaseChaosConditionType `json:"type"`
	Status corev1.ConditionStatus     `json:"status"`
	Reason string                     `json:"reason,omitempty"`
}

type DatabaseChaosRecord struct {
	ExecutionCount int                  `json:"executionCount"`
	Events         []DatabaseChaosEvent `json:"events,omitempty"`
}

type DatabaseChaosEvent struct {
	Type      RecordEventType `json:"type"`
	Message   string          `json:"message,omitempty"`
	Timestamp *metav1.Time    `json:"timestamp"`
}

type RecordEventType string

const (
	TypeSucceeded RecordEventType = "Succeeded"
	TypeFailed    RecordEventType = "Failed"
)

// +kubebuilder:object:root=true
type DatabaseChaosList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DatabaseChaos `json:"items"`
}
