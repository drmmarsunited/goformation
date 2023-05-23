// Code generated by "go generate". Please don't change this file directly.

package datasync

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// LocationFSxOpenZFS_Protocol AWS CloudFormation Resource (AWS::DataSync::LocationFSxOpenZFS.Protocol)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datasync-locationfsxopenzfs-protocol.html
type LocationFSxOpenZFS_Protocol struct {

	// NFS AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datasync-locationfsxopenzfs-protocol.html#cfn-datasync-locationfsxopenzfs-protocol-nfs
	NFS *LocationFSxOpenZFS_NFS `json:"NFS,omitempty"`

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
func (r *LocationFSxOpenZFS_Protocol) AWSCloudFormationType() string {
	return "AWS::DataSync::LocationFSxOpenZFS.Protocol"
}
