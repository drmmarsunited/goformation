// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// Dashboard_DashboardPublishOptions AWS CloudFormation Resource (AWS::QuickSight::Dashboard.DashboardPublishOptions)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-dashboardpublishoptions.html
type Dashboard_DashboardPublishOptions struct {

	// AdHocFilteringOption AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-dashboardpublishoptions.html#cfn-quicksight-dashboard-dashboardpublishoptions-adhocfilteringoption
	AdHocFilteringOption *Dashboard_AdHocFilteringOption `json:"AdHocFilteringOption,omitempty"`

	// DataPointDrillUpDownOption AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-dashboardpublishoptions.html#cfn-quicksight-dashboard-dashboardpublishoptions-datapointdrillupdownoption
	DataPointDrillUpDownOption *Dashboard_DataPointDrillUpDownOption `json:"DataPointDrillUpDownOption,omitempty"`

	// DataPointMenuLabelOption AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-dashboardpublishoptions.html#cfn-quicksight-dashboard-dashboardpublishoptions-datapointmenulabeloption
	DataPointMenuLabelOption *Dashboard_DataPointMenuLabelOption `json:"DataPointMenuLabelOption,omitempty"`

	// DataPointTooltipOption AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-dashboardpublishoptions.html#cfn-quicksight-dashboard-dashboardpublishoptions-datapointtooltipoption
	DataPointTooltipOption *Dashboard_DataPointTooltipOption `json:"DataPointTooltipOption,omitempty"`

	// ExportToCSVOption AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-dashboardpublishoptions.html#cfn-quicksight-dashboard-dashboardpublishoptions-exporttocsvoption
	ExportToCSVOption *Dashboard_ExportToCSVOption `json:"ExportToCSVOption,omitempty"`

	// ExportWithHiddenFieldsOption AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-dashboardpublishoptions.html#cfn-quicksight-dashboard-dashboardpublishoptions-exportwithhiddenfieldsoption
	ExportWithHiddenFieldsOption *Dashboard_ExportWithHiddenFieldsOption `json:"ExportWithHiddenFieldsOption,omitempty"`

	// SheetControlsOption AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-dashboardpublishoptions.html#cfn-quicksight-dashboard-dashboardpublishoptions-sheetcontrolsoption
	SheetControlsOption *Dashboard_SheetControlsOption `json:"SheetControlsOption,omitempty"`

	// SheetLayoutElementMaximizationOption AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-dashboardpublishoptions.html#cfn-quicksight-dashboard-dashboardpublishoptions-sheetlayoutelementmaximizationoption
	SheetLayoutElementMaximizationOption *Dashboard_SheetLayoutElementMaximizationOption `json:"SheetLayoutElementMaximizationOption,omitempty"`

	// VisualAxisSortOption AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-dashboardpublishoptions.html#cfn-quicksight-dashboard-dashboardpublishoptions-visualaxissortoption
	VisualAxisSortOption *Dashboard_VisualAxisSortOption `json:"VisualAxisSortOption,omitempty"`

	// VisualMenuOption AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-dashboardpublishoptions.html#cfn-quicksight-dashboard-dashboardpublishoptions-visualmenuoption
	VisualMenuOption *Dashboard_VisualMenuOption `json:"VisualMenuOption,omitempty"`

	// VisualPublishOptions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-dashboardpublishoptions.html#cfn-quicksight-dashboard-dashboardpublishoptions-visualpublishoptions
	VisualPublishOptions *Dashboard_DashboardVisualPublishOptions `json:"VisualPublishOptions,omitempty"`

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
func (r *Dashboard_DashboardPublishOptions) AWSCloudFormationType() string {
	return "AWS::QuickSight::Dashboard.DashboardPublishOptions"
}
