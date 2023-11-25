package main

import (
	"context"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
)

var ginLambda *ginadapter.GinLambda

func init() {
	//Set the router as the default one provided by Gin
	router := gin.Default()
	// Setup route group for the API
	api := router.Group("/api")
	{
		api.GET("/", RootHandler)
		api.GET("/notes", NotesHandler)
	}
	// Start and run the server
	ginLambda = ginadapter.New(router)
}

// AWS Lambda handler
func PingGetHandler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return ginLambda.ProxyWithContext(ctx, request)
}

// Handler for API group ping
func RootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "The beauty of Serverless is that you don't have to worry about the server.",
	})
}

// Handler for API group ping
func NotesHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"status":  "success",
		"message": "Sense notes fetched!",
		"notes":   []string{"Lorem ipsum dolor sit amet", "The beautiful thing is that no one can take it away from you", "The beautiful thing is that no one can take it away from you"},
	})
}

func main() {
	lambda.Start(PingGetHandler)
}
