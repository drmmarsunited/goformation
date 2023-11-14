// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// Template_TemplateVersion AWS CloudFormation Resource (AWS::QuickSight::Template.TemplateVersion)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-templateversion.html
type Template_TemplateVersion[T any] struct {

	// CreatedTime AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-templateversion.html#cfn-quicksight-template-templateversion-createdtime
	CreatedTime *T `json:"CreatedTime,omitempty"`

	// DataSetConfigurations AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-templateversion.html#cfn-quicksight-template-templateversion-datasetconfigurations
	DataSetConfigurations []Template_DataSetConfiguration[any] `json:"DataSetConfigurations,omitempty"`

	// Description AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-templateversion.html#cfn-quicksight-template-templateversion-description
	Description *T `json:"Description,omitempty"`

	// Errors AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-templateversion.html#cfn-quicksight-template-templateversion-errors
	Errors []Template_TemplateError[any] `json:"Errors,omitempty"`

	// Options AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-templateversion.html#cfn-quicksight-template-templateversion-options
	Options *Template_AssetOptions[any] `json:"Options,omitempty"`

	// Sheets AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-templateversion.html#cfn-quicksight-template-templateversion-sheets
	Sheets []Template_Sheet[any] `json:"Sheets,omitempty"`

	// SourceEntityArn AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-templateversion.html#cfn-quicksight-template-templateversion-sourceentityarn
	SourceEntityArn *T `json:"SourceEntityArn,omitempty"`

	// Status AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-templateversion.html#cfn-quicksight-template-templateversion-status
	Status *T `json:"Status,omitempty"`

	// ThemeArn AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-templateversion.html#cfn-quicksight-template-templateversion-themearn
	ThemeArn *T `json:"ThemeArn,omitempty"`

	// VersionNumber AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-templateversion.html#cfn-quicksight-template-templateversion-versionnumber
	VersionNumber *T `json:"VersionNumber,omitempty"`

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
func (r *Template_TemplateVersion[any]) AWSCloudFormationType() string {
	return "AWS::QuickSight::Template.TemplateVersion"
}
