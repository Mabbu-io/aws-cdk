package awscloudfront


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var bucket bucket
//   var function_ function
//   var keyGroup keyGroup
//   var originAccessIdentity originAccessIdentity
//   var version version
//
//   sourceConfiguration := &sourceConfiguration{
//   	behaviors: []behavior{
//   		&behavior{
//   			allowedMethods: awscdk.Aws_cloudfront.cloudFrontAllowedMethods_GET_HEAD,
//   			cachedMethods: awscdk.*Aws_cloudfront.cloudFrontAllowedCachedMethods_GET_HEAD,
//   			compress: jsii.Boolean(false),
//   			defaultTtl: cdk.duration.minutes(jsii.Number(30)),
//   			forwardedValues: &forwardedValuesProperty{
//   				queryString: jsii.Boolean(false),
//
//   				// the properties below are optional
//   				cookies: &cookiesProperty{
//   					forward: jsii.String("forward"),
//
//   					// the properties below are optional
//   					whitelistedNames: []*string{
//   						jsii.String("whitelistedNames"),
//   					},
//   				},
//   				headers: []*string{
//   					jsii.String("headers"),
//   				},
//   				queryStringCacheKeys: []*string{
//   					jsii.String("queryStringCacheKeys"),
//   				},
//   			},
//   			functionAssociations: []functionAssociation{
//   				&functionAssociation{
//   					eventType: awscdk.*Aws_cloudfront.functionEventType_VIEWER_REQUEST,
//   					function: function_,
//   				},
//   			},
//   			isDefaultBehavior: jsii.Boolean(false),
//   			lambdaFunctionAssociations: []lambdaFunctionAssociation{
//   				&lambdaFunctionAssociation{
//   					eventType: awscdk.*Aws_cloudfront.lambdaEdgeEventType_ORIGIN_REQUEST,
//   					lambdaFunction: version,
//
//   					// the properties below are optional
//   					includeBody: jsii.Boolean(false),
//   				},
//   			},
//   			maxTtl: cdk.*duration.minutes(jsii.Number(30)),
//   			minTtl: cdk.*duration.minutes(jsii.Number(30)),
//   			pathPattern: jsii.String("pathPattern"),
//   			trustedKeyGroups: []iKeyGroup{
//   				keyGroup,
//   			},
//   			trustedSigners: []*string{
//   				jsii.String("trustedSigners"),
//   			},
//   			viewerProtocolPolicy: awscdk.*Aws_cloudfront.viewerProtocolPolicy_HTTPS_ONLY,
//   		},
//   	},
//
//   	// the properties below are optional
//   	connectionAttempts: jsii.Number(123),
//   	connectionTimeout: cdk.*duration.minutes(jsii.Number(30)),
//   	customOriginSource: &customOriginConfig{
//   		domainName: jsii.String("domainName"),
//
//   		// the properties below are optional
//   		allowedOriginSSLVersions: []originSslPolicy{
//   			awscdk.*Aws_cloudfront.*originSslPolicy_SSL_V3,
//   		},
//   		httpPort: jsii.Number(123),
//   		httpsPort: jsii.Number(123),
//   		originHeaders: map[string]*string{
//   			"originHeadersKey": jsii.String("originHeaders"),
//   		},
//   		originKeepaliveTimeout: cdk.*duration.minutes(jsii.Number(30)),
//   		originPath: jsii.String("originPath"),
//   		originProtocolPolicy: awscdk.*Aws_cloudfront.originProtocolPolicy_HTTP_ONLY,
//   		originReadTimeout: cdk.*duration.minutes(jsii.Number(30)),
//   		originShieldRegion: jsii.String("originShieldRegion"),
//   	},
//   	failoverCriteriaStatusCodes: []failoverStatusCode{
//   		awscdk.*Aws_cloudfront.*failoverStatusCode_FORBIDDEN,
//   	},
//   	failoverCustomOriginSource: &customOriginConfig{
//   		domainName: jsii.String("domainName"),
//
//   		// the properties below are optional
//   		allowedOriginSSLVersions: []*originSslPolicy{
//   			awscdk.*Aws_cloudfront.*originSslPolicy_SSL_V3,
//   		},
//   		httpPort: jsii.Number(123),
//   		httpsPort: jsii.Number(123),
//   		originHeaders: map[string]*string{
//   			"originHeadersKey": jsii.String("originHeaders"),
//   		},
//   		originKeepaliveTimeout: cdk.*duration.minutes(jsii.Number(30)),
//   		originPath: jsii.String("originPath"),
//   		originProtocolPolicy: awscdk.*Aws_cloudfront.*originProtocolPolicy_HTTP_ONLY,
//   		originReadTimeout: cdk.*duration.minutes(jsii.Number(30)),
//   		originShieldRegion: jsii.String("originShieldRegion"),
//   	},
//   	failoverS3OriginSource: &s3OriginConfig{
//   		s3BucketSource: bucket,
//
//   		// the properties below are optional
//   		originAccessIdentity: originAccessIdentity,
//   		originHeaders: map[string]*string{
//   			"originHeadersKey": jsii.String("originHeaders"),
//   		},
//   		originPath: jsii.String("originPath"),
//   		originShieldRegion: jsii.String("originShieldRegion"),
//   	},
//   	originShieldRegion: jsii.String("originShieldRegion"),
//   	s3OriginSource: &s3OriginConfig{
//   		s3BucketSource: bucket,
//
//   		// the properties below are optional
//   		originAccessIdentity: originAccessIdentity,
//   		originHeaders: map[string]*string{
//   			"originHeadersKey": jsii.String("originHeaders"),
//   		},
//   		originPath: jsii.String("originPath"),
//   		originShieldRegion: jsii.String("originShieldRegion"),
//   	},
//   }
//
type OriginSslPolicy string

const (
	OriginSslPolicy_SSL_V3 OriginSslPolicy = "SSL_V3"
	OriginSslPolicy_TLS_V1 OriginSslPolicy = "TLS_V1"
	OriginSslPolicy_TLS_V1_1 OriginSslPolicy = "TLS_V1_1"
	OriginSslPolicy_TLS_V1_2 OriginSslPolicy = "TLS_V1_2"
)

