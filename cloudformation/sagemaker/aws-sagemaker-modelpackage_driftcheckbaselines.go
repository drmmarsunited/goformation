// Code generated by "go generate". Please don't change this file directly.

package sagemaker

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// ModelPackage_DriftCheckBaselines AWS CloudFormation Resource (AWS::SageMaker::ModelPackage.DriftCheckBaselines)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelpackage-driftcheckbaselines.html
type ModelPackage_DriftCheckBaselines[T any] struct {

	// Bias AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelpackage-driftcheckbaselines.html#cfn-sagemaker-modelpackage-driftcheckbaselines-bias
	Bias *ModelPackage_DriftCheckBias[any] `json:"Bias,omitempty"`

	// Explainability AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelpackage-driftcheckbaselines.html#cfn-sagemaker-modelpackage-driftcheckbaselines-explainability
	Explainability *ModelPackage_DriftCheckExplainability[any] `json:"Explainability,omitempty"`

	// ModelDataQuality AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelpackage-driftcheckbaselines.html#cfn-sagemaker-modelpackage-driftcheckbaselines-modeldataquality
	ModelDataQuality *ModelPackage_DriftCheckModelDataQuality[any] `json:"ModelDataQuality,omitempty"`

	// ModelQuality AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelpackage-driftcheckbaselines.html#cfn-sagemaker-modelpackage-driftcheckbaselines-modelquality
	ModelQuality *ModelPackage_DriftCheckModelQuality[any] `json:"ModelQuality,omitempty"`

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
func (r *ModelPackage_DriftCheckBaselines[any]) AWSCloudFormationType() string {
	return "AWS::SageMaker::ModelPackage.DriftCheckBaselines"
}
