// Version 2 of the AWS Cloud Development Kit library
package awscdk


// Supported Docker volume consistency types.
//
// Only valid on macOS due to the way file storage works on Mac.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var dockerImage dockerImage
//   var localBundling iLocalBundling
//
//   assetStagingProps := &assetStagingProps{
//   	sourcePath: jsii.String("sourcePath"),
//
//   	// the properties below are optional
//   	assetHash: jsii.String("assetHash"),
//   	assetHashType: cdk.assetHashType_SOURCE,
//   	bundling: &bundlingOptions{
//   		image: dockerImage,
//
//   		// the properties below are optional
//   		command: []*string{
//   			jsii.String("command"),
//   		},
//   		entrypoint: []*string{
//   			jsii.String("entrypoint"),
//   		},
//   		environment: map[string]*string{
//   			"environmentKey": jsii.String("environment"),
//   		},
//   		local: localBundling,
//   		network: jsii.String("network"),
//   		outputType: cdk.bundlingOutput_ARCHIVED,
//   		securityOpt: jsii.String("securityOpt"),
//   		user: jsii.String("user"),
//   		volumes: []dockerVolume{
//   			&dockerVolume{
//   				containerPath: jsii.String("containerPath"),
//   				hostPath: jsii.String("hostPath"),
//
//   				// the properties below are optional
//   				consistency: cdk.dockerVolumeConsistency_CONSISTENT,
//   			},
//   		},
//   		volumesFrom: []*string{
//   			jsii.String("volumesFrom"),
//   		},
//   		workingDirectory: jsii.String("workingDirectory"),
//   	},
//   	exclude: []*string{
//   		jsii.String("exclude"),
//   	},
//   	extraHash: jsii.String("extraHash"),
//   	follow: cdk.symlinkFollowMode_NEVER,
//   	ignoreMode: cdk.ignoreMode_GLOB,
//   }
//
type DockerVolumeConsistency string

const (
	// Read/write operations inside the Docker container are applied immediately on the mounted host machine volumes.
	DockerVolumeConsistency_CONSISTENT DockerVolumeConsistency = "CONSISTENT"
	// Read/write operations on mounted Docker volumes are first written inside the container and then synchronized to the host machine.
	DockerVolumeConsistency_DELEGATED DockerVolumeConsistency = "DELEGATED"
	// Read/write operations on mounted Docker volumes are first applied on the host machine and then synchronized to the container.
	DockerVolumeConsistency_CACHED DockerVolumeConsistency = "CACHED"
)

