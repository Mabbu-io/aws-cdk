package awssqs

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnQueue`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var redriveAllowPolicy interface{}
//   var redrivePolicy interface{}
//
//   cfnQueueProps := &cfnQueueProps{
//   	contentBasedDeduplication: jsii.Boolean(false),
//   	deduplicationScope: jsii.String("deduplicationScope"),
//   	delaySeconds: jsii.Number(123),
//   	fifoQueue: jsii.Boolean(false),
//   	fifoThroughputLimit: jsii.String("fifoThroughputLimit"),
//   	kmsDataKeyReusePeriodSeconds: jsii.Number(123),
//   	kmsMasterKeyId: jsii.String("kmsMasterKeyId"),
//   	maximumMessageSize: jsii.Number(123),
//   	messageRetentionPeriod: jsii.Number(123),
//   	queueName: jsii.String("queueName"),
//   	receiveMessageWaitTimeSeconds: jsii.Number(123),
//   	redriveAllowPolicy: redriveAllowPolicy,
//   	redrivePolicy: redrivePolicy,
//   	sqsManagedSseEnabled: jsii.Boolean(false),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	visibilityTimeout: jsii.Number(123),
//   }
//
type CfnQueueProps struct {
	// For first-in-first-out (FIFO) queues, specifies whether to enable content-based deduplication.
	//
	// During the deduplication interval, Amazon SQS treats messages that are sent with identical content as duplicates and delivers only one copy of the message. For more information, see the `ContentBasedDeduplication` attribute for the `[CreateQueue](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/APIReference/API_CreateQueue.html)` action in the *Amazon SQS API Reference* .
	ContentBasedDeduplication interface{} `field:"optional" json:"contentBasedDeduplication" yaml:"contentBasedDeduplication"`
	// For high throughput for FIFO queues, specifies whether message deduplication occurs at the message group or queue level.
	//
	// Valid values are `messageGroup` and `queue` .
	//
	// To enable high throughput for a FIFO queue, set this attribute to `messageGroup` *and* set the `FifoThroughputLimit` attribute to `perMessageGroupId` . If you set these attributes to anything other than these values, normal throughput is in effect and deduplication occurs as specified. For more information, see [High throughput for FIFO queues](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/high-throughput-fifo.html) and [Quotas related to messages](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/quotas-messages.html) in the *Amazon SQS Developer Guide* .
	DeduplicationScope *string `field:"optional" json:"deduplicationScope" yaml:"deduplicationScope"`
	// The time in seconds for which the delivery of all messages in the queue is delayed.
	//
	// You can specify an integer value of `0` to `900` (15 minutes). The default value is `0` .
	DelaySeconds *float64 `field:"optional" json:"delaySeconds" yaml:"delaySeconds"`
	// If set to true, creates a FIFO queue.
	//
	// If you don't specify this property, Amazon SQS creates a standard queue. For more information, see [FIFO queues](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/FIFO-queues.html) in the *Amazon SQS Developer Guide* .
	FifoQueue interface{} `field:"optional" json:"fifoQueue" yaml:"fifoQueue"`
	// For high throughput for FIFO queues, specifies whether the FIFO queue throughput quota applies to the entire queue or per message group.
	//
	// Valid values are `perQueue` and `perMessageGroupId` .
	//
	// To enable high throughput for a FIFO queue, set this attribute to `perMessageGroupId` *and* set the `DeduplicationScope` attribute to `messageGroup` . If you set these attributes to anything other than these values, normal throughput is in effect and deduplication occurs as specified. For more information, see [High throughput for FIFO queues](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/high-throughput-fifo.html) and [Quotas related to messages](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/quotas-messages.html) in the *Amazon SQS Developer Guide* .
	FifoThroughputLimit *string `field:"optional" json:"fifoThroughputLimit" yaml:"fifoThroughputLimit"`
	// The length of time in seconds for which Amazon SQS can reuse a data key to encrypt or decrypt messages before calling AWS KMS again.
	//
	// The value must be an integer between 60 (1 minute) and 86,400 (24 hours). The default is 300 (5 minutes).
	//
	// > A shorter time period provides better security, but results in more calls to AWS KMS , which might incur charges after Free Tier. For more information, see [Encryption at rest](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-server-side-encryption.html#sqs-how-does-the-data-key-reuse-period-work) in the *Amazon SQS Developer Guide* .
	KmsDataKeyReusePeriodSeconds *float64 `field:"optional" json:"kmsDataKeyReusePeriodSeconds" yaml:"kmsDataKeyReusePeriodSeconds"`
	// The ID of an AWS managed customer master key (CMK) for Amazon SQS or a custom CMK.
	//
	// To use the AWS managed CMK for Amazon SQS , specify the (default) alias `alias/aws/sqs` . For more information, see the following:
	//
	// - [Encryption at rest](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-server-side-encryption.html) in the *Amazon SQS Developer Guide*
	// - [CreateQueue](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/APIReference/API_CreateQueue.html) in the *Amazon SQS API Reference*
	// - The Customer Master Keys section of the [AWS Key Management Service Best Practices](https://docs.aws.amazon.com/https://d0.awsstatic.com/whitepapers/aws-kms-best-practices.pdf) whitepaper
	KmsMasterKeyId *string `field:"optional" json:"kmsMasterKeyId" yaml:"kmsMasterKeyId"`
	// The limit of how many bytes that a message can contain before Amazon SQS rejects it.
	//
	// You can specify an integer value from `1,024` bytes (1 KiB) to `262,144` bytes (256 KiB). The default value is `262,144` (256 KiB).
	MaximumMessageSize *float64 `field:"optional" json:"maximumMessageSize" yaml:"maximumMessageSize"`
	// The number of seconds that Amazon SQS retains a message.
	//
	// You can specify an integer value from `60` seconds (1 minute) to `1,209,600` seconds (14 days). The default value is `345,600` seconds (4 days).
	MessageRetentionPeriod *float64 `field:"optional" json:"messageRetentionPeriod" yaml:"messageRetentionPeriod"`
	// A name for the queue.
	//
	// To create a FIFO queue, the name of your FIFO queue must end with the `.fifo` suffix. For more information, see [FIFO queues](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/FIFO-queues.html) in the *Amazon SQS Developer Guide* .
	//
	// If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the queue name. For more information, see [Name type](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html) in the *AWS CloudFormation User Guide* .
	//
	// > If you specify a name, you can't perform updates that require replacement of this resource. You can perform updates that require no or some interruption. If you must replace the resource, specify a new name.
	QueueName *string `field:"optional" json:"queueName" yaml:"queueName"`
	// Specifies the duration, in seconds, that the ReceiveMessage action call waits until a message is in the queue in order to include it in the response, rather than returning an empty response if a message isn't yet available.
	//
	// You can specify an integer from 1 to 20. Short polling is used as the default or when you specify 0 for this property. For more information, see [Consuming messages using long polling](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-short-and-long-polling.html#sqs-long-polling) in the *Amazon SQS Developer Guide* .
	ReceiveMessageWaitTimeSeconds *float64 `field:"optional" json:"receiveMessageWaitTimeSeconds" yaml:"receiveMessageWaitTimeSeconds"`
	// The string that includes the parameters for the permissions for the dead-letter queue redrive permission and which source queues can specify dead-letter queues as a JSON object.
	//
	// The parameters are as follows:
	//
	// - `redrivePermission` : The permission type that defines which source queues can specify the current queue as the dead-letter queue. Valid values are:
	//
	// - `allowAll` : (Default) Any source queues in this AWS account in the same Region can specify this queue as the dead-letter queue.
	// - `denyAll` : No source queues can specify this queue as the dead-letter queue.
	// - `byQueue` : Only queues specified by the `sourceQueueArns` parameter can specify this queue as the dead-letter queue.
	// - `sourceQueueArns` : The Amazon Resource Names (ARN)s of the source queues that can specify this queue as the dead-letter queue and redrive messages. You can specify this parameter only when the `redrivePermission` parameter is set to `byQueue` . You can specify up to 10 source queue ARNs. To allow more than 10 source queues to specify dead-letter queues, set the `redrivePermission` parameter to `allowAll` .
	RedriveAllowPolicy interface{} `field:"optional" json:"redriveAllowPolicy" yaml:"redriveAllowPolicy"`
	// The string that includes the parameters for the dead-letter queue functionality of the source queue as a JSON object.
	//
	// The parameters are as follows:
	//
	// - `deadLetterTargetArn` : The Amazon Resource Name (ARN) of the dead-letter queue to which Amazon SQS moves messages after the value of `maxReceiveCount` is exceeded.
	// - `maxReceiveCount` : The number of times a message is delivered to the source queue before being moved to the dead-letter queue. When the `ReceiveCount` for a message exceeds the `maxReceiveCount` for a queue, Amazon SQS moves the message to the dead-letter-queue.
	//
	// > The dead-letter queue of a FIFO queue must also be a FIFO queue. Similarly, the dead-letter queue of a standard queue must also be a standard queue.
	//
	// *JSON*
	//
	// `{ "deadLetterTargetArn" : *String* , "maxReceiveCount" : *Integer* }`
	//
	// *YAML*
	//
	// `deadLetterTargetArn : *String*`
	//
	// `maxReceiveCount : *Integer*`.
	RedrivePolicy interface{} `field:"optional" json:"redrivePolicy" yaml:"redrivePolicy"`
	// `AWS::SQS::Queue.SqsManagedSseEnabled`.
	SqsManagedSseEnabled interface{} `field:"optional" json:"sqsManagedSseEnabled" yaml:"sqsManagedSseEnabled"`
	// The tags that you attach to this queue.
	//
	// For more information, see [Resource tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) in the *AWS CloudFormation User Guide* .
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The length of time during which a message will be unavailable after a message is delivered from the queue.
	//
	// This blocks other components from receiving the same message and gives the initial component time to process and delete the message from the queue.
	//
	// Values must be from 0 to 43,200 seconds (12 hours). If you don't specify a value, AWS CloudFormation uses the default value of 30 seconds.
	//
	// For more information about Amazon SQS queue visibility timeouts, see [Visibility timeout](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-visibility-timeout.html) in the *Amazon SQS Developer Guide* .
	VisibilityTimeout *float64 `field:"optional" json:"visibilityTimeout" yaml:"visibilityTimeout"`
}

