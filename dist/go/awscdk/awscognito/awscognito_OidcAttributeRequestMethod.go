package awscognito


// The method to use to request attributes.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var providerAttribute providerAttribute
//   var userPool userPool
//
//   userPoolIdentityProviderOidc := awscdk.Aws_cognito.NewUserPoolIdentityProviderOidc(this, jsii.String("MyUserPoolIdentityProviderOidc"), &userPoolIdentityProviderOidcProps{
//   	clientId: jsii.String("clientId"),
//   	clientSecret: jsii.String("clientSecret"),
//   	issuerUrl: jsii.String("issuerUrl"),
//   	userPool: userPool,
//
//   	// the properties below are optional
//   	attributeMapping: &attributeMapping{
//   		address: providerAttribute,
//   		birthdate: providerAttribute,
//   		custom: map[string]*providerAttribute{
//   			"customKey": providerAttribute,
//   		},
//   		email: providerAttribute,
//   		familyName: providerAttribute,
//   		fullname: providerAttribute,
//   		gender: providerAttribute,
//   		givenName: providerAttribute,
//   		lastUpdateTime: providerAttribute,
//   		locale: providerAttribute,
//   		middleName: providerAttribute,
//   		nickname: providerAttribute,
//   		phoneNumber: providerAttribute,
//   		preferredUsername: providerAttribute,
//   		profilePage: providerAttribute,
//   		profilePicture: providerAttribute,
//   		timezone: providerAttribute,
//   		website: providerAttribute,
//   	},
//   	attributeRequestMethod: awscdk.*Aws_cognito.oidcAttributeRequestMethod_GET,
//   	endpoints: &oidcEndpoints{
//   		authorization: jsii.String("authorization"),
//   		jwksUri: jsii.String("jwksUri"),
//   		token: jsii.String("token"),
//   		userInfo: jsii.String("userInfo"),
//   	},
//   	identifiers: []*string{
//   		jsii.String("identifiers"),
//   	},
//   	name: jsii.String("name"),
//   	scopes: []*string{
//   		jsii.String("scopes"),
//   	},
//   })
//
type OidcAttributeRequestMethod string

const (
	// GET.
	OidcAttributeRequestMethod_GET OidcAttributeRequestMethod = "GET"
	// POST.
	OidcAttributeRequestMethod_POST OidcAttributeRequestMethod = "POST"
)

