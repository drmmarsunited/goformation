// Code generated by "go generate". Please don't change this file directly.

package ec2

import (
	"github.com/drmmarsunited/goformation/v7/cloudformation/policies"
)

// NetworkInsightsAnalysis_Explanation AWS CloudFormation Resource (AWS::EC2::NetworkInsightsAnalysis.Explanation)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsanalysis-explanation.html
type NetworkInsightsAnalysis_Explanation[T any] struct {

	// Acl AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsanalysis-explanation.html#cfn-ec2-networkinsightsanalysis-explanation-acl
	Acl *NetworkInsightsAnalysis_AnalysisComponent[any] `json:"Acl,omitempty"`

	// AclRule AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsanalysis-explanation.html#cfn-ec2-networkinsightsanalysis-explanation-aclrule
	AclRule *NetworkInsightsAnalysis_AnalysisAclRule[any] `json:"AclRule,omitempty"`

	// Address AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsanalysis-explanation.html#cfn-ec2-networkinsightsanalysis-explanation-address
	Address *T `json:"Address,omitempty"`

	// Addresses AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsanalysis-explanation.html#cfn-ec2-networkinsightsanalysis-explanation-addresses
	Addresses []T `json:"Addresses,omitempty"`

	// AttachedTo AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsanalysis-explanation.html#cfn-ec2-networkinsightsanalysis-explanation-attachedto
	AttachedTo *NetworkInsightsAnalysis_AnalysisComponent[any] `json:"AttachedTo,omitempty"`

	// AvailabilityZones AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsanalysis-explanation.html#cfn-ec2-networkinsightsanalysis-explanation-availabilityzones
	AvailabilityZones []T `json:"AvailabilityZones,omitempty"`

	// Cidrs AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsanalysis-explanation.html#cfn-ec2-networkinsightsanalysis-explanation-cidrs
	Cidrs []T `json:"Cidrs,omitempty"`

	// ClassicLoadBalancerListener AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsanalysis-explanation.html#cfn-ec2-networkinsightsanalysis-explanation-classicloadbalancerlistener
	ClassicLoadBalancerListener *NetworkInsightsAnalysis_AnalysisLoadBalancerListener[any] `json:"ClassicLoadBalancerListener,omitempty"`

	// Component AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsanalysis-explanation.html#cfn-ec2-networkinsightsanalysis-explanation-component
	Component *NetworkInsightsAnalysis_AnalysisComponent[any] `json:"Component,omitempty"`

	// ComponentAccount AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsanalysis-explanation.html#cfn-ec2-networkinsightsanalysis-explanation-componentaccount
	ComponentAccount *T `json:"ComponentAccount,omitempty"`

	// ComponentRegion AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsanalysis-explanation.html#cfn-ec2-networkinsightsanalysis-explanation-componentregion
	ComponentRegion *T `json:"ComponentRegion,omitempty"`

	// CustomerGateway AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsanalysis-explanation.html#cfn-ec2-networkinsightsanalysis-explanation-customergateway
	CustomerGateway *NetworkInsightsAnalysis_AnalysisComponent[any] `json:"CustomerGateway,omitempty"`

	// Destination AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsanalysis-explanation.html#cfn-ec2-networkinsightsanalysis-explanation-destination
	Destination *NetworkInsightsAnalysis_AnalysisComponent[any] `json:"Destination,omitempty"`

	// DestinationVpc AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsanalysis-explanation.html#cfn-ec2-networkinsightsanalysis-explanation-destinationvpc
	DestinationVpc *NetworkInsightsAnalysis_AnalysisComponent[any] `json:"DestinationVpc,omitempty"`

	// Direction AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsanalysis-explanation.html#cfn-ec2-networkinsightsanalysis-explanation-direction
	Direction *T `json:"Direction,omitempty"`

	// ElasticLoadBalancerListener AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsanalysis-explanation.html#cfn-ec2-networkinsightsanalysis-explanation-elasticloadbalancerlistener
	ElasticLoadBalancerListener *NetworkInsightsAnalysis_AnalysisComponent[any] `json:"ElasticLoadBalancerListener,omitempty"`

	// ExplanationCode AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsanalysis-explanation.html#cfn-ec2-networkinsightsanalysis-explanation-explanationcode
	ExplanationCode *T `json:"ExplanationCode,omitempty"`

	// IngressRouteTable AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsanalysis-explanation.html#cfn-ec2-networkinsightsanalysis-explanation-ingressroutetable
	IngressRouteTable *NetworkInsightsAnalysis_AnalysisComponent[any] `json:"IngressRouteTable,omitempty"`

	// InternetGateway AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsanalysis-explanation.html#cfn-ec2-networkinsightsanalysis-explanation-internetgateway
	InternetGateway *NetworkInsightsAnalysis_AnalysisComponent[any] `json:"InternetGateway,omitempty"`

	// LoadBalancerArn AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsanalysis-explanation.html#cfn-ec2-networkinsightsanalysis-explanation-loadbalancerarn
	LoadBalancerArn *T `json:"LoadBalancerArn,omitempty"`

	// LoadBalancerListenerPort AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsanalysis-explanation.html#cfn-ec2-networkinsightsanalysis-explanation-loadbalancerlistenerport
	LoadBalancerListenerPort *T `json:"LoadBalancerListenerPort,omitempty"`

	// LoadBalancerTarget AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsanalysis-explanation.html#cfn-ec2-networkinsightsanalysis-explanation-loadbalancertarget
	LoadBalancerTarget *NetworkInsightsAnalysis_AnalysisLoadBalancerTarget[any] `json:"LoadBalancerTarget,omitempty"`

	// LoadBalancerTargetGroup AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsanalysis-explanation.html#cfn-ec2-networkinsightsanalysis-explanation-loadbalancertargetgroup
	LoadBalancerTargetGroup *NetworkInsightsAnalysis_AnalysisComponent[any] `json:"LoadBalancerTargetGroup,omitempty"`

	// LoadBalancerTargetGroups AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsanalysis-explanation.html#cfn-ec2-networkinsightsanalysis-explanation-loadbalancertargetgroups
	LoadBalancerTargetGroups []NetworkInsightsAnalysis_AnalysisComponent[any] `json:"LoadBalancerTargetGroups,omitempty"`

	// LoadBalancerTargetPort AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsanalysis-explanation.html#cfn-ec2-networkinsightsanalysis-explanation-loadbalancertargetport
	LoadBalancerTargetPort *T `json:"LoadBalancerTargetPort,omitempty"`

	// MissingComponent AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsanalysis-explanation.html#cfn-ec2-networkinsightsanalysis-explanation-missingcomponent
	MissingComponent *T `json:"MissingComponent,omitempty"`

	// NatGateway AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsanalysis-explanation.html#cfn-ec2-networkinsightsanalysis-explanation-natgateway
	NatGateway *NetworkInsightsAnalysis_AnalysisComponent[any] `json:"NatGateway,omitempty"`

	// NetworkInterface AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsanalysis-explanation.html#cfn-ec2-networkinsightsanalysis-explanation-networkinterface
	NetworkInterface *NetworkInsightsAnalysis_AnalysisComponent[any] `json:"NetworkInterface,omitempty"`

	// PacketField AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsanalysis-explanation.html#cfn-ec2-networkinsightsanalysis-explanation-packetfield
	PacketField *T `json:"PacketField,omitempty"`

	// Port AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsanalysis-explanation.html#cfn-ec2-networkinsightsanalysis-explanation-port
	Port *T `json:"Port,omitempty"`

	// PortRanges AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsanalysis-explanation.html#cfn-ec2-networkinsightsanalysis-explanation-portranges
	PortRanges []NetworkInsightsAnalysis_PortRange[any] `json:"PortRanges,omitempty"`

	// PrefixList AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsanalysis-explanation.html#cfn-ec2-networkinsightsanalysis-explanation-prefixlist
	PrefixList *NetworkInsightsAnalysis_AnalysisComponent[any] `json:"PrefixList,omitempty"`

	// Protocols AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsanalysis-explanation.html#cfn-ec2-networkinsightsanalysis-explanation-protocols
	Protocols []T `json:"Protocols,omitempty"`

	// RouteTable AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsanalysis-explanation.html#cfn-ec2-networkinsightsanalysis-explanation-routetable
	RouteTable *NetworkInsightsAnalysis_AnalysisComponent[any] `json:"RouteTable,omitempty"`

	// RouteTableRoute AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsanalysis-explanation.html#cfn-ec2-networkinsightsanalysis-explanation-routetableroute
	RouteTableRoute *NetworkInsightsAnalysis_AnalysisRouteTableRoute[any] `json:"RouteTableRoute,omitempty"`

	// SecurityGroup AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsanalysis-explanation.html#cfn-ec2-networkinsightsanalysis-explanation-securitygroup
	SecurityGroup *NetworkInsightsAnalysis_AnalysisComponent[any] `json:"SecurityGroup,omitempty"`

	// SecurityGroupRule AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsanalysis-explanation.html#cfn-ec2-networkinsightsanalysis-explanation-securitygrouprule
	SecurityGroupRule *NetworkInsightsAnalysis_AnalysisSecurityGroupRule[any] `json:"SecurityGroupRule,omitempty"`

	// SecurityGroups AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsanalysis-explanation.html#cfn-ec2-networkinsightsanalysis-explanation-securitygroups
	SecurityGroups []NetworkInsightsAnalysis_AnalysisComponent[any] `json:"SecurityGroups,omitempty"`

	// SourceVpc AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsanalysis-explanation.html#cfn-ec2-networkinsightsanalysis-explanation-sourcevpc
	SourceVpc *NetworkInsightsAnalysis_AnalysisComponent[any] `json:"SourceVpc,omitempty"`

	// State AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsanalysis-explanation.html#cfn-ec2-networkinsightsanalysis-explanation-state
	State *T `json:"State,omitempty"`

	// Subnet AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsanalysis-explanation.html#cfn-ec2-networkinsightsanalysis-explanation-subnet
	Subnet *NetworkInsightsAnalysis_AnalysisComponent[any] `json:"Subnet,omitempty"`

	// SubnetRouteTable AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsanalysis-explanation.html#cfn-ec2-networkinsightsanalysis-explanation-subnetroutetable
	SubnetRouteTable *NetworkInsightsAnalysis_AnalysisComponent[any] `json:"SubnetRouteTable,omitempty"`

	// TransitGateway AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsanalysis-explanation.html#cfn-ec2-networkinsightsanalysis-explanation-transitgateway
	TransitGateway *NetworkInsightsAnalysis_AnalysisComponent[any] `json:"TransitGateway,omitempty"`

	// TransitGatewayAttachment AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsanalysis-explanation.html#cfn-ec2-networkinsightsanalysis-explanation-transitgatewayattachment
	TransitGatewayAttachment *NetworkInsightsAnalysis_AnalysisComponent[any] `json:"TransitGatewayAttachment,omitempty"`

	// TransitGatewayRouteTable AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsanalysis-explanation.html#cfn-ec2-networkinsightsanalysis-explanation-transitgatewayroutetable
	TransitGatewayRouteTable *NetworkInsightsAnalysis_AnalysisComponent[any] `json:"TransitGatewayRouteTable,omitempty"`

	// TransitGatewayRouteTableRoute AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsanalysis-explanation.html#cfn-ec2-networkinsightsanalysis-explanation-transitgatewayroutetableroute
	TransitGatewayRouteTableRoute *NetworkInsightsAnalysis_TransitGatewayRouteTableRoute[any] `json:"TransitGatewayRouteTableRoute,omitempty"`

	// Vpc AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsanalysis-explanation.html#cfn-ec2-networkinsightsanalysis-explanation-vpc
	Vpc *NetworkInsightsAnalysis_AnalysisComponent[any] `json:"Vpc,omitempty"`

	// VpcPeeringConnection AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsanalysis-explanation.html#cfn-ec2-networkinsightsanalysis-explanation-vpcpeeringconnection
	VpcPeeringConnection *NetworkInsightsAnalysis_AnalysisComponent[any] `json:"VpcPeeringConnection,omitempty"`

	// VpnConnection AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsanalysis-explanation.html#cfn-ec2-networkinsightsanalysis-explanation-vpnconnection
	VpnConnection *NetworkInsightsAnalysis_AnalysisComponent[any] `json:"VpnConnection,omitempty"`

	// VpnGateway AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsanalysis-explanation.html#cfn-ec2-networkinsightsanalysis-explanation-vpngateway
	VpnGateway *NetworkInsightsAnalysis_AnalysisComponent[any] `json:"VpnGateway,omitempty"`

	// vpcEndpoint AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsanalysis-explanation.html#cfn-ec2-networkinsightsanalysis-explanation-vpcendpoint
	vpcEndpoint *NetworkInsightsAnalysis_AnalysisComponent[any] `json:"vpcEndpoint,omitempty"`

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
func (r *NetworkInsightsAnalysis_Explanation[any]) AWSCloudFormationType() string {
	return "AWS::EC2::NetworkInsightsAnalysis.Explanation"
}
