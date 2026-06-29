package main

import (
	"net/http"

	"github.com/MelinaClouet/go-native-docker-api/models"
	"github.com/MelinaClouet/go-native-docker-api/services"
	"github.com/MelinaClouet/go-native-docker-api/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	utils.InitLogger()
	utils.Logger.Println("[SERVER] starting on :8080")

	r := gin.Default()

	// Middleware HTTP logs (TRÈS IMPORTANT)
	r.Use(func(c *gin.Context) {
		utils.Logger.Printf("[HTTP] %s %s", c.Request.Method, c.Request.URL.Path)

		c.Next()

		utils.Logger.Printf("[HTTP] %s %s -> %d",
			c.Request.Method,
			c.Request.URL.Path,
			c.Writer.Status(),
		)
	})

	// --- DOCKER INFO ---
	r.GET("/docker-info", func(c *gin.Context) {
		utils.Logger.Println("[API] docker-info requested")

		info, err := services.GetDockerInfo()
		if err != nil {
			utils.Logger.Printf("[ERROR][API][docker-info] %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, info)
	})

	// --- PROJECTS LIST ---
	r.GET("/projects", func(c *gin.Context) {
		utils.Logger.Println("[API] get projects")

		projects, err := services.GetProjects()
		if err != nil {
			utils.Logger.Printf("[ERROR][API][projects] %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, projects)
	})

	// --- CREATE PROJECT ---
	r.POST("/projects", func(c *gin.Context) {
		var project models.Project

		utils.Logger.Println("[API] create project")

		if err := c.ShouldBindJSON(&project); err != nil {
			utils.Logger.Printf("[ERROR][API][create-project] invalid body: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		utils.Logger.Printf("[API] saving project name=%s", project.Name)

		if err := services.SaveProject(project); err != nil {
			utils.Logger.Printf("[ERROR][API][save-project] %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		utils.Logger.Printf("[API] project saved name=%s", project.Name)

		c.JSON(http.StatusOK, gin.H{"message": "Projet sauvegardé avec succès"})
	})

	// --- DEPLOY ---
	r.POST("/deploy", func(c *gin.Context) {
		var project models.Project

		utils.Logger.Println("[API] deploy request")

		if err := c.ShouldBindJSON(&project); err != nil {
			utils.Logger.Printf("[ERROR][API][deploy] invalid body: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		utils.Logger.Printf("[API] deploying project=%s engine=%s", project.Name, project.Engine.Host)

		err := services.DeployProject(project)
		if err != nil {
			utils.Logger.Printf("[ERROR][API][deploy] project=%s err=%v", project.Name, err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		utils.Logger.Printf("[API] deploy success project=%s", project.Name)

		c.JSON(http.StatusOK, gin.H{"message": "project deployed"})
	})

	// --- STATUS ---
	r.GET("/status/:name", func(c *gin.Context) {
		name := c.Param("name")

		utils.Logger.Printf("[API] status request container=%s", name)

		status, err := services.GetDeploymentStatus(name)
		if err != nil {
			utils.Logger.Printf("[ERROR][API][status] %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		utils.Logger.Printf("[API] status result name=%s status=%s", name, status)

		c.JSON(http.StatusOK, gin.H{"status": status})
	})

	utils.Logger.Println("[SERVER] ready")

	r.Run(":8080")
}
