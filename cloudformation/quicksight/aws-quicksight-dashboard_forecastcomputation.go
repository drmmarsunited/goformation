// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// Dashboard_ForecastComputation AWS CloudFormation Resource (AWS::QuickSight::Dashboard.ForecastComputation)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-forecastcomputation.html
type Dashboard_ForecastComputation[T any] struct {

	// ComputationId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-forecastcomputation.html#cfn-quicksight-dashboard-forecastcomputation-computationid
	ComputationId T `json:"ComputationId"`

	// CustomSeasonalityValue AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-forecastcomputation.html#cfn-quicksight-dashboard-forecastcomputation-customseasonalityvalue
	CustomSeasonalityValue *T `json:"CustomSeasonalityValue,omitempty"`

	// LowerBoundary AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-forecastcomputation.html#cfn-quicksight-dashboard-forecastcomputation-lowerboundary
	LowerBoundary *T `json:"LowerBoundary,omitempty"`

	// Name AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-forecastcomputation.html#cfn-quicksight-dashboard-forecastcomputation-name
	Name *T `json:"Name,omitempty"`

	// PeriodsBackward AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-forecastcomputation.html#cfn-quicksight-dashboard-forecastcomputation-periodsbackward
	PeriodsBackward *T `json:"PeriodsBackward,omitempty"`

	// PeriodsForward AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-forecastcomputation.html#cfn-quicksight-dashboard-forecastcomputation-periodsforward
	PeriodsForward *T `json:"PeriodsForward,omitempty"`

	// PredictionInterval AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-forecastcomputation.html#cfn-quicksight-dashboard-forecastcomputation-predictioninterval
	PredictionInterval *T `json:"PredictionInterval,omitempty"`

	// Seasonality AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-forecastcomputation.html#cfn-quicksight-dashboard-forecastcomputation-seasonality
	Seasonality *T `json:"Seasonality,omitempty"`

	// Time AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-forecastcomputation.html#cfn-quicksight-dashboard-forecastcomputation-time
	Time *Dashboard_DimensionField[any] `json:"Time,omitempty"`

	// UpperBoundary AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-forecastcomputation.html#cfn-quicksight-dashboard-forecastcomputation-upperboundary
	UpperBoundary *T `json:"UpperBoundary,omitempty"`

	// Value AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-forecastcomputation.html#cfn-quicksight-dashboard-forecastcomputation-value
	Value *Dashboard_MeasureField[any] `json:"Value,omitempty"`

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
func (r *Dashboard_ForecastComputation[any]) AWSCloudFormationType() string {
	return "AWS::QuickSight::Dashboard.ForecastComputation"
}
