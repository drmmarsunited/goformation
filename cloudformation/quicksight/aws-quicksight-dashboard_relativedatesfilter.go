// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// Dashboard_RelativeDatesFilter AWS CloudFormation Resource (AWS::QuickSight::Dashboard.RelativeDatesFilter)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-relativedatesfilter.html
type Dashboard_RelativeDatesFilter struct {

	// AnchorDateConfiguration AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-relativedatesfilter.html#cfn-quicksight-dashboard-relativedatesfilter-anchordateconfiguration
	AnchorDateConfiguration *Dashboard_AnchorDateConfiguration `json:"AnchorDateConfiguration"`

	// Column AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-relativedatesfilter.html#cfn-quicksight-dashboard-relativedatesfilter-column
	Column *Dashboard_ColumnIdentifier `json:"Column"`

	// ExcludePeriodConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-relativedatesfilter.html#cfn-quicksight-dashboard-relativedatesfilter-excludeperiodconfiguration
	ExcludePeriodConfiguration *Dashboard_ExcludePeriodConfiguration `json:"ExcludePeriodConfiguration,omitempty"`

	// FilterId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-relativedatesfilter.html#cfn-quicksight-dashboard-relativedatesfilter-filterid
	FilterId string `json:"FilterId"`

	// MinimumGranularity AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-relativedatesfilter.html#cfn-quicksight-dashboard-relativedatesfilter-minimumgranularity
	MinimumGranularity *string `json:"MinimumGranularity,omitempty"`

	// NullOption AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-relativedatesfilter.html#cfn-quicksight-dashboard-relativedatesfilter-nulloption
	NullOption string `json:"NullOption"`

	// ParameterName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-relativedatesfilter.html#cfn-quicksight-dashboard-relativedatesfilter-parametername
	ParameterName *string `json:"ParameterName,omitempty"`

	// RelativeDateType AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-relativedatesfilter.html#cfn-quicksight-dashboard-relativedatesfilter-relativedatetype
	RelativeDateType string `json:"RelativeDateType"`

	// RelativeDateValue AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-relativedatesfilter.html#cfn-quicksight-dashboard-relativedatesfilter-relativedatevalue
	RelativeDateValue *float64 `json:"RelativeDateValue,omitempty"`

	// TimeGranularity AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-relativedatesfilter.html#cfn-quicksight-dashboard-relativedatesfilter-timegranularity
	TimeGranularity string `json:"TimeGranularity"`

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
func (r *Dashboard_RelativeDatesFilter) AWSCloudFormationType() string {
	return "AWS::QuickSight::Dashboard.RelativeDatesFilter"
}
