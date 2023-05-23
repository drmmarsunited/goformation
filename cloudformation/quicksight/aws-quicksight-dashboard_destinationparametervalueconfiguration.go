// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// Dashboard_DestinationParameterValueConfiguration AWS CloudFormation Resource (AWS::QuickSight::Dashboard.DestinationParameterValueConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-destinationparametervalueconfiguration.html
type Dashboard_DestinationParameterValueConfiguration struct {

	// CustomValuesConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-destinationparametervalueconfiguration.html#cfn-quicksight-dashboard-destinationparametervalueconfiguration-customvaluesconfiguration
	CustomValuesConfiguration *Dashboard_CustomValuesConfiguration `json:"CustomValuesConfiguration,omitempty"`

	// SelectAllValueOptions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-destinationparametervalueconfiguration.html#cfn-quicksight-dashboard-destinationparametervalueconfiguration-selectallvalueoptions
	SelectAllValueOptions *string `json:"SelectAllValueOptions,omitempty"`

	// SourceField AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-destinationparametervalueconfiguration.html#cfn-quicksight-dashboard-destinationparametervalueconfiguration-sourcefield
	SourceField *string `json:"SourceField,omitempty"`

	// SourceParameterName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-destinationparametervalueconfiguration.html#cfn-quicksight-dashboard-destinationparametervalueconfiguration-sourceparametername
	SourceParameterName *string `json:"SourceParameterName,omitempty"`

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
func (r *Dashboard_DestinationParameterValueConfiguration) AWSCloudFormationType() string {
	return "AWS::QuickSight::Dashboard.DestinationParameterValueConfiguration"
}
