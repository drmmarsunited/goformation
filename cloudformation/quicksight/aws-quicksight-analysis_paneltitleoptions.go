// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// Analysis_PanelTitleOptions AWS CloudFormation Resource (AWS::QuickSight::Analysis.PanelTitleOptions)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-paneltitleoptions.html
type Analysis_PanelTitleOptions struct {

	// FontConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-paneltitleoptions.html#cfn-quicksight-analysis-paneltitleoptions-fontconfiguration
	FontConfiguration *Analysis_FontConfiguration `json:"FontConfiguration,omitempty"`

	// HorizontalTextAlignment AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-paneltitleoptions.html#cfn-quicksight-analysis-paneltitleoptions-horizontaltextalignment
	HorizontalTextAlignment *string `json:"HorizontalTextAlignment,omitempty"`

	// Visibility AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-paneltitleoptions.html#cfn-quicksight-analysis-paneltitleoptions-visibility
	Visibility *string `json:"Visibility,omitempty"`

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
func (r *Analysis_PanelTitleOptions) AWSCloudFormationType() string {
	return "AWS::QuickSight::Analysis.PanelTitleOptions"
}
