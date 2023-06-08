// Code generated by "go generate". Please don't change this file directly.

package lambda

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// EventSourceMapping_Endpoints AWS CloudFormation Resource (AWS::Lambda::EventSourceMapping.Endpoints)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-eventsourcemapping-endpoints.html
type EventSourceMapping_Endpoints[T any] struct {

	// KafkaBootstrapServers AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-eventsourcemapping-endpoints.html#cfn-lambda-eventsourcemapping-endpoints-kafkabootstrapservers
	KafkaBootstrapServers []string `json:"KafkaBootstrapServers,omitempty"`

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
func (r *EventSourceMapping_Endpoints[any]) AWSCloudFormationType() string {
	return "AWS::Lambda::EventSourceMapping.Endpoints"
}
