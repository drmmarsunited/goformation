// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// Dashboard_GaugeChartFieldWells AWS CloudFormation Resource (AWS::QuickSight::Dashboard.GaugeChartFieldWells)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-gaugechartfieldwells.html
type Dashboard_GaugeChartFieldWells[T any] struct {

	// TargetValues AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-gaugechartfieldwells.html#cfn-quicksight-dashboard-gaugechartfieldwells-targetvalues
	TargetValues []Dashboard_MeasureField[any] `json:"TargetValues,omitempty"`

	// Values AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-gaugechartfieldwells.html#cfn-quicksight-dashboard-gaugechartfieldwells-values
	Values []Dashboard_MeasureField[any] `json:"Values,omitempty"`

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
func (r *Dashboard_GaugeChartFieldWells[any]) AWSCloudFormationType() string {
	return "AWS::QuickSight::Dashboard.GaugeChartFieldWells"
}
