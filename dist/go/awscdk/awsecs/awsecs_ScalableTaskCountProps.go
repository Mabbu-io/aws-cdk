package awsecs

import (
	"github.com/aws/aws-cdk-go/awscdk/awsapplicationautoscaling"
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
)

// The properties of a scalable attribute representing task count.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var role role
//
//   scalableTaskCountProps := &scalableTaskCountProps{
//   	dimension: jsii.String("dimension"),
//   	maxCapacity: jsii.Number(123),
//   	resourceId: jsii.String("resourceId"),
//   	role: role,
//   	serviceNamespace: awscdk.Aws_applicationautoscaling.serviceNamespace_ECS,
//
//   	// the properties below are optional
//   	minCapacity: jsii.Number(123),
//   }
//
type ScalableTaskCountProps struct {
	// Maximum capacity to scale to.
	MaxCapacity *float64 `field:"required" json:"maxCapacity" yaml:"maxCapacity"`
	// Minimum capacity to scale to.
	MinCapacity *float64 `field:"optional" json:"minCapacity" yaml:"minCapacity"`
	// Scalable dimension of the attribute.
	Dimension *string `field:"required" json:"dimension" yaml:"dimension"`
	// Resource ID of the attribute.
	ResourceId *string `field:"required" json:"resourceId" yaml:"resourceId"`
	// Role to use for scaling.
	Role awsiam.IRole `field:"required" json:"role" yaml:"role"`
	// Service namespace of the scalable attribute.
	ServiceNamespace awsapplicationautoscaling.ServiceNamespace `field:"required" json:"serviceNamespace" yaml:"serviceNamespace"`
}

