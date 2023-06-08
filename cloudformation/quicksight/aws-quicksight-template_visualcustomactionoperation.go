// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// Template_VisualCustomActionOperation AWS CloudFormation Resource (AWS::QuickSight::Template.VisualCustomActionOperation)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-visualcustomactionoperation.html
type Template_VisualCustomActionOperation[T any] struct {

	// FilterOperation AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-visualcustomactionoperation.html#cfn-quicksight-template-visualcustomactionoperation-filteroperation
	FilterOperation *Template_CustomActionFilterOperation[any] `json:"FilterOperation,omitempty"`

	// NavigationOperation AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-visualcustomactionoperation.html#cfn-quicksight-template-visualcustomactionoperation-navigationoperation
	NavigationOperation *Template_CustomActionNavigationOperation[any] `json:"NavigationOperation,omitempty"`

	// SetParametersOperation AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-visualcustomactionoperation.html#cfn-quicksight-template-visualcustomactionoperation-setparametersoperation
	SetParametersOperation *Template_CustomActionSetParametersOperation[any] `json:"SetParametersOperation,omitempty"`

	// URLOperation AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-visualcustomactionoperation.html#cfn-quicksight-template-visualcustomactionoperation-urloperation
	URLOperation *Template_CustomActionURLOperation[any] `json:"URLOperation,omitempty"`

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
func (r *Template_VisualCustomActionOperation[any]) AWSCloudFormationType() string {
	return "AWS::QuickSight::Template.VisualCustomActionOperation"
}
