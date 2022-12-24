package awsecs


// Permissions for device access.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   device := &device{
//   	hostPath: jsii.String("hostPath"),
//
//   	// the properties below are optional
//   	containerPath: jsii.String("containerPath"),
//   	permissions: []devicePermission{
//   		awscdk.Aws_ecs.*devicePermission_READ,
//   	},
//   }
//
type DevicePermission string

const (
	// Read.
	DevicePermission_READ DevicePermission = "READ"
	// Write.
	DevicePermission_WRITE DevicePermission = "WRITE"
	// Make a node.
	DevicePermission_MKNOD DevicePermission = "MKNOD"
)

