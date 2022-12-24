package cloudassemblyschema


// The protocol for connections from clients to the load balancer.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   loadBalancerListenerContextQuery := &loadBalancerListenerContextQuery{
//   	account: jsii.String("account"),
//   	loadBalancerType: awscdk.Cloud_assembly_schema.loadBalancerType_NETWORK,
//   	region: jsii.String("region"),
//
//   	// the properties below are optional
//   	listenerArn: jsii.String("listenerArn"),
//   	listenerPort: jsii.Number(123),
//   	listenerProtocol: awscdk.*Cloud_assembly_schema.loadBalancerListenerProtocol_HTTP,
//   	loadBalancerArn: jsii.String("loadBalancerArn"),
//   	loadBalancerTags: []tag{
//   		&tag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	lookupRoleArn: jsii.String("lookupRoleArn"),
//   }
//
type LoadBalancerListenerProtocol string

const (
	// HTTP protocol.
	LoadBalancerListenerProtocol_HTTP LoadBalancerListenerProtocol = "HTTP"
	// HTTPS protocol.
	LoadBalancerListenerProtocol_HTTPS LoadBalancerListenerProtocol = "HTTPS"
	// TCP protocol.
	LoadBalancerListenerProtocol_TCP LoadBalancerListenerProtocol = "TCP"
	// TLS protocol.
	LoadBalancerListenerProtocol_TLS LoadBalancerListenerProtocol = "TLS"
	// UDP protocol.
	LoadBalancerListenerProtocol_UDP LoadBalancerListenerProtocol = "UDP"
	// TCP and UDP protocol.
	LoadBalancerListenerProtocol_TCP_UDP LoadBalancerListenerProtocol = "TCP_UDP"
)

