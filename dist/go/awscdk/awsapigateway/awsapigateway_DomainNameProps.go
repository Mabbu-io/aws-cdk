package awsapigateway

import (
	"github.com/aws/aws-cdk-go/awscdk/awscertificatemanager"
)

// Example:
//   var acmCertificateForExampleCom interface{}
//   var restApi restApi
//
//
//   apigateway.NewDomainName(this, jsii.String("custom-domain"), &domainNameProps{
//   	domainName: jsii.String("example.com"),
//   	certificate: acmCertificateForExampleCom,
//   	mapping: restApi,
//   	basePath: jsii.String("orders/v1/api"),
//   })
//
type DomainNameProps struct {
	// The reference to an AWS-managed certificate for use by the edge-optimized endpoint for the domain name.
	//
	// For "EDGE" domain names, the certificate
	// needs to be in the US East (N. Virginia) region.
	Certificate awscertificatemanager.ICertificate `field:"required" json:"certificate" yaml:"certificate"`
	// The custom domain name for your API.
	//
	// Uppercase letters are not supported.
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// The base path name that callers of the API must provide in the URL after the domain name (e.g. `example.com/base-path`). If you specify this property, it can't be an empty string.
	BasePath *string `field:"optional" json:"basePath" yaml:"basePath"`
	// The type of endpoint for this DomainName.
	EndpointType EndpointType `field:"optional" json:"endpointType" yaml:"endpointType"`
	// The mutual TLS authentication configuration for a custom domain name.
	Mtls *MTLSConfig `field:"optional" json:"mtls" yaml:"mtls"`
	// The Transport Layer Security (TLS) version + cipher suite for this domain name.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-domainname.html
	//
	SecurityPolicy SecurityPolicy `field:"optional" json:"securityPolicy" yaml:"securityPolicy"`
	// If specified, all requests to this domain will be mapped to the production deployment of this API.
	//
	// If you wish to map this domain to multiple APIs
	// with different base paths, use `addBasePathMapping` or `addApiMapping`.
	Mapping IRestApi `field:"optional" json:"mapping" yaml:"mapping"`
}

