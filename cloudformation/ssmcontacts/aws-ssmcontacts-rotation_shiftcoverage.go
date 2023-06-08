// Code generated by "go generate". Please don't change this file directly.

package ssmcontacts

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// Rotation_ShiftCoverage AWS CloudFormation Resource (AWS::SSMContacts::Rotation.ShiftCoverage)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssmcontacts-rotation-shiftcoverage.html
type Rotation_ShiftCoverage[T any] struct {

	// CoverageTimes AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssmcontacts-rotation-shiftcoverage.html#cfn-ssmcontacts-rotation-shiftcoverage-coveragetimes
	CoverageTimes []Rotation_CoverageTime[any] `json:"CoverageTimes"`

	// DayOfWeek AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssmcontacts-rotation-shiftcoverage.html#cfn-ssmcontacts-rotation-shiftcoverage-dayofweek
	DayOfWeek string `json:"DayOfWeek"`

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
func (r *Rotation_ShiftCoverage[any]) AWSCloudFormationType() string {
	return "AWS::SSMContacts::Rotation.ShiftCoverage"
}
