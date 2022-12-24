package awsautoscaling


// The monitoring mode for instances launched in an autoscaling group.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var blockDeviceVolume blockDeviceVolume
//   var groupMetrics groupMetrics
//   var healthCheck healthCheck
//   var scalingEvents scalingEvents
//   var signals signals
//   var subnet subnet
//   var subnetFilter subnetFilter
//   var topic topic
//   var updatePolicy updatePolicy
//
//   commonAutoScalingGroupProps := &commonAutoScalingGroupProps{
//   	allowAllOutbound: jsii.Boolean(false),
//   	associatePublicIpAddress: jsii.Boolean(false),
//   	autoScalingGroupName: jsii.String("autoScalingGroupName"),
//   	blockDevices: []blockDevice{
//   		&blockDevice{
//   			deviceName: jsii.String("deviceName"),
//   			volume: blockDeviceVolume,
//   		},
//   	},
//   	cooldown: cdk.duration.minutes(jsii.Number(30)),
//   	defaultInstanceWarmup: cdk.*duration.minutes(jsii.Number(30)),
//   	desiredCapacity: jsii.Number(123),
//   	groupMetrics: []*groupMetrics{
//   		groupMetrics,
//   	},
//   	healthCheck: healthCheck,
//   	ignoreUnmodifiedSizeProperties: jsii.Boolean(false),
//   	instanceMonitoring: awscdk.Aws_autoscaling.monitoring_BASIC,
//   	keyName: jsii.String("keyName"),
//   	maxCapacity: jsii.Number(123),
//   	maxInstanceLifetime: cdk.*duration.minutes(jsii.Number(30)),
//   	minCapacity: jsii.Number(123),
//   	newInstancesProtectedFromScaleIn: jsii.Boolean(false),
//   	notifications: []notificationConfiguration{
//   		&notificationConfiguration{
//   			topic: topic,
//
//   			// the properties below are optional
//   			scalingEvents: scalingEvents,
//   		},
//   	},
//   	signals: signals,
//   	spotPrice: jsii.String("spotPrice"),
//   	terminationPolicies: []terminationPolicy{
//   		awscdk.*Aws_autoscaling.*terminationPolicy_ALLOCATION_STRATEGY,
//   	},
//   	updatePolicy: updatePolicy,
//   	vpcSubnets: &subnetSelection{
//   		availabilityZones: []*string{
//   			jsii.String("availabilityZones"),
//   		},
//   		onePerAz: jsii.Boolean(false),
//   		subnetFilters: []*subnetFilter{
//   			subnetFilter,
//   		},
//   		subnetGroupName: jsii.String("subnetGroupName"),
//   		subnets: []iSubnet{
//   			subnet,
//   		},
//   		subnetType: awscdk.Aws_ec2.subnetType_PRIVATE_ISOLATED,
//   	},
//   }
//
type Monitoring string

const (
	// Generates metrics every 5 minutes.
	Monitoring_BASIC Monitoring = "BASIC"
	// Generates metrics every minute.
	Monitoring_DETAILED Monitoring = "DETAILED"
)

