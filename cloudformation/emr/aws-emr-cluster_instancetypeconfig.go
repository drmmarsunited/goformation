// Code generated by "go generate". Please don't change this file directly.

package emr

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// Cluster_InstanceTypeConfig AWS CloudFormation Resource (AWS::EMR::Cluster.InstanceTypeConfig)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-instancetypeconfig.html
type Cluster_InstanceTypeConfig[T any] struct {

	// BidPrice AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-instancetypeconfig.html#cfn-elasticmapreduce-cluster-instancetypeconfig-bidprice
	BidPrice *T `json:"BidPrice,omitempty"`

	// BidPriceAsPercentageOfOnDemandPrice AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-instancetypeconfig.html#cfn-elasticmapreduce-cluster-instancetypeconfig-bidpriceaspercentageofondemandprice
	BidPriceAsPercentageOfOnDemandPrice *T `json:"BidPriceAsPercentageOfOnDemandPrice,omitempty"`

	// Configurations AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-instancetypeconfig.html#cfn-elasticmapreduce-cluster-instancetypeconfig-configurations
	Configurations []Cluster_Configuration[any] `json:"Configurations,omitempty"`

	// CustomAmiId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-instancetypeconfig.html#cfn-elasticmapreduce-cluster-instancetypeconfig-customamiid
	CustomAmiId *T `json:"CustomAmiId,omitempty"`

	// EbsConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-instancetypeconfig.html#cfn-elasticmapreduce-cluster-instancetypeconfig-ebsconfiguration
	EbsConfiguration *Cluster_EbsConfiguration[any] `json:"EbsConfiguration,omitempty"`

	// InstanceType AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-instancetypeconfig.html#cfn-elasticmapreduce-cluster-instancetypeconfig-instancetype
	InstanceType T `json:"InstanceType"`

	// WeightedCapacity AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-instancetypeconfig.html#cfn-elasticmapreduce-cluster-instancetypeconfig-weightedcapacity
	WeightedCapacity *T `json:"WeightedCapacity,omitempty"`

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
func (r *Cluster_InstanceTypeConfig[any]) AWSCloudFormationType() string {
	return "AWS::EMR::Cluster.InstanceTypeConfig"
}
