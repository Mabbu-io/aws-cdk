package awsautoscaling


// The instance state in the warm pool.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var autoScalingGroup autoScalingGroup
//
//   warmPool := awscdk.Aws_autoscaling.NewWarmPool(this, jsii.String("MyWarmPool"), &warmPoolProps{
//   	autoScalingGroup: autoScalingGroup,
//
//   	// the properties below are optional
//   	maxGroupPreparedCapacity: jsii.Number(123),
//   	minSize: jsii.Number(123),
//   	poolState: awscdk.*Aws_autoscaling.poolState_HIBERNATED,
//   	reuseOnScaleIn: jsii.Boolean(false),
//   })
//
type PoolState string

const (
	// Hibernated.
	//
	// To use this state, prerequisites must be in place.
	// See: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/hibernating-prerequisites.html
	//
	PoolState_HIBERNATED PoolState = "HIBERNATED"
	// Running.
	PoolState_RUNNING PoolState = "RUNNING"
	// Stopped.
	PoolState_STOPPED PoolState = "STOPPED"
)

