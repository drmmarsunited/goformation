// Code generated by "go generate". Please don't change this file directly.

package servicecatalog

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// CloudFormationProduct_CodeStarParameters AWS CloudFormation Resource (AWS::ServiceCatalog::CloudFormationProduct.CodeStarParameters)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-servicecatalog-cloudformationproduct-codestarparameters.html
type CloudFormationProduct_CodeStarParameters[T any] struct {

	// ArtifactPath AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-servicecatalog-cloudformationproduct-codestarparameters.html#cfn-servicecatalog-cloudformationproduct-codestarparameters-artifactpath
	ArtifactPath T `json:"ArtifactPath"`

	// Branch AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-servicecatalog-cloudformationproduct-codestarparameters.html#cfn-servicecatalog-cloudformationproduct-codestarparameters-branch
	Branch T `json:"Branch"`

	// ConnectionArn AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-servicecatalog-cloudformationproduct-codestarparameters.html#cfn-servicecatalog-cloudformationproduct-codestarparameters-connectionarn
	ConnectionArn T `json:"ConnectionArn"`

	// Repository AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-servicecatalog-cloudformationproduct-codestarparameters.html#cfn-servicecatalog-cloudformationproduct-codestarparameters-repository
	Repository T `json:"Repository"`

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
func (r *CloudFormationProduct_CodeStarParameters[any]) AWSCloudFormationType() string {
	return "AWS::ServiceCatalog::CloudFormationProduct.CodeStarParameters"
}
