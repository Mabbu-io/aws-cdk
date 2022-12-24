package cloudassemblyschema


// In what scenarios should the CLI ask for approval.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cdkCommands := &cdkCommands{
//   	deploy: &deployCommand{
//   		args: &deployOptions{
//   			all: jsii.Boolean(false),
//   			app: jsii.String("app"),
//   			assetMetadata: jsii.Boolean(false),
//   			caBundlePath: jsii.String("caBundlePath"),
//   			changeSetName: jsii.String("changeSetName"),
//   			ci: jsii.Boolean(false),
//   			color: jsii.Boolean(false),
//   			context: map[string]*string{
//   				"contextKey": jsii.String("context"),
//   			},
//   			debug: jsii.Boolean(false),
//   			ec2Creds: jsii.Boolean(false),
//   			exclusively: jsii.Boolean(false),
//   			execute: jsii.Boolean(false),
//   			force: jsii.Boolean(false),
//   			ignoreErrors: jsii.Boolean(false),
//   			json: jsii.Boolean(false),
//   			lookups: jsii.Boolean(false),
//   			notices: jsii.Boolean(false),
//   			notificationArns: []*string{
//   				jsii.String("notificationArns"),
//   			},
//   			output: jsii.String("output"),
//   			outputsFile: jsii.String("outputsFile"),
//   			parameters: map[string]*string{
//   				"parametersKey": jsii.String("parameters"),
//   			},
//   			pathMetadata: jsii.Boolean(false),
//   			profile: jsii.String("profile"),
//   			proxy: jsii.String("proxy"),
//   			requireApproval: awscdk.Cloud_assembly_schema.requireApproval_NEVER,
//   			reuseAssets: []*string{
//   				jsii.String("reuseAssets"),
//   			},
//   			roleArn: jsii.String("roleArn"),
//   			rollback: jsii.Boolean(false),
//   			stacks: []*string{
//   				jsii.String("stacks"),
//   			},
//   			staging: jsii.Boolean(false),
//   			strict: jsii.Boolean(false),
//   			toolkitStackName: jsii.String("toolkitStackName"),
//   			trace: jsii.Boolean(false),
//   			usePreviousParameters: jsii.Boolean(false),
//   			verbose: jsii.Boolean(false),
//   			versionReporting: jsii.Boolean(false),
//   		},
//   		enabled: jsii.Boolean(false),
//   		expectedMessage: jsii.String("expectedMessage"),
//   		expectError: jsii.Boolean(false),
//   	},
//   	destroy: &destroyCommand{
//   		args: &destroyOptions{
//   			all: jsii.Boolean(false),
//   			app: jsii.String("app"),
//   			assetMetadata: jsii.Boolean(false),
//   			caBundlePath: jsii.String("caBundlePath"),
//   			color: jsii.Boolean(false),
//   			context: map[string]*string{
//   				"contextKey": jsii.String("context"),
//   			},
//   			debug: jsii.Boolean(false),
//   			ec2Creds: jsii.Boolean(false),
//   			exclusively: jsii.Boolean(false),
//   			force: jsii.Boolean(false),
//   			ignoreErrors: jsii.Boolean(false),
//   			json: jsii.Boolean(false),
//   			lookups: jsii.Boolean(false),
//   			notices: jsii.Boolean(false),
//   			output: jsii.String("output"),
//   			pathMetadata: jsii.Boolean(false),
//   			profile: jsii.String("profile"),
//   			proxy: jsii.String("proxy"),
//   			roleArn: jsii.String("roleArn"),
//   			stacks: []*string{
//   				jsii.String("stacks"),
//   			},
//   			staging: jsii.Boolean(false),
//   			strict: jsii.Boolean(false),
//   			trace: jsii.Boolean(false),
//   			verbose: jsii.Boolean(false),
//   			versionReporting: jsii.Boolean(false),
//   		},
//   		enabled: jsii.Boolean(false),
//   		expectedMessage: jsii.String("expectedMessage"),
//   		expectError: jsii.Boolean(false),
//   	},
//   }
//
type RequireApproval string

const (
	// Never ask for approval.
	RequireApproval_NEVER RequireApproval = "NEVER"
	// Prompt for approval for any type  of change to the stack.
	RequireApproval_ANYCHANGE RequireApproval = "ANYCHANGE"
	// Only prompt for approval if there are security related changes.
	RequireApproval_BROADENING RequireApproval = "BROADENING"
)

