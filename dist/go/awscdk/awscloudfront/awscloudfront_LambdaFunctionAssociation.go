package awscloudfront

import (
	"github.com/aws/aws-cdk-go/awscdk/awslambda"
)

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var version version
//
//   lambdaFunctionAssociation := &lambdaFunctionAssociation{
//   	eventType: awscdk.Aws_cloudfront.lambdaEdgeEventType_ORIGIN_REQUEST,
//   	lambdaFunction: version,
//
//   	// the properties below are optional
//   	includeBody: jsii.Boolean(false),
//   }
//
type LambdaFunctionAssociation struct {
	// The lambda event type defines at which event the lambda is called during the request lifecycle.
	EventType LambdaEdgeEventType `field:"required" json:"eventType" yaml:"eventType"`
	// A version of the lambda to associate.
	LambdaFunction awslambda.IVersion `field:"required" json:"lambdaFunction" yaml:"lambdaFunction"`
	// Allows a Lambda function to have read access to the body content.
	//
	// Only valid for "request" event types (`ORIGIN_REQUEST` or `VIEWER_REQUEST`).
	// See https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/lambda-include-body-access.html
	IncludeBody *bool `field:"optional" json:"includeBody" yaml:"includeBody"`
}

