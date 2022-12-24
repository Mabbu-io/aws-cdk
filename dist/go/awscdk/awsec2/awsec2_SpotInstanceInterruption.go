package awsec2


// Provides the options for the types of interruption for spot instances.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var expiration expiration
//
//   launchTemplateSpotOptions := &launchTemplateSpotOptions{
//   	blockDuration: cdk.duration.minutes(jsii.Number(30)),
//   	interruptionBehavior: awscdk.Aws_ec2.spotInstanceInterruption_STOP,
//   	maxPrice: jsii.Number(123),
//   	requestType: awscdk.*Aws_ec2.spotRequestType_ONE_TIME,
//   	validUntil: expiration,
//   }
//
type SpotInstanceInterruption string

const (
	// The instance will stop when interrupted.
	SpotInstanceInterruption_STOP SpotInstanceInterruption = "STOP"
	// The instance will be terminated when interrupted.
	SpotInstanceInterruption_TERMINATE SpotInstanceInterruption = "TERMINATE"
	// The instance will hibernate when interrupted.
	SpotInstanceInterruption_HIBERNATE SpotInstanceInterruption = "HIBERNATE"
)

