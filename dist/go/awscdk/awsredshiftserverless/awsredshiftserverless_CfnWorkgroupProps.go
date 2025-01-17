package awsredshiftserverless

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnWorkgroup`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnWorkgroupProps := &cfnWorkgroupProps{
//   	workgroupName: jsii.String("workgroupName"),
//
//   	// the properties below are optional
//   	baseCapacity: jsii.Number(123),
//   	configParameters: []interface{}{
//   		&configParameterProperty{
//   			parameterKey: jsii.String("parameterKey"),
//   			parameterValue: jsii.String("parameterValue"),
//   		},
//   	},
//   	enhancedVpcRouting: jsii.Boolean(false),
//   	namespaceName: jsii.String("namespaceName"),
//   	publiclyAccessible: jsii.Boolean(false),
//   	securityGroupIds: []*string{
//   		jsii.String("securityGroupIds"),
//   	},
//   	subnetIds: []*string{
//   		jsii.String("subnetIds"),
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	workgroup: &workgroupProperty{
//   		baseCapacity: jsii.Number(123),
//   		configParameters: []interface{}{
//   			&configParameterProperty{
//   				parameterKey: jsii.String("parameterKey"),
//   				parameterValue: jsii.String("parameterValue"),
//   			},
//   		},
//   		creationDate: jsii.String("creationDate"),
//   		endpoint: &endpointProperty{
//   			address: jsii.String("address"),
//   			port: jsii.Number(123),
//   			vpcEndpoints: []interface{}{
//   				&vpcEndpointProperty{
//   					networkInterfaces: []interface{}{
//   						&networkInterfaceProperty{
//   							availabilityZone: jsii.String("availabilityZone"),
//   							networkInterfaceId: jsii.String("networkInterfaceId"),
//   							privateIpAddress: jsii.String("privateIpAddress"),
//   							subnetId: jsii.String("subnetId"),
//   						},
//   					},
//   					vpcEndpointId: jsii.String("vpcEndpointId"),
//   					vpcId: jsii.String("vpcId"),
//   				},
//   			},
//   		},
//   		enhancedVpcRouting: jsii.Boolean(false),
//   		namespaceName: jsii.String("namespaceName"),
//   		publiclyAccessible: jsii.Boolean(false),
//   		securityGroupIds: []*string{
//   			jsii.String("securityGroupIds"),
//   		},
//   		status: jsii.String("status"),
//   		subnetIds: []*string{
//   			jsii.String("subnetIds"),
//   		},
//   		workgroupArn: jsii.String("workgroupArn"),
//   		workgroupId: jsii.String("workgroupId"),
//   		workgroupName: jsii.String("workgroupName"),
//   	},
//   }
//
type CfnWorkgroupProps struct {
	// `AWS::RedshiftServerless::Workgroup.WorkgroupName`.
	WorkgroupName *string `field:"required" json:"workgroupName" yaml:"workgroupName"`
	// `AWS::RedshiftServerless::Workgroup.BaseCapacity`.
	BaseCapacity *float64 `field:"optional" json:"baseCapacity" yaml:"baseCapacity"`
	// `AWS::RedshiftServerless::Workgroup.ConfigParameters`.
	ConfigParameters interface{} `field:"optional" json:"configParameters" yaml:"configParameters"`
	// `AWS::RedshiftServerless::Workgroup.EnhancedVpcRouting`.
	EnhancedVpcRouting interface{} `field:"optional" json:"enhancedVpcRouting" yaml:"enhancedVpcRouting"`
	// `AWS::RedshiftServerless::Workgroup.NamespaceName`.
	NamespaceName *string `field:"optional" json:"namespaceName" yaml:"namespaceName"`
	// `AWS::RedshiftServerless::Workgroup.PubliclyAccessible`.
	PubliclyAccessible interface{} `field:"optional" json:"publiclyAccessible" yaml:"publiclyAccessible"`
	// `AWS::RedshiftServerless::Workgroup.SecurityGroupIds`.
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	// `AWS::RedshiftServerless::Workgroup.SubnetIds`.
	SubnetIds *[]*string `field:"optional" json:"subnetIds" yaml:"subnetIds"`
	// `AWS::RedshiftServerless::Workgroup.Tags`.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// `AWS::RedshiftServerless::Workgroup.Workgroup`.
	Workgroup interface{} `field:"optional" json:"workgroup" yaml:"workgroup"`
}

