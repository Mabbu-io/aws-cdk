package awslogs

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for a LogStream.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var logGroup logGroup
//
//   logStreamProps := &logStreamProps{
//   	logGroup: logGroup,
//
//   	// the properties below are optional
//   	logStreamName: jsii.String("logStreamName"),
//   	removalPolicy: cdk.removalPolicy_DESTROY,
//   }
//
type LogStreamProps struct {
	// The log group to create a log stream for.
	LogGroup ILogGroup `field:"required" json:"logGroup" yaml:"logGroup"`
	// The name of the log stream to create.
	//
	// The name must be unique within the log group.
	LogStreamName *string `field:"optional" json:"logStreamName" yaml:"logStreamName"`
	// Determine what happens when the log stream resource is removed from the app.
	//
	// Normally you want to retain the log stream so you can diagnose issues from
	// logs even after a deployment that no longer includes the log stream.
	//
	// The date-based retention policy of your log group will age out the logs
	// after a certain time.
	RemovalPolicy awscdk.RemovalPolicy `field:"optional" json:"removalPolicy" yaml:"removalPolicy"`
}

