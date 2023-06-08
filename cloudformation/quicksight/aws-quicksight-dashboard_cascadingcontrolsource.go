// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// Dashboard_CascadingControlSource AWS CloudFormation Resource (AWS::QuickSight::Dashboard.CascadingControlSource)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-cascadingcontrolsource.html
type Dashboard_CascadingControlSource[T any] struct {

	// ColumnToMatch AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-cascadingcontrolsource.html#cfn-quicksight-dashboard-cascadingcontrolsource-columntomatch
	ColumnToMatch *Dashboard_ColumnIdentifier[any] `json:"ColumnToMatch,omitempty"`

	// SourceSheetControlId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-cascadingcontrolsource.html#cfn-quicksight-dashboard-cascadingcontrolsource-sourcesheetcontrolid
	SourceSheetControlId *string `json:"SourceSheetControlId,omitempty"`

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
func (r *Dashboard_CascadingControlSource[any]) AWSCloudFormationType() string {
	return "AWS::QuickSight::Dashboard.CascadingControlSource"
}
