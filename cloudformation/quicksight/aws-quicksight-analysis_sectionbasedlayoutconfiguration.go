// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// Analysis_SectionBasedLayoutConfiguration AWS CloudFormation Resource (AWS::QuickSight::Analysis.SectionBasedLayoutConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-sectionbasedlayoutconfiguration.html
type Analysis_SectionBasedLayoutConfiguration[T any] struct {

	// BodySections AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-sectionbasedlayoutconfiguration.html#cfn-quicksight-analysis-sectionbasedlayoutconfiguration-bodysections
	BodySections []Analysis_BodySectionConfiguration[any] `json:"BodySections"`

	// CanvasSizeOptions AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-sectionbasedlayoutconfiguration.html#cfn-quicksight-analysis-sectionbasedlayoutconfiguration-canvassizeoptions
	CanvasSizeOptions *Analysis_SectionBasedLayoutCanvasSizeOptions[any] `json:"CanvasSizeOptions"`

	// FooterSections AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-sectionbasedlayoutconfiguration.html#cfn-quicksight-analysis-sectionbasedlayoutconfiguration-footersections
	FooterSections []Analysis_HeaderFooterSectionConfiguration[any] `json:"FooterSections"`

	// HeaderSections AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-sectionbasedlayoutconfiguration.html#cfn-quicksight-analysis-sectionbasedlayoutconfiguration-headersections
	HeaderSections []Analysis_HeaderFooterSectionConfiguration[any] `json:"HeaderSections"`

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
func (r *Analysis_SectionBasedLayoutConfiguration[any]) AWSCloudFormationType() string {
	return "AWS::QuickSight::Analysis.SectionBasedLayoutConfiguration"
}
