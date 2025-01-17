package awsautoscaling

import (
	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awscloudwatch"
)

// Example:
//   var autoScalingGroup autoScalingGroup
//
//
//   workerUtilizationMetric := cloudwatch.NewMetric(&metricProps{
//   	namespace: jsii.String("MyService"),
//   	metricName: jsii.String("WorkerUtilization"),
//   })
//
//   autoScalingGroup.scaleOnMetric(jsii.String("ScaleToCPU"), &basicStepScalingPolicyProps{
//   	metric: workerUtilizationMetric,
//   	scalingSteps: []scalingInterval{
//   		&scalingInterval{
//   			upper: jsii.Number(10),
//   			change: -jsii.Number(1),
//   		},
//   		&scalingInterval{
//   			lower: jsii.Number(50),
//   			change: +jsii.Number(1),
//   		},
//   		&scalingInterval{
//   			lower: jsii.Number(70),
//   			change: +jsii.Number(3),
//   		},
//   	},
//
//   	// Change this to AdjustmentType.PERCENT_CHANGE_IN_CAPACITY to interpret the
//   	// 'change' numbers before as percentages instead of capacity counts.
//   	adjustmentType: autoscaling.adjustmentType_CHANGE_IN_CAPACITY,
//   })
//
type BasicStepScalingPolicyProps struct {
	// Metric to scale on.
	Metric awscloudwatch.IMetric `field:"required" json:"metric" yaml:"metric"`
	// The intervals for scaling.
	//
	// Maps a range of metric values to a particular scaling behavior.
	ScalingSteps *[]*ScalingInterval `field:"required" json:"scalingSteps" yaml:"scalingSteps"`
	// How the adjustment numbers inside 'intervals' are interpreted.
	AdjustmentType AdjustmentType `field:"optional" json:"adjustmentType" yaml:"adjustmentType"`
	// Grace period after scaling activity.
	Cooldown awscdk.Duration `field:"optional" json:"cooldown" yaml:"cooldown"`
	// Estimated time until a newly launched instance can send metrics to CloudWatch.
	EstimatedInstanceWarmup awscdk.Duration `field:"optional" json:"estimatedInstanceWarmup" yaml:"estimatedInstanceWarmup"`
	// How many evaluation periods of the metric to wait before triggering a scaling action.
	//
	// Raising this value can be used to smooth out the metric, at the expense
	// of slower response times.
	EvaluationPeriods *float64 `field:"optional" json:"evaluationPeriods" yaml:"evaluationPeriods"`
	// Aggregation to apply to all data points over the evaluation periods.
	//
	// Only has meaning if `evaluationPeriods != 1`.
	MetricAggregationType MetricAggregationType `field:"optional" json:"metricAggregationType" yaml:"metricAggregationType"`
	// Minimum absolute number to adjust capacity with as result of percentage scaling.
	//
	// Only when using AdjustmentType = PercentChangeInCapacity, this number controls
	// the minimum absolute effect size.
	MinAdjustmentMagnitude *float64 `field:"optional" json:"minAdjustmentMagnitude" yaml:"minAdjustmentMagnitude"`
}

