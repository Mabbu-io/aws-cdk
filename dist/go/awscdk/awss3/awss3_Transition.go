package awss3

import (
	"time"

	"github.com/aws/aws-cdk-go/awscdk"
)

// Describes when an object transitions to a specified storage class.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var storageClass storageClass
//
//   transition := &transition{
//   	storageClass: storageClass,
//
//   	// the properties below are optional
//   	transitionAfter: cdk.duration.minutes(jsii.Number(30)),
//   	transitionDate: NewDate(),
//   }
//
type Transition struct {
	// The storage class to which you want the object to transition.
	StorageClass StorageClass `field:"required" json:"storageClass" yaml:"storageClass"`
	// Indicates the number of days after creation when objects are transitioned to the specified storage class.
	TransitionAfter awscdk.Duration `field:"optional" json:"transitionAfter" yaml:"transitionAfter"`
	// Indicates when objects are transitioned to the specified storage class.
	//
	// The date value must be in ISO 8601 format. The time is always midnight UTC.
	TransitionDate *time.Time `field:"optional" json:"transitionDate" yaml:"transitionDate"`
}

