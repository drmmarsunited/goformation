// Code generated by "go generate". Please don't change this file directly.

package pinpoint

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// InAppTemplate_OverrideButtonConfiguration AWS CloudFormation Resource (AWS::Pinpoint::InAppTemplate.OverrideButtonConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-inapptemplate-overridebuttonconfiguration.html
type InAppTemplate_OverrideButtonConfiguration[T any] struct {

	// ButtonAction AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-inapptemplate-overridebuttonconfiguration.html#cfn-pinpoint-inapptemplate-overridebuttonconfiguration-buttonaction
	ButtonAction *T `json:"ButtonAction,omitempty"`

	// Link AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-inapptemplate-overridebuttonconfiguration.html#cfn-pinpoint-inapptemplate-overridebuttonconfiguration-link
	Link *T `json:"Link,omitempty"`

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
func (r *InAppTemplate_OverrideButtonConfiguration[any]) AWSCloudFormationType() string {
	return "AWS::Pinpoint::InAppTemplate.OverrideButtonConfiguration"
}
