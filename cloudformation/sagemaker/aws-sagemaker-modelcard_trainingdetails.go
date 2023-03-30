// Code generated by "go generate". Please don't change this file directly.

package sagemaker

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// ModelCard_TrainingDetails AWS CloudFormation Resource (AWS::SageMaker::ModelCard.TrainingDetails)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelcard-trainingdetails.html
type ModelCard_TrainingDetails struct {

	// ObjectiveFunction AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelcard-trainingdetails.html#cfn-sagemaker-modelcard-trainingdetails-objectivefunction
	ObjectiveFunction *ModelCard_ObjectiveFunction `json:"ObjectiveFunction,omitempty"`

	// TrainingJobDetails AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelcard-trainingdetails.html#cfn-sagemaker-modelcard-trainingdetails-trainingjobdetails
	TrainingJobDetails *ModelCard_TrainingJobDetails `json:"TrainingJobDetails,omitempty"`

	// TrainingObservations AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelcard-trainingdetails.html#cfn-sagemaker-modelcard-trainingdetails-trainingobservations
	TrainingObservations *string `json:"TrainingObservations,omitempty"`

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
func (r *ModelCard_TrainingDetails) AWSCloudFormationType() string {
	return "AWS::SageMaker::ModelCard.TrainingDetails"
}