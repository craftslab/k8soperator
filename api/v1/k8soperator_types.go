/*
Copyright 2023 craftslab.

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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// K8sOperatorSpec defines the desired state of K8sOperator
type K8sOperatorSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	Image    string `json:"image,omitempty"`
	Replicas int    `json:"replicas,omitempty"`
}

// K8sOperatorStatus defines the observed state of K8sOperator
type K8sOperatorStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// K8sOperator is the Schema for the k8soperators API
type K8sOperator struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   K8sOperatorSpec   `json:"spec,omitempty"`
	Status K8sOperatorStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// K8sOperatorList contains a list of K8sOperator
type K8sOperatorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []K8sOperator `json:"items"`
}

func init() {
	SchemeBuilder.Register(&K8sOperator{}, &K8sOperatorList{})
}
