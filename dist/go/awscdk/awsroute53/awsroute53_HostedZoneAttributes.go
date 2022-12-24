package awsroute53


// Reference to a hosted zone.
//
// Example:
//   zone := route53.hostedZone.fromHostedZoneAttributes(this, jsii.String("MyZone"), &hostedZoneAttributes{
//   	zoneName: jsii.String("example.com"),
//   	hostedZoneId: jsii.String("ZOJJZC49E0EPZ"),
//   })
//
type HostedZoneAttributes struct {
	// Identifier of the hosted zone.
	HostedZoneId *string `field:"required" json:"hostedZoneId" yaml:"hostedZoneId"`
	// Name of the hosted zone.
	ZoneName *string `field:"required" json:"zoneName" yaml:"zoneName"`
}

