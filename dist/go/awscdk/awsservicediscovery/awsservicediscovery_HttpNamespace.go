package awsservicediscovery

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsservicediscovery/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Define an HTTP Namespace.
//
// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import servicediscovery "github.com/aws/aws-cdk-go/awscdk"
//
//   app := cdk.NewApp()
//   stack := cdk.NewStack(app, jsii.String("aws-servicediscovery-integ"))
//
//   namespace := servicediscovery.NewHttpNamespace(stack, jsii.String("MyNamespace"), &httpNamespaceProps{
//   	name: jsii.String("MyHTTPNamespace"),
//   })
//
//   service1 := namespace.createService(jsii.String("NonIpService"), &baseServiceProps{
//   	description: jsii.String("service registering non-ip instances"),
//   })
//
//   service1.registerNonIpInstance(jsii.String("NonIpInstance"), &nonIpInstanceBaseProps{
//   	customAttributes: map[string]*string{
//   		"arn": jsii.String("arn:aws:s3:::mybucket"),
//   	},
//   })
//
//   service2 := namespace.createService(jsii.String("IpService"), &baseServiceProps{
//   	description: jsii.String("service registering ip instances"),
//   	healthCheck: &healthCheckConfig{
//   		type: servicediscovery.healthCheckType_HTTP,
//   		resourcePath: jsii.String("/check"),
//   	},
//   })
//
//   service2.registerIpInstance(jsii.String("IpInstance"), &ipInstanceBaseProps{
//   	ipv4: jsii.String("54.239.25.192"),
//   })
//
//   app.synth()
//
type HttpNamespace interface {
	awscdk.Resource
	IHttpNamespace
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	Env() *awscdk.ResourceEnvironment
	HttpNamespaceArn() *string
	HttpNamespaceId() *string
	HttpNamespaceName() *string
	// Namespace Arn for the namespace.
	NamespaceArn() *string
	// Namespace Id for the namespace.
	NamespaceId() *string
	// A name for the namespace.
	NamespaceName() *string
	// The tree node.
	Node() constructs.Node
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//    cross-environment scenarios.
	PhysicalName() *string
	// The stack in which this resource is defined.
	Stack() awscdk.Stack
	// Type of the namespace.
	Type() NamespaceType
	// Apply the given removal policy to this resource.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	// Creates a service within the namespace.
	CreateService(id *string, props *BaseServiceProps) Service
	GeneratePhysicalName() *string
	// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
	//
	// Normally, this token will resolve to `arnAttr`, but if the resource is
	// referenced across environments, `arnComponents` will be used to synthesize
	// a concrete ARN with the resource's physical name. Make sure to reference
	// `this.physicalName` in `arnComponents`.
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
	//
	// Normally, this token will resolve to `nameAttr`, but if the resource is
	// referenced across environments, it will be resolved to `this.physicalName`,
	// which will be a concrete name.
	GetResourceNameAttribute(nameAttr *string) *string
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for HttpNamespace
type jsiiProxy_HttpNamespace struct {
	internal.Type__awscdkResource
	jsiiProxy_IHttpNamespace
}

func (j *jsiiProxy_HttpNamespace) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HttpNamespace) HttpNamespaceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpNamespaceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HttpNamespace) HttpNamespaceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpNamespaceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HttpNamespace) HttpNamespaceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpNamespaceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HttpNamespace) NamespaceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HttpNamespace) NamespaceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HttpNamespace) NamespaceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HttpNamespace) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HttpNamespace) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HttpNamespace) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HttpNamespace) Type() NamespaceType {
	var returns NamespaceType
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}


func NewHttpNamespace(scope constructs.Construct, id *string, props *HttpNamespaceProps) HttpNamespace {
	_init_.Initialize()

	if err := validateNewHttpNamespaceParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_HttpNamespace{}

	_jsii_.Create(
		"aws-cdk-lib.aws_servicediscovery.HttpNamespace",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewHttpNamespace_Override(h HttpNamespace, scope constructs.Construct, id *string, props *HttpNamespaceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_servicediscovery.HttpNamespace",
		[]interface{}{scope, id, props},
		h,
	)
}

func HttpNamespace_FromHttpNamespaceAttributes(scope constructs.Construct, id *string, attrs *HttpNamespaceAttributes) IHttpNamespace {
	_init_.Initialize()

	if err := validateHttpNamespace_FromHttpNamespaceAttributesParameters(scope, id, attrs); err != nil {
		panic(err)
	}
	var returns IHttpNamespace

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_servicediscovery.HttpNamespace",
		"fromHttpNamespaceAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func HttpNamespace_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateHttpNamespace_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_servicediscovery.HttpNamespace",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
func HttpNamespace_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateHttpNamespace_IsOwnedResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_servicediscovery.HttpNamespace",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func HttpNamespace_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateHttpNamespace_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_servicediscovery.HttpNamespace",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HttpNamespace) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := h.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (h *jsiiProxy_HttpNamespace) CreateService(id *string, props *BaseServiceProps) Service {
	if err := h.validateCreateServiceParameters(id, props); err != nil {
		panic(err)
	}
	var returns Service

	_jsii_.Invoke(
		h,
		"createService",
		[]interface{}{id, props},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HttpNamespace) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HttpNamespace) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	if err := h.validateGetResourceArnAttributeParameters(arnAttr, arnComponents); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		h,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HttpNamespace) GetResourceNameAttribute(nameAttr *string) *string {
	if err := h.validateGetResourceNameAttributeParameters(nameAttr); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		h,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HttpNamespace) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

