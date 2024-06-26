// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// Dashboard_ForecastConfiguration AWS CloudFormation Resource (AWS::QuickSight::Dashboard.ForecastConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-forecastconfiguration.html
type Dashboard_ForecastConfiguration[T any] struct {

	// ForecastProperties AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-forecastconfiguration.html#cfn-quicksight-dashboard-forecastconfiguration-forecastproperties
	ForecastProperties *Dashboard_TimeBasedForecastProperties[any] `json:"ForecastProperties,omitempty"`

	// Scenario AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-forecastconfiguration.html#cfn-quicksight-dashboard-forecastconfiguration-scenario
	Scenario *Dashboard_ForecastScenario[any] `json:"Scenario,omitempty"`

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
func (r *Dashboard_ForecastConfiguration[any]) AWSCloudFormationType() string {
	return "AWS::QuickSight::Dashboard.ForecastConfiguration"
}
