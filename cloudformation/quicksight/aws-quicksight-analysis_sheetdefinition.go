// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// Analysis_SheetDefinition AWS CloudFormation Resource (AWS::QuickSight::Analysis.SheetDefinition)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-sheetdefinition.html
type Analysis_SheetDefinition[T any] struct {

	// ContentType AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-sheetdefinition.html#cfn-quicksight-analysis-sheetdefinition-contenttype
	ContentType *T `json:"ContentType,omitempty"`

	// Description AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-sheetdefinition.html#cfn-quicksight-analysis-sheetdefinition-description
	Description *T `json:"Description,omitempty"`

	// FilterControls AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-sheetdefinition.html#cfn-quicksight-analysis-sheetdefinition-filtercontrols
	FilterControls []Analysis_FilterControl[any] `json:"FilterControls,omitempty"`

	// Layouts AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-sheetdefinition.html#cfn-quicksight-analysis-sheetdefinition-layouts
	Layouts []Analysis_Layout[any] `json:"Layouts,omitempty"`

	// Name AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-sheetdefinition.html#cfn-quicksight-analysis-sheetdefinition-name
	Name *T `json:"Name,omitempty"`

	// ParameterControls AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-sheetdefinition.html#cfn-quicksight-analysis-sheetdefinition-parametercontrols
	ParameterControls []Analysis_ParameterControl[any] `json:"ParameterControls,omitempty"`

	// SheetControlLayouts AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-sheetdefinition.html#cfn-quicksight-analysis-sheetdefinition-sheetcontrollayouts
	SheetControlLayouts []Analysis_SheetControlLayout[any] `json:"SheetControlLayouts,omitempty"`

	// SheetId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-sheetdefinition.html#cfn-quicksight-analysis-sheetdefinition-sheetid
	SheetId T `json:"SheetId"`

	// TextBoxes AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-sheetdefinition.html#cfn-quicksight-analysis-sheetdefinition-textboxes
	TextBoxes []Analysis_SheetTextBox[any] `json:"TextBoxes,omitempty"`

	// Title AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-sheetdefinition.html#cfn-quicksight-analysis-sheetdefinition-title
	Title *T `json:"Title,omitempty"`

	// Visuals AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-sheetdefinition.html#cfn-quicksight-analysis-sheetdefinition-visuals
	Visuals []Analysis_Visual[any] `json:"Visuals,omitempty"`

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
func (r *Analysis_SheetDefinition[any]) AWSCloudFormationType() string {
	return "AWS::QuickSight::Analysis.SheetDefinition"
}
