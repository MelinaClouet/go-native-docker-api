package services

import (
	"encoding/json"
	"os"

	"github.com/MelinaClouet/go-native-docker-api/models"
	"github.com/MelinaClouet/go-native-docker-api/utils"
)

const projectFile = "storage/projects.json"

// SaveProject sauvegarde un projet dans le fichier JSON local.
//
// Cette fonction :
//  1. charge les projets existants
//  2. ajoute le nouveau projet
//  3. sauvegarde l’ensemble dans storage/projects.json
//
// Elle agit comme une "mini base de données locale".
func SaveProject(project models.Project) error {

	utils.Logger.Printf("[PROJECT][SAVE] loading existing projects file=%s", projectFile)

	// -------------------------
	// LOAD EXISTING DATA
	// -------------------------
	projects, err := loadProjects()
	if err != nil {
		utils.Logger.Printf("[ERROR][PROJECT][LOAD] failed err=%v", err)
		return err
	}

	utils.Logger.Printf("[PROJECT][SAVE] existing projects=%d", len(projects))

	// -------------------------
	// APPEND NEW PROJECT
	// -------------------------
	projects = append(projects, project)

	utils.Logger.Printf("[PROJECT][SAVE] adding project name=%s total=%d",
		project.Name, len(projects),
	)

	// -------------------------
	// SERIALIZE DATA
	// -------------------------
	data, err := json.MarshalIndent(projects, "", "  ")
	if err != nil {
		utils.Logger.Printf("[ERROR][PROJECT][MARSHAL] project=%s err=%v",
			project.Name, err,
		)
		return err
	}

	// -------------------------
	// WRITE FILE
	// -------------------------
	err = os.WriteFile(projectFile, data, 0644)
	if err != nil {
		utils.Logger.Printf("[ERROR][PROJECT][WRITE] file=%s err=%v",
			projectFile, err,
		)
		return err
	}

	utils.Logger.Printf("[PROJECT][SAVE] success project=%s file=%s",
		project.Name, projectFile,
	)

	return nil
}

// loadProjects charge tous les projets depuis le fichier JSON.
//
// Si le fichier n’existe pas, retourne une liste vide (cas normal au démarrage).
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

// GetProjects retourne tous les projets enregistrés.
//
// Cette fonction est utilisée par l’API REST (GET /projects).
func GetProjects() ([]models.Project, error) {

	utils.Logger.Println("[PROJECT][GET] fetching all projects")

	return loadProjects()
}
