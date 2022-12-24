package awsconfig

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsevents"
	"github.com/aws/constructs-go/constructs/v10"
)

// Checks whether your CloudFormation stacks are sending event notifications to a SNS topic.
//
// Optionally checks whether specified SNS topics are used.
//
// Example:
//   // topics to which CloudFormation stacks may send event notifications
//   topic1 := sns.NewTopic(this, jsii.String("AllowedTopic1"))
//   topic2 := sns.NewTopic(this, jsii.String("AllowedTopic2"))
//
//   // non-compliant if CloudFormation stack does not send notifications to 'topic1' or 'topic2'
//   // non-compliant if CloudFormation stack does not send notifications to 'topic1' or 'topic2'
//   config.NewCloudFormationStackNotificationCheck(this, jsii.String("NotificationCheck"), &cloudFormationStackNotificationCheckProps{
//   	topics: []iTopic{
//   		topic1,
//   		topic2,
//   	},
//   })
//
// See: https://docs.aws.amazon.com/config/latest/developerguide/cloudformation-stack-notification-check.html
//
type CloudFormationStackNotificationCheck interface {
	ManagedRule
	// The arn of the rule.
	ConfigRuleArn() *string
	// The compliance status of the rule.
	ConfigRuleComplianceType() *string
	// The id of the rule.
	ConfigRuleId() *string
	// The name of the rule.
	ConfigRuleName() *string
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	Env() *awscdk.ResourceEnvironment
	IsCustomWithChanges() *bool
	SetIsCustomWithChanges(val *bool)
	IsManaged() *bool
	SetIsManaged(val *bool)
	// The tree node.
	Node() constructs.Node
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//    cross-environment scenarios.
	PhysicalName() *string
	RuleScope() RuleScope
	SetRuleScope(val RuleScope)
	// The stack in which this resource is defined.
	Stack() awscdk.Stack
	// Apply the given removal policy to this resource.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
	//
	// Normally, this token will resolve to `arnAttr`, but if the resource is
	// referenced across environments, `arnComponents` will be used to synthesize
	// a concrete ARN with the resource's physical name. Make sure to reference
	// `this.physicalName` in `arnComponents`.
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
	//
	// Normally, this token will resolve to `nameAttr`, but if the resource is
	// referenced across environments, it will be resolved to `this.physicalName`,
	// which will be a concrete name.
	GetResourceNameAttribute(nameAttr *string) *string
	// Defines an EventBridge event rule which triggers for rule compliance events.
	OnComplianceChange(id *string, options *awsevents.OnEventOptions) awsevents.Rule
	// Defines an EventBridge event rule which triggers for rule events.
	//
	// Use
	// `rule.addEventPattern(pattern)` to specify a filter.
	OnEvent(id *string, options *awsevents.OnEventOptions) awsevents.Rule
	// Defines an EventBridge event rule which triggers for rule re-evaluation status events.
	OnReEvaluationStatus(id *string, options *awsevents.OnEventOptions) awsevents.Rule
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for CloudFormationStackNotificationCheck
type jsiiProxy_CloudFormationStackNotificationCheck struct {
	jsiiProxy_ManagedRule
}

func (j *jsiiProxy_CloudFormationStackNotificationCheck) ConfigRuleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configRuleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudFormationStackNotificationCheck) ConfigRuleComplianceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configRuleComplianceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudFormationStackNotificationCheck) ConfigRuleId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configRuleId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudFormationStackNotificationCheck) ConfigRuleName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configRuleName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudFormationStackNotificationCheck) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudFormationStackNotificationCheck) IsCustomWithChanges() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isCustomWithChanges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudFormationStackNotificationCheck) IsManaged() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isManaged",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudFormationStackNotificationCheck) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudFormationStackNotificationCheck) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudFormationStackNotificationCheck) RuleScope() RuleScope {
	var returns RuleScope
	_jsii_.Get(
		j,
		"ruleScope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudFormationStackNotificationCheck) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


func NewCloudFormationStackNotificationCheck(scope constructs.Construct, id *string, props *CloudFormationStackNotificationCheckProps) CloudFormationStackNotificationCheck {
	_init_.Initialize()

	if err := validateNewCloudFormationStackNotificationCheckParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CloudFormationStackNotificationCheck{}

	_jsii_.Create(
		"aws-cdk-lib.aws_config.CloudFormationStackNotificationCheck",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewCloudFormationStackNotificationCheck_Override(c CloudFormationStackNotificationCheck, scope constructs.Construct, id *string, props *CloudFormationStackNotificationCheckProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_config.CloudFormationStackNotificationCheck",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CloudFormationStackNotificationCheck)SetIsCustomWithChanges(val *bool) {
	_jsii_.Set(
		j,
		"isCustomWithChanges",
		val,
	)
}

func (j *jsiiProxy_CloudFormationStackNotificationCheck)SetIsManaged(val *bool) {
	_jsii_.Set(
		j,
		"isManaged",
		val,
	)
}

func (j *jsiiProxy_CloudFormationStackNotificationCheck)SetRuleScope(val RuleScope) {
	_jsii_.Set(
		j,
		"ruleScope",
		val,
	)
}

// Imports an existing rule.
func CloudFormationStackNotificationCheck_FromConfigRuleName(scope constructs.Construct, id *string, configRuleName *string) IRule {
	_init_.Initialize()

	if err := validateCloudFormationStackNotificationCheck_FromConfigRuleNameParameters(scope, id, configRuleName); err != nil {
		panic(err)
	}
	var returns IRule

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_config.CloudFormationStackNotificationCheck",
		"fromConfigRuleName",
		[]interface{}{scope, id, configRuleName},
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
func CloudFormationStackNotificationCheck_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCloudFormationStackNotificationCheck_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_config.CloudFormationStackNotificationCheck",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
func CloudFormationStackNotificationCheck_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateCloudFormationStackNotificationCheck_IsOwnedResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_config.CloudFormationStackNotificationCheck",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func CloudFormationStackNotificationCheck_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateCloudFormationStackNotificationCheck_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_config.CloudFormationStackNotificationCheck",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudFormationStackNotificationCheck) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := c.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (c *jsiiProxy_CloudFormationStackNotificationCheck) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudFormationStackNotificationCheck) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	if err := c.validateGetResourceArnAttributeParameters(arnAttr, arnComponents); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudFormationStackNotificationCheck) GetResourceNameAttribute(nameAttr *string) *string {
	if err := c.validateGetResourceNameAttributeParameters(nameAttr); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudFormationStackNotificationCheck) OnComplianceChange(id *string, options *awsevents.OnEventOptions) awsevents.Rule {
	if err := c.validateOnComplianceChangeParameters(id, options); err != nil {
		panic(err)
	}
	var returns awsevents.Rule

	_jsii_.Invoke(
		c,
		"onComplianceChange",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudFormationStackNotificationCheck) OnEvent(id *string, options *awsevents.OnEventOptions) awsevents.Rule {
	if err := c.validateOnEventParameters(id, options); err != nil {
		panic(err)
	}
	var returns awsevents.Rule

	_jsii_.Invoke(
		c,
		"onEvent",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudFormationStackNotificationCheck) OnReEvaluationStatus(id *string, options *awsevents.OnEventOptions) awsevents.Rule {
	if err := c.validateOnReEvaluationStatusParameters(id, options); err != nil {
		panic(err)
	}
	var returns awsevents.Rule

	_jsii_.Invoke(
		c,
		"onReEvaluationStatus",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudFormationStackNotificationCheck) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

