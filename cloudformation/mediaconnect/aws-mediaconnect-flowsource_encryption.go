// Code generated by "go generate". Please don't change this file directly.

package mediaconnect

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// FlowSource_Encryption AWS CloudFormation Resource (AWS::MediaConnect::FlowSource.Encryption)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flowsource-encryption.html
type FlowSource_Encryption[T any] struct {

	// Algorithm AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flowsource-encryption.html#cfn-mediaconnect-flowsource-encryption-algorithm
	Algorithm *T `json:"Algorithm,omitempty"`

	// ConstantInitializationVector AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flowsource-encryption.html#cfn-mediaconnect-flowsource-encryption-constantinitializationvector
	ConstantInitializationVector *T `json:"ConstantInitializationVector,omitempty"`

	// DeviceId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flowsource-encryption.html#cfn-mediaconnect-flowsource-encryption-deviceid
	DeviceId *T `json:"DeviceId,omitempty"`

	// KeyType AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flowsource-encryption.html#cfn-mediaconnect-flowsource-encryption-keytype
	KeyType *T `json:"KeyType,omitempty"`

	// Region AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flowsource-encryption.html#cfn-mediaconnect-flowsource-encryption-region
	Region *T `json:"Region,omitempty"`

	// ResourceId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flowsource-encryption.html#cfn-mediaconnect-flowsource-encryption-resourceid
	ResourceId *T `json:"ResourceId,omitempty"`

	// RoleArn AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flowsource-encryption.html#cfn-mediaconnect-flowsource-encryption-rolearn
	RoleArn T `json:"RoleArn"`

	// SecretArn AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flowsource-encryption.html#cfn-mediaconnect-flowsource-encryption-secretarn
	SecretArn *T `json:"SecretArn,omitempty"`

	// Url AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flowsource-encryption.html#cfn-mediaconnect-flowsource-encryption-url
	Url *T `json:"Url,omitempty"`

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
func (r *FlowSource_Encryption[any]) AWSCloudFormationType() string {
	return "AWS::MediaConnect::FlowSource.Encryption"
}
