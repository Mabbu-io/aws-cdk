package awsrds

import (
	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/awskms"
	"github.com/aws/aws-cdk-go/awscdk/awslogs"
	"github.com/aws/aws-cdk-go/awscdk/awss3"
)

// Construction properties for a DatabaseInstanceReadReplica.
//
// Example:
//   var vpc vpc
//
//   var sourceInstance databaseInstance
//
//   rds.NewDatabaseInstanceFromSnapshot(this, jsii.String("Instance"), &databaseInstanceFromSnapshotProps{
//   	snapshotIdentifier: jsii.String("my-snapshot"),
//   	engine: rds.databaseInstanceEngine.postgres(&postgresInstanceEngineProps{
//   		version: rds.postgresEngineVersion_VER_12_3(),
//   	}),
//   	// optional, defaults to m5.large
//   	instanceType: ec2.instanceType.of(ec2.instanceClass_BURSTABLE2, ec2.instanceSize_LARGE),
//   	vpc: vpc,
//   })
//   rds.NewDatabaseInstanceReadReplica(this, jsii.String("ReadReplica"), &databaseInstanceReadReplicaProps{
//   	sourceDatabaseInstance: sourceInstance,
//   	instanceType: ec2.*instanceType.of(ec2.*instanceClass_BURSTABLE2, ec2.*instanceSize_LARGE),
//   	vpc: vpc,
//   })
//
type DatabaseInstanceReadReplicaProps struct {
	// The VPC network where the DB subnet group should be created.
	Vpc awsec2.IVpc `field:"required" json:"vpc" yaml:"vpc"`
	// Indicates that minor engine upgrades are applied automatically to the DB instance during the maintenance window.
	AutoMinorVersionUpgrade *bool `field:"optional" json:"autoMinorVersionUpgrade" yaml:"autoMinorVersionUpgrade"`
	// The name of the Availability Zone where the DB instance will be located.
	AvailabilityZone *string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// The number of days during which automatic DB snapshots are retained.
	//
	// Set to zero to disable backups.
	// When creating a read replica, you must enable automatic backups on the source
	// database instance by setting the backup retention to a value other than zero.
	BackupRetention awscdk.Duration `field:"optional" json:"backupRetention" yaml:"backupRetention"`
	// The list of log types that need to be enabled for exporting to CloudWatch Logs.
	CloudwatchLogsExports *[]*string `field:"optional" json:"cloudwatchLogsExports" yaml:"cloudwatchLogsExports"`
	// The number of days log events are kept in CloudWatch Logs.
	//
	// When updating
	// this property, unsetting it doesn't remove the log retention policy. To
	// remove the retention policy, set the value to `Infinity`.
	CloudwatchLogsRetention awslogs.RetentionDays `field:"optional" json:"cloudwatchLogsRetention" yaml:"cloudwatchLogsRetention"`
	// The IAM role for the Lambda function associated with the custom resource that sets the retention policy.
	CloudwatchLogsRetentionRole awsiam.IRole `field:"optional" json:"cloudwatchLogsRetentionRole" yaml:"cloudwatchLogsRetentionRole"`
	// Indicates whether to copy all of the user-defined tags from the DB instance to snapshots of the DB instance.
	CopyTagsToSnapshot *bool `field:"optional" json:"copyTagsToSnapshot" yaml:"copyTagsToSnapshot"`
	// Indicates whether automated backups should be deleted or retained when you delete a DB instance.
	DeleteAutomatedBackups *bool `field:"optional" json:"deleteAutomatedBackups" yaml:"deleteAutomatedBackups"`
	// Indicates whether the DB instance should have deletion protection enabled.
	DeletionProtection *bool `field:"optional" json:"deletionProtection" yaml:"deletionProtection"`
	// The Active Directory directory ID to create the DB instance in.
	Domain *string `field:"optional" json:"domain" yaml:"domain"`
	// The IAM role to be used when making API calls to the Directory Service.
	//
	// The role needs the AWS-managed policy
	// AmazonRDSDirectoryServiceAccess or equivalent.
	DomainRole awsiam.IRole `field:"optional" json:"domainRole" yaml:"domainRole"`
	// Whether to enable Performance Insights for the DB instance.
	EnablePerformanceInsights *bool `field:"optional" json:"enablePerformanceInsights" yaml:"enablePerformanceInsights"`
	// Whether to enable mapping of AWS Identity and Access Management (IAM) accounts to database accounts.
	IamAuthentication *bool `field:"optional" json:"iamAuthentication" yaml:"iamAuthentication"`
	// A name for the DB instance.
	//
	// If you specify a name, AWS CloudFormation
	// converts it to lowercase.
	InstanceIdentifier *string `field:"optional" json:"instanceIdentifier" yaml:"instanceIdentifier"`
	// The number of I/O operations per second (IOPS) that the database provisions.
	//
	// The value must be equal to or greater than 1000.
	Iops *float64 `field:"optional" json:"iops" yaml:"iops"`
	// Upper limit to which RDS can scale the storage in GiB(Gibibyte).
	// See: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_PIOPS.StorageTypes.html#USER_PIOPS.Autoscaling
	//
	MaxAllocatedStorage *float64 `field:"optional" json:"maxAllocatedStorage" yaml:"maxAllocatedStorage"`
	// The interval, in seconds, between points when Amazon RDS collects enhanced monitoring metrics for the DB instance.
	MonitoringInterval awscdk.Duration `field:"optional" json:"monitoringInterval" yaml:"monitoringInterval"`
	// Role that will be used to manage DB instance monitoring.
	MonitoringRole awsiam.IRole `field:"optional" json:"monitoringRole" yaml:"monitoringRole"`
	// Specifies if the database instance is a multiple Availability Zone deployment.
	MultiAz *bool `field:"optional" json:"multiAz" yaml:"multiAz"`
	// The network type of the DB instance.
	NetworkType NetworkType `field:"optional" json:"networkType" yaml:"networkType"`
	// The option group to associate with the instance.
	OptionGroup IOptionGroup `field:"optional" json:"optionGroup" yaml:"optionGroup"`
	// The DB parameter group to associate with the instance.
	ParameterGroup IParameterGroup `field:"optional" json:"parameterGroup" yaml:"parameterGroup"`
	// The AWS KMS key for encryption of Performance Insights data.
	PerformanceInsightEncryptionKey awskms.IKey `field:"optional" json:"performanceInsightEncryptionKey" yaml:"performanceInsightEncryptionKey"`
	// The amount of time, in days, to retain Performance Insights data.
	PerformanceInsightRetention PerformanceInsightRetention `field:"optional" json:"performanceInsightRetention" yaml:"performanceInsightRetention"`
	// The port for the instance.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// The daily time range during which automated backups are performed.
	//
	// Constraints:
	// - Must be in the format `hh24:mi-hh24:mi`.
	// - Must be in Universal Coordinated Time (UTC).
	// - Must not conflict with the preferred maintenance window.
	// - Must be at least 30 minutes.
	PreferredBackupWindow *string `field:"optional" json:"preferredBackupWindow" yaml:"preferredBackupWindow"`
	// The weekly time range (in UTC) during which system maintenance can occur.
	//
	// Format: `ddd:hh24:mi-ddd:hh24:mi`
	// Constraint: Minimum 30-minute window.
	PreferredMaintenanceWindow *string `field:"optional" json:"preferredMaintenanceWindow" yaml:"preferredMaintenanceWindow"`
	// The number of CPU cores and the number of threads per core.
	ProcessorFeatures *ProcessorFeatures `field:"optional" json:"processorFeatures" yaml:"processorFeatures"`
	// Indicates whether the DB instance is an internet-facing instance.
	PubliclyAccessible *bool `field:"optional" json:"publiclyAccessible" yaml:"publiclyAccessible"`
	// The CloudFormation policy to apply when the instance is removed from the stack or replaced during an update.
	RemovalPolicy awscdk.RemovalPolicy `field:"optional" json:"removalPolicy" yaml:"removalPolicy"`
	// S3 buckets that you want to load data into.
	//
	// This property must not be used if `s3ExportRole` is used.
	//
	// For Microsoft SQL Server:.
	// See: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/oracle-s3-integration.html
	//
	S3ExportBuckets *[]awss3.IBucket `field:"optional" json:"s3ExportBuckets" yaml:"s3ExportBuckets"`
	// Role that will be associated with this DB instance to enable S3 export.
	//
	// This property must not be used if `s3ExportBuckets` is used.
	//
	// For Microsoft SQL Server:.
	// See: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/oracle-s3-integration.html
	//
	S3ExportRole awsiam.IRole `field:"optional" json:"s3ExportRole" yaml:"s3ExportRole"`
	// S3 buckets that you want to load data from.
	//
	// This feature is only supported by the Microsoft SQL Server, Oracle, and PostgreSQL engines.
	//
	// This property must not be used if `s3ImportRole` is used.
	//
	// For Microsoft SQL Server:.
	// See: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/PostgreSQL.Procedural.Importing.html
	//
	S3ImportBuckets *[]awss3.IBucket `field:"optional" json:"s3ImportBuckets" yaml:"s3ImportBuckets"`
	// Role that will be associated with this DB instance to enable S3 import.
	//
	// This feature is only supported by the Microsoft SQL Server, Oracle, and PostgreSQL engines.
	//
	// This property must not be used if `s3ImportBuckets` is used.
	//
	// For Microsoft SQL Server:.
	// See: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/PostgreSQL.Procedural.Importing.html
	//
	S3ImportRole awsiam.IRole `field:"optional" json:"s3ImportRole" yaml:"s3ImportRole"`
	// The security groups to assign to the DB instance.
	SecurityGroups *[]awsec2.ISecurityGroup `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// The storage throughput, specified in mebibytes per second (MiBps).
	//
	// Only applicable for GP3.
	// See: https://docs.aws.amazon.com//AmazonRDS/latest/UserGuide/CHAP_Storage.html#gp3-storage
	//
	StorageThroughput *float64 `field:"optional" json:"storageThroughput" yaml:"storageThroughput"`
	// The storage type.
	//
	// Storage types supported are gp2, io1, standard.
	// See: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_Storage.html#Concepts.Storage.GeneralSSD
	//
	StorageType StorageType `field:"optional" json:"storageType" yaml:"storageType"`
	// Existing subnet group for the instance.
	SubnetGroup ISubnetGroup `field:"optional" json:"subnetGroup" yaml:"subnetGroup"`
	// The type of subnets to add to the created DB subnet group.
	VpcSubnets *awsec2.SubnetSelection `field:"optional" json:"vpcSubnets" yaml:"vpcSubnets"`
	// The name of the compute and memory capacity classes.
	InstanceType awsec2.InstanceType `field:"required" json:"instanceType" yaml:"instanceType"`
	// The source database instance.
	//
	// Each DB instance can have a limited number of read replicas. For more
	// information, see https://docs.aws.amazon.com/AmazonRDS/latest/DeveloperGuide/USER_ReadRepl.html.
	SourceDatabaseInstance IDatabaseInstance `field:"required" json:"sourceDatabaseInstance" yaml:"sourceDatabaseInstance"`
	// Indicates whether the DB instance is encrypted.
	StorageEncrypted *bool `field:"optional" json:"storageEncrypted" yaml:"storageEncrypted"`
	// The KMS key that's used to encrypt the DB instance.
	StorageEncryptionKey awskms.IKey `field:"optional" json:"storageEncryptionKey" yaml:"storageEncryptionKey"`
}

