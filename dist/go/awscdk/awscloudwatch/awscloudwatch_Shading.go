package awscloudwatch


// Fill shading options that will be used with an annotation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   horizontalAnnotation := &horizontalAnnotation{
//   	value: jsii.Number(123),
//
//   	// the properties below are optional
//   	color: jsii.String("color"),
//   	fill: awscdk.Aws_cloudwatch.shading_NONE,
//   	label: jsii.String("label"),
//   	visible: jsii.Boolean(false),
//   }
//
type Shading string

const (
	// Don't add shading.
	Shading_NONE Shading = "NONE"
	// Add shading above the annotation.
	Shading_ABOVE Shading = "ABOVE"
	// Add shading below the annotation.
	Shading_BELOW Shading = "BELOW"
)

