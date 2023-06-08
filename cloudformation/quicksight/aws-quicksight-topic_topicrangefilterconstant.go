// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// Topic_TopicRangeFilterConstant AWS CloudFormation Resource (AWS::QuickSight::Topic.TopicRangeFilterConstant)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-topicrangefilterconstant.html
type Topic_TopicRangeFilterConstant[T any] struct {

	// ConstantType AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-topicrangefilterconstant.html#cfn-quicksight-topic-topicrangefilterconstant-constanttype
	ConstantType *string `json:"ConstantType,omitempty"`

	// RangeConstant AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-topicrangefilterconstant.html#cfn-quicksight-topic-topicrangefilterconstant-rangeconstant
	RangeConstant *Topic_RangeConstant[any] `json:"RangeConstant,omitempty"`

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
func (r *Topic_TopicRangeFilterConstant[any]) AWSCloudFormationType() string {
	return "AWS::QuickSight::Topic.TopicRangeFilterConstant"
}
