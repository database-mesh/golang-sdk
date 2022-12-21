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
type DataShard struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DataShardSpec   `json:"spec,omitempty"`
	Status            DataShardStatus `json:"status,omitempty"`
}

// DataShardSpec is spec for DataShard
type DataShardSpec struct {
	Rules []ShardingRule `json:"rules"`
}

type ShardingRule struct {
	TableName               string                    `json:"tableName"`
	ReadWriteSplittingGroup []ReadWriteSplittingGroup `json:"readWriteSplittingGroup,omitempty"`
	ActualDatanodes         ActualDatanodesValue      `json:"actualDatanodes"`
	TableStrategy           *TableStrategy            `json:"tableStrategy,omitempty"`
	DatabaseStrategy        *DatabaseStrategy         `json:"databaseStrategy,omitempty"`
	DatabaseTableStrategy   *DatabaseTableStrategy    `json:"databaseTableStrategy,omitempty"`
}

type TableStrategy struct {
	TableShardingAlgorithmName string `json:"tableShardingAlgorithmName"`
	TableShardingColumn        string `json:"tableShardingColumn"`
	ShardingCount              uint32 `json:"shardingCount"`
}

type DatabaseStrategy struct {
	DatabaseShardingAlgorithmName string `json:"databaseShardingAlgorithmName"`
	DatabaseShardingColumn        string `json:"databaseShardingColumn"`
}

type DatabaseTableStrategy struct {
	TableStrategy    `json:",inline"`
	DatabaseStrategy `json:",inline"`
}

type ActualDatanodesValue struct {
	ValueSource *ValueSourceType `json:"valueSource"`
}

type ValueSourceType struct {
	*ActualDatanodesExpressionValue `json:",inline"`
	*ActualDatanodesNodeValue       `json:",inline"`
}

type ActualDatanodesExpressionValue struct {
	Expression string `json:"expression,omitempty"`
}

type ActualDatanodesNodeValue struct {
	Nodes []ValueFrom `json:"nodes,omitempty"`
}

type ValueFrom struct {
	Value                       string                       `json:"value,omitempty"`
	ValueFromReadWriteSplitting *ValueFromReadWriteSplitting `json:"valueFromReadWriteSplitting,omitempty"`
}

type ValueFromReadWriteSplitting struct {
	Name string `json:"name,omitempty"`
}

type ReadWriteSplittingGroup struct {
	Name  string                   `json:"name"`
	Rules []ReadWriteSplittingRule `json:"rules,omitempty"`
}

type DataShardStatus struct {
}

// +kubebuilder:object:root=true
type DataShardList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DataShard `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DataShard{}, &DataShardList{})
}
