package awselasticloadbalancingv2

import (
	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsec2"
)

// Properties for a new Network Target Group.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var networkLoadBalancerTarget iNetworkLoadBalancerTarget
//   var vpc vpc
//
//   networkTargetGroupProps := &networkTargetGroupProps{
//   	port: jsii.Number(123),
//
//   	// the properties below are optional
//   	connectionTermination: jsii.Boolean(false),
//   	deregistrationDelay: cdk.duration.minutes(jsii.Number(30)),
//   	healthCheck: &healthCheck{
//   		enabled: jsii.Boolean(false),
//   		healthyGrpcCodes: jsii.String("healthyGrpcCodes"),
//   		healthyHttpCodes: jsii.String("healthyHttpCodes"),
//   		healthyThresholdCount: jsii.Number(123),
//   		interval: cdk.*duration.minutes(jsii.Number(30)),
//   		path: jsii.String("path"),
//   		port: jsii.String("port"),
//   		protocol: awscdk.Aws_elasticloadbalancingv2.protocol_HTTP,
//   		timeout: cdk.*duration.minutes(jsii.Number(30)),
//   		unhealthyThresholdCount: jsii.Number(123),
//   	},
//   	preserveClientIp: jsii.Boolean(false),
//   	protocol: awscdk.*Aws_elasticloadbalancingv2.*protocol_HTTP,
//   	proxyProtocolV2: jsii.Boolean(false),
//   	targetGroupName: jsii.String("targetGroupName"),
//   	targets: []*iNetworkLoadBalancerTarget{
//   		networkLoadBalancerTarget,
//   	},
//   	targetType: awscdk.*Aws_elasticloadbalancingv2.targetType_INSTANCE,
//   	vpc: vpc,
//   }
//
type NetworkTargetGroupProps struct {
	// The amount of time for Elastic Load Balancing to wait before deregistering a target.
	//
	// The range is 0-3600 seconds.
	DeregistrationDelay awscdk.Duration `field:"optional" json:"deregistrationDelay" yaml:"deregistrationDelay"`
	// Health check configuration.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-targetgroup.html#aws-resource-elasticloadbalancingv2-targetgroup-properties
	//
	HealthCheck *HealthCheck `field:"optional" json:"healthCheck" yaml:"healthCheck"`
	// The name of the target group.
	//
	// This name must be unique per region per account, can have a maximum of
	// 32 characters, must contain only alphanumeric characters or hyphens, and
	// must not begin or end with a hyphen.
	TargetGroupName *string `field:"optional" json:"targetGroupName" yaml:"targetGroupName"`
	// The type of targets registered to this TargetGroup, either IP or Instance.
	//
	// All targets registered into the group must be of this type. If you
	// register targets to the TargetGroup in the CDK app, the TargetType is
	// determined automatically.
	TargetType TargetType `field:"optional" json:"targetType" yaml:"targetType"`
	// The virtual private cloud (VPC).
	//
	// only if `TargetType` is `Ip` or `InstanceId`.
	Vpc awsec2.IVpc `field:"optional" json:"vpc" yaml:"vpc"`
	// The port on which the listener listens for requests.
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// Indicates whether the load balancer terminates connections at the end of the deregistration timeout.
	ConnectionTermination *bool `field:"optional" json:"connectionTermination" yaml:"connectionTermination"`
	// Indicates whether client IP preservation is enabled.
	PreserveClientIp *bool `field:"optional" json:"preserveClientIp" yaml:"preserveClientIp"`
	// Protocol for target group, expects TCP, TLS, UDP, or TCP_UDP.
	Protocol Protocol `field:"optional" json:"protocol" yaml:"protocol"`
	// Indicates whether Proxy Protocol version 2 is enabled.
	ProxyProtocolV2 *bool `field:"optional" json:"proxyProtocolV2" yaml:"proxyProtocolV2"`
	// The targets to add to this target group.
	//
	// Can be `Instance`, `IPAddress`, or any self-registering load balancing
	// target. If you use either `Instance` or `IPAddress` as targets, all
	// target must be of the same type.
	Targets *[]INetworkLoadBalancerTarget `field:"optional" json:"targets" yaml:"targets"`
}

