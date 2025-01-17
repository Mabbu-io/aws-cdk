package awsautoscaling

import (
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
)

// Options needed to bind a target to a lifecycle hook.
//
// [disable-awslint:ref-via-interface] The lifecycle hook to attach to and an IRole to use.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var lifecycleHook lifecycleHook
//   var role role
//
//   bindHookTargetOptions := &bindHookTargetOptions{
//   	lifecycleHook: lifecycleHook,
//
//   	// the properties below are optional
//   	role: role,
//   }
//
type BindHookTargetOptions struct {
	// The lifecycle hook to attach to.
	//
	// [disable-awslint:ref-via-interface].
	LifecycleHook LifecycleHook `field:"required" json:"lifecycleHook" yaml:"lifecycleHook"`
	// The role to use when attaching to the lifecycle hook.
	//
	// [disable-awslint:ref-via-interface].
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
}

