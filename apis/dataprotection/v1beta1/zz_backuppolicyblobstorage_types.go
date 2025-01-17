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

type BackupPolicyBlobStorageObservation struct {

	// The ID of the Backup Policy Blob Storage.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Duration of deletion after given timespan. It should follow ISO 8601 duration format. Changing this forces a new Backup Policy Blob Storage to be created.
	RetentionDuration *string `json:"retentionDuration,omitempty" tf:"retention_duration,omitempty"`

	// The ID of the Backup Vault within which the Backup Policy Blob Storage should exist. Changing this forces a new Backup Policy Blob Storage to be created.
	VaultID *string `json:"vaultId,omitempty" tf:"vault_id,omitempty"`
}

type BackupPolicyBlobStorageParameters struct {

	// Duration of deletion after given timespan. It should follow ISO 8601 duration format. Changing this forces a new Backup Policy Blob Storage to be created.
	// +kubebuilder:validation:Optional
	RetentionDuration *string `json:"retentionDuration,omitempty" tf:"retention_duration,omitempty"`

	// The ID of the Backup Vault within which the Backup Policy Blob Storage should exist. Changing this forces a new Backup Policy Blob Storage to be created.
	// +crossplane:generate:reference:type=BackupVault
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	VaultID *string `json:"vaultId,omitempty" tf:"vault_id,omitempty"`

	// Reference to a BackupVault to populate vaultId.
	// +kubebuilder:validation:Optional
	VaultIDRef *v1.Reference `json:"vaultIdRef,omitempty" tf:"-"`

	// Selector for a BackupVault to populate vaultId.
	// +kubebuilder:validation:Optional
	VaultIDSelector *v1.Selector `json:"vaultIdSelector,omitempty" tf:"-"`
}

// BackupPolicyBlobStorageSpec defines the desired state of BackupPolicyBlobStorage
type BackupPolicyBlobStorageSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BackupPolicyBlobStorageParameters `json:"forProvider"`
}

// BackupPolicyBlobStorageStatus defines the observed state of BackupPolicyBlobStorage.
type BackupPolicyBlobStorageStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BackupPolicyBlobStorageObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// BackupPolicyBlobStorage is the Schema for the BackupPolicyBlobStorages API. Manages a Backup Policy Blob Storage.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type BackupPolicyBlobStorage struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.retentionDuration)",message="retentionDuration is a required parameter"
	Spec   BackupPolicyBlobStorageSpec   `json:"spec"`
	Status BackupPolicyBlobStorageStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BackupPolicyBlobStorageList contains a list of BackupPolicyBlobStorages
type BackupPolicyBlobStorageList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BackupPolicyBlobStorage `json:"items"`
}

// Repository type metadata.
var (
	BackupPolicyBlobStorage_Kind             = "BackupPolicyBlobStorage"
	BackupPolicyBlobStorage_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: BackupPolicyBlobStorage_Kind}.String()
	BackupPolicyBlobStorage_KindAPIVersion   = BackupPolicyBlobStorage_Kind + "." + CRDGroupVersion.String()
	BackupPolicyBlobStorage_GroupVersionKind = CRDGroupVersion.WithKind(BackupPolicyBlobStorage_Kind)
)

func init() {
	SchemeBuilder.Register(&BackupPolicyBlobStorage{}, &BackupPolicyBlobStorageList{})
}
