package awsfsx


// Enum for representing all the days of the week.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   lustreMaintenanceTime := awscdk.Aws_fsx.NewLustreMaintenanceTime(&lustreMaintenanceTimeProps{
//   	day: awscdk.*Aws_fsx.weekday_MONDAY,
//   	hour: jsii.Number(123),
//   	minute: jsii.Number(123),
//   })
//
type Weekday string

const (
	// Monday.
	Weekday_MONDAY Weekday = "MONDAY"
	// Tuesday.
	Weekday_TUESDAY Weekday = "TUESDAY"
	// Wednesday.
	Weekday_WEDNESDAY Weekday = "WEDNESDAY"
	// Thursday.
	Weekday_THURSDAY Weekday = "THURSDAY"
	// Friday.
	Weekday_FRIDAY Weekday = "FRIDAY"
	// Saturday.
	Weekday_SATURDAY Weekday = "SATURDAY"
	// Sunday.
	Weekday_SUNDAY Weekday = "SUNDAY"
)

