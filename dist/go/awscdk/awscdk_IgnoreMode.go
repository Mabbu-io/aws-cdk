// Version 2 of the AWS Cloud Development Kit library
package awscdk


// Determines the ignore behavior to use.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var networkMode networkMode
//   var platform platform
//
//   dockerImageAssetOptions := &dockerImageAssetOptions{
//   	buildArgs: map[string]*string{
//   		"buildArgsKey": jsii.String("buildArgs"),
//   	},
//   	exclude: []*string{
//   		jsii.String("exclude"),
//   	},
//   	extraHash: jsii.String("extraHash"),
//   	file: jsii.String("file"),
//   	followSymlinks: cdk.symlinkFollowMode_NEVER,
//   	ignoreMode: cdk.ignoreMode_GLOB,
//   	invalidation: &dockerImageAssetInvalidationOptions{
//   		buildArgs: jsii.Boolean(false),
//   		extraHash: jsii.Boolean(false),
//   		file: jsii.Boolean(false),
//   		networkMode: jsii.Boolean(false),
//   		platform: jsii.Boolean(false),
//   		repositoryName: jsii.Boolean(false),
//   		target: jsii.Boolean(false),
//   	},
//   	networkMode: networkMode,
//   	platform: platform,
//   	target: jsii.String("target"),
//   }
//
type IgnoreMode string

const (
	// Ignores file paths based on simple glob patterns.
	//
	// This is the default for file assets.
	//
	// It is also the default for Docker image assets, unless the '@aws-cdk/aws-ecr-assets:dockerIgnoreSupport'
	// context flag is set.
	IgnoreMode_GLOB IgnoreMode = "GLOB"
	// Ignores file paths based on the [`.gitignore specification`](https://git-scm.com/docs/gitignore).
	IgnoreMode_GIT IgnoreMode = "GIT"
	// Ignores file paths based on the [`.dockerignore specification`](https://docs.docker.com/engine/reference/builder/#dockerignore-file).
	//
	// This is the default for Docker image assets if the '@aws-cdk/aws-ecr-assets:dockerIgnoreSupport'
	// context flag is set.
	IgnoreMode_DOCKER IgnoreMode = "DOCKER"
)

