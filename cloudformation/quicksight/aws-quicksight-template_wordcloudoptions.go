// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// Template_WordCloudOptions AWS CloudFormation Resource (AWS::QuickSight::Template.WordCloudOptions)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-wordcloudoptions.html
type Template_WordCloudOptions[T any] struct {

	// CloudLayout AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-wordcloudoptions.html#cfn-quicksight-template-wordcloudoptions-cloudlayout
	CloudLayout *T `json:"CloudLayout,omitempty"`

	// MaximumStringLength AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-wordcloudoptions.html#cfn-quicksight-template-wordcloudoptions-maximumstringlength
	MaximumStringLength *T `json:"MaximumStringLength,omitempty"`

	// WordCasing AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-wordcloudoptions.html#cfn-quicksight-template-wordcloudoptions-wordcasing
	WordCasing *T `json:"WordCasing,omitempty"`

	// WordOrientation AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-wordcloudoptions.html#cfn-quicksight-template-wordcloudoptions-wordorientation
	WordOrientation *T `json:"WordOrientation,omitempty"`

	// WordPadding AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-wordcloudoptions.html#cfn-quicksight-template-wordcloudoptions-wordpadding
	WordPadding *T `json:"WordPadding,omitempty"`

	// WordScaling AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-wordcloudoptions.html#cfn-quicksight-template-wordcloudoptions-wordscaling
	WordScaling *T `json:"WordScaling,omitempty"`

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
func (r *Template_WordCloudOptions[any]) AWSCloudFormationType() string {
	return "AWS::QuickSight::Template.WordCloudOptions"
}
