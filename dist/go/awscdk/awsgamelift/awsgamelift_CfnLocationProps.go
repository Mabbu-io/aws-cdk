package awsgamelift

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnLocation`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnLocationProps := &cfnLocationProps{
//   	locationName: jsii.String("locationName"),
//
//   	// the properties below are optional
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnLocationProps struct {
	// `AWS::GameLift::Location.LocationName`.
	LocationName *string `field:"required" json:"locationName" yaml:"locationName"`
	// `AWS::GameLift::Location.Tags`.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

