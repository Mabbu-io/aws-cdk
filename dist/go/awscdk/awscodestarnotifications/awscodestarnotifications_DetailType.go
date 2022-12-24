package awscodestarnotifications


// The level of detail to include in the notifications for this resource.
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
type DetailType string

const (
	// BASIC will include only the contents of the event as it would appear in AWS CloudWatch.
	DetailType_BASIC DetailType = "BASIC"
	// FULL will include any supplemental information provided by AWS CodeStar Notifications and/or the service for the resource for which the notification is created.
	DetailType_FULL DetailType = "FULL"
)

