package awseventstargets

import (
	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awssqs"
)

// The generic properties for an RuleTarget.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var queue queue
//
//   targetBaseProps := &targetBaseProps{
//   	deadLetterQueue: queue,
//   	maxEventAge: cdk.duration.minutes(jsii.Number(30)),
//   	retryAttempts: jsii.Number(123),
//   }
//
type TargetBaseProps struct {
	// The SQS queue to be used as deadLetterQueue. Check out the [considerations for using a dead-letter queue](https://docs.aws.amazon.com/eventbridge/latest/userguide/rule-dlq.html#dlq-considerations).
	//
	// The events not successfully delivered are automatically retried for a specified period of time,
	// depending on the retry policy of the target.
	// If an event is not delivered before all retry attempts are exhausted, it will be sent to the dead letter queue.
	DeadLetterQueue awssqs.IQueue `field:"optional" json:"deadLetterQueue" yaml:"deadLetterQueue"`
	// The maximum age of a request that Lambda sends to a function for processing.
	//
	// Minimum value of 60.
	// Maximum value of 86400.
	MaxEventAge awscdk.Duration `field:"optional" json:"maxEventAge" yaml:"maxEventAge"`
	// The maximum number of times to retry when the function returns an error.
	//
	// Minimum value of 0.
	// Maximum value of 185.
	RetryAttempts *float64 `field:"optional" json:"retryAttempts" yaml:"retryAttempts"`
}

