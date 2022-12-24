package awslambda


// The type of destination.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   destinationOptions := &destinationOptions{
//   	type: awscdk.Aws_lambda.destinationType_FAILURE,
//   }
//
type DestinationType string

const (
	// Failure.
	DestinationType_FAILURE DestinationType = "FAILURE"
	// Success.
	DestinationType_SUCCESS DestinationType = "SUCCESS"
)

