// Code generated by "go generate". Please don't change this file directly.

package timestream

import (
	"bytes"
	"encoding/json"

	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
	"github.com/drmmarsunited/goformation/v7/cloudformation/tags"
)

// ScheduledQuery AWS CloudFormation Resource (AWS::Timestream::ScheduledQuery)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-timestream-scheduledquery.html
type ScheduledQuery[T any] struct {

	// ClientToken AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-timestream-scheduledquery.html#cfn-timestream-scheduledquery-clienttoken
	ClientToken *string `json:"ClientToken,omitempty"`

	// ErrorReportConfiguration AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-timestream-scheduledquery.html#cfn-timestream-scheduledquery-errorreportconfiguration
	ErrorReportConfiguration *ScheduledQuery_ErrorReportConfiguration[any] `json:"ErrorReportConfiguration"`

	// KmsKeyId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-timestream-scheduledquery.html#cfn-timestream-scheduledquery-kmskeyid
	KmsKeyId *string `json:"KmsKeyId,omitempty"`

	// NotificationConfiguration AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-timestream-scheduledquery.html#cfn-timestream-scheduledquery-notificationconfiguration
	NotificationConfiguration *ScheduledQuery_NotificationConfiguration[any] `json:"NotificationConfiguration"`

	// QueryString AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-timestream-scheduledquery.html#cfn-timestream-scheduledquery-querystring
	QueryString string `json:"QueryString"`

	// ScheduleConfiguration AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-timestream-scheduledquery.html#cfn-timestream-scheduledquery-scheduleconfiguration
	ScheduleConfiguration *ScheduledQuery_ScheduleConfiguration[any] `json:"ScheduleConfiguration"`

	// ScheduledQueryExecutionRoleArn AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-timestream-scheduledquery.html#cfn-timestream-scheduledquery-scheduledqueryexecutionrolearn
	ScheduledQueryExecutionRoleArn string `json:"ScheduledQueryExecutionRoleArn"`

	// ScheduledQueryName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-timestream-scheduledquery.html#cfn-timestream-scheduledquery-scheduledqueryname
	ScheduledQueryName *string `json:"ScheduledQueryName,omitempty"`

	// Tags AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-timestream-scheduledquery.html#cfn-timestream-scheduledquery-tags
	Tags []tags.Tag `json:"Tags,omitempty"`

	// TargetConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-timestream-scheduledquery.html#cfn-timestream-scheduledquery-targetconfiguration
	TargetConfiguration *ScheduledQuery_TargetConfiguration[any] `json:"TargetConfiguration,omitempty"`

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
func (r *ScheduledQuery[any]) AWSCloudFormationType() string {
	return "AWS::Timestream::ScheduledQuery"
}

// MarshalJSON is a custom JSON marshalling hook that embeds this object into
// an AWS CloudFormation JSON resource's 'Properties' field and adds a 'Type'.
func (r ScheduledQuery[any]) MarshalJSON() ([]byte, error) {
	type Properties ScheduledQuery[any]

	return json.Marshal(&struct {
		Type                string
		Properties          Properties
		DependsOn           []string                     `json:"DependsOn,omitempty"`
		Metadata            map[string]interface{}       `json:"Metadata,omitempty"`
		DeletionPolicy      policies.DeletionPolicy      `json:"DeletionPolicy,omitempty"`
		UpdateReplacePolicy policies.UpdateReplacePolicy `json:"UpdateReplacePolicy,omitempty"`
		Condition           string                       `json:"Condition,omitempty"`
	}{
		Type:                r.AWSCloudFormationType(),
		Properties:          (Properties)(r),
		DependsOn:           r.AWSCloudFormationDependsOn,
		Metadata:            r.AWSCloudFormationMetadata,
		DeletionPolicy:      r.AWSCloudFormationDeletionPolicy,
		UpdateReplacePolicy: r.AWSCloudFormationUpdateReplacePolicy,
		Condition:           r.AWSCloudFormationCondition,
	})
}

// UnmarshalJSON is a custom JSON unmarshalling hook that strips the outer
// AWS CloudFormation resource object, and just keeps the 'Properties' field.
func (r *ScheduledQuery[any]) UnmarshalJSON(b []byte) error {
	type Properties ScheduledQuery[any]

	res := &struct {
		Type                string
		Properties          *Properties
		DependsOn           interface{}
		Metadata            map[string]interface{}
		DeletionPolicy      string
		UpdateReplacePolicy string
		Condition           string
	}{}

	dec := json.NewDecoder(bytes.NewReader(b))
	dec.DisallowUnknownFields() // Force error if unknown field is found

	if err := dec.Decode(&res); err != nil {
		return err
	}

	// If the resource has no Properties set, it could be nil
	if res.Properties != nil {
		*r = ScheduledQuery[any](*res.Properties)
	}
	if res.DependsOn != nil {
		switch obj := res.DependsOn.(type) {
		case string:
			r.AWSCloudFormationDependsOn = []string{obj}
		case []interface{}:
			s := make([]string, 0, len(obj))
			for _, v := range obj {
				if value, ok := v.(string); ok {
					s = append(s, value)
				}
			}
			r.AWSCloudFormationDependsOn = s
		}
	}
	if res.Metadata != nil {
		r.AWSCloudFormationMetadata = res.Metadata
	}
	if res.DeletionPolicy != "" {
		r.AWSCloudFormationDeletionPolicy = policies.DeletionPolicy(res.DeletionPolicy)
	}
	if res.UpdateReplacePolicy != "" {
		r.AWSCloudFormationUpdateReplacePolicy = policies.UpdateReplacePolicy(res.UpdateReplacePolicy)
	}
	if res.Condition != "" {
		r.AWSCloudFormationCondition = res.Condition
	}
	return nil
}
