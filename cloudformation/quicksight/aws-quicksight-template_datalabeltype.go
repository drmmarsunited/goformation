// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// Template_DataLabelType AWS CloudFormation Resource (AWS::QuickSight::Template.DataLabelType)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-datalabeltype.html
type Template_DataLabelType struct {

	// DataPathLabelType AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-datalabeltype.html#cfn-quicksight-template-datalabeltype-datapathlabeltype
	DataPathLabelType *Template_DataPathLabelType `json:"DataPathLabelType,omitempty"`

	// FieldLabelType AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-datalabeltype.html#cfn-quicksight-template-datalabeltype-fieldlabeltype
	FieldLabelType *Template_FieldLabelType `json:"FieldLabelType,omitempty"`

	// MaximumLabelType AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-datalabeltype.html#cfn-quicksight-template-datalabeltype-maximumlabeltype
	MaximumLabelType *Template_MaximumLabelType `json:"MaximumLabelType,omitempty"`

	// MinimumLabelType AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-datalabeltype.html#cfn-quicksight-template-datalabeltype-minimumlabeltype
	MinimumLabelType *Template_MinimumLabelType `json:"MinimumLabelType,omitempty"`

	// RangeEndsLabelType AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-datalabeltype.html#cfn-quicksight-template-datalabeltype-rangeendslabeltype
	RangeEndsLabelType *Template_RangeEndsLabelType `json:"RangeEndsLabelType,omitempty"`

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
func (r *Template_DataLabelType) AWSCloudFormationType() string {
	return "AWS::QuickSight::Template.DataLabelType"
}
