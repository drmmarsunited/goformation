// Code generated by "go generate". Please don't change this file directly.

package appflow

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// ConnectorProfile_CustomConnectorProfileProperties AWS CloudFormation Resource (AWS::AppFlow::ConnectorProfile.CustomConnectorProfileProperties)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-customconnectorprofileproperties.html
type ConnectorProfile_CustomConnectorProfileProperties[T any] struct {

	// OAuth2Properties AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-customconnectorprofileproperties.html#cfn-appflow-connectorprofile-customconnectorprofileproperties-oauth2properties
	OAuth2Properties *ConnectorProfile_OAuth2Properties[any] `json:"OAuth2Properties,omitempty"`

	// ProfileProperties AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-customconnectorprofileproperties.html#cfn-appflow-connectorprofile-customconnectorprofileproperties-profileproperties
	ProfileProperties map[string]T `json:"ProfileProperties,omitempty"`

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
func (r *ConnectorProfile_CustomConnectorProfileProperties[any]) AWSCloudFormationType() string {
	return "AWS::AppFlow::ConnectorProfile.CustomConnectorProfileProperties"
}
