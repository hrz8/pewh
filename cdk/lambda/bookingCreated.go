package lambda

import (
	"strings"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3assets"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssns"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssnssubscriptions"
	"github.com/aws/jsii-runtime-go"
)

func NewBookingCreatedLambda(stack awscdk.Stack, stage string) awslambda.Function {
	functionID := jsii.String("BookingCreated")
	topicID := jsii.String("BookingCreatedTopic")
	topicName := jsii.String("booking-created-" + strings.ToLower(stage))

	function := awslambda.NewFunction(stack, functionID, &awslambda.FunctionProps{
		Runtime:      awslambda.Runtime_PROVIDED_AL2(),
		Architecture: awslambda.Architecture_ARM_64(),
		MemorySize:   jsii.Number(128),
		Handler:      jsii.String("bootstrap"),
		Code:         awslambda.Code_FromAsset(jsii.String("./bin/bookingCreated.zip"), &awss3assets.AssetOptions{}),
	})

	topic := awssns.NewTopic(stack, topicID, &awssns.TopicProps{
		TopicName: topicName,
	})

	topic.AddSubscription(awssnssubscriptions.NewLambdaSubscription(function, &awssnssubscriptions.LambdaSubscriptionProps{}))

	return function
}
