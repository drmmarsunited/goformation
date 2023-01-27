// Code generated by "go generate". Please don't change this file directly.

package lex

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Bot_InitialResponseSetting AWS CloudFormation Resource (AWS::Lex::Bot.InitialResponseSetting)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-initialresponsesetting.html
type Bot_InitialResponseSetting struct {

	// CodeHook AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-initialresponsesetting.html#cfn-lex-bot-initialresponsesetting-codehook
	CodeHook *Bot_DialogCodeHookInvocationSetting `json:"CodeHook,omitempty"`

	// Conditional AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-initialresponsesetting.html#cfn-lex-bot-initialresponsesetting-conditional
	Conditional *Bot_ConditionalSpecification `json:"Conditional,omitempty"`

	// InitialResponse AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-initialresponsesetting.html#cfn-lex-bot-initialresponsesetting-initialresponse
	InitialResponse *Bot_ResponseSpecification `json:"InitialResponse,omitempty"`

	// NextStep AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-initialresponsesetting.html#cfn-lex-bot-initialresponsesetting-nextstep
	NextStep *Bot_DialogState `json:"NextStep,omitempty"`

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
func (r *Bot_InitialResponseSetting) AWSCloudFormationType() string {
	return "AWS::Lex::Bot.InitialResponseSetting"
}
