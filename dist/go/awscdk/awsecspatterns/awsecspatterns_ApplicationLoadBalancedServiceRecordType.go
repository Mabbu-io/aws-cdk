package awsecspatterns


// Describes the type of DNS record the service should create.
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
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var applicationLoadBalancer applicationLoadBalancer
//   var certificate certificate
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
//   applicationLoadBalancedServiceBaseProps := &applicationLoadBalancedServiceBaseProps{
//   	capacityProviderStrategies: []capacityProviderStrategy{
//   		&capacityProviderStrategy{
//   			capacityProvider: jsii.String("capacityProvider"),
//
//   			// the properties below are optional
//   			base: jsii.Number(123),
//   			weight: jsii.Number(123),
//   		},
//   	},
//   	certificate: certificate,
//   	circuitBreaker: &deploymentCircuitBreaker{
//   		rollback: jsii.Boolean(false),
//   	},
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
//   	deploymentController: &deploymentController{
//   		type: awscdk.Aws_ecs.deploymentControllerType_ECS,
//   	},
//   	desiredCount: jsii.Number(123),
//   	domainName: jsii.String("domainName"),
//   	domainZone: hostedZone,
//   	enableECSManagedTags: jsii.Boolean(false),
//   	enableExecuteCommand: jsii.Boolean(false),
//   	healthCheckGracePeriod: cdk.*duration.minutes(jsii.Number(30)),
//   	idleTimeout: cdk.*duration.minutes(jsii.Number(30)),
//   	listenerPort: jsii.Number(123),
//   	loadBalancer: applicationLoadBalancer,
//   	loadBalancerName: jsii.String("loadBalancerName"),
//   	maxHealthyPercent: jsii.Number(123),
//   	minHealthyPercent: jsii.Number(123),
//   	openListener: jsii.Boolean(false),
//   	propagateTags: awscdk.*Aws_ecs.propagatedTagSource_SERVICE,
//   	protocol: awscdk.Aws_elasticloadbalancingv2.applicationProtocol_HTTP,
//   	protocolVersion: awscdk.*Aws_elasticloadbalancingv2.applicationProtocolVersion_GRPC,
//   	publicLoadBalancer: jsii.Boolean(false),
//   	recordType: awscdk.Aws_ecs_patterns.applicationLoadBalancedServiceRecordType_ALIAS,
//   	redirectHTTP: jsii.Boolean(false),
//   	serviceName: jsii.String("serviceName"),
//   	sslPolicy: awscdk.*Aws_elasticloadbalancingv2.sslPolicy_RECOMMENDED_TLS,
//   	targetProtocol: awscdk.*Aws_elasticloadbalancingv2.*applicationProtocol_HTTP,
//   	taskImageOptions: &applicationLoadBalancedTaskImageOptions{
//   		image: containerImage,
//
//   		// the properties below are optional
//   		command: []*string{
//   			jsii.String("command"),
//   		},
//   		containerName: jsii.String("containerName"),
//   		containerPort: jsii.Number(123),
//   		dockerLabels: map[string]*string{
//   			"dockerLabelsKey": jsii.String("dockerLabels"),
//   		},
//   		enableLogging: jsii.Boolean(false),
//   		entryPoint: []*string{
//   			jsii.String("entryPoint"),
//   		},
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
type ApplicationLoadBalancedServiceRecordType string

const (
	// Create Route53 A Alias record.
	ApplicationLoadBalancedServiceRecordType_ALIAS ApplicationLoadBalancedServiceRecordType = "ALIAS"
	// Create a CNAME record.
	ApplicationLoadBalancedServiceRecordType_CNAME ApplicationLoadBalancedServiceRecordType = "CNAME"
	// Do not create any DNS records.
	ApplicationLoadBalancedServiceRecordType_NONE ApplicationLoadBalancedServiceRecordType = "NONE"
)

