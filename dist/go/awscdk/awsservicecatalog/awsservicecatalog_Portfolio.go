package awsservicecatalog

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/awsservicecatalog/internal"
	"github.com/aws/aws-cdk-go/awscdk/awssns"
	"github.com/aws/constructs-go/constructs/v10"
)

// A Service Catalog portfolio.
//
// Example:
//   servicecatalog.NewPortfolio(this, jsii.String("Portfolio"), &portfolioProps{
//   	displayName: jsii.String("MyPortfolio"),
//   	providerName: jsii.String("MyTeam"),
//   })
//
type Portfolio interface {
	awscdk.Resource
	IPortfolio
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	Env() *awscdk.ResourceEnvironment
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
	// The ARN of the portfolio.
	PortfolioArn() *string
	// The ID of the portfolio.
	PortfolioId() *string
	// The stack in which this resource is defined.
	Stack() awscdk.Stack
	// Associate portfolio with the given product.
	AddProduct(product IProduct)
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
	// Associate Tag Options.
	//
	// A TagOption is a key-value pair managed in AWS Service Catalog.
	// It is not an AWS tag, but serves as a template for creating an AWS tag based on the TagOption.
	AssociateTagOptions(tagOptions TagOptions)
	// Set provisioning rules for the product.
	ConstrainCloudFormationParameters(product IProduct, options *CloudFormationRuleConstraintOptions)
	// Add a Resource Update Constraint.
	ConstrainTagUpdates(product IProduct, options *TagUpdateConstraintOptions)
	// Configure deployment options using AWS Cloudformation StackSets.
	DeployWithStackSets(product IProduct, options *StackSetsConstraintOptions)
	GeneratePhysicalName() *string
	// Create a unique id based off the L1 CfnPortfolio or the arn of an imported portfolio.
	GenerateUniqueHash(value *string) *string
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
	// Associate portfolio with an IAM Group.
	GiveAccessToGroup(group awsiam.IGroup)
	// Associate portfolio with an IAM Role.
	GiveAccessToRole(role awsiam.IRole)
	// Associate portfolio with an IAM User.
	GiveAccessToUser(user awsiam.IUser)
	// Add notifications for supplied topics on the provisioned product.
	NotifyOnStackEvents(product IProduct, topic awssns.ITopic, options *CommonConstraintOptions)
	// Force users to assume a certain role when launching a product.
	//
	// This sets the launch role using the role arn which is tied to the account this role exists in.
	// This is useful if you will be provisioning products from the account where this role exists.
	// If you intend to share the portfolio across accounts, use a local launch role.
	SetLaunchRole(product IProduct, launchRole awsiam.IRole, options *CommonConstraintOptions)
	// Force users to assume a certain role when launching a product.
	//
	// The role name will be referenced by in the local account and must be set explicitly.
	// This is useful when sharing the portfolio with multiple accounts.
	SetLocalLaunchRole(product IProduct, launchRole awsiam.IRole, options *CommonConstraintOptions)
	// Force users to assume a certain role when launching a product.
	//
	// The role will be referenced by name in the local account instead of a static role arn.
	// A role with this name will automatically be created and assumable by Service Catalog in this account.
	// This is useful when sharing the portfolio with multiple accounts.
	SetLocalLaunchRoleName(product IProduct, launchRoleName *string, options *CommonConstraintOptions) awsiam.IRole
	// Initiate a portfolio share with another account.
	ShareWithAccount(accountId *string, options *PortfolioShareOptions)
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for Portfolio
type jsiiProxy_Portfolio struct {
	internal.Type__awscdkResource
	jsiiProxy_IPortfolio
}

func (j *jsiiProxy_Portfolio) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Portfolio) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Portfolio) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Portfolio) PortfolioArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"portfolioArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Portfolio) PortfolioId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"portfolioId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Portfolio) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


func NewPortfolio(scope constructs.Construct, id *string, props *PortfolioProps) Portfolio {
	_init_.Initialize()

	if err := validateNewPortfolioParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_Portfolio{}

	_jsii_.Create(
		"aws-cdk-lib.aws_servicecatalog.Portfolio",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewPortfolio_Override(p Portfolio, scope constructs.Construct, id *string, props *PortfolioProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_servicecatalog.Portfolio",
		[]interface{}{scope, id, props},
		p,
	)
}

// Creates a Portfolio construct that represents an external portfolio.
func Portfolio_FromPortfolioArn(scope constructs.Construct, id *string, portfolioArn *string) IPortfolio {
	_init_.Initialize()

	if err := validatePortfolio_FromPortfolioArnParameters(scope, id, portfolioArn); err != nil {
		panic(err)
	}
	var returns IPortfolio

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_servicecatalog.Portfolio",
		"fromPortfolioArn",
		[]interface{}{scope, id, portfolioArn},
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
func Portfolio_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePortfolio_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_servicecatalog.Portfolio",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
func Portfolio_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validatePortfolio_IsOwnedResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_servicecatalog.Portfolio",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func Portfolio_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validatePortfolio_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_servicecatalog.Portfolio",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Portfolio) AddProduct(product IProduct) {
	if err := p.validateAddProductParameters(product); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addProduct",
		[]interface{}{product},
	)
}

