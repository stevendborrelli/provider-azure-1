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

type APIOperationTagObservation struct {

	// The ID of the API Management API Operation. Changing this forces a new API Management API Operation Tag to be created.
	APIOperationID *string `json:"apiOperationId,omitempty" tf:"api_operation_id,omitempty"`

	// The display name of the API Management API Operation Tag.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// The ID of the API Management API Operation Tag.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type APIOperationTagParameters struct {

	// The ID of the API Management API Operation. Changing this forces a new API Management API Operation Tag to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/apimanagement/v1beta1.APIOperation
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	APIOperationID *string `json:"apiOperationId,omitempty" tf:"api_operation_id,omitempty"`

	// Reference to a APIOperation in apimanagement to populate apiOperationId.
	// +kubebuilder:validation:Optional
	APIOperationIDRef *v1.Reference `json:"apiOperationIdRef,omitempty" tf:"-"`

	// Selector for a APIOperation in apimanagement to populate apiOperationId.
	// +kubebuilder:validation:Optional
	APIOperationIDSelector *v1.Selector `json:"apiOperationIdSelector,omitempty" tf:"-"`

	// The display name of the API Management API Operation Tag.
	// +kubebuilder:validation:Optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`
}

// APIOperationTagSpec defines the desired state of APIOperationTag
type APIOperationTagSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     APIOperationTagParameters `json:"forProvider"`
}

// APIOperationTagStatus defines the observed state of APIOperationTag.
type APIOperationTagStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        APIOperationTagObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// APIOperationTag is the Schema for the APIOperationTags API. Manages a API Management API Operation Tag.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type APIOperationTag struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.displayName)",message="displayName is a required parameter"
	Spec   APIOperationTagSpec   `json:"spec"`
	Status APIOperationTagStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// APIOperationTagList contains a list of APIOperationTags
type APIOperationTagList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []APIOperationTag `json:"items"`
}

// Repository type metadata.
var (
	APIOperationTag_Kind             = "APIOperationTag"
	APIOperationTag_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: APIOperationTag_Kind}.String()
	APIOperationTag_KindAPIVersion   = APIOperationTag_Kind + "." + CRDGroupVersion.String()
	APIOperationTag_GroupVersionKind = CRDGroupVersion.WithKind(APIOperationTag_Kind)
)

func init() {
	SchemeBuilder.Register(&APIOperationTag{}, &APIOperationTagList{})
}
