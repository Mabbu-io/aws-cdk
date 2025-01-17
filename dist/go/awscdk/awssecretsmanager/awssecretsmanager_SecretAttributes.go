package awssecretsmanager

import (
	"github.com/aws/aws-cdk-go/awscdk/awskms"
)

// Attributes required to import an existing secret into the Stack.
//
// One ARN format (`secretArn`, `secretCompleteArn`, `secretPartialArn`) must be provided.
//
// Example:
//   var encryptionKey key
//
//   secret := secretsmanager.secret.fromSecretAttributes(this, jsii.String("ImportedSecret"), &secretAttributes{
//   	secretArn: jsii.String("arn:aws:secretsmanager:<region>:<account-id-number>:secret:<secret-name>-<random-6-characters>"),
//   	// If the secret is encrypted using a KMS-hosted CMK, either import or reference that key:
//   	encryptionKey: encryptionKey,
//   })
//
type SecretAttributes struct {
	// The encryption key that is used to encrypt the secret, unless the default SecretsManager key is used.
	EncryptionKey awskms.IKey `field:"optional" json:"encryptionKey" yaml:"encryptionKey"`
	// The complete ARN of the secret in SecretsManager.
	//
	// This is the ARN including the Secrets Manager 6-character suffix.
	// Cannot be used with `secretArn` or `secretPartialArn`.
	SecretCompleteArn *string `field:"optional" json:"secretCompleteArn" yaml:"secretCompleteArn"`
	// The partial ARN of the secret in SecretsManager.
	//
	// This is the ARN without the Secrets Manager 6-character suffix.
	// Cannot be used with `secretArn` or `secretCompleteArn`.
	SecretPartialArn *string `field:"optional" json:"secretPartialArn" yaml:"secretPartialArn"`
}

