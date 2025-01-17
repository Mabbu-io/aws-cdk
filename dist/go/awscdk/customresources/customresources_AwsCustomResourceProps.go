package customresources

import (
	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/awslogs"
)

// Properties for AwsCustomResource.
//
// Note that at least onCreate, onUpdate or onDelete must be specified.
//
// Example:
//   getParameter := cr.NewAwsCustomResource(this, jsii.String("GetParameter"), &awsCustomResourceProps{
//   	onUpdate: &awsSdkCall{
//   		 // will also be called for a CREATE event
//   		service: jsii.String("SSM"),
//   		action: jsii.String("getParameter"),
//   		parameters: map[string]interface{}{
//   			"Name": jsii.String("my-parameter"),
//   			"WithDecryption": jsii.Boolean(true),
//   		},
//   		physicalResourceId: cr.physicalResourceId.of(date.now().toString()),
//   	},
//   	policy: cr.awsCustomResourcePolicy.fromSdkCalls(&sdkCallsPolicyOptions{
//   		resources: cr.*awsCustomResourcePolicy_ANY_RESOURCE(),
//   	}),
//   })
//
//   // Use the value in another construct with
//   getParameter.getResponseField(jsii.String("Parameter.Value"))
//
type AwsCustomResourceProps struct {
	// A name for the singleton Lambda function implementing this custom resource.
	//
	// The function name will remain the same after the first AwsCustomResource is created in a stack.
	FunctionName *string `field:"optional" json:"functionName" yaml:"functionName"`
	// Whether to install the latest AWS SDK v2. Allows to use the latest API calls documented at https://docs.aws.amazon.com/AWSJavaScriptSDK/latest/index.html.
	//
	// The installation takes around 60 seconds.
	InstallLatestAwsSdk *bool `field:"optional" json:"installLatestAwsSdk" yaml:"installLatestAwsSdk"`
	// The number of days log events of the singleton Lambda function implementing this custom resource are kept in CloudWatch Logs.
	LogRetention awslogs.RetentionDays `field:"optional" json:"logRetention" yaml:"logRetention"`
	// The AWS SDK call to make when the resource is created.
	OnCreate *AwsSdkCall `field:"optional" json:"onCreate" yaml:"onCreate"`
	// The AWS SDK call to make when the resource is deleted.
	OnDelete *AwsSdkCall `field:"optional" json:"onDelete" yaml:"onDelete"`
	// The AWS SDK call to make when the resource is updated.
	OnUpdate *AwsSdkCall `field:"optional" json:"onUpdate" yaml:"onUpdate"`
	// The policy that will be added to the execution role of the Lambda function implementing this custom resource provider.
	//
	// The custom resource also implements `iam.IGrantable`, making it possible
	// to use the `grantXxx()` methods.
	//
	// As this custom resource uses a singleton Lambda function, it's important
	// to note the that function's role will eventually accumulate the
	// permissions/grants from all resources.
	//
	// Note that a policy must be specified if `role` is not provided, as
	// by default a new role is created which requires policy changes to access
	// resources.
	// See: Policy.fromSdkCalls
	//
	Policy AwsCustomResourcePolicy `field:"optional" json:"policy" yaml:"policy"`
	// Cloudformation Resource type.
	ResourceType *string `field:"optional" json:"resourceType" yaml:"resourceType"`
	// The execution role for the singleton Lambda function implementing this custom resource provider.
	//
	// This role will apply to all `AwsCustomResource`
	// instances in the stack. The role must be assumable by the
	// `lambda.amazonaws.com` service principal.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// The timeout for the singleton Lambda function implementing this custom resource.
	Timeout awscdk.Duration `field:"optional" json:"timeout" yaml:"timeout"`
	// The vpc to provision the lambda function in.
	Vpc awsec2.IVpc `field:"optional" json:"vpc" yaml:"vpc"`
	// Which subnets from the VPC to place the lambda function in.
	//
	// Only used if 'vpc' is supplied. Note: internet access for Lambdas
	// requires a NAT gateway, so picking Public subnets is not allowed.
	VpcSubnets *awsec2.SubnetSelection `field:"optional" json:"vpcSubnets" yaml:"vpcSubnets"`
}

