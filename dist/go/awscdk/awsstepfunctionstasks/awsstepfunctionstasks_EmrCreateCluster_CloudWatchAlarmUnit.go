package awsstepfunctionstasks


// CloudWatch Alarm Units.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   autoScalingPolicyProperty := &autoScalingPolicyProperty{
//   	constraints: &scalingConstraintsProperty{
//   		maxCapacity: jsii.Number(123),
//   		minCapacity: jsii.Number(123),
//   	},
//   	rules: []scalingRuleProperty{
//   		&scalingRuleProperty{
//   			action: &scalingActionProperty{
//   				simpleScalingPolicyConfiguration: &simpleScalingPolicyConfigurationProperty{
//   					scalingAdjustment: jsii.Number(123),
//
//   					// the properties below are optional
//   					adjustmentType: awscdk.Aws_stepfunctions_tasks.emrCreateCluster.scalingAdjustmentType_CHANGE_IN_CAPACITY,
//   					coolDown: jsii.Number(123),
//   				},
//
//   				// the properties below are optional
//   				market: awscdk.*Aws_stepfunctions_tasks.*emrCreateCluster.instanceMarket_ON_DEMAND,
//   			},
//   			name: jsii.String("name"),
//   			trigger: &scalingTriggerProperty{
//   				cloudWatchAlarmDefinition: &cloudWatchAlarmDefinitionProperty{
//   					comparisonOperator: awscdk.*Aws_stepfunctions_tasks.*emrCreateCluster.cloudWatchAlarmComparisonOperator_GREATER_THAN_OR_EQUAL,
//   					metricName: jsii.String("metricName"),
//   					period: cdk.duration.minutes(jsii.Number(30)),
//
//   					// the properties below are optional
//   					dimensions: []metricDimensionProperty{
//   						&metricDimensionProperty{
//   							key: jsii.String("key"),
//   							value: jsii.String("value"),
//   						},
//   					},
//   					evaluationPeriods: jsii.Number(123),
//   					namespace: jsii.String("namespace"),
//   					statistic: awscdk.*Aws_stepfunctions_tasks.*emrCreateCluster.cloudWatchAlarmStatistic_SAMPLE_COUNT,
//   					threshold: jsii.Number(123),
//   					unit: awscdk.*Aws_stepfunctions_tasks.*emrCreateCluster.cloudWatchAlarmUnit_NONE,
//   				},
//   			},
//
//   			// the properties below are optional
//   			description: jsii.String("description"),
//   		},
//   	},
//   }
//
type EmrCreateCluster_CloudWatchAlarmUnit string

