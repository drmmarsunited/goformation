// Code generated by "go generate". Please don't change this file directly.

package resiliencehub

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// App_EventSubscription AWS CloudFormation Resource (AWS::ResilienceHub::App.EventSubscription)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resiliencehub-app-eventsubscription.html
type App_EventSubscription[T any] struct {

	// EventType AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resiliencehub-app-eventsubscription.html#cfn-resiliencehub-app-eventsubscription-eventtype
	EventType T `json:"EventType"`

	// Name AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resiliencehub-app-eventsubscription.html#cfn-resiliencehub-app-eventsubscription-name
	Name T `json:"Name"`

	// SnsTopicArn AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resiliencehub-app-eventsubscription.html#cfn-resiliencehub-app-eventsubscription-snstopicarn
	SnsTopicArn *T `json:"SnsTopicArn,omitempty"`

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
func (r *App_EventSubscription[any]) AWSCloudFormationType() string {
	return "AWS::ResilienceHub::App.EventSubscription"
}
