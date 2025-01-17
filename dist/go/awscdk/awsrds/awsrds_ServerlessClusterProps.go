package awsrds

import (
	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/awskms"
)

// Properties for a new Aurora Serverless Cluster.
//
// Example:
//   var vpc vpc
//
//   var code code
//
//
//   cluster := rds.NewServerlessCluster(this, jsii.String("AnotherCluster"), &serverlessClusterProps{
//   	engine: rds.databaseClusterEngine_AURORA_MYSQL(),
//   	vpc: vpc,
//   	 // this parameter is optional for serverless Clusters
//   	enableDataApi: jsii.Boolean(true),
//   })
//   fn := lambda.NewFunction(this, jsii.String("MyFunction"), &functionProps{
//   	runtime: lambda.runtime_NODEJS_14_X(),
//   	handler: jsii.String("index.handler"),
//   	code: code,
//   	environment: map[string]*string{
//   		"CLUSTER_ARN": cluster.clusterArn,
//   		"SECRET_ARN": cluster.secret.secretArn,
//   	},
//   })
//   cluster.grantDataApiAccess(fn)
//
type ServerlessClusterProps struct {
	// What kind of database to start.
	Engine IClusterEngine `field:"required" json:"engine" yaml:"engine"`
	// The number of days during which automatic DB snapshots are retained.
	//
	// Automatic backup retention cannot be disabled on serverless clusters.
	// Must be a value from 1 day to 35 days.
	BackupRetention awscdk.Duration `field:"optional" json:"backupRetention" yaml:"backupRetention"`
	// An optional identifier for the cluster.
	ClusterIdentifier *string `field:"optional" json:"clusterIdentifier" yaml:"clusterIdentifier"`
	// Whether to copy tags to the snapshot when a snapshot is created.
	CopyTagsToSnapshot *bool `field:"optional" json:"copyTagsToSnapshot" yaml:"copyTagsToSnapshot"`
	// Credentials for the administrative user.
	Credentials Credentials `field:"optional" json:"credentials" yaml:"credentials"`
	// Name of a database which is automatically created inside the cluster.
	DefaultDatabaseName *string `field:"optional" json:"defaultDatabaseName" yaml:"defaultDatabaseName"`
	// Indicates whether the DB cluster should have deletion protection enabled.
	DeletionProtection *bool `field:"optional" json:"deletionProtection" yaml:"deletionProtection"`
	// Whether to enable the Data API.
	// See: https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/data-api.html
	//
	EnableDataApi *bool `field:"optional" json:"enableDataApi" yaml:"enableDataApi"`
	// Additional parameters to pass to the database engine.
	ParameterGroup IParameterGroup `field:"optional" json:"parameterGroup" yaml:"parameterGroup"`
	// The removal policy to apply when the cluster and its instances are removed from the stack or replaced during an update.
	RemovalPolicy awscdk.RemovalPolicy `field:"optional" json:"removalPolicy" yaml:"removalPolicy"`
	// Scaling configuration of an Aurora Serverless database cluster.
	Scaling *ServerlessScalingOptions `field:"optional" json:"scaling" yaml:"scaling"`
	// Security group.
	SecurityGroups *[]awsec2.ISecurityGroup `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// The KMS key for storage encryption.
	StorageEncryptionKey awskms.IKey `field:"optional" json:"storageEncryptionKey" yaml:"storageEncryptionKey"`
	// Existing subnet group for the cluster.
	SubnetGroup ISubnetGroup `field:"optional" json:"subnetGroup" yaml:"subnetGroup"`
	// The VPC that this Aurora Serverless cluster has been created in.
	Vpc awsec2.IVpc `field:"optional" json:"vpc" yaml:"vpc"`
	// Where to place the instances within the VPC.
	//
	// If provided, the `vpc` property must also be specified.
	VpcSubnets *awsec2.SubnetSelection `field:"optional" json:"vpcSubnets" yaml:"vpcSubnets"`
}

