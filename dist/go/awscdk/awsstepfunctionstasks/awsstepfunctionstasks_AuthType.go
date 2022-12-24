package awsstepfunctionstasks


// The authentication method used to call the endpoint.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var resultSelector interface{}
//   var taskInput taskInput
//   var taskRole taskRole
//
//   callApiGatewayEndpointBaseProps := &callApiGatewayEndpointBaseProps{
//   	method: awscdk.Aws_stepfunctions_tasks.httpMethod_GET,
//
//   	// the properties below are optional
//   	apiPath: jsii.String("apiPath"),
//   	authType: awscdk.*Aws_stepfunctions_tasks.authType_NO_AUTH,
//   	comment: jsii.String("comment"),
//   	credentials: &credentials{
//   		role: taskRole,
//   	},
//   	headers: taskInput,
//   	heartbeat: cdk.duration.minutes(jsii.Number(30)),
//   	inputPath: jsii.String("inputPath"),
//   	integrationPattern: awscdk.Aws_stepfunctions.integrationPattern_REQUEST_RESPONSE,
//   	outputPath: jsii.String("outputPath"),
//   	queryParameters: taskInput,
//   	requestBody: taskInput,
//   	resultPath: jsii.String("resultPath"),
//   	resultSelector: map[string]interface{}{
//   		"resultSelectorKey": resultSelector,
//   	},
//   	timeout: cdk.*duration.minutes(jsii.Number(30)),
//   }
//
type AuthType string

const (
	// Call the API direclty with no authorization method.
	AuthType_NO_AUTH AuthType = "NO_AUTH"
	// Use the IAM role associated with the current state machine for authorization.
	AuthType_IAM_ROLE AuthType = "IAM_ROLE"
	// Use the resource policy of the API for authorization.
	AuthType_RESOURCE_POLICY AuthType = "RESOURCE_POLICY"
)

