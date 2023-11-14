// Code generated by "go generate". Please don't change this file directly.

package appsync

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// GraphQLApi_LambdaAuthorizerConfig AWS CloudFormation Resource (AWS::AppSync::GraphQLApi.LambdaAuthorizerConfig)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-graphqlapi-lambdaauthorizerconfig.html
type GraphQLApi_LambdaAuthorizerConfig[T any] struct {

	// AuthorizerResultTtlInSeconds AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-graphqlapi-lambdaauthorizerconfig.html#cfn-appsync-graphqlapi-lambdaauthorizerconfig-authorizerresultttlinseconds
	AuthorizerResultTtlInSeconds *T `json:"AuthorizerResultTtlInSeconds,omitempty"`

	// AuthorizerUri AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-graphqlapi-lambdaauthorizerconfig.html#cfn-appsync-graphqlapi-lambdaauthorizerconfig-authorizeruri
	AuthorizerUri *T `json:"AuthorizerUri,omitempty"`

	// IdentityValidationExpression AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-graphqlapi-lambdaauthorizerconfig.html#cfn-appsync-graphqlapi-lambdaauthorizerconfig-identityvalidationexpression
	IdentityValidationExpression *T `json:"IdentityValidationExpression,omitempty"`

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
func (r *GraphQLApi_LambdaAuthorizerConfig[any]) AWSCloudFormationType() string {
	return "AWS::AppSync::GraphQLApi.LambdaAuthorizerConfig"
}
