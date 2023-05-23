// Code generated by "go generate". Please don't change this file directly.

package rum

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// AppMonitor_CustomEvents AWS CloudFormation Resource (AWS::RUM::AppMonitor.CustomEvents)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rum-appmonitor-customevents.html
type AppMonitor_CustomEvents struct {

	// Status AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rum-appmonitor-customevents.html#cfn-rum-appmonitor-customevents-status
	Status *string `json:"Status,omitempty"`

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
func (r *AppMonitor_CustomEvents) AWSCloudFormationType() string {
	return "AWS::RUM::AppMonitor.CustomEvents"
}
