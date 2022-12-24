package awssecretsmanager


// The type of service or database that's being associated with the secret.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   secretAttachmentTargetProps := &secretAttachmentTargetProps{
//   	targetId: jsii.String("targetId"),
//   	targetType: awscdk.Aws_secretsmanager.attachmentTargetType_RDS_DB_INSTANCE,
//   }
//
type AttachmentTargetType string

const (
	// AWS::RDS::DBInstance.
	AttachmentTargetType_RDS_DB_INSTANCE AttachmentTargetType = "RDS_DB_INSTANCE"
	// AWS::RDS::DBCluster.
	AttachmentTargetType_RDS_DB_CLUSTER AttachmentTargetType = "RDS_DB_CLUSTER"
	// AWS::RDS::DBProxy.
	AttachmentTargetType_RDS_DB_PROXY AttachmentTargetType = "RDS_DB_PROXY"
	// AWS::Redshift::Cluster.
	AttachmentTargetType_REDSHIFT_CLUSTER AttachmentTargetType = "REDSHIFT_CLUSTER"
	// AWS::DocDB::DBInstance.
	AttachmentTargetType_DOCDB_DB_INSTANCE AttachmentTargetType = "DOCDB_DB_INSTANCE"
	// AWS::DocDB::DBCluster.
	AttachmentTargetType_DOCDB_DB_CLUSTER AttachmentTargetType = "DOCDB_DB_CLUSTER"
)

