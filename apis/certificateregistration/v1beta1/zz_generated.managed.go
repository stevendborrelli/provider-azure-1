/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1beta1

import xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"

// GetCondition of this AppServiceCertificateOrder.
func (mg *AppServiceCertificateOrder) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this AppServiceCertificateOrder.
func (mg *AppServiceCertificateOrder) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this AppServiceCertificateOrder.
func (mg *AppServiceCertificateOrder) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this AppServiceCertificateOrder.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *AppServiceCertificateOrder) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetPublishConnectionDetailsTo of this AppServiceCertificateOrder.
func (mg *AppServiceCertificateOrder) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this AppServiceCertificateOrder.
func (mg *AppServiceCertificateOrder) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this AppServiceCertificateOrder.
func (mg *AppServiceCertificateOrder) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this AppServiceCertificateOrder.
func (mg *AppServiceCertificateOrder) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this AppServiceCertificateOrder.
func (mg *AppServiceCertificateOrder) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this AppServiceCertificateOrder.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *AppServiceCertificateOrder) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetPublishConnectionDetailsTo of this AppServiceCertificateOrder.
func (mg *AppServiceCertificateOrder) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this AppServiceCertificateOrder.
func (mg *AppServiceCertificateOrder) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}
