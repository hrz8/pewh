package main

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
)

type StageProps struct {
	awscdk.StageProps
}

func NewStage(scope constructs.Construct, id string, props *StageProps) {
	stage := awscdk.NewStage(scope, &id, &props.StageProps)

	// register link-service
	NewLinkService(stage, &ServiceLinkProps{
		StackProps: awscdk.StackProps{
			Env: props.Env,
		},
		Stage: id,
	})
}
