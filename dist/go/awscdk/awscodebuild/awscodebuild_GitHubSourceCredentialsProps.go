package awscodebuild

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Creation properties for {@link GitHubSourceCredentials}.
//
// Example:
//   codebuild.NewGitHubSourceCredentials(this, jsii.String("CodeBuildGitHubCreds"), &gitHubSourceCredentialsProps{
//   	accessToken: awscdk.SecretValue.secretsManager(jsii.String("my-token")),
//   })
//
type GitHubSourceCredentialsProps struct {
	// The personal access token to use when contacting the GitHub API.
	AccessToken awscdk.SecretValue `field:"required" json:"accessToken" yaml:"accessToken"`
}

