// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// Template_ColumnConfiguration AWS CloudFormation Resource (AWS::QuickSight::Template.ColumnConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-columnconfiguration.html
type Template_ColumnConfiguration[T any] struct {

	// ColorsConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-columnconfiguration.html#cfn-quicksight-template-columnconfiguration-colorsconfiguration
	ColorsConfiguration *Template_ColorsConfiguration[any] `json:"ColorsConfiguration,omitempty"`

	// Column AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-columnconfiguration.html#cfn-quicksight-template-columnconfiguration-column
	Column *Template_ColumnIdentifier[any] `json:"Column"`

	// FormatConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-columnconfiguration.html#cfn-quicksight-template-columnconfiguration-formatconfiguration
	FormatConfiguration *Template_FormatConfiguration[any] `json:"FormatConfiguration,omitempty"`

	// Role AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-columnconfiguration.html#cfn-quicksight-template-columnconfiguration-role
	Role *T `json:"Role,omitempty"`

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
func (r *Template_ColumnConfiguration[any]) AWSCloudFormationType() string {
	return "AWS::QuickSight::Template.ColumnConfiguration"
}
