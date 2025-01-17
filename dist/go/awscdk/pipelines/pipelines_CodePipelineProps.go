package pipelines

import (
	"github.com/aws/aws-cdk-go/awscdk/awscodepipeline"
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
)

// Properties for a `CodePipeline`.
//
// Example:
//   var codePipeline pipeline
//
//
//   pipeline := pipelines.NewCodePipeline(this, jsii.String("Pipeline"), &codePipelineProps{
//   	synth: pipelines.NewShellStep(jsii.String("Synth"), &shellStepProps{
//   		input: pipelines.codePipelineSource.connection(jsii.String("my-org/my-app"), jsii.String("main"), &connectionSourceOptions{
//   			connectionArn: jsii.String("arn:aws:codestar-connections:us-east-1:222222222222:connection/7d2469ff-514a-4e4f-9003-5ca4a43cdc41"),
//   		}),
//   		commands: []*string{
//   			jsii.String("npm ci"),
//   			jsii.String("npm run build"),
//   			jsii.String("npx cdk synth"),
//   		},
//   	}),
//   	codePipeline: codePipeline,
//   })
//
type CodePipelineProps struct {
	// The build step that produces the CDK Cloud Assembly.
	//
	// The primary output of this step needs to be the `cdk.out` directory
	// generated by the `cdk synth` command.
	//
	// If you use a `ShellStep` here and you don't configure an output directory,
	// the output directory will automatically be assumed to be `cdk.out`.
	Synth IFileSetProducer `field:"required" json:"synth" yaml:"synth"`
	// Additional customizations to apply to the asset publishing CodeBuild projects.
	AssetPublishingCodeBuildDefaults *CodeBuildOptions `field:"optional" json:"assetPublishingCodeBuildDefaults" yaml:"assetPublishingCodeBuildDefaults"`
	// CDK CLI version to use in self-mutation and asset publishing steps.
	//
	// If you want to lock the CDK CLI version used in the pipeline, by steps
	// that are automatically generated for you, specify the version here.
	//
	// We recommend you do not specify this value, as not specifying it always
	// uses the latest CLI version which is backwards compatible with old versions.
	//
	// If you do specify it, be aware that this version should always be equal to or higher than the
	// version of the CDK framework used by the CDK app, when the CDK commands are
	// run during your pipeline execution. When you change this version, the *next
	// time* the `SelfMutate` step runs it will still be using the CLI of the the
	// *previous* version that was in this property: it will only start using the
	// new version after `SelfMutate` completes successfully. That means that if
	// you want to update both framework and CLI version, you should update the
	// CLI version first, commit, push and deploy, and only then update the
	// framework version.
	CliVersion *string `field:"optional" json:"cliVersion" yaml:"cliVersion"`
	// Customize the CodeBuild projects created for this pipeline.
	CodeBuildDefaults *CodeBuildOptions `field:"optional" json:"codeBuildDefaults" yaml:"codeBuildDefaults"`
	// An existing Pipeline to be reused and built upon.
	//
	// [disable-awslint:ref-via-interface].
	CodePipeline awscodepipeline.Pipeline `field:"optional" json:"codePipeline" yaml:"codePipeline"`
	// Create KMS keys for the artifact buckets, allowing cross-account deployments.
	//
	// The artifact buckets have to be encrypted to support deploying CDK apps to
	// another account, so if you want to do that or want to have your artifact
	// buckets encrypted, be sure to set this value to `true`.
	//
	// Be aware there is a cost associated with maintaining the KMS keys.
	CrossAccountKeys *bool `field:"optional" json:"crossAccountKeys" yaml:"crossAccountKeys"`
	// A list of credentials used to authenticate to Docker registries.
	//
	// Specify any credentials necessary within the pipeline to build, synth, update, or publish assets.
	DockerCredentials *[]DockerCredential `field:"optional" json:"dockerCredentials" yaml:"dockerCredentials"`
	// Enable Docker for the self-mutate step.
	//
	// Set this to true if the pipeline itself uses Docker container assets
	// (for example, if you use `LinuxBuildImage.fromAsset()` as the build
	// image of a CodeBuild step in the pipeline).
	//
	// You do not need to set it if you build Docker image assets in the
	// application Stages and Stacks that are *deployed* by this pipeline.
	//
	// Configures privileged mode for the self-mutation CodeBuild action.
	//
	// If you are about to turn this on in an already-deployed Pipeline,
	// set the value to `true` first, commit and allow the pipeline to
	// self-update, and only then use the Docker asset in the pipeline.
	DockerEnabledForSelfMutation *bool `field:"optional" json:"dockerEnabledForSelfMutation" yaml:"dockerEnabledForSelfMutation"`
	// Enable Docker for the 'synth' step.
	//
	// Set this to true if you are using file assets that require
	// "bundling" anywhere in your application (meaning an asset
	// compilation step will be run with the tools provided by
	// a Docker image), both for the Pipeline stack as well as the
	// application stacks.
	//
	// A common way to use bundling assets in your application is by
	// using the `@aws-cdk/aws-lambda-nodejs` library.
	//
	// Configures privileged mode for the synth CodeBuild action.
	//
	// If you are about to turn this on in an already-deployed Pipeline,
	// set the value to `true` first, commit and allow the pipeline to
	// self-update, and only then use the bundled asset.
	DockerEnabledForSynth *bool `field:"optional" json:"dockerEnabledForSynth" yaml:"dockerEnabledForSynth"`
	// The name of the CodePipeline pipeline.
	PipelineName *string `field:"optional" json:"pipelineName" yaml:"pipelineName"`
	// Publish assets in multiple CodeBuild projects.
	//
	// If set to false, use one Project per type to publish all assets.
	//
	// Publishing in parallel improves concurrency and may reduce publishing
	// latency, but may also increase overall provisioning time of the CodeBuild
	// projects.
	//
	// Experiment and see what value works best for you.
	PublishAssetsInParallel *bool `field:"optional" json:"publishAssetsInParallel" yaml:"publishAssetsInParallel"`
	// Reuse the same cross region support stack for all pipelines in the App.
	ReuseCrossRegionSupportStacks *bool `field:"optional" json:"reuseCrossRegionSupportStacks" yaml:"reuseCrossRegionSupportStacks"`
	// The IAM role to be assumed by this Pipeline.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// Whether the pipeline will update itself.
	//
	// This needs to be set to `true` to allow the pipeline to reconfigure
	// itself when assets or stages are being added to it, and `true` is the
	// recommended setting.
	//
	// You can temporarily set this to `false` while you are iterating
	// on the pipeline itself and prefer to deploy changes using `cdk deploy`.
	SelfMutation *bool `field:"optional" json:"selfMutation" yaml:"selfMutation"`
	// Additional customizations to apply to the self mutation CodeBuild projects.
	SelfMutationCodeBuildDefaults *CodeBuildOptions `field:"optional" json:"selfMutationCodeBuildDefaults" yaml:"selfMutationCodeBuildDefaults"`
	// Additional customizations to apply to the synthesize CodeBuild projects.
	SynthCodeBuildDefaults *CodeBuildOptions `field:"optional" json:"synthCodeBuildDefaults" yaml:"synthCodeBuildDefaults"`
	// Deploy every stack by creating a change set and executing it.
	//
	// When enabled, creates a "Prepare" and "Execute" action for each stack. Disable
	// to deploy the stack in one pipeline action.
	UseChangeSets *bool `field:"optional" json:"useChangeSets" yaml:"useChangeSets"`
}

