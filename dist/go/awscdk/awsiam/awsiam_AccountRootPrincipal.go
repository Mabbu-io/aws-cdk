package awsiam

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Use the AWS account into which a stack is deployed as the principal entity in a policy.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   bucket := s3.NewBucket(this, jsii.String("MyBucket"))
//   result := bucket.addToResourcePolicy(iam.NewPolicyStatement(&policyStatementProps{
//   	actions: []*string{
//   		jsii.String("s3:GetObject"),
//   	},
//   	resources: []*string{
//   		bucket.arnForObjects(jsii.String("file.txt")),
//   	},
//   	principals: []iPrincipal{
//   		iam.NewAccountRootPrincipal(),
//   	},
//   }))
//
type AccountRootPrincipal interface {
	AccountPrincipal
	// AWS account ID (i.e. '123456789012').
	AccountId() interface{}
	// Amazon Resource Name (ARN) of the principal entity (i.e. arn:aws:iam::123456789012:user/user-name).
	Arn() *string
	// When this Principal is used in an AssumeRole policy, the action to use.
	AssumeRoleAction() *string
	// The principal to grant permissions to.
	GrantPrincipal() IPrincipal
	// Return the policy fragment that identifies this principal in a Policy.
	PolicyFragment() PrincipalPolicyFragment
	// The AWS account ID of this principal.
	//
	// Can be undefined when the account is not known
	// (for example, for service principals).
	// Can be a Token - in that case,
	// it's assumed to be AWS::AccountId.
	PrincipalAccount() *string
	// Add the princpial to the AssumeRolePolicyDocument.
	//
	// Add the statements to the AssumeRolePolicyDocument necessary to give this principal
	// permissions to assume the given role.
	AddToAssumeRolePolicy(document PolicyDocument)
	// Add to the policy of this principal.
	AddToPolicy(statement PolicyStatement) *bool
	// Add to the policy of this principal.
	AddToPrincipalPolicy(_statement PolicyStatement) *AddToPrincipalPolicyResult
	// Return whether or not this principal is equal to the given principal.
	DedupeString() *string
	// A convenience method for adding a condition that the principal is part of the specified AWS Organization.
	InOrganization(organizationId *string) PrincipalBase
	// JSON-ify the principal.
	//
	// Used when JSON.stringify() is called
	ToJSON() *map[string]*[]*string
	// Returns a string representation of an object.
	ToString() *string
	// Returns a new PrincipalWithConditions using this principal as the base, with the passed conditions added.
	//
	// When there is a value for the same operator and key in both the principal and the
	// conditions parameter, the value from the conditions parameter will be used.
	//
	// Returns: a new PrincipalWithConditions object.
	WithConditions(conditions *map[string]interface{}) PrincipalBase
	// Returns a new principal using this principal as the base, with session tags enabled.
	//
	// Returns: a new SessionTagsPrincipal object.
	WithSessionTags() PrincipalBase
}

// The jsii proxy struct for AccountRootPrincipal
type jsiiProxy_AccountRootPrincipal struct {
	jsiiProxy_AccountPrincipal
}

func (j *jsiiProxy_AccountRootPrincipal) AccountId() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"accountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountRootPrincipal) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountRootPrincipal) AssumeRoleAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"assumeRoleAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountRootPrincipal) GrantPrincipal() IPrincipal {
	var returns IPrincipal
	_jsii_.Get(
		j,
		"grantPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountRootPrincipal) PolicyFragment() PrincipalPolicyFragment {
	var returns PrincipalPolicyFragment
	_jsii_.Get(
		j,
		"policyFragment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountRootPrincipal) PrincipalAccount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"principalAccount",
		&returns,
	)
	return returns
}


func NewAccountRootPrincipal() AccountRootPrincipal {
	_init_.Initialize()

	j := jsiiProxy_AccountRootPrincipal{}

	_jsii_.Create(
		"aws-cdk-lib.aws_iam.AccountRootPrincipal",
		nil, // no parameters
		&j,
	)

	return &j
}

func NewAccountRootPrincipal_Override(a AccountRootPrincipal) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_iam.AccountRootPrincipal",
		nil, // no parameters
		a,
	)
}

func (a *jsiiProxy_AccountRootPrincipal) AddToAssumeRolePolicy(document PolicyDocument) {
	if err := a.validateAddToAssumeRolePolicyParameters(document); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addToAssumeRolePolicy",
		[]interface{}{document},
	)
}

func (a *jsiiProxy_AccountRootPrincipal) AddToPolicy(statement PolicyStatement) *bool {
	if err := a.validateAddToPolicyParameters(statement); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.Invoke(
		a,
		"addToPolicy",
		[]interface{}{statement},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccountRootPrincipal) AddToPrincipalPolicy(_statement PolicyStatement) *AddToPrincipalPolicyResult {
	if err := a.validateAddToPrincipalPolicyParameters(_statement); err != nil {
		panic(err)
	}
	var returns *AddToPrincipalPolicyResult

	_jsii_.Invoke(
		a,
		"addToPrincipalPolicy",
		[]interface{}{_statement},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccountRootPrincipal) DedupeString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"dedupeString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccountRootPrincipal) InOrganization(organizationId *string) PrincipalBase {
	if err := a.validateInOrganizationParameters(organizationId); err != nil {
		panic(err)
	}
	var returns PrincipalBase

	_jsii_.Invoke(
		a,
		"inOrganization",
		[]interface{}{organizationId},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccountRootPrincipal) ToJSON() *map[string]*[]*string {
	var returns *map[string]*[]*string

	_jsii_.Invoke(
		a,
		"toJSON",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccountRootPrincipal) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccountRootPrincipal) WithConditions(conditions *map[string]interface{}) PrincipalBase {
	if err := a.validateWithConditionsParameters(conditions); err != nil {
		panic(err)
	}
	var returns PrincipalBase

	_jsii_.Invoke(
		a,
		"withConditions",
		[]interface{}{conditions},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccountRootPrincipal) WithSessionTags() PrincipalBase {
	var returns PrincipalBase

	_jsii_.Invoke(
		a,
		"withSessionTags",
		nil, // no parameters
		&returns,
	)

	return returns
}

