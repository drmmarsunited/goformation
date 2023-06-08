// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// Analysis_PieChartAggregatedFieldWells AWS CloudFormation Resource (AWS::QuickSight::Analysis.PieChartAggregatedFieldWells)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-piechartaggregatedfieldwells.html
type Analysis_PieChartAggregatedFieldWells[T any] struct {

	// Category AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-piechartaggregatedfieldwells.html#cfn-quicksight-analysis-piechartaggregatedfieldwells-category
	Category []Analysis_DimensionField[any] `json:"Category,omitempty"`

	// SmallMultiples AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-piechartaggregatedfieldwells.html#cfn-quicksight-analysis-piechartaggregatedfieldwells-smallmultiples
	SmallMultiples []Analysis_DimensionField[any] `json:"SmallMultiples,omitempty"`

	// Values AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-piechartaggregatedfieldwells.html#cfn-quicksight-analysis-piechartaggregatedfieldwells-values
	Values []Analysis_MeasureField[any] `json:"Values,omitempty"`

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
func (r *Analysis_PieChartAggregatedFieldWells[any]) AWSCloudFormationType() string {
	return "AWS::QuickSight::Analysis.PieChartAggregatedFieldWells"
}
