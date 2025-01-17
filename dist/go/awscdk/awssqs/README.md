# Amazon Simple Queue Service Construct Library

Amazon Simple Queue Service (SQS) is a fully managed message queuing service that
enables you to decouple and scale microservices, distributed systems, and serverless
applications. SQS eliminates the complexity and overhead associated with managing and
operating message oriented middleware, and empowers developers to focus on differentiating work.
Using SQS, you can send, store, and receive messages between software components at any volume,
without losing messages or requiring other services to be available.

## Installation

Import to your project:

```go
import sqs "github.com/aws/aws-cdk-go/awscdk"
```

## Basic usage

Here's how to add a basic queue to your application:

```go
sqs.NewQueue(this, jsii.String("Queue"))
```

## Encryption

By default queues are encrypted using SSE-SQS. If you want to change the encryption mode, set the `encryption` property.
The following encryption modes are supported:

* KMS key that SQS manages for you
* KMS key that you can managed yourself
* Server-side encryption managed by SQS (SSE-SQS)
* Unencrypted

To learn more about SSE-SQS on Amazon SQS, please visit the
[Amazon SQS documentation](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-server-side-encryption.html).

```go
// Use managed key
// Use managed key
sqs.NewQueue(this, jsii.String("Queue"), &queueProps{
	encryption: sqs.queueEncryption_KMS_MANAGED,
})

// Use custom key
myKey := kms.NewKey(this, jsii.String("Key"))

sqs.NewQueue(this, jsii.String("Queue"), &queueProps{
	encryption: sqs.*queueEncryption_KMS,
	encryptionMasterKey: myKey,
})

// Use SQS managed server side encryption (SSE-SQS)
// Use SQS managed server side encryption (SSE-SQS)
sqs.NewQueue(this, jsii.String("Queue"), &queueProps{
	encryption: sqs.*queueEncryption_SQS_MANAGED,
})

// Unencrypted queue
// Unencrypted queue
sqs.NewQueue(this, jsii.String("Queue"), &queueProps{
	encryption: sqs.*queueEncryption_UNENCRYPTED,
})
```

## Encryption in transit

If you want to enforce encryption of data in transit, set the `enforceSSL` property to `true`.
A resource policy statement that allows only encrypted connections over HTTPS (TLS)
will be added to the queue.

```go
sqs.NewQueue(this, jsii.String("Queue"), &queueProps{
	enforceSSL: jsii.Boolean(true),
})
```

## First-In-First-Out (FIFO) queues

FIFO queues give guarantees on the order in which messages are dequeued, and have additional
features in order to help guarantee exactly-once processing. For more information, see
the SQS manual. Note that FIFO queues are not available in all AWS regions.

A queue can be made a FIFO queue by either setting `fifo: true`, giving it a name which ends
in `".fifo"`, or by enabling a FIFO specific feature such as: content-based deduplication,
deduplication scope or fifo throughput limit.
