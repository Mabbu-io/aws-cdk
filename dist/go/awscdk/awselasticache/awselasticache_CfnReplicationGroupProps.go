package awselasticache

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnReplicationGroup`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnReplicationGroupProps := &cfnReplicationGroupProps{
//   	replicationGroupDescription: jsii.String("replicationGroupDescription"),
//
//   	// the properties below are optional
//   	atRestEncryptionEnabled: jsii.Boolean(false),
//   	authToken: jsii.String("authToken"),
//   	automaticFailoverEnabled: jsii.Boolean(false),
//   	autoMinorVersionUpgrade: jsii.Boolean(false),
//   	cacheNodeType: jsii.String("cacheNodeType"),
//   	cacheParameterGroupName: jsii.String("cacheParameterGroupName"),
//   	cacheSecurityGroupNames: []*string{
//   		jsii.String("cacheSecurityGroupNames"),
//   	},
//   	cacheSubnetGroupName: jsii.String("cacheSubnetGroupName"),
//   	dataTieringEnabled: jsii.Boolean(false),
//   	engine: jsii.String("engine"),
//   	engineVersion: jsii.String("engineVersion"),
//   	globalReplicationGroupId: jsii.String("globalReplicationGroupId"),
//   	ipDiscovery: jsii.String("ipDiscovery"),
//   	kmsKeyId: jsii.String("kmsKeyId"),
//   	logDeliveryConfigurations: []interface{}{
//   		&logDeliveryConfigurationRequestProperty{
//   			destinationDetails: &destinationDetailsProperty{
//   				cloudWatchLogsDetails: &cloudWatchLogsDestinationDetailsProperty{
//   					logGroup: jsii.String("logGroup"),
//   				},
//   				kinesisFirehoseDetails: &kinesisFirehoseDestinationDetailsProperty{
//   					deliveryStream: jsii.String("deliveryStream"),
//   				},
//   			},
//   			destinationType: jsii.String("destinationType"),
//   			logFormat: jsii.String("logFormat"),
//   			logType: jsii.String("logType"),
//   		},
//   	},
//   	multiAzEnabled: jsii.Boolean(false),
//   	networkType: jsii.String("networkType"),
//   	nodeGroupConfiguration: []interface{}{
//   		&nodeGroupConfigurationProperty{
//   			nodeGroupId: jsii.String("nodeGroupId"),
//   			primaryAvailabilityZone: jsii.String("primaryAvailabilityZone"),
//   			replicaAvailabilityZones: []*string{
//   				jsii.String("replicaAvailabilityZones"),
//   			},
//   			replicaCount: jsii.Number(123),
//   			slots: jsii.String("slots"),
//   		},
//   	},
//   	notificationTopicArn: jsii.String("notificationTopicArn"),
//   	numCacheClusters: jsii.Number(123),
//   	numNodeGroups: jsii.Number(123),
//   	port: jsii.Number(123),
//   	preferredCacheClusterAZs: []*string{
//   		jsii.String("preferredCacheClusterAZs"),
//   	},
//   	preferredMaintenanceWindow: jsii.String("preferredMaintenanceWindow"),
//   	primaryClusterId: jsii.String("primaryClusterId"),
//   	replicasPerNodeGroup: jsii.Number(123),
//   	replicationGroupId: jsii.String("replicationGroupId"),
//   	securityGroupIds: []*string{
//   		jsii.String("securityGroupIds"),
//   	},
//   	snapshotArns: []*string{
//   		jsii.String("snapshotArns"),
//   	},
//   	snapshotName: jsii.String("snapshotName"),
//   	snapshotRetentionLimit: jsii.Number(123),
//   	snapshottingClusterId: jsii.String("snapshottingClusterId"),
//   	snapshotWindow: jsii.String("snapshotWindow"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	transitEncryptionEnabled: jsii.Boolean(false),
//   	userGroupIds: []*string{
//   		jsii.String("userGroupIds"),
//   	},
//   }
//
type CfnReplicationGroupProps struct {
	// A user-created description for the replication group.
	ReplicationGroupDescription *string `field:"required" json:"replicationGroupDescription" yaml:"replicationGroupDescription"`
	// A flag that enables encryption at rest when set to `true` .
	//
	// You cannot modify the value of `AtRestEncryptionEnabled` after the replication group is created. To enable encryption at rest on a replication group you must set `AtRestEncryptionEnabled` to `true` when you create the replication group.
	//
	// *Required:* Only available when creating a replication group in an Amazon VPC using redis version `3.2.6` or `4.x` onward.
	//
	// Default: `false`.
	AtRestEncryptionEnabled interface{} `field:"optional" json:"atRestEncryptionEnabled" yaml:"atRestEncryptionEnabled"`
	// *Reserved parameter.* The password used to access a password protected server.
	//
	// `AuthToken` can be specified only on replication groups where `TransitEncryptionEnabled` is `true` . For more information, see [Authenticating Users with the Redis AUTH Command](https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/auth.html) .
	//
	// > For HIPAA compliance, you must specify `TransitEncryptionEnabled` as `true` , an `AuthToken` , and a `CacheSubnetGroup` .
	//
	// Password constraints:
	//
	// - Must be only printable ASCII characters.
	// - Must be at least 16 characters and no more than 128 characters in length.
	// - Nonalphanumeric characters are restricted to (!, &, #, $, ^, <, >, -, ).
	//
	// For more information, see [AUTH password](https://docs.aws.amazon.com/http://redis.io/commands/AUTH) at http://redis.io/commands/AUTH.
	AuthToken *string `field:"optional" json:"authToken" yaml:"authToken"`
	// Specifies whether a read-only replica is automatically promoted to read/write primary if the existing primary fails.
	//
	// `AutomaticFailoverEnabled` must be enabled for Redis (cluster mode enabled) replication groups.
	//
	// Default: false.
	AutomaticFailoverEnabled interface{} `field:"optional" json:"automaticFailoverEnabled" yaml:"automaticFailoverEnabled"`
	// If you are running Redis engine version 6.0 or later, set this parameter to yes if you want to opt-in to the next minor version upgrade campaign. This parameter is disabled for previous versions.
	AutoMinorVersionUpgrade interface{} `field:"optional" json:"autoMinorVersionUpgrade" yaml:"autoMinorVersionUpgrade"`
	// The compute and memory capacity of the nodes in the node group (shard).
	//
	// The following node types are supported by ElastiCache. Generally speaking, the current generation types provide more memory and computational power at lower cost when compared to their equivalent previous generation counterparts.
	//
	// - General purpose:
	//
	// - Current generation:
	//
	// *M6g node types:* `cache.m6g.large` , `cache.m6g.xlarge` , `cache.m6g.2xlarge` , `cache.m6g.4xlarge` , `cache.m6g.12xlarge` , `cache.m6g.24xlarge`
	//
	// *M5 node types:* `cache.m5.large` , `cache.m5.xlarge` , `cache.m5.2xlarge` , `cache.m5.4xlarge` , `cache.m5.12xlarge` , `cache.m5.24xlarge`
	//
	// *M4 node types:* `cache.m4.large` , `cache.m4.xlarge` , `cache.m4.2xlarge` , `cache.m4.4xlarge` , `cache.m4.10xlarge`
	//
	// *T4g node types:* `cache.t4g.micro` , `cache.t4g.small` , `cache.t4g.medium`
	//
	// *T3 node types:* `cache.t3.micro` , `cache.t3.small` , `cache.t3.medium`
	//
	// *T2 node types:* `cache.t2.micro` , `cache.t2.small` , `cache.t2.medium`
	// - Previous generation: (not recommended)
	//
	// *T1 node types:* `cache.t1.micro`
	//
	// *M1 node types:* `cache.m1.small` , `cache.m1.medium` , `cache.m1.large` , `cache.m1.xlarge`
	//
	// *M3 node types:* `cache.m3.medium` , `cache.m3.large` , `cache.m3.xlarge` , `cache.m3.2xlarge`
	// - Compute optimized:
	//
	// - Previous generation: (not recommended)
	//
	// *C1 node types:* `cache.c1.xlarge`
	// - Memory optimized:
	//
	// - Current generation:
	//
	// *R6gd node types:* `cache.r6gd.xlarge` , `cache.r6gd.2xlarge` , `cache.r6gd.4xlarge` , `cache.r6gd.8xlarge` , `cache.r6gd.12xlarge` , `cache.r6gd.16xlarge`
	//
	// > The `r6gd` family is available in the following regions: `us-east-2` , `us-east-1` , `us-west-2` , `us-west-1` , `eu-west-1` , `eu-central-1` , `ap-northeast-1` , `ap-southeast-1` , `ap-southeast-2` .
	//
	// *R6g node types:* `cache.r6g.large` , `cache.r6g.xlarge` , `cache.r6g.2xlarge` , `cache.r6g.4xlarge` , `cache.r6g.12xlarge` , `cache.r6g.24xlarge`
	//
	// *R5 node types:* `cache.r5.large` , `cache.r5.xlarge` , `cache.r5.2xlarge` , `cache.r5.4xlarge` , `cache.r5.12xlarge` , `cache.r5.24xlarge`
	//
	// *R4 node types:* `cache.r4.large` , `cache.r4.xlarge` , `cache.r4.2xlarge` , `cache.r4.4xlarge` , `cache.r4.8xlarge` , `cache.r4.16xlarge`
	// - Previous generation: (not recommended)
	//
	// *M2 node types:* `cache.m2.xlarge` , `cache.m2.2xlarge` , `cache.m2.4xlarge`
	//
	// *R3 node types:* `cache.r3.large` , `cache.r3.xlarge` , `cache.r3.2xlarge` , `cache.r3.4xlarge` , `cache.r3.8xlarge`
	//
	// For region availability, see [Supported Node Types by Amazon Region](https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/CacheNodes.SupportedTypes.html#CacheNodes.SupportedTypesByRegion)
	CacheNodeType *string `field:"optional" json:"cacheNodeType" yaml:"cacheNodeType"`
	// The name of the parameter group to associate with this replication group.
	//
	// If this argument is omitted, the default cache parameter group for the specified engine is used.
	//
	// If you are running Redis version 3.2.4 or later, only one node group (shard), and want to use a default parameter group, we recommend that you specify the parameter group by name.
	//
	// - To create a Redis (cluster mode disabled) replication group, use `CacheParameterGroupName=default.redis3.2` .
	// - To create a Redis (cluster mode enabled) replication group, use `CacheParameterGroupName=default.redis3.2.cluster.on` .
	CacheParameterGroupName *string `field:"optional" json:"cacheParameterGroupName" yaml:"cacheParameterGroupName"`
	// A list of cache security group names to associate with this replication group.
	CacheSecurityGroupNames *[]*string `field:"optional" json:"cacheSecurityGroupNames" yaml:"cacheSecurityGroupNames"`
	// The name of the cache subnet group to be used for the replication group.
	//
	// > If you're going to launch your cluster in an Amazon VPC, you need to create a subnet group before you start creating a cluster. For more information, see [AWS::ElastiCache::SubnetGroup](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-subnetgroup.html) .
	CacheSubnetGroupName *string `field:"optional" json:"cacheSubnetGroupName" yaml:"cacheSubnetGroupName"`
	// Enables data tiering.
	//
	// Data tiering is only supported for replication groups using the r6gd node type. This parameter must be set to true when using r6gd nodes. For more information, see [Data tiering](https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/data-tiering.html) .
	DataTieringEnabled interface{} `field:"optional" json:"dataTieringEnabled" yaml:"dataTieringEnabled"`
	// The name of the cache engine to be used for the clusters in this replication group.
	//
	// Must be Redis.
	Engine *string `field:"optional" json:"engine" yaml:"engine"`
	// The version number of the cache engine to be used for the clusters in this replication group.
	//
	// To view the supported cache engine versions, use the `DescribeCacheEngineVersions` operation.
	//
	// *Important:* You can upgrade to a newer engine version (see [Selecting a Cache Engine and Version](https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/SelectEngine.html#VersionManagement) ) in the *ElastiCache User Guide* , but you cannot downgrade to an earlier engine version. If you want to use an earlier engine version, you must delete the existing cluster or replication group and create it anew with the earlier engine version.
	EngineVersion *string `field:"optional" json:"engineVersion" yaml:"engineVersion"`
	// The name of the Global datastore.
	GlobalReplicationGroupId *string `field:"optional" json:"globalReplicationGroupId" yaml:"globalReplicationGroupId"`
	// `AWS::ElastiCache::ReplicationGroup.IpDiscovery`.
	IpDiscovery *string `field:"optional" json:"ipDiscovery" yaml:"ipDiscovery"`
	// The ID of the KMS key used to encrypt the disk on the cluster.
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// Specifies the destination, format and type of the logs.
	LogDeliveryConfigurations interface{} `field:"optional" json:"logDeliveryConfigurations" yaml:"logDeliveryConfigurations"`
	// A flag indicating if you have Multi-AZ enabled to enhance fault tolerance.
	//
	// For more information, see [Minimizing Downtime: Multi-AZ](https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/AutoFailover.html) .
	MultiAzEnabled interface{} `field:"optional" json:"multiAzEnabled" yaml:"multiAzEnabled"`
	// `AWS::ElastiCache::ReplicationGroup.NetworkType`.
	NetworkType *string `field:"optional" json:"networkType" yaml:"networkType"`
	// `NodeGroupConfiguration` is a property of the `AWS::ElastiCache::ReplicationGroup` resource that configures an Amazon ElastiCache (ElastiCache) Redis cluster node group.
	//
	// If you set [UseOnlineResharding](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-updatepolicy.html#cfn-attributes-updatepolicy-useonlineresharding) to `true` , you can update `NodeGroupConfiguration` without interruption. When `UseOnlineResharding` is set to `false` , or is not specified, updating `NodeGroupConfiguration` results in [replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement) .
	NodeGroupConfiguration interface{} `field:"optional" json:"nodeGroupConfiguration" yaml:"nodeGroupConfiguration"`
	// The Amazon Resource Name (ARN) of the Amazon Simple Notification Service (SNS) topic to which notifications are sent.
	//
	// > The Amazon SNS topic owner must be the same as the cluster owner.
	NotificationTopicArn *string `field:"optional" json:"notificationTopicArn" yaml:"notificationTopicArn"`
	// The number of clusters this replication group initially has.
	//
	// This parameter is not used if there is more than one node group (shard). You should use `ReplicasPerNodeGroup` instead.
	//
	// If `AutomaticFailoverEnabled` is `true` , the value of this parameter must be at least 2. If `AutomaticFailoverEnabled` is `false` you can omit this parameter (it will default to 1), or you can explicitly set it to a value between 2 and 6.
	//
	// The maximum permitted value for `NumCacheClusters` is 6 (1 primary plus 5 replicas).
	NumCacheClusters *float64 `field:"optional" json:"numCacheClusters" yaml:"numCacheClusters"`
	// An optional parameter that specifies the number of node groups (shards) for this Redis (cluster mode enabled) replication group.
	//
	// For Redis (cluster mode disabled) either omit this parameter or set it to 1.
	//
	// If you set [UseOnlineResharding](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-updatepolicy.html#cfn-attributes-updatepolicy-useonlineresharding) to `true` , you can update `NumNodeGroups` without interruption. When `UseOnlineResharding` is set to `false` , or is not specified, updating `NumNodeGroups` results in [replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement) .
	//
	// Default: 1.
	NumNodeGroups *float64 `field:"optional" json:"numNodeGroups" yaml:"numNodeGroups"`
	// The port number on which each member of the replication group accepts connections.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// A list of EC2 Availability Zones in which the replication group's clusters are created.
	//
	// The order of the Availability Zones in the list is the order in which clusters are allocated. The primary cluster is created in the first AZ in the list.
	//
	// This parameter is not used if there is more than one node group (shard). You should use `NodeGroupConfiguration` instead.
	//
	// > If you are creating your replication group in an Amazon VPC (recommended), you can only locate clusters in Availability Zones associated with the subnets in the selected subnet group.
	// >
	// > The number of Availability Zones listed must equal the value of `NumCacheClusters` .
	//
	// Default: system chosen Availability Zones.
	PreferredCacheClusterAZs *[]*string `field:"optional" json:"preferredCacheClusterAZs" yaml:"preferredCacheClusterAZs"`
	// Specifies the weekly time range during which maintenance on the cluster is performed.
	//
	// It is specified as a range in the format ddd:hh24:mi-ddd:hh24:mi (24H Clock UTC). The minimum maintenance window is a 60 minute period.
	//
	// Valid values for `ddd` are:
	//
	// - `sun`
	// - `mon`
	// - `tue`
	// - `wed`
	// - `thu`
	// - `fri`
	// - `sat`
	//
	// Example: `sun:23:00-mon:01:30`.
	PreferredMaintenanceWindow *string `field:"optional" json:"preferredMaintenanceWindow" yaml:"preferredMaintenanceWindow"`
	// The identifier of the cluster that serves as the primary for this replication group.
	//
	// This cluster must already exist and have a status of `available` .
	//
	// This parameter is not required if `NumCacheClusters` , `NumNodeGroups` , or `ReplicasPerNodeGroup` is specified.
	PrimaryClusterId *string `field:"optional" json:"primaryClusterId" yaml:"primaryClusterId"`
	// An optional parameter that specifies the number of replica nodes in each node group (shard).
	//
	// Valid values are 0 to 5.
	ReplicasPerNodeGroup *float64 `field:"optional" json:"replicasPerNodeGroup" yaml:"replicasPerNodeGroup"`
	// The replication group identifier. This parameter is stored as a lowercase string.
	//
	// Constraints:
	//
	// - A name must contain from 1 to 40 alphanumeric characters or hyphens.
	// - The first character must be a letter.
	// - A name cannot end with a hyphen or contain two consecutive hyphens.
	ReplicationGroupId *string `field:"optional" json:"replicationGroupId" yaml:"replicationGroupId"`
	// One or more Amazon VPC security groups associated with this replication group.
	//
	// Use this parameter only when you are creating a replication group in an Amazon Virtual Private Cloud (Amazon VPC).
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	// A list of Amazon Resource Names (ARN) that uniquely identify the Redis RDB snapshot files stored in Amazon S3.
	//
	// The snapshot files are used to populate the new replication group. The Amazon S3 object name in the ARN cannot contain any commas. The new replication group will have the number of node groups (console: shards) specified by the parameter *NumNodeGroups* or the number of node groups configured by *NodeGroupConfiguration* regardless of the number of ARNs specified here.
	//
	// Example of an Amazon S3 ARN: `arn:aws:s3:::my_bucket/snapshot1.rdb`
	SnapshotArns *[]*string `field:"optional" json:"snapshotArns" yaml:"snapshotArns"`
	// The name of a snapshot from which to restore data into the new replication group.
	//
	// The snapshot status changes to `restoring` while the new replication group is being created.
	SnapshotName *string `field:"optional" json:"snapshotName" yaml:"snapshotName"`
	// The number of days for which ElastiCache retains automatic snapshots before deleting them.
	//
	// For example, if you set `SnapshotRetentionLimit` to 5, a snapshot that was taken today is retained for 5 days before being deleted.
	//
	// Default: 0 (i.e., automatic backups are disabled for this cluster).
	SnapshotRetentionLimit *float64 `field:"optional" json:"snapshotRetentionLimit" yaml:"snapshotRetentionLimit"`
	// The cluster ID that is used as the daily snapshot source for the replication group.
	//
	// This parameter cannot be set for Redis (cluster mode enabled) replication groups.
	SnapshottingClusterId *string `field:"optional" json:"snapshottingClusterId" yaml:"snapshottingClusterId"`
	// The daily time range (in UTC) during which ElastiCache begins taking a daily snapshot of your node group (shard).
	//
	// Example: `05:00-09:00`
	//
	// If you do not specify this parameter, ElastiCache automatically chooses an appropriate time range.
	SnapshotWindow *string `field:"optional" json:"snapshotWindow" yaml:"snapshotWindow"`
	// A list of tags to be added to this resource.
	//
	// Tags are comma-separated key,value pairs (e.g. Key= `myKey` , Value= `myKeyValue` . You can include multiple tags as shown following: Key= `myKey` , Value= `myKeyValue` Key= `mySecondKey` , Value= `mySecondKeyValue` . Tags on replication groups will be replicated to all nodes.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// A flag that enables in-transit encryption when set to `true` .
	//
	// You cannot modify the value of `TransitEncryptionEnabled` after the cluster is created. To enable in-transit encryption on a cluster you must set `TransitEncryptionEnabled` to `true` when you create a cluster.
	//
	// This parameter is valid only if the `Engine` parameter is `redis` , the `EngineVersion` parameter is `3.2.6` or `4.x` onward, and the cluster is being created in an Amazon VPC.
	//
	// If you enable in-transit encryption, you must also specify a value for `CacheSubnetGroup` .
	//
	// *Required:* Only available when creating a replication group in an Amazon VPC using redis version `3.2.6` or `4.x` onward.
	//
	// Default: `false`
	//
	// > For HIPAA compliance, you must specify `TransitEncryptionEnabled` as `true` , an `AuthToken` , and a `CacheSubnetGroup` .
	TransitEncryptionEnabled interface{} `field:"optional" json:"transitEncryptionEnabled" yaml:"transitEncryptionEnabled"`
	// The list of user groups to associate with the replication group.
	UserGroupIds *[]*string `field:"optional" json:"userGroupIds" yaml:"userGroupIds"`
}

