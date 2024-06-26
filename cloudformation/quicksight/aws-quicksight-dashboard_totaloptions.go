// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// Dashboard_TotalOptions AWS CloudFormation Resource (AWS::QuickSight::Dashboard.TotalOptions)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-totaloptions.html
type Dashboard_TotalOptions[T any] struct {

	// CustomLabel AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-totaloptions.html#cfn-quicksight-dashboard-totaloptions-customlabel
	CustomLabel *T `json:"CustomLabel,omitempty"`

	// Placement AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-totaloptions.html#cfn-quicksight-dashboard-totaloptions-placement
	Placement *T `json:"Placement,omitempty"`

	// ScrollStatus AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-totaloptions.html#cfn-quicksight-dashboard-totaloptions-scrollstatus
	ScrollStatus *T `json:"ScrollStatus,omitempty"`

	// TotalAggregationOptions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-totaloptions.html#cfn-quicksight-dashboard-totaloptions-totalaggregationoptions
	TotalAggregationOptions []Dashboard_TotalAggregationOption[any] `json:"TotalAggregationOptions,omitempty"`

	// TotalCellStyle AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-totaloptions.html#cfn-quicksight-dashboard-totaloptions-totalcellstyle
	TotalCellStyle *Dashboard_TableCellStyle[any] `json:"TotalCellStyle,omitempty"`

	// TotalsVisibility AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-totaloptions.html#cfn-quicksight-dashboard-totaloptions-totalsvisibility
	TotalsVisibility *T `json:"TotalsVisibility,omitempty"`

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
func (r *Dashboard_TotalOptions[any]) AWSCloudFormationType() string {
	return "AWS::QuickSight::Dashboard.TotalOptions"
}
