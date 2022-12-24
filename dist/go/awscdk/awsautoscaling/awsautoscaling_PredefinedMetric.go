package awsautoscaling


// One of the predefined autoscaling metrics.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var autoScalingGroup autoScalingGroup
//   var metric metric
//
//   targetTrackingScalingPolicyProps := &targetTrackingScalingPolicyProps{
//   	autoScalingGroup: autoScalingGroup,
//   	targetValue: jsii.Number(123),
//
//   	// the properties below are optional
//   	cooldown: cdk.duration.minutes(jsii.Number(30)),
//   	customMetric: metric,
//   	disableScaleIn: jsii.Boolean(false),
//   	estimatedInstanceWarmup: cdk.*duration.minutes(jsii.Number(30)),
//   	predefinedMetric: awscdk.Aws_autoscaling.predefinedMetric_ASG_AVERAGE_CPU_UTILIZATION,
//   	resourceLabel: jsii.String("resourceLabel"),
//   }
//
type PredefinedMetric string

const (
	// Average CPU utilization of the Auto Scaling group.
	PredefinedMetric_ASG_AVERAGE_CPU_UTILIZATION PredefinedMetric = "ASG_AVERAGE_CPU_UTILIZATION"
	// Average number of bytes received on all network interfaces by the Auto Scaling group.
	PredefinedMetric_ASG_AVERAGE_NETWORK_IN PredefinedMetric = "ASG_AVERAGE_NETWORK_IN"
	// Average number of bytes sent out on all network interfaces by the Auto Scaling group.
	PredefinedMetric_ASG_AVERAGE_NETWORK_OUT PredefinedMetric = "ASG_AVERAGE_NETWORK_OUT"
	// Number of requests completed per target in an Application Load Balancer target group.
	//
	// Specify the ALB to look at in the `resourceLabel` field.
	PredefinedMetric_ALB_REQUEST_COUNT_PER_TARGET PredefinedMetric = "ALB_REQUEST_COUNT_PER_TARGET"
)

