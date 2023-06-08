// Code generated by "go generate". Please don't change this file directly.

package lakeformation

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// DataLakeSettings_CreateDatabaseDefaultPermissions AWS CloudFormation Resource (AWS::LakeFormation::DataLakeSettings.CreateDatabaseDefaultPermissions)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lakeformation-datalakesettings-createdatabasedefaultpermissions.html
type DataLakeSettings_CreateDatabaseDefaultPermissions[T any] struct {

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
func (r *DataLakeSettings_CreateDatabaseDefaultPermissions[any]) AWSCloudFormationType() string {
	return "AWS::LakeFormation::DataLakeSettings.CreateDatabaseDefaultPermissions"
}
