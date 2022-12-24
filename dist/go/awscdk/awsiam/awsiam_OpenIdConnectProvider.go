package awsiam

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsiam/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// IAM OIDC identity providers are entities in IAM that describe an external identity provider (IdP) service that supports the OpenID Connect (OIDC) standard, such as Google or Salesforce.
//
// You use an IAM OIDC identity provider
// when you want to establish trust between an OIDC-compatible IdP and your AWS
// account. This is useful when creating a mobile app or web application that
// requires access to AWS resources, but you don't want to create custom sign-in
// code or manage your own user identities.
//
// Example:
//   provider := iam.NewOpenIdConnectProvider(this, jsii.String("MyProvider"), &openIdConnectProviderProps{
//   	url: jsii.String("https://openid/connect"),
//   	clientIds: []*string{
//   		jsii.String("myclient1"),
//   		jsii.String("myclient2"),
//   	},
//   })
//
// See: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_providers_oidc.html
//
type OpenIdConnectProvider interface {
	awscdk.Resource
	IOpenIdConnectProvider
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	Env() *awscdk.ResourceEnvironment
	// The tree node.
	Node() constructs.Node
	// The Amazon Resource Name (ARN) of the IAM OpenID Connect provider.
	OpenIdConnectProviderArn() *string
	// The issuer for OIDC Provider.
	OpenIdConnectProviderIssuer() *string
	// The thumbprints configured for this provider.
	OpenIdConnectProviderthumbprints() *string
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

// The jsii proxy struct for OpenIdConnectProvider
type jsiiProxy_OpenIdConnectProvider struct {
	internal.Type__awscdkResource
	jsiiProxy_IOpenIdConnectProvider
}

func (j *jsiiProxy_OpenIdConnectProvider) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpenIdConnectProvider) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpenIdConnectProvider) OpenIdConnectProviderArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"openIdConnectProviderArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpenIdConnectProvider) OpenIdConnectProviderIssuer() *string {
	var returns *string
	_jsii_.Get(
		j,
		"openIdConnectProviderIssuer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpenIdConnectProvider) OpenIdConnectProviderthumbprints() *string {
	var returns *string
	_jsii_.Get(
		j,
		"openIdConnectProviderthumbprints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpenIdConnectProvider) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpenIdConnectProvider) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Defines an OpenID Connect provider.
func NewOpenIdConnectProvider(scope constructs.Construct, id *string, props *OpenIdConnectProviderProps) OpenIdConnectProvider {
	_init_.Initialize()

	if err := validateNewOpenIdConnectProviderParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_OpenIdConnectProvider{}

	_jsii_.Create(
		"aws-cdk-lib.aws_iam.OpenIdConnectProvider",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Defines an OpenID Connect provider.
func NewOpenIdConnectProvider_Override(o OpenIdConnectProvider, scope constructs.Construct, id *string, props *OpenIdConnectProviderProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_iam.OpenIdConnectProvider",
		[]interface{}{scope, id, props},
		o,
	)
}

// Imports an Open ID connect provider from an ARN.
func OpenIdConnectProvider_FromOpenIdConnectProviderArn(scope constructs.Construct, id *string, openIdConnectProviderArn *string) IOpenIdConnectProvider {
	_init_.Initialize()

	if err := validateOpenIdConnectProvider_FromOpenIdConnectProviderArnParameters(scope, id, openIdConnectProviderArn); err != nil {
		panic(err)
	}
	var returns IOpenIdConnectProvider

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iam.OpenIdConnectProvider",
		"fromOpenIdConnectProviderArn",
		[]interface{}{scope, id, openIdConnectProviderArn},
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
func OpenIdConnectProvider_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateOpenIdConnectProvider_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iam.OpenIdConnectProvider",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
func OpenIdConnectProvider_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateOpenIdConnectProvider_IsOwnedResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iam.OpenIdConnectProvider",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func OpenIdConnectProvider_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateOpenIdConnectProvider_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iam.OpenIdConnectProvider",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpenIdConnectProvider) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := o.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (o *jsiiProxy_OpenIdConnectProvider) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpenIdConnectProvider) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	if err := o.validateGetResourceArnAttributeParameters(arnAttr, arnComponents); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		o,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpenIdConnectProvider) GetResourceNameAttribute(nameAttr *string) *string {
	if err := o.validateGetResourceNameAttributeParameters(nameAttr); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		o,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpenIdConnectProvider) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

