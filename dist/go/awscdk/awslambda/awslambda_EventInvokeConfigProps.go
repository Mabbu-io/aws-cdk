package awslambda

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for an EventInvokeConfig.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var destination iDestination
//   var function_ function
//
//   eventInvokeConfigProps := &eventInvokeConfigProps{
//   	function: function_,
//
//   	// the properties below are optional
//   	maxEventAge: cdk.duration.minutes(jsii.Number(30)),
//   	onFailure: destination,
//   	onSuccess: destination,
//   	qualifier: jsii.String("qualifier"),
//   	retryAttempts: jsii.Number(123),
//   }
//
type EventInvokeConfigProps struct {
	// The maximum age of a request that Lambda sends to a function for processing.
	//
	// Minimum: 60 seconds
	// Maximum: 6 hours.
	MaxEventAge awscdk.Duration `field:"optional" json:"maxEventAge" yaml:"maxEventAge"`
	// The destination for failed invocations.
	OnFailure IDestination `field:"optional" json:"onFailure" yaml:"onFailure"`
	// The destination for successful invocations.
	OnSuccess IDestination `field:"optional" json:"onSuccess" yaml:"onSuccess"`
	// The maximum number of times to retry when the function returns an error.
	//
	// Minimum: 0
	// Maximum: 2.
	RetryAttempts *float64 `field:"optional" json:"retryAttempts" yaml:"retryAttempts"`
	// The Lambda function.
	Function IFunction `field:"required" json:"function" yaml:"function"`
	// The qualifier.
	Qualifier *string `field:"optional" json:"qualifier" yaml:"qualifier"`
}

