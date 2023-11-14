// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// Analysis_ReferenceLineDataConfiguration AWS CloudFormation Resource (AWS::QuickSight::Analysis.ReferenceLineDataConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-referencelinedataconfiguration.html
type Analysis_ReferenceLineDataConfiguration[T any] struct {

	// AxisBinding AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-referencelinedataconfiguration.html#cfn-quicksight-analysis-referencelinedataconfiguration-axisbinding
	AxisBinding *T `json:"AxisBinding,omitempty"`

	// DynamicConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-referencelinedataconfiguration.html#cfn-quicksight-analysis-referencelinedataconfiguration-dynamicconfiguration
	DynamicConfiguration *Analysis_ReferenceLineDynamicDataConfiguration[any] `json:"DynamicConfiguration,omitempty"`

	// SeriesType AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-referencelinedataconfiguration.html#cfn-quicksight-analysis-referencelinedataconfiguration-seriestype
	SeriesType *T `json:"SeriesType,omitempty"`

	// StaticConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-referencelinedataconfiguration.html#cfn-quicksight-analysis-referencelinedataconfiguration-staticconfiguration
	StaticConfiguration *Analysis_ReferenceLineStaticDataConfiguration[any] `json:"StaticConfiguration,omitempty"`

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
func (r *Analysis_ReferenceLineDataConfiguration[any]) AWSCloudFormationType() string {
	return "AWS::QuickSight::Analysis.ReferenceLineDataConfiguration"
}
