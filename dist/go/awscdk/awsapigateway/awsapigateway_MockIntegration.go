package awsapigateway

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// This type of integration lets API Gateway return a response without sending the request further to the backend.
//
// This is useful for API testing because it
// can be used to test the integration set up without incurring charges for
// using the backend and to enable collaborative development of an API. In
// collaborative development, a team can isolate their development effort by
// setting up simulations of API components owned by other teams by using the
// MOCK integrations. It is also used to return CORS-related headers to ensure
// that the API method permits CORS access. In fact, the API Gateway console
// integrates the OPTIONS method to support CORS with a mock integration.
// Gateway responses are other examples of mock integrations.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/constructs-go/constructs"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   /**
//    * This file showcases how to split up a RestApi's Resources and Methods across nested stacks.
//    *
//    * The root stack 'RootStack' first defines a RestApi.
//    * Two nested stacks BooksStack and PetsStack, create corresponding Resources '/books' and '/pets'.
//    * They are then deployed to a 'prod' Stage via a third nested stack - DeployStack.
//    *
//    * To verify this worked, go to the APIGateway
//    */
//
//   type rootStack struct {
//   	stack
//   }
//
//   func newRootStack(scope construct) *rootStack {
//   	this := &rootStack{}
//   	newStack_Override(this, scope, jsii.String("integ-restapi-import-RootStack"))
//
//   	restApi := awscdk.NewRestApi(this, jsii.String("RestApi"), &restApiProps{
//   		cloudWatchRole: jsii.Boolean(true),
//   		deploy: jsii.Boolean(false),
//   	})
//   	restApi.root.addMethod(jsii.String("ANY"))
//
//   	petsStack := NewPetsStack(this, &resourceNestedStackProps{
//   		restApiId: restApi.restApiId,
//   		rootResourceId: restApi.restApiRootResourceId,
//   	})
//   	booksStack := NewBooksStack(this, &resourceNestedStackProps{
//   		restApiId: restApi.restApiId,
//   		rootResourceId: restApi.restApiRootResourceId,
//   	})
//   	NewDeployStack(this, &deployStackProps{
//   		restApiId: restApi.restApiId,
//   		methods: petsStack.methods.concat(booksStack.methods),
//   	})
//
//   	awscdk.NewCfnOutput(this, jsii.String("PetsURL"), &cfnOutputProps{
//   		value: fmt.Sprintf("https://%v.execute-api.%v.amazonaws.com/prod/pets", restApi.restApiId, this.region),
//   	})
//
//   	awscdk.NewCfnOutput(this, jsii.String("BooksURL"), &cfnOutputProps{
//   		value: fmt.Sprintf("https://%v.execute-api.%v.amazonaws.com/prod/books", restApi.restApiId, this.region),
//   	})
//   	return this
//   }
//
//   type resourceNestedStackProps struct {
//   	nestedStackProps
//   	restApiId *string
//   	rootResourceId *string
//   }
//
//   type petsStack struct {
//   	nestedStack
//   	methods []method
//   }
//
//   func newPetsStack(scope construct, props resourceNestedStackProps) *petsStack {
//   	this := &petsStack{}
//   	newNestedStack_Override(this, scope, jsii.String("integ-restapi-import-PetsStack"), props)
//
//   	api := awscdk.RestApi.fromRestApiAttributes(this, jsii.String("RestApi"), &restApiAttributes{
//   		restApiId: props.restApiId,
//   		rootResourceId: props.rootResourceId,
//   	})
//
//   	method := api.root.addResource(jsii.String("pets")).addMethod(jsii.String("GET"), awscdk.NewMockIntegration(&integrationOptions{
//   		integrationResponses: []integrationResponse{
//   			&integrationResponse{
//   				statusCode: jsii.String("200"),
//   			},
//   		},
//   		passthroughBehavior: awscdk.PassthroughBehavior_NEVER,
//   		requestTemplates: map[string]*string{
//   			"application/json": jsii.String("{ \"statusCode\": 200 }"),
//   		},
//   	}), &methodOptions{
//   		methodResponses: []methodResponse{
//   			&methodResponse{
//   				statusCode: jsii.String("200"),
//   			},
//   		},
//   	})
//
//   	this.methods.push(method)
//   	return this
//   }
//
//   type booksStack struct {
//   	nestedStack
//   	methods []*method
//   }
//
//   func newBooksStack(scope construct, props resourceNestedStackProps) *booksStack {
//   	this := &booksStack{}
//   	newNestedStack_Override(this, scope, jsii.String("integ-restapi-import-BooksStack"), props)
//
//   	api := awscdk.RestApi.fromRestApiAttributes(this, jsii.String("RestApi"), &restApiAttributes{
//   		restApiId: props.restApiId,
//   		rootResourceId: props.rootResourceId,
//   	})
//
//   	method := api.root.addResource(jsii.String("books")).addMethod(jsii.String("GET"), awscdk.NewMockIntegration(&integrationOptions{
//   		integrationResponses: []*integrationResponse{
//   			&integrationResponse{
//   				statusCode: jsii.String("200"),
//   			},
//   		},
//   		passthroughBehavior: awscdk.PassthroughBehavior_NEVER,
//   		requestTemplates: map[string]*string{
//   			"application/json": jsii.String("{ \"statusCode\": 200 }"),
//   		},
//   	}), &methodOptions{
//   		methodResponses: []*methodResponse{
//   			&methodResponse{
//   				statusCode: jsii.String("200"),
//   			},
//   		},
//   	})
//
//   	this.methods.push(method)
//   	return this
//   }
//
//   type deployStackProps struct {
//   	nestedStackProps
//   	restApiId *string
//   	methods []*method
//   }
//
//   type deployStack struct {
//   	nestedStack
//   }
//
//   func newDeployStack(scope construct, props deployStackProps) *deployStack {
//   	this := &deployStack{}
//   	newNestedStack_Override(this, scope, jsii.String("integ-restapi-import-DeployStack"), props)
//
//   	deployment := awscdk.NewDeployment(this, jsii.String("Deployment"), &deploymentProps{
//   		api: awscdk.RestApi.fromRestApiId(this, jsii.String("RestApi"), props.restApiId),
//   	})
//   	if *props.methods {
//   		for _, method := range *props.methods {
//   			deployment.node.addDependency(method)
//   		}
//   	}
//   	awscdk.NewStage(this, jsii.String("Stage"), &stageProps{
//   		deployment: deployment,
//   	})
//   	return this
//   }
//
//   NewRootStack(awscdk.NewApp())
//
type MockIntegration interface {
	Integration
	// Can be overridden by subclasses to allow the integration to interact with the method being integrated, access the REST API object, method ARNs, etc.
	Bind(_method Method) *IntegrationConfig
}

// The jsii proxy struct for MockIntegration
type jsiiProxy_MockIntegration struct {
	jsiiProxy_Integration
}

func NewMockIntegration(options *IntegrationOptions) MockIntegration {
	_init_.Initialize()

	if err := validateNewMockIntegrationParameters(options); err != nil {
		panic(err)
	}
	j := jsiiProxy_MockIntegration{}

	_jsii_.Create(
		"aws-cdk-lib.aws_apigateway.MockIntegration",
		[]interface{}{options},
		&j,
	)

	return &j
}

func NewMockIntegration_Override(m MockIntegration, options *IntegrationOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_apigateway.MockIntegration",
		[]interface{}{options},
		m,
	)
}

func (m *jsiiProxy_MockIntegration) Bind(_method Method) *IntegrationConfig {
	if err := m.validateBindParameters(_method); err != nil {
		panic(err)
	}
	var returns *IntegrationConfig

	_jsii_.Invoke(
		m,
		"bind",
		[]interface{}{_method},
		&returns,
	)

	return returns
}

