package awselasticloadbalancingv2

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Options for `NetworkListenerAction.forward()`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   networkForwardOptions := &networkForwardOptions{
//   	stickinessDuration: cdk.duration.minutes(jsii.Number(30)),
//   }
//
type NetworkForwardOptions struct {
	// For how long clients should be directed to the same target group.
	//
	// Range between 1 second and 7 days.
	StickinessDuration awscdk.Duration `field:"optional" json:"stickinessDuration" yaml:"stickinessDuration"`
}

