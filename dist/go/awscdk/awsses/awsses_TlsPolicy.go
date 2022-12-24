package awsses


// The type of TLS policy for a receipt rule.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var receiptRule receiptRule
//   var receiptRuleAction iReceiptRuleAction
//   var receiptRuleSet receiptRuleSet
//
//   dropSpamReceiptRuleProps := &dropSpamReceiptRuleProps{
//   	ruleSet: receiptRuleSet,
//
//   	// the properties below are optional
//   	actions: []*iReceiptRuleAction{
//   		receiptRuleAction,
//   	},
//   	after: receiptRule,
//   	enabled: jsii.Boolean(false),
//   	receiptRuleName: jsii.String("receiptRuleName"),
//   	recipients: []*string{
//   		jsii.String("recipients"),
//   	},
//   	scanEnabled: jsii.Boolean(false),
//   	tlsPolicy: awscdk.Aws_ses.tlsPolicy_OPTIONAL,
//   }
//
type TlsPolicy string

const (
	// Do not check for TLS.
	TlsPolicy_OPTIONAL TlsPolicy = "OPTIONAL"
	// Bounce emails that are not received over TLS.
	TlsPolicy_REQUIRE TlsPolicy = "REQUIRE"
)

