package awscodedeploy


// Lambda Deployment config type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customLambdaDeploymentConfig := awscdk.Aws_codedeploy.NewCustomLambdaDeploymentConfig(this, jsii.String("MyCustomLambdaDeploymentConfig"), &customLambdaDeploymentConfigProps{
//   	interval: cdk.duration.minutes(jsii.Number(30)),
//   	percentage: jsii.Number(123),
//   	type: awscdk.*Aws_codedeploy.customLambdaDeploymentConfigType_CANARY,
//
//   	// the properties below are optional
//   	deploymentConfigName: jsii.String("deploymentConfigName"),
//   })
//
// Deprecated: Use `LambdaDeploymentConfig`.
type CustomLambdaDeploymentConfigType string

const (
	// Canary deployment type.
	// Deprecated: Use `LambdaDeploymentConfig`.
	CustomLambdaDeploymentConfigType_CANARY CustomLambdaDeploymentConfigType = "CANARY"
	// Linear deployment type.
	// Deprecated: Use `LambdaDeploymentConfig`.
	CustomLambdaDeploymentConfigType_LINEAR CustomLambdaDeploymentConfigType = "LINEAR"
)

