package players

import (
	"net/http"

	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"

	database "github.com/baimamboukar/go-serverless-api/src/database"
	"github.com/baimamboukar/go-serverless-api/src/models"
)

var ginLambda *ginadapter.GinLambda

func PlayersIntroHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Welcome to the players endpoint",
		"data": gin.H{
			"endpoints": gin.H{
				"players": []gin.H{
					{
						"method": "GET",
						"path":   "/api/v1/players",
						"desc":   "Get all players",
					},
					{
						"method": "GET",
						"path":   "/api/v1/players/:id",
						"desc":   "Get player Informations",
					},
					{
						"method": "POST",
						"path":   "/api/v1/players/create",
						"desc":   "Create a player",
					},
					{
						"method": "PUT",
						"path":   "/api/v1/players/update/:d",
						"desc":   "Create a player",
					},
				},
			},
		},
	})
}

// CharactersHandler handles the /api/v1/characters/create endpoint
// It creates a new character from the request body and saves it to the Databse via GORM
func CreatePlayerHandler(c *gin.Context) {
	// get the body of our POST request
	// unmarshal this into a new Character struct

	// Create a new Character instance to unmarshal the request body into
	var player models.KenganPlayer

	// Bind the request body to the newCharacter struct
	if err := c.BindJSON(&player); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload! All fields are required"})
		return
	}

	// Save the character to the database via GORM
	err := savePlayer(&player)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save character to the database"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"status": "success", "message": "Character created successfully!", "data": player})

}

func savePlayer(player *models.KenganPlayer) error {
	db := database.GetDatabaseInstance()
	result := db.Create(&player)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
