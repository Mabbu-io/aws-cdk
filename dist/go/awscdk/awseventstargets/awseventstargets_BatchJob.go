package awseventstargets

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/awseventstargets/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Use an AWS Batch Job / Queue as an event rule target.
//
// Most likely the code will look something like this:
// `new BatchJob(jobQueue.jobQueueArn, jobQueue, jobDefinition.jobDefinitionArn, jobDefinition)`
//
// In the future this API will be improved to be fully typed.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   import batch "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//
//   jobQueue := batch.NewJobQueue(this, jsii.String("MyQueue"), map[string][]map[string]interface{}{
//   	"computeEnvironments": []map[string]interface{}{
//   		map[string]interface{}{
//   			"computeEnvironment": batch.NewComputeEnvironment(this, jsii.String("ComputeEnvironment"), map[string]*bool{
//   				"managed": jsii.Boolean(false),
//   			}),
//   			"order": jsii.Number(1),
//   		},
//   	},
//   })
//
//   jobDefinition := batch.NewJobDefinition(this, jsii.String("MyJob"), map[string]map[string]repositoryImage{
//   	"container": map[string]repositoryImage{
//   		"image": awscdk.ContainerImage.fromRegistry(jsii.String("test-repo")),
//   	},
//   })
//
//   queue := sqs.NewQueue(this, jsii.String("Queue"))
//
//   rule := events.NewRule(this, jsii.String("Rule"), &ruleProps{
//   	schedule: events.schedule.rate(cdk.duration.hours(jsii.Number(1))),
//   })
//
//   rule.addTarget(targets.NewBatchJob(jobQueue.jobQueueArn, jobQueue, jobDefinition.jobDefinitionArn, jobDefinition, &batchJobProps{
//   	deadLetterQueue: queue,
//   	event: events.ruleTargetInput.fromObject(map[string]*string{
//   		"SomeParam": jsii.String("SomeValue"),
//   	}),
//   	retryAttempts: jsii.Number(2),
//   	maxEventAge: cdk.*duration.hours(jsii.Number(2)),
//   }))
//
type BatchJob interface {
	awsevents.IRuleTarget
	// Returns a RuleTarget that can be used to trigger queue this batch job as a result from an EventBridge event.
	Bind(rule awsevents.IRule, _id *string) *awsevents.RuleTargetConfig
}

// The jsii proxy struct for BatchJob
type jsiiProxy_BatchJob struct {
	internal.Type__awseventsIRuleTarget
}

func NewBatchJob(jobQueueArn *string, jobQueueScope constructs.IConstruct, jobDefinitionArn *string, jobDefinitionScope constructs.IConstruct, props *BatchJobProps) BatchJob {
	_init_.Initialize()

	if err := validateNewBatchJobParameters(jobQueueArn, jobQueueScope, jobDefinitionArn, jobDefinitionScope, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_BatchJob{}

	_jsii_.Create(
		"aws-cdk-lib.aws_events_targets.BatchJob",
		[]interface{}{jobQueueArn, jobQueueScope, jobDefinitionArn, jobDefinitionScope, props},
		&j,
	)

	return &j
}

func NewBatchJob_Override(b BatchJob, jobQueueArn *string, jobQueueScope constructs.IConstruct, jobDefinitionArn *string, jobDefinitionScope constructs.IConstruct, props *BatchJobProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_events_targets.BatchJob",
		[]interface{}{jobQueueArn, jobQueueScope, jobDefinitionArn, jobDefinitionScope, props},
		b,
	)
}

func (b *jsiiProxy_BatchJob) Bind(rule awsevents.IRule, _id *string) *awsevents.RuleTargetConfig {
	if err := b.validateBindParameters(rule); err != nil {
		panic(err)
	}
	var returns *awsevents.RuleTargetConfig

	_jsii_.Invoke(
		b,
		"bind",
		[]interface{}{rule, _id},
		&returns,
	)

	return returns
}

