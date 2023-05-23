// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// Template_ReferenceLineValueLabelConfiguration AWS CloudFormation Resource (AWS::QuickSight::Template.ReferenceLineValueLabelConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-referencelinevaluelabelconfiguration.html
type Template_ReferenceLineValueLabelConfiguration struct {

	// FormatConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-referencelinevaluelabelconfiguration.html#cfn-quicksight-template-referencelinevaluelabelconfiguration-formatconfiguration
	FormatConfiguration *Template_NumericFormatConfiguration `json:"FormatConfiguration,omitempty"`

	// RelativePosition AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-referencelinevaluelabelconfiguration.html#cfn-quicksight-template-referencelinevaluelabelconfiguration-relativeposition
	RelativePosition *string `json:"RelativePosition,omitempty"`

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
func (r *Template_ReferenceLineValueLabelConfiguration) AWSCloudFormationType() string {
	return "AWS::QuickSight::Template.ReferenceLineValueLabelConfiguration"
}
