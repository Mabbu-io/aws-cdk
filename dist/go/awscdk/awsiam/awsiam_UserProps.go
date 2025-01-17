package awsiam

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining an IAM user.
//
// Example:
//   user := awscdk.NewUser(this, jsii.String("MyUser"), &userProps{
//   	password: awscdk.SecretValue.plainText(jsii.String("1234")),
//   })
//   group := awscdk.NewGroup(this, jsii.String("MyGroup"))
//
//   policy := awscdk.NewPolicy(this, jsii.String("MyPolicy"))
//   policy.attachToUser(user)
//   group.attachInlinePolicy(policy)
//
type UserProps struct {
	// Groups to add this user to.
	//
	// You can also use `addToGroup` to add this
	// user to a group.
	Groups *[]IGroup `field:"optional" json:"groups" yaml:"groups"`
	// A list of managed policies associated with this role.
	//
	// You can add managed policies later using
	// `addManagedPolicy(ManagedPolicy.fromAwsManagedPolicyName(policyName))`.
	ManagedPolicies *[]IManagedPolicy `field:"optional" json:"managedPolicies" yaml:"managedPolicies"`
	// The password for the user. This is required so the user can access the AWS Management Console.
	//
	// You can use `SecretValue.unsafePlainText` to specify a password in plain text or
	// use `secretsmanager.Secret.fromSecretAttributes` to reference a secret in
	// Secrets Manager.
	Password awscdk.SecretValue `field:"optional" json:"password" yaml:"password"`
	// Specifies whether the user is required to set a new password the next time the user logs in to the AWS Management Console.
	//
	// If this is set to 'true', you must also specify "initialPassword".
	PasswordResetRequired *bool `field:"optional" json:"passwordResetRequired" yaml:"passwordResetRequired"`
	// The path for the user name.
	//
	// For more information about paths, see IAM
	// Identifiers in the IAM User Guide.
	Path *string `field:"optional" json:"path" yaml:"path"`
	// AWS supports permissions boundaries for IAM entities (users or roles).
	//
	// A permissions boundary is an advanced feature for using a managed policy
	// to set the maximum permissions that an identity-based policy can grant to
	// an IAM entity. An entity's permissions boundary allows it to perform only
	// the actions that are allowed by both its identity-based policies and its
	// permissions boundaries.
	PermissionsBoundary IManagedPolicy `field:"optional" json:"permissionsBoundary" yaml:"permissionsBoundary"`
	// A name for the IAM user.
	//
	// For valid values, see the UserName parameter for
	// the CreateUser action in the IAM API Reference. If you don't specify a
	// name, AWS CloudFormation generates a unique physical ID and uses that ID
	// for the user name.
	//
	// If you specify a name, you cannot perform updates that require
	// replacement of this resource. You can perform updates that require no or
	// some interruption. If you must replace the resource, specify a new name.
	//
	// If you specify a name, you must specify the CAPABILITY_NAMED_IAM value to
	// acknowledge your template's capabilities. For more information, see
	// Acknowledging IAM Resources in AWS CloudFormation Templates.
	UserName *string `field:"optional" json:"userName" yaml:"userName"`
}

