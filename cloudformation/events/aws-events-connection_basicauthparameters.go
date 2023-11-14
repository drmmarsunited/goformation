// Code generated by "go generate". Please don't change this file directly.

package events

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// Connection_BasicAuthParameters AWS CloudFormation Resource (AWS::Events::Connection.BasicAuthParameters)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-connection-basicauthparameters.html
type Connection_BasicAuthParameters[T any] struct {

	// Password AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-connection-basicauthparameters.html#cfn-events-connection-basicauthparameters-password
	Password T `json:"Password"`

	// Username AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-connection-basicauthparameters.html#cfn-events-connection-basicauthparameters-username
	Username T `json:"Username"`

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
func (r *Connection_BasicAuthParameters[any]) AWSCloudFormationType() string {
	return "AWS::Events::Connection.BasicAuthParameters"
}
