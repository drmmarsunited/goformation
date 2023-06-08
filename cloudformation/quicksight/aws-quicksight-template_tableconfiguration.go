// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// Template_TableConfiguration AWS CloudFormation Resource (AWS::QuickSight::Template.TableConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-tableconfiguration.html
type Template_TableConfiguration[T any] struct {

	// FieldOptions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-tableconfiguration.html#cfn-quicksight-template-tableconfiguration-fieldoptions
	FieldOptions *Template_TableFieldOptions[any] `json:"FieldOptions,omitempty"`

	// FieldWells AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-tableconfiguration.html#cfn-quicksight-template-tableconfiguration-fieldwells
	FieldWells *Template_TableFieldWells[any] `json:"FieldWells,omitempty"`

	// PaginatedReportOptions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-tableconfiguration.html#cfn-quicksight-template-tableconfiguration-paginatedreportoptions
	PaginatedReportOptions *Template_TablePaginatedReportOptions[any] `json:"PaginatedReportOptions,omitempty"`

	// SortConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-tableconfiguration.html#cfn-quicksight-template-tableconfiguration-sortconfiguration
	SortConfiguration *Template_TableSortConfiguration[any] `json:"SortConfiguration,omitempty"`

	// TableInlineVisualizations AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-tableconfiguration.html#cfn-quicksight-template-tableconfiguration-tableinlinevisualizations
	TableInlineVisualizations []Template_TableInlineVisualization[any] `json:"TableInlineVisualizations,omitempty"`

	// TableOptions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-tableconfiguration.html#cfn-quicksight-template-tableconfiguration-tableoptions
	TableOptions *Template_TableOptions[any] `json:"TableOptions,omitempty"`

	// TotalOptions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-tableconfiguration.html#cfn-quicksight-template-tableconfiguration-totaloptions
	TotalOptions *Template_TotalOptions[any] `json:"TotalOptions,omitempty"`

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
func (r *Template_TableConfiguration[any]) AWSCloudFormationType() string {
	return "AWS::QuickSight::Template.TableConfiguration"
}
