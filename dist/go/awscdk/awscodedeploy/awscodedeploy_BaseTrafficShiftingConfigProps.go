package awscodedeploy

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Common properties of traffic shifting routing configurations.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   baseTrafficShiftingConfigProps := &baseTrafficShiftingConfigProps{
//   	interval: cdk.duration.minutes(jsii.Number(30)),
//   	percentage: jsii.Number(123),
//   }
//
type BaseTrafficShiftingConfigProps struct {
	// The amount of time between traffic shifts.
	Interval awscdk.Duration `field:"required" json:"interval" yaml:"interval"`
	// The percentage to increase traffic on each traffic shift.
	Percentage *float64 `field:"required" json:"percentage" yaml:"percentage"`
}

