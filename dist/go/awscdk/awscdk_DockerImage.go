// Version 2 of the AWS Cloud Development Kit library
package awscdk

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// A Docker image.
//
// Example:
//   asset := assets.NewAsset(this, jsii.String("BundledAsset"), &assetProps{
//   	path: path.join(__dirname, jsii.String("markdown-asset")),
//   	 // /asset-input and working directory in the container
//   	bundling: &bundlingOptions{
//   		image: awscdk.DockerImage.fromBuild(path.join(__dirname, jsii.String("alpine-markdown"))),
//   		 // Build an image
//   		command: []*string{
//   			jsii.String("sh"),
//   			jsii.String("-c"),
//   			jsii.String("\n            markdown index.md > /asset-output/index.html\n          "),
//   		},
//   	},
//   })
//
type DockerImage interface {
	// The Docker image.
	Image() *string
	// Copies a file or directory out of the Docker image to the local filesystem.
	//
	// If `outputPath` is omitted the destination path is a temporary directory.
	//
	// Returns: the destination path.
	Cp(imagePath *string, outputPath *string) *string
	// Runs a Docker image.
	Run(options *DockerRunOptions)
	// Provides a stable representation of this image for JSON serialization.
	//
	// Returns: The overridden image name if set or image hash name in that order.
	ToJSON() *string
}

// The jsii proxy struct for DockerImage
type jsiiProxy_DockerImage struct {
	_ byte // padding
}

func (j *jsiiProxy_DockerImage) Image() *string {
	var returns *string
	_jsii_.Get(
		j,
		"image",
		&returns,
	)
	return returns
}


func NewDockerImage(image *string, _imageHash *string) DockerImage {
	_init_.Initialize()

	if err := validateNewDockerImageParameters(image); err != nil {
		panic(err)
	}
	j := jsiiProxy_DockerImage{}

	_jsii_.Create(
		"aws-cdk-lib.DockerImage",
		[]interface{}{image, _imageHash},
		&j,
	)

	return &j
}

func NewDockerImage_Override(d DockerImage, image *string, _imageHash *string) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.DockerImage",
		[]interface{}{image, _imageHash},
		d,
	)
}

// Builds a Docker image.
func DockerImage_FromBuild(path *string, options *DockerBuildOptions) DockerImage {
	_init_.Initialize()

	if err := validateDockerImage_FromBuildParameters(path, options); err != nil {
		panic(err)
	}
	var returns DockerImage

	_jsii_.StaticInvoke(
		"aws-cdk-lib.DockerImage",
		"fromBuild",
		[]interface{}{path, options},
		&returns,
	)

	return returns
}

// Reference an image on DockerHub or another online registry.
func DockerImage_FromRegistry(image *string) DockerImage {
	_init_.Initialize()

	if err := validateDockerImage_FromRegistryParameters(image); err != nil {
		panic(err)
	}
	var returns DockerImage

	_jsii_.StaticInvoke(
		"aws-cdk-lib.DockerImage",
		"fromRegistry",
		[]interface{}{image},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DockerImage) Cp(imagePath *string, outputPath *string) *string {
	if err := d.validateCpParameters(imagePath); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"cp",
		[]interface{}{imagePath, outputPath},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DockerImage) Run(options *DockerRunOptions) {
	if err := d.validateRunParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"run",
		[]interface{}{options},
	)
}

func (d *jsiiProxy_DockerImage) ToJSON() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toJSON",
		nil, // no parameters
		&returns,
	)

	return returns
}

