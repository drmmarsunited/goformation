// Code generated by "go generate". Please don't change this file directly.

package networkmanager

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
	"github.com/drmmarsunited/goformation/v7/cloudformation/tags"
)

// VpcAttachment_ProposedSegmentChange AWS CloudFormation Resource (AWS::NetworkManager::VpcAttachment.ProposedSegmentChange)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkmanager-vpcattachment-proposedsegmentchange.html
type VpcAttachment_ProposedSegmentChange[T any] struct {

	// AttachmentPolicyRuleNumber AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkmanager-vpcattachment-proposedsegmentchange.html#cfn-networkmanager-vpcattachment-proposedsegmentchange-attachmentpolicyrulenumber
	AttachmentPolicyRuleNumber *T `json:"AttachmentPolicyRuleNumber,omitempty"`

	// SegmentName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkmanager-vpcattachment-proposedsegmentchange.html#cfn-networkmanager-vpcattachment-proposedsegmentchange-segmentname
	SegmentName *string `json:"SegmentName,omitempty"`

	// Tags AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkmanager-vpcattachment-proposedsegmentchange.html#cfn-networkmanager-vpcattachment-proposedsegmentchange-tags
	Tags []tags.Tag `json:"Tags,omitempty"`

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
func (r *VpcAttachment_ProposedSegmentChange[any]) AWSCloudFormationType() string {
	return "AWS::NetworkManager::VpcAttachment.ProposedSegmentChange"
}
