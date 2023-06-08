// Code generated by "go generate". Please don't change this file directly.

package iot

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// DomainConfiguration_TlsConfig AWS CloudFormation Resource (AWS::IoT::DomainConfiguration.TlsConfig)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-domainconfiguration-tlsconfig.html
type DomainConfiguration_TlsConfig[T any] struct {

	// SecurityPolicy AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-domainconfiguration-tlsconfig.html#cfn-iot-domainconfiguration-tlsconfig-securitypolicy
	SecurityPolicy *string `json:"SecurityPolicy,omitempty"`

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
func (r *DomainConfiguration_TlsConfig[any]) AWSCloudFormationType() string {
	return "AWS::IoT::DomainConfiguration.TlsConfig"
}
