// Code generated by "go generate". Please don't change this file directly.

package connect

import (
	"bytes"
	"encoding/json"

	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// InstanceStorageConfig AWS CloudFormation Resource (AWS::Connect::InstanceStorageConfig)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-instancestorageconfig.html
type InstanceStorageConfig[T any] struct {

	// InstanceArn AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-instancestorageconfig.html#cfn-connect-instancestorageconfig-instancearn
	InstanceArn string `json:"InstanceArn"`

	// KinesisFirehoseConfig AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-instancestorageconfig.html#cfn-connect-instancestorageconfig-kinesisfirehoseconfig
	KinesisFirehoseConfig *InstanceStorageConfig_KinesisFirehoseConfig[any] `json:"KinesisFirehoseConfig,omitempty"`

	// KinesisStreamConfig AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-instancestorageconfig.html#cfn-connect-instancestorageconfig-kinesisstreamconfig
	KinesisStreamConfig *InstanceStorageConfig_KinesisStreamConfig[any] `json:"KinesisStreamConfig,omitempty"`

	// KinesisVideoStreamConfig AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-instancestorageconfig.html#cfn-connect-instancestorageconfig-kinesisvideostreamconfig
	KinesisVideoStreamConfig *InstanceStorageConfig_KinesisVideoStreamConfig[any] `json:"KinesisVideoStreamConfig,omitempty"`

	// ResourceType AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-instancestorageconfig.html#cfn-connect-instancestorageconfig-resourcetype
	ResourceType string `json:"ResourceType"`

	// S3Config AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-instancestorageconfig.html#cfn-connect-instancestorageconfig-s3config
	S3Config *InstanceStorageConfig_S3Config[any] `json:"S3Config,omitempty"`

	// StorageType AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-instancestorageconfig.html#cfn-connect-instancestorageconfig-storagetype
	StorageType string `json:"StorageType"`

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
func (r *InstanceStorageConfig[any]) AWSCloudFormationType() string {
	return "AWS::Connect::InstanceStorageConfig"
}

// MarshalJSON is a custom JSON marshalling hook that embeds this object into
// an AWS CloudFormation JSON resource's 'Properties' field and adds a 'Type'.
func (r InstanceStorageConfig[any]) MarshalJSON() ([]byte, error) {
	type Properties InstanceStorageConfig[any]

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
func (r *InstanceStorageConfig[any]) UnmarshalJSON(b []byte) error {
	type Properties InstanceStorageConfig[any]

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
		*r = InstanceStorageConfig[any](*res.Properties)
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
