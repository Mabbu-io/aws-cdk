package awslambda

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// The collection of versions of the ADOT Lambda Layer for Python SDK.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   adotLambdaLayerPythonSdkVersion := awscdk.Aws_lambda.adotLambdaLayerPythonSdkVersion_LATEST()
//
type AdotLambdaLayerPythonSdkVersion interface {
	LayerVersion() *string
	// The ARN of the Lambda layer.
	LayerArn(scope constructs.IConstruct, architecture Architecture) *string
}

// The jsii proxy struct for AdotLambdaLayerPythonSdkVersion
type jsiiProxy_AdotLambdaLayerPythonSdkVersion struct {
	_ byte // padding
}

func (j *jsiiProxy_AdotLambdaLayerPythonSdkVersion) LayerVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"layerVersion",
		&returns,
	)
	return returns
}


func AdotLambdaLayerPythonSdkVersion_LATEST() AdotLambdaLayerPythonSdkVersion {
	_init_.Initialize()
	var returns AdotLambdaLayerPythonSdkVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.AdotLambdaLayerPythonSdkVersion",
		"LATEST",
		&returns,
	)
	return returns
}

func AdotLambdaLayerPythonSdkVersion_V1_13_0() AdotLambdaLayerPythonSdkVersion {
	_init_.Initialize()
	var returns AdotLambdaLayerPythonSdkVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.AdotLambdaLayerPythonSdkVersion",
		"V1_13_0",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AdotLambdaLayerPythonSdkVersion) LayerArn(scope constructs.IConstruct, architecture Architecture) *string {
	if err := a.validateLayerArnParameters(scope, architecture); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"layerArn",
		[]interface{}{scope, architecture},
		&returns,
	)

	return returns
}

