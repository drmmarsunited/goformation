// Code generated by "go generate". Please don't change this file directly.

package msk

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// Replicator_ConsumerGroupReplication AWS CloudFormation Resource (AWS::MSK::Replicator.ConsumerGroupReplication)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-replicator-consumergroupreplication.html
type Replicator_ConsumerGroupReplication[T any] struct {

	// ConsumerGroupsToExclude AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-replicator-consumergroupreplication.html#cfn-msk-replicator-consumergroupreplication-consumergroupstoexclude
	ConsumerGroupsToExclude []T `json:"ConsumerGroupsToExclude,omitempty"`

	// ConsumerGroupsToReplicate AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-replicator-consumergroupreplication.html#cfn-msk-replicator-consumergroupreplication-consumergroupstoreplicate
	ConsumerGroupsToReplicate []T `json:"ConsumerGroupsToReplicate"`

	// DetectAndCopyNewConsumerGroups AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-replicator-consumergroupreplication.html#cfn-msk-replicator-consumergroupreplication-detectandcopynewconsumergroups
	DetectAndCopyNewConsumerGroups *T `json:"DetectAndCopyNewConsumerGroups,omitempty"`

	// SynchroniseConsumerGroupOffsets AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-replicator-consumergroupreplication.html#cfn-msk-replicator-consumergroupreplication-synchroniseconsumergroupoffsets
	SynchroniseConsumerGroupOffsets *T `json:"SynchroniseConsumerGroupOffsets,omitempty"`

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
func (r *Replicator_ConsumerGroupReplication[any]) AWSCloudFormationType() string {
	return "AWS::MSK::Replicator.ConsumerGroupReplication"
}
