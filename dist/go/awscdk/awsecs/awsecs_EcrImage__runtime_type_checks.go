//go:build !no_runtime_type_checking

package awsecs

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/awsecr"
	"github.com/aws/aws-cdk-go/awscdk/awsecrassets"
	"github.com/aws/constructs-go/constructs/v10"
)

func (e *jsiiProxy_EcrImage) validateBindParameters(_scope constructs.Construct, containerDefinition ContainerDefinition) error {
	if _scope == nil {
		return fmt.Errorf("parameter _scope is required, but nil was provided")
	}

	if containerDefinition == nil {
		return fmt.Errorf("parameter containerDefinition is required, but nil was provided")
	}

	return nil
}

func validateEcrImage_FromAssetParameters(directory *string, props *AssetImageProps) error {
	if directory == nil {
		return fmt.Errorf("parameter directory is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func validateEcrImage_FromDockerImageAssetParameters(asset awsecrassets.DockerImageAsset) error {
	if asset == nil {
		return fmt.Errorf("parameter asset is required, but nil was provided")
	}

	return nil
}

func validateEcrImage_FromEcrRepositoryParameters(repository awsecr.IRepository) error {
	if repository == nil {
		return fmt.Errorf("parameter repository is required, but nil was provided")
	}

	return nil
}

func validateEcrImage_FromRegistryParameters(name *string, props *RepositoryImageProps) error {
	if name == nil {
		return fmt.Errorf("parameter name is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func validateEcrImage_FromTarballParameters(tarballFile *string) error {
	if tarballFile == nil {
		return fmt.Errorf("parameter tarballFile is required, but nil was provided")
	}

	return nil
}

func validateNewEcrImageParameters(repository awsecr.IRepository, tagOrDigest *string) error {
	if repository == nil {
		return fmt.Errorf("parameter repository is required, but nil was provided")
	}

	if tagOrDigest == nil {
		return fmt.Errorf("parameter tagOrDigest is required, but nil was provided")
	}

	return nil
}

