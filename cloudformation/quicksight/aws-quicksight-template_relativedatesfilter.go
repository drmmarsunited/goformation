// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// Template_RelativeDatesFilter AWS CloudFormation Resource (AWS::QuickSight::Template.RelativeDatesFilter)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-relativedatesfilter.html
type Template_RelativeDatesFilter[T any] struct {

	// AnchorDateConfiguration AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-relativedatesfilter.html#cfn-quicksight-template-relativedatesfilter-anchordateconfiguration
	AnchorDateConfiguration *Template_AnchorDateConfiguration[any] `json:"AnchorDateConfiguration"`

	// Column AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-relativedatesfilter.html#cfn-quicksight-template-relativedatesfilter-column
	Column *Template_ColumnIdentifier[any] `json:"Column"`

	// ExcludePeriodConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-relativedatesfilter.html#cfn-quicksight-template-relativedatesfilter-excludeperiodconfiguration
	ExcludePeriodConfiguration *Template_ExcludePeriodConfiguration[any] `json:"ExcludePeriodConfiguration,omitempty"`

	// FilterId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-relativedatesfilter.html#cfn-quicksight-template-relativedatesfilter-filterid
	FilterId T `json:"FilterId"`

	// MinimumGranularity AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-relativedatesfilter.html#cfn-quicksight-template-relativedatesfilter-minimumgranularity
	MinimumGranularity *T `json:"MinimumGranularity,omitempty"`

	// NullOption AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-relativedatesfilter.html#cfn-quicksight-template-relativedatesfilter-nulloption
	NullOption T `json:"NullOption"`

	// ParameterName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-relativedatesfilter.html#cfn-quicksight-template-relativedatesfilter-parametername
	ParameterName *T `json:"ParameterName,omitempty"`

	// RelativeDateType AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-relativedatesfilter.html#cfn-quicksight-template-relativedatesfilter-relativedatetype
	RelativeDateType T `json:"RelativeDateType"`

	// RelativeDateValue AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-relativedatesfilter.html#cfn-quicksight-template-relativedatesfilter-relativedatevalue
	RelativeDateValue *T `json:"RelativeDateValue,omitempty"`

	// TimeGranularity AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-relativedatesfilter.html#cfn-quicksight-template-relativedatesfilter-timegranularity
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
func (r *Template_RelativeDatesFilter[any]) AWSCloudFormationType() string {
	return "AWS::QuickSight::Template.RelativeDatesFilter"
}