func (p *jsiiProxy_Portfolio) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := p.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (p *jsiiProxy_Portfolio) AssociateTagOptions(tagOptions TagOptions) {
	if err := p.validateAssociateTagOptionsParameters(tagOptions); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"associateTagOptions",
		[]interface{}{tagOptions},
	)
}

func (p *jsiiProxy_Portfolio) ConstrainCloudFormationParameters(product IProduct, options *CloudFormationRuleConstraintOptions) {
	if err := p.validateConstrainCloudFormationParametersParameters(product, options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"constrainCloudFormationParameters",
		[]interface{}{product, options},
	)
}

func (p *jsiiProxy_Portfolio) ConstrainTagUpdates(product IProduct, options *TagUpdateConstraintOptions) {
	if err := p.validateConstrainTagUpdatesParameters(product, options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"constrainTagUpdates",
		[]interface{}{product, options},
	)
}

func (p *jsiiProxy_Portfolio) DeployWithStackSets(product IProduct, options *StackSetsConstraintOptions) {
	if err := p.validateDeployWithStackSetsParameters(product, options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"deployWithStackSets",
		[]interface{}{product, options},
	)
}

func (p *jsiiProxy_Portfolio) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Portfolio) GenerateUniqueHash(value *string) *string {
	if err := p.validateGenerateUniqueHashParameters(value); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		p,
		"generateUniqueHash",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Portfolio) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	if err := p.validateGetResourceArnAttributeParameters(arnAttr, arnComponents); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		p,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Portfolio) GetResourceNameAttribute(nameAttr *string) *string {
	if err := p.validateGetResourceNameAttributeParameters(nameAttr); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		p,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Portfolio) GiveAccessToGroup(group awsiam.IGroup) {
	if err := p.validateGiveAccessToGroupParameters(group); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"giveAccessToGroup",
		[]interface{}{group},
	)
}

func (p *jsiiProxy_Portfolio) GiveAccessToRole(role awsiam.IRole) {
	if err := p.validateGiveAccessToRoleParameters(role); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"giveAccessToRole",
		[]interface{}{role},
	)
}

func (p *jsiiProxy_Portfolio) GiveAccessToUser(user awsiam.IUser) {
	if err := p.validateGiveAccessToUserParameters(user); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"giveAccessToUser",
		[]interface{}{user},
	)
}

func (p *jsiiProxy_Portfolio) NotifyOnStackEvents(product IProduct, topic awssns.ITopic, options *CommonConstraintOptions) {
	if err := p.validateNotifyOnStackEventsParameters(product, topic, options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"notifyOnStackEvents",
		[]interface{}{product, topic, options},
	)
}

func (p *jsiiProxy_Portfolio) SetLaunchRole(product IProduct, launchRole awsiam.IRole, options *CommonConstraintOptions) {
	if err := p.validateSetLaunchRoleParameters(product, launchRole, options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"setLaunchRole",
		[]interface{}{product, launchRole, options},
	)
}

func (p *jsiiProxy_Portfolio) SetLocalLaunchRole(product IProduct, launchRole awsiam.IRole, options *CommonConstraintOptions) {
	if err := p.validateSetLocalLaunchRoleParameters(product, launchRole, options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"setLocalLaunchRole",
		[]interface{}{product, launchRole, options},
	)
}

func (p *jsiiProxy_Portfolio) SetLocalLaunchRoleName(product IProduct, launchRoleName *string, options *CommonConstraintOptions) awsiam.IRole {
	if err := p.validateSetLocalLaunchRoleNameParameters(product, launchRoleName, options); err != nil {
		panic(err)
	}
	var returns awsiam.IRole

	_jsii_.Invoke(
		p,
		"setLocalLaunchRoleName",
		[]interface{}{product, launchRoleName, options},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Portfolio) ShareWithAccount(accountId *string, options *PortfolioShareOptions) {
	if err := p.validateShareWithAccountParameters(accountId, options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"shareWithAccount",
		[]interface{}{accountId, options},
	)
}

func (p *jsiiProxy_Portfolio) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

