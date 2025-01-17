package awscognito

import (
	"github.com/aws/aws-cdk-go/awscdk/awscertificatemanager"
)

// Options while specifying custom domain.
//
// Example:
//   pool := cognito.NewUserPool(this, jsii.String("Pool"))
//
//   pool.addDomain(jsii.String("CognitoDomain"), &userPoolDomainOptions{
//   	cognitoDomain: &cognitoDomainOptions{
//   		domainPrefix: jsii.String("my-awesome-app"),
//   	},
//   })
//
//   certificateArn := "arn:aws:acm:us-east-1:123456789012:certificate/11-3336f1-44483d-adc7-9cd375c5169d"
//
//   domainCert := certificatemanager.certificate.fromCertificateArn(this, jsii.String("domainCert"), certificateArn)
//   pool.addDomain(jsii.String("CustomDomain"), &userPoolDomainOptions{
//   	customDomain: &customDomainOptions{
//   		domainName: jsii.String("user.myapp.com"),
//   		certificate: domainCert,
//   	},
//   })
//
// See: https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-pools-add-custom-domain.html
//
type CustomDomainOptions struct {
	// The certificate to associate with this domain.
	Certificate awscertificatemanager.ICertificate `field:"required" json:"certificate" yaml:"certificate"`
	// The custom domain name that you would like to associate with this User Pool.
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
}

