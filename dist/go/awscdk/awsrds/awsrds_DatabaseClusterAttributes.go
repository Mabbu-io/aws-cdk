package awsrds

import (
	"github.com/aws/aws-cdk-go/awscdk/awsec2"
)

// Properties that describe an existing cluster instance.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var clusterEngine iClusterEngine
//   var securityGroup securityGroup
//
//   databaseClusterAttributes := &databaseClusterAttributes{
//   	clusterIdentifier: jsii.String("clusterIdentifier"),
//
//   	// the properties below are optional
//   	clusterEndpointAddress: jsii.String("clusterEndpointAddress"),
//   	engine: clusterEngine,
//   	instanceEndpointAddresses: []*string{
//   		jsii.String("instanceEndpointAddresses"),
//   	},
//   	instanceIdentifiers: []*string{
//   		jsii.String("instanceIdentifiers"),
//   	},
//   	port: jsii.Number(123),
//   	readerEndpointAddress: jsii.String("readerEndpointAddress"),
//   	securityGroups: []iSecurityGroup{
//   		securityGroup,
//   	},
//   }
//
type DatabaseClusterAttributes struct {
	// Identifier for the cluster.
	ClusterIdentifier *string `field:"required" json:"clusterIdentifier" yaml:"clusterIdentifier"`
	// Cluster endpoint address.
	ClusterEndpointAddress *string `field:"optional" json:"clusterEndpointAddress" yaml:"clusterEndpointAddress"`
	// The engine of the existing Cluster.
	Engine IClusterEngine `field:"optional" json:"engine" yaml:"engine"`
	// Endpoint addresses of individual instances.
	InstanceEndpointAddresses *[]*string `field:"optional" json:"instanceEndpointAddresses" yaml:"instanceEndpointAddresses"`
	// Identifier for the instances.
	InstanceIdentifiers *[]*string `field:"optional" json:"instanceIdentifiers" yaml:"instanceIdentifiers"`
	// The database port.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// Reader endpoint address.
	ReaderEndpointAddress *string `field:"optional" json:"readerEndpointAddress" yaml:"readerEndpointAddress"`
	// The security groups of the database cluster.
	SecurityGroups *[]awsec2.ISecurityGroup `field:"optional" json:"securityGroups" yaml:"securityGroups"`
}

