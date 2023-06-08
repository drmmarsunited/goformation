// Code generated by "go generate". Please don't change this file directly.

package kendra

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// DataSource_DocumentAttributeCondition AWS CloudFormation Resource (AWS::Kendra::DataSource.DocumentAttributeCondition)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-documentattributecondition.html
type DataSource_DocumentAttributeCondition[T any] struct {

	// ConditionDocumentAttributeKey AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-documentattributecondition.html#cfn-kendra-datasource-documentattributecondition-conditiondocumentattributekey
	ConditionDocumentAttributeKey string `json:"ConditionDocumentAttributeKey"`

	// ConditionOnValue AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-documentattributecondition.html#cfn-kendra-datasource-documentattributecondition-conditiononvalue
	ConditionOnValue *DataSource_DocumentAttributeValue[any] `json:"ConditionOnValue,omitempty"`

	// Operator AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-documentattributecondition.html#cfn-kendra-datasource-documentattributecondition-operator
	Operator string `json:"Operator"`

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
func (r *DataSource_DocumentAttributeCondition[any]) AWSCloudFormationType() string {
	return "AWS::Kendra::DataSource.DocumentAttributeCondition"
}
