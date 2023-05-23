// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// Dashboard_SankeyDiagramSortConfiguration AWS CloudFormation Resource (AWS::QuickSight::Dashboard.SankeyDiagramSortConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-sankeydiagramsortconfiguration.html
type Dashboard_SankeyDiagramSortConfiguration struct {

	// DestinationItemsLimit AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-sankeydiagramsortconfiguration.html#cfn-quicksight-dashboard-sankeydiagramsortconfiguration-destinationitemslimit
	DestinationItemsLimit *Dashboard_ItemsLimitConfiguration `json:"DestinationItemsLimit,omitempty"`

	// SourceItemsLimit AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-sankeydiagramsortconfiguration.html#cfn-quicksight-dashboard-sankeydiagramsortconfiguration-sourceitemslimit
	SourceItemsLimit *Dashboard_ItemsLimitConfiguration `json:"SourceItemsLimit,omitempty"`

	// WeightSort AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-sankeydiagramsortconfiguration.html#cfn-quicksight-dashboard-sankeydiagramsortconfiguration-weightsort
	WeightSort []Dashboard_FieldSortOptions `json:"WeightSort,omitempty"`

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
func (r *Dashboard_SankeyDiagramSortConfiguration) AWSCloudFormationType() string {
	return "AWS::QuickSight::Dashboard.SankeyDiagramSortConfiguration"
}
