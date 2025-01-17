package awscloudfront

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// A CloudFront behavior wrapper.
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
type Behavior struct {
	// The method this CloudFront distribution responds do.
	AllowedMethods CloudFrontAllowedMethods `field:"optional" json:"allowedMethods" yaml:"allowedMethods"`
	// Which methods are cached by CloudFront by default.
	CachedMethods CloudFrontAllowedCachedMethods `field:"optional" json:"cachedMethods" yaml:"cachedMethods"`
	// If CloudFront should automatically compress some content types.
	Compress *bool `field:"optional" json:"compress" yaml:"compress"`
	// The default amount of time CloudFront will cache an object.
	//
	// This value applies only when your custom origin does not add HTTP headers,
	// such as Cache-Control max-age, Cache-Control s-maxage, and Expires to objects.
	DefaultTtl awscdk.Duration `field:"optional" json:"defaultTtl" yaml:"defaultTtl"`
	// The values CloudFront will forward to the origin when making a request.
	ForwardedValues *CfnDistribution_ForwardedValuesProperty `field:"optional" json:"forwardedValues" yaml:"forwardedValues"`
	// The CloudFront functions to invoke before serving the contents.
	FunctionAssociations *[]*FunctionAssociation `field:"optional" json:"functionAssociations" yaml:"functionAssociations"`
	// If this behavior is the default behavior for the distribution.
	//
	// You must specify exactly one default distribution per CloudFront distribution.
	// The default behavior is allowed to omit the "path" property.
	IsDefaultBehavior *bool `field:"optional" json:"isDefaultBehavior" yaml:"isDefaultBehavior"`
	// Declares associated lambda@edge functions for this distribution behaviour.
	LambdaFunctionAssociations *[]*LambdaFunctionAssociation `field:"optional" json:"lambdaFunctionAssociations" yaml:"lambdaFunctionAssociations"`
	// The max amount of time you want objects to stay in the cache before CloudFront queries your origin.
	MaxTtl awscdk.Duration `field:"optional" json:"maxTtl" yaml:"maxTtl"`
	// The minimum amount of time that you want objects to stay in the cache before CloudFront queries your origin.
	MinTtl awscdk.Duration `field:"optional" json:"minTtl" yaml:"minTtl"`
	// The path this behavior responds to.
	//
	// Required for all non-default behaviors. (The default behavior implicitly has "*" as the path pattern. )
	PathPattern *string `field:"optional" json:"pathPattern" yaml:"pathPattern"`
	// A list of Key Groups that CloudFront can use to validate signed URLs or signed cookies.
	// See: https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/PrivateContent.html
	//
	TrustedKeyGroups *[]IKeyGroup `field:"optional" json:"trustedKeyGroups" yaml:"trustedKeyGroups"`
	// Trusted signers is how CloudFront allows you to serve private content.
	//
	// The signers are the account IDs that are allowed to sign cookies/presigned URLs for this distribution.
	//
	// If you pass a non empty value, all requests for this behavior must be signed (no public access will be allowed).
	// Deprecated: - We recommend using trustedKeyGroups instead of trustedSigners.
	TrustedSigners *[]*string `field:"optional" json:"trustedSigners" yaml:"trustedSigners"`
	// The viewer policy for this behavior.
	ViewerProtocolPolicy ViewerProtocolPolicy `field:"optional" json:"viewerProtocolPolicy" yaml:"viewerProtocolPolicy"`
}

