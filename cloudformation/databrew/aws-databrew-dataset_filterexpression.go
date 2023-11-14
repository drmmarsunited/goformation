// Code generated by "go generate". Please don't change this file directly.

package databrew

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// Dataset_FilterExpression AWS CloudFormation Resource (AWS::DataBrew::Dataset.FilterExpression)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-dataset-filterexpression.html
type Dataset_FilterExpression[T any] struct {

	// Expression AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-dataset-filterexpression.html#cfn-databrew-dataset-filterexpression-expression
	Expression T `json:"Expression"`

	// ValuesMap AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-dataset-filterexpression.html#cfn-databrew-dataset-filterexpression-valuesmap
	ValuesMap []Dataset_FilterValue[any] `json:"ValuesMap"`

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
func (r *Dataset_FilterExpression[any]) AWSCloudFormationType() string {
	return "AWS::DataBrew::Dataset.FilterExpression"
}
