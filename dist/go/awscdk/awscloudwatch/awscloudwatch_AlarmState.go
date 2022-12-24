package awscloudwatch


// Enumeration indicates state of Alarm used in building Alarm Rule.
//
// Example:
//   var alarm1 alarm
//   var alarm2 alarm
//   var alarm3 alarm
//   var alarm4 alarm
//
//
//   alarmRule := cloudwatch.alarmRule.anyOf(cloudwatch.alarmRule.allOf(cloudwatch.alarmRule.anyOf(alarm1, cloudwatch.alarmRule.fromAlarm(alarm2, cloudwatch.alarmState_OK), alarm3), cloudwatch.alarmRule.not(cloudwatch.alarmRule.fromAlarm(alarm4, cloudwatch.alarmState_INSUFFICIENT_DATA))), cloudwatch.alarmRule.fromBoolean(jsii.Boolean(false)))
//
//   cloudwatch.NewCompositeAlarm(this, jsii.String("MyAwesomeCompositeAlarm"), &compositeAlarmProps{
//   	alarmRule: alarmRule,
//   })
//
type AlarmState string

const (
	// State indicates resource is in ALARM.
	AlarmState_ALARM AlarmState = "ALARM"
	// State indicates resource is not in ALARM.
	AlarmState_OK AlarmState = "OK"
	// State indicates there is not enough data to determine is resource is in ALARM.
	AlarmState_INSUFFICIENT_DATA AlarmState = "INSUFFICIENT_DATA"
)

