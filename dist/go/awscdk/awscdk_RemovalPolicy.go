// Version 2 of the AWS Cloud Development Kit library
package awscdk


// Possible values for a resource's Removal Policy.
//
// The removal policy controls what happens to the resource if it stops being
// managed by CloudFormation. This can happen in one of three situations:
//
// - The resource is removed from the template, so CloudFormation stops managing it;
// - A change to the resource is made that requires it to be replaced, so CloudFormation stops
//    managing it;
// - The stack is deleted, so CloudFormation stops managing all resources in it.
//
// The Removal Policy applies to all above cases.
//
// Many stateful resources in the AWS Construct Library will accept a
// `removalPolicy` as a property, typically defaulting it to `RETAIN`.
//
// If the AWS Construct Library resource does not accept a `removalPolicy`
// argument, you can always configure it by using the escape hatch mechanism,
// as shown in the following example:
//
// ```ts
// declare const bucket: s3.Bucket;
//
// const cfnBucket = bucket.node.findChild('Resource') as CfnResource;
// cfnBucket.applyRemovalPolicy(RemovalPolicy.DESTROY);
// ```.
//
// Example:
//   readScaling := table.autoScaleReadCapacity(&enableScalingProps{
//   	minCapacity: jsii.Number(1),
//   	maxCapacity: jsii.Number(50),
//   })
//
//   readScaling.scaleOnUtilization(&utilizationScalingProps{
//   	targetUtilizationPercent: jsii.Number(50),
//   })
//
//   readScaling.scaleOnSchedule(jsii.String("ScaleUpInTheMorning"), &scalingSchedule{
//   	schedule: appscaling.schedule.cron(&cronOptions{
//   		hour: jsii.String("8"),
//   		minute: jsii.String("0"),
//   	}),
//   	minCapacity: jsii.Number(20),
//   })
//
//   readScaling.scaleOnSchedule(jsii.String("ScaleDownAtNight"), &scalingSchedule{
//   	schedule: appscaling.*schedule.cron(&cronOptions{
//   		hour: jsii.String("20"),
//   		minute: jsii.String("0"),
//   	}),
//   	maxCapacity: jsii.Number(20),
//   })
//
type RemovalPolicy string

const (
	// This is the default removal policy.
	//
	// It means that when the resource is
	// removed from the app, it will be physically destroyed.
	RemovalPolicy_DESTROY RemovalPolicy = "DESTROY"
	// This uses the 'Retain' DeletionPolicy, which will cause the resource to be retained in the account, but orphaned from the stack.
	RemovalPolicy_RETAIN RemovalPolicy = "RETAIN"
	// This retention policy deletes the resource, but saves a snapshot of its data before deleting, so that it can be re-created later.
	//
	// Only available for some stateful resources,
	// like databases, EC2 volumes, etc.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html#aws-attribute-deletionpolicy-options
	//
	RemovalPolicy_SNAPSHOT RemovalPolicy = "SNAPSHOT"
)

