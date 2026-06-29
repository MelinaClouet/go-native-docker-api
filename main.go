package main

import (
	"net/http"

	"github.com/MelinaClouet/go-native-docker-api/services"
	"github.com/gin-gonic/gin"
)

func main() {
	println("Serveur démarré")
	r := gin.Default()

	// Route pour récupérer les informations Docker
	r.GET("/docker-info", func(c *gin.Context) {
		info, err := services.GetDockerInfo()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, info)
	})

	r.Run(":8080") // Démarre le serveur sur le port 8080
}
