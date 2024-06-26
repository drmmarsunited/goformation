// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// Dashboard_DecimalParameterDeclaration AWS CloudFormation Resource (AWS::QuickSight::Dashboard.DecimalParameterDeclaration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-decimalparameterdeclaration.html
type Dashboard_DecimalParameterDeclaration[T any] struct {

	// DefaultValues AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-decimalparameterdeclaration.html#cfn-quicksight-dashboard-decimalparameterdeclaration-defaultvalues
	DefaultValues *Dashboard_DecimalDefaultValues[any] `json:"DefaultValues,omitempty"`

	// MappedDataSetParameters AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-decimalparameterdeclaration.html#cfn-quicksight-dashboard-decimalparameterdeclaration-mappeddatasetparameters
	MappedDataSetParameters []Dashboard_MappedDataSetParameter[any] `json:"MappedDataSetParameters,omitempty"`

	// Name AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-decimalparameterdeclaration.html#cfn-quicksight-dashboard-decimalparameterdeclaration-name
	Name T `json:"Name"`

	// ParameterValueType AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-decimalparameterdeclaration.html#cfn-quicksight-dashboard-decimalparameterdeclaration-parametervaluetype
	ParameterValueType T `json:"ParameterValueType"`

	// ValueWhenUnset AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-decimalparameterdeclaration.html#cfn-quicksight-dashboard-decimalparameterdeclaration-valuewhenunset
	ValueWhenUnset *Dashboard_DecimalValueWhenUnsetConfiguration[any] `json:"ValueWhenUnset,omitempty"`

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
func (r *Dashboard_DecimalParameterDeclaration[any]) AWSCloudFormationType() string {
	return "AWS::QuickSight::Dashboard.DecimalParameterDeclaration"
}
