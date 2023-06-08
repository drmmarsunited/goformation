// Code generated by "go generate". Please don't change this file directly.

package connect

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// TaskTemplate_ReadOnlyFieldInfo AWS CloudFormation Resource (AWS::Connect::TaskTemplate.ReadOnlyFieldInfo)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-tasktemplate-readonlyfieldinfo.html
type TaskTemplate_ReadOnlyFieldInfo[T any] struct {

	// Id AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-tasktemplate-readonlyfieldinfo.html#cfn-connect-tasktemplate-readonlyfieldinfo-id
	Id *TaskTemplate_FieldIdentifier[any] `json:"Id"`

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
func (r *TaskTemplate_ReadOnlyFieldInfo[any]) AWSCloudFormationType() string {
	return "AWS::Connect::TaskTemplate.ReadOnlyFieldInfo"
}
