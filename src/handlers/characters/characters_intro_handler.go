package characters

import (
	"net/http"

	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
)

var ginLambda *ginadapter.GinLambda

// CharactersIntroHandler handles the /api/v1/characters endpoint
func charactersIntroHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Welcome to the Charcater endpoint",
		"data": gin.H{
			"endpoints": gin.H{
				"players": []gin.H{
					{
						"method": "GET",
						"path":   "/api/v1/chatacters",
						"desc":   "Get all chatacters",
					},
					{
						"method": "GET",
						"path":   "/api/v1/chatacter/:id",
						"desc":   "Get chatacter Informations",
					},
					{
						"method": "POST",
						"path":   "/api/v1/chatacters/create",
						"desc":   "Create a chatacter",
					},
					{
						"method": "PUT",
						"path":   "/api/v1/chatacters/update/:id",
						"desc":   "Create a chatacter",
					},
					{
						"method": "DELETE",
						"path":   "/api/v1/chatacters/update/:id",
						"desc":   "Deletes a chatacter",
					},
				},
			},
		},
	})
}
