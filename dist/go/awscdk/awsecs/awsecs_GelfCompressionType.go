package awsecs


// The type of compression the GELF driver uses to compress each log message.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   gelfLogDriver := awscdk.Aws_ecs.NewGelfLogDriver(&gelfLogDriverProps{
//   	address: jsii.String("address"),
//
//   	// the properties below are optional
//   	compressionLevel: jsii.Number(123),
//   	compressionType: awscdk.*Aws_ecs.gelfCompressionType_GZIP,
//   	env: []*string{
//   		jsii.String("env"),
//   	},
//   	envRegex: jsii.String("envRegex"),
//   	labels: []*string{
//   		jsii.String("labels"),
//   	},
//   	tag: jsii.String("tag"),
//   	tcpMaxReconnect: jsii.Number(123),
//   	tcpReconnectDelay: cdk.duration.minutes(jsii.Number(30)),
//   })
//
type GelfCompressionType string

const (
	GelfCompressionType_GZIP GelfCompressionType = "GZIP"
	GelfCompressionType_ZLIB GelfCompressionType = "ZLIB"
	GelfCompressionType_NONE GelfCompressionType = "NONE"
)

