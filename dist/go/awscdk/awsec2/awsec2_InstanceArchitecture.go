package awsec2


// Identifies an instance's CPU architecture.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   bottleRocketImageProps := &bottleRocketImageProps{
//   	architecture: awscdk.Aws_ec2.instanceArchitecture_ARM_64,
//   	cachedInContext: jsii.Boolean(false),
//   	variant: awscdk.Aws_ecs.bottlerocketEcsVariant_AWS_ECS_1,
//   }
//
type InstanceArchitecture string

const (
	// ARM64 architecture.
	InstanceArchitecture_ARM_64 InstanceArchitecture = "ARM_64"
	// x86-64 architecture.
	InstanceArchitecture_X86_64 InstanceArchitecture = "X86_64"
)

