// Code generated by "go generate". Please don't change this file directly.

package imagebuilder

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// ContainerRecipe_EbsInstanceBlockDeviceSpecification AWS CloudFormation Resource (AWS::ImageBuilder::ContainerRecipe.EbsInstanceBlockDeviceSpecification)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-containerrecipe-ebsinstanceblockdevicespecification.html
type ContainerRecipe_EbsInstanceBlockDeviceSpecification[T any] struct {

	// DeleteOnTermination AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-containerrecipe-ebsinstanceblockdevicespecification.html#cfn-imagebuilder-containerrecipe-ebsinstanceblockdevicespecification-deleteontermination
	DeleteOnTermination *T `json:"DeleteOnTermination,omitempty"`

	// Encrypted AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-containerrecipe-ebsinstanceblockdevicespecification.html#cfn-imagebuilder-containerrecipe-ebsinstanceblockdevicespecification-encrypted
	Encrypted *T `json:"Encrypted,omitempty"`

	// Iops AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-containerrecipe-ebsinstanceblockdevicespecification.html#cfn-imagebuilder-containerrecipe-ebsinstanceblockdevicespecification-iops
	Iops *T `json:"Iops,omitempty"`

	// KmsKeyId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-containerrecipe-ebsinstanceblockdevicespecification.html#cfn-imagebuilder-containerrecipe-ebsinstanceblockdevicespecification-kmskeyid
	KmsKeyId *T `json:"KmsKeyId,omitempty"`

	// SnapshotId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-containerrecipe-ebsinstanceblockdevicespecification.html#cfn-imagebuilder-containerrecipe-ebsinstanceblockdevicespecification-snapshotid
	SnapshotId *T `json:"SnapshotId,omitempty"`

	// Throughput AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-containerrecipe-ebsinstanceblockdevicespecification.html#cfn-imagebuilder-containerrecipe-ebsinstanceblockdevicespecification-throughput
	Throughput *T `json:"Throughput,omitempty"`

	// VolumeSize AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-containerrecipe-ebsinstanceblockdevicespecification.html#cfn-imagebuilder-containerrecipe-ebsinstanceblockdevicespecification-volumesize
	VolumeSize *T `json:"VolumeSize,omitempty"`

	// VolumeType AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-containerrecipe-ebsinstanceblockdevicespecification.html#cfn-imagebuilder-containerrecipe-ebsinstanceblockdevicespecification-volumetype
	VolumeType *T `json:"VolumeType,omitempty"`

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
func (r *ContainerRecipe_EbsInstanceBlockDeviceSpecification[any]) AWSCloudFormationType() string {
	return "AWS::ImageBuilder::ContainerRecipe.EbsInstanceBlockDeviceSpecification"
}
