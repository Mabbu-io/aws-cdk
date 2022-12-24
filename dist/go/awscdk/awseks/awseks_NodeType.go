package awseks


// Whether the worker nodes should support GPU or just standard instances.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   eksOptimizedImage := awscdk.Aws_eks.NewEksOptimizedImage(&eksOptimizedImageProps{
//   	cpuArch: awscdk.*Aws_eks.cpuArch_ARM_64,
//   	kubernetesVersion: jsii.String("kubernetesVersion"),
//   	nodeType: awscdk.*Aws_eks.nodeType_STANDARD,
//   })
//
type NodeType string

const (
	// Standard instances.
	NodeType_STANDARD NodeType = "STANDARD"
	// GPU instances.
	NodeType_GPU NodeType = "GPU"
	// Inferentia instances.
	NodeType_INFERENTIA NodeType = "INFERENTIA"
)

