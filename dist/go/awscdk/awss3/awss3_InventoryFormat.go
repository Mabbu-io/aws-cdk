package awss3


// All supported inventory list formats.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var bucket bucket
//
//   inventory := &inventory{
//   	destination: &inventoryDestination{
//   		bucket: bucket,
//
//   		// the properties below are optional
//   		bucketOwner: jsii.String("bucketOwner"),
//   		prefix: jsii.String("prefix"),
//   	},
//
//   	// the properties below are optional
//   	enabled: jsii.Boolean(false),
//   	format: awscdk.Aws_s3.inventoryFormat_CSV,
//   	frequency: awscdk.*Aws_s3.inventoryFrequency_DAILY,
//   	includeObjectVersions: awscdk.*Aws_s3.inventoryObjectVersion_ALL,
//   	inventoryId: jsii.String("inventoryId"),
//   	objectsPrefix: jsii.String("objectsPrefix"),
//   	optionalFields: []*string{
//   		jsii.String("optionalFields"),
//   	},
//   }
//
type InventoryFormat string

const (
	// Generate the inventory list as CSV.
	InventoryFormat_CSV InventoryFormat = "CSV"
	// Generate the inventory list as Parquet.
	InventoryFormat_PARQUET InventoryFormat = "PARQUET"
	// Generate the inventory list as ORC.
	InventoryFormat_ORC InventoryFormat = "ORC"
)

