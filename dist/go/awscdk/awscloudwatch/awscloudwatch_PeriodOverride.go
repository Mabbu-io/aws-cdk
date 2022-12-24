package awscloudwatch


// Specify the period for graphs when the CloudWatch dashboard loads.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var widget iWidget
//
//   dashboard := awscdk.Aws_cloudwatch.NewDashboard(this, jsii.String("MyDashboard"), &dashboardProps{
//   	dashboardName: jsii.String("dashboardName"),
//   	end: jsii.String("end"),
//   	periodOverride: awscdk.*Aws_cloudwatch.periodOverride_AUTO,
//   	start: jsii.String("start"),
//   	widgets: [][]*iWidget{
//   		[]*iWidget{
//   			widget,
//   		},
//   	},
//   })
//
type PeriodOverride string

const (
	// Period of all graphs on the dashboard automatically adapt to the time range of the dashboard.
	PeriodOverride_AUTO PeriodOverride = "AUTO"
	// Period set for each graph will be used.
	PeriodOverride_INHERIT PeriodOverride = "INHERIT"
)

