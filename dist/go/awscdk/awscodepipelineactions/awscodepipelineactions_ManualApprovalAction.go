package awscodepipelineactions

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/awscodepipeline"
	"github.com/aws/aws-cdk-go/awscdk/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/awssns"
	"github.com/aws/constructs-go/constructs/v10"
)

// Manual approval action.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   pipeline := codepipeline.NewPipeline(this, jsii.String("MyPipeline"))
//   approveStage := pipeline.addStage(&stageOptions{
//   	stageName: jsii.String("Approve"),
//   })
//   manualApprovalAction := codepipeline_actions.NewManualApprovalAction(&manualApprovalActionProps{
//   	actionName: jsii.String("Approve"),
//   })
//   approveStage.addAction(manualApprovalAction)
//
//   role := iam.role.fromRoleArn(this, jsii.String("Admin"), awscdk.Arn.format(&arnComponents{
//   	service: jsii.String("iam"),
//   	resource: jsii.String("role"),
//   	resourceName: jsii.String("Admin"),
//   }, this))
//   manualApprovalAction.grantManualApproval(role)
//
type ManualApprovalAction interface {
	Action
	// The simple properties of the Action, like its Owner, name, etc.
	//
	// Note that this accessor will be called before the {@link bind} callback.
	ActionProperties() *awscodepipeline.ActionProperties
	NotificationTopic() awssns.ITopic
	// This is a renamed version of the {@link IAction.actionProperties} property.
	ProvidedActionProperties() *awscodepipeline.ActionProperties
	// The callback invoked when this Action is added to a Pipeline.
	Bind(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig
	// This is a renamed version of the {@link IAction.bind} method.
	Bound(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig
	// grant the provided principal the permissions to approve or reject this manual approval action.
	//
	// For more info see:
	// https://docs.aws.amazon.com/codepipeline/latest/userguide/approvals-iam-permissions.html
	GrantManualApproval(grantable awsiam.IGrantable)
	// Creates an Event that will be triggered whenever the state of this Action changes.
	OnStateChange(name *string, target awsevents.IRuleTarget, options *awsevents.RuleProps) awsevents.Rule
	VariableExpression(variableName *string) *string
}

// The jsii proxy struct for ManualApprovalAction
type jsiiProxy_ManualApprovalAction struct {
	jsiiProxy_Action
}

func (j *jsiiProxy_ManualApprovalAction) ActionProperties() *awscodepipeline.ActionProperties {
	var returns *awscodepipeline.ActionProperties
	_jsii_.Get(
		j,
		"actionProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManualApprovalAction) NotificationTopic() awssns.ITopic {
	var returns awssns.ITopic
	_jsii_.Get(
		j,
		"notificationTopic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManualApprovalAction) ProvidedActionProperties() *awscodepipeline.ActionProperties {
	var returns *awscodepipeline.ActionProperties
	_jsii_.Get(
		j,
		"providedActionProperties",
		&returns,
	)
	return returns
}


func NewManualApprovalAction(props *ManualApprovalActionProps) ManualApprovalAction {
	_init_.Initialize()

	if err := validateNewManualApprovalActionParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_ManualApprovalAction{}

	_jsii_.Create(
		"aws-cdk-lib.aws_codepipeline_actions.ManualApprovalAction",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewManualApprovalAction_Override(m ManualApprovalAction, props *ManualApprovalActionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_codepipeline_actions.ManualApprovalAction",
		[]interface{}{props},
		m,
	)
}

func (m *jsiiProxy_ManualApprovalAction) Bind(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig {
	if err := m.validateBindParameters(scope, stage, options); err != nil {
		panic(err)
	}
	var returns *awscodepipeline.ActionConfig

	_jsii_.Invoke(
		m,
		"bind",
		[]interface{}{scope, stage, options},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManualApprovalAction) Bound(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) *awscodepipeline.ActionConfig {
	if err := m.validateBoundParameters(scope, stage, options); err != nil {
		panic(err)
	}
	var returns *awscodepipeline.ActionConfig

	_jsii_.Invoke(
		m,
		"bound",
		[]interface{}{scope, stage, options},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManualApprovalAction) GrantManualApproval(grantable awsiam.IGrantable) {
	if err := m.validateGrantManualApprovalParameters(grantable); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"grantManualApproval",
		[]interface{}{grantable},
	)
}

func (m *jsiiProxy_ManualApprovalAction) OnStateChange(name *string, target awsevents.IRuleTarget, options *awsevents.RuleProps) awsevents.Rule {
	if err := m.validateOnStateChangeParameters(name, options); err != nil {
		panic(err)
	}
	var returns awsevents.Rule

	_jsii_.Invoke(
		m,
		"onStateChange",
		[]interface{}{name, target, options},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManualApprovalAction) VariableExpression(variableName *string) *string {
	if err := m.validateVariableExpressionParameters(variableName); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		m,
		"variableExpression",
		[]interface{}{variableName},
		&returns,
	)

	return returns
}

