package awsecspatterns

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/awsecs"
	"github.com/aws/aws-cdk-go/awscdk/awsecspatterns/internal"
	"github.com/aws/aws-cdk-go/awscdk/awselasticloadbalancingv2"
	"github.com/aws/constructs-go/constructs/v10"
)

// The base class for NetworkLoadBalancedEc2Service and NetworkLoadBalancedFargateService services.
type NetworkLoadBalancedServiceBase interface {
	constructs.Construct
	// The cluster that hosts the service.
	Cluster() awsecs.ICluster
	// The desired number of instantiations of the task definition to keep running on the service.
	//
	// The default is 1 for all new services and uses the existing services desired count
	// when updating an existing service, if one is not provided.
	InternalDesiredCount() *float64
	// The listener for the service.
	Listener() awselasticloadbalancingv2.NetworkListener
	// The Network Load Balancer for the service.
	LoadBalancer() awselasticloadbalancingv2.NetworkLoadBalancer
	// The tree node.
	Node() constructs.Node
	// The target group for the service.
	TargetGroup() awselasticloadbalancingv2.NetworkTargetGroup
	// Adds service as a target of the target group.
	AddServiceAsTarget(service awsecs.BaseService)
	CreateAWSLogDriver(prefix *string) awsecs.AwsLogDriver
	// Returns the default cluster.
	GetDefaultCluster(scope constructs.Construct, vpc awsec2.IVpc) awsecs.Cluster
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for NetworkLoadBalancedServiceBase
type jsiiProxy_NetworkLoadBalancedServiceBase struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_NetworkLoadBalancedServiceBase) Cluster() awsecs.ICluster {
	var returns awsecs.ICluster
	_jsii_.Get(
		j,
		"cluster",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkLoadBalancedServiceBase) InternalDesiredCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"internalDesiredCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkLoadBalancedServiceBase) Listener() awselasticloadbalancingv2.NetworkListener {
	var returns awselasticloadbalancingv2.NetworkListener
	_jsii_.Get(
		j,
		"listener",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkLoadBalancedServiceBase) LoadBalancer() awselasticloadbalancingv2.NetworkLoadBalancer {
	var returns awselasticloadbalancingv2.NetworkLoadBalancer
	_jsii_.Get(
		j,
		"loadBalancer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkLoadBalancedServiceBase) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkLoadBalancedServiceBase) TargetGroup() awselasticloadbalancingv2.NetworkTargetGroup {
	var returns awselasticloadbalancingv2.NetworkTargetGroup
	_jsii_.Get(
		j,
		"targetGroup",
		&returns,
	)
	return returns
}


// Constructs a new instance of the NetworkLoadBalancedServiceBase class.
func NewNetworkLoadBalancedServiceBase_Override(n NetworkLoadBalancedServiceBase, scope constructs.Construct, id *string, props *NetworkLoadBalancedServiceBaseProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs_patterns.NetworkLoadBalancedServiceBase",
		[]interface{}{scope, id, props},
		n,
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
func NetworkLoadBalancedServiceBase_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNetworkLoadBalancedServiceBase_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs_patterns.NetworkLoadBalancedServiceBase",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkLoadBalancedServiceBase) AddServiceAsTarget(service awsecs.BaseService) {
	if err := n.validateAddServiceAsTargetParameters(service); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"addServiceAsTarget",
		[]interface{}{service},
	)
}

func (n *jsiiProxy_NetworkLoadBalancedServiceBase) CreateAWSLogDriver(prefix *string) awsecs.AwsLogDriver {
	if err := n.validateCreateAWSLogDriverParameters(prefix); err != nil {
		panic(err)
	}
	var returns awsecs.AwsLogDriver

	_jsii_.Invoke(
		n,
		"createAWSLogDriver",
		[]interface{}{prefix},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkLoadBalancedServiceBase) GetDefaultCluster(scope constructs.Construct, vpc awsec2.IVpc) awsecs.Cluster {
	if err := n.validateGetDefaultClusterParameters(scope); err != nil {
		panic(err)
	}
	var returns awsecs.Cluster

	_jsii_.Invoke(
		n,
		"getDefaultCluster",
		[]interface{}{scope, vpc},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkLoadBalancedServiceBase) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

