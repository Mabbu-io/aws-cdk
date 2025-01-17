package awsconnect

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnRule`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var assignContactCategoryActions interface{}
//
//   cfnRuleProps := &cfnRuleProps{
//   	actions: &actionsProperty{
//   		assignContactCategoryActions: []interface{}{
//   			assignContactCategoryActions,
//   		},
//   		eventBridgeActions: []interface{}{
//   			&eventBridgeActionProperty{
//   				name: jsii.String("name"),
//   			},
//   		},
//   		sendNotificationActions: []interface{}{
//   			&sendNotificationActionProperty{
//   				content: jsii.String("content"),
//   				contentType: jsii.String("contentType"),
//   				deliveryMethod: jsii.String("deliveryMethod"),
//   				recipient: &notificationRecipientTypeProperty{
//   					userArns: []*string{
//   						jsii.String("userArns"),
//   					},
//   					userTags: map[string]*string{
//   						"userTagsKey": jsii.String("userTags"),
//   					},
//   				},
//
//   				// the properties below are optional
//   				subject: jsii.String("subject"),
//   			},
//   		},
//   		taskActions: []interface{}{
//   			&taskActionProperty{
//   				contactFlowArn: jsii.String("contactFlowArn"),
//   				name: jsii.String("name"),
//
//   				// the properties below are optional
//   				description: jsii.String("description"),
//   				references: map[string]interface{}{
//   					"referencesKey": &ReferenceProperty{
//   						"type": jsii.String("type"),
//   						"value": jsii.String("value"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   	function: jsii.String("function"),
//   	instanceArn: jsii.String("instanceArn"),
//   	name: jsii.String("name"),
//   	publishStatus: jsii.String("publishStatus"),
//   	triggerEventSource: &ruleTriggerEventSourceProperty{
//   		eventSourceName: jsii.String("eventSourceName"),
//
//   		// the properties below are optional
//   		integrationAssociationArn: jsii.String("integrationAssociationArn"),
//   	},
//
//   	// the properties below are optional
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnRuleProps struct {
	// `AWS::Connect::Rule.Actions`.
	Actions interface{} `field:"required" json:"actions" yaml:"actions"`
	// `AWS::Connect::Rule.Function`.
	Function *string `field:"required" json:"function" yaml:"function"`
	// `AWS::Connect::Rule.InstanceArn`.
	InstanceArn *string `field:"required" json:"instanceArn" yaml:"instanceArn"`
	// `AWS::Connect::Rule.Name`.
	Name *string `field:"required" json:"name" yaml:"name"`
	// `AWS::Connect::Rule.PublishStatus`.
	PublishStatus *string `field:"required" json:"publishStatus" yaml:"publishStatus"`
	// `AWS::Connect::Rule.TriggerEventSource`.
	TriggerEventSource interface{} `field:"required" json:"triggerEventSource" yaml:"triggerEventSource"`
	// `AWS::Connect::Rule.Tags`.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

