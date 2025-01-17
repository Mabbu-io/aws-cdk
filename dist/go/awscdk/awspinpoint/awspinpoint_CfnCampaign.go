package awspinpoint

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awspinpoint/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A CloudFormation `AWS::Pinpoint::Campaign`.
//
// Specifies the settings for a campaign. A *campaign* is a messaging initiative that engages a specific segment of users for an Amazon Pinpoint application.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var attributes interface{}
//   var customConfig interface{}
//   var metrics interface{}
//   var tags interface{}
//
//   cfnCampaign := awscdk.Aws_pinpoint.NewCfnCampaign(this, jsii.String("MyCfnCampaign"), &cfnCampaignProps{
//   	applicationId: jsii.String("applicationId"),
//   	name: jsii.String("name"),
//   	schedule: &scheduleProperty{
//   		endTime: jsii.String("endTime"),
//   		eventFilter: &campaignEventFilterProperty{
//   			dimensions: &eventDimensionsProperty{
//   				attributes: attributes,
//   				eventType: &setDimensionProperty{
//   					dimensionType: jsii.String("dimensionType"),
//   					values: []*string{
//   						jsii.String("values"),
//   					},
//   				},
//   				metrics: metrics,
//   			},
//   			filterType: jsii.String("filterType"),
//   		},
//   		frequency: jsii.String("frequency"),
//   		isLocalTime: jsii.Boolean(false),
//   		quietTime: &quietTimeProperty{
//   			end: jsii.String("end"),
//   			start: jsii.String("start"),
//   		},
//   		startTime: jsii.String("startTime"),
//   		timeZone: jsii.String("timeZone"),
//   	},
//   	segmentId: jsii.String("segmentId"),
//
//   	// the properties below are optional
//   	additionalTreatments: []interface{}{
//   		&writeTreatmentResourceProperty{
//   			customDeliveryConfiguration: &customDeliveryConfigurationProperty{
//   				deliveryUri: jsii.String("deliveryUri"),
//   				endpointTypes: []*string{
//   					jsii.String("endpointTypes"),
//   				},
//   			},
//   			messageConfiguration: &messageConfigurationProperty{
//   				admMessage: &messageProperty{
//   					action: jsii.String("action"),
//   					body: jsii.String("body"),
//   					imageIconUrl: jsii.String("imageIconUrl"),
//   					imageSmallIconUrl: jsii.String("imageSmallIconUrl"),
//   					imageUrl: jsii.String("imageUrl"),
//   					jsonBody: jsii.String("jsonBody"),
//   					mediaUrl: jsii.String("mediaUrl"),
//   					rawContent: jsii.String("rawContent"),
//   					silentPush: jsii.Boolean(false),
//   					timeToLive: jsii.Number(123),
//   					title: jsii.String("title"),
//   					url: jsii.String("url"),
//   				},
//   				apnsMessage: &messageProperty{
//   					action: jsii.String("action"),
//   					body: jsii.String("body"),
//   					imageIconUrl: jsii.String("imageIconUrl"),
//   					imageSmallIconUrl: jsii.String("imageSmallIconUrl"),
//   					imageUrl: jsii.String("imageUrl"),
//   					jsonBody: jsii.String("jsonBody"),
//   					mediaUrl: jsii.String("mediaUrl"),
//   					rawContent: jsii.String("rawContent"),
//   					silentPush: jsii.Boolean(false),
//   					timeToLive: jsii.Number(123),
//   					title: jsii.String("title"),
//   					url: jsii.String("url"),
//   				},
//   				baiduMessage: &messageProperty{
//   					action: jsii.String("action"),
//   					body: jsii.String("body"),
//   					imageIconUrl: jsii.String("imageIconUrl"),
//   					imageSmallIconUrl: jsii.String("imageSmallIconUrl"),
//   					imageUrl: jsii.String("imageUrl"),
//   					jsonBody: jsii.String("jsonBody"),
//   					mediaUrl: jsii.String("mediaUrl"),
//   					rawContent: jsii.String("rawContent"),
//   					silentPush: jsii.Boolean(false),
//   					timeToLive: jsii.Number(123),
//   					title: jsii.String("title"),
//   					url: jsii.String("url"),
//   				},
//   				customMessage: &campaignCustomMessageProperty{
//   					data: jsii.String("data"),
//   				},
//   				defaultMessage: &messageProperty{
//   					action: jsii.String("action"),
//   					body: jsii.String("body"),
//   					imageIconUrl: jsii.String("imageIconUrl"),
//   					imageSmallIconUrl: jsii.String("imageSmallIconUrl"),
//   					imageUrl: jsii.String("imageUrl"),
//   					jsonBody: jsii.String("jsonBody"),
//   					mediaUrl: jsii.String("mediaUrl"),
//   					rawContent: jsii.String("rawContent"),
//   					silentPush: jsii.Boolean(false),
//   					timeToLive: jsii.Number(123),
//   					title: jsii.String("title"),
//   					url: jsii.String("url"),
//   				},
//   				emailMessage: &campaignEmailMessageProperty{
//   					body: jsii.String("body"),
//   					fromAddress: jsii.String("fromAddress"),
//   					htmlBody: jsii.String("htmlBody"),
//   					title: jsii.String("title"),
//   				},
//   				gcmMessage: &messageProperty{
//   					action: jsii.String("action"),
//   					body: jsii.String("body"),
//   					imageIconUrl: jsii.String("imageIconUrl"),
//   					imageSmallIconUrl: jsii.String("imageSmallIconUrl"),
//   					imageUrl: jsii.String("imageUrl"),
//   					jsonBody: jsii.String("jsonBody"),
//   					mediaUrl: jsii.String("mediaUrl"),
//   					rawContent: jsii.String("rawContent"),
//   					silentPush: jsii.Boolean(false),
//   					timeToLive: jsii.Number(123),
//   					title: jsii.String("title"),
//   					url: jsii.String("url"),
//   				},
//   				inAppMessage: &campaignInAppMessageProperty{
//   					content: []interface{}{
//   						&inAppMessageContentProperty{
//   							backgroundColor: jsii.String("backgroundColor"),
//   							bodyConfig: &inAppMessageBodyConfigProperty{
//   								alignment: jsii.String("alignment"),
//   								body: jsii.String("body"),
//   								textColor: jsii.String("textColor"),
//   							},
//   							headerConfig: &inAppMessageHeaderConfigProperty{
//   								alignment: jsii.String("alignment"),
//   								header: jsii.String("header"),
//   								textColor: jsii.String("textColor"),
//   							},
//   							imageUrl: jsii.String("imageUrl"),
//   							primaryBtn: &inAppMessageButtonProperty{
//   								android: &overrideButtonConfigurationProperty{
//   									buttonAction: jsii.String("buttonAction"),
//   									link: jsii.String("link"),
//   								},
//   								defaultConfig: &defaultButtonConfigurationProperty{
//   									backgroundColor: jsii.String("backgroundColor"),
//   									borderRadius: jsii.Number(123),
//   									buttonAction: jsii.String("buttonAction"),
//   									link: jsii.String("link"),
//   									text: jsii.String("text"),
//   									textColor: jsii.String("textColor"),
//   								},
//   								ios: &overrideButtonConfigurationProperty{
//   									buttonAction: jsii.String("buttonAction"),
//   									link: jsii.String("link"),
//   								},
//   								web: &overrideButtonConfigurationProperty{
//   									buttonAction: jsii.String("buttonAction"),
//   									link: jsii.String("link"),
//   								},
//   							},
//   							secondaryBtn: &inAppMessageButtonProperty{
//   								android: &overrideButtonConfigurationProperty{
//   									buttonAction: jsii.String("buttonAction"),
//   									link: jsii.String("link"),
//   								},
//   								defaultConfig: &defaultButtonConfigurationProperty{
//   									backgroundColor: jsii.String("backgroundColor"),
//   									borderRadius: jsii.Number(123),
//   									buttonAction: jsii.String("buttonAction"),
//   									link: jsii.String("link"),
//   									text: jsii.String("text"),
//   									textColor: jsii.String("textColor"),
//   								},
//   								ios: &overrideButtonConfigurationProperty{
//   									buttonAction: jsii.String("buttonAction"),
//   									link: jsii.String("link"),
//   								},
//   								web: &overrideButtonConfigurationProperty{
//   									buttonAction: jsii.String("buttonAction"),
//   									link: jsii.String("link"),
//   								},
//   							},
//   						},
//   					},
//   					customConfig: customConfig,
//   					layout: jsii.String("layout"),
//   				},
//   				smsMessage: &campaignSmsMessageProperty{
//   					body: jsii.String("body"),
//   					entityId: jsii.String("entityId"),
//   					messageType: jsii.String("messageType"),
//   					originationNumber: jsii.String("originationNumber"),
//   					senderId: jsii.String("senderId"),
//   					templateId: jsii.String("templateId"),
//   				},
//   			},
//   			schedule: &scheduleProperty{
//   				endTime: jsii.String("endTime"),
//   				eventFilter: &campaignEventFilterProperty{
//   					dimensions: &eventDimensionsProperty{
//   						attributes: attributes,
//   						eventType: &setDimensionProperty{
//   							dimensionType: jsii.String("dimensionType"),
//   							values: []*string{
//   								jsii.String("values"),
//   							},
//   						},
//   						metrics: metrics,
//   					},
//   					filterType: jsii.String("filterType"),
//   				},
//   				frequency: jsii.String("frequency"),
//   				isLocalTime: jsii.Boolean(false),
//   				quietTime: &quietTimeProperty{
//   					end: jsii.String("end"),
//   					start: jsii.String("start"),
//   				},
//   				startTime: jsii.String("startTime"),
//   				timeZone: jsii.String("timeZone"),
//   			},
//   			sizePercent: jsii.Number(123),
//   			templateConfiguration: &templateConfigurationProperty{
//   				emailTemplate: &templateProperty{
//   					name: jsii.String("name"),
//   					version: jsii.String("version"),
//   				},
//   				pushTemplate: &templateProperty{
//   					name: jsii.String("name"),
//   					version: jsii.String("version"),
//   				},
//   				smsTemplate: &templateProperty{
//   					name: jsii.String("name"),
//   					version: jsii.String("version"),
//   				},
//   				voiceTemplate: &templateProperty{
//   					name: jsii.String("name"),
//   					version: jsii.String("version"),
//   				},
//   			},
//   			treatmentDescription: jsii.String("treatmentDescription"),
//   			treatmentName: jsii.String("treatmentName"),
//   		},
//   	},
//   	campaignHook: &campaignHookProperty{
//   		lambdaFunctionName: jsii.String("lambdaFunctionName"),
//   		mode: jsii.String("mode"),
//   		webUrl: jsii.String("webUrl"),
//   	},
//   	customDeliveryConfiguration: &customDeliveryConfigurationProperty{
//   		deliveryUri: jsii.String("deliveryUri"),
//   		endpointTypes: []*string{
//   			jsii.String("endpointTypes"),
//   		},
//   	},
//   	description: jsii.String("description"),
//   	holdoutPercent: jsii.Number(123),
//   	isPaused: jsii.Boolean(false),
//   	limits: &limitsProperty{
//   		daily: jsii.Number(123),
//   		maximumDuration: jsii.Number(123),
//   		messagesPerSecond: jsii.Number(123),
//   		session: jsii.Number(123),
//   		total: jsii.Number(123),
//   	},
//   	messageConfiguration: &messageConfigurationProperty{
//   		admMessage: &messageProperty{
//   			action: jsii.String("action"),
//   			body: jsii.String("body"),
//   			imageIconUrl: jsii.String("imageIconUrl"),
//   			imageSmallIconUrl: jsii.String("imageSmallIconUrl"),
//   			imageUrl: jsii.String("imageUrl"),
//   			jsonBody: jsii.String("jsonBody"),
//   			mediaUrl: jsii.String("mediaUrl"),
//   			rawContent: jsii.String("rawContent"),
//   			silentPush: jsii.Boolean(false),
//   			timeToLive: jsii.Number(123),
//   			title: jsii.String("title"),
//   			url: jsii.String("url"),
//   		},
//   		apnsMessage: &messageProperty{
//   			action: jsii.String("action"),
//   			body: jsii.String("body"),
//   			imageIconUrl: jsii.String("imageIconUrl"),
//   			imageSmallIconUrl: jsii.String("imageSmallIconUrl"),
//   			imageUrl: jsii.String("imageUrl"),
//   			jsonBody: jsii.String("jsonBody"),
//   			mediaUrl: jsii.String("mediaUrl"),
//   			rawContent: jsii.String("rawContent"),
//   			silentPush: jsii.Boolean(false),
//   			timeToLive: jsii.Number(123),
//   			title: jsii.String("title"),
//   			url: jsii.String("url"),
//   		},
//   		baiduMessage: &messageProperty{
//   			action: jsii.String("action"),
//   			body: jsii.String("body"),
//   			imageIconUrl: jsii.String("imageIconUrl"),
//   			imageSmallIconUrl: jsii.String("imageSmallIconUrl"),
//   			imageUrl: jsii.String("imageUrl"),
//   			jsonBody: jsii.String("jsonBody"),
//   			mediaUrl: jsii.String("mediaUrl"),
//   			rawContent: jsii.String("rawContent"),
//   			silentPush: jsii.Boolean(false),
//   			timeToLive: jsii.Number(123),
//   			title: jsii.String("title"),
//   			url: jsii.String("url"),
//   		},
//   		customMessage: &campaignCustomMessageProperty{
//   			data: jsii.String("data"),
//   		},
//   		defaultMessage: &messageProperty{
//   			action: jsii.String("action"),
//   			body: jsii.String("body"),
//   			imageIconUrl: jsii.String("imageIconUrl"),
//   			imageSmallIconUrl: jsii.String("imageSmallIconUrl"),
//   			imageUrl: jsii.String("imageUrl"),
//   			jsonBody: jsii.String("jsonBody"),
//   			mediaUrl: jsii.String("mediaUrl"),
//   			rawContent: jsii.String("rawContent"),
//   			silentPush: jsii.Boolean(false),
//   			timeToLive: jsii.Number(123),
//   			title: jsii.String("title"),
//   			url: jsii.String("url"),
//   		},
//   		emailMessage: &campaignEmailMessageProperty{
//   			body: jsii.String("body"),
//   			fromAddress: jsii.String("fromAddress"),
//   			htmlBody: jsii.String("htmlBody"),
//   			title: jsii.String("title"),
//   		},
//   		gcmMessage: &messageProperty{
//   			action: jsii.String("action"),
//   			body: jsii.String("body"),
//   			imageIconUrl: jsii.String("imageIconUrl"),
//   			imageSmallIconUrl: jsii.String("imageSmallIconUrl"),
//   			imageUrl: jsii.String("imageUrl"),
//   			jsonBody: jsii.String("jsonBody"),
//   			mediaUrl: jsii.String("mediaUrl"),
//   			rawContent: jsii.String("rawContent"),
//   			silentPush: jsii.Boolean(false),
//   			timeToLive: jsii.Number(123),
//   			title: jsii.String("title"),
//   			url: jsii.String("url"),
//   		},
//   		inAppMessage: &campaignInAppMessageProperty{
//   			content: []interface{}{
//   				&inAppMessageContentProperty{
//   					backgroundColor: jsii.String("backgroundColor"),
//   					bodyConfig: &inAppMessageBodyConfigProperty{
//   						alignment: jsii.String("alignment"),
//   						body: jsii.String("body"),
//   						textColor: jsii.String("textColor"),
//   					},
//   					headerConfig: &inAppMessageHeaderConfigProperty{
//   						alignment: jsii.String("alignment"),
//   						header: jsii.String("header"),
//   						textColor: jsii.String("textColor"),
//   					},
//   					imageUrl: jsii.String("imageUrl"),
//   					primaryBtn: &inAppMessageButtonProperty{
//   						android: &overrideButtonConfigurationProperty{
//   							buttonAction: jsii.String("buttonAction"),
//   							link: jsii.String("link"),
//   						},
//   						defaultConfig: &defaultButtonConfigurationProperty{
//   							backgroundColor: jsii.String("backgroundColor"),
//   							borderRadius: jsii.Number(123),
//   							buttonAction: jsii.String("buttonAction"),
//   							link: jsii.String("link"),
//   							text: jsii.String("text"),
//   							textColor: jsii.String("textColor"),
//   						},
//   						ios: &overrideButtonConfigurationProperty{
//   							buttonAction: jsii.String("buttonAction"),
//   							link: jsii.String("link"),
//   						},
//   						web: &overrideButtonConfigurationProperty{
//   							buttonAction: jsii.String("buttonAction"),
//   							link: jsii.String("link"),
//   						},
//   					},
//   					secondaryBtn: &inAppMessageButtonProperty{
//   						android: &overrideButtonConfigurationProperty{
//   							buttonAction: jsii.String("buttonAction"),
//   							link: jsii.String("link"),
//   						},
//   						defaultConfig: &defaultButtonConfigurationProperty{
//   							backgroundColor: jsii.String("backgroundColor"),
//   							borderRadius: jsii.Number(123),
//   							buttonAction: jsii.String("buttonAction"),
//   							link: jsii.String("link"),
//   							text: jsii.String("text"),
//   							textColor: jsii.String("textColor"),
//   						},
//   						ios: &overrideButtonConfigurationProperty{
//   							buttonAction: jsii.String("buttonAction"),
//   							link: jsii.String("link"),
//   						},
//   						web: &overrideButtonConfigurationProperty{
//   							buttonAction: jsii.String("buttonAction"),
//   							link: jsii.String("link"),
//   						},
//   					},
//   				},
//   			},
//   			customConfig: customConfig,
//   			layout: jsii.String("layout"),
//   		},
//   		smsMessage: &campaignSmsMessageProperty{
//   			body: jsii.String("body"),
//   			entityId: jsii.String("entityId"),
//   			messageType: jsii.String("messageType"),
//   			originationNumber: jsii.String("originationNumber"),
//   			senderId: jsii.String("senderId"),
//   			templateId: jsii.String("templateId"),
//   		},
//   	},
//   	priority: jsii.Number(123),
//   	segmentVersion: jsii.Number(123),
//   	tags: tags,
//   	templateConfiguration: &templateConfigurationProperty{
//   		emailTemplate: &templateProperty{
//   			name: jsii.String("name"),
//   			version: jsii.String("version"),
//   		},
//   		pushTemplate: &templateProperty{
//   			name: jsii.String("name"),
//   			version: jsii.String("version"),
//   		},
//   		smsTemplate: &templateProperty{
//   			name: jsii.String("name"),
//   			version: jsii.String("version"),
//   		},
//   		voiceTemplate: &templateProperty{
//   			name: jsii.String("name"),
//   			version: jsii.String("version"),
//   		},
//   	},
//   	treatmentDescription: jsii.String("treatmentDescription"),
//   	treatmentName: jsii.String("treatmentName"),
//   })
//
type CfnCampaign interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// An array of requests that defines additional treatments for the campaign, in addition to the default treatment for the campaign.
	AdditionalTreatments() interface{}
	SetAdditionalTreatments(val interface{})
	// The unique identifier for the Amazon Pinpoint application that the campaign is associated with.
	ApplicationId() *string
	SetApplicationId(val *string)
	// The Amazon Resource Name (ARN) of the campaign.
	AttrArn() *string
	// The unique identifier for the campaign.
	AttrCampaignId() *string
	// Specifies the Lambda function to use as a code hook for a campaign.
	CampaignHook() interface{}
	SetCampaignHook(val interface{})
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The delivery configuration settings for sending the treatment through a custom channel.
	//
	// This object is required if the `MessageConfiguration` object for the treatment specifies a `CustomMessage` object.
	CustomDeliveryConfiguration() interface{}
	SetCustomDeliveryConfiguration(val interface{})
	// A custom description of the campaign.
	Description() *string
	SetDescription(val *string)
	// The allocated percentage of users (segment members) who shouldn't receive messages from the campaign.
	HoldoutPercent() *float64
	SetHoldoutPercent(val *float64)
	// Specifies whether to pause the campaign.
	//
	// A paused campaign doesn't run unless you resume it by changing this value to `false` . If you restart a campaign, the campaign restarts from the beginning and not at the point you paused it.
	IsPaused() interface{}
	SetIsPaused(val interface{})
	// The messaging limits for the campaign.
	Limits() interface{}
	SetLimits(val interface{})
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
	// The message configuration settings for the campaign.
	MessageConfiguration() interface{}
	SetMessageConfiguration(val interface{})
	// The name of the campaign.
	Name() *string
	SetName(val *string)
	// The tree node.
	Node() constructs.Node
	// An integer between 1 and 5, inclusive, that represents the priority of the in-app message campaign, where 1 is the highest priority and 5 is the lowest.
	//
	// If there are multiple messages scheduled to be displayed at the same time, the priority determines the order in which those messages are displayed.
	Priority() *float64
	SetPriority(val *float64)
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The schedule settings for the campaign.
	Schedule() interface{}
	SetSchedule(val interface{})
	// The unique identifier for the segment to associate with the campaign.
	SegmentId() *string
	SetSegmentId(val *string)
	// The version of the segment to associate with the campaign.
	SegmentVersion() *float64
	SetSegmentVersion(val *float64)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags() awscdk.TagManager
	// The message template to use for the treatment.
	TemplateConfiguration() interface{}
	SetTemplateConfiguration(val interface{})
	// A custom description of the default treatment for the campaign.
	TreatmentDescription() *string
	SetTreatmentDescription(val *string)
	// A custom name of the default treatment for the campaign, if the campaign has multiple treatments.
	//
	// A *treatment* is a variation of a campaign that's used for A/B testing.
	TreatmentName() *string
	SetTreatmentName(val *string)
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

