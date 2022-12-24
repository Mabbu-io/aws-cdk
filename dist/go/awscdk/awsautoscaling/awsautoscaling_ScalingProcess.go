package awsautoscaling


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   rollingUpdateOptions := &rollingUpdateOptions{
//   	maxBatchSize: jsii.Number(123),
//   	minInstancesInService: jsii.Number(123),
//   	minSuccessPercentage: jsii.Number(123),
//   	pauseTime: cdk.duration.minutes(jsii.Number(30)),
//   	suspendProcesses: []scalingProcess{
//   		awscdk.Aws_autoscaling.*scalingProcess_LAUNCH,
//   	},
//   	waitOnResourceSignals: jsii.Boolean(false),
//   }
//
type ScalingProcess string

const (
	ScalingProcess_LAUNCH ScalingProcess = "LAUNCH"
	ScalingProcess_TERMINATE ScalingProcess = "TERMINATE"
	ScalingProcess_HEALTH_CHECK ScalingProcess = "HEALTH_CHECK"
	ScalingProcess_REPLACE_UNHEALTHY ScalingProcess = "REPLACE_UNHEALTHY"
	ScalingProcess_AZ_REBALANCE ScalingProcess = "AZ_REBALANCE"
	ScalingProcess_ALARM_NOTIFICATION ScalingProcess = "ALARM_NOTIFICATION"
	ScalingProcess_SCHEDULED_ACTIONS ScalingProcess = "SCHEDULED_ACTIONS"
	ScalingProcess_ADD_TO_LOAD_BALANCER ScalingProcess = "ADD_TO_LOAD_BALANCER"
)

