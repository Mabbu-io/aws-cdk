package awscodepipelineactions


// Describes whether AWS CloudFormation StackSets automatically deploys to AWS Organizations accounts that are added to a target organization or organizational unit (OU).
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   organizationsDeploymentProps := &organizationsDeploymentProps{
//   	autoDeployment: awscdk.Aws_codepipeline_actions.stackSetOrganizationsAutoDeployment_ENABLED,
//   }
//
type StackSetOrganizationsAutoDeployment string

const (
	// StackSets automatically deploys additional stack instances to AWS Organizations accounts that are added to a target organization or organizational unit (OU) in the specified Regions.
	//
	// If an account is removed from a target organization or OU, AWS CloudFormation StackSets
	// deletes stack instances from the account in the specified Regions.
	StackSetOrganizationsAutoDeployment_ENABLED StackSetOrganizationsAutoDeployment = "ENABLED"
	// StackSets does not automatically deploy additional stack instances to AWS Organizations accounts that are added to a target organization or organizational unit (OU) in the specified Regions.
	StackSetOrganizationsAutoDeployment_DISABLED StackSetOrganizationsAutoDeployment = "DISABLED"
	// Stack resources are retained when an account is removed from a target organization or OU.
	StackSetOrganizationsAutoDeployment_ENABLED_WITH_STACK_RETENTION StackSetOrganizationsAutoDeployment = "ENABLED_WITH_STACK_RETENTION"
)

