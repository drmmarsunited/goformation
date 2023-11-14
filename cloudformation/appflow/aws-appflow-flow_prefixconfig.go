// Code generated by "go generate". Please don't change this file directly.

package appflow

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// Flow_PrefixConfig AWS CloudFormation Resource (AWS::AppFlow::Flow.PrefixConfig)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-prefixconfig.html
type Flow_PrefixConfig[T any] struct {

	// PathPrefixHierarchy AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-prefixconfig.html#cfn-appflow-flow-prefixconfig-pathprefixhierarchy
	PathPrefixHierarchy []T `json:"PathPrefixHierarchy,omitempty"`

	// PrefixFormat AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-prefixconfig.html#cfn-appflow-flow-prefixconfig-prefixformat
	PrefixFormat *T `json:"PrefixFormat,omitempty"`

	// PrefixType AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-prefixconfig.html#cfn-appflow-flow-prefixconfig-prefixtype
	PrefixType *T `json:"PrefixType,omitempty"`

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
func (r *Flow_PrefixConfig[any]) AWSCloudFormationType() string {
	return "AWS::AppFlow::Flow.PrefixConfig"
}
