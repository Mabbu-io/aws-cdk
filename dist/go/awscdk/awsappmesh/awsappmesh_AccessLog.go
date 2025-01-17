package awsappmesh

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// Configuration for Envoy Access logs for mesh endpoints.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   var mesh mesh
//   vpc := ec2.NewVpc(this, jsii.String("vpc"))
//   namespace := cloudmap.NewPrivateDnsNamespace(this, jsii.String("test-namespace"), &privateDnsNamespaceProps{
//   	vpc: vpc,
//   	name: jsii.String("domain.local"),
//   })
//   service := namespace.createService(jsii.String("Svc"))
//   node := mesh.addVirtualNode(jsii.String("virtual-node"), &virtualNodeBaseProps{
//   	serviceDiscovery: appmesh.serviceDiscovery.cloudMap(service),
//   	listeners: []virtualNodeListener{
//   		appmesh.*virtualNodeListener.http(&httpVirtualNodeListenerOptions{
//   			port: jsii.Number(8081),
//   			healthCheck: appmesh.healthCheck.http(&httpHealthCheckOptions{
//   				healthyThreshold: jsii.Number(3),
//   				interval: cdk.duration.seconds(jsii.Number(5)),
//   				 // minimum
//   				path: jsii.String("/health-check-path"),
//   				timeout: cdk.*duration.seconds(jsii.Number(2)),
//   				 // minimum
//   				unhealthyThreshold: jsii.Number(2),
//   			}),
//   		}),
//   	},
//   	accessLog: appmesh.accessLog.fromFilePath(jsii.String("/dev/stdout")),
//   })
//
type AccessLog interface {
	// Called when the AccessLog type is initialized.
	//
	// Can be used to enforce
	// mutual exclusivity with future properties.
	Bind(scope constructs.Construct) *AccessLogConfig
}

// The jsii proxy struct for AccessLog
type jsiiProxy_AccessLog struct {
	_ byte // padding
}

func NewAccessLog_Override(a AccessLog) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_appmesh.AccessLog",
		nil, // no parameters
		a,
	)
}

// Path to a file to write access logs to.
func AccessLog_FromFilePath(filePath *string) AccessLog {
	_init_.Initialize()

	if err := validateAccessLog_FromFilePathParameters(filePath); err != nil {
		panic(err)
	}
	var returns AccessLog

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.AccessLog",
		"fromFilePath",
		[]interface{}{filePath},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessLog) Bind(scope constructs.Construct) *AccessLogConfig {
	if err := a.validateBindParameters(scope); err != nil {
		panic(err)
	}
	var returns *AccessLogConfig

	_jsii_.Invoke(
		a,
		"bind",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

