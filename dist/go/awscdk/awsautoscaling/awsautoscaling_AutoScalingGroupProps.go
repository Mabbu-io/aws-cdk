package awsautoscaling

import (
	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
)

// Properties of a Fleet.
//
// Example:
//   var vpc vpc
//   var instanceType instanceType
//   var machineImage iMachineImage
//
//
//   autoscaling.NewAutoScalingGroup(this, jsii.String("ASG"), &autoScalingGroupProps{
//   	vpc: vpc,
//   	instanceType: instanceType,
//   	machineImage: machineImage,
//
//   	// ...
//
//   	init: ec2.cloudFormationInit.fromElements(ec2.initFile.fromString(jsii.String("/etc/my_instance"), jsii.String("This got written during instance startup"))),
//   	signals: autoscaling.signals.waitForAll(&signalsOptions{
//   		timeout: awscdk.Duration.minutes(jsii.Number(10)),
//   	}),
//   })
//
type AutoScalingGroupProps struct {
	// Whether the instances can initiate connections to anywhere by default.
	AllowAllOutbound *bool `field:"optional" json:"allowAllOutbound" yaml:"allowAllOutbound"`
	// Whether instances in the Auto Scaling Group should have public IP addresses associated with them.
	//
	// `launchTemplate` and `mixedInstancesPolicy` must not be specified when this property is specified.
	AssociatePublicIpAddress *bool `field:"optional" json:"associatePublicIpAddress" yaml:"associatePublicIpAddress"`
	// The name of the Auto Scaling group.
	//
	// This name must be unique per Region per account.
	AutoScalingGroupName *string `field:"optional" json:"autoScalingGroupName" yaml:"autoScalingGroupName"`
	// Specifies how block devices are exposed to the instance. You can specify virtual devices and EBS volumes.
	//
	// Each instance that is launched has an associated root device volume,
	// either an Amazon EBS volume or an instance store volume.
	// You can use block device mappings to specify additional EBS volumes or
	// instance store volumes to attach to an instance when it is launched.
	//
	// `launchTemplate` and `mixedInstancesPolicy` must not be specified when this property is specified.
	// See: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/block-device-mapping-concepts.html
	//
	BlockDevices *[]*BlockDevice `field:"optional" json:"blockDevices" yaml:"blockDevices"`
	// Default scaling cooldown for this AutoScalingGroup.
	Cooldown awscdk.Duration `field:"optional" json:"cooldown" yaml:"cooldown"`
	// The amount of time, in seconds, until a newly launched instance can contribute to the Amazon CloudWatch metrics.
	//
	// This delay lets an instance finish initializing before Amazon EC2 Auto Scaling aggregates instance metrics,
	// resulting in more reliable usage data. Set this value equal to the amount of time that it takes for resource
	// consumption to become stable after an instance reaches the InService state.
	//
	// To optimize the performance of scaling policies that scale continuously, such as target tracking and
	// step scaling policies, we strongly recommend that you enable the default instance warmup, even if its value is set to 0 seconds
	//
	// Default instance warmup will not be added if no value is specified.
	// See: https://docs.aws.amazon.com/autoscaling/ec2/userguide/ec2-auto-scaling-default-instance-warmup.html
	//
	DefaultInstanceWarmup awscdk.Duration `field:"optional" json:"defaultInstanceWarmup" yaml:"defaultInstanceWarmup"`
	// Initial amount of instances in the fleet.
	//
	// If this is set to a number, every deployment will reset the amount of
	// instances to this number. It is recommended to leave this value blank.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-group.html#cfn-as-group-desiredcapacity
	//
	DesiredCapacity *float64 `field:"optional" json:"desiredCapacity" yaml:"desiredCapacity"`
	// Enable monitoring for group metrics, these metrics describe the group rather than any of its instances.
	//
	// To report all group metrics use `GroupMetrics.all()`
	// Group metrics are reported in a granularity of 1 minute at no additional charge.
	GroupMetrics *[]GroupMetrics `field:"optional" json:"groupMetrics" yaml:"groupMetrics"`
	// Configuration for health checks.
	HealthCheck HealthCheck `field:"optional" json:"healthCheck" yaml:"healthCheck"`
	// If the ASG has scheduled actions, don't reset unchanged group sizes.
	//
	// Only used if the ASG has scheduled actions (which may scale your ASG up
	// or down regardless of cdk deployments). If true, the size of the group
	// will only be reset if it has been changed in the CDK app. If false, the
	// sizes will always be changed back to what they were in the CDK app
	// on deployment.
	IgnoreUnmodifiedSizeProperties *bool `field:"optional" json:"ignoreUnmodifiedSizeProperties" yaml:"ignoreUnmodifiedSizeProperties"`
	// Controls whether instances in this group are launched with detailed or basic monitoring.
	//
	// When detailed monitoring is enabled, Amazon CloudWatch generates metrics every minute and your account
	// is charged a fee. When you disable detailed monitoring, CloudWatch generates metrics every 5 minutes.
	//
	// `launchTemplate` and `mixedInstancesPolicy` must not be specified when this property is specified.
	// See: https://docs.aws.amazon.com/autoscaling/latest/userguide/as-instance-monitoring.html#enable-as-instance-metrics
	//
	InstanceMonitoring Monitoring `field:"optional" json:"instanceMonitoring" yaml:"instanceMonitoring"`
	// Name of SSH keypair to grant access to instances.
	//
	// `launchTemplate` and `mixedInstancesPolicy` must not be specified when this property is specified.
	KeyName *string `field:"optional" json:"keyName" yaml:"keyName"`
	// Maximum number of instances in the fleet.
	MaxCapacity *float64 `field:"optional" json:"maxCapacity" yaml:"maxCapacity"`
	// The maximum amount of time that an instance can be in service.
	//
	// The maximum duration applies
	// to all current and future instances in the group. As an instance approaches its maximum duration,
	// it is terminated and replaced, and cannot be used again.
	//
	// You must specify a value of at least 604,800 seconds (7 days). To clear a previously set value,
	// leave this property undefined.
	// See: https://docs.aws.amazon.com/autoscaling/ec2/userguide/asg-max-instance-lifetime.html
	//
	MaxInstanceLifetime awscdk.Duration `field:"optional" json:"maxInstanceLifetime" yaml:"maxInstanceLifetime"`
	// Minimum number of instances in the fleet.
	MinCapacity *float64 `field:"optional" json:"minCapacity" yaml:"minCapacity"`
	// Whether newly-launched instances are protected from termination by Amazon EC2 Auto Scaling when scaling in.
	//
	// By default, Auto Scaling can terminate an instance at any time after launch
	// when scaling in an Auto Scaling Group, subject to the group's termination
	// policy. However, you may wish to protect newly-launched instances from
	// being scaled in if they are going to run critical applications that should
	// not be prematurely terminated.
	//
	// This flag must be enabled if the Auto Scaling Group will be associated with
	// an ECS Capacity Provider with managed termination protection.
	NewInstancesProtectedFromScaleIn *bool `field:"optional" json:"newInstancesProtectedFromScaleIn" yaml:"newInstancesProtectedFromScaleIn"`
	// Configure autoscaling group to send notifications about fleet changes to an SNS topic(s).
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-group.html#cfn-as-group-notificationconfigurations
	//
	Notifications *[]*NotificationConfiguration `field:"optional" json:"notifications" yaml:"notifications"`
	// Configure waiting for signals during deployment.
	//
	// Use this to pause the CloudFormation deployment to wait for the instances
	// in the AutoScalingGroup to report successful startup during
	// creation and updates. The UserData script needs to invoke `cfn-signal`
	// with a success or failure code after it is done setting up the instance.
	//
	// Without waiting for signals, the CloudFormation deployment will proceed as
	// soon as the AutoScalingGroup has been created or updated but before the
	// instances in the group have been started.
	//
	// For example, to have instances wait for an Elastic Load Balancing health check before
	// they signal success, add a health-check verification by using the
	// cfn-init helper script. For an example, see the verify_instance_health
	// command in the Auto Scaling rolling updates sample template:
	//
	// https://github.com/awslabs/aws-cloudformation-templates/blob/master/aws/services/AutoScaling/AutoScalingRollingUpdates.yaml
	Signals Signals `field:"optional" json:"signals" yaml:"signals"`
	// The maximum hourly price (in USD) to be paid for any Spot Instance launched to fulfill the request.
	//
	// Spot Instances are
	// launched when the price you specify exceeds the current Spot market price.
	//
	// `launchTemplate` and `mixedInstancesPolicy` must not be specified when this property is specified.
	SpotPrice *string `field:"optional" json:"spotPrice" yaml:"spotPrice"`
	// A policy or a list of policies that are used to select the instances to terminate.
	//
	// The policies are executed in the order that you list them.
	// See: https://docs.aws.amazon.com/autoscaling/ec2/userguide/as-instance-termination.html
	//
	TerminationPolicies *[]TerminationPolicy `field:"optional" json:"terminationPolicies" yaml:"terminationPolicies"`
	// What to do when an AutoScalingGroup's instance configuration is changed.
	//
	// This is applied when any of the settings on the ASG are changed that
	// affect how the instances should be created (VPC, instance type, startup
	// scripts, etc.). It indicates how the existing instances should be
	// replaced with new instances matching the new config. By default, nothing
	// is done and only new instances are launched with the new config.
	UpdatePolicy UpdatePolicy `field:"optional" json:"updatePolicy" yaml:"updatePolicy"`
	// Where to place instances within the VPC.
	VpcSubnets *awsec2.SubnetSelection `field:"optional" json:"vpcSubnets" yaml:"vpcSubnets"`
	// VPC to launch these instances in.
	Vpc awsec2.IVpc `field:"required" json:"vpc" yaml:"vpc"`
	// Apply the given CloudFormation Init configuration to the instances in the AutoScalingGroup at startup.
	//
	// If you specify `init`, you must also specify `signals` to configure
	// the number of instances to wait for and the timeout for waiting for the
	// init process.
	Init awsec2.CloudFormationInit `field:"optional" json:"init" yaml:"init"`
	// Use the given options for applying CloudFormation Init.
	//
	// Describes the configsets to use and the timeout to wait.
	InitOptions *ApplyCloudFormationInitOptions `field:"optional" json:"initOptions" yaml:"initOptions"`
	// Type of instance to launch.
	//
	// `launchTemplate` and `mixedInstancesPolicy` must not be specified when this property is specified.
	InstanceType awsec2.InstanceType `field:"optional" json:"instanceType" yaml:"instanceType"`
	// Launch template to use.
	//
	// Launch configuration related settings and MixedInstancesPolicy must not be specified when a
	// launch template is specified.
	LaunchTemplate awsec2.ILaunchTemplate `field:"optional" json:"launchTemplate" yaml:"launchTemplate"`
	// AMI to launch.
	//
	// `launchTemplate` and `mixedInstancesPolicy` must not be specified when this property is specified.
	MachineImage awsec2.IMachineImage `field:"optional" json:"machineImage" yaml:"machineImage"`
	// Mixed Instances Policy to use.
	//
	// Launch configuration related settings and Launch Template  must not be specified when a
	// MixedInstancesPolicy is specified.
	MixedInstancesPolicy *MixedInstancesPolicy `field:"optional" json:"mixedInstancesPolicy" yaml:"mixedInstancesPolicy"`
	// Whether IMDSv2 should be required on launched instances.
	RequireImdsv2 *bool `field:"optional" json:"requireImdsv2" yaml:"requireImdsv2"`
	// An IAM role to associate with the instance profile assigned to this Auto Scaling Group.
	//
	// The role must be assumable by the service principal `ec2.amazonaws.com`:
	//
	// `launchTemplate` and `mixedInstancesPolicy` must not be specified when this property is specified.
	//
	// Example:
	//   role := iam.NewRole(this, jsii.String("MyRole"), &roleProps{
	//   	assumedBy: iam.NewServicePrincipal(jsii.String("ec2.amazonaws.com")),
	//   })
	//
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// Security group to launch the instances in.
	//
	// `launchTemplate` and `mixedInstancesPolicy` must not be specified when this property is specified.
	SecurityGroup awsec2.ISecurityGroup `field:"optional" json:"securityGroup" yaml:"securityGroup"`
	// Specific UserData to use.
	//
	// The UserData may still be mutated after creation.
	//
	// `launchTemplate` and `mixedInstancesPolicy` must not be specified when this property is specified.
	UserData awsec2.UserData `field:"optional" json:"userData" yaml:"userData"`
}

