/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	"github.com/pkg/errors"

	"github.com/upbound/upjet/pkg/resource"
	"github.com/upbound/upjet/pkg/resource/json"
)

// GetTerraformResourceType returns Terraform resource type for this DesktopScalingPlan
func (mg *DesktopScalingPlan) GetTerraformResourceType() string {
	return "azurerm_virtual_desktop_scaling_plan"
}

// GetConnectionDetailsMapping for this DesktopScalingPlan
func (tr *DesktopScalingPlan) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this DesktopScalingPlan
func (tr *DesktopScalingPlan) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this DesktopScalingPlan
func (tr *DesktopScalingPlan) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this DesktopScalingPlan
func (tr *DesktopScalingPlan) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this DesktopScalingPlan
func (tr *DesktopScalingPlan) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this DesktopScalingPlan
func (tr *DesktopScalingPlan) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this DesktopScalingPlan using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *DesktopScalingPlan) LateInitialize(attrs []byte) (bool, error) {
	params := &DesktopScalingPlanParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *DesktopScalingPlan) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this NetworkGatewayNATRule
func (mg *NetworkGatewayNATRule) GetTerraformResourceType() string {
	return "azurerm_virtual_network_gateway_nat_rule"
}

// GetConnectionDetailsMapping for this NetworkGatewayNATRule
func (tr *NetworkGatewayNATRule) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this NetworkGatewayNATRule
func (tr *NetworkGatewayNATRule) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this NetworkGatewayNATRule
func (tr *NetworkGatewayNATRule) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this NetworkGatewayNATRule
func (tr *NetworkGatewayNATRule) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this NetworkGatewayNATRule
func (tr *NetworkGatewayNATRule) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this NetworkGatewayNATRule
func (tr *NetworkGatewayNATRule) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this NetworkGatewayNATRule using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *NetworkGatewayNATRule) LateInitialize(attrs []byte) (bool, error) {
	params := &NetworkGatewayNATRuleParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *NetworkGatewayNATRule) GetTerraformSchemaVersion() int {
	return 0
}