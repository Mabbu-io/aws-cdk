package awss3notifications

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/awss3"
	"github.com/aws/aws-cdk-go/awscdk/awss3notifications/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Use a Lambda function as a bucket notification destination.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   var myLambda function
//
//   bucket := s3.bucket.fromBucketAttributes(this, jsii.String("ImportedBucket"), &bucketAttributes{
//   	bucketArn: jsii.String("arn:aws:s3:::my-bucket"),
//   })
//
//   // now you can just call methods on the bucket
//   bucket.addEventNotification(s3.eventType_OBJECT_CREATED, s3n.NewLambdaDestination(myLambda), &notificationKeyFilter{
//   	prefix: jsii.String("home/myusername/*"),
//   })
//
type LambdaDestination interface {
	awss3.IBucketNotificationDestination
	// Registers this resource to receive notifications for the specified bucket.
	//
	// This method will only be called once for each destination/bucket
	// pair and the result will be cached, so there is no need to implement
	// idempotency in each destination.
	Bind(_scope constructs.Construct, bucket awss3.IBucket) *awss3.BucketNotificationDestinationConfig
}

// The jsii proxy struct for LambdaDestination
type jsiiProxy_LambdaDestination struct {
	internal.Type__awss3IBucketNotificationDestination
}

func NewLambdaDestination(fn awslambda.IFunction) LambdaDestination {
	_init_.Initialize()

	if err := validateNewLambdaDestinationParameters(fn); err != nil {
		panic(err)
	}
	j := jsiiProxy_LambdaDestination{}

	_jsii_.Create(
		"aws-cdk-lib.aws_s3_notifications.LambdaDestination",
		[]interface{}{fn},
		&j,
	)

	return &j
}

func NewLambdaDestination_Override(l LambdaDestination, fn awslambda.IFunction) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_s3_notifications.LambdaDestination",
		[]interface{}{fn},
		l,
	)
}

func (l *jsiiProxy_LambdaDestination) Bind(_scope constructs.Construct, bucket awss3.IBucket) *awss3.BucketNotificationDestinationConfig {
	if err := l.validateBindParameters(_scope, bucket); err != nil {
		panic(err)
	}
	var returns *awss3.BucketNotificationDestinationConfig

	_jsii_.Invoke(
		l,
		"bind",
		[]interface{}{_scope, bucket},
		&returns,
	)

	return returns
}

