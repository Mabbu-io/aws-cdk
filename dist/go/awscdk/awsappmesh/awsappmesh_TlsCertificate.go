package awsappmesh

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/awscertificatemanager"
	"github.com/aws/constructs-go/constructs/v10"
)

// Represents a TLS certificate.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   // A Virtual Node with listener TLS from an ACM provided certificate
//   var cert certificate
//   var mesh mesh
//
//
//   node := appmesh.NewVirtualNode(this, jsii.String("node"), &virtualNodeProps{
//   	mesh: mesh,
//   	serviceDiscovery: appmesh.serviceDiscovery.dns(jsii.String("node")),
//   	listeners: []virtualNodeListener{
//   		appmesh.*virtualNodeListener.grpc(&grpcVirtualNodeListenerOptions{
//   			port: jsii.Number(80),
//   			tls: &listenerTlsOptions{
//   				mode: appmesh.tlsMode_STRICT,
//   				certificate: appmesh.tlsCertificate.acm(cert),
//   			},
//   		}),
//   	},
//   })
//
//   // A Virtual Gateway with listener TLS from a customer provided file certificate
//   gateway := appmesh.NewVirtualGateway(this, jsii.String("gateway"), &virtualGatewayProps{
//   	mesh: mesh,
//   	listeners: []virtualGatewayListener{
//   		appmesh.*virtualGatewayListener.grpc(&grpcGatewayListenerOptions{
//   			port: jsii.Number(8080),
//   			tls: &listenerTlsOptions{
//   				mode: appmesh.*tlsMode_STRICT,
//   				certificate: appmesh.*tlsCertificate.file(jsii.String("path/to/certChain"), jsii.String("path/to/privateKey")),
//   			},
//   		}),
//   	},
//   	virtualGatewayName: jsii.String("gateway"),
//   })
//
//   // A Virtual Gateway with listener TLS from a SDS provided certificate
//   gateway2 := appmesh.NewVirtualGateway(this, jsii.String("gateway2"), &virtualGatewayProps{
//   	mesh: mesh,
//   	listeners: []*virtualGatewayListener{
//   		appmesh.*virtualGatewayListener.http2(&http2GatewayListenerOptions{
//   			port: jsii.Number(8080),
//   			tls: &listenerTlsOptions{
//   				mode: appmesh.*tlsMode_STRICT,
//   				certificate: appmesh.*tlsCertificate.sds(jsii.String("secrete_certificate")),
//   			},
//   		}),
//   	},
//   	virtualGatewayName: jsii.String("gateway2"),
//   })
//
type TlsCertificate interface {
	// Returns TLS certificate based provider.
	Bind(_scope constructs.Construct) *TlsCertificateConfig
}

// The jsii proxy struct for TlsCertificate
type jsiiProxy_TlsCertificate struct {
	_ byte // padding
}

func NewTlsCertificate_Override(t TlsCertificate) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_appmesh.TlsCertificate",
		nil, // no parameters
		t,
	)
}

// Returns an ACM TLS Certificate.
func TlsCertificate_Acm(certificate awscertificatemanager.ICertificate) TlsCertificate {
	_init_.Initialize()

	if err := validateTlsCertificate_AcmParameters(certificate); err != nil {
		panic(err)
	}
	var returns TlsCertificate

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.TlsCertificate",
		"acm",
		[]interface{}{certificate},
		&returns,
	)

	return returns
}

// Returns an File TLS Certificate.
func TlsCertificate_File(certificateChainPath *string, privateKeyPath *string) MutualTlsCertificate {
	_init_.Initialize()

	if err := validateTlsCertificate_FileParameters(certificateChainPath, privateKeyPath); err != nil {
		panic(err)
	}
	var returns MutualTlsCertificate

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.TlsCertificate",
		"file",
		[]interface{}{certificateChainPath, privateKeyPath},
		&returns,
	)

	return returns
}

// Returns an SDS TLS Certificate.
func TlsCertificate_Sds(secretName *string) MutualTlsCertificate {
	_init_.Initialize()

	if err := validateTlsCertificate_SdsParameters(secretName); err != nil {
		panic(err)
	}
	var returns MutualTlsCertificate

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.TlsCertificate",
		"sds",
		[]interface{}{secretName},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TlsCertificate) Bind(_scope constructs.Construct) *TlsCertificateConfig {
	if err := t.validateBindParameters(_scope); err != nil {
		panic(err)
	}
	var returns *TlsCertificateConfig

	_jsii_.Invoke(
		t,
		"bind",
		[]interface{}{_scope},
		&returns,
	)

	return returns
}

