package awsecs


// The `ContainerDependency` property specifies the dependencies defined for container startup and shutdown.
//
// A container can contain multiple dependencies. When a dependency is defined for container startup, for container shutdown it is reversed.
//
// Your Amazon ECS container instances require at least version 1.26.0 of the container agent to enable container dependencies. However, we recommend using the latest container agent version. For information about checking your agent version and updating to the latest version, see [Updating the Amazon ECS Container Agent](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-agent-update.html) in the *Amazon Elastic Container Service Developer Guide* . If you are using an Amazon ECS-optimized Linux AMI, your instance needs at least version 1.26.0-1 of the `ecs-init` package. If your container instances are launched from version `20190301` or later, then they contain the required versions of the container agent and `ecs-init` . For more information, see [Amazon ECS-optimized Linux AMI](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-optimized_AMI.html) in the *Amazon Elastic Container Service Developer Guide* .
//
// > For tasks using the Fargate launch type, this parameter requires that the task or service uses platform version 1.3.0 or later.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   containerDependencyProperty := &containerDependencyProperty{
//   	condition: jsii.String("condition"),
//   	containerName: jsii.String("containerName"),
//   }
//
type CfnTaskDefinition_ContainerDependencyProperty struct {
	// The dependency condition of the container. The following are the available conditions and their behavior:.
	//
	// - `START` - This condition emulates the behavior of links and volumes today. It validates that a dependent container is started before permitting other containers to start.
	// - `COMPLETE` - This condition validates that a dependent container runs to completion (exits) before permitting other containers to start. This can be useful for nonessential containers that run a script and then exit. This condition can't be set on an essential container.
	// - `SUCCESS` - This condition is the same as `COMPLETE` , but it also requires that the container exits with a `zero` status. This condition can't be set on an essential container.
	// - `HEALTHY` - This condition validates that the dependent container passes its Docker health check before permitting other containers to start. This requires that the dependent container has health checks configured. This condition is confirmed only at task startup.
	Condition *string `field:"optional" json:"condition" yaml:"condition"`
	// The name of a container.
	ContainerName *string `field:"optional" json:"containerName" yaml:"containerName"`
}

