package awsroute53targets

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/awsroute53"
	"github.com/aws/aws-cdk-go/awscdk/awsroute53targets/internal"
)

// Defines an API Gateway V2 domain name as the alias target.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   import apigwv2 "github.com/aws/aws-cdk-go/awscdk"
//
//   var zone hostedZone
//   var domainName apigwv2.DomainName
//
//
//   route53.NewARecord(this, jsii.String("AliasRecord"), &aRecordProps{
//   	zone: zone,
//   	target: route53.recordTarget.fromAlias(targets.NewApiGatewayv2DomainProperties(domainName.regionalDomainName, domainName.regionalHostedZoneId)),
//   })
//
type ApiGatewayv2DomainProperties interface {
	awsroute53.IAliasRecordTarget
	// Return hosted zone ID and DNS name, usable for Route53 alias targets.
	Bind(_record awsroute53.IRecordSet, _zone awsroute53.IHostedZone) *awsroute53.AliasRecordTargetConfig
}

// The jsii proxy struct for ApiGatewayv2DomainProperties
type jsiiProxy_ApiGatewayv2DomainProperties struct {
	internal.Type__awsroute53IAliasRecordTarget
}

func NewApiGatewayv2DomainProperties(regionalDomainName *string, regionalHostedZoneId *string) ApiGatewayv2DomainProperties {
	_init_.Initialize()

	if err := validateNewApiGatewayv2DomainPropertiesParameters(regionalDomainName, regionalHostedZoneId); err != nil {
		panic(err)
	}
	j := jsiiProxy_ApiGatewayv2DomainProperties{}

	_jsii_.Create(
		"aws-cdk-lib.aws_route53_targets.ApiGatewayv2DomainProperties",
		[]interface{}{regionalDomainName, regionalHostedZoneId},
		&j,
	)

	return &j
}

func NewApiGatewayv2DomainProperties_Override(a ApiGatewayv2DomainProperties, regionalDomainName *string, regionalHostedZoneId *string) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_route53_targets.ApiGatewayv2DomainProperties",
		[]interface{}{regionalDomainName, regionalHostedZoneId},
		a,
	)
}

func (a *jsiiProxy_ApiGatewayv2DomainProperties) Bind(_record awsroute53.IRecordSet, _zone awsroute53.IHostedZone) *awsroute53.AliasRecordTargetConfig {
	if err := a.validateBindParameters(_record); err != nil {
		panic(err)
	}
	var returns *awsroute53.AliasRecordTargetConfig

	_jsii_.Invoke(
		a,
		"bind",
		[]interface{}{_record, _zone},
		&returns,
	)

	return returns
}

