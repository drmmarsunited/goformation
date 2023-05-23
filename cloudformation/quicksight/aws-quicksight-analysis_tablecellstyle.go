// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// Analysis_TableCellStyle AWS CloudFormation Resource (AWS::QuickSight::Analysis.TableCellStyle)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-tablecellstyle.html
type Analysis_TableCellStyle struct {

	// BackgroundColor AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-tablecellstyle.html#cfn-quicksight-analysis-tablecellstyle-backgroundcolor
	BackgroundColor *string `json:"BackgroundColor,omitempty"`

	// Border AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-tablecellstyle.html#cfn-quicksight-analysis-tablecellstyle-border
	Border *Analysis_GlobalTableBorderOptions `json:"Border,omitempty"`

	// FontConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-tablecellstyle.html#cfn-quicksight-analysis-tablecellstyle-fontconfiguration
	FontConfiguration *Analysis_FontConfiguration `json:"FontConfiguration,omitempty"`

	// Height AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-tablecellstyle.html#cfn-quicksight-analysis-tablecellstyle-height
	Height *float64 `json:"Height,omitempty"`

	// HorizontalTextAlignment AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-tablecellstyle.html#cfn-quicksight-analysis-tablecellstyle-horizontaltextalignment
	HorizontalTextAlignment *string `json:"HorizontalTextAlignment,omitempty"`

	// TextWrap AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-tablecellstyle.html#cfn-quicksight-analysis-tablecellstyle-textwrap
	TextWrap *string `json:"TextWrap,omitempty"`

	// VerticalTextAlignment AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-tablecellstyle.html#cfn-quicksight-analysis-tablecellstyle-verticaltextalignment
	VerticalTextAlignment *string `json:"VerticalTextAlignment,omitempty"`

	// Visibility AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-tablecellstyle.html#cfn-quicksight-analysis-tablecellstyle-visibility
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
func (r *Analysis_TableCellStyle) AWSCloudFormationType() string {
	return "AWS::QuickSight::Analysis.TableCellStyle"
}
