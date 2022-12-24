package awsecs


// Log Message Format.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var secret secret
//
//   splunkLogDriver := awscdk.Aws_ecs.NewSplunkLogDriver(&splunkLogDriverProps{
//   	secretToken: secret,
//   	url: jsii.String("url"),
//
//   	// the properties below are optional
//   	caName: jsii.String("caName"),
//   	caPath: jsii.String("caPath"),
//   	env: []*string{
//   		jsii.String("env"),
//   	},
//   	envRegex: jsii.String("envRegex"),
//   	format: awscdk.*Aws_ecs.splunkLogFormat_INLINE,
//   	gzip: jsii.Boolean(false),
//   	gzipLevel: jsii.Number(123),
//   	index: jsii.String("index"),
//   	insecureSkipVerify: jsii.String("insecureSkipVerify"),
//   	labels: []*string{
//   		jsii.String("labels"),
//   	},
//   	source: jsii.String("source"),
//   	sourceType: jsii.String("sourceType"),
//   	tag: jsii.String("tag"),
//   	verifyConnection: jsii.Boolean(false),
//   })
//
type SplunkLogFormat string

const (
	SplunkLogFormat_INLINE SplunkLogFormat = "INLINE"
	SplunkLogFormat_JSON SplunkLogFormat = "JSON"
	SplunkLogFormat_RAW SplunkLogFormat = "RAW"
)

