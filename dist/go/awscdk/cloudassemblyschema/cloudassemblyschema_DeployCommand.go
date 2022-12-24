package cloudassemblyschema


// Represents a cdk deploy command.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   deployCommand := &deployCommand{
//   	args: &deployOptions{
//   		all: jsii.Boolean(false),
//   		app: jsii.String("app"),
//   		assetMetadata: jsii.Boolean(false),
//   		caBundlePath: jsii.String("caBundlePath"),
//   		changeSetName: jsii.String("changeSetName"),
//   		ci: jsii.Boolean(false),
//   		color: jsii.Boolean(false),
//   		context: map[string]*string{
//   			"contextKey": jsii.String("context"),
//   		},
//   		debug: jsii.Boolean(false),
//   		ec2Creds: jsii.Boolean(false),
//   		exclusively: jsii.Boolean(false),
//   		execute: jsii.Boolean(false),
//   		force: jsii.Boolean(false),
//   		ignoreErrors: jsii.Boolean(false),
//   		json: jsii.Boolean(false),
//   		lookups: jsii.Boolean(false),
//   		notices: jsii.Boolean(false),
//   		notificationArns: []*string{
//   			jsii.String("notificationArns"),
//   		},
//   		output: jsii.String("output"),
//   		outputsFile: jsii.String("outputsFile"),
//   		parameters: map[string]*string{
//   			"parametersKey": jsii.String("parameters"),
//   		},
//   		pathMetadata: jsii.Boolean(false),
//   		profile: jsii.String("profile"),
//   		proxy: jsii.String("proxy"),
//   		requireApproval: awscdk.Cloud_assembly_schema.requireApproval_NEVER,
//   		reuseAssets: []*string{
//   			jsii.String("reuseAssets"),
//   		},
//   		roleArn: jsii.String("roleArn"),
//   		rollback: jsii.Boolean(false),
//   		stacks: []*string{
//   			jsii.String("stacks"),
//   		},
//   		staging: jsii.Boolean(false),
//   		strict: jsii.Boolean(false),
//   		toolkitStackName: jsii.String("toolkitStackName"),
//   		trace: jsii.Boolean(false),
//   		usePreviousParameters: jsii.Boolean(false),
//   		verbose: jsii.Boolean(false),
//   		versionReporting: jsii.Boolean(false),
//   	},
//   	enabled: jsii.Boolean(false),
//   	expectedMessage: jsii.String("expectedMessage"),
//   	expectError: jsii.Boolean(false),
//   }
//
type DeployCommand struct {
	// Whether or not to run this command as part of the workflow This can be used if you only want to test some of the workflow for example enable `synth` and disable `deploy` & `destroy` in order to limit the test to synthesis.
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
	// This can be used in combination with `expectedError` to validate that a specific message is returned.
	ExpectedMessage *string `field:"optional" json:"expectedMessage" yaml:"expectedMessage"`
	// If the runner should expect this command to fail.
	ExpectError *bool `field:"optional" json:"expectError" yaml:"expectError"`
	// Additional arguments to pass to the command This can be used to test specific CLI functionality.
	Args *DeployOptions `field:"optional" json:"args" yaml:"args"`
}

