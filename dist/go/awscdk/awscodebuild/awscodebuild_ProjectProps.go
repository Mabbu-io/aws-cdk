package awscodebuild

import (
	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/awskms"
)

// Example:
//   var bucket bucket
//
//
//   project := codebuild.NewProject(this, jsii.String("MyProject"), &projectProps{
//   	buildSpec: codebuild.buildSpec.fromObject(map[string]interface{}{
//   		"version": jsii.String("0.2"),
//   	}),
//   	artifacts: codebuild.artifacts.s3(&s3ArtifactsProps{
//   		bucket: bucket,
//   		includeBuildId: jsii.Boolean(false),
//   		packageZip: jsii.Boolean(true),
//   		path: jsii.String("another/path"),
//   		identifier: jsii.String("AddArtifact1"),
//   	}),
//   })
//
type ProjectProps struct {
	// Whether to allow the CodeBuild to send all network traffic.
	//
	// If set to false, you must individually add traffic rules to allow the
	// CodeBuild project to connect to network targets.
	//
	// Only used if 'vpc' is supplied.
	AllowAllOutbound *bool `field:"optional" json:"allowAllOutbound" yaml:"allowAllOutbound"`
	// Indicates whether AWS CodeBuild generates a publicly accessible URL for your project's build badge.
	//
	// For more information, see Build Badges Sample
	// in the AWS CodeBuild User Guide.
	Badge *bool `field:"optional" json:"badge" yaml:"badge"`
	// Filename or contents of buildspec in JSON format.
	// See: https://docs.aws.amazon.com/codebuild/latest/userguide/build-spec-ref.html#build-spec-ref-example
	//
	BuildSpec BuildSpec `field:"optional" json:"buildSpec" yaml:"buildSpec"`
	// Caching strategy to use.
	Cache Cache `field:"optional" json:"cache" yaml:"cache"`
	// Whether to check for the presence of any secrets in the environment variables of the default type, BuildEnvironmentVariableType.PLAINTEXT. Since using a secret for the value of that kind of variable would result in it being displayed in plain text in the AWS Console, the construct will throw an exception if it detects a secret was passed there. Pass this property as false if you want to skip this validation, and keep using a secret in a plain text environment variable.
	CheckSecretsInPlainTextEnvVariables *bool `field:"optional" json:"checkSecretsInPlainTextEnvVariables" yaml:"checkSecretsInPlainTextEnvVariables"`
	// Maximum number of concurrent builds.
	//
	// Minimum value is 1 and maximum is account build limit.
	ConcurrentBuildLimit *float64 `field:"optional" json:"concurrentBuildLimit" yaml:"concurrentBuildLimit"`
	// A description of the project.
	//
	// Use the description to identify the purpose
	// of the project.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Encryption key to use to read and write artifacts.
	EncryptionKey awskms.IKey `field:"optional" json:"encryptionKey" yaml:"encryptionKey"`
	// Build environment to use for the build.
	Environment *BuildEnvironment `field:"optional" json:"environment" yaml:"environment"`
	// Additional environment variables to add to the build environment.
	EnvironmentVariables *map[string]*BuildEnvironmentVariable `field:"optional" json:"environmentVariables" yaml:"environmentVariables"`
	// An  ProjectFileSystemLocation objects for a CodeBuild build project.
	//
	// A ProjectFileSystemLocation object specifies the identifier, location, mountOptions, mountPoint,
	// and type of a file system created using Amazon Elastic File System.
	FileSystemLocations *[]IFileSystemLocation `field:"optional" json:"fileSystemLocations" yaml:"fileSystemLocations"`
	// Add permissions to this project's role to create and use test report groups with name starting with the name of this project.
	//
	// That is the standard report group that gets created when a simple name
	// (in contrast to an ARN)
	// is used in the 'reports' section of the buildspec of this project.
	// This is usually harmless, but you can turn these off if you don't plan on using test
	// reports in this project.
	// See: https://docs.aws.amazon.com/codebuild/latest/userguide/test-report-group-naming.html
	//
	GrantReportGroupPermissions *bool `field:"optional" json:"grantReportGroupPermissions" yaml:"grantReportGroupPermissions"`
	// Information about logs for the build project.
	//
	// A project can create logs in Amazon CloudWatch Logs, an S3 bucket, or both.
	Logging *LoggingOptions `field:"optional" json:"logging" yaml:"logging"`
	// The physical, human-readable name of the CodeBuild Project.
	ProjectName *string `field:"optional" json:"projectName" yaml:"projectName"`
	// The number of minutes after which AWS CodeBuild stops the build if it's still in queue.
	//
	// For valid values, see the timeoutInMinutes field in the AWS
	// CodeBuild User Guide.
	QueuedTimeout awscdk.Duration `field:"optional" json:"queuedTimeout" yaml:"queuedTimeout"`
	// Service Role to assume while running the build.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// What security group to associate with the codebuild project's network interfaces.
	//
	// If no security group is identified, one will be created automatically.
	//
	// Only used if 'vpc' is supplied.
	SecurityGroups *[]awsec2.ISecurityGroup `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// Add the permissions necessary for debugging builds with SSM Session Manager.
	//
	// If the following prerequisites have been met:
	//
	// - The necessary permissions have been added by setting this flag to true.
	// - The build image has the SSM agent installed (true for default CodeBuild images).
	// - The build is started with [debugSessionEnabled](https://docs.aws.amazon.com/codebuild/latest/APIReference/API_StartBuild.html#CodeBuild-StartBuild-request-debugSessionEnabled) set to true.
	//
	// Then the build container can be paused and inspected using Session Manager
	// by invoking the `codebuild-breakpoint` command somewhere during the build.
	//
	// `codebuild-breakpoint` commands will be ignored if the build is not started
	// with `debugSessionEnabled=true`.
	// See: https://docs.aws.amazon.com/codebuild/latest/userguide/session-manager.html
	//
	SsmSessionPermissions *bool `field:"optional" json:"ssmSessionPermissions" yaml:"ssmSessionPermissions"`
	// Where to place the network interfaces within the VPC.
	//
	// Only used if 'vpc' is supplied.
	SubnetSelection *awsec2.SubnetSelection `field:"optional" json:"subnetSelection" yaml:"subnetSelection"`
	// The number of minutes after which AWS CodeBuild stops the build if it's not complete.
	//
	// For valid values, see the timeoutInMinutes field in the AWS
	// CodeBuild User Guide.
	Timeout awscdk.Duration `field:"optional" json:"timeout" yaml:"timeout"`
	// VPC network to place codebuild network interfaces.
	//
	// Specify this if the codebuild project needs to access resources in a VPC.
	Vpc awsec2.IVpc `field:"optional" json:"vpc" yaml:"vpc"`
	// Defines where build artifacts will be stored.
	//
	// Could be: PipelineBuildArtifacts, NoArtifacts and S3Artifacts.
	Artifacts IArtifacts `field:"optional" json:"artifacts" yaml:"artifacts"`
	// The secondary artifacts for the Project.
	//
	// Can also be added after the Project has been created by using the {@link Project#addSecondaryArtifact} method.
	// See: https://docs.aws.amazon.com/codebuild/latest/userguide/sample-multi-in-out.html
	//
	SecondaryArtifacts *[]IArtifacts `field:"optional" json:"secondaryArtifacts" yaml:"secondaryArtifacts"`
	// The secondary sources for the Project.
	//
	// Can be also added after the Project has been created by using the {@link Project#addSecondarySource} method.
	// See: https://docs.aws.amazon.com/codebuild/latest/userguide/sample-multi-in-out.html
	//
	SecondarySources *[]ISource `field:"optional" json:"secondarySources" yaml:"secondarySources"`
	// The source of the build.
	//
	// *Note*: if {@link NoSource} is given as the source,
	// then you need to provide an explicit `buildSpec`.
	Source ISource `field:"optional" json:"source" yaml:"source"`
}

