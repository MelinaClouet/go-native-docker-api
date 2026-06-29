package services

import (
	"encoding/json"
	"os"

	"github.com/MelinaClouet/go-native-docker-api/models"
)

const projectFile = "storage/projects.json"

// SaveProject sauvegarde un projet dans un fichier JSON
func SaveProject(project models.Project) error {
	// Lire les données existantes
	projects, err := loadProjects()
	if err != nil {
		return err
	}

	// Ajouter le nouveau projet
	projects = append(projects, project)

	// Sauvegarder dans le fichier
	data, err := json.MarshalIndent(projects, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(projectFile, data, 0644)
}

// loadProjects charge les projets d'un fichier JSON
func loadProjects() ([]models.Project, error) {
	data, err := os.ReadFile(projectFile)
	if os.IsNotExist(err) {
		return []models.Project{}, nil
	}
	if err != nil {
		return nil, err
	}

	var projects []models.Project
	err = json.Unmarshal(data, &projects)
	if err != nil {
		return nil, err
	}

	return projects, nil
}
