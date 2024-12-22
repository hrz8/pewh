package lambda

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigatewayv2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigatewayv2integrations"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3assets"
	"github.com/aws/jsii-runtime-go"
)

func NewWebSocketLambda(stack awscdk.Stack) awslambda.Function {
	functionID := jsii.String("Websocket")
	wsApigwID := jsii.String("WebsocketApi")
	wsApigwIntegrationID := jsii.String("WebsocketApiIntegration")

	function := awslambda.NewFunction(stack, functionID, &awslambda.FunctionProps{
		Runtime:      awslambda.Runtime_PROVIDED_AL2(),
		Architecture: awslambda.Architecture_ARM_64(),
		MemorySize:   jsii.Number(128),
		Handler:      jsii.String("bootstrap"),
		Code:         awslambda.Code_FromAsset(jsii.String("./bin/websocket.zip"), &awss3assets.AssetOptions{}),
	})

	wsApi := awsapigatewayv2.NewWebSocketApi(stack, wsApigwID, &awsapigatewayv2.WebSocketApiProps{
		RouteSelectionExpression: jsii.String(
			"$request.body.action",
		),
	})

	wsInteg := awsapigatewayv2integrations.NewWebSocketLambdaIntegration(
		wsApigwIntegrationID, function, &awsapigatewayv2integrations.WebSocketLambdaIntegrationProps{},
	)

	wsApi.AddRoute(jsii.String("$connect"), &awsapigatewayv2.WebSocketRouteOptions{
		Integration: wsInteg,
	})
	wsApi.AddRoute(jsii.String("$disconnect"), &awsapigatewayv2.WebSocketRouteOptions{
		Integration: wsInteg,
	})
	wsApi.AddRoute(jsii.String("ping"), &awsapigatewayv2.WebSocketRouteOptions{
		Integration: wsInteg,
	})
	wsApi.AddRoute(jsii.String("publish"), &awsapigatewayv2.WebSocketRouteOptions{
		Integration: wsInteg,
	})

	return function
}
