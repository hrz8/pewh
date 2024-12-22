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

func NewCheckInStartedLambda(stack awscdk.Stack, stage string) awslambda.Function {
	stageLower := strings.ToLower(stage)

	functionID := jsii.String("CheckInStarted")
	queueID := jsii.String("CheckInStartedQueue")
	queueName := jsii.String("check-in-started-" + stageLower + ".fifo")
	dlqID := jsii.String("CheckInStartedDlq")
	dlqName := jsii.String("check-in-started-dlq-" + stageLower + ".fifo")

	function := awslambda.NewFunction(stack, functionID, &awslambda.FunctionProps{
		Runtime:                      awslambda.Runtime_PROVIDED_AL2(),
		Architecture:                 awslambda.Architecture_ARM_64(),
		MemorySize:                   jsii.Number(128),
		Handler:                      jsii.String("bootstrap"),
		Code:                         awslambda.Code_FromAsset(jsii.String("./bin/checkInStarted.zip"), &awss3assets.AssetOptions{}),
		ReservedConcurrentExecutions: jsii.Number(30),
	})

	dlq := awssqs.NewQueue(stack, dlqID, &awssqs.QueueProps{
		QueueName:                 dlqName,
		Fifo:                      jsii.Bool(true),
		ContentBasedDeduplication: jsii.Bool(true),
		RetentionPeriod:           awscdk.Duration_Seconds(jsii.Number(1209600)), // 14 days
	})

	queue := awssqs.NewQueue(stack, queueID, &awssqs.QueueProps{
		QueueName:                 queueName,
		Fifo:                      jsii.Bool(true),
		ContentBasedDeduplication: jsii.Bool(true),
		RetentionPeriod:           awscdk.Duration_Seconds(jsii.Number(21600)), // 6 hours
		VisibilityTimeout:         awscdk.Duration_Seconds(jsii.Number(15)),    // 15 seconds
		DeadLetterQueue: &awssqs.DeadLetterQueue{
			MaxReceiveCount: jsii.Number(5),
			Queue:           dlq,
		},
	})

	function.AddEventSource(awslambdaeventsources.NewSqsEventSource(queue, &awslambdaeventsources.SqsEventSourceProps{
		BatchSize: jsii.Number(10),
	}))

	queue.GrantConsumeMessages(function)

	return function
}
