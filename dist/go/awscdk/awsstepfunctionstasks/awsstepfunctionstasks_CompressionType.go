package awsstepfunctionstasks


// Compression type of the data.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var s3Location s3Location
//
//   channel := &channel{
//   	channelName: jsii.String("channelName"),
//   	dataSource: &dataSource{
//   		s3DataSource: &s3DataSource{
//   			s3Location: s3Location,
//
//   			// the properties below are optional
//   			attributeNames: []*string{
//   				jsii.String("attributeNames"),
//   			},
//   			s3DataDistributionType: awscdk.Aws_stepfunctions_tasks.s3DataDistributionType_FULLY_REPLICATED,
//   			s3DataType: awscdk.*Aws_stepfunctions_tasks.s3DataType_MANIFEST_FILE,
//   		},
//   	},
//
//   	// the properties below are optional
//   	compressionType: awscdk.*Aws_stepfunctions_tasks.compressionType_NONE,
//   	contentType: jsii.String("contentType"),
//   	inputMode: awscdk.*Aws_stepfunctions_tasks.inputMode_PIPE,
//   	recordWrapperType: awscdk.*Aws_stepfunctions_tasks.recordWrapperType_NONE,
//   	shuffleConfig: &shuffleConfig{
//   		seed: jsii.Number(123),
//   	},
//   }
//
type CompressionType string

const (
	// None compression type.
	CompressionType_NONE CompressionType = "NONE"
	// Gzip compression type.
	CompressionType_GZIP CompressionType = "GZIP"
)

