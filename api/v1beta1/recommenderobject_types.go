/*
Copyright 2023.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1beta1

import (
	resource "k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// Mode describes how the recommender will work.
// Only one of the following mode may be specified.
// If none of the following mode is specified, the default one
// is Suggestion.
type Mode string

const (
	// Suggestion mode to run recommender for providing suggestion for resources.
	Suggestion Mode = "Suggestion"

	// Enforce mode to run recommender to enforce the suggestion for resources.
	Enforce Mode = "Enforce"

	// Suspend mode to suspend the function of recommender.
	Suspend Mode = "Suspend"
)

type Target struct {
	Name string `json:"name"`

	APIVersion string `json:"apiVersion,omitempty"`

	Kind string `json:"kind,omitempty"`
}

type Type string

const (
	Cpu    Type = "cpu"
	Memory Type = "memory"
)

type Restrictions struct {
	Type  Type               `json:"type"`
	Value *resource.Quantity `json:"value"`
}

type ResourceRequests struct {
	Type  Type               `json:"type"`
	Value *resource.Quantity `json:"value"`
}

type ResourceLimits struct {
	Type  Type               `json:"type"`
	Value *resource.Quantity `json:"value"`
}

// RecommenderObjectSpec defines the desired state of RecommenderObject
type RecommenderObjectSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Specifies the mode of running of recommender.
	Mode Mode `json:"mode,omitempty"`

	// Specifiy the target referece
	TargetRef *Target `json:"targetRef"`

	// Specify the interval in hours to implement the recommendation
	// If not specified, the default value is
	// 12
	Interval *int32 `json:"interval,omitempty"`

	// Specify threshold percentage difference b/w current and new resources metrics
	Threshold *int32 `json:"threshold"`

	// Specify restrictions for recommender for resource changes
	Restrictions []Restrictions `json:"restrictions,omitempty"`
}

// RecommenderObjectStatus defines the observed state of RecommenderObject
type RecommenderObjectStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// +optional
	TargetKind string `json:"targetKind"`
	// +optional
	TargetName string `json:"targetName"`
	// +optional
	OriginalResouceRequests []ResourceRequests `json:"originalResouceRequests"`
	// +optional
	OriginalResouceLimits []ResourceLimits `json:"originalResouceLimits"`
	// +optional
	LastUpdateTime *metav1.Time `json:"lastUpdateTime,omitempty"`
	// +optional
	LatestResouceRequests []ResourceRequests `json:"latestResouceRequests"`
	// +optional
	LatestResouceLimits []ResourceLimits `json:"latestResouceLimits"`
	// +optional
	Restrictions Restrictions `json:"restrictions,omitempty"`
	// +optional
	Mode Mode `json:"mode"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// RecommenderObject is the Schema for the recommenderobjects API
type RecommenderObject struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RecommenderObjectSpec   `json:"spec,omitempty"`
	Status RecommenderObjectStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// RecommenderObjectList contains a list of RecommenderObject
type RecommenderObjectList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RecommenderObject `json:"items"`
}

func init() {
	SchemeBuilder.Register(&RecommenderObject{}, &RecommenderObjectList{})
}
