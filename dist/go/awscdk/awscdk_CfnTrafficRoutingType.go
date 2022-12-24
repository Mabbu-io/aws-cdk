// Version 2 of the AWS Cloud Development Kit library
package awscdk


// The possible types of traffic shifting for the blue-green deployment configuration.
//
// The type of the {@link CfnTrafficRoutingConfig.type} property.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnCodeDeployBlueGreenHookProps := &cfnCodeDeployBlueGreenHookProps{
//   	applications: []cfnCodeDeployBlueGreenApplication{
//   		&cfnCodeDeployBlueGreenApplication{
//   			ecsAttributes: &cfnCodeDeployBlueGreenEcsAttributes{
//   				taskDefinitions: []*string{
//   					jsii.String("taskDefinitions"),
//   				},
//   				taskSets: []*string{
//   					jsii.String("taskSets"),
//   				},
//   				trafficRouting: &cfnTrafficRouting{
//   					prodTrafficRoute: &cfnTrafficRoute{
//   						logicalId: jsii.String("logicalId"),
//   						type: jsii.String("type"),
//   					},
//   					targetGroups: []*string{
//   						jsii.String("targetGroups"),
//   					},
//   					testTrafficRoute: &cfnTrafficRoute{
//   						logicalId: jsii.String("logicalId"),
//   						type: jsii.String("type"),
//   					},
//   				},
//   			},
//   			target: &cfnCodeDeployBlueGreenApplicationTarget{
//   				logicalId: jsii.String("logicalId"),
//   				type: jsii.String("type"),
//   			},
//   		},
//   	},
//   	serviceRole: jsii.String("serviceRole"),
//
//   	// the properties below are optional
//   	additionalOptions: &cfnCodeDeployBlueGreenAdditionalOptions{
//   		terminationWaitTimeInMinutes: jsii.Number(123),
//   	},
//   	lifecycleEventHooks: &cfnCodeDeployBlueGreenLifecycleEventHooks{
//   		afterAllowTestTraffic: jsii.String("afterAllowTestTraffic"),
//   		afterAllowTraffic: jsii.String("afterAllowTraffic"),
//   		afterInstall: jsii.String("afterInstall"),
//   		beforeAllowTraffic: jsii.String("beforeAllowTraffic"),
//   		beforeInstall: jsii.String("beforeInstall"),
//   	},
//   	trafficRoutingConfig: &cfnTrafficRoutingConfig{
//   		type: cdk.cfnTrafficRoutingType_ALL_AT_ONCE,
//
//   		// the properties below are optional
//   		timeBasedCanary: &cfnTrafficRoutingTimeBasedCanary{
//   			bakeTimeMins: jsii.Number(123),
//   			stepPercentage: jsii.Number(123),
//   		},
//   		timeBasedLinear: &cfnTrafficRoutingTimeBasedLinear{
//   			bakeTimeMins: jsii.Number(123),
//   			stepPercentage: jsii.Number(123),
//   		},
//   	},
//   }
//
type CfnTrafficRoutingType string

const (
	// Switch from blue to green at once.
	CfnTrafficRoutingType_ALL_AT_ONCE CfnTrafficRoutingType = "ALL_AT_ONCE"
	// Specifies a configuration that shifts traffic from blue to green in two increments.
	CfnTrafficRoutingType_TIME_BASED_CANARY CfnTrafficRoutingType = "TIME_BASED_CANARY"
	// Specifies a configuration that shifts traffic from blue to green in equal increments, with an equal number of minutes between each increment.
	CfnTrafficRoutingType_TIME_BASED_LINEAR CfnTrafficRoutingType = "TIME_BASED_LINEAR"
)

