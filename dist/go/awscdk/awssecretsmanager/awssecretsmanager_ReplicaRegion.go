package awssecretsmanager

import (
	"github.com/aws/aws-cdk-go/awscdk/awskms"
)

// Secret replica region.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var key key
//
//   replicaRegion := &replicaRegion{
//   	region: jsii.String("region"),
//
//   	// the properties below are optional
//   	encryptionKey: key,
//   }
//
type ReplicaRegion struct {
	// The name of the region.
	Region *string `field:"required" json:"region" yaml:"region"`
	// The customer-managed encryption key to use for encrypting the secret value.
	EncryptionKey awskms.IKey `field:"optional" json:"encryptionKey" yaml:"encryptionKey"`
}

