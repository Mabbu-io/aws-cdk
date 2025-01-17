package awscodebuild

import (
	"github.com/aws/aws-cdk-go/awscdk/awss3"
)

// Construction properties for {@link S3Source}.
//
// Example:
//   bucket := s3.NewBucket(this, jsii.String("MyBucket"))
//
//   codebuild.NewProject(this, jsii.String("MyProject"), &projectProps{
//   	source: codebuild.source.s3(&s3SourceProps{
//   		bucket: bucket,
//   		path: jsii.String("path/to/file.zip"),
//   	}),
//   })
//
type S3SourceProps struct {
	// The source identifier.
	//
	// This property is required on secondary sources.
	Identifier *string `field:"optional" json:"identifier" yaml:"identifier"`
	Bucket awss3.IBucket `field:"required" json:"bucket" yaml:"bucket"`
	Path *string `field:"required" json:"path" yaml:"path"`
	// The version ID of the object that represents the build input ZIP file to use.
	Version *string `field:"optional" json:"version" yaml:"version"`
}

