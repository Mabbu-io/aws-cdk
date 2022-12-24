package awslogs

import (
	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awskms"
)

// Properties for a LogGroup.
//
// Example:
//   import logs "github.com/aws/aws-cdk-go/awscdk"
//
//
//   logGroup := logs.NewLogGroup(this, jsii.String("MyLogGroup"), &logGroupProps{
//   	logGroupName: jsii.String("MyLogGroup"),
//   })
//
//   rule := events.NewRule(this, jsii.String("rule"), &ruleProps{
//   	eventPattern: &eventPattern{
//   		source: []*string{
//   			jsii.String("aws.ec2"),
//   		},
//   	},
//   })
//
//   rule.addTarget(targets.NewCloudWatchLogGroup(logGroup))
//
type LogGroupProps struct {
	// The KMS customer managed key to encrypt the log group with.
	EncryptionKey awskms.IKey `field:"optional" json:"encryptionKey" yaml:"encryptionKey"`
	// Name of the log group.
	LogGroupName *string `field:"optional" json:"logGroupName" yaml:"logGroupName"`
	// Determine the removal policy of this log group.
	//
	// Normally you want to retain the log group so you can diagnose issues
	// from logs even after a deployment that no longer includes the log group.
	// In that case, use the normal date-based retention policy to age out your
	// logs.
	RemovalPolicy awscdk.RemovalPolicy `field:"optional" json:"removalPolicy" yaml:"removalPolicy"`
	// How long, in days, the log contents will be retained.
	//
	// To retain all logs, set this value to RetentionDays.INFINITE.
	Retention RetentionDays `field:"optional" json:"retention" yaml:"retention"`
}

