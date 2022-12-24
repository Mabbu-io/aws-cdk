package awsdynamodb


// The set of attributes that are projected into the index.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   localSecondaryIndexProps := &localSecondaryIndexProps{
//   	indexName: jsii.String("indexName"),
//   	sortKey: &attribute{
//   		name: jsii.String("name"),
//   		type: awscdk.Aws_dynamodb.attributeType_BINARY,
//   	},
//
//   	// the properties below are optional
//   	nonKeyAttributes: []*string{
//   		jsii.String("nonKeyAttributes"),
//   	},
//   	projectionType: awscdk.*Aws_dynamodb.projectionType_KEYS_ONLY,
//   }
//
// See: https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_Projection.html
//
type ProjectionType string

const (
	// Only the index and primary keys are projected into the index.
	ProjectionType_KEYS_ONLY ProjectionType = "KEYS_ONLY"
	// Only the specified table attributes are projected into the index.
	//
	// The list of projected attributes is in `nonKeyAttributes`.
	ProjectionType_INCLUDE ProjectionType = "INCLUDE"
	// All of the table attributes are projected into the index.
	ProjectionType_ALL ProjectionType = "ALL"
)

