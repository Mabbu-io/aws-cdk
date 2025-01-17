package awscodedeploy

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/awselasticloadbalancing"
	"github.com/aws/aws-cdk-go/awscdk/awselasticloadbalancingv2"
)

// An interface of an abstract load balancer, as needed by CodeDeploy.
//
// Create instances using the static factory methods:
// {@link #classic}, {@link #application} and {@link #network}.
//
// Example:
//   import elb "github.com/aws/aws-cdk-go/awscdk"
//
//   var lb loadBalancer
//
//   lb.addListener(&loadBalancerListener{
//   	externalPort: jsii.Number(80),
//   })
//
//   deploymentGroup := codedeploy.NewServerDeploymentGroup(this, jsii.String("DeploymentGroup"), &serverDeploymentGroupProps{
//   	loadBalancer: codedeploy.loadBalancer.classic(lb),
//   })
//
type LoadBalancer interface {
	Generation() LoadBalancerGeneration
	Name() *string
}

// The jsii proxy struct for LoadBalancer
type jsiiProxy_LoadBalancer struct {
	_ byte // padding
}

func (j *jsiiProxy_LoadBalancer) Generation() LoadBalancerGeneration {
	var returns LoadBalancerGeneration
	_jsii_.Get(
		j,
		"generation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LoadBalancer) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}


func NewLoadBalancer_Override(l LoadBalancer) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_codedeploy.LoadBalancer",
		nil, // no parameters
		l,
	)
}

// Creates a new CodeDeploy load balancer from an Application Load Balancer Target Group.
func LoadBalancer_Application(albTargetGroup awselasticloadbalancingv2.IApplicationTargetGroup) LoadBalancer {
	_init_.Initialize()

	if err := validateLoadBalancer_ApplicationParameters(albTargetGroup); err != nil {
		panic(err)
	}
	var returns LoadBalancer

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codedeploy.LoadBalancer",
		"application",
		[]interface{}{albTargetGroup},
		&returns,
	)

	return returns
}

// Creates a new CodeDeploy load balancer from a Classic ELB Load Balancer.
func LoadBalancer_Classic(loadBalancer awselasticloadbalancing.LoadBalancer) LoadBalancer {
	_init_.Initialize()

	if err := validateLoadBalancer_ClassicParameters(loadBalancer); err != nil {
		panic(err)
	}
	var returns LoadBalancer

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codedeploy.LoadBalancer",
		"classic",
		[]interface{}{loadBalancer},
		&returns,
	)

	return returns
}

// Creates a new CodeDeploy load balancer from a Network Load Balancer Target Group.
func LoadBalancer_Network(nlbTargetGroup awselasticloadbalancingv2.INetworkTargetGroup) LoadBalancer {
	_init_.Initialize()

	if err := validateLoadBalancer_NetworkParameters(nlbTargetGroup); err != nil {
		panic(err)
	}
	var returns LoadBalancer

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codedeploy.LoadBalancer",
		"network",
		[]interface{}{nlbTargetGroup},
		&returns,
	)

	return returns
}

