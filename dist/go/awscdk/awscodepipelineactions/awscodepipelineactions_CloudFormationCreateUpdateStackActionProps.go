package awscodepipelineactions

import (
	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awscodepipeline"
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
)

// Properties for the CloudFormationCreateUpdateStackAction.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   // in stack for account 123456789012...
//   var otherAccountStack stack
//
//   actionRole := iam.NewRole(otherAccountStack, jsii.String("ActionRole"), &roleProps{
//   	assumedBy: iam.NewAccountPrincipal(jsii.String("123456789012")),
//   	// the role has to have a physical name set
//   	roleName: awscdk.PhysicalName_GENERATE_IF_NEEDED(),
//   })
//
//   // in the pipeline stack...
//   sourceOutput := codepipeline.NewArtifact()
//   codepipeline_actions.NewCloudFormationCreateUpdateStackAction(&cloudFormationCreateUpdateStackActionProps{
//   	actionName: jsii.String("CloudFormationCreateUpdate"),
//   	stackName: jsii.String("MyStackName"),
//   	adminPermissions: jsii.Boolean(true),
//   	templatePath: sourceOutput.atPath(jsii.String("template.yaml")),
//   	role: actionRole,
//   })
//
type CloudFormationCreateUpdateStackActionProps struct {
	// The physical, human-readable name of the Action.
	//
	// Note that Action names must be unique within a single Stage.
	ActionName *string `field:"required" json:"actionName" yaml:"actionName"`
	// The runOrder property for this Action.
	//
	// RunOrder determines the relative order in which multiple Actions in the same Stage execute.
	// See: https://docs.aws.amazon.com/codepipeline/latest/userguide/reference-pipeline-structure.html
	//
	RunOrder *float64 `field:"optional" json:"runOrder" yaml:"runOrder"`
	// The name of the namespace to use for variables emitted by this action.
	VariablesNamespace *string `field:"optional" json:"variablesNamespace" yaml:"variablesNamespace"`
	// The Role in which context's this Action will be executing in.
	//
	// The Pipeline's Role will assume this Role
	// (the required permissions for that will be granted automatically)
	// right before executing this Action.
	// This Action will be passed into your {@link IAction.bind}
	// method in the {@link ActionBindOptions.role} property.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// Whether to grant full permissions to CloudFormation while deploying this template.
	//
	// Setting this to `true` affects the defaults for `role` and `capabilities`, if you
	// don't specify any alternatives.
	//
	// The default role that will be created for you will have full (i.e., `*`)
	// permissions on all resources, and the deployment will have named IAM
	// capabilities (i.e., able to create all IAM resources).
	//
	// This is a shorthand that you can use if you fully trust the templates that
	// are deployed in this pipeline. If you want more fine-grained permissions,
	// use `addToRolePolicy` and `capabilities` to control what the CloudFormation
	// deployment is allowed to do.
	AdminPermissions *bool `field:"required" json:"adminPermissions" yaml:"adminPermissions"`
	// The name of the stack to apply this action to.
	StackName *string `field:"required" json:"stackName" yaml:"stackName"`
	// Input artifact with the CloudFormation template to deploy.
	TemplatePath awscodepipeline.ArtifactPath `field:"required" json:"templatePath" yaml:"templatePath"`
	// The AWS account this Action is supposed to operate in.
	//
	// **Note**: if you specify the `role` property,
	// this is ignored - the action will operate in the same region the passed role does.
	Account *string `field:"optional" json:"account" yaml:"account"`
	// Acknowledge certain changes made as part of deployment.
	//
	// For stacks that contain certain resources,
	// explicit acknowledgement is required that AWS CloudFormation might create or update those resources.
	// For example, you must specify `ANONYMOUS_IAM` or `NAMED_IAM` if your stack template contains AWS
	// Identity and Access Management (IAM) resources.
	// For more information, see the link below.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-iam-template.html#using-iam-capabilities
	//
	CfnCapabilities *[]awscdk.CfnCapabilities `field:"optional" json:"cfnCapabilities" yaml:"cfnCapabilities"`
	// IAM role to assume when deploying changes.
	//
	// If not specified, a fresh role is created. The role is created with zero
	// permissions unless `adminPermissions` is true, in which case the role will have
	// full permissions.
	DeploymentRole awsiam.IRole `field:"optional" json:"deploymentRole" yaml:"deploymentRole"`
	// The list of additional input Artifacts for this Action.
	//
	// This is especially useful when used in conjunction with the `parameterOverrides` property.
	// For example, if you have:
	//
	//    parameterOverrides: {
	//      'Param1': action1.outputArtifact.bucketName,
	//      'Param2': action2.outputArtifact.objectKey,
	//    }
	//
	// , if the output Artifacts of `action1` and `action2` were not used to
	// set either the `templateConfiguration` or the `templatePath` properties,
	// you need to make sure to include them in the `extraInputs` -
	// otherwise, you'll get an "unrecognized Artifact" error during your Pipeline's execution.
	ExtraInputs *[]awscodepipeline.Artifact `field:"optional" json:"extraInputs" yaml:"extraInputs"`
	// The name of the output artifact to generate.
	//
	// Only applied if `outputFileName` is set as well.
	Output awscodepipeline.Artifact `field:"optional" json:"output" yaml:"output"`
	// A name for the filename in the output artifact to store the AWS CloudFormation call's result.
	//
	// The file will contain the result of the call to AWS CloudFormation (for example
	// the call to UpdateStack or CreateChangeSet).
	//
	// AWS CodePipeline adds the file to the output artifact after performing
	// the specified action.
	OutputFileName *string `field:"optional" json:"outputFileName" yaml:"outputFileName"`
	// Additional template parameters.
	//
	// Template parameters specified here take precedence over template parameters
	// found in the artifact specified by the `templateConfiguration` property.
	//
	// We recommend that you use the template configuration file to specify
	// most of your parameter values. Use parameter overrides to specify only
	// dynamic parameter values (values that are unknown until you run the
	// pipeline).
	//
	// All parameter names must be present in the stack template.
	//
	// Note: the entire object cannot be more than 1kB.
	ParameterOverrides *map[string]interface{} `field:"optional" json:"parameterOverrides" yaml:"parameterOverrides"`
	// The AWS region the given Action resides in.
	//
	// Note that a cross-region Pipeline requires replication buckets to function correctly.
	// You can provide their names with the {@link PipelineProps#crossRegionReplicationBuckets} property.
	// If you don't, the CodePipeline Construct will create new Stacks in your CDK app containing those buckets,
	// that you will need to `cdk deploy` before deploying the main, Pipeline-containing Stack.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Replace the stack if it's in a failed state.
	//
	// If this is set to true and the stack is in a failed state (one of
	// ROLLBACK_COMPLETE, ROLLBACK_FAILED, CREATE_FAILED, DELETE_FAILED, or
	// UPDATE_ROLLBACK_FAILED), AWS CloudFormation deletes the stack and then
	// creates a new stack.
	//
	// If this is not set to true and the stack is in a failed state,
	// the deployment fails.
	ReplaceOnFailure *bool `field:"optional" json:"replaceOnFailure" yaml:"replaceOnFailure"`
	// Input artifact to use for template parameters values and stack policy.
	//
	// The template configuration file should contain a JSON object that should look like this:
	// `{ "Parameters": {...}, "Tags": {...}, "StackPolicy": {... }}`. For more information,
	// see [AWS CloudFormation Artifacts](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/continuous-delivery-codepipeline-cfn-artifacts.html).
	//
	// Note that if you include sensitive information, such as passwords, restrict access to this
	// file.
	TemplateConfiguration awscodepipeline.ArtifactPath `field:"optional" json:"templateConfiguration" yaml:"templateConfiguration"`
}

