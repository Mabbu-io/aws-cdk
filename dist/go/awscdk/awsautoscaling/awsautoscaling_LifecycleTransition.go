package awsautoscaling


// What instance transition to attach the hook to.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var autoScalingGroup autoScalingGroup
//   var lifecycleHookTarget iLifecycleHookTarget
//   var role role
//
//   lifecycleHookProps := &lifecycleHookProps{
//   	autoScalingGroup: autoScalingGroup,
//   	lifecycleTransition: awscdk.Aws_autoscaling.lifecycleTransition_INSTANCE_LAUNCHING,
//
//   	// the properties below are optional
//   	defaultResult: awscdk.*Aws_autoscaling.defaultResult_CONTINUE,
//   	heartbeatTimeout: cdk.duration.minutes(jsii.Number(30)),
//   	lifecycleHookName: jsii.String("lifecycleHookName"),
//   	notificationMetadata: jsii.String("notificationMetadata"),
//   	notificationTarget: lifecycleHookTarget,
//   	role: role,
//   }
//
type LifecycleTransition string

const (
	// Execute the hook when an instance is about to be added.
	LifecycleTransition_INSTANCE_LAUNCHING LifecycleTransition = "INSTANCE_LAUNCHING"
	// Execute the hook when an instance is about to be terminated.
	LifecycleTransition_INSTANCE_TERMINATING LifecycleTransition = "INSTANCE_TERMINATING"
)

