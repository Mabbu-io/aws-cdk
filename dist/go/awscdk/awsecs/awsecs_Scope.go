package awsecs


// The scope for the Docker volume that determines its lifecycle.
//
// Docker volumes that are scoped to a task are automatically provisioned when the task starts and destroyed when the task stops.
// Docker volumes that are scoped as shared persist after the task stops.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var proxyConfiguration proxyConfiguration
//   var role role
//
//   commonTaskDefinitionProps := &commonTaskDefinitionProps{
//   	executionRole: role,
//   	family: jsii.String("family"),
//   	proxyConfiguration: proxyConfiguration,
//   	taskRole: role,
//   	volumes: []volume{
//   		&volume{
//   			name: jsii.String("name"),
//
//   			// the properties below are optional
//   			dockerVolumeConfiguration: &dockerVolumeConfiguration{
//   				driver: jsii.String("driver"),
//   				scope: awscdk.Aws_ecs.scope_TASK,
//
//   				// the properties below are optional
//   				autoprovision: jsii.Boolean(false),
//   				driverOpts: map[string]*string{
//   					"driverOptsKey": jsii.String("driverOpts"),
//   				},
//   				labels: map[string]*string{
//   					"labelsKey": jsii.String("labels"),
//   				},
//   			},
//   			efsVolumeConfiguration: &efsVolumeConfiguration{
//   				fileSystemId: jsii.String("fileSystemId"),
//
//   				// the properties below are optional
//   				authorizationConfig: &authorizationConfig{
//   					accessPointId: jsii.String("accessPointId"),
//   					iam: jsii.String("iam"),
//   				},
//   				rootDirectory: jsii.String("rootDirectory"),
//   				transitEncryption: jsii.String("transitEncryption"),
//   				transitEncryptionPort: jsii.Number(123),
//   			},
//   			host: &host{
//   				sourcePath: jsii.String("sourcePath"),
//   			},
//   		},
//   	},
//   }
//
type Scope string

const (
	// Docker volumes that are scoped to a task are automatically provisioned when the task starts and destroyed when the task stops.
	Scope_TASK Scope = "TASK"
	// Docker volumes that are scoped as shared persist after the task stops.
	Scope_SHARED Scope = "SHARED"
)

