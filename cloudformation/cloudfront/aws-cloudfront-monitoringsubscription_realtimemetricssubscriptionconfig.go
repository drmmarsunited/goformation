// Code generated by "go generate". Please don't change this file directly.

package cloudfront

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// MonitoringSubscription_RealtimeMetricsSubscriptionConfig AWS CloudFormation Resource (AWS::CloudFront::MonitoringSubscription.RealtimeMetricsSubscriptionConfig)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-monitoringsubscription-realtimemetricssubscriptionconfig.html
type MonitoringSubscription_RealtimeMetricsSubscriptionConfig struct {

	// RealtimeMetricsSubscriptionStatus AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-monitoringsubscription-realtimemetricssubscriptionconfig.html#cfn-cloudfront-monitoringsubscription-realtimemetricssubscriptionconfig-realtimemetricssubscriptionstatus
	RealtimeMetricsSubscriptionStatus string `json:"RealtimeMetricsSubscriptionStatus"`

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
func (r *MonitoringSubscription_RealtimeMetricsSubscriptionConfig) AWSCloudFormationType() string {
	return "AWS::CloudFront::MonitoringSubscription.RealtimeMetricsSubscriptionConfig"
}