// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// Dashboard_VisualCustomActionOperation AWS CloudFormation Resource (AWS::QuickSight::Dashboard.VisualCustomActionOperation)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-visualcustomactionoperation.html
type Dashboard_VisualCustomActionOperation struct {

	// FilterOperation AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-visualcustomactionoperation.html#cfn-quicksight-dashboard-visualcustomactionoperation-filteroperation
	FilterOperation *Dashboard_CustomActionFilterOperation `json:"FilterOperation,omitempty"`

	// NavigationOperation AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-visualcustomactionoperation.html#cfn-quicksight-dashboard-visualcustomactionoperation-navigationoperation
	NavigationOperation *Dashboard_CustomActionNavigationOperation `json:"NavigationOperation,omitempty"`

	// SetParametersOperation AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-visualcustomactionoperation.html#cfn-quicksight-dashboard-visualcustomactionoperation-setparametersoperation
	SetParametersOperation *Dashboard_CustomActionSetParametersOperation `json:"SetParametersOperation,omitempty"`

	// URLOperation AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-visualcustomactionoperation.html#cfn-quicksight-dashboard-visualcustomactionoperation-urloperation
	URLOperation *Dashboard_CustomActionURLOperation `json:"URLOperation,omitempty"`

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
func (r *Dashboard_VisualCustomActionOperation) AWSCloudFormationType() string {
	return "AWS::QuickSight::Dashboard.VisualCustomActionOperation"
}
