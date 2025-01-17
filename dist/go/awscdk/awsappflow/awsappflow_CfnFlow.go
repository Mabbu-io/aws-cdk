package awsappflow

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsappflow/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A CloudFormation `AWS::AppFlow::Flow`.
//
// The `AWS::AppFlow::Flow` resource is an Amazon AppFlow resource type that specifies a new flow.
//
// > If you want to use AWS CloudFormation to create a connector profile for connectors that implement OAuth (such as Salesforce, Slack, Zendesk, and Google Analytics), you must fetch the access and refresh tokens. You can do this by implementing your own UI for OAuth, or by retrieving the tokens from elsewhere. Alternatively, you can use the Amazon AppFlow console to create the connector profile, and then use that connector profile in the flow creation CloudFormation template.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnFlow := awscdk.Aws_appflow.NewCfnFlow(this, jsii.String("MyCfnFlow"), &cfnFlowProps{
//   	destinationFlowConfigList: []interface{}{
//   		&destinationFlowConfigProperty{
//   			connectorType: jsii.String("connectorType"),
//   			destinationConnectorProperties: &destinationConnectorPropertiesProperty{
//   				customConnector: &customConnectorDestinationPropertiesProperty{
//   					entityName: jsii.String("entityName"),
//
//   					// the properties below are optional
//   					customProperties: map[string]*string{
//   						"customPropertiesKey": jsii.String("customProperties"),
//   					},
//   					errorHandlingConfig: &errorHandlingConfigProperty{
//   						bucketName: jsii.String("bucketName"),
//   						bucketPrefix: jsii.String("bucketPrefix"),
//   						failOnFirstError: jsii.Boolean(false),
//   					},
//   					idFieldNames: []*string{
//   						jsii.String("idFieldNames"),
//   					},
//   					writeOperationType: jsii.String("writeOperationType"),
//   				},
//   				eventBridge: &eventBridgeDestinationPropertiesProperty{
//   					object: jsii.String("object"),
//
//   					// the properties below are optional
//   					errorHandlingConfig: &errorHandlingConfigProperty{
//   						bucketName: jsii.String("bucketName"),
//   						bucketPrefix: jsii.String("bucketPrefix"),
//   						failOnFirstError: jsii.Boolean(false),
//   					},
//   				},
//   				lookoutMetrics: &lookoutMetricsDestinationPropertiesProperty{
//   					object: jsii.String("object"),
//   				},
//   				marketo: &marketoDestinationPropertiesProperty{
//   					object: jsii.String("object"),
//
//   					// the properties below are optional
//   					errorHandlingConfig: &errorHandlingConfigProperty{
//   						bucketName: jsii.String("bucketName"),
//   						bucketPrefix: jsii.String("bucketPrefix"),
//   						failOnFirstError: jsii.Boolean(false),
//   					},
//   				},
//   				redshift: &redshiftDestinationPropertiesProperty{
//   					intermediateBucketName: jsii.String("intermediateBucketName"),
//   					object: jsii.String("object"),
//
//   					// the properties below are optional
//   					bucketPrefix: jsii.String("bucketPrefix"),
//   					errorHandlingConfig: &errorHandlingConfigProperty{
//   						bucketName: jsii.String("bucketName"),
//   						bucketPrefix: jsii.String("bucketPrefix"),
//   						failOnFirstError: jsii.Boolean(false),
//   					},
//   				},
//   				s3: &s3DestinationPropertiesProperty{
//   					bucketName: jsii.String("bucketName"),
//
//   					// the properties below are optional
//   					bucketPrefix: jsii.String("bucketPrefix"),
//   					s3OutputFormatConfig: &s3OutputFormatConfigProperty{
//   						aggregationConfig: &aggregationConfigProperty{
//   							aggregationType: jsii.String("aggregationType"),
//   							targetFileSize: jsii.Number(123),
//   						},
//   						fileType: jsii.String("fileType"),
//   						prefixConfig: &prefixConfigProperty{
//   							pathPrefixHierarchy: []*string{
//   								jsii.String("pathPrefixHierarchy"),
//   							},
//   							prefixFormat: jsii.String("prefixFormat"),
//   							prefixType: jsii.String("prefixType"),
//   						},
//   						preserveSourceDataTyping: jsii.Boolean(false),
//   					},
//   				},
//   				salesforce: &salesforceDestinationPropertiesProperty{
//   					object: jsii.String("object"),
//
//   					// the properties below are optional
//   					dataTransferApi: jsii.String("dataTransferApi"),
//   					errorHandlingConfig: &errorHandlingConfigProperty{
//   						bucketName: jsii.String("bucketName"),
//   						bucketPrefix: jsii.String("bucketPrefix"),
//   						failOnFirstError: jsii.Boolean(false),
//   					},
//   					idFieldNames: []*string{
//   						jsii.String("idFieldNames"),
//   					},
//   					writeOperationType: jsii.String("writeOperationType"),
//   				},
//   				sapoData: &sAPODataDestinationPropertiesProperty{
//   					objectPath: jsii.String("objectPath"),
//
//   					// the properties below are optional
//   					errorHandlingConfig: &errorHandlingConfigProperty{
//   						bucketName: jsii.String("bucketName"),
//   						bucketPrefix: jsii.String("bucketPrefix"),
//   						failOnFirstError: jsii.Boolean(false),
//   					},
//   					idFieldNames: []*string{
//   						jsii.String("idFieldNames"),
//   					},
//   					successResponseHandlingConfig: &successResponseHandlingConfigProperty{
//   						bucketName: jsii.String("bucketName"),
//   						bucketPrefix: jsii.String("bucketPrefix"),
//   					},
//   					writeOperationType: jsii.String("writeOperationType"),
//   				},
//   				snowflake: &snowflakeDestinationPropertiesProperty{
//   					intermediateBucketName: jsii.String("intermediateBucketName"),
//   					object: jsii.String("object"),
//
//   					// the properties below are optional
//   					bucketPrefix: jsii.String("bucketPrefix"),
//   					errorHandlingConfig: &errorHandlingConfigProperty{
//   						bucketName: jsii.String("bucketName"),
//   						bucketPrefix: jsii.String("bucketPrefix"),
//   						failOnFirstError: jsii.Boolean(false),
//   					},
//   				},
//   				upsolver: &upsolverDestinationPropertiesProperty{
//   					bucketName: jsii.String("bucketName"),
//   					s3OutputFormatConfig: &upsolverS3OutputFormatConfigProperty{
//   						prefixConfig: &prefixConfigProperty{
//   							pathPrefixHierarchy: []*string{
//   								jsii.String("pathPrefixHierarchy"),
//   							},
//   							prefixFormat: jsii.String("prefixFormat"),
//   							prefixType: jsii.String("prefixType"),
//   						},
//
//   						// the properties below are optional
//   						aggregationConfig: &aggregationConfigProperty{
//   							aggregationType: jsii.String("aggregationType"),
//   							targetFileSize: jsii.Number(123),
//   						},
//   						fileType: jsii.String("fileType"),
//   					},
//
//   					// the properties below are optional
//   					bucketPrefix: jsii.String("bucketPrefix"),
//   				},
//   				zendesk: &zendeskDestinationPropertiesProperty{
//   					object: jsii.String("object"),
//
//   					// the properties below are optional
//   					errorHandlingConfig: &errorHandlingConfigProperty{
//   						bucketName: jsii.String("bucketName"),
//   						bucketPrefix: jsii.String("bucketPrefix"),
//   						failOnFirstError: jsii.Boolean(false),
//   					},
//   					idFieldNames: []*string{
//   						jsii.String("idFieldNames"),
//   					},
//   					writeOperationType: jsii.String("writeOperationType"),
//   				},
//   			},
//
//   			// the properties below are optional
//   			apiVersion: jsii.String("apiVersion"),
//   			connectorProfileName: jsii.String("connectorProfileName"),
//   		},
//   	},
//   	flowName: jsii.String("flowName"),
//   	sourceFlowConfig: &sourceFlowConfigProperty{
//   		connectorType: jsii.String("connectorType"),
//   		sourceConnectorProperties: &sourceConnectorPropertiesProperty{
//   			amplitude: &amplitudeSourcePropertiesProperty{
//   				object: jsii.String("object"),
//   			},
//   			customConnector: &customConnectorSourcePropertiesProperty{
//   				entityName: jsii.String("entityName"),
//
//   				// the properties below are optional
//   				customProperties: map[string]*string{
//   					"customPropertiesKey": jsii.String("customProperties"),
//   				},
//   			},
//   			datadog: &datadogSourcePropertiesProperty{
//   				object: jsii.String("object"),
//   			},
//   			dynatrace: &dynatraceSourcePropertiesProperty{
//   				object: jsii.String("object"),
//   			},
//   			googleAnalytics: &googleAnalyticsSourcePropertiesProperty{
//   				object: jsii.String("object"),
//   			},
//   			inforNexus: &inforNexusSourcePropertiesProperty{
//   				object: jsii.String("object"),
//   			},
//   			marketo: &marketoSourcePropertiesProperty{
//   				object: jsii.String("object"),
//   			},
//   			s3: &s3SourcePropertiesProperty{
//   				bucketName: jsii.String("bucketName"),
//   				bucketPrefix: jsii.String("bucketPrefix"),
//
//   				// the properties below are optional
//   				s3InputFormatConfig: &s3InputFormatConfigProperty{
//   					s3InputFileType: jsii.String("s3InputFileType"),
//   				},
//   			},
//   			salesforce: &salesforceSourcePropertiesProperty{
//   				object: jsii.String("object"),
//
//   				// the properties below are optional
//   				dataTransferApi: jsii.String("dataTransferApi"),
//   				enableDynamicFieldUpdate: jsii.Boolean(false),
//   				includeDeletedRecords: jsii.Boolean(false),
//   			},
//   			sapoData: &sAPODataSourcePropertiesProperty{
//   				objectPath: jsii.String("objectPath"),
//   			},
//   			serviceNow: &serviceNowSourcePropertiesProperty{
//   				object: jsii.String("object"),
//   			},
//   			singular: &singularSourcePropertiesProperty{
//   				object: jsii.String("object"),
//   			},
//   			slack: &slackSourcePropertiesProperty{
//   				object: jsii.String("object"),
//   			},
//   			trendmicro: &trendmicroSourcePropertiesProperty{
//   				object: jsii.String("object"),
//   			},
//   			veeva: &veevaSourcePropertiesProperty{
//   				object: jsii.String("object"),
//
//   				// the properties below are optional
//   				documentType: jsii.String("documentType"),
//   				includeAllVersions: jsii.Boolean(false),
//   				includeRenditions: jsii.Boolean(false),
//   				includeSourceFiles: jsii.Boolean(false),
//   			},
//   			zendesk: &zendeskSourcePropertiesProperty{
//   				object: jsii.String("object"),
//   			},
//   		},
//
//   		// the properties below are optional
//   		apiVersion: jsii.String("apiVersion"),
//   		connectorProfileName: jsii.String("connectorProfileName"),
//   		incrementalPullConfig: &incrementalPullConfigProperty{
//   			datetimeTypeFieldName: jsii.String("datetimeTypeFieldName"),
//   		},
//   	},
//   	tasks: []interface{}{
//   		&taskProperty{
//   			sourceFields: []*string{
//   				jsii.String("sourceFields"),
//   			},
//   			taskType: jsii.String("taskType"),
//
//   			// the properties below are optional
//   			connectorOperator: &connectorOperatorProperty{
//   				amplitude: jsii.String("amplitude"),
//   				customConnector: jsii.String("customConnector"),
//   				datadog: jsii.String("datadog"),
//   				dynatrace: jsii.String("dynatrace"),
//   				googleAnalytics: jsii.String("googleAnalytics"),
//   				inforNexus: jsii.String("inforNexus"),
//   				marketo: jsii.String("marketo"),
//   				s3: jsii.String("s3"),
//   				salesforce: jsii.String("salesforce"),
//   				sapoData: jsii.String("sapoData"),
//   				serviceNow: jsii.String("serviceNow"),
//   				singular: jsii.String("singular"),
//   				slack: jsii.String("slack"),
//   				trendmicro: jsii.String("trendmicro"),
//   				veeva: jsii.String("veeva"),
//   				zendesk: jsii.String("zendesk"),
//   			},
//   			destinationField: jsii.String("destinationField"),
//   			taskProperties: []interface{}{
//   				&taskPropertiesObjectProperty{
//   					key: jsii.String("key"),
//   					value: jsii.String("value"),
//   				},
//   			},
//   		},
//   	},
//   	triggerConfig: &triggerConfigProperty{
//   		triggerType: jsii.String("triggerType"),
//
//   		// the properties below are optional
//   		triggerProperties: &scheduledTriggerPropertiesProperty{
//   			scheduleExpression: jsii.String("scheduleExpression"),
//
//   			// the properties below are optional
//   			dataPullMode: jsii.String("dataPullMode"),
//   			firstExecutionFrom: jsii.Number(123),
//   			flowErrorDeactivationThreshold: jsii.Number(123),
//   			scheduleEndTime: jsii.Number(123),
//   			scheduleOffset: jsii.Number(123),
//   			scheduleStartTime: jsii.Number(123),
//   			timeZone: jsii.String("timeZone"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	kmsArn: jsii.String("kmsArn"),
//   	metadataCatalogConfig: &metadataCatalogConfigProperty{
//   		glueDataCatalog: &glueDataCatalogProperty{
//   			databaseName: jsii.String("databaseName"),
//   			roleArn: jsii.String("roleArn"),
//   			tablePrefix: jsii.String("tablePrefix"),
//   		},
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   })
//
type CfnFlow interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The flow's Amazon Resource Name (ARN).
	AttrFlowArn() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// A user-entered description of the flow.
	Description() *string
	SetDescription(val *string)
	// The configuration that controls how Amazon AppFlow places data in the destination connector.
	DestinationFlowConfigList() interface{}
	SetDestinationFlowConfigList(val interface{})
	// The specified name of the flow.
	//
	// Spaces are not allowed. Use underscores (_) or hyphens (-) only.
	FlowName() *string
	SetFlowName(val *string)
	// The ARN (Amazon Resource Name) of the Key Management Service (KMS) key you provide for encryption.
	//
	// This is required if you do not want to use the Amazon AppFlow-managed KMS key. If you don't provide anything here, Amazon AppFlow uses the Amazon AppFlow-managed KMS key.
	KmsArn() *string
	SetKmsArn(val *string)
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	LogicalId() *string
	// `AWS::AppFlow::Flow.MetadataCatalogConfig`.
	MetadataCatalogConfig() interface{}
	SetMetadataCatalogConfig(val interface{})
	// The tree node.
	Node() constructs.Node
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// Contains information about the configuration of the source connector used in the flow.
	SourceFlowConfig() interface{}
	SetSourceFlowConfig(val interface{})
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// The tags used to organize, track, or control access for your flow.
	Tags() awscdk.TagManager
	// A list of tasks that Amazon AppFlow performs while transferring the data in the flow run.
	Tasks() interface{}
	SetTasks(val interface{})
	// The trigger settings that determine how and when Amazon AppFlow runs the specified flow.
	TriggerConfig() interface{}
	SetTriggerConfig(val interface{})
	// Deprecated.
	// Deprecated: use `updatedProperties`
	//
	// Return properties modified after initiation
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperties() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	AddDependency(target awscdk.CfnResource)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	// Deprecated: use addDependency.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	AddMetadata(key *string, value interface{})
	// Adds an override to the synthesized CloudFormation resource.
	//
	// To add a
	// property override, either use `addPropertyOverride` or prefix `path` with
	// "Properties." (i.e. `Properties.TopicName`).
	//
	// If the override is nested, separate each nested level using a dot (.) in the path parameter.
	// If there is an array as part of the nesting, specify the index in the path.
	//
	// To include a literal `.` in the property name, prefix with a `\`. In most
	// programming languages you will need to write this as `"\\."` because the
	// `\` itself will need to be escaped.
	//
	// For example,
	// ```typescript
	// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.0.Projection.NonKeyAttributes', ['myattribute']);
	// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.1.ProjectionType', 'INCLUDE');
	// ```
	// would add the overrides
	// ```json
	// "Properties": {
	//    "GlobalSecondaryIndexes": [
	//      {
	//        "Projection": {
	//          "NonKeyAttributes": [ "myattribute" ]
	//          ...
	//        }
	//        ...
	//      },
	//      {
	//        "ProjectionType": "INCLUDE"
	//        ...
	//      },
	//    ]
	//    ...
	// }
	// ```
	//
	// The `value` argument to `addOverride` will not be processed or translated
	// in any way. Pass raw JSON values in here with the correct capitalization
	// for CloudFormation. If you pass CDK classes or structs, they will be
	// rendered with lowercased key names, and CloudFormation will reject the
	// template.
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	AddPropertyOverride(propertyPath *string, value interface{})
	// Sets the deletion policy of the resource based on the removal policy specified.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`). In some
	// cases, a snapshot can be taken of the resource prior to deletion
	// (`RemovalPolicy.SNAPSHOT`). A list of resources that support this policy
	// can be found in the following link:.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html#aws-attribute-deletionpolicy-options
	//
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Retrieves an array of resources this resource depends on.
	//
	// This assembles dependencies on resources across stacks (including nested stacks)
	// automatically.
	ObtainDependencies() *[]interface{}
	// Get a shallow copy of dependencies between this resource and other resources in the same stack.
	ObtainResourceDependencies() *[]awscdk.CfnResource
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	// Indicates that this resource no longer depends on another resource.
	//
	// This can be used for resources across stacks (including nested stacks)
	// and the dependency will automatically be removed from the relevant scope.
	RemoveDependency(target awscdk.CfnResource)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Replaces one dependency with another.
	ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource)
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	ShouldSynthesize() *bool
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnFlow
type jsiiProxy_CfnFlow struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnFlow) AttrFlowArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrFlowArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlow) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlow) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlow) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlow) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlow) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlow) DestinationFlowConfigList() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"destinationFlowConfigList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlow) FlowName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"flowName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlow) KmsArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlow) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlow) MetadataCatalogConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"metadataCatalogConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlow) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlow) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlow) SourceFlowConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sourceFlowConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlow) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlow) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlow) Tasks() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tasks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlow) TriggerConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"triggerConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlow) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlow) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}


