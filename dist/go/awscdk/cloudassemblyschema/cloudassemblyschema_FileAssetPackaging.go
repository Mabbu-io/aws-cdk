package cloudassemblyschema


// Packaging strategy for file assets.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   fileAsset := &fileAsset{
//   	destinations: map[string]fileDestination{
//   		"destinationsKey": &fileDestination{
//   			"bucketName": jsii.String("bucketName"),
//   			"objectKey": jsii.String("objectKey"),
//
//   			// the properties below are optional
//   			"assumeRoleArn": jsii.String("assumeRoleArn"),
//   			"assumeRoleExternalId": jsii.String("assumeRoleExternalId"),
//   			"region": jsii.String("region"),
//   		},
//   	},
//   	source: &fileSource{
//   		executable: []*string{
//   			jsii.String("executable"),
//   		},
//   		packaging: awscdk.Cloud_assembly_schema.fileAssetPackaging_FILE,
//   		path: jsii.String("path"),
//   	},
//   }
//
type FileAssetPackaging string

const (
	// Upload the given path as a file.
	FileAssetPackaging_FILE FileAssetPackaging = "FILE"
	// The given path is a directory, zip it and upload.
	FileAssetPackaging_ZIP_DIRECTORY FileAssetPackaging = "ZIP_DIRECTORY"
)

