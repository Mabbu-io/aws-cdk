package awscodecommit


// List of event types for AWS CodeCommit.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   repositoryNotifyOnOptions := &repositoryNotifyOnOptions{
//   	events: []repositoryNotificationEvents{
//   		awscdk.Aws_codecommit.*repositoryNotificationEvents_COMMIT_COMMENT,
//   	},
//
//   	// the properties below are optional
//   	detailType: awscdk.Aws_codestarnotifications.detailType_BASIC,
//   	enabled: jsii.Boolean(false),
//   	notificationRuleName: jsii.String("notificationRuleName"),
//   }
//
// See: https://docs.aws.amazon.com/dtconsole/latest/userguide/concepts.html#events-ref-repositories
//
type RepositoryNotificationEvents string

const (
	// Trigger notication when comment made on commit.
	RepositoryNotificationEvents_COMMIT_COMMENT RepositoryNotificationEvents = "COMMIT_COMMENT"
	// Trigger notification when comment made on pull request.
	RepositoryNotificationEvents_PULL_REQUEST_COMMENT RepositoryNotificationEvents = "PULL_REQUEST_COMMENT"
	// Trigger notification when approval status changed.
	RepositoryNotificationEvents_APPROVAL_STATUS_CHANGED RepositoryNotificationEvents = "APPROVAL_STATUS_CHANGED"
	// Trigger notifications when approval rule is overridden.
	RepositoryNotificationEvents_APPROVAL_RULE_OVERRIDDEN RepositoryNotificationEvents = "APPROVAL_RULE_OVERRIDDEN"
	// Trigger notification when pull request created.
	RepositoryNotificationEvents_PULL_REQUEST_CREATED RepositoryNotificationEvents = "PULL_REQUEST_CREATED"
	// Trigger notification when pull request source updated.
	RepositoryNotificationEvents_PULL_REQUEST_SOURCE_UPDATED RepositoryNotificationEvents = "PULL_REQUEST_SOURCE_UPDATED"
	// Trigger notification when pull request status is changed.
	RepositoryNotificationEvents_PULL_REQUEST_STATUS_CHANGED RepositoryNotificationEvents = "PULL_REQUEST_STATUS_CHANGED"
	// Trigger notification when pull requset is merged.
	RepositoryNotificationEvents_PULL_REQUEST_MERGED RepositoryNotificationEvents = "PULL_REQUEST_MERGED"
	// Trigger notification when a branch or tag is created.
	RepositoryNotificationEvents_BRANCH_OR_TAG_CREATED RepositoryNotificationEvents = "BRANCH_OR_TAG_CREATED"
	// Trigger notification when a branch or tag is deleted.
	RepositoryNotificationEvents_BRANCH_OR_TAG_DELETED RepositoryNotificationEvents = "BRANCH_OR_TAG_DELETED"
	// Trigger notification when a branch or tag is updated.
	RepositoryNotificationEvents_BRANCH_OR_TAG_UPDATED RepositoryNotificationEvents = "BRANCH_OR_TAG_UPDATED"
)

