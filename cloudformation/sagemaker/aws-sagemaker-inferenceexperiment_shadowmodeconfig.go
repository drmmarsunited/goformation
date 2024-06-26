// Code generated by "go generate". Please don't change this file directly.

package sagemaker

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// InferenceExperiment_ShadowModeConfig AWS CloudFormation Resource (AWS::SageMaker::InferenceExperiment.ShadowModeConfig)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-inferenceexperiment-shadowmodeconfig.html
type InferenceExperiment_ShadowModeConfig[T any] struct {

	// ShadowModelVariants AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-inferenceexperiment-shadowmodeconfig.html#cfn-sagemaker-inferenceexperiment-shadowmodeconfig-shadowmodelvariants
	ShadowModelVariants []InferenceExperiment_ShadowModelVariantConfig[any] `json:"ShadowModelVariants"`

	// SourceModelVariantName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-inferenceexperiment-shadowmodeconfig.html#cfn-sagemaker-inferenceexperiment-shadowmodeconfig-sourcemodelvariantname
	SourceModelVariantName T `json:"SourceModelVariantName"`

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
func (r *InferenceExperiment_ShadowModeConfig[any]) AWSCloudFormationType() string {
	return "AWS::SageMaker::InferenceExperiment.ShadowModeConfig"
}
