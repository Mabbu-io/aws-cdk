package cxapi


// The type of subnet group.
//
// Same as SubnetType in the @aws-cdk/aws-ec2 package,
// but we can't use that because of cyclical dependencies.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   vpcContextResponse := &vpcContextResponse{
//   	availabilityZones: []*string{
//   		jsii.String("availabilityZones"),
//   	},
//   	vpcId: jsii.String("vpcId"),
//
//   	// the properties below are optional
//   	isolatedSubnetIds: []*string{
//   		jsii.String("isolatedSubnetIds"),
//   	},
//   	isolatedSubnetNames: []*string{
//   		jsii.String("isolatedSubnetNames"),
//   	},
//   	isolatedSubnetRouteTableIds: []*string{
//   		jsii.String("isolatedSubnetRouteTableIds"),
//   	},
//   	privateSubnetIds: []*string{
//   		jsii.String("privateSubnetIds"),
//   	},
//   	privateSubnetNames: []*string{
//   		jsii.String("privateSubnetNames"),
//   	},
//   	privateSubnetRouteTableIds: []*string{
//   		jsii.String("privateSubnetRouteTableIds"),
//   	},
//   	publicSubnetIds: []*string{
//   		jsii.String("publicSubnetIds"),
//   	},
//   	publicSubnetNames: []*string{
//   		jsii.String("publicSubnetNames"),
//   	},
//   	publicSubnetRouteTableIds: []*string{
//   		jsii.String("publicSubnetRouteTableIds"),
//   	},
//   	region: jsii.String("region"),
//   	subnetGroups: []vpcSubnetGroup{
//   		&vpcSubnetGroup{
//   			name: jsii.String("name"),
//   			subnets: []vpcSubnet{
//   				&vpcSubnet{
//   					availabilityZone: jsii.String("availabilityZone"),
//   					routeTableId: jsii.String("routeTableId"),
//   					subnetId: jsii.String("subnetId"),
//
//   					// the properties below are optional
//   					cidr: jsii.String("cidr"),
//   				},
//   			},
//   			type: awscdk.Cx_api.vpcSubnetGroupType_PUBLIC,
//   		},
//   	},
//   	vpcCidrBlock: jsii.String("vpcCidrBlock"),
//   	vpnGatewayId: jsii.String("vpnGatewayId"),
//   }
//
type VpcSubnetGroupType string

const (
	// Public subnet group type.
	VpcSubnetGroupType_PUBLIC VpcSubnetGroupType = "PUBLIC"
	// Private subnet group type.
	VpcSubnetGroupType_PRIVATE VpcSubnetGroupType = "PRIVATE"
	// Isolated subnet group type.
	VpcSubnetGroupType_ISOLATED VpcSubnetGroupType = "ISOLATED"
)

