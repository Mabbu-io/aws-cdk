package awsecs

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// A log driver that sends log information to syslog Logs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   syslogLogDriver := awscdk.Aws_ecs.NewSyslogLogDriver(&syslogLogDriverProps{
//   	address: jsii.String("address"),
//   	env: []*string{
//   		jsii.String("env"),
//   	},
//   	envRegex: jsii.String("envRegex"),
//   	facility: jsii.String("facility"),
//   	format: jsii.String("format"),
//   	labels: []*string{
//   		jsii.String("labels"),
//   	},
//   	tag: jsii.String("tag"),
//   	tlsCaCert: jsii.String("tlsCaCert"),
//   	tlsCert: jsii.String("tlsCert"),
//   	tlsKey: jsii.String("tlsKey"),
//   	tlsSkipVerify: jsii.Boolean(false),
//   })
//
type SyslogLogDriver interface {
	LogDriver
	// Called when the log driver is configured on a container.
	Bind(_scope constructs.Construct, _containerDefinition ContainerDefinition) *LogDriverConfig
}

// The jsii proxy struct for SyslogLogDriver
type jsiiProxy_SyslogLogDriver struct {
	jsiiProxy_LogDriver
}

// Constructs a new instance of the SyslogLogDriver class.
func NewSyslogLogDriver(props *SyslogLogDriverProps) SyslogLogDriver {
	_init_.Initialize()

	if err := validateNewSyslogLogDriverParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_SyslogLogDriver{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.SyslogLogDriver",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Constructs a new instance of the SyslogLogDriver class.
func NewSyslogLogDriver_Override(s SyslogLogDriver, props *SyslogLogDriverProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.SyslogLogDriver",
		[]interface{}{props},
		s,
	)
}

// Creates a log driver configuration that sends log information to CloudWatch Logs.
func SyslogLogDriver_AwsLogs(props *AwsLogDriverProps) LogDriver {
	_init_.Initialize()

	if err := validateSyslogLogDriver_AwsLogsParameters(props); err != nil {
		panic(err)
	}
	var returns LogDriver

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.SyslogLogDriver",
		"awsLogs",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyslogLogDriver) Bind(_scope constructs.Construct, _containerDefinition ContainerDefinition) *LogDriverConfig {
	if err := s.validateBindParameters(_scope, _containerDefinition); err != nil {
		panic(err)
	}
	var returns *LogDriverConfig

	_jsii_.Invoke(
		s,
		"bind",
		[]interface{}{_scope, _containerDefinition},
		&returns,
	)

	return returns
}

