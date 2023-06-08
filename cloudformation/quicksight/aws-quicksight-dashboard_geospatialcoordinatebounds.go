// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// Dashboard_GeospatialCoordinateBounds AWS CloudFormation Resource (AWS::QuickSight::Dashboard.GeospatialCoordinateBounds)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-geospatialcoordinatebounds.html
type Dashboard_GeospatialCoordinateBounds[T any] struct {

	// East AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-geospatialcoordinatebounds.html#cfn-quicksight-dashboard-geospatialcoordinatebounds-east
	East T `json:"East"`

	// North AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-geospatialcoordinatebounds.html#cfn-quicksight-dashboard-geospatialcoordinatebounds-north
	North T `json:"North"`

	// South AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-geospatialcoordinatebounds.html#cfn-quicksight-dashboard-geospatialcoordinatebounds-south
	South T `json:"South"`

	// West AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-geospatialcoordinatebounds.html#cfn-quicksight-dashboard-geospatialcoordinatebounds-west
	West T `json:"West"`

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
func (r *Dashboard_GeospatialCoordinateBounds[any]) AWSCloudFormationType() string {
	return "AWS::QuickSight::Dashboard.GeospatialCoordinateBounds"
}
