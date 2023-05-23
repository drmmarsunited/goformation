// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// Dashboard_MeasureField AWS CloudFormation Resource (AWS::QuickSight::Dashboard.MeasureField)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-measurefield.html
type Dashboard_MeasureField struct {

	// CalculatedMeasureField AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-measurefield.html#cfn-quicksight-dashboard-measurefield-calculatedmeasurefield
	CalculatedMeasureField *Dashboard_CalculatedMeasureField `json:"CalculatedMeasureField,omitempty"`

	// CategoricalMeasureField AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-measurefield.html#cfn-quicksight-dashboard-measurefield-categoricalmeasurefield
	CategoricalMeasureField *Dashboard_CategoricalMeasureField `json:"CategoricalMeasureField,omitempty"`

	// DateMeasureField AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-measurefield.html#cfn-quicksight-dashboard-measurefield-datemeasurefield
	DateMeasureField *Dashboard_DateMeasureField `json:"DateMeasureField,omitempty"`

	// NumericalMeasureField AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-measurefield.html#cfn-quicksight-dashboard-measurefield-numericalmeasurefield
	NumericalMeasureField *Dashboard_NumericalMeasureField `json:"NumericalMeasureField,omitempty"`

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
func (r *Dashboard_MeasureField) AWSCloudFormationType() string {
	return "AWS::QuickSight::Dashboard.MeasureField"
}
