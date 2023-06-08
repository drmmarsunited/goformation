// Code generated by "go generate". Please don't change this file directly.

package evidently

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// Launch_SegmentOverride AWS CloudFormation Resource (AWS::Evidently::Launch.SegmentOverride)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-evidently-launch-segmentoverride.html
type Launch_SegmentOverride[T any] struct {

	// EvaluationOrder AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-evidently-launch-segmentoverride.html#cfn-evidently-launch-segmentoverride-evaluationorder
	EvaluationOrder T `json:"EvaluationOrder"`

	// Segment AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-evidently-launch-segmentoverride.html#cfn-evidently-launch-segmentoverride-segment
	Segment string `json:"Segment"`

	// Weights AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-evidently-launch-segmentoverride.html#cfn-evidently-launch-segmentoverride-weights
	Weights []Launch_GroupToWeight[any] `json:"Weights"`

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
func (r *Launch_SegmentOverride[any]) AWSCloudFormationType() string {
	return "AWS::Evidently::Launch.SegmentOverride"
}
