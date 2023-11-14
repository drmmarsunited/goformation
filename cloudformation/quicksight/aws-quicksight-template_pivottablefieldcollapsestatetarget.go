// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// Template_PivotTableFieldCollapseStateTarget AWS CloudFormation Resource (AWS::QuickSight::Template.PivotTableFieldCollapseStateTarget)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-pivottablefieldcollapsestatetarget.html
type Template_PivotTableFieldCollapseStateTarget[T any] struct {

	// FieldDataPathValues AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-pivottablefieldcollapsestatetarget.html#cfn-quicksight-template-pivottablefieldcollapsestatetarget-fielddatapathvalues
	FieldDataPathValues []Template_DataPathValue[any] `json:"FieldDataPathValues,omitempty"`

	// FieldId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-pivottablefieldcollapsestatetarget.html#cfn-quicksight-template-pivottablefieldcollapsestatetarget-fieldid
	FieldId *T `json:"FieldId,omitempty"`

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
func (r *Template_PivotTableFieldCollapseStateTarget[any]) AWSCloudFormationType() string {
	return "AWS::QuickSight::Template.PivotTableFieldCollapseStateTarget"
}