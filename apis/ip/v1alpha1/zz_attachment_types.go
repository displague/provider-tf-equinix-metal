/*
Copyright 2021 The Crossplane Authors.

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

// Code generated by terrajet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type AttachmentObservation struct {
	Address *string `json:"address,omitempty" tf:"address,omitempty"`

	AddressFamily *int64 `json:"addressFamily,omitempty" tf:"address_family,omitempty"`

	Cidr *int64 `json:"cidr,omitempty" tf:"cidr,omitempty"`

	Gateway *string `json:"gateway,omitempty" tf:"gateway,omitempty"`

	Global *bool `json:"global,omitempty" tf:"global,omitempty"`

	Manageable *bool `json:"manageable,omitempty" tf:"manageable,omitempty"`

	Management *bool `json:"management,omitempty" tf:"management,omitempty"`

	Netmask *string `json:"netmask,omitempty" tf:"netmask,omitempty"`

	Network *string `json:"network,omitempty" tf:"network,omitempty"`

	Public *bool `json:"public,omitempty" tf:"public,omitempty"`
}

type AttachmentParameters struct {

	// +kubebuilder:validation:Required
	CidrNotation *string `json:"cidrNotation" tf:"cidr_notation,omitempty"`

	// +kubebuilder:validation:Required
	DeviceID *string `json:"deviceId" tf:"device_id,omitempty"`
}

// AttachmentSpec defines the desired state of Attachment
type AttachmentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AttachmentParameters `json:"forProvider"`
}

// AttachmentStatus defines the observed state of Attachment.
type AttachmentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AttachmentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Attachment is the Schema for the Attachments API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfequinixmetal}
type Attachment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AttachmentSpec   `json:"spec"`
	Status            AttachmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AttachmentList contains a list of Attachments
type AttachmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Attachment `json:"items"`
}

// Repository type metadata.
var (
	AttachmentKind             = "Attachment"
	AttachmentGroupKind        = schema.GroupKind{Group: Group, Kind: AttachmentKind}.String()
	AttachmentKindAPIVersion   = AttachmentKind + "." + GroupVersion.String()
	AttachmentGroupVersionKind = GroupVersion.WithKind(AttachmentKind)
)

func init() {
	SchemeBuilder.Register(&Attachment{}, &AttachmentList{})
}
