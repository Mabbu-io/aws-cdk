package awsrds

import (
	"github.com/aws/aws-cdk-go/awscdk/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/awssecretsmanager"
)

// Properties that describe an existing cluster instance.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var secret secret
//   var securityGroup securityGroup
//
//   serverlessClusterAttributes := &serverlessClusterAttributes{
//   	clusterIdentifier: jsii.String("clusterIdentifier"),
//
//   	// the properties below are optional
//   	clusterEndpointAddress: jsii.String("clusterEndpointAddress"),
//   	port: jsii.Number(123),
//   	readerEndpointAddress: jsii.String("readerEndpointAddress"),
//   	secret: secret,
//   	securityGroups: []iSecurityGroup{
//   		securityGroup,
//   	},
//   }
//
type ServerlessClusterAttributes struct {
	// Identifier for the cluster.
	ClusterIdentifier *string `field:"required" json:"clusterIdentifier" yaml:"clusterIdentifier"`
	// Cluster endpoint address.
	ClusterEndpointAddress *string `field:"optional" json:"clusterEndpointAddress" yaml:"clusterEndpointAddress"`
	// The database port.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// Reader endpoint address.
	ReaderEndpointAddress *string `field:"optional" json:"readerEndpointAddress" yaml:"readerEndpointAddress"`
	// The secret attached to the database cluster.
	Secret awssecretsmanager.ISecret `field:"optional" json:"secret" yaml:"secret"`
	// The security groups of the database cluster.
	SecurityGroups *[]awsec2.ISecurityGroup `field:"optional" json:"securityGroups" yaml:"securityGroups"`
}

