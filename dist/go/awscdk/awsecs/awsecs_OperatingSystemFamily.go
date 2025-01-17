package awsecs

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// The operating system for Fargate Runtime Platform.
//
// Example:
//   // Create a Task Definition for the Windows container to start
//   taskDefinition := ecs.NewFargateTaskDefinition(this, jsii.String("TaskDef"), &fargateTaskDefinitionProps{
//   	runtimePlatform: &runtimePlatform{
//   		operatingSystemFamily: ecs.operatingSystemFamily_WINDOWS_SERVER_2019_CORE(),
//   		cpuArchitecture: ecs.cpuArchitecture_X86_64(),
//   	},
//   	cpu: jsii.Number(1024),
//   	memoryLimitMiB: jsii.Number(2048),
//   })
//
//   taskDefinition.addContainer(jsii.String("windowsservercore"), &containerDefinitionOptions{
//   	logging: ecs.logDriver.awsLogs(&awsLogDriverProps{
//   		streamPrefix: jsii.String("win-iis-on-fargate"),
//   	}),
//   	portMappings: []portMapping{
//   		&portMapping{
//   			containerPort: jsii.Number(80),
//   		},
//   	},
//   	image: ecs.containerImage.fromRegistry(jsii.String("mcr.microsoft.com/windows/servercore/iis:windowsservercore-ltsc2019")),
//   })
//
type OperatingSystemFamily interface {
}

// The jsii proxy struct for OperatingSystemFamily
type jsiiProxy_OperatingSystemFamily struct {
	_ byte // padding
}

// Other operating system family.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-runtimeplatform.html#cfn-ecs-taskdefinition-runtimeplatform-operatingsystemfamily for all available operating system family.
//
func OperatingSystemFamily_Of(family *string) OperatingSystemFamily {
	_init_.Initialize()

	if err := validateOperatingSystemFamily_OfParameters(family); err != nil {
		panic(err)
	}
	var returns OperatingSystemFamily

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.OperatingSystemFamily",
		"of",
		[]interface{}{family},
		&returns,
	)

	return returns
}

func OperatingSystemFamily_LINUX() OperatingSystemFamily {
	_init_.Initialize()
	var returns OperatingSystemFamily
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ecs.OperatingSystemFamily",
		"LINUX",
		&returns,
	)
	return returns
}

func OperatingSystemFamily_WINDOWS_SERVER_2004_CORE() OperatingSystemFamily {
	_init_.Initialize()
	var returns OperatingSystemFamily
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ecs.OperatingSystemFamily",
		"WINDOWS_SERVER_2004_CORE",
		&returns,
	)
	return returns
}

func OperatingSystemFamily_WINDOWS_SERVER_2016_FULL() OperatingSystemFamily {
	_init_.Initialize()
	var returns OperatingSystemFamily
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ecs.OperatingSystemFamily",
		"WINDOWS_SERVER_2016_FULL",
		&returns,
	)
	return returns
}

func OperatingSystemFamily_WINDOWS_SERVER_2019_CORE() OperatingSystemFamily {
	_init_.Initialize()
	var returns OperatingSystemFamily
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ecs.OperatingSystemFamily",
		"WINDOWS_SERVER_2019_CORE",
		&returns,
	)
	return returns
}

func OperatingSystemFamily_WINDOWS_SERVER_2019_FULL() OperatingSystemFamily {
	_init_.Initialize()
	var returns OperatingSystemFamily
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ecs.OperatingSystemFamily",
		"WINDOWS_SERVER_2019_FULL",
		&returns,
	)
	return returns
}

func OperatingSystemFamily_WINDOWS_SERVER_2022_CORE() OperatingSystemFamily {
	_init_.Initialize()
	var returns OperatingSystemFamily
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ecs.OperatingSystemFamily",
		"WINDOWS_SERVER_2022_CORE",
		&returns,
	)
	return returns
}

func OperatingSystemFamily_WINDOWS_SERVER_2022_FULL() OperatingSystemFamily {
	_init_.Initialize()
	var returns OperatingSystemFamily
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ecs.OperatingSystemFamily",
		"WINDOWS_SERVER_2022_FULL",
		&returns,
	)
	return returns
}

func OperatingSystemFamily_WINDOWS_SERVER_20H2_CORE() OperatingSystemFamily {
	_init_.Initialize()
	var returns OperatingSystemFamily
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ecs.OperatingSystemFamily",
		"WINDOWS_SERVER_20H2_CORE",
		&returns,
	)
	return returns
}

