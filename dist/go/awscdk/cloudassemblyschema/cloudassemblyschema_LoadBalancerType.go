package cloudassemblyschema


// Type of load balancer.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   loadBalancerContextQuery := &loadBalancerContextQuery{
//   	account: jsii.String("account"),
//   	loadBalancerType: awscdk.Cloud_assembly_schema.loadBalancerType_NETWORK,
//   	region: jsii.String("region"),
//
//   	// the properties below are optional
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
type LoadBalancerType string

const (
	// Network load balancer.
	LoadBalancerType_NETWORK LoadBalancerType = "NETWORK"
	// Application load balancer.
	LoadBalancerType_APPLICATION LoadBalancerType = "APPLICATION"
)

