package awsecs

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// A log driver that sends log information to journald Logs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   fluentdLogDriver := awscdk.Aws_ecs.NewFluentdLogDriver(&fluentdLogDriverProps{
//   	address: jsii.String("address"),
//   	asyncConnect: jsii.Boolean(false),
//   	bufferLimit: jsii.Number(123),
//   	env: []*string{
//   		jsii.String("env"),
//   	},
//   	envRegex: jsii.String("envRegex"),
//   	labels: []*string{
//   		jsii.String("labels"),
//   	},
//   	maxRetries: jsii.Number(123),
//   	retryWait: cdk.duration.minutes(jsii.Number(30)),
//   	subSecondPrecision: jsii.Boolean(false),
//   	tag: jsii.String("tag"),
//   })
//
type FluentdLogDriver interface {
	LogDriver
	// Called when the log driver is configured on a container.
	Bind(_scope constructs.Construct, _containerDefinition ContainerDefinition) *LogDriverConfig
}

// The jsii proxy struct for FluentdLogDriver
type jsiiProxy_FluentdLogDriver struct {
	jsiiProxy_LogDriver
}

// Constructs a new instance of the FluentdLogDriver class.
func NewFluentdLogDriver(props *FluentdLogDriverProps) FluentdLogDriver {
	_init_.Initialize()

	if err := validateNewFluentdLogDriverParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_FluentdLogDriver{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.FluentdLogDriver",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Constructs a new instance of the FluentdLogDriver class.
func NewFluentdLogDriver_Override(f FluentdLogDriver, props *FluentdLogDriverProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.FluentdLogDriver",
		[]interface{}{props},
		f,
	)
}

// Creates a log driver configuration that sends log information to CloudWatch Logs.
func FluentdLogDriver_AwsLogs(props *AwsLogDriverProps) LogDriver {
	_init_.Initialize()

	if err := validateFluentdLogDriver_AwsLogsParameters(props); err != nil {
		panic(err)
	}
	var returns LogDriver

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.FluentdLogDriver",
		"awsLogs",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FluentdLogDriver) Bind(_scope constructs.Construct, _containerDefinition ContainerDefinition) *LogDriverConfig {
	if err := f.validateBindParameters(_scope, _containerDefinition); err != nil {
		panic(err)
	}
	var returns *LogDriverConfig

	_jsii_.Invoke(
		f,
		"bind",
		[]interface{}{_scope, _containerDefinition},
		&returns,
	)

	return returns
}

