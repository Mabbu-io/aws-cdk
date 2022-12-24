package awsautoscaling


// How the scaling metric is going to be aggregated.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var autoScalingGroup autoScalingGroup
//
//   stepScalingActionProps := &stepScalingActionProps{
//   	autoScalingGroup: autoScalingGroup,
//
//   	// the properties below are optional
//   	adjustmentType: awscdk.Aws_autoscaling.adjustmentType_CHANGE_IN_CAPACITY,
//   	cooldown: cdk.duration.minutes(jsii.Number(30)),
//   	estimatedInstanceWarmup: cdk.*duration.minutes(jsii.Number(30)),
//   	metricAggregationType: awscdk.*Aws_autoscaling.metricAggregationType_AVERAGE,
//   	minAdjustmentMagnitude: jsii.Number(123),
//   }
//
type MetricAggregationType string

const (
	// Average.
	MetricAggregationType_AVERAGE MetricAggregationType = "AVERAGE"
	// Minimum.
	MetricAggregationType_MINIMUM MetricAggregationType = "MINIMUM"
	// Maximum.
	MetricAggregationType_MAXIMUM MetricAggregationType = "MAXIMUM"
)

