// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// Dashboard_PieChartAggregatedFieldWells AWS CloudFormation Resource (AWS::QuickSight::Dashboard.PieChartAggregatedFieldWells)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-piechartaggregatedfieldwells.html
type Dashboard_PieChartAggregatedFieldWells struct {

	// Category AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-piechartaggregatedfieldwells.html#cfn-quicksight-dashboard-piechartaggregatedfieldwells-category
	Category []Dashboard_DimensionField `json:"Category,omitempty"`

	// SmallMultiples AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-piechartaggregatedfieldwells.html#cfn-quicksight-dashboard-piechartaggregatedfieldwells-smallmultiples
	SmallMultiples []Dashboard_DimensionField `json:"SmallMultiples,omitempty"`

	// Values AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-piechartaggregatedfieldwells.html#cfn-quicksight-dashboard-piechartaggregatedfieldwells-values
	Values []Dashboard_MeasureField `json:"Values,omitempty"`

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
func (r *Dashboard_PieChartAggregatedFieldWells) AWSCloudFormationType() string {
	return "AWS::QuickSight::Dashboard.PieChartAggregatedFieldWells"
}
