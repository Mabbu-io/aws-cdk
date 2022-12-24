package pipelines


// Type of the asset that is being published.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   stackAsset := &stackAsset{
//   	assetId: jsii.String("assetId"),
//   	assetManifestPath: jsii.String("assetManifestPath"),
//   	assetSelector: jsii.String("assetSelector"),
//   	assetType: awscdk.Pipelines.assetType_FILE,
//   	isTemplate: jsii.Boolean(false),
//
//   	// the properties below are optional
//   	assetPublishingRoleArn: jsii.String("assetPublishingRoleArn"),
//   }
//
type AssetType string

const (
	// A file.
	AssetType_FILE AssetType = "FILE"
	// A Docker image.
	AssetType_DOCKER_IMAGE AssetType = "DOCKER_IMAGE"
)

