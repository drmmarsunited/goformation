// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// Dashboard_FilterOperationSelectedFieldsConfiguration AWS CloudFormation Resource (AWS::QuickSight::Dashboard.FilterOperationSelectedFieldsConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-filteroperationselectedfieldsconfiguration.html
type Dashboard_FilterOperationSelectedFieldsConfiguration[T any] struct {

	// SelectedFieldOptions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-filteroperationselectedfieldsconfiguration.html#cfn-quicksight-dashboard-filteroperationselectedfieldsconfiguration-selectedfieldoptions
	SelectedFieldOptions *string `json:"SelectedFieldOptions,omitempty"`

	// SelectedFields AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-filteroperationselectedfieldsconfiguration.html#cfn-quicksight-dashboard-filteroperationselectedfieldsconfiguration-selectedfields
	SelectedFields []string `json:"SelectedFields,omitempty"`

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
func (r *Dashboard_FilterOperationSelectedFieldsConfiguration[any]) AWSCloudFormationType() string {
	return "AWS::QuickSight::Dashboard.FilterOperationSelectedFieldsConfiguration"
}
