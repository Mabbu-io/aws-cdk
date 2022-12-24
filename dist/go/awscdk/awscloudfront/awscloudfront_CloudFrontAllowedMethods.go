package awscloudfront


// An enum for the supported methods to a CloudFront distribution.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var function_ function
//   var keyGroup keyGroup
//   var version version
//
//   behavior := &behavior{
//   	allowedMethods: awscdk.Aws_cloudfront.cloudFrontAllowedMethods_GET_HEAD,
//   	cachedMethods: awscdk.*Aws_cloudfront.cloudFrontAllowedCachedMethods_GET_HEAD,
//   	compress: jsii.Boolean(false),
//   	defaultTtl: cdk.duration.minutes(jsii.Number(30)),
//   	forwardedValues: &forwardedValuesProperty{
//   		queryString: jsii.Boolean(false),
//
//   		// the properties below are optional
//   		cookies: &cookiesProperty{
//   			forward: jsii.String("forward"),
//
//   			// the properties below are optional
//   			whitelistedNames: []*string{
//   				jsii.String("whitelistedNames"),
//   			},
//   		},
//   		headers: []*string{
//   			jsii.String("headers"),
//   		},
//   		queryStringCacheKeys: []*string{
//   			jsii.String("queryStringCacheKeys"),
//   		},
//   	},
//   	functionAssociations: []functionAssociation{
//   		&functionAssociation{
//   			eventType: awscdk.*Aws_cloudfront.functionEventType_VIEWER_REQUEST,
//   			function: function_,
//   		},
//   	},
//   	isDefaultBehavior: jsii.Boolean(false),
//   	lambdaFunctionAssociations: []lambdaFunctionAssociation{
//   		&lambdaFunctionAssociation{
//   			eventType: awscdk.*Aws_cloudfront.lambdaEdgeEventType_ORIGIN_REQUEST,
//   			lambdaFunction: version,
//
//   			// the properties below are optional
//   			includeBody: jsii.Boolean(false),
//   		},
//   	},
//   	maxTtl: cdk.*duration.minutes(jsii.Number(30)),
//   	minTtl: cdk.*duration.minutes(jsii.Number(30)),
//   	pathPattern: jsii.String("pathPattern"),
//   	trustedKeyGroups: []iKeyGroup{
//   		keyGroup,
//   	},
//   	trustedSigners: []*string{
//   		jsii.String("trustedSigners"),
//   	},
//   	viewerProtocolPolicy: awscdk.*Aws_cloudfront.viewerProtocolPolicy_HTTPS_ONLY,
//   }
//
type CloudFrontAllowedMethods string

const (
	CloudFrontAllowedMethods_GET_HEAD CloudFrontAllowedMethods = "GET_HEAD"
	CloudFrontAllowedMethods_GET_HEAD_OPTIONS CloudFrontAllowedMethods = "GET_HEAD_OPTIONS"
	CloudFrontAllowedMethods_ALL CloudFrontAllowedMethods = "ALL"
)

