package awscloudwatch


// Statistic to use over the aggregation period.
// See: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/Statistics-definitions.html
//
// Deprecated: Use one of the factory methods on `Stats` to produce statistics strings.
type Statistic string

const (
	// The count (number) of data points used for the statistical calculation.
	// Deprecated: Use one of the factory methods on `Stats` to produce statistics strings.
	Statistic_SAMPLE_COUNT Statistic = "SAMPLE_COUNT"
	// The value of Sum / SampleCount during the specified period.
	// Deprecated: Use one of the factory methods on `Stats` to produce statistics strings.
	Statistic_AVERAGE Statistic = "AVERAGE"
	// All values submitted for the matching metric added together.
	//
	// This statistic can be useful for determining the total volume of a metric.
	// Deprecated: Use one of the factory methods on `Stats` to produce statistics strings.
	Statistic_SUM Statistic = "SUM"
	// The lowest value observed during the specified period.
	//
	// You can use this value to determine low volumes of activity for your application.
	// Deprecated: Use one of the factory methods on `Stats` to produce statistics strings.
	Statistic_MINIMUM Statistic = "MINIMUM"
	// The highest value observed during the specified period.
	//
	// You can use this value to determine high volumes of activity for your application.
	// Deprecated: Use one of the factory methods on `Stats` to produce statistics strings.
	Statistic_MAXIMUM Statistic = "MAXIMUM"
)

