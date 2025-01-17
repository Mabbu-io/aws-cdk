package pipelines


// Properties for a `Pipeline`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var fileSetProducer iFileSetProducer
//
//   pipelineBaseProps := &pipelineBaseProps{
//   	synth: fileSetProducer,
//   }
//
type PipelineBaseProps struct {
	// The build step that produces the CDK Cloud Assembly.
	//
	// The primary output of this step needs to be the `cdk.out` directory
	// generated by the `cdk synth` command.
	//
	// If you use a `ShellStep` here and you don't configure an output directory,
	// the output directory will automatically be assumed to be `cdk.out`.
	Synth IFileSetProducer `field:"required" json:"synth" yaml:"synth"`
}

