package awsec2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Factory functions for standard Amazon Machine Image objects.
//
// Example:
//   // Pick the right Amazon Linux edition. All arguments shown are optional
//   // and will default to these values when omitted.
//   amznLinux := ec2.machineImage.latestAmazonLinux(&amazonLinuxImageProps{
//   	generation: ec2.amazonLinuxGeneration_AMAZON_LINUX,
//   	edition: ec2.amazonLinuxEdition_STANDARD,
//   	virtualization: ec2.amazonLinuxVirt_HVM,
//   	storage: ec2.amazonLinuxStorage_GENERAL_PURPOSE,
//   	cpuType: ec2.amazonLinuxCpuType_X86_64,
//   })
//
//   // Pick a Windows edition to use
//   windows := ec2.machineImage.latestWindows(ec2.windowsVersion_WINDOWS_SERVER_2019_ENGLISH_FULL_BASE)
//
//   // Read AMI id from SSM parameter store
//   ssm := ec2.machineImage.fromSsmParameter(jsii.String("/my/ami"), &ssmParameterImageOptions{
//   	os: ec2.operatingSystemType_LINUX,
//   })
//
//   // Look up the most recent image matching a set of AMI filters.
//   // In this case, look up the NAT instance AMI, by using a wildcard
//   // in the 'name' field:
//   natAmi := ec2.machineImage.lookup(&lookupMachineImageProps{
//   	name: jsii.String("amzn-ami-vpc-nat-*"),
//   	owners: []*string{
//   		jsii.String("amazon"),
//   	},
//   })
//
//   // For other custom (Linux) images, instantiate a `GenericLinuxImage` with
//   // a map giving the AMI to in for each region:
//   linux := ec2.machineImage.genericLinux(map[string]*string{
//   	"us-east-1": jsii.String("ami-97785bed"),
//   	"eu-west-1": jsii.String("ami-12345678"),
//   })
//
//   // For other custom (Windows) images, instantiate a `GenericWindowsImage` with
//   // a map giving the AMI to in for each region:
//   genericWindows := ec2.machineImage.genericWindows(map[string]*string{
//   	"us-east-1": jsii.String("ami-97785bed"),
//   	"eu-west-1": jsii.String("ami-12345678"),
//   })
//
type MachineImage interface {
}

// The jsii proxy struct for MachineImage
type jsiiProxy_MachineImage struct {
	_ byte // padding
}

func NewMachineImage_Override(m MachineImage) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ec2.MachineImage",
		nil, // no parameters
		m,
	)
}

// An image specified in SSM parameter store.
//
// By default, the SSM parameter is refreshed at every deployment,
// causing your instances to be replaced whenever a new version of the AMI
// is released.
//
// Pass `{ cachedInContext: true }` to keep the AMI ID stable. If you do, you
// will have to remember to periodically invalidate the context to refresh
// to the newest AMI ID.
func MachineImage_FromSsmParameter(parameterName *string, options *SsmParameterImageOptions) IMachineImage {
	_init_.Initialize()

	if err := validateMachineImage_FromSsmParameterParameters(parameterName, options); err != nil {
		panic(err)
	}
	var returns IMachineImage

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.MachineImage",
		"fromSsmParameter",
		[]interface{}{parameterName, options},
		&returns,
	)

	return returns
}

// A Linux image where you specify the AMI ID for every region.
func MachineImage_GenericLinux(amiMap *map[string]*string, props *GenericLinuxImageProps) IMachineImage {
	_init_.Initialize()

	if err := validateMachineImage_GenericLinuxParameters(amiMap, props); err != nil {
		panic(err)
	}
	var returns IMachineImage

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.MachineImage",
		"genericLinux",
		[]interface{}{amiMap, props},
		&returns,
	)

	return returns
}

// A Windows image where you specify the AMI ID for every region.
func MachineImage_GenericWindows(amiMap *map[string]*string, props *GenericWindowsImageProps) IMachineImage {
	_init_.Initialize()

	if err := validateMachineImage_GenericWindowsParameters(amiMap, props); err != nil {
		panic(err)
	}
	var returns IMachineImage

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.MachineImage",
		"genericWindows",
		[]interface{}{amiMap, props},
		&returns,
	)

	return returns
}

// An Amazon Linux image that is automatically kept up-to-date.
//
// This Machine Image automatically updates to the latest version on every
// deployment. Be aware this will cause your instances to be replaced when a
// new version of the image becomes available. Do not store stateful information
// on the instance if you are using this image.
func MachineImage_LatestAmazonLinux(props *AmazonLinuxImageProps) IMachineImage {
	_init_.Initialize()

	if err := validateMachineImage_LatestAmazonLinuxParameters(props); err != nil {
		panic(err)
	}
	var returns IMachineImage

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.MachineImage",
		"latestAmazonLinux",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// A Windows image that is automatically kept up-to-date.
//
// This Machine Image automatically updates to the latest version on every
// deployment. Be aware this will cause your instances to be replaced when a
// new version of the image becomes available. Do not store stateful information
// on the instance if you are using this image.
func MachineImage_LatestWindows(version WindowsVersion, props *WindowsImageProps) IMachineImage {
	_init_.Initialize()

	if err := validateMachineImage_LatestWindowsParameters(version, props); err != nil {
		panic(err)
	}
	var returns IMachineImage

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.MachineImage",
		"latestWindows",
		[]interface{}{version, props},
		&returns,
	)

	return returns
}

// Look up a shared Machine Image using DescribeImages.
//
// The most recent, available, launchable image matching the given filter
// criteria will be used. Looking up AMIs may take a long time; specify
// as many filter criteria as possible to narrow down the search.
//
// The AMI selected will be cached in `cdk.context.json` and the same value
// will be used on future runs. To refresh the AMI lookup, you will have to
// evict the value from the cache using the `cdk context` command. See
// https://docs.aws.amazon.com/cdk/latest/guide/context.html for more information.
//
// This function can not be used in environment-agnostic stacks.
func MachineImage_Lookup(props *LookupMachineImageProps) IMachineImage {
	_init_.Initialize()

	if err := validateMachineImage_LookupParameters(props); err != nil {
		panic(err)
	}
	var returns IMachineImage

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.MachineImage",
		"lookup",
		[]interface{}{props},
		&returns,
	)

	return returns
}

