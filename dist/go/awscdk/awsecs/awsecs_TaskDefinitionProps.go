package awsecs

import (
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
)

// The properties for task definitions.
//
// Example:
//   vpc := ec2.vpc.fromLookup(this, jsii.String("Vpc"), &vpcLookupOptions{
//   	isDefault: jsii.Boolean(true),
//   })
//
//   cluster := ecs.NewCluster(this, jsii.String("Ec2Cluster"), &clusterProps{
//   	vpc: vpc,
//   })
//   cluster.addCapacity(jsii.String("DefaultAutoScalingGroup"), &addCapacityOptions{
//   	instanceType: ec2.NewInstanceType(jsii.String("t2.micro")),
//   	vpcSubnets: &subnetSelection{
//   		subnetType: ec2.subnetType_PUBLIC,
//   	},
//   })
//
//   taskDefinition := ecs.NewTaskDefinition(this, jsii.String("TD"), &taskDefinitionProps{
//   	compatibility: ecs.compatibility_EC2,
//   })
//
//   taskDefinition.addContainer(jsii.String("TheContainer"), &containerDefinitionOptions{
//   	image: ecs.containerImage.fromRegistry(jsii.String("foo/bar")),
//   	memoryLimitMiB: jsii.Number(256),
//   })
//
//   runTask := tasks.NewEcsRunTask(this, jsii.String("Run"), &ecsRunTaskProps{
//   	integrationPattern: sfn.integrationPattern_RUN_JOB,
//   	cluster: cluster,
//   	taskDefinition: taskDefinition,
//   	launchTarget: tasks.NewEcsEc2LaunchTarget(&ecsEc2LaunchTargetOptions{
//   		placementStrategies: []placementStrategy{
//   			ecs.*placementStrategy.spreadAcrossInstances(),
//   			ecs.*placementStrategy.packedByCpu(),
//   			ecs.*placementStrategy.randomly(),
//   		},
//   		placementConstraints: []placementConstraint{
//   			ecs.*placementConstraint.memberOf(jsii.String("blieptuut")),
//   		},
//   	}),
//   })
//
type TaskDefinitionProps struct {
	// The name of the IAM task execution role that grants the ECS agent permission to call AWS APIs on your behalf.
	//
	// The role will be used to retrieve container images from ECR and create CloudWatch log groups.
	ExecutionRole awsiam.IRole `field:"optional" json:"executionRole" yaml:"executionRole"`
	// The name of a family that this task definition is registered to.
	//
	// A family groups multiple versions of a task definition.
	Family *string `field:"optional" json:"family" yaml:"family"`
	// The configuration details for the App Mesh proxy.
	ProxyConfiguration ProxyConfiguration `field:"optional" json:"proxyConfiguration" yaml:"proxyConfiguration"`
	// The name of the IAM role that grants containers in the task permission to call AWS APIs on your behalf.
	TaskRole awsiam.IRole `field:"optional" json:"taskRole" yaml:"taskRole"`
	// The list of volume definitions for the task.
	//
	// For more information, see
	// [Task Definition Parameter Volumes](https://docs.aws.amazon.com/AmazonECS/latest/developerguide//task_definition_parameters.html#volumes).
	Volumes *[]*Volume `field:"optional" json:"volumes" yaml:"volumes"`
	// The task launch type compatiblity requirement.
	Compatibility Compatibility `field:"required" json:"compatibility" yaml:"compatibility"`
	// The number of cpu units used by the task.
	//
	// If you are using the EC2 launch type, this field is optional and any value can be used.
	// If you are using the Fargate launch type, this field is required and you must use one of the following values,
	// which determines your range of valid values for the memory parameter:
	//
	// 256 (.25 vCPU) - Available memory values: 512 (0.5 GB), 1024 (1 GB), 2048 (2 GB)
	//
	// 512 (.5 vCPU) - Available memory values: 1024 (1 GB), 2048 (2 GB), 3072 (3 GB), 4096 (4 GB)
	//
	// 1024 (1 vCPU) - Available memory values: 2048 (2 GB), 3072 (3 GB), 4096 (4 GB), 5120 (5 GB), 6144 (6 GB), 7168 (7 GB), 8192 (8 GB)
	//
	// 2048 (2 vCPU) - Available memory values: Between 4096 (4 GB) and 16384 (16 GB) in increments of 1024 (1 GB)
	//
	// 4096 (4 vCPU) - Available memory values: Between 8192 (8 GB) and 30720 (30 GB) in increments of 1024 (1 GB)
	//
	// 8192 (8 vCPU) - Available memory values: Between 16384 (16 GB) and 61440 (60 GB) in increments of 4096 (4 GB)
	//
	// 16384 (16 vCPU) - Available memory values: Between 32768 (32 GB) and 122880 (120 GB) in increments of 8192 (8 GB).
	Cpu *string `field:"optional" json:"cpu" yaml:"cpu"`
	// The amount (in GiB) of ephemeral storage to be allocated to the task.
	//
	// Only supported in Fargate platform version 1.4.0 or later.
	EphemeralStorageGiB *float64 `field:"optional" json:"ephemeralStorageGiB" yaml:"ephemeralStorageGiB"`
	// The inference accelerators to use for the containers in the task.
	//
	// Not supported in Fargate.
	InferenceAccelerators *[]*InferenceAccelerator `field:"optional" json:"inferenceAccelerators" yaml:"inferenceAccelerators"`
	// The IPC resource namespace to use for the containers in the task.
	//
	// Not supported in Fargate and Windows containers.
	IpcMode IpcMode `field:"optional" json:"ipcMode" yaml:"ipcMode"`
	// The amount (in MiB) of memory used by the task.
	//
	// If using the EC2 launch type, this field is optional and any value can be used.
	// If using the Fargate launch type, this field is required and you must use one of the following values,
	// which determines your range of valid values for the cpu parameter:
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
	// Between 32768 (32 GB) and 122880 (120 GB) in increments of 8192 (8 GB) - Available cpu values: 16384 (16 vCPU).
	MemoryMiB *string `field:"optional" json:"memoryMiB" yaml:"memoryMiB"`
	// The networking mode to use for the containers in the task.
	//
	// On Fargate, the only supported networking mode is AwsVpc.
	NetworkMode NetworkMode `field:"optional" json:"networkMode" yaml:"networkMode"`
	// The process namespace to use for the containers in the task.
	//
	// Not supported in Fargate and Windows containers.
	PidMode PidMode `field:"optional" json:"pidMode" yaml:"pidMode"`
	// The placement constraints to use for tasks in the service.
	//
	// You can specify a maximum of 10 constraints per task (this limit includes
	// constraints in the task definition and those specified at run time).
	//
	// Not supported in Fargate.
	PlacementConstraints *[]PlacementConstraint `field:"optional" json:"placementConstraints" yaml:"placementConstraints"`
	// The operating system that your task definitions are running on.
	//
	// A runtimePlatform is supported only for tasks using the Fargate launch type.
	RuntimePlatform *RuntimePlatform `field:"optional" json:"runtimePlatform" yaml:"runtimePlatform"`
}

