package awseks

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Options for fetching a ServiceLoadBalancerAddress.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   serviceLoadBalancerAddressOptions := &serviceLoadBalancerAddressOptions{
//   	namespace: jsii.String("namespace"),
//   	timeout: cdk.duration.minutes(jsii.Number(30)),
//   }
//
type ServiceLoadBalancerAddressOptions struct {
	// The namespace the service belongs to.
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Timeout for waiting on the load balancer address.
	Timeout awscdk.Duration `field:"optional" json:"timeout" yaml:"timeout"`
}

