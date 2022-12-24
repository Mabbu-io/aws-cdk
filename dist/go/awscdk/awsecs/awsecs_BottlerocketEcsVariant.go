package awsecs


// Amazon ECS variant.
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
type BottlerocketEcsVariant string

const (
	// aws-ecs-1 variant.
	BottlerocketEcsVariant_AWS_ECS_1 BottlerocketEcsVariant = "AWS_ECS_1"
)

