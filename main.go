package main

import (
	"net/http"

	"github.com/MelinaClouet/go-native-docker-api/models"
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

	r.GET("/projects", func(c *gin.Context) {
		projects, err := services.GetProjects()
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, projects)
	})

	r.POST("/projects", func(c *gin.Context) {
		var project models.Project
		if err := c.ShouldBindJSON(&project); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Sauvegarde le projet via le service
		if err := services.SaveProject(project); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Projet sauvegardé avec succès"})
	})

	r.Run(":8080") // Démarre le serveur sur le port 8080
}
