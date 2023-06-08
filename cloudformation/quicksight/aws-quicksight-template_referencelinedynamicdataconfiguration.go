// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// Template_ReferenceLineDynamicDataConfiguration AWS CloudFormation Resource (AWS::QuickSight::Template.ReferenceLineDynamicDataConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-referencelinedynamicdataconfiguration.html
type Template_ReferenceLineDynamicDataConfiguration[T any] struct {

	// Calculation AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-referencelinedynamicdataconfiguration.html#cfn-quicksight-template-referencelinedynamicdataconfiguration-calculation
	Calculation *Template_NumericalAggregationFunction[any] `json:"Calculation"`

	// Column AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-referencelinedynamicdataconfiguration.html#cfn-quicksight-template-referencelinedynamicdataconfiguration-column
	Column *Template_ColumnIdentifier[any] `json:"Column"`

	// MeasureAggregationFunction AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-referencelinedynamicdataconfiguration.html#cfn-quicksight-template-referencelinedynamicdataconfiguration-measureaggregationfunction
	MeasureAggregationFunction *Template_AggregationFunction[any] `json:"MeasureAggregationFunction"`

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
func (r *Template_ReferenceLineDynamicDataConfiguration[any]) AWSCloudFormationType() string {
	return "AWS::QuickSight::Template.ReferenceLineDynamicDataConfiguration"
}
