package awsstepfunctions


// Three ways to call an integrated service: Request Response, Run a Job and Wait for a Callback with Task Token.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var cluster cluster
//   var containerDefinition containerDefinition
//   var taskDefinition taskDefinition
//
//   commonEcsRunTaskProps := &commonEcsRunTaskProps{
//   	cluster: cluster,
//   	taskDefinition: taskDefinition,
//
//   	// the properties below are optional
//   	containerOverrides: []containerOverride{
//   		&containerOverride{
//   			containerDefinition: containerDefinition,
//
//   			// the properties below are optional
//   			command: []*string{
//   				jsii.String("command"),
//   			},
//   			cpu: jsii.Number(123),
//   			environment: []taskEnvironmentVariable{
//   				&taskEnvironmentVariable{
//   					name: jsii.String("name"),
//   					value: jsii.String("value"),
//   				},
//   			},
//   			memoryLimit: jsii.Number(123),
//   			memoryReservation: jsii.Number(123),
//   		},
//   	},
//   	integrationPattern: awscdk.Aws_stepfunctions.serviceIntegrationPattern_FIRE_AND_FORGET,
//   }
//
// See: https://docs.aws.amazon.com/step-functions/latest/dg/connect-to-resource.html
//
// Here, they are named as FIRE_AND_FORGET, SYNC and WAIT_FOR_TASK_TOKEN respectfully.
//
type ServiceIntegrationPattern string

const (
	// Call a service and progress to the next state immediately after the API call completes.
	ServiceIntegrationPattern_FIRE_AND_FORGET ServiceIntegrationPattern = "FIRE_AND_FORGET"
	// Call a service and wait for a job to complete.
	ServiceIntegrationPattern_SYNC ServiceIntegrationPattern = "SYNC"
	// Call a service with a task token and wait until that token is returned by SendTaskSuccess/SendTaskFailure with payload.
	ServiceIntegrationPattern_WAIT_FOR_TASK_TOKEN ServiceIntegrationPattern = "WAIT_FOR_TASK_TOKEN"
)

