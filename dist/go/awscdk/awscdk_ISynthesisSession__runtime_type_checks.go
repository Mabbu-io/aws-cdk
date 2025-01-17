//go:build !no_runtime_type_checking

// Version 2 of the AWS Cloud Development Kit library
package awscdk

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/cxapi"
)

func (j *jsiiProxy_ISynthesisSession) validateSetAssemblyParameters(val cxapi.CloudAssemblyBuilder) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_ISynthesisSession) validateSetOutdirParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

