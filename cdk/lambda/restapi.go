package lambda

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigatewayv2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigatewayv2integrations"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3assets"
	"github.com/aws/jsii-runtime-go"
)

func NewRestapiLambda(stack awscdk.Stack) awslambda.Function {
	functionID := jsii.String("Restapi")
	apigwID := jsii.String("RestapiApi")
	apigwIntegrationID := jsii.String("RestapiApiIntegration")

	function := awslambda.NewFunction(stack, functionID, &awslambda.FunctionProps{
		Runtime:      awslambda.Runtime_PROVIDED_AL2(),
		Architecture: awslambda.Architecture_ARM_64(),
		MemorySize:   jsii.Number(128),
		Handler:      jsii.String("bootstrap"),
		Code:         awslambda.Code_FromAsset(jsii.String("./bin/restapi.zip"), &awss3assets.AssetOptions{}),
	})

	api := awsapigatewayv2.NewHttpApi(stack, apigwID, &awsapigatewayv2.HttpApiProps{
		CorsPreflight: &awsapigatewayv2.CorsPreflightOptions{
			AllowOrigins: &[]*string{jsii.String("*")},
		},
	})

	integ := awsapigatewayv2integrations.NewHttpLambdaIntegration(
		apigwIntegrationID, function, &awsapigatewayv2integrations.HttpLambdaIntegrationProps{},
	)

	api.AddRoutes(&awsapigatewayv2.AddRoutesOptions{
		Integration: integ,
		Methods:     &[]awsapigatewayv2.HttpMethod{awsapigatewayv2.HttpMethod_GET},
		Path:        jsii.String("/api/v1/rooms"),
	})
	api.AddRoutes(&awsapigatewayv2.AddRoutesOptions{
		Integration: integ,
		Methods:     &[]awsapigatewayv2.HttpMethod{awsapigatewayv2.HttpMethod_POST},
		Path:        jsii.String("/api/v1/rooms"),
	})
	api.AddRoutes(&awsapigatewayv2.AddRoutesOptions{
		Integration: integ,
		Methods:     &[]awsapigatewayv2.HttpMethod{awsapigatewayv2.HttpMethod_GET},
		Path:        jsii.String("/api/v1/rooms/another"),
	})
	api.AddRoutes(&awsapigatewayv2.AddRoutesOptions{
		Integration: integ,
		Methods:     &[]awsapigatewayv2.HttpMethod{awsapigatewayv2.HttpMethod_GET},
		Path:        jsii.String("/api/v1/rooms/log"),
	})

	return function
}
