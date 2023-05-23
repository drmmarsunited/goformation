// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// Dashboard_ColumnHierarchy AWS CloudFormation Resource (AWS::QuickSight::Dashboard.ColumnHierarchy)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-columnhierarchy.html
type Dashboard_ColumnHierarchy struct {

	// DateTimeHierarchy AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-columnhierarchy.html#cfn-quicksight-dashboard-columnhierarchy-datetimehierarchy
	DateTimeHierarchy *Dashboard_DateTimeHierarchy `json:"DateTimeHierarchy,omitempty"`

	// ExplicitHierarchy AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-columnhierarchy.html#cfn-quicksight-dashboard-columnhierarchy-explicithierarchy
	ExplicitHierarchy *Dashboard_ExplicitHierarchy `json:"ExplicitHierarchy,omitempty"`

	// PredefinedHierarchy AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-columnhierarchy.html#cfn-quicksight-dashboard-columnhierarchy-predefinedhierarchy
	PredefinedHierarchy *Dashboard_PredefinedHierarchy `json:"PredefinedHierarchy,omitempty"`

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
func (r *Dashboard_ColumnHierarchy) AWSCloudFormationType() string {
	return "AWS::QuickSight::Dashboard.ColumnHierarchy"
}
