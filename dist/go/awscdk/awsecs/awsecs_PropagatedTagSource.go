package awsecs


// Propagate tags from either service or task definition.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var cluster cluster
//   var containerDefinition containerDefinition
//   var containerImage containerImage
//   var hostedZone hostedZone
//   var logDriver logDriver
//   var namespace iNamespace
//   var role role
//   var secret secret
//   var vpc vpc
//
//   networkMultipleTargetGroupsServiceBaseProps := &networkMultipleTargetGroupsServiceBaseProps{
//   	cloudMapOptions: &cloudMapOptions{
//   		cloudMapNamespace: namespace,
//   		container: containerDefinition,
//   		containerPort: jsii.Number(123),
//   		dnsRecordType: awscdk.Aws_servicediscovery.dnsRecordType_A,
//   		dnsTtl: cdk.duration.minutes(jsii.Number(30)),
//   		failureThreshold: jsii.Number(123),
//   		name: jsii.String("name"),
//   	},
//   	cluster: cluster,
//   	desiredCount: jsii.Number(123),
//   	enableECSManagedTags: jsii.Boolean(false),
//   	enableExecuteCommand: jsii.Boolean(false),
//   	healthCheckGracePeriod: cdk.*duration.minutes(jsii.Number(30)),
//   	loadBalancers: []networkLoadBalancerProps{
//   		&networkLoadBalancerProps{
//   			listeners: []networkListenerProps{
//   				&networkListenerProps{
//   					name: jsii.String("name"),
//
//   					// the properties below are optional
//   					port: jsii.Number(123),
//   				},
//   			},
//   			name: jsii.String("name"),
//
//   			// the properties below are optional
//   			domainName: jsii.String("domainName"),
//   			domainZone: hostedZone,
//   			publicLoadBalancer: jsii.Boolean(false),
//   		},
//   	},
//   	propagateTags: awscdk.Aws_ecs.propagatedTagSource_SERVICE,
//   	serviceName: jsii.String("serviceName"),
//   	targetGroups: []networkTargetProps{
//   		&networkTargetProps{
//   			containerPort: jsii.Number(123),
//
//   			// the properties below are optional
//   			listener: jsii.String("listener"),
//   		},
//   	},
//   	taskImageOptions: &networkLoadBalancedTaskImageProps{
//   		image: containerImage,
//
//   		// the properties below are optional
//   		containerName: jsii.String("containerName"),
//   		containerPorts: []*f64{
//   			jsii.Number(123),
//   		},
//   		dockerLabels: map[string]*string{
//   			"dockerLabelsKey": jsii.String("dockerLabels"),
//   		},
//   		enableLogging: jsii.Boolean(false),
//   		environment: map[string]*string{
//   			"environmentKey": jsii.String("environment"),
//   		},
//   		executionRole: role,
//   		family: jsii.String("family"),
//   		logDriver: logDriver,
//   		secrets: map[string]*secret{
//   			"secretsKey": secret,
//   		},
//   		taskRole: role,
//   	},
//   	vpc: vpc,
//   }
//
type PropagatedTagSource string

const (
	// Propagate tags from service.
	PropagatedTagSource_SERVICE PropagatedTagSource = "SERVICE"
	// Propagate tags from task definition.
	PropagatedTagSource_TASK_DEFINITION PropagatedTagSource = "TASK_DEFINITION"
	// Do not propagate.
	PropagatedTagSource_NONE PropagatedTagSource = "NONE"
)

