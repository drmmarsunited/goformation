// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// Analysis_BodySectionConfiguration AWS CloudFormation Resource (AWS::QuickSight::Analysis.BodySectionConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-bodysectionconfiguration.html
type Analysis_BodySectionConfiguration[T any] struct {

	// Content AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-bodysectionconfiguration.html#cfn-quicksight-analysis-bodysectionconfiguration-content
	Content *Analysis_BodySectionContent[any] `json:"Content"`

	// PageBreakConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-bodysectionconfiguration.html#cfn-quicksight-analysis-bodysectionconfiguration-pagebreakconfiguration
	PageBreakConfiguration *Analysis_SectionPageBreakConfiguration[any] `json:"PageBreakConfiguration,omitempty"`

	// SectionId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-bodysectionconfiguration.html#cfn-quicksight-analysis-bodysectionconfiguration-sectionid
	SectionId T `json:"SectionId"`

	// Style AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-bodysectionconfiguration.html#cfn-quicksight-analysis-bodysectionconfiguration-style
	Style *Analysis_SectionStyle[any] `json:"Style,omitempty"`

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
func (r *Analysis_BodySectionConfiguration[any]) AWSCloudFormationType() string {
	return "AWS::QuickSight::Analysis.BodySectionConfiguration"
}
