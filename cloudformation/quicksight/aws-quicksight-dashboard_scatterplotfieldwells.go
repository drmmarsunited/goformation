// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// Dashboard_ScatterPlotFieldWells AWS CloudFormation Resource (AWS::QuickSight::Dashboard.ScatterPlotFieldWells)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-scatterplotfieldwells.html
type Dashboard_ScatterPlotFieldWells struct {

	// ScatterPlotCategoricallyAggregatedFieldWells AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-scatterplotfieldwells.html#cfn-quicksight-dashboard-scatterplotfieldwells-scatterplotcategoricallyaggregatedfieldwells
	ScatterPlotCategoricallyAggregatedFieldWells *Dashboard_ScatterPlotCategoricallyAggregatedFieldWells `json:"ScatterPlotCategoricallyAggregatedFieldWells,omitempty"`

	// ScatterPlotUnaggregatedFieldWells AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-scatterplotfieldwells.html#cfn-quicksight-dashboard-scatterplotfieldwells-scatterplotunaggregatedfieldwells
	ScatterPlotUnaggregatedFieldWells *Dashboard_ScatterPlotUnaggregatedFieldWells `json:"ScatterPlotUnaggregatedFieldWells,omitempty"`

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
func (r *Dashboard_ScatterPlotFieldWells) AWSCloudFormationType() string {
	return "AWS::QuickSight::Dashboard.ScatterPlotFieldWells"
}
