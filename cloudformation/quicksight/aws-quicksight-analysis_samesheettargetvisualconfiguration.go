// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// Analysis_SameSheetTargetVisualConfiguration AWS CloudFormation Resource (AWS::QuickSight::Analysis.SameSheetTargetVisualConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-samesheettargetvisualconfiguration.html
type Analysis_SameSheetTargetVisualConfiguration struct {

	// TargetVisualOptions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-samesheettargetvisualconfiguration.html#cfn-quicksight-analysis-samesheettargetvisualconfiguration-targetvisualoptions
	TargetVisualOptions *string `json:"TargetVisualOptions,omitempty"`

	// TargetVisuals AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-samesheettargetvisualconfiguration.html#cfn-quicksight-analysis-samesheettargetvisualconfiguration-targetvisuals
	TargetVisuals []string `json:"TargetVisuals,omitempty"`

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
func (r *Analysis_SameSheetTargetVisualConfiguration) AWSCloudFormationType() string {
	return "AWS::QuickSight::Analysis.SameSheetTargetVisualConfiguration"
}
