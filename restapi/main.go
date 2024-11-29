package main

import (
	"context"
	"net/http"

	adapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var ginLambda *adapter.GinLambda

func init() {
	router := gin.Default()
	router.RedirectFixedPath = false

	api := router.Group("/api/v1")
	{
		rooms := api.Group("/rooms")
		rooms.GET("", func(c *gin.Context) {
			c.JSON(http.StatusOK, []gin.H{{"name": "Room 1"}, {"name": "Room 2"}})
		})
		rooms.GET("/another", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"foo": "bar"})
		})
	}

	ginLambda = adapter.New(router)
}

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return ginLambda.ProxyWithContext(ctx, request)
}

func main() {
	lambda.Start(handler)
}
