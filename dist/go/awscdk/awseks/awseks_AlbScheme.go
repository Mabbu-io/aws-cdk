package awseks


// ALB Scheme.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   kubernetesManifestOptions := &kubernetesManifestOptions{
//   	ingressAlb: jsii.Boolean(false),
//   	ingressAlbScheme: awscdk.Aws_eks.albScheme_INTERNAL,
//   	prune: jsii.Boolean(false),
//   	skipValidation: jsii.Boolean(false),
//   }
//
// See: https://kubernetes-sigs.github.io/aws-load-balancer-controller/v2.3/guide/ingress/annotations/#scheme
//
type AlbScheme string

const (
	// The nodes of an internal load balancer have only private IP addresses.
	//
	// The DNS name of an internal load balancer is publicly resolvable to the private IP addresses of the nodes.
	// Therefore, internal load balancers can only route requests from clients with access to the VPC for the load balancer.
	AlbScheme_INTERNAL AlbScheme = "INTERNAL"
	// An internet-facing load balancer has a publicly resolvable DNS name, so it can route requests from clients over the internet to the EC2 instances that are registered with the load balancer.
	AlbScheme_INTERNET_FACING AlbScheme = "INTERNET_FACING"
)

