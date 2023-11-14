// Code generated by "go generate". Please don't change this file directly.

package devopsguru

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// NotificationChannel_SnsChannelConfig AWS CloudFormation Resource (AWS::DevOpsGuru::NotificationChannel.SnsChannelConfig)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsguru-notificationchannel-snschannelconfig.html
type NotificationChannel_SnsChannelConfig[T any] struct {

	// TopicArn AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsguru-notificationchannel-snschannelconfig.html#cfn-devopsguru-notificationchannel-snschannelconfig-topicarn
	TopicArn *T `json:"TopicArn,omitempty"`

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
func (r *NotificationChannel_SnsChannelConfig[any]) AWSCloudFormationType() string {
	return "AWS::DevOpsGuru::NotificationChannel.SnsChannelConfig"
}
