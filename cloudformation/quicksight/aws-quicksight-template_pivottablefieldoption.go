// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// Template_PivotTableFieldOption AWS CloudFormation Resource (AWS::QuickSight::Template.PivotTableFieldOption)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-pivottablefieldoption.html
type Template_PivotTableFieldOption[T any] struct {

	// CustomLabel AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-pivottablefieldoption.html#cfn-quicksight-template-pivottablefieldoption-customlabel
	CustomLabel *string `json:"CustomLabel,omitempty"`

	// FieldId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-pivottablefieldoption.html#cfn-quicksight-template-pivottablefieldoption-fieldid
	FieldId string `json:"FieldId"`

	// Visibility AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-pivottablefieldoption.html#cfn-quicksight-template-pivottablefieldoption-visibility
	Visibility *string `json:"Visibility,omitempty"`

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
func (r *Template_PivotTableFieldOption[any]) AWSCloudFormationType() string {
	return "AWS::QuickSight::Template.PivotTableFieldOption"
}
