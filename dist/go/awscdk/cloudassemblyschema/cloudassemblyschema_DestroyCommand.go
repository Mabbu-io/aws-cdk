package cloudassemblyschema


// Represents a cdk destroy command.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   destroyCommand := &destroyCommand{
//   	args: &destroyOptions{
//   		all: jsii.Boolean(false),
//   		app: jsii.String("app"),
//   		assetMetadata: jsii.Boolean(false),
//   		caBundlePath: jsii.String("caBundlePath"),
//   		color: jsii.Boolean(false),
//   		context: map[string]*string{
//   			"contextKey": jsii.String("context"),
//   		},
//   		debug: jsii.Boolean(false),
//   		ec2Creds: jsii.Boolean(false),
//   		exclusively: jsii.Boolean(false),
//   		force: jsii.Boolean(false),
//   		ignoreErrors: jsii.Boolean(false),
//   		json: jsii.Boolean(false),
//   		lookups: jsii.Boolean(false),
//   		notices: jsii.Boolean(false),
//   		output: jsii.String("output"),
//   		pathMetadata: jsii.Boolean(false),
//   		profile: jsii.String("profile"),
//   		proxy: jsii.String("proxy"),
//   		roleArn: jsii.String("roleArn"),
//   		stacks: []*string{
//   			jsii.String("stacks"),
//   		},
//   		staging: jsii.Boolean(false),
//   		strict: jsii.Boolean(false),
//   		trace: jsii.Boolean(false),
//   		verbose: jsii.Boolean(false),
//   		versionReporting: jsii.Boolean(false),
//   	},
//   	enabled: jsii.Boolean(false),
//   	expectedMessage: jsii.String("expectedMessage"),
//   	expectError: jsii.Boolean(false),
//   }
//
type DestroyCommand struct {
	// Whether or not to run this command as part of the workflow This can be used if you only want to test some of the workflow for example enable `synth` and disable `deploy` & `destroy` in order to limit the test to synthesis.
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
	// This can be used in combination with `expectedError` to validate that a specific message is returned.
	ExpectedMessage *string `field:"optional" json:"expectedMessage" yaml:"expectedMessage"`
	// If the runner should expect this command to fail.
	ExpectError *bool `field:"optional" json:"expectError" yaml:"expectError"`
	// Additional arguments to pass to the command This can be used to test specific CLI functionality.
	Args *DestroyOptions `field:"optional" json:"args" yaml:"args"`
}

