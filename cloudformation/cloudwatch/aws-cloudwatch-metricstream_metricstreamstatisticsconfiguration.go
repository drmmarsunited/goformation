// Code generated by "go generate". Please don't change this file directly.

package cloudwatch

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// MetricStream_MetricStreamStatisticsConfiguration AWS CloudFormation Resource (AWS::CloudWatch::MetricStream.MetricStreamStatisticsConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-metricstream-metricstreamstatisticsconfiguration.html
type MetricStream_MetricStreamStatisticsConfiguration[T any] struct {

	// AdditionalStatistics AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-metricstream-metricstreamstatisticsconfiguration.html#cfn-cloudwatch-metricstream-metricstreamstatisticsconfiguration-additionalstatistics
	AdditionalStatistics []T `json:"AdditionalStatistics"`

	// IncludeMetrics AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-metricstream-metricstreamstatisticsconfiguration.html#cfn-cloudwatch-metricstream-metricstreamstatisticsconfiguration-includemetrics
	IncludeMetrics []MetricStream_MetricStreamStatisticsMetric[any] `json:"IncludeMetrics"`

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
func (r *MetricStream_MetricStreamStatisticsConfiguration[any]) AWSCloudFormationType() string {
	return "AWS::CloudWatch::MetricStream.MetricStreamStatisticsConfiguration"
}
