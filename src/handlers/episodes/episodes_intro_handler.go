package episodes

import (
	"net/http"

	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
)

var ginLambda *ginadapter.GinLambda

func EpisodesIntroHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Welcome to the episodes endpoint",
		"data": gin.H{
			"endpoints": gin.H{
				"players": []gin.H{
					{
						"method": "GET",
						"path":   "/api/v1/episodes",
						"desc":   "Get all episodes",
					},
					{
						"method": "GET",
						"path":   "/api/v1/episodes/:id",
						"desc":   "Get episodes Informations",
					},
					{
						"method": "POST",
						"path":   "/api/v1/episodes/create",
						"desc":   "Create an episode",
					},
					{
						"method": "PUT",
						"path":   "/api/v1/episodes/update/:id",
						"desc":   "Updates a episodes",
					},
					{
						"method": "DELETE",
						"path":   "/api/v1/episodes/update/:id",
						"desc":   "Deletes an episode",
					},
				},
			},
		},
	})
}
