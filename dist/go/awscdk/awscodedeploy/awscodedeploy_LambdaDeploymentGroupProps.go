package awscodedeploy

import (
	"github.com/aws/aws-cdk-go/awscdk/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/awslambda"
)

// Construction properties for {@link LambdaDeploymentGroup}.
//
// Example:
//   var myApplication lambdaApplication
//   var func function
//
//   version := func.currentVersion
//   version1Alias := lambda.NewAlias(this, jsii.String("alias"), &aliasProps{
//   	aliasName: jsii.String("prod"),
//   	version: version,
//   })
//
//   deploymentGroup := codedeploy.NewLambdaDeploymentGroup(this, jsii.String("BlueGreenDeployment"), &lambdaDeploymentGroupProps{
//   	application: myApplication,
//   	 // optional property: one will be created for you if not provided
//   	alias: version1Alias,
//   	deploymentConfig: codedeploy.lambdaDeploymentConfig_LINEAR_10PERCENT_EVERY_1MINUTE(),
//   })
//
type LambdaDeploymentGroupProps struct {
	// Lambda Alias to shift traffic. Updating the version of the alias will trigger a CodeDeploy deployment.
	//
	// [disable-awslint:ref-via-interface] since we need to modify the alias CFN resource update policy.
	Alias awslambda.Alias `field:"required" json:"alias" yaml:"alias"`
	// The CloudWatch alarms associated with this Deployment Group.
	//
	// CodeDeploy will stop (and optionally roll back)
	// a deployment if during it any of the alarms trigger.
	//
	// Alarms can also be added after the Deployment Group is created using the {@link #addAlarm} method.
	// See: https://docs.aws.amazon.com/codedeploy/latest/userguide/monitoring-create-alarms.html
	//
	Alarms *[]awscloudwatch.IAlarm `field:"optional" json:"alarms" yaml:"alarms"`
	// The reference to the CodeDeploy Lambda Application that this Deployment Group belongs to.
	Application ILambdaApplication `field:"optional" json:"application" yaml:"application"`
	// The auto-rollback configuration for this Deployment Group.
	AutoRollback *AutoRollbackConfig `field:"optional" json:"autoRollback" yaml:"autoRollback"`
	// The Deployment Configuration this Deployment Group uses.
	DeploymentConfig ILambdaDeploymentConfig `field:"optional" json:"deploymentConfig" yaml:"deploymentConfig"`
	// The physical, human-readable name of the CodeDeploy Deployment Group.
	DeploymentGroupName *string `field:"optional" json:"deploymentGroupName" yaml:"deploymentGroupName"`
	// Whether to continue a deployment even if fetching the alarm status from CloudWatch failed.
	IgnorePollAlarmsFailure *bool `field:"optional" json:"ignorePollAlarmsFailure" yaml:"ignorePollAlarmsFailure"`
	// The Lambda function to run after traffic routing starts.
	PostHook awslambda.IFunction `field:"optional" json:"postHook" yaml:"postHook"`
	// The Lambda function to run before traffic routing starts.
	PreHook awslambda.IFunction `field:"optional" json:"preHook" yaml:"preHook"`
	// The service Role of this Deployment Group.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
}

