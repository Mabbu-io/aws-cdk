package awss3


// Supported types of notification destinations.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import constructs "github.com/aws/constructs-go/constructs"
//
//   var dependable iDependable
//
//   bucketNotificationDestinationConfig := &bucketNotificationDestinationConfig{
//   	arn: jsii.String("arn"),
//   	type: awscdk.Aws_s3.bucketNotificationDestinationType_LAMBDA,
//
//   	// the properties below are optional
//   	dependencies: []*iDependable{
//   		dependable,
//   	},
//   }
//
type BucketNotificationDestinationType string

const (
	BucketNotificationDestinationType_LAMBDA BucketNotificationDestinationType = "LAMBDA"
	BucketNotificationDestinationType_QUEUE BucketNotificationDestinationType = "QUEUE"
	BucketNotificationDestinationType_TOPIC BucketNotificationDestinationType = "TOPIC"
)

