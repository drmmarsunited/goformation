// Code generated by "go generate". Please don't change this file directly.

package grafana

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// Workspace_SamlConfiguration AWS CloudFormation Resource (AWS::Grafana::Workspace.SamlConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-grafana-workspace-samlconfiguration.html
type Workspace_SamlConfiguration struct {

	// AllowedOrganizations AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-grafana-workspace-samlconfiguration.html#cfn-grafana-workspace-samlconfiguration-allowedorganizations
	AllowedOrganizations []string `json:"AllowedOrganizations,omitempty"`

	// AssertionAttributes AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-grafana-workspace-samlconfiguration.html#cfn-grafana-workspace-samlconfiguration-assertionattributes
	AssertionAttributes *Workspace_AssertionAttributes `json:"AssertionAttributes,omitempty"`

	// IdpMetadata AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-grafana-workspace-samlconfiguration.html#cfn-grafana-workspace-samlconfiguration-idpmetadata
	IdpMetadata *Workspace_IdpMetadata `json:"IdpMetadata"`

	// LoginValidityDuration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-grafana-workspace-samlconfiguration.html#cfn-grafana-workspace-samlconfiguration-loginvalidityduration
	LoginValidityDuration *float64 `json:"LoginValidityDuration,omitempty"`

	// RoleValues AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-grafana-workspace-samlconfiguration.html#cfn-grafana-workspace-samlconfiguration-rolevalues
	RoleValues *Workspace_RoleValues `json:"RoleValues,omitempty"`

	// AWSCloudFormationDeletionPolicy represents a CloudFormation DeletionPolicy
	AWSCloudFormationDeletionPolicy policies.DeletionPolicy `json:"-"`

	// AWSCloudFormationUpdateReplacePolicy represents a CloudFormation UpdateReplacePolicy
	AWSCloudFormationUpdateReplacePolicy policies.UpdateReplacePolicy `json:"-"`

	// AWSCloudFormationDependsOn stores the logical ID of the resources to be created before this resource
	AWSCloudFormationDependsOn []string `json:"-"`

	// AWSCloudFormationMetadata stores structured data associated with this resource
	AWSCloudFormationMetadata map[string]interface{} `json:"-"`

	// AWSCloudFormationCondition stores the logical ID of the condition that must be satisfied for this resource to be created
	AWSCloudFormationCondition string `json:"-"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *Workspace_SamlConfiguration) AWSCloudFormationType() string {
	return "AWS::Grafana::Workspace.SamlConfiguration"
}
