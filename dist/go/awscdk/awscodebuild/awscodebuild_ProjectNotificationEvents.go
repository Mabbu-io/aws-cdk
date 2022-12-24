package awscodebuild


// The list of event types for AWS Codebuild.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   projectNotifyOnOptions := &projectNotifyOnOptions{
//   	events: []projectNotificationEvents{
//   		awscdk.Aws_codebuild.*projectNotificationEvents_BUILD_FAILED,
//   	},
//
//   	// the properties below are optional
//   	detailType: awscdk.Aws_codestarnotifications.detailType_BASIC,
//   	enabled: jsii.Boolean(false),
//   	notificationRuleName: jsii.String("notificationRuleName"),
//   }
//
// See: https://docs.aws.amazon.com/dtconsole/latest/userguide/concepts.html#events-ref-buildproject
//
type ProjectNotificationEvents string

const (
	// Trigger notification when project build state failed.
	ProjectNotificationEvents_BUILD_FAILED ProjectNotificationEvents = "BUILD_FAILED"
	// Trigger notification when project build state succeeded.
	ProjectNotificationEvents_BUILD_SUCCEEDED ProjectNotificationEvents = "BUILD_SUCCEEDED"
	// Trigger notification when project build state in progress.
	ProjectNotificationEvents_BUILD_IN_PROGRESS ProjectNotificationEvents = "BUILD_IN_PROGRESS"
	// Trigger notification when project build state stopped.
	ProjectNotificationEvents_BUILD_STOPPED ProjectNotificationEvents = "BUILD_STOPPED"
	// Trigger notification when project build phase failure.
	ProjectNotificationEvents_BUILD_PHASE_FAILED ProjectNotificationEvents = "BUILD_PHASE_FAILED"
	// Trigger notification when project build phase success.
	ProjectNotificationEvents_BUILD_PHASE_SUCCEEDED ProjectNotificationEvents = "BUILD_PHASE_SUCCEEDED"
)

