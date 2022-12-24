package awsecs


// The launch type of an ECS service.
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
type LaunchType string

const (
	// The service will be launched using the EC2 launch type.
	LaunchType_EC2 LaunchType = "EC2"
	// The service will be launched using the FARGATE launch type.
	LaunchType_FARGATE LaunchType = "FARGATE"
	// The service will be launched using the EXTERNAL launch type.
	LaunchType_EXTERNAL LaunchType = "EXTERNAL"
)

