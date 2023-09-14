// Code generated by "go generate". Please don't change this file directly.

package internetmonitor

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// Monitor_HealthEventsConfig AWS CloudFormation Resource (AWS::InternetMonitor::Monitor.HealthEventsConfig)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-internetmonitor-monitor-healtheventsconfig.html
type Monitor_HealthEventsConfig[T any] struct {

	// AvailabilityLocalHealthEventsConfig AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-internetmonitor-monitor-healtheventsconfig.html#cfn-internetmonitor-monitor-healtheventsconfig-availabilitylocalhealtheventsconfig
	AvailabilityLocalHealthEventsConfig *Monitor_LocalHealthEventsConfig[any] `json:"AvailabilityLocalHealthEventsConfig,omitempty"`

	// AvailabilityScoreThreshold AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-internetmonitor-monitor-healtheventsconfig.html#cfn-internetmonitor-monitor-healtheventsconfig-availabilityscorethreshold
	AvailabilityScoreThreshold *T `json:"AvailabilityScoreThreshold,omitempty"`

	// PerformanceLocalHealthEventsConfig AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-internetmonitor-monitor-healtheventsconfig.html#cfn-internetmonitor-monitor-healtheventsconfig-performancelocalhealtheventsconfig
	PerformanceLocalHealthEventsConfig *Monitor_LocalHealthEventsConfig[any] `json:"PerformanceLocalHealthEventsConfig,omitempty"`

	// PerformanceScoreThreshold AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-internetmonitor-monitor-healtheventsconfig.html#cfn-internetmonitor-monitor-healtheventsconfig-performancescorethreshold
	PerformanceScoreThreshold *T `json:"PerformanceScoreThreshold,omitempty"`

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
func (r *Monitor_HealthEventsConfig[any]) AWSCloudFormationType() string {
	return "AWS::InternetMonitor::Monitor.HealthEventsConfig"
}
