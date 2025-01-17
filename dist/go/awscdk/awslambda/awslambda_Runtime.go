package awslambda

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
)

// Lambda function runtime environment.
//
// If you need to use a runtime name that doesn't exist as a static member, you
// can instantiate a `Runtime` object, e.g: `new Runtime('nodejs99.99')`.
//
// Example:
//   import lambda "github.com/aws/aws-cdk-go/awscdk"
//
//
//   fn := lambda.NewFunction(this, jsii.String("MyFunc"), &functionProps{
//   	runtime: lambda.runtime_NODEJS_14_X(),
//   	handler: jsii.String("index.handler"),
//   	code: lambda.code.fromInline(jsii.String("exports.handler = handler.toString()")),
//   })
//
//   rule := events.NewRule(this, jsii.String("rule"), &ruleProps{
//   	eventPattern: &eventPattern{
//   		source: []*string{
//   			jsii.String("aws.ec2"),
//   		},
//   	},
//   })
//
//   queue := sqs.NewQueue(this, jsii.String("Queue"))
//
//   rule.addTarget(targets.NewLambdaFunction(fn, &lambdaFunctionProps{
//   	deadLetterQueue: queue,
//   	 // Optional: add a dead letter queue
//   	maxEventAge: cdk.duration.hours(jsii.Number(2)),
//   	 // Optional: set the maxEventAge retry policy
//   	retryAttempts: jsii.Number(2),
//   }))
//
type Runtime interface {
	// The bundling Docker image for this runtime.
	BundlingImage() awscdk.DockerImage
	// The runtime family.
	Family() RuntimeFamily
	// The name of this runtime, as expected by the Lambda resource.
	Name() *string
	// Whether this runtime is integrated with and supported for profiling using Amazon CodeGuru Profiler.
	SupportsCodeGuruProfiling() *bool
	// Whether the ``ZipFile`` (aka inline code) property can be used with this runtime.
	SupportsInlineCode() *bool
	RuntimeEquals(other Runtime) *bool
	ToString() *string
}

// The jsii proxy struct for Runtime
type jsiiProxy_Runtime struct {
	_ byte // padding
}

func (j *jsiiProxy_Runtime) BundlingImage() awscdk.DockerImage {
	var returns awscdk.DockerImage
	_jsii_.Get(
		j,
		"bundlingImage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Runtime) Family() RuntimeFamily {
	var returns RuntimeFamily
	_jsii_.Get(
		j,
		"family",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Runtime) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Runtime) SupportsCodeGuruProfiling() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"supportsCodeGuruProfiling",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Runtime) SupportsInlineCode() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"supportsInlineCode",
		&returns,
	)
	return returns
}


func NewRuntime(name *string, family RuntimeFamily, props *LambdaRuntimeProps) Runtime {
	_init_.Initialize()

	if err := validateNewRuntimeParameters(name, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_Runtime{}

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda.Runtime",
		[]interface{}{name, family, props},
		&j,
	)

	return &j
}

func NewRuntime_Override(r Runtime, name *string, family RuntimeFamily, props *LambdaRuntimeProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda.Runtime",
		[]interface{}{name, family, props},
		r,
	)
}

func Runtime_ALL() *[]Runtime {
	_init_.Initialize()
	var returns *[]Runtime
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.Runtime",
		"ALL",
		&returns,
	)
	return returns
}

func Runtime_DOTNET_6() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.Runtime",
		"DOTNET_6",
		&returns,
	)
	return returns
}

func Runtime_DOTNET_CORE_1() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.Runtime",
		"DOTNET_CORE_1",
		&returns,
	)
	return returns
}

func Runtime_DOTNET_CORE_2() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.Runtime",
		"DOTNET_CORE_2",
		&returns,
	)
	return returns
}

func Runtime_DOTNET_CORE_2_1() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.Runtime",
		"DOTNET_CORE_2_1",
		&returns,
	)
	return returns
}

func Runtime_DOTNET_CORE_3_1() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.Runtime",
		"DOTNET_CORE_3_1",
		&returns,
	)
	return returns
}

func Runtime_FROM_IMAGE() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.Runtime",
		"FROM_IMAGE",
		&returns,
	)
	return returns
}

func Runtime_GO_1_X() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.Runtime",
		"GO_1_X",
		&returns,
	)
	return returns
}

func Runtime_JAVA_11() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.Runtime",
		"JAVA_11",
		&returns,
	)
	return returns
}

func Runtime_JAVA_8() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.Runtime",
		"JAVA_8",
		&returns,
	)
	return returns
}

func Runtime_JAVA_8_CORRETTO() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.Runtime",
		"JAVA_8_CORRETTO",
		&returns,
	)
	return returns
}

func Runtime_NODEJS() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.Runtime",
		"NODEJS",
		&returns,
	)
	return returns
}

func Runtime_NODEJS_10_X() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.Runtime",
		"NODEJS_10_X",
		&returns,
	)
	return returns
}

func Runtime_NODEJS_12_X() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.Runtime",
		"NODEJS_12_X",
		&returns,
	)
	return returns
}

func Runtime_NODEJS_14_X() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.Runtime",
		"NODEJS_14_X",
		&returns,
	)
	return returns
}

func Runtime_NODEJS_16_X() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.Runtime",
		"NODEJS_16_X",
		&returns,
	)
	return returns
}

func Runtime_NODEJS_18_X() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.Runtime",
		"NODEJS_18_X",
		&returns,
	)
	return returns
}

func Runtime_NODEJS_4_3() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.Runtime",
		"NODEJS_4_3",
		&returns,
	)
	return returns
}

func Runtime_NODEJS_6_10() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.Runtime",
		"NODEJS_6_10",
		&returns,
	)
	return returns
}

func Runtime_NODEJS_8_10() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.Runtime",
		"NODEJS_8_10",
		&returns,
	)
	return returns
}

func Runtime_PROVIDED() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.Runtime",
		"PROVIDED",
		&returns,
	)
	return returns
}

func Runtime_PROVIDED_AL2() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.Runtime",
		"PROVIDED_AL2",
		&returns,
	)
	return returns
}

func Runtime_PYTHON_2_7() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.Runtime",
		"PYTHON_2_7",
		&returns,
	)
	return returns
}

func Runtime_PYTHON_3_6() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.Runtime",
		"PYTHON_3_6",
		&returns,
	)
	return returns
}

func Runtime_PYTHON_3_7() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.Runtime",
		"PYTHON_3_7",
		&returns,
	)
	return returns
}

func Runtime_PYTHON_3_8() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.Runtime",
		"PYTHON_3_8",
		&returns,
	)
	return returns
}

func Runtime_PYTHON_3_9() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.Runtime",
		"PYTHON_3_9",
		&returns,
	)
	return returns
}

func Runtime_RUBY_2_5() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.Runtime",
		"RUBY_2_5",
		&returns,
	)
	return returns
}

func Runtime_RUBY_2_7() Runtime {
	_init_.Initialize()
	var returns Runtime
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.Runtime",
		"RUBY_2_7",
		&returns,
	)
	return returns
}

func (r *jsiiProxy_Runtime) RuntimeEquals(other Runtime) *bool {
	if err := r.validateRuntimeEqualsParameters(other); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.Invoke(
		r,
		"runtimeEquals",
		[]interface{}{other},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Runtime) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

