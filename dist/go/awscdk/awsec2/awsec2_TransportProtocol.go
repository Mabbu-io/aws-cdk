package awsec2


// Transport protocol for client VPN.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var clientVpnConnectionHandler iClientVpnConnectionHandler
//   var clientVpnUserBasedAuthentication clientVpnUserBasedAuthentication
//   var logGroup logGroup
//   var logStream logStream
//   var securityGroup securityGroup
//   var subnet subnet
//   var subnetFilter subnetFilter
//   var vpc vpc
//
//   clientVpnEndpointProps := &clientVpnEndpointProps{
//   	cidr: jsii.String("cidr"),
//   	serverCertificateArn: jsii.String("serverCertificateArn"),
//   	vpc: vpc,
//
//   	// the properties below are optional
//   	authorizeAllUsersToVpcCidr: jsii.Boolean(false),
//   	clientCertificateArn: jsii.String("clientCertificateArn"),
//   	clientConnectionHandler: clientVpnConnectionHandler,
//   	clientLoginBanner: jsii.String("clientLoginBanner"),
//   	description: jsii.String("description"),
//   	dnsServers: []*string{
//   		jsii.String("dnsServers"),
//   	},
//   	logging: jsii.Boolean(false),
//   	logGroup: logGroup,
//   	logStream: logStream,
//   	port: awscdk.Aws_ec2.vpnPort_HTTPS,
//   	securityGroups: []iSecurityGroup{
//   		securityGroup,
//   	},
//   	selfServicePortal: jsii.Boolean(false),
//   	sessionTimeout: awscdk.*Aws_ec2.clientVpnSessionTimeout_EIGHT_HOURS,
//   	splitTunnel: jsii.Boolean(false),
//   	transportProtocol: awscdk.*Aws_ec2.transportProtocol_TCP,
//   	userBasedAuthentication: clientVpnUserBasedAuthentication,
//   	vpcSubnets: &subnetSelection{
//   		availabilityZones: []*string{
//   			jsii.String("availabilityZones"),
//   		},
//   		onePerAz: jsii.Boolean(false),
//   		subnetFilters: []*subnetFilter{
//   			subnetFilter,
//   		},
//   		subnetGroupName: jsii.String("subnetGroupName"),
//   		subnets: []iSubnet{
//   			subnet,
//   		},
//   		subnetType: awscdk.*Aws_ec2.subnetType_PRIVATE_ISOLATED,
//   	},
//   }
//
type TransportProtocol string

const (
	// Transmission Control Protocol (TCP).
	TransportProtocol_TCP TransportProtocol = "TCP"
	// User Datagram Protocol (UDP).
	TransportProtocol_UDP TransportProtocol = "UDP"
)

