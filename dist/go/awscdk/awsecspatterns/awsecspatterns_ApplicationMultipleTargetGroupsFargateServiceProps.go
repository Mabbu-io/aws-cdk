package awsecspatterns

import (
	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/awsecs"
)

// The properties for the ApplicationMultipleTargetGroupsFargateService service.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   vpc := ec2.NewVpc(this, jsii.String("Vpc"), &vpcProps{
//   	maxAzs: jsii.Number(1),
//   })
//   loadBalancedFargateService := ecsPatterns.NewApplicationMultipleTargetGroupsFargateService(this, jsii.String("myService"), &applicationMultipleTargetGroupsFargateServiceProps{
//   	cluster: ecs.NewCluster(this, jsii.String("EcsCluster"), &clusterProps{
//   		vpc: vpc,
//   	}),
//   	memoryLimitMiB: jsii.Number(256),
//   	taskImageOptions: &applicationLoadBalancedTaskImageProps{
//   		image: ecs.containerImage.fromRegistry(jsii.String("amazon/amazon-ecs-sample")),
//   	},
//   	enableExecuteCommand: jsii.Boolean(true),
//   	loadBalancers: []applicationLoadBalancerProps{
//   		&applicationLoadBalancerProps{
//   			name: jsii.String("lb"),
//   			idleTimeout: awscdk.Duration.seconds(jsii.Number(400)),
//   			domainName: jsii.String("api.example.com"),
//   			domainZone: awscdk.NewPublicHostedZone(this, jsii.String("HostedZone"), &publicHostedZoneProps{
//   				zoneName: jsii.String("example.com"),
//   			}),
//   			listeners: []applicationListenerProps{
//   				&applicationListenerProps{
//   					name: jsii.String("listener"),
//   					protocol: awscdk.ApplicationProtocol_HTTPS,
//   					certificate: awscdk.Certificate.fromCertificateArn(this, jsii.String("Cert"), jsii.String("helloworld")),
//   					sslPolicy: awscdk.SslPolicy_TLS12_EXT,
//   				},
//   			},
//   		},
//   		&applicationLoadBalancerProps{
//   			name: jsii.String("lb2"),
//   			idleTimeout: awscdk.Duration.seconds(jsii.Number(120)),
//   			domainName: jsii.String("frontend.com"),
//   			domainZone: awscdk.NewPublicHostedZone(this, jsii.String("HostedZone"), &publicHostedZoneProps{
//   				zoneName: jsii.String("frontend.com"),
//   			}),
//   			listeners: []*applicationListenerProps{
//   				&applicationListenerProps{
//   					name: jsii.String("listener2"),
//   					protocol: awscdk.ApplicationProtocol_HTTPS,
//   					certificate: awscdk.Certificate.fromCertificateArn(this, jsii.String("Cert2"), jsii.String("helloworld")),
//   					sslPolicy: awscdk.SslPolicy_TLS12_EXT,
//   				},
//   			},
//   		},
//   	},
//   	targetGroups: []applicationTargetProps{
//   		&applicationTargetProps{
//   			containerPort: jsii.Number(80),
//   			listener: jsii.String("listener"),
//   		},
//   		&applicationTargetProps{
//   			containerPort: jsii.Number(90),
//   			pathPattern: jsii.String("a/b/c"),
//   			priority: jsii.Number(10),
//   			listener: jsii.String("listener"),
//   		},
//   		&applicationTargetProps{
//   			containerPort: jsii.Number(443),
//   			listener: jsii.String("listener2"),
//   		},
//   		&applicationTargetProps{
//   			containerPort: jsii.Number(80),
//   			pathPattern: jsii.String("a/b/c"),
//   			priority: jsii.Number(10),
//   			listener: jsii.String("listener2"),
//   		},
//   	},
//   })
//
type ApplicationMultipleTargetGroupsFargateServiceProps struct {
	// The options for configuring an Amazon ECS service to use service discovery.
	CloudMapOptions *awsecs.CloudMapOptions `field:"optional" json:"cloudMapOptions" yaml:"cloudMapOptions"`
	// The name of the cluster that hosts the service.
	//
	// If a cluster is specified, the vpc construct should be omitted. Alternatively, you can omit both cluster and vpc.
	Cluster awsecs.ICluster `field:"optional" json:"cluster" yaml:"cluster"`
	// The desired number of instantiations of the task definition to keep running on the service.
	DesiredCount *float64 `field:"optional" json:"desiredCount" yaml:"desiredCount"`
	// Specifies whether to enable Amazon ECS managed tags for the tasks within the service.
	//
	// For more information, see
	// [Tagging Your Amazon ECS Resources](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-using-tags.html)
	EnableECSManagedTags *bool `field:"optional" json:"enableECSManagedTags" yaml:"enableECSManagedTags"`
	// Whether ECS Exec should be enabled.
	EnableExecuteCommand *bool `field:"optional" json:"enableExecuteCommand" yaml:"enableExecuteCommand"`
	// The period of time, in seconds, that the Amazon ECS service scheduler ignores unhealthy Elastic Load Balancing target health checks after a task has first started.
	HealthCheckGracePeriod awscdk.Duration `field:"optional" json:"healthCheckGracePeriod" yaml:"healthCheckGracePeriod"`
	// The application load balancer that will serve traffic to the service.
	LoadBalancers *[]*ApplicationLoadBalancerProps `field:"optional" json:"loadBalancers" yaml:"loadBalancers"`
	// Specifies whether to propagate the tags from the task definition or the service to the tasks in the service.
	//
	// Tags can only be propagated to the tasks within the service during service creation.
	PropagateTags awsecs.PropagatedTagSource `field:"optional" json:"propagateTags" yaml:"propagateTags"`
	// The name of the service.
	ServiceName *string `field:"optional" json:"serviceName" yaml:"serviceName"`
	// Properties to specify ALB target groups.
	TargetGroups *[]*ApplicationTargetProps `field:"optional" json:"targetGroups" yaml:"targetGroups"`
	// The properties required to create a new task definition.
	//
	// Only one of TaskDefinition or TaskImageOptions must be specified.
	TaskImageOptions *ApplicationLoadBalancedTaskImageProps `field:"optional" json:"taskImageOptions" yaml:"taskImageOptions"`
	// The VPC where the container instances will be launched or the elastic network interfaces (ENIs) will be deployed.
	//
	// If a vpc is specified, the cluster construct should be omitted. Alternatively, you can omit both vpc and cluster.
	Vpc awsec2.IVpc `field:"optional" json:"vpc" yaml:"vpc"`
	// The number of cpu units used by the task.
	//
	// Valid values, which determines your range of valid values for the memory parameter:
	//
	// 256 (.25 vCPU) - Available memory values: 0.5GB, 1GB, 2GB
	//
	// 512 (.5 vCPU) - Available memory values: 1GB, 2GB, 3GB, 4GB
	//
	// 1024 (1 vCPU) - Available memory values: 2GB, 3GB, 4GB, 5GB, 6GB, 7GB, 8GB
	//
	// 2048 (2 vCPU) - Available memory values: Between 4GB and 16GB in 1GB increments
	//
	// 4096 (4 vCPU) - Available memory values: Between 8GB and 30GB in 1GB increments
	//
	// 8192 (8 vCPU) - Available memory values: Between 16GB and 60GB in 4GB increments
	//
	// 16384 (16 vCPU) - Available memory values: Between 32GB and 120GB in 8GB increments
	//
	// This default is set in the underlying FargateTaskDefinition construct.
	Cpu *float64 `field:"optional" json:"cpu" yaml:"cpu"`
	// The amount (in MiB) of memory used by the task.
	//
	// This field is required and you must use one of the following values, which determines your range of valid values
	// for the cpu parameter:
	//
	// 512 (0.5 GB), 1024 (1 GB), 2048 (2 GB) - Available cpu values: 256 (.25 vCPU)
	//
	// 1024 (1 GB), 2048 (2 GB), 3072 (3 GB), 4096 (4 GB) - Available cpu values: 512 (.5 vCPU)
	//
	// 2048 (2 GB), 3072 (3 GB), 4096 (4 GB), 5120 (5 GB), 6144 (6 GB), 7168 (7 GB), 8192 (8 GB) - Available cpu values: 1024 (1 vCPU)
	//
	// Between 4096 (4 GB) and 16384 (16 GB) in increments of 1024 (1 GB) - Available cpu values: 2048 (2 vCPU)
	//
	// Between 8192 (8 GB) and 30720 (30 GB) in increments of 1024 (1 GB) - Available cpu values: 4096 (4 vCPU)
	//
	// Between 16384 (16 GB) and 61440 (60 GB) in increments of 4096 (4 GB) - Available cpu values: 8192 (8 vCPU)
	//
	// Between 32768 (32 GB) and 122880 (120 GB) in increments of 8192 (8 GB) - Available cpu values: 16384 (16 vCPU)
	//
	// This default is set in the underlying FargateTaskDefinition construct.
	MemoryLimitMiB *float64 `field:"optional" json:"memoryLimitMiB" yaml:"memoryLimitMiB"`
	// The platform version on which to run your service.
	//
	// If one is not specified, the LATEST platform version is used by default. For more information, see
	// [AWS Fargate Platform Versions](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/platform_versions.html)
	// in the Amazon Elastic Container Service Developer Guide.
	PlatformVersion awsecs.FargatePlatformVersion `field:"optional" json:"platformVersion" yaml:"platformVersion"`
	// The runtime platform of the task definition.
	RuntimePlatform *awsecs.RuntimePlatform `field:"optional" json:"runtimePlatform" yaml:"runtimePlatform"`
	// The task definition to use for tasks in the service. TaskDefinition or TaskImageOptions must be specified, but not both.
	//
	// [disable-awslint:ref-via-interface].
	TaskDefinition awsecs.FargateTaskDefinition `field:"optional" json:"taskDefinition" yaml:"taskDefinition"`
	// Determines whether the service will be assigned a public IP address.
	AssignPublicIp *bool `field:"optional" json:"assignPublicIp" yaml:"assignPublicIp"`
}

