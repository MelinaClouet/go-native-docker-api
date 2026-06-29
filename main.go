package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	// Route de santé
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "API is running",
		})
	})

	// Lancement du serveur sur le port 8080
	r.Run(":8080")
}
