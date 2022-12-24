package awselasticloadbalancingv2


// Basic properties for a Network Listener.
//
// Example:
//   var vpc vpc
//   var asg autoScalingGroup
//
//
//   // Create the load balancer in a VPC. 'internetFacing' is 'false'
//   // by default, which creates an internal load balancer.
//   lb := elbv2.NewNetworkLoadBalancer(this, jsii.String("LB"), &networkLoadBalancerProps{
//   	vpc: vpc,
//   	internetFacing: jsii.Boolean(true),
//   })
//
//   // Add a listener on a particular port.
//   listener := lb.addListener(jsii.String("Listener"), &baseNetworkListenerProps{
//   	port: jsii.Number(443),
//   })
//
//   // Add targets on a particular port.
//   listener.addTargets(jsii.String("AppFleet"), &addNetworkTargetsProps{
//   	port: jsii.Number(443),
//   	targets: []iNetworkLoadBalancerTarget{
//   		asg,
//   	},
//   })
//
type BaseNetworkListenerProps struct {
	// The port on which the listener listens for requests.
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// Application-Layer Protocol Negotiation (ALPN) is a TLS extension that is sent on the initial TLS handshake hello messages.
	//
	// ALPN enables the application layer to negotiate which protocols should be used over a secure connection, such as HTTP/1 and HTTP/2.
	//
	// Can only be specified together with Protocol TLS.
	AlpnPolicy AlpnPolicy `field:"optional" json:"alpnPolicy" yaml:"alpnPolicy"`
	// Certificate list of ACM cert ARNs.
	//
	// You must provide exactly one certificate if the listener protocol is HTTPS or TLS.
	Certificates *[]IListenerCertificate `field:"optional" json:"certificates" yaml:"certificates"`
	// Default action to take for requests to this listener.
	//
	// This allows full control of the default Action of the load balancer,
	// including weighted forwarding. See the `NetworkListenerAction` class for
	// all options.
	//
	// Cannot be specified together with `defaultTargetGroups`.
	DefaultAction NetworkListenerAction `field:"optional" json:"defaultAction" yaml:"defaultAction"`
	// Default target groups to load balance to.
	//
	// All target groups will be load balanced to with equal weight and without
	// stickiness. For a more complex configuration than that, use
	// either `defaultAction` or `addAction()`.
	//
	// Cannot be specified together with `defaultAction`.
	DefaultTargetGroups *[]INetworkTargetGroup `field:"optional" json:"defaultTargetGroups" yaml:"defaultTargetGroups"`
	// Protocol for listener, expects TCP, TLS, UDP, or TCP_UDP.
	Protocol Protocol `field:"optional" json:"protocol" yaml:"protocol"`
	// SSL Policy.
	SslPolicy SslPolicy `field:"optional" json:"sslPolicy" yaml:"sslPolicy"`
}

