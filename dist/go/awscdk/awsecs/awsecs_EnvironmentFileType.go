package awsecs


// Type of environment file to be included in the container definition.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   environmentFileConfig := &environmentFileConfig{
//   	fileType: awscdk.Aws_ecs.environmentFileType_S3,
//   	s3Location: &location{
//   		bucketName: jsii.String("bucketName"),
//   		objectKey: jsii.String("objectKey"),
//
//   		// the properties below are optional
//   		objectVersion: jsii.String("objectVersion"),
//   	},
//   }
//
type EnvironmentFileType string

const (
	// Environment file hosted on S3, referenced by object ARN.
	EnvironmentFileType_S3 EnvironmentFileType = "S3"
)

