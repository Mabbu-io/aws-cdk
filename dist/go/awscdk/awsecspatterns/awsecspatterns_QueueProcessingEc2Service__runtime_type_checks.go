//go:build !no_runtime_type_checking

package awsecspatterns

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/awsecs"
	"github.com/aws/constructs-go/constructs/v10"
)

func (q *jsiiProxy_QueueProcessingEc2Service) validateConfigureAutoscalingForServiceParameters(service awsecs.BaseService) error {
	if service == nil {
		return fmt.Errorf("parameter service is required, but nil was provided")
	}

	return nil
}

func (q *jsiiProxy_QueueProcessingEc2Service) validateGetDefaultClusterParameters(scope constructs.Construct) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	return nil
}

func (q *jsiiProxy_QueueProcessingEc2Service) validateGrantPermissionsToServiceParameters(service awsecs.BaseService) error {
	if service == nil {
		return fmt.Errorf("parameter service is required, but nil was provided")
	}

	return nil
}

func validateQueueProcessingEc2Service_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateNewQueueProcessingEc2ServiceParameters(scope constructs.Construct, id *string, props *QueueProcessingEc2ServiceProps) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

