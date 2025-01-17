package awseks

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Controller version.
//
// Corresponds to the image tag of 'amazon/aws-load-balancer-controller' image.
//
// Example:
//   eks.NewCluster(this, jsii.String("HelloEKS"), &clusterProps{
//   	version: eks.kubernetesVersion_V1_21(),
//   	albController: &albControllerOptions{
//   		version: eks.albControllerVersion_V2_4_1(),
//   	},
//   })
//
type AlbControllerVersion interface {
	// Whether or not its a custom version.
	Custom() *bool
	// The version string.
	Version() *string
}

// The jsii proxy struct for AlbControllerVersion
type jsiiProxy_AlbControllerVersion struct {
	_ byte // padding
}

func (j *jsiiProxy_AlbControllerVersion) Custom() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"custom",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbControllerVersion) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}


// Specify a custom version.
//
// Use this if the version you need is not available in one of the predefined versions.
// Note that in this case, you will also need to provide an IAM policy in the controller options.
func AlbControllerVersion_Of(version *string) AlbControllerVersion {
	_init_.Initialize()

	if err := validateAlbControllerVersion_OfParameters(version); err != nil {
		panic(err)
	}
	var returns AlbControllerVersion

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_eks.AlbControllerVersion",
		"of",
		[]interface{}{version},
		&returns,
	)

	return returns
}

func AlbControllerVersion_V2_0_0() AlbControllerVersion {
	_init_.Initialize()
	var returns AlbControllerVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_eks.AlbControllerVersion",
		"V2_0_0",
		&returns,
	)
	return returns
}

func AlbControllerVersion_V2_0_1() AlbControllerVersion {
	_init_.Initialize()
	var returns AlbControllerVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_eks.AlbControllerVersion",
		"V2_0_1",
		&returns,
	)
	return returns
}

func AlbControllerVersion_V2_1_0() AlbControllerVersion {
	_init_.Initialize()
	var returns AlbControllerVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_eks.AlbControllerVersion",
		"V2_1_0",
		&returns,
	)
	return returns
}

func AlbControllerVersion_V2_1_1() AlbControllerVersion {
	_init_.Initialize()
	var returns AlbControllerVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_eks.AlbControllerVersion",
		"V2_1_1",
		&returns,
	)
	return returns
}

func AlbControllerVersion_V2_1_2() AlbControllerVersion {
	_init_.Initialize()
	var returns AlbControllerVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_eks.AlbControllerVersion",
		"V2_1_2",
		&returns,
	)
	return returns
}

func AlbControllerVersion_V2_1_3() AlbControllerVersion {
	_init_.Initialize()
	var returns AlbControllerVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_eks.AlbControllerVersion",
		"V2_1_3",
		&returns,
	)
	return returns
}

func AlbControllerVersion_V2_2_0() AlbControllerVersion {
	_init_.Initialize()
	var returns AlbControllerVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_eks.AlbControllerVersion",
		"V2_2_0",
		&returns,
	)
	return returns
}

func AlbControllerVersion_V2_2_1() AlbControllerVersion {
	_init_.Initialize()
	var returns AlbControllerVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_eks.AlbControllerVersion",
		"V2_2_1",
		&returns,
	)
	return returns
}

func AlbControllerVersion_V2_2_2() AlbControllerVersion {
	_init_.Initialize()
	var returns AlbControllerVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_eks.AlbControllerVersion",
		"V2_2_2",
		&returns,
	)
	return returns
}

func AlbControllerVersion_V2_2_3() AlbControllerVersion {
	_init_.Initialize()
	var returns AlbControllerVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_eks.AlbControllerVersion",
		"V2_2_3",
		&returns,
	)
	return returns
}

func AlbControllerVersion_V2_2_4() AlbControllerVersion {
	_init_.Initialize()
	var returns AlbControllerVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_eks.AlbControllerVersion",
		"V2_2_4",
		&returns,
	)
	return returns
}

func AlbControllerVersion_V2_3_0() AlbControllerVersion {
	_init_.Initialize()
	var returns AlbControllerVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_eks.AlbControllerVersion",
		"V2_3_0",
		&returns,
	)
	return returns
}

func AlbControllerVersion_V2_3_1() AlbControllerVersion {
	_init_.Initialize()
	var returns AlbControllerVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_eks.AlbControllerVersion",
		"V2_3_1",
		&returns,
	)
	return returns
}

func AlbControllerVersion_V2_4_1() AlbControllerVersion {
	_init_.Initialize()
	var returns AlbControllerVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_eks.AlbControllerVersion",
		"V2_4_1",
		&returns,
	)
	return returns
}

