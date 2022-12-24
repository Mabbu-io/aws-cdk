package awsrds

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awskms"
	"github.com/aws/aws-cdk-go/awscdk/awssecretsmanager"
)

// Username and password combination.
//
// Example:
//   var vpc vpc
//
//
//   // Simple secret
//   secret := secretsmanager.NewSecret(this, jsii.String("Secret"))
//   // Using the secret
//   instance1 := rds.NewDatabaseInstance(this, jsii.String("PostgresInstance1"), &databaseInstanceProps{
//   	engine: rds.databaseInstanceEngine_POSTGRES(),
//   	credentials: rds.credentials.fromSecret(secret),
//   	vpc: vpc,
//   })
//   // Templated secret with username and password fields
//   templatedSecret := secretsmanager.NewSecret(this, jsii.String("TemplatedSecret"), &secretProps{
//   	generateSecretString: &secretStringGenerator{
//   		secretStringTemplate: jSON.stringify(map[string]*string{
//   			"username": jsii.String("postgres"),
//   		}),
//   		generateStringKey: jsii.String("password"),
//   	},
//   })
//   // Using the templated secret as credentials
//   instance2 := rds.NewDatabaseInstance(this, jsii.String("PostgresInstance2"), &databaseInstanceProps{
//   	engine: rds.*databaseInstanceEngine_POSTGRES(),
//   	credentials: map[string]interface{}{
//   		"username": templatedSecret.secretValueFromJson(jsii.String("username")).toString(),
//   		"password": templatedSecret.secretValueFromJson(jsii.String("password")),
//   	},
//   	vpc: vpc,
//   })
//
type Credentials interface {
	// KMS encryption key to encrypt the generated secret.
	EncryptionKey() awskms.IKey
	// The characters to exclude from the generated password.
	//
	// Only used if {@link password} has not been set.
	ExcludeCharacters() *string
	// Password.
	//
	// Do not put passwords in your CDK code directly.
	Password() awscdk.SecretValue
	// A list of regions where to replicate the generated secret.
	ReplicaRegions() *[]*awssecretsmanager.ReplicaRegion
	// Secret used to instantiate this Login.
	Secret() awssecretsmanager.ISecret
	// The name to use for the Secret if a new Secret is to be generated in SecretsManager for these Credentials.
	SecretName() *string
	// Username.
	Username() *string
	// Whether the username should be referenced as a string and not as a dynamic reference to the username in the secret.
	UsernameAsString() *bool
}

// The jsii proxy struct for Credentials
type jsiiProxy_Credentials struct {
	_ byte // padding
}

func (j *jsiiProxy_Credentials) EncryptionKey() awskms.IKey {
	var returns awskms.IKey
	_jsii_.Get(
		j,
		"encryptionKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Credentials) ExcludeCharacters() *string {
	var returns *string
	_jsii_.Get(
		j,
		"excludeCharacters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Credentials) Password() awscdk.SecretValue {
	var returns awscdk.SecretValue
	_jsii_.Get(
		j,
		"password",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Credentials) ReplicaRegions() *[]*awssecretsmanager.ReplicaRegion {
	var returns *[]*awssecretsmanager.ReplicaRegion
	_jsii_.Get(
		j,
		"replicaRegions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Credentials) Secret() awssecretsmanager.ISecret {
	var returns awssecretsmanager.ISecret
	_jsii_.Get(
		j,
		"secret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Credentials) SecretName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Credentials) Username() *string {
	var returns *string
	_jsii_.Get(
		j,
		"username",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Credentials) UsernameAsString() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"usernameAsString",
		&returns,
	)
	return returns
}


func NewCredentials_Override(c Credentials) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_rds.Credentials",
		nil, // no parameters
		c,
	)
}

// Creates Credentials with a password generated and stored in Secrets Manager.
func Credentials_FromGeneratedSecret(username *string, options *CredentialsBaseOptions) Credentials {
	_init_.Initialize()

	if err := validateCredentials_FromGeneratedSecretParameters(username, options); err != nil {
		panic(err)
	}
	var returns Credentials

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.Credentials",
		"fromGeneratedSecret",
		[]interface{}{username, options},
		&returns,
	)

	return returns
}

// Creates Credentials from a password.
//
// Do not put passwords in your CDK code directly.
func Credentials_FromPassword(username *string, password awscdk.SecretValue) Credentials {
	_init_.Initialize()

	if err := validateCredentials_FromPasswordParameters(username, password); err != nil {
		panic(err)
	}
	var returns Credentials

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.Credentials",
		"fromPassword",
		[]interface{}{username, password},
		&returns,
	)

	return returns
}

// Creates Credentials from an existing Secrets Manager ``Secret`` (or ``DatabaseSecret``).
//
// The Secret must be a JSON string with a ``username`` and ``password`` field:
// ```
// {
//    ...
//    "username": <required: username>,
//    "password": <required: password>,
// }
// ```.
func Credentials_FromSecret(secret awssecretsmanager.ISecret, username *string) Credentials {
	_init_.Initialize()

	if err := validateCredentials_FromSecretParameters(secret); err != nil {
		panic(err)
	}
	var returns Credentials

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.Credentials",
		"fromSecret",
		[]interface{}{secret, username},
		&returns,
	)

	return returns
}

// Creates Credentials for the given username, and optional password and key.
//
// If no password is provided, one will be generated and stored in Secrets Manager.
func Credentials_FromUsername(username *string, options *CredentialsFromUsernameOptions) Credentials {
	_init_.Initialize()

	if err := validateCredentials_FromUsernameParameters(username, options); err != nil {
		panic(err)
	}
	var returns Credentials

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.Credentials",
		"fromUsername",
		[]interface{}{username, options},
		&returns,
	)

	return returns
}

