// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// DataSet_DateTimeDatasetParameter AWS CloudFormation Resource (AWS::QuickSight::DataSet.DateTimeDatasetParameter)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-datetimedatasetparameter.html
type DataSet_DateTimeDatasetParameter[T any] struct {

	// DefaultValues AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-datetimedatasetparameter.html#cfn-quicksight-dataset-datetimedatasetparameter-defaultvalues
	DefaultValues *DataSet_DateTimeDatasetParameterDefaultValues[any] `json:"DefaultValues,omitempty"`

	// Id AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-datetimedatasetparameter.html#cfn-quicksight-dataset-datetimedatasetparameter-id
	Id T `json:"Id"`

	// Name AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-datetimedatasetparameter.html#cfn-quicksight-dataset-datetimedatasetparameter-name
	Name T `json:"Name"`

	// TimeGranularity AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-datetimedatasetparameter.html#cfn-quicksight-dataset-datetimedatasetparameter-timegranularity
	TimeGranularity *T `json:"TimeGranularity,omitempty"`

	// ValueType AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-datetimedatasetparameter.html#cfn-quicksight-dataset-datetimedatasetparameter-valuetype
	ValueType T `json:"ValueType"`

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
func (r *DataSet_DateTimeDatasetParameter[any]) AWSCloudFormationType() string {
	return "AWS::QuickSight::DataSet.DateTimeDatasetParameter"
}