// Create a new `AWS::AppFlow::Flow`.
func NewCfnFlow(scope constructs.Construct, id *string, props *CfnFlowProps) CfnFlow {
	_init_.Initialize()

	if err := validateNewCfnFlowParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnFlow{}

	_jsii_.Create(
		"aws-cdk-lib.aws_appflow.CfnFlow",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::AppFlow::Flow`.
func NewCfnFlow_Override(c CfnFlow, scope constructs.Construct, id *string, props *CfnFlowProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_appflow.CfnFlow",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnFlow)SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnFlow)SetDestinationFlowConfigList(val interface{}) {
	if err := j.validateSetDestinationFlowConfigListParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationFlowConfigList",
		val,
	)
}

func (j *jsiiProxy_CfnFlow)SetFlowName(val *string) {
	if err := j.validateSetFlowNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"flowName",
		val,
	)
}

func (j *jsiiProxy_CfnFlow)SetKmsArn(val *string) {
	_jsii_.Set(
		j,
		"kmsArn",
		val,
	)
}

func (j *jsiiProxy_CfnFlow)SetMetadataCatalogConfig(val interface{}) {
	if err := j.validateSetMetadataCatalogConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metadataCatalogConfig",
		val,
	)
}

func (j *jsiiProxy_CfnFlow)SetSourceFlowConfig(val interface{}) {
	if err := j.validateSetSourceFlowConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceFlowConfig",
		val,
	)
}

func (j *jsiiProxy_CfnFlow)SetTasks(val interface{}) {
	if err := j.validateSetTasksParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tasks",
		val,
	)
}

func (j *jsiiProxy_CfnFlow)SetTriggerConfig(val interface{}) {
	if err := j.validateSetTriggerConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"triggerConfig",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnFlow_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnFlow_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appflow.CfnFlow",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnFlow_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateCfnFlow_IsCfnResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appflow.CfnFlow",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func CfnFlow_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnFlow_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appflow.CfnFlow",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnFlow_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_appflow.CfnFlow",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnFlow) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnFlow) AddDependency(target awscdk.CfnResource) {
	if err := c.validateAddDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnFlow) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnFlow) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnFlow) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnFlow) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnFlow) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnFlow) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnFlow) GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference {
	if err := c.validateGetAttParameters(attributeName); err != nil {
		panic(err)
	}
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName, typeHint},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnFlow) GetMetadata(key *string) interface{} {
	if err := c.validateGetMetadataParameters(key); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnFlow) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnFlow) ObtainDependencies() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		c,
		"obtainDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnFlow) ObtainResourceDependencies() *[]awscdk.CfnResource {
	var returns *[]awscdk.CfnResource

	_jsii_.Invoke(
		c,
		"obtainResourceDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnFlow) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnFlow) RemoveDependency(target awscdk.CfnResource) {
	if err := c.validateRemoveDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"removeDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnFlow) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	if err := c.validateRenderPropertiesParameters(props); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnFlow) ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource) {
	if err := c.validateReplaceDependencyParameters(target, newTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"replaceDependency",
		[]interface{}{target, newTarget},
	)
}

func (c *jsiiProxy_CfnFlow) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnFlow) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnFlow) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

