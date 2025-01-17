package awsecs

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Complete base service properties that are required to be supplied by the implementation of the BaseService class.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var cluster cluster
//   var containerDefinition containerDefinition
//   var logDriver logDriver
//   var namespace iNamespace
//
//   baseServiceProps := &baseServiceProps{
//   	cluster: cluster,
//   	launchType: awscdk.Aws_ecs.launchType_EC2,
//
//   	// the properties below are optional
//   	capacityProviderStrategies: []capacityProviderStrategy{
//   		&capacityProviderStrategy{
//   			capacityProvider: jsii.String("capacityProvider"),
//
//   			// the properties below are optional
//   			base: jsii.Number(123),
//   			weight: jsii.Number(123),
//   		},
//   	},
//   	circuitBreaker: &deploymentCircuitBreaker{
//   		rollback: jsii.Boolean(false),
//   	},
//   	cloudMapOptions: &cloudMapOptions{
//   		cloudMapNamespace: namespace,
//   		container: containerDefinition,
//   		containerPort: jsii.Number(123),
//   		dnsRecordType: awscdk.Aws_servicediscovery.dnsRecordType_A,
//   		dnsTtl: cdk.duration.minutes(jsii.Number(30)),
//   		failureThreshold: jsii.Number(123),
//   		name: jsii.String("name"),
//   	},
//   	deploymentController: &deploymentController{
//   		type: awscdk.*Aws_ecs.deploymentControllerType_ECS,
//   	},
//   	desiredCount: jsii.Number(123),
//   	enableECSManagedTags: jsii.Boolean(false),
//   	enableExecuteCommand: jsii.Boolean(false),
//   	healthCheckGracePeriod: cdk.*duration.minutes(jsii.Number(30)),
//   	maxHealthyPercent: jsii.Number(123),
//   	minHealthyPercent: jsii.Number(123),
//   	propagateTags: awscdk.*Aws_ecs.propagatedTagSource_SERVICE,
//   	serviceConnectConfiguration: &serviceConnectProps{
//   		logDriver: logDriver,
//   		namespace: jsii.String("namespace"),
//   		services: []serviceConnectService{
//   			&serviceConnectService{
//   				portMappingName: jsii.String("portMappingName"),
//
//   				// the properties below are optional
//   				discoveryName: jsii.String("discoveryName"),
//   				dnsName: jsii.String("dnsName"),
//   				ingressPortOverride: jsii.Number(123),
//   				port: jsii.Number(123),
//   			},
//   		},
//   	},
//   	serviceName: jsii.String("serviceName"),
//   }
//
type BaseServiceProps struct {
	// The name of the cluster that hosts the service.
	Cluster ICluster `field:"required" json:"cluster" yaml:"cluster"`
	// A list of Capacity Provider strategies used to place a service.
	CapacityProviderStrategies *[]*CapacityProviderStrategy `field:"optional" json:"capacityProviderStrategies" yaml:"capacityProviderStrategies"`
	// Whether to enable the deployment circuit breaker.
	//
	// If this property is defined, circuit breaker will be implicitly
	// enabled.
	CircuitBreaker *DeploymentCircuitBreaker `field:"optional" json:"circuitBreaker" yaml:"circuitBreaker"`
	// The options for configuring an Amazon ECS service to use service discovery.
	CloudMapOptions *CloudMapOptions `field:"optional" json:"cloudMapOptions" yaml:"cloudMapOptions"`
	// Specifies which deployment controller to use for the service.
	//
	// For more information, see
	// [Amazon ECS Deployment Types](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/deployment-types.html)
	DeploymentController *DeploymentController `field:"optional" json:"deploymentController" yaml:"deploymentController"`
	// The desired number of instantiations of the task definition to keep running on the service.
	DesiredCount *float64 `field:"optional" json:"desiredCount" yaml:"desiredCount"`
	// Specifies whether to enable Amazon ECS managed tags for the tasks within the service.
	//
	// For more information, see
	// [Tagging Your Amazon ECS Resources](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-using-tags.html)
	EnableECSManagedTags *bool `field:"optional" json:"enableECSManagedTags" yaml:"enableECSManagedTags"`
	// Whether to enable the ability to execute into a container.
	EnableExecuteCommand *bool `field:"optional" json:"enableExecuteCommand" yaml:"enableExecuteCommand"`
	// The period of time, in seconds, that the Amazon ECS service scheduler ignores unhealthy Elastic Load Balancing target health checks after a task has first started.
	HealthCheckGracePeriod awscdk.Duration `field:"optional" json:"healthCheckGracePeriod" yaml:"healthCheckGracePeriod"`
	// The maximum number of tasks, specified as a percentage of the Amazon ECS service's DesiredCount value, that can run in a service during a deployment.
	MaxHealthyPercent *float64 `field:"optional" json:"maxHealthyPercent" yaml:"maxHealthyPercent"`
	// The minimum number of tasks, specified as a percentage of the Amazon ECS service's DesiredCount value, that must continue to run and remain healthy during a deployment.
	MinHealthyPercent *float64 `field:"optional" json:"minHealthyPercent" yaml:"minHealthyPercent"`
	// Specifies whether to propagate the tags from the task definition or the service to the tasks in the service.
	//
	// Valid values are: PropagatedTagSource.SERVICE, PropagatedTagSource.TASK_DEFINITION or PropagatedTagSource.NONE
	PropagateTags PropagatedTagSource `field:"optional" json:"propagateTags" yaml:"propagateTags"`
	// Configuration for Service Connect.
	ServiceConnectConfiguration *ServiceConnectProps `field:"optional" json:"serviceConnectConfiguration" yaml:"serviceConnectConfiguration"`
	// The name of the service.
	ServiceName *string `field:"optional" json:"serviceName" yaml:"serviceName"`
	// The launch type on which to run your service.
	//
	// LaunchType will be omitted if capacity provider strategies are specified on the service.
	// See: - https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-service.html#cfn-ecs-service-capacityproviderstrategy
	//
	// Valid values are: LaunchType.ECS or LaunchType.FARGATE or LaunchType.EXTERNAL
	//
	LaunchType LaunchType `field:"required" json:"launchType" yaml:"launchType"`
}

