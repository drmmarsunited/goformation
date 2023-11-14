// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// Analysis_RelativeDatesFilter AWS CloudFormation Resource (AWS::QuickSight::Analysis.RelativeDatesFilter)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-relativedatesfilter.html
type Analysis_RelativeDatesFilter[T any] struct {

	// AnchorDateConfiguration AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-relativedatesfilter.html#cfn-quicksight-analysis-relativedatesfilter-anchordateconfiguration
	AnchorDateConfiguration *Analysis_AnchorDateConfiguration[any] `json:"AnchorDateConfiguration"`

	// Column AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-relativedatesfilter.html#cfn-quicksight-analysis-relativedatesfilter-column
	Column *Analysis_ColumnIdentifier[any] `json:"Column"`

	// ExcludePeriodConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-relativedatesfilter.html#cfn-quicksight-analysis-relativedatesfilter-excludeperiodconfiguration
	ExcludePeriodConfiguration *Analysis_ExcludePeriodConfiguration[any] `json:"ExcludePeriodConfiguration,omitempty"`

	// FilterId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-relativedatesfilter.html#cfn-quicksight-analysis-relativedatesfilter-filterid
	FilterId T `json:"FilterId"`

	// MinimumGranularity AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-relativedatesfilter.html#cfn-quicksight-analysis-relativedatesfilter-minimumgranularity
	MinimumGranularity *T `json:"MinimumGranularity,omitempty"`

	// NullOption AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-relativedatesfilter.html#cfn-quicksight-analysis-relativedatesfilter-nulloption
	NullOption T `json:"NullOption"`

	// ParameterName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-relativedatesfilter.html#cfn-quicksight-analysis-relativedatesfilter-parametername
	ParameterName *T `json:"ParameterName,omitempty"`

	// RelativeDateType AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-relativedatesfilter.html#cfn-quicksight-analysis-relativedatesfilter-relativedatetype
	RelativeDateType T `json:"RelativeDateType"`

	// RelativeDateValue AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-relativedatesfilter.html#cfn-quicksight-analysis-relativedatesfilter-relativedatevalue
	RelativeDateValue *T `json:"RelativeDateValue,omitempty"`

	// TimeGranularity AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-relativedatesfilter.html#cfn-quicksight-analysis-relativedatesfilter-timegranularity
	TimeGranularity T `json:"TimeGranularity"`

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
func (r *Analysis_RelativeDatesFilter[any]) AWSCloudFormationType() string {
	return "AWS::QuickSight::Analysis.RelativeDatesFilter"
}
