// Code generated by "go generate". Please don't change this file directly.

package timestream

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// Table_MagneticStoreWriteProperties AWS CloudFormation Resource (AWS::Timestream::Table.MagneticStoreWriteProperties)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-timestream-table-magneticstorewriteproperties.html
type Table_MagneticStoreWriteProperties struct {

	// EnableMagneticStoreWrites AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-timestream-table-magneticstorewriteproperties.html#cfn-timestream-table-magneticstorewriteproperties-enablemagneticstorewrites
	EnableMagneticStoreWrites bool `json:"EnableMagneticStoreWrites"`

	// MagneticStoreRejectedDataLocation AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-timestream-table-magneticstorewriteproperties.html#cfn-timestream-table-magneticstorewriteproperties-magneticstorerejecteddatalocation
	MagneticStoreRejectedDataLocation *Table_MagneticStoreRejectedDataLocation `json:"MagneticStoreRejectedDataLocation,omitempty"`

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
func (r *Table_MagneticStoreWriteProperties) AWSCloudFormationType() string {
	return "AWS::Timestream::Table.MagneticStoreWriteProperties"
}