const (
	// NONE.
	EmrCreateCluster_CloudWatchAlarmUnit_NONE EmrCreateCluster_CloudWatchAlarmUnit = "NONE"
	// SECONDS.
	EmrCreateCluster_CloudWatchAlarmUnit_SECONDS EmrCreateCluster_CloudWatchAlarmUnit = "SECONDS"
	// MICRO_SECONDS.
	EmrCreateCluster_CloudWatchAlarmUnit_MICRO_SECONDS EmrCreateCluster_CloudWatchAlarmUnit = "MICRO_SECONDS"
	// MILLI_SECONDS.
	EmrCreateCluster_CloudWatchAlarmUnit_MILLI_SECONDS EmrCreateCluster_CloudWatchAlarmUnit = "MILLI_SECONDS"
	// BYTES.
	EmrCreateCluster_CloudWatchAlarmUnit_BYTES EmrCreateCluster_CloudWatchAlarmUnit = "BYTES"
	// KILO_BYTES.
	EmrCreateCluster_CloudWatchAlarmUnit_KILO_BYTES EmrCreateCluster_CloudWatchAlarmUnit = "KILO_BYTES"
	// MEGA_BYTES.
	EmrCreateCluster_CloudWatchAlarmUnit_MEGA_BYTES EmrCreateCluster_CloudWatchAlarmUnit = "MEGA_BYTES"
	// GIGA_BYTES.
	EmrCreateCluster_CloudWatchAlarmUnit_GIGA_BYTES EmrCreateCluster_CloudWatchAlarmUnit = "GIGA_BYTES"
	// TERA_BYTES.
	EmrCreateCluster_CloudWatchAlarmUnit_TERA_BYTES EmrCreateCluster_CloudWatchAlarmUnit = "TERA_BYTES"
	// BITS.
	EmrCreateCluster_CloudWatchAlarmUnit_BITS EmrCreateCluster_CloudWatchAlarmUnit = "BITS"
	// KILO_BITS.
	EmrCreateCluster_CloudWatchAlarmUnit_KILO_BITS EmrCreateCluster_CloudWatchAlarmUnit = "KILO_BITS"
	// MEGA_BITS.
	EmrCreateCluster_CloudWatchAlarmUnit_MEGA_BITS EmrCreateCluster_CloudWatchAlarmUnit = "MEGA_BITS"
	// GIGA_BITS.
	EmrCreateCluster_CloudWatchAlarmUnit_GIGA_BITS EmrCreateCluster_CloudWatchAlarmUnit = "GIGA_BITS"
	// TERA_BITS.
	EmrCreateCluster_CloudWatchAlarmUnit_TERA_BITS EmrCreateCluster_CloudWatchAlarmUnit = "TERA_BITS"
	// PERCENT.
	EmrCreateCluster_CloudWatchAlarmUnit_PERCENT EmrCreateCluster_CloudWatchAlarmUnit = "PERCENT"
	// COUNT.
	EmrCreateCluster_CloudWatchAlarmUnit_COUNT EmrCreateCluster_CloudWatchAlarmUnit = "COUNT"
	// BYTES_PER_SECOND.
	EmrCreateCluster_CloudWatchAlarmUnit_BYTES_PER_SECOND EmrCreateCluster_CloudWatchAlarmUnit = "BYTES_PER_SECOND"
	// KILO_BYTES_PER_SECOND.
	EmrCreateCluster_CloudWatchAlarmUnit_KILO_BYTES_PER_SECOND EmrCreateCluster_CloudWatchAlarmUnit = "KILO_BYTES_PER_SECOND"
	// MEGA_BYTES_PER_SECOND.
	EmrCreateCluster_CloudWatchAlarmUnit_MEGA_BYTES_PER_SECOND EmrCreateCluster_CloudWatchAlarmUnit = "MEGA_BYTES_PER_SECOND"
	// GIGA_BYTES_PER_SECOND.
	EmrCreateCluster_CloudWatchAlarmUnit_GIGA_BYTES_PER_SECOND EmrCreateCluster_CloudWatchAlarmUnit = "GIGA_BYTES_PER_SECOND"
	// TERA_BYTES_PER_SECOND.
	EmrCreateCluster_CloudWatchAlarmUnit_TERA_BYTES_PER_SECOND EmrCreateCluster_CloudWatchAlarmUnit = "TERA_BYTES_PER_SECOND"
	// BITS_PER_SECOND.
	EmrCreateCluster_CloudWatchAlarmUnit_BITS_PER_SECOND EmrCreateCluster_CloudWatchAlarmUnit = "BITS_PER_SECOND"
	// KILO_BITS_PER_SECOND.
	EmrCreateCluster_CloudWatchAlarmUnit_KILO_BITS_PER_SECOND EmrCreateCluster_CloudWatchAlarmUnit = "KILO_BITS_PER_SECOND"
	// MEGA_BITS_PER_SECOND.
	EmrCreateCluster_CloudWatchAlarmUnit_MEGA_BITS_PER_SECOND EmrCreateCluster_CloudWatchAlarmUnit = "MEGA_BITS_PER_SECOND"
	// GIGA_BITS_PER_SECOND.
	EmrCreateCluster_CloudWatchAlarmUnit_GIGA_BITS_PER_SECOND EmrCreateCluster_CloudWatchAlarmUnit = "GIGA_BITS_PER_SECOND"
	// TERA_BITS_PER_SECOND.
	EmrCreateCluster_CloudWatchAlarmUnit_TERA_BITS_PER_SECOND EmrCreateCluster_CloudWatchAlarmUnit = "TERA_BITS_PER_SECOND"
	// COUNT_PER_SECOND.
	EmrCreateCluster_CloudWatchAlarmUnit_COUNT_PER_SECOND EmrCreateCluster_CloudWatchAlarmUnit = "COUNT_PER_SECOND"
)

