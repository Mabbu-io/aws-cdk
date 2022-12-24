package awscertificatemanager


// Method used to assert ownership of the domain.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var hostedZone hostedZone
//
//   certificationValidationProps := &certificationValidationProps{
//   	hostedZone: hostedZone,
//   	hostedZones: map[string]iHostedZone{
//   		"hostedZonesKey": hostedZone,
//   	},
//   	method: awscdk.Aws_certificatemanager.validationMethod_EMAIL,
//   	validationDomains: map[string]*string{
//   		"validationDomainsKey": jsii.String("validationDomains"),
//   	},
//   }
//
type ValidationMethod string

const (
	// Send email to a number of email addresses associated with the domain.
	// See: https://docs.aws.amazon.com/acm/latest/userguide/gs-acm-validate-email.html
	//
	ValidationMethod_EMAIL ValidationMethod = "EMAIL"
	// Validate ownership by adding appropriate DNS records.
	// See: https://docs.aws.amazon.com/acm/latest/userguide/gs-acm-validate-dns.html
	//
	ValidationMethod_DNS ValidationMethod = "DNS"
)

