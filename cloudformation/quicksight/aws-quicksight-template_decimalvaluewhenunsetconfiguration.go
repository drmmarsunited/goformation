// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// Template_DecimalValueWhenUnsetConfiguration AWS CloudFormation Resource (AWS::QuickSight::Template.DecimalValueWhenUnsetConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-decimalvaluewhenunsetconfiguration.html
type Template_DecimalValueWhenUnsetConfiguration[T any] struct {

	// CustomValue AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-decimalvaluewhenunsetconfiguration.html#cfn-quicksight-template-decimalvaluewhenunsetconfiguration-customvalue
	CustomValue *T `json:"CustomValue,omitempty"`

	// ValueWhenUnsetOption AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-decimalvaluewhenunsetconfiguration.html#cfn-quicksight-template-decimalvaluewhenunsetconfiguration-valuewhenunsetoption
	ValueWhenUnsetOption *T `json:"ValueWhenUnsetOption,omitempty"`

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
func (r *Template_DecimalValueWhenUnsetConfiguration[any]) AWSCloudFormationType() string {
	return "AWS::QuickSight::Template.DecimalValueWhenUnsetConfiguration"
}
