package awsecs

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// The base class for proxy configurations.
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
type ProxyConfiguration interface {
	// Called when the proxy configuration is configured on a task definition.
	Bind(_scope constructs.Construct, _taskDefinition TaskDefinition) *CfnTaskDefinition_ProxyConfigurationProperty
}

// The jsii proxy struct for ProxyConfiguration
type jsiiProxy_ProxyConfiguration struct {
	_ byte // padding
}

func NewProxyConfiguration_Override(p ProxyConfiguration) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.ProxyConfiguration",
		nil, // no parameters
		p,
	)
}

func (p *jsiiProxy_ProxyConfiguration) Bind(_scope constructs.Construct, _taskDefinition TaskDefinition) *CfnTaskDefinition_ProxyConfigurationProperty {
	if err := p.validateBindParameters(_scope, _taskDefinition); err != nil {
		panic(err)
	}
	var returns *CfnTaskDefinition_ProxyConfigurationProperty

	_jsii_.Invoke(
		p,
		"bind",
		[]interface{}{_scope, _taskDefinition},
		&returns,
	)

	return returns
}

