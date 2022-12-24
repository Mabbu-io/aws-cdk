package awssesactions


// The type of invocation to use for a Lambda Action.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var function_ function
//   var topic topic
//
//   lambda := awscdk.Aws_ses_actions.NewLambda(&lambdaProps{
//   	function: function_,
//
//   	// the properties below are optional
//   	invocationType: awscdk.*Aws_ses_actions.lambdaInvocationType_EVENT,
//   	topic: topic,
//   })
//
type LambdaInvocationType string

const (
	// The function will be invoked asynchronously.
	LambdaInvocationType_EVENT LambdaInvocationType = "EVENT"
	// The function will be invoked sychronously.
	//
	// Use RequestResponse only when
	// you want to make a mail flow decision, such as whether to stop the receipt
	// rule or the receipt rule set.
	LambdaInvocationType_REQUEST_RESPONSE LambdaInvocationType = "REQUEST_RESPONSE"
)

