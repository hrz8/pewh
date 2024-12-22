package main

import (
	"pewh/cdk/lambda"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type ServiceLinkProps struct {
	awscdk.StackProps
	Stage string
}

func NewLinkService(scope constructs.Construct, props *ServiceLinkProps) awscdk.Stack {
	stackID := jsii.String("LinkServiceStack")

	stack := awscdk.NewStack(scope, stackID, &props.StackProps)

	// register reources of this service
	lambda.NewRestapiLambda(stack)
	lambda.NewWebSocketLambda(stack)
	lambda.NewBookingCreatedLambda(stack, props.Stage)
	lambda.NewBookingPaidLambda(stack, props.Stage)
	lambda.NewBookingCanceledLambda(stack, props.Stage)
	lambda.NewCheckInStartedLambda(stack, props.Stage)

	return stack
}
