// Code generated by "go generate". Please don't change this file directly.

package iotwireless

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// WirelessDeviceImportTask_Sidewalk AWS CloudFormation Resource (AWS::IoTWireless::WirelessDeviceImportTask.Sidewalk)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotwireless-wirelessdeviceimporttask-sidewalk.html
type WirelessDeviceImportTask_Sidewalk[T any] struct {

	// DeviceCreationFile AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotwireless-wirelessdeviceimporttask-sidewalk.html#cfn-iotwireless-wirelessdeviceimporttask-sidewalk-devicecreationfile
	DeviceCreationFile *string `json:"DeviceCreationFile,omitempty"`

	// DeviceCreationFileList AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotwireless-wirelessdeviceimporttask-sidewalk.html#cfn-iotwireless-wirelessdeviceimporttask-sidewalk-devicecreationfilelist
	DeviceCreationFileList []string `json:"DeviceCreationFileList,omitempty"`

	// Role AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotwireless-wirelessdeviceimporttask-sidewalk.html#cfn-iotwireless-wirelessdeviceimporttask-sidewalk-role
	Role *string `json:"Role,omitempty"`

	// SidewalkManufacturingSn AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotwireless-wirelessdeviceimporttask-sidewalk.html#cfn-iotwireless-wirelessdeviceimporttask-sidewalk-sidewalkmanufacturingsn
	SidewalkManufacturingSn *string `json:"SidewalkManufacturingSn,omitempty"`

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
func (r *WirelessDeviceImportTask_Sidewalk[any]) AWSCloudFormationType() string {
	return "AWS::IoTWireless::WirelessDeviceImportTask.Sidewalk"
}
