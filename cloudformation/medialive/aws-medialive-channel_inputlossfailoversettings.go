// Code generated by "go generate". Please don't change this file directly.

package medialive

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// Channel_InputLossFailoverSettings AWS CloudFormation Resource (AWS::MediaLive::Channel.InputLossFailoverSettings)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-inputlossfailoversettings.html
type Channel_InputLossFailoverSettings[T any] struct {

	// InputLossThresholdMsec AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-inputlossfailoversettings.html#cfn-medialive-channel-inputlossfailoversettings-inputlossthresholdmsec
	InputLossThresholdMsec *T `json:"InputLossThresholdMsec,omitempty"`

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
func (r *Channel_InputLossFailoverSettings[any]) AWSCloudFormationType() string {
	return "AWS::MediaLive::Channel.InputLossFailoverSettings"
}
