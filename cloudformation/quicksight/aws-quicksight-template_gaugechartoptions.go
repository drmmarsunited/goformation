// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// Template_GaugeChartOptions AWS CloudFormation Resource (AWS::QuickSight::Template.GaugeChartOptions)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-gaugechartoptions.html
type Template_GaugeChartOptions[T any] struct {

	// Arc AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-gaugechartoptions.html#cfn-quicksight-template-gaugechartoptions-arc
	Arc *Template_ArcConfiguration[any] `json:"Arc,omitempty"`

	// ArcAxis AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-gaugechartoptions.html#cfn-quicksight-template-gaugechartoptions-arcaxis
	ArcAxis *Template_ArcAxisConfiguration[any] `json:"ArcAxis,omitempty"`

	// Comparison AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-gaugechartoptions.html#cfn-quicksight-template-gaugechartoptions-comparison
	Comparison *Template_ComparisonConfiguration[any] `json:"Comparison,omitempty"`

	// PrimaryValueDisplayType AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-gaugechartoptions.html#cfn-quicksight-template-gaugechartoptions-primaryvaluedisplaytype
	PrimaryValueDisplayType *T `json:"PrimaryValueDisplayType,omitempty"`

	// PrimaryValueFontConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-gaugechartoptions.html#cfn-quicksight-template-gaugechartoptions-primaryvaluefontconfiguration
	PrimaryValueFontConfiguration *Template_FontConfiguration[any] `json:"PrimaryValueFontConfiguration,omitempty"`

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
func (r *Template_GaugeChartOptions[any]) AWSCloudFormationType() string {
	return "AWS::QuickSight::Template.GaugeChartOptions"
}
