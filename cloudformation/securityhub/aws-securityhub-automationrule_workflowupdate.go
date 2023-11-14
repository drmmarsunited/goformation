// Code generated by "go generate". Please don't change this file directly.

package securityhub

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// AutomationRule_WorkflowUpdate AWS CloudFormation Resource (AWS::SecurityHub::AutomationRule.WorkflowUpdate)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-automationrule-workflowupdate.html
type AutomationRule_WorkflowUpdate[T any] struct {

	// Status AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-automationrule-workflowupdate.html#cfn-securityhub-automationrule-workflowupdate-status
	Status T `json:"Status"`

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
func (r *AutomationRule_WorkflowUpdate[any]) AWSCloudFormationType() string {
	return "AWS::SecurityHub::AutomationRule.WorkflowUpdate"
}
