package awsautoscaling


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
type DefaultResult string

const (
	DefaultResult_CONTINUE DefaultResult = "CONTINUE"
	DefaultResult_ABANDON DefaultResult = "ABANDON"
)

