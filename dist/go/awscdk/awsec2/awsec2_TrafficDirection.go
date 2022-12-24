package awsec2


// Direction of traffic the AclEntry applies to.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var aclCidr aclCidr
//   var aclTraffic aclTraffic
//   var networkAcl networkAcl
//
//   networkAclEntryProps := &networkAclEntryProps{
//   	cidr: aclCidr,
//   	networkAcl: networkAcl,
//   	ruleNumber: jsii.Number(123),
//   	traffic: aclTraffic,
//
//   	// the properties below are optional
//   	direction: awscdk.Aws_ec2.trafficDirection_EGRESS,
//   	networkAclEntryName: jsii.String("networkAclEntryName"),
//   	ruleAction: awscdk.*Aws_ec2.action_ALLOW,
//   }
//
type TrafficDirection string

const (
	// Traffic leaving the subnet.
	TrafficDirection_EGRESS TrafficDirection = "EGRESS"
	// Traffic entering the subnet.
	TrafficDirection_INGRESS TrafficDirection = "INGRESS"
)

