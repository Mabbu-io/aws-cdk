// Version 2 of the AWS Cloud Development Kit library
package awscdk


// Type hints for resolved values.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var value interface{}
//
//   intrinsic := cdk.NewIntrinsic(value, &intrinsicProps{
//   	stackTrace: jsii.Boolean(false),
//   	typeHint: cdk.resolutionTypeHint_STRING,
//   })
//
type ResolutionTypeHint string

const (
	// This value is expected to resolve to a String.
	ResolutionTypeHint_STRING ResolutionTypeHint = "STRING"
	// This value is expected to resolve to a Number.
	ResolutionTypeHint_NUMBER ResolutionTypeHint = "NUMBER"
	// This value is expected to resolve to a String List.
	ResolutionTypeHint_STRING_LIST ResolutionTypeHint = "STRING_LIST"
)

