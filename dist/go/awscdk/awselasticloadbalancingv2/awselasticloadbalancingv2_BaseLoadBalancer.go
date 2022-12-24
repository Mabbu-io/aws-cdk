package awselasticloadbalancingv2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/awselasticloadbalancingv2/internal"
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/awss3"
	"github.com/aws/constructs-go/constructs/v10"
)

// Base class for both Application and Network Load Balancers.
type BaseLoadBalancer interface {
	awscdk.Resource
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	Env() *awscdk.ResourceEnvironment
	// The ARN of this load balancer.
	//
	// Example value: `arn:aws:elasticloadbalancing:us-west-2:123456789012:loadbalancer/app/my-internal-load-balancer/50dc6c495c0c9188`.
	LoadBalancerArn() *string
	// The canonical hosted zone ID of this load balancer.
	//
	// Example value: `Z2P70J7EXAMPLE`.
	LoadBalancerCanonicalHostedZoneId() *string
	// The DNS name of this load balancer.
	//
	// Example value: `my-load-balancer-424835706.us-west-2.elb.amazonaws.com`
	LoadBalancerDnsName() *string
	// The full name of this load balancer.
	//
	// Example value: `app/my-load-balancer/50dc6c495c0c9188`.
	LoadBalancerFullName() *string
	// The name of this load balancer.
	//
	// Example value: `my-load-balancer`.
	LoadBalancerName() *string
	LoadBalancerSecurityGroups() *[]*string
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
	// The VPC this load balancer has been created in.
	//
	// This property is always defined (not `null` or `undefined`) for sub-classes of `BaseLoadBalancer`.
	Vpc() awsec2.IVpc
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
	// Enable access logging for this load balancer.
	//
	// A region must be specified on the stack containing the load balancer; you cannot enable logging on
	// environment-agnostic stacks. See https://docs.aws.amazon.com/cdk/latest/guide/environments.html
	LogAccessLogs(bucket awss3.IBucket, prefix *string)
	// Remove an attribute from the load balancer.
	RemoveAttribute(key *string)
	ResourcePolicyPrincipal() awsiam.IPrincipal
	// Set a non-standard attribute on the load balancer.
	// See: https://docs.aws.amazon.com/elasticloadbalancing/latest/application/application-load-balancers.html#load-balancer-attributes
	//
	SetAttribute(key *string, value *string)
	// Returns a string representation of this construct.
	ToString() *string
	ValidateLoadBalancer() *[]*string
}

// The jsii proxy struct for BaseLoadBalancer
type jsiiProxy_BaseLoadBalancer struct {
	internal.Type__awscdkResource
}

func (j *jsiiProxy_BaseLoadBalancer) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BaseLoadBalancer) LoadBalancerArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancerArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BaseLoadBalancer) LoadBalancerCanonicalHostedZoneId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancerCanonicalHostedZoneId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BaseLoadBalancer) LoadBalancerDnsName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancerDnsName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BaseLoadBalancer) LoadBalancerFullName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancerFullName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BaseLoadBalancer) LoadBalancerName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancerName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BaseLoadBalancer) LoadBalancerSecurityGroups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"loadBalancerSecurityGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BaseLoadBalancer) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BaseLoadBalancer) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BaseLoadBalancer) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BaseLoadBalancer) Vpc() awsec2.IVpc {
	var returns awsec2.IVpc
	_jsii_.Get(
		j,
		"vpc",
		&returns,
	)
	return returns
}


func NewBaseLoadBalancer_Override(b BaseLoadBalancer, scope constructs.Construct, id *string, baseProps *BaseLoadBalancerProps, additionalProps interface{}) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_elasticloadbalancingv2.BaseLoadBalancer",
		[]interface{}{scope, id, baseProps, additionalProps},
		b,
	)
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
func BaseLoadBalancer_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateBaseLoadBalancer_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_elasticloadbalancingv2.BaseLoadBalancer",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
func BaseLoadBalancer_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateBaseLoadBalancer_IsOwnedResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_elasticloadbalancingv2.BaseLoadBalancer",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func BaseLoadBalancer_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateBaseLoadBalancer_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_elasticloadbalancingv2.BaseLoadBalancer",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BaseLoadBalancer) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := b.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (b *jsiiProxy_BaseLoadBalancer) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BaseLoadBalancer) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	if err := b.validateGetResourceArnAttributeParameters(arnAttr, arnComponents); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		b,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BaseLoadBalancer) GetResourceNameAttribute(nameAttr *string) *string {
	if err := b.validateGetResourceNameAttributeParameters(nameAttr); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		b,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BaseLoadBalancer) LogAccessLogs(bucket awss3.IBucket, prefix *string) {
	if err := b.validateLogAccessLogsParameters(bucket); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"logAccessLogs",
		[]interface{}{bucket, prefix},
	)
}

func (b *jsiiProxy_BaseLoadBalancer) RemoveAttribute(key *string) {
	if err := b.validateRemoveAttributeParameters(key); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"removeAttribute",
		[]interface{}{key},
	)
}

func (b *jsiiProxy_BaseLoadBalancer) ResourcePolicyPrincipal() awsiam.IPrincipal {
	var returns awsiam.IPrincipal

	_jsii_.Invoke(
		b,
		"resourcePolicyPrincipal",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BaseLoadBalancer) SetAttribute(key *string, value *string) {
	if err := b.validateSetAttributeParameters(key); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"setAttribute",
		[]interface{}{key, value},
	)
}

func (b *jsiiProxy_BaseLoadBalancer) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BaseLoadBalancer) ValidateLoadBalancer() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"validateLoadBalancer",
		nil, // no parameters
		&returns,
	)

	return returns
}

