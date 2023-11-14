// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// Dashboard_NumberDisplayFormatConfiguration AWS CloudFormation Resource (AWS::QuickSight::Dashboard.NumberDisplayFormatConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-numberdisplayformatconfiguration.html
type Dashboard_NumberDisplayFormatConfiguration[T any] struct {

	// DecimalPlacesConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-numberdisplayformatconfiguration.html#cfn-quicksight-dashboard-numberdisplayformatconfiguration-decimalplacesconfiguration
	DecimalPlacesConfiguration *Dashboard_DecimalPlacesConfiguration[any] `json:"DecimalPlacesConfiguration,omitempty"`

	// NegativeValueConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-numberdisplayformatconfiguration.html#cfn-quicksight-dashboard-numberdisplayformatconfiguration-negativevalueconfiguration
	NegativeValueConfiguration *Dashboard_NegativeValueConfiguration[any] `json:"NegativeValueConfiguration,omitempty"`

	// NullValueFormatConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-numberdisplayformatconfiguration.html#cfn-quicksight-dashboard-numberdisplayformatconfiguration-nullvalueformatconfiguration
	NullValueFormatConfiguration *Dashboard_NullValueFormatConfiguration[any] `json:"NullValueFormatConfiguration,omitempty"`

	// NumberScale AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-numberdisplayformatconfiguration.html#cfn-quicksight-dashboard-numberdisplayformatconfiguration-numberscale
	NumberScale *T `json:"NumberScale,omitempty"`

	// Prefix AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-numberdisplayformatconfiguration.html#cfn-quicksight-dashboard-numberdisplayformatconfiguration-prefix
	Prefix *T `json:"Prefix,omitempty"`

	// SeparatorConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-numberdisplayformatconfiguration.html#cfn-quicksight-dashboard-numberdisplayformatconfiguration-separatorconfiguration
	SeparatorConfiguration *Dashboard_NumericSeparatorConfiguration[any] `json:"SeparatorConfiguration,omitempty"`

	// Suffix AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-numberdisplayformatconfiguration.html#cfn-quicksight-dashboard-numberdisplayformatconfiguration-suffix
	Suffix *T `json:"Suffix,omitempty"`

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
func (r *Dashboard_NumberDisplayFormatConfiguration[any]) AWSCloudFormationType() string {
	return "AWS::QuickSight::Dashboard.NumberDisplayFormatConfiguration"
}
