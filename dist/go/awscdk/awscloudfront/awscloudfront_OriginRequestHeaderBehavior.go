package awscloudfront

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Determines whether any HTTP headers (and if so, which headers) are included in requests that CloudFront sends to the origin.
//
// Example:
//   // Creating a custom origin request policy for a Distribution -- all parameters optional
//   var bucketOrigin s3Origin
//
//   myOriginRequestPolicy := cloudfront.NewOriginRequestPolicy(this, jsii.String("OriginRequestPolicy"), &originRequestPolicyProps{
//   	originRequestPolicyName: jsii.String("MyPolicy"),
//   	comment: jsii.String("A default policy"),
//   	cookieBehavior: cloudfront.originRequestCookieBehavior.none(),
//   	headerBehavior: cloudfront.originRequestHeaderBehavior.all(jsii.String("CloudFront-Is-Android-Viewer")),
//   	queryStringBehavior: cloudfront.originRequestQueryStringBehavior.allowList(jsii.String("username")),
//   })
//
//   cloudfront.NewDistribution(this, jsii.String("myDistCustomPolicy"), &distributionProps{
//   	defaultBehavior: &behaviorOptions{
//   		origin: bucketOrigin,
//   		originRequestPolicy: myOriginRequestPolicy,
//   	},
//   })
//
type OriginRequestHeaderBehavior interface {
	// The behavior of headers: allow all, none or an allow list.
	Behavior() *string
	// The headers for the allow list or the included CloudFront headers, if applicable.
	Headers() *[]*string
}

// The jsii proxy struct for OriginRequestHeaderBehavior
type jsiiProxy_OriginRequestHeaderBehavior struct {
	_ byte // padding
}

func (j *jsiiProxy_OriginRequestHeaderBehavior) Behavior() *string {
	var returns *string
	_jsii_.Get(
		j,
		"behavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OriginRequestHeaderBehavior) Headers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"headers",
		&returns,
	)
	return returns
}


// All HTTP headers in viewer requests are included in requests that CloudFront sends to the origin.
//
// Additionally, any additional CloudFront headers provided are included; the additional headers are added by CloudFront.
// See: https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/using-cloudfront-headers.html
//
func OriginRequestHeaderBehavior_All(cloudfrontHeaders ...*string) OriginRequestHeaderBehavior {
	_init_.Initialize()

	args := []interface{}{}
	for _, a := range cloudfrontHeaders {
		args = append(args, a)
	}

	var returns OriginRequestHeaderBehavior

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudfront.OriginRequestHeaderBehavior",
		"all",
		args,
		&returns,
	)

	return returns
}

// Listed headers are included in requests that CloudFront sends to the origin.
func OriginRequestHeaderBehavior_AllowList(headers ...*string) OriginRequestHeaderBehavior {
	_init_.Initialize()

	args := []interface{}{}
	for _, a := range headers {
		args = append(args, a)
	}

	var returns OriginRequestHeaderBehavior

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudfront.OriginRequestHeaderBehavior",
		"allowList",
		args,
		&returns,
	)

	return returns
}

// HTTP headers are not included in requests that CloudFront sends to the origin.
//
// Any headers that are listed in a CachePolicy are still included in origin requests.
func OriginRequestHeaderBehavior_None() OriginRequestHeaderBehavior {
	_init_.Initialize()

	var returns OriginRequestHeaderBehavior

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudfront.OriginRequestHeaderBehavior",
		"none",
		nil, // no parameters
		&returns,
	)

	return returns
}

