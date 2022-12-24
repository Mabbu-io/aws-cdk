package awselasticloadbalancingv2


// Application-Layer Protocol Negotiation Policies for network load balancers.
//
// Which protocols should be used over a secure connection.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var listenerCertificate listenerCertificate
//   var networkListenerAction networkListenerAction
//   var networkLoadBalancer networkLoadBalancer
//   var networkTargetGroup networkTargetGroup
//
//   networkListenerProps := &networkListenerProps{
//   	loadBalancer: networkLoadBalancer,
//   	port: jsii.Number(123),
//
//   	// the properties below are optional
//   	alpnPolicy: awscdk.Aws_elasticloadbalancingv2.alpnPolicy_HTTP1_ONLY,
//   	certificates: []iListenerCertificate{
//   		listenerCertificate,
//   	},
//   	defaultAction: networkListenerAction,
//   	defaultTargetGroups: []iNetworkTargetGroup{
//   		networkTargetGroup,
//   	},
//   	protocol: awscdk.*Aws_elasticloadbalancingv2.protocol_HTTP,
//   	sslPolicy: awscdk.*Aws_elasticloadbalancingv2.sslPolicy_RECOMMENDED_TLS,
//   }
//
type AlpnPolicy string

const (
	// Negotiate only HTTP/1.*. The ALPN preference list is http/1.1, http/1.0.
	AlpnPolicy_HTTP1_ONLY AlpnPolicy = "HTTP1_ONLY"
	// Negotiate only HTTP/2.
	//
	// The ALPN preference list is h2.
	AlpnPolicy_HTTP2_ONLY AlpnPolicy = "HTTP2_ONLY"
	// Prefer HTTP/1.* over HTTP/2 (which can be useful for HTTP/2 testing). The ALPN preference list is http/1.1, http/1.0, h2.
	AlpnPolicy_HTTP2_OPTIONAL AlpnPolicy = "HTTP2_OPTIONAL"
	// Prefer HTTP/2 over HTTP/1.*. The ALPN preference list is h2, http/1.1, http/1.0.
	AlpnPolicy_HTTP2_PREFERRED AlpnPolicy = "HTTP2_PREFERRED"
	// Do not negotiate ALPN.
	AlpnPolicy_NONE AlpnPolicy = "NONE"
)

