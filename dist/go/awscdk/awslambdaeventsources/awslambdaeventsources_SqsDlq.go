package awslambdaeventsources

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/awslambdaeventsources/internal"
	"github.com/aws/aws-cdk-go/awscdk/awssqs"
)

// An SQS dead letter queue destination configuration for a Lambda event source.
//
// Example:
//   import dynamodb "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var table table
//
//   var fn function
//
//
//   deadLetterQueue := sqs.NewQueue(this, jsii.String("deadLetterQueue"))
//   fn.addEventSource(awscdk.NewDynamoEventSource(table, &dynamoEventSourceProps{
//   	startingPosition: lambda.startingPosition_TRIM_HORIZON,
//   	batchSize: jsii.Number(5),
//   	bisectBatchOnError: jsii.Boolean(true),
//   	onFailure: awscdk.NewSqsDlq(deadLetterQueue),
//   	retryAttempts: jsii.Number(10),
//   }))
//
type SqsDlq interface {
	awslambda.IEventSourceDlq
	// Returns a destination configuration for the DLQ.
	Bind(_target awslambda.IEventSourceMapping, targetHandler awslambda.IFunction) *awslambda.DlqDestinationConfig
}

// The jsii proxy struct for SqsDlq
type jsiiProxy_SqsDlq struct {
	internal.Type__awslambdaIEventSourceDlq
}

func NewSqsDlq(queue awssqs.IQueue) SqsDlq {
	_init_.Initialize()

	if err := validateNewSqsDlqParameters(queue); err != nil {
		panic(err)
	}
	j := jsiiProxy_SqsDlq{}

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda_event_sources.SqsDlq",
		[]interface{}{queue},
		&j,
	)

	return &j
}

func NewSqsDlq_Override(s SqsDlq, queue awssqs.IQueue) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda_event_sources.SqsDlq",
		[]interface{}{queue},
		s,
	)
}

func (s *jsiiProxy_SqsDlq) Bind(_target awslambda.IEventSourceMapping, targetHandler awslambda.IFunction) *awslambda.DlqDestinationConfig {
	if err := s.validateBindParameters(_target, targetHandler); err != nil {
		panic(err)
	}
	var returns *awslambda.DlqDestinationConfig

	_jsii_.Invoke(
		s,
		"bind",
		[]interface{}{_target, targetHandler},
		&returns,
	)

	return returns
}

