package customresources

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/customresources/internal"
)

// Reference to the physical resource id that can be passed to the AWS operation as a parameter.
//
// Example:
//   awsCustom := cr.NewAwsCustomResource(this, jsii.String("aws-custom"), &awsCustomResourceProps{
//   	onCreate: &awsSdkCall{
//   		service: jsii.String("..."),
//   		action: jsii.String("..."),
//   		parameters: map[string]*string{
//   			"text": jsii.String("..."),
//   		},
//   		physicalResourceId: cr.physicalResourceId.of(jsii.String("...")),
//   	},
//   	onUpdate: &awsSdkCall{
//   		service: jsii.String("..."),
//   		action: jsii.String("..."),
//   		parameters: map[string]interface{}{
//   			"text": jsii.String("..."),
//   			"resourceId": cr.NewPhysicalResourceIdReference(),
//   		},
//   	},
//   	policy: cr.awsCustomResourcePolicy.fromSdkCalls(&sdkCallsPolicyOptions{
//   		resources: cr.*awsCustomResourcePolicy_ANY_RESOURCE(),
//   	}),
//   })
//
type PhysicalResourceIdReference interface {
	awscdk.IResolvable
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// This may return an array with a single informational element indicating how
	// to get this property populated, if it was skipped for performance reasons.
	CreationStack() *[]*string
	// Produce the Token's value at resolution time.
	Resolve(_arg awscdk.IResolveContext) interface{}
	// toJSON serialization to replace `PhysicalResourceIdReference` with a magic string.
	ToJSON() *string
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	ToString() *string
}

// The jsii proxy struct for PhysicalResourceIdReference
type jsiiProxy_PhysicalResourceIdReference struct {
	internal.Type__awscdkIResolvable
}

func (j *jsiiProxy_PhysicalResourceIdReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}


func NewPhysicalResourceIdReference() PhysicalResourceIdReference {
	_init_.Initialize()

	j := jsiiProxy_PhysicalResourceIdReference{}

	_jsii_.Create(
		"aws-cdk-lib.custom_resources.PhysicalResourceIdReference",
		nil, // no parameters
		&j,
	)

	return &j
}

func NewPhysicalResourceIdReference_Override(p PhysicalResourceIdReference) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.custom_resources.PhysicalResourceIdReference",
		nil, // no parameters
		p,
	)
}

func (p *jsiiProxy_PhysicalResourceIdReference) Resolve(_arg awscdk.IResolveContext) interface{} {
	if err := p.validateResolveParameters(_arg); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		p,
		"resolve",
		[]interface{}{_arg},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PhysicalResourceIdReference) ToJSON() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toJSON",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PhysicalResourceIdReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

