//go:build !no_runtime_type_checking

package awscloudwatchactions

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/awscloudwatch"
	"github.com/aws/constructs-go/constructs/v10"
)

func (e *jsiiProxy_Ec2Action) validateBindParameters(_scope constructs.Construct, _alarm awscloudwatch.IAlarm) error {
	if _scope == nil {
		return fmt.Errorf("parameter _scope is required, but nil was provided")
	}

	if _alarm == nil {
		return fmt.Errorf("parameter _alarm is required, but nil was provided")
	}

	return nil
}

func validateNewEc2ActionParameters(instanceAction Ec2InstanceAction) error {
	if instanceAction == "" {
		return fmt.Errorf("parameter instanceAction is required, but nil was provided")
	}

	return nil
}

