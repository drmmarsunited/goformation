// Code generated by "go generate". Please don't change this file directly.

package fsx

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// StorageVirtualMachine_ActiveDirectoryConfiguration AWS CloudFormation Resource (AWS::FSx::StorageVirtualMachine.ActiveDirectoryConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-storagevirtualmachine-activedirectoryconfiguration.html
type StorageVirtualMachine_ActiveDirectoryConfiguration[T any] struct {

	// NetBiosName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-storagevirtualmachine-activedirectoryconfiguration.html#cfn-fsx-storagevirtualmachine-activedirectoryconfiguration-netbiosname
	NetBiosName *T `json:"NetBiosName,omitempty"`

	// SelfManagedActiveDirectoryConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-storagevirtualmachine-activedirectoryconfiguration.html#cfn-fsx-storagevirtualmachine-activedirectoryconfiguration-selfmanagedactivedirectoryconfiguration
	SelfManagedActiveDirectoryConfiguration *StorageVirtualMachine_SelfManagedActiveDirectoryConfiguration[any] `json:"SelfManagedActiveDirectoryConfiguration,omitempty"`

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
func (r *StorageVirtualMachine_ActiveDirectoryConfiguration[any]) AWSCloudFormationType() string {
	return "AWS::FSx::StorageVirtualMachine.ActiveDirectoryConfiguration"
}
