// Code generated by "go generate". Please don't change this file directly.

package fsx

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// FileSystem_WindowsConfiguration AWS CloudFormation Resource (AWS::FSx::FileSystem.WindowsConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-filesystem-windowsconfiguration.html
type FileSystem_WindowsConfiguration[T any] struct {

	// ActiveDirectoryId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-filesystem-windowsconfiguration.html#cfn-fsx-filesystem-windowsconfiguration-activedirectoryid
	ActiveDirectoryId *string `json:"ActiveDirectoryId,omitempty"`

	// Aliases AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-filesystem-windowsconfiguration.html#cfn-fsx-filesystem-windowsconfiguration-aliases
	Aliases []string `json:"Aliases,omitempty"`

	// AuditLogConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-filesystem-windowsconfiguration.html#cfn-fsx-filesystem-windowsconfiguration-auditlogconfiguration
	AuditLogConfiguration *FileSystem_AuditLogConfiguration[any] `json:"AuditLogConfiguration,omitempty"`

	// AutomaticBackupRetentionDays AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-filesystem-windowsconfiguration.html#cfn-fsx-filesystem-windowsconfiguration-automaticbackupretentiondays
	AutomaticBackupRetentionDays *T `json:"AutomaticBackupRetentionDays,omitempty"`

	// CopyTagsToBackups AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-filesystem-windowsconfiguration.html#cfn-fsx-filesystem-windowsconfiguration-copytagstobackups
	CopyTagsToBackups *T `json:"CopyTagsToBackups,omitempty"`

	// DailyAutomaticBackupStartTime AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-filesystem-windowsconfiguration.html#cfn-fsx-filesystem-windowsconfiguration-dailyautomaticbackupstarttime
	DailyAutomaticBackupStartTime *string `json:"DailyAutomaticBackupStartTime,omitempty"`

	// DeploymentType AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-filesystem-windowsconfiguration.html#cfn-fsx-filesystem-windowsconfiguration-deploymenttype
	DeploymentType *string `json:"DeploymentType,omitempty"`

	// DiskIopsConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-filesystem-windowsconfiguration.html#cfn-fsx-filesystem-windowsconfiguration-diskiopsconfiguration
	DiskIopsConfiguration *FileSystem_DiskIopsConfiguration[any] `json:"DiskIopsConfiguration,omitempty"`

	// PreferredSubnetId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-filesystem-windowsconfiguration.html#cfn-fsx-filesystem-windowsconfiguration-preferredsubnetid
	PreferredSubnetId *string `json:"PreferredSubnetId,omitempty"`

	// SelfManagedActiveDirectoryConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-filesystem-windowsconfiguration.html#cfn-fsx-filesystem-windowsconfiguration-selfmanagedactivedirectoryconfiguration
	SelfManagedActiveDirectoryConfiguration *FileSystem_SelfManagedActiveDirectoryConfiguration[any] `json:"SelfManagedActiveDirectoryConfiguration,omitempty"`

	// ThroughputCapacity AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-filesystem-windowsconfiguration.html#cfn-fsx-filesystem-windowsconfiguration-throughputcapacity
	ThroughputCapacity T `json:"ThroughputCapacity"`

	// WeeklyMaintenanceStartTime AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-filesystem-windowsconfiguration.html#cfn-fsx-filesystem-windowsconfiguration-weeklymaintenancestarttime
	WeeklyMaintenanceStartTime *string `json:"WeeklyMaintenanceStartTime,omitempty"`

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
func (r *FileSystem_WindowsConfiguration[any]) AWSCloudFormationType() string {
	return "AWS::FSx::FileSystem.WindowsConfiguration"
}
