package services

import (
	"encoding/json"
	"os"

	"github.com/MelinaClouet/go-native-docker-api/models"
	"github.com/MelinaClouet/go-native-docker-api/utils"
)

const projectFile = "storage/projects.json"

// SaveProject sauvegarde un projet dans un fichier JSON
func SaveProject(project models.Project) error {
	utils.Logger.Printf("[PROJECT][SAVE] loading existing projects file=%s", projectFile)

	// Lire les données existantes
	projects, err := loadProjects()
	if err != nil {
		utils.Logger.Printf("[ERROR][PROJECT][LOAD] failed err=%v", err)
		return err
	}

	utils.Logger.Printf("[PROJECT][SAVE] existing projects=%d", len(projects))

	// Ajouter le nouveau projet
	projects = append(projects, project)

	utils.Logger.Printf("[PROJECT][SAVE] adding project name=%s total=%d", project.Name, len(projects))

	// Sauvegarder dans le fichier
	data, err := json.MarshalIndent(projects, "", "  ")
	if err != nil {
		utils.Logger.Printf("[ERROR][PROJECT][MARSHAL] project=%s err=%v", project.Name, err)
		return err
	}

	err = os.WriteFile(projectFile, data, 0644)
	if err != nil {
		utils.Logger.Printf("[ERROR][PROJECT][WRITE] file=%s err=%v", projectFile, err)
		return err
	}

	utils.Logger.Printf("[PROJECT][SAVE] success project=%s file=%s", project.Name, projectFile)

	return nil
}

// loadProjects charge les projets d'un fichier JSON
func loadProjects() ([]models.Project, error) {
	utils.Logger.Printf("[PROJECT][LOAD] reading file=%s", projectFile)

	data, err := os.ReadFile(projectFile)
	if os.IsNotExist(err) {
		utils.Logger.Printf("[PROJECT][LOAD] file not found -> returning empty list")
		return []models.Project{}, nil
	}
	if err != nil {
		utils.Logger.Printf("[ERROR][PROJECT][LOAD] read file failed err=%v", err)
		return nil, err
	}

	var projects []models.Project
	err = json.Unmarshal(data, &projects)
	if err != nil {
		utils.Logger.Printf("[ERROR][PROJECT][UNMARSHAL] err=%v", err)
		return nil, err
	}

	utils.Logger.Printf("[PROJECT][LOAD] success projects=%d", len(projects))

	return projects, nil
}

func GetProjects() ([]models.Project, error) {
	utils.Logger.Println("[PROJECT][GET] fetching all projects")
	return loadProjects()
}
