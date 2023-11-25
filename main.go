package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "200",
			"message": "success",
			"data":    "Welcome to Go Serverless with AWS Lambda & API Gateway!",
		})
	})
	r.Run()
}
