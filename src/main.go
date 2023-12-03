package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"

	EpisodesRequestsHandler "github.com/baimamboukar/go-serverless-api/src/handlers/episodes"
	PlayerRequestsHandler "github.com/baimamboukar/go-serverless-api/src/handlers/players"
)

var ginLambda *ginadapter.GinLambda

func init() {
	//Set the router as the default one provided by Gin
	router := gin.Default()
	// Setup route group for the API
	api := router.Group("/api/v1")
	{
		players := api.Group("/players")
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

// AWS Lambda handler
func GinRequestHandler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return ginLambda.ProxyWithContext(ctx, request)
}

func main() {
	//init()
	lambda.Start(GinRequestHandler)

}
