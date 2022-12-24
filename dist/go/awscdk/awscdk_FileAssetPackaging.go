// Version 2 of the AWS Cloud Development Kit library
package awscdk


// Packaging modes for file assets.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   fileAssetSource := &fileAssetSource{
//   	sourceHash: jsii.String("sourceHash"),
//
//   	// the properties below are optional
//   	executable: []*string{
//   		jsii.String("executable"),
//   	},
//   	fileName: jsii.String("fileName"),
//   	packaging: cdk.fileAssetPackaging_ZIP_DIRECTORY,
//   }
//
type FileAssetPackaging string

const (
	// The asset source path points to a directory, which should be archived using zip and and then uploaded to Amazon S3.
	FileAssetPackaging_ZIP_DIRECTORY FileAssetPackaging = "ZIP_DIRECTORY"
	// The asset source path points to a single file, which should be uploaded to Amazon S3.
	FileAssetPackaging_FILE FileAssetPackaging = "FILE"
)

