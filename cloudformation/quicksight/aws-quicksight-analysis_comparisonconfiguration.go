// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// Analysis_ComparisonConfiguration AWS CloudFormation Resource (AWS::QuickSight::Analysis.ComparisonConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-comparisonconfiguration.html
type Analysis_ComparisonConfiguration[T any] struct {

	// ComparisonFormat AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-comparisonconfiguration.html#cfn-quicksight-analysis-comparisonconfiguration-comparisonformat
	ComparisonFormat *Analysis_ComparisonFormatConfiguration[any] `json:"ComparisonFormat,omitempty"`

	// ComparisonMethod AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-comparisonconfiguration.html#cfn-quicksight-analysis-comparisonconfiguration-comparisonmethod
	ComparisonMethod *T `json:"ComparisonMethod,omitempty"`

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
func (r *Analysis_ComparisonConfiguration[any]) AWSCloudFormationType() string {
	return "AWS::QuickSight::Analysis.ComparisonConfiguration"
}
