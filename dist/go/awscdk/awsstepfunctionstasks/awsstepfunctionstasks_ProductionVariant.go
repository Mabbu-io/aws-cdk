package awsstepfunctionstasks

import (
	"github.com/aws/aws-cdk-go/awscdk/awsec2"
)

// Identifies a model that you want to host and the resources to deploy for hosting it.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var acceleratorType acceleratorType
//   var instanceType instanceType
//
//   productionVariant := &productionVariant{
//   	instanceType: instanceType,
//   	modelName: jsii.String("modelName"),
//   	variantName: jsii.String("variantName"),
//
//   	// the properties below are optional
//   	acceleratorType: acceleratorType,
//   	initialInstanceCount: jsii.Number(123),
//   	initialVariantWeight: jsii.Number(123),
//   }
//
// See: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_ProductionVariant.html
//
type ProductionVariant struct {
	// The ML compute instance type.
	InstanceType awsec2.InstanceType `field:"required" json:"instanceType" yaml:"instanceType"`
	// The name of the model that you want to host.
	//
	// This is the name that you specified when creating the model.
	ModelName *string `field:"required" json:"modelName" yaml:"modelName"`
	// The name of the production variant.
	VariantName *string `field:"required" json:"variantName" yaml:"variantName"`
	// The size of the Elastic Inference (EI) instance to use for the production variant.
	AcceleratorType AcceleratorType `field:"optional" json:"acceleratorType" yaml:"acceleratorType"`
	// Number of instances to launch initially.
	InitialInstanceCount *float64 `field:"optional" json:"initialInstanceCount" yaml:"initialInstanceCount"`
	// Determines initial traffic distribution among all of the models that you specify in the endpoint configuration.
	InitialVariantWeight *float64 `field:"optional" json:"initialVariantWeight" yaml:"initialVariantWeight"`
}

