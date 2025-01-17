package awsroute53targets

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/awscognito"
	"github.com/aws/aws-cdk-go/awscdk/awsroute53"
	"github.com/aws/aws-cdk-go/awscdk/awsroute53targets/internal"
)

// Use a user pool domain as an alias record target.
//
// Example:
//   import cognito "github.com/aws/aws-cdk-go/awscdk"
//
//   var zone hostedZone
//   var domain userPoolDomain
//
//   route53.NewARecord(this, jsii.String("AliasRecord"), &aRecordProps{
//   	zone: zone,
//   	target: route53.recordTarget.fromAlias(targets.NewUserPoolDomainTarget(domain)),
//   })
//
type UserPoolDomainTarget interface {
	awsroute53.IAliasRecordTarget
	// Return hosted zone ID and DNS name, usable for Route53 alias targets.
	Bind(_record awsroute53.IRecordSet, _zone awsroute53.IHostedZone) *awsroute53.AliasRecordTargetConfig
}

// The jsii proxy struct for UserPoolDomainTarget
type jsiiProxy_UserPoolDomainTarget struct {
	internal.Type__awsroute53IAliasRecordTarget
}

func NewUserPoolDomainTarget(domain awscognito.UserPoolDomain) UserPoolDomainTarget {
	_init_.Initialize()

	if err := validateNewUserPoolDomainTargetParameters(domain); err != nil {
		panic(err)
	}
	j := jsiiProxy_UserPoolDomainTarget{}

	_jsii_.Create(
		"aws-cdk-lib.aws_route53_targets.UserPoolDomainTarget",
		[]interface{}{domain},
		&j,
	)

	return &j
}

func NewUserPoolDomainTarget_Override(u UserPoolDomainTarget, domain awscognito.UserPoolDomain) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_route53_targets.UserPoolDomainTarget",
		[]interface{}{domain},
		u,
	)
}

func (u *jsiiProxy_UserPoolDomainTarget) Bind(_record awsroute53.IRecordSet, _zone awsroute53.IHostedZone) *awsroute53.AliasRecordTargetConfig {
	if err := u.validateBindParameters(_record); err != nil {
		panic(err)
	}
	var returns *awsroute53.AliasRecordTargetConfig

	_jsii_.Invoke(
		u,
		"bind",
		[]interface{}{_record, _zone},
		&returns,
	)

	return returns
}

