package awscodedeploy


// The compute platform of a deployment configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var minimumHealthyHosts minimumHealthyHosts
//   var trafficRouting trafficRouting
//
//   baseDeploymentConfigProps := &baseDeploymentConfigProps{
//   	computePlatform: awscdk.Aws_codedeploy.computePlatform_SERVER,
//   	deploymentConfigName: jsii.String("deploymentConfigName"),
//   	minimumHealthyHosts: minimumHealthyHosts,
//   	trafficRouting: trafficRouting,
//   }
//
type ComputePlatform string

const (
	// The deployment will target EC2 instances or on-premise servers.
	ComputePlatform_SERVER ComputePlatform = "SERVER"
	// The deployment will target a Lambda function.
	ComputePlatform_LAMBDA ComputePlatform = "LAMBDA"
	// The deployment will target an ECS server.
	ComputePlatform_ECS ComputePlatform = "ECS"
)

