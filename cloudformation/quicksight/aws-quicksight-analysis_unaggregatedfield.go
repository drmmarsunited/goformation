// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// Analysis_UnaggregatedField AWS CloudFormation Resource (AWS::QuickSight::Analysis.UnaggregatedField)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-unaggregatedfield.html
type Analysis_UnaggregatedField[T any] struct {

	// Column AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-unaggregatedfield.html#cfn-quicksight-analysis-unaggregatedfield-column
	Column *Analysis_ColumnIdentifier[any] `json:"Column"`

	// FieldId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-unaggregatedfield.html#cfn-quicksight-analysis-unaggregatedfield-fieldid
	FieldId T `json:"FieldId"`

	// FormatConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-unaggregatedfield.html#cfn-quicksight-analysis-unaggregatedfield-formatconfiguration
	FormatConfiguration *Analysis_FormatConfiguration[any] `json:"FormatConfiguration,omitempty"`

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
func (r *Analysis_UnaggregatedField[any]) AWSCloudFormationType() string {
	return "AWS::QuickSight::Analysis.UnaggregatedField"
}
