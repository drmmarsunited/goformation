// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// Topic_TopicRelativeDateFilter AWS CloudFormation Resource (AWS::QuickSight::Topic.TopicRelativeDateFilter)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-topicrelativedatefilter.html
type Topic_TopicRelativeDateFilter[T any] struct {

	// Constant AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-topicrelativedatefilter.html#cfn-quicksight-topic-topicrelativedatefilter-constant
	Constant *Topic_TopicSingularFilterConstant[any] `json:"Constant,omitempty"`

	// RelativeDateFilterFunction AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-topicrelativedatefilter.html#cfn-quicksight-topic-topicrelativedatefilter-relativedatefilterfunction
	RelativeDateFilterFunction *string `json:"RelativeDateFilterFunction,omitempty"`

	// TimeGranularity AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-topicrelativedatefilter.html#cfn-quicksight-topic-topicrelativedatefilter-timegranularity
	TimeGranularity *string `json:"TimeGranularity,omitempty"`

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
func (r *Topic_TopicRelativeDateFilter[any]) AWSCloudFormationType() string {
	return "AWS::QuickSight::Topic.TopicRelativeDateFilter"
}
