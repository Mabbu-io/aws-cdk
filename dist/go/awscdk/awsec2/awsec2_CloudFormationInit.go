package awsec2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
)

// A CloudFormation-init configuration.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   var myBucket bucket
//
//
//   handle := ec2.NewInitServiceRestartHandle()
//
//   ec2.cloudFormationInit.fromElements(ec2.initFile.fromString(jsii.String("/etc/nginx/nginx.conf"), jsii.String("..."), &initFileOptions{
//   	serviceRestartHandles: []initServiceRestartHandle{
//   		handle,
//   	},
//   }), ec2.initSource.fromS3Object(jsii.String("/var/www/html"), myBucket, jsii.String("html.zip"), &initSourceOptions{
//   	serviceRestartHandles: []*initServiceRestartHandle{
//   		handle,
//   	},
//   }), ec2.initService.enable(jsii.String("nginx"), &initServiceOptions{
//   	serviceRestartHandle: handle,
//   }))
//
type CloudFormationInit interface {
	// Add a config with the given name to this CloudFormationInit object.
	AddConfig(configName *string, config InitConfig)
	// Add a config set with the given name to this CloudFormationInit object.
	//
	// The new configset will reference the given configs in the given order.
	AddConfigSet(configSetName *string, configNames *[]*string)
	// Attach the CloudFormation Init config to the given resource.
	//
	// As an app builder, use `instance.applyCloudFormationInit()` or
	// `autoScalingGroup.applyCloudFormationInit()` to trigger this method.
	//
	// This method does the following:
	//
	// - Renders the `AWS::CloudFormation::Init` object to the given resource's
	//    metadata, potentially adding a `AWS::CloudFormation::Authentication` object
	//    next to it if required.
	// - Updates the instance role policy to be able to call the APIs required for
	//    `cfn-init` and `cfn-signal` to work, and potentially add permissions to download
	//    referenced asset and bucket resources.
	// - Updates the given UserData with commands to execute the `cfn-init` script.
	Attach(attachedResource awscdk.CfnResource, attachOptions *AttachInitOptions)
}

// The jsii proxy struct for CloudFormationInit
type jsiiProxy_CloudFormationInit struct {
	_ byte // padding
}

// Use an existing InitConfig object as the default and only config.
func CloudFormationInit_FromConfig(config InitConfig) CloudFormationInit {
	_init_.Initialize()

	if err := validateCloudFormationInit_FromConfigParameters(config); err != nil {
		panic(err)
	}
	var returns CloudFormationInit

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.CloudFormationInit",
		"fromConfig",
		[]interface{}{config},
		&returns,
	)

	return returns
}

// Build a CloudFormationInit from config sets.
func CloudFormationInit_FromConfigSets(props *ConfigSetProps) CloudFormationInit {
	_init_.Initialize()

	if err := validateCloudFormationInit_FromConfigSetsParameters(props); err != nil {
		panic(err)
	}
	var returns CloudFormationInit

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.CloudFormationInit",
		"fromConfigSets",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Build a new config from a set of Init Elements.
func CloudFormationInit_FromElements(elements ...InitElement) CloudFormationInit {
	_init_.Initialize()

	args := []interface{}{}
	for _, a := range elements {
		args = append(args, a)
	}

	var returns CloudFormationInit

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.CloudFormationInit",
		"fromElements",
		args,
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudFormationInit) AddConfig(configName *string, config InitConfig) {
	if err := c.validateAddConfigParameters(configName, config); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addConfig",
		[]interface{}{configName, config},
	)
}

func (c *jsiiProxy_CloudFormationInit) AddConfigSet(configSetName *string, configNames *[]*string) {
	if err := c.validateAddConfigSetParameters(configSetName); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addConfigSet",
		[]interface{}{configSetName, configNames},
	)
}

func (c *jsiiProxy_CloudFormationInit) Attach(attachedResource awscdk.CfnResource, attachOptions *AttachInitOptions) {
	if err := c.validateAttachParameters(attachedResource, attachOptions); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"attach",
		[]interface{}{attachedResource, attachOptions},
	)
}

