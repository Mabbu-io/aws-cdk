package cxapi


// Load balancer ip address type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   loadBalancerContextResponse := &loadBalancerContextResponse{
//   	ipAddressType: awscdk.Cx_api.loadBalancerIpAddressType_IPV4,
//   	loadBalancerArn: jsii.String("loadBalancerArn"),
//   	loadBalancerCanonicalHostedZoneId: jsii.String("loadBalancerCanonicalHostedZoneId"),
//   	loadBalancerDnsName: jsii.String("loadBalancerDnsName"),
//   	securityGroupIds: []*string{
//   		jsii.String("securityGroupIds"),
//   	},
//   	vpcId: jsii.String("vpcId"),
//   }
//
type LoadBalancerIpAddressType string

const (
	// IPV4 ip address.
	LoadBalancerIpAddressType_IPV4 LoadBalancerIpAddressType = "IPV4"
	// Dual stack address.
	LoadBalancerIpAddressType_DUAL_STACK LoadBalancerIpAddressType = "DUAL_STACK"
)

