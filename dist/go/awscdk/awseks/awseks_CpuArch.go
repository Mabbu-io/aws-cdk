package awseks


// CPU architecture.
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
type CpuArch string

const (
	// arm64 CPU type.
	CpuArch_ARM_64 CpuArch = "ARM_64"
	// x86_64 CPU type.
	CpuArch_X86_64 CpuArch = "X86_64"
)

