// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// Topic_TopicCalculatedField AWS CloudFormation Resource (AWS::QuickSight::Topic.TopicCalculatedField)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-topiccalculatedfield.html
type Topic_TopicCalculatedField[T any] struct {

	// Aggregation AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-topiccalculatedfield.html#cfn-quicksight-topic-topiccalculatedfield-aggregation
	Aggregation *string `json:"Aggregation,omitempty"`

	// AllowedAggregations AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-topiccalculatedfield.html#cfn-quicksight-topic-topiccalculatedfield-allowedaggregations
	AllowedAggregations []string `json:"AllowedAggregations,omitempty"`

	// CalculatedFieldDescription AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-topiccalculatedfield.html#cfn-quicksight-topic-topiccalculatedfield-calculatedfielddescription
	CalculatedFieldDescription *string `json:"CalculatedFieldDescription,omitempty"`

	// CalculatedFieldName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-topiccalculatedfield.html#cfn-quicksight-topic-topiccalculatedfield-calculatedfieldname
	CalculatedFieldName string `json:"CalculatedFieldName"`

	// CalculatedFieldSynonyms AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-topiccalculatedfield.html#cfn-quicksight-topic-topiccalculatedfield-calculatedfieldsynonyms
	CalculatedFieldSynonyms []string `json:"CalculatedFieldSynonyms,omitempty"`

	// CellValueSynonyms AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-topiccalculatedfield.html#cfn-quicksight-topic-topiccalculatedfield-cellvaluesynonyms
	CellValueSynonyms []Topic_CellValueSynonym[any] `json:"CellValueSynonyms,omitempty"`

	// ColumnDataRole AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-topiccalculatedfield.html#cfn-quicksight-topic-topiccalculatedfield-columndatarole
	ColumnDataRole *string `json:"ColumnDataRole,omitempty"`

	// ComparativeOrder AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-topiccalculatedfield.html#cfn-quicksight-topic-topiccalculatedfield-comparativeorder
	ComparativeOrder *Topic_ComparativeOrder[any] `json:"ComparativeOrder,omitempty"`

	// DefaultFormatting AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-topiccalculatedfield.html#cfn-quicksight-topic-topiccalculatedfield-defaultformatting
	DefaultFormatting *Topic_DefaultFormatting[any] `json:"DefaultFormatting,omitempty"`

	// Expression AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-topiccalculatedfield.html#cfn-quicksight-topic-topiccalculatedfield-expression
	Expression string `json:"Expression"`

	// IsIncludedInTopic AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-topiccalculatedfield.html#cfn-quicksight-topic-topiccalculatedfield-isincludedintopic
	IsIncludedInTopic *T `json:"IsIncludedInTopic,omitempty"`

	// NeverAggregateInFilter AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-topiccalculatedfield.html#cfn-quicksight-topic-topiccalculatedfield-neveraggregateinfilter
	NeverAggregateInFilter *T `json:"NeverAggregateInFilter,omitempty"`

	// NotAllowedAggregations AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-topiccalculatedfield.html#cfn-quicksight-topic-topiccalculatedfield-notallowedaggregations
	NotAllowedAggregations []string `json:"NotAllowedAggregations,omitempty"`

	// SemanticType AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-topiccalculatedfield.html#cfn-quicksight-topic-topiccalculatedfield-semantictype
	SemanticType *Topic_SemanticType[any] `json:"SemanticType,omitempty"`

	// TimeGranularity AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-topiccalculatedfield.html#cfn-quicksight-topic-topiccalculatedfield-timegranularity
	TimeGranularity *string `json:"TimeGranularity,omitempty"`

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
func (r *Topic_TopicCalculatedField[any]) AWSCloudFormationType() string {
	return "AWS::QuickSight::Topic.TopicCalculatedField"
}