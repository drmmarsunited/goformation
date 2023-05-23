// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// Template_PivotTableOptions AWS CloudFormation Resource (AWS::QuickSight::Template.PivotTableOptions)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-pivottableoptions.html
type Template_PivotTableOptions struct {

	// CellStyle AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-pivottableoptions.html#cfn-quicksight-template-pivottableoptions-cellstyle
	CellStyle *Template_TableCellStyle `json:"CellStyle,omitempty"`

	// ColumnHeaderStyle AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-pivottableoptions.html#cfn-quicksight-template-pivottableoptions-columnheaderstyle
	ColumnHeaderStyle *Template_TableCellStyle `json:"ColumnHeaderStyle,omitempty"`

	// ColumnNamesVisibility AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-pivottableoptions.html#cfn-quicksight-template-pivottableoptions-columnnamesvisibility
	ColumnNamesVisibility *string `json:"ColumnNamesVisibility,omitempty"`

	// MetricPlacement AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-pivottableoptions.html#cfn-quicksight-template-pivottableoptions-metricplacement
	MetricPlacement *string `json:"MetricPlacement,omitempty"`

	// RowAlternateColorOptions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-pivottableoptions.html#cfn-quicksight-template-pivottableoptions-rowalternatecoloroptions
	RowAlternateColorOptions *Template_RowAlternateColorOptions `json:"RowAlternateColorOptions,omitempty"`

	// RowFieldNamesStyle AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-pivottableoptions.html#cfn-quicksight-template-pivottableoptions-rowfieldnamesstyle
	RowFieldNamesStyle *Template_TableCellStyle `json:"RowFieldNamesStyle,omitempty"`

	// RowHeaderStyle AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-pivottableoptions.html#cfn-quicksight-template-pivottableoptions-rowheaderstyle
	RowHeaderStyle *Template_TableCellStyle `json:"RowHeaderStyle,omitempty"`

	// SingleMetricVisibility AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-pivottableoptions.html#cfn-quicksight-template-pivottableoptions-singlemetricvisibility
	SingleMetricVisibility *string `json:"SingleMetricVisibility,omitempty"`

	// ToggleButtonsVisibility AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-pivottableoptions.html#cfn-quicksight-template-pivottableoptions-togglebuttonsvisibility
	ToggleButtonsVisibility *string `json:"ToggleButtonsVisibility,omitempty"`

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
func (r *Template_PivotTableOptions) AWSCloudFormationType() string {
	return "AWS::QuickSight::Template.PivotTableOptions"
}
