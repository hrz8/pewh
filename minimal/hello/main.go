package main

import (
	"context"
	"net/http"

	adapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var ginLambda *adapter.GinLambdaV2

func init() {
	router := gin.Default()
	router.RedirectFixedPath = false

	api := router.Group("/hello")
	api.GET("", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"foo": "bar"})
	})

	ginLambda = adapter.NewV2(router)
}

func handler(ctx context.Context, request events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	return ginLambda.ProxyWithContext(ctx, request)
}

func main() {
	lambda.Start(handler)
}
