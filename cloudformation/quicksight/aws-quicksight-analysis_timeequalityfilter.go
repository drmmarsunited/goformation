// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// Analysis_TimeEqualityFilter AWS CloudFormation Resource (AWS::QuickSight::Analysis.TimeEqualityFilter)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-timeequalityfilter.html
type Analysis_TimeEqualityFilter struct {

	// Column AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-timeequalityfilter.html#cfn-quicksight-analysis-timeequalityfilter-column
	Column *Analysis_ColumnIdentifier `json:"Column"`

	// FilterId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-timeequalityfilter.html#cfn-quicksight-analysis-timeequalityfilter-filterid
	FilterId string `json:"FilterId"`

	// ParameterName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-timeequalityfilter.html#cfn-quicksight-analysis-timeequalityfilter-parametername
	ParameterName *string `json:"ParameterName,omitempty"`

	// TimeGranularity AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-timeequalityfilter.html#cfn-quicksight-analysis-timeequalityfilter-timegranularity
	TimeGranularity *string `json:"TimeGranularity,omitempty"`

	// Value AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-timeequalityfilter.html#cfn-quicksight-analysis-timeequalityfilter-value
	Value *string `json:"Value,omitempty"`

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
func (r *Analysis_TimeEqualityFilter) AWSCloudFormationType() string {
	return "AWS::QuickSight::Analysis.TimeEqualityFilter"
}
