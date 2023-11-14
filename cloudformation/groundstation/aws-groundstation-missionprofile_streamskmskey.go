// Code generated by "go generate". Please don't change this file directly.

package groundstation

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// MissionProfile_StreamsKmsKey AWS CloudFormation Resource (AWS::GroundStation::MissionProfile.StreamsKmsKey)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-missionprofile-streamskmskey.html
type MissionProfile_StreamsKmsKey[T any] struct {

	// KmsAliasArn AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-missionprofile-streamskmskey.html#cfn-groundstation-missionprofile-streamskmskey-kmsaliasarn
	KmsAliasArn *T `json:"KmsAliasArn,omitempty"`

	// KmsKeyArn AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-missionprofile-streamskmskey.html#cfn-groundstation-missionprofile-streamskmskey-kmskeyarn
	KmsKeyArn *T `json:"KmsKeyArn,omitempty"`

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
func (r *MissionProfile_StreamsKmsKey[any]) AWSCloudFormationType() string {
	return "AWS::GroundStation::MissionProfile.StreamsKmsKey"
}
