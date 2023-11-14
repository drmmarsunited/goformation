// Code generated by "go generate". Please don't change this file directly.

package billingconductor

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// CustomLineItem_BillingPeriodRange AWS CloudFormation Resource (AWS::BillingConductor::CustomLineItem.BillingPeriodRange)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-billingconductor-customlineitem-billingperiodrange.html
type CustomLineItem_BillingPeriodRange[T any] struct {

	// ExclusiveEndBillingPeriod AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-billingconductor-customlineitem-billingperiodrange.html#cfn-billingconductor-customlineitem-billingperiodrange-exclusiveendbillingperiod
	ExclusiveEndBillingPeriod *T `json:"ExclusiveEndBillingPeriod,omitempty"`

	// InclusiveStartBillingPeriod AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-billingconductor-customlineitem-billingperiodrange.html#cfn-billingconductor-customlineitem-billingperiodrange-inclusivestartbillingperiod
	InclusiveStartBillingPeriod *T `json:"InclusiveStartBillingPeriod,omitempty"`

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
func (r *CustomLineItem_BillingPeriodRange[any]) AWSCloudFormationType() string {
	return "AWS::BillingConductor::CustomLineItem.BillingPeriodRange"
}
