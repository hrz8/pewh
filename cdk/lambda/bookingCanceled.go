package lambda

import (
	"strings"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambdaeventsources"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3assets"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
	"github.com/aws/jsii-runtime-go"
)

func NewBookingCanceledLambda(stack awscdk.Stack, stage string) awslambda.Function {
	stageLower := strings.ToLower(stage)

	functionID := jsii.String("BookingCanceled")
	queueID := jsii.String("BookingCanceledQueue")
	queueName := jsii.String("booking-canceled-" + stageLower)

	function := awslambda.NewFunction(stack, functionID, &awslambda.FunctionProps{
		Runtime:                      awslambda.Runtime_PROVIDED_AL2(),
		Architecture:                 awslambda.Architecture_ARM_64(),
		MemorySize:                   jsii.Number(128),
		Handler:                      jsii.String("bootstrap"),
		Code:                         awslambda.Code_FromAsset(jsii.String("./bin/bookingCanceled.zip"), &awss3assets.AssetOptions{}),
		ReservedConcurrentExecutions: jsii.Number(1),
	})

	queue := awssqs.NewQueue(stack, queueID, &awssqs.QueueProps{
		QueueName:         queueName,
		RetentionPeriod:   awscdk.Duration_Seconds(jsii.Number(86400)), // 24 hours
		VisibilityTimeout: awscdk.Duration_Seconds(jsii.Number(120)),   // 2 minutes
	})

	function.AddEventSource(awslambdaeventsources.NewSqsEventSource(queue, &awslambdaeventsources.SqsEventSourceProps{
		BatchSize: jsii.Number(10),
	}))

	queue.GrantConsumeMessages(function)

	return function
}
