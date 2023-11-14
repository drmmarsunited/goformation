// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// Dashboard_TableCellStyle AWS CloudFormation Resource (AWS::QuickSight::Dashboard.TableCellStyle)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-tablecellstyle.html
type Dashboard_TableCellStyle[T any] struct {

	// BackgroundColor AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-tablecellstyle.html#cfn-quicksight-dashboard-tablecellstyle-backgroundcolor
	BackgroundColor *T `json:"BackgroundColor,omitempty"`

	// Border AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-tablecellstyle.html#cfn-quicksight-dashboard-tablecellstyle-border
	Border *Dashboard_GlobalTableBorderOptions[any] `json:"Border,omitempty"`

	// FontConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-tablecellstyle.html#cfn-quicksight-dashboard-tablecellstyle-fontconfiguration
	FontConfiguration *Dashboard_FontConfiguration[any] `json:"FontConfiguration,omitempty"`

	// Height AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-tablecellstyle.html#cfn-quicksight-dashboard-tablecellstyle-height
	Height *T `json:"Height,omitempty"`

	// HorizontalTextAlignment AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-tablecellstyle.html#cfn-quicksight-dashboard-tablecellstyle-horizontaltextalignment
	HorizontalTextAlignment *T `json:"HorizontalTextAlignment,omitempty"`

	// TextWrap AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-tablecellstyle.html#cfn-quicksight-dashboard-tablecellstyle-textwrap
	TextWrap *T `json:"TextWrap,omitempty"`

	// VerticalTextAlignment AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-tablecellstyle.html#cfn-quicksight-dashboard-tablecellstyle-verticaltextalignment
	VerticalTextAlignment *T `json:"VerticalTextAlignment,omitempty"`

	// Visibility AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-tablecellstyle.html#cfn-quicksight-dashboard-tablecellstyle-visibility
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
func (r *Dashboard_TableCellStyle[any]) AWSCloudFormationType() string {
	return "AWS::QuickSight::Dashboard.TableCellStyle"
}
