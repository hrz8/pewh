package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"pewh/awssdk"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var awsCli *awssdk.AWSSDKClient

func handler(_ context.Context, request events.APIGatewayWebsocketProxyRequest) (events.APIGatewayProxyResponse, error) {
	awsCli = awssdk.New()

	appId := request.QueryStringParameters["appid"]
	fmt.Printf("App ID: %s\n", appId)

	connectionID := request.RequestContext.ConnectionID

	switch request.RequestContext.RouteKey {
	case "$connect":
		log.Printf("Connection ID: %s connected", connectionID)
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
		}, nil

	case "$disconnect":
		log.Printf("Connection ID: %s disconnected", connectionID)
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
		}, nil

	case "publish":
		log.Printf("Received message: %s", request.Body)
		go func() {
			time.Sleep(5 * time.Second)
			log.Printf("Replying message for: %s", connectionID)
			awsCli.PostToConnection(connectionID, struct {
				Reply string `json:"reply"`
			}{
				Reply: "pong!",
			})
		}()
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
		}, nil

	default:
		log.Printf("Unknown route: %s", request.RequestContext.RouteKey)
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body:       "Unknown route",
		}, nil
	}
}

func main() {
	lambda.Start(handler)
}
