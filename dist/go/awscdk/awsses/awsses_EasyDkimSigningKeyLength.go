package awsses


// The signing key length for Easy DKIM.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dkimIdentityConfig := &dkimIdentityConfig{
//   	domainSigningPrivateKey: jsii.String("domainSigningPrivateKey"),
//   	domainSigningSelector: jsii.String("domainSigningSelector"),
//   	nextSigningKeyLength: awscdk.Aws_ses.easyDkimSigningKeyLength_RSA_1024_BIT,
//   }
//
type EasyDkimSigningKeyLength string

const (
	// RSA 1024-bit.
	EasyDkimSigningKeyLength_RSA_1024_BIT EasyDkimSigningKeyLength = "RSA_1024_BIT"
	// RSA 2048-bit.
	EasyDkimSigningKeyLength_RSA_2048_BIT EasyDkimSigningKeyLength = "RSA_2048_BIT"
)

