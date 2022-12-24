package awsevents


// Supported HTTP operations.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var httpParameter httpParameter
//   var secretValue secretValue
//
//   oAuthAuthorizationProps := &oAuthAuthorizationProps{
//   	authorizationEndpoint: jsii.String("authorizationEndpoint"),
//   	clientId: jsii.String("clientId"),
//   	clientSecret: secretValue,
//   	httpMethod: awscdk.Aws_events.httpMethod_POST,
//
//   	// the properties below are optional
//   	bodyParameters: map[string]*httpParameter{
//   		"bodyParametersKey": httpParameter,
//   	},
//   	headerParameters: map[string]*httpParameter{
//   		"headerParametersKey": httpParameter,
//   	},
//   	queryStringParameters: map[string]*httpParameter{
//   		"queryStringParametersKey": httpParameter,
//   	},
//   }
//
type HttpMethod string

const (
	// POST.
	HttpMethod_POST HttpMethod = "POST"
	// GET.
	HttpMethod_GET HttpMethod = "GET"
	// HEAD.
	HttpMethod_HEAD HttpMethod = "HEAD"
	// OPTIONS.
	HttpMethod_OPTIONS HttpMethod = "OPTIONS"
	// PUT.
	HttpMethod_PUT HttpMethod = "PUT"
	// PATCH.
	HttpMethod_PATCH HttpMethod = "PATCH"
	// DELETE.
	HttpMethod_DELETE HttpMethod = "DELETE"
)

