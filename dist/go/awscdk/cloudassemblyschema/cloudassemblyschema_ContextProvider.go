package cloudassemblyschema


// Identifier for the context provider.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   assemblyManifest := &assemblyManifest{
//   	version: jsii.String("version"),
//
//   	// the properties below are optional
//   	artifacts: map[string]artifactManifest{
//   		"artifactsKey": &artifactManifest{
//   			"type": awscdk.cloud_assembly_schema.ArtifactType_NONE,
//
//   			// the properties below are optional
//   			"dependencies": []*string{
//   				jsii.String("dependencies"),
//   			},
//   			"displayName": jsii.String("displayName"),
//   			"environment": jsii.String("environment"),
//   			"metadata": map[string][]MetadataEntry{
//   				"metadataKey": []MetadataEntry{
//   					&MetadataEntry{
//   						"type": jsii.String("type"),
//
//   						// the properties below are optional
//   						"data": jsii.String("data"),
//   						"trace": []*string{
//   							jsii.String("trace"),
//   						},
//   					},
//   				},
//   			},
//   			"properties": &AwsCloudFormationStackProperties{
//   				"templateFile": jsii.String("templateFile"),
//
//   				// the properties below are optional
//   				"assumeRoleArn": jsii.String("assumeRoleArn"),
//   				"assumeRoleExternalId": jsii.String("assumeRoleExternalId"),
//   				"bootstrapStackVersionSsmParameter": jsii.String("bootstrapStackVersionSsmParameter"),
//   				"cloudFormationExecutionRoleArn": jsii.String("cloudFormationExecutionRoleArn"),
//   				"lookupRole": &BootstrapRole{
//   					"arn": jsii.String("arn"),
//
//   					// the properties below are optional
//   					"assumeRoleExternalId": jsii.String("assumeRoleExternalId"),
//   					"bootstrapStackVersionSsmParameter": jsii.String("bootstrapStackVersionSsmParameter"),
//   					"requiresBootstrapStackVersion": jsii.Number(123),
//   				},
//   				"parameters": map[string]*string{
//   					"parametersKey": jsii.String("parameters"),
//   				},
//   				"requiresBootstrapStackVersion": jsii.Number(123),
//   				"stackName": jsii.String("stackName"),
//   				"stackTemplateAssetObjectUrl": jsii.String("stackTemplateAssetObjectUrl"),
//   				"tags": map[string]*string{
//   					"tagsKey": jsii.String("tags"),
//   				},
//   				"terminationProtection": jsii.Boolean(false),
//   				"validateOnSynth": jsii.Boolean(false),
//   			},
//   		},
//   	},
//   	missing: []missingContext{
//   		&missingContext{
//   			key: jsii.String("key"),
//   			props: &amiContextQuery{
//   				account: jsii.String("account"),
//   				filters: map[string][]*string{
//   					"filtersKey": []*string{
//   						jsii.String("filters"),
//   					},
//   				},
//   				region: jsii.String("region"),
//
//   				// the properties below are optional
//   				lookupRoleArn: jsii.String("lookupRoleArn"),
//   				owners: []interface{}{
//   					jsii.String("owners"),
//   				},
//   			},
//   			provider: awscdk.Cloud_assembly_schema.contextProvider_AMI_PROVIDER,
//   		},
//   	},
//   	runtime: &runtimeInfo{
//   		libraries: map[string]*string{
//   			"librariesKey": jsii.String("libraries"),
//   		},
//   	},
//   }
//
type ContextProvider string

const (
	// AMI provider.
	ContextProvider_AMI_PROVIDER ContextProvider = "AMI_PROVIDER"
	// AZ provider.
	ContextProvider_AVAILABILITY_ZONE_PROVIDER ContextProvider = "AVAILABILITY_ZONE_PROVIDER"
	// Route53 Hosted Zone provider.
	ContextProvider_HOSTED_ZONE_PROVIDER ContextProvider = "HOSTED_ZONE_PROVIDER"
	// SSM Parameter Provider.
	ContextProvider_SSM_PARAMETER_PROVIDER ContextProvider = "SSM_PARAMETER_PROVIDER"
	// VPC Provider.
	ContextProvider_VPC_PROVIDER ContextProvider = "VPC_PROVIDER"
	// VPC Endpoint Service AZ Provider.
	ContextProvider_ENDPOINT_SERVICE_AVAILABILITY_ZONE_PROVIDER ContextProvider = "ENDPOINT_SERVICE_AVAILABILITY_ZONE_PROVIDER"
	// Load balancer provider.
	ContextProvider_LOAD_BALANCER_PROVIDER ContextProvider = "LOAD_BALANCER_PROVIDER"
	// Load balancer listener provider.
	ContextProvider_LOAD_BALANCER_LISTENER_PROVIDER ContextProvider = "LOAD_BALANCER_LISTENER_PROVIDER"
	// Security group provider.
	ContextProvider_SECURITY_GROUP_PROVIDER ContextProvider = "SECURITY_GROUP_PROVIDER"
	// KMS Key Provider.
	ContextProvider_KEY_PROVIDER ContextProvider = "KEY_PROVIDER"
	// A plugin provider (the actual plugin name will be in the properties).
	ContextProvider_PLUGIN ContextProvider = "PLUGIN"
)

