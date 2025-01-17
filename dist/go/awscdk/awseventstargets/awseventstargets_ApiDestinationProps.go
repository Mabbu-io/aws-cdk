package awseventstargets

import (
	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/awssqs"
)

// Customize the EventBridge Api Destinations Target.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var queue queue
//   var role role
//   var ruleTargetInput ruleTargetInput
//
//   apiDestinationProps := &apiDestinationProps{
//   	deadLetterQueue: queue,
//   	event: ruleTargetInput,
//   	eventRole: role,
//   	headerParameters: map[string]*string{
//   		"headerParametersKey": jsii.String("headerParameters"),
//   	},
//   	maxEventAge: cdk.duration.minutes(jsii.Number(30)),
//   	pathParameterValues: []*string{
//   		jsii.String("pathParameterValues"),
//   	},
//   	queryStringParameters: map[string]*string{
//   		"queryStringParametersKey": jsii.String("queryStringParameters"),
//   	},
//   	retryAttempts: jsii.Number(123),
//   }
//
type ApiDestinationProps struct {
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
	// The event to send.
	Event awsevents.RuleTargetInput `field:"optional" json:"event" yaml:"event"`
	// The role to assume before invoking the target.
	EventRole awsiam.IRole `field:"optional" json:"eventRole" yaml:"eventRole"`
	// Additional headers sent to the API Destination.
	//
	// These are merged with headers specified on the Connection, with
	// the headers on the Connection taking precedence.
	//
	// You can only specify secret values on the Connection.
	HeaderParameters *map[string]*string `field:"optional" json:"headerParameters" yaml:"headerParameters"`
	// Path parameters to insert in place of path wildcards (`*`).
	//
	// If the API destination has a wilcard in the path, these path parts
	// will be inserted in that place.
	PathParameterValues *[]*string `field:"optional" json:"pathParameterValues" yaml:"pathParameterValues"`
	// Additional query string parameters sent to the API Destination.
	//
	// These are merged with headers specified on the Connection, with
	// the headers on the Connection taking precedence.
	//
	// You can only specify secret values on the Connection.
	QueryStringParameters *map[string]*string `field:"optional" json:"queryStringParameters" yaml:"queryStringParameters"`
}

