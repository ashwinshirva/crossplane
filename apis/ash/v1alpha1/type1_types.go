/*
Copyright 2022 The Crossplane Authors.

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

package v1alpha1

import (
	"reflect"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

// type1Parameters are the configurable fields of a type1.
type type1Parameters struct {
	ConfigurableField string `json:"configurableField"`
}

// type1Observation are the observable fields of a type1.
type type1Observation struct {
	ObservableField string `json:"observableField,omitempty"`
}

// A type1Spec defines the desired state of a type1.
type type1Spec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       type1Parameters `json:"forProvider"`
}

// A type1Status represents the observed state of a type1.
type type1Status struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          type1Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// A type1 is an example API type.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,ashprovider}
type type1 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   type1Spec   `json:"spec"`
	Status type1Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// type1List contains a list of type1
type type1List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []type1 `json:"items"`
}

// type1 type metadata.
var (
	type1Kind             = reflect.TypeOf(type1{}).Name()
	type1GroupKind        = schema.GroupKind{Group: Group, Kind: type1Kind}.String()
	type1KindAPIVersion   = type1Kind + "." + SchemeGroupVersion.String()
	type1GroupVersionKind = SchemeGroupVersion.WithKind(type1Kind)
)

func init() {
	SchemeBuilder.Register(&type1{}, &type1List{})
}
