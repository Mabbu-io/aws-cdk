package pipelines

import (
	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awssns"
)

// Properties for a `PermissionsBroadeningCheck`.
//
// Example:
//   var pipeline codePipeline
//
//   stage := NewMyApplicationStage(this, jsii.String("MyApplication"))
//   pipeline.addStage(stage, &addStageOpts{
//   	pre: []step{
//   		pipelines.NewConfirmPermissionsBroadening(jsii.String("Check"), &permissionsBroadeningCheckProps{
//   			stage: stage,
//   		}),
//   	},
//   })
//
type PermissionsBroadeningCheckProps struct {
	// The CDK Stage object to check the stacks of.
	//
	// This should be the same Stage object you are passing to `addStage()`.
	Stage awscdk.Stage `field:"required" json:"stage" yaml:"stage"`
	// Topic to send notifications when a human needs to give manual confirmation.
	NotificationTopic awssns.ITopic `field:"optional" json:"notificationTopic" yaml:"notificationTopic"`
}