// The jsii proxy struct for CfnCampaign
type jsiiProxy_CfnCampaign struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnCampaign) AdditionalTreatments() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"additionalTreatments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCampaign) ApplicationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCampaign) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCampaign) AttrCampaignId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrCampaignId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCampaign) CampaignHook() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"campaignHook",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCampaign) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCampaign) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCampaign) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCampaign) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCampaign) CustomDeliveryConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customDeliveryConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCampaign) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCampaign) HoldoutPercent() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"holdoutPercent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCampaign) IsPaused() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isPaused",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCampaign) Limits() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"limits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCampaign) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCampaign) MessageConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"messageConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCampaign) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCampaign) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCampaign) Priority() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCampaign) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCampaign) Schedule() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"schedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCampaign) SegmentId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"segmentId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCampaign) SegmentVersion() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"segmentVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCampaign) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCampaign) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCampaign) TemplateConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"templateConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCampaign) TreatmentDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"treatmentDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCampaign) TreatmentName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"treatmentName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCampaign) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCampaign) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}


// Create a new `AWS::Pinpoint::Campaign`.
func NewCfnCampaign(scope constructs.Construct, id *string, props *CfnCampaignProps) CfnCampaign {
	_init_.Initialize()

	if err := validateNewCfnCampaignParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnCampaign{}

	_jsii_.Create(
		"aws-cdk-lib.aws_pinpoint.CfnCampaign",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Pinpoint::Campaign`.
func NewCfnCampaign_Override(c CfnCampaign, scope constructs.Construct, id *string, props *CfnCampaignProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_pinpoint.CfnCampaign",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnCampaign)SetAdditionalTreatments(val interface{}) {
	if err := j.validateSetAdditionalTreatmentsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"additionalTreatments",
		val,
	)
}

func (j *jsiiProxy_CfnCampaign)SetApplicationId(val *string) {
	if err := j.validateSetApplicationIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"applicationId",
		val,
	)
}

func (j *jsiiProxy_CfnCampaign)SetCampaignHook(val interface{}) {
	if err := j.validateSetCampaignHookParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"campaignHook",
		val,
	)
}

func (j *jsiiProxy_CfnCampaign)SetCustomDeliveryConfiguration(val interface{}) {
	if err := j.validateSetCustomDeliveryConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customDeliveryConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnCampaign)SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnCampaign)SetHoldoutPercent(val *float64) {
	_jsii_.Set(
		j,
		"holdoutPercent",
		val,
	)
}

func (j *jsiiProxy_CfnCampaign)SetIsPaused(val interface{}) {
	if err := j.validateSetIsPausedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isPaused",
		val,
	)
}

func (j *jsiiProxy_CfnCampaign)SetLimits(val interface{}) {
	if err := j.validateSetLimitsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"limits",
		val,
	)
}

func (j *jsiiProxy_CfnCampaign)SetMessageConfiguration(val interface{}) {
	if err := j.validateSetMessageConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"messageConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnCampaign)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnCampaign)SetPriority(val *float64) {
	_jsii_.Set(
		j,
		"priority",
		val,
	)
}

func (j *jsiiProxy_CfnCampaign)SetSchedule(val interface{}) {
	if err := j.validateSetScheduleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"schedule",
		val,
	)
}

func (j *jsiiProxy_CfnCampaign)SetSegmentId(val *string) {
	if err := j.validateSetSegmentIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"segmentId",
		val,
	)
}

func (j *jsiiProxy_CfnCampaign)SetSegmentVersion(val *float64) {
	_jsii_.Set(
		j,
		"segmentVersion",
		val,
	)
}

func (j *jsiiProxy_CfnCampaign)SetTemplateConfiguration(val interface{}) {
	if err := j.validateSetTemplateConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"templateConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnCampaign)SetTreatmentDescription(val *string) {
	_jsii_.Set(
		j,
		"treatmentDescription",
		val,
	)
}

func (j *jsiiProxy_CfnCampaign)SetTreatmentName(val *string) {
	_jsii_.Set(
		j,
		"treatmentName",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnCampaign_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnCampaign_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_pinpoint.CfnCampaign",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnCampaign_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateCfnCampaign_IsCfnResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_pinpoint.CfnCampaign",
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
func CfnCampaign_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnCampaign_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_pinpoint.CfnCampaign",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnCampaign_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_pinpoint.CfnCampaign",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnCampaign) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnCampaign) AddDependency(target awscdk.CfnResource) {
	if err := c.validateAddDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnCampaign) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnCampaign) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnCampaign) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnCampaign) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnCampaign) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnCampaign) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnCampaign) GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference {
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

func (c *jsiiProxy_CfnCampaign) GetMetadata(key *string) interface{} {
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

func (c *jsiiProxy_CfnCampaign) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnCampaign) ObtainDependencies() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		c,
		"obtainDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCampaign) ObtainResourceDependencies() *[]awscdk.CfnResource {
	var returns *[]awscdk.CfnResource

	_jsii_.Invoke(
		c,
		"obtainResourceDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCampaign) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnCampaign) RemoveDependency(target awscdk.CfnResource) {
	if err := c.validateRemoveDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"removeDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnCampaign) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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

func (c *jsiiProxy_CfnCampaign) ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource) {
	if err := c.validateReplaceDependencyParameters(target, newTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"replaceDependency",
		[]interface{}{target, newTarget},
	)
}

func (c *jsiiProxy_CfnCampaign) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCampaign) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCampaign) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

