// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// Template_PercentageDisplayFormatConfiguration AWS CloudFormation Resource (AWS::QuickSight::Template.PercentageDisplayFormatConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-percentagedisplayformatconfiguration.html
type Template_PercentageDisplayFormatConfiguration[T any] struct {

	// DecimalPlacesConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-percentagedisplayformatconfiguration.html#cfn-quicksight-template-percentagedisplayformatconfiguration-decimalplacesconfiguration
	DecimalPlacesConfiguration *Template_DecimalPlacesConfiguration[any] `json:"DecimalPlacesConfiguration,omitempty"`

	// NegativeValueConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-percentagedisplayformatconfiguration.html#cfn-quicksight-template-percentagedisplayformatconfiguration-negativevalueconfiguration
	NegativeValueConfiguration *Template_NegativeValueConfiguration[any] `json:"NegativeValueConfiguration,omitempty"`

	// NullValueFormatConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-percentagedisplayformatconfiguration.html#cfn-quicksight-template-percentagedisplayformatconfiguration-nullvalueformatconfiguration
	NullValueFormatConfiguration *Template_NullValueFormatConfiguration[any] `json:"NullValueFormatConfiguration,omitempty"`

	// Prefix AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-percentagedisplayformatconfiguration.html#cfn-quicksight-template-percentagedisplayformatconfiguration-prefix
	Prefix *string `json:"Prefix,omitempty"`

	// SeparatorConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-percentagedisplayformatconfiguration.html#cfn-quicksight-template-percentagedisplayformatconfiguration-separatorconfiguration
	SeparatorConfiguration *Template_NumericSeparatorConfiguration[any] `json:"SeparatorConfiguration,omitempty"`

	// Suffix AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-percentagedisplayformatconfiguration.html#cfn-quicksight-template-percentagedisplayformatconfiguration-suffix
	Suffix *string `json:"Suffix,omitempty"`

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
func (r *Template_PercentageDisplayFormatConfiguration[any]) AWSCloudFormationType() string {
	return "AWS::QuickSight::Template.PercentageDisplayFormatConfiguration"
}
