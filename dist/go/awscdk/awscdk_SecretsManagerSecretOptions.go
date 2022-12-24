// Version 2 of the AWS Cloud Development Kit library
package awscdk


// Options for referencing a secret value from Secrets Manager.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   secret := awscdk.SecretValue.secretsManager(jsii.String("secretId"), &secretsManagerSecretOptions{
//   	jsonField: jsii.String("password"),
//   	 // optional: key of a JSON field to retrieve (defaults to all content),
//   	versionId: jsii.String("id"),
//   	 // optional: id of the version (default AWSCURRENT)
//   	versionStage: jsii.String("stage"),
//   })
//
type SecretsManagerSecretOptions struct {
	// The key of a JSON field to retrieve.
	//
	// This can only be used if the secret
	// stores a JSON object.
	JsonField *string `field:"optional" json:"jsonField" yaml:"jsonField"`
	// Specifies the unique identifier of the version of the secret you want to use.
	//
	// Can specify at most one of `versionId` and `versionStage`.
	VersionId *string `field:"optional" json:"versionId" yaml:"versionId"`
	// Specifies the secret version that you want to retrieve by the staging label attached to the version.
	//
	// Can specify at most one of `versionId` and `versionStage`.
	VersionStage *string `field:"optional" json:"versionStage" yaml:"versionStage"`
}

