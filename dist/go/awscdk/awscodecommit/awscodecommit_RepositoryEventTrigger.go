package awscodecommit


// Repository events that will cause the trigger to run actions in another service.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   repositoryTriggerOptions := &repositoryTriggerOptions{
//   	branches: []*string{
//   		jsii.String("branches"),
//   	},
//   	customData: jsii.String("customData"),
//   	events: []repositoryEventTrigger{
//   		awscdk.Aws_codecommit.*repositoryEventTrigger_ALL,
//   	},
//   	name: jsii.String("name"),
//   }
//
type RepositoryEventTrigger string

const (
	RepositoryEventTrigger_ALL RepositoryEventTrigger = "ALL"
	RepositoryEventTrigger_UPDATE_REF RepositoryEventTrigger = "UPDATE_REF"
	RepositoryEventTrigger_CREATE_REF RepositoryEventTrigger = "CREATE_REF"
	RepositoryEventTrigger_DELETE_REF RepositoryEventTrigger = "DELETE_REF"
)

