package awskendra


// Provides the configuration information for crawling knowledge articles in the ServiceNow site.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   serviceNowKnowledgeArticleConfigurationProperty := &serviceNowKnowledgeArticleConfigurationProperty{
//   	documentDataFieldName: jsii.String("documentDataFieldName"),
//
//   	// the properties below are optional
//   	crawlAttachments: jsii.Boolean(false),
//   	documentTitleFieldName: jsii.String("documentTitleFieldName"),
//   	excludeAttachmentFilePatterns: []*string{
//   		jsii.String("excludeAttachmentFilePatterns"),
//   	},
//   	fieldMappings: []interface{}{
//   		&dataSourceToIndexFieldMappingProperty{
//   			dataSourceFieldName: jsii.String("dataSourceFieldName"),
//   			indexFieldName: jsii.String("indexFieldName"),
//
//   			// the properties below are optional
//   			dateFieldFormat: jsii.String("dateFieldFormat"),
//   		},
//   	},
//   	filterQuery: jsii.String("filterQuery"),
//   	includeAttachmentFilePatterns: []*string{
//   		jsii.String("includeAttachmentFilePatterns"),
//   	},
//   }
//
type CfnDataSource_ServiceNowKnowledgeArticleConfigurationProperty struct {
	// The name of the ServiceNow field that is mapped to the index document contents field in the Amazon Kendra index.
	DocumentDataFieldName *string `field:"required" json:"documentDataFieldName" yaml:"documentDataFieldName"`
	// Indicates whether Amazon Kendra should index attachments to knowledge articles.
	CrawlAttachments interface{} `field:"optional" json:"crawlAttachments" yaml:"crawlAttachments"`
	// The name of the ServiceNow field that is mapped to the index document title field.
	DocumentTitleFieldName *string `field:"optional" json:"documentTitleFieldName" yaml:"documentTitleFieldName"`
	// A list of regular expression patterns to exclude certain attachments of knowledge articles in your ServiceNow.
	//
	// Item that match the patterns are excluded from the index. Items that don't match the patterns are included in the index. If an item matches both an inclusion and exclusion pattern, the exclusion pattern takes precedence and the item isn't included in the index.
	//
	// The regex is applied to the field specified in the `PatternTargetField` .
	ExcludeAttachmentFilePatterns *[]*string `field:"optional" json:"excludeAttachmentFilePatterns" yaml:"excludeAttachmentFilePatterns"`
	// Maps attributes or field names of knoweldge articles to Amazon Kendra index field names.
	//
	// To create custom fields, use the `UpdateIndex` API before you map to ServiceNow fields. For more information, see [Mapping data source fields](https://docs.aws.amazon.com/kendra/latest/dg/field-mapping.html) . The ServiceNow data source field names must exist in your ServiceNow custom metadata.
	FieldMappings interface{} `field:"optional" json:"fieldMappings" yaml:"fieldMappings"`
	// A query that selects the knowledge articles to index.
	//
	// The query can return articles from multiple knowledge bases, and the knowledge bases can be public or private.
	//
	// The query string must be one generated by the ServiceNow console. For more information, see [Specifying documents to index with a query](https://docs.aws.amazon.com/kendra/latest/dg/servicenow-query.html) .
	FilterQuery *string `field:"optional" json:"filterQuery" yaml:"filterQuery"`
	// A list of regular expression patterns to include certain attachments of knowledge articles in your ServiceNow.
	//
	// Item that match the patterns are included in the index. Items that don't match the patterns are excluded from the index. If an item matches both an inclusion and exclusion pattern, the exclusion pattern takes precedence and the item isn't included in the index.
	//
	// The regex is applied to the field specified in the `PatternTargetField` .
	IncludeAttachmentFilePatterns *[]*string `field:"optional" json:"includeAttachmentFilePatterns" yaml:"includeAttachmentFilePatterns"`
}

