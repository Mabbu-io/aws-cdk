package awscodepipelineactions

import (
	"github.com/aws/aws-cdk-go/awscdk/awscodepipeline"
	"github.com/aws/aws-cdk-go/awscdk/awsecr"
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
)

// Construction properties of {@link EcrSourceAction}.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   import ecr "github.com/aws/aws-cdk-go/awscdk"
//
//   var ecrRepository repository
//
//   pipeline := codepipeline.NewPipeline(this, jsii.String("MyPipeline"))
//   sourceOutput := codepipeline.NewArtifact()
//   sourceAction := codepipeline_actions.NewEcrSourceAction(&ecrSourceActionProps{
//   	actionName: jsii.String("ECR"),
//   	repository: ecrRepository,
//   	imageTag: jsii.String("some-tag"),
//   	 // optional, default: 'latest'
//   	output: sourceOutput,
//   })
//   pipeline.addStage(&stageOptions{
//   	stageName: jsii.String("Source"),
//   	actions: []iAction{
//   		sourceAction,
//   	},
//   })
//
type EcrSourceActionProps struct {
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
	Output awscodepipeline.Artifact `field:"required" json:"output" yaml:"output"`
	// The repository that will be watched for changes.
	Repository awsecr.IRepository `field:"required" json:"repository" yaml:"repository"`
	// The image tag that will be checked for changes.
	//
	// It is not possible to trigger on changes to more than one tag.
	ImageTag *string `field:"optional" json:"imageTag" yaml:"imageTag"`
}

