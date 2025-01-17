package awsfsx

import (
	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/awskms"
)

// Properties specific to the Lustre version of the FSx file system.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   var vpc vpc
//   var bucket s3.Bucket
//
//
//   lustreConfiguration := map[string]interface{}{
//   	"deploymentType": fsx.LustreDeploymentType_SCRATCH_2,
//   	"exportPath": bucket.s3UrlForObject(),
//   	"importPath": bucket.s3UrlForObject(),
//   	"autoImportPolicy": fsx.LustreAutoImportPolicy_NEW_CHANGED_DELETED,
//   }
//
//   fs := fsx.NewLustreFileSystem(this, jsii.String("FsxLustreFileSystem"), &lustreFileSystemProps{
//   	vpc: vpc,
//   	vpcSubnet: vpc.privateSubnets[jsii.Number(0)],
//   	storageCapacityGiB: jsii.Number(1200),
//   	lustreConfiguration: lustreConfiguration,
//   })
//
type LustreFileSystemProps struct {
	// The storage capacity of the file system being created.
	//
	// For Windows file systems, valid values are 32 GiB to 65,536 GiB.
	// For SCRATCH_1 deployment types, valid values are 1,200, 2,400, 3,600, then continuing in increments of 3,600 GiB.
	// For SCRATCH_2 and PERSISTENT_1 types, valid values are 1,200, 2,400, then continuing in increments of 2,400 GiB.
	StorageCapacityGiB *float64 `field:"required" json:"storageCapacityGiB" yaml:"storageCapacityGiB"`
	// The VPC to launch the file system in.
	Vpc awsec2.IVpc `field:"required" json:"vpc" yaml:"vpc"`
	// The ID of the backup.
	//
	// Specifies the backup to use if you're creating a file system from an existing backup.
	BackupId *string `field:"optional" json:"backupId" yaml:"backupId"`
	// The KMS key used for encryption to protect your data at rest.
	KmsKey awskms.IKey `field:"optional" json:"kmsKey" yaml:"kmsKey"`
	// Policy to apply when the file system is removed from the stack.
	RemovalPolicy awscdk.RemovalPolicy `field:"optional" json:"removalPolicy" yaml:"removalPolicy"`
	// Security Group to assign to this file system.
	SecurityGroup awsec2.ISecurityGroup `field:"optional" json:"securityGroup" yaml:"securityGroup"`
	// Additional configuration for FSx specific to Lustre.
	LustreConfiguration *LustreConfiguration `field:"required" json:"lustreConfiguration" yaml:"lustreConfiguration"`
	// The subnet that the file system will be accessible from.
	VpcSubnet awsec2.ISubnet `field:"required" json:"vpcSubnet" yaml:"vpcSubnet"`
}

