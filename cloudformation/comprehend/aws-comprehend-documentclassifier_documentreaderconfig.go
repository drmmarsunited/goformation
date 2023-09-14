// Code generated by "go generate". Please don't change this file directly.

package comprehend

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// DocumentClassifier_DocumentReaderConfig AWS CloudFormation Resource (AWS::Comprehend::DocumentClassifier.DocumentReaderConfig)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-comprehend-documentclassifier-documentreaderconfig.html
type DocumentClassifier_DocumentReaderConfig[T any] struct {

	// DocumentReadAction AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-comprehend-documentclassifier-documentreaderconfig.html#cfn-comprehend-documentclassifier-documentreaderconfig-documentreadaction
	DocumentReadAction string `json:"DocumentReadAction"`

	// DocumentReadMode AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-comprehend-documentclassifier-documentreaderconfig.html#cfn-comprehend-documentclassifier-documentreaderconfig-documentreadmode
	DocumentReadMode *string `json:"DocumentReadMode,omitempty"`

	// FeatureTypes AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-comprehend-documentclassifier-documentreaderconfig.html#cfn-comprehend-documentclassifier-documentreaderconfig-featuretypes
	FeatureTypes []string `json:"FeatureTypes,omitempty"`

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
func (r *DocumentClassifier_DocumentReaderConfig[any]) AWSCloudFormationType() string {
	return "AWS::Comprehend::DocumentClassifier.DocumentReaderConfig"
}
