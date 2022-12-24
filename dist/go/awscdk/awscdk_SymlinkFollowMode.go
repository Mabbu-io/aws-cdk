// Version 2 of the AWS Cloud Development Kit library
package awscdk


// Determines how symlinks are followed.
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
type SymlinkFollowMode string

const (
	// Never follow symlinks.
	SymlinkFollowMode_NEVER SymlinkFollowMode = "NEVER"
	// Materialize all symlinks, whether they are internal or external to the source directory.
	SymlinkFollowMode_ALWAYS SymlinkFollowMode = "ALWAYS"
	// Only follows symlinks that are external to the source directory.
	SymlinkFollowMode_EXTERNAL SymlinkFollowMode = "EXTERNAL"
	// Forbids source from having any symlinks pointing outside of the source tree.
	//
	// This is the safest mode of operation as it ensures that copy operations
	// won't materialize files from the user's file system. Internal symlinks are
	// not followed.
	//
	// If the copy operation runs into an external symlink, it will fail.
	SymlinkFollowMode_BLOCK_EXTERNAL SymlinkFollowMode = "BLOCK_EXTERNAL"
)

