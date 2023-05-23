// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// Template_TopBottomMoversComputation AWS CloudFormation Resource (AWS::QuickSight::Template.TopBottomMoversComputation)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-topbottommoverscomputation.html
type Template_TopBottomMoversComputation struct {

	// Category AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-topbottommoverscomputation.html#cfn-quicksight-template-topbottommoverscomputation-category
	Category *Template_DimensionField `json:"Category"`

	// ComputationId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-topbottommoverscomputation.html#cfn-quicksight-template-topbottommoverscomputation-computationid
	ComputationId string `json:"ComputationId"`

	// MoverSize AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-topbottommoverscomputation.html#cfn-quicksight-template-topbottommoverscomputation-moversize
	MoverSize *float64 `json:"MoverSize,omitempty"`

	// Name AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-topbottommoverscomputation.html#cfn-quicksight-template-topbottommoverscomputation-name
	Name *string `json:"Name,omitempty"`

	// SortOrder AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-topbottommoverscomputation.html#cfn-quicksight-template-topbottommoverscomputation-sortorder
	SortOrder *string `json:"SortOrder,omitempty"`

	// Time AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-topbottommoverscomputation.html#cfn-quicksight-template-topbottommoverscomputation-time
	Time *Template_DimensionField `json:"Time"`

	// Type AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-topbottommoverscomputation.html#cfn-quicksight-template-topbottommoverscomputation-type
	Type string `json:"Type"`

	// Value AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-topbottommoverscomputation.html#cfn-quicksight-template-topbottommoverscomputation-value
	Value *Template_MeasureField `json:"Value,omitempty"`

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
func (r *Template_TopBottomMoversComputation) AWSCloudFormationType() string {
	return "AWS::QuickSight::Template.TopBottomMoversComputation"
}
