package awsrds


// The retention period for Performance Insight.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var bucket bucket
//   var key key
//   var optionGroup optionGroup
//   var parameterGroup parameterGroup
//   var role role
//   var securityGroup securityGroup
//   var subnet subnet
//   var subnetFilter subnetFilter
//   var subnetGroup subnetGroup
//   var vpc vpc
//
//   databaseInstanceNewProps := &databaseInstanceNewProps{
//   	vpc: vpc,
//
//   	// the properties below are optional
//   	autoMinorVersionUpgrade: jsii.Boolean(false),
//   	availabilityZone: jsii.String("availabilityZone"),
//   	backupRetention: cdk.duration.minutes(jsii.Number(30)),
//   	cloudwatchLogsExports: []*string{
//   		jsii.String("cloudwatchLogsExports"),
//   	},
//   	cloudwatchLogsRetention: awscdk.Aws_logs.retentionDays_ONE_DAY,
//   	cloudwatchLogsRetentionRole: role,
//   	copyTagsToSnapshot: jsii.Boolean(false),
//   	deleteAutomatedBackups: jsii.Boolean(false),
//   	deletionProtection: jsii.Boolean(false),
//   	domain: jsii.String("domain"),
//   	domainRole: role,
//   	enablePerformanceInsights: jsii.Boolean(false),
//   	iamAuthentication: jsii.Boolean(false),
//   	instanceIdentifier: jsii.String("instanceIdentifier"),
//   	iops: jsii.Number(123),
//   	maxAllocatedStorage: jsii.Number(123),
//   	monitoringInterval: cdk.*duration.minutes(jsii.Number(30)),
//   	monitoringRole: role,
//   	multiAz: jsii.Boolean(false),
//   	networkType: awscdk.Aws_rds.networkType_IPV4,
//   	optionGroup: optionGroup,
//   	parameterGroup: parameterGroup,
//   	performanceInsightEncryptionKey: key,
//   	performanceInsightRetention: awscdk.*Aws_rds.performanceInsightRetention_DEFAULT,
//   	port: jsii.Number(123),
//   	preferredBackupWindow: jsii.String("preferredBackupWindow"),
//   	preferredMaintenanceWindow: jsii.String("preferredMaintenanceWindow"),
//   	processorFeatures: &processorFeatures{
//   		coreCount: jsii.Number(123),
//   		threadsPerCore: jsii.Number(123),
//   	},
//   	publiclyAccessible: jsii.Boolean(false),
//   	removalPolicy: cdk.removalPolicy_DESTROY,
//   	s3ExportBuckets: []iBucket{
//   		bucket,
//   	},
//   	s3ExportRole: role,
//   	s3ImportBuckets: []*iBucket{
//   		bucket,
//   	},
//   	s3ImportRole: role,
//   	securityGroups: []iSecurityGroup{
//   		securityGroup,
//   	},
//   	storageThroughput: jsii.Number(123),
//   	storageType: awscdk.*Aws_rds.storageType_STANDARD,
//   	subnetGroup: subnetGroup,
//   	vpcSubnets: &subnetSelection{
//   		availabilityZones: []*string{
//   			jsii.String("availabilityZones"),
//   		},
//   		onePerAz: jsii.Boolean(false),
//   		subnetFilters: []*subnetFilter{
//   			subnetFilter,
//   		},
//   		subnetGroupName: jsii.String("subnetGroupName"),
//   		subnets: []iSubnet{
//   			subnet,
//   		},
//   		subnetType: awscdk.Aws_ec2.subnetType_PRIVATE_ISOLATED,
//   	},
//   }
//
type PerformanceInsightRetention string

const (
	// Default retention period of 7 days.
	PerformanceInsightRetention_DEFAULT PerformanceInsightRetention = "DEFAULT"
	// Long term retention period of 2 years.
	PerformanceInsightRetention_LONG_TERM PerformanceInsightRetention = "LONG_TERM"
)

