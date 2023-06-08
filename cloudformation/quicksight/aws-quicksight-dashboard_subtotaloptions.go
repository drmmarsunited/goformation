// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// Dashboard_SubtotalOptions AWS CloudFormation Resource (AWS::QuickSight::Dashboard.SubtotalOptions)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-subtotaloptions.html
type Dashboard_SubtotalOptions[T any] struct {

	// CustomLabel AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-subtotaloptions.html#cfn-quicksight-dashboard-subtotaloptions-customlabel
	CustomLabel *string `json:"CustomLabel,omitempty"`

	// FieldLevel AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-subtotaloptions.html#cfn-quicksight-dashboard-subtotaloptions-fieldlevel
	FieldLevel *string `json:"FieldLevel,omitempty"`

	// FieldLevelOptions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-subtotaloptions.html#cfn-quicksight-dashboard-subtotaloptions-fieldleveloptions
	FieldLevelOptions []Dashboard_PivotTableFieldSubtotalOptions[any] `json:"FieldLevelOptions,omitempty"`

	// MetricHeaderCellStyle AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-subtotaloptions.html#cfn-quicksight-dashboard-subtotaloptions-metricheadercellstyle
	MetricHeaderCellStyle *Dashboard_TableCellStyle[any] `json:"MetricHeaderCellStyle,omitempty"`

	// TotalCellStyle AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-subtotaloptions.html#cfn-quicksight-dashboard-subtotaloptions-totalcellstyle
	TotalCellStyle *Dashboard_TableCellStyle[any] `json:"TotalCellStyle,omitempty"`

	// TotalsVisibility AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-subtotaloptions.html#cfn-quicksight-dashboard-subtotaloptions-totalsvisibility
	TotalsVisibility *string `json:"TotalsVisibility,omitempty"`

	// ValueCellStyle AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-subtotaloptions.html#cfn-quicksight-dashboard-subtotaloptions-valuecellstyle
	ValueCellStyle *Dashboard_TableCellStyle[any] `json:"ValueCellStyle,omitempty"`

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
func (r *Dashboard_SubtotalOptions[any]) AWSCloudFormationType() string {
	return "AWS::QuickSight::Dashboard.SubtotalOptions"
}
