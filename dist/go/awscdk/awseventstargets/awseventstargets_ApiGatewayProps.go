package awseventstargets

import (
	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/awssqs"
)

// Customize the API Gateway Event Target.
//
// Example:
//   import api "github.com/aws/aws-cdk-go/awscdk"
//   import lambda "github.com/aws/aws-cdk-go/awscdk"
//
//
//   rule := events.NewRule(this, jsii.String("Rule"), &ruleProps{
//   	schedule: events.schedule.rate(cdk.duration.minutes(jsii.Number(1))),
//   })
//
//   fn := lambda.NewFunction(this, jsii.String("MyFunc"), &functionProps{
//   	handler: jsii.String("index.handler"),
//   	runtime: lambda.runtime_NODEJS_14_X(),
//   	code: lambda.code.fromInline(jsii.String("exports.handler = e => {}")),
//   })
//
//   restApi := api.NewLambdaRestApi(this, jsii.String("MyRestAPI"), &lambdaRestApiProps{
//   	handler: fn,
//   })
//
//   dlq := sqs.NewQueue(this, jsii.String("DeadLetterQueue"))
//
//   rule.addTarget(
//   targets.NewApiGateway(restApi, &apiGatewayProps{
//   	path: jsii.String("/*/test"),
//   	method: jsii.String("GET"),
//   	stage: jsii.String("prod"),
//   	pathParameterValues: []*string{
//   		jsii.String("path-value"),
//   	},
//   	headerParameters: map[string]*string{
//   		"Header1": jsii.String("header1"),
//   	},
//   	queryStringParameters: map[string]*string{
//   		"QueryParam1": jsii.String("query-param-1"),
//   	},
//   	deadLetterQueue: dlq,
//   }))
//
type ApiGatewayProps struct {
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
	// The role to assume before invoking the target (i.e., the pipeline) when the given rule is triggered.
	EventRole awsiam.IRole `field:"optional" json:"eventRole" yaml:"eventRole"`
	// The headers to be set when requesting API.
	HeaderParameters *map[string]*string `field:"optional" json:"headerParameters" yaml:"headerParameters"`
	// The method for api resource invoked by the rule.
	Method *string `field:"optional" json:"method" yaml:"method"`
	// The api resource invoked by the rule.
	//
	// We can use wildcards('*') to specify the path. In that case,
	// an equal number of real values must be specified for pathParameterValues.
	Path *string `field:"optional" json:"path" yaml:"path"`
	// The path parameter values to be used to populate to wildcards("*") of requesting api path.
	PathParameterValues *[]*string `field:"optional" json:"pathParameterValues" yaml:"pathParameterValues"`
	// This will be the post request body send to the API.
	PostBody awsevents.RuleTargetInput `field:"optional" json:"postBody" yaml:"postBody"`
	// The query parameters to be set when requesting API.
	QueryStringParameters *map[string]*string `field:"optional" json:"queryStringParameters" yaml:"queryStringParameters"`
	// The deploy stage of api gateway invoked by the rule.
	Stage *string `field:"optional" json:"stage" yaml:"stage"`
}

