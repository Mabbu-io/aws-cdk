package awsglobalaccelerator


// The protocol for the connections from clients to the accelerator.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var endpoint iEndpoint
//   var listener listener
//
//   endpointGroupProps := &endpointGroupProps{
//   	listener: listener,
//
//   	// the properties below are optional
//   	endpointGroupName: jsii.String("endpointGroupName"),
//   	endpoints: []*iEndpoint{
//   		endpoint,
//   	},
//   	healthCheckInterval: cdk.duration.minutes(jsii.Number(30)),
//   	healthCheckPath: jsii.String("healthCheckPath"),
//   	healthCheckPort: jsii.Number(123),
//   	healthCheckProtocol: awscdk.Aws_globalaccelerator.healthCheckProtocol_TCP,
//   	healthCheckThreshold: jsii.Number(123),
//   	portOverrides: []portOverride{
//   		&portOverride{
//   			endpointPort: jsii.Number(123),
//   			listenerPort: jsii.Number(123),
//   		},
//   	},
//   	region: jsii.String("region"),
//   	trafficDialPercentage: jsii.Number(123),
//   }
//
type HealthCheckProtocol string

const (
	// TCP.
	HealthCheckProtocol_TCP HealthCheckProtocol = "TCP"
	// HTTP.
	HealthCheckProtocol_HTTP HealthCheckProtocol = "HTTP"
	// HTTPS.
	HealthCheckProtocol_HTTPS HealthCheckProtocol = "HTTPS"
)

