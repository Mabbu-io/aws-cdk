package awselasticloadbalancingv2


// What to do with unauthenticated requests.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var listenerAction listenerAction
//   var secretValue secretValue
//
//   authenticateCognitoAction := awscdk.Aws_elasticloadbalancingv2_actions.authenticateCognitoAction.authenticateOidc(&authenticateOidcOptions{
//   	authorizationEndpoint: jsii.String("authorizationEndpoint"),
//   	clientId: jsii.String("clientId"),
//   	clientSecret: secretValue,
//   	issuer: jsii.String("issuer"),
//   	next: listenerAction,
//   	tokenEndpoint: jsii.String("tokenEndpoint"),
//   	userInfoEndpoint: jsii.String("userInfoEndpoint"),
//
//   	// the properties below are optional
//   	authenticationRequestExtraParams: map[string]*string{
//   		"authenticationRequestExtraParamsKey": jsii.String("authenticationRequestExtraParams"),
//   	},
//   	onUnauthenticatedRequest: awscdk.Aws_elasticloadbalancingv2.unauthenticatedAction_DENY,
//   	scope: jsii.String("scope"),
//   	sessionCookieName: jsii.String("sessionCookieName"),
//   	sessionTimeout: cdk.duration.minutes(jsii.Number(30)),
//   })
//
type UnauthenticatedAction string

const (
	// Return an HTTP 401 Unauthorized error.
	UnauthenticatedAction_DENY UnauthenticatedAction = "DENY"
	// Allow the request to be forwarded to the target.
	UnauthenticatedAction_ALLOW UnauthenticatedAction = "ALLOW"
	// Redirect the request to the IdP authorization endpoint.
	UnauthenticatedAction_AUTHENTICATE UnauthenticatedAction = "AUTHENTICATE"
)

