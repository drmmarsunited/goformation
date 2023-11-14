// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// Dashboard_DefaultNewSheetConfiguration AWS CloudFormation Resource (AWS::QuickSight::Dashboard.DefaultNewSheetConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-defaultnewsheetconfiguration.html
type Dashboard_DefaultNewSheetConfiguration[T any] struct {

	// InteractiveLayoutConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-defaultnewsheetconfiguration.html#cfn-quicksight-dashboard-defaultnewsheetconfiguration-interactivelayoutconfiguration
	InteractiveLayoutConfiguration *Dashboard_DefaultInteractiveLayoutConfiguration[any] `json:"InteractiveLayoutConfiguration,omitempty"`

	// PaginatedLayoutConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-defaultnewsheetconfiguration.html#cfn-quicksight-dashboard-defaultnewsheetconfiguration-paginatedlayoutconfiguration
	PaginatedLayoutConfiguration *Dashboard_DefaultPaginatedLayoutConfiguration[any] `json:"PaginatedLayoutConfiguration,omitempty"`

	// SheetContentType AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-defaultnewsheetconfiguration.html#cfn-quicksight-dashboard-defaultnewsheetconfiguration-sheetcontenttype
	SheetContentType *T `json:"SheetContentType,omitempty"`

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
func (r *Dashboard_DefaultNewSheetConfiguration[any]) AWSCloudFormationType() string {
	return "AWS::QuickSight::Dashboard.DefaultNewSheetConfiguration"
}
