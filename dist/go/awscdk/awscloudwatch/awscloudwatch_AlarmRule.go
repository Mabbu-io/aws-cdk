package awscloudwatch

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Class with static functions to build AlarmRule for Composite Alarms.
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
type AlarmRule interface {
}

// The jsii proxy struct for AlarmRule
type jsiiProxy_AlarmRule struct {
	_ byte // padding
}

func NewAlarmRule() AlarmRule {
	_init_.Initialize()

	j := jsiiProxy_AlarmRule{}

	_jsii_.Create(
		"aws-cdk-lib.aws_cloudwatch.AlarmRule",
		nil, // no parameters
		&j,
	)

	return &j
}

func NewAlarmRule_Override(a AlarmRule) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_cloudwatch.AlarmRule",
		nil, // no parameters
		a,
	)
}

// function to join all provided AlarmRules with AND operator.
func AlarmRule_AllOf(operands ...IAlarmRule) IAlarmRule {
	_init_.Initialize()

	args := []interface{}{}
	for _, a := range operands {
		args = append(args, a)
	}

	var returns IAlarmRule

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudwatch.AlarmRule",
		"allOf",
		args,
		&returns,
	)

	return returns
}

// function to join all provided AlarmRules with OR operator.
func AlarmRule_AnyOf(operands ...IAlarmRule) IAlarmRule {
	_init_.Initialize()

	args := []interface{}{}
	for _, a := range operands {
		args = append(args, a)
	}

	var returns IAlarmRule

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudwatch.AlarmRule",
		"anyOf",
		args,
		&returns,
	)

	return returns
}

// function to build Rule Expression for given IAlarm and AlarmState.
func AlarmRule_FromAlarm(alarm IAlarm, alarmState AlarmState) IAlarmRule {
	_init_.Initialize()

	if err := validateAlarmRule_FromAlarmParameters(alarm, alarmState); err != nil {
		panic(err)
	}
	var returns IAlarmRule

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudwatch.AlarmRule",
		"fromAlarm",
		[]interface{}{alarm, alarmState},
		&returns,
	)

	return returns
}

// function to build TRUE/FALSE intent for Rule Expression.
func AlarmRule_FromBoolean(value *bool) IAlarmRule {
	_init_.Initialize()

	if err := validateAlarmRule_FromBooleanParameters(value); err != nil {
		panic(err)
	}
	var returns IAlarmRule

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudwatch.AlarmRule",
		"fromBoolean",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// function to build Rule Expression for given Alarm Rule string.
func AlarmRule_FromString(alarmRule *string) IAlarmRule {
	_init_.Initialize()

	if err := validateAlarmRule_FromStringParameters(alarmRule); err != nil {
		panic(err)
	}
	var returns IAlarmRule

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudwatch.AlarmRule",
		"fromString",
		[]interface{}{alarmRule},
		&returns,
	)

	return returns
}

// function to wrap provided AlarmRule in NOT operator.
func AlarmRule_Not(operand IAlarmRule) IAlarmRule {
	_init_.Initialize()

	if err := validateAlarmRule_NotParameters(operand); err != nil {
		panic(err)
	}
	var returns IAlarmRule

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudwatch.AlarmRule",
		"not",
		[]interface{}{operand},
		&returns,
	)

	return returns
}

