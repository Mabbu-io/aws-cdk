package awss3


// All http request methods.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   corsRule := &corsRule{
//   	allowedMethods: []httpMethods{
//   		awscdk.Aws_s3.*httpMethods_GET,
//   	},
//   	allowedOrigins: []*string{
//   		jsii.String("allowedOrigins"),
//   	},
//
//   	// the properties below are optional
//   	allowedHeaders: []*string{
//   		jsii.String("allowedHeaders"),
//   	},
//   	exposedHeaders: []*string{
//   		jsii.String("exposedHeaders"),
//   	},
//   	id: jsii.String("id"),
//   	maxAge: jsii.Number(123),
//   }
//
type HttpMethods string

const (
	// The GET method requests a representation of the specified resource.
	HttpMethods_GET HttpMethods = "GET"
	// The PUT method replaces all current representations of the target resource with the request payload.
	HttpMethods_PUT HttpMethods = "PUT"
	// The HEAD method asks for a response identical to that of a GET request, but without the response body.
	HttpMethods_HEAD HttpMethods = "HEAD"
	// The POST method is used to submit an entity to the specified resource, often causing a change in state or side effects on the server.
	HttpMethods_POST HttpMethods = "POST"
	// The DELETE method deletes the specified resource.
	HttpMethods_DELETE HttpMethods = "DELETE"
)

