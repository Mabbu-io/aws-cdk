package awsecs


// Type of resource to set a limit on.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ulimit := &ulimit{
//   	hardLimit: jsii.Number(123),
//   	name: awscdk.Aws_ecs.ulimitName_CORE,
//   	softLimit: jsii.Number(123),
//   }
//
type UlimitName string

const (
	UlimitName_CORE UlimitName = "CORE"
	UlimitName_CPU UlimitName = "CPU"
	UlimitName_DATA UlimitName = "DATA"
	UlimitName_FSIZE UlimitName = "FSIZE"
	UlimitName_LOCKS UlimitName = "LOCKS"
	UlimitName_MEMLOCK UlimitName = "MEMLOCK"
	UlimitName_MSGQUEUE UlimitName = "MSGQUEUE"
	UlimitName_NICE UlimitName = "NICE"
	UlimitName_NOFILE UlimitName = "NOFILE"
	UlimitName_NPROC UlimitName = "NPROC"
	UlimitName_RSS UlimitName = "RSS"
	UlimitName_RTPRIO UlimitName = "RTPRIO"
	UlimitName_RTTIME UlimitName = "RTTIME"
	UlimitName_SIGPENDING UlimitName = "SIGPENDING"
	UlimitName_STACK UlimitName = "STACK"
)

