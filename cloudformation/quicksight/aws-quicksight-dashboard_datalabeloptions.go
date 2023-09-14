// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// Dashboard_DataLabelOptions AWS CloudFormation Resource (AWS::QuickSight::Dashboard.DataLabelOptions)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-datalabeloptions.html
type Dashboard_DataLabelOptions[T any] struct {

	// CategoryLabelVisibility AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-datalabeloptions.html#cfn-quicksight-dashboard-datalabeloptions-categorylabelvisibility
	CategoryLabelVisibility *string `json:"CategoryLabelVisibility,omitempty"`

	// DataLabelTypes AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-datalabeloptions.html#cfn-quicksight-dashboard-datalabeloptions-datalabeltypes
	DataLabelTypes []Dashboard_DataLabelType[any] `json:"DataLabelTypes,omitempty"`

	// LabelColor AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-datalabeloptions.html#cfn-quicksight-dashboard-datalabeloptions-labelcolor
	LabelColor *string `json:"LabelColor,omitempty"`

	// LabelContent AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-datalabeloptions.html#cfn-quicksight-dashboard-datalabeloptions-labelcontent
	LabelContent *string `json:"LabelContent,omitempty"`

	// LabelFontConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-datalabeloptions.html#cfn-quicksight-dashboard-datalabeloptions-labelfontconfiguration
	LabelFontConfiguration *Dashboard_FontConfiguration[any] `json:"LabelFontConfiguration,omitempty"`

	// MeasureLabelVisibility AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-datalabeloptions.html#cfn-quicksight-dashboard-datalabeloptions-measurelabelvisibility
	MeasureLabelVisibility *string `json:"MeasureLabelVisibility,omitempty"`

	// Overlap AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-datalabeloptions.html#cfn-quicksight-dashboard-datalabeloptions-overlap
	Overlap *string `json:"Overlap,omitempty"`

	// Position AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-datalabeloptions.html#cfn-quicksight-dashboard-datalabeloptions-position
	Position *string `json:"Position,omitempty"`

	// TotalsVisibility AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-datalabeloptions.html#cfn-quicksight-dashboard-datalabeloptions-totalsvisibility
	TotalsVisibility *string `json:"TotalsVisibility,omitempty"`

	// Visibility AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-datalabeloptions.html#cfn-quicksight-dashboard-datalabeloptions-visibility
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
func (r *Dashboard_DataLabelOptions[any]) AWSCloudFormationType() string {
	return "AWS::QuickSight::Dashboard.DataLabelOptions"
}
