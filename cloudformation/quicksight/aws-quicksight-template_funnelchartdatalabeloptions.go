// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// Template_FunnelChartDataLabelOptions AWS CloudFormation Resource (AWS::QuickSight::Template.FunnelChartDataLabelOptions)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-funnelchartdatalabeloptions.html
type Template_FunnelChartDataLabelOptions[T any] struct {

	// CategoryLabelVisibility AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-funnelchartdatalabeloptions.html#cfn-quicksight-template-funnelchartdatalabeloptions-categorylabelvisibility
	CategoryLabelVisibility *T `json:"CategoryLabelVisibility,omitempty"`

	// LabelColor AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-funnelchartdatalabeloptions.html#cfn-quicksight-template-funnelchartdatalabeloptions-labelcolor
	LabelColor *T `json:"LabelColor,omitempty"`

	// LabelFontConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-funnelchartdatalabeloptions.html#cfn-quicksight-template-funnelchartdatalabeloptions-labelfontconfiguration
	LabelFontConfiguration *Template_FontConfiguration[any] `json:"LabelFontConfiguration,omitempty"`

	// MeasureDataLabelStyle AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-funnelchartdatalabeloptions.html#cfn-quicksight-template-funnelchartdatalabeloptions-measuredatalabelstyle
	MeasureDataLabelStyle *T `json:"MeasureDataLabelStyle,omitempty"`

	// MeasureLabelVisibility AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-funnelchartdatalabeloptions.html#cfn-quicksight-template-funnelchartdatalabeloptions-measurelabelvisibility
	MeasureLabelVisibility *T `json:"MeasureLabelVisibility,omitempty"`

	// Position AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-funnelchartdatalabeloptions.html#cfn-quicksight-template-funnelchartdatalabeloptions-position
	Position *T `json:"Position,omitempty"`

	// Visibility AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-funnelchartdatalabeloptions.html#cfn-quicksight-template-funnelchartdatalabeloptions-visibility
	Visibility *T `json:"Visibility,omitempty"`

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
func (r *Template_FunnelChartDataLabelOptions[any]) AWSCloudFormationType() string {
	return "AWS::QuickSight::Template.FunnelChartDataLabelOptions"
}
