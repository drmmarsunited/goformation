// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// Template_WordCloudAggregatedFieldWells AWS CloudFormation Resource (AWS::QuickSight::Template.WordCloudAggregatedFieldWells)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-wordcloudaggregatedfieldwells.html
type Template_WordCloudAggregatedFieldWells struct {

	// GroupBy AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-wordcloudaggregatedfieldwells.html#cfn-quicksight-template-wordcloudaggregatedfieldwells-groupby
	GroupBy []Template_DimensionField `json:"GroupBy,omitempty"`

	// Size AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-wordcloudaggregatedfieldwells.html#cfn-quicksight-template-wordcloudaggregatedfieldwells-size
	Size []Template_MeasureField `json:"Size,omitempty"`

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
func (r *Template_WordCloudAggregatedFieldWells) AWSCloudFormationType() string {
	return "AWS::QuickSight::Template.WordCloudAggregatedFieldWells"
}
