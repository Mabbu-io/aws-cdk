package awslambda

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsecr"
	"github.com/aws/aws-cdk-go/awscdk/awss3"
	"github.com/aws/aws-cdk-go/awscdk/awss3assets"
	"github.com/aws/constructs-go/constructs/v10"
)

// Represents an ECR image that will be constructed from the specified asset and can be bound as Lambda code.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var networkMode networkMode
//   var platform platform
//
//   assetImageCode := awscdk.Aws_lambda.NewAssetImageCode(jsii.String("directory"), &assetImageCodeProps{
//   	buildArgs: map[string]*string{
//   		"buildArgsKey": jsii.String("buildArgs"),
//   	},
//   	cmd: []*string{
//   		jsii.String("cmd"),
//   	},
//   	entrypoint: []*string{
//   		jsii.String("entrypoint"),
//   	},
//   	exclude: []*string{
//   		jsii.String("exclude"),
//   	},
//   	extraHash: jsii.String("extraHash"),
//   	file: jsii.String("file"),
//   	followSymlinks: cdk.symlinkFollowMode_NEVER,
//   	ignoreMode: cdk.ignoreMode_GLOB,
//   	invalidation: &dockerImageAssetInvalidationOptions{
//   		buildArgs: jsii.Boolean(false),
//   		extraHash: jsii.Boolean(false),
//   		file: jsii.Boolean(false),
//   		networkMode: jsii.Boolean(false),
//   		platform: jsii.Boolean(false),
//   		repositoryName: jsii.Boolean(false),
//   		target: jsii.Boolean(false),
//   	},
//   	networkMode: networkMode,
//   	platform: platform,
//   	target: jsii.String("target"),
//   	workingDirectory: jsii.String("workingDirectory"),
//   })
//
type AssetImageCode interface {
	Code
	// Determines whether this Code is inline code or not.
	IsInline() *bool
	// Called when the lambda or layer is initialized to allow this object to bind to the stack, add resources and have fun.
	Bind(scope constructs.Construct) *CodeConfig
	// Called after the CFN function resource has been created to allow the code class to bind to it.
	//
	// Specifically it's required to allow assets to add
	// metadata for tooling like SAM CLI to be able to find their origins.
	BindToResource(resource awscdk.CfnResource, options *ResourceBindOptions)
}

// The jsii proxy struct for AssetImageCode
type jsiiProxy_AssetImageCode struct {
	jsiiProxy_Code
}

func (j *jsiiProxy_AssetImageCode) IsInline() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isInline",
		&returns,
	)
	return returns
}


func NewAssetImageCode(directory *string, props *AssetImageCodeProps) AssetImageCode {
	_init_.Initialize()

	if err := validateNewAssetImageCodeParameters(directory, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_AssetImageCode{}

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda.AssetImageCode",
		[]interface{}{directory, props},
		&j,
	)

	return &j
}

func NewAssetImageCode_Override(a AssetImageCode, directory *string, props *AssetImageCodeProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda.AssetImageCode",
		[]interface{}{directory, props},
		a,
	)
}

// Loads the function code from a local disk path.
func AssetImageCode_FromAsset(path *string, options *awss3assets.AssetOptions) AssetCode {
	_init_.Initialize()

	if err := validateAssetImageCode_FromAssetParameters(path, options); err != nil {
		panic(err)
	}
	var returns AssetCode

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.AssetImageCode",
		"fromAsset",
		[]interface{}{path, options},
		&returns,
	)

	return returns
}

// Create an ECR image from the specified asset and bind it as the Lambda code.
func AssetImageCode_FromAssetImage(directory *string, props *AssetImageCodeProps) AssetImageCode {
	_init_.Initialize()

	if err := validateAssetImageCode_FromAssetImageParameters(directory, props); err != nil {
		panic(err)
	}
	var returns AssetImageCode

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.AssetImageCode",
		"fromAssetImage",
		[]interface{}{directory, props},
		&returns,
	)

	return returns
}

// Lambda handler code as an S3 object.
func AssetImageCode_FromBucket(bucket awss3.IBucket, key *string, objectVersion *string) S3Code {
	_init_.Initialize()

	if err := validateAssetImageCode_FromBucketParameters(bucket, key); err != nil {
		panic(err)
	}
	var returns S3Code

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.AssetImageCode",
		"fromBucket",
		[]interface{}{bucket, key, objectVersion},
		&returns,
	)

	return returns
}

// Creates a new Lambda source defined using CloudFormation parameters.
//
// Returns: a new instance of `CfnParametersCode`.
func AssetImageCode_FromCfnParameters(props *CfnParametersCodeProps) CfnParametersCode {
	_init_.Initialize()

	if err := validateAssetImageCode_FromCfnParametersParameters(props); err != nil {
		panic(err)
	}
	var returns CfnParametersCode

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.AssetImageCode",
		"fromCfnParameters",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Loads the function code from an asset created by a Docker build.
//
// By default, the asset is expected to be located at `/asset` in the
// image.
func AssetImageCode_FromDockerBuild(path *string, options *DockerBuildAssetOptions) AssetCode {
	_init_.Initialize()

	if err := validateAssetImageCode_FromDockerBuildParameters(path, options); err != nil {
		panic(err)
	}
	var returns AssetCode

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.AssetImageCode",
		"fromDockerBuild",
		[]interface{}{path, options},
		&returns,
	)

	return returns
}

// Use an existing ECR image as the Lambda code.
func AssetImageCode_FromEcrImage(repository awsecr.IRepository, props *EcrImageCodeProps) EcrImageCode {
	_init_.Initialize()

	if err := validateAssetImageCode_FromEcrImageParameters(repository, props); err != nil {
		panic(err)
	}
	var returns EcrImageCode

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.AssetImageCode",
		"fromEcrImage",
		[]interface{}{repository, props},
		&returns,
	)

	return returns
}

// Inline code for Lambda handler.
//
// Returns: `LambdaInlineCode` with inline code.
func AssetImageCode_FromInline(code *string) InlineCode {
	_init_.Initialize()

	if err := validateAssetImageCode_FromInlineParameters(code); err != nil {
		panic(err)
	}
	var returns InlineCode

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.AssetImageCode",
		"fromInline",
		[]interface{}{code},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AssetImageCode) Bind(scope constructs.Construct) *CodeConfig {
	if err := a.validateBindParameters(scope); err != nil {
		panic(err)
	}
	var returns *CodeConfig

	_jsii_.Invoke(
		a,
		"bind",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AssetImageCode) BindToResource(resource awscdk.CfnResource, options *ResourceBindOptions) {
	if err := a.validateBindToResourceParameters(resource, options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"bindToResource",
		[]interface{}{resource, options},
	)
}

