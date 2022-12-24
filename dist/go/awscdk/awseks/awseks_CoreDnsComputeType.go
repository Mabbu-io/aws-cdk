package awseks


// The type of compute resources to use for CoreDNS.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var albControllerVersion albControllerVersion
//   var endpointAccess endpointAccess
//   var key key
//   var kubernetesVersion kubernetesVersion
//   var layerVersion layerVersion
//   var policy interface{}
//   var role role
//   var securityGroup securityGroup
//   var size size
//   var subnet subnet
//   var subnetFilter subnetFilter
//   var vpc vpc
//
//   clusterOptions := &clusterOptions{
//   	version: kubernetesVersion,
//
//   	// the properties below are optional
//   	albController: &albControllerOptions{
//   		version: albControllerVersion,
//
//   		// the properties below are optional
//   		policy: policy,
//   		repository: jsii.String("repository"),
//   	},
//   	awscliLayer: layerVersion,
//   	clusterHandlerEnvironment: map[string]*string{
//   		"clusterHandlerEnvironmentKey": jsii.String("clusterHandlerEnvironment"),
//   	},
//   	clusterHandlerSecurityGroup: securityGroup,
//   	clusterLogging: []clusterLoggingTypes{
//   		awscdk.Aws_eks.*clusterLoggingTypes_API,
//   	},
//   	clusterName: jsii.String("clusterName"),
//   	coreDnsComputeType: awscdk.*Aws_eks.coreDnsComputeType_EC2,
//   	endpointAccess: endpointAccess,
//   	kubectlEnvironment: map[string]*string{
//   		"kubectlEnvironmentKey": jsii.String("kubectlEnvironment"),
//   	},
//   	kubectlLayer: layerVersion,
//   	kubectlMemory: size,
//   	mastersRole: role,
//   	onEventLayer: layerVersion,
//   	outputClusterName: jsii.Boolean(false),
//   	outputConfigCommand: jsii.Boolean(false),
//   	outputMastersRoleArn: jsii.Boolean(false),
//   	placeClusterHandlerInVpc: jsii.Boolean(false),
//   	prune: jsii.Boolean(false),
//   	role: role,
//   	secretsEncryptionKey: key,
//   	securityGroup: securityGroup,
//   	serviceIpv4Cidr: jsii.String("serviceIpv4Cidr"),
//   	vpc: vpc,
//   	vpcSubnets: []subnetSelection{
//   		&subnetSelection{
//   			availabilityZones: []*string{
//   				jsii.String("availabilityZones"),
//   			},
//   			onePerAz: jsii.Boolean(false),
//   			subnetFilters: []*subnetFilter{
//   				subnetFilter,
//   			},
//   			subnetGroupName: jsii.String("subnetGroupName"),
//   			subnets: []iSubnet{
//   				subnet,
//   			},
//   			subnetType: awscdk.Aws_ec2.subnetType_PRIVATE_ISOLATED,
//   		},
//   	},
//   }
//
type CoreDnsComputeType string

const (
	// Deploy CoreDNS on EC2 instances.
	CoreDnsComputeType_EC2 CoreDnsComputeType = "EC2"
	// Deploy CoreDNS on Fargate-managed instances.
	CoreDnsComputeType_FARGATE CoreDnsComputeType = "FARGATE"
)

