// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// Dashboard_UniqueValuesComputation AWS CloudFormation Resource (AWS::QuickSight::Dashboard.UniqueValuesComputation)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-uniquevaluescomputation.html
type Dashboard_UniqueValuesComputation struct {

	// Category AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-uniquevaluescomputation.html#cfn-quicksight-dashboard-uniquevaluescomputation-category
	Category *Dashboard_DimensionField `json:"Category"`

	// ComputationId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-uniquevaluescomputation.html#cfn-quicksight-dashboard-uniquevaluescomputation-computationid
	ComputationId string `json:"ComputationId"`

	// Name AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-uniquevaluescomputation.html#cfn-quicksight-dashboard-uniquevaluescomputation-name
	Name *string `json:"Name,omitempty"`

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
func (r *Dashboard_UniqueValuesComputation) AWSCloudFormationType() string {
	return "AWS::QuickSight::Dashboard.UniqueValuesComputation"
}
