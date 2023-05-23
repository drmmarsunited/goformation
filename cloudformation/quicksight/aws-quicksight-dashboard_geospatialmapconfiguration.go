// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// Dashboard_GeospatialMapConfiguration AWS CloudFormation Resource (AWS::QuickSight::Dashboard.GeospatialMapConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-geospatialmapconfiguration.html
type Dashboard_GeospatialMapConfiguration struct {

	// FieldWells AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-geospatialmapconfiguration.html#cfn-quicksight-dashboard-geospatialmapconfiguration-fieldwells
	FieldWells *Dashboard_GeospatialMapFieldWells `json:"FieldWells,omitempty"`

	// Legend AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-geospatialmapconfiguration.html#cfn-quicksight-dashboard-geospatialmapconfiguration-legend
	Legend *Dashboard_LegendOptions `json:"Legend,omitempty"`

	// MapStyleOptions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-geospatialmapconfiguration.html#cfn-quicksight-dashboard-geospatialmapconfiguration-mapstyleoptions
	MapStyleOptions *Dashboard_GeospatialMapStyleOptions `json:"MapStyleOptions,omitempty"`

	// PointStyleOptions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-geospatialmapconfiguration.html#cfn-quicksight-dashboard-geospatialmapconfiguration-pointstyleoptions
	PointStyleOptions *Dashboard_GeospatialPointStyleOptions `json:"PointStyleOptions,omitempty"`

	// Tooltip AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-geospatialmapconfiguration.html#cfn-quicksight-dashboard-geospatialmapconfiguration-tooltip
	Tooltip *Dashboard_TooltipOptions `json:"Tooltip,omitempty"`

	// VisualPalette AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-geospatialmapconfiguration.html#cfn-quicksight-dashboard-geospatialmapconfiguration-visualpalette
	VisualPalette *Dashboard_VisualPalette `json:"VisualPalette,omitempty"`

	// WindowOptions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-geospatialmapconfiguration.html#cfn-quicksight-dashboard-geospatialmapconfiguration-windowoptions
	WindowOptions *Dashboard_GeospatialWindowOptions `json:"WindowOptions,omitempty"`

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
func (r *Dashboard_GeospatialMapConfiguration) AWSCloudFormationType() string {
	return "AWS::QuickSight::Dashboard.GeospatialMapConfiguration"
}
