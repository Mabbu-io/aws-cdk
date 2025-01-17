package awslambdaeventsources

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/awsdynamodb"
	"github.com/aws/aws-cdk-go/awscdk/awslambda"
)

// Use an Amazon DynamoDB stream as an event source for AWS Lambda.
//
// Example:
//   import eventsources "github.com/aws/aws-cdk-go/awscdk"
//   import dynamodb "github.com/aws/aws-cdk-go/awscdk"
//
//   var fn function
//
//   table := dynamodb.NewTable(this, jsii.String("Table"), &tableProps{
//   	partitionKey: &attribute{
//   		name: jsii.String("id"),
//   		type: dynamodb.attributeType_STRING,
//   	},
//   	stream: dynamodb.streamViewType_NEW_IMAGE,
//   })
//   fn.addEventSource(eventsources.NewDynamoEventSource(table, &dynamoEventSourceProps{
//   	startingPosition: lambda.startingPosition_LATEST,
//   	filters: []map[string]interface{}{
//   		lambda.filterCriteria.filter(map[string]interface{}{
//   			"eventName": lambda.FilterRule.isEqual(jsii.String("INSERT")),
//   		}),
//   	},
//   }))
//
type DynamoEventSource interface {
	StreamEventSource
	// The identifier for this EventSourceMapping.
	EventSourceMappingId() *string
	Props() *StreamEventSourceProps
	// Called by `lambda.addEventSource` to allow the event source to bind to this function.
	Bind(target awslambda.IFunction)
	EnrichMappingOptions(options *awslambda.EventSourceMappingOptions) *awslambda.EventSourceMappingOptions
}

// The jsii proxy struct for DynamoEventSource
type jsiiProxy_DynamoEventSource struct {
	jsiiProxy_StreamEventSource
}

func (j *jsiiProxy_DynamoEventSource) EventSourceMappingId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventSourceMappingId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamoEventSource) Props() *StreamEventSourceProps {
	var returns *StreamEventSourceProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}


func NewDynamoEventSource(table awsdynamodb.ITable, props *DynamoEventSourceProps) DynamoEventSource {
	_init_.Initialize()

	if err := validateNewDynamoEventSourceParameters(table, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_DynamoEventSource{}

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda_event_sources.DynamoEventSource",
		[]interface{}{table, props},
		&j,
	)

	return &j
}

func NewDynamoEventSource_Override(d DynamoEventSource, table awsdynamodb.ITable, props *DynamoEventSourceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda_event_sources.DynamoEventSource",
		[]interface{}{table, props},
		d,
	)
}

func (d *jsiiProxy_DynamoEventSource) Bind(target awslambda.IFunction) {
	if err := d.validateBindParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"bind",
		[]interface{}{target},
	)
}

func (d *jsiiProxy_DynamoEventSource) EnrichMappingOptions(options *awslambda.EventSourceMappingOptions) *awslambda.EventSourceMappingOptions {
	if err := d.validateEnrichMappingOptionsParameters(options); err != nil {
		panic(err)
	}
	var returns *awslambda.EventSourceMappingOptions

	_jsii_.Invoke(
		d,
		"enrichMappingOptions",
		[]interface{}{options},
		&returns,
	)

	return returns
}

