package awsservicediscovery


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var namespace iNamespace
//
//   serviceAttributes := &serviceAttributes{
//   	dnsRecordType: awscdk.Aws_servicediscovery.dnsRecordType_A,
//   	namespace: namespace,
//   	routingPolicy: awscdk.*Aws_servicediscovery.routingPolicy_WEIGHTED,
//   	serviceArn: jsii.String("serviceArn"),
//   	serviceId: jsii.String("serviceId"),
//   	serviceName: jsii.String("serviceName"),
//
//   	// the properties below are optional
//   	discoveryType: awscdk.*Aws_servicediscovery.discoveryType_API,
//   }
//
type RoutingPolicy string

const (
	// Route 53 returns the applicable value from one randomly selected instance from among the instances that you registered using the same service.
	RoutingPolicy_WEIGHTED RoutingPolicy = "WEIGHTED"
	// If you define a health check for the service and the health check is healthy, Route 53 returns the applicable value for up to eight instances.
	RoutingPolicy_MULTIVALUE RoutingPolicy = "MULTIVALUE"
)

