// Code generated by "go generate". Please don't change this file directly.

package cleanrooms

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// ConfiguredTable_AggregateColumn AWS CloudFormation Resource (AWS::CleanRooms::ConfiguredTable.AggregateColumn)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-configuredtable-aggregatecolumn.html
type ConfiguredTable_AggregateColumn[T any] struct {

	// ColumnNames AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-configuredtable-aggregatecolumn.html#cfn-cleanrooms-configuredtable-aggregatecolumn-columnnames
	ColumnNames []T `json:"ColumnNames"`

	// Function AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-configuredtable-aggregatecolumn.html#cfn-cleanrooms-configuredtable-aggregatecolumn-function
	Function T `json:"Function"`

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
func (r *ConfiguredTable_AggregateColumn[any]) AWSCloudFormationType() string {
	return "AWS::CleanRooms::ConfiguredTable.AggregateColumn"
}
