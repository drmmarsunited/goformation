// Code generated by "go generate". Please don't change this file directly.

package kendra

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// DataSource_ColumnConfiguration AWS CloudFormation Resource (AWS::Kendra::DataSource.ColumnConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-columnconfiguration.html
type DataSource_ColumnConfiguration[T any] struct {

	// ChangeDetectingColumns AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-columnconfiguration.html#cfn-kendra-datasource-columnconfiguration-changedetectingcolumns
	ChangeDetectingColumns []string `json:"ChangeDetectingColumns"`

	// DocumentDataColumnName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-columnconfiguration.html#cfn-kendra-datasource-columnconfiguration-documentdatacolumnname
	DocumentDataColumnName string `json:"DocumentDataColumnName"`

	// DocumentIdColumnName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-columnconfiguration.html#cfn-kendra-datasource-columnconfiguration-documentidcolumnname
	DocumentIdColumnName string `json:"DocumentIdColumnName"`

	// DocumentTitleColumnName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-columnconfiguration.html#cfn-kendra-datasource-columnconfiguration-documenttitlecolumnname
	DocumentTitleColumnName *string `json:"DocumentTitleColumnName,omitempty"`

	// FieldMappings AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-columnconfiguration.html#cfn-kendra-datasource-columnconfiguration-fieldmappings
	FieldMappings []DataSource_DataSourceToIndexFieldMapping[any] `json:"FieldMappings,omitempty"`

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
func (r *DataSource_ColumnConfiguration[any]) AWSCloudFormationType() string {
	return "AWS::Kendra::DataSource.ColumnConfiguration"
}
