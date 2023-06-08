// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// Analysis_GeospatialWindowOptions AWS CloudFormation Resource (AWS::QuickSight::Analysis.GeospatialWindowOptions)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-geospatialwindowoptions.html
type Analysis_GeospatialWindowOptions[T any] struct {

	// Bounds AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-geospatialwindowoptions.html#cfn-quicksight-analysis-geospatialwindowoptions-bounds
	Bounds *Analysis_GeospatialCoordinateBounds[any] `json:"Bounds,omitempty"`

	// MapZoomMode AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-geospatialwindowoptions.html#cfn-quicksight-analysis-geospatialwindowoptions-mapzoommode
	MapZoomMode *string `json:"MapZoomMode,omitempty"`

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
func (r *Analysis_GeospatialWindowOptions[any]) AWSCloudFormationType() string {
	return "AWS::QuickSight::Analysis.GeospatialWindowOptions"
}
