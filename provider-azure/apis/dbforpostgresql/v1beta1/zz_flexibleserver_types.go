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

type FlexibleServerObservation struct {
	Fqdn *string `json:"fqdn,omitempty" tf:"fqdn,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	PublicNetworkAccessEnabled *bool `json:"publicNetworkAccessEnabled,omitempty" tf:"public_network_access_enabled,omitempty"`
}

type FlexibleServerParameters struct {

	// +kubebuilder:validation:Optional
	AdministratorLogin *string `json:"administratorLogin,omitempty" tf:"administrator_login,omitempty"`

	// +kubebuilder:validation:Optional
	AdministratorPasswordSecretRef *v1.SecretKeySelector `json:"administratorPasswordSecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	BackupRetentionDays *float64 `json:"backupRetentionDays,omitempty" tf:"backup_retention_days,omitempty"`

	// +kubebuilder:validation:Optional
	CreateMode *string `json:"createMode,omitempty" tf:"create_mode,omitempty"`

	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-azure/apis/network/v1beta1.Subnet
	// +crossplane:generate:reference:extractor=github.com/upbound/official-providers/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	DelegatedSubnetID *string `json:"delegatedSubnetId,omitempty" tf:"delegated_subnet_id,omitempty"`

	// +kubebuilder:validation:Optional
	DelegatedSubnetIDRef *v1.Reference `json:"delegatedSubnetIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	DelegatedSubnetIDSelector *v1.Selector `json:"delegatedSubnetIdSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	GeoRedundantBackupEnabled *bool `json:"geoRedundantBackupEnabled,omitempty" tf:"geo_redundant_backup_enabled,omitempty"`

	// +kubebuilder:validation:Optional
	HighAvailability []HighAvailabilityParameters `json:"highAvailability,omitempty" tf:"high_availability,omitempty"`

	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// +kubebuilder:validation:Optional
	MaintenanceWindow []MaintenanceWindowParameters `json:"maintenanceWindow,omitempty" tf:"maintenance_window,omitempty"`

	// +kubebuilder:validation:Optional
	PointInTimeRestoreTimeInUtc *string `json:"pointInTimeRestoreTimeInUtc,omitempty" tf:"point_in_time_restore_time_in_utc,omitempty"`

	// +kubebuilder:validation:Optional
	PrivateDNSZoneID *string `json:"privateDnsZoneId,omitempty" tf:"private_dns_zone_id,omitempty"`

	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	SkuName *string `json:"skuName,omitempty" tf:"sku_name,omitempty"`

	// +kubebuilder:validation:Optional
	SourceServerID *string `json:"sourceServerId,omitempty" tf:"source_server_id,omitempty"`

	// +kubebuilder:validation:Optional
	StorageMb *float64 `json:"storageMb,omitempty" tf:"storage_mb,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Optional
	Version *string `json:"version,omitempty" tf:"version,omitempty"`

	// +kubebuilder:validation:Optional
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

type HighAvailabilityObservation struct {
}

type HighAvailabilityParameters struct {

	// +kubebuilder:validation:Required
	Mode *string `json:"mode" tf:"mode,omitempty"`

	// +kubebuilder:validation:Optional
	StandbyAvailabilityZone *string `json:"standbyAvailabilityZone,omitempty" tf:"standby_availability_zone,omitempty"`
}

type MaintenanceWindowObservation struct {
}

type MaintenanceWindowParameters struct {

	// +kubebuilder:validation:Optional
	DayOfWeek *float64 `json:"dayOfWeek,omitempty" tf:"day_of_week,omitempty"`

	// +kubebuilder:validation:Optional
	StartHour *float64 `json:"startHour,omitempty" tf:"start_hour,omitempty"`

	// +kubebuilder:validation:Optional
	StartMinute *float64 `json:"startMinute,omitempty" tf:"start_minute,omitempty"`
}

// FlexibleServerSpec defines the desired state of FlexibleServer
type FlexibleServerSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FlexibleServerParameters `json:"forProvider"`
}

// FlexibleServerStatus defines the observed state of FlexibleServer.
type FlexibleServerStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FlexibleServerObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// FlexibleServer is the Schema for the FlexibleServers API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type FlexibleServer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FlexibleServerSpec   `json:"spec"`
	Status            FlexibleServerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FlexibleServerList contains a list of FlexibleServers
type FlexibleServerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FlexibleServer `json:"items"`
}

// Repository type metadata.
var (
	FlexibleServer_Kind             = "FlexibleServer"
	FlexibleServer_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: FlexibleServer_Kind}.String()
	FlexibleServer_KindAPIVersion   = FlexibleServer_Kind + "." + CRDGroupVersion.String()
	FlexibleServer_GroupVersionKind = CRDGroupVersion.WithKind(FlexibleServer_Kind)
)

func init() {
	SchemeBuilder.Register(&FlexibleServer{}, &FlexibleServerList{})
}