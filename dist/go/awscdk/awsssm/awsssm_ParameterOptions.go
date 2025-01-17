package awsssm


// Properties needed to create a new SSM Parameter.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   parameterOptions := &parameterOptions{
//   	allowedPattern: jsii.String("allowedPattern"),
//   	description: jsii.String("description"),
//   	parameterName: jsii.String("parameterName"),
//   	simpleName: jsii.Boolean(false),
//   	tier: awscdk.Aws_ssm.parameterTier_ADVANCED,
//   }
//
type ParameterOptions struct {
	// A regular expression used to validate the parameter value.
	//
	// For example, for String types with values restricted to
	// numbers, you can specify the following: ``^\d+$``.
	AllowedPattern *string `field:"optional" json:"allowedPattern" yaml:"allowedPattern"`
	// Information about the parameter that you want to add to the system.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the parameter.
	ParameterName *string `field:"optional" json:"parameterName" yaml:"parameterName"`
	// Indicates of the parameter name is a simple name (i.e. does not include "/" separators).
	//
	// This is only required only if `parameterName` is a token, which means we
	// are unable to detect if the name is simple or "path-like" for the purpose
	// of rendering SSM parameter ARNs.
	//
	// If `parameterName` is not specified, `simpleName` must be `true` (or
	// undefined) since the name generated by AWS CloudFormation is always a
	// simple name.
	SimpleName *bool `field:"optional" json:"simpleName" yaml:"simpleName"`
	// The tier of the string parameter.
	Tier ParameterTier `field:"optional" json:"tier" yaml:"tier"`
}

