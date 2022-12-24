package awsstepfunctionstasks


// Spot Timeout Actions.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   instanceFleetProvisioningSpecificationsProperty := &instanceFleetProvisioningSpecificationsProperty{
//   	spotSpecification: &spotProvisioningSpecificationProperty{
//   		timeoutAction: awscdk.Aws_stepfunctions_tasks.emrCreateCluster.spotTimeoutAction_SWITCH_TO_ON_DEMAND,
//   		timeoutDurationMinutes: jsii.Number(123),
//
//   		// the properties below are optional
//   		allocationStrategy: awscdk.*Aws_stepfunctions_tasks.*emrCreateCluster.spotAllocationStrategy_CAPACITY_OPTIMIZED,
//   		blockDurationMinutes: jsii.Number(123),
//   	},
//   }
//
type EmrCreateCluster_SpotTimeoutAction string

const (
	// SWITCH_TO_ON_DEMAND.
	EmrCreateCluster_SpotTimeoutAction_SWITCH_TO_ON_DEMAND EmrCreateCluster_SpotTimeoutAction = "SWITCH_TO_ON_DEMAND"
	// TERMINATE_CLUSTER.
	EmrCreateCluster_SpotTimeoutAction_TERMINATE_CLUSTER EmrCreateCluster_SpotTimeoutAction = "TERMINATE_CLUSTER"
)

