package awslambda


// All http request methods.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var function_ function
//
//   functionUrlProps := &functionUrlProps{
//   	function: function_,
//
//   	// the properties below are optional
//   	authType: awscdk.Aws_lambda.functionUrlAuthType_AWS_IAM,
//   	cors: &functionUrlCorsOptions{
//   		allowCredentials: jsii.Boolean(false),
//   		allowedHeaders: []*string{
//   			jsii.String("allowedHeaders"),
//   		},
//   		allowedMethods: []httpMethod{
//   			awscdk.*Aws_lambda.*httpMethod_GET,
//   		},
//   		allowedOrigins: []*string{
//   			jsii.String("allowedOrigins"),
//   		},
//   		exposedHeaders: []*string{
//   			jsii.String("exposedHeaders"),
//   		},
//   		maxAge: cdk.duration.minutes(jsii.Number(30)),
//   	},
//   }
//
type HttpMethod string

const (
	// The GET method requests a representation of the specified resource.
	HttpMethod_GET HttpMethod = "GET"
	// The PUT method replaces all current representations of the target resource with the request payload.
	HttpMethod_PUT HttpMethod = "PUT"
	// The HEAD method asks for a response identical to that of a GET request, but without the response body.
	HttpMethod_HEAD HttpMethod = "HEAD"
	// The POST method is used to submit an entity to the specified resource, often causing a change in state or side effects on the server.
	HttpMethod_POST HttpMethod = "POST"
	// The DELETE method deletes the specified resource.
	HttpMethod_DELETE HttpMethod = "DELETE"
	// The PATCH method applies partial modifications to a resource.
	HttpMethod_PATCH HttpMethod = "PATCH"
	// The OPTIONS method describes the communication options for the target resource.
	HttpMethod_OPTIONS HttpMethod = "OPTIONS"
	// The wildcard entry to allow all methods.
	HttpMethod_ALL HttpMethod = "ALL"
)

