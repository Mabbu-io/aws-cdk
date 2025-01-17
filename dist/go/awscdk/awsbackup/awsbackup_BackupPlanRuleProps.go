package awsbackup

import (
	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsevents"
)

// Properties for a BackupPlanRule.
//
// Example:
//   var plan backupPlan
//   var secondaryVault backupVault
//
//   plan.addRule(backup.NewBackupPlanRule(&backupPlanRuleProps{
//   	copyActions: []backupPlanCopyActionProps{
//   		&backupPlanCopyActionProps{
//   			destinationBackupVault: secondaryVault,
//   			moveToColdStorageAfter: awscdk.Duration.days(jsii.Number(30)),
//   			deleteAfter: awscdk.Duration.days(jsii.Number(120)),
//   		},
//   	},
//   }))
//
type BackupPlanRuleProps struct {
	// The backup vault where backups are.
	BackupVault IBackupVault `field:"optional" json:"backupVault" yaml:"backupVault"`
	// The duration after a backup job is successfully started before it must be completed or it is canceled by AWS Backup.
	CompletionWindow awscdk.Duration `field:"optional" json:"completionWindow" yaml:"completionWindow"`
	// Copy operations to perform on recovery points created by this rule.
	CopyActions *[]*BackupPlanCopyActionProps `field:"optional" json:"copyActions" yaml:"copyActions"`
	// Specifies the duration after creation that a recovery point is deleted.
	//
	// Must be greater than `moveToColdStorageAfter`.
	DeleteAfter awscdk.Duration `field:"optional" json:"deleteAfter" yaml:"deleteAfter"`
	// Enables continuous backup and point-in-time restores (PITR).
	//
	// Property `deleteAfter` defines the retention period for the backup. It is mandatory if PITR is enabled.
	// If no value is specified, the retention period is set to 35 days which is the maximum retention period supported by PITR.
	//
	// Property `moveToColdStorageAfter` must not be specified because PITR does not support this option.
	EnableContinuousBackup *bool `field:"optional" json:"enableContinuousBackup" yaml:"enableContinuousBackup"`
	// Specifies the duration after creation that a recovery point is moved to cold storage.
	MoveToColdStorageAfter awscdk.Duration `field:"optional" json:"moveToColdStorageAfter" yaml:"moveToColdStorageAfter"`
	// A display name for the backup rule.
	RuleName *string `field:"optional" json:"ruleName" yaml:"ruleName"`
	// A CRON expression specifying when AWS Backup initiates a backup job.
	ScheduleExpression awsevents.Schedule `field:"optional" json:"scheduleExpression" yaml:"scheduleExpression"`
	// The duration after a backup is scheduled before a job is canceled if it doesn't start successfully.
	StartWindow awscdk.Duration `field:"optional" json:"startWindow" yaml:"startWindow"`
}

