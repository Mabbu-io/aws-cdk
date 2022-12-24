package awsstepfunctionstasks


// Instance Role Types.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var configurationProperty_ configurationProperty
//   var size size
//
//   instanceFleetConfigProperty := &instanceFleetConfigProperty{
//   	instanceFleetType: awscdk.Aws_stepfunctions_tasks.emrCreateCluster.instanceRoleType_MASTER,
//
//   	// the properties below are optional
//   	instanceTypeConfigs: []instanceTypeConfigProperty{
//   		&instanceTypeConfigProperty{
//   			instanceType: jsii.String("instanceType"),
//
//   			// the properties below are optional
//   			bidPrice: jsii.String("bidPrice"),
//   			bidPriceAsPercentageOfOnDemandPrice: jsii.Number(123),
//   			configurations: []*configurationProperty{
//   				&configurationProperty{
//   					classification: jsii.String("classification"),
//   					configurations: []*configurationProperty{
//   						configurationProperty_,
//   					},
//   					properties: map[string]*string{
//   						"propertiesKey": jsii.String("properties"),
//   					},
//   				},
//   			},
//   			ebsConfiguration: &ebsConfigurationProperty{
//   				ebsBlockDeviceConfigs: []ebsBlockDeviceConfigProperty{
//   					&ebsBlockDeviceConfigProperty{
//   						volumeSpecification: &volumeSpecificationProperty{
//   							volumeSize: size,
//   							volumeType: awscdk.*Aws_stepfunctions_tasks.*emrCreateCluster.ebsBlockDeviceVolumeType_GP2,
//
//   							// the properties below are optional
//   							iops: jsii.Number(123),
//   						},
//
//   						// the properties below are optional
//   						volumesPerInstance: jsii.Number(123),
//   					},
//   				},
//   				ebsOptimized: jsii.Boolean(false),
//   			},
//   			weightedCapacity: jsii.Number(123),
//   		},
//   	},
//   	launchSpecifications: &instanceFleetProvisioningSpecificationsProperty{
//   		spotSpecification: &spotProvisioningSpecificationProperty{
//   			timeoutAction: awscdk.*Aws_stepfunctions_tasks.*emrCreateCluster.spotTimeoutAction_SWITCH_TO_ON_DEMAND,
//   			timeoutDurationMinutes: jsii.Number(123),
//
//   			// the properties below are optional
//   			allocationStrategy: awscdk.*Aws_stepfunctions_tasks.*emrCreateCluster.spotAllocationStrategy_CAPACITY_OPTIMIZED,
//   			blockDurationMinutes: jsii.Number(123),
//   		},
//   	},
//   	name: jsii.String("name"),
//   	targetOnDemandCapacity: jsii.Number(123),
//   	targetSpotCapacity: jsii.Number(123),
//   }
//
type EmrCreateCluster_InstanceRoleType string

const (
	// Master Node.
	EmrCreateCluster_InstanceRoleType_MASTER EmrCreateCluster_InstanceRoleType = "MASTER"
	// Core Node.
	EmrCreateCluster_InstanceRoleType_CORE EmrCreateCluster_InstanceRoleType = "CORE"
	// Task Node.
	EmrCreateCluster_InstanceRoleType_TASK EmrCreateCluster_InstanceRoleType = "TASK"
)

