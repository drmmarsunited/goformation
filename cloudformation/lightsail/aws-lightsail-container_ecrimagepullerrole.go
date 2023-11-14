// Code generated by "go generate". Please don't change this file directly.

package lightsail

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// Container_EcrImagePullerRole AWS CloudFormation Resource (AWS::Lightsail::Container.EcrImagePullerRole)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lightsail-container-ecrimagepullerrole.html
type Container_EcrImagePullerRole[T any] struct {

	// IsActive AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lightsail-container-ecrimagepullerrole.html#cfn-lightsail-container-ecrimagepullerrole-isactive
	IsActive *T `json:"IsActive,omitempty"`

	// PrincipalArn AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lightsail-container-ecrimagepullerrole.html#cfn-lightsail-container-ecrimagepullerrole-principalarn
	PrincipalArn *T `json:"PrincipalArn,omitempty"`

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
func (r *Container_EcrImagePullerRole[any]) AWSCloudFormationType() string {
	return "AWS::Lightsail::Container.EcrImagePullerRole"
}
