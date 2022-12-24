package awscognito

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Password policy for User Pools.
//
// Example:
//   cognito.NewUserPool(this, jsii.String("myuserpool"), &userPoolProps{
//   	// ...
//   	passwordPolicy: &passwordPolicy{
//   		minLength: jsii.Number(12),
//   		requireLowercase: jsii.Boolean(true),
//   		requireUppercase: jsii.Boolean(true),
//   		requireDigits: jsii.Boolean(true),
//   		requireSymbols: jsii.Boolean(true),
//   		tempPasswordValidity: awscdk.Duration.days(jsii.Number(3)),
//   	},
//   })
//
type PasswordPolicy struct {
	// Minimum length required for a user's password.
	MinLength *float64 `field:"optional" json:"minLength" yaml:"minLength"`
	// Whether the user is required to have digits in their password.
	RequireDigits *bool `field:"optional" json:"requireDigits" yaml:"requireDigits"`
	// Whether the user is required to have lowercase characters in their password.
	RequireLowercase *bool `field:"optional" json:"requireLowercase" yaml:"requireLowercase"`
	// Whether the user is required to have symbols in their password.
	RequireSymbols *bool `field:"optional" json:"requireSymbols" yaml:"requireSymbols"`
	// Whether the user is required to have uppercase characters in their password.
	RequireUppercase *bool `field:"optional" json:"requireUppercase" yaml:"requireUppercase"`
	// The length of time the temporary password generated by an admin is valid.
	//
	// This must be provided as whole days, like Duration.days(3) or Duration.hours(48).
	// Fractional days, such as Duration.hours(20), will generate an error.
	TempPasswordValidity awscdk.Duration `field:"optional" json:"tempPasswordValidity" yaml:"tempPasswordValidity"`
}

