package awscodecommit

import (
	"github.com/aws/aws-cdk-go/awscdk/awsevents"
	"github.com/aws/constructs-go/constructs/v10"
)

// Options for the onCommit() method.
//
// Example:
//   import codecommit "github.com/aws/aws-cdk-go/awscdk"
//   import targets "github.com/aws/aws-cdk-go/awscdk"
//
//   var repo repository
//
//   myTopic := sns.NewTopic(this, jsii.String("Topic"))
//
//   repo.onCommit(jsii.String("OnCommit"), &onCommitOptions{
//   	target: targets.NewSnsTopic(myTopic),
//   })
//
type OnCommitOptions struct {
	// The scope to use if the source of the rule and its target are in different Stacks (but in the same account & region).
	//
	// This helps dealing with cycles that often arise in these situations.
	CrossStackScope constructs.Construct `field:"optional" json:"crossStackScope" yaml:"crossStackScope"`
	// A description of the rule's purpose.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Additional restrictions for the event to route to the specified target.
	//
	// The method that generates the rule probably imposes some type of event
	// filtering. The filtering implied by what you pass here is added
	// on top of that filtering.
	// See: https://docs.aws.amazon.com/eventbridge/latest/userguide/eventbridge-and-event-patterns.html
	//
	EventPattern *awsevents.EventPattern `field:"optional" json:"eventPattern" yaml:"eventPattern"`
	// A name for the rule.
	RuleName *string `field:"optional" json:"ruleName" yaml:"ruleName"`
	// The target to register for the event.
	Target awsevents.IRuleTarget `field:"optional" json:"target" yaml:"target"`
	// The branch to monitor.
	Branches *[]*string `field:"optional" json:"branches" yaml:"branches"`
}

