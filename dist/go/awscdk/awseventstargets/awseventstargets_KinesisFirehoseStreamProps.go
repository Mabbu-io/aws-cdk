package awseventstargets

import (
	"github.com/aws/aws-cdk-go/awscdk/awsevents"
)

// Customize the Firehose Stream Event Target.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var ruleTargetInput ruleTargetInput
//
//   kinesisFirehoseStreamProps := &kinesisFirehoseStreamProps{
//   	message: ruleTargetInput,
//   }
//
type KinesisFirehoseStreamProps struct {
	// The message to send to the stream.
	//
	// Must be a valid JSON text passed to the target stream.
	Message awsevents.RuleTargetInput `field:"optional" json:"message" yaml:"message"`
}

