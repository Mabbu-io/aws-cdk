package awsecs

import (
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
)

// The common task definition attributes used across all types of task definitions.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var role role
//
//   commonTaskDefinitionAttributes := &commonTaskDefinitionAttributes{
//   	taskDefinitionArn: jsii.String("taskDefinitionArn"),
//
//   	// the properties below are optional
//   	networkMode: awscdk.Aws_ecs.networkMode_NONE,
//   	taskRole: role,
//   }
//
type CommonTaskDefinitionAttributes struct {
	// The arn of the task definition.
	TaskDefinitionArn *string `field:"required" json:"taskDefinitionArn" yaml:"taskDefinitionArn"`
	// The networking mode to use for the containers in the task.
	NetworkMode NetworkMode `field:"optional" json:"networkMode" yaml:"networkMode"`
	// The name of the IAM role that grants containers in the task permission to call AWS APIs on your behalf.
	TaskRole awsiam.IRole `field:"optional" json:"taskRole" yaml:"taskRole"`
}

