package awsglobalaccelerator


// Client affinity gives you control over whether to always route each client to the same specific endpoint.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var accelerator accelerator
//
//   listenerProps := &listenerProps{
//   	accelerator: accelerator,
//   	portRanges: []portRange{
//   		&portRange{
//   			fromPort: jsii.Number(123),
//
//   			// the properties below are optional
//   			toPort: jsii.Number(123),
//   		},
//   	},
//
//   	// the properties below are optional
//   	clientAffinity: awscdk.Aws_globalaccelerator.clientAffinity_NONE,
//   	listenerName: jsii.String("listenerName"),
//   	protocol: awscdk.*Aws_globalaccelerator.connectionProtocol_TCP,
//   }
//
// See: https://docs.aws.amazon.com/global-accelerator/latest/dg/about-listeners.html#about-listeners-client-affinity
//
type ClientAffinity string

const (
	// Route traffic based on the 5-tuple `(source IP, source port, destination IP, destination port, protocol)`.
	ClientAffinity_NONE ClientAffinity = "NONE"
	// Route traffic based on the 2-tuple `(source IP, destination IP)`.
	//
	// The result is that multiple connections from the same client will be routed the same.
	ClientAffinity_SOURCE_IP ClientAffinity = "SOURCE_IP"
)

