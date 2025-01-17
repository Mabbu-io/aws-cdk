package awsevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
)

// Authorization type for an API Destination Connection.
//
// Example:
//   connection := events.NewConnection(this, jsii.String("Connection"), &connectionProps{
//   	authorization: events.authorization.apiKey(jsii.String("x-api-key"), awscdk.SecretValue.secretsManager(jsii.String("ApiSecretName"))),
//   	description: jsii.String("Connection with API Key x-api-key"),
//   })
//
//   destination := events.NewApiDestination(this, jsii.String("Destination"), &apiDestinationProps{
//   	connection: connection,
//   	endpoint: jsii.String("https://example.com"),
//   	description: jsii.String("Calling example.com with API key x-api-key"),
//   })
//
//   rule := events.NewRule(this, jsii.String("Rule"), &ruleProps{
//   	schedule: events.schedule.rate(cdk.duration.minutes(jsii.Number(1))),
//   	targets: []iRuleTarget{
//   		targets.NewApiDestination(destination),
//   	},
//   })
//
type Authorization interface {
}

// The jsii proxy struct for Authorization
type jsiiProxy_Authorization struct {
	_ byte // padding
}

func NewAuthorization_Override(a Authorization) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_events.Authorization",
		nil, // no parameters
		a,
	)
}

// Use API key authorization.
//
// API key authorization has two components: an API key name and an API key value.
// What these are depends on the target of your connection.
func Authorization_ApiKey(apiKeyName *string, apiKeyValue awscdk.SecretValue) Authorization {
	_init_.Initialize()

	if err := validateAuthorization_ApiKeyParameters(apiKeyName, apiKeyValue); err != nil {
		panic(err)
	}
	var returns Authorization

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_events.Authorization",
		"apiKey",
		[]interface{}{apiKeyName, apiKeyValue},
		&returns,
	)

	return returns
}

// Use username and password authorization.
func Authorization_Basic(username *string, password awscdk.SecretValue) Authorization {
	_init_.Initialize()

	if err := validateAuthorization_BasicParameters(username, password); err != nil {
		panic(err)
	}
	var returns Authorization

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_events.Authorization",
		"basic",
		[]interface{}{username, password},
		&returns,
	)

	return returns
}

// Use OAuth authorization.
func Authorization_Oauth(props *OAuthAuthorizationProps) Authorization {
	_init_.Initialize()

	if err := validateAuthorization_OauthParameters(props); err != nil {
		panic(err)
	}
	var returns Authorization

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_events.Authorization",
		"oauth",
		[]interface{}{props},
		&returns,
	)

	return returns
}

