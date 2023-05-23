// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// Analysis_DecimalParameterDeclaration AWS CloudFormation Resource (AWS::QuickSight::Analysis.DecimalParameterDeclaration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-decimalparameterdeclaration.html
type Analysis_DecimalParameterDeclaration struct {

	// DefaultValues AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-decimalparameterdeclaration.html#cfn-quicksight-analysis-decimalparameterdeclaration-defaultvalues
	DefaultValues *Analysis_DecimalDefaultValues `json:"DefaultValues,omitempty"`

	// MappedDataSetParameters AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-decimalparameterdeclaration.html#cfn-quicksight-analysis-decimalparameterdeclaration-mappeddatasetparameters
	MappedDataSetParameters []Analysis_MappedDataSetParameter `json:"MappedDataSetParameters,omitempty"`

	// Name AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-decimalparameterdeclaration.html#cfn-quicksight-analysis-decimalparameterdeclaration-name
	Name string `json:"Name"`

	// ParameterValueType AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-decimalparameterdeclaration.html#cfn-quicksight-analysis-decimalparameterdeclaration-parametervaluetype
	ParameterValueType string `json:"ParameterValueType"`

	// ValueWhenUnset AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-decimalparameterdeclaration.html#cfn-quicksight-analysis-decimalparameterdeclaration-valuewhenunset
	ValueWhenUnset *Analysis_DecimalValueWhenUnsetConfiguration `json:"ValueWhenUnset,omitempty"`

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
func (r *Analysis_DecimalParameterDeclaration) AWSCloudFormationType() string {
	return "AWS::QuickSight::Analysis.DecimalParameterDeclaration"
}
