package services

import (
	"context"

	"github.com/MelinaClouet/go-native-docker-api/utils"
	"github.com/moby/moby/client"
)

// DockerInfo regroupe les informations du Docker Engine.
//
// Elle contient la liste des conteneurs et des images
// récupérées depuis le daemon Docker (local ou distant).
type DockerInfo struct {
	Containers client.ContainerListResult
	Images     client.ImageListResult
}

// GetDockerInfo retourne un snapshot de l'état actuel du Docker Engine.
//
// Elle récupère :
//   - la liste des conteneurs
//   - la liste des images
//
// Cette fonction est utilisée par la route /docker-info.
func GetDockerInfo() (*DockerInfo, error) {

	utils.Logger.Println("[DOCKER_INFO] fetching docker info")

	// Initialisation du client Docker depuis l'environnement système
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		utils.Logger.Printf("[ERROR][DOCKER_INFO] docker client init failed: %v", err)
		return nil, err
	}

	// -------------------------
	// LIST CONTAINERS
	// -------------------------
	utils.Logger.Println("[DOCKER_INFO] listing containers")

	containers, err := cli.ContainerList(context.Background(), client.ContainerListOptions{})
	if err != nil {
		utils.Logger.Printf("[ERROR][DOCKER_INFO] container list failed: %v", err)
		return nil, err
	}

	utils.Logger.Printf("[DOCKER_INFO] containers found=%d", len(containers.Items))

	// -------------------------
	// LIST IMAGES
	// -------------------------
	utils.Logger.Println("[DOCKER_INFO] listing images")

	images, err := cli.ImageList(context.Background(), client.ImageListOptions{})
	if err != nil {
		utils.Logger.Printf("[ERROR][DOCKER_INFO] image list failed: %v", err)
		return nil, err
	}

	utils.Logger.Printf("[DOCKER_INFO] images found=%d", len(images.Items))

	return &DockerInfo{
		Containers: containers,
		Images:     images,
	}, nil
}

// GetDeploymentStatus retourne l'état d'un conteneur Docker.
//
// Elle recherche un conteneur par son nom et retourne son statut :
//   - running
//   - exited
//   - created
//   - not-found
func GetDeploymentStatus(name string) (string, error) {

	utils.Logger.Printf("[STATUS] checking container name=%s", name)

	// Client Docker (local ou remote selon env)
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		utils.Logger.Printf("[ERROR][STATUS] docker client init failed: %v", err)
		return "", err
	}

	utils.Logger.Println("[STATUS] listing containers (All=true)")

	// Récupération de tous les conteneurs (actifs + arrêtés)
	result, err := cli.ContainerList(context.Background(), client.ContainerListOptions{
		All: true,
	})
	if err != nil {
		utils.Logger.Printf("[ERROR][STATUS] container list failed: %v", err)
		return "", err
	}

	utils.Logger.Printf("[STATUS] total containers=%d", len(result.Items))

	// -------------------------
	// SEARCH CONTAINER
	// -------------------------
	for _, ctr := range result.Items {
		for _, n := range ctr.Names {
			if n == "/"+name {
				utils.Logger.Printf("[STATUS] found container name=%s status=%s", name, ctr.Status)
				return ctr.Status, nil
			}
		}
	}

	utils.Logger.Printf("[STATUS] container not found name=%s", name)

	return "not-found", nil
}
