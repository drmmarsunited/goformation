// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// Template_DateTimeParameterDeclaration AWS CloudFormation Resource (AWS::QuickSight::Template.DateTimeParameterDeclaration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-datetimeparameterdeclaration.html
type Template_DateTimeParameterDeclaration struct {

	// DefaultValues AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-datetimeparameterdeclaration.html#cfn-quicksight-template-datetimeparameterdeclaration-defaultvalues
	DefaultValues *Template_DateTimeDefaultValues `json:"DefaultValues,omitempty"`

	// MappedDataSetParameters AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-datetimeparameterdeclaration.html#cfn-quicksight-template-datetimeparameterdeclaration-mappeddatasetparameters
	MappedDataSetParameters []Template_MappedDataSetParameter `json:"MappedDataSetParameters,omitempty"`

	// Name AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-datetimeparameterdeclaration.html#cfn-quicksight-template-datetimeparameterdeclaration-name
	Name string `json:"Name"`

	// TimeGranularity AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-datetimeparameterdeclaration.html#cfn-quicksight-template-datetimeparameterdeclaration-timegranularity
	TimeGranularity *string `json:"TimeGranularity,omitempty"`

	// ValueWhenUnset AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-datetimeparameterdeclaration.html#cfn-quicksight-template-datetimeparameterdeclaration-valuewhenunset
	ValueWhenUnset *Template_DateTimeValueWhenUnsetConfiguration `json:"ValueWhenUnset,omitempty"`

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
func (r *Template_DateTimeParameterDeclaration) AWSCloudFormationType() string {
	return "AWS::QuickSight::Template.DateTimeParameterDeclaration"
}
