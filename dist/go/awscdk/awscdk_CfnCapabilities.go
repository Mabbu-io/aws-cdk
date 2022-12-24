// Version 2 of the AWS Cloud Development Kit library
package awscdk


// Capabilities that affect whether CloudFormation is allowed to change IAM resources.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var artifact artifact
//   var artifactPath artifactPath
//   var parameterOverrides interface{}
//   var role role
//
//   cloudFormationDeleteStackAction := awscdk.Aws_codepipeline_actions.NewCloudFormationDeleteStackAction(&cloudFormationDeleteStackActionProps{
//   	actionName: jsii.String("actionName"),
//   	adminPermissions: jsii.Boolean(false),
//   	stackName: jsii.String("stackName"),
//
//   	// the properties below are optional
//   	account: jsii.String("account"),
//   	cfnCapabilities: []cfnCapabilities{
//   		cdk.*cfnCapabilities_NONE,
//   	},
//   	deploymentRole: role,
//   	extraInputs: []*artifact{
//   		artifact,
//   	},
//   	output: artifact,
//   	outputFileName: jsii.String("outputFileName"),
//   	parameterOverrides: map[string]interface{}{
//   		"parameterOverridesKey": parameterOverrides,
//   	},
//   	region: jsii.String("region"),
//   	role: role,
//   	runOrder: jsii.Number(123),
//   	templateConfiguration: artifactPath,
//   	variablesNamespace: jsii.String("variablesNamespace"),
//   })
//
type CfnCapabilities string

const (
	// No IAM Capabilities.
	//
	// Pass this capability if you wish to block the creation IAM resources.
	CfnCapabilities_NONE CfnCapabilities = "NONE"
	// Capability to create anonymous IAM resources.
	//
	// Pass this capability if you're only creating anonymous resources.
	CfnCapabilities_ANONYMOUS_IAM CfnCapabilities = "ANONYMOUS_IAM"
	// Capability to create named IAM resources.
	//
	// Pass this capability if you're creating IAM resources that have physical
	// names.
	//
	// `CloudFormationCapabilities.NamedIAM` implies `CloudFormationCapabilities.IAM`; you don't have to pass both.
	CfnCapabilities_NAMED_IAM CfnCapabilities = "NAMED_IAM"
	// Capability to run CloudFormation macros.
	//
	// Pass this capability if your template includes macros, for example AWS::Include or AWS::Serverless.
	CfnCapabilities_AUTO_EXPAND CfnCapabilities = "AUTO_EXPAND"
)

