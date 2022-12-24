package awsbackup


// An operation that is applied to a key-value pair.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tagCondition := &tagCondition{
//   	key: jsii.String("key"),
//   	value: jsii.String("value"),
//
//   	// the properties below are optional
//   	operation: awscdk.Aws_backup.tagOperation_STRING_EQUALS,
//   }
//
type TagOperation string

const (
	// StringEquals.
	TagOperation_STRING_EQUALS TagOperation = "STRING_EQUALS"
	// Dummy member.
	TagOperation_DUMMY TagOperation = "DUMMY"
)

