package awsroute53

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Options available when creating a delegation relationship from one PublicHostedZone to another.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   zoneDelegationOptions := &zoneDelegationOptions{
//   	comment: jsii.String("comment"),
//   	ttl: cdk.duration.minutes(jsii.Number(30)),
//   }
//
type ZoneDelegationOptions struct {
	// A comment to add on the DNS record created to incorporate the delegation.
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// The TTL (Time To Live) of the DNS delegation record in DNS caches.
	Ttl awscdk.Duration `field:"optional" json:"ttl" yaml:"ttl"`
}

