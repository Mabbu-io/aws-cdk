package awscodebuild

import (
	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awss3"
)

// Construction properties for {@link ReportGroup}.
//
// Example:
//   var source source
//
//
//   // create a new ReportGroup
//   reportGroup := codebuild.NewReportGroup(this, jsii.String("ReportGroup"), &reportGroupProps{
//   	type: codebuild.reportGroupType_CODE_COVERAGE,
//   })
//
//   project := codebuild.NewProject(this, jsii.String("Project"), &projectProps{
//   	source: source,
//   	buildSpec: codebuild.buildSpec.fromObject(map[string]interface{}{
//   		// ...
//   		"reports": map[string]map[string]*string{
//   			reportGroup.reportGroupArn: map[string]*string{
//   				"files": jsii.String("**/*"),
//   				"base-directory": jsii.String("build/coverage-report.xml"),
//   				"file-format": jsii.String("JACOCOXML"),
//   			},
//   		},
//   	}),
//   })
//
type ReportGroupProps struct {
	// An optional S3 bucket to export the reports to.
	ExportBucket awss3.IBucket `field:"optional" json:"exportBucket" yaml:"exportBucket"`
	// What to do when this resource is deleted from a stack.
	//
	// As CodeBuild does not allow deleting a ResourceGroup that has reports inside of it,
	// this is set to retain the resource by default.
	RemovalPolicy awscdk.RemovalPolicy `field:"optional" json:"removalPolicy" yaml:"removalPolicy"`
	// The physical name of the report group.
	ReportGroupName *string `field:"optional" json:"reportGroupName" yaml:"reportGroupName"`
	// The type of report group. This can be one of the following values:.
	//
	// - **TEST** - The report group contains test reports.
	// - **CODE_COVERAGE** - The report group contains code coverage reports.
	Type ReportGroupType `field:"optional" json:"type" yaml:"type"`
	// Whether to output the report files into the export bucket as-is, or create a ZIP from them before doing the export.
	//
	// Ignored if {@link exportBucket} has not been provided.
	ZipExport *bool `field:"optional" json:"zipExport" yaml:"zipExport"`
}

