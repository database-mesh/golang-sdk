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
package client

import (
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type TrafficStrategyList struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Items             []TrafficStrategy `json:"items"`
}

type TrafficStrategy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TrafficStrategySpec `json:"spec,omitempty"`
}

// TrafficStrategySpec defines the desired state of TrafficStrategy
type TrafficStrategySpec struct {
	Selector            *metav1.LabelSelector `json:"selector"`
	LoadBalance         *LoadBalance          `json:"loadBalance,omitempty"`
	CircuitBreaks       []CircuitBreak        `json:"circuitBreaks,omitempty"`
	ConcurrencyControls []ConcurrencyControl  `json:"concurrencyControls,omitempty"`
}

// LoadBalance The choice of load balancing strategy, currently supported: SimpleLoadBalance
type LoadBalance struct {
	ReadWriteSplitting *ReadWriteSplitting `json:"readWriteSplitting,omitempty"`
	SimpleLoadBalance  *SimpleLoadBalance  `json:"simpleLoadBalance,omitempty"`
}

// ReadWriteSplitting support static and dynamic read-write splitting algorithm
type ReadWriteSplitting struct {
	Static  *ReadWriteSplittingStatic  `json:"static"`
	Dynamic *ReadWriteSplittingDynamic `json:"dynamic"`
}

// ReadWriteSplittingStatic defines static rules based read-write splitting algorithm
type ReadWriteSplittingStatic struct {
	DefaultTarget string                   `json:"defaultTarget,omitempty"`
	Rules         []ReadWriteSplittingRule `json:"rules,omitempty"`
}

// ReadWriteSplittingStaticRule defines static rules
type ReadWriteSplittingRule struct {
	Name          string               `json:"name"`
	Type          RuleType             `json:"type"`
	Regex         []string             `json:"regex,omitempty"`
	Target        string               `json:"target,omitempty"`
	AlgorithmName LoadBalanceAlgorithm `json:"algorithmName"`
}

type ReadWriteSplittingDynamic struct {
	DefaultTarget string                   `json:"defaultTarget,omitempty"`
	Rules         []ReadWriteSplittingRule `json:"rules,omitempty"`
	Discovery     ReadWriteDiscovery       `json:"discovery"`
}

type ReadWriteDiscovery struct {
	MasterHighAvailability *MasterHighAvailability `json:"masterHighAvailability,omitempty"`
}

type MasterHighAvailability struct {
	User                string               `json:"user"`
	Password            string               `json:"password"`
	MonitorInterval     uint64               `json:"monitorInterval"`
	ConnectionProbe     *ConnectionProbe     `json:"connectionProbe"`
	PingProbe           *PingProbe           `json:"pingProbe"`
	ReplicationLagProbe *ReplicationLagProbe `json:"replicationLagProbe"`
	ReadOnlyProbe       *ReadOnlyProbe       `json:"readOnlyProbe"`
}

type ReadOnlyProbe struct {
	*Probe `json:",inline"`
}

type ReplicationLagProbe struct {
	*Probe            `json:",inline"`
	MaxReplicationLag uint64 `json:"maxReplicationLag"`
}

type PingProbe struct {
	*Probe `json:",inline"`
}

type ConnectionProbe struct {
	*Probe `json:",inline"`
}

type Probe struct {
	PeriodMilliseconds  uint64 `json:"periodMilliseconds"`
	TimeoutMilliseconds uint64 `json:"timeoutMilliseconds"`
	FailureThreshold    uint64 `json:"failureThreshold"`
	SuccessThreshold    uint64 `json:"successThreshold"`
}

// RuleType defines the type of static rule
type RuleType string

const (
	RuleTypeRegex = "regex"
)

// LoadBalanceAlgorithm defines the name of managed loadbalance algorithm
type LoadBalanceAlgorithm string

const (
	LoadBalanceAlgorithmRandom     = "random"
	LoadBalanceAlgorithmRoundRobin = "roundrobin"
)

// SimpleLoadBalance support load balancing type: 1. random 2. roundrobin
type SimpleLoadBalance struct {
	Kind LoadBalanceAlgorithm `json:"kind"`
}

// CircuitBreak works with regular expressions.
// SQL statements that conform to regular expressions will be denied.
type CircuitBreak struct {
	Regex []string `json:"regex"`
}

// ConcurrencyControl works according to regular expressions.
// SQL statements that meet the regular conditions will be blown after the maximum concurrency limit is exceeded.
type ConcurrencyControl struct {
	Regex          []string      `json:"regex"`
	Duration       time.Duration `json:"duration"` // Issue: Duration:1ns in fmt.Print
	MaxConcurrency int           `json:"maxConcurrency"`
}
