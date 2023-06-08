// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// Template_SheetDefinition AWS CloudFormation Resource (AWS::QuickSight::Template.SheetDefinition)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-sheetdefinition.html
type Template_SheetDefinition[T any] struct {

	// ContentType AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-sheetdefinition.html#cfn-quicksight-template-sheetdefinition-contenttype
	ContentType *string `json:"ContentType,omitempty"`

	// Description AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-sheetdefinition.html#cfn-quicksight-template-sheetdefinition-description
	Description *string `json:"Description,omitempty"`

	// FilterControls AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-sheetdefinition.html#cfn-quicksight-template-sheetdefinition-filtercontrols
	FilterControls []Template_FilterControl[any] `json:"FilterControls,omitempty"`

	// Layouts AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-sheetdefinition.html#cfn-quicksight-template-sheetdefinition-layouts
	Layouts []Template_Layout[any] `json:"Layouts,omitempty"`

	// Name AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-sheetdefinition.html#cfn-quicksight-template-sheetdefinition-name
	Name *string `json:"Name,omitempty"`

	// ParameterControls AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-sheetdefinition.html#cfn-quicksight-template-sheetdefinition-parametercontrols
	ParameterControls []Template_ParameterControl[any] `json:"ParameterControls,omitempty"`

	// SheetControlLayouts AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-sheetdefinition.html#cfn-quicksight-template-sheetdefinition-sheetcontrollayouts
	SheetControlLayouts []Template_SheetControlLayout[any] `json:"SheetControlLayouts,omitempty"`

	// SheetId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-sheetdefinition.html#cfn-quicksight-template-sheetdefinition-sheetid
	SheetId string `json:"SheetId"`

	// TextBoxes AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-sheetdefinition.html#cfn-quicksight-template-sheetdefinition-textboxes
	TextBoxes []Template_SheetTextBox[any] `json:"TextBoxes,omitempty"`

	// Title AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-sheetdefinition.html#cfn-quicksight-template-sheetdefinition-title
	Title *string `json:"Title,omitempty"`

	// Visuals AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-sheetdefinition.html#cfn-quicksight-template-sheetdefinition-visuals
	Visuals []Template_Visual[any] `json:"Visuals,omitempty"`

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
func (r *Template_SheetDefinition[any]) AWSCloudFormationType() string {
	return "AWS::QuickSight::Template.SheetDefinition"
}
