package awsec2


// The available destination types for Flow Logs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var bucket bucket
//   var logGroup logGroup
//   var role role
//
//   flowLogDestinationConfig := &flowLogDestinationConfig{
//   	logDestinationType: awscdk.Aws_ec2.flowLogDestinationType_CLOUD_WATCH_LOGS,
//
//   	// the properties below are optional
//   	destinationOptions: &destinationOptions{
//   		fileFormat: awscdk.*Aws_ec2.flowLogFileFormat_PLAIN_TEXT,
//   		hiveCompatiblePartitions: jsii.Boolean(false),
//   		perHourPartition: jsii.Boolean(false),
//   	},
//   	iamRole: role,
//   	keyPrefix: jsii.String("keyPrefix"),
//   	logGroup: logGroup,
//   	s3Bucket: bucket,
//   }
//
type FlowLogDestinationType string

const (
	// Send flow logs to CloudWatch Logs Group.
	FlowLogDestinationType_CLOUD_WATCH_LOGS FlowLogDestinationType = "CLOUD_WATCH_LOGS"
	// Send flow logs to S3 Bucket.
	FlowLogDestinationType_S3 FlowLogDestinationType = "S3"
)

