/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type FeatureObservation struct {

	// Specifies the name of the feature to register.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Should this feature be Registered or Unregistered?
	Registered *bool `json:"registered,omitempty" tf:"registered,omitempty"`
}

type FeatureParameters struct {

	// Specifies the name of the feature to register.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Should this feature be Registered or Unregistered?
	// +kubebuilder:validation:Required
	Registered *bool `json:"registered" tf:"registered,omitempty"`
}

type ResourceProviderRegistrationObservation struct {

	// A list of feature blocks as defined below.
	Feature []FeatureObservation `json:"feature,omitempty" tf:"feature,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The namespace of the Resource Provider which should be registered. Changing this forces a new resource to be created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type ResourceProviderRegistrationParameters struct {

	// A list of feature blocks as defined below.
	// +kubebuilder:validation:Optional
	Feature []FeatureParameters `json:"feature,omitempty" tf:"feature,omitempty"`

	// The namespace of the Resource Provider which should be registered. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

// ResourceProviderRegistrationSpec defines the desired state of ResourceProviderRegistration
type ResourceProviderRegistrationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ResourceProviderRegistrationParameters `json:"forProvider"`
}

// ResourceProviderRegistrationStatus defines the observed state of ResourceProviderRegistration.
type ResourceProviderRegistrationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ResourceProviderRegistrationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ResourceProviderRegistration is the Schema for the ResourceProviderRegistrations API. Manages the Registration of a Resource Provider.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type ResourceProviderRegistration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.name)",message="name is a required parameter"
	Spec   ResourceProviderRegistrationSpec   `json:"spec"`
	Status ResourceProviderRegistrationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ResourceProviderRegistrationList contains a list of ResourceProviderRegistrations
type ResourceProviderRegistrationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ResourceProviderRegistration `json:"items"`
}

// Repository type metadata.
var (
	ResourceProviderRegistration_Kind             = "ResourceProviderRegistration"
	ResourceProviderRegistration_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ResourceProviderRegistration_Kind}.String()
	ResourceProviderRegistration_KindAPIVersion   = ResourceProviderRegistration_Kind + "." + CRDGroupVersion.String()
	ResourceProviderRegistration_GroupVersionKind = CRDGroupVersion.WithKind(ResourceProviderRegistration_Kind)
)

func init() {
	SchemeBuilder.Register(&ResourceProviderRegistration{}, &ResourceProviderRegistrationList{})
}
