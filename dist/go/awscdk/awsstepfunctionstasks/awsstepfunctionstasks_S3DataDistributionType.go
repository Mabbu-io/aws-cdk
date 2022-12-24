package awsstepfunctionstasks


// S3 Data Distribution Type.
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
type S3DataDistributionType string

const (
	// Fully replicated S3 Data Distribution Type.
	S3DataDistributionType_FULLY_REPLICATED S3DataDistributionType = "FULLY_REPLICATED"
	// Sharded By S3 Key Data Distribution Type.
	S3DataDistributionType_SHARDED_BY_S3_KEY S3DataDistributionType = "SHARDED_BY_S3_KEY"
)

