package awsstepfunctionstasks


// EBS Volume Types.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var configurationProperty_ configurationProperty
//   var size size
//
//   instanceTypeConfigProperty := &instanceTypeConfigProperty{
//   	instanceType: jsii.String("instanceType"),
//
//   	// the properties below are optional
//   	bidPrice: jsii.String("bidPrice"),
//   	bidPriceAsPercentageOfOnDemandPrice: jsii.Number(123),
//   	configurations: []*configurationProperty{
//   		&configurationProperty{
//   			classification: jsii.String("classification"),
//   			configurations: []*configurationProperty{
//   				configurationProperty_,
//   			},
//   			properties: map[string]*string{
//   				"propertiesKey": jsii.String("properties"),
//   			},
//   		},
//   	},
//   	ebsConfiguration: &ebsConfigurationProperty{
//   		ebsBlockDeviceConfigs: []ebsBlockDeviceConfigProperty{
//   			&ebsBlockDeviceConfigProperty{
//   				volumeSpecification: &volumeSpecificationProperty{
//   					volumeSize: size,
//   					volumeType: awscdk.Aws_stepfunctions_tasks.emrCreateCluster.ebsBlockDeviceVolumeType_GP2,
//
//   					// the properties below are optional
//   					iops: jsii.Number(123),
//   				},
//
//   				// the properties below are optional
//   				volumesPerInstance: jsii.Number(123),
//   			},
//   		},
//   		ebsOptimized: jsii.Boolean(false),
//   	},
//   	weightedCapacity: jsii.Number(123),
//   }
//
type EmrCreateCluster_EbsBlockDeviceVolumeType string

const (
	// gp2 Volume Type.
	EmrCreateCluster_EbsBlockDeviceVolumeType_GP2 EmrCreateCluster_EbsBlockDeviceVolumeType = "GP2"
	// io1 Volume Type.
	EmrCreateCluster_EbsBlockDeviceVolumeType_IO1 EmrCreateCluster_EbsBlockDeviceVolumeType = "IO1"
	// Standard Volume Type.
	EmrCreateCluster_EbsBlockDeviceVolumeType_STANDARD EmrCreateCluster_EbsBlockDeviceVolumeType = "STANDARD"
)

