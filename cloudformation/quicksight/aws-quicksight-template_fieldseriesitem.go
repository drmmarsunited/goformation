// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// Template_FieldSeriesItem AWS CloudFormation Resource (AWS::QuickSight::Template.FieldSeriesItem)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-fieldseriesitem.html
type Template_FieldSeriesItem[T any] struct {

	// AxisBinding AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-fieldseriesitem.html#cfn-quicksight-template-fieldseriesitem-axisbinding
	AxisBinding string `json:"AxisBinding"`

	// FieldId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-fieldseriesitem.html#cfn-quicksight-template-fieldseriesitem-fieldid
	FieldId string `json:"FieldId"`

	// Settings AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-fieldseriesitem.html#cfn-quicksight-template-fieldseriesitem-settings
	Settings *Template_LineChartSeriesSettings[any] `json:"Settings,omitempty"`

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
func (r *Template_FieldSeriesItem[any]) AWSCloudFormationType() string {
	return "AWS::QuickSight::Template.FieldSeriesItem"
}
