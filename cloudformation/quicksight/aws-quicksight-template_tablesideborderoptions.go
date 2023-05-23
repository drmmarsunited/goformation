// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// Template_TableSideBorderOptions AWS CloudFormation Resource (AWS::QuickSight::Template.TableSideBorderOptions)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-tablesideborderoptions.html
type Template_TableSideBorderOptions struct {

	// Bottom AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-tablesideborderoptions.html#cfn-quicksight-template-tablesideborderoptions-bottom
	Bottom *Template_TableBorderOptions `json:"Bottom,omitempty"`

	// InnerHorizontal AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-tablesideborderoptions.html#cfn-quicksight-template-tablesideborderoptions-innerhorizontal
	InnerHorizontal *Template_TableBorderOptions `json:"InnerHorizontal,omitempty"`

	// InnerVertical AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-tablesideborderoptions.html#cfn-quicksight-template-tablesideborderoptions-innervertical
	InnerVertical *Template_TableBorderOptions `json:"InnerVertical,omitempty"`

	// Left AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-tablesideborderoptions.html#cfn-quicksight-template-tablesideborderoptions-left
	Left *Template_TableBorderOptions `json:"Left,omitempty"`

	// Right AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-tablesideborderoptions.html#cfn-quicksight-template-tablesideborderoptions-right
	Right *Template_TableBorderOptions `json:"Right,omitempty"`

	// Top AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-tablesideborderoptions.html#cfn-quicksight-template-tablesideborderoptions-top
	Top *Template_TableBorderOptions `json:"Top,omitempty"`

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
func (r *Template_TableSideBorderOptions) AWSCloudFormationType() string {
	return "AWS::QuickSight::Template.TableSideBorderOptions"
}
