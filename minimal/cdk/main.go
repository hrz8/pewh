package main

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigatewayv2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigatewayv2integrations"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3assets"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

func main() {
	defer jsii.Close()

	app := awscdk.NewApp(nil)

	NewStage(app, "Local", &awscdk.StageProps{
		Env: &awscdk.Environment{
			Account: jsii.String("000000000000"),
			Region:  jsii.String("us-east-1"),
		},
	})

	app.Synth(nil)
}

func NewStage(scope constructs.Construct, id string, props *awscdk.StageProps) {
	stage := awscdk.NewStage(scope, &id, props)

	stack := awscdk.NewStack(stage, jsii.String("GreeterService"), &awscdk.StackProps{
		Env: props.Env,
	})

	function := awslambda.NewFunction(stack, jsii.String("Hello"), &awslambda.FunctionProps{
		Runtime:      awslambda.Runtime_PROVIDED_AL2(),
		Architecture: awslambda.Architecture_ARM_64(),
		MemorySize:   jsii.Number(128),
		Handler:      jsii.String("bootstrap"),
		Code: awslambda.Code_FromAsset(
			jsii.String("./bin/hello.zip"),
			&awss3assets.AssetOptions{},
		),
	})

	api := awsapigatewayv2.NewHttpApi(stack, jsii.String("HelloApi"), &awsapigatewayv2.HttpApiProps{
		CorsPreflight: &awsapigatewayv2.CorsPreflightOptions{
			AllowOrigins: &[]*string{jsii.String("*")},
		},
	})

	integration := awsapigatewayv2integrations.NewHttpLambdaIntegration(
		jsii.String("HelloApiIntegration"), function, &awsapigatewayv2integrations.HttpLambdaIntegrationProps{},
	)

	api.AddRoutes(&awsapigatewayv2.AddRoutesOptions{
		Integration: integration,
		Methods:     &[]awsapigatewayv2.HttpMethod{awsapigatewayv2.HttpMethod_GET},
		Path:        jsii.String("/hello"),
	})
}
