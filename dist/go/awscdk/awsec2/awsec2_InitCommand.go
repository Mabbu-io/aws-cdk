package awsec2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Command to execute on the instance.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   handle := ec2.NewInitServiceRestartHandle()
//   ec2.cloudFormationInit.fromElements(ec2.initCommand.shellCommand(jsii.String("/usr/bin/custom-nginx-install.sh"), &initCommandOptions{
//   	serviceRestartHandles: []initServiceRestartHandle{
//   		handle,
//   	},
//   }), ec2.initService.enable(jsii.String("nginx"), &initServiceOptions{
//   	serviceRestartHandle: handle,
//   }))
//
type InitCommand interface {
	InitElement
	// Returns the init element type for this element.
	ElementType() *string
}

// The jsii proxy struct for InitCommand
type jsiiProxy_InitCommand struct {
	jsiiProxy_InitElement
}

func (j *jsiiProxy_InitCommand) ElementType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"elementType",
		&returns,
	)
	return returns
}


// Run a command from an argv array.
//
// You do not need to escape space characters or enclose command parameters in quotes.
func InitCommand_ArgvCommand(argv *[]*string, options *InitCommandOptions) InitCommand {
	_init_.Initialize()

	if err := validateInitCommand_ArgvCommandParameters(argv, options); err != nil {
		panic(err)
	}
	var returns InitCommand

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.InitCommand",
		"argvCommand",
		[]interface{}{argv, options},
		&returns,
	)

	return returns
}

// Run a shell command.
//
// Remember that some characters like `&`, `|`, `;`, `>` etc. have special meaning in a shell and
// need to be preceded by a `\` if you want to treat them as part of a filename.
func InitCommand_ShellCommand(shellCommand *string, options *InitCommandOptions) InitCommand {
	_init_.Initialize()

	if err := validateInitCommand_ShellCommandParameters(shellCommand, options); err != nil {
		panic(err)
	}
	var returns InitCommand

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.InitCommand",
		"shellCommand",
		[]interface{}{shellCommand, options},
		&returns,
	)

	return returns
}

