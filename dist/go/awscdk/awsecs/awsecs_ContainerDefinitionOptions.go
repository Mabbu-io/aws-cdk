package awsecs

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Example:
//   var taskDefinition taskDefinition
//   var cluster cluster
//
//
//   // Add a container to the task definition
//   specificContainer := taskDefinition.addContainer(jsii.String("Container"), &containerDefinitionOptions{
//   	image: ecs.containerImage.fromRegistry(jsii.String("/aws/aws-example-app")),
//   	memoryLimitMiB: jsii.Number(2048),
//   })
//
//   // Add a port mapping
//   specificContainer.addPortMappings(&portMapping{
//   	containerPort: jsii.Number(7600),
//   	protocol: ecs.protocol_TCP,
//   })
//
//   ecs.NewEc2Service(this, jsii.String("Service"), &ec2ServiceProps{
//   	cluster: cluster,
//   	taskDefinition: taskDefinition,
//   	cloudMapOptions: &cloudMapOptions{
//   		// Create SRV records - useful for bridge networking
//   		dnsRecordType: cloudmap.dnsRecordType_SRV,
//   		// Targets port TCP port 7600 `specificContainer`
//   		container: specificContainer,
//   		containerPort: jsii.Number(7600),
//   	},
//   })
//
type ContainerDefinitionOptions struct {
	// The image used to start a container.
	//
	// This string is passed directly to the Docker daemon.
	// Images in the Docker Hub registry are available by default.
	// Other repositories are specified with either repository-url/image:tag or repository-url/image@digest.
	// TODO: Update these to specify using classes of IContainerImage.
	Image ContainerImage `field:"required" json:"image" yaml:"image"`
	// The command that is passed to the container.
	//
	// If you provide a shell command as a single string, you have to quote command-line arguments.
	Command *[]*string `field:"optional" json:"command" yaml:"command"`
	// The name of the container.
	ContainerName *string `field:"optional" json:"containerName" yaml:"containerName"`
	// The minimum number of CPU units to reserve for the container.
	Cpu *float64 `field:"optional" json:"cpu" yaml:"cpu"`
	// Specifies whether networking is disabled within the container.
	//
	// When this parameter is true, networking is disabled within the container.
	DisableNetworking *bool `field:"optional" json:"disableNetworking" yaml:"disableNetworking"`
	// A list of DNS search domains that are presented to the container.
	DnsSearchDomains *[]*string `field:"optional" json:"dnsSearchDomains" yaml:"dnsSearchDomains"`
	// A list of DNS servers that are presented to the container.
	DnsServers *[]*string `field:"optional" json:"dnsServers" yaml:"dnsServers"`
	// A key/value map of labels to add to the container.
	DockerLabels *map[string]*string `field:"optional" json:"dockerLabels" yaml:"dockerLabels"`
	// A list of strings to provide custom labels for SELinux and AppArmor multi-level security systems.
	DockerSecurityOptions *[]*string `field:"optional" json:"dockerSecurityOptions" yaml:"dockerSecurityOptions"`
	// The ENTRYPOINT value to pass to the container.
	// See: https://docs.docker.com/engine/reference/builder/#entrypoint
	//
	EntryPoint *[]*string `field:"optional" json:"entryPoint" yaml:"entryPoint"`
	// The environment variables to pass to the container.
	Environment *map[string]*string `field:"optional" json:"environment" yaml:"environment"`
	// The environment files to pass to the container.
	// See: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/taskdef-envfiles.html
	//
	EnvironmentFiles *[]EnvironmentFile `field:"optional" json:"environmentFiles" yaml:"environmentFiles"`
	// Specifies whether the container is marked essential.
	//
	// If the essential parameter of a container is marked as true, and that container fails
	// or stops for any reason, all other containers that are part of the task are stopped.
	// If the essential parameter of a container is marked as false, then its failure does not
	// affect the rest of the containers in a task. All tasks must have at least one essential container.
	//
	// If this parameter is omitted, a container is assumed to be essential.
	Essential *bool `field:"optional" json:"essential" yaml:"essential"`
	// A list of hostnames and IP address mappings to append to the /etc/hosts file on the container.
	ExtraHosts *map[string]*string `field:"optional" json:"extraHosts" yaml:"extraHosts"`
	// The number of GPUs assigned to the container.
	GpuCount *float64 `field:"optional" json:"gpuCount" yaml:"gpuCount"`
	// The health check command and associated configuration parameters for the container.
	HealthCheck *HealthCheck `field:"optional" json:"healthCheck" yaml:"healthCheck"`
	// The hostname to use for your container.
	Hostname *string `field:"optional" json:"hostname" yaml:"hostname"`
	// The inference accelerators referenced by the container.
	InferenceAcceleratorResources *[]*string `field:"optional" json:"inferenceAcceleratorResources" yaml:"inferenceAcceleratorResources"`
	// Linux-specific modifications that are applied to the container, such as Linux kernel capabilities.
	//
	// For more information see [KernelCapabilities](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_KernelCapabilities.html).
	LinuxParameters LinuxParameters `field:"optional" json:"linuxParameters" yaml:"linuxParameters"`
	// The log configuration specification for the container.
	Logging LogDriver `field:"optional" json:"logging" yaml:"logging"`
	// The amount (in MiB) of memory to present to the container.
	//
	// If your container attempts to exceed the allocated memory, the container
	// is terminated.
	//
	// At least one of memoryLimitMiB and memoryReservationMiB is required for non-Fargate services.
	MemoryLimitMiB *float64 `field:"optional" json:"memoryLimitMiB" yaml:"memoryLimitMiB"`
	// The soft limit (in MiB) of memory to reserve for the container.
	//
	// When system memory is under heavy contention, Docker attempts to keep the
	// container memory to this soft limit. However, your container can consume more
	// memory when it needs to, up to either the hard limit specified with the memory
	// parameter (if applicable), or all of the available memory on the container
	// instance, whichever comes first.
	//
	// At least one of memoryLimitMiB and memoryReservationMiB is required for non-Fargate services.
	MemoryReservationMiB *float64 `field:"optional" json:"memoryReservationMiB" yaml:"memoryReservationMiB"`
	// The port mappings to add to the container definition.
	PortMappings *[]*PortMapping `field:"optional" json:"portMappings" yaml:"portMappings"`
	// Specifies whether the container is marked as privileged.
	//
	// When this parameter is true, the container is given elevated privileges on the host container instance (similar to the root user).
	Privileged *bool `field:"optional" json:"privileged" yaml:"privileged"`
	// When this parameter is true, the container is given read-only access to its root file system.
	ReadonlyRootFilesystem *bool `field:"optional" json:"readonlyRootFilesystem" yaml:"readonlyRootFilesystem"`
	// The secret environment variables to pass to the container.
	Secrets *map[string]Secret `field:"optional" json:"secrets" yaml:"secrets"`
	// Time duration (in seconds) to wait before giving up on resolving dependencies for a container.
	StartTimeout awscdk.Duration `field:"optional" json:"startTimeout" yaml:"startTimeout"`
	// Time duration (in seconds) to wait before the container is forcefully killed if it doesn't exit normally on its own.
	StopTimeout awscdk.Duration `field:"optional" json:"stopTimeout" yaml:"stopTimeout"`
	// A list of namespaced kernel parameters to set in the container.
	// See: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task_definition_parameters.html#container_definition_systemcontrols
	//
	SystemControls *[]*SystemControl `field:"optional" json:"systemControls" yaml:"systemControls"`
	// The user name to use inside the container.
	User *string `field:"optional" json:"user" yaml:"user"`
	// The working directory in which to run commands inside the container.
	WorkingDirectory *string `field:"optional" json:"workingDirectory" yaml:"workingDirectory"`
}

