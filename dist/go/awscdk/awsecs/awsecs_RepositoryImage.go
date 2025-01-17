package awsecs

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/awsecr"
	"github.com/aws/aws-cdk-go/awscdk/awsecrassets"
	"github.com/aws/constructs-go/constructs/v10"
)

// An image hosted in a public or private repository.
//
// For images hosted in Amazon ECR, see
// [EcrImage](https://docs.aws.amazon.com/AmazonECR/latest/userguide/images.html).
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   import batch "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//
//   jobQueue := batch.NewJobQueue(this, jsii.String("MyQueue"), map[string][]map[string]interface{}{
//   	"computeEnvironments": []map[string]interface{}{
//   		map[string]interface{}{
//   			"computeEnvironment": batch.NewComputeEnvironment(this, jsii.String("ComputeEnvironment"), map[string]*bool{
//   				"managed": jsii.Boolean(false),
//   			}),
//   			"order": jsii.Number(1),
//   		},
//   	},
//   })
//
//   jobDefinition := batch.NewJobDefinition(this, jsii.String("MyJob"), map[string]map[string]repositoryImage{
//   	"container": map[string]repositoryImage{
//   		"image": awscdk.ContainerImage.fromRegistry(jsii.String("test-repo")),
//   	},
//   })
//
//   queue := sqs.NewQueue(this, jsii.String("Queue"))
//
//   rule := events.NewRule(this, jsii.String("Rule"), &ruleProps{
//   	schedule: events.schedule.rate(cdk.duration.hours(jsii.Number(1))),
//   })
//
//   rule.addTarget(targets.NewBatchJob(jobQueue.jobQueueArn, jobQueue, jobDefinition.jobDefinitionArn, jobDefinition, &batchJobProps{
//   	deadLetterQueue: queue,
//   	event: events.ruleTargetInput.fromObject(map[string]*string{
//   		"SomeParam": jsii.String("SomeValue"),
//   	}),
//   	retryAttempts: jsii.Number(2),
//   	maxEventAge: cdk.*duration.hours(jsii.Number(2)),
//   }))
//
type RepositoryImage interface {
	ContainerImage
	// Called when the image is used by a ContainerDefinition.
	Bind(scope constructs.Construct, containerDefinition ContainerDefinition) *ContainerImageConfig
}

// The jsii proxy struct for RepositoryImage
type jsiiProxy_RepositoryImage struct {
	jsiiProxy_ContainerImage
}

// Constructs a new instance of the RepositoryImage class.
func NewRepositoryImage(imageName *string, props *RepositoryImageProps) RepositoryImage {
	_init_.Initialize()

	if err := validateNewRepositoryImageParameters(imageName, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_RepositoryImage{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.RepositoryImage",
		[]interface{}{imageName, props},
		&j,
	)

	return &j
}

// Constructs a new instance of the RepositoryImage class.
func NewRepositoryImage_Override(r RepositoryImage, imageName *string, props *RepositoryImageProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.RepositoryImage",
		[]interface{}{imageName, props},
		r,
	)
}

// Reference an image that's constructed directly from sources on disk.
//
// If you already have a `DockerImageAsset` instance, you can use the
// `ContainerImage.fromDockerImageAsset` method instead.
func RepositoryImage_FromAsset(directory *string, props *AssetImageProps) AssetImage {
	_init_.Initialize()

	if err := validateRepositoryImage_FromAssetParameters(directory, props); err != nil {
		panic(err)
	}
	var returns AssetImage

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.RepositoryImage",
		"fromAsset",
		[]interface{}{directory, props},
		&returns,
	)

	return returns
}

// Use an existing `DockerImageAsset` for this container image.
func RepositoryImage_FromDockerImageAsset(asset awsecrassets.DockerImageAsset) ContainerImage {
	_init_.Initialize()

	if err := validateRepositoryImage_FromDockerImageAssetParameters(asset); err != nil {
		panic(err)
	}
	var returns ContainerImage

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.RepositoryImage",
		"fromDockerImageAsset",
		[]interface{}{asset},
		&returns,
	)

	return returns
}

// Reference an image in an ECR repository.
func RepositoryImage_FromEcrRepository(repository awsecr.IRepository, tag *string) EcrImage {
	_init_.Initialize()

	if err := validateRepositoryImage_FromEcrRepositoryParameters(repository); err != nil {
		panic(err)
	}
	var returns EcrImage

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.RepositoryImage",
		"fromEcrRepository",
		[]interface{}{repository, tag},
		&returns,
	)

	return returns
}

// Reference an image on DockerHub or another online registry.
func RepositoryImage_FromRegistry(name *string, props *RepositoryImageProps) RepositoryImage {
	_init_.Initialize()

	if err := validateRepositoryImage_FromRegistryParameters(name, props); err != nil {
		panic(err)
	}
	var returns RepositoryImage

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.RepositoryImage",
		"fromRegistry",
		[]interface{}{name, props},
		&returns,
	)

	return returns
}

// Use an existing tarball for this container image.
//
// Use this method if the container image has already been created by another process (e.g. jib)
// and you want to add it as a container image asset.
func RepositoryImage_FromTarball(tarballFile *string) ContainerImage {
	_init_.Initialize()

	if err := validateRepositoryImage_FromTarballParameters(tarballFile); err != nil {
		panic(err)
	}
	var returns ContainerImage

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.RepositoryImage",
		"fromTarball",
		[]interface{}{tarballFile},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RepositoryImage) Bind(scope constructs.Construct, containerDefinition ContainerDefinition) *ContainerImageConfig {
	if err := r.validateBindParameters(scope, containerDefinition); err != nil {
		panic(err)
	}
	var returns *ContainerImageConfig

	_jsii_.Invoke(
		r,
		"bind",
		[]interface{}{scope, containerDefinition},
		&returns,
	)

	return returns
}

