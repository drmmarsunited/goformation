// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// Dashboard_CustomFilterConfiguration AWS CloudFormation Resource (AWS::QuickSight::Dashboard.CustomFilterConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-customfilterconfiguration.html
type Dashboard_CustomFilterConfiguration[T any] struct {

	// CategoryValue AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-customfilterconfiguration.html#cfn-quicksight-dashboard-customfilterconfiguration-categoryvalue
	CategoryValue *T `json:"CategoryValue,omitempty"`

	// MatchOperator AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-customfilterconfiguration.html#cfn-quicksight-dashboard-customfilterconfiguration-matchoperator
	MatchOperator T `json:"MatchOperator"`

	// NullOption AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-customfilterconfiguration.html#cfn-quicksight-dashboard-customfilterconfiguration-nulloption
	NullOption T `json:"NullOption"`

	// ParameterName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-customfilterconfiguration.html#cfn-quicksight-dashboard-customfilterconfiguration-parametername
	ParameterName *T `json:"ParameterName,omitempty"`

	// SelectAllOptions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-customfilterconfiguration.html#cfn-quicksight-dashboard-customfilterconfiguration-selectalloptions
	SelectAllOptions *T `json:"SelectAllOptions,omitempty"`

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
func (r *Dashboard_CustomFilterConfiguration[any]) AWSCloudFormationType() string {
	return "AWS::QuickSight::Dashboard.CustomFilterConfiguration"
}
