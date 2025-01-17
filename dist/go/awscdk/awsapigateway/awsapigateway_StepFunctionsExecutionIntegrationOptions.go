package awsapigateway

import (
	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
)

// Options when configuring Step Functions synchronous integration with Rest API.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var role role
//   var vpcLink vpcLink
//
//   stepFunctionsExecutionIntegrationOptions := &stepFunctionsExecutionIntegrationOptions{
//   	authorizer: jsii.Boolean(false),
//   	cacheKeyParameters: []*string{
//   		jsii.String("cacheKeyParameters"),
//   	},
//   	cacheNamespace: jsii.String("cacheNamespace"),
//   	connectionType: awscdk.Aws_apigateway.connectionType_INTERNET,
//   	contentHandling: awscdk.*Aws_apigateway.contentHandling_CONVERT_TO_BINARY,
//   	credentialsPassthrough: jsii.Boolean(false),
//   	credentialsRole: role,
//   	headers: jsii.Boolean(false),
//   	integrationResponses: []integrationResponse{
//   		&integrationResponse{
//   			statusCode: jsii.String("statusCode"),
//
//   			// the properties below are optional
//   			contentHandling: awscdk.*Aws_apigateway.*contentHandling_CONVERT_TO_BINARY,
//   			responseParameters: map[string]*string{
//   				"responseParametersKey": jsii.String("responseParameters"),
//   			},
//   			responseTemplates: map[string]*string{
//   				"responseTemplatesKey": jsii.String("responseTemplates"),
//   			},
//   			selectionPattern: jsii.String("selectionPattern"),
//   		},
//   	},
//   	passthroughBehavior: awscdk.*Aws_apigateway.passthroughBehavior_WHEN_NO_MATCH,
//   	path: jsii.Boolean(false),
//   	querystring: jsii.Boolean(false),
//   	requestContext: &requestContext{
//   		accountId: jsii.Boolean(false),
//   		apiId: jsii.Boolean(false),
//   		apiKey: jsii.Boolean(false),
//   		authorizerPrincipalId: jsii.Boolean(false),
//   		caller: jsii.Boolean(false),
//   		cognitoAuthenticationProvider: jsii.Boolean(false),
//   		cognitoAuthenticationType: jsii.Boolean(false),
//   		cognitoIdentityId: jsii.Boolean(false),
//   		cognitoIdentityPoolId: jsii.Boolean(false),
//   		httpMethod: jsii.Boolean(false),
//   		requestId: jsii.Boolean(false),
//   		resourceId: jsii.Boolean(false),
//   		resourcePath: jsii.Boolean(false),
//   		sourceIp: jsii.Boolean(false),
//   		stage: jsii.Boolean(false),
//   		user: jsii.Boolean(false),
//   		userAgent: jsii.Boolean(false),
//   		userArn: jsii.Boolean(false),
//   	},
//   	requestParameters: map[string]*string{
//   		"requestParametersKey": jsii.String("requestParameters"),
//   	},
//   	requestTemplates: map[string]*string{
//   		"requestTemplatesKey": jsii.String("requestTemplates"),
//   	},
//   	timeout: cdk.duration.minutes(jsii.Number(30)),
//   	vpcLink: vpcLink,
//   }
//
type StepFunctionsExecutionIntegrationOptions struct {
	// A list of request parameters whose values are to be cached.
	//
	// It determines
	// request parameters that will make it into the cache key.
	CacheKeyParameters *[]*string `field:"optional" json:"cacheKeyParameters" yaml:"cacheKeyParameters"`
	// An API-specific tag group of related cached parameters.
	CacheNamespace *string `field:"optional" json:"cacheNamespace" yaml:"cacheNamespace"`
	// The type of network connection to the integration endpoint.
	ConnectionType ConnectionType `field:"optional" json:"connectionType" yaml:"connectionType"`
	// Specifies how to handle request payload content type conversions.
	ContentHandling ContentHandling `field:"optional" json:"contentHandling" yaml:"contentHandling"`
	// Requires that the caller's identity be passed through from the request.
	CredentialsPassthrough *bool `field:"optional" json:"credentialsPassthrough" yaml:"credentialsPassthrough"`
	// An IAM role that API Gateway assumes.
	//
	// Mutually exclusive with `credentialsPassThrough`.
	CredentialsRole awsiam.IRole `field:"optional" json:"credentialsRole" yaml:"credentialsRole"`
	// The response that API Gateway provides after a method's backend completes processing a request.
	//
	// API Gateway intercepts the response from the
	// backend so that you can control how API Gateway surfaces backend
	// responses. For example, you can map the backend status codes to codes
	// that you define.
	IntegrationResponses *[]*IntegrationResponse `field:"optional" json:"integrationResponses" yaml:"integrationResponses"`
	// Specifies the pass-through behavior for incoming requests based on the Content-Type header in the request, and the available mapping templates specified as the requestTemplates property on the Integration resource.
	//
	// There are three valid values: WHEN_NO_MATCH, WHEN_NO_TEMPLATES, and
	// NEVER.
	PassthroughBehavior PassthroughBehavior `field:"optional" json:"passthroughBehavior" yaml:"passthroughBehavior"`
	// The request parameters that API Gateway sends with the backend request.
	//
	// Specify request parameters as key-value pairs (string-to-string
	// mappings), with a destination as the key and a source as the value.
	//
	// Specify the destination by using the following pattern
	// integration.request.location.name, where location is querystring, path,
	// or header, and name is a valid, unique parameter name.
	//
	// The source must be an existing method request parameter or a static
	// value. You must enclose static values in single quotation marks and
	// pre-encode these values based on their destination in the request.
	RequestParameters *map[string]*string `field:"optional" json:"requestParameters" yaml:"requestParameters"`
	// A map of Apache Velocity templates that are applied on the request payload.
	//
	// The template that API Gateway uses is based on the value of the
	// Content-Type header that's sent by the client. The content type value is
	// the key, and the template is the value (specified as a string), such as
	// the following snippet:
	//
	// ```
	//    { "application/json": "{ \"statusCode\": 200 }" }
	// ```.
	// See: http://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-mapping-template-reference.html
	//
	RequestTemplates *map[string]*string `field:"optional" json:"requestTemplates" yaml:"requestTemplates"`
	// The maximum amount of time an integration will run before it returns without a response.
	//
	// Must be between 50 milliseconds and 29 seconds.
	Timeout awscdk.Duration `field:"optional" json:"timeout" yaml:"timeout"`
	// The VpcLink used for the integration.
	//
	// Required if connectionType is VPC_LINK.
	VpcLink IVpcLink `field:"optional" json:"vpcLink" yaml:"vpcLink"`
	// If the whole authorizer object, including custom context values should be in the execution input.
	//
	// The execution input will include a new key `authorizer`:
	//
	// {
	//    "body": {},
	//    "authorizer": {
	//      "key": "value"
	//    }
	// }.
	Authorizer *bool `field:"optional" json:"authorizer" yaml:"authorizer"`
	// Check if header is to be included inside the execution input.
	//
	// The execution input will include a new key `headers`:
	//
	// {
	//    "body": {},
	//    "headers": {
	//       "header1": "value",
	//       "header2": "value"
	//    }
	// }.
	Headers *bool `field:"optional" json:"headers" yaml:"headers"`
	// Check if path is to be included inside the execution input.
	//
	// The execution input will include a new key `path`:
	//
	// {
	//    "body": {},
	//    "path": {
	//      "resourceName": "resourceValue"
	//    }
	// }.
	Path *bool `field:"optional" json:"path" yaml:"path"`
	// Check if querystring is to be included inside the execution input.
	//
	// The execution input will include a new key `queryString`:
	//
	// {
	//    "body": {},
	//    "querystring": {
	//      "key": "value"
	//    }
	// }.
	Querystring *bool `field:"optional" json:"querystring" yaml:"querystring"`
	// Which details of the incoming request must be passed onto the underlying state machine, such as, account id, user identity, request id, etc.
	//
	// The execution input will include a new key `requestContext`:
	//
	// {
	//    "body": {},
	//    "requestContext": {
	//        "key": "value"
	//    }
	// }.
	RequestContext *RequestContext `field:"optional" json:"requestContext" yaml:"requestContext"`
}

