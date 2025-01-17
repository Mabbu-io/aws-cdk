package awsrds

import (
	"github.com/aws/aws-cdk-go/awscdk/awsec2"
)

// Configuration properties for an option.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var securityGroup securityGroup
//   var vpc vpc
//
//   optionConfiguration := &optionConfiguration{
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	port: jsii.Number(123),
//   	securityGroups: []iSecurityGroup{
//   		securityGroup,
//   	},
//   	settings: map[string]*string{
//   		"settingsKey": jsii.String("settings"),
//   	},
//   	version: jsii.String("version"),
//   	vpc: vpc,
//   }
//
type OptionConfiguration struct {
	// The name of the option.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The port number that this option uses.
	//
	// If `port` is specified then `vpc`
	// must also be specified.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// Optional list of security groups to use for this option, if `vpc` is specified.
	//
	// If no groups are provided, a default one will be created.
	SecurityGroups *[]awsec2.ISecurityGroup `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// The settings for the option.
	Settings *map[string]*string `field:"optional" json:"settings" yaml:"settings"`
	// The version for the option.
	Version *string `field:"optional" json:"version" yaml:"version"`
	// The VPC where a security group should be created for this option.
	//
	// If `vpc`
	// is specified then `port` must also be specified.
	Vpc awsec2.IVpc `field:"optional" json:"vpc" yaml:"vpc"`
}

