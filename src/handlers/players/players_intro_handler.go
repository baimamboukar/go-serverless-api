package players

import (
	"net/http"
	"strconv"

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

// deletePlayer deletes a player from the database
// Handles enpoint api/v1/players/delete/:id
func DeletePlayerHandler(c *gin.Context) {
	// Get the id of the player to delete
	id := c.Param("id")
	// Convert the id to an integer
	intID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid player id provided"})
		return
	}

	failed := deletePlayer(intID)
	if failed != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete player"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Player deleted successfully!"})
}
func deletePlayer(id int) error {
	db := database.GetDatabaseInstance()
	result := db.Delete(&models.KenganPlayer{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// GetPlayerHandler gets a player from the database
// Handles enpoint api/v1/players/:id

func GetPlayerHandler(c *gin.Context) {
	// Get the id of the player to delete
	id := c.Param("id")
	// Convert the id to an integer
	intID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid player id provided"})
		return
	}

	player, err := getPlayer(intID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get player"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Player retrieved successfully!", "data": player})
}

func getPlayer(id int) (*models.KenganPlayer, error) {
	db := database.GetDatabaseInstance()
	var player models.KenganPlayer
	result := db.First(&player, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &player, nil
}

func UpdatePlayerHandler(c *gin.Context) {
	// Get the id of the player to delete
	id := c.Param("id")
	// Convert the id to an integer
	intID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid player id provided"})
		return
	}

	// Create a new Character instance to unmarshal the request body into
	var player models.KenganPlayer

	// Bind the request body to the newCharacter struct
	if err := c.BindJSON(&player); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload! All fields are required"})
		return
	}

	failed := updatePlayer(intID, &player)
	if failed != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update player"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Player updated successfully!"})
}

func updatePlayer(id int, player *models.KenganPlayer) error {
	db := database.GetDatabaseInstance()
	result := db.Model(&models.KenganPlayer{}).Where("id = ?", id).Updates(&player)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetAllPlayersHandler(c *gin.Context) {
	players, err := getAllPlayers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get players"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Players retrieved successfully!", "data": players})

}

func getAllPlayers() ([]models.KenganPlayer, error) {
	db := database.GetDatabaseInstance()
	var players []models.KenganPlayer
	result := db.Find(&players)
	if result.Error != nil {
		return nil, nil
	}
	return players, nil
}
