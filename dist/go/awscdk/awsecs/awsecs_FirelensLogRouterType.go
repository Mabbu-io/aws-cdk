package awsecs


// Firelens log router type, fluentbit or fluentd.
//
// https://docs.aws.amazon.com/AmazonECS/latest/developerguide/using_firelens.html
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var appProtocol appProtocol
//   var containerImage containerImage
//   var environmentFile environmentFile
//   var linuxParameters linuxParameters
//   var logDriver logDriver
//   var secret secret
//
//   firelensLogRouterDefinitionOptions := &firelensLogRouterDefinitionOptions{
//   	firelensConfig: &firelensConfig{
//   		type: awscdk.Aws_ecs.firelensLogRouterType_FLUENTBIT,
//
//   		// the properties below are optional
//   		options: &firelensOptions{
//   			configFileType: awscdk.*Aws_ecs.firelensConfigFileType_S3,
//   			configFileValue: jsii.String("configFileValue"),
//   			enableECSLogMetadata: jsii.Boolean(false),
//   		},
//   	},
//   	image: containerImage,
//
//   	// the properties below are optional
//   	command: []*string{
//   		jsii.String("command"),
//   	},
//   	containerName: jsii.String("containerName"),
//   	cpu: jsii.Number(123),
//   	disableNetworking: jsii.Boolean(false),
//   	dnsSearchDomains: []*string{
//   		jsii.String("dnsSearchDomains"),
//   	},
//   	dnsServers: []*string{
//   		jsii.String("dnsServers"),
//   	},
//   	dockerLabels: map[string]*string{
//   		"dockerLabelsKey": jsii.String("dockerLabels"),
//   	},
//   	dockerSecurityOptions: []*string{
//   		jsii.String("dockerSecurityOptions"),
//   	},
//   	entryPoint: []*string{
//   		jsii.String("entryPoint"),
//   	},
//   	environment: map[string]*string{
//   		"environmentKey": jsii.String("environment"),
//   	},
//   	environmentFiles: []*environmentFile{
//   		environmentFile,
//   	},
//   	essential: jsii.Boolean(false),
//   	extraHosts: map[string]*string{
//   		"extraHostsKey": jsii.String("extraHosts"),
//   	},
//   	gpuCount: jsii.Number(123),
//   	healthCheck: &healthCheck{
//   		command: []*string{
//   			jsii.String("command"),
//   		},
//
//   		// the properties below are optional
//   		interval: cdk.duration.minutes(jsii.Number(30)),
//   		retries: jsii.Number(123),
//   		startPeriod: cdk.*duration.minutes(jsii.Number(30)),
//   		timeout: cdk.*duration.minutes(jsii.Number(30)),
//   	},
//   	hostname: jsii.String("hostname"),
//   	inferenceAcceleratorResources: []*string{
//   		jsii.String("inferenceAcceleratorResources"),
//   	},
//   	linuxParameters: linuxParameters,
//   	logging: logDriver,
//   	memoryLimitMiB: jsii.Number(123),
//   	memoryReservationMiB: jsii.Number(123),
//   	portMappings: []portMapping{
//   		&portMapping{
//   			containerPort: jsii.Number(123),
//
//   			// the properties below are optional
//   			appProtocol: appProtocol,
//   			hostPort: jsii.Number(123),
//   			name: jsii.String("name"),
//   			protocol: awscdk.*Aws_ecs.protocol_TCP,
//   		},
//   	},
//   	privileged: jsii.Boolean(false),
//   	readonlyRootFilesystem: jsii.Boolean(false),
//   	secrets: map[string]*secret{
//   		"secretsKey": secret,
//   	},
//   	startTimeout: cdk.*duration.minutes(jsii.Number(30)),
//   	stopTimeout: cdk.*duration.minutes(jsii.Number(30)),
//   	systemControls: []systemControl{
//   		&systemControl{
//   			namespace: jsii.String("namespace"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	user: jsii.String("user"),
//   	workingDirectory: jsii.String("workingDirectory"),
//   }
//
type FirelensLogRouterType string

const (
	// fluentbit.
	FirelensLogRouterType_FLUENTBIT FirelensLogRouterType = "FLUENTBIT"
	// fluentd.
	FirelensLogRouterType_FLUENTD FirelensLogRouterType = "FLUENTD"
)

