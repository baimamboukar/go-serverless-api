package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	EpisodesRequestsHandler "github.com/baimamboukar/go-serverless-api/src/handlers/episodes"
	PlayerRequestsHandler "github.com/baimamboukar/go-serverless-api/src/handlers/players"
)

// This function loads the `.env` file.
//
// # The file should contain your AWS credentials
//
// AWS_REGION - AWS_ACCESS_KEY - AWS_ACCESS_KEY_ID
func LoadDotEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file:", err)
	}
}

// AWS gin lambda adapter
var ginLambda *ginadapter.GinLambda

func init() {

	// Load .env file
	LoadDotEnv()

	//Set the router as the default one provided by Gin
	router := gin.Default()

	// Setup route group for KenganAshura Players
	// Here we create a group `/api/v1`
	// All the requests will be mapped to this group
	// localhost:8080/api/v1/getAll

	api := router.Group("/api/v1")
	{
		players := api.Group("/players")

		// Mapping Player routes to their handlers
		players.GET("/", PlayerRequestsHandler.PlayersIntroHandler)
		players.GET("/get/:id", PlayerRequestsHandler.GetPlayerHandler)
		players.GET("/getAll", PlayerRequestsHandler.GetAllPlayersHandler)
		players.POST("/create", PlayerRequestsHandler.CreatePlayerHandler)
		players.PATCH("/update/:id", PlayerRequestsHandler.UpdatePlayerHandler)
		players.DELETE("/delete/:id", PlayerRequestsHandler.DeletePlayerHandler)

		// Setup route group for characters
		//	characters := api.Group("/characters")
		//characters.GET("/", CharacterRequestsHandler.)
		// characters.POST("/", CreateCharacterHandler)
		// characters.PUT("/:id", UpdateCharacterHandler)
		// characters.DELETE("/:id", DeleteCharacterHandler)

		// Setup routes for episodes
		episodes := api.Group("/episodes")
		episodes.GET("/", EpisodesRequestsHandler.EpisodesIntroHandler)
		// episodes.GET("/:id", EpisodeHandler)
		// episodes.POST("/", CreateEpisodeHandler)
		// episodes.PUT("/:id", UpdateEpisodeHandler)
		// episodes.DELETE("/:id", DeleteEpisodeHandler)

	}
	// Start and run the server
	ginLambda = ginadapter.New(router)
	//router.Run(":8000")
}

// AWS Lambda Proxy Handler
// This handler acts like a bridge between AWS Lambda and our Local GIn server
// It maps each GIN route to a Lambda function as handler
//
// This is useful to make our function execution possible.
func GinRequestHandler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return ginLambda.ProxyWithContext(ctx, request)
}

func main() {
	// Starts Lambda server
	lambda.Start(GinRequestHandler)

}
