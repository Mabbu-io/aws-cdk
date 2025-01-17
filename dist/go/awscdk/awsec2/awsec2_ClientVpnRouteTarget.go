package awsec2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Target for a client VPN route.
//
// Example:
//   endpoint := vpc.addClientVpnEndpoint(jsii.String("Endpoint"), &clientVpnEndpointOptions{
//   	cidr: jsii.String("10.100.0.0/16"),
//   	serverCertificateArn: jsii.String("arn:aws:acm:us-east-1:123456789012:certificate/server-certificate-id"),
//   	userBasedAuthentication: ec2.clientVpnUserBasedAuthentication.federated(samlProvider),
//   })
//
//   // Client-to-client access
//   endpoint.addRoute(jsii.String("Route"), &clientVpnRouteOptions{
//   	cidr: jsii.String("10.100.0.0/16"),
//   	target: ec2.clientVpnRouteTarget.local(),
//   })
//
type ClientVpnRouteTarget interface {
	// The subnet ID.
	SubnetId() *string
}

// The jsii proxy struct for ClientVpnRouteTarget
type jsiiProxy_ClientVpnRouteTarget struct {
	_ byte // padding
}

func (j *jsiiProxy_ClientVpnRouteTarget) SubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetId",
		&returns,
	)
	return returns
}


func NewClientVpnRouteTarget_Override(c ClientVpnRouteTarget) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ec2.ClientVpnRouteTarget",
		nil, // no parameters
		c,
	)
}

// Local network.
func ClientVpnRouteTarget_Local() ClientVpnRouteTarget {
	_init_.Initialize()

	var returns ClientVpnRouteTarget

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.ClientVpnRouteTarget",
		"local",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Subnet.
//
// The specified subnet must be an existing target network of the client VPN
// endpoint.
func ClientVpnRouteTarget_Subnet(subnet ISubnet) ClientVpnRouteTarget {
	_init_.Initialize()

	if err := validateClientVpnRouteTarget_SubnetParameters(subnet); err != nil {
		panic(err)
	}
	var returns ClientVpnRouteTarget

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.ClientVpnRouteTarget",
		"subnet",
		[]interface{}{subnet},
		&returns,
	)

	return returns
}

