package awsglobalaccelerator


// The protocol for the connections from clients to the accelerator.
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
type ConnectionProtocol string

const (
	// TCP.
	ConnectionProtocol_TCP ConnectionProtocol = "TCP"
	// UDP.
	ConnectionProtocol_UDP ConnectionProtocol = "UDP"
)

