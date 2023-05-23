// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// Analysis_DefaultNewSheetConfiguration AWS CloudFormation Resource (AWS::QuickSight::Analysis.DefaultNewSheetConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-defaultnewsheetconfiguration.html
type Analysis_DefaultNewSheetConfiguration struct {

	// InteractiveLayoutConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-defaultnewsheetconfiguration.html#cfn-quicksight-analysis-defaultnewsheetconfiguration-interactivelayoutconfiguration
	InteractiveLayoutConfiguration *Analysis_DefaultInteractiveLayoutConfiguration `json:"InteractiveLayoutConfiguration,omitempty"`

	// PaginatedLayoutConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-defaultnewsheetconfiguration.html#cfn-quicksight-analysis-defaultnewsheetconfiguration-paginatedlayoutconfiguration
	PaginatedLayoutConfiguration *Analysis_DefaultPaginatedLayoutConfiguration `json:"PaginatedLayoutConfiguration,omitempty"`

	// SheetContentType AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-defaultnewsheetconfiguration.html#cfn-quicksight-analysis-defaultnewsheetconfiguration-sheetcontenttype
	SheetContentType *string `json:"SheetContentType,omitempty"`

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
func (r *Analysis_DefaultNewSheetConfiguration) AWSCloudFormationType() string {
	return "AWS::QuickSight::Analysis.DefaultNewSheetConfiguration"
}
