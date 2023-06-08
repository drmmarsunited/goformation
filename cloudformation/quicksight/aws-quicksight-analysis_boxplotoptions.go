// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// Analysis_BoxPlotOptions AWS CloudFormation Resource (AWS::QuickSight::Analysis.BoxPlotOptions)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-boxplotoptions.html
type Analysis_BoxPlotOptions[T any] struct {

	// AllDataPointsVisibility AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-boxplotoptions.html#cfn-quicksight-analysis-boxplotoptions-alldatapointsvisibility
	AllDataPointsVisibility *string `json:"AllDataPointsVisibility,omitempty"`

	// OutlierVisibility AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-boxplotoptions.html#cfn-quicksight-analysis-boxplotoptions-outliervisibility
	OutlierVisibility *string `json:"OutlierVisibility,omitempty"`

	// StyleOptions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-boxplotoptions.html#cfn-quicksight-analysis-boxplotoptions-styleoptions
	StyleOptions *Analysis_BoxPlotStyleOptions[any] `json:"StyleOptions,omitempty"`

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
func (r *Analysis_BoxPlotOptions[any]) AWSCloudFormationType() string {
	return "AWS::QuickSight::Analysis.BoxPlotOptions"
}
