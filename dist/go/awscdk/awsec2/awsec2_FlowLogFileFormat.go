package awsec2


// The file format for flow logs written to an S3 bucket destination.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   destinationOptions := &destinationOptions{
//   	fileFormat: awscdk.Aws_ec2.flowLogFileFormat_PLAIN_TEXT,
//   	hiveCompatiblePartitions: jsii.Boolean(false),
//   	perHourPartition: jsii.Boolean(false),
//   }
//
type FlowLogFileFormat string

const (
	// File will be written as plain text.
	//
	// This is the default value.
	FlowLogFileFormat_PLAIN_TEXT FlowLogFileFormat = "PLAIN_TEXT"
	// File will be written in parquet format.
	FlowLogFileFormat_PARQUET FlowLogFileFormat = "PARQUET"
)

