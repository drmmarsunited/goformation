// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// Analysis_GeospatialPointStyleOptions AWS CloudFormation Resource (AWS::QuickSight::Analysis.GeospatialPointStyleOptions)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-geospatialpointstyleoptions.html
type Analysis_GeospatialPointStyleOptions[T any] struct {

	// ClusterMarkerConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-geospatialpointstyleoptions.html#cfn-quicksight-analysis-geospatialpointstyleoptions-clustermarkerconfiguration
	ClusterMarkerConfiguration *Analysis_ClusterMarkerConfiguration[any] `json:"ClusterMarkerConfiguration,omitempty"`

	// HeatmapConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-geospatialpointstyleoptions.html#cfn-quicksight-analysis-geospatialpointstyleoptions-heatmapconfiguration
	HeatmapConfiguration *Analysis_GeospatialHeatmapConfiguration[any] `json:"HeatmapConfiguration,omitempty"`

	// SelectedPointStyle AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-geospatialpointstyleoptions.html#cfn-quicksight-analysis-geospatialpointstyleoptions-selectedpointstyle
	SelectedPointStyle *string `json:"SelectedPointStyle,omitempty"`

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
func (r *Analysis_GeospatialPointStyleOptions[any]) AWSCloudFormationType() string {
	return "AWS::QuickSight::Analysis.GeospatialPointStyleOptions"
}
