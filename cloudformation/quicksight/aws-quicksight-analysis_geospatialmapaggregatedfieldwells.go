// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// Analysis_GeospatialMapAggregatedFieldWells AWS CloudFormation Resource (AWS::QuickSight::Analysis.GeospatialMapAggregatedFieldWells)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-geospatialmapaggregatedfieldwells.html
type Analysis_GeospatialMapAggregatedFieldWells struct {

	// Colors AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-geospatialmapaggregatedfieldwells.html#cfn-quicksight-analysis-geospatialmapaggregatedfieldwells-colors
	Colors []Analysis_DimensionField `json:"Colors,omitempty"`

	// Geospatial AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-geospatialmapaggregatedfieldwells.html#cfn-quicksight-analysis-geospatialmapaggregatedfieldwells-geospatial
	Geospatial []Analysis_DimensionField `json:"Geospatial,omitempty"`

	// Values AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-geospatialmapaggregatedfieldwells.html#cfn-quicksight-analysis-geospatialmapaggregatedfieldwells-values
	Values []Analysis_MeasureField `json:"Values,omitempty"`

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
func (r *Analysis_GeospatialMapAggregatedFieldWells) AWSCloudFormationType() string {
	return "AWS::QuickSight::Analysis.GeospatialMapAggregatedFieldWells"
}
