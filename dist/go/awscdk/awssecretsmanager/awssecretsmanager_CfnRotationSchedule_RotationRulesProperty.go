package awssecretsmanager


// The rotation schedule and window.
//
// We recommend you use `ScheduleExpression` to set a cron or rate expression for the schedule and `Duration` to set the length of the rotation window.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   rotationRulesProperty := &rotationRulesProperty{
//   	automaticallyAfterDays: jsii.Number(123),
//   	duration: jsii.String("duration"),
//   	scheduleExpression: jsii.String("scheduleExpression"),
//   }
//
type CfnRotationSchedule_RotationRulesProperty struct {
	// The number of days between automatic scheduled rotations of the secret.
	//
	// You can use this value to check that your secret meets your compliance guidelines for how often secrets must be rotated.
	//
	// In `DescribeSecret` and `ListSecrets` , this value is calculated from the rotation schedule after every successful rotation. In `RotateSecret` , you can set the rotation schedule in `RotationRules` with `AutomaticallyAfterDays` or `ScheduleExpression` , but not both.
	AutomaticallyAfterDays *float64 `field:"optional" json:"automaticallyAfterDays" yaml:"automaticallyAfterDays"`
	// The length of the rotation window in hours, for example `3h` for a three hour window.
	//
	// Secrets Manager rotates your secret at any time during this window. The window must not go into the next UTC day. If you don't specify this value, the window automatically ends at the end of the UTC day. The window begins according to the `ScheduleExpression` . For more information, including examples, see [Schedule expressions in Secrets Manager rotation](https://docs.aws.amazon.com/secretsmanager/latest/userguide/rotate-secrets_schedule.html) .
	Duration *string `field:"optional" json:"duration" yaml:"duration"`
	// A `cron()` or `rate()` expression that defines the schedule for rotating your secret.
	//
	// Secrets Manager rotation schedules use UTC time zone.
	//
	// Secrets Manager `rate()` expressions represent the interval in days that you want to rotate your secret, for example `rate(10 days)` . If you use a `rate()` expression, the rotation window opens at midnight, and Secrets Manager rotates your secret any time that day after midnight. You can set a `Duration` to shorten the rotation window.
	//
	// You can use a `cron()` expression to create rotation schedules that are more detailed than a rotation interval. For more information, including examples, see [Schedule expressions in Secrets Manager rotation](https://docs.aws.amazon.com/secretsmanager/latest/userguide/rotate-secrets_schedule.html) . If you use a `cron()` expression, Secrets Manager rotates your secret any time during that day after the window opens. For example, `cron(0 8 1 * ? *)` represents a rotation window that occurs on the first day of every month beginning at 8:00 AM UTC. Secrets Manager rotates the secret any time that day after 8:00 AM. You can set a `Duration` to shorten the rotation window.
	ScheduleExpression *string `field:"optional" json:"scheduleExpression" yaml:"scheduleExpression"`
}

