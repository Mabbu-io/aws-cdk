package awsstepfunctionstasks


// Spot Allocation Strategies.
//
// Specifies the strategy to use in launching Spot Instance fleets. For example, "capacity-optimized" launches instances from Spot Instance pools with optimal capacity for the number of instances that are launching.
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
// See: https://docs.aws.amazon.com/emr/latest/APIReference/API_SpotProvisioningSpecification.html
//
type EmrCreateCluster_SpotAllocationStrategy string

const (
	// Capacity-optimized, which launches instances from Spot Instance pools with optimal capacity for the number of instances that are launching.
	EmrCreateCluster_SpotAllocationStrategy_CAPACITY_OPTIMIZED EmrCreateCluster_SpotAllocationStrategy = "CAPACITY_OPTIMIZED"
)

