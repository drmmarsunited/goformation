// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// Analysis_LineChartSeriesSettings AWS CloudFormation Resource (AWS::QuickSight::Analysis.LineChartSeriesSettings)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-linechartseriessettings.html
type Analysis_LineChartSeriesSettings[T any] struct {

	// LineStyleSettings AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-linechartseriessettings.html#cfn-quicksight-analysis-linechartseriessettings-linestylesettings
	LineStyleSettings *Analysis_LineChartLineStyleSettings[any] `json:"LineStyleSettings,omitempty"`

	// MarkerStyleSettings AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-linechartseriessettings.html#cfn-quicksight-analysis-linechartseriessettings-markerstylesettings
	MarkerStyleSettings *Analysis_LineChartMarkerStyleSettings[any] `json:"MarkerStyleSettings,omitempty"`

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
func (r *Analysis_LineChartSeriesSettings[any]) AWSCloudFormationType() string {
	return "AWS::QuickSight::Analysis.LineChartSeriesSettings"
}
