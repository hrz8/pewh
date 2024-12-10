package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"pewh/awssdk"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var awsCli *awssdk.AWSSDKClient

type Body struct {
	Action string `json:"action"`
	Data   struct {
		Message string `json:"message"`
	} `json:"data"`
}

func handler(_ context.Context, request events.APIGatewayWebsocketProxyRequest) (events.APIGatewayProxyResponse, error) {
	awsCli = awssdk.New()

	appId := request.QueryStringParameters["appid"]
	fmt.Printf("App ID: %s\n", appId)

	connectionID := request.RequestContext.ConnectionID

	switch request.RequestContext.RouteKey {
	case "$connect":
		log.Printf("Connection ID: %s connected\n", connectionID)
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
		}, nil

	case "$disconnect":
		log.Printf("Connection ID: %s disconnected\n", connectionID)
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
		}, nil

	case "ping":
		awsCli.PostToConnection(connectionID, struct {
			Reply string `json:"reply"`
		}{
			Reply: "pong!",
		})
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Body:       "success",
		}, nil

	case "publish":
		log.Printf("Received body raw: %s\n", request.Body)
		var body Body
		if err := json.Unmarshal([]byte(request.Body), &body); err != nil {
			log.Printf("Error parsing body: %v\n", err)
		}
		fmt.Printf("Received parsed body: %v\n", body)
		go func() {
			if body.Data.Message == "graceful" {
				return
			}
			time.Sleep(5 * time.Second)
			log.Printf("Replying message for: %s\n", connectionID)
			awsCli.PostToConnection(connectionID, struct {
				Reply string `json:"reply"`
			}{
				Reply: "this is the reply!",
			})
		}()
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Body:       "success",
		}, nil

	default:
		log.Printf("Unknown route: %s\n", request.RequestContext.RouteKey)
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body:       "Unknown route",
		}, nil
	}
}

func main() {
	lambda.Start(handler)
}
