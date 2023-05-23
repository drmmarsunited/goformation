// Code generated by "go generate". Please don't change this file directly.

package fis

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// ExperimentTemplate_ExperimentTemplateLogConfiguration AWS CloudFormation Resource (AWS::FIS::ExperimentTemplate.ExperimentTemplateLogConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fis-experimenttemplate-experimenttemplatelogconfiguration.html
type ExperimentTemplate_ExperimentTemplateLogConfiguration struct {

	// CloudWatchLogsConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fis-experimenttemplate-experimenttemplatelogconfiguration.html#cfn-fis-experimenttemplate-experimenttemplatelogconfiguration-cloudwatchlogsconfiguration
	CloudWatchLogsConfiguration *ExperimentTemplate_CloudWatchLogsConfiguration `json:"CloudWatchLogsConfiguration,omitempty"`

	// LogSchemaVersion AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fis-experimenttemplate-experimenttemplatelogconfiguration.html#cfn-fis-experimenttemplate-experimenttemplatelogconfiguration-logschemaversion
	LogSchemaVersion int `json:"LogSchemaVersion"`

	// S3Configuration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fis-experimenttemplate-experimenttemplatelogconfiguration.html#cfn-fis-experimenttemplate-experimenttemplatelogconfiguration-s3configuration
	S3Configuration *ExperimentTemplate_S3Configuration `json:"S3Configuration,omitempty"`

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
func (r *ExperimentTemplate_ExperimentTemplateLogConfiguration) AWSCloudFormationType() string {
	return "AWS::FIS::ExperimentTemplate.ExperimentTemplateLogConfiguration"
}
