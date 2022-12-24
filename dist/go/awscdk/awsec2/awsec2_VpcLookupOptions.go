package awsec2


// Properties for looking up an existing VPC.
//
// The combination of properties must specify filter down to exactly one
// non-default VPC, otherwise an error is raised.
//
// Example:
//   vpc := ec2.vpc.fromLookup(this, jsii.String("Vpc"), &vpcLookupOptions{
//   	isDefault: jsii.Boolean(true),
//   })
//
//   cluster := ecs.NewCluster(this, jsii.String("FargateCluster"), &clusterProps{
//   	vpc: vpc,
//   })
//
//   taskDefinition := ecs.NewTaskDefinition(this, jsii.String("TD"), &taskDefinitionProps{
//   	memoryMiB: jsii.String("512"),
//   	cpu: jsii.String("256"),
//   	compatibility: ecs.compatibility_FARGATE,
//   })
//
//   containerDefinition := taskDefinition.addContainer(jsii.String("TheContainer"), &containerDefinitionOptions{
//   	image: ecs.containerImage.fromRegistry(jsii.String("foo/bar")),
//   	memoryLimitMiB: jsii.Number(256),
//   })
//
//   runTask := tasks.NewEcsRunTask(this, jsii.String("RunFargate"), &ecsRunTaskProps{
//   	integrationPattern: sfn.integrationPattern_RUN_JOB,
//   	cluster: cluster,
//   	taskDefinition: taskDefinition,
//   	assignPublicIp: jsii.Boolean(true),
//   	containerOverrides: []containerOverride{
//   		&containerOverride{
//   			containerDefinition: containerDefinition,
//   			environment: []taskEnvironmentVariable{
//   				&taskEnvironmentVariable{
//   					name: jsii.String("SOME_KEY"),
//   					value: sfn.jsonPath.stringAt(jsii.String("$.SomeKey")),
//   				},
//   			},
//   		},
//   	},
//   	launchTarget: tasks.NewEcsFargateLaunchTarget(),
//   })
//
type VpcLookupOptions struct {
	// Whether to match the default VPC.
	IsDefault *bool `field:"optional" json:"isDefault" yaml:"isDefault"`
	// Optional to override inferred region.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Optional tag for subnet group name.
	//
	// If not provided, we'll look at the aws-cdk:subnet-name tag.
	// If the subnet does not have the specified tag,
	// we'll use its type as the name.
	SubnetGroupNameTag *string `field:"optional" json:"subnetGroupNameTag" yaml:"subnetGroupNameTag"`
	// Tags on the VPC.
	//
	// The VPC must have all of these tags.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// The ID of the VPC.
	//
	// If given, will import exactly this VPC.
	VpcId *string `field:"optional" json:"vpcId" yaml:"vpcId"`
	// The name of the VPC.
	//
	// If given, will import the VPC with this name.
	VpcName *string `field:"optional" json:"vpcName" yaml:"vpcName"`
}

