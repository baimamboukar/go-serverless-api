package main

import (
	"context"
	"net/http"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/gin-gonic/gin"
)

func init() {
	//Set the router as the default one provided by Gin
	router := gin.Default()
	// Setup route group for the API
	api := router.Group("/api")
	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
	}
	// Start and run the server
	router.Run(":3000")
}

func main() {
	// engine := gin.Default()
	// engine.GET("/ping", PingGetHandler)
	//engine.POST("/ping", PingPostHandler)
	// engine.Run()
	lambda.Start(PingPostHandler)
}

func PingPostHandler(ctx context.Context, name events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	time.Sleep(5 * time.Second)
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       "Welcome to Go Serverless with AWS Lambda & API Gateway!",
	}, nil
}
func PingGetHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"status":  200,
		"message": "success",
		"data":    "Welcome to Go Serverless with AWS Lambda & API Gateway!",
	})
}
