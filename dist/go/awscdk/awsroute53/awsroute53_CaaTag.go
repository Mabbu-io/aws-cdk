package awsroute53


// The CAA tag.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var hostedZone hostedZone
//
//   caaRecordProps := &caaRecordProps{
//   	values: []caaRecordValue{
//   		&caaRecordValue{
//   			flag: jsii.Number(123),
//   			tag: awscdk.Aws_route53.caaTag_ISSUE,
//   			value: jsii.String("value"),
//   		},
//   	},
//   	zone: hostedZone,
//
//   	// the properties below are optional
//   	comment: jsii.String("comment"),
//   	deleteExisting: jsii.Boolean(false),
//   	recordName: jsii.String("recordName"),
//   	ttl: cdk.duration.minutes(jsii.Number(30)),
//   }
//
type CaaTag string

const (
	// Explicity authorizes a single certificate authority to issue a certificate (any type) for the hostname.
	CaaTag_ISSUE CaaTag = "ISSUE"
	// Explicity authorizes a single certificate authority to issue a wildcard certificate (and only wildcard) for the hostname.
	CaaTag_ISSUEWILD CaaTag = "ISSUEWILD"
	// Specifies a URL to which a certificate authority may report policy violations.
	CaaTag_IODEF CaaTag = "IODEF"
)

