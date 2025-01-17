//go:build !no_runtime_type_checking

// Version 2 of the AWS Cloud Development Kit library
package awscdk

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/cloudassemblyschema"
)

func (c *jsiiProxy_CliCredentialsStackSynthesizer) validateAddBootstrapVersionRuleParameters(requiredVersion *float64, bootstrapStackVersionSsmParameter *string) error {
	if requiredVersion == nil {
		return fmt.Errorf("parameter requiredVersion is required, but nil was provided")
	}

	if bootstrapStackVersionSsmParameter == nil {
		return fmt.Errorf("parameter bootstrapStackVersionSsmParameter is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CliCredentialsStackSynthesizer) validateAddDockerImageAssetParameters(asset *DockerImageAssetSource) error {
	if asset == nil {
		return fmt.Errorf("parameter asset is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(asset, func() string { return "parameter asset" }); err != nil {
		return err
	}

	return nil
}

func (c *jsiiProxy_CliCredentialsStackSynthesizer) validateAddFileAssetParameters(asset *FileAssetSource) error {
	if asset == nil {
		return fmt.Errorf("parameter asset is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(asset, func() string { return "parameter asset" }); err != nil {
		return err
	}

	return nil
}

func (c *jsiiProxy_CliCredentialsStackSynthesizer) validateBindParameters(stack Stack) error {
	if stack == nil {
		return fmt.Errorf("parameter stack is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CliCredentialsStackSynthesizer) validateCloudFormationLocationFromDockerImageAssetParameters(dest *cloudassemblyschema.DockerImageDestination) error {
	if dest == nil {
		return fmt.Errorf("parameter dest is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(dest, func() string { return "parameter dest" }); err != nil {
		return err
	}

	return nil
}

func (c *jsiiProxy_CliCredentialsStackSynthesizer) validateCloudFormationLocationFromFileAssetParameters(location *cloudassemblyschema.FileDestination) error {
	if location == nil {
		return fmt.Errorf("parameter location is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(location, func() string { return "parameter location" }); err != nil {
		return err
	}

	return nil
}

func (c *jsiiProxy_CliCredentialsStackSynthesizer) validateEmitArtifactParameters(session ISynthesisSession, options *SynthesizeStackArtifactOptions) error {
	if session == nil {
		return fmt.Errorf("parameter session is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (c *jsiiProxy_CliCredentialsStackSynthesizer) validateEmitStackArtifactParameters(stack Stack, session ISynthesisSession, options *SynthesizeStackArtifactOptions) error {
	if stack == nil {
		return fmt.Errorf("parameter stack is required, but nil was provided")
	}

	if session == nil {
		return fmt.Errorf("parameter session is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (c *jsiiProxy_CliCredentialsStackSynthesizer) validateSynthesizeParameters(session ISynthesisSession) error {
	if session == nil {
		return fmt.Errorf("parameter session is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CliCredentialsStackSynthesizer) validateSynthesizeStackTemplateParameters(stack Stack, session ISynthesisSession) error {
	if stack == nil {
		return fmt.Errorf("parameter stack is required, but nil was provided")
	}

	if session == nil {
		return fmt.Errorf("parameter session is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CliCredentialsStackSynthesizer) validateSynthesizeTemplateParameters(session ISynthesisSession) error {
	if session == nil {
		return fmt.Errorf("parameter session is required, but nil was provided")
	}

	return nil
}

func validateNewCliCredentialsStackSynthesizerParameters(props *CliCredentialsStackSynthesizerProps) error {
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

